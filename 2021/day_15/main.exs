defmodule AOC do
  def star1 do
    stream_file()
    |> generate_map()
    |> find_paths()
    |> with_max_pos()
    |> then(fn {map, max_pos} -> Map.get(map, max_pos) end)
  end

  def star2 do
    stream_file()
    |> generate_map()
    |> grow_map(5)
    |> find_paths()
    |> with_max_pos()
    |> then(fn {map, max_pos} -> Map.get(map, max_pos) end)
  end

  def find_paths(map), do: find_paths(map, %{{0, 0} => 0}, [{0, 0}], [])
  def find_paths(_, paths, [], []), do: paths
  def find_paths(map, paths, [], next_layer), do: find_paths(map, paths, next_layer, [])
  def find_paths(map, paths, [pos | curr_layer], next_layer) do
    neighbours =
      find_nearby(map, pos)
      |> Enum.filter(fn {p, v} -> !Map.has_key?(paths, p) || paths[pos] + v < paths[p] end)

    paths =
      neighbours
      |> Enum.reduce(paths, fn {p, v}, paths -> Map.put(paths, p, paths[pos] + v) end)

    find_paths(map, paths, curr_layer, next_layer ++ Enum.map(neighbours, &(elem(&1, 0))))
  end

  def find_nearby(map, pos) do
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

  def with_max_pos(map) do
    max_pos =
      map
      |> Map.keys()
      |> Enum.max_by(fn {x, y} -> x + y end)
    {map, max_pos}
  end

  def grow_map(map, times) do
    {_, max_pos} = with_max_pos(map)
    x_max = elem(max_pos, 0) + 1
    y_max = elem(max_pos, 1) + 1

    map
    |> Enum.reduce(%{}, fn {{x, y}, v}, map ->
      Enum.reduce(0..(times - 1), map, fn ry, map ->
        Enum.reduce(0..(times - 1), map, fn rx, map ->
          Map.put(map, {x + (rx * x_max), y + (ry * y_max)}, rem(v + (rx + ry) - 1, 9) + 1)
        end)
      end)
    end)
  end

  def generate_map(lines) do
    lines
    |> Enum.reduce(%{}, fn {line, y}, map ->
      Map.merge(
        map,
        Map.new(line, fn {v, x} -> {{x, y}, v} end)
      )
    end)
  end

  def print_map(map) do
    IO.write("\n")
    {_, {max_x, max_y}} = with_max_pos(map)


    Enum.each(0..max_y, fn y ->
      Enum.each(0..max_x, fn x ->
        IO.write(map[{x, y}] || "#")
      end)
      IO.write("\n")
    end)
    IO.write("\n")
    map
  end

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Enum.filter(&(&1 != ""))
    |> Stream.map(&parse_line/1)
    |> Enum.with_index()
  end

  def parse_line(line) do
    line
    |> String.trim()
    |> String.graphemes()
    |> Enum.map(&String.to_integer/1)
    |> Enum.with_index()
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
