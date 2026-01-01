// main.go
package main

import(
	"fmt"
	"net/http"
	"html/template"
)

// "real-time-forum/internal/database"
func main() {
	// Entry point for the server
	// database.InitDB()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../web/css"))))

	http.HandleFunc("/", homeHandler)



	// Start server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse template file
	tmpl, err := template.ParseFiles("../web/index.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render template
	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Println("ansabsbas",err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}