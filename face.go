package kit

import (
	face "github.com/Kagami/go-face"
	utils "github.com/xuender/go-utils"
)

// Face 脸
type Face struct {
	ID   utils.ID  // ID
	PID  []byte    // 照片文件ID
	Face face.Face // 人脸识别
}

// NewFace 新建脸
func NewFace(f face.Face, pid []byte) *Face {
	return &Face{
		Face: f,
		ID:   utils.NewID(FacePrefix),
		PID:  pid,
	}
}
