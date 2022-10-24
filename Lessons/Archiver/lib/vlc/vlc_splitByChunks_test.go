package vlc

import "testing"

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
			for i, got := range splitByChunks(tt.args.bStr, tt.args.chunkSize) {
				if got != tt.want[i] {
					t.Errorf("splitByChunks(). Chunk N#{i}: got - #{got}, want - #{tt.want[i]}")
				}
			}
		})
	}
}
