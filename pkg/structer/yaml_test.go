// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import (
	"reflect"
	"testing"
)

func TestYamlToJsonDecoder_Decode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		e       YamlToJsonDecoder
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.Decode(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("YamlToJsonDecoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YamlToJsonDecoder.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlToJsonDecoder_convert(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		e    YamlToJsonDecoder
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.convert(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YamlToJsonDecoder.convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
