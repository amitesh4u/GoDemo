package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func testHttpCall() {
	makeGetCall()
}

func makeGetCall() {
	res, getErr := http.Get("https://google.com")
	if getErr != nil {
		log.Fatal(getErr)
		return
	}
	fmt.Println(res)

	IOReader(res)

	ResponseReader(res)

	CustomLogWriter(res)
}

type customWriter struct{}

func (customWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func CustomLogWriter(res *http.Response) {
	cw := customWriter{}
	dataLen, err := io.Copy(cw, res.Body)

	fmt.Printf("\nWritten %d bytes using custom writer\n", dataLen)

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseReader(res *http.Response) {
	// Using Reader of Response. Even a successful read returns err == EOF which is reported as Error
	bs := make([]byte, 99999)
	dataLen, err := res.Body.Read(bs)

	// Make sure to process the return bytes before handling error
	fmt.Println(string(bs))
	fmt.Printf("\nWritten %d bytes using Response Body Reader\n", dataLen)

	if err != nil {
		log.Fatal(err)
	}
}

func IOReader(res *http.Response) {
	// Using IO util. A successful copy returns err == nil and not err == EOF
	dataLen, ioErr := io.Copy(os.Stdout, res.Body)
	fmt.Printf("Written %d bytes using IO Copy Reader", dataLen)
	if ioErr != nil {
		log.Fatal(ioErr)
	}
}
