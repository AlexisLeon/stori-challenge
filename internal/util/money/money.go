package money

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"math"
)

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

func (m Money) Format() string {
	s := humanize.Commaf(math.Abs(float64(m)) / 100)
	if m < 0 {
		return fmt.Sprintf("-$%s", s)
	}

	return fmt.Sprintf("$%s", s)
}
