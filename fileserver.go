// © 2013 QuanQinle. Licensed under the MIT license.


/*
http文件服务器
通过浏览器查看文件内容，输入形如 http://127.0.0.1:8080/D:/chelloworld.TXT
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    var (
        servHost := "127.0.0.1"
        servPort := ":8080"
        rootPath := "."
        )

    if len(os.Args) != 2 {
        _, filename := filepath.Split(os.Args[0])
        fmt.Println("USAGE:\n\t", filename, "files-server-root-path")
        os.Exit(1)
    }

    fmt.Println("请用浏览器访问：", servHost, servPort)

    // deliver files from the directory /var/www
    var root http.FileSystem
    rootPath = os.Args[1]
    root = http.Dir(rootPath)

    fileServer := http.FileServer(root)
    // register the handler and deliver requests to it
    err1 := http.ListenAndServe(servPort, fileServer)
    if err1 != nil {
        log.Fatal(err1)
    }
}

/*
//另一种实现方法（非原创）
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func showLocalFile(w http.ResponseWriter, r *http.Request) {
    contents, err := ioutil.ReadFile(r.URL.Path[1:])
    if err != nil {
        fmt.Fprintf(w, "404\n请检查url参数是否是正确的文件路径！")
        return
    }
    fmt.Fprintf(w, "%s!\n", contents)
}

func main() {
    http.HandleFunc("/", showLocalFile)
    http.ListenAndServe(":8080", nil)
}
*/
