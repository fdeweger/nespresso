package main

type Memory struct {
    cells [65536]uint16
}

func (r *Memory) Read(address uint16) uint16 {
    return r.cells[address]
}

func (r *Memory) Write(address, value uint16) {
    r.cells[address] = value
}
