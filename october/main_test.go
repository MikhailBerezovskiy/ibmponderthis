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

func Test_gridConstraint(t *testing.T) {
	type args struct {
		t    int
		b    int
		r    int
		want int
	}
}

func Test_set_copy(t *testing.T) {
	s1 := newSet()
	s1.collinearX[1] = 1

	s2 := s1.copy()
	s2.collinearX[1] = 2
	s1.p = append(s1.p, point{1, 1})

	if s1.collinearX[1] != 1 {
		t.Errorf("want s1.colX = 1, got=%d", s1.collinearX[1])
	}

	if len(s2.p) != 0 {
		t.Errorf("want len s2.p = 0, got %d", len(s2.p))
	}
}

func Test_checkAreas(t *testing.T) {
	s := newSet()

	type args struct {
		name string
		d    []point
		want bool
	}

	tests := []args{
		args{"true test", []point{{1, 0}, {1, 1}, {0, 3}, {3, 3}}, true},
		args{"false test", []point{{0, 1}, {1, 1}, {1, 0}, {2, 0}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.p = tt.d
			got := s.checkAreas()
			if got != tt.want {
				t.Errorf("%v should be %v", tt.d, tt.want)
			}
		})
	}
}
