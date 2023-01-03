package main

import (
	"flag"
	"github.com/SandersN79/parallelChecker/core"
	"os"
	"strings"
	"testing"
)

func Test_duplicatesCheck(t *testing.T) {
	path, _ := os.Getwd()
	path = strings.Replace(path, "cmd/checker", "testData/", -1)
	tests := []struct {
		name    string
		want    string
		wantErr bool
		osArgs  []string // The command arguments used for this test
	}{
		{
			"no duplicates",
			"",
			false,
			[]string{"test", "-path", path, "-files", "TestProcessEligibleChannel2_1_TestProcessEligibleChannel2_1_CODES.csv"},
		},
		{
			"duplicates",
			"context canceled",
			true,
			[]string{"test", "-path", path, "-files", "*"},
		},
	}
	// Iterating over the previous test slice
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOsArgs := os.Args
			defer func() {
				os.Args = actualOsArgs                                           // Restoring the original os.Args reference
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Resetting the Flag command line. So that we can parse flags again
			}()
			os.Args = tt.osArgs // Setting the specific command args for this test
			inputs, err := getCliInputs()
			err = core.DuplicateCheck(inputs.filePaths)
			if (err != nil) != tt.wantErr {
				t.Errorf("DuplicateCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if err.Error() != tt.want {
					t.Errorf("DuplicateCheck() = %v, want %v", err.Error(), tt.want)
				}
			}
		})
	}
}
