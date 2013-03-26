package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "time"
    "loglog"
)
func main() {
    filename := "pg100.txt"

    f, _ := os.Open(filename)
    reader := bufio.NewReader(f)
    i := 0
    N := 1000000
    t := time.Now().UnixNano()
    t0 := float64(t)
    counter := loglog.NewCounter(12)
    for {
        line, err := reader.ReadString('\n')
        if err != nil { break }
        for _, word := range strings.Fields(line) {
            i += 1
            counter.Digest(word)
            if i % N == 0 {
                d := time.Now().UnixNano() - t
                fmt.Printf("Scanned %d in %d ns.\n", N, d)
                t = time.Now().UnixNano()
            }
        }
    }
    fmt.Printf("Total fields: %d in %.2f s\n", i, (float64(time.Now().UnixNano()) - t0)/1000000000.0)
    fmt.Printf("Estimated distinct: %f\n", counter.Estimate())
}
