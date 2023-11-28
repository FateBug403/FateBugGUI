package global

import (
	"github.com/adjust/rmq/v4"
	"gorm.io/gorm"
)

var (
	GVA_RmqC rmq.Connection
	GVA_DB     *gorm.DB

)