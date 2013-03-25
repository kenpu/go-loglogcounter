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
    hv := hashValue(s)
    t.Logf("\"%s\" -> %d, with bitlen=%d\n", s, uint64(hv.Int64()), hv.BitLen())
    mv, hv2 := valueSplit(hv, 10)
    t.Logf("the first 10 bits are: %d\n", mv)
    t.Logf("the reminder is: %d with bitlen=%d\n", uint64(hv2.Int64()), hv2.BitLen())
    t.Logf("rank = %d\n", rank(hv2.Lsh(hv2, 4)))
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