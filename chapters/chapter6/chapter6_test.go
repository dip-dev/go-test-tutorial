package chapter6

import (
	"testing"
)

func TestPrintString(t *testing.T) {
	cases := map[string]struct {
		arg string
	}{
		"print string 1": {
			arg: "為せば成る",
		},
		"print string 2": {
			arg: "為さねば成らぬ",
		},
		"print string 3": {
			arg: "何事も成らぬは人の為さぬなり",
		},
	}
	for tt, tc := range cases {
		t.Run(tt, func(t *testing.T) {
			dispArg(tc.arg)
		})
	}
}

func TestPrintInt(t *testing.T) {
	cases := map[string]struct {
		arg int
	}{
		"年": {
			arg: 1997,
		},
		"月": {
			arg: 3,
		},
		"日": {
			arg: 14,
		},
	}
	for tt, tc := range cases {
		tc := tc
		t.Run(tt, func(t *testing.T) {
			t.Parallel()
			dispArg(tc.arg)
		})
	}
}
