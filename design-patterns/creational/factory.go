package creational

import "fmt"

type PaymentGateway interface {
	Pay(amount float64) string
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

type Stripe struct{}

func (s *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Stripe", amount)
}

func PaymentFactory(gateway string) PaymentGateway {
	switch gateway {
	case "paypal":
		return &PayPal{}
	case "stripe":
		return &Stripe{}
	}
	return nil
}

func Factory() {
	var gateWays = []string{"paypal", "stripe"}
	for _, g := range gateWays {
		p := PaymentFactory(g)
		if p != nil {
			fmt.Println(p.Pay(100.50))
		} else {
			fmt.Println("Unknown payment gateway:", g)
		}
	}
}
