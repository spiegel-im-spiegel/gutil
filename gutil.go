/**
 * Go-lang Miscellaneous Utility Library
 *
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed
 */

//Go-lang Miscellaneous Utility Library
package gutil

import (
	"runtime"
)

const (
	//Line-Ending for Windows
	LineEndingWindows = "\r\n"
	//Line-Ending for not-Windows (Linux, OS X,...)
	LineEndingNowin   = "\n"
)

//LineEnding is function that getting line-ending.
func LineEnding() string {
	if runtime.GOOS == "windows" { //maybe Windows, but...
		return LineEndingWindows
	} else {
		return LineEndingNowin
	}
}
