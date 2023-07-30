// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"sync"

// 	"github.com/julienschmidt/httprouter"
// )

// type AuthInfo struct {
// 	Key  string `json:"key"`
// 	ID   string `json:"id"`
// 	Info string `json:"info"`
// }

// type AuthReponse struct {
// 	Ok       bool     `json:"ok"`
// 	AuthInfo AuthInfo `json:"authinfo"`
// }

// var (
// 	keyStore     map[string]AuthInfo
// 	keyStoreLock sync.RWMutex
// )

// func AddKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	var authInfo AuthInfo
// 	err := json.NewDecoder(r.Body).Decode(&authInfo)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}
// 	keyStoreLock.Lock()
// 	keyStore[authInfo.Key] = authInfo
// 	keyStoreLock.Unlock()
// }

// func ValidateKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	key := ps.ByName("key")

// 	keyStoreLock.RLock()
// 	authInfo, exists := keyStore[key]
// 	keyStoreLock.RUnlock()

// 	authResponse := AuthReponse{
// 		Ok:       exists,
// 		AuthInfo: authInfo,
// 	}

// 	jsonString, err := json.Marshal(authResponse)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, string(jsonString))
// }

// func serveIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Println("serving shit")
// 	// http.ServeFile(w, r, "./static/index.html")
// 	http.ServeFile(w, r, "./static/index.html")
// }

// // func main() {
// // 	keyStore = make(map[string]AuthInfo)
// // 	router := httprouter.New()
// // 	router.POST("/addKey", AddKey)
// // 	router.GET("/validateKey/:key", ValidateKey)
// // 	router.GET("/", serveIndex)

// // 	fmt.Println("App running on port 8080")
// // 	log.Fatal(http.ListenAndServe(":8080", router))

// // }

// // func main() {
// // 	_, b, _, _ := runtime.Caller(0)
// // 	basepath := filepath.Dir(b)
// // 	staticDir := filepath.Join(basepath, "static")

// // 	router := httprouter.New()
// // 	router.ServeFiles("/static/*filepath", http.Dir(staticDir))

// // 	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// // 		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
// // 	})

// // 	http.ListenAndServe(":8000", router)
// // }

package main

import (
	"encoding/hex"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var (
	apiKeys = []map[string]string{}
	mutex   sync.Mutex
)

// func createAPIKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	description := r.FormValue("description")
// 	// Generate a random API key
// 	keyValue := generateAPIKey()

// 	key := map[string]string{"key": keyValue, "description": description}

// 	mutex.Lock()
// 	apiKeys = append(apiKeys, key)
// 	mutex.Unlock()

// 	tmpl := `<table>{{ range . }}<tr><td>{{ .key }}</td><td>{{ .description }}</td></tr>{{ end }}</table>`
// 	t, _ := template.New("keys").Parse(tmpl)
// 	t.Execute(w, apiKeys)
// }

func createAPIKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	description := r.FormValue("description")
	keyValue := generateAPIKey()

	key := map[string]string{"key": keyValue, "description": description}

	mutex.Lock()
	apiKeys = append(apiKeys, key)
	mutex.Unlock()

	tmpl := `<table>
                <tr>
                    <th>Service/User</th>
                    <th>Key</th>
                    <th>Action</th>
                </tr>
                {{ range $index, $element := . }}
                <tr>
                    <td>{{ .description }}</td>
                    <td>{{ .key }}</td>
                    <td><button onclick="deleteKey({{$index}})">Delete</button></td>
                </tr>
                {{ end }}
            </table>`
	t, _ := template.New("keys").Parse(tmpl)
	t.Execute(w, apiKeys)
}

func deleteAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, err := strconv.Atoi(ps.ByName("index"))
	if err != nil || index < 0 || index >= len(apiKeys) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	apiKeys = append(apiKeys[:index], apiKeys[index+1:]...)
	mutex.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// generateAPIKey creates a random string to be used as an API key
func generateAPIKey() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func main() {
	router := httprouter.New()

	// Serve index.html at the root
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "./static/index.html")
	})

	router.POST("/create", createAPIKey)
	router.GET("/delete/:index", deleteAPIKey)

	fmt.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", router)
}

// func main() {
// 	_, b, _, _ := runtime.Caller(0)
// 	basepath := filepath.Dir(b)
// 	staticDir := filepath.Join(basepath, "static")

// 	fmt.Println(staticDir)

// 	router := httprouter.New()
// 	router.ServeFiles("/static/*filepath", http.Dir(staticDir))

// 	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
// 	})

// 	http.ListenAndServe(":8000", router)
// }

// func main() {
// 	router := httprouter.New()

// 	// Serve index.html at the root
// 	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 		http.ServeFile(w, r, "./static/index.html")
// 	})

// 	// // Serve static files
// 	// staticDir := filepath.Join(".", "static")
// 	// router.ServeFiles("/static/*filepath", http.Dir(staticDir))

// 	http.ListenAndServe(":8000", router)
// }
