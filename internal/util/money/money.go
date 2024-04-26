package money

type Money int64

func Float64(amount float64) Money {
	// converting to float32 to avoid precision loss when converting to int64
	return Money(float32(amount) * 100)
}

func (m Money) Int64() int64 {
	return int64(m)
}

func (m Money) Int() int {
	return int(m)
}
