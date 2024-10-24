package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func p[T any](v T) *T {
	return &v
}

func TestGame_checkWinner(t *testing.T) {
	type fields struct {
		Board [3][3]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "first line X",
			fields: fields{
				Board: [3][3]string{
					{"X", "X", "X"},
					{" ", " ", " "},
					{" ", " ", " "},
				},
			},
			want: p("X"),
		},
		{
			name: "first line O",
			fields: fields{
				Board: [3][3]string{
					{"O", "O", "O"},
					{" ", " ", " "},
					{" ", " ", " "},
				},
			},
			want: p("O"),
		},
		{
			name: "second line",
			fields: fields{
				Board: [3][3]string{
					{" ", " ", " "},
					{"X", "X", "X"},
					{" ", " ", " "},
				},
			},
			want: p("X"),
		},
		{
			name: "third line",
			fields: fields{
				Board: [3][3]string{
					{" ", " ", " "},
					{" ", " ", " "},
					{"X", "X", "X"},
				},
			},
			want: p("X"),
		},
		{
			name: "diagonale left",
			fields: fields{
				Board: [3][3]string{
					{"X", " ", " "},
					{" ", "X", " "},
					{" ", " ", "X"},
				},
			},
			want: p("X"),
		},
		{
			name: "third right",
			fields: fields{
				Board: [3][3]string{
					{" ", " ", "X"},
					{" ", "X", " "},
					{"X", " ", " "},
				},
			},
			want: p("X"),
		},
		{
			name: "no one",
			fields: fields{
				Board: [3][3]string{
					{" ", " ", " "},
					{" ", " ", " "},
					{" ", " ", " "},
				},
			},
			want: nil,
		},
		{
			name: "no one 2",
			fields: fields{
				Board: [3][3]string{
					{"X", "0", "X"},
					{"0", "0", "X"},
					{"X", "X", "0"},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Board: tt.fields.Board,
			}
			winner := g.checkWinner()
			assert.Equal(t, tt.want, winner)

		})
	}
}
