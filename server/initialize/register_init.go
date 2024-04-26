package initialize

import (
	"fmt"
	// _ "github.com/51w/server-go/server/source/example"
	// _ "github.com/51w/server-go/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
	fmt.Println("initialize register_init init.")
}
