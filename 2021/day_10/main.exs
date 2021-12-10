defmodule AOC do

  @opening ["(", "[", "{", "<"]
  @closing [")", "]", "}", ">"]
  @mapping Enum.zip(@opening, @closing) |> Map.new()
  @scoring %{")" => 3, "]" => 57, "}" => 1197, ">" => 25137}
  @scoring_incomplete %{")" => 1, "]" => 2, "}" => 3, ">" => 4}

  def star1 do
    stream_file()
    |> Enum.map(&parse_line/1)
    |> Enum.filter(fn res -> elem(res, 0) == :error end)
    |> Enum.map(fn {:error, sym} -> @scoring[sym] end)
    |> Enum.sum()
  end

  def star2 do
    stream_file()
    |> Enum.map(&parse_line/1)
    |> Enum.filter(fn res -> elem(res, 0) == :ok end)
    |> Enum.map(&score_incomplete/1)
    |> Enum.sort()
    |> then(fn scores -> Enum.at(scores, length(scores) |> div(2)) end)
  end

  def parse_line(line), do: parse_line(line, [])

  def parse_line([sym | line], stack) when sym in @opening do
    parse_line(line, [@mapping[sym] | stack])
  end

  def parse_line([sym | line], [sym | stack]) when sym in @closing do
    parse_line(line, stack)
  end

  def parse_line([sym | _], _), do: {:error, sym}

  def parse_line([], stack), do: {:ok, stack}

  def score_incomplete({:ok, incomplete}) do
    incomplete
    |> Enum.reduce(0, fn sym, acc ->
      (acc * 5) + @scoring_incomplete[sym]
    end)
  end

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.graphemes/1)
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
