package main


import ("flag"; "fmt")

func main() {
    flag.Parse()

    if flag.NArg() == 0 {
        fmt.Println("本校验程序需要一个文件参数，用法： CheckFile 文件名 ")
        return
    }

    fmt.Println("校验数据结束")
}
