package main

import "fmt"

type Cpu struct {
    A   uint8
    X   uint8
    Y   uint8
    P   uint8
    Ram Memory
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
    if val > 0xFF {
        c.setCarry()
    } else {
        c.clearCarry()
    }
}

func (c *Cpu) testAndSetCarrySubtraction(val int) {
    if val < 0 {
        c.clearCarry()
    } else {
        c.setCarry()
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

func (c *Cpu) testAndSetOverflowSubtraction(a, b, d uint8) {
    if ((a^b)&0x80 == 0x80) && ((a^d)&0x80 == 0x80) {
        c.setOverflow()
    } else {
        c.clearOverflow()
    }
}

func (c *Cpu) Adc(val uint8) {
    old := c.A

    c.A += val + (c.P & 0x01)
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
    c.testAndSetCarryAddition(int(val) + int(old) + int(c.P&0x01))
    c.testAndSetOverflowAddition(old, val, c.A)
}

func (c *Cpu) And(val uint8) {
    c.A = c.A & val
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) aslBase(val uint8) uint8 {
    if val&0x80 == 0x80 {
        c.setCarry()
    } else {
        c.clearCarry()
    }

    val = val << 1
    c.testAndSetNegative(val)
    c.testAndSetZero(val)
    return val
}

func (c *Cpu) Asl(loc uint16) {
    c.Ram.Write(loc, c.aslBase(c.Ram.Read(loc)))
}

func (c *Cpu) AslAcc() {
    c.A = c.aslBase(c.A)
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

func (c *Cpu) Cmp(val uint8) {
    res := c.A - val
    c.testAndSetNegative(res)
    c.testAndSetZero(res)
    c.testAndSetCarrySubtraction(int(c.A) - int(val))
}

func (c *Cpu) Cpx(val uint8) {
    res := c.X - val
    c.testAndSetNegative(res)
    c.testAndSetZero(res)
    c.testAndSetCarrySubtraction(int(c.X) - int(val))
}

func (c *Cpu) Cpy(val uint8) {
    res := c.Y - val
    c.testAndSetNegative(res)
    c.testAndSetZero(res)
    c.testAndSetCarrySubtraction(int(c.Y) - int(val))
}

func (c *Cpu) Dec(loc uint16) {
    val := c.Ram.Read(loc)
    val--

    c.testAndSetNegative(val)
    c.testAndSetZero(val)
    c.Ram.Write(loc, val)
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

func (c *Cpu) Eor(val uint8) {
    c.A = c.A ^ val
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) Inc(loc uint16) {
    val := c.Ram.Read(loc)
    val++

    c.testAndSetNegative(val)
    c.testAndSetZero(val)
    c.Ram.Write(loc, val)
}

func (c *Cpu) Inx() {
    c.X++
    c.testAndSetNegative(c.X)
    c.testAndSetZero(c.X)
}

func (c *Cpu) Lda(val uint8) {
    c.A = val
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) Ldx(val uint8) {
    c.X = val
    c.testAndSetNegative(c.X)
    c.testAndSetZero(c.X)
}

func (c *Cpu) Ldy(val uint8) {
    c.Y = val
    c.testAndSetNegative(c.Y)
    c.testAndSetZero(c.Y)
}

func (c *Cpu) lsrBase(val uint8) uint8 {
    if val&0x01 == 0x01 {
        c.setCarry()
    } else {
        c.clearCarry()
    }

    val = val >> 1
    c.testAndSetNegative(val)
    c.testAndSetZero(val)
    return val
}

func (c *Cpu) LsrAcc() {
    c.A = c.lsrBase(c.A)
}

func (c *Cpu) Lsr(loc uint16) {
    c.Ram.Write(loc, c.lsrBase(c.Ram.Read(loc)))
}

func (c *Cpu) Iny() {
    c.Y++
    c.testAndSetNegative(c.Y)
    c.testAndSetZero(c.Y)
}

func (c *Cpu) Nop() {
    return
}

func (c *Cpu) Ora(val uint8) {
    c.A = c.A | val
    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
}

func (c *Cpu) rolBase(val uint8) uint8 {
    carry := val & 0x80
    
    val = val << 1
    
    if c.getCarry() {
        val += 1
    }
    
    if carry == 0x80 {
        c.setCarry()
    } else  {
        c.clearCarry();
    }
    
    c.testAndSetNegative(val)
    c.testAndSetZero(val)
    
    return val 
}

func (c *Cpu) RolAcc() {
    c.A = c.rolBase(c.A)
}

func (c *Cpu) Rol(loc uint16) {
    c.Ram.Write(loc, c.rolBase(c.Ram.Read(loc)))
}

func (c *Cpu) rorBase(val uint8) uint8 {
    carry := val & 0x01
    
    val = val >> 1
    
    if c.getCarry() {
        val = val | 0x80
    }
    
    if carry == 0x01 {
        c.setCarry()
    } else  {
        c.clearCarry();
    }
    
    c.testAndSetNegative(val)
    c.testAndSetZero(val)
    
    return val 
}

func (c *Cpu) RorAcc() {
    c.A = c.rorBase(c.A)
}

func (c *Cpu) Ror(loc uint16) {
    c.Ram.Write(loc, c.rorBase(c.Ram.Read(loc)))
}

func (c *Cpu) Sbc(val uint8) {
    old := c.A //0x10
    c.A = old - val
    c.A = c.A - (1 - c.P&0x01) //0xf0

    c.testAndSetNegative(c.A)
    c.testAndSetZero(c.A)
    c.testAndSetOverflowSubtraction(old, val, c.A)
    c.testAndSetCarrySubtraction(int(old) - int(val) - (1 - int(c.P&0x01)))
}

func (c *Cpu) Sec() {
    c.setCarry()
}

func (c *Cpu) Sei() {
    c.setInteruptDisable()
}

func (c *Cpu) Sta(loc uint16) {
    c.Ram.Write(loc, c.A)
}

func (c *Cpu) Stx(loc uint16) {
    c.Ram.Write(loc, c.X)
}

func (c *Cpu) Sty(loc uint16) {
    c.Ram.Write(loc, c.Y)
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
    return fmt.Sprintf("X: %#X\tY: %#X\nA: %#X\tP: %#X\n", c.A, c.X, c.Y, c.P)
}
