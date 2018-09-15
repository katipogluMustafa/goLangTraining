package math

// types

type IntPoint struct {
	X, Y int
}

// set implementation for small number of items
type IntPointSet struct {
	slice []IntPoint
}

// functions

func (p1 IntPoint) Equals(p2 IntPoint) bool {
	return (p1.X == p2.X) && (p1.Y == p2.Y)
}

func (set *IntPointSet) Add(p IntPoint) {
	if !set.Contains(p) {
		set.slice = append(set.slice, p)
	}
}

func (set IntPointSet) Contains(p IntPoint) bool {
	for _, v := range set.slice {
		if v.Equals(p) {
			return true
		}
	}
	return false
}

func (set IntPointSet) NumElements() int {
	return len(set.slice)
}

func NewIntPointSet() IntPointSet {
	return IntPointSet{(make([]IntPoint, 0, 10))}
}
