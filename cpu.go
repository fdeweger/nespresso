package main

import "fmt"

type Cpu struct {
    A   uint8
    X   uint8
    Y   uint8
    P   uint8
}

func (c *Cpu) getCarry() bool {
    return c.P&0x01 == 0x01
}

func (c *Cpu) setCarry() {
    c.P = c.P | 0x01
}

func (c *Cpu) clearCarry() {
    c.P = c.P & 0xFE
}

func (c *Cpu) getZero() bool {
    return c.P&0x02 == 0x02
}

func (c *Cpu) setZero() {
    c.P = c.P | 0x02
}

func (c *Cpu) clearZero() {
    c.P = c.P & 0xFD
}

func (c *Cpu) getInteruptDisable() bool {
    return c.P&0x04 == 0x04
}

func (c *Cpu) setInteruptDisable() {
    c.P = c.P | 0x04
}

func (c *Cpu) clearInteruptDisable() {
    c.P = c.P & 0xFB
}

func (c *Cpu) getBrk() bool {
    return c.P&0x10 == 0x10
}

func (c *Cpu) setBrk() {
    c.P = c.P | 0x10
}

func (c *Cpu) clearBrk() {
    c.P = c.P & 0xEF
}

func (c *Cpu) getOverflow() bool {
    return c.P&0x40 == 0x40
}

func (c *Cpu) setOverflow() {
    c.P = c.P | 0x40
}

func (c *Cpu) clearOverflow() {
    c.P = c.P & 0xBF
}

func (c *Cpu) getNegative() bool {
    return c.P&0x80 == 0x80
}

func (c *Cpu) setNegative() {
    c.P = c.P | 0x80
}

func (c *Cpu) clearNegative() {
    c.P = c.P & 0x7F
}

func (c *Cpu) testAndSetNegative(val uint8) {
    if val&0x80 == 0x80 {
        c.setNegative()
    } else {
        c.clearNegative()
    }
}

func (c *Cpu) testAndSetZero(val uint8) {
    if val == 0x00 {
        c.setZero()
    } else {
        c.clearZero()
    }
}

func (c *Cpu) testAndSetCarryAddition(val int) {
    if val > 0x80 {
        c.setCarry()
    } else {
        c.clearCarry()
    }
}

//see http://teaching.idallen.com/dat2343/10f/notes/040_overflow.txt for an excellent explanation
func (c *Cpu) testAndSetOverflowAddition(a, b, d uint8) {
    if ((a^b)&0x80 == 0x00) && ((a^d)&0x80 == 0x80) {
        c.setOverflow()
    } else {
        c.clearOverflow()
    }
}

func (c *Cpu) Adc(val uint8) {
    old := c.A

    c.A += old + (c.P & 0x01)
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
    c.testAndSetCarryAddition(int(c.A) + int(old) + int(c.P&0x01))
    c.testAndSetOverflowAddition(old, val, c.A)
}

func (c *Cpu) And(val uint8) {
    c.A = c.A & val
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) Clc() {
    c.clearCarry()
}

func (c *Cpu) Cli() {
    c.clearInteruptDisable()
}

func (c *Cpu) Clv() {
    c.clearOverflow()
}

func (c *Cpu) Dex() {
    c.X--
    c.testAndSetNegative(c.X)
    c.testAndSetZero(c.X)
}

func (c *Cpu) Dey() {
    c.Y--
    c.testAndSetNegative(c.Y)
    c.testAndSetZero(c.Y)
}

func (c *Cpu) Inx() {
    c.X++
    c.testAndSetNegative(c.X)
    c.testAndSetZero(c.X)
}

func (c *Cpu) Iny() {
    c.Y++
    c.testAndSetNegative(c.Y)
    c.testAndSetZero(c.Y)
}

func (c *Cpu) Tax() {
    c.X = c.A
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) Tay() {
    c.Y = c.A
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) Txa() {
    c.A = c.X
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) Tya() {
    c.A = c.Y
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) Dump() string {
    return fmt.Sprintf("X: %#X\tY: %#X\nA: %#X\tP: %#X\n", c.X, c.Y, c.A, c.P)
}
