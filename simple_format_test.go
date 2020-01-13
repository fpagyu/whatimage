package whatimage

import (
	"os"
	"testing"
)

func TestIdentifyJPEG(t *testing.T) {
	f, err := os.Open("/home/jvm/testfile/result.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	r := IdentifyJpeg(f)
	t.Log("format: ", r)
}
