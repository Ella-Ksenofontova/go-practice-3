package main

import (
	"strings"
)

type Battery struct {
	dischargedPart int
	chargedPart    int
}

func (b *Battery) fillParts(sequence string) {
	numberOfZeros := 0
	numberOfOnes := 0

	for _, sym := range sequence {
		switch sym {
		case '0':
			numberOfZeros++
		case '1':
			numberOfOnes++
		}
	}

	b.dischargedPart = numberOfZeros
	b.chargedPart = numberOfOnes
}

func (b Battery) String() string {
	return "[" + strings.Repeat(" ", b.dischargedPart) + strings.Repeat("X", b.chargedPart) + "]"
}
