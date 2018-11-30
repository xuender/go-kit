package kit

import (
	face "github.com/Kagami/go-face"
	utils "github.com/xuender/go-utils"
)

// Classifier 人脸分类器
type Classifier struct {
	Named   map[string][]Face // 命名的人脸
	Waiting [][]Face          // 待命名的
	Rec     *face.Recognizer  // 人脸识别
}

// SetName 设置名称
func (c *Classifier) SetName(id utils.ID, name string) {
	// TODO 设置名称
}

// Classify 分类
func (c *Classifier) Classify(faces []Face) {
	// TODO 分类
	// 大量照片数据分析
	// 测试大量照片分类性能
}
