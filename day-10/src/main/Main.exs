defmodule Point do
  defstruct x: 0, y: 0
end

defmodule Utils do
  def prepareAsteroids(content) do
    lines = String.split(String.trim(content), "\n")

    Stream.with_index(lines, 0)
    |> Enum.map(fn {line, i} ->
      Stream.with_index(String.split(line, ""), 0)
      |> Enum.map(fn {mark, j} ->
        if mark == "#" do
          {j, i}
        end
      end)
    end)
    |> List.flatten()
    |> Enum.filter(&(!is_nil(&1)))
  end

  def angle({x, y}, {a, b}) do
    angle =
      :math.atan2(x - a, b - y) * 180 /
        :math.pi()

    if angle < 0 do
      360 + angle
    else
      angle
    end
  end

  def distance({x, y}, {a, b}) do
    :math.sqrt((a - x) * (a - x) + (b - y) * (b - y))
  end
end

if length(System.argv()) != 1 do
  IO.puts("You should pass path to input file!")
  System.halt(1)
end

filePath = Enum.at(System.argv(), 0)

asteroids =
  case File.read(filePath) do
    {:ok, content} ->
      Utils.prepareAsteroids(content)

    {:error, :enoent} ->
      IO.puts("Could not read file " + filePath)
      System.halt(1)
  end

# Part One
IO.puts("Part One (x, y, visible asteroids):")

bestLocation =
  Enum.map(asteroids, fn {x, y} ->
    {x, y,
     Enum.map(asteroids, fn {a, b} ->
       if a == x && b == y do
         nil
       else
         Utils.angle({x, y}, {a, b})
       end
     end)
     |> Enum.filter(&(!is_nil(&1)))
     |> Enum.uniq()
     |> length}
  end)
  |> Enum.sort_by(&elem(&1, 2))
  |> List.last()
  |> Tuple.to_list()

bestLocation |> Enum.join(", ") |> IO.puts()

# Part Two
location = {Enum.at(bestLocation, 0), Enum.at(bestLocation, 1)}

asteroids
|> Stream.reject(&(&1 == location))
|> Enum.map(fn {x, y} ->
  {
    {x, y},
    Utils.distance({x, y}, location),
    Utils.angle({x, y}, location)
  }
end)
|> Enum.sort_by(fn {_, distance, angle} -> {angle, distance} end)
|> Enum.chunk_by(fn {_, _, angle} -> angle end)
|> Enum.map(&hd(&1))
|> Enum.map(fn {{x, y}, _, _} -> x * 100 + y end)
|> Enum.to_list()
|> Enum.at(196)
|> IO.puts()
