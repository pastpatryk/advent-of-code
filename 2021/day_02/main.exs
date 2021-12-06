defmodule AOC do
  def star1 do
    stream_file()
    |> Enum.reduce({0, 0}, &move/2)
  end

  def star2 do
    stream_file()
    |> Enum.reduce({0, 0, 0}, &move2/2)
  end

  def move({"forward", val}, {x, y}), do: {x + val, y}
  def move({"down", val}, {x, y}), do: {x, y + val}
  def move({"up", val}, {x, y}), do: {x, y - val}

  def move2({"forward", val}, {x, y, aim}), do: {x + val, y + (aim * val), aim}
  def move2({"down", val}, {x, y, aim}), do: {x, y, aim + val}
  def move2({"up", val}, {x, y, aim}), do: {x, y, aim - val}

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Stream.map(&parse_line/1)
  end

  def parse_line(line) do
    [cmd, valueStr] = String.split(line, " ")
    {cmd, String.to_integer(valueStr)}
  end
end

IO.puts("⭐️")
{x, y} = AOC.star1() |> IO.inspect()
IO.puts(x * y)

IO.puts("⭐️⭐️")
{x, y, _} = AOC.star2() |> IO.inspect()
IO.puts(x * y)
