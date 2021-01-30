package conf

import (
	"os"
	"path/filepath"
)

const (
	ConfFileName = "minigui.conf"
)

func GetDirList() []string {
	return []string{
		GetExecDir(),
		GetConfDir(),
	}
}

func GetExecDir() string {
	return filepath.Dir(os.Args[0])
}

func GetConfDir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return userConfigDir
}

// 加载资源
//func LoadFile() {
//fyne.LoadResourceFromPath()
//fyne.LoadResourceFromURLString()
//fyne.NewStaticResource()
//}
