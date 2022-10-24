package vlc

import "testing"

func Test_prepareText(t *testing.T) {
	tests := []struct{
		name string
		str string
		want string
	}{
		{
			name: "base test",
			str: "My name is World",
			want: "!my name is !world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = #{got}, want #{tt.want}")
			}
		})
	}
}