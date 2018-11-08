package main

import (
	"testing"
)

func Test_getArea(t *testing.T) {
	type args struct {
		a point
		b point
		c point
	}

	p := []point{point{1, 0}, point{1, 1}, point{0, 3}, point{3, 3}}

	tests := []struct {
		name  string
		args  args
		wantS float64
	}{
		{"t1", args{p[0], p[1], p[2]}, 0.5},
		{"t2", args{p[0], p[1], p[3]}, 1},
		{"t3", args{p[0], p[2], p[3]}, 4.5},
		{"t3", args{p[1], p[2], p[3]}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := getArea(tt.args.a, tt.args.b, tt.args.c); gotS != tt.wantS {
				t.Errorf("getArea() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
