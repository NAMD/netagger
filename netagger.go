package main

import (
    "bufio"
    "io"
    "io/ioutil"
    "fmt"
    "log"
    "path/filepath"
    "os"
    "strings"
)

func open_file(path string) *os.File {
    fp, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    return fp
}

func load_data(paths []string) map[string]bool {
    data := make(map[string]bool)
    for _, path := range paths {
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
    }
    return data
}

func get_text_from_file(path string) string {
    content, _ := ioutil.ReadFile(path)
    text := strings.TrimSpace(string(content))
    return text
}

func main() {
    entity_list_files, _ := filepath.Glob("KnownLists/en/*")

    entities := load_data(entity_list_files)

    text := get_text_from_file("input.txt")

    tokens := strings.Split(text, " ")

    for _, token := range tokens {
        if _, exists := entities[token]; exists {
            fmt.Println(token)
        }
    }
}
