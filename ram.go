package main

type Memory interface {
    Read(uint16) uint8
    Write(uint16, uint8)
}

type SimpleMemory struct {
    cells [65536]uint8
}

func (r *SimpleMemory) Read(address uint16) uint8 {
    return r.cells[address]
}

func (r *SimpleMemory) Write(address uint16, value uint8) {
    r.cells[address] = value
}
