package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperate(t *testing.T) {
	input := []string{
		"old + 5",
		"old * 5",
		"old - 5",
		"old / 5",
	}
	olds := []uint64{
		5,
		5,
		5,
		5,
	}
	wants := []uint64{
		10,
		25,
		0,
		1,
	}
	for testIndex, want := range wants {
		result := Operate(input[testIndex], olds[testIndex])
		assert.Equal(t, want, result)
	}
}

func TestIsDivisibleBy(t *testing.T) {
	input := []uint64{
		3,
		5,
		3,
		5,
		7,
	}
	olds := []uint64{
		5,
		5,
		10,
		10,
		7,
	}
	wants := []bool{
		false,
		true,
		false,
		true,
		true,
	}
	for testIndex, want := range wants {
		result := IsDivisibleBy(input[testIndex], olds[testIndex])
		assert.Equal(t, want, result)
	}
}

func TestInspect(t *testing.T) {
	input := []string{
		"old + 5",
		"old * 5",
		"old - 5",
		"old / 5",
	}
	olds := []uint64{
		5,
		5,
		5,
		5,
	}
	wants := []uint64{
		10,
		25,
		0,
		1,
	}
	for testIndex, want := range wants {
		result := Inspect(input[testIndex], olds[testIndex])
		assert.Equal(t, want, result)
	}
}

func TestRelief(t *testing.T) {
	input := []uint64{
		3,
		5,
		3,
		5,
		7,
		18,
		19,
	}
	wants := []uint64{
		1,
		1,
		1,
		1,
		2,
		6,
		6,
	}
	for testIndex, want := range wants {
		result := Relief(input[testIndex])
		assert.Equal(t, want, result)
	}
}

func TestTest(t *testing.T) {
	input := []string{
		"3",
		"5",
		"3",
		"5",
		"7",
	}
	olds := []uint64{
		5,
		5,
		10,
		10,
		7,
	}
	wants := []bool{
		false,
		true,
		false,
		true,
		true,
	}
	for testIndex, want := range wants {
		result := Test(input[testIndex], olds[testIndex])
		assert.Equal(t, want, result)
	}
}

func TestRound(t *testing.T) {
	input := []string{
		"Monkey 0:\n" +
			"  Starting items: 79, 98\n" +
			"  Operation: new = old * 19\n" +
			"  Test: divisible by 23\n" +
			"    If true: throw to monkey 2\n" +
			"    If false: throw to monkey 3\n\n" +

			"Monkey 1:\n" +
			"  Starting items: 54, 65, 75, 74\n" +
			"  Operation: new = old + 6\n" +
			"  Test: divisible by 19\n" +
			"    If true: throw to monkey 2\n" +
			"    If false: throw to monkey 0\n\n" +

			"Monkey 2:\n" +
			"  Starting items: 79, 60, 97\n" +
			"  Operation: new = old * old\n" +
			"  Test: divisible by 13\n" +
			"    If true: throw to monkey 1\n" +
			"    If false: throw to monkey 3\n\n" +

			"Monkey 3:\n" +
			"  Starting items: 74\n" +
			"  Operation: new = old + 3\n" +
			"  Test: divisible by 17\n" +
			"    If true: throw to monkey 0\n" +
			"    If false: throw to monkey 1\n",
		"Monkey 0:\n" +
			"  Starting items: 79, 98\n" +
			"  Operation: new = old * 19\n" +
			"  Test: divisible by 23\n" +
			"    If true: throw to monkey 2\n" +
			"    If false: throw to monkey 3\n\n" +

			"Monkey 1:\n" +
			"  Starting items: 54, 65, 75, 74\n" +
			"  Operation: new = old + 6\n" +
			"  Test: divisible by 19\n" +
			"    If true: throw to monkey 2\n" +
			"    If false: throw to monkey 0\n\n" +

			"Monkey 2:\n" +
			"  Starting items: 79, 60, 97\n" +
			"  Operation: new = old * old\n" +
			"  Test: divisible by 13\n" +
			"    If true: throw to monkey 1\n" +
			"    If false: throw to monkey 3\n\n" +

			"Monkey 3:\n" +
			"  Starting items: 74\n" +
			"  Operation: new = old + 3\n" +
			"  Test: divisible by 17\n" +
			"    If true: throw to monkey 0\n" +
			"    If false: throw to monkey 1\n",
	}
	reliefs := []bool{
		true,
		false,
	}
	wants := []string{
		"Monkey 0: 20, 23, 27, 26\n" +
			"Monkey 1: 2080, 25, 167, 207, 401, 1046\n" +
			"Monkey 2:\n" +
			"Monkey 3:\n",
		"Monkey 0: 60, 71, 81, 80\n" +
			"Monkey 1: 77, 1504, 1865, 6244, 3603, 9412\n" +
			"Monkey 2:\n" +
			"Monkey 3:\n",
	}
	for testIndex, want := range wants {
		result := FormatMonkeys(Round(Load(input[testIndex]), reliefs[testIndex]))
		assert.Equal(t, want, result)
	}
}
