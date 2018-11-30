package kit

import (
	face "github.com/Kagami/go-face"
	utils "github.com/xuender/go-utils"
)

// Face 脸
type Face struct {
	ID   utils.ID  // ID
	FID  []byte    // 照片文件ID
	Face face.Face // 人脸识别
}
