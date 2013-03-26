package main

import (
    "strings"
    "loglog"
    "os"
    "bufio"
    "fmt"
    "time"
)

const (
    WORKERS = 8
    MBITS   = 12
)


// stream the input
func streamWords(filename string) (chan string) {
    ch := make(chan string)

    go func() {
        f, _ := os.Open(filename)
        reader := bufio.NewReader(f)
        for {
            line, err := reader.ReadString('\n')
            if err != nil { break }
            for _, word := range strings.Fields(line) {
                ch <- word
            }
        }
        close(ch)
    }()
    return ch
}

// scatter and gather

func scatter(in chan string, k int, counter *loglog.Counter) (chan *loglog.Entry) {
    out := make(chan *loglog.Entry)
    done := make(chan int)
    // start k workers
    for i:=0; i < k; i++ {
        go func(i int) {
            fmt.Println("Worker:", i)
            for {
                var entry loglog.Entry

                x, ok := <-in
                if !ok { break }

                hv := loglog.HashValue(x)
                loglog.SetEntry(hv, counter.MBits, &entry)
                out <- &entry
            }
            done <- 1
        }(i)
    }
    go func() {
        for i:=0; i < k; i++ {
            <-done
        }
        close(out)
    }()

    return out
}

func main() {
    c := loglog.NewCounter(MBITS)
    filename := "pg100.txt"
    words := streamWords(filename)
    hashvalues := scatter(words, WORKERS, c)

    t := time.Now().UnixNano()

    for {
        entry, ok := <-hashvalues
        if !ok {break}
        c.DigestEntry(entry)
    }

    fmt.Println(c.Estimate())
    d := float64(time.Now().UnixNano() - t)/1000000000.0
    fmt.Printf("In %.2f seconds.\n", d)
}
