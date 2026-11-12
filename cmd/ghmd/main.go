package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"ghmd"
	"ghmd/internal/diff"
	"ghmd/internal/lint"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	command := "render"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		switch args[0] {
		case "render", "diff", "lint":
			command = args[0]
			args = args[1:]
		}
	}
	switch command {
	case "render":
		return renderCommand(args)
	case "diff":
		return diffCommand(args)
	case "lint":
		return lintCommand(args)
	}
	return errors.New("Usage: ghmd [render|diff|lint] [options] [input.md|-]")
}

func renderCommand(args []string) error {
	runID, path, err := parseRenderArgs(args)
	if err != nil {
		return err
	}
	source, err := readInput(path)
	if err != nil {
		return err
	}
	html, err := ghmd.Render(source, ghmd.RenderOptions{RunID: runID})
	if err != nil {
		return err
	}
	_, err = os.Stdout.WriteString(html)
	return err
}

func diffCommand(args []string) error {
	if len(args) != 1 || args[0] == "-h" || args[0] == "--help" {
		return errors.New("Usage: ghmd diff [input.md|-]")
	}
	source, err := readInput(args[0])
	if err != nil {
		return err
	}
	findings, err := diff.GitHubParsers(source)
	if err != nil {
		return err
	}
	for _, finding := range findings {
		fmt.Println(finding)
	}
	if len(findings) > 0 {
		os.Exit(1)
	}
	return nil
}

func lintCommand(args []string) error {
	if len(args) > 1 || len(args) == 1 && (args[0] == "-h" || args[0] == "--help" || strings.HasPrefix(args[0], "-") && args[0] != "-") {
		return errors.New("Usage: ghmd lint [input.md|-]")
	}
	path := ""
	if len(args) == 1 {
		path = args[0]
	}
	source, err := readInput(path)
	if err != nil {
		return err
	}
	findings, err := lint.Source(source)
	if err != nil {
		return err
	}
	displayPath := path
	if displayPath == "" || displayPath == "-" {
		displayPath = "<stdin>"
	}
	for _, finding := range findings {
		fmt.Printf("%s:%s\n", displayPath, finding)
	}
	if len(findings) > 0 {
		os.Exit(1)
	}
	return nil
}

func parseRenderArgs(args []string) (runID, path string, err error) {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--run-id" {
			if runID != "" || i+1 >= len(args) {
				return "", "", errors.New("--run-id requires one value")
			}
			i++
			runID = args[i]
		} else if (strings.HasPrefix(arg, "-") && arg != "-") || path != "" {
			return "", "", errors.New("Usage: render [--run-id ID] [input.md|-]")
		} else {
			path = arg
		}
	}
	return runID, path, nil
}

func readInput(path string) (string, error) {
	var data []byte
	var err error
	if path == "" || path == "-" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(path)
	}
	return string(data), err
}
