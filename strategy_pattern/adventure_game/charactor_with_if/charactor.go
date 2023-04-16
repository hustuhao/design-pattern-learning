package charactor_with_if

import (
	"fmt"
)

// 本例演示了一个角色使用不同的武器的攻击策略。
// 角色用不同的武器时，
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
