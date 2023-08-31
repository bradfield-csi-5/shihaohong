package metrics

import (
	"math"
	"testing"
)

func BenchmarkMetrics(b *testing.B) {
	users := LoadData()
	userCount := len(users)
	age := make([]float64, 0, len(users))
	dollarAmounts := make([]DollarAmount, 0)
	for _, u := range users {
		age = append(age, float64(u.age))
		for _, payments := range u.payments {
			dollarAmounts = append(dollarAmounts, payments.amount)
		}
	}

	b.Run("Average age", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = AverageAge(age)
		}
		expected := 59.62
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected average age to be around %.2f, not %.3f", expected, actual)
		}
	})

	b.Run("Average payment", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = AveragePaymentAmount(dollarAmounts)
		}

		expected := 499850.559
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected average payment amount to be around %.2f, not %.3f", expected, actual)
		}
	})

	b.Run("Payment stddev", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = StdDevPaymentAmount(users, dollarAmounts, userCount)
		}
		expected := 288684.850
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected standard deviation to be around %.2f, not %.3f", expected, actual)
		}
	})

}
