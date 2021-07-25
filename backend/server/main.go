package main

import (
	"github.com/kkhan01/caputo/backend/goutils"
	metausecases "github.com/kkhan01/caputo/backend/server/usecases/meta"
)

// Version and BuildData get replaced during build with the commit hash and time
// of build.
var (
	CommitHash = ""
	BuildDate  = ""
)

func main() {
	goutils.PrettyPrint(metausecases.GetMeta(CommitHash, BuildDate))
}
