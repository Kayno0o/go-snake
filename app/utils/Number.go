package utils

func Max[T int | int8 | int16 | int32 | int64](values ...T) T {
	var maxValue *T

	for _, value := range values {
		if maxValue == nil || value > *maxValue {
			maxValue = &value
		}
	}

	return *maxValue
}
