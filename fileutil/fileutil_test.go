package fileutil_test

import (
	"reflect"
	"testing"

	"github.com/Ibotta/go-commons/sopstool/fileutil"
)

func TestNormalizeToPlaintextFile(t *testing.T) {
	type args struct {
		fn string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "With the segment", args: args{fn: "filename.sops.yaml"}, want: "filename.yaml"},
		{name: "Without the segment", args: args{fn: "filename.yaml"}, want: "filename.yaml"},
		{name: "With the segment twice", args: args{fn: "filename.sops.something.sops.yaml"}, want: "filename.sops.something.yaml"},
		{name: "With the segment last", args: args{fn: "something.bin.sops"}, want: "something.bin"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileutil.NormalizeToPlaintextFile(tt.args.fn); got != tt.want {
				t.Errorf("NormalizeToPlaintextFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeToSopsFile(t *testing.T) {
	type args struct {
		fn string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "With the segment", args: args{fn: "filename.sops.yaml"}, want: "filename.sops.yaml"},
		{name: "Without the segment", args: args{fn: "filename.yaml"}, want: "filename.sops.yaml"},
		{name: "With the segment twice", args: args{fn: "filename.sops.something.sops.yaml"}, want: "filename.sops.something.sops.yaml"},
		{name: "With the segment in the wrong place", args: args{fn: "filename.sops.something.yaml"}, want: "filename.sops.something.sops.yaml"},
		{name: "Ends with segment", args: args{fn: "filename.yaml.sops"}, want: "filename.yaml.sops"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileutil.NormalizeToSopsFile(tt.args.fn); got != tt.want {
				t.Errorf("NormalizeToSopsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListIndexOf(t *testing.T) {
	type args struct {
		files []string
		fn    string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found in first",
			args: args{
				files: []string{"filename.ext", "filename2.ext"},
				fn:    "filename.ext",
			},
			want: 0,
		},
		{
			name: "found in last",
			args: args{
				files: []string{"filename.ext", "filename2.ext"},
				fn:    "filename2.ext",
			},
			want: 1,
		},
		{
			name: "not found",
			args: args{
				files: []string{"filename.ext", "filename2.ext"},
				fn:    "different.ext",
			},
			want: -1,
		},
		{
			name: "normalized",
			args: args{
				files: []string{"filename.ext", "filename2.ext"},
				fn:    "filename.sops.ext",
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileutil.ListIndexOf(tt.args.files, tt.args.fn); got != tt.want {
				t.Errorf("ListIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeOrAllFiles(t *testing.T) {
	type args struct {
		args     []string
		encFiles []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "all files",
			args: args{
				args:     []string{},
				encFiles: []string{"filename.ext", "filename2.ext", "filename3.ext"},
			},
			want:    []string{"filename.ext", "filename2.ext", "filename3.ext"},
			wantErr: false,
		},
		{
			name: "one file",
			args: args{
				args:     []string{"filename.ext"},
				encFiles: []string{"filename.ext", "filename2.ext", "filename3.ext"},
			},
			want:    []string{"filename.ext"},
			wantErr: false,
		},
		{
			name: "one file normalized",
			args: args{
				args:     []string{"filename.sops.ext"},
				encFiles: []string{"filename.ext", "filename2.ext", "filename3.ext"},
			},
			want:    []string{"filename.ext"},
			wantErr: false,
		},
		{
			name: "two files",
			args: args{
				args:     []string{"filename.ext", "filename2.ext"},
				encFiles: []string{"filename.ext", "filename2.ext", "filename3.ext"},
			},
			want:    []string{"filename.ext", "filename2.ext"},
			wantErr: false,
		},
		{
			name: "missing files",
			args: args{
				args:     []string{"filename.ext", "different.ext"},
				encFiles: []string{"filename.ext", "filename2.ext", "filename3.ext"},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileutil.SomeOrAllFiles(tt.args.args, tt.args.encFiles)
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeOrAllFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeOrAllFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
