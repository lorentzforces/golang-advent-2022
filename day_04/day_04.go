package day_04

import (
	"errors"
	"fmt"
	"local-advent-2022/util"
	"strconv"
	"strings"
)

func PartOne(input string) uint {
	lines := util.AsLines(input)
	count := uint(0)
	for _, line := range lines {
		spans, err := parseSpanPair(line)
		if err != nil {
			panic(err)
		}

		if spansAreContained(spans) {
			count += 1
		}
	}

	return count
}

func PartTwo(input string) uint {
	lines := util.AsLines(input)
	count := uint(0)
	for _, line := range lines {
		spans, err := parseSpanPair(line)
		if err != nil {
			panic(err)
		}

		if overlap(spans) {
			count += 1
		}
	}


	return count
}

type spanPair struct {
	left [2]uint
	right [2]uint
}

func spansAreContained(spans spanPair) bool {
	return contains(spans.left, spans.right) || contains(spans.right, spans.left)
}

func contains(left, right [2]uint) bool {
	return left[0] <= right[0] && right[1] <= left[1]
}

// spans overlap if either endpoint of a span is contained in the other span
func overlap(spans spanPair) bool {
	return (spans.left[0] <= spans.right[0] && spans.right[0] <= spans.left[1]) ||
		(spans.left[0] <= spans.right[1] && spans.right[1] <= spans.left[1]) ||
		(spans.right[0] <= spans.left[0] && spans.left[0] <= spans.right[1]) ||
		(spans.right[0] <= spans.left[1] && spans.left[1] <= spans.right[1])
}

func parseSpanPair(s string) (spanPair, error) {
	strSpans := strings.Split(s, ",")
	pair := spanPair {
		left: [2]uint{0, 0},
		right: [2]uint{0, 0},
	}
	if len(strSpans) != 2 {
		err := errors.New(fmt.Sprintf(
			"Expected exactly two values for a parsed span pair from \"%s\", but saw %d values",
			s,
			len(strSpans),
		))
		return pair, err
	}

	left, err := parseSpan(strSpans[0])
	if err != nil {
		return pair, err
	}
	right, err := parseSpan(strSpans[1])
	if err != nil {
		return pair, err
	}

	pair.left = left
	pair.right = right
	return pair, nil
}

func parseSpan(s string) ([2]uint, error) {
	span := [2]uint{0, 0}
	numbers := strings.Split(s, "-")
	if len(numbers) != 2 {
		err := errors.New(fmt.Sprintf(
			"Expected exactly two values for parsed numbers from span \"%s\", but saw %d values",
			s,
			len(numbers),
		))
		return span, err
	}

	for i, valStr := range numbers {
		val, err := strconv.ParseUint(valStr, 10, 0)
		if err != nil {
			return span, err
		}
		span[i] = uint(val)
	}

	// I think all the inputs are in sorted order, but it's not stated explicitly so...
	if span[0] > span[1] {
		temp := span[0]
		span[0] = span[1]
		span[1] = temp
	}

	return span, nil
}
