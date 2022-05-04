package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// TODO is it worth using ast instead of strings, even with go fmt regularity?

func main() {
	if len(os.Args) == 1 || (len(os.Args) > 1 && (os.Args[0] == "-h" || os.Args[0] == "--help" || os.Args[0] == "help")) {
		fmt.Println("\nlinter to check that all switch statements have a default case")
		fmt.Println("\nusage: go-hasdefault $(find -type f -name '*.go')")
		os.Exit(1)
	}
	fail := false
	for _, filePath := range os.Args {
		if strings.HasSuffix(filePath, ".go") {
			var stdout bytes.Buffer
			cmd := exec.Command("gofmt", filePath)
			cmd.Stdout = &stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("fatal: gofmt failed on:", filePath)
				os.Exit(1)
			}
			switchesWithoutDefault := make(map[int]int) // indentLevel:lineNum
			lines := strings.Split(stdout.String(), "\n")
			for lineNum, line := range lines {
				token := strings.Split(strings.Trim(line, "\t"), " ")[0]
				indentLevel := 0
				for _, char := range line {
					if char != '\t' {
						break
					}
					indentLevel++
				}
				switch token {
				case "switch":
					switchesWithoutDefault[indentLevel] = lineNum
				case "default:":
					delete(switchesWithoutDefault, indentLevel)
				}
			}
			if len(switchesWithoutDefault) != 0 {
				fail = true
				for _, lineNum := range switchesWithoutDefault {
					fmt.Println(filePath + ":" + fmt.Sprint(lineNum) + ": switch statement missing default case")
				}
			}
		}
	}
	if fail {
		os.Exit(1)
	}
}
