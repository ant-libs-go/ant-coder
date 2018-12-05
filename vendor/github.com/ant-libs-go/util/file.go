/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 04:58:45
# File Name: file.go
# Description:
####################################################################### */

package util

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func PathExists(path string) (exists bool, isdir bool, err error) {
	f, err := os.Stat(path)
	if err == nil {
		return true, f.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, false, nil
	}
	return false, false, err
}

// 遍历目录下所有指定后缀的文件
func WalkDir(path string, suffixes []string) (files []string, err error) {
	for k, suffix := range suffixes {
		suffixes[k] = strings.ToUpper(suffix)
	}

	err = filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}

		if len(suffixes) == 0 {
			files = append(files, filename)
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), fmt.Sprintf(".%s", suffix)) {
				files = append(files, filename)
			}
		}
		return nil
	})
	return files, err
}

//MD5
func Md5File(file string) (string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}

	h := md5.New()
	_, err = io.Copy(h, bufio.NewReader(f))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//SHA1
func SHA1File(file string) (string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}

	h := sha1.New()
	_, err = io.Copy(h, bufio.NewReader(f))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//SHA256
func SHA256File(file string) (string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}

	h := sha256.New()
	_, err = io.Copy(h, bufio.NewReader(f))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func ReadFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(data string, file string) error {
	return ioutil.WriteFile(file, []byte(data), 0666)
}
