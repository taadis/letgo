package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	//"os"
)

//
func init() {
	log.SetFlags(0)
}

func main() {
	fmt.Println("wechat_dat_to_jpg.go")
	// 微信 dat 文件目录
	// C:\Users\taadis\Documents\WeChat Files\taadis\FileStorage\Image\2019-07

	// 源目录
	dirname := `E:\\dats`
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Panicln(err)
	}
	for key, fileInfo := range fileInfos {
		// 判断是目录还是文件
		if fileInfo.IsDir() {
			log.Println("dir: ", key)
		} else {
			// 判断是不是 .dat 后缀的文件
			if !strings.HasSuffix(fileInfo.Name(), ".dat") {
				break
			}

			// 读取文件字节数据
			inFilename := fmt.Sprintf("%s\\%s", dirname, fileInfo.Name())
			dat, err := ioutil.ReadFile(inFilename) //
			if err != nil {
				log.Println(err)
			}

			img, err := WechatDatToImage(dat)
			if err != nil {
				log.Panicln(err)
			}

			outFilename := fmt.Sprintf("%s.jpg", inFilename)
			err = ioutil.WriteFile(outFilename, img, 0777)
			if err != nil {
				log.Panicln(err)
			}
			log.Printf("compeleted: %d %s", key, outFilename)
		}
	}
}

// 微信 dat 字节数组转图片字节数组
func WechatDatToImage(dat []byte) (img []byte, err error) {
	var buffer bytes.Buffer
	for _, value := range dat {
		//err = buffer.WriteByte(value ^ 0x75) // 网上的不对
		err = buffer.WriteByte(value ^ 0xF2) // ok
		if err != nil {
			return
		}
	}
	img = buffer.Bytes()
	return
}
