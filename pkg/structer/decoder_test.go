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

func TestNewDecoderFactory(t *testing.T) {
	tests := []struct {
		name string
		want DecoderFactory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecoderFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecoderFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoderFactory_Create(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		f       DecoderFactory
		args    args
		want    ToJsonDecoder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Create(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecoderFactory.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecoderFactory.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoderFactory_find(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		f       DecoderFactory
		args    args
		want    ToJsonDecoder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.find(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecoderFactory.find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecoderFactory.find() = %v, want %v", got, tt.want)
			}
		})
	}
}
