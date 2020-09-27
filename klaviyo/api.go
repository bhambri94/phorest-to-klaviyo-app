package klaviyo

import (
	"fmt"
	"net/http"
)

func TrackEventOnKlaviyo(dataRequest string) {

	resp, err := http.Get("https://a.klaviyo.com/api/track?data=" + dataRequest)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
}
