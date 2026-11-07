package main_test

import (
	"testing"

	main "empijei.science/go-101/1/5-control-flow"
)

func TestExecuteSimple(t *testing.T) {
	got, err := main.Execute(2, 7, "*")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	want := 14
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name, op   string
		a, b, want int
	}{
		{
			name: "mul",
			op:   "*",
			a:    2,
			b:    10,
			want: 20,
		},
		{
			name: "div",
			op:   "/",
			a:    6,
			b:    3,
			want: 2,
		},
		{
			name: "sub",
			op:   "-",
			a:    10,
			b:    8,
			want: 2,
		},
		{
			name: "sum",
			op:   "+",
			a:    2,
			b:    2,
			want: 4,
		},
		/*
			{
				name: "mod",
				op:   "%",
				a:    8,
				b:    3,
				want: 2,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := main.Execute(tt.a, tt.b, tt.op)
			if err != nil {
				t.Fatalf("Unexpected err %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, tt.want %v", got, tt.want)
			}
		})
	}
}
