package dto

import (
    "github.com/baowk/dilu-core/core/base"
    "dilu/modules/sys/models"
)

type CounterGetPageReq struct {
	base.ReqPage `query:"-"`
    SortOrder  string `json:"-" query:"type:order;column:key"`
    
    
}

func (CounterGetPageReq) TableName() string {
	return models.TBCounter
}


//计数器
type CounterDto struct {
    
    Key string `json:"key"` //主键
    Val int `json:"val"` //值 
}