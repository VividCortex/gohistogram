package gohistogram

// Histogram is the interface that wraps the Add and Quantile methods.
type Histogram interface {
	// Add adds a new value, n, to the histogram. Trimming is done
	// automatically.
	Add(n float64)

	// Quantile returns an approximation.
	Quantile(n float64) (q float64)
}

type bin struct {
	value float64
	count float64
}
