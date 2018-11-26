package kit

import (
	"fmt"
	"os"
	"time"

	"github.com/Kagami/go-face"
	"github.com/h2non/filetype"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

// Photo 照片处理
type Photo struct {
	Size  int64       // 尺寸
	Ca    time.Time   // 创建时间
	Faces []face.Face // 人脸识别
	Make  string      // 相机
	Model string      // 型号
}

// NewPhoto 照片信息
func NewPhoto(file string, rec *face.Recognizer) (photo *Photo, err error) {
	// 图片类型
	t, err := filetype.MatchFile(file)
	if err != nil {
		return
	}
	if t.MIME.Type != "image" {
		err = fmt.Errorf("file type is not image")
		return
	}
	if t.MIME.Subtype != "jpeg" {
		err = fmt.Errorf("image is not JPEG")
		return
	}
	// 人脸识别
	faces, err := rec.RecognizeFile(file)
	if err != nil {
		return
	}
	photo = &Photo{
		Faces: faces,
	}
	// 基本信息
	if s, err := os.Stat(file); err == nil {
		photo.Ca = s.ModTime()
		photo.Size = s.Size()
	}
	// 扩展信息
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if x, e := exif.Decode(f); e == nil {
		if date, e := x.DateTime(); e == nil {
			photo.Ca = date
		}
		if make, e := x.Get(exif.Make); e == nil {
			photo.Make, _ = make.StringVal()
		}
		if model, e := x.Get(exif.Model); e == nil {
			photo.Model, _ = model.StringVal()
		}
	}
	return
}
func init() {
	exif.RegisterParsers(mknote.All...)
}
