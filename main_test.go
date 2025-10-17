package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		size    int
		wantLen int
	}{
		// Корректные значения
		{
			size:    0,
			wantLen: 0,
		},
		{
			size:    -10,
			wantLen: 0,
		},
		{
			size:    -10000,
			wantLen: 0,
		},
		{
			size:    1,
			wantLen: 1,
		},
		{
			size:    5,
			wantLen: 5,
		},
		{
			size:    1_000_000,
			wantLen: 1_000_000,
		},
	}

	for _, tt := range tests {
		numbers := generateRandomElements(tt.size)
		require.Len(t, numbers, tt.wantLen)
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		slice   []int
		wantMax int
	}{
		{
			slice:   []int{},
			wantMax: 0,
		},
		{
			slice:   nil,
			wantMax: 0,
		},
		{
			slice:   []int{0, 1, 2, 10, -5},
			wantMax: 10,
		},
		{
			slice:   []int{4},
			wantMax: 4,
		},
		{
			slice:   []int{1000, 50, 3, 14532},
			wantMax: 14532,
		},
		{
			slice:   []int{100, 100, 100},
			wantMax: 100,
		},
		{
			slice:   []int{-100, -10, -25},
			wantMax: -10,
		},
	}

	for _, tt := range tests {
		max := maximum(tt.slice)
		require.Equal(t, tt.wantMax, max)
	}
}
