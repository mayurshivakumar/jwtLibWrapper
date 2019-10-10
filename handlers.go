package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	// assuming authentation is already done.
	jwtLib := GetJwtLib(os.Getenv(`mySigningKey`))
	claims := &Claims{
		UserID: "1",
		Admin:  true,
	}
	claims.StandardClaims.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
	tokenStr, err := jwtLib.Generate(claims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenStr,
	})

}

// Home ...
func Home(w http.ResponseWriter, r *http.Request) {
	//sample json out put
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
	w.Header().Set("Content-Type", "application/json")
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
		Phone: "000099999",
	}

	json.NewEncoder(w).Encode(user)
}

// isAuthorized ...
func isAuthorized(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := c.Value
		jwtLib := GetJwtLib(os.Getenv(`mySigningKey`))
		valid, err := jwtLib.Validate(token)
		if err != nil {
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		endpoint.ServeHTTP(w, r)
	})
}
