import java.io.File

val path = if (args.contains("-f")) args[1 + args.indexOf("-f")] else "input.txt"

val input = File(path).useLines { it.first() }.split(",").map { it.toInt() }

fun partOne(input: List<Int>) {
    val mutableInput = input.toMutableList()
    mutableInput[1] = 12
    mutableInput[2] = 2
    val output = intcode(mutableInput)
    println("Part One result = ${output[0]}")
}

fun partTwo(input: List<Int>) {
    (0..100).forEach { noun ->
        run {
            (0..100).forEach { verb ->
                run {
                    val mutableInput = input.toMutableList()
                    mutableInput[1] = noun
                    mutableInput[2] = verb
                    if (intcode(mutableInput)[0] == 19690720) println("Part Two result = ${100 * noun + verb}")
                }
            }
        }
    }
}

fun intcode(input: MutableList<Int>): MutableList<Int> {
    var counter = 0
    while (counter < input.size) {
        when (input[counter]) {
            99 -> counter = input.size
            1 -> {
                val sum = input[input[counter + 1]] + input[input[counter + 2]]
                input[input[counter + 3]] = sum
                counter += 4
            }
            2 -> {
                val product = input[input[counter + 1]] * input[input[counter + 2]]
                input[input[counter + 3]] = product
                counter += 4
            }
            else -> throw RuntimeException("Unknown opcode!")
        }
    }
    return input
}

partOne(input)
partTwo(input)