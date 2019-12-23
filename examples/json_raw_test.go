package examples

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestJsonRaw(t *testing.T) {
	type Message struct {
		ObjectType string          `json:"object_type"`
		SKU        int64           `json:"id,string"`
		Event      string          `json:"event"`
		CreatedAt  time.Time       `json:"created_at"`
		Data       json.RawMessage `json:"data"`
	}
	type MsgData struct {
		IsEnabled bool   `json:"is_enabled"`
		CompanyID int64  `json:"company_id"`
		ProductID int64  `json:"product_id"`
		Source    string `json:"source"`
	}
	value := []byte(`{"data":{"is_enabled":true,"company_id":1,"product_id":16001371,"source":"fbo"},"updated_at":"2019-12-23T11:53:02Z","object_type":"sku","id":"159943916","event":"skuEnabledChange"}`)
	msg := &Message{}
	err := json.Unmarshal(value, msg)
	if err != nil {
		fmt.Printf("error %#v\n", err)
	}
	fmt.Printf("msg %#v\n", msg)
	msgData := &MsgData{}
	err = json.Unmarshal(msg.Data, msgData)
	if err != nil {
		fmt.Printf("error %#v\n", err)
	}
	fmt.Printf("msgData %#v\n", msgData)
}
