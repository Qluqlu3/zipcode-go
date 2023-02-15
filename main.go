package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("郵便番号を入力：")
	zipCode, _ := reader.ReadString('\n')
	zipCode = zipCode[:len(zipCode)-1]

	if i, err := strconv.Atoi(zipCode); err != nil {
		fmt.Println("数字で入力してください。")
		return
	} else if len(zipCode) != 7 {
		fmt.Println("7桁で入力してください。")
		return
	}

	apiUrl := "https://zipcloud.ibsnet.co.jp/api/search?zipcode=" + zipCode
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println(err)
		return
	}

	results, ok := response["results"].([]interface{})
	if !ok || len(results) == 0 {
		fmt.Println("その郵便番号に対応した住所は存在しません。")
		return
	}
	result := results[0].(map[string]interface{})
	address1, _ := result["address1"].(string)
	address2, _ := result["address2"].(string)
	address3, _ := result["address3"].(string)
	fmt.Printf("住所: %s %s %s\n", address1, address2, address3)
}
