/**
 * Go-lang Miscellaneous Utility Library
 *
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed
 */

//Go-lang Miscellaneous Utility Library
package gutil

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var inputMsgs []string
var inputMsg string

func TestMain(m *testing.M) {
	//initialization
	inputMsgs = []string{
		"Take the Go-lang!",
		"Go言語で行こう！",
	}
	inputMsg = strings.Join(inputMsgs, "\n")

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestCliContextOutput(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliContext{Writer: outBuf}

	cliio.Output(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliContext.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliContextOutputln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliContext{Writer: outBuf}

	cliio.Outputln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliContext.Output = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestCliContextOutputBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliContext{Writer: outBuf}

	cliio.OutputBytes([]byte(inputMsg))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliContext.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliContextOutputErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliContext{ErrorWriter: outBuf}

	cliio.OutputErr(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliContext.OutputErr = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliContextOutputErrln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliContext{ErrorWriter: outBuf}

	cliio.OutputErrln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliContext.OutputErr = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestCliContextRefresh(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	cliio := CliContext{}
	cliio.Input(inBuf)
	if err := cliio.Refresh(); err != nil {
		t.Errorf("CliContext.Refresh = \"%v\", want nil.", err)
	}
}

func TestCliContextNewReader(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliContext{Reader: inBuf, ErrorWriter: outBuf}

	buf := new(bytes.Buffer)
	r, err := cliio.NewReader()
	if err != nil {
		t.Errorf("CliContext.NewReader = \"%v\", want nil.", err)
	}
	if _, err := buf.ReadFrom(r); err != nil {
		t.Errorf("CliContext.NewReader = \"%v\", want nil.", err)
	}
	result := buf.String()
	if result != inputMsg {
		t.Errorf("CliContext.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliContextCopyData(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliContext{Reader: inBuf, ErrorWriter: outBuf}

	result := bytes.NewBuffer(cliio.CopyData()).String()
	if result != inputMsg {
		t.Errorf("CliContext.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliContextData2String(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliContext{Reader: inBuf, ErrorWriter: outBuf}

	result := cliio.Data2String()
	if result != inputMsg {
		t.Errorf("CliContext.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliContextData2StringLines(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliContext{Reader: inBuf, ErrorWriter: outBuf}

	lines := cliio.Data2StringLines()
	result := strings.Join(lines, "\n")
	if result != inputMsg {
		t.Errorf("CliContext.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}
