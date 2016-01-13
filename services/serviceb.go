package services

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func StartServiceB() {
	viper.BindEnv("version")
	viper.BindEnv("port")
	http.HandleFunc("/", serviceBHandler)
	http.ListenAndServe(":"+viper.GetString("port"), nil)
}

func serviceBHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm version %s of Service B\n", viper.GetString("version"))
}
