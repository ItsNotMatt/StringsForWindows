package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

func main() {
    dir, err := filepath.Abs(filepath.Dir("."))
    if err != nil {
        log.Fatal("Error getting current dir")
    }

    filepath := filepath.Join(dir, "test_exe", "strings_windows.exe")

    content, err := os.ReadFile(filepath)
    if err != nil {
        log.Fatal("Error reading file", err)
    }

    parseFile(string(content), true)
}

func parseFile(content string, writeToFile bool) {
    chars := []rune(content)
    var tempChars []rune
    var strings []string

    validSequenceCount := 0
    addToStrings := false

    for _, c := range chars {
        if unicode.IsLetter(c) {
            tempChars = append(tempChars, c)
            validSequenceCount++
            if validSequenceCount >= 4 {
                addToStrings = true
            }
        } else if addToStrings {
            strings = append(strings, string(tempChars))

            tempChars = nil
            validSequenceCount = 0
            addToStrings = false
        }
    }

    if writeToFile {
        dir, _ := filepath.Abs(filepath.Dir("."))
        filepath := filepath.Join(dir, "dump.txt")

        file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
        if err != nil {
            log.Fatal("Error opening file", err)
        }
        defer file.Close()

        for _, str := range strings {
            _, err := file.WriteString(str + "\n")
            if err != nil {
                log.Fatal("Err writing to file", err)
            }
        } 
    } else {
        for _, str := range strings {
            fmt.Println(str)
        }
    }

}
