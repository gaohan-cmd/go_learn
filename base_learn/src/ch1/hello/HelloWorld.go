package main

// 程序入口 main包 pkg和文件夹名字可以不一致，main函数不支持任何返回值
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World!", os.Args[1])
	}
	os.Exit(0)

}
