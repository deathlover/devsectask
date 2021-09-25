package client

import (
	"math/rand"
	"strings"
	"sync"
)

func (c *Job) AttackTranshID(url string, data []string) {
	var wg sync.WaitGroup
	limiter := make(chan bool, 5)

	for _, v := range data {
		limiter <- true
		wg.Add(1)
		finalURL := BaseURL + strings.ReplaceAll(url+"/{id}", "{id}", v)
		go func() {
			defer func() { <-limiter }()
			defer wg.Done()
			c.SimpleRunner(finalURL)
		}()
	}
	wg.Wait()
}

func (c *Job) AttackAssetID(data string) {
	var wg sync.WaitGroup
	limiter := make(chan bool, 5)

	limiter <- true
	wg.Add(1)
	finalURL := BaseURL + "/assets/" + data
	go func() {
		defer func() { <-limiter }()
		defer wg.Done()
		c.SimpleRunner(finalURL)
	}()
}

func (c *Job) StupidAttacker() {
	var finalURL string

	for _, urltransh := range c.DataTransh {
		if strings.Contains(urltransh, "{id}") {
			continue
		} else {
			urltransh = strings.ReplaceAll(urltransh, "\"", "")
			finalURL = BaseURL + urltransh + "?"

			for _, v := range c.Params {
				finalURL += v + "=FAKE&"
			}
			c.SimpleRunner(finalURL)
		}
	}
}

func (c *Job) MindAttacker() {
	var finalURL string

	for transhurl, params := range c.ParamsMind {
		if strings.Contains(transhurl, "{id}") {
			continue
		} else {
			transhurl = strings.ReplaceAll(transhurl, "\"", "")
			finalURL = BaseURL + transhurl + "?"

			for ix := 0; ix < len(params)-1; ix++ {
				if params[ix] == "" {
					continue
				}
				switch params[ix] {
				case "sender":
					randomIndex := rand.Intn(len(c.SenderTransh[transhurl]))
					pick := c.SenderTransh[transhurl][randomIndex]
					finalURL += "sender=" + pick + "&"
				case "recipient":
					randomIndex := rand.Intn(len(c.RecipientTransh[transhurl]))
					pick := c.RecipientTransh[transhurl][randomIndex]
					finalURL += "recipient=" + pick + "&"
				case "limit":
					finalURL += "limit=100&"
				case "sort":
					finalURL += "sort=desc&"
				default:
					finalURL += params[ix] + "=FAKE&"
				}
			}
			c.SimpleRunner(finalURL)
		}
	}
}
