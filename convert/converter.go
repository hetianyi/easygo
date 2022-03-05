// Copyright (C) 2019 tisnyo <tisnyo@gmail.com>.
//
// package convert helps you convert types to another type.
// license that can be found in the LICENSE file.

package convert

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

const (
	Byte    byte    = 0
	Int     int     = 0
	Uint    uint    = 0
	Int8    int8    = 0
	Uint8   uint8   = 0
	Int16   int16   = 0
	Uint16  uint16  = 0
	Int32   int32   = 0
	Uint32  uint32  = 0
	Int64   int64   = 0
	Uint64  uint64  = 0
	Float32 float32 = 0
	Float64 float64 = 0
)

type Number interface {
	byte | int | uint | int8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func NumberToStr[T Number](t T) string {
	var v = any(t)
	switch v.(type) {
	case byte, int, int8, int16, int32, int64:
		return strconv.FormatInt(int64(t), 10)
	case uint, uint16, uint32, uint64:
		return strconv.FormatUint(uint64(t), 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(t), 'f', -1, 64)
	default:
		return "" // never reached here
	}
}

func StrToNumber[T Number](s string, t T) (T, error) {
	switch v := any(t).(type) {
	case int:
		r, err := strconv.Atoi(s)
		return any(r).(T), err
	case uint:
		r, err := strconv.ParseUint(s, 10, 32)
		return any(uint(r)).(T), err
	case int8:
		r, err := strconv.ParseInt(s, 10, 8)
		return any(int8(r)).(T), err
	case uint8:
		r, err := strconv.ParseUint(s, 10, 8)
		return any(uint8(r)).(T), err
	case int16:
		r, err := strconv.ParseInt(s, 10, 16)
		return any(int16(r)).(T), err
	case uint16:
		r, err := strconv.ParseUint(s, 10, 16)
		return any(uint16(r)).(T), err
	case int32:
		r, err := strconv.ParseInt(s, 10, 32)
		return any(int32(r)).(T), err
	case uint32:
		r, err := strconv.ParseUint(s, 10, 32)
		return any(uint32(r)).(T), err
	case int64:
		r, err := strconv.ParseInt(s, 10, 64)
		return any(r).(T), err
	case uint64:
		r, err := strconv.ParseUint(s, 10, 64)
		return any(uint64(r)).(T), err
	case float32:
		r, err := strconv.ParseFloat(s, 32)
		return any(float32(r)).(T), err
	case float64:
		r, err := strconv.ParseFloat(s, 64)
		return any(r).(T), err
	default:
		panic(fmt.Sprintf("conversion not supported for type %v", v))
	}
}

// BoolToStr converts bool to string.
func BoolToStr(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

// StrToBool converts string to bool.
func StrToBool(value string) (bool, error) {
	return strconv.ParseBool(strings.ToLower(value))
}

// Length2Bytes converts an int64 value to a byte array.
func Length2Bytes(len int64, buffer []byte) []byte {
	binary.BigEndian.PutUint64(buffer, uint64(len))
	return buffer
}

// Bytes2Length converts a byte array to an int64 value.
func Bytes2Length(ret []byte) int64 {
	return int64(binary.BigEndian.Uint64(ret))
}
