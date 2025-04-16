package dog

import (
	"fmt"
	"strings"
)

func WhenGrownUp(s string) string {
	fmt.Println("This is dog version: 1.2.0")
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
