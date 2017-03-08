package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/jlaffaye/ftp"
	"log"
	"sync"
)

var (
	ip   = flag.String("ip", "ftp 的地址", "use -ip <127.0.0.1>")
	user = flag.String("user", "ftp的用户名", "Use -user <admin>")
	pw   = flag.String("passwd", "ftp的密码", "Use -passwd <admin>")
	lp   = flag.String("local", "本地所要保存的目录", "Use -local <filepath>")
	rp   = flag.String("remote", "远程所要同步的目录", "Use -remote <filepath>")
)

type connPool struct {
	Dial func() (*ftp.ServerConn, error)
    maxActive int
	conn chan *ftp.ServerConn
}

func (p *connPool) initPool() {
	p.conn = make(chan *ftp.ServerConn, p.maxActive)
	for i := 0 ; i < p.maxActive; i++ {
		conn , err := p.Dial()
		if err != nil {
			log.Println("获取链接失败！")
			continue
		}
		p.conn <- conn
	}
}

func (p *connPool) Get() *ftp.ServerConn{

	if p.conn == nil {
		p.initPool()
	}
	return <- p.conn
}

func (p *connPool) release(conn *ftp.ServerConn) {
	p.conn <- conn
}

var p *connPool
var wg sync.WaitGroup

func main() {
	flag.Parse()
	
	p = newConnPool()
	wg.Add(1)
	go func() {
		listDir(*rp, *lp)
	}()
	wg.Wait()
	
}

func newConnPool() *connPool {
	return &connPool{
		maxActive: 20,
		Dial: func() (*ftp.ServerConn, error) {
			conn, err := ftp.Connect(*ip)
			if err != nil {
				return nil ,err
			}
			err = conn.Login(*user, *pw)
			if err != nil {
				return nil ,err
			}
			return conn, nil
		},
	}
}

//listDir :遍历目录，如果不是隐藏文件，且不ignoretxt文件中，则下载，如果是目录则创建目录，并迭代遍历
func listDir(_rpath, lpath string) {
	conn := p.Get()
	
	//创建或打开ignoretxt文件，获取里面的已有文件列表
	ig, err := NewIgnore(lpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	localList := ig.read()

	//需要保存到文件中新的文件的列表
	sList := make([]string, 0, 0)

	entries, err := conn.List(_rpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, et := range entries {
		//判断是否是目录
		if et.Type == ftp.EntryTypeFolder {
			//判断是否是已经在ignoretxt中的文件
			if !isInclude(et.Name, localList) {
				//判断是否是隐藏目录
				if isHideFile(et.Name) {
					continue
				}
				tmpDir := filepath.Join(lpath, et.Name)
				err := os.Mkdir(tmpDir, 0666)
				if err != nil {
					fmt.Println(err)
				}
				sList = append(sList, et.Name)
			}
			tmpDir := filepath.Join(lpath, et.Name)
			rDir := _rpath + "/" + et.Name
			wg.Add(1)
			go listDir(rDir, tmpDir)
		} else { //是文件执行下载操作
			if !isInclude(et.Name, localList) {
				sp := filepath.Join(lpath, et.Name)
				rf := _rpath + "/" + et.Name
				wg.Add(1)
				go download(rf, sp)
				sList = append(sList, et.Name)
			}
		}
	}
	ig.write(sList)
	ig.iClose()
	wg.Done()
	p.release(conn)
}

//isHideFile：是否是隐藏文件
func isHideFile(s string) bool {
	if strings.EqualFold(s, "..") || strings.EqualFold(s, ".") {
		return true
	}
	return false
}

//下载文件
func download(_path, savePath string) {
	defer wg.Done()
	conn := p.Get()
	fmt.Println("save: ", savePath)
	rc, err := conn.Retr(_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	w, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
		rc.Close()
		return
	}
	io.Copy(w, rc)
	w.Close()
	rc.Close()
	p.release(conn)
}

//判断字符串是否包含在字符串列表中
func isInclude(name string, list []string) bool {
	judge := false
	for _, v := range list {
		if strings.EqualFold(v, name) {
			judge = true
		}
	}
	return judge
}

type Ignore struct {
	f *os.File
}

func (i *Ignore) write(list []string) {
	var buf bytes.Buffer
	for _, v := range list {
		buf.WriteString(v)
		buf.WriteString("\n")
		i.f.WriteString(buf.String())
		buf.Reset()
	}
}

func (i *Ignore) read() []string {
	list := make([]string, 0, 0)
	buf := bufio.NewReader(i.f)
	for {
		bs, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		list = append(list, string(bs))
	}
	return list
}

func (i *Ignore) iClose() error {
	return i.f.Close()
}

//判断文件是否存在
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func NewIgnore(dp string) (*Ignore, error) {
	const ignoreName = ".ignoretxt"
	_path := filepath.Join(dp, ignoreName)
	if !PathExist(_path) {
		fp, err := os.OpenFile(_path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		return &Ignore{f: fp}, nil
	}
	fp, err := os.OpenFile(_path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Ignore{f: fp}, nil
}
