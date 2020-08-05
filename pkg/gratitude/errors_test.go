package gratitude

import (
	"fmt"
	"testing"

	errs "errors"

	"github.com/go-playground/assert/v2"
)

func TestNew(t *testing.T) {
	tt := []struct {
		name     string
		code     int
		msg      string
		expected Error
	}{
		{name: "new stackstr error", code: 400, msg: "invalid request", expected: Error{
			code:    400,
			message: "invalid request",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var e *Error
			got := NewGratitudeError(tc.code, tc.msg)

			assert.Equal(t, errs.As(got, &e), true)

			if e, ok := got.(*Error); ok {
				fmt.Println(e)
				//assert.Equal(t, e.code, tc.code)
				//assert.Equal(t, e.message, tc.msg)
			}
		})
	}
}

func Test_ErrorError(t *testing.T) {
	tt := []struct {
		name     string
		code     int
		msg      string
		expected Error
	}{
		{name: "Error.Error", code: 400, msg: "invalid request", expected: Error{
			code:    400,
			message: "invalid request",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var e *Error
			got := NewGratitudeError(tc.code, tc.msg)

			assert.Equal(t, errs.As(got, &e), true)

			if e, ok := got.(*Error); ok {
				assert.Equal(t, e.Error(), tc.msg)
			}
		})
	}
}

func Test_ErrorStatus(t *testing.T) {
	tt := []struct {
		name     string
		code     int
		msg      string
		expected Error
	}{
		{name: "Error.Status", code: 400, msg: "invalid request", expected: Error{
			code:    400,
			message: "invalid request",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var e *Error
			got := NewGratitudeError(tc.code, tc.msg)

			assert.Equal(t, errs.As(got, &e), true)

			if e, ok := got.(*Error); ok {
				assert.Equal(t, e.Status(), tc.code)
			}
		})
	}
}
