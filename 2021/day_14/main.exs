defmodule AOC do
  def star1 do
    stream_file()
    |> grow(10)
    |> Enum.reduce(%{}, fn el, acc -> map_inc(acc, el) end)
    |> Enum.min_max_by(fn {_, v} -> v end)
    |> then(fn {{_, min}, {_, max}} -> max - min end)
  end

  def star2 do
    stream_file()
    |> then(fn {polymer, mapping} -> {pairs_count(polymer), mapping} end)
    |> grow_counts(40)
    |> Enum.reduce(%{}, fn {[a, b], count}, acc ->
      acc
      |> map_inc(a, count / 2)
      |> map_inc(b, count / 2)
    end)
    |> Enum.map(fn {_, v} -> ceil(v) end)
    |> Enum.min_max()
    |> then(fn {min, max} -> max - min end)
  end

  def grow({polymer, mapping}, steps), do: grow(polymer, mapping, steps)
  def grow(polymer, _, 0), do: polymer
  def grow(polymer, mapping, steps) do
    polymer
    |> Enum.chunk_every(2, 1)
    |> Enum.flat_map(&(new_pair(&1, mapping)))
    |> grow(mapping, steps - 1)
  end

  def new_pair([a], _), do: [a]
  def new_pair([a, b], mapping) when is_map_key(mapping, a <> b) do
    [a, mapping[a <> b]]
  end

  def pairs_count(polymer) do
    polymer
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.reduce(%{}, fn pair, acc -> map_inc(acc, pair) end)
  end

  def grow_counts({counts, mapping}, steps), do: grow_counts(counts, mapping, steps)
  def grow_counts(counts, _, 0), do: counts
  def grow_counts(counts, mapping, steps) do
    counts
    |> Enum.reduce(%{}, fn {[a, b], count}, acc ->
      c = mapping[a <> b]
      acc
      |> map_inc([a, c], count)
      |> map_inc([c, b], count)
    end)
    |> grow_counts(mapping, steps - 1)
  end

  def map_inc(map, key, inc \\ 1), do: Map.update(map, key, inc, &(&1 + inc))

  def stream_file(name \\ "input.txt") do
    [template | lines] =
      File.stream!(name)
      |> Enum.map(&String.trim/1)
      |> Enum.filter(&(&1 != ""))

    mapping =
      Enum.map(lines, &(String.split(&1, " -> ")))
      |> Map.new(fn [k, v] -> {k, v} end)

    {String.graphemes(template), mapping}
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
