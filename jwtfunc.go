package seamless

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	appID = "1e0yVt1zUCgFKOhvCNk1"
)

func CreateJWT(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username required"))
		return
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": appID,
		"sub": username,
	})

	// Sign and get the complete encoded token as a string using the secret
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatalln("Must set SECRET env var")
	}
	tokenString, err := token.SignedString([]byte(secret))
	fmt.Println(tokenString, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("problem signing token"))
		return
	}
	resp := map[string]interface{}{}
	resp["token"] = tokenString
	b, err := json.Marshal(resp)
	if err != nil {
		log.Println("error marshalling response!", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("problem signing token"))
		return
	}
	fmt.Fprint(w, string(b))
}
