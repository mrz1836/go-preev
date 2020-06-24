package preev

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetPairs this endpoint retrieves all active pairs
//
// For more information: https://preev.pro/api/
func (c *Client) GetPairs() (pairList *PairList, err error) {

	var resp string
	// https://api.preev.pro/v1/pairs
	if resp, err = c.Request(fmt.Sprintf("%s/pairs", apiEndpoint), http.MethodGet, nil); err != nil {
		return
	}

	// Status was not as expected?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error from Preev: %s", resp)
		return
	}

	err = json.Unmarshal([]byte(resp), &pairList)
	return
}

// GetPair this endpoint retrieves the requested pair
//
// For more information: https://preev.pro/api/
func (c *Client) GetPair(pairID string) (pair *Pair, err error) {

	var resp string
	// https://api.preev.pro/v1/pairs/<pair_id>
	if resp, err = c.Request(fmt.Sprintf("%s/pairs/%s", apiEndpoint, pairID), http.MethodGet, nil); err != nil {
		return
	}

	// Status was not as expected?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error from Preev: %s", resp)
		return
	}

	err = json.Unmarshal([]byte(resp), &pair)
	return
}