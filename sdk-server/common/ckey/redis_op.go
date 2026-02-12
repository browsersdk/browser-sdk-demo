package ckey

import (
	"fmt"
)

func GetPhoneCodeKey(phone string) string {
	return "login:code:" + phone
}

func TeamMemberKey(teamId, userId int) string {
	return fmt.Sprintf("t:m:%d:%d", teamId, userId)
}

func LoginFeishuVerifyKey(state string) string {
	return fmt.Sprintf("login:feishu:verify:%s", state)
}

func DefendCodeStatusKey(username string) string {
	return fmt.Sprintf("defend:code:status:%s", username)
}

// func ChatSessionKey(sessionId string) string {
// 	return fmt.Sprintf("chat:session:%s", sessionId)
// }

// func UserChatKey(userId int) string {
// 	return fmt.Sprintf("user:chat:%d", userId)
// }

func AdminKey(id int) string {
	return fmt.Sprintf("admin:%d", id)
}

func RoleUserMenu(platform, roleId int) string {
	return fmt.Sprintf("role:menu:%d:%d", platform, roleId)
}

func RoleApis(platform, roleId int) string {
	return fmt.Sprintf("role:apis:%d:%d", platform, roleId)
}

func ApiKey(permType int) string {
	return fmt.Sprintf("sys:api:%d", permType)
}

func MemberNotice(teamId, uid int) string {
	return fmt.Sprintf("member:notice:%d:%d", teamId, uid)
}

func DeviceCategoryGroups(productId int, intentionCategoryId string) string {
	return fmt.Sprintf("device:category:groups:%d:%s", productId, intentionCategoryId)
}

func DeviceSessionLock(uid, sessionId string) string {
	return fmt.Sprintf("device:session:lock:%s:%s", uid, sessionId)
}

func ContinueMissIntentionCount(sessionId string) string {
	return "continue:miss:intention:count:" + sessionId
}

func UserLastCategory(uid, sessionId string) string {
	return fmt.Sprintf("user:last:category:%s:%s", uid, sessionId)
}

func UserLastCategoryAlpha(uid, sessionId string) string {
	return fmt.Sprintf("user:last:category:alpha:%s:%s", uid, sessionId)
}

func UserLastGroupId(uid, sessionId string) string {
	return fmt.Sprintf("user:last:group:id:%s:%s", uid, sessionId)
}

func UserLast2Category(uid, sessionId string) string {
	return fmt.Sprintf("user:last2:category:%s:%s", uid, sessionId)
}

func UserLast2CategoryAlpha(uid, sessionId string) string {
	return fmt.Sprintf("user:last2:category:alpha:%s:%s", uid, sessionId)
}

func UserLast2GroupId(uid, sessionId string) string {
	return fmt.Sprintf("user:last2:group:id:%s:%s", uid, sessionId)
}

func AppKey(appId int) string {
	return fmt.Sprintf("app:%d", appId)
}

func UserInfoKey(appId, uid int) string {
	return fmt.Sprintf("user:info:%d:%d", appId, uid)
}
