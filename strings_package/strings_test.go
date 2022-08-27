// strings パッケージをご覧ください。(https://pkg.go.dev/strings)
// 便利だと思われる関数を見つけて、ここにあるようなテストを作成して実験してください。
// 標準ライブラリの学習に時間を費やすことは、時間の経過とともに本当に実を結びます。

package strings_package

import (
	"testing"
	"strings"
)

func TestIndex(t *testing.T) {
	t.Run("A matching string exists", func (t *testing.T) {
		got := strings.Index("Nakamura", "kamu")
		want := 2
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
	t.Run("Return -1 if the strings does not exits", func (t *testing.T) {
		got := strings.Index("Nakamura", "x")
		want := -1
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}