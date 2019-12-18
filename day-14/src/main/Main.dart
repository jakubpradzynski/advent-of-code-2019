import 'dart:collection';
import 'dart:io';

const FUEL = "FUEL";
const ORE = "ORE";
const TRILLION = 1000000000000;

class Chemical {
  int amount;
  String name;

  Chemical(this.amount, this.name);

  String toString() {
    return 'Chemical(amount= ${amount}, name = ${name}';
  }  
}

class Reaction {
  Chemical output;
  List<Chemical> inputs;

  Reaction(this.output, this.inputs);

  String toString() {
    return 'Reaction(output = ${output}, inputs = ${inputs}';
  }
}

class Nanofactory {
  List<Reaction> reactions;
  HashMap<String, int> fuelComponents;
  HashMap<String, int> leftovers;
  Nanofactory(this.reactions) {
    this.fuelComponents = new HashMap();
    this.leftovers = new HashMap();
  }
  
  int cost(String name, int n) {
    leftovers.putIfAbsent(name, () => 0);
    if (name == ORE) {
      return n;
    }
    if (leftovers[name] >= n) {
      leftovers[name] -= n;
      return 0;
    }
    if (leftovers[name] > 0) {
      n -= leftovers[name];
      leftovers[name] = 0;
    }

    var reaction = findReactionByOutputName(name);
    var iterationsCount = ((n.toDouble() / reaction.output.amount.toDouble()).ceil()).toInt();

    var oreCount = reaction.inputs.map((input) {
      return cost(input.name, input.amount * iterationsCount);
    }).reduce((a,b) => a + b);

    leftovers[name] += reaction.output.amount * iterationsCount - n;
    return oreCount;
  }

  int cleanAndCalculateCost(String name, int n) {
    leftovers = new HashMap();
    fuelComponents = new HashMap();
    return cost(name, n);
  }

  Reaction findReactionByOutputName(String name) {
    return reactions.firstWhere((reaction) {
      return reaction.output.name == name;
    });
  }

}

List<Reaction> getReactions(List<String> fileLines) {
  return fileLines.map((line) {
    var split = line.split(" => ");
    var output = split[1].trim().split(" ");
    var inputs = split[0].split(", ");
    return new Reaction(new Chemical(int.parse(output[0]), output[1]), inputs.map((input) {
      var split = input.trim().split(" ");
      return new Chemical(int.parse(split[0]), split[1]);
    }).toList());
  }).toList();
}

void main(List<String> args) {
  var filePath = "day-14/src/resources/input.txt";
  if (args.length != 0) {
    filePath = args[0];
  }
  var inputLines = new File(filePath).readAsLinesSync();
  var reactions = getReactions(inputLines);
  var nanofactory = new Nanofactory(reactions);
  var cost = nanofactory.cost(FUEL, 1);
  print("Part One - " + cost.toString());

  var maximumAmountOfFuel = 0;
  for (var i = 0; i < TRILLION; i++) {
    if (nanofactory.cleanAndCalculateCost(FUEL, i) >= TRILLION) {
      maximumAmountOfFuel = i - 1;
      break;
    }
  }
  // var maximumAmountOfFuel = new List<bool>.generate(TRILLION, (i) => nanofactory.cleanAndCalculateCost(FUEL, i) >= TRILLION).indexOf(true) - 1;
  print("Part Two - " + maximumAmountOfFuel.toString());
}
