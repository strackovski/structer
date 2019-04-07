// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import "testing"

func TestFile(t *testing.T) {
	type args struct {
		input      []byte
		outputName string
		outputPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := File(tt.args.input, tt.args.outputName, tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("File() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.e)
		})
	}
}
