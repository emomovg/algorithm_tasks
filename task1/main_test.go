package main

import (
	"reflect"
	"testing"
)

func TestProcessRequests(t *testing.T) {
	want := []int{0, 1, 4, 3, 5}
	request := [][]int{
		{1, 4},
		{2, 5},
		{1, 6},
		{2, 4},
		{1, 0},
		{1, 2},
		{2, 2},
		{2, 1},
		{1, 0},
		{2, 3},
	}

	got := ProcessRequests(request)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}
