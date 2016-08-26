package main

import (
    "archive/zip"
    "flag"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    rawUrl := flag.String("url", "", "download zip file")
    flag.Parse()
    log.Println("Downloading ", *rawUrl)

    Download(*rawUrl)
    destPath := Unzip("/tmp/downloaded.zip", "/")
    log.Println("Unzipped ", *destPath)
    os.Symlink(*destPath, "/data")

    http.Handle("/", http.FileServer(http.Dir("/data")))

    log.Println("Serving /data on HTTP port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func Download(rawUrl string) {
    os.MkdirAll("/tmp", 0755)
    fileName := "/tmp/downloaded.zip"

    file, err := os.Create(fileName)
    if err != nil {
     log.Println(err)
     panic(err)
    }
    defer file.Close()

    check := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }

    resp, err := check.Get(rawUrl)
    if err != nil {
        log.Println(err)
        panic(err)
    }
    defer resp.Body.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        panic(err)
    }
}

func Unzip(src, dest string) *string {
    r, err := zip.OpenReader(src)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := r.Close(); err != nil {
            panic(err)
        }
    }()

    // Closure to address file descriptors issue with all the deferred .Close() methods
    extractAndWriteFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer func() {
            if err := rc.Close(); err != nil {
                panic(err)
            }
        }()

        path := filepath.Join(dest, f.Name)

        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer func() {
                if err := f.Close(); err != nil {
                    panic(err)
                }
            }()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }

    for _, f := range r.File {
        err := extractAndWriteFile(f)
        if err != nil {
            panic(err)
        }
    }

    result := filepath.Join(dest, r.File[0].Name)
    return *result
}
