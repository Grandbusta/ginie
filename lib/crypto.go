package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Cryptodata struct {
	Usd           float32 `json:"usd"`
	Gbp           float32 `json:"gbp"`
	Eur           float32 `json:"eur"`
	LastUpdatedAt int64   `json:"last_updated_at"`
}

var dsn string = fmt.Sprintf("https://graph.facebook.com/v13.0/110769228316637/messages?access_token=%v", os.Getenv("WA_KEY"))

func GetCryptoPrice(phone_number_id, from, profile_name string, msg_split []string) int {
	var body string
	if len(msg_split) < 2 {
		body = fmt.Sprintf("Hello %v..👋🏽 \n\nNo token name in request.", profile_name)
	} else {
		currency := msg_split[1]
		req_url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%v&vs_currencies=usd,gbp,eur&include_last_updated_at=true", currency)
		v, _ := http.Get(req_url)
		fmt.Println("str", v.Status)
		bodys, _ := ioutil.ReadAll(v.Body)
		data := make(map[string]Cryptodata, 0)
		json.Unmarshal(bodys, &data)
		if data[currency].Eur == 0 {
			body = fmt.Sprintf("Hello %v..👋🏽 \n\nThe specified coin/token was NOT FOUND.\nAre you sure it is in full?\ne.g: Bitcoin instead of BTC", profile_name)
		} else {
			body = fmt.Sprintf("Hello %v..👋🏽 \n\n*%v Price today..*💪🏽\nUSD🇺🇸 --> $%v\nGPB🇬🇧 --> £%v\nEURO🇪🇺 --> €%v\n_Last Updated: %v_", profile_name, strings.Title(currency), data[currency].Usd, data[currency].Gbp, data[currency].Eur, time.Unix(data[currency].LastUpdatedAt, 0))
		}
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
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return resp.StatusCode
}
