package share

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func MkDir(dst string) (bool, error) {
	dstArr := strings.Split(dst, string(os.PathSeparator))
	dstLen := len(dstArr)
	fPath := string("")
	if dstLen < 1 {
		return false, nil
	}
	for key, temp := range dstArr {

		if key == dstLen-1 {
			break
		}

		if fPath == "" {
			fPath = temp
		} else {
			fPath = fPath + string(os.PathSeparator) + temp
		}

		exist, isDir, err := PathExists(fPath)
		if err != nil {
			return false, err
		}

		if exist || isDir {
			continue
		}

		if !exist {
			err := os.Mkdir(fPath, os.ModePerm)
			if err != nil {
				return false, err
			}
		}
	}
	return false, nil
}

func CopyFile(src, dst string) (bool, error) {

	_, err := MkDir(dst)
	if err != nil {
		fmt.Print(err)
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return false, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return false, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return false, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return false, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes > 0, err
}

func PathExists(path string) (bool, bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return true, info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, false, nil
	}
	return false, false, err
}

func InArray(value string, arr []string) bool {

	var result = false

	for _, v := range arr {
		if v == value {
			result = true
			break
		}
	}
	return result
}

func FindRecentModifyFile(fileName string, interval int64, timestamp int64, noWatch []string, result *[]string) []string {

	list, err := ioutil.ReadDir(fileName)
	if err != nil {
		fmt.Println("err = ", err)
		return *result
	}

	for _, fi := range list {
		currentFile := fi.Name()
		currentFile = fileName + string(os.PathSeparator) + fi.Name()

		if InArray(fi.Name(), noWatch) {
			continue
		}

		if fi.IsDir() {
			FindRecentModifyFile(currentFile, interval, timestamp, noWatch, result)
			continue
		}

		//当 interval 为 0 时，上传所有文件
		if timestamp-fi.ModTime().Unix() < interval || interval == 0 {
			*result = append(*result, currentFile)
		}
	}
	return *result
}

func GetRunPath() string {
	file, _ := exec.LookPath(os.Args[0])
	deployPath, _ := filepath.Abs(file)
	deployHome := filepath.Dir(deployPath)
	return deployHome +  string(os.PathSeparator)
}

func GetDefaultConfigPath() string {
	return GetRunPath() + "config" + string(os.PathSeparator) + "kael.ini"
}

func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return value, nil
}

func MD5Bytes(s []byte) string {
	ret := md5.Sum(s)
	return hex.EncodeToString(ret[:])
}

//计算字符串MD5值
func MD5(s string) string {
	return MD5Bytes([]byte(s))
}

//计算文件MD5值
func MD5File(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return MD5Bytes(data), nil
}
