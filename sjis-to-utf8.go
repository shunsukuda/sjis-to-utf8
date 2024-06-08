package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: sjis-to-utf8 <file1> <file2> ...")
		return
	}
	args := os.Args[1:]
	for i := range args {
		sjisFile, err := os.Open(args[i])
		if err != nil {
			log.Fatal(err)
		}
		defer sjisFile.Close()

		lidx := strings.LastIndexByte(args[i], '.')
		fname := args[i][:lidx]
		ext := args[i][lidx:]
		reader := transform.NewReader(sjisFile, japanese.ShiftJIS.NewDecoder())
		utf8File, err := os.Create(fname + "_utf8" + ext)
		if err != nil {
			log.Fatal(err)
		}
		defer utf8File.Close()

		tee := io.TeeReader(reader, utf8File)
		s := bufio.NewScanner(tee)
		for s.Scan() {
		}
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		log.Println("Transform:" + args[i] + " -> " + fname + "_utf8" + ext)
	}

}
