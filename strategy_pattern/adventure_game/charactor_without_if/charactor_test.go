package charactor_without_if

import "testing"

// 你作为一个游戏主角，可以自由切换武器进行攻击
func TestUseWeapon1(t *testing.T) {
	c := new(Character)
	// 使用小刀
	knife := new(KnifeBehavior)
	c.SetWeaponBehavior(knife)
	c.Fight()

	// 切换到斧头
	axe := new(AxeBehavior)
	c.SetWeaponBehavior(axe)
	c.Fight()

}
