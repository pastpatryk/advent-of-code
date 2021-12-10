defmodule AOC do
  def star1 do
    stream_file()
    |> find_lowest()
    |> Enum.map(fn {_, v} -> v + 1 end)
    |> Enum.sum()
  end

  def star2 do
    stream_file()
    |> find_basins()
    |> IO.inspect()
    |> Enum.map(&Enum.count/1)
    |> Enum.sort(:desc)
    |> Enum.take(3)
    |> Enum.reduce(&(&1 * &2))
  end

  def find_lowest(map) do
    map
    |> Enum.filter(fn {pos, val} -> check_lowest({pos, val}, map) end)
  end

  def check_lowest({pos, val}, map) do
    find_nearby(pos, map)
    |> Enum.all?(fn {_, v} -> v > val end)
  end

  def find_basins(map) do
    map
    |> find_lowest()
    |> Enum.map(fn {pos, val} -> find_basin({pos, val}, map) end)
  end

  def find_basin({pos, val}, map) do
    find_nearby(pos, map)
    |> Enum.filter(fn {_, v} -> v > val && v < 9 end)
    |> Enum.reduce(MapSet.new(), fn {p, v}, b -> MapSet.union(b, find_basin({p, v}, map)) end)
    |> MapSet.put({pos, val})
  end

  def find_nearby(pos, map) do
    nearby_positions(pos)
    |> Enum.map(fn p -> {p, Map.get(map, p)} end)
    |> Enum.filter(fn {_, v} -> v != nil end)
  end

  def nearby_positions({x, y}) do
    [
      {x - 1, y},
      {x + 1, y},
      {x, y - 1},
      {x, y + 1}
    ]
  end

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Stream.filter(fn line -> line != "\n" end)
    |> Stream.map(&parse_line/1)
    |> generate_map()
  end

  def parse_line(line) do
    line
    |> String.trim()
    |> String.graphemes()
    |> Enum.map(&String.to_integer/1)
  end

  def generate_map(lines) do
    lines
    |> Stream.with_index()
    |> Enum.reduce(%{}, &generate_map_line/2)
  end

  def generate_map_line({line, y}, map) do
    line
    |> Enum.with_index()
    |> Enum.reduce(map, fn {height, x}, map -> Map.put(map, {x, y}, height) end)
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
