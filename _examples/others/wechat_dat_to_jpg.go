package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//
func init() {
	log.SetFlags(0)
}

func main() {
	// 微信 .dat 文件目录
	// C:\Users\taadis\Documents\WeChat Files\taadis\FileStorage\Image\2019-07

	// 输入目录
	dirname, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panicln(err)
	}
	// 输出目录
	outputDirname := filepath.Join(dirname, "output")
	log.Printf("当前程序执行路径: %s", dirname)
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Panicln(err)
	}
	var datFiles []string // dat 文件集合
	for key, fileInfo := range fileInfos {
		// 判断是目录还是文件
		if fileInfo.IsDir() {
			log.Println("dir: ", key)
		} else {
			// 判断是不是 .dat 后缀的文件
			if !strings.HasSuffix(fileInfo.Name(), ".dat") {
				break
			}
			inFilename := fmt.Sprintf("%s\\%s", dirname, fileInfo.Name())
			datFiles = append(datFiles, inFilename)
		}
	}

	//
	if len(datFiles) < 1 {
		log.Println("当前路径下没有找到需要处理的 .dat 文件")
		return
	}

	//
	for key, datFile := range datFiles {
		datBytes, err := ioutil.ReadFile(datFile) //
		if err != nil {
			log.Println(err)
		}

		imgBytes, err := WechatDatToImage(datBytes)
		if err != nil {
			log.Panicln(err)
		}

		datFilename := filepath.Base(datFile) + ".jpg"
		outFilename := filepath.Join(outputDirname, datFilename)
		err = ensureDirExist(outFilename)
		if err != nil {
			log.Panicln(err)
		}
		err = ioutil.WriteFile(outFilename, imgBytes, 0777)
		if err != nil {
			log.Panicln(err)
		}
		log.Printf("compeleted: %d %s", key, outFilename)
	}
}

// 微信 dat 字节数组转图片字节数组
func WechatDatToImage(dat []byte) (img []byte, err error) {
	var buffer bytes.Buffer
	for _, value := range dat {
		err = buffer.WriteByte(value ^ 0xF2)
		if err != nil {
			return
		}
	}
	img = buffer.Bytes()
	return
}

// 写个函数，确保文件夹存在，省的重复写
func ensureDirExist(path string) error {
	dir := filepath.Dir(path)
	exists := isPathExists(dir)
	if !exists {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 判断路径是否存在
func isPathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
