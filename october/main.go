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

const (
	// GRID max grid size
	GRID int = 20
	// N max number of points
	N int = 4
)

var solutions = make([][]point, 0)

func main() {
	// RUN
	pts := gridConstaint(0, 0, 0)
	fmt.Printf("len pts:%d\n", len(pts))
	for i := 0; i < N; i++ {
		initSet := newSet()
		Search(initSet, point{i, 0})
	}
}

// Search function is a recursive search
func Search(s *set, pt point) {
	// create deep copy of set
	sc := s.copy()

	// add point to solution
	sc.p = append(sc.p, pt)
	sc.seen[pt] = true

	//fmt.Scanln()
	//fmt.Printf("N=%d pt=%v\n", len(sc.p), sc.p)

	// init slice of available points
	var pts []point

	// we ensure constraints satisfaction before calling search function
	// that means pt point already satisfied constraints
	// except areas

	// now search for next available points

	// CONSTRAINT#1 - grid
	// adjust current grid border values for Top, Bottom and Right with new point
	if pt.y > sc.gT {
		sc.gT = pt.y
	}
	if pt.y < sc.gB {
		sc.gB = pt.y
	}
	if pt.x > sc.gR {
		sc.gR = pt.y
	}
	//pts = gridConstaint(sc.gT, sc.gB, sc.gR) // test constraints and filter remain points
	pts = gridConstaint(0, 0, 0)

	// remove seen points from pts
	temp := make([]point, 0)
	for _, po := range pts {
		_, ok := sc.seen[po]
		if ok {
			continue
		}
		temp = append(temp, po)
	}
	pts = temp

	//fmt.Printf("grid = %+v\n", pts)

	// CONSTRAINT#2 - collinearity
	// adjust constraints map,
	sc.collinearX[pt.x]++
	sc.collinearY[pt.y]++
	pts = collinearityConstraint(sc.collinearX, sc.collinearY, pts) // test collinearity

	//fmt.Printf("coll = %+v\n", pts)

	// CONSTRAINT#3 - areas
	// check areas only when we have at least 3 points
	if len(sc.p) >= 3 {
		if !sc.checkAreas() {
			return
		}
	}

	// termination
	// if solution is found, len(s.p) == N
	if len(sc.p) == N {
		solutions = append(solutions, sc.p)
		fmt.Printf("new solution = %+v\n", sc.p)
		return
	}

	// if there available points continue search
	// otherwise deadend solution
	if len(pts) == 0 {
		return
	}
	for _, avPoint := range pts {
		Search(sc, avPoint)
	}
}

/*
gridConstraint(t, b, r, lim)
	res = []int
	for x = r to lim
		//above x-axis
		for y = t to lim
			if x*y < lim then
				res.append((x,y))
		// below x
		for y = b to -lim
			if x*y < -lim then
				res.append((x,y))
	return res
*/
func gridConstaint(t, b, r int) []point {
	res := make([]point, 0)
	for x := r; x <= GRID; x++ {
		for y := t; y <= GRID; y++ {
			if (x)*(y) <= GRID {
				res = append(res, point{x, y})
			}
		}
		for y := b; y >= -GRID; y-- {
			if (x)*(y) >= GRID {
				res = append(res, point{x, y})
			}
		}
	}
	return res
}

/*
collinearityConstraint(xConstr, yConstr, []points)
	filtered = []points
	for point in []points:
		if xConstr[point.x] == 2 || yConstr[point.y] == 2:
			continue
		filtered.append(point)
	return filtered
*/
func collinearityConstraint(xConstr, yConstr map[int]int, pts []point) []point {
	res := make([]point, 0)
	for _, pt := range pts {
		if xConstr[pt.x] == 2 || yConstr[pt.y] == 2 {
			continue
		}
		res = append(res, pt)
	}
	return res
}

// set is an snapshot of current search progress
type set struct {
	gT            int
	gB            int
	gR            int
	collinearX    map[int]int
	collinearY    map[int]int
	triangleAreas map[float64]int
	p             []point
	seen          map[point]bool
}

// x,y coordinates
type point struct {
	x int
	y int
}

// create new set of values
func newSet() *set {
	s := &set{
		collinearX:    make(map[int]int),
		collinearY:    make(map[int]int),
		triangleAreas: make(map[float64]int),
		p:             make([]point, 0),
		seen:          make(map[point]bool),
	}
	return s
}

// creates deep copy of current set
func (s *set) copy() *set {
	t := &set{
		gB:            s.gB,
		gR:            s.gR,
		gT:            s.gT,
		collinearX:    make(map[int]int),
		collinearY:    make(map[int]int),
		triangleAreas: make(map[float64]int),
		seen:          make(map[point]bool),
		p:             make([]point, 0),
	}

	for k, v := range s.collinearX {
		t.collinearX[k] = v
	}
	for k, v := range s.collinearY {
		t.collinearY[k] = v
	}
	for k, v := range s.triangleAreas {
		t.triangleAreas[k] = v
	}
	for k, v := range s.seen {
		t.seen[k] = v
	}

	t.p = append(t.p, s.p...)

	return t
}

// test collinearity constraint for new point
func (s *set) isCollinear(p point) bool {
	if s.collinearX[p.x] == 2 || s.collinearY[p.y] == 2 {
		return false
	}
	return true
}

// test areas constraint for new point
func (s *set) checkAreas() bool {
	combinations := getCombinations(s.p)
	for _, comb := range combinations {
		S := getArea(comb[0], comb[1], comb[2])
		_, ok := s.triangleAreas[S]
		if ok {
			return false
		}
		s.triangleAreas[S]++
	}
	return true
}

// return area of triangle by 3 points
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
