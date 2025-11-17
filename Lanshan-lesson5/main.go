// 8个小时拼尽全力无法战胜（划）
package main

import (
	"Lesson_1/Lanshan-lesson5/service"
	"Lesson_1/Lanshan-lesson5/workerPool"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("食用方法: %s [目录路径] [搜索关键词]\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]
	keyword := os.Args[2]

	//验证目录是否合法
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("目录 '%s' 不存在\n", dir)
		os.Exit(1)
	}

	workCount := runtime.NumCPU() * 2
	fmt.Printf("检索到目录：'%s'，开始搜索关键词 '%s'\n", dir, keyword)

	startTime := time.Now()

	//初始化协程池
	pool := workerPool.NewWorkerPool(workCount, workCount*100)

	//遍历目录，储存路径
	paths, err := rangeFiles(dir)
	if err != nil {
		fmt.Println("遍历目录时出错: %v\n", err)
		os.Exit(1)
	}

	service.SetTotal(int64(len(paths)))
	fmt.Printf("文件数: %d\n", len(paths))

	// 收集结果
	results := make(chan workerPool.TaskResult, workCount*100)
	done := make(chan bool, 1)
	go GetResults(results, done, len(paths))

	finished := 0
	for _, path := range paths {
		//执行任务
		task := service.Task{path, keyword}
		err := pool.Produce(func() (interface{}, error) {
			return service.Search(task)
		})

		if err != nil {
			fmt.Printf("produced failed: %v\n", err)
		} else {
			finished++
		}
	}

	fmt.Printf("任务数: %d\n", finished)

	// 收集channel
	go TaskResults(pool.GetResults(), results)

	pool.Close()
	//time.Sleep(100 * time.Millisecond)                Debug
	close(results)

	<-done

	total, found, lines := service.GetInfo()
	times := time.Since(startTime)

	fmt.Printf("\n============= 搜索完成 =============\n")
	fmt.Printf("总文件数: %d\n", total)
	fmt.Printf("包含关键词的文件数: %d\n", found)
	fmt.Printf("总匹配行数: %d\n", lines)
	fmt.Printf("耗时: %v\n", times)

	proSum, conSum := pool.GetInfo()
	fmt.Printf("任务提交数: %v\n", proSum)
	fmt.Printf("任务完成数: %v\n", conSum)
}

func rangeFiles(dir string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("无法访问 '%s' : %v\n", path, err)
			return nil
		}

		if !d.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func TaskResults(taskResults <-chan workerPool.TaskResult, results chan<- workerPool.TaskResult) {
	for taskResult := range taskResults {
		results <- taskResult
	}
}

func GetResults(results <-chan workerPool.TaskResult, done chan<- bool, total int) {
	var finalResult []service.Result
	finished := 0

	for result := range results {
		finished++

		if finished%100 == 0 {
			fmt.Printf("处理进度: %d/%d (%.2f%%)\n", finished, total, float64(finished)/float64(total))
		}

		// 断言结果类型
		if searchResult, ok := result.Result.(service.Result); ok {
			finalResult = append(finalResult, searchResult)

			if len(searchResult.Info) > 0 {
				service.AddFound(1)
				service.AddLines(int64(len(searchResult.Info)))
			}
		}
	}

	fmt.Printf("\n================搜索结果================\n")
	printAll(finalResult)
	done <- true
}

func printAll(result []service.Result) {
	sort.Slice(result, func(i, j int) bool {
		return result[i].Path < result[j].Path
	})

	for _, result := range result {
		if result.Err != nil {
			fmt.Printf("\n错误: %s - %v\n", result.Path, result.Err)
			continue
		}

		if len(result.Info) > 0 {
			fmt.Printf("\n%s:\n", result.Path)
			for _, info := range result.Info {
				fmt.Printf("    %d: %s\n", info.Line, info.Content)
			}
		}
	}
}
