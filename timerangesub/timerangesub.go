package timerangesub

import (
	"sort"
)

func sortRanges(ranges [][]int) {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][1] > ranges[j][0] {
			return false
		}

		return true
	})
}

func TimeRangesSubtract(rangesSub, rangesMinuend [][]int) [][]int {
	result := [][]int{}

	sortRanges(rangesSub)
	sortRanges(rangesMinuend)

	subCount, minuCount := len(rangesSub), len(rangesMinuend)

	i, j := 0, 0
	for i < subCount && j < minuCount {
		s, m := rangesSub[i], rangesMinuend[j]
		//check if there is no overlap.
		if s[0] > m[1] || s[1] < m[0] {
			result = append(result, s)
			i++
			continue
		} else {
			if s[0] < m[0] {
				result = append(result, []int{s[0], m[0]})
				s[0] = m[0]
				continue
			}

			if s[1] > m[1] {
				s[0] = m[1]
				j++
			} else {
				i++
			}
		}
	}

	if i < subCount {
		result = append(result, rangesSub[i:]...)
	}

	return result
}
