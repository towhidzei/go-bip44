package bitcoin_address

type HDStartPath struct {
	PurposeIndex  uint32 `json:"purpose_index"`
	CoinTypeIndex uint32 `json:"coin_type"`
	AccountIndex  uint32 `json:"account_index"`
}

type HDEndPath struct {
	ChangeIndex  uint32 `json:"change_index"`
	AddressIndex uint32 `json:"address_index"`
}

type Address struct {
	HDStartPath HDStartPath `json:"hd_path"`
	HDEndPath   HDEndPath   `json:"hd_path"`
	Value       string      `json:"value"`
}
