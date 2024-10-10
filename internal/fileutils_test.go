package internal

import (
	"testing"
)

func TestCountLinesInFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "count lines in file",
			args: args{
				filename: "test.txt",
			},
			want: 3,
		},
		{
			name: "no such file",
			args: args{
				filename: "test2.txt",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountLinesInFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountLinesInFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountLinesInFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLinesFromStdin(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "count lines from stdin",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLinesFromStdin(); got != tt.want {
				t.Errorf("CountLinesFromStdin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateChecksum(t *testing.T) {
	type args struct {
		filename  string
		algorithm string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "calculate md5",
			args: args{
				filename:  "test.txt",
				algorithm: "md5",
			},
			want:    "bfb77520994c313d1abff83000f19dc3",
			wantErr: false,
		},
		{
			name: "calculate sha1",
			args: args{
				filename:  "test.txt",
				algorithm: "sha1",
			},
			want:    "9e17c07645660e62ccbfc742c38db923701abc11",
			wantErr: false,
		},
		{
			name: "calculate sha256",
			args: args{
				filename:  "test.txt",
				algorithm: "sha256",
			},
			want:    "ad53e8806d17c82d38902738d1d47d96bddaade27513466322efa0f793149dd0",
			wantErr: false,
		},
		{
			name: "unsupported algorithm",
			args: args{
				filename:  "test.txt",
				algorithm: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "no such file",
			args: args{
				filename:  "test2.txt",
				algorithm: "md5",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateChecksum(tt.args.filename, tt.args.algorithm)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateChecksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksumFromStdin(t *testing.T) {
	type args struct {
		algorithm string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "calculate md5",
			args: args{
				algorithm: "md5",
			},
			want: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name: "calculate sha1",
			args: args{
				algorithm: "sha1",
			},
			want: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			name: "calculate sha256",
			args: args{
				algorithm: "sha256",
			},
			want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChecksumFromStdin(tt.args.algorithm); got != tt.want {
				t.Errorf("ChecksumFromStdin() = %v, want %v", got, tt.want)
			}
		})
	}
}
