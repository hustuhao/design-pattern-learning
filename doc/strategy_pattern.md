# 策略模式

代码目录：strategy_pattern

策略模式定义了算法族，分别封装起来，让他们之间可以相互替换。此模式让算法的变化独立于使用算法的客户端。其解耦的是策略的定义、创建和使用三部分。

```
├── adventure_game                    # 冒险游戏: 角色和装备切换、攻击。
│ ├── charactor_with_if               # 使用 IF 语句来实现「使用装备攻击」
│ │ ├── charactor.go                   
│ │ └── charactor_test.go
│ └── charactor_without_if            # 使用策略模式实现「使用装备攻击」
│     ├── axe_behavior.go
│     ├── character.go
│     ├── charactor_test.go
│     ├── knife_behavior.go
│     └── weapon_behavior.go
├── config
└── order                             # 使用策略模式，实现订单折扣
    ├── discount_strategy.go          # 折扣策略实现
    ├── discount_strategy_test.go     # 折扣策略
    ├── order.go                      # 订单
    └── order_service.go              # 订单服务
```