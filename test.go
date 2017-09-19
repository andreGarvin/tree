package main

import "fmt"

import "./allPaths"

func main() {

    paths, err := allPaths.All("../../")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(paths)
    }

    paths, err := allPaths.Dirs("../")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(paths)
    }

    paths, err := allPaths.WithExt("../GNU/src", "go")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(paths)
    }

}
