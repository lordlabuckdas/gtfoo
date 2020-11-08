package gtfobins

import (
	"fmt"

	"github.com/lordlabuckdas/gtfoo/utils"
)

func Greet(name string) string {
	return fmt.Sprintf("%sHi, %v%s", utils.ColorBlue, name, utils.ColorReset)
}
