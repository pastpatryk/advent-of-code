defmodule AOC do
  def star1 do
    stream_file()
    |> generate_graph()
    |> find_paths()
    |> Enum.count()
  end

  def star2 do
    stream_file()
    |> generate_graph()
    |> find_paths2()
    |> Enum.count()
  end

  def find_paths(graph), do: find_paths(graph, ["start"], &can_visit?/2)
  def find_paths2(graph), do: find_paths(graph, ["start"], &can_visit2?/2)

  def find_paths(_, ["end" | _] = path, _), do: [path]

  def find_paths(graph, [node | _] = path, can_visit_fn) do
    graph[node]
    |> Enum.filter(fn n -> can_visit_fn.(n, path) end)
    |> Enum.flat_map(fn n -> find_paths(graph, [n | path], can_visit_fn) end)
  end

  def can_visit?(node, path) do
    !(node == String.downcase(node) && Enum.member?(path, node))
  end

  def can_visit2?("start", _), do: false
  def can_visit2?(node, path) do
    !(lowercase?(node) && Enum.member?(path, node) && small_visited_twice?(path))
  end

  def small_visited_twice?(path) do
    path
      |> Enum.filter(&lowercase?/1)
      |> Enum.group_by(&(&1))
      |> Enum.any?(fn {_, v} -> length(v) == 2 end)
  end

  def lowercase?(text), do: text == String.downcase(text)

  def stream_file(name \\ "input.txt") do
    File.stream!(name)
    |> Enum.filter(&(&1 != ""))
    |> Stream.map(&parse_line/1)
  end

  def parse_line(line) do
    line
    |> String.trim()
    |> String.split("-")
  end

  def generate_graph(lines) do
    lines
    |> Enum.reduce(%{}, fn [n1, n2], graph ->
      graph
      |> add_edge(n1, n2)
      |> add_edge(n2, n1)
    end)
  end

  def add_edge(graph, n1, n2) do
    Map.update(graph, n1, MapSet.new([n2]), fn set ->
      MapSet.put(set, n2)
    end)
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
