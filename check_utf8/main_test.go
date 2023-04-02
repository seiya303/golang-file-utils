package main

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsEncodedInUtf8(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name       string
		fileName   string
		args       args
		wantIsUtf8 bool
		wantErr    error
	}{
		{
			name:     "SUCCESS(when encoded in utf-8)",
			fileName: "./testdata/utf8_sample.txt",
			args: args{
				file: nil,
			},
			wantIsUtf8: true,
			wantErr:    nil,
		},
		{
			name:     "SUCCESS(when encoded in non-utf-8)",
			fileName: "./testdata/utf16_sample.txt",
			args: args{
				file: nil,
			},
			wantIsUtf8: false,
			wantErr:    nil,
		},
		{
			name:     "FAIL(when file doesn't exist)",
			fileName: "./testdata/non_existent.txt",
			args: args{
				file: nil,
			},
			wantIsUtf8: false,
			wantErr:    errNonExistentFile,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dir, _ := os.Getwd()
			filePath := path.Join(dir, tt.fileName)
			fp, _ := os.Open(filePath)
			defer fp.Close()

			tt.args.file = io.Reader(fp)

			gotIsUtf8, err := isEncodedInUtf8(tt.args.file)
			if err != tt.wantErr {
				t.Errorf("IsEncodedInUtf8() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(gotIsUtf8, tt.wantIsUtf8); diff != "" {
				t.Errorf("gotIsUtf8 value is mismatch (-IsEncodedInUtf8() +want):\n%s", diff)
			}
		})
	}
}
