# INTM

CodeKata to implement an interval merge in golang.

## Usage

    go run ./cmd/intm -h
    Usage of /tmp/go-build3227434488/b001/exe/intm:
      -alg string
            used algorithm for merging (default "simple")
      -path string
            file to merge (default "examples/test.txt")

### Map Reduce / Re-reduce

With large data files multiple instances of the program can be started. Each instance
can handle a part of the problem, the final set of files can be concatenated and passed
ans input. This mechanism allows the usage or more compute and memory resources then
available on a single machine.

    go build -o intm ./cmd/intm
    ./intm -path examples/part1.txt >> /tmp/parts
    ./intm -path examples/part2.txt >> /tmp/parts
    ./intm -path /tmp/parts

## Problem

Find all overlapping intervals and return a list of the consolidated overlapping intervals.

## Assumptions

* Input list of intervals is not sorted
* The Input list can contain invalid records, these should be skipped
  * `Start == End` -> is understood as error as the interval size (`End - Start`) would be 0
  * `End > Start` -> is understood as error, that could be auto-corrected depending on the input

## Architecture

The implementation follows the hexagonal architecture pattern.

The `cmd/intm/main.go` integrates all dependencies. The main starts reading the input file
and parses the individual intervals. All intervals are passed to the merging function.
After all intervals were processed by the merger the result is returned on stdout.

### Merger

* `noop` implements a simple none merge implementation
* `simple` sorts the list and uses a simple extension loop
* `btree` sorts while merging in order to save memory

### Reader

* `reader` currently a simple line format based reader is implemented
