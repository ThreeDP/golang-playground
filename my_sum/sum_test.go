package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Coleção de 5 números", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15
	
		if expected != result {
			t.Errorf("Result %d, expected %d, numbers %v", result, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1,2}, []int{0,9})
	expected := []int{3, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result %v expected %v", result, expected)
	}
}

func TestSumAllRest(t *testing.T) {
	checkSums := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v, expected %v", result, expected)
		}
	}
	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, result, expected)
	})
	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, result, expected)
	})
}