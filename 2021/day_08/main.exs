defmodule AOC do
  @positions ["a", "b", "c", "d", "e", "f", "g"]

  @codes %{
    0 => MapSet.new(["a", "b", "c", "e", "f", "g"]),
    1 => MapSet.new(["c", "f"]),
    2 => MapSet.new(["a", "c", "d", "e", "g"]),
    3 => MapSet.new(["a", "c", "d", "f", "g"]),
    4 => MapSet.new(["b", "c", "d", "f"]),
    5 => MapSet.new(["a", "b", "d", "f", "g"]),
    6 => MapSet.new(["a", "b", "d", "e", "f", "g"]),
    7 => MapSet.new(["a", "c", "f"]),
    8 => MapSet.new(["a", "b", "c", "d", "e", "f", "g"]),
    9 => MapSet.new(["a", "b", "c", "d", "f", "g"]),
  }

  @codes_map Map.new(@codes, fn {key, val} -> {val, key} end)

  def star1 do
    stream_file()
    |> Stream.map(&count_known/1)
    |> Enum.sum()
  end

  def star2 do
    all_mappings =
      permutations(@positions)
      |> Stream.map(fn perm -> Map.new(Enum.zip(perm, @positions)) end)

    stream_file()
    |> Enum.map(fn codes -> find_output(all_mappings, codes) end)
    |> Enum.sum()
  end

  def find_output(all_mappings, [signal, output]) do
    mapping = find_mapping(all_mappings, signal ++ output)
    output
    |> Enum.map(fn code -> Map.get(@codes_map, map_code(mapping, code)) end)
    |> Enum.join("")
    |> String.to_integer()
  end

  def find_mapping(all_mappings, codes) do
    all_mappings
    |> Enum.find(fn mapping -> valid_mapping?(mapping, codes) end)
  end

  def valid_mapping?(mapping, codes) do
    codes
    |> Enum.all?(fn code -> Map.has_key?(@codes_map, map_code(mapping, code)) end)
  end

  def map_code(mapping, code) do
    String.graphemes(code)
    |> Enum.map(&(Map.get(mapping, &1)))
    |> MapSet.new()
  end

  def permutations([]), do: [[]]
  def permutations(list), do: for elem <- list, rest <- permutations(list--[elem]), do: [elem|rest]

  def count_known([_singnal, output]) do
    output
    |> Enum.count(fn code -> Enum.member?([2,3,4,7], String.length(code)) end)
  end

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Stream.map(&parse_line/1)
  end

  def parse_line(line) do
    line
    |> String.split(" | ")
    |> Enum.map(&String.split/1)
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
