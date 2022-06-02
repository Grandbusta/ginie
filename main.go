package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"ginie/structs"

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
			dsn := fmt.Sprintf("https://graph.facebook.com/v13.0/110769228316637/messages?access_token=%v", "EAATQ0QIR0scBAFSrrKIJAfbQcMMmiTlZBZAbUnqcvDrK7Tp8WyPoU7BAxm4JXGEE5srwAiOnpEJY2ZBbMEudQDiFvHUIyGgdBq1upSEcSAs6pmcmNDDmHgMzdIOdetZAN6O6C1LZCFqsc5wB121fYK8ZAOmPZATVxG5Q346zC2FZCZAO7YEdCqctAerFcnZCB1SkQhpZBOuKG1JZCAZDZD")
			fmt.Println(phone_number_id, from, msg_body, dsn)
			var body string
			// @Todo: logoic to work on
			if strings.ToLower(msg_body) == "crypto" {
				body = "Yes, i will supply you btc"
			} else {
				body = "i don't know you before"
			}
			json_data, err := json.Marshal(map[string]interface{}{
				"messaging_product": "whatsapp",
				"to":                from,
				"recipient_type":    "individual",
				"type":              "text",
				"text": map[string]interface{}{
					"preview_url": true,
					"body":        body,
				},
			})
			if err != nil {
				log.Fatal(err)
			}
			resp, err := http.Post(dsn, "application/json", bytes.NewBuffer(json_data))
			fmt.Println(resp)
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
