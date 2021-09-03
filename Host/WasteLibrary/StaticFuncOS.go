package WasteLibrary

import (
	"os"
	"os/exec"
	"os/user"
	"time"
)

//GetCurrentUser
func GetCurrentUser() string {
	user, err := user.Current()
	if err != nil {
		LogErr(err)
	}

	username := user.Username
	return username
}

//IsFileExists
func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//RemoveFile
func RemoveFile(filePath string) {
	LogStr("Remove File : " + filePath)
	cmdRm := exec.Command("rm", filePath)
	errRm := cmdRm.Start()
	if errRm != nil {
		LogErr(errRm)
	}
}

//GetTime
func GetTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05")
}