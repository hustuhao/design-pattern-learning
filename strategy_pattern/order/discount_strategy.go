package order

// 定义折扣策略和折扣策略工厂。

type DiscountStrategy interface {
	CalDiscount() float64
}

var dsFactory *DiscountStrategyFactory

type DiscountStrategyFactory struct {
	strategies map[int]DiscountStrategy // 策略
}

func (f *DiscountStrategyFactory) GetStrategy(orderType int) DiscountStrategy {
	return f.strategies[orderType]
}

// 具体的折扣策略

// NormalDiscountStrategy 普通订单折扣
type NormalDiscountStrategy struct{}

func (s *NormalDiscountStrategy) CalDiscount() float64 {
	return 1
}

// GrouponDiscountStrategy 拼团订单折扣
type GrouponDiscountStrategy struct{}

func (s *GrouponDiscountStrategy) CalDiscount() float64 {
	return 0.95
}

// PromotionDiscountStrategy  促销订单折扣
type PromotionDiscountStrategy struct{}

func (s *PromotionDiscountStrategy) CalDiscount() float64 {
	return 0.8
}

func NewDiscountStrategyFactory() *DiscountStrategyFactory {
	f := new(DiscountStrategyFactory)
	f.strategies = make(map[int]DiscountStrategy)
	//f.strategies = make(map[int]*DiscountStrategy)
	// var s1 *DiscountStrategy = new(NormalDiscountStrategy) // 编译不通过 var s1 DiscountStrategy 已经是指针了
	f.strategies[ORDER_TYPE_NORMAL] = new(NormalDiscountStrategy)
	f.strategies[ORDER_TYPE_GROUPON] = new(GrouponDiscountStrategy)
	f.strategies[ORDER_TYPE_PROMOTION] = new(PromotionDiscountStrategy)
	return f
}

func init() {
	dsFactory = NewDiscountStrategyFactory()
}
