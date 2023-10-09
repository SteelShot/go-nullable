/*
 * Copyright (C) 2023 Džiugas Eiva
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package internal

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// FormatString provides fmt.FormatString implementation before Go 1.20
func FormatString(state fmt.State, verb rune) string {
	format := append([]byte{}, 37)

	for _, flag := range []int{32, 35, 43, 45, 48} {
		if state.Flag(flag) {
			format = append(format, byte(flag))
		}
	}

	if wid, ok := state.Width(); ok {
		format = strconv.AppendInt(format, int64(wid), 10)
	}

	if prec, ok := state.Precision(); ok {
		format = strconv.AppendInt(append(format, 46), int64(prec), 10)
	}

	return string(utf8.AppendRune(format, verb))
}
