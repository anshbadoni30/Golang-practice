package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")
	if err!= nil{
		panic(err)
	}
	defer f.Close()
	fileinfo, err:= f.Stat()
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.ModTime())

	//Reading a File
	f1,errr := os.ReadFile("hello.txt")
	if errr!= nil{
		panic(errr)
	}
	fmt.Println(string(f1)) 

	//Reading a folder
	dir, e2:=os.Open(".")
	if e2!= nil{
		panic(e2)
	}
	defer dir.Close()
	dirinfo , _ := dir.ReadDir(-1) 
	for _,a:=range dirinfo{
		fmt.Println(a)
	}

	// Creating a File
	f3, err:=os.Create("example.txt")
	if err!=nil{
		panic(err)
	}
	defer f3.Close()
	//1st method
	//f3.WriteString("Hello guys")
	//f3.WriteString("yoo bitch")

	//2nd method
	bytes:= []byte("Hello world")
	f3.Write(bytes)

	//Reading a file and write content in another file
	readfile, er:=os.Open("example.txt")
	if er!=nil{
		panic(er)
	}
	defer readfile.Close()

	writefile, err:= os.Create("example2.txt")
	if err!=nil{
		panic(err)
	}
	defer writefile.Close()
	reader := bufio.NewReader(readfile)
	writer :=bufio.NewWriter(writefile)

	for{
		f4,l:=reader.ReadByte()
		if l!=nil{
			if l.Error()!="EOF"{
				panic(l)
			}
			break
		}

		em:= writer.WriteByte(f4)
		if em!=nil{
			panic(em)
		}
	}
	writer.Flush()

	//deleting file
	er=os.Remove("example.txt")
	if er!=nil{
		panic(er)
	}
	fmt.Println("file removed succesfully")


}