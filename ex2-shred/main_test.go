package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func Test_shred(t *testing.T) {
	type args struct {
		path       string
		iterations uint
	}
	tests := []struct {
		name         string
		args         args
		createTarget bool
		wantErr      error
	}{
		{
			name: "shred successfully",
			args: args{
				path:       "tests/test.txt",
				iterations: 5,
			},
			createTarget: true,
			wantErr:      nil,
		},

		{
			name: "empty path",
			args: args{
				path:       "",
				iterations: 5,
			},
			wantErr: ErrEmptyPath,
		},
		{
			name: "0 iterations",
			args: args{
				path:       "tests/test.txt",
				iterations: 0,
			},
			wantErr: ErrNoIterration,
		},
		{
			name: "impossible to shred a directory",
			args: args{
				path:       "tests/",
				iterations: 5,
			},
			wantErr: ErrPathIsDir,
		},
		{
			name: "inexistent file",
			args: args{
				path:       "tests/non-existent.txt",
				iterations: 5,
			},
			wantErr: os.ErrNotExist,
		},
	}
	for _, tt := range tests {
		fileContent := []byte("test")
		if tt.createTarget {
			err := os.WriteFile(tt.args.path, []byte("test"), 0655)
			if err != nil {
				t.Fatal(err)
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			err := shred(tt.args.path, tt.args.iterations)
			if (err != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("shred() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				afterShred, err := os.ReadFile(tt.args.path)
				if err != nil {
					t.Fatal(err)
				}
				if bytes.Equal(afterShred, fileContent) {
					t.Error("no shred happened")
				}
			}
		})
	}
}
