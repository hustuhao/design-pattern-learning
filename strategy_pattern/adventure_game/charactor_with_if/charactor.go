package charactor_with_if

import (
	"fmt"
)

// CASE 1、2：游戏角色
// CASE 3 :订单折扣
type Character struct {
	Weapon string
}

func (c *Character) SetWeapon(w string) {
	c.Weapon = w
}

func (c *Character) UseWeapon() {
	if c.Weapon == "小刀" {
		fmt.Printf("使用小刀攻击\n")
		return
	}

	if c.Weapon == "斧头" {
		fmt.Printf("使用斧头攻击\n")
		return
	}
}
