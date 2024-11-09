package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goschool/crud/types"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func HandleEchoUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "this is the email: %s", string(user.Email))
}
