package main

import "io"

type xmlWriter interface {
	WriteXML(w io.Writer) error
}

func StreamXML(v interface{}, w io.Writer) error {
	if xw, ok := v.(xmlWriter); ok {
		return xw.WriteXML(w)
	}
	return encodeToXML(v, w)
}
