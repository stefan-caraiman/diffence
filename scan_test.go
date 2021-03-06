package diffence

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_ScanDiffs(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "ScanDiffs() split fn with Commit ID",
			args: args{r: bytes.NewReader(
				[]byte(
					"17949087b8e0c9179345e8dbb7b6705b49c93c77 Adds results logger" +
						"\n" +
						"diff --git a/README.md b/README.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644" +
						"\n" +
						"22949087b8e0c9179345e8dbb7b6705b49c93c88 Does something else" +
						"\n" +
						"diff --git a/TODO.md b/TODO.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			)},
			want: [][]byte{
				[]byte(
					"17949087b8e0c9179345e8dbb7b6705b49c93c77 Adds results logger" +
						"\n" +
						"diff --git a/README.md b/README.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
				[]byte(
					"22949087b8e0c9179345e8dbb7b6705b49c93c88 Does something else" +
						"\n" +
						"diff --git a/TODO.md b/TODO.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			},
		},
		{
			name: "ScanDiffs() split fn",
			args: args{r: bytes.NewReader(
				[]byte(
					"diff --git a/NOHASH.md b/NOHASH.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644" +
						"\n" +
						"diff --git a/STILL.md b/STILL.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			)},
			want: [][]byte{
				[]byte(
					"diff --git a/NOHASH.md b/NOHASH.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
				[]byte(
					"diff --git a/STILL.md b/STILL.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := [][]byte{}
			scanner := bufio.NewScanner(tt.args.r)
			scanner.Split(ScanDiffs)
			for scanner.Scan() {
				got = append(got, scanner.Bytes())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("ScanDiffs()")
				// t.Errorf("\n\nFull::: \n\nGOT: %s\n\nWANT: %s", got, tt.want)
				t.Log("GOT::::")
				for i, g := range got {
					t.Logf("%d: %s\n", i, g)
				}
				t.Log("WANT::::")
				for i, g := range tt.want {
					t.Logf("%d: %s\n", i, g)
				}

			}
		})
	}
}
