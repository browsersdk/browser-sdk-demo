package feishus

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkauthen "github.com/larksuite/oapi-sdk-go/v3/service/authen/v1"
	"golang.org/x/oauth2"
)

type Client struct {
	client    *lark.Client
	appId     string
	appSecret string
}

func NewClient(appId, appSecret string) *Client {
	client := lark.NewClient(appId, appSecret,
		//lark.WithMarketplaceApp(),// 设置为商店应用
		lark.WithLogLevel(larkcore.LogLevelDebug),
		lark.WithEnableTokenCache(true),
		//lark.WithHelpdeskCredential("id", "token"),
	) // 默认配置为自建应用

	return &Client{
		client:    client,
		appId:     appId,
		appSecret: appSecret,
	}

}

// 自建应用获取访问token
func (cli *Client) GetTenantAccessToken() {
	// 发起 API 请求
	resp, err := cli.client.GetTenantAccessTokenBySelfBuiltApp(context.Background(), &larkcore.SelfBuiltTenantAccessTokenReq{
		AppID:     cli.appId,
		AppSecret: cli.appSecret,
	})
	// 错误处理
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取请求 ID
	fmt.Println(resp.LogId())
	// 处理请求结果
	fmt.Println(resp.StatusCode) // http status code
	fmt.Println(resp.TenantAccessToken, resp.Expire)
}

// func (cli *Client) GetUserAccessToken(code string) {
// 	resp, err := cli.client.Authen.AccessToken.GetUserAccessToken(context.Background(), &larkcore.UserAccessTokenReq{
// 		UserAccessToken: userAccessToken,
// 	})
// }

var oauthEndpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.feishu.cn/open-apis/authen/v1/authorize",
	TokenURL: "https://open.feishu.cn/open-apis/authen/v2/oauth/token",
}

func OAuth2Config(appId, appSecret string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     appId,
		ClientSecret: appSecret,
		RedirectURL:  "http://localhost:8080/callback", // 请先添加该重定向 URL，配置路径：开发者后台 -> 开发配置 -> 安全设置 -> 重定向 URL -> 添加
		Endpoint:     oauthEndpoint,
		Scopes:       []string{"offline_access"}, // 如果你不需要 refresh_token，请注释掉该行，否则你需要先申请 offline_access 权限方可使用，配置路径：开发者后台 -> 开发配置 -> 权限管理
	}
}

func (cli *Client) GetUserInfo(userAccessToken string) (*larkauthen.GetUserInfoRespData, error) {
	resp, err := cli.client.Authen.V1.UserInfo.Get(context.Background(), larkcore.WithUserAccessToken(userAccessToken))
	// 处理错误
	if err != nil {
		return nil, err
	}

	// 服务端错误处理
	if !resp.Success() {
		return nil, fmt.Errorf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
	return resp.Data, nil
}
