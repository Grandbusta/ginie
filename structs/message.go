package structs

type Response struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}
type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}
type Profile struct {
	Name string `json:"name"`
}
type Contacts struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}
type Text struct {
	Body string `json:"body"`
}

type Image struct {
	Caption  string `json:"caption"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}
type Messages struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Text      Text   `json:"text"`
	Image     Image  `json:"image"`
}

type Value struct {
	MessagingProduct string     `json:"messaging_product"`
	Metadata         Metadata   `json:"metadata"`
	Contacts         []Contacts `json:"contacts"`
	Messages         []Messages `json:"messages"`
}
type Changes struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}
type Entry struct {
	ID      string    `json:"id"`
	Changes []Changes `json:"changes"`
}
