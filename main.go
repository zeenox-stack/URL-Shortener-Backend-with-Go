package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/rs/cors"
)

var (
	m   sync.Mutex
	urls = make(map[string]string)
)

type BodyType struct {
	Url string `json:"url"`
}

func generateKey(url string) string {
	hash := sha256.Sum256([]byte(url))
	return hex.EncodeToString(hash[:])[:6]
}

func reset() {
	m.Lock()
	defer m.Unlock()

	urls = make(map[string]string)
}

func redirect(res http.ResponseWriter, req *http.Request) {
	key := req.URL.Path[1:]
	m.Lock()
	defer m.Unlock()

	url, ok := urls[key]
	if !ok {
		http.NotFound(res, req)
		return
	}

	http.Redirect(res, req, url, http.StatusFound)
}

func shorten(res http.ResponseWriter, req *http.Request) {
	var body BodyType

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	m.Lock()
	defer m.Unlock()

	if len(urls) >= 1000 {
		reset()
	}

	if _, ok := urls[body.Url]; ok {
		http.Error(res, "URL already exists", http.StatusBadRequest)
		return
	}

	key := generateKey(body.Url)

	urls[key] = body.Url
	response := map[string]string{"shortened_url": "https://shortenit.up.railway.app/" + key}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", shorten)
	mux.HandleFunc("/", redirect)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, 
		AllowedMethods: []string{"GET", "POST"}, 
		AllowedHeaders: []string{"Content-Type"}, 
		AllowCredentials: true,
	}); 

	handle := c.Handler(mux);

	fmt.Printf("Server is running at http://localhost:8000\n")
	if err := http.ListenAndServe(":8000", handle); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
