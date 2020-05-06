package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func main() {
	os.Exit(Run())
}

func Run() int {
	flag.Parse()

	argc := flag.NArg()

	switch argc {
	case 0:
		useTUI()
	case 1:
		target := flag.Arg(0)
		useCLI(target)
	default:
		fmt.Fprintf(os.Stderr, "Error: A number of argments must be 0 or 1\n")
		return ExitCodeError
	}

	return ExitCodeOK
}

func useTUI() {
	fmt.Fprintf(os.Stderr, "Error: Not implemented\n")
}

func useCLI(target string) {

	var i int64

	switch {
	case strings.HasPrefix(target, "0x"):
		tmp, err := strconv.ParseInt(target[2:], 16, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid hex value.")
			return
		}
		i = tmp
	case strings.HasPrefix(target, "0b"):
		target = strings.Replace(target, "_", "", -1)
		tmp, err := strconv.ParseInt(target[2:], 2, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid binary value.")
			return
		}
		i = tmp
	case strings.HasPrefix(target, "0"):
		tmp, err := strconv.ParseInt(target[1:], 8, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid octal value.")
			return
		}
		i = tmp
	default:
		tmp, err := strconv.ParseInt(target, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid decimal value.")
			return
		}
		i = tmp
	}

	binTmp := strconv.FormatInt(i, 2)
	var padding int
	if len(binTmp)%4 == 0 {
		padding = 0
	} else {
		padding = 4 - (len(binTmp) % 4)
	}

	bin := ""
	bin += strings.Repeat("0", padding)
	for index, char := range binTmp {
		bin += string(char)
		if (padding+index)%4 == 3 && index+1 != len(binTmp) {
			bin += "_"
		}
	}

	fmt.Printf("bin[2]  : %s\n", bin)
	oct := strconv.FormatInt(i, 8)
	fmt.Printf("oct[8]  : %s\n", oct)
	dec := strconv.FormatInt(i, 10)
	fmt.Printf("dec[10] : %s\n", dec)
	hex := strconv.FormatInt(i, 16)
	fmt.Printf("hex[16] : %s\n", hex)
}
