package errors

import (
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

// Error is a wrapper around parent error with an optional title and unique id.
type Error struct {
	ID     string `json:"id"`
	Title  string `json:"title,omitempty"`
	Parent error  `json:"parent,omitempty"`
}

// From creates a new Error from parent err.
func From(err error) *Error {
	return &Error{
		ID:     uuid.NewV4().String(),
		Parent: err,
	}
}

// FromWithTitle creates a new Error from parent err with title.
func FromWithTitle(title string, err error) *Error {
	e := From(err)
	e.Title = title
	return e
}

// Error returns the string representation of parent error with title and parent err
// if they are not cleared out.
func (e *Error) Error() string {
	if e.Title == "" || e.Parent == nil {
		return e.ID
	}
	return fmt.Sprintf("(%s) %s: %s", e.ID, e.Title, e.Parent)
}

// Log logs Error into stdout and returns it.
func (e *Error) Log() *Error {
	log.Println(e)
	return e
}

// Logger is a widely known logger interface.
type Logger interface {
	Log(v ...interface{})
}

// LogWith logs Error by using given logger and returns it.
func (e *Error) LogWith(l Logger) *Error {
	l.Log(e)
	return e
}

// Clear clears title and parent err info from Error and returns it.
// this is useful to only send the error's id to end user while hiding confidential info.
func (e *Error) Clear() *Error {
	e.Title = ""
	e.Parent = nil
	return e
}
