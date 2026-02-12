package enums

const (
	RolePlatformSuper = 1
	RoleTeamOwner     = 1
)

type PostType struct {
	Id   int
	Name string
	Key  string
}

var Admin = PostType{
	Id:   1,
	Name: "主管",
	Key:  "Admin",
}

var DeptHead = PostType{
	Id:   2,
	Name: "子管理员",
	Key:  "DeptHead",
}

// var DeptDeputy = PostType{
// 	Id:   4,
// 	Name: "主管",
// 	Key:  "Head",
// }

var Staff = PostType{
	Id:   4,
	Name: "职员",
	Key:  "Staff",
}
