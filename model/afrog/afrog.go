// 自动生成模板Afrog
package afrog

// Afrog 结构体
type Afrog struct {
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
}


// TableName Afrog 表名
func (Afrog) TableName() string {
  return "afrog"
}

