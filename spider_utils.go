package raselper

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func ListFromNodes(sources []*html.Node, pattern string) []*html.Node {
	findNodes := make([]*html.Node, 10)

	for _, source := range sources {
		nodes := htmlquery.Find(source, pattern)
		for _, node := range nodes {
			findNodes = append(findNodes, node)
		}
	}

	return findNodes
}

func DownLoadFile(url string, targetPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("获取链接失败")
	}
	defer resp.Body.Close()

	_, fileName := filepath.Split(url)
	path := filepath.Join(targetPath, fileName)
	file, _ := os.Create(path)
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return errors.New("写入文件失败: " + err.Error())
	}

	return nil
}

func SaveUrlContent(url string, targetPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
