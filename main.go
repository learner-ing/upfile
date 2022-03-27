package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"log"
	"flag"
)

const tarFirstDir = "file"

var dir,host string

func Init() {
	flag.StringVar(&dir, "dir", "", "Compressed folder,[/var/www]")
	flag.StringVar(&host, "host", "", "Server IP port,[192.168.1.1:8888]")
}
func main() {
	Init()
	flag.Parse()
	if dir=="" ||host=="" {
		fmt.Println("upfile -dir /var/www -host 192.168.1.1:8888")
		fmt.Println("upfile -dir /var/www -host https://192.168.1.1 还未开发")
		flag.PrintDefaults()
		os.Exit(0)
	}
	_,err := os.Stat(dir)
	if err!=nil {
		log.Fatalln("-dir error,",err)
	}
	start(host,dir,"nc")
}

func start(host,dir,module string)  {
	switch module {
	case "nc":
		conn, err := net.Dial("tcp", host)
		defer conn.Close()
		if err != nil {
			log.Fatalln("Connect "+host+"error,",err)
		}
		err = sendData(dir, conn)
		if err!=nil {
			log.Fatalln("sendData error,",err)
		}
	default:
		log.Fatalln("功能还未开发")
	}

}

//发送压缩文件
func sendData(path string, conn net.Conn) error {
	gw := gzip.NewWriter(conn)
	tw := tar.NewWriter(gw)
	defer gw.Close()
	defer tw.Close()
	return filepath.Walk(path, func(fileName string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//兼容windows
		fileName = strings.ReplaceAll(fileName, "\\", "/")
		fileHeader, err := tar.FileInfoHeader(fileInfo, "")
		if err != nil {
			return err
		}
		//替换绝对路径
		fileName = filepath.Clean(fileName)
		fileHeader.Name = tarFirstDir+"/"+strings.TrimPrefix(strings.ReplaceAll(fileName, filepath.Dir(fileName), ""), string(filepath.Separator))
		fileHeader.Format = tar.FormatGNU
		if err := tw.WriteHeader(fileHeader); err != nil {
			return err
		}
		if !fileInfo.Mode().IsRegular() {
			return nil
		}
		fileRead, err := os.Open(fileName)
		defer fileRead.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, fileRead)
		if err != nil {
			return err
		}
		return nil
	})
}
