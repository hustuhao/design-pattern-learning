package order

import "testing"

func TestDiscount(t *testing.T) {
	s := new(Service)
	order := new(Order)
	order.OrderType = 1
	d1 := s.Discount1(order)
	d2 := s.Discount2(order)
	t.Logf("d1=%f,d2=%f\n", d1, d2)
	order.OrderType = 2
	d3 := s.Discount1(order)
	d4 := s.Discount2(order)
	t.Logf("d1=%f,d2=%f\n", d3, d4)
	order.OrderType = 3
	d5 := s.Discount1(order)
	d6 := s.Discount2(order)
	t.Logf("d1=%f,d2=%f\n", d5, d6)

}
