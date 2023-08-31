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

type DollarAmount struct {
	dollars, cents uint64
}

type Payment struct {
	amount DollarAmount
	time   time.Time
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
	for ; i < count-4; i += 4 {
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

func AveragePaymentAmount(amount []DollarAmount) float64 {
	count := len(amount)
	amount1, amount2, amount3, amount4 := 0.0, 0.0, 0.0, 0.0
	i := 0
	for ; i < count-4; i += 4 {
		amount1 += float64(amount[i].dollars) + float64(amount[i].cents)/100
		amount2 += float64(amount[i+1].dollars) + float64(amount[i+1].cents)/100
		amount3 += float64(amount[i+2].dollars) + float64(amount[i+2].cents)/100
		amount4 += float64(amount[i+3].dollars) + float64(amount[i+3].cents)/100
	}

	for ; i < count; i++ {
		amount1 += float64(amount[i].dollars) + float64(amount[i].cents)/100
	}

	return (amount1 + amount2 + amount3 + amount4) / float64(count)
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(users UserMap, amount []DollarAmount, userCount int) float64 {
	mean := AveragePaymentAmount(amount)
	squaredDiffs, count := 0.0, 0.0
	for _, u := range users {
		for _, p := range u.payments {
			count += 1
			amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
			diff := amount - mean
			squaredDiffs += diff * diff
		}
	}
	return math.Sqrt(squaredDiffs / count)
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
			DollarAmount{uint64(paymentCents / 100), uint64(paymentCents % 100)},
			datetime,
		})
	}

	return users
}
