package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
	err := decodeJSONBody(res, req, &shout)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(res, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
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
		return
	} else {
		echo := Echo{
			Echo: "How are uuuuu~",
		}
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-type", "application/json; charset=UTF-8")
		json.NewEncoder(res).Encode(echo)
		return
	}
}
