package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"time"
)

var (
	//最多支持20个并发
	sema = make(chan struct{}, 20)
	//用来判断是否终止创建go runtime，并且能够做到优雅的退出
	done = make(chan struct{})
	//接收ctrl+c的终止信号
	stop = make(chan os.Signal, 1)
)

//创建一个监视对象
type Monitor struct {
	Dir  string
	Size int64
	Time time.Duration
}

var (
	d = flag.String("d", ".", "要监控的目录")
	s = flag.Int64("s", 200, "默认最小值为200M")
	t = flag.Duration("t", 1*time.Second, "间隔多久时间遍历目录大小")
)

var usage = `Usage: monitor [option...]

option:
    -d    要监控的目录,默认为当前目录
    -s    默认最小值为200M，超过200M，会执行删除操作
    -t    间隔多久时间遍历目录大小，默认值为1s
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()

	//指定捕捉ctrl+c信号
	signal.Notify(stop, os.Interrupt)

	(&Monitor{
		Dir:  *d,
		Size: *s,
		Time: *t,
	}).Run()

}

func (m *Monitor) Run() {
	tick := time.Tick(m.Time)

	HandleDir(m.Dir, m.Size)
loop:
	for {
		select {
		case <-done:
			GetDirSize(m.Dir) //等待返回，不做任何操作
			break loop
		case <-tick:
			HandleDir(m.Dir, m.Size)
		case <-stop:
			fmt.Println("good bye\n")
			close(done)
		}
	}

}

//处理目录
func HandleDir(dir string, size int64) {
	sum := float64(GetDirSize(dir)) / 1e6
	if sum >= float64(size) {
		fmt.Fprintf(os.Stdout, "该目录大小：%f,大于%d MB了，准备删除\n", sum, size)
		Remove(dir)
	} else {
		fmt.Fprintf(os.Stdout, "目录大小为%f MB\n", sum)
	}
}

//删除对应目录下的文件
func Remove(dir string) {
	for _, file := range GetFileList(dir) {
		path := filepath.Join(dir, file.Name())
		if err := os.RemoveAll(path); err != nil {
			fmt.Fprintf(os.Stderr, "删除文件夹%s内容失败，原因是:%v\n", dir, err)
		}
	}
}

//获取目录大小
func GetDirSize(dir string) int64 {
	var (
		sum int64
		w   sync.WaitGroup
	)

	filesizes := make(chan int64)

	w.Add(1)
	go WalkDir(dir, filesizes, &w)

	go func() {
		w.Wait()
		close(filesizes)
	}()

loop:
	for {
		select {
		//判断是否是被中断了，如果被中断，则优雅的等待channel循环结束
		case <-done:
			for range filesizes {
			}
			break loop
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			sum += size
		}
	}

	return sum
}

//迭代获取目录中文件的大小，写入channel filesizes中
func WalkDir(dir string, filesizes chan<- int64, w *sync.WaitGroup) {
	defer w.Done()
	for _, file := range GetFileList(dir) {
		if file.IsDir() {
			dirpath := filepath.Join(dir, file.Name())
			w.Add(1)
			go WalkDir(dirpath, filesizes, w)
		} else {
			filesizes <- file.Size()
		}
	}
}

//获取目录的文件列表
func GetFileList(dir string) []os.FileInfo {
	select {
	//增加一个判断，判断是否是要中断结束
	case <-done:
		return nil
	default:
		sema <- struct{}{}
	}
	defer func() {
		<-sema
	}()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "获取文件列表失败: %v\n", err)
		return nil
	}
	return files
}
