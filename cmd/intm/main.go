package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/threez/intm/internal/adapter/merger"
	"github.com/threez/intm/internal/adapter/reader"
)

var path string

func main() {
	flag.StringVar(&path, "path", "examples/test.txt", "file to merge")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open input file %q: %v", path, err)
		os.Exit(1)
	}
	defer f.Close()

	m := &merger.Noop{}
	r := reader.NewReader(f)
	for {
		i, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "SKIP %v\n", err)
			continue
		}

		m.MergeInterval(i)
	}

	// simple writer
	for _, i := range m.Result() {
		fmt.Println(i.String())
	}
}
