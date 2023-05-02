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

// NewCounter creates a new counter metric, without any labels.
// Use NewCounterGroup for metrics with labels.
func NewCounter[V Value](name string, cfg CounterConfig) (_ *Counter[V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L24-L26
	doPanic("encore apps must be run using the encore command")
	return
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
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L42-L47
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
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L51-L59
	doPanic("encore apps must be run using the encore command")
	return
}

// NewCounterGroup creates a new counter group with a set of labels,
// where each unique combination of labels becomes its own counter.
//
// The Labels type must be a named struct, where each field corresponds to
// a single label. Each field must be of type string.
func NewCounterGroup[L Labels, V Value](name string, cfg CounterConfig) (_ *CounterGroup[L, V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L66-L68
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
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L86-L89
	doPanic("encore apps must be run using the encore command")
	return
}

// GaugeConfig configures a gauge.
// It's currently a placeholder as there is not yet any additional configuration.
type GaugeConfig struct {
	_ int // for godoc to show unexported fields
}

// NewGauge creates a new counter metric, without any labels.
// Use NewGaugeGroup for metrics with labels.
func NewGauge[V Value](name string, cfg GaugeConfig) (_ *Gauge[V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L111-L113
	doPanic("encore apps must be run using the encore command")
	return
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
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L129-L134
	doPanic("encore apps must be run using the encore command")
	return
}

func (*Gauge[V]) Add(val V) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L136-L141
	doPanic("encore apps must be run using the encore command")
	return
}

// NewGaugeGroup creates a new gauge group with a set of labels,
// where each unique combination of labels becomes its own gauge.
//
// The Labels type must be a named struct, where each field corresponds to
// a single label. Each field must be of type string.
func NewGaugeGroup[L Labels, V Value](name string, cfg GaugeConfig) (_ *GaugeGroup[L, V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L148-L150
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
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/metrics/metrics.go#L163-L166
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
