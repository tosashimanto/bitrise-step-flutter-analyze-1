package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
	shellquote "github.com/kballard/go-shellquote"
)

var severityRegExp = map[string]string{
	"error":   "error",
	"warning": "(error|warning)",
	"info":    "(error|warning|info)",
}

type config struct {
	AdditionalParams string `env:"additional_params"`
	ProjectLocation  string `env:"project_location,dir"`
	FailSeverity     string `env:"fail_severity,opt[error,warning,info]"`
}

func failf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
	os.Exit(1)
}

func hasAnalyzeError(cmdOutput string, failSeverity string) bool {
	// example: error • Undefined class 'function' • lib/package.dart:3:1 • undefined_class
	severityPattern := severityRegExp[failSeverity]
	pattern := fmt.Sprintf(`%s.+\.dart:\d+:\d+`, severityPattern)
	analyzeErrorPattern := regexp.MustCompile(pattern)
	return analyzeErrorPattern.MatchString(cmdOutput)
}

func hasOtherError(cmdOutput string) bool {
	analyzeErrorPattern := regexp.MustCompile(`(error|warning|info).+\.dart:\d+:\d+`)
	return !analyzeErrorPattern.MatchString(cmdOutput)
}

func main() {
	var cfg config
	if err := stepconf.Parse(&cfg); err != nil {
		failf("Issue with input: %s", err)
	}
	stepconf.Print(cfg)

	additionalParams, err := shellquote.Split(cfg.AdditionalParams)
	if err != nil {
		failf("Failed to parse additional parameters, error: %s", err)
	}

	fmt.Println()
	log.Infof("Running analyze")

	var b bytes.Buffer
	multiwr := io.MultiWriter(os.Stdout, &b)
	analyzeCmd := command.New("flutter", append([]string{"analyze"}, additionalParams...)...).
		SetDir(cfg.ProjectLocation).
		SetStdout(multiwr).
		SetStderr(os.Stderr)

	fmt.Println()
	log.Donef("$ %s", analyzeCmd.PrintableCommandArgs())
	fmt.Println()

	if err := analyzeCmd.Run(); err != nil {
		if hasAnalyzeError(b.String(), cfg.FailSeverity) {
			log.Errorf("flutter analyze found errors: %s", err)
			os.Exit(1)
		} else if hasOtherError(b.String()) {
			failf("step failed with error: %s", err)
		}
	}
}
