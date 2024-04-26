package money

import "testing"

func TestMoneyFromFloat64(t *testing.T) {
	tests := []struct {
		amount float64
		expect Money
	}{
		{
			amount: 12.34,
			expect: 1234,
		},
		{
			amount: 0.01,
			expect: 1,
		},
		{
			amount: 0.00,
			expect: 0,
		},
		{
			amount: 0.99,
			expect: 99,
		},
		{
			amount: 1.999,
			expect: 199,
		},
		{
			amount: 1.990001,
			expect: 199,
		},
	}

	for _, tt := range tests {
		if got := Float64(tt.amount); got != tt.expect {
			t.Errorf("Float64(%f) got %v, expected %v", tt.amount, got, tt.expect)
		}
	}
}
