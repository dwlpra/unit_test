package mymath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return added number",
			args: args{10, 20},
			want: 30,
		},
		{
			name: "adding two negative numbers",
			args: args{-10, -5},
			want: -15,
		},
		{
			name: "adding positive and negative number",
			args: args{10, -5},
			want: 5,
		},
		{
			name: "adding two zeros",
			args: args{0, 0},
			want: 0,
		},
		{
			name: "adding number with zero",
			args: args{10, 0},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
