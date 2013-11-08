package gohistogram

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
		bins:    make([]bin, 0),
		maxbins: n,
		total:   0,
	}
}

func (h *NumericHistogram) Add(n float64) {
	defer h.trim()
	h.total++
	for i := range h.bins {
		if h.bins[i].value == n {
			h.bins[i].count++
			return
		}

		if h.bins[i].value > n {

			newbin := bin{value: n, count: 1}
			head := append(make([]bin, 0), h.bins[0:i]...)

			head = append(head, newbin)
			tail := h.bins[i:]
			h.bins = append(head, tail...)
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

// trim merges adjacent bins to decrease the bin count to the maximum value
func (h *NumericHistogram) trim() {
	for len(h.bins) > h.maxbins {
		// Find closest bins in terms of value
		minDelta := 1e99
		minDeltaIndex := 0
		for i := range h.bins {
			if i == 0 {
				continue
			}

			if delta := h.bins[i].value - h.bins[i-1].value; delta < minDelta {
				minDelta = delta
				minDeltaIndex = i
			}
		}

		// We need to merge bins minDeltaIndex-1 and minDeltaIndex
		mergedbin := bin{
			value: (h.bins[minDeltaIndex-1].value + h.bins[minDeltaIndex].value) / 2, // average value
			count: h.bins[minDeltaIndex-1].count + h.bins[minDeltaIndex].count,       // summed heights
		}
		head := append(make([]bin, 0), h.bins[0:minDeltaIndex-1]...)
		tail := append([]bin{mergedbin}, h.bins[minDeltaIndex+1:]...)
		h.bins = append(head, tail...)
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
