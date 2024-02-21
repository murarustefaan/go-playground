/**
 * Read from stdin, normalize and print the top 10 most frequent words.
 * Usage: `go run cmd/slices/slices.go < util/words.txt`
 */

package main

import (
	"bufio"
	"fmt"
	"go-playground/pkg/words"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	counts := make(map[string]int)
	for scan.Scan() {
		normalized, ok := words.NormalizeWord(scan.Text())
		if !ok {
			continue
		}

		counts[normalized]++
	}

	fmt.Println("Unique words:\t", len(counts))

	var temp []words.KV
	for word, count := range counts {
		temp = append(temp, words.KV{Key: word, Value: count})
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i].Value > temp[j].Value
	})
	for _, kv := range temp[:10] {
		fmt.Println(kv.Key+"\t", kv.Value)
	}
}
