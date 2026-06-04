package main

import (
	"fmt"
	"os"

	"github.com/siddharthkundu/fintech-ledger/internal/version"
)

func main() {
	fmt.Println(version.Banner())
	os.Exit(0)
}
