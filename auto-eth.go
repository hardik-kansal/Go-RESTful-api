package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hardik-kansal/gobank/ewt"
)

func WithEWTAuth(handlerFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetTokenFromRequest(r)
		check,token := validateEWT(tokenString)
		if check == false {
			log.Printf("failed to validate token")
			permissionDenied(w)
			return
		}
		log.Printf("address of user authorized is %v",token.SigResponse.Address)
		handlerFunc(w, r)
	}
}

func CreateEWT(addr string, sig string, mssg string) string {
	sigRes := ewt.SignatureResponse{
		Address: addr,
		Msg:     mssg,
		Sig:     sig,
		Version: "2",
	}
	token := ewt.Token{
		SigResponse: sigRes,
		Expirydate:  time.Now().Add(10 * 24 * time.Hour),
		Valid:       true,
	}
	return token.String()
}

func validateEWT(token1 string) (bool,ewt.Token) {
	var token ewt.Token
	token.FromJSON(token1)
	check := ewt.VerifySig(token.SigResponse.Sig, token.SigResponse.Address, token.SigResponse.Msg)
	if time.Now().Before(token.Expirydate) {
		token.Valid = true
	}

	return check,token
}
func createAndSetEWTCookie(addr string, mssg string, sig string, w http.ResponseWriter) (string, error) {
	token := CreateEWT(addr, sig, mssg)
	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
