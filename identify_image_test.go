package whatimage

import "testing"

import "os"

func TestIdentifyImage(t *testing.T) {
	files := []string{
		"/home/jvm/testfile/32851.jpg",
		"/home/jvm/testfile/32851.png",
		"/home/jvm/IMG_1383.HEIC",
	}

	for i := range files {
		f, err := os.Open(files[i])
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("format: ", IdentifyImage(f))
	}
}
