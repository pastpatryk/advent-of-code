defmodule AOC do
  def star1 do
    {map, folds} = stream_file()

    fold(Enum.at(folds, 0), map)
    |> count_dots()
  end

  def star2 do
    {map, folds} = stream_file()

    folds
    |> Enum.reduce(map, &fold/2)
    |> print()
  end

  def count_dots(map) do
    map
    |> Enum.count(fn {_, v} -> v end)
  end

  def print(map) do
    max_x = map |> Enum.map(fn {{x, _}, _} -> x end) |> Enum.max()
    max_y = map |> Enum.map(fn {{_, y}, _} -> y end) |> Enum.max()

    Enum.each(0..max_y, fn y ->
      Enum.each(0..max_x, fn x ->
         if Map.get(map, {x, y}) do
          IO.write("#")
         else
          IO.write(".")
         end
        end)
        IO.write("\n")
    end)
  end

  def fold({:fold, :x, fold_x}, map) do
    Enum.reduce(map, %{}, fn {{x, y}, dot}, map ->
      if x < fold_x do
        Map.put(map, {x, y}, dot)
      else
        Map.put(map, {fold_x - (x - fold_x), y}, dot)
      end
    end)
  end

  def fold({:fold, :y, fold_y}, map) do
    Enum.reduce(map, %{}, fn {{x, y}, dot}, map ->
      if y < fold_y do
        Map.put(map, {x, y}, dot)
      else
        Map.put(map, {x, fold_y - (y - fold_y)}, dot)
      end
    end)
  end

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Enum.filter(&(&1 != "" && &1 != "\n"))
    |> Stream.map(&parse_line/1)
    |> generate_map_and_folds()
  end

  def generate_map_and_folds(lines) do
    %{fold: folds, dot: dots} = Enum.group_by(lines, &line_type/1)
    map =
      dots
      |> Enum.reduce(%{}, fn [x, y], map ->
        Map.put(map, {x, y}, true)
      end)
    {map, folds}
  end

  def line_type({:fold, _, _}), do: :fold
  def line_type(_), do: :dot

  def parse_line("fold along x=" <> fold_line) do
    {:fold, :x, parse_fold_line(fold_line)}
  end

  def parse_line("fold along y=" <> fold_line) do
    {:fold, :y, parse_fold_line(fold_line)}
  end

  def parse_line(line) do
    line
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
  end

  def parse_fold_line(fold_line) do
    fold_line
    |> String.trim()
    |> String.to_integer()
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
