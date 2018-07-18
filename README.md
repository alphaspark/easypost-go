# EasyPost API v2 client for Go

Currently, this is a minimal implementation of the EasyPost API. It will be expanded as time permits. It's in use in production, but the API is not stable currently.

## Getting Started

More defined models should be added later. Maps are used for now for simplicity.

```golang
c = easypost.NewClient(secretKey)

fromAddressId := "adr_..."
toAddress := map[string]interface{}{
	// Match documented API fields
}
parcel := map[string]interface{}{
	// Match documented API fields
}
shipment, err := c.NewShipment(fromAddressId, toAddress, parcel)

customsInfo := map[string]interface{}{
	// Match documented API fields
}
internationalShipment, isErr := c.NewShipmentWithCustoms(fromAddressId, toAddress, parcel, customsInfo)

selectedRate := "rate_..."
updatedShipment, rErr :=c.BuyShipmentRate(shipment, selectedRate)
```