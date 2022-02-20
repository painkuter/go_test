package examples

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type RfbsReturnMessage struct {
	Weight       float64   `json:"weightG"`
	Amount       string    `json:"amountForDelivery"`
	EventType    string    `json:"eventType"`
	CreatedDate  time.Time `json:"createdAt"`
	ReturnNumber string    `json:"returnNumber"`
}

func TestJSON(t *testing.T) {
	testData := `{ 
    "eventType": "OnWayToSeller3PL", 
    "createdAt": "2022-02-15T011:33:25", 
    "returnNumber": "39145354-R3",
    "trackNumber" : "CL089475504RU", 
    "amountForDelivery" : "12345"
  }`
	result := RfbsReturnMessage{}
	err := json.Unmarshal([]byte(testData), &result)
	assert.Nil(t, err)
	fmt.Println(result)
}
