package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Tile represent single image tile
type Tile [][]byte

// Print prints tile
func (t Tile) Print() {
	for _, row := range t {
		fmt.Println(string(row))
	}
	fmt.Println("")
}

// Edges returns list of tile edges
func (t Tile) Edges() [4]string {
	var colL, colR []byte
	for _, row := range t {
		colL = append(colL, row[0])
		colR = append(colR, row[len(row)-1])
	}
	return [4]string{string(t[0]), string(colR), string(t[len(t)-1]), string(colL)}
}

// Rotate returns copy of tile rotated by 90 degree CW
func (t Tile) Rotate() Tile {
	n := len(t)
	newTile := emptyTile(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newTile[i][j] = t[n-j-1][i]
		}
	}

	return newTile
}

// RotateTimes rotates tile given number of times
func (t Tile) RotateTimes(n int) Tile {
	tile := t
	for i := 0; i < n; i++ {
		tile = tile.Rotate()
	}
	return tile
}

// FlipHorizontal flips tile horizontaly
func (t Tile) FlipHorizontal() Tile {
	n := len(t)
	newTile := emptyTile(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newTile[i][j] = t[i][n-j-1]
		}
	}
	return newTile
}

// FlipVertical flips tile vertically
func (t Tile) FlipVertical() Tile {
	n := len(t)
	newTile := emptyTile(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newTile[i][j] = t[n-i-1][j]
		}
	}
	return newTile
}

func emptyTile(n int) Tile {
	tile := make(Tile, n)
	for i := range tile {
		tile[i] = make([]byte, n)
	}
	return tile
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	tiles, err := readTiles(file)

	fmt.Printf("Tiles: %d\n", len(tiles))

	edgesMap := make(map[string][]int)
	for id, tile := range tiles {
		for _, edge := range tile.Edges() {
			if _, ok := edgesMap[reverse(edge)]; ok {
				edge = reverse(edge)
			}
			edgesMap[edge] = append(edgesMap[edge], id)
		}
	}

	singleEdgeTiles := make(map[int][]string)
	for edge, ids := range edgesMap {
		if len(ids) == 1 {
			for _, id := range ids {
				singleEdgeTiles[id] = append(singleEdgeTiles[id], edge)
			}
		}
	}

	var cornerTiles []int
	for id, edges := range singleEdgeTiles {
		fmt.Printf("%d, ", id)
		if len(edges) == 2 {
			cornerTiles = append(cornerTiles, id)
		}
	}

	fmt.Printf("Corner tiles: %v\n", cornerTiles)
	fmt.Printf("Result: %d\n", cornerTiles[0]*cornerTiles[1]*cornerTiles[2]*cornerTiles[3])

	n := int(math.Sqrt(float64(len(tiles))))
	fmt.Printf("Image size: %dx%d\n", n, n)

	img := make([][]Tile, n)
	for i := range img {
		img[i] = make([]Tile, n)
	}

	cornerID := cornerTiles[0]
	outlineEdges := singleEdgeTiles[cornerID]
	corner := tiles[cornerID]

	delete(tiles, cornerID)
	delete(singleEdgeTiles, cornerID)

	allCornerEdges := corner.Edges()
	idx1 := findEdge(outlineEdges[0], allCornerEdges)
	idx2 := findEdge(outlineEdges[1], allCornerEdges)

	if idx2 < idx1 || (idx2 == 3 && idx1 == 0) {
		idx1, idx2 = idx2, idx1
		outlineEdges[0], outlineEdges[1] = outlineEdges[1], outlineEdges[0]
	}

	corner = corner.RotateTimes(calcRotations(idx1))
	img[0][n-1] = corner
	corner.Print()

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if img[i][j] != nil {
				continue
			}
			neighbourEdges := getNeighbourEdges(img, i, j)

			for id, tile := range tiles {
				tile := matchEdges(tile, neighbourEdges)
				if tile != nil {
					// fmt.Printf("Found matching tile %d at [%d, %d]\n", id, i, j)
					// tile.Print()
					delete(tiles, id)
					img[i][j] = tile
					break
				}
			}

			if img[i][j] == nil {
				fmt.Printf("Tile not found for: %d, %d\n", i, j)
				panic("not found")
			}
		}
	}

	fmt.Println("")

	trimedTileSize := len(corner) - 2
	imgSize := n * trimedTileSize
	finalImg := emptyTile(imgSize)
	for i := 0; i < imgSize; i++ {
		ti := i / trimedTileSize
		for j := 0; j < imgSize; j++ {
			tj := j / trimedTileSize
			tile := img[ti][tj]
			finalImg[i][j] = tile[(i%trimedTileSize)+1][(j%trimedTileSize)+1]
		}
	}

	finalImg.Print()

	monstersMap := make(map[[2]int]bool)

	finalImg = findPossibleMonsters(finalImg, monstersMap)

	finalImg.Print()

	notMonster := 0
	for i := 0; i < imgSize; i++ {
		for j := 0; j < imgSize; j++ {
			_, isMonster := monstersMap[[2]int{i, j}]
			if finalImg[i][j] == '#' && !isMonster {
				notMonster++
			}
		}
	}

	fmt.Printf("Not monster fields: %d\n", notMonster)
}

func findPossibleMonsters(tile Tile, monstersMap map[[2]int]bool) Tile {
	for f := 0; f < 2; f++ {
		for r := 0; r < 4; r++ {
			monsters := findMonsters(tile, monstersMap)
			if len(monsters) > 0 {
				fmt.Printf("Found monster at %v\n", monsters)
				return tile
			}
			tile = tile.Rotate()
		}
		tile = tile.FlipHorizontal()
	}
	return nil
}

// Monster represents pattern for monster
var Monster = [3]string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func findMonsters(tile Tile, monsterMap map[[2]int]bool) [][2]int {
	var monsters [][2]int
	for i := 0; i <= len(tile)-len(Monster); i++ {
		for j := 0; j <= len(tile)-len(Monster[0]); j++ {
			if checkMonster(tile, i, j, monsterMap) {
				monsters = append(monsters, [2]int{i, j})
			}
		}
	}
	return monsters
}

func checkMonster(tile Tile, i, j int, monsterMap map[[2]int]bool) bool {
	for mi := range Monster {
		for mj, char := range Monster[mi] {
			if char == '#' && tile[i+mi][j+mj] != byte(char) {
				return false
			}
		}
	}

	for mi := range Monster {
		for mj, char := range Monster[mi] {
			if char == '#' && tile[i+mi][j+mj] == byte(char) {
				monsterMap[[2]int{i + mi, j + mj}] = true
			}
		}
	}

	return true
}

func getNeighbours(img [][]Tile, i, j int) [4]Tile {
	var neighbours [4]Tile
	if i > 0 {
		neighbours[0] = img[i-1][j]
	}
	if (j + 1) < len(img) {
		neighbours[1] = img[i][j+1]
	}
	if (i + 1) < len(img) {
		neighbours[2] = img[i+1][j]
	}
	if j > 0 {
		neighbours[3] = img[i][j-1]
	}
	return neighbours
}

func getNeighbourEdges(img [][]Tile, i, j int) [4]string {
	var edges [4]string
	for i, n := range getNeighbours(img, i, j) {
		if n != nil {
			edges[i] = n.Edges()[(i+2)%4]
		}
	}
	return edges
}

func matchEdges(tile Tile, edges [4]string) Tile {
	t := tile
	for f := 0; f < 2; f++ {
		for r := 0; r < 4; r++ {
			if matchingEdges(t.Edges(), edges) {
				return t
			}
			t = t.Rotate()
		}
		t = t.FlipHorizontal()
	}

	return nil
}

func matchingEdges(tileEdges, neighbourEdges [4]string) bool {
	for i, edge := range tileEdges {
		if neighbourEdges[i] != "" && neighbourEdges[i] != edge {
			return false
		}
	}
	return true
}

func calcRotations(idx int) int {
	return (4 - idx) % 4
}

func findEdge(edge string, edges [4]string) int {
	for idx, e := range edges {
		if sameEdge(e, edge) {
			return idx
		}
	}
	return -1
}

func readTiles(reader io.Reader) (map[int]Tile, error) {
	var err error
	tiles := make(map[int]Tile)
	scanner := bufio.NewScanner(reader)

	var id int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Tile") {
			id, err = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(line, "Tile "), ":"))
			if err != nil {
				return nil, err
			}
			tiles[id] = make(Tile, 0, 10)
		} else if line != "" {
			tiles[id] = append(tiles[id], []byte(line))
		}
	}

	return tiles, nil
}

func sameEdge(edge1, edge2 string) bool {
	return edge1 == edge2 || edge1 == reverse(edge2)
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
