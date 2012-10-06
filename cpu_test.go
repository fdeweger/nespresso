package main

import "testing"

func TestCarry(t *testing.T) {
    cpu := Cpu{}

    cpu.setCarry()
    if cpu.P != 0x01 {
        t.Errorf("Expected state of P register after setCarry is 0x01, got %#X", cpu.P)
    }

    if !cpu.getCarry() {
        t.Errorf("Expected value of getCarry is true, got false")
    }

    cpu.clearCarry()
    if cpu.P != 0x00 {
        t.Errorf("Expected state of P register after clearCarry is 0x00, got %#X", cpu.P)
    }
}

func TestZero(t *testing.T) {
    cpu := Cpu{}

    cpu.setZero()
    if cpu.P != 0x02 {
        t.Errorf("Expected state of P register after setZero is 0x02, got %#X", cpu.P)
    }

    if !cpu.getZero() {
        t.Errorf("Expected value of getZero is true, got false")
    }

    cpu.clearZero()
    if cpu.P != 0x00 {
        t.Errorf("Expected state of P register after clearZero is 0x00, got %#X", cpu.P)
    }
}

func TestInteruptDisable(t *testing.T) {
    cpu := Cpu{}

    cpu.setInteruptDisable()
    if cpu.P != 0x04 {
        t.Errorf("Expected state of P register after setInteruptDisable is 0x04, got %#X", cpu.P)
    }

    if !cpu.getInteruptDisable() {
        t.Errorf("Expected value of getInteruptDisable is true, got false")
    }

    cpu.clearInteruptDisable()
    if cpu.P != 0x00 {
        t.Errorf("Expected state of P register after clearInteruptDisable is 0x00, got %#X", cpu.P)
    }
}

func TestBrk(t *testing.T) {
    cpu := Cpu{}

    cpu.setBrk()
    if cpu.P != 0x10 {
        t.Errorf("Expected state of P register after setBrk is 0x10, got %#X", cpu.P)
    }

    if !cpu.getBrk() {
        t.Errorf("Expected value of getBrk is true, got false")
    }

    cpu.clearBrk()
    if cpu.P != 0x00 {
        t.Errorf("Expected state of P register after clearBrk is 0x00, got %#X", cpu.P)
    }
}

func TestOverflow(t *testing.T) {
    cpu := Cpu{}

    cpu.setOverflow()
    if cpu.P != 0x40 {
        t.Errorf("Expected state of P register after setOverflow is 0x40, got %#X", cpu.P)
    }

    if !cpu.getOverflow() {
        t.Errorf("Expected value of getOverflow is true, got false")
    }

    cpu.clearOverflow()
    if cpu.P != 0x00 {
        t.Errorf("Expected state of P register after clearOverflow is 0x00, got %#X", cpu.P)
    }
}

func TestNegative(t *testing.T) {
    cpu := Cpu{}

    cpu.setNegative()
    if cpu.P != 0x80 {
        t.Errorf("Expected state of P register after setNegative is 0x80, got %#X", cpu.P)
    }

    if !cpu.getNegative() {
        t.Errorf("Expected value of getNegative is true, got false")
    }

    cpu.clearNegative()
    if cpu.P != 0x00 {
        t.Errorf("Expected state of P register after clearOverflow is 0x00, got %#X", cpu.P)
    }
}

func TestTestAndSetNegative(t *testing.T) {
    cpu := Cpu{}
    cpu.testAndSetNegative(0x80)
    if !cpu.getNegative() {
        t.Error("Expected getNegative to return true")
    }

    cpu.testAndSetNegative(0x79)
    if cpu.getNegative() {
        t.Error("Expected getNegative to return false")
    }

}

func TestTestAndSetZero(t *testing.T) {
    cpu := Cpu{}
    cpu.testAndSetZero(0x00)
    if !cpu.getZero() {
        t.Error("Expected getZero to return true")
    }

    cpu.testAndSetZero(0x01)
    if cpu.getZero() {
        t.Error("Expected getZero to return false")
    }
}

func TestTestAndSetCarryAddition(t *testing.T) {
    cpu := Cpu{}

    cpu.testAndSetCarryAddition(0x81)
    if !cpu.getCarry() {
        t.Error("Expected getCarry to return true")
    }

    cpu.testAndSetCarryAddition(0x80)
    if cpu.getCarry() {
        t.Error("Expected getCarry to return false")
    }
}

var overflowAdditionTests = []struct {
    a, b, c  uint8
    expected bool
}{
    {0x80, 0x7f, 0x00, false},
    {0x80, 0x80, 0x00, true},
    {0x80, 0x80, 0x80, false},
    {0x00, 0x00, 0x80, true},
}

func TestTestAndSetOverflowAddition(t *testing.T) {
    cpu := new(Cpu)
    for _, val := range overflowAdditionTests {
        cpu.testAndSetOverflowAddition(val.a, val.b, val.c)
        if val.expected != cpu.getOverflow() {
            t.Errorf("Expected %t for %#X, %#X and %#X, got %t", val.expected, val.a, val.b, val.c, cpu.getOverflow())
        }
    }
}
