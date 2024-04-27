package request

import (
	"github.com/51w/server-go/server/model/common/request"
	"github.com/51w/server-go/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
