package charactor_without_if

import "fmt"

type KnifeBehavior struct {
}

func (k *KnifeBehavior) UseWeapon() {
	fmt.Printf("Use a Knife.\n")
}
