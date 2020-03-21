package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateUserInfo(userName, userPass string) bool {
	filePath := "./.user/" + userName + ".pass"

	// check directory first
	ensureDir := ensureDir(filePath)
	if !ensureDir {
		return false
	}

	// open or create file second
	fileHandler, error := openOrCreate(filePath)
	if error != nil {
		fmt.Printf("CreateUserInfo Error：can not open or create file %s\n", filePath)
		return false
	}

	//defer fileHandler.Close()

	// TODO encrypt the userpass

	// put userPass to the open file third
	_, writeErr := fileHandler.WriteString(userPass)
	if writeErr != nil {
		fmt.Printf("CreateUserInfo Error: can not put content to the file %s, %s\n", filePath, writeErr)
		return false
	}

	return true
}

// open the giving file return file
func openOrCreate(filePath string) (*os.File, error) {
	fileSource, openErr := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	return fileSource, openErr
}

// init user info directory
func ensureDir(filePath string) bool {
	// check directory first
	dirName := filepath.Dir(filePath)
	dirErr := os.MkdirAll(dirName, 0777)

	if dirErr != nil {
		fmt.Printf("ensureDir Error: can not create directory %s \n", dirName)
		return false
	}


	fmt.Printf("ensureDir Ok: %s \n", dirName)
	return true
}