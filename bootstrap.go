// +build bootstrap

// If you do not have mage installed you can run this.
// go run bootstrap.go
// go run bootstrap.go <task>
package main

import (
	"github.com/magefile/mage/mage"
	"os"
)

func main() {
	os.Exit(mage.Main())
}
