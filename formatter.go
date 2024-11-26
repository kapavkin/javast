package javast

import (
	"io"
	"strings"
)

const (
	Identation string = "    "
	LineLength int    = 80
)

type Formatter struct {
	Writer  io.Writer
	Options FormatterOptions
	state   FormatterState
}

func (f *Formatter) Write(p []byte) (int, error) {
	writer, options := f.Writer, f.Options
	written := 0
	switch f.state.LastToken {
	case ".":
		if n, err := writer.Write(p); err != nil {
			return written + n, err
		} else {
			written += n
		}
	case "<":
		if n, err := writer.Write(p); err != nil {
			return written + n, err
		} else {
			written += n
		}
	case ";":
		switch string(p) {
		case "}":
			f.state.Identation = strings.Replace(f.state.Identation, options.Identation, "", 1)
			if n, err := writer.Write([]byte{'\n'}); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write([]byte(f.state.Identation)); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		default:
			if n, err := writer.Write([]byte{'\n'}); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write([]byte(f.state.Identation)); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		}
	case "{":
		switch string(p) {
		case "}":
			f.state.Identation = strings.Replace(f.state.Identation, options.Identation, "", 1)
			if n, err := writer.Write([]byte{'\n'}); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write([]byte(f.state.Identation)); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		default:
			f.state.Identation += options.Identation
			if n, err := writer.Write([]byte{'\n'}); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write([]byte(f.state.Identation)); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		}
	default:
		switch string(p) {
		case ".":
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		case ">":
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		case "}":
			f.state.Identation = strings.Replace(f.state.Identation, options.Identation, "", 1)
			if n, err := writer.Write([]byte{'\n'}); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write([]byte(f.state.Identation)); err != nil {
				return written + n, err
			} else {
				written += n
			}
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		default:
			if len(f.state.LastToken) > 0 {
				if n, err := writer.Write([]byte{' '}); err != nil {
					return written + n, err
				} else {
					written += n
				}
			}
			if n, err := writer.Write(p); err != nil {
				return written + n, err
			} else {
				written += n
			}
		}
	}
	f.state.LastToken = string(p)
	return written, nil
}

type FormatterOptions struct {
	Identation string
	LineLength int
}

type FormatterState struct {
	Identation string
	LastToken  string
}
