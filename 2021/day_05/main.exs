defmodule AOC do
  def star1 do
    stream_file()
    |> Enum.filter(fn [{x1, y1}, {x2, y2}] -> x1 == x2 || y1 == y2 end)
    |> Enum.reduce(%{}, &mark_line/2)
    |> Enum.count(fn {_, val} -> val > 1 end)
  end

  def star2 do
    stream_file()
    |> Enum.reduce(%{}, &mark_line/2)
    |> Enum.count(fn {_, val} -> val > 1 end)
  end

  def mark_line([p1, p2], space) do
    points_on_line(p1, p2)
    |> Enum.reduce(space, fn p, acc -> Map.update(acc, p, 1, &(&1 + 1)) end)
  end

  def points_on_line({x1, y}, {x2, y}), do: Enum.map(x1..x2, fn x -> {x, y} end)
  def points_on_line({x, y1}, {x, y2}), do: Enum.map(y1..y2, fn y -> {x, y} end)
  def points_on_line({x1, y1}, {x2, y2}), do: Enum.zip(x1..x2, y1..y2)

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Enum.filter(&(&1 != ""))
    |> Stream.map(&parse_line/1)
  end

  def parse_line(line) do
    line
    |> String.trim()
    |> String.split(" -> ")
    |> Enum.map(&parse_coords/1)
  end

  def parse_coords(coords) do
    [x, y] =
      coords
      |> String.split(",")
      |> Enum.map(&String.to_integer/1)

    {x, y}
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
