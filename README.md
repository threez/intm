# INTM

CodeKata to implement an interval merge in golang.

## Usage

    go run ./cmd/intm

## Problem

Find all overlapping intervals and return a list of the consolidated overlapping intervals.

## Assumptions

* Input list of intervals is not sorted
* The Input list can contain invalid records, these should be skipped
  * `Start == End` -> is understood as error as the interval size (`End - Start`) would be 0
  * `End > Start` -> is understood as error, that could be autocorrected depending on the input
