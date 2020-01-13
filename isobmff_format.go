package whatimage

import (
	"bytes"
	"encoding/binary"
	"io"
)

type identify_isobmff_func func(major_brand []byte, minor_version uint32, compatible_brands []byte) string

func identifyHeic(major_brand []byte, minor_version uint32, compatible_brands []byte) string {
	container_brands := [][]byte{[]byte("mif1"), []byte("msf1")}
	coding_brands := [][]byte{[]byte("heic"), []byte("heix"), []byte("hevc"), []byte("hevx")}

	for i := range coding_brands {
		if bytes.Compare(major_brand, coding_brands[i]) == 0 {
			return "heic"
		}
	}

	for i := range container_brands {
		if bytes.Compare(major_brand, container_brands[i]) == 0 {
			for n := 0; n < len(compatible_brands)-4; n += 4 {
				cb := compatible_brands[n*4 : n*4+4]
				for j := range coding_brands {
					if bytes.Compare(cb, coding_brands[j]) == 0 {
						return "heic"
					}
				}
			}
			break
		}
	}

	return ""
}

func identifyAvif(major_brand []byte, minor_version uint32, compatible_brands []byte) string {
	container_brands := [][]byte{[]byte("mif1"), []byte("msf1")}
	coding_brands := [][]byte{[]byte("avif"), []byte("avis")}

	for i := range coding_brands {
		if bytes.Compare(major_brand, coding_brands[i]) == 0 {
			return "avif"
		}
	}

	for i := range container_brands {
		if bytes.Compare(major_brand, container_brands[i]) == 0 {
			for n := 0; n < len(compatible_brands)-4; n += 4 {
				cb := compatible_brands[n*4 : n*4+4]
				for j := range coding_brands {
					if bytes.Compare(cb, coding_brands[j]) == 0 {
						return "avif"
					}
				}
			}
			break
		}
	}

	return ""
}

func identifyISOBmff(reader io.Reader, funcs ...identify_isobmff_func) string {
	buf := make([]byte, 16)
	n, err := reader.Read(buf)
	if err != nil || n < 16 {
		return ""
	}

	if bytes.Compare(buf[4:8], []byte{'f', 't', 'y', 'p'}) != 0 {
		return ""
	}

	ftyp_len := binary.BigEndian.Uint32(buf[0:4])
	major_brand := buf[8:12]
	minor_version := binary.BigEndian.Uint32(buf[12:16])

	buf = make([]byte, ftyp_len-16)
	n, err = reader.Read(buf)
	if err != nil || uint32(n) < (ftyp_len-16) {
		return ""
	}

	for _, f := range funcs {
		r := f(major_brand, minor_version, buf)
		if r != "" {
			return r
		}
	}

	return ""
}

func IdentifyHeic(reader io.Reader) bool {
	return identifyISOBmff(reader, identifyHeic) == "heic"
}

func IdentifyAvif(reader io.Reader) bool {
	return identifyISOBmff(reader, identifyAvif) == "avif"
}

func IdentifyISOBmff(reader io.Reader) string {
	funcs := []identify_isobmff_func{
		identifyHeic,
		identifyAvif,
	}

	return identifyISOBmff(reader, funcs...)
}
