package metrics

type Value interface {
	uint64 | int64 | float64
}
