package goutils

import (
	"encoding/json"
)

func PrettyFormat(e interface{}) []byte {
	output, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}

	return output
}
