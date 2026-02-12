package tencent

import (
	"context"
	"dilu/common/third/file_store/cfg"
	"dilu/common/utils"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

type COSHandler struct {
	cfg *cfg.FileStoreCfg
	//client *cos.Client
}

func NewCOSHandler(fsc *cfg.FileStoreCfg) (*COSHandler, error) {
	return &COSHandler{
		cfg: fsc,
	}, nil
}

func (h *COSHandler) UploadFile(fileName, prefix, bucketName string) (filePath string, fileKey string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error("Failed to open file", "err", err)
		return "", "", err
	}
	defer file.Close()

	filePath, fileKey, err = h.Upload(file, prefix, fileName, bucketName)
	return
}

func (h *COSHandler) UploadFileHeader(fileHeader *multipart.FileHeader, prefix, bucketName string) (filePath string, fileKey string, err error) {
	//fileName := FileHeader.Filename
	file, err := fileHeader.Open()
	if err != nil {
		slog.Error("Failed to open file", "err", err)
		return "", "", err
	}
	defer file.Close()

	filePath, fileKey, err = h.Upload(file, prefix, fileHeader.Filename, bucketName)
	return
}

func (h *COSHandler) Upload(r io.Reader, prefix, fileName string, bucket string) (string, string, error) {
	if bucket == "" {
		bucket = h.cfg.Bucket
	}

	fileKey := h.cfg.GetFileKey(prefix, utils.GetFileExtension(fileName))

	c, err := getCosClient(h.cfg.SecretID, h.cfg.SecretKey, h.cfg.Region, bucket)
	if err != nil {
		return "", "", err
	}

	result, err := c.Object.Put(context.Background(), fileKey, r, nil)
	if err != nil {
		slog.Error("Failed to upload", "err", err)
		return "", "", err
	}
	defer result.Body.Close()
	if result.StatusCode != 200 {
		slog.Error("Failed to upload", "statusCode", result.StatusCode)
		return "", "", fmt.Errorf("failed to upload, statusCode: %d", result.StatusCode)
	}
	return h.cfg.Domain + fileKey, fileKey, nil
}

// func (h *COSHandler) Upload(region, prefix, bucket, fileName string, r io.Reader) (string, string, error) {

// 	if bucket == "" {
// 		bucket = h.cfg.Bucket
// 	}

// 	fileKey := h.cfg.GetFileKey(prefix, utils.GetFileExtension(fileName))

// 	c, err := getCosClient(h.cfg.SecretID, h.cfg.SecretKey, h.cfg.Region, bucket)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	result, err := c.Object.Put(context.Background(), fileKey, r, nil)
// 	if err != nil {
// 		slog.Error("Failed to upload", "err", err)
// 		return "", "", err
// 	}
// 	defer result.Body.Close()
// 	if result.StatusCode != 200 {
// 		slog.Error("Failed to upload", "statusCode", result.StatusCode)
// 		return "", "", fmt.Errorf("failed to upload, statusCode: %d", result.StatusCode)
// 	}
// 	return h.cfg.Domain + fileKey, fileKey, nil
// }

func (h *COSHandler) DeleteFile(key, bucketName string) error {
	if bucketName == "" {
		bucketName = h.cfg.Bucket
	}

	c, err := getCosClient(h.cfg.SecretID, h.cfg.SecretKey, h.cfg.Region, bucketName)
	if err != nil {
		return err
	}
	_, err = c.Object.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func (h *COSHandler) DownloadToByte(region, bucket string, fileName string) ([]byte, error) {
	if bucket == "" {
		bucket = h.cfg.Bucket
	}

	c, err := getCosClient(h.cfg.SecretID, h.cfg.SecretKey, h.cfg.Region, bucket)
	if err != nil {
		return nil, err
	}

	resp, err := c.Object.Get(context.Background(), fileName, nil)
	if err != nil {
		return nil, err
	}
	bs, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return bs, nil
}

func (h *COSHandler) DownloadToFile(region, bucket string, fileName, saveName string) error {
	if bucket == "" {
		bucket = h.cfg.Bucket
	}

	c, err := getCosClient(h.cfg.SecretID, h.cfg.SecretKey, h.cfg.Region, bucket)
	if err != nil {
		return err
	}

	_, err = c.Object.GetToFile(context.Background(), fileName, saveName, nil)
	if err != nil {
		return err
	}
	return nil
}

func (h *COSHandler) Download(region, bucket string, fileName, saveName string) (io.ReadCloser, error) {
	if bucket == "" {
		bucket = h.cfg.Bucket
	}

	c, err := getCosClient(h.cfg.SecretID, h.cfg.SecretKey, h.cfg.Region, bucket)
	if err != nil {
		return nil, err
	}

	resp, err := c.Object.Get(context.Background(), fileName, nil)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func getCosClient(secretId, secretKey, region, bucket string) (*cos.Client, error) {
	u, err := url.Parse(getUrl(region, bucket))
	if err != nil {
		return nil, err
	}
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey,
		},
	}), nil
}

// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
func getUrl(region, bucket string) string {
	return fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region)
}
