/**
 * Go-lang Miscellaneous Utility Library
 *
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed
 */

//Go-lang Miscellaneous Utility Library
package gutil

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

//Context for command-line tools.
//Reader/Writer assumes standard I/Os.
type CliContext struct {
	//Refresh flag (for Reader stream)
	refresh bool

	//Data from Reader stream
	inputData []byte

	//Input Stream.
	Reader io.Reader

	//Output streams.
	//(Error information is written to ErrorWriter)
	Writer, ErrorWriter io.Writer
}

//Reset Reader stream.
func (c *CliContext) Input(reader io.Reader) {
	c.refresh = false
	c.Reader = reader
}

//Refresh Read buffer
func (c *CliContext) Refresh() error {
	if !c.refresh {
		c.refresh = true
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(c.Reader); err != nil {
			return err
		}
		c.inputData = buf.Bytes()
	}
	return nil
}

//New Stream for inputData
func (c *CliContext) NewReader() (io.Reader, error) {
	if err := c.Refresh(); err != nil { //read from Reader, if not read.
		return nil, err
	}
	return bytes.NewReader(c.inputData), nil
}

//New Stream for inputData
func (c *CliContext) CopyData() []byte {
	if err := c.Refresh(); err != nil { //read from Reader, if not read.
		return make([]byte, 0)
	}
	dst := make([]byte, len(c.inputData))
	if len(c.inputData) > 0 {
		copy(dst, c.inputData)
	}
	return dst
}

//New Stream for inputData
func (c *CliContext) Data2String() string {
	if err := c.Refresh(); err != nil { //read from Reader, if not read.
		return ""
	}
	return bytes.NewBuffer(c.inputData).String()
}

//New Stream for inputData
func (c *CliContext) Data2StringLines() []string {
	if err := c.Refresh(); err != nil { //read from Reader, if not read.
		return []string{}
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(c.inputData))
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

//Output to Writer stream.
func (c *CliContext) Output(a ...interface{}) error {
	return c.doOutput(c.Writer, a)
}

//Output to ErrorWriter stream.
func (c *CliContext) OutputErr(a ...interface{}) error {
	return c.doOutput(c.ErrorWriter, a)
}

//Output to ErrorWriter stream.
func (c *CliContext) doOutput(writer io.Writer, a []interface{}) error {
	_, err := fmt.Fprint(writer, a...)
	return err
}
