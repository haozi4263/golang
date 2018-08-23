package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DockerLog struct {
	Log       string `json:"log"`
	Stream    string `json:"stream"`
	Source    string `json:"__source__"`
	TimeStamp string `json:"@timestamp"`
	Topic     string `json:"__topic__"`
	Host      string `json:"host"`
	Pod       string `json:"k8s_pod"`
	Time      int64  `json:"__time__"`
	Container string `json:"docker_container"`
}

type DockerLogs []*DockerLog

func (a DockerLogs) Len() int           { return len(a) }
func (a DockerLogs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DockerLogs) Less(i, j int) bool { return a[i].Time <= a[j].Time }

type FileCtx struct {
	Target *os.File
	Name   string
}

func TestJsonParse(t *testing.T) {
	file1, err := os.Open("C:/Users/bingh/Desktop/log/im-proxy.txt0") // For read access.
	assert.NoError(t, err)
	defer file1.Close()

	file2, err := os.Open("C:/Users/bingh/Desktop/log/im-proxy.txt1") // For read access.
	assert.NoError(t, err)
	defer file2.Close()

	fileList := []*os.File{file1, file2}
	fileMap := make(map[string]*FileCtx)

	for _, v := range fileList {
		rd := bufio.NewReader(v)
		for {
			dl := new(DockerLog)
			line, err := rd.ReadBytes('\n') //以'\n'为结束符读入一行

			if err != nil || io.EOF == err {
				break
			}
			fmt.Println(string(line))
			err = json.Unmarshal(line, dl)
			assert.NoError(t, err)
			ctx, ok := fileMap[dl.Container]
			if ok {
				ctx.Target.WriteString(string(line))
			} else {
				target, err := os.OpenFile(dl.Container, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
				assert.NoError(t, err)
				fileMap[dl.Container] = &FileCtx{target, dl.Container}
				target.WriteString(string(line))
			}
		}
	}

	targetDir := "E:/gocode/src/test_leak/target/"

	for key, _ := range fileMap {
		file, err := os.Open(key)
		assert.NoError(t, err)
		defer file.Close()

		sortFile(t, file, targetDir+key)
	}

}

func sortFile(t *testing.T, source *os.File, targetDir string) {
	file, err := os.OpenFile(targetDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) // For read access.
	assert.NoError(t, err)
	defer file.Close()

	var lineList DockerLogs
	lineList = make([]*DockerLog, 0, 10000)
	rd := bufio.NewReader(source)
	for {
		dl := new(DockerLog)
		line, err := rd.ReadBytes('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		err = json.Unmarshal(line, dl)
		assert.NoError(t, err)
		lineList = append(lineList, dl)
	}

	fmt.Println(len(lineList))
	sort.Sort(lineList)
	for _, v := range lineList {
		file.WriteString(v.Log)
	}
}
