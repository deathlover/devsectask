package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/buger/jsonparser"
)

const (
	BaseURL       = "https://api.wavesplatform.com/v0"
	BaseURLAssets = "https://api.wavesplatform.com/v0/assets?ticker=%2A&search=bitc&limit=100"
	BaseListURL   = "https://api.wavesplatform.com/v0/docs/openapi.json"
)

func (c *Job) SimpleRunner(uri string) []byte {
	// fmt.Println(uri)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyreader, _ := io.ReadAll(resp.Body)

	return bodyreader
}

func (c *Job) GetAssetsID() {
	var asset AssetsResponse

	bodyreader := c.SimpleRunner(BaseURLAssets)
	parser, _, _, _ := jsonparser.Get(bodyreader, "data")
	json.Unmarshal(parser, &asset)

	for i := 0; i < len(asset); i++ {
		c.AssetID = append(c.AssetID, asset[i].Data.ID)
	}
}

func (c *Job) ExecAssetsURL() {
	for _, urlTransh := range c.DataTransh {
		if strings.Contains(urlTransh, "{id}") {
			for _, v := range c.AssetID {
				c.AttackAssetID(v)
			}
		} else {
			continue
		}
	}
}

func (c *Job) ExecTranshURL() {
	for _, urlTransh := range c.DataTransh {
		if strings.Contains(urlTransh, "{id}") {
			transhurl := strings.ReplaceAll(urlTransh, "\"", "")
			for i, v := range c.TranshID {
				if i+"/{id}" == transhurl {
					c.AttackTranshID(i, v)
				}
			}
		} else {
			continue
		}
	}
}

func (c *Job) GetTranshLinks() {
	fmt.Printf("Parsing Information...\n")
	var transactioninfo TranshInfo
	for _, transhurl := range c.DataTransh {
		if strings.Contains(transhurl, "{id}") {
			continue
		} else {
			transhurl = strings.ReplaceAll(transhurl, "\"", "")
			fmt.Println(BaseURL + transhurl)

			resp := c.SimpleRunner(BaseURL + transhurl)
			json.Unmarshal(resp, &transactioninfo)
			for i := 0; i < len(transactioninfo.Data); i++ {
				c.TranshID[transhurl] = append(c.TranshID[transhurl], transactioninfo.Data[i].Data.ID)
				c.RecipientTransh[transhurl] = append(c.RecipientTransh[transhurl], transactioninfo.Data[i].Data.Recipient)
				c.SenderTransh[transhurl] = append(c.SenderTransh[transhurl], transactioninfo.Data[i].Data.Sender)
			}
		}
	}
}

func (c *Job) GetList() {
	var transh ParamResponse
	re := regexp.MustCompile(`"\/transaction.*"`)
	bodyreader := c.SimpleRunner(BaseListURL)
	transactionlinks := re.FindAllString(string(bodyreader), -1)
	c.DataTransh = transactionlinks

	for _, v := range transactionlinks {
		url := strings.ReplaceAll(v, "\"", "")
		paramfindjson, _, _, _ := jsonparser.Get(bodyreader, "paths", url, "get", "parameters")
		json.Unmarshal(paramfindjson, &transh)
		for i := 0; i < len(transh); i++ {
			c.Params = append(c.Params, transh[i].Name)

			c.ParamsMind[v] = append(c.ParamsMind[v], transh[i].Name)
		}
	}
}
