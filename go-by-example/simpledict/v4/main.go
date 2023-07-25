package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

type DicStruct struct {
	LoggedIn            bool          `json:"loggedIn"`
	Notifications       []interface{} `json:"notifications"`
	UserName            interface{}   `json:"userName"`
	RandomImageQuizHook struct {
		Category   string `json:"category"`
		Level      int    `json:"level"`
		Filename   string `json:"filename"`
		QuizID     string `json:"quizId"`
		Copyright1 string `json:"copyright1"`
		Copyright2 string `json:"copyright2"`
		Copyright3 string `json:"copyright3"`
	} `json:"randomImageQuizHook"`
	RecommendedQuizzes []struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		Level     int    `json:"level"`
		Image     string `json:"image"`
		Copyright string `json:"copyright"`
		Shield1   bool   `json:"shield1"`
		Shield2   bool   `json:"shield2"`
		Shield3   bool   `json:"shield3"`
		Shield4   bool   `json:"shield4"`
	} `json:"recommendedQuizzes"`
	Wordlists             interface{} `json:"wordlists"`
	WordlistHomepage      interface{} `json:"wordlistHomepage"`
	PreferredDictionaries []struct {
		DataCode string `json:"dataCode"`
		Name     string `json:"name"`
		Selected bool   `json:"selected"`
	} `json:"preferredDictionaries"`
	TranslatePanelDefaultEntry struct {
		Dataset     string `json:"dataset"`
		DatasetText string `json:"datasetText"`
		EntryRight  string `json:"entryRight"`
		EntryLeft   string `json:"entryLeft"`
		EntryURL    string `json:"entryUrl"`
	} `json:"translatePanelDefaultEntry"`
	DisplayLoginPopup    bool `json:"displayLoginPopup"`
	DisplayClassicSurvey bool `json:"displayClassicSurvey"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
			example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	go query(word)
	query2(word)
}

func query2(word string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://dictionary.cambridge.org/zhs/auth/info?rid=amp-EjIuKiidMrzRZ4MzEp9L8g&url=https%3A%2F%2Fdictionary.cambridge.org%2Fzhs%2F%25E8%25AF%258D%25E5%2585%25B8%2F%25E8%258B%25B1%25E8%25AF%25AD-%25E6%25B1%2589%25E8%25AF%25AD-%25E7%25AE%2580%25E4%25BD%2593%2Fsleep&ref=https%3A%2F%2Fdictionary.cambridge.org%2Fzhs%2F%25E8%25AF%258D%25E5%2585%25B8%2F%25E8%258B%25B1%25E8%25AF%25AD-%25E6%25B1%2589%25E8%25AF%25AD-%25E7%25AE%2580%25E4%25BD%2593%2Fsalad&type=ENTRY_TRANSLATE&v1=english-chinese-simplified&"+
		"v2="+word+
		"&v3=&v4=english-chinese-simplified&v5=&v6=&_=0.7942914152585214&__amp_source_origin=https%3A%2F%2Fdictionary.cambridge.org", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("AMP-Same-Origin", "true")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", `preferredDictionaries="english-chinese-simplified,english-chinese-traditional,english,british-grammar"; XSRF-TOKEN=81341a01-59df-4e29-acb5-6210298e7d3f; amp-access=amp-EjIuKiidMrzRZ4MzEp9L8g; loginPopup=5`)
	req.Header.Set("Referer", "https://dictionary.cambridge.org/zhs/%E8%AF%8D%E5%85%B8/%E8%8B%B1%E8%AF%AD-%E6%B1%89%E8%AF%AD-%E7%AE%80%E4%BD%93/sleep")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var res DicStruct
	if err := json.Unmarshal(bodyText, &res); err != nil {
		log.Fatal("Error for unmarshalling json result")
		return
	}
	fmt.Println(res.TranslatePanelDefaultEntry.DatasetText, ": ", res.TranslatePanelDefaultEntry.EntryLeft)
}
