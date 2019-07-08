package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	cmd := "ping"
	timeout := 5 * time.Second

	// The command line tool "ping" is executed for 5 seconds
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	proc := exec.CommandContext(ctx, cmd, "google.com")

	// The process output is obtained
	// in form of io.ReadCloser. The underlying
	// implementation use the os.Pipe
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	// Start the process
	golog.DoItOrDie(proc.Start(), "starting process")

	// For more comfortable reading the bufio.Scanner is used. The read call is blocking.
	s := bufio.NewScanner(stdout)
	sum := 0.0
	count := 0
	for s.Scan() {
		fmt.Println(s.Text())
		words := regexp.MustCompile("[*,% ()]+").Split(s.Text(), -1)
		for idx, word := range words {
			if len(word) > 0 {
				fmt.Printf("Word %d is: %s\n", idx, word)
				if strings.HasPrefix(word, "time=") {
					parts := strings.Split(word, "=")
					if s, err := strconv.ParseFloat(parts[1], 64); err == nil {
						sum += s
						count += 1
					}
				}
			}
		}

	}
	fmt.Printf("Sended %d pings avg time was %v", count, sum/float64(count))
}
