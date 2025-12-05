package main

type Range struct {
	Start int
	End   int
}

func NewRange(start, end int) *Range {
	return &Range{Start: start, End: end}
}

func (r *Range) Contains(value int) bool {
	return value >= r.Start && value <= r.End
}

//=================================

type RangeSet struct {
	ranges []Range
}

func NewRangeSet() *RangeSet {
	return &RangeSet{ranges: make([]Range, 0)}
}

func (rs *RangeSet) Contains(value int) bool {
	for _, r := range rs.ranges {
		if r.Contains(value) {
			return true
		}
	}
	return false
}

func (rs *RangeSet) Add(r Range) {
	rs.ranges = append(rs.ranges, r)
}

func CountFreshIngredients(ingredients []int, rangeSet *RangeSet) int {
	sum := 0
	for _, i := range ingredients {
		if rangeSet.Contains(i) {
			sum++
		}
	}
	return sum
}
