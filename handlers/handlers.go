package handlers

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// IndexHandler - Handle the call for /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Dash Button service is running.")
}

// DashHandler - Handle the call for /dash
func DashHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	mac := strings.ToLower(r.Form.Get("mac"))
	mode := strings.ToLower(r.Form.Get("mode"))
	ip := r.Form.Get("ip")
	host := strings.ToLower(r.Form.Get("host"))

	env := os.Environ()
	env = append(env, fmt.Sprintf("MAC=%s", mac))
	env = append(env, fmt.Sprintf("MODE=%s", mode))
	env = append(env, fmt.Sprintf("IP=%s", ip))
	env = append(env, fmt.Sprintf("HOST=%s", host))

	stringSliceContains := func(list []string, item string) bool {
		for _, listItem := range list {
			if listItem == item {
				return true
			}
		}

		return false
	}

	for _, nestedKey := range viper.AllKeys() {
		nestedKeyParts := strings.Split(nestedKey, ".")

		if len(nestedKeyParts) < 2 {
			continue
		}

		macHostKey := nestedKeyParts[0]
		macHostKeyParts := strings.Split(macHostKey, "+")

		if !stringSliceContains(macHostKeyParts, mac) &&
			!stringSliceContains(macHostKeyParts, host) {
			continue
		}

		modesKey := nestedKeyParts[1]
		modesKeyParts := strings.Split(modesKey, "+")

		if !stringSliceContains(modesKeyParts, mode) {
			continue
		}

		cmdParts := viper.GetStringSlice(nestedKey)

		runnable := exec.Command(cmdParts[0], cmdParts[1:]...)
		runnable.Env = env

		fmt.Fprint(w, strings.Join(cmdParts, " "))
		go func() {
			if err := runnable.Run(); err != nil {
				log.Println(err)
			}
		}()
	}
}
