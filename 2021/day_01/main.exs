defmodule AOC do
  def star1 do
    stream_file()
    |> Stream.chunk_every(2, 1, :discard)
    |> Enum.reduce(0, &count_inc/2)
  end

  def star2 do
    stream_file()
    |> Stream.chunk_every(3, 1, :discard)
    |> Stream.map(&Enum.sum/1)
    |> Stream.chunk_every(2, 1, :discard)
    |> Enum.reduce(0, &count_inc/2)
  end

  def count_inc([prev, curr], total) when curr > prev, do: total + 1
  def count_inc(_, total), do: total

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Enum.filter(&(&1 != ""))
    |> Stream.map(&String.to_integer/1)
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
