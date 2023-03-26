package charactor_with_if

import "testing"

// 你作为一个游戏主角，可以自由切换武器进行攻击
func TestUseWeapon1(t *testing.T) {
	c := new(Character)
	c.Weapon = "小刀"
	c.UseWeapon()

	c.Weapon = "斧头"
	c.UseWeapon()

}
