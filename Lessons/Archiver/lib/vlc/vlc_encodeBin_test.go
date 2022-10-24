package vlc

import "testing"

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!my name",
			want: "00100000001100000011110000011000011101",
			},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := encodeBin(tt.str); got != tt.want {
						t.Errorf("encodeBin() = #{got}, want #{tt.want}")
					}
				})
			}
}