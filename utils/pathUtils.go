package utils

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]
	return path
}

func CheckExist(fileName string) string {
	return path.Ext(fileName)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}
