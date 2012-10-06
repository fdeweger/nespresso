package main

import "testing"

func TestRam(t *testing.T) {
    r := new(Memory)
    if r.Read(0) != 0 {
        t.Fail()
    }

    r.Write(0, 1)

    if r.Read(0) != 1 {
        t.Fail()
    }
}
