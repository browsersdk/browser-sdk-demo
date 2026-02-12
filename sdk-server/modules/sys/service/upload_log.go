package service

import (
	"dilu/common/config"
	"dilu/common/third/file_store"
	"dilu/modules/sys/models"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"sync"

	"github.com/baowk/dilu-core/core/base"
)

type UploadLogService struct {
	*base.BaseService
}

var SerUploadLog = UploadLogService{
	base.NewService("sys"),
}

var once sync.Once

var fileStore file_store.FSHandler

func GetFsHandler() file_store.FSHandler {
	once.Do(
		func() {
			fileStore, _ = file_store.Setup(&config.Ext.FSCfg)
		},
	)
	return fileStore
}

func (s *UploadLogService) UploadFile(file *multipart.FileHeader, prefix string) (string, error) {
	if filePath, _, err := GetFsHandler().UploadFileHeader(file, prefix, ""); err != nil {
		return "", err
	} else {
		go s.Create(&models.UploadLog{
			Provider:   string(config.Ext.FSCfg.Select),
			Url:        filePath,
			Status:     1,
			SyncStatus: 0,
		})
		return filePath, nil
	}
}

func (s *UploadLogService) UploadReader(reader io.Reader, prefix string, fileName string) (string, error) {
	if filePath, _, err := GetFsHandler().Upload(reader, prefix, fileName, ""); err != nil {
		return "", err
	} else {
		go s.Create(&models.UploadLog{
			Provider:   string(config.Ext.FSCfg.Select),
			Url:        filePath,
			Status:     1,
			SyncStatus: 0,
		})
		return filePath, nil
	}
}

func (s *UploadLogService) DelFile(urlPath string) error {
	if urlPath == "" {
		return nil
	}
	var ul models.UploadLog
	err := s.DB().Where("url = ?", urlPath).Find(&ul).Error
	if err != nil {
		return err
	}
	if ul.Id == 0 {
		return nil
	}
	if ul.Status != 1 {
		return nil
	}
	ula := models.UploadLog{
		Id:         ul.Id,
		Status:     -1,
		SyncStatus: 1,
	}
	if err := s.UpdateById(&ula); err != nil {
		return err
	}
	key := urlPath
	if strings.HasPrefix(urlPath, "http") {
		curl, _ := url.Parse(urlPath)
		key = curl.Path
	}
	if err := GetFsHandler().DeleteFile(key, ""); err == nil {
		ula.SyncStatus = 3
		if err := s.UpdateById(&ula); err != nil {
			return err
		}
	} else {
		ula.SyncStatus = -1
		if err := s.UpdateById(&ula); err != nil {
			return err
		}
		return err
	}
	return nil
}
