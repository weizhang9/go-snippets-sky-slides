package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var b = Buffer{buffer: bytes.NewBuffer([]byte{})}
	var wf WriterFlusher = &b
	var URL string
	fmt.Print("Enter the website address you would like to curl: ")
	fmt.Scan(&URL)
	wf.Write(curlWeb(URL))
	wf.Flush()
}

type Writer interface {
	Write([]byte) (int, error)
}

type Flusher interface {
	Flush() error
}

type WriterFlusher interface {
	Writer
	Flusher
}

type Buffer struct {
	buffer *bytes.Buffer
}

func (b *Buffer) Write(p []byte) (int, error) {
	n, err := b.buffer.Write(p)
	if err != nil {
		return 0, err
	}

	cut := make([]byte, 8)
	for b.buffer.Len() > 8 {
		_, err := b.buffer.Read(cut)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(cut))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (b *Buffer) Flush() error {
	for b.buffer.Len() > 0 {
		p := b.buffer.Next(8)
		_, err := fmt.Println(string(p))
		if err != nil {
			return err
		}
	}
	return nil
}

func curlWeb(link string) []byte {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err.Error)
	}
	defer res.Body.Close()
	p, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error)
	}
	return p
}
