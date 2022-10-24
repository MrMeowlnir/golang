package vlc

import "testing"

func Test_Execute(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is World",
			want: "20 30 3C 18 77 4A E4 03 8A 09 28",
		},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Execute(tt.str); got != tt.want {
					t.Errorf("Execute() = %s, want %s", got, tt.want)
				}
			})
		}
}