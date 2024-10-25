package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"encoding/json"
)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSum(arr[][] int)int{
	n := len(arr)
	for i := n-2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			arr[i][j] += max(arr[i+1][j], arr[i+1][j+1])
		}
	}
	return(arr[0][0])
}

func main() {
	file, err := os.Open("hard.json")
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    
	bytes, err := io.ReadAll(file)
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }


    var triangle [][]int
    err = json.Unmarshal(bytes, &triangle)
    if err != nil {
        log.Fatalf("Failed to parse JSON: %v", err)
    }
	var test = [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	
	fmt.Println("max sum Test is: ", maxSum(test))
	fmt.Println("max sum Triangle is: ", maxSum(triangle))
}
