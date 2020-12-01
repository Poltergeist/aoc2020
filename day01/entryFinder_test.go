package day01

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	type args struct {
		sum  int
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				sum: 2020,
				list: []int{1721,
					979,
					366,
					299,
					675,
					1456},
			},
			want: 514579},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checksum(tt.args.sum, tt.args.list); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksumThree(t *testing.T) {
	type args struct {
		sum  int
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				sum: 2020,
				list: []int{1721,
					979,
					366,
					299,
					675,
					1456},
			},
			want: 241861950},	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChecksumThree(tt.args.sum, tt.args.list); got != tt.want {
				t.Errorf("ChecksumThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
