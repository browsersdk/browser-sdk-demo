package dto

type TeamReqIds struct {
	Ids    []int `json:"ids" form:"ids[]"`     //多id
	TeamId int   `json:"teamId" form:"teamId"` //团队id
}
