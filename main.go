package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"ginie/structs"

	"ginie/lib"

	"github.com/gorilla/mux"
)

func webhookPOST(w http.ResponseWriter, r *http.Request) {
	var response structs.Response
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &response)
	if len(response.Object) > 0 {
		if len(response.Entry) > 0 && len(response.Entry[0].Changes) > 0 && len(response.Entry[0].Changes[0].Value.Messages) > 0 {
			phone_number_id := response.Entry[0].Changes[0].Value.Metadata.PhoneNumberID
			from := response.Entry[0].Changes[0].Value.Messages[0].From
			msg_body := response.Entry[0].Changes[0].Value.Messages[0].Text.Body
			fmt.Println(phone_number_id, from, msg_body)
			msg_split := strings.Split(msg_body, " ")
			if strings.ToLower(msg_split[0]) == "crypto" {
				lib.GetCryptoPrice(phone_number_id, from, msg_body)
			}
		}
	}
}

func webhookGET(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("b"))
	query := r.URL.Query()
	mode := query.Get("hub.mode")
	token := query.Get("hub.verify_token")
	challenge := query.Get("hub.challenge")
	fmt.Println(mode, token, challenge)
	if len(token) > 0 && len(mode) > 0 {
		if mode == "subscribe" && token == "ginie" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(challenge))
		}

	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook", webhookPOST).Methods("POST")
	r.HandleFunc("/webhook", webhookGET).Methods("GET")
	r.Handle("/", r)
	http.ListenAndServe(":80", r)
}
