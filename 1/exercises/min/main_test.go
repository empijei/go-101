package main_test

import (
	"testing"

	main "empijei.science/go-101/1/exercises/min"
)

func TestComputeMin(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		want    int
		wantErr bool
	}{
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name: "two vals",
			data: []int{1, 2},
			want: 1,
		},
		{
			name: "two vals, inverted",
			data: []int{2, 1},
			want: 1,
		},
		{
			name: "repeated",
			data: []int{2, 1, 2},
			want: 1,
		},
		{
			name: "neg",
			data: []int{2, 1, -1, 1, 2},
			want: -1,
		},
		{
			name: "many",
			data: []int{20, 12, -10, 0, 2, 5, -3},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := main.ComputeMin(tt.data)
			if tt.wantErr != (err != nil) {
				t.Fatalf("Error expectation mismatch: want error %v but got %v", tt.wantErr, err)
			}
			if got != tt.want {
				t.Errorf("ComputeMin(%v): got %v want %v", tt.data, got, tt.want)
			}
		})
	}
}
