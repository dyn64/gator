package main

// unit test for the config internals

import (
	"testing"

	"github.com/dyn64/gator/internal/config"
	"github.com/google/go-cmp/cmp"
)

func TestSetusername(t *testing.T) {
	cases := []struct {
		input    string
		expected config.Config
	}{
		{
			input: "test_user",
			expected: config.Config{
				CurrentUserName: "test_user",
			},
		},
		{
			input: "",
			expected: config.Config{
				CurrentUserName: "erik",
			},
		},
	}

	for _, c := range cases {
		conf, err := config.Read()
		if err != nil {
			t.Errorf("Error reading config file")
		}
		err = conf.SetUser(c.input)
		conf, err = config.Read()
		if err != nil {
			t.Errorf("Error re-reading config file")
		}
		actual := conf.CurrentUserName
		diff := cmp.Diff(actual, c.expected.CurrentUserName)
		if diff != "" {
			t.Fatalf("%s", diff)
		}
	}
}
