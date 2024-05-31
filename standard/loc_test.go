package standard

import (
	"testing"
	"time"
)

func TestTz(t *testing.T) {
	tz()
}

func TestCompare(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	n := time.Now()
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1和t2相同",
			args: args{
				t1: n.AddDate(0, 0, 1),
				t2: n.Add(24 * time.Hour),
			},
			want: 0,
		},
		{
			name: "t1比t2小",
			args: args{
				t1: n.AddDate(0, 0, 1),
				t2: n.AddDate(0, 0, 2),
			},
			want: -1,
		},
		{
			name: "t1比t2大",
			args: args{
				t1: n.AddDate(0, 0, 5),
				t2: n.AddDate(0, 0, 2),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t1.Compare(tt.args.t2); got != tt.want {
				t.Errorf("t1.Compare(t2) = %v, want = %v", got, tt.want)
			}
		})
	}
}
