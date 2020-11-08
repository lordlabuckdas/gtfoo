package lolbas

import (
	"fmt"

	"github.com/lordlabuckdas/gtfoo/utils"
)

func Hello(name string) string {
	return fmt.Sprintf("%sHello, %v%s", utils.ColorRed, name, utils.ColorReset)
}
