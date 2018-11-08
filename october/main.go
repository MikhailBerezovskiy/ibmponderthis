package main

import "fmt"

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

	for x := 0; x < 600; x++ {
		for y := 0; y < 600; y++ {
			if s.done {
				fmt.Printf("Points: %v\n", s.p)
				fmt.Printf("Triangle areas: %v\n", s.triangleAreas)
				fmt.Printf("Total area: %f\n", s.area)
				return
			}
			s.addOne(point{x, y})
		}
	}

}

type set struct {
	gridX         int
	gridY         int
	collinearX    map[int]int
	collinearY    map[int]int
	triangleAreas map[float32]int
	p             [11]point
	i             int
	area          float32
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
		triangleAreas: make(map[float32]int),
	}
	return s
}

func (s *set) addOne(p point) {
	if !s.checkCollinear(p) {
		return
	}
	s.collinearX[p.x]++
	s.collinearY[p.y]++
	s.p[s.i] = p
	s.i++

	if s.i == 11 {
		s.done = true
	}
}

func (s *set) checkCollinear(p point) bool {
	if s.collinearX[p.x] == 2 || s.collinearY[p.y] == 2 {
		return false
	}
	return true
}

// calculate areas
func getAreasTest() map[float32]bool {
	areas := make(map[float32]bool)

	p := []point{
		point{1, 0},
		point{1, 1},
		point{0, 3},
		point{3, 3},
	}

	fmt.Printf("p = %+v\n", p)

	return areas
}
