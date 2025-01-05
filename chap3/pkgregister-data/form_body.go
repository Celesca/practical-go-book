package pkgregister

import "io"

type pkgData struct {
	Name     string
	Version  string
	Filename string
	Bytes    io.Reader
}
