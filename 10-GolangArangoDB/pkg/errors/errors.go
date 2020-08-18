package errors

import (
	"encoding/json"
	"runtime"
)


// ErrorJSON contains more information than standard go exception, and allows more descriptive exception messages.
type ErrorJSON struct {
	Cause    string `json:"cause"`
	Message  string `json:"msg"`
	File     string `json:"file"`
	Function string `json:"method"`
	Line     int    `json:"line"`
}

func New(e error, customMessage string) error {
	pc, f, l, _ := runtime.Caller(1)

	errorMessage := "**NIL**"
	if e != nil {
		errorMessage = e.Error()
	}
	return &ErrorJSON{
		Cause:    errorMessage,
		Message:  customMessage,
		File:     f,
		Function: runtime.FuncForPC(pc).Name(),
		Line:     l,
	}
}

// Error Required method to implement exception
func (e *ErrorJSON) Error() string {
	if e == nil {
		return "** NIL **"
	}

	out, err := json.Marshal(e)
	if err != nil {
		return e.Cause
	}

	return string(out)
}

func (e *ErrorJSON) ExternalError() interface{} {
	if e == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{"cause": e.Cause, "msg": e.Message}
}

func (e *ErrorJSON) InternalError() interface{} {
	if e == nil {
		return map[string]interface{}{}
	}
	out, err := json.Marshal(e)
	if err != nil {
		// TODO not a good idea to hide the error
		return map[string]interface{}{}
	}

	return out
}