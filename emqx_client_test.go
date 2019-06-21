package emqx

import (
	"testing"
)

func TestCreateSubscription(t *testing.T) {
	c := ClientConfig{
		BaseURL:   "http://localhost:8080",
		AppID:     "a9ba10e1b3774",
		AppSecret: "Mjg3NzE0ODEwMTAxMTI4Mzk4NTEzOTI5MDUxMzQ1Mzg3NTC",
	}
	a := NewAPIClient(c)
	req := CreateSubscriptionRequestV3{
		Topic:    "abc",
		Qos:      0,
		ClientID: "asss",
	}
	resp, err := a.CreateSubscription(&req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Code != 0 {
		t.Fatal(resp.Code)
	}
}
