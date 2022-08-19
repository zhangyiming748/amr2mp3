package amr

import "testing"

func TestShell(t *testing.T) {
	src := "/Users/zen/Github/Tools/amr/before"
	dst := "/Users/zen/Github/Tools/amr/after"
	ConvertMP3(src, dst)
}
