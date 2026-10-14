package main

import (
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
		{"размер 1", 1, 1},
		{"размер 100", 100, 100},
		{"размер 1000", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("generateRandomElements(%d) вернул слайс длиной %d, ожидалось %d",
					tt.size, len(got), tt.want)
			}
		})
	}

	// Проверяем, что все элементы неотрицательны.
	t.Run("все элементы положительные", func(t *testing.T) {
		data := generateRandomElements(1000)
		for i, v := range data {
			if v < 0 {
				t.Errorf("элемент с индексом %d отрицательный: %d", i, v)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"максимум в начале", []int{10, 5, 3, 1}, 10},
		{"максимум в конце", []int{1, 3, 5, 10}, 10},
		{"максимум в середине", []int{1, 10, 5, 3}, 10},
		{"все одинаковые", []int{7, 7, 7, 7}, 7},
		{"отрицательные числа", []int{-5, -10, -3}, -3},
		{"большой слайс", func() []int {
			s := make([]int, 1000)
			for i := range s {
				s[i] = i
			}
			return s
		}(), 999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"меньше чем CHUNKS", []int{1, 2, 3}, 3},
		{"ровно CHUNKS", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"больше чем CHUNKS", []int{1, 5, 3, 2, 9, 4, 7, 6, 8, 0}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}

	// Сравнение с однопоточной версией на большом объёме данных.
	// Используем локальный генератор вместо устаревшего rand.Seed.
	t.Run("совпадает с maximum на большом слайсе", func(t *testing.T) {
		r := rand.New(rand.NewSource(1))
		data := make([]int, 100_000)
		for i := range data {
			data[i] = r.Int()
		}

		want := maximum(data)
		got := maxChunks(data)
		if got != want {
			t.Errorf("maxChunks() = %d, maximum() = %d — результаты не совпадают", got, want)
		}
	})
}
