package v1

import (
	"github.com/51w/server-go/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
