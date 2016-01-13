package services

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func StartServiceA() {
	viper.BindEnv("version")
	viper.BindEnv("port")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+viper.GetString("port"), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm version %s of Service A\n", viper.GetString("version"))
	url := "http://serviceb:8383/"
	fmt.Fprintf(w, "I'm calling service B at this location: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, "Problem getting %s, error: %s\n", url, err)
	}
	if resp == nil || resp.Body == nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "\n\nResponse:\n %s", bytes.NewBuffer(body))
}
