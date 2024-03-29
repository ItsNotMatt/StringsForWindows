package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

func main() {
    dir, err := filepath.Abs(filepath.Dir("."))
    if err != nil {
        log.Fatal("Error getting current dir")
    }

    path, cont := parseFirstArg(os.Args[1])
    if !cont {
        return
    }
    path = filepath.Join(dir, path)
    

    var option string
    var additionalOptions bool
    var dumpFile *os.File
    var dump bool
    if len(os.Args) > 2 {
        option, additionalOptions = parseOptions(os.Args[2])
        if !additionalOptions {
            dumpFile, dump = parseSecondArg(dir, os.Args[2]) 
            option, additionalOptions = parseOptions(os.Args[4])
        }
    }

    content, err := os.ReadFile(path)
    if err != nil {
        log.Fatal("Error reading file", err)
    }
    strings := parseFile(string(content), true)

    if additionalOptions {
        if option == "-f" {
            strings = formatStrings(strings)
        }
    }


    if dump {
        for _, str := range strings {
            _, err := dumpFile.WriteString(str + "\n")
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

func parseFile(content string, writeToFile bool) []string {
    chars := []rune(content)
    var tempChars []rune
    var strings []string

    validSequenceCount := 0
    addToStrings := false

    for _, c := range chars {
        if unicode.IsSpace(c) {
            continue
        }
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
        } else {
            tempChars = nil
            validSequenceCount = 0
        }
    }
    return strings
}

func formatStrings(strings []string) []string {
    hmap := make(map[string]int)
    var newStrings []string

    for _, str := range strings {
        _, ok := hmap[str]
        if ok {
            hmap[str]++
        } else {
            hmap[str] = 1
        }
    }

    for key, value := range hmap {
        num := strconv.Itoa(value)
        str := key + " " + num
        newStrings = append(newStrings, str)
    }
    
    return newStrings
}


func parseFirstArg(arg string) (string, bool) {
    if arg == "help" {
        fmt.Println("Commands: ")
        return "", false
    } else {
        //assuming its a filename now
        return arg, true
    }
}

func parseSecondArg(dir string, arg string) (*os.File, bool) {
    if arg == ">" {
        filepath := filepath.Join(dir, os.Args[3])
        file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
        if err != nil {
            log.Fatal("Error opening file", err)
        }
        defer file.Close()
        return file, true
    } else {
        return nil, false
    }
}

func parseOptions(arg string) (string, bool) {
    if arg == "-f" {
        return arg, true
    } else {
        return "", false
    }
}



