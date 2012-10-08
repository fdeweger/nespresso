package main

import "testing"

/**
 * Register tests
 */
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

    cpu.testAndSetCarryAddition(0x100)
    if !cpu.getCarry() {
        t.Error("Expected getCarry to return true")
    }

    cpu.testAndSetCarryAddition(0x80)
    if cpu.getCarry() {
        t.Error("Expected getCarry to return false")
    }
}

func TestTestAndSetCarrySubtraction(t *testing.T) {
    cpu := Cpu{}

    cpu.testAndSetCarrySubtraction(0)
    if !cpu.getCarry() {
        t.Error("Expected getCarry to return true")
    }

    cpu.testAndSetCarrySubtraction(-1)
    if cpu.getCarry() {
        t.Error("Expected getCarry to return false")
    }

}

func TestTestAndSetOverflowAddition(t *testing.T) {
    var overflowAdditionTests = []struct {
        a, b, c  uint8
        expected bool
    }{
        {0x80, 0x7f, 0x00, false},
        {0x80, 0x80, 0x00, true},
        {0x80, 0x80, 0x80, false},
        {0x00, 0x00, 0x80, true},
    }

    cpu := new(Cpu)
    for _, val := range overflowAdditionTests {
        cpu.testAndSetOverflowAddition(val.a, val.b, val.c)
        if val.expected != cpu.getOverflow() {
            t.Errorf("Expected %t for %#X, %#X and %#X, got %t", val.expected, val.a, val.b, val.c, cpu.getOverflow())
        }
    }
}

func TestTestAndSetOverflowSubtraction(t *testing.T) {
    var overflowAdditionTests = []struct {
        a, b, c  uint8
        expected bool
    }{
        {0x80, 0x7f, 0x00, true},
        {0x80, 0x80, 0x00, false},
        {0x80, 0x80, 0x80, false},
        {0x00, 0x00, 0x80, false},
    }

    cpu := new(Cpu)
    for _, val := range overflowAdditionTests {
        cpu.testAndSetOverflowSubtraction(val.a, val.b, val.c)
        if val.expected != cpu.getOverflow() {
            t.Errorf("Expected %t for %#X, %#X and %#X, got %t", val.expected, val.a, val.b, val.c, cpu.getOverflow())
        }
    }
}

/**
 * Opcode tests
 */
type cpuTest struct {
    a1, x1, y1, p1 uint8
    a2, x2, y2, p2 uint8
    val            uint8
    cpu            *Cpu
}

func (c *cpuTest) setup() {
    c.cpu.A = c.a1
    c.cpu.X = c.x1
    c.cpu.Y = c.y1
    c.cpu.P = c.p1
    c.cpu.Ram = new(SimpleMemory)
}

func (c *cpuTest) test(t *testing.T) {
    if c.cpu.A != c.a2 || c.cpu.X != c.x2 || c.cpu.Y != c.y2 || c.cpu.P != c.p2 {
        t.Errorf("cpuTest:\nA: %#X\tX: %#X\tY: %#X\tP: %#X \nA: %#X\tX: %#X\tY: %#X\tP: %#X", c.a2, c.x2, c.y2, c.p2, c.cpu.A, c.cpu.X, c.cpu.Y, c.cpu.P)
    }
}

type memTest struct {
    loc           uint16
    before, after uint8
    p             uint8
    cpu           *Cpu
}

func (m *memTest) setup() {
    m.cpu.Ram = new(SimpleMemory)
    m.cpu.Ram.Write(m.loc, m.before)
}

func (m *memTest) test(t *testing.T) {
    if m.cpu.Ram.Read(m.loc) != m.after {
        t.Errorf("memTest: expected %#X, got %#X", m.after, m.cpu.Ram.Read(m.loc))
    }

    if m.cpu.P != m.p {
        t.Errorf("memTest: expected register state %#X, got %#X", m.p, m.cpu.P)
    }
}

func TestAdc(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x04, cpu},
        {0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x43, 0x80, cpu},
        {0x80, 0x00, 0x00, 0x00, 0x81, 0x00, 0x00, 0x80, 0x01, cpu},
        {0x80, 0x00, 0x00, 0x01, 0x82, 0x00, 0x00, 0x80, 0x01, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Adc(test.val)
        test.test(t)
    }
}

func TestAnd(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x88, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x08, cpu},
        {0x80, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x80, 0x80, cpu},
        {0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x40, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.And(test.val)
        test.test(t)
    }
}

func TestAsl(t *testing.T) {
    cpu := new(Cpu)
    tests := []memTest{
        {0x00, 0x01, 0x02, 0x00, cpu},
        {0x00, 0x40, 0x80, 0x80, cpu},
        {0x00, 0x80, 0x00, 0x03, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Asl(test.loc)
        test.test(t)
    }
}

func TestAslAcc(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, cpu},
        {0x20, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x40, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x80, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.AslAcc()
        test.test(t)
    }
}

func TestClc(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFE, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Clc()
        test.test(t)
    }
}

func TestCli(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFB, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Cli()
        test.test(t)
    }
}

func TestClv(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xBF, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Clv()
        test.test(t)
    }
}

func TestCmp(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x40, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x01, 0x20, cpu},
        {0x40, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x03, 0x40, cpu},
        {0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x81, 0x7F, cpu},
        {0x10, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x80, 0x20, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Cmp(test.val)
        test.test(t)
    }
}

func TestCpx(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x40, 0x00, 0x00, 0x00, 0x40, 0x00, 0x01, 0x20, cpu},
        {0x00, 0x40, 0x00, 0x00, 0x00, 0x40, 0x00, 0x03, 0x40, cpu},
        {0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x81, 0x7F, cpu},
        {0x00, 0x10, 0x00, 0x00, 0x00, 0x10, 0x00, 0x80, 0x20, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Cpx(test.val)
        test.test(t)
    }
}

func TestCpy(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x40, 0x01, 0x20, cpu},
        {0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x40, 0x03, 0x40, cpu},
        {0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x81, 0x7F, cpu},
        {0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x10, 0x80, 0x20, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Cpy(test.val)
        test.test(t)
    }
}

func TestDec(t *testing.T) {
    cpu := new(Cpu)

    tests := []memTest{
        {0x00, 0x02, 0x01, 0x00, cpu},
        {0x00, 0x81, 0x80, 0x80, cpu},
        {0x00, 0x01, 0x00, 0x02, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Dec(test.loc)
        test.test(t)
    }
}

func TestDex(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x80, 0x00, cpu},
        {0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Dex()
        test.test(t)
    }
}

func TestDey(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x00, cpu},
        {0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Dey()
        test.test(t)
    }
}

func TestEor(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0xFF, 0x00, 0x00, 0x00, 0xFE, 0x00, 0x00, 0x80, 0x01, cpu},
        {0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xFF, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Eor(test.val)
        test.test(t)
    }
}

func TestInc(t *testing.T) {
    cpu := new(Cpu)
    tests := []memTest{
        {0x00, 0x01, 0x02, 0x00, cpu},
        {0x00, 0x7f, 0x80, 0x80, cpu},
        {0x00, 0xff, 0x00, 0x02, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Inc(test.loc)
        test.test(t)
    }
}

func TestInx(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x7F, 0x00, 0x00, 0x00, 0x80, 0x00, 0x80, 0x00, cpu},
        {0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Inx()
        test.test(t)
    }
}

func TestIny(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x00, 0x00, 0x7F, 0x00, 0x00, 0x00, 0x80, 0x80, 0x00, cpu},
        {0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Iny()
        test.test(t)
    }
}

func TestLda(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x01, 0x02, 0x03, 0x00, 0x00, 0x02, 0x03, 0x02, 0x00, cpu},
        {0x01, 0x02, 0x03, 0x00, 0xF0, 0x02, 0x03, 0x80, 0xF0, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Lda(test.val)
        test.test(t)
    }
}

func TestLdx(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x01, 0x02, 0x03, 0x00, 0x01, 0x00, 0x03, 0x02, 0x00, cpu},
        {0x01, 0x02, 0x03, 0x00, 0x01, 0xF0, 0x03, 0x80, 0xF0, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Ldx(test.val)
        test.test(t)
    }
}

func TestLdy(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x01, 0x02, 0x03, 0x00, 0x01, 0x02, 0x00, 0x02, 0x00, cpu},
        {0x01, 0x02, 0x03, 0x00, 0x01, 0x02, 0xF0, 0x80, 0xF0, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Ldy(test.val)
        test.test(t)
    }
}

func TestLsr(t *testing.T) {
    //no need to test register states here, TestLsrAcc already does that implicitly
    cpu := new(Cpu)
    test := memTest{0x00, 0x02, 0x01, 0x00, cpu}
    test.setup()
    cpu.Lsr(test.loc)
    test.test(t)
}

func TestLsrAcc(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x80, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, cpu},
        {0x11, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x01, 0x00, cpu},
        {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.LsrAcc()
        test.test(t)
    }
}

func TestNop(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x01, 0x02, 0x03, 0x04, 0x01, 0x02, 0x03, 0x04, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Nop()
        test.test(t)
    }
}

func TestOra(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x7E, 0x00, 0x00, 0x00, 0x7F, 0x00, 0x00, 0x00, 0x01, cpu},
        {0xFE, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x80, 0x01, cpu},
        {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Ora(test.val)
        test.test(t)
    }
}

func TestSbc(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x10, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x00, 0x01, 0x01, cpu},
        {0x10, 0x00, 0x00, 0x01, 0x0F, 0x00, 0x00, 0x01, 0x01, cpu},
        {0x10, 0x00, 0x00, 0x01, 0xF0, 0x00, 0x00, 0x80, 0x20, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Sbc(test.val)
        test.test(t)
    }
}

func TestSec(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x01, 0x02, 0x03, 0x04, 0x01, 0x02, 0x03, 0x05, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Sec()
        test.test(t)
    }
}

func TestSei(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x01, 0x02, 0x03, 0x01, 0x01, 0x02, 0x03, 0x05, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Sei()
        test.test(t)
    }
}

func TestSta(t *testing.T) {
    cpu := new(Cpu)
    cpu.A = 0x03
    test := memTest{0x00, 0x00, 0x03, 0x00, cpu}
    test.setup()
    cpu.Sta(test.loc)
    test.test(t)
}

func TestStx(t *testing.T) {
    cpu := new(Cpu)
    cpu.X = 0x04
    test := memTest{0x00, 0x00, 0x04, 0x00, cpu}
    test.setup()
    cpu.Stx(test.loc)
    test.test(t)
}

func TestSty(t *testing.T) {
    cpu := new(Cpu)
    cpu.Y = 0x08
    test := memTest{0x00, 0x00, 0x08, 0x00, cpu}
    test.setup()
    cpu.Sty(test.loc)
    test.test(t)
}

func TestTax(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0xF0, 0x40, 0x00, 0x00, 0xF0, 0xF0, 0x00, 0x80, 0x00, cpu},
        {0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Tax()
        test.test(t)
    }
}

func TestTay(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0xF0, 0x40, 0x20, 0x00, 0xF0, 0x40, 0xF0, 0x80, 0x00, cpu},
        {0x00, 0x40, 0x20, 0x00, 0x00, 0x40, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Tay()
        test.test(t)
    }
}

func TestTxa(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x40, 0xF0, 0x00, 0x00, 0xF0, 0xF0, 0x00, 0x80, 0x00, cpu},
        {0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Txa()
        test.test(t)
    }
}

func TestTya(t *testing.T) {
    cpu := new(Cpu)
    tests := []cpuTest{
        {0x40, 0x00, 0xF0, 0x00, 0xF0, 0x00, 0xF0, 0x80, 0x00, cpu},
        {0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, cpu},
    }

    for _, test := range tests {
        test.setup()
        cpu.Tya()
        test.test(t)
    }
}
