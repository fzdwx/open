package cons

import "errors"

var (
	ClipboardEmptyError = errors.New("clipboard is empty")
	StdinEmptyError     = errors.New("stdin is empty")
)
