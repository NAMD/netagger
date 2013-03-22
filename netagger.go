package main

import (
    "bufio"
    "io"
    "fmt"
    "log"
    "os"
)

func load_data(path string) map[string]string {
    data := make(map[string]string)
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

        word := string(line)
        data[word] = word
    }
    return data
}

func main() {
    countries := load_data("KnownLists/en/known_country.lst")
    fmt.Println(countries["Brazil"])
}
