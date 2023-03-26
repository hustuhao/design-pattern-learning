package charactor_without_if

type Character struct {
	Weapon WeaponBehavior
}

func (c *Character) Fight() {
	c.Weapon.UseWeapon()
}

func (c *Character) SetWeaponBehavior(w WeaponBehavior) {
	c.Weapon = w
}
