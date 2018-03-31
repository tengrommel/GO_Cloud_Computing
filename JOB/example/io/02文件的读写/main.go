package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func WriteFile(path string)  {
	// 打开文件
	f, err := os.Create(path)
	if err != nil{
		fmt.Println("err = ", err)
		return
	}
	// 使用完毕，需要关闭文件
	defer f.Close()

	var buf string

	for i:= 0; i<10; i++{
		// "i = 1\n", 这个字符串存储在buf中
		buf = fmt.Sprintf("i = %d\n", i)
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("n = ", n)
	}
}

func ReadFile(path string)  {
	// 打开文件
	f, err := os.Open(path)
	if err != nil{
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer f.Close()
	buf := make([]byte, 2048)
	// n 代表内容读取的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF{
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}
// 每次读取一行

func ReadFileL(path string)  {
	// 打开文件ine
	f, err := os.Open(path)
	if err != nil{
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 新建一个缓冲区，把内容先放在缓冲去
	r := bufio.NewReader(f)
	for {
		// 遇到'\n'结束读取,但是`\n`也读取进入
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF{
				break
			}else {
				fmt.Println(err)
			}
			fmt.Printf("buf = #%s#\n", string(buf))
		}
	}
}

func main() {
	path := "./demo.txt"
	//WriteFile(path)
	ReadFile(path)
}
