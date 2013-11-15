package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"github.com/metakeule/loop"
	"io"
	"os"
	"os/exec"
)

func run(sound []byte) error {
	cmd := exec.Command("aplay", "-f", "cdr")
	cmd.Stdin = loop.New(sound)
	return cmd.Run()
}

func decompress() ([]byte, error) {
	b := new(bytes.Buffer)
	r := bzip2.NewReader(bytes.NewBuffer(noiseAiffBz2))
	_, err := io.Copy(b, r)
	return b.Bytes(), err
}

func main() {
	noiseAiff, err := decompress()
	if err != nil {
		fmt.Printf("Can't copy sound: %s\n", err)
		os.Exit(1)
	}
	start := 112
	end := len(noiseAiff) - 10
	err = run(noiseAiff[start:end])
	if err != nil {
		fmt.Printf("Error while running 'aplay -f cdr': %s\n", err)
		os.Exit(1)
	}
}
