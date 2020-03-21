package lib

import (
	"fmt"
	user "gin-rest-api/models"
	"io/ioutil"
	"os"
	"path/filepath"
)

const UserInfoDir = "./.user/"

func GetUserName(userName string) bool {
	filePath := UserInfoDir + userName + ".pass"

	// check user file exists
	_, error := os.Stat(filePath)
	if error == nil {
		return true
	}

	fmt.Printf("GetUser Error: user file not exists %s, %s", filePath, error)
	return false
}

func GetUserPass(userName string) string  {
	var userPass = ""
	// get file content from user.pass file first
	filePath := UserInfoDir + userName + ".pass"
	fileContent, error := ioutil.ReadFile(filePath)
	if error != nil {
		fmt.Printf("GetUserPass Error：can not get content from file %s, %s\n", filePath, error)
		return userPass
	}

	userPass = string(fileContent)

	return userPass
}

func CreateUserInfo(userName, userPass string) bool {
	filePath := UserInfoDir + userName + ".pass"

	// check directory first
	ensureDir := ensureDir(filePath)
	if !ensureDir {
		return false
	}

	// open or create file second
	fileHandler, error := openOrCreate(filePath)
	if error != nil {
		fileHandler.Close()
		fmt.Printf("CreateUserInfo Error：can not open or create file %s, %s\n", filePath, error)
		return false
	}

	// encrypt the userpass
	encryptUserPass := user.EncryptUserPass(userPass)
	// put userPass to the open file third
	_, writeErr := fileHandler.Write(encryptUserPass)
	if writeErr != nil {
		fileHandler.Close()
		fmt.Printf("CreateUserInfo Error: can not put content to the file %s, %s\n", filePath, writeErr)
		return false
	}

	fileHandler.Close()
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