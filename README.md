gohistogram
=======
![histogram](http://i.imgur.com/5OplaRs.png)

The histograms in this package are based on the algorithms found in
Ben-Haim & Tom-Tov's *A Streaming Parallel Decision Tree Algorithm*
([PDF](http://jmlr.org/papers/volume11/ben-haim10a/ben-haim10a.pdf)).
Another implementation can be found in the Apache Hive project (see
NumericHistogram).

The accurate method of calculating quantiles (like percentiles) requires
data to be sorted. Streaming histograms make it possible to approximate
quantiles without sorting (or even individually storing) values.

`NumericHistogram` is the more basic implementation of a streaming
histogram. `WeightedHistogram` implements bin values as exponentially-weighted
moving averages.

A maximum bin size is passed as an argument to the constructor methods. A
larger bin size yields more accurate approximations at the cost of increased
memory utilization and performance.

### License
    Copyright (c) 2013 VividCortex

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.
