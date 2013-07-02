package gohistogram

// Histogram is the interface that wraps the Add and Quantile methods.
//
// Add adds a new value, n, to the histogram. Trimming is done automatically.
//
// Quantile returns an approximation.
type Histogram interface {
	Add(n float64)
	Quantile(n float64) (q float64)
}

type bin struct {
	value float64
	count float64
}
