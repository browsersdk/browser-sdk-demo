package feishus

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func NoticeCard(webhook string, content string, imageKey string, headerContent string) {
	body := map[string]interface{}{
		"msg_type": "interactive",
		"card": map[string]interface{}{
			"config": map[string]bool{
				"wide_screen_mode": true,
			},
			"elements": []map[string]interface{}{
				{
					"alt": map[string]string{
						"content": "",
						"tag":     "plain_text",
					},
					"img_key": imageKey,
					"tag":     "img",
				},
				{
					"tag": "div",
					"text": map[string]string{
						"content": content,
						"tag":     "lark_md",
					},
				},
			},
			"header": map[string]interface{}{
				"template": "blue",
				"title": map[string]string{
					"content": headerContent,
					"tag":     "plain_text",
				},
			},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		slog.Error("failed to marshal data", "err", err)
		return
	}

	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("failed to post to Feishu", "err", err)
		return
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	slog.Debug("response from Feishu", "body", respBody)
}

func NoticeText(webhookURL, message string) {
	reqBody, err := json.Marshal(struct {
		MsgType string `json:"msg_type"`
		Content struct {
			Text string `json:"text"`
		} `json:"content"`
	}{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: message,
		},
	})
	if err != nil {
		slog.Error("failed to marshal request body", "err", err)
		return
	}
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		slog.Error("Failed to send message", "err", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	slog.Debug("response from Feishu", "body", respBody)
}
