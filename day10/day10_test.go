package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {
	instruction := []string{
		"noop",
		"noop",
		"addx 1",
		"addx 15",
		"addx -15",
		"addx -15",
		"addx 15",
		"addx -15",
		"addx 15",
	}
	register := []int{
		0,
		-5,
		0,
		0,
		0,
		10,
		10,
		-10,
		-10,
	}
	wants := []int{
		0,
		-5,
		1,
		15,
		-15,
		-5,
		25,
		-25,
		5,
	}

	for testIndex, want := range wants {
		result := Process(instruction[testIndex], register[testIndex])
		assert.Equal(t, want, result)
	}
}

func TestRunInstructions(t *testing.T) {
    input:= [][]string{
        {
            "addx 1",
            "noop",
            "noop",
            "noop",
            "addx 5",
            "addx 5",
            "noop",
            "noop",
            "addx 9",
            "addx -5",
            "addx 1",
            "addx 4",
            "noop",
            "noop",
            "noop",
            "addx 6",
            "addx -1",
            "noop",
            "addx 5",
            "addx -2",
            "addx 7",
            "noop",
            "addx 3",
            "addx -2",
            "addx -38",
            "noop",
            "noop",
            "addx 32",
            "addx -22",
            "noop",
            "addx 2",
            "addx 3",
            "noop",
            "addx 2",
            "addx -2",
            "addx 7",
            "addx -2",
            "noop",
            "addx 3",
            "addx 2",
            "addx 5",
            "addx 2",
            "addx -5",
            "addx 10",
            "noop",
            "addx 3",
            "noop",
            "addx -38",
            "addx 1",
            "addx 27",
            "noop",
            "addx -20",
            "noop",
            "addx 2",
            "addx 27",
            "noop",
            "addx -22",
            "noop",
            "noop",
            "noop",
            "noop",
            "addx 3",
            "addx 5",
            "addx 2",
            "addx -11",
            "addx 16",
            "addx -2",
            "addx -17",
            "addx 24",
            "noop",
            "noop",
            "addx 1",
            "addx -38",
            "addx 15",
            "addx 10",
            "addx -15",
            "noop",
            "addx 2",
            "addx 26",
            "noop",
            "addx -21",
            "addx 19",
            "addx -33",
            "addx 19",
            "noop",
            "addx -6",
            "addx 9",
            "addx 3",
            "addx 4",
            "addx -21",
            "addx 4",
            "addx 20",
            "noop",
            "addx 3",
            "addx -38",
            "addx 28",
            "addx -21",
            "addx 9",
            "addx -8",
            "addx 2",
            "addx 5",
            "addx 2",
            "addx -9",
            "addx 14",
            "addx -2",
            "addx -5",
            "addx 12",
            "addx 3",
            "addx -2",
            "addx 2",
            "addx 7",
            "noop",
            "noop",
            "addx -27",
            "addx 28",
            "addx -36",
            "noop",
            "addx 1",
            "addx 5",
            "addx -1",
            "noop",
            "addx 6",
            "addx -1",
            "addx 5",
            "addx 5",
            "noop",
            "noop",
            "addx -2",
            "addx 20",
            "addx -10",
            "addx -3",
            "addx 1",
            "addx 3",
            "addx 2",
            "addx 4",
            "addx 3",
            "noop",
            "addx -30",
            "noop",
        },
    }
    screens:= []string{
            "\t       0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39 \n\t" +
            "   1   R   #   #   L               #   L           R   #           R   #   #   L       R   #   L               #   L           R                   R           L      40 \n\t" +
            "  41   L                               R       R           R       R                   L           L       L           R       R                   R           R      80 \n\t" +
            "  81   R   R   #                       R       R                   R   #   L           L           R       R                   L                   R   #   L   L     120 \n\t" +
            " 121   L                               R       R                   R                   R   #   L           L       #   L       R                   L           L     160 \n\t" +
            " 161   L                   R           L       L           R       R                   L                   R           L       L                   R           L     200 \n\t" +
            " 201   R   #   L   L           R   L               R   L           R                   L                       #   L   L       #   L   #   L       R           R     240 \n\t" +
            "       0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39",
    }
    sums:= []int{
        11960,
    }
    registers:= []int{
        9,
    }
    for testIndex, screen := range screens {
        actualRegister, actualScreen, actualSum := RunInstructions(input[testIndex])
        assert.Equal(t, screen, actualScreen)
        assert.Equal(t, sums[testIndex], actualSum)
        assert.Equal(t, registers[testIndex], actualRegister)
    }
}