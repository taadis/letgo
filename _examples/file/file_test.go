package file

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// FindDirectories 打印指定目录下所有子目录名称
func FindDirectories(pwd string) {
	fileInfos, err := ioutil.ReadDir(pwd)
	if err != nil {
		fmt.Printf("ioutil.ReadDir error:%+v", err)
		return
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			wd := fmt.Sprintf("%s/%s", pwd, fileInfo.Name())
			fmt.Println(wd)
			FindDirectories(wd)
		}
	}
}

func TestFindDirectories(t *testing.T) {
	FindDirectories("/Users/taadis/my/letgo")
}
