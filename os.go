package kit

import (
	"fmt"
	"os"
)

// Mkdir 创建目录
func Mkdir(dir string) error {
	if fi, err := os.Stat(dir); err == nil {
		if !fi.IsDir() {
			return fmt.Errorf("创建目录失败： %s 已经存在，并且不是目录。", dir)
		}
	} else if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
