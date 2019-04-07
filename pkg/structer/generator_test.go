// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import (
	"io"
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		name string
		want TypeGenerator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_Generate(t *testing.T) {
	type args struct {
		input      io.Reader
		structName string
		pkgName    string
	}
	tests := []struct {
		name    string
		d       typeGenerator
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Generate(tt.args.input, tt.args.structName, tt.args.pkgName)
			if (err != nil) != tt.wantErr {
				t.Errorf("typeGenerator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("typeGenerator.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_generate(t *testing.T) {
	type args struct {
		input      io.Reader
		structName string
		pkgName    string
	}
	tests := []struct {
		name    string
		d       typeGenerator
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.generate(tt.args.input, tt.args.structName, tt.args.pkgName)
			if (err != nil) != tt.wantErr {
				t.Errorf("typeGenerator.generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("typeGenerator.generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_generateTypes(t *testing.T) {
	type args struct {
		obj   map[string]interface{}
		depth int
	}
	tests := []struct {
		name string
		d    typeGenerator
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.generateTypes(tt.args.obj, tt.args.depth); got != tt.want {
				t.Errorf("typeGenerator.generateTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_fmtFieldName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		d    typeGenerator
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.fmtFieldName(tt.args.s); got != tt.want {
				t.Errorf("typeGenerator.fmtFieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_lintFieldName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		d    typeGenerator
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.lintFieldName(tt.args.name); got != tt.want {
				t.Errorf("typeGenerator.lintFieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_typeForValue(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		d    typeGenerator
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.typeForValue(tt.args.value); got != tt.want {
				t.Errorf("typeGenerator.typeForValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_disambiguateFloatInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		d    typeGenerator
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.disambiguateFloatInt(tt.args.value); got != tt.want {
				t.Errorf("typeGenerator.disambiguateFloatInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_stringifyFirstChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		d    typeGenerator
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.stringifyFirstChar(tt.args.str); got != tt.want {
				t.Errorf("typeGenerator.stringifyFirstChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeGenerator_generateStruct(t *testing.T) {
	type args struct {
		obj   map[string]interface{}
		depth int
	}
	tests := []struct {
		name string
		d    typeGenerator
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.generateStruct(tt.args.obj, tt.args.depth); got != tt.want {
				t.Errorf("typeGenerator.generateStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
