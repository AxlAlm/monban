package webapp

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// func createAPIKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	description := r.FormValue("description")
// 	keyValue := generateAPIKey()

// 	key := map[string]string{"key": keyValue, "description": description}

// 	tmpl := `<table>
//                 <tr>
//                     <th>Service/User</th>
//                     <th>Key</th>
//                     <th>Action</th>
//                 </tr>
//                 {{ range $index, $element := . }}
//                 <tr>
//                     <td>{{ .description }}</td>
//                     <td>{{ .key }}</td>
//                     <td><button onclick="deleteKey({{$index}})">Delete</button></td>
//                 </tr>
//                 {{ end }}
//             </table>`
// 	t, _ := template.New("keys").Parse(tmpl)
// 	t.Execute(w, apiKeys)
// }

// func deleteAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	index, err := strconv.Atoi(ps.ByName("index"))
// 	if err != nil || index < 0 || index >= len(apiKeys) {
// 		http.Error(w, "Invalid index", http.StatusBadRequest)
// 		return
// 	}

// 	mutex.Lock()
// 	apiKeys = append(apiKeys[:index], apiKeys[index+1:]...)
// 	mutex.Unlock()

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

// // generateAPIKey creates a random string to be used as an API key
// func generateAPIKey() string {
// 	bytes := make([]byte, 32)
// 	rand.Read(bytes)
// 	return hex.EncodeToString(bytes)
// }

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	content := `
    <!doctype html>
    <html lang="en">
      <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>MonBan</title>
        <script src="https://unpkg.com/htmx.org@1.9.8"></script>
        <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.15/dist/tailwind.min.css" rel="stylesheet">
      </head>
    <body class="bg-gray-100">
        <div class="container mx-auto mt-5">
            <h1 class="text-3xl mb-4">Sleek HTMX Table</h1>
            <table class="bg-white rounded-lg shadow-lg overflow-hidden">
                <thead class="bg-gray-200">
                    <tr>
                        <th class="px-6 py-3 text-left font-semibold text-gray-700">Name</th>
                        <th class="px-6 py-3 text-left font-semibold text-gray-700">Age</th>
                        <th class="px-6 py-3 text-left font-semibold text-gray-700">Email</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td class="px-6 py-4 whitespace-nowrap">John Doe</td>
                        <td class="px-6 py-4 whitespace-nowrap">30</td>
                        <td class="px-6 py-4 whitespace-nowrap">john@example.com</td>
                    </tr>
                    <tr>
                        <td class="px-6 py-4 whitespace-nowrap">Jane Smith</td>
                        <td class="px-6 py-4 whitespace-nowrap">28</td>
                        <td class="px-6 py-4 whitespace-nowrap">jane@example.com</td>
                    </tr>
                    <!-- Add more rows as needed -->
                </tbody>
            </table>
        </div>
    </body>
    </html>
    `

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte(content))
}

func SetupWebApp() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Get("/", HomeHandler)
	return r
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
