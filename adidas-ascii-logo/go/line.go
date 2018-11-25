package main

import (
	"strings"
)

type Block struct {
	maxColumns int
	column     int
	width      int
}

func (l *Block) Draw(spaces int, symbol string) string {
	gaps := strings.Repeat(" ", (l.maxColumns-l.column)*l.width)
	line := strings.Repeat("@", l.width)

	lines := make([]string, l.column)
	for i := range lines {
		lines[i] = line
	}

	return gaps + strings.Join(lines, strings.Repeat(" ", spaces))
}
