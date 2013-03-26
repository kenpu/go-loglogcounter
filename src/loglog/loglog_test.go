package loglog

import (
    "testing"
    "math/big"
)

func Test_max(t *testing.T) {
    if max(0, 2) != 2 {
        t.Fail()
    }
}
func Test_rank(t *testing.T) {
    i := big.NewInt(1)
    if rank(i) != 1 {
        t.Fail()
    }
}
func Test_hash(t *testing.T) {
    s := "hello world"
    hv := HashValue(s)
    t.Logf("\"%s\" -> %d, with bitlen=%d\n", s, uint64(hv.Int64()), hv.BitLen())
}

func Test_Counter(t *testing.T) {
    m := 5
    c := NewCounter(m)
    c.Digest("hello")
    c.Digest("world")
    c.Digest("again")
    t.Logf("c.table = %#v\n", c.Table)
    t.Logf("estimate = %f\n", c.Estimate())
}