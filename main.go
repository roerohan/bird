package main

import (
	"bufio"
	"flag"
	"os"
	"sync"

	"github.com/roerohan/bird/brutus"
	"github.com/roerohan/bird/logger"
	"github.com/roerohan/bird/progress"
)

func main() {
	flag.Parse()

	if len(urls) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if wordlist == "" {
		flag.Usage()
		os.Exit(1)
	}

	if len(success) == 0 {
		success = append(success, "200")
	}

	workers := 4
	logger.Info("Starting 4 worker threads...")

	successCodes := make(map[string]bool)
	for _, code := range success {
		successCodes[code] = true
	}

	var wg sync.WaitGroup
	var bar progress.Progress
	c := make(chan *brutus.Brute)
	logs := make(chan logger.Log)

	go logger.Start(logs)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for b := range c {
				b.Try(successCodes, logs)
			}
		}()
	}

	file, err := os.Open(wordlist)
	if err != nil {
		logger.Fatal("Could not open wordlist: " + wordlist)
	}

	defer file.Close()

	stat, _ := file.Stat()
	size := stat.Size()

	bar.New(0, int(size))

	count := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		for _, url := range urls {
			c <- brutus.New(url, text)
		}

		count += len(text) + 1
		bar.Play(count, logs)
	}

	close(c)

	wg.Wait()
}
