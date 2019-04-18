package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		input   string
		verbose bool
	)

	flag.StringVar(&input, "i", "main.go", "set input file")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	f, err := os.OpenFile(input, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	}()

	out := strings.Replace(input, filepath.Ext(input), ".pgo", 1)

	buf := bytes.NewBuffer([]byte{})

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		line := scan.Text()
		def := strings.TrimSpace(line)
		if strings.HasPrefix(def, "//#") {
			def = strings.Replace(def, "//", "", 1)
			buf.WriteString(def + "\n")
			if verbose {
				fmt.Println(def)
			}
		} else {
			buf.WriteString(line + "\n")
			if verbose {
				fmt.Println(line)
			}
		}
	}

	fmt.Println(out)

	if err := ioutil.WriteFile(out, buf.Bytes(), os.ModePerm); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

}
