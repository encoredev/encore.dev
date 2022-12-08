package metrics

type Labels interface {
}

// CounterConfig configures a counter.
// It's currently a placeholder as there is not yet any additional configuration.
type CounterConfig struct {
	_ int // for godoc to show unexported fields
}

// NewCounter creates a new counter metric, without any labels.
// Use NewCounterGroup for metrics with labels.
func NewCounter[V Value](name string, cfg CounterConfig) *Counter[V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L23-L25
	panic("encore apps must be run using the encore command")
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
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L41-L45
	panic("encore apps must be run using the encore command")
}

// Add adds an arbitrary, non-negative value to the counter.
// It panics if the delta is < 0.
func (*Counter[V]) Add(delta V) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L49-L56
	panic("encore apps must be run using the encore command")
}

// NewCounterGroup creates a new counter group with a set of labels,
// where each unique combination of labels becomes its own counter.
//
// The Labels type must be a named struct, where each field corresponds to
// a single label. Each field must be of type string.
func NewCounterGroup[L Labels, V Value](name string, cfg CounterConfig) *CounterGroup[L, V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L63-L65
	panic("encore apps must be run using the encore command")
}

type CounterGroup[L Labels, V Value] struct {
	_ int // for godoc to show unexported fields
}

func (*CounterGroup[L, V]) With(labels L) *Counter[V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L78-L81
	panic("encore apps must be run using the encore command")
}

// GaugeConfig configures a gauge.
// It's currently a placeholder as there is not yet any additional configuration.
type GaugeConfig struct {
	_ int // for godoc to show unexported fields
}

// NewGauge creates a new counter metric, without any labels.
// Use NewGaugeGroup for metrics with labels.
func NewGauge[V Value](name string, cfg GaugeConfig) *Gauge[V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L103-L105
	panic("encore apps must be run using the encore command")
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
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L121-L125
	panic("encore apps must be run using the encore command")
}

func (*Gauge[V]) Add(val V) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L127-L131
	panic("encore apps must be run using the encore command")
}

// NewGaugeGroup creates a new gauge group with a set of labels,
// where each unique combination of labels becomes its own gauge.
//
// The Labels type must be a named struct, where each field corresponds to
// a single label. Each field must be of type string.
func NewGaugeGroup[L Labels, V Value](name string, cfg GaugeConfig) *GaugeGroup[L, V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L138-L140
	panic("encore apps must be run using the encore command")
}

type GaugeGroup[L Labels, V Value] struct {
	_ int // for godoc to show unexported fields
}

func (*GaugeGroup[L, V]) With(labels L) *Gauge[V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/metrics/metrics.go#L153-L156
	panic("encore apps must be run using the encore command")
}
