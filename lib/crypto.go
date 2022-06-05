package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Cryptodata struct {
	Usd float32 `json:"usd"`
}

var dsn string = fmt.Sprintf("https://graph.facebook.com/v13.0/110769228316637/messages?access_token=%v", "EAATQ0QIR0scBAFSrrKIJAfbQcMMmiTlZBZAbUnqcvDrK7Tp8WyPoU7BAxm4JXGEE5srwAiOnpEJY2ZBbMEudQDiFvHUIyGgdBq1upSEcSAs6pmcmNDDmHgMzdIOdetZAN6O6C1LZCFqsc5wB121fYK8ZAOmPZATVxG5Q346zC2FZCZAO7YEdCqctAerFcnZCB1SkQhpZBOuKG1JZCAZDZD")

func GetCryptoPrice(phone_number_id, from, msg_body string) {
	var body string
	msg_split := strings.Split(msg_body, " ")
	req_url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%v&vs_currencies=usd", strings.ToLower(msg_split[2]))
	v, _ := http.Get(req_url)
	bodys, _ := ioutil.ReadAll(v.Body)
	data := make(map[string]Cryptodata, 0)
	json.Unmarshal(bodys, &data)
	body = fmt.Sprintf("Name: %v \nPrice: %v", strings.ToLower(msg_split[2]), data[strings.ToLower(msg_split[2])].Usd)
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
	fmt.Println(resp.Status)
}
