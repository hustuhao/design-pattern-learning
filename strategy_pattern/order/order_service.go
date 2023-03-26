package order

// 定义折扣Service，根据订单类型获取相对应的折扣。

type Service struct {
}

// Discount1 获取订单的折扣
// 新增加订单类型，则需要改写下面的代码。起始也还行
func (s *Service) Discount1(o *Order) float64 {
	if o.OrderType == ORDER_TYPE_NORMAL {
		return 1
	} else if o.OrderType == ORDER_TYPE_GROUPON {
		return 0.95
	} else if o.OrderType == ORDER_TYPE_PROMOTION {
		return 0.80
	} else {
		return 1
	}
}

// Discount2 获取订单的折扣
func (s *Service) Discount2(o *Order) float64 {
	strategy := dsFactory.GetStrategy(o.OrderType)
	return strategy.CalDiscount()
}
