package main

import (
    "bufio"
    "io"
    "fmt"
    "log"
    "os"
)

func load_data(path string) map[string]bool {
    data := make(map[string]bool)
    fp, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    bf := bufio.NewReader(fp)
    for {
        line, isPrefix, err := bf.ReadLine()

        if err == io.EOF {
            break
        }

        if isPrefix {
            log.Fatal("Line too long")
        }

        data[string(line)] = true
    }
    return data
}

func main() {
    countries := load_data("KnownLists/en/known_country.lst")
    fmt.Println(countries["Brazil"])
}
