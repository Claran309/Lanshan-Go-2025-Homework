package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
)

type Result struct {
	Path string
	Info []LineInfo
	Err  error
}

type LineInfo struct {
	Line    int
	Content string
}

type Task struct {
	Path    string
	Keyword string
}

func Search(task Task) (interface{}, error) {
	// 打开文件
	file, err := os.Open(task.Path)
	if err != nil {
		return Result{
			Path: task.Path,
			Err:  fmt.Errorf("无法打开文件: %v\n", err),
		}, err
	}
	defer file.Close()

	var info []LineInfo
	scanner := bufio.NewScanner(file)
	cnt := 0

	// 持续读入
	for scanner.Scan() {
		cnt++
		line := scanner.Text()

		if strings.Contains(line, task.Keyword) {
			info = append(info, LineInfo{Line: cnt, Content: line})
		}
	}

	if err := scanner.Err(); err != nil {
		return Result{
			Path: task.Path,
			Err:  fmt.Errorf("读取文件错误: %v", err),
		}, err
	}

	return Result{
		Path: task.Path,
		Info: info,
	}, nil
}

// 统计信息
var (
	total int64
	found int64
	lines int64
)

func SetTotal(num int64) {
	atomic.StoreInt64(&total, num)
}

func AddFound(num int64) {
	atomic.AddInt64(&found, num)
}

func AddLines(num int64) {
	atomic.AddInt64(&lines, num)
}

func GetInfo() (int, int, int) {
	return int(atomic.LoadInt64(&total)), int(atomic.LoadInt64(&found)), int(atomic.LoadInt64(&lines))
}
