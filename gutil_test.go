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
	"runtime"
	"strings"
	"testing"
)

var inputMsgs []string
var inputMsg string
var lineEnding string

func TestMain(m *testing.M) {
	//initialization
	inputMsgs = []string{
		"Take the Go-lang!",
		"Go言語で行こう！",
	}
	inputMsg = strings.Join(inputMsgs, "\n")
	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	} else {
		lineEnding = "\n"
	}

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestCliUiOutput(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliUi{Writer: outBuf}

	cliio.Output(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiOutputln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliUi{Writer: outBuf}

	cliio.Outputln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestCliUiOutputBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliUi{Writer: outBuf}

	cliio.OutputBytes([]byte(inputMsg))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiOutputErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliUi{ErrorWriter: outBuf}

	cliio.OutputErr(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.OutputErr = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiOutputErrln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliio := CliUi{ErrorWriter: outBuf}

	cliio.OutputErrln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliUi.OutputErr = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestCliUiRefresh(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	cliio := CliUi{}
	cliio.Input(inBuf)
	if err := cliio.Refresh(); err != nil {
		t.Errorf("CliUi.Refresh = \"%v\", want nil.", err)
	}
}

func TestCliUiNewReader(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliUi{Reader: inBuf, ErrorWriter: outBuf}

	buf := new(bytes.Buffer)
	r, err := cliio.NewReader()
	if err != nil {
		t.Errorf("CliUi.NewReader = \"%v\", want nil.", err)
	}
	if _, err := buf.ReadFrom(r); err != nil {
		t.Errorf("CliUi.NewReader = \"%v\", want nil.", err)
	}
	result := buf.String()
	if result != inputMsg {
		t.Errorf("CliUi.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiCopyData(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliUi{Reader: inBuf, ErrorWriter: outBuf}

	result := bytes.NewBuffer(cliio.CopyData()).String()
	if result != inputMsg {
		t.Errorf("CliUi.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiData2String(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliUi{Reader: inBuf, ErrorWriter: outBuf}

	result := cliio.Data2String()
	if result != inputMsg {
		t.Errorf("CliUi.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiData2StringLines(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliio := CliUi{Reader: inBuf, ErrorWriter: outBuf}

	lines := cliio.Data2StringLines()
	result := strings.Join(lines, "\n")
	if result != inputMsg {
		t.Errorf("CliUi.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestLineEnding(t *testing.T) {
	result := LineEnding()
	if result != lineEnding {
		t.Errorf("LineEnding() = [\"%s\"], want [\"%s\"].", result, lineEnding)
	}
}
