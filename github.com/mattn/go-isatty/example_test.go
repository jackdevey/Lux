package isatty_test

import (
	"fmt"
	"os"
)

func Example() {
	if IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Terminal")
	} else if IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Cygwin/MSYS2 Terminal")
	} else {
		fmt.Println("Is Not Terminal")
	}
}
