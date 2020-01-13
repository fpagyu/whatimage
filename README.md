# whatimage
identify image format

```Simple Demo

import (
    "os"
    "log"
    "github.com/fpagyu/whatimage"
)

func main() {
	files := []string{
		"32851.jpg",
		"32851.png",
		"32851.HEIC",
	}

	for i := range files {
		f, err := os.Open(files[i])
		if err != nil {
			log.Println("error info: ", err)
		    continue
		}
		log.Logf("%s's format: %s", files[i], IdentifyImage(f))
	}
}

```
