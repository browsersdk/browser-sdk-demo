package aliyun

import (
	"dilu/common/third/sms/config"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	extCfg "dilu/common/config"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/aliyun/credentials-go/credentials"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliyunSms struct {
	smsCfg *config.SmsConfig
}

func NewAliyunSms(cfg *config.SmsConfig) *AliyunSms {
	return &AliyunSms{smsCfg: cfg}
}

func (s *AliyunSms) Send(phone string, code string, tmpId string) error {
	//参数一：连接的节点地址（有很多节点选择，这里我选择杭州）
	//参数二：AccessKey ID
	//参数三：AccessKey Secret
	client, err := dysmsapi.NewClientWithAccessKey(s.smsCfg.Region, s.smsCfg.SecretId, s.smsCfg.SecretKey)
	if err != nil {
		slog.Error("AliYun sms new client fail", "error", fmt.Sprintf("%+v", err)) //处理错误
		return err
	}

	request := dysmsapi.CreateSendSmsRequest() //创建请求
	request.Scheme = "https"                   //请求协议
	request.PhoneNumbers = phone               //接收短信的手机号码
	request.SignName = s.smsCfg.Sign           //短信签名名称
	if tmpId != "" {
		request.TemplateCode = tmpId //短信模板ID
	} else {
		request.TemplateCode = s.smsCfg.TemplateId //短信模板ID
	}

	par, err := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": code,
	})
	if err != nil {
		return err
	}
	request.TemplateParam = string(par) //将短信模板参数传入短信模板

	response, err := client.SendSms(request) //调用阿里云API发送信息
	if err != nil {
		slog.Error("AliYun sms fail", "error", err, "req", request, "resp", response) //处理错误
		return err
	}
	return nil
}

// Description:
//
// 使用凭据初始化账号Client
//
// @return Client
//
// @throws Exception
func CreateClient() (_result *openapi.Client, _err error) {
	// 工程代码建议使用更安全的无AK方式，凭据配置方式请参见：https://help.aliyun.com/document_detail/378661.html。
	// credential, _err := credential.NewCredential(nil)
	// if _err != nil {
	// 	return _result, _err
	// }

	crtConfig := new(credentials.Config).
		SetType("access_key").
		SetAccessKeyId(extCfg.Ext.Sms.SecretId).
		SetAccessKeySecret(extCfg.Ext.Sms.SecretKey)

	akCredential, err := credentials.NewCredential(crtConfig)
	if err != nil {
		return
	}

	config := &openapi.Config{
		Credential: akCredential,
	}

	// Endpoint 请参考 https://api.aliyun.com/product/Dypnsapi
	config.Endpoint = tea.String("dypnsapi.aliyuncs.com")
	_result = &openapi.Client{}
	_result, _err = openapi.NewClient(config)
	return _result, _err
}

const (
	AndroidSchemeCode  = "FA000000007380804002"
	AndroidPackageName = "com.taoxinkey.ai"
	AndroidPackageSign = "13f4dbc362ce1e7a09720d40ff48a942"

	IOSSchemeCode = "FA000000007340904002"
	IOSBundleId   = "com.taoxinkey.ai"
)

func CreateApiInfo(action string) (_result *openapi.Params) {
	params := &openapi.Params{
		// 接口名称
		Action: tea.String(action),
		// 接口版本
		Version: tea.String("2017-05-25"),
		// 接口协议
		Protocol: tea.String("HTTPS"),
		// 接口 HTTP 方法
		Method:   tea.String("POST"),
		AuthType: tea.String("AK"),
		Style:    tea.String("RPC"),
		// 接口 PATH
		Pathname: tea.String("/"),
		// 接口请求体内容格式
		ReqBodyType: tea.String("json"),
		// 接口响应体内容格式
		BodyType: tea.String("json"),
	}
	_result = params
	return _result
}

func GetFusionAuthToken(platform string) (data string, _err error) {
	client, _err := CreateClient()
	if _err != nil {
		return "", _err
	}

	params := CreateApiInfo("GetFusionAuthToken")

	// query params
	queries := map[string]interface{}{}
	if platform == "Android" {
		queries["SchemeCode"] = tea.String(AndroidSchemeCode)
		queries["PackageName"] = tea.String(AndroidPackageName)
		queries["PackageSign"] = tea.String(AndroidPackageSign)
		queries["Platform"] = tea.String("Android")
	} else if platform == "iOS" {
		queries["SchemeCode"] = tea.String(IOSSchemeCode)
		queries["BundleId"] = tea.String(IOSBundleId)
		queries["Platform"] = tea.String("iOS")
	} else {
		return "", _err
	}
	queries["DurationSeconds"] = tea.Int(900)
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值实际为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	var _result map[string]interface{}
	_result, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return "", _err
	}
	fmt.Println("GetOneClickToken result:", _result)
	if body, ok := _result["body"]; ok {
		bodyMap, ok := body.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("unexpected body type: %T", body)
		}
		if code, ok := bodyMap["Code"]; ok && code == "OK" {
			if model, ok := bodyMap["Model"]; ok {
				return model.(string), nil
			}
			return "", fmt.Errorf("Model not found in response")
		} else if message, ok := bodyMap["Message"]; ok {
			return "", fmt.Errorf("error from API: %s", message.(string))
		}
	}

	return "", fmt.Errorf("unexpected response format: %v", _result)
}

func VerifyWithFusionAuthToken(verifyToken string) (phone string, _err error) {
	client, _err := CreateClient()
	if _err != nil {
		return "", _err
	}

	params := CreateApiInfo("VerifyWithFusionAuthToken")
	// query params
	queries := map[string]interface{}{}
	queries["VerifyToken"] = tea.String(verifyToken)
	// if platform == "Android" {
	// 	queries["SchemeCode"] = tea.String(AndroidSchemeCode)
	// } else if platform == "iOS" {
	// 	queries["SchemeCode"] = tea.String(IOSSchemeCode)
	// }
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值实际为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	var _result map[string]interface{}
	_result, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return "", _err
	}
	fmt.Println("OneClickVerify result:", _result)
	if body, ok := _result["body"]; ok {
		data, err := json.Marshal(body)
		if err != nil {
			return "", err
		}
		var res Res
		if err = json.Unmarshal(data, &res); err != nil {
			return "", err
		}

		fmt.Println("OneClickVerify res:", res)
		if res.Code == "OK" {
			phone = res.Model.PhoneNumber
			if phone == "" || phone == "UNKNOWN" {
				return "", errors.New("phone number is empty or unknown")
			}
			return
		}
		return "", errors.New(res.Message)
	}

	return "", fmt.Errorf("unexpected response format: %v", _result)
}

func VerifyMobile(accessCode string, phoneNumber string, outId string) (*GateVerifyResultDTO, error) {
	client, _err := CreateClient()
	if _err != nil {
		return nil, _err
	}

	params := CreateApiInfo("VerifyMobile")
	// query params
	queries := map[string]interface{}{}
	queries["AccessCode"] = tea.String(accessCode)
	queries["PhoneNumber"] = tea.String(phoneNumber)
	queries["OutId"] = tea.String(outId)
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	var _result map[string]interface{}
	_result, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return nil, _err
	}
	fmt.Println("VerifyMobile result:", _result)
	if body, ok := _result["body"]; ok {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		var res Res
		if err = json.Unmarshal(data, &res); err != nil {
			return nil, err
		}

		fmt.Println("VerifyMobile res:", res)
		if res.Code == "OK" {
			return &res.GateVerifyResultDTO, nil
		}
		return nil, errors.New(res.Message)
	}
	return nil, fmt.Errorf("unexpected response format: %v", _result)
}

func GetMobile(token string, outId string) (string, error) {
	client, _err := CreateClient()
	if _err != nil {
		return "", _err
	}

	params := CreateApiInfo("GetMobile")
	// query params
	queries := map[string]interface{}{}
	queries["AccessToken"] = tea.String(token)
	queries["OutId"] = tea.String(outId)
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	var _result map[string]interface{}
	_result, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return "", _err
	}
	fmt.Println("GetMobile result:", _result)
	if body, ok := _result["body"]; ok {
		data, err := json.Marshal(body)
		if err != nil {
			return "", err
		}
		var res Res
		if err = json.Unmarshal(data, &res); err != nil {
			return "", err
		}

		fmt.Println("GetMobile res:", res)
		if res.Code == "OK" {
			return res.GetMobileResultDTO.Mobile, nil
		}
		return "", errors.New(res.Message)
	}
	return "", fmt.Errorf("unexpected response format: %v", _result)
}

type Res struct {
	Message   string `json:"Message"` //状态码的描述。
	RequestId string `json:"RequestId"`
	Model     struct {
		PhoneNumber  string `json:"PhoneNumber"`  // 手机号码
		VerifyResult string `json:"VerifyResult"` // 验证结果
		PhoneScore   int64  `json:"PhoneScore"`   // 手机号码评分
	} `json:"Model,omitempty"`
	GateVerifyResultDTO GateVerifyResultDTO `json:"GateVerifyResultDTO,omitempty"`
	GetMobileResultDTO  GetMobileResultDTO  `json:"GetMobileResultDTO,omitempty"`
	Code                string              `json:"Code"`    //状态码，取值：OK、InvalidAccessKeyId、InvalidSignature、SignatureDoesNotMatch、InvalidAction、InvalidVersion、InvalidParameter、InternalError等。
	Success             bool                `json:"Success"` //是否成功，取值：true、false。
}

type GateVerifyResultDTO struct {
	VerifyResult string `json:"VerifyResult"`
	VerifyId     string `json:"VerifyId"`
}

type GetMobileResultDTO struct {
	Mobile string `json:"Mobile"`
}
