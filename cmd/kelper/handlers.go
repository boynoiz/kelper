package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Welcome!\n")
}

func PingHandler(res http.ResponseWriter, req *http.Request) {
	ping := Ping{
		"Pong!",
		200}
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-type", "application/json; charset=UTF-8")
	json.NewEncoder(res).Encode(ping)
}

func ShoutHandler(res http.ResponseWriter, req *http.Request) {
	var shout Shout
	if req.Body == nil {
		http.Error(res, "Please give me a shout!\n", 400)
		return
	}
	decoder := json.NewDecoder(req.Body).Decode(&shout)
	if decoder != nil {
		http.Error(res, decoder.Error(), 400)
		return
	}
	hello, _ := regexp.MatchString("([H|h][E|e][L|l]{2}[O|o]).*", shout.Shout)

	if hello {
		echo := Echo{
			Echo: "Oh hey~",
		}
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-type", "application/json; charset=UTF-8")
		json.NewEncoder(res).Encode(echo)
	} else {
		echo := Echo{
			Echo: "How are uuuuu~",
		}
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-type", "application/json; charset=UTF-8")
		json.NewEncoder(res).Encode(echo)
	}
}
