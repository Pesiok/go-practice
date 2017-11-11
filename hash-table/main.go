package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatal(err)
	}

	numberOfBuckets := 12

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanning
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([]int, numberOfBuckets)

	// loop over the words
	// scanner.Scan is an iterator
	for scanner.Scan() {
		// get location
		n := HashBucket(scanner.Text(), numberOfBuckets)
		// increment value at n location (for getting number of elements)
		buckets[n]++
	}

	fmt.Println(buckets)

}

func HashBucket(word string, buckets int) int {
	var sum int
	// for each word, convert it to int and add all their values
	for _, value := range word {
		sum += int(value)
	}
	// ensure more even distribution
	// 34 % 3 = 1 -> first bucket
	return sum % buckets
}
