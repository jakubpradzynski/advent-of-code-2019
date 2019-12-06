import 'dart:collection';
import 'dart:io';

import 'dart:math';

void main(List<String> args) {
  new File(args[0]).readAsLines().then((List<String> wires) {
    var firstWirePoints = generateWirePoints(wires[0]);
    var secondWirePoints = generateWirePoints(wires[1]);
    var intersections = firstWirePoints.toSet().intersection(secondWirePoints.toSet());
    var intersectionsManhattanDistance = intersections.map((intersection) {
      return intersection.x.abs() + intersection.y.abs();
    });
    print("Result Part One = " +
        intersectionsManhattanDistance.reduce(min).toString());
        
    var intersectionsDistanceOnWirePath = intersections.map((intersection) {
      return 2 +
          firstWirePoints.toList().indexOf(intersection) +
          secondWirePoints.toList().indexOf(intersection);
    });
    print("Result Part Two = " +
        intersectionsDistanceOnWirePath.reduce(min).toString());
  });
}

Queue<Point> generateWirePoints(String wire) {
  var lastPoint = new Point(0, 0);
  var points = new Queue<Point<int>>();
  wire.split(",").forEach((step) {
    var direction = step[0];
    var value = int.parse(step.substring(1, step.length));
    for (var i = 1; i <= value; i++) {
      var newPoint = new Point(
          lastPoint.x + (direction == "R" ? 1 : (direction == "L" ? -1 : 0)),
          lastPoint.y + (direction == "U" ? 1 : (direction == "D" ? -1 : 0)));
      points.add(newPoint);
      lastPoint = newPoint;
    }
  });
  return points;
}
