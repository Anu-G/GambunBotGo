package random_pics

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetCats() (image, preview, errs string) {
	var client http.Client
	resp, err := client.Get("http://aws.random.cat/meow")
	if err != nil {
		log.Fatal(err)
	}
	// defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		bodyString := string(bodyBytes)
		temp1 := strings.Split(bodyString, `"`)
		temp2 := strings.Split(temp1[3], "\\")
		for _, data := range temp2 {
			image += data
		}

		return image, image, ""
	} else {
		out := fmt.Sprint(resp.StatusCode)
		return "", "", "cat server error " + out
	}
}
