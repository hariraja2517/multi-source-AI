package main

import (
	"fmt"
	"net/http"
)

type train struct {
	Tag       string `json:"tag"`
	Patterns  string `json:"patterns"`
	Responses string `json:"responses"`
}

type rMessage struct {
	Query string `json:"query"`
	Lan   string `json:"lan"`
}
type sMessage struct {
	Message string `json:"message"`
}
type rsMessage struct {
	Message string `json:"message"`
	Lan     string `json:"lan"`
}

var smessage *sMessage

func test(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	// if r.Method != "OPTIONS" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	fmt.Fprintf(w, "StatusMethodNotAllowed")
	// 	return
	// }
	fmt.Fprintf(w, "test123")
}

func getData(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "StatusMethodNotAllowed")
		return
	}
	writeJson(w, 200, db)
}

func message(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "StatusMethodNotAllowed")
		return
	}
	var rmessage *rMessage
	bindJson(r, &rmessage)

	dbhandler(*rmessage, "message")
	if smessage.Message != "not found" {
		writeJson(w, http.StatusAccepted, smessage)
	} else {
		botTrainer(*rmessage)
		if len(rmessage.Lan) > 0 {
			translateData := &rsMessage{
				Message: smessage.Message,
				Lan:     rmessage.Lan,
			}
			translate(translateData)
		}
		writeJson(w, http.StatusAccepted, smessage)
	}

}

func trainBot(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "StatusMethodNotAllowed")
		return
	}
	var data *train
	bindJson(r, &data)

	dbhandler(*data, "train")

	if len(data.Tag) > 0 && len(data.Patterns) > 0 && len(data.Responses) > 0 {
		writeJson(w, http.StatusCreated, *data)
	} else {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "struct err"})
	}
}
