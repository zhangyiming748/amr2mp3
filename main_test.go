package amr2mp3

import (
	"github.com/zhangyiming748/amr2mp3/amr"
	"testing"
)

func TestMaster(t *testing.T) {
	src := "amr/before"
	dst := "amr/after"

	amr.ConvertMP3(src, dst)
}
