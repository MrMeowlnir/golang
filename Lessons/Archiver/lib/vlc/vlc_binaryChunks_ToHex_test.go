package vlc

import (
	"reflect"
	"testing"
)

func TestBinaryChunk_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs binaryChunks
		want hexChunks
	}{
		{
			name: "base test",
			bcs: binaryChunks{"00100000", "00110000", "00111100", "00011000", "01110100"},
			want: hexChunks{"20", "30", "3C", "18", "74"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got:=tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryChunks.ToHex()=%v, want - %v", got, tt.want)
			}
		})
	}
}
