defmodule AOC do
  @packet_type %{
    0 => :sum,
    1 => :product,
    2 => :min,
    3 => :max,
    5 => :gt,
    6 => :lt,
    7 => :eq,
  }

  def star1 do
    read_transmission()
    |> parse_packet()
    |> elem(1)
    |> sum_versions
  end

  def star2 do
    read_transmission()
    |> parse_packet()
    |> elem(1)
    |> calculate()
  end

  def sum_versions({version, :literal, _}), do: version
  def sum_versions({version, _, subpackets}) do
    version + (Enum.map(subpackets, &sum_versions/1) |> Enum.sum())
  end

  def calculate({_, :literal, val}), do: val
  def calculate({_, :sum, subpackets}), do: reduce(subpackets, &Enum.sum/1)
  def calculate({_, :product, subpackets}), do: reduce(subpackets, &Enum.product/1)
  def calculate({_, :min, subpackets}), do: reduce(subpackets, &Enum.min/1)
  def calculate({_, :max, subpackets}), do: reduce(subpackets, &Enum.max/1)
  def calculate({_, :gt, [p1, p2]}), do: bool_to_int(calculate(p1) > calculate(p2))
  def calculate({_, :lt, [p1, p2]}), do: bool_to_int(calculate(p1) < calculate(p2))
  def calculate({_, :eq, [p1, p2]}), do: bool_to_int(calculate(p1) == calculate(p2))

  def reduce(subpackets, fun), do: subpackets |> Enum.map(&calculate/1) |> fun.()

  def bool_to_int(true), do: 1
  def bool_to_int(false), do: 0

  def parse_packet(<< version::3, 1::1, 0::1, 0::1, trans::bitstring  >>) do
    {trans, literal} = parse_literal(trans)
    {trans, {version, :literal, literal}}
  end

  def parse_packet(<< version::3, packet_type::3, 0::1, packets_len::15, packets::size(packets_len), trans::bitstring >>) do
    << subpackets::bitstring >> = << packets::size(packets_len) >>
    {trans, {version, @packet_type[packet_type], parse_subpackets(subpackets)}}
  end

  def parse_packet(<< version::3, packet_type::3, 1::1, packets_count::11, trans::bitstring >>) do
    {trans, subpackets} = parse_subpackets(trans, packets_count)
    {trans, {version, @packet_type[packet_type], subpackets}}
  end

  def parse_literal(trans), do: parse_literal(trans, <<>>)
  def parse_literal(<< 1::1, num::4, trans::bitstring >>, literal) do
    parse_literal(trans, <<literal::bitstring, num::4>>)
  end
  def parse_literal(<< 0::1, num::4, trans::bitstring >>, literal) do
    literal = <<literal::bitstring, num::4>>
    s = bit_size(literal)
    << x::size(s) >> = literal
    {trans, x}
  end

  def parse_subpackets(<<>>), do: []
  def parse_subpackets(trans) do
    {trans, packet} = parse_packet(trans)
    [packet | parse_subpackets(trans)]
  end

  def parse_subpackets(trans, 0), do: {trans, []}
  def parse_subpackets(trans, count) do
    {trans, packet} = parse_packet(trans)
    {trans, subpackets} = parse_subpackets(trans, count - 1)
    {trans, [packet | subpackets]}
  end

  def read_transmission(name \\ "input.txt") do
    File.read!(name)
    |> String.replace("\n", "")
    |> Base.decode16!()
  end
end

IO.puts("⭐️")
AOC.star1() |> IO.inspect()

IO.puts("⭐️⭐️")
AOC.star2() |> IO.inspect()
