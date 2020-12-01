package misc

import (
	"reflect"
	"testing"
)

func TestConvertStringsToInt(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "happy path",
			args: args{
				list: []string{"1", "2", "3", "\n"},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringsToInt(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringsToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
