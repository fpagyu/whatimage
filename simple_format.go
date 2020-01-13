package whatimage

import (
	"bytes"
	"encoding/binary"
	"io"
)

func IdentifyJpeg(reader io.Reader) bool {
	buf := make([]byte, 3)
	n, err := reader.Read(buf)
	if err != nil || n < 3 {
		return false
	}

	return buf[0] == 0xff && buf[1] == 0xd8 && buf[2] == 0xff
}

func IdentifyPng(reader io.Reader) bool {
	buf := make([]byte, 4)
	n, err := reader.Read(buf)
	if err != nil || n < 4 {
		return false
	}

	return buf[0] == 0x89 && bytes.Compare(buf[1:4], []byte{'P', 'N', 'G'}) == 0
}

func IdentifyGif(reader io.Reader) bool {
	buf := make([]byte, 6)
	n, err := reader.Read(buf)
	if err != nil || n < 6 {
		return false
	}

	return bytes.Compare(buf, []byte{'G', 'I', 'F', '8', '7', 'a'}) == 0 ||
		bytes.Compare(buf, []byte{'G', 'I', 'F', '8', '9', 'a'}) == 0
}

func IdentifyWebp(reader io.Reader) bool {
	buf := make([]byte, 12)
	n, err := reader.Read(buf)
	if err != nil || n < 12 {
		return false
	}

	return bytes.Compare(buf[0:4], []byte{'R', 'I', 'F', 'F'}) == 0 &&
		bytes.Compare(buf[8:12], []byte{'W', 'E', 'B', 'P'}) == 0
}

func IdentifyBmp(reader io.Reader) bool {
	buf := make([]byte, 2)
	n, err := reader.Read(buf)
	if err != nil || n < 2 {
		return false
	}

	return bytes.Compare(buf, []byte{'B', 'M'}) == 0
}

func IdentifyPbm(reader io.Reader) bool {
	buf := make([]byte, 2)
	n, err := reader.Read(buf)
	if err != nil || n < 2 {
		return false
	}

	return bytes.Compare(buf, []byte{'P', '4'}) == 0
}

func IdentifyPgm(reader io.Reader) bool {
	buf := make([]byte, 2)
	n, err := reader.Read(buf)
	if err != nil || n < 2 {
		return false
	}

	return bytes.Compare(buf, []byte{'P', '5'}) == 0
}

func IdentifyPpm(reader io.Reader) bool {
	buf := make([]byte, 2)
	n, err := reader.Read(buf)
	if err != nil || n < 2 {
		return false
	}

	return bytes.Compare(buf, []byte{'P', '6'}) == 0
}

func IdentifyTiff(reader io.Reader) bool {
	buf := make([]byte, 4)
	n, err := reader.Read(buf)
	if err != nil || n < 4 {
		return false
	}

	if bytes.Compare(buf[0:2], []byte{'M', 'M'}) == 0 {
		// 判断buf[2:4]的大端字节序 == 42
		return binary.BigEndian.Uint16(buf[2:4]) == 42
	}

	if bytes.Compare(buf[0:2], []byte{'I', 'I'}) == 0 {
		// 判断buf[2:4]的小端字节序 == 42
		return binary.LittleEndian.Uint16(buf[2:4]) == 42
	}

	return false
}

func identifyJpeg(reader io.Reader) string {
	if IdentifyJpeg(reader) {
		return "jpeg"
	}
	return ""
}

func identifyPng(reader io.Reader) string {
	if IdentifyPng(reader) {
		return "png"
	}
	return ""
}

func identifyGif(reader io.Reader) string {
	if IdentifyGif(reader) {
		return "gif"
	}
	return ""
}

func identifyWebp(reader io.Reader) string {
	if IdentifyWebp(reader) {
		return "webp"
	}
	return ""
}

func identifyBmp(reader io.Reader) string {
	if IdentifyBmp(reader) {
		return "bmp"
	}
	return ""
}

func identifyPbm(reader io.Reader) string {
	if IdentifyPbm(reader) {
		return "pbm"
	}
	return ""
}

func identifyPgm(reader io.Reader) string {
	if IdentifyPgm(reader) {
		return "pgm"
	}
	return ""
}

func identifyPpm(reader io.Reader) string {
	if IdentifyPpm(reader) {
		return "ppm"
	}
	return ""
}

func identifyTiff(reader io.Reader) string {
	if IdentifyTiff(reader) {
		return "tiff"
	}
	return ""
}

func IdentifySimple(reader io.ReadSeeker) string {
	funcs := []func(io.Reader) string{
		identifyJpeg,
		identifyPng,
		identifyGif,
		identifyWebp,
		identifyBmp,
		identifyPbm,
		identifyPgm,
		identifyPpm,
		identifyTiff,
	}

	for _, f := range funcs {
		reader.Seek(0, io.SeekStart)
		if r := f(reader); r != "" {
			return r
		}
	}

	return ""
}
