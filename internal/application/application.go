package application

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tsha-en/Calculator_App/pkg/calculator"
)

type Config struct {
	Addr string
}

func ConfigFromFile() *Config {
	config := new(Config)
	config.Addr = "8080"

	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromFile(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {

	request := new(Request)
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `{
    "error": "Internal server error"
}`, http.StatusInternalServerError)
		return
	}

	result, err := calculator.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{
    "error": "Expression is not valid"
}`)

	} else {

		fmt.Fprintf(w, `{
	"result": "%g"
}`, result)
	}
}

func RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":80", nil)
}
