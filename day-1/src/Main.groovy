class Main {
    static void main(String[] args) {
        partOne(args[0])
        partTwo(args[0])
    }

    private static void partOne(String filepath) {
        def sumOfModulesFuel = 0
        new File(filepath).eachLine { line ->
            sumOfModulesFuel += calculateFuelForModule(line as long)
        }
        println("Part One result = $sumOfModulesFuel")
    }

    private static void partTwo(String filepath) {
        def sumOfModulesFuel = 0
        new File(filepath).eachLine { line ->
            def moduleMass = line as long
            while (true) {
                def moduleFuel = calculateFuelForModule(moduleMass)
                if (moduleFuel > 0) {
                    sumOfModulesFuel += moduleFuel
                    moduleMass = moduleFuel
                } else {
                    break
                }
            }
        }
        println("Part Two result = $sumOfModulesFuel")
    }

    private static long calculateFuelForModule(long moduleMass) {
        return Math.floor(moduleMass / 3) - 2
    }
}
