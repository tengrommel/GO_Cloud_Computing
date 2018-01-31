package main

import (
	"os"
	"log"
	"io"
	"bufio"
	"fmt"
)

func main() {
	f1, err := os.Open("test1.txt")
	PrintFatalError(err)
	defer f1.Close()

	f2, err := os.Create("text2.txt")
	PrintFatalError(err)
	defer f2.Close()

	f3, err := os.OpenFile("text3.txt", os.O_APPEND|os.O_RDWR, 0666)
	PrintFatalError(err)
	defer f3.Close()

	err = os.Rename("test1.txt","test1New.txt")
	PrintFatalError(err)

	err = os.Rename("./test1.txt", "./testfolder/test1.txt")
	PrintFatalError(err)

	CopyFile("test3.txt", "./testfolder/test3.txt")

	scanner := bufio.NewScanner(f3)
	count := 0
	for scanner.Scan(){
		count ++
		fmt.Println("Found line:", count, scanner.Text())
	}

	writebuffer := bufio.NewWriter(f3)
	for i:=1;i<5;i++ {
		writebuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	writebuffer.Flush()
}

func CopyFile(fname1, fname2 string) {
	fOld, err := os.Open(fname1)
	PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	_, err = io.Copy(fNew, fOld)
	PrintFatalError(err)

	err = fNew.Sync()
	PrintFatalError(err)
}

func PrintFatalError(e error) {
	if e != nil{
		log.Fatal("Error happened while processing file", e)
	}
}


