package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/threez/intm/internal/adapter/merger"
	"github.com/threez/intm/internal/adapter/reader"
	"github.com/threez/intm/internal/port"
)

var (
	path string
	alg  string
)

func main() {
	flag.StringVar(&path, "path", "examples/test.txt", "file to merge")
	flag.StringVar(&alg, "alg", "simple", "used algorithm for merging")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open input file %q: %v", path, err)
		os.Exit(1)
	}
	defer f.Close()

	var m port.Merger
	switch alg {
	case "noop":
		m = &merger.Noop{}
	case "simple":
		m = merger.NewSimple()
	case "btree":
		m = merger.NewBtree()
	default:
		fmt.Fprintf(os.Stderr, "unknown merge algorithm %q", alg)
		os.Exit(1)
	}

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
		err = i.Validate()
		if err != nil {
			fmt.Fprintf(os.Stderr, "INVALID %v\n", err)
			continue
		}

		m.MergeInterval(i)
	}

	// simple writer
	for _, i := range m.Result() {
		fmt.Println(i.String())
	}
}
