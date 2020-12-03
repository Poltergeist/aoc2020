package day03

import "testing"

func TestTreeCount(t *testing.T) {

	type args struct {
		treeMap []string
		m       Move
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				treeMap: []string{
					"..##.......",
					"#...#...#..",
					".#....#..#.",
					"..#.#...#.#",
					".#...##..#.",
					"..#.##.....",
					".#.#.#....#",
					".#........#",
					"#.##...#...",
					"#...##....#",
					".#..#...#.#",
				},
				m: Move{
					R: 3,
					D: 1,
				},
			},
			want: 7,
		},
		{
			name: "happy path",
			args: args{
				treeMap: []string{
					"..##.......",
					"#...#...#..",
					".#....#..#.",
					"..#.#...#.#",
					".#...##..#.",
					"..#.##.....",
					".#.#.#....#",
					".#........#",
					"#.##...#...",
					"#...##....#",
					".#..#...#.#",
				},
				m: Move{
					R: 1,
					D: 2,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeCount(tt.args.treeMap, tt.args.m); got != tt.want {
				t.Errorf("TreeCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
