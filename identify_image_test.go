package whatimage

import (
	"os"
	"testing"
)

func TestIdentifyImage(t *testing.T) {
	files := []string{
		//"/home/jvm/testfile/32851.jpg",
		//"/home/jvm/testfile/32851.png",
		//"/home/jvm/IMG_1383.HEIC",
		"/Users/fpgayu/Downloads/pexels-photo-2082087.jpg",
	}

	for i := range files {
		f, err := os.Open(files[i])
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("format: ", IdentifyImage(f))
		f.Close()
	}
}
