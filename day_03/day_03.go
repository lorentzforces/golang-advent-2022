package day_03

import (
	"errors"
	"fmt"
	"local-advent-2022/util"
)

type groupInventory []ruckContents
type ruckContents map[byte]struct{}

func PartOne(input string) uint {
	groups, err := parseHalvedStringGroups(util.AsLines(input))
	if err != nil {
		panic(err)
	}

	prioritySum := uint(0)
	for _, group := range groups {
		prioritySum += calcSharedPriorities(group)
	}
	return prioritySum
}

func PartTwo(input string) uint {
	groups, err := parseTripleLineGroups(util.AsLines(input))
	if err != nil {
		panic(err)
	}

	prioritySum := uint(0)
	for _, group := range groups {
		prioritySum += calcSharedPriorities(group)
	}
	return prioritySum
}

func calcSharedPriorities(group groupInventory) uint {
	groupItemsPresence := make(map[byte]int)
	for _, ruck := range group {
		for itemPrio := range ruck {
			currentVal := 0
			v, present := groupItemsPresence[itemPrio]
			if present {
				currentVal = v
			}
			groupItemsPresence[itemPrio] = currentVal + 1
		}
	}

	numRucks := len(group)
	totalPrio := uint(0)
	for itemPrio, countSeen := range groupItemsPresence {
		// if the item is in all rucks in the group
		if countSeen == numRucks {
			totalPrio += uint(itemPrio)
		}
	}
	return totalPrio
}

func parseTripleLineGroups(lines []string) ([]groupInventory, error) {
	if len(lines) % 3 != 0 {
		err := errors.New(fmt.Sprintf(
			"Input should be in groups of 3 lines, but found %d lines instead",
			len(lines),
		))
		return nil, err
	}

	groups := make([]groupInventory, 0, len(lines) / 3)
	for i := 0; i < len(lines); i += 3 {
		group := make(groupInventory, 3)
		for j := 0; j < 3; j++ {
			inventory, err := charValues(lines[i + j])
			if err != nil {
				return nil, err
			}
			group[j] = inventory
		}
		groups = append(groups, group)
	}

	return groups, nil
}


func parseHalvedStringGroups(lines []string) ([]groupInventory, error) {
	groups := make([]groupInventory, 0, len(lines))
	for _, line := range lines {
		lineGroups, err := parseHalvedStringInventories(line)
		if err != nil {
			return make([]groupInventory, 0), err
		}
		groups = append(groups, lineGroups)
	}
	return groups, nil
}

func parseHalvedStringInventories(s string) ([]ruckContents, error) {
	if len(s) % 2 != 0 {
		err := errors.New(fmt.Sprintf("Group inventories string not an even number: %d", len(s)))
		return nil, err
	}
	eachSize := len(s) / 2
	left, err := charValues(s[:eachSize])
	if err != nil {
		return nil, err
	}
	right, err := charValues(s[eachSize:])
	if err != nil {
		return nil, err
	}
	return []ruckContents{left, right}, nil
}

// From an input string:
// - remove duplicate characters
// - translate characters to their "priority" values
func charValues(s string) (ruckContents, error) {
	presentChars := make(map[byte]struct{})
	for i :=0; i < len(s); i++ {
		c := s[i]
		prio, err := itemPriority(c)
		if err != nil {
			return nil, err
		}
		presentChars[prio] = util.Empty
	}

	return presentChars, nil
}

type inventory struct {
	left string
	right string
	both string
}

func itemPriority(char byte) (byte, error) {
	if 64 < char && char < 91 {
		return byte(char - 64 + 26), nil
	} else if 96 < char && char < 123 {
		return byte(char - 96), nil
	} else {
		return 0, errors.New(fmt.Sprintf("Character given outside ASCII range: %c", char))
	}
}
