const PHASE_COUNT = 100;
const REPEAT_NUMBER = 10000;
const PATTERN = [0, 1, 0, -1];
var INPUT = "59718730609456731351293131043954182702121108074562978243742884161871544398977055503320958653307507508966449714414337735187580549358362555889812919496045724040642138706110661041990885362374435198119936583163910712480088609327792784217885605021161016819501165393890652993818130542242768441596060007838133531024988331598293657823801146846652173678159937295632636340994166521987674402071483406418370292035144241585262551324299766286455164775266890428904814988362921594953203336562273760946178800473700853809323954113201123479775212494228741821718730597221148998454224256326346654873824296052279974200167736410629219931381311353792034748731880630444730593"
  .split("").map((digit) => int.parse(digit)).toList();


List<int> getPattern(int position, List<int> input) {
  var pattern = new List<int>();
  while (pattern.length <= input.length) {
    for (var i = 0; i < position; i++) {
      pattern.add(PATTERN[0]);
    }
    for (var i = 0; i < position; i++) {
      pattern.add(PATTERN[1]);
    }
    for (var i = 0; i < position; i++) {
      pattern.add(PATTERN[2]);
    }
    for (var i = 0; i < position; i++) {
      pattern.add(PATTERN[3]);
    }
  }
  return pattern.skip(1).take(input.length).toList();
}

int calculateNewValue(List<int> input, List<int> pattern) {
  var result = 0;
  for (var i = 0; i < input.length; i++) {
    result += input[i] * pattern[i];
  }
  return result.abs() % 10;
}

List<int> calculate(List<int> input) {
  var output = new List<int>();
  for (var i = 1; i <= input.length; i++) {
    var pattern = getPattern(i, input);
    output.add(calculateNewValue(input, pattern));
  }
  return output;
}

void main(List<String> args) {
  var input = INPUT;
  for (var i = 0; i < PHASE_COUNT; i++) {
    input = calculate(input);
  }
  print("Part One: " + input.take(8).map((digit) => digit.toString()).join(""));

  input = [];
  var offset = int.parse(INPUT.take(7).map((d) => d.toString()).join(""));
  for (var i = 0; i < REPEAT_NUMBER; i++) {
    input.addAll(INPUT);
  }
  input = input.skip(offset).toList();
  for (var i = 1; i <= PHASE_COUNT; i++) {
    var result = 0;
    for (var j = input.length - 1; j >= 0; j--) {
      result += input[j];
      result = result.abs() % 10;
      input[j] = result;
    }
  }
  print("Part Two: " + input.take(8).map((digit) => digit.toString()).join(""));
}
