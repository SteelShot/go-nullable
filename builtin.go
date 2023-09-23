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

package nullable

// Real number alias
type (
	Int     = Any[int]
	Int8    = Any[int8]
	Int16   = Any[int16]
	Int32   = Any[int32]
	Int64   = Any[int64]
	Uint    = Any[uint]
	Uint8   = Any[uint8]
	Uint16  = Any[uint16]
	Uint32  = Any[uint32]
	Uint64  = Any[uint64]
	Float32 = Any[float32]
	Float64 = Any[float64]
)

// Arbitrary type alias
type (
	Bool   = Any[bool]
	String = Any[string]
)

// Builtin alias' alias
type (
	Byte = Uint8
	Rune = Int32
)
