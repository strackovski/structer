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

func TestDefault(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_Initialisms(t *testing.T) {
	tests := []struct {
		name string
		c    config
		want map[string]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Initialisms(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("config.Initialisms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_OutputDir(t *testing.T) {
	tests := []struct {
		name string
		c    config
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.OutputDir(); got != tt.want {
				t.Errorf("config.OutputDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_IntWordMap(t *testing.T) {
	tests := []struct {
		name string
		c    config
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IntWordMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("config.IntWordMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
