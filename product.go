package apiTokenControl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

func DeleteAllItems(s Token)(bool, error) {
		
	ids, err := Token.GetProductItemsIds(s)
	
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(ids); i++ {	
		fmt.Println("https://api.mercadolibre.com/items/" + ids[i])
		errClosed:=deleteItem(ids[i], `{"status": "closed"}`,s.Access_token)

		if errClosed != nil {
			return false, errClosed 
		}

		errDeleted:=deleteItem(ids[i], `{"deleted": "true"}`,s.Access_token)

		if errDeleted != nil {
			return false, errDeleted 
		}

	}
	return true, nil 
}

func deleteItem(item string, data string, access_token string )(error) {

	url := "https://api.mercadolibre.com/items/" + item
	 
	client := &http.Client{}

	req, _ := http.NewRequest("PUT", url, bytes.NewReader([]byte(data)))

	req.Header.Add("Authorization", "Bearer "+ access_token)

	resp, err := client.Do(req)

	if err != nil {
		return err 
	}

	defer resp.Body.Close()
	return nil 
}