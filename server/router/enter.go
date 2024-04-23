package router

import (
	"github.com/51w/server-go/server/router/example"
	"github.com/51w/server-go/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
