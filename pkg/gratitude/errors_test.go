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
		expected GratitudeError
	}{
		{name: "new stackstr error", code: 400, msg: "invalid request", expected: GratitudeError{
			code:    400,
			message: "invalid request",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var e *GratitudeError
			got := NewGratitudeError(tc.code, tc.msg)

			assert.Equal(t, errs.As(got, &e), true)

			if e, ok := got.(*GratitudeError); ok {
				fmt.Println(e)
				//assert.Equal(t, e.code, tc.code)
				//assert.Equal(t, e.message, tc.msg)
			}
		})
	}
}

func Test_GratitudeErrorError(t *testing.T) {
	tt := []struct {
		name     string
		code     int
		msg      string
		expected GratitudeError
	}{
		{name: "GratitudeError.Error", code: 400, msg: "invalid request", expected: GratitudeError{
			code:    400,
			message: "invalid request",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var e *GratitudeError
			got := NewGratitudeError(tc.code, tc.msg)

			assert.Equal(t, errs.As(got, &e), true)

			if e, ok := got.(*GratitudeError); ok {
				assert.Equal(t, e.Error(), tc.msg)
			}
		})
	}
}

func Test_GratitudeErrorStatus(t *testing.T) {
	tt := []struct {
		name     string
		code     int
		msg      string
		expected GratitudeError
	}{
		{name: "GratitudeError.Status", code: 400, msg: "invalid request", expected: GratitudeError{
			code:    400,
			message: "invalid request",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var e *GratitudeError
			got := NewGratitudeError(tc.code, tc.msg)

			assert.Equal(t, errs.As(got, &e), true)

			if e, ok := got.(*GratitudeError); ok {
				assert.Equal(t, e.Status(), tc.code)
			}
		})
	}
}
