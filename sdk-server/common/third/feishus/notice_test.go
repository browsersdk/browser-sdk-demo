package feishus

import "testing"

const webhook = "https://open.feishu.cn/open-apis/bot/v2/hook/77b95b3d-ff1e-4a2d-af4c-7bd1733bb0fc"

func TestNoticeCard(t *testing.T) {
	NoticeCard(webhook, "test", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "test")
}

func TestNoticeText(t *testing.T) {
	NoticeText(webhook, "test")
}
