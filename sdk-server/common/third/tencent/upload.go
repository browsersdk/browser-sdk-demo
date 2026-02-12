package tencent

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cos "github.com/tencentyun/cos-go-sdk-v5"
)

const CDN_DOMAIN = ""

// 新腾讯云配置
const TENCENT_SECRET_ID = ""
const TENCENT_SECRET_KEY = ""

// 腾讯云语音识别文字接口
const TENCENT_ASR_SECRET_ID = ""
const TENCENT_ASR_SECRET_KEY = ""

type TencentResourceEnum struct {
	Region string
	Bucket string
	Path   string
}

func VoiceToText(url, data, region string) (string, error) {
	credential := common.NewCredential(
		TENCENT_ASR_SECRET_ID,
		TENCENT_ASR_SECRET_KEY,
	)

	cpf := profile.NewClientProfile()
	client, err := asr.NewClient(credential, region, cpf)
	if err != nil {
		log.Printf("Failed to create ASR client: %v", err)
		return "", err
	}

	req := asr.NewSentenceRecognitionRequest()
	voiceFormat := strings.TrimPrefix(filepath.Ext(url), ".")
	if data == "" {
		fileContent, err := os.ReadFile(url)
		if err != nil {
			log.Printf("Failed to read file: %v", err)
			return "", err
		}
		data = base64.StdEncoding.EncodeToString(fileContent)
	}

	params := map[string]interface{}{
		"EngSerViceType": "16k_en",
		"SubServiceType": 2,
		"ProjectId":      0,
		"SourceType":     1,
		"VoiceFormat":    voiceFormat,
		"Data":           data,
		"DataLen":        len(data) * 3 / 4,
	}

	paramsBytes, err := json.Marshal(params)
	if err != nil {
		log.Printf("Failed to marshal params: %v", err)
		return "", err
	}

	req.FromJsonString(string(paramsBytes))

	resp, err := client.SentenceRecognition(req)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		log.Printf("Failed to recognize sentence: %v", err)
		return "", err
	}

	return *resp.Response.Result, nil
}

func UploadFile(region, bucket, fileName, srcfile string) (string, error) {
	f, err := os.Open(srcfile)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return "", err
	}
	defer f.Close()

	return Upload(region, bucket, fileName, f)
}

func UploadContent(region, bucket, fileName string, content []byte) (string, error) {
	buf := bytes.NewBuffer(content)
	return Upload(region, bucket, fileName, buf)
}

func Upload(region, bucket, fileName string, r io.Reader) (string, error) {
	c := GetCosClient(region, bucket)
	fileName = strings.TrimPrefix(fileName, "/")
	result, err := c.Object.Put(context.Background(), fileName, r, nil)
	if err != nil {
		slog.Error("Failed to upload", "err", err)
		return "", err
	}
	defer result.Body.Close()
	if result.StatusCode != 200 {
		slog.Error("Failed to upload", "statusCode", result.StatusCode)
		return "", fmt.Errorf("failed to upload, statusCode: %d", result.StatusCode)
	}
	return CDN_DOMAIN + fileName, nil
}

func DelFile(region, bucket string, fileName string) error {
	c := GetCosClient(region, bucket)
	_, err := c.Object.Delete(context.Background(), fileName)
	if err != nil {
		return err
	}
	return nil
}

func DownloadToByte(region, bucket string, fileName string) ([]byte, error) {
	c := GetCosClient(region, bucket)
	resp, err := c.Object.Get(context.Background(), fileName, nil)
	if err != nil {
		return nil, err
	}
	bs, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return bs, nil
}

func DownloadToFile(region, bucket string, fileName, saveName string) error {
	c := GetCosClient(region, bucket)
	_, err := c.Object.GetToFile(context.Background(), fileName, saveName, nil)
	if err != nil {
		return err
	}
	return nil
}

func Download(region, bucket string, fileName, saveName string) (io.ReadCloser, error) {
	c := GetCosClient(region, bucket)
	resp, err := c.Object.Get(context.Background(), fileName, nil)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func GetCosClient(region, bucket string) *cos.Client {
	u, _ := url.Parse(getUrl(region, bucket))
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  TENCENT_SECRET_ID,
			SecretKey: TENCENT_SECRET_KEY,
		},
	})
}

// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
func getUrl(region, bucket string) string {
	return fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region)
}
