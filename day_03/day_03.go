package day_03

import (
	"errors"
	"fmt"
	"sort"
	"local-advent-2022/util"
)

func PartOne(input string) uint {
	prioritySum := uint(0)
	for _, line := range util.AsLines(input) {
		inv, err := parseInventory(line)
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(inv.both); i++ {
			prio, err := itemPriority(inv.both[i])
			if err != nil {
				panic(err)
			}
			prioritySum += prio
		}
	}

	return prioritySum
}

type inventory struct {
	left string
	right string
	both string
}

func parseInventory(s string) (inventory, error) {
	inv := inventory {
		left: "",
		right: "",
		both: "",
	}

	if len(s) % 2 != 0 {
		return inv, errors.New(fmt.Sprintf("Inventory string not an even number: %d", len(s)))
	}

	sideSize := len(s) / 2

	leftCharMap := make(map[byte]struct{})
	for i := 0; i < sideSize; i++ {
		c := s[i]
		leftCharMap[c] = util.Empty
	}
	rightCharMap := make(map[byte]struct{})
	for i := sideSize; i < len(s); i++ {
		c := s[i]
		rightCharMap[c] = util.Empty
	}

	leftChars := make([]byte, 0, len(leftCharMap))
	rightChars := make([]byte, 0, len(rightCharMap))
	bothChars := make([]byte, 0)
	for c := range leftCharMap {
		leftChars = append(leftChars, c)
		_, inRight := rightCharMap[c]
		if inRight {
			bothChars = append(bothChars, c)
		}

	}
	for c := range rightCharMap {
		rightChars = append(rightChars, c)
	}

	sort.Slice(leftChars, func(i, j int) bool {return leftChars[i] < leftChars[j]})
	sort.Slice(rightChars, func(i, j int) bool {return rightChars[i] < rightChars[j]})
	sort.Slice(bothChars, func(i, j int) bool {return bothChars[i] < bothChars[j]})

	inv.left = string(leftChars)
	inv.right = string(rightChars)
	inv.both = string(bothChars)
	return inv, nil
}

func itemPriority(char byte) (uint, error) {
	if 64 < char && char < 91 {
		return uint(char - 64 + 26), nil
	} else if 96 < char && char < 123 {
		return uint(char - 96), nil
	} else {
		return 0, errors.New(fmt.Sprintf("Character given outside ASCII range: %c", char))
	}
}
