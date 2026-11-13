//go:build js && wasm

// Command ghmd-wasm exposes render, lint, and diff through JavaScript.
package main

import (
	"encoding/json"
	"syscall/js"

	"ghmd"
	"ghmd/internal/diff"
	"ghmd/internal/lint"
)

type lintFinding struct {
	Line     int    `json:"line"`
	Column   int    `json:"column"`
	Severity string `json:"severity"`
	Rule     string `json:"rule"`
	Message  string `json:"message"`
}

type diffFinding struct {
	Path string `json:"path"`
	Want string `json:"want"`
	Got  string `json:"got"`
}

type response struct {
	Error string        `json:"error,omitempty"`
	HTML  string        `json:"html,omitempty"`
	Lint  []lintFinding `json:"lint,omitempty"`
	Diff  []diffFinding `json:"diff,omitempty"`
}

func main() {
	js.Global().Set("ghmdRun", js.FuncOf(run))
	if ready := js.Global().Get("ghmdReady"); ready.Type() == js.TypeFunction {
		ready.Invoke()
	}
	select {}
}

func run(_ js.Value, args []js.Value) any {
	if len(args) != 2 {
		return encode(response{Error: "ghmdRun requires mode and source"})
	}
	mode, source := args[0].String(), args[1].String()
	if len(source) > 1<<20 {
		return encode(response{Error: "input exceeds the 1 MiB limit"})
	}
	var result response
	switch mode {
	case "render":
		html, err := ghmd.Render(source, ghmd.RenderOptions{})
		if err != nil {
			result.Error = err.Error()
		} else {
			result.HTML = html
		}
	case "lint":
		findings, err := lint.Source(source)
		if err != nil {
			result.Error = err.Error()
			break
		}
		result.Lint = make([]lintFinding, len(findings))
		for i, finding := range findings {
			result.Lint[i] = lintFinding{
				Line: finding.Line, Column: finding.Column,
				Severity: finding.Severity.String(), Rule: finding.Rule, Message: finding.Message,
			}
		}
	case "diff":
		findings, err := diff.GitHubParsers(source)
		if err != nil {
			result.Error = err.Error()
			break
		}
		result.Diff = make([]diffFinding, len(findings))
		for i, finding := range findings {
			result.Diff[i] = diffFinding{Path: finding.Path, Want: finding.Want, Got: finding.Got}
		}
	default:
		result.Error = "unknown mode: " + mode
	}
	return encode(result)
}

func encode(value response) string {
	data, err := json.Marshal(value)
	if err != nil {
		return `{"error":"failed to encode response"}`
	}
	return string(data)
}
