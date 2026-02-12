package models

// 计数器
type Counter struct {
	Key string `json:"key" gorm:"type:varchar(100);primaryKey;autoIncrement;comment:主键"` //主键
	Seq int64  `json:"seq" gorm:"type:bigint unsigned;comment:值"`                        //值
}

const TBCounter = "counter"

func (Counter) TableName() string {
	return TBCounter
}

const (
	CounterKeyUserNo = "user_no"
	MinUserNo        = 10000000
)
