package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type ReqParams struct {
	Status string `json:"status"`
}
type ReqInput struct {
	Parameters ReqParams `json:"parameters"`
}

type InputParams struct {
	ApplicationSetName string   `json:"applicationSetName"`
	Input              ReqInput `json:"input"`
}

type Param struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Label     string `json:"label"`
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
	// Set some defaults
	var label string = "devops-is-awesome"

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

	// Read the parameters from the request in a unstructured way
	// {"applicationSetName":"myappset","input":{"parameters":{"status":"gitops"}}}
	var result InputParams
	b, _ := io.ReadAll(r.Body)
	r.Body.Close()
	if err = json.Unmarshal(b, &result); err != nil {
		fmt.Println(err)
	} else {
		// print out the value of .input.parameters.status
		if result.Input.Parameters.Status == "gitops" {
			label = "gitops-is-the-way"
		}
	}

	// Log the request
	log.Info(string(b))

	// set op equal to a new OutputParams struct with dummy data
	op := OutputParams{
		Output: Output{
			Parameters: []Param{
				{
					Name:      "foo",
					Namespace: "bar",
					Label:     label,
				},
				{
					Name:      "fuzz",
					Namespace: "bazz",
					Label:     label,
				},
			},
		},
	}

	// set the right header for JSON
	w.Header().Add("Content-Type", "application/json")

	// encode struct as JSON and send it back
	json.NewEncoder(w).Encode(op)
}
