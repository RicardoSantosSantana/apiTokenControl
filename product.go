package apiTokenControl

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (s Token) GetProductItemsIds() ([]string, error) {

	product := ItemsIds{}

	var url string = "https://api.mercadolibre.com/users/" + strconv.Itoa(s.User_id) + "/items/search?access_token=" + s.Access_token

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if body != nil && strings.Contains(string(body), "results") {
		err := json.Unmarshal(body, &product)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	defer resp.Body.Close()
	return product.Results, err
}

func GetItemsDetails(itemId string) (Items, error) {

	items := Items{}

	url := "https://api.mercadolibre.com/items/" + itemId

	resp, err := http.Get(url)
	if err != nil {
		return Items{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Items{}, err
	}

	if body != nil && strings.Contains(string(body), "site_id") {

		err := json.Unmarshal(body, &items)
		if err != nil {
			return Items{}, err
		}

	} else {
		err := errors.New("no valid results on GetProductItems")
		if err != nil {
			return Items{}, err
		}
	}

	return items, nil
}

func GetItemsDescription(itemId string) (Description, error) {

	description := Description{}

	url := "https://api.mercadolibre.com/items/" + itemId + "/description"

	resp, err := http.Get(url)
	if err != nil {
		return description, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return description, err
	}

	err_json := json.Unmarshal(body, &description)
	if err_json != nil {
		return description, err_json
	}

	defer resp.Body.Close()
	return description, err
}
