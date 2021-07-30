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

func MebibyteToBytes(mb int64) int64 {
	return mb << 20
}
