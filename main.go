package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-l] <color spec>\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		`Where <color spec> is one of:
	#hexcode       - Typical HTML Hex color code
	<R> <G> <B>    - Decimal RGB values
	0.00,0.00,0.00 - LaTeX color spec

If -l is not given, then hexcodes and LaTeX codes give decimal RGB values, and vice versa.
`)
	os.Exit(1)
}

type RGB [3]int

var showLatexFmt bool
var showHelp bool

func init() {
	flag.BoolVar(&showLatexFmt, "l", false, "Output in LaTeX Color spec format")
	flag.BoolVar(&showHelp, "h", false, "Provide help")
	flag.Usage = usage
	flag.Parse()
}

func parseError(err error) {
	if err != nil {

		fmt.Fprintf(os.Stderr, "Error parsing color: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {

	if showHelp {
		usage()
	}

	var color RGB
	if flag.NArg() < 1 {
		usage()
	}

	if flag.NArg() == 1 {
		if strings.Contains(flag.Arg(0), ",") {
			parseError(parseLatex(color[:]))
			showOutput(outputRGB, color[:])
		} else {
			parseError(parseHex(color[:]))
			showOutput(outputRGB, color[:])
		}
	} else if flag.NArg() == 3 {
		parseError(parseRGB(color[:]))
		showOutput(outputHex, color[:])
	} else {
		usage()
	}
}

func parseHex(parts []int) error {
	str := flag.Arg(0)
	if len(str) == 7 {
		str = str[1:]
	}

	for i := 0; i < len(str); i += 2 {
		if part, err := strconv.ParseInt(str[i:i+2], 16, 32); err == nil {
			parts[i/2] = int(part)
		} else {
			return err
		}
	}

	return nil
}

func parseLatex(color []int) error {
	parts := strings.SplitN(flag.Arg(0), ",", 3)
	for i, p := range parts {
		if num, err := strconv.ParseFloat(p, 32); err == nil {
			color[i] = int(num * 255.0)
		} else {
			return err
		}
	}
	return nil
}

func parseRGB(parts []int) error {
	for i := 0; i < flag.NArg(); i++ {
		if part, err := strconv.Atoi(flag.Arg(i)); err == nil {
			parts[i] = part
		} else {
			return err
		}
	}
	return nil
}

func showOutput(fmtFunc func(parts []int), color []int) {
	if showLatexFmt {
		outputLatex(color)
		return
	}
	fmtFunc(color)
}

func outputHex(parts []int) {
	fmt.Printf("#%02X%02X%02X\n", parts[0], parts[1], parts[2])
}

func outputRGB(parts []int) {
	fmt.Printf("%d %d %d\n", parts[0], parts[1], parts[2])
}

func outputLatex(parts []int) {
	pcts := [3]float32{}
	for i := range pcts {
		f := float32(parts[i])
		pcts[i] = f / 255.0
	}
	fmt.Printf("%.2f,%.2f,%.2f\n", pcts[0], pcts[1], pcts[2])
}
