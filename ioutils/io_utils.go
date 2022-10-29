package ioutils

import "io"

var CloserFunc = func(closer io.Closer) {
	err := closer.Close()
	if err != nil {

	}
}
