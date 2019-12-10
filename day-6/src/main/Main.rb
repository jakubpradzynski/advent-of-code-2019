#!/usr/bin/ruby

class ObjectInSpace 
    def initialize(name, onOrbit)
        @name = name
        @onOrbit = onOrbit
    end

    def name
        @name
    end

    def onOrbit
        @onOrbit
    end
end

# Read input file
inputFilePath = ARGV[0]
inputFileLines = IO.readlines(inputFilePath)

# Convert lines to objects ObjectInSpace
objects = []
inputFileLines.each() do |line|
    split = line.delete("\n").split(")")
    objects << ObjectInSpace.new(split[1], split[0])
end

# Calculate checksum
sum = 0
objects.each() do |object|
    previousObjectName = object.onOrbit
    sum += 1
    while previousObjectName != "COM" do
        previousObject = objects.find { |o| o.name == previousObjectName }
        sum += 1
        previousObjectName = previousObject.onOrbit
    end
end
puts "Part One"
puts sum


class GraphObject
    def initialize(name)
        @name = name
        @neighbours = []
    end

    def name
        @name
    end

    def neighbours
        @neighbours
    end 
end

# Calculate minimal orbit transfers
# Create graph
graph = []
objects << ObjectInSpace.new("COM", nil)
objects.each() do |object|
    graphObject = GraphObject.new(object.name)
    if object.name != "COM" then
        onOrbitObject = objects.find { |o| o.name == object.onOrbit }
        graphObject.neighbours << onOrbitObject
    end
    objects.select { |o| o.onOrbit == object.name }.each() do |o|
        graphObject.neighbours << o
    end
    graph << graphObject
end
COM = GraphObject.new("COM")
COM.neighbours << objects.select { |o| o.onOrbit == "COM" }
graph << COM


# BFS
YOU = graph.find { |g| g.name == "YOU" }
SAN = graph.find { |g| g.name == "SAN" }
visited = Hash.new
distance = Hash.new
queue = Queue.new()
distance[YOU.name] = 0
queue << YOU
visited[YOU.name] = true
while !queue.empty? do
    x = queue.pop()
    length = x.neighbours.length
    for i in 0..length-1
        if !visited[x.neighbours[i].name] then
            distance[x.neighbours[i].name] = distance[x.name] + 1
            queue << graph.find { |g| g.name == x.neighbours[i].name }
            visited[x.neighbours[i].name] = true
        end
    end
end
puts "Part Two"
puts distance[SAN.name] - 2