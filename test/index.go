package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type ReleaseLog struct {
	Deployer     string
	LineMessages []string
	CreatedAt    string
	Status       string
}

type PageData struct {
	Email         string
	Services      []string
	OptEnvs       []string
	Service       string
	OptEnv        string
	ServiceOwners string
	ToTag         string
	FromTag       string
	ReleaseLogs   []ReleaseLog
}

func CreateServiceArray(count int) []string {
	var services []string
	for i := 1; i <= count; i++ {
		services = append(services, "Service "+strconv.Itoa(i))
	}
	return services
}

func CreateReleaseLogArray(count int) []ReleaseLog {
	var releaseLogs []ReleaseLog
	for i := 1; i <= count; i++ {
		releaseLogs = append(releaseLogs, ReleaseLog{
			Deployer:     "Duy Bui " + strconv.Itoa(i),
			LineMessages: []string{"Message 1", "Message 2", "Message 3"},
			CreatedAt:    "2024-01-01 00:00:00",
			Status:       "Success",
		})
	}
	return releaseLogs
}

var pageResult = PageData{
	Email:         "test@example.com",
	Services:      CreateServiceArray(55),
	OptEnvs:       []string{"develop", "prod", "stage"},
	Service:       "demo",
	OptEnv:        "prod",
	ServiceOwners: "Kiwi 1, Kiwi 2, Kiwi 3",
	ToTag:         "v1.0.0",
	FromTag:       "v1.0.1",
	ReleaseLogs:   CreateReleaseLogArray(5),
}

// Mock data for testing template

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../release/index.html"))
		data := pageResult
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/services/*", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../release/option.html"))
		data := pageResult
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/services/demo/detail", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../release/detail.html"))
		data := pageResult
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/services/demo/confirm", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../release/release.html"))
		data := pageResult
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
