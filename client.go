package easypost

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	apiKey  string
	baseURL string
}

func NewClient(key string) *Client {
	if key == "" {
		return nil
	}
	return &Client{key, "https://api.easypost.com/v2"}
}

type Address map[string]interface{}
type Parcel map[string]interface{}
type CustomsItem map[string]interface{}
type CustomsInfo map[string]interface{}
type Shipment map[string]interface{}

func (c *Client) NewShipment(fromAddressId string, toAddress Address, parcel Parcel) (s Shipment, e error) {
	shipmentBody := map[string]interface{}{
		"shipment": Shipment{
			"from_address": Address{"id": fromAddressId},
			"to_address":   toAddress,
			"parcel":       parcel,
		},
	}
	// Encode JSON body
	jsonStr, err := json.Marshal(shipmentBody)
	if err != nil {
		return nil, err
	}
	// POST to /shipments
	req, rErr := http.NewRequest("POST", fmt.Sprintf("%s%s", c.baseURL, "/shipments"), bytes.NewBuffer(jsonStr))
	if rErr != nil {
		return nil, rErr
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.apiKey, "")

	hc := &http.Client{}
	resp, rsErr := hc.Do(req)
	if rsErr != nil {
		return nil, rsErr
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusCreated {
		err := json.Unmarshal(body, &s)
		if err != nil {
			return nil, err
		}
		return
	}

	return nil, errors.New(fmt.Sprintf("Problem creating shipment: %s", body))
}

func (c *Client) NewShipmentWithCustoms(fromAddressId string, toAddress Address, parcel Parcel, customsInfo CustomsInfo) (s Shipment, e error) {
	shipmentBody := map[string]interface{}{
		"shipment": Shipment{
			"from_address": Address{"id": fromAddressId},
			"to_address":   toAddress,
			"parcel":       parcel,
			"customs_info": customsInfo,
		},
	}
	// Encode JSON body
	jsonStr, err := json.Marshal(shipmentBody)
	if err != nil {
		return nil, err
	}

	// POST to /shipments
	req, rErr := http.NewRequest("POST", fmt.Sprintf("%s%s", c.baseURL, "/shipments"), bytes.NewBuffer(jsonStr))
	if rErr != nil {
		return nil, rErr
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.apiKey, "")

	hc := &http.Client{}
	resp, rsErr := hc.Do(req)
	if rsErr != nil {
		return nil, rsErr
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusCreated {
		err := json.Unmarshal(body, &s)
		if err != nil {
			return nil, err
		}
		return
	}

	return nil, errors.New(fmt.Sprintf("Problem creating shipment: %s", body))
}

func (c *Client) BuyShipment(shipment map[string]interface{}, rateId string) (e error) {
	return
}
