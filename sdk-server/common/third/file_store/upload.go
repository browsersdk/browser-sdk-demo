package file_store

import (
	"dilu/common/third/file_store/cfg"
	"dilu/common/third/file_store/oss"
	"dilu/common/third/file_store/tencent"
	"io"
	"mime/multipart"
)

// OSS 对象存储接口
type FSHandler interface {
	UploadFile(fileName, rootPath, bucketName string) (filePath string, fileKey string, err error)
	UploadFileHeader(file *multipart.FileHeader, rootPath, bucketName string) (filePath string, fileKey string, err error)
	Upload(reader io.Reader, prefix, fileName string, bucketName string) (filePath string, fileKey string, err error)
	DeleteFile(key, bucketName string) error
}

func Setup(fsCfg *cfg.FileStoreCfg) (FSHandler, error) {
	var fs FSHandler
	var err error
	switch fsCfg.Select {

	case cfg.OssStore:
		fs, err = oss.New(fsCfg)
	case cfg.TencentStore:
		fs, err = tencent.NewCOSHandler(fsCfg)
		// case cfg.LocalStore:
		// 	return local.NewLocalHandler(cfg.RootPath)
	}
	return fs, err
}
