package example

import "testing"

func TestSum(t *testing.T) {
	testsets := map[string]struct {
		i    int
		j    int
		want int
	}{
		"test1": {
			i:    1,
			j:    2,
			want: 3,
		},
		"test2": {
			i:    0,
			j:    0,
			want: 0,
		},
	}

	for label, tt := range testsets {
		t.Run(label, func(t *testing.T) {
			if got := Sum(tt.i, tt.j); got != tt.want {
				t.Errorf("got = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	testsets := map[string]struct {
		i    int
		j    int
		want int
	}{
		"test1": {
			i:    3,
			j:    1,
			want: 2,
		},
		"test2": {
			i:    5,
			j:    6,
			want: 1,
		},
	}

	for label, tt := range testsets {
		t.Run(label, func(t *testing.T) {
			if got := Diff(tt.i, tt.j); got != tt.want {
				t.Errorf("got = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	testsets := map[string]struct {
		i    int
		j    int
		want int
	}{
		"test1": {
			i:    1,
			j:    2,
			want: 1,
		},
		"test2": {
			i:    2,
			j:    0,
			want: 1,
		},
		"test3": {
			i:    3,
			j:    3,
			want: 27,
		},
	}

	for label, tt := range testsets {
		t.Run(label, func(t *testing.T) {
			if got := Pow(tt.i, tt.j); got != tt.want {
				t.Errorf("got = %d, want %d", got, tt.want)
			}
		})
	}
}
