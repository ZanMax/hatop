package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const socketPath = "/run/haproxy/admin.sock"

func fetchAndParseStats() ([][]string, int, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	fmt.Fprintln(conn, "show stat")

	reader := csv.NewReader(bufio.NewReader(conn))
	data, err := reader.ReadAll()
	if err != nil {
		return nil, 0, err
	}

	var results [][]string
	totalQueue := 0
	if len(data) > 0 {
		for _, row := range data[1:] {
			if len(row) >= 18 {
				qcur := row[2]
				if qcur != "" {
					qnum, err := strconv.Atoi(qcur)
					if err == nil {
						totalQueue += qnum
					}
				}
				results = append(results, []string{row[0], row[1], row[4], row[17], row[34]})
			}
		}
	}
	return results, totalQueue, nil
}

func printTable(data [][]string, totalQueue int) {
	headers := []string{"PxName", "SvName", "Scur", "Status", "LastChk"}
	fmt.Println("HAProxy Statistics (refreshing every 2 seconds):")
	fmt.Println(strings.Join(headers, " | "))
	fmt.Println(strings.Repeat("-", 70))

	for _, row := range data {
		fmt.Println(strings.Join(row, " | "))
	}

	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("Queue: %d\n", totalQueue)
}

func main() {
	for {
		stats, totalQueue, err := fetchAndParseStats()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching stats: %v\n", err)
			os.Exit(1)
		}

		fmt.Print("\033[H\033[2J")
		printTable(stats, totalQueue)
		time.Sleep(2 * time.Second)
	}
}
