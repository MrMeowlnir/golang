package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want binaryChunks
	}{
		{
			name: "base test",
			args: args{	bStr: "00100000001100000011110000011000011101",
						chunkSize: 8,
					},
			want: binaryChunks{"00100000", "00110000", "00111100", "00011000", "01110100"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got:=splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryChunks.ToHex(%s, %d) = %v, want %v", tt.args.bStr, tt.args.chunkSize, got, tt.want)
			}
		})
	}
}
