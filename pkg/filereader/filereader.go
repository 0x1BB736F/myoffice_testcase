package filereader

import (
	"bufio"
	"io"
	"time"
)

type Reader struct {
	// for line by line scanning
	scanner bufio.Scanner
	// for async handling
	wChan chan string
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
		// close to stop for-range loop
		close(r.wChan)
	}()

	return r.wChan
}
