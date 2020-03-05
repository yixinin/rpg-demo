package card

import (
	"fmt"
	"strconv"
)

type Card uint8 // 0=癞子
type Color uint8
type Value uint8

const (
	Diamond = 0x10 // 方块 ♦
	Club    = 0x20 // 梅花 ♣
	Heart   = 0x30 // 红桃 ♥
	Spade   = 0x40 // 黑桃 ♠

	DiamondAny = 0xa0 // 癞子方块 ♦
	ClubAny    = 0xb0 // 癞子梅花 ♣
	HeartAny   = 0xc0 // 癞子红桃 ♥
	SpadeAny   = 0xd0 // 癞子黑桃 ♠

	YungJoker = 0x50 //小王
	OldJoker  = 0x60 //大王
)

func (c Color) String() string {
	switch c {
	case Spade:
		return "♠"
	case Heart:
		return "♥"
	case Club:
		return "♣"
	case Diamond:
		return "♦"

	case YungJoker:
		return "小"
	case OldJoker:
		return "大"

	case SpadeAny:
		return "癞子♠"
	case HeartAny:
		return "癞子♥"
	case ClubAny:
		return "癞子♣"
	case DiamondAny:
		return "癞子♦"
	}

	return "*"
}

func (c Value) String() string {
	switch c {
	case 0x01:
		return "A"
	case 0x0a:
		return "10"
	case 0x0b:
		return "J"
	case 0x0c:
		return "Q"
	case 0x0d:
		return "K"
	case 0x0e:
		return "🃏"

	}
	return strconv.FormatUint(uint64(c), 10)
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.GetColor().String(), c.GetValue().String())
}

func (c Card) GetColor() Color {
	return Color(uint8(c) & 0xf0)
}

func (c Card) GetValue() Value {
	return Value(uint8(c) & 0x0f)
}
