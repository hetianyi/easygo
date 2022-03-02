// Copyright (C) 2019 tisnyo <tisnyo@gmail.com>.
//
// This file simulate try-catch structure.

package base

// Try simulates try catch
func Try(f func(), catcher func(e interface{})) {
	defer func() {
		if err := recover(); err != nil && catcher != nil {
			catcher(err)
		}
	}()
	f()
}
