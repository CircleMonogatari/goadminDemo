package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
)



func(a *API)USERS()[]string{
	b, err := a.GET("http://localhost:8000/users", nil)
	if err != nil {
		return nil
	}

	users := User{}

	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil
	}
	if users.Statuc != "ok" {
		return nil
	}
	return users.Data
}

type Transactionlist struct {
	Data []Data `json:"data"`
}

func(a *API)Transactionlist(address string) string{
	b, err := a.GET("http://localhost:8000/transactionlist", nil)
	if err != nil {
		return err.Error()
	}
	//TODO:数据树排列
	return string(b)
}


func(a *API)VERSION()int{
	b, err := a.GET("http://localhost:8000/version", nil)
	if err != nil {
		return 0
	}

	v := Version{}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return 0
	}

	return v.Version
}

func(a *API)Balance(address string)*Balance{

	p := url2.Values{}
	p.Add("address", address)
	b, err := a.GET("http://localhost:8000/balance", p)
	if err != nil {
		return nil
	}

	v := Balance{}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil
	}

	return &v
}

type Balancedetailed struct {
	Data   []Data `json:"data"`
	Status string `json:"status"`
}

func(a *API)Balancedetailed(address string)*Balancedetailed{

	p := url2.Values{}
	p.Add("address", address)
	b, err := a.GET("http://localhost:8000/balancedetailed", p)
	if err != nil {
		return nil
	}

	v := Balancedetailed{}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil
	}
	return &v
}


//录入数据
func(a *API)Entry(address, data string)error{
	body := make(map[string]interface{})
	body["address"] = address
	body["amount "] = 1
	body["data "] = data

	b, err := a.POST("http://localhost:8000/entry", nil, body)
	if err != nil {
		return err
	}

	v := make(map[string]interface{})
	err = json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	if v["status"] != "ok" {
		return fmt.Errorf("Server return Error")
	}
	return nil
}

//转发
func(a *API)Repost(address, to, data string)error{
	body := make(map[string]interface{})
	body["from"] = address
	body["to"] = to
	body["amount"] = 1
	body["data"] = data

	b, err := a.POST("http://localhost:8000/transaction", nil, body)
	if err != nil {
		return err
	}

	v := make(map[string]interface{})
	err = json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	if v["status"] != "ok" {
		return fmt.Errorf("Server return Error")
	}
	return nil
}

func(a *API)GET(url string, Query url2.Values)([]byte, error){
	// Get
	u, _ := url2.Parse(url)
	if Query != nil {
		u.RawQuery = Query.Encode()
	}

	res, err := http.Get(u.String());
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}


func(a *API)POST(url string, Query url2.Values, s map[string]interface{})([]byte, error){
	// Get
	u, _ := url2.Parse(url)
	if Query != nil {
		u.RawQuery = Query.Encode()
	}
	//body
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer([]byte(b))

	res, err := http.Post(u.String(), "application/json;charset=utf-8", body);
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

