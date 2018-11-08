package main

import (
	"fmt"
	"math"
)

/*
Put N points on integer coordinates of a rectangular grid
of dimension L1xL2, so that no three points are collinear
and the areas of the triangles formed by the C(N,3)
possible triplets are all different.

For example, for N=4 and L1=L2=3,
the following set of points is a solution: [[1,0],[1,1],[0,3],[3,3]]

The areas formed by the 4 triangles are:

[1, 0] [1, 1] [0, 3] 0.5
[1, 0] [1, 1] [3, 3] 1
[1, 0] [0, 3] [3, 3] 4.5
[1, 1] [0, 3] [3, 3] 3
Find a solution for N=11 and the L1*L2<=600.
*/

func main() {

	s := newSet()

	for i := 0; i < 20; i++ {
		for j := 0; j < 30; j++ {
			s.addOne(point{i, j})
			if s.done {
				fmt.Printf("points: %v\n", s.p)
				fmt.Printf("triangles: %v\n", s.triangleAreas)
				return
			}
		}
	}
}

type set struct {
	gridX         int
	gridY         int
	collinearX    map[int]int
	collinearY    map[int]int
	triangleAreas map[float64]int
	p             []point
	i             int
	area          float64
	done          bool
}

type point struct {
	x int
	y int
}

func newSet() *set {
	s := &set{
		collinearX:    make(map[int]int),
		collinearY:    make(map[int]int),
		triangleAreas: make(map[float64]int),
		p:             make([]point, 0),
	}
	return s
}

func (s *set) addOne(p point) {
	if !s.isCollinear(p) {
		return
	}
	if s.i >= 2 {
		if !s.checkAreas(p) {
			return
		}
	}
	s.collinearX[p.x]++
	s.collinearY[p.y]++
	s.p = append(s.p, p)
	s.i++

	if s.i == 5 {
		s.done = true
	}
}

func (s *set) isCollinear(p point) bool {
	if s.collinearX[p.x] == 2 || s.collinearY[p.y] == 2 {
		return false
	}
	return true
}

func (s *set) checkAreas(p point) bool {
	points := make([]point, len(s.p)+1)
	copy(points, s.p)
	points[len(points)-1] = p
	areas := make(map[float64]int)
	for k, v := range s.triangleAreas {
		areas[k] = v
	}
	combinations := getCombinations(points)
	for _, comb := range combinations {
		S := getArea(comb[0], comb[1], comb[2])
		_, ok := areas[S]
		if ok {
			return false
		}
		areas[S]++
	}
	s.triangleAreas = areas
	return true
}

func getArea(a, b, c point) (s float64) {
	s = math.Abs(float64(a.x*(b.y-c.y)-b.x*(a.y-c.y)+c.x*(a.y-b.y))) / 2
	return
}

// get all posible combinations of set.points by 3
func getCombinations(input []point) [][]point {
	var result = make([][]point, 0)
	var recf func([]point, int, int)
	recf = func(comb []point, low, high int) {
		if high > len(input) {
			result = append(result, comb)
			return
		}
		for i := low; i < high; i++ {
			combCopy := make([]point, len(comb))
			copy(combCopy, comb)
			combCopy = append(combCopy, input[i])
			recf(combCopy, i+1, high+1)
		}
	}

	recf(make([]point, 0), 0, len(input)-2)
	return result
}
