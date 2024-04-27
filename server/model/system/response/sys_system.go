package response

import "github.com/51w/server-go/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
