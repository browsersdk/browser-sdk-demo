package oss

import (
	"dilu/common/third/file_store/cfg"
	"dilu/common/utils"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/baowk/dilu-core/core"
)

func New(cfg *cfg.FileStoreCfg) (*AliyunOSS, error) {
	client, err := newClient(cfg)
	if err != nil {
		return nil, err
	}

	return &AliyunOSS{
		cfg:    cfg,
		client: client,
	}, nil
}

type AliyunOSS struct {
	cfg    *cfg.FileStoreCfg
	client *oss.Client
}

func (e *AliyunOSS) UploadFile(fileName, prefix, bucketName string) (filePath string, fileKey string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error("Failed to open file", "err", err)
		return "", "", err
	}
	defer file.Close()

	return e.Upload(file, prefix, fileName, bucketName)
}

func (e *AliyunOSS) UploadFileHeader(fileHeader *multipart.FileHeader, prefix, bucketName string) (filePath string, fileKey string, err error) {
	// 读取本地文件。
	file, err := fileHeader.Open()
	if err != nil {
		slog.Error("Failed to open file", "err", err)
		return "", "", err
	}

	defer file.Close() // 创建文件 defer 关闭
	return e.Upload(file, prefix, fileHeader.Filename, bucketName)
}

func (e *AliyunOSS) Upload(reader io.Reader, prefix, fileName string, bucketName string) (filePath string, fileKey string, err error) {
	fileKey = e.cfg.GetFileKey(prefix, utils.GetFileExtension(fileName))
	if bucketName == "" {
		bucketName = e.cfg.Bucket
	}
	bucket, err := e.getBucket(bucketName)
	err = bucket.PutObject(fileKey, reader)
	if err != nil {
		slog.Error("Failed to upload file", "err", err)
		return "", "", err
	}
	if e.cfg.Domain == "" {
		filePath = getDomain(e.cfg.Region, bucketName) + fileKey
	} else {
		filePath = e.cfg.Domain + fileKey

	}
	return
}

func (e *AliyunOSS) DeleteFile(key, bucketName string) error {
	if bucketName == "" {
		bucketName = e.cfg.Bucket
	}
	bucket, err := e.getBucket(bucketName)
	if err != nil {
		slog.Error("function AliyunOSS.NewBucket() Failed", "err", err)
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		core.Log.Error("function bucketManager.Delete() Filed", err)
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func (e *AliyunOSS) getBucket(name string) (*oss.Bucket, error) {
	if name == "" {
		name = e.cfg.Bucket
	}
	// 获取存储空间。
	bucket, err := e.client.Bucket(name)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func newClient(cfg *cfg.FileStoreCfg) (*oss.Client, error) {
	var endpoint string
	if cfg.Internal {
		endpoint = fmt.Sprintf("https://oss-%s-internal.aliyuncs.com", cfg.Region)
	} else {
		endpoint = fmt.Sprintf("https://oss-%s.aliyuncs.com", cfg.Region)
	}

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, cfg.SecretID, cfg.SecretKey)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getDomain(region, bucketName string) string {
	return fmt.Sprintf("https://%s.oss-%s.aliyuncs.com/", bucketName, region)
}
