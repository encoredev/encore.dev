package metrics

import "os"

type Labels interface {
	comparable
}

// CounterConfig configures a counter.
// It's currently a placeholder as there is not yet any additional configuration.
type CounterConfig struct {
	_ int // for godoc to show unexported fields
}

type Counter[V Value] struct {
	_ int // for godoc to show unexported fields
}

// Increment increments the counter by 1.
func (*Counter[V]) Increment() {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.0/runtimes/go/metrics/metrics.go#L36-L41
	doPanic("encore apps must be run using the encore command")
	return
}

// Add adds an arbitrary, non-negative value to the counter.
// It panics if the delta is < 0.
func (*Counter[V]) Add(delta V) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.0/runtimes/go/metrics/metrics.go#L45-L53
	doPanic("encore apps must be run using the encore command")
	return
}

type CounterGroup[L Labels, V Value] struct {
	_ int // for godoc to show unexported fields
}

func (*CounterGroup[L, V]) With(labels L) (_ *Counter[V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.0/runtimes/go/metrics/metrics.go#L71-L74
	doPanic("encore apps must be run using the encore command")
	return
}

// GaugeConfig configures a gauge.
// It's currently a placeholder as there is not yet any additional configuration.
type GaugeConfig struct {
	_ int // for godoc to show unexported fields
}

type Gauge[V Value] struct {
	_ int // for godoc to show unexported fields
}

func (*Gauge[V]) Set(val V) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.0/runtimes/go/metrics/metrics.go#L111-L116
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Gauge[V]) Add(val V) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.0/runtimes/go/metrics/metrics.go#L118-L123
	doPanic("encore apps must be run using the encore command")
	return
}

type GaugeGroup[L Labels, V Value] struct {
	_ int // for godoc to show unexported fields
}

func (*GaugeGroup[L, V]) With(labels L) (_ *Gauge[V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.0/runtimes/go/metrics/metrics.go#L136-L139
	doPanic("encore apps must be run using the encore command")
	return
}

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking Encore APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("ENCORERUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
