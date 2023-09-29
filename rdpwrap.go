package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
)

func main() {
    // 设置路由
    http.HandleFunc("/", indexHandler)

    // 启动服务器
    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // 判断请求方式
    if r.Method == "GET" {
        // 判断请求路径是否为/rdpwrap.ini
        if r.URL.Path == "/rdpwrap.ini" {
            // 读取当前目录下的rdpwrap.ini文件内容
            content, err := ioutil.ReadFile("rdpwrap.ini")
            if err != nil {
                // 如果出错，返回404错误码和错误信息
                http.NotFound(w, r)
                return
            }
            // 将rdpwrap.ini文件内容渲染到模板中
            t, err := template.New("rdpwrap").Parse(string(content))
            if err != nil {
                // 如果出错，返回500错误码和错误信息
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            t.Execute(w, nil)
            return
        }
        // 如果路径不是/rdpwrap.ini，返回404错误码
        http.NotFound(w, r)
        return
    }
    // 如果请求方式不是GET，返回405错误码和错误信息
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

