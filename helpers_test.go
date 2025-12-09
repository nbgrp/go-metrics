// SPDX-License-Identifier: BSD-3-Clause

package metrics

import (
	"errors"
	"testing"
	"time"
)

func TestSeconds(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		duration time.Duration
		want     float64
	}{
		{
			name:     "zero duration",
			duration: 0,
			want:     0,
		},
		{
			name:     "1 second",
			duration: time.Second,
			want:     1.0,
		},
		{
			name:     "1 millisecond",
			duration: time.Millisecond,
			want:     0.001,
		},
		{
			name:     "1 microsecond",
			duration: time.Microsecond,
			want:     0.000001,
		},
		{
			name:     "1 nanosecond",
			duration: time.Nanosecond,
			want:     0.000000001,
		},
		{
			name:     "1 minute",
			duration: time.Minute,
			want:     60.0,
		},
		{
			name:     "1 hour",
			duration: time.Hour,
			want:     3600.0,
		},
		{
			name:     "500 milliseconds",
			duration: 500 * time.Millisecond,
			want:     0.5,
		},
		{
			name:     "1.5 seconds",
			duration: 1500 * time.Millisecond,
			want:     1.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Seconds(tt.duration)
			if got != tt.want {
				t.Errorf("Seconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinceSeconds(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		started time.Time
		check   func(float64) bool
	}{
		{
			name:    "now should be close to zero",
			started: time.Now(),
			check: func(result float64) bool {
				return result >= 0 && result < 0.01
			},
		},
		{
			name:    "one second ago",
			started: time.Now().Add(-time.Second),
			check: func(result float64) bool {
				return result >= 1.0 && result < 1.1
			},
		},
		{
			name:    "100 milliseconds ago",
			started: time.Now().Add(-100 * time.Millisecond),
			check: func(result float64) bool {
				return result >= 0.1 && result < 0.2
			},
		},
		{
			name:    "5 seconds ago",
			started: time.Now().Add(-5 * time.Second),
			check: func(result float64) bool {
				return result >= 5.0 && result < 5.1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := SinceSeconds(tt.started)
			if !tt.check(got) {
				t.Errorf("SinceSeconds() = %v, failed check", got)
			}
		})
	}
}

func TestIsError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "nil error",
			err:  nil,
			want: "0",
		},
		{
			name: "non-nil error",
			err:  errors.New("some error"),
			want: "1",
		},
		{
			name: "empty error message",
			err:  errors.New(""),
			want: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := IsError(tt.err)
			if got != tt.want {
				t.Errorf("IsError() = %v, want %v", got, tt.want)
			}
		})
	}
}
