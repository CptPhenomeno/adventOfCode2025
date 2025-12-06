package main

import "sort"

type Range struct {
	Start uint64
	End   uint64
}

func NewRange(start, end uint64) *Range {
	return &Range{Start: start, End: end}
}

func (r *Range) Contains(value uint64) bool {
	return value >= r.Start && value <= r.End
}

func (r *Range) Size() uint64 {
	return r.End - r.Start + 1
}

//=================================

type RangeSet struct {
	ranges []Range
}

func NewRangeSet() *RangeSet {
	return &RangeSet{ranges: make([]Range, 0)}
}

func (rs *RangeSet) Contains(value uint64) bool {
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

func (rs *RangeSet) Sort() {
	sort.Slice(rs.ranges, func(i, j int) bool {
		if rs.ranges[i].Start < rs.ranges[j].Start {
			return true
		} else if rs.ranges[i].Start == rs.ranges[j].Start {
			return rs.ranges[i].End < rs.ranges[j].End
		}

		return false
	})
}

func (rs *RangeSet) Flatten() *RangeSet {
	rs.Sort()
	flattenRangeSet := NewRangeSet()

	insertPending := true
	actualRange := NewRange(rs.ranges[0].Start, rs.ranges[0].End)
	for _, rg := range rs.ranges[1:] {
		if actualRange.Contains(rg.Start) {
			actualRange.End = max(actualRange.End, rg.End)
			insertPending = true
		} else {
			flattenRangeSet.Add(*actualRange)
			actualRange = NewRange(rg.Start, rg.End)
			insertPending = false
		}
	}
	if insertPending {
		flattenRangeSet.Add(*actualRange)
	}

	return flattenRangeSet
}

func (rs *RangeSet) Size() uint64 {
	var size uint64 = 0
	for _, r := range rs.ranges {
		size += r.Size()
	}
	return size
}

func CountFreshIngredients(ingredients []uint64, rangeSet *RangeSet) uint64 {
	var sum uint64 = 0
	for _, i := range ingredients {
		if rangeSet.Contains(i) {
			sum++
		}
	}
	return sum
}
