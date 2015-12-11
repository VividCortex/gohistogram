package gohistogram

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

import (
	"fmt"
)

type NumericHistogram struct {
	bins    []bin
	maxbins int
	total   uint64
}

// NewHistogram returns a new NumericHistogram with a maximum of n bins.
//
// There is no "optimal" bin count, but somewhere between 20 and 80 bins
// should be sufficient.
func NewHistogram(n int) *NumericHistogram {
	return &NumericHistogram{
		bins:    make([]bin, 0, n+1),
		maxbins: n,
		total:   0,
	}
}

func (h *NumericHistogram) Add(n float64) {
	defer h.trim()
	h.total++
	for i, bini := range h.bins {
		if bini.value == n {
			h.bins[i].count++
			return
		}

		if bini.value > n {

			newbin := bin{value: n, count: 1}
			tail := h.bins[i:]
			h.bins = h.bins[0 : len(h.bins)+1]
			copy(h.bins[i+1:], tail)
			h.bins[i] = newbin

			return
		}
	}

	h.bins = append(h.bins, bin{count: 1, value: n})
}

func (h *NumericHistogram) Quantile(q float64) float64 {
	count := q * float64(h.total)
	for i := range h.bins {
		count -= float64(h.bins[i].count)

		if count <= 0 {
			return h.bins[i].value
		}
	}

	return -1
}

// CDF returns the value of the cumulative distribution function
// at x
func (h *NumericHistogram) CDF(x float64) float64 {
	count := 0.0
	for i := range h.bins {
		if h.bins[i].value <= x {
			count += float64(h.bins[i].count)
		}
	}

	return count / float64(h.total)
}

// Mean returns the sample mean of the distribution
func (h *NumericHistogram) Mean() float64 {
	if h.total == 0 {
		return 0
	}

	sum := 0.0

	for i := range h.bins {
		sum += h.bins[i].value * h.bins[i].count
	}

	return sum / float64(h.total)
}

// Variance returns the variance of the distribution
func (h *NumericHistogram) Variance() float64 {
	if h.total == 0 {
		return 0
	}

	sum := 0.0
	mean := h.Mean()

	for i := range h.bins {
		sum += (h.bins[i].count * (h.bins[i].value - mean) * (h.bins[i].value - mean))
	}

	return sum / float64(h.total)
}

func (h *NumericHistogram) Count() float64 {
	return float64(h.total)
}

// trim merges adjacent bins to decrease the bin count to the maximum value
func (h *NumericHistogram) trim() {
	for len(h.bins) > h.maxbins {
		// Find closest bins in terms of value
		minDelta := 1e99
		minDeltaIndex := 0
		for i, bini := range h.bins {
			if i == 0 {
				continue
			}

			if delta := bini.value - h.bins[i-1].value; delta < minDelta {
				minDelta = delta
				minDeltaIndex = i
			}
		}

		binMinDeltaIdxSub1 := h.bins[minDeltaIndex-1]
		binMinDeltaIdx := h.bins[minDeltaIndex]

		// We need to merge bins minDeltaIndex-1 and minDeltaIndex
		totalCount := binMinDeltaIdxSub1.count + binMinDeltaIdx.count
		mergedbin := bin{
			value: (binMinDeltaIdxSub1.value*
				binMinDeltaIdxSub1.count +
				binMinDeltaIdx.value*
					binMinDeltaIdx.count) /
				totalCount, // weighted average
			count: totalCount, // summed heights
		}
		head := h.bins[0:minDeltaIndex]
		tail := h.bins[minDeltaIndex+1:]
		head[minDeltaIndex-1] = mergedbin
		copy(h.bins[minDeltaIndex:], tail)
		h.bins = h.bins[:h.maxbins]
	}
}

// String returns a string reprentation of the histogram,
// which is useful for printing to a terminal.
func (h *NumericHistogram) String() (str string) {
	str += fmt.Sprintln("Total:", h.total)

	for i := range h.bins {
		var bar string
		for j := 0; j < int(float64(h.bins[i].count)/float64(h.total)*200); j++ {
			bar += "."
		}
		str += fmt.Sprintln(h.bins[i].value, "\t", bar)
	}

	return
}
