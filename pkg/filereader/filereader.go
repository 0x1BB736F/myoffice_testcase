package filereader

import (
	"bufio"
	"io"
	"time"
)

type Reader struct {
	scanner bufio.Scanner
	wChan   chan string
}

func New(rc io.Reader) *Reader {
	return &Reader{
		scanner: *bufio.NewScanner(rc),
		wChan:   make(chan string, 10),
	}
}

func (r *Reader) ReadChan() chan string {
	go func() {
		for r.scanner.Scan() {
			// add some delay
			<-time.Tick(time.Millisecond * 50)
			r.wChan <- r.scanner.Text()
		}
		close(r.wChan)
	}()

	return r.wChan
}
