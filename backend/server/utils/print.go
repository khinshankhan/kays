package utils

import (
	"fmt"
)

func PrettyPrint(e interface{}) {
	output := PrettyFormat(e)
	fmt.Printf("%v\n", string(output))
}
