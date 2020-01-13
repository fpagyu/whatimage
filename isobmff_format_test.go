package whatimage

import (
	"os"
	"testing"
)

func TestIdentifyISOBmff(t *testing.T) {
	reader, err := os.Open("/home/jvm/IMG_1383.HEIC")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("========= format ==========: ", IdentifyISOBmff(reader))
}
