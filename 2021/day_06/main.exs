defmodule AOC do

  def star1 do
    read_file()
    |> simulate_days(80)
    |> Enum.count()
  end

  def star2 do
    read_file()
    |> map_fishes()
    |> simulate_days(256)
    |> Map.values()
    |> Enum.sum()
  end

  # Star 1
  def simulate_days(fishes, 0), do: fishes
  def simulate_days(fishes, days) when is_list(fishes) do
    fishes
    |> Enum.flat_map(&simulate_fish_day/1)
    |> simulate_days(days - 1)
  end

  def simulate_fish_day(0), do: [6, 8]
  def simulate_fish_day(fish), do: [fish - 1]

  # Star 2
  def map_fishes(fishes) do
    Enum.reduce(
      fishes,
      Map.new(0..8, &({&1, 0})),
      fn fish, acc -> Map.update!(acc, fish, &(&1 + 1)) end
    )
  end

  def simulate_days(fishes, days) when is_map(fishes) do
    fishes
    |> Map.new(fn {k, v} -> {k - 1, v} end)
    |> create_new_fishes()
    |> simulate_days(days - 1)
  end

  def create_new_fishes(fishes) do
    {new, fishes} = Map.pop(fishes, -1, 0)
    fishes
    |> Map.put(8, new)
    |> Map.update!(6, &(&1 + new))
  end

  def read_file(name \\ "input.txt") do
    File.read!(name)
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
