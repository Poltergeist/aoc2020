package passwordCheck

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				list: []string{
					"1-3 a: abcde",
					"1-3 b: cdefg",
					"2-9 c: ccccccccc",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check(tt.args.list); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
