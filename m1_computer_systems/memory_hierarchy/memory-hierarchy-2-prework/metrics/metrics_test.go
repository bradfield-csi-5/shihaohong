package metrics

import (
	"math"
	"testing"
)

func BenchmarkMetrics(b *testing.B) {
	users := LoadData()
	userCount := len(users)
	age := make([]float64, 0, len(users))
	payments := make([]Payment, 0, len(users))
	for _, u := range users {
		age = append(age, float64(u.age))
		payments = append(payments, u.payments...)
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
			actual = AveragePaymentAmount(payments)
		}

		expected := 499850.559
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected average payment amount to be around %.2f, not %.3f", expected, actual)
		}
	})

	b.Run("Payment stddev", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = StdDevPaymentAmount(users, payments, userCount)
		}
		expected := 288684.850
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected standard deviation to be around %.2f, not %.3f", expected, actual)
		}
	})

}
