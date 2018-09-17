package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {

	fromCommitArg := flag.String("fromCommit", "...", "From commit...")
	toCommitArg := flag.String("toCommit", "...", "To commit...")
	copyToArg := flag.String("copyTo", "./tmp", "Copy to destination directory")

	flag.Parse()

	green := color.New(color.FgHiGreen).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()
	cyan := color.New(color.FgHiCyan).SprintFunc()

	fmt.Fprintf(color.Output, "From commit: %s\n", green(*fromCommitArg))
	fmt.Fprintf(color.Output, "To commit: %s\n", green(*toCommitArg))
	fmt.Fprintf(color.Output, "Copy files to: %s\n", green(*copyToArg))

	cmd := exec.Command("git", "diff", "--name-only", *fromCommitArg, *toCommitArg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	arrow := "→"
	if runtime.GOOS == "windows" {
		arrow = "=>"
	}

	lines := strings.Split(string(out), "\n")
	for i, line := range lines {
		if len(line) > 0 {
			d, f := filepath.Split(line)
			// Windows compatibility
			// git return list of files in UNIX slash but os.PathSeparator return "\"
			d = strings.Replace(d, "/", string(os.PathSeparator), -1)
			dirs := strings.Split(d, string(os.PathSeparator))
			schema := dirs[0]
			fmt.Fprintf(color.Output, "  [%3.3d] %v %v %v\n", i, cyan(schema), arrow, yellow(f))

			// Create new directory
			newDir := *copyToArg + string(os.PathSeparator) + schema
			err = os.MkdirAll(newDir, 0770)
			if err != nil {
				panic(err)
			}

			// Copy file to...
			fileCopyTo := newDir + string(os.PathSeparator) + f
			if _, err := os.Stat(line); err == nil {
				_, err = copy(line, fileCopyTo)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
