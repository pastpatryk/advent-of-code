defmodule AOC do
  def star1 do
    read_file()
    |> find_best_position(&dist_abs/2)
  end

  def star2 do
    read_file()
    |> find_best_position(&dist_sum/2)
  end

  def find_best_position(crabs, dist_fn) do
    {min, max} = Enum.min_max(crabs)
    (min..max)
    |> Enum.map(fn pos -> crabs_cost(crabs, pos, dist_fn) end)
    |> Enum.min_by(&(elem(&1, 1)))
  end

  def crabs_cost(crabs, pos, dist_fn) do
    cost = Enum.reduce(crabs, 0, fn crab, acc -> acc + dist_fn.(crab, pos) end)
    {pos, cost}
  end

  def dist_abs(a, b), do: abs(a - b)

  def dist_sum(a, b), do: Enum.sum(1..abs(a - b))

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
