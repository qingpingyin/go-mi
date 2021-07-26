package upload

import (
	"MI/pkg/setting"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func UploadFile(file multipart.File,fileSize int64,fileName string)(string,error){
	putPolicy := storage.PutPolicy{
		Scope: setting.QiNiuYunConf.Bucket,
	}
	mac := qbox.NewMac(setting.QiNiuYunConf.AccessKey,setting.QiNiuYunConf.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{}


	err := formUploader.Put(context.Background(), &ret, upToken,fileName, file, fileSize, &putExtra)
	if err != nil {
		return "",err
	}
	return setting.QiNiuYunConf.Server+ret.Key ,nil
}
