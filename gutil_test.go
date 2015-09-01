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
	cliui := CliUi{}
	cliui.ResetWriter(outBuf)

	cliui.Output(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiOutputln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliui := CliUi{Writer: outBuf}

	cliui.Outputln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestCliUiOutputBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliui := CliUi{Writer: outBuf}

	cliui.OutputBytes([]byte(inputMsg))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiOutputErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliui := CliUi{}
	cliui.ResetErrorWriter(outBuf)

	cliui.OutputErr(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.OutputErr = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiOutputErrln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cliui := CliUi{ErrorWriter: outBuf}

	cliui.OutputErrln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliUi.OutputErr = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestCliUiRefresh(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	cliui := CliUi{}
	cliui.ResetReader(inBuf)
	cliui.ModeStream()
	if err := cliui.Refresh(); err != nil {
		t.Errorf("CliUi.Refresh() = \"%v\", want nil.", err)
	}
}

func TestCliUiRefreshNG(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	cliui := CliUi{}
	cliui.ResetReader(inBuf)
	cliui.ModeInteract()
	if err := cliui.Refresh(); err != ModeErrorStream {
		t.Errorf("CliUi.Refresh = error \"%v\", want error \"%v\".", err, ModeErrorStream)
	}
}

func TestCliUiNewReader(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliui := CliUi{Reader: inBuf, Writer: outBuf}

	buf := new(bytes.Buffer)
	r, err := cliui.NewReader()
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
	cliui := CliUi{Reader: inBuf, Writer: outBuf}

	result := bytes.NewBuffer(cliui.CopyData()).String()
	if result != inputMsg {
		t.Errorf("CliUi.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiData2String(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliui := CliUi{Reader: inBuf, Writer: outBuf}

	result := cliui.Data2String()
	if result != inputMsg {
		t.Errorf("CliUi.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiData2StringLines(t *testing.T) {
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	cliui := CliUi{Reader: inBuf, Writer: outBuf}

	lines := cliui.Data2StringLines()
	result := strings.Join(lines, "\n")
	if result != inputMsg {
		t.Errorf("CliUi.NewReader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestCliUiPrompt(t *testing.T) {
	prompt := "prompt>"
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cliui := CliUi{Reader: inBuf, Writer: outBuf, ErrorWriter: errBuf}
	cliui.ModeInteract()
	result, err := cliui.Prompt(prompt)
	p := errBuf.String()
	if err != nil {
		t.Errorf("CliUi.Prompt() = \"%v\", want nil.", err)
	} else {
		if p != prompt {
			t.Errorf("CliUi.Prompt(): Prompt = \"%v\", want \"%v\".", p, prompt)
		}
		if result != inputMsgs[0] {
			t.Errorf("CliUi.Prompt() = \"%s\", want \"%s\".", result, inputMsgs[0])
		}
	}
}

func TestCliUiPromptNG(t *testing.T) {
	prompt := "prompt>"
	inBuf := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cliui := CliUi{Reader: inBuf, Writer: outBuf, ErrorWriter: errBuf}
	cliui.ModeStream()
	_, err := cliui.Prompt(prompt)
	if err == nil {
		t.Errorf("CliUi.Prompt() = %v, want error.", err)
	}
}

func TestLineEnding(t *testing.T) {
	result := LineEnding()
	if result != lineEnding {
		t.Errorf("LineEnding() = [\"%s\"], want [\"%s\"].", result, lineEnding)
	}
}
