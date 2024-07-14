package main

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"os"
)

func main() {
	vi, err := (&youtube.Client{}).GetVideo("8LMJiFnnI2I")
	fmt.Println(vi, err)
	//exitOnError(rootCmd.Execute())
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
