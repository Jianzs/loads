package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

type keyValue struct {
	key   string
	value int64
}

func mapFn(in []byte, jobMd5 string) []keyValue {
	kvs := []keyValue{}

	curBeg := 0
	for curIdx := 0; curIdx < len(in); curIdx++ {
		if isDigit
	}
}

func redFn() {

}

var inPath = flag.String("in", "in.csv", "input file path")
var outPath = flag.String("out", "output.csv", "output file path")
var mapSize = flag.Int("msize", 64, "the map size in MB")
var noMap = flag.Int("mnum", 1, "the number of mapper")
var noRed = flag.Int("rnum", 1, "the number of reducer")

func main() {
	if _, err := os.Stat(*inPath); os.IsNotExist(err) {
		fmt.Printf("no such file: %s\n", *inPath)
		return
	}

	inFile, err := os.OpenFile(*inPath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("open file failed, %s\n", err.Error())
		return
	}
	defer inFile.Close()

	md := md5.Sum([]byte(*outPath))
	jobMd5 := hex.EncodeToString(md[:])[:6]

	mapCh := make(chan struct{}, *noMap)
	var wg sync.WaitGroup
	iMap := 0
	for {
		buf := make([]byte, *mapSize)
		n, err := inFile.ReadAt(buf, int64(*mapSize*iMap))
		if n == 0 && err == io.EOF {
			break
		}

		go func(iMap int) {
			wg.Add(1)
			mapCh <- struct{}{}
			defer func() { <-mapCh; wg.Done() }()

			mapFn(buf, jobMd5)
		}(iMap)

		iMap++
	}
	wg.Wait()
}
