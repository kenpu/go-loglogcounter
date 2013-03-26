package loglog

import (
    "crypto/md5"
    "io"
    "hash"
    "math/big"
    "math"
)

const (
    MAX_LEN = 128
    ALPHA = 0.39701
    BETA = 1.29806
)
func HashValue(s string) *big.Int {
    var (
        h hash.Hash = md5.New()
        x []byte
        value *big.Int
    )
    io.WriteString(h, s)
    x = h.Sum(nil)
    value = new(big.Int)
    value.SetBytes(x)
    return value
}
func rank(hv *big.Int) int {
    for i:=0; i < hv.BitLen(); i++ {
        if hv.Bit(i) > 0 {
            return i+1
        }
    }
    return MAX_LEN
}
func max(v1 int, vs ...int) (r int) {
    r = v1
    for _, v := range vs {
        if r < v {
            r = v
        }
    }

    return
}

type Entry struct {
    MValue uint64
    Rank int
}

func SetEntry(hv *big.Int, numBits int, entry *Entry) {
    mask := new(big.Int)
    for i:=0; i < numBits; i++ {
        mask.SetBit(mask, i, 1)
    }
    mv := uint64(mask.And(mask, hv).Int64())
    hv2 := new(big.Int).Rsh(hv, uint(numBits))

    entry.MValue = mv
    entry.Rank = rank(hv2)
}



type Counter struct {
    Table []int     // mValue -> rank
    MBits int       // bitlen of mValue
    entry Entry     // entry
}

func NewCounter(mBits int) *Counter {
    var N uint = 1 << uint(mBits)
    t := Counter{
        make([]int, N),
        mBits,
        Entry{0, 0},
    }
    return &t
}


func (counter *Counter) Digest(s string) {
    SetEntry(HashValue(s), counter.MBits, &counter.entry)
    counter.DigestEntry(&counter.entry)
    return
}

func (counter *Counter) DigestEntry(entry *Entry) {
    counter.Table[entry.MValue] = max(counter.Table[entry.MValue], entry.Rank)

    return
}

func (counter *Counter) Estimate() (E float64) {
    var (
        N uint = (1 << uint(counter.MBits))
        sum float64
    )

    for _, v := range(counter.Table) {
        sum += float64(v)
    }
    E = math.Exp2(sum / float64(N)) * ALPHA * float64(N)

    return
}