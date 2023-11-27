package handler

import "net/http"

func GetHomeHandler(rw http.ResponseWriter, r *http.Request) {
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
                  <tr hx-get="/apikey"" hx-trigger="load">
                    <td class="px-6 py-4 whitespace-nowrap">John Doe</td>
                    <td class="px-6 py-4 whitespace-nowrap">30</td>
                    <td class="px-6 py-4 whitespace-nowrap">john@example.com</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <button hx-get="/delete-row/1" hx-confirm="Are you sure?" hx-target="this" hx-swap="outerHTML" class="text-red-600">X</button>
                    </td>
                  </tr>
                  <!-- Add more rows as needed -->
                  <tr>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <button hx-get="/add-row" hx-trigger="click" class="text-green-600">+</button>
                    </td>
                  </tr>
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
