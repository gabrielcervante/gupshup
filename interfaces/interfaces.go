package interfaces

import "encoding/json"

type Whatsapp interface {
	SendMessage(string, string) (json.RawMessage, error)
}
