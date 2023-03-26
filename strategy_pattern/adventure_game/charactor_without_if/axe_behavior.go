package charactor_without_if

import "fmt"

type AxeBehavior struct {
}

func (a *AxeBehavior) UseWeapon() {
	fmt.Printf("Use a axe.\n")
}
