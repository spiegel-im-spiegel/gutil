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

//User-Interface for command-line tools.
//Reader/Writer assumes standard I/Os.
type CliUi struct {
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
func (c *CliUi) Input(reader io.Reader) {
	c.refresh = false
	c.Reader = reader
}

//Refresh Read buffer
func (c *CliUi) Refresh() error {
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

//New buffer stream for inputData
func (c *CliUi) NewReader() (*bytes.Reader, error) {
	if err := c.Refresh(); err != nil { //read from Reader, if not read.
		return nil, err
	}
	return bytes.NewReader(c.inputData), nil
}

//Copy inputData to new []byte
func (c *CliUi) CopyData() []byte {
	if err := c.Refresh(); err != nil { //read from Reader, if not read.
		return make([]byte, 0)
	}
	dst := make([]byte, len(c.inputData))
	if len(c.inputData) > 0 {
		copy(dst, c.inputData)
	}
	return dst
}

//Copy inputData to string
func (c *CliUi) Data2String() string {
	if err := c.Refresh(); err != nil { //read from Reader, if not read.
		return ""
	}
	return bytes.NewBuffer(c.inputData).String()
}

//Copy inputData to strings (split by line-ending).
func (c *CliUi) Data2StringLines() []string {
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
func (c *CliUi) Output(val ...interface{}) error {
	return c.doOutput(c.Writer, val)
}

//Output to Writer stream (add line-ending).
func (c *CliUi) Outputln(val ...interface{}) error {
	return c.doOutputln(c.Writer, val)
}

//Output to Writer stream ([]byte data).
func (c *CliUi) OutputBytes(data []byte) error {
	writer := bufio.NewWriter(c.Writer)
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return writer.Flush()
}

//Output to ErrorWriter stream.
func (c *CliUi) OutputErr(val ...interface{}) error {
	return c.doOutput(c.ErrorWriter, val)
}

//Output to ErrorWriter stream (add line-ending).
func (c *CliUi) OutputErrln(val ...interface{}) error {
	return c.doOutputln(c.ErrorWriter, val)
}

//Output to ErrorWriter stream.
func (c *CliUi) doOutput(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprint(writer, val...)
	return err
}

//Output to ErrorWriter stream (add line-ending).
func (c *CliUi) doOutputln(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprintln(writer, val...)
	return err
}
