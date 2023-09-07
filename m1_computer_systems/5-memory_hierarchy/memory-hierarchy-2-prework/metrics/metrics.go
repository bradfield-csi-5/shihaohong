package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type UserId int
type UserMap map[UserId]*User

type Address struct {
	fullAddress string
	zip         int
}

type Payment struct {
	cents uint32
	time  time.Time
}

type User struct {
	id       UserId
	name     string
	age      int
	address  Address
	payments []Payment
}

func AverageAge(ages []float64) float64 {
	count := len(ages)
	age1, age2, age3, age4 := 0.0, 0.0, 0.0, 0.0
	i := 0
	for ; i < count-3; i += 4 {
		age1 += ages[i]
		age2 += ages[i+1]
		age3 += ages[i+2]
		age4 += ages[i+3]
	}

	for ; i < count; i++ {
		age1 += ages[i]
	}

	return (age1 + age2 + age3 + age4) / float64(count)
}

func AveragePaymentAmount(amounts []uint32) float64 {
	count := len(amounts)
	amount1, amount2, amount3, amount4 := 0.0, 0.0, 0.0, 0.0
	i := 0
	for ; i < count-3; i += 4 {
		amount1 += float64(amounts[i])
		amount2 += float64(amounts[i+1])
		amount3 += float64(amounts[i+2])
		amount4 += float64(amounts[i+3])
	}
	for ; i < count; i++ {
		amount1 += float64(amounts[i])
	}

	return ((amount1 + amount2 + amount3 + amount4) * 0.01 / float64(count))
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(amounts []uint32) float64 {
	sumSquare, sum := float64(0), float64(0)
	sumSquare2, sum2 := float64(0), float64(0)
	count := len(amounts)
	i := 0
	for ; i < count-1; i += 2 {
		x := float64(amounts[i]) * 0.01
		sumSquare += x * x
		sum += x
		x2 := float64(amounts[i+1]) * 0.01
		sumSquare2 += x2 * x2
		sum2 += x2
	}

	avgSquare := (sumSquare + sumSquare2) / float64(count)
	avg := (sum + sum2) / float64(count)
	return math.Sqrt(avgSquare - avg*avg)
}

func LoadData() UserMap {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	users := make(UserMap, len(userLines))
	for _, line := range userLines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		age, _ := strconv.Atoi(line[2])
		address := line[3]
		zip, _ := strconv.Atoi(line[3])
		users[UserId(id)] = &User{UserId(id), name, age, Address{address, zip}, []Payment{}}
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	for _, line := range paymentLines {
		userId, _ := strconv.Atoi(line[2])
		paymentCents, _ := strconv.Atoi(line[0])
		datetime, _ := time.Parse(time.RFC3339, line[1])
		users[UserId(userId)].payments = append(users[UserId(userId)].payments, Payment{
			uint32(paymentCents),
			datetime,
		})
	}

	return users
}
