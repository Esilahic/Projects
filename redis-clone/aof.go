package main

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type AOF struct {
	file *os.File
	rd   *bufio.Reader
	mu   sync.Mutex
}

func NewAof(path string) (*AOF, error) {
	// create the file if it doesn’t exist or open it if it does.
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	// read the file
	aof := &AOF{file: f, rd: bufio.NewReader(f)}

	// goroutine to sync aof file every 1 sec
	go func() {
		for {
			aof.mu.Lock()
			aof.file.Sync()
			aof.mu.Unlock()
			time.Sleep(time.Second)
		}
	}()
	return aof, nil
}

func (aof *AOF) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	return aof.file.Close()
}

// write the command to the AOF file whenever we receive a request
func (aof *AOF) Write(value Value) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	// we use value.Marshal() to write the command to the file in the same RESP format that we receive.
	_, err := aof.file.Write(value.Marshal())
	if err != nil {
		return err
	}
	return nil
}
