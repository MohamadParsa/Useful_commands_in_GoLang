package core

import "encoding/json"

func SearchVersion1() []byte {
	type solution struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	res := solution{Name: "solution1", Address: "address1"}
	jsonResult, _ := json.Marshal(res)
	return jsonResult
}
func SearchVersion2() []byte {
	type solution struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Number  string `json:"number"`
	}
	res := solution{Name: "solution2", Address: "address2", Number: "number2"}
	jsonResult, _ := json.Marshal(res)
	return jsonResult
}
