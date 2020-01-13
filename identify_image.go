package whatimage

import (
	"io"
)

func IdentifyImage(reader io.ReadSeeker) string {
	r := IdentifySimple(reader)
	if r != "" {
		return r
	}

	reader.Seek(0, io.SeekStart)
	return IdentifyISOBmff(reader)
}
