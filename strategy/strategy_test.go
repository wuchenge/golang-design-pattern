package strategy

import "testing"

func TestPayCash(t *testing.T) {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
}

func TestPayBank(t *testing.T) {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
}
