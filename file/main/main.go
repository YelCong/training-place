package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
)

var filepath string = "/tmp/t"

func main() {
	//readOnce()
	//readWithBuf()
	//writeWithBuf()
	//writeOnce()
	//fileCopy()
	//fileExist()
	//fileCopyWithBuf()
	//args()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}

func args()  {
	fmt.Println(os.Args)
	var host string
	var port int64
	flag.StringVar(&host, "h", "127.0.0.1", "usage string")
	flag.Int64Var(&port,"p",3306,"set port")
	flag.Parse()
	fmt.Println(host,port)
}

func readOnce() {
	filepath := "/tmp/t"
	//隐式的完成open.close操作 适用于小文件
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(fileContent))
}

//适用于大文件
func readWithBuf() {
	filepath := "/tmp/t"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("file close err", err) //如果已经关闭了则显示  file already closed
			return
		}
	}()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //表示结束
			break
		}
		fmt.Print(str)
	}
	fmt.Println("ending")
}

func writeWithBuf() {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("os.OpenFile Err", err)
	}
	defer func() {
		if err1 := file.Close(); err1 != nil {
			//err1的变量作用域范围就在这里，出了这里都不能用
			fmt.Println("file close err", err1)
			return
		}
	}()
	//先读
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //表示结束
			break
		}
		fmt.Print(str)
	}
	//后写
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		j, _ := writer.WriteString("f5\n")
		fmt.Println(j)
	}
	if err := writer.Flush(); err != nil {
		fmt.Println("file flush err", err)
	}
}

func writeOnce() {
	filepath := "/tmp/t2"
	data := []byte("阿里巴巴\n")
	err := ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ending")
}

func fileCopy() {
	from := "/tmp/t"
	to := "/tmp/t3"
	data, err := ioutil.ReadFile(from)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := ioutil.WriteFile(to, data, 0666); err != nil {
		fmt.Println(err)
		return
	}
}

func fileExist() {
	filepath += "aasd"
	fmt.Println(filepath)
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileInfo)
}

func fileCopyWithBuf() {
	dst := "/tmp/t4"
	src := "/tmp/t"

	file, err := os.Open(src)
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	file2, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file2.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file2)

	i, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
