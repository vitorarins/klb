package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

type (
	testTbl struct {
		lookFor     string
		outExpected string
		errExpected string
		errStr      string
	}
)

func dotest(t *testing.T, test testTbl) {
	testdir := "/src/github.com/NeowayLabs/klb/cmd/klb/testcases/"
	nashpath := os.Getenv("GOPATH") + testdir
	os.Setenv("NASHPATH", nashpath)

	var (
		stdout, stderr bytes.Buffer
	)

	err := doc(&stdout, &stderr, []string{"doc", test.lookFor})

	if err != nil {
		if test.errStr != "" {
			if err.Error() != test.errStr {
				t.Fatalf("Expected error '%s', but got '%s'",
					test.errStr, err.Error())
			}
		} else {
			t.Fatal(err)
		}
	}

	gotOut := string(stdout.Bytes())
	gotErr := string(stderr.Bytes())

	if test.outExpected != gotOut {
		t.Fatalf("Stdout differs: '%s' != '%s'", test.outExpected, gotOut)
	}

	if test.errExpected != gotErr {
		t.Fatalf("Stderr differs: '%s' != '%s'", test.errExpected, gotErr)
	}
}

func TestDoc(t *testing.T) {
	for _, test := range []testTbl{
		{
			"somepkg.somefn",
			`fn somefn(a, b, c, d)
	somefn is a testcase function
	multiples comments
	etc
	etc
`, ``, "",
		},
		// Test empty pkg and func
		{
			"",
			"",
			fmt.Sprintf("Usage: %s doc <package>.<fn name>\n", os.Args[0]),
			"",
		},

		// test non existent package
		{
			"blahbleh.bli",
			"",
			"",
			"Function blahbleh.bli not found",
		},

		{
			"a.a",
			`fn a()
`, ``, "",
		},

		{
			"a.b",
			`fn b(a)
	bleh
`,
			``, "",
		},
		{
			"a.c",
			`fn c(a, b)
	carrr
`,
			``, "",
		},
	} {
		dotest(t, test)
	}
}