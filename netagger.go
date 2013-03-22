package main

import (
    "bufio"
    "io"
    "io/ioutil"
    "fmt"
    "log"
    "strings"
    "os"
)

func open_file(path string) *os.File {
    fp, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    return fp
}

func load_data(path string) map[string]bool {
    data := make(map[string]bool)
    fp := open_file(path)
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
    fp.Close()
    return data
}

func get_text_from_file(path string) string {
    content, _ := ioutil.ReadFile(path)
    text := strings.Replace(string(content), "\n", "", -1)
    return text
}

func main() {
    countries := load_data("KnownLists/en/known_country.lst")

    text := get_text_from_file("input.txt")

    tokens := strings.Split(text, " ")

    for _, token := range tokens {
        if _, exists := countries[token]; exists {
            fmt.Println(token)
        }
    }
}
