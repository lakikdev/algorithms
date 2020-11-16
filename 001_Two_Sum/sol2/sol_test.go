package sol2

import (
	"reflect"
	"testing"
)

func TestTwoNums(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{[]int{2, 7, 1, 15}, 9}, []int{0, 1}},
		{"test 2", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"test 3", args{[]int{3, 3}, 6}, []int{0, 1}},
		{"test 4 no answer", args{[]int{3, 2}, 6}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoNums(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoNums() = %v, want %v", got, tt.want)
			}
		})
	}
}





