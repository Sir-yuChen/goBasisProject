package InitLib1

import (
	_ "../InitLib2"
	"fmt"
)

func init() {
	fmt.Println("IntLib1=====>init")
}
