package poker

import (
	"fmt"
	"strconv"
)

type Card struct {
	S uint8 //源值
	A bool  //是否为癞子
	U uint8 //癞子值
}

type Color uint8
type Value uint8

const (
	Diamond = 0x10 // 方块 ♦
	Club    = 0x20 // 梅花 ♣
	Heart   = 0x30 // 红桃 ♥
	Spade   = 0x40 // 黑桃 ♠

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

func (c *Card) String() string {
	var base = fmt.Sprintf("%s%s", c.GetColor().String(), c.GetValue().String())

	if c.U != c.S {
		u := fmt.Sprintf("%s%s", c.GetColor().String(), c.GetAnyValue().String())
		return fmt.Sprintf("癞子%s(%s)", u, base)
	}

	if c.IsAny() {
		base = "癞子" + base
	}
	return base
}

func (c *Card) GetColor() Color {
	return Color(uint8(c.S) & 0xf0)
}

func (c *Card) GetValue() Value {
	return Value(uint8(c.S) & 0x0f)
}

func (c *Card) GetAnyValue() Value {
	return Value(uint8(c.U) & 0x0f)
}

func (c *Card) SetAnyValue(v Value) *Card {
	c.U = uint8(c.GetColor()) | uint8(v)
	c.A = true
	return c
}

func (c *Card) ToAny() *Card {
	c.A = true
	return c
}
func (c *Card) IsAny() bool {
	return c.A
}

func NewColorCard(c Color, v Value) *Card {
	s := uint8(c) | uint8(v)
	return &Card{
		S: s,
		U: s,
	}
}

func NewCard(v uint8) *Card {
	return &Card{
		S: v,
		U: v,
	}
}
