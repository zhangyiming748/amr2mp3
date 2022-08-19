package amr2mp3

import "testing"

func TestShell(t *testing.T) {
	src := "before"
	dst := "after"
	ConvertMP3(src, dst)
}
