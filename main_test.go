package amr2mp3

import (
	"testing"
)

func TestMaster(t *testing.T) {
	src := "amr/before"
	dst := "amr/after"

	ConvertMP3(src, dst)
}
