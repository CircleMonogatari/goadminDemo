package api

type API struct {
}

type User struct {
	Data   []string `json:"data"`
	Statuc string   `json:"statuc"`
}

type Version struct {
	Version int `json:"version"`
}


type Balance struct {
	Data []struct {
		Value        int    `json:"Value"`
		ScriptPubKey string `json:"ScriptPubKey"`
	} `json:"data"`
	Status string `json:"status"`
}

type Vin struct {
	Txid      interface{} `json:"Txid"`
	Vout      int         `json:"Vout"`
	ScriptSig string      `json:"ScriptSig"`
}
type Vout struct {
	Value        int    `json:"Value"`
	ScriptPubKey string `json:"ScriptPubKey"`
}
type Data struct {
	ID   interface{} `json:"ID"`
	Vin  []Vin       `json:"Vin"`
	Vout []Vout      `json:"Vout"`
	Data string      `json:"Data"`
}
