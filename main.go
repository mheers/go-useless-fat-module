package useless

import (
	_ "embed"
	"fmt"
)

//go:embed 10mbRandom.dump
var randomBytes []byte

func Useless() {
	fmt.Printf("do not use me, i just have %d bytes additionally embedded \n", len(randomBytes))
}
