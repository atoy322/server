package main

import (
    "fmt"
    "net/http"
    "os"
)

type MyHandler struct {
    Count int
}

func IsExists(name string) bool {
    _, err := os.Stat(name)
    if !os.IsNotExist(err) {
        return true
    } else {
        return false
    }
}

func TransportFile(w http.ResponseWriter, name string) error {
    f, err := os.Open(name)
    if err != nil {
        return err
    }
    defer f.Close()

    buf := make([]byte, 1024)
    num := 0
    for {
        num, err = f.Read(buf)
        if (err != nil) || (num == 0) {
            break
        }
        w.Write(buf[:num])
    }
    return nil
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    name := "." + r.URL.String()

    if name == "./" {
        name = "./index.html"
    }

    exists := IsExists(name)
    if exists {
        w.WriteHeader(200)
        TransportFile(w, name)
        fmt.Printf("%3d: [OK] %s\n", h.Count, name)
    } else {
        w.WriteHeader(404)
        if IsExists("notfound.html") {
            TransportFile(w, "notfound.html")
        } else {
            w.Write([]byte("NOT FOUND"))
        }
        fmt.Printf("%3d: [NOT FOUND] %s\n", h.Count, name)
    }
    h.Count += 1
}

func main() {
    h := &MyHandler{Count: 0}
    fmt.Println("Start")

    srv := &http.Server{
        Addr: "0.0.0.0:8000",
        Handler: h,
    }

    srv.ListenAndServe()
}
