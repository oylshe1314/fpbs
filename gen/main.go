package main

import (
	"fmt"
	"fpbs"
	"fpbs/gen/parser/json"
	"fpbs/gen/writer/csharp"
	"fpbs/gen/writer/golang"
	"fpbs/util"
	"os"
	"strings"
)

type _Struct struct {
	pack   string
	name   string
	orders []string
	fields map[string]*util.Pair[fpbs.FieldKey, fpbs.FieldType]
}

func main() {
	goDir, csDir, incDirs, files := parseArgs()
	incDirs = append(incDirs, "./")

	var parser = json.NewParser()

	packages, err := parser.ParseFiles(files, incDirs)
	if err != nil {
		fmt.Println(err)
		return
	}

	if goDir != "" {
		var goWriter = golang.NewWriter()
		err = goWriter.Write(goDir, packages)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if csDir != "" {
		var csWriter = csharp.NewWriter()
		err = csWriter.Write(csDir, packages)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func parseArgs() (goDir, csDir string, incDirs []string, files []string) {
	if len(os.Args) < 2 {
		fmt.Println("Missing input file.")
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		var arg = os.Args[i]

		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-go":
				if i == len(os.Args)-1 || strings.HasPrefix(os.Args[i+1], "-") {
					fmt.Println("Missing parameter for option '-go'")
					os.Exit(1)
				}
				goDir = os.Args[i+1]
				i++
			case "-cs":
				if i == len(os.Args)-1 || strings.HasPrefix(os.Args[i+1], "-") {
					fmt.Println("Missing parameter for option '-cs'")
					os.Exit(1)
				}
				csDir = os.Args[i+1]
				i++
			case "-I":
				if i == len(os.Args)-1 || strings.HasPrefix(os.Args[i+1], "-") {
					fmt.Println("Missing parameter for option '-I'")
					os.Exit(1)
				}

				if strings.HasSuffix(os.Args[i+1], "/") && strings.HasSuffix(os.Args[i+1], "\\") {
					if os.Args[i+1] == "./" || os.Args[i+1] == ".\\" {
						continue
					}
					incDirs = append(incDirs, os.Args[i+1])
				} else {
					if os.Args[i+1] == "." {
						continue
					}
					incDirs = append(incDirs, os.Args[i+1]+"/")
				}
				i++
			}
		} else {
			files = append(files, arg)
		}
	}
	if goDir == "" && csDir == "" {
		fmt.Println("Missing output directives.")
		os.Exit(1)
	}
	if len(files) == 0 {
		fmt.Println("Missing input file.")
		os.Exit(1)
	}
	return
}
