package cmd

import (
	"fmt"
)

type Baker struct {
}

func (t Baker) BakeChicken() {
	fmt.Println("烤鸡翅")
}

func (t Baker) BakeSheep() {
	fmt.Println("烤羊肉串")
}
