package array

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []*Interval) []*Interval {
	// write code here
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var stack []*Interval
	for _, v := range intervals {
		if len(stack) == 0 {
			stack = append(stack, v)
		} else if stack[len(stack)-1].End >= v.Start && stack[len(stack)-1].End < v.End {
			stack[len(stack)-1].End = v.End
		} else if stack[len(stack)-1].End >= v.End {
			continue
		} else {
			stack = append(stack, v)
		}
	}
	return stack
}
