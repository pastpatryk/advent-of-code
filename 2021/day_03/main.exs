defmodule AOC do
  def star1 do
    common_bits = stream_file() |> most_common_bits()

    gamma = common_bits |> rate_from_bits(true)
    epsilon = common_bits |> rate_from_bits(false)

    gamma * epsilon
  end

  def star2 do
    lines = stream_file() |> Enum.to_list()

    ox_rating = lines |> find_rating(true)
    co2_rating = lines |> find_rating(false)

    ox_rating * co2_rating
  end

  def find_rating(bits_list, flip), do: find_rating(bits_list, flip, 0)
  def find_rating([bits], _, _), do: bits |> :erlang.list_to_integer(2)
  def find_rating(bits_list, flip, pos) do
    common = bits_list |> most_common_at(pos) |> bit_for_val(flip)

    bits_list
    |> Enum.filter(fn bits -> has_bit?(bits, pos, common) end)
    |> find_rating(flip, pos + 1)
  end

  def has_bit?(bits, pos, 0), do: Enum.at(bits, pos) == ?0
  def has_bit?(bits, pos, 1), do: Enum.at(bits, pos) == ?1

  def rate_from_bits(bits, flip) do
    bits
    |> Enum.map(fn {_, val} -> bit_for_val(val, flip) end)
    |> Enum.join()
    |> String.to_integer(2)
  end

  def bit_for_val(val, true) when val >= 0, do: 1
  def bit_for_val(_, true), do: 0
  def bit_for_val(val, false) when val >= 0, do: 0
  def bit_for_val(_, false), do: 1

  def most_common_bits(bits_list), do: Enum.reduce(bits_list, %{}, &most_common_bit/2)

  def most_common_bit(bits, counts) do
    bits
    |> Enum.with_index()
    |> Enum.reduce(counts, &update_count/2)
  end

  def update_count({bit, idx}, counts),
    do: Map.update(counts, idx, value_for(bit), &(&1 + value_for(bit)))

  def most_common_at(bits_list, pos) do
    bits_list
    |> Enum.map(&(Enum.at(&1, pos)))
    |> Enum.reduce(0, fn bit, acc -> acc + value_for(bit) end)
  end

  def value_for(?0), do: -1
  def value_for(?1), do: 1

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.to_charlist/1)
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
