package app

import (
	"encoding/json"
	"net/http"
	"os"
)

type Param struct {
	Foo  string `json:"foo"`
	Bizz string `json:"bizz"`
}

type Output struct {
	Parameters []Param `json:"parameters"`
}

type OutputParams struct {
	Output Output `json:"output"`
}

// unsupported returns a 404 error
func unsupported(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unsupported request", http.StatusNotFound)
}

// getparams returns a JSON response with the parameters passed in
func getparams(w http.ResponseWriter, r *http.Request) {
	// Load token information from file
	tokenfile := "/var/run/argo/token"
	token, err := os.ReadFile(tokenfile)
	if err != nil {
		http.Error(w, "System Error", http.StatusInternalServerError)
		return
	}

	// Check to see if they provided the right token
	t := r.Header.Get("Authorization")
	if t != "Bearer "+string(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// set op equal to a new OutputParams struct with dummy data
	op := OutputParams{
		Output: Output{
			Parameters: []Param{
				{
					Foo:  "bar",
					Bizz: "buzz",
				},
			},
		},
	}

	// set the right header for JSON
	w.Header().Add("Content-Type", "application/json")

	// encode struct as JSON and send it back
	json.NewEncoder(w).Encode(op)
}
