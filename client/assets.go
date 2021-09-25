package client

type AssetsResponse []struct {
	Data struct {
		Description     string `json:"description"`
		HasScript       bool   `json:"hasScript"`
		Height          int64  `json:"height"`
		ID              string `json:"id"`
		MinSponsoredFee int64  `json:"minSponsoredFee"`
		Name            string `json:"name"`
		Precision       int64  `json:"precision"`
		Quantity        int64  `json:"quantity"`
		Reissuable      bool   `json:"reissuable"`
		Sender          string `json:"sender"`
		Ticker          string `json:"ticker"`
		Timestamp       string `json:"timestamp"`
	} `json:"data"`
}

type ParamResponse []struct {
	Description string `json:"description"`
	In          string `json:"in"`
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Schema      struct {
		Type string `json:"type"`
	} `json:"schema"`
}

type TranshInfo struct {
	Data []struct {
		Data struct {
			Amount            float64 `json:"amount"`
			ApplicationStatus string  `json:"applicationStatus"`
			Fee               float64 `json:"fee"`
			Height            int64   `json:"height"`
			ID                string  `json:"id"`
			Recipient         string  `json:"recipient"`
			Sender            string  `json:"sender"`
			SenderPublicKey   string  `json:"senderPublicKey"`
			Signature         string  `json:"signature"`
			Timestamp         string  `json:"timestamp"`
			Type              int64   `json:"type"`
		} `json:"data"`
	} `json:"data"`
}
