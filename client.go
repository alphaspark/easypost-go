package easypost

type Client struct {
	apiKey  string
	baseURL string
}

func NewClient(key string) *Client {
	return &Client{key, "https://api.easypost.com/v2"}
}

type Address map[string]interface{}
type Parcel map[string]interface{}
type CustomsItem map[string]interface{}
type CustomsInfo map[string]interface{}
type Shipment map[string]interface{}

func (c *Client) NewShipment(fromAddressId string, toAddress *Address, parcel *Parcel) (s *Shipment, e error) {
	return
}

func (c *Client) NewShipmentWithCustoms(fromAddressId string, toAddress *Address, parcel *Parcel, customsInfo *CustomsInfo) (s *Shipment, e error) {
	return
}

func (c *Client) BuyShipment(shipment map[string]interface{}, rateId string) (e error) {
	return
}
