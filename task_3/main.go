package main

import "errors"

func Pop(inp []int) ([]int, error) {
	if len(inp) == 0 {
		return []int{}, errors.New("empty input array")
	}
	return inp[:len(inp)-1], nil
}
