# whatimage
identify image format

It's golang implemention of [whatimage]('https://github.com/david-poirier-csn/whatimage')

```Simple Demo

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
		f.Close()
	}
}

```
