package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysCfgService struct {
	*base.BaseService
}

var SerSysCfg = SysCfgService{
	base.NewService("sys"),
}

// func (s *SysCfgService) TTS(req *dto.TTSReq) (string, error) {
// 	fac, err := tts.NewFactory(config.Ext.TTS)
// 	if err != nil {
// 		return "", err
// 	}
// 	sMetaData := &sentence.MetaData{Sid: 1, Stage: sentence.BeforeSendToRTC}
// 	ctx := aigcCtx.NewContext(context.WithValue(context.Background(), enums.SentenceMetaData, sMetaData), sMetaData)

// 	ttsClient, err := fac.CreateTTS(ctx, "", req.VoiceId, "", 0)

// 	ttsClient.Send(context.Background(), 1, req.Content, true)

// 	filePath := "tmp" + time.Now().Format("20060102150405") + ".mp3"

// 	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()

// 	for {
// 		data, ok := <-ttsClient.GetResult()
// 		if ok {
// 			_, err = f.Write(data)
// 			if err != nil {
// 				fmt.Println("写入文件失败：", err)
// 				return "", err
// 			}
// 		} else {
// 			break
// 		}
// 	}

// 	urlPath, _, err := GetFsHandler().UploadFile(filePath, "tts", "")
// 	defer os.Remove(filePath)
// 	if err != nil {
// 		return "", err
// 	} else {
// 		return urlPath, nil
// 	}
// }
