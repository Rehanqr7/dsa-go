package stacks

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	cases := []struct{
		in []int
		exp []int
	}{
		{[]int{2,1,2,4,3}, []int{4,2,4,-1,-1}},
		{[]int{1,3,2,4}, []int{3,4,4,-1}},
	}
	for _, c := range cases {
		if got := NextGreaterElements(c.in); !reflect.DeepEqual(got, c.exp) {
			t.Fatalf("NextGreaterElements(%v)=%v want %v", c.in, got, c.exp)
		}
	}
}
