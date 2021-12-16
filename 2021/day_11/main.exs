defmodule AOC do
  def star1 do
    stream_file()
    |> tap(&print/1)
    |> simulate_days(1000)
    |> elem(1)
  end

  def star2 do
    stream_file("test.txt")
  end

  def simulate_days(map, days), do: simulate_days(map, 0, days)

  def simulate_days(map, flashes_count, 0), do: {map, flashes_count}

  def simulate_days(map, flashes_count, days) do
    {map, count} =
      map
      |> inc_energy()
      |> flash()

      IO.puts("[Day: #{days}] Flashes: #{flashes_count + count}")
      print(map)

      if count == 100 do
        raise "All flashed"
      end

    simulate_days(map, flashes_count + count, days - 1)
  end

  def inc_energy(map) do
    map
    |> Enum.reduce(%{}, fn {pos, val}, map ->
      Map.put(map, pos, val + 1)
    end)
  end

  def flash(map), do: flash(map, [], -1)

  def flash(map, total_flashed, 0 = _last_count) do
    map =
      total_flashed
      |> Enum.reduce(map, fn pos, map ->
        Map.put(map, pos, 0)
      end)

    {map, length(total_flashed)}
  end

  def flash(map, total_flashed, _last_count) do
    flashes =
      map
      |> Enum.filter(fn {_, val} -> val > 9 end)
      |> Enum.map(fn {pos, _} -> pos end)

    new_flashes = flashes -- total_flashed

    map =
      new_flashes
      |> Enum.reduce(map, &flash_single/2)

    flash(map, total_flashed ++ new_flashes, length(new_flashes))
  end

  def flash_single(pos, map) do
    find_nearby(pos, map)
    |> Enum.reduce(map, fn {pos, _}, map ->
      Map.update!(map, pos, &(&1 + 1))
    end)
  end

  def find_nearby(pos, map) do
    nearby_positions(pos)
    |> Enum.map(fn p -> {p, Map.get(map, p)} end)
    |> Enum.filter(fn {_, v} -> v != nil end)
  end

  def nearby_positions({x, y}) do
    [
      {x - 1, y - 1},
      {x - 1, y},
      {x - 1, y + 1},
      {x, y - 1},
      {x, y + 1},
      {x + 1, y - 1},
      {x + 1, y},
      {x + 1, y + 1},
    ]
  end

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Stream.filter(fn line -> line != "\n" end)
    |> Stream.map(&parse_line/1)
    |> generate_map()
  end

  def parse_line(line) do
    line
    |> String.trim()
    |> String.graphemes()
    |> Enum.map(&String.to_integer/1)
  end

  def generate_map(lines) do
    lines
    |> Stream.with_index()
    |> Enum.reduce(%{}, &generate_map_line/2)
  end

  def generate_map_line({line, y}, map) do
    line
    |> Enum.with_index()
    |> Enum.reduce(map, fn {height, x}, map -> Map.put(map, {x, y}, height) end)
  end

  def print(map) do
    IO.puts("-------")
    for y <- 0..9, x <- 0..9 do
      IO.write(map[{x, y}])
      if x == 9 do
        IO.write("\n")
      end
    end
    IO.puts("-------")
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
