package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatches(t *testing.T) {
	packets := [][]string{
		{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
		},
		{
			"[1],[2,3,4]]",
			"[[1],4]",
		},
		{
			"[9]",
			"[[8,7,6]]",
		},
		{
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
		},
		{
			"[7,7,7,7]",
			"[7,7,7]",
		},
		{
			"[]",
			"[3]",
		},
		{
			"[[[]]]",
			"[[]]",
		},
		{
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		},
	}
	wants := []bool{
		true,
		true,
		false,
		true,
		false,
		true,
		false,
		false,
	}
	for testIndex, want := range wants {
        result := Matches(packets[testIndex][0], packets[testIndex][1])
		assert.Equal(t, want, result)
	}
}

//
//points :=[][]point{{},{}}
//for po, input := range []string{p,o}{
//    depth:=-1
//    breadth:=0
//    for charIndex:=0; charIndex < len(input);charIndex++{
//        inputEnd := charIndex+1
//        if input[charIndex] == '['{
//            depth++
//        }else if input[charIndex] == ']'{
//            depth--
//            breadth = 0
//        }else if input[charIndex] ==','{
//            breadth++
//        }else{
//            if input[inputEnd]!='[' && input[inputEnd]!=']' && input[inputEnd]!=','{
//                inputEnd++
//            }
//            points[po] = append(points[po], point{depth: depth,breadth: breadth,val: input[charIndex:inputEnd]})
//            if input[charIndex+1]!='[' && input[charIndex+1]!=']' && input[charIndex+1]!=','{
//                charIndex++
//            }
//        }
//    }
//}