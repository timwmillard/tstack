package admin

import (
	"fmt"

	"github.com/a-h/templ"
)

//go:generate templ generate

func safeURL(format string, a ...any) templ.SafeURL {
	return templ.SafeURL(fmt.Sprintf(format, a...))
}

func url(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
