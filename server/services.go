package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

//	type rbt struct {
//		Text string `json:"text"`
//	}
type jb3 struct {
	Content string `json:"content"`
}
type jb2 struct {
	Message jb3 `json:"message"`
}
type jb struct {
	Choices []jb2 `json:"choices"`
}
type sbt struct {
	Query string
	Mesg  string
}

type rt struct {
	Trans                string  `json:"trans"`
	Source_language_code string  `json:"source_language_code"`
	Source_language      string  `json:"source_language"`
	Trust_level          float64 `json:"Trust_level"`
}

func translate(data *rsMessage) {

	url := "https://google-translate113.p.rapidapi.com/api/v1/translator/text"

	payload := strings.NewReader("from=auto&to=" + data.Lan + "&text=" + data.Message)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", "bdc65970b5mshb17e38c69631846p134748jsnf942ce2e206a")
	req.Header.Add("X-RapidAPI-Host", "google-translate113.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var rt *rt
	json.Unmarshal(body, &rt)
	smessage = &sMessage{
		Message: rt.Trans,
	}

}

func botTrainer(data rMessage) {

	// url := "https://ai-api-textgen.p.rapidapi.com/completions"
	url := "https://open-ai34.p.rapidapi.com/api/v1/chat/completions"
	// d2 := rbt2{
	// 	Role:    "user",
	// 	Content: data.Query,
	// }

	d := map[string]any{
		"model": "mixtral-8x7b",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": data.Query,
			},
		},
	}
	jsonValue, _ := json.Marshal(d)

	body := bytes.NewReader(jsonValue)

	req, _ := http.NewRequest("POST", url, body)

	// req.Header.Add("content-type", "application/json")
	// req.Header.Add("X-RapidAPI-Key", "bdc65970b5mshb17e38c69631846p134748jsnf942ce2e206a")
	// req.Header.Add("X-RapidAPI-Host", "ai-api-textgen.p.rapidapi.com")

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "bdc65970b5mshb17e38c69631846p134748jsnf942ce2e206a")
	req.Header.Add("X-RapidAPI-Host", "open-ai34.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	b, _ := io.ReadAll(res.Body)
	var jb *jb
	json.Unmarshal(b, &jb)

	// r := strings.Trim(string(b), "\" ")
	smessage = &sMessage{
		Message: jb.Choices[0].Message.Content,
	}
	defer res.Body.Close()

	dbdata := sbt{
		Query: data.Query,
		Mesg:  jb.Choices[0].Message.Content,
	}
	dbhandler(dbdata, "botTrain")

}
