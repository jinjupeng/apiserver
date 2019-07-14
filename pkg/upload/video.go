package upload

import (
	"apiserver/pkg/file"
	"apiserver/util"
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// 获取视频完整访问URL
func GetVideoFullUrl(name string) string {
	return viper.GetString("video.VideoPrefixUrl") + "/" + GetVideoPath() + name
}

// 获取视频名称
func GetVideoName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// 获取视频路径
func GetVideoPath() string {
	return viper.GetString("video.RuntimeRootPath")
}

// 获取视频完整路径
func GetVideoFullPath() string {
	return viper.GetString("video.RuntimeRootPath") + GetVideoPath()
}

// 检查视频后缀名
func CheckVideoExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range viper.GetStringSlice("video.VideoAllowExts") {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// 检查视频大小
func CheckVideoSize(f multipart.File) bool {
	_, err := file.GetSize(f)
	if err != nil {
		log.Fatal("视频检查异常", err)
		return false
	}

	// return size <= viper.GetInt("video.VideoMaxSize")
	return true
}

// 检查视频
func CheckVideo(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
