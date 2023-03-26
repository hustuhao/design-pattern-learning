package order

const (
	// 订单类型
	ORDER_TYPE_NORMAL    = 1 // 普通订单
	ORDER_TYPE_GROUPON   = 2 // 拼团订单
	ORDER_TYPE_PROMOTION = 3 // 促销订单
)

type Order struct {
	ID        int64 // 订单id
	OrderType int   // 订单类型

	// ......
}
