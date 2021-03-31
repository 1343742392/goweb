package filetool

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func ParsePath(path string)string{
	if runtime.GOOS == "windows" {
		return 	path[0:strings.LastIndex(path, `\`)]
	}else{
		return path[0:strings.LastIndex(path, `/`)]
	}
}

func GetNowPath() string {
	filePath := os.Args[0]
	if  runtime.GOOS == "windows" {
		filePath = strings.ReplaceAll(filePath, "/", "\\")
	} 
	return ParsePath(filePath);
}

func GetFileString(path string) (string, error) {
	bytes, err := GetFileBytes(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func SetFile(path string, content string) error {
	var f *os.File
	var err1 error
	if CheckFileIsExist(path) {
		os.Remove(path)
		f, err1 = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModeAppend|os.ModePerm) //打开文件
	} else {
		f, err1 = os.Create(path)
	}
	if err1 != nil {
		return err1
	}
	_, err2 := io.WriteString(f, content)
	if err2 != nil {
		return err1
	}
	return nil
}

func GetFileBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, rerr := file.Read(buffer)
	if rerr != nil {
		fmt.Println(rerr)
		return nil, rerr
	}
	return buffer, nil
}

/*
C:/a-b
	|-src
	|-bin
	|-run.go
input :
	c:/a/b
output:
	src,bin
*/
func GetDirs(path string) []string {
	fileInfoList, _ := ioutil.ReadDir(path)
	res := []string{}
	for _, f := range fileInfoList {
		if f.IsDir() {
			res = append(res, f.Name())
		}
	}
	return res
}
