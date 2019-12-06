import sys

input_range = sys.argv[1].split("-")


def value_never_decreases(value):
    last_digit = 0
    for digit in value:
        if (last_digit > int(digit)):
            return False
        last_digit = int(digit)
    return True


def value_has_two_the_same_adjecent_digits(value):
    result = False
    for i in range(0, len(value) - 1):
        if value[i] == value[i + 1]:
            result = True
    return result


def value_has_exactly_two_the_same_adjecent_digits(value):
    counter = 1
    result = False
    for i in range(0, len(value) - 1):
        if value[i] == value[i + 1]:
            counter = counter + 1
        else:
            if (counter == 2):
                result = True
            counter = 1
    if (counter == 2):
        result = True
    return result


counter = 0
for value in range(int(input_range[0]), int(input_range[1])):
    if value_never_decreases(str(value)) & value_has_two_the_same_adjecent_digits(str(value)):
        counter = counter + 1

print("Part One result = " + str(counter))


counter = 0
for value in range(int(input_range[0]), int(input_range[1])):
    if value_never_decreases(str(value)) & value_has_exactly_two_the_same_adjecent_digits(str(value)):
        counter = counter + 1

print("Part Two result = " + str(counter))

assert value_has_exactly_two_the_same_adjecent_digits('112233') == True
assert value_has_exactly_two_the_same_adjecent_digits('123444') == False
assert value_has_exactly_two_the_same_adjecent_digits('111122') == True
assert value_has_exactly_two_the_same_adjecent_digits('111221') == True
