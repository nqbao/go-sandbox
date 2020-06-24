package restapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/gorilla/context"
	"github.com/nqbao/learn-go/dynonote/service"
)

type tokenRequest struct {
	IDToken string `json:"id_token"`
}

type _authHandler struct {
	handler http.Handler
}

func (h _authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: don't restrict access to all endpoint
	if r.Method != "OPTIONS" && r.URL.Path != "/api/tokensignin" {
		authHeader := r.Header["Authorization"]
		if len(authHeader) > 0 {
			svc := cognitoidentity.New(service.NewSession())

			logins := map[string]*string{
				"accounts.google.com": aws.String(authHeader[0]),
			}

			ids, err := svc.GetId(&cognitoidentity.GetIdInput{
				IdentityPoolId: aws.String("us-east-1:660f9f63-aae4-4954-96a6-12589e06e40c"),
				Logins:         logins,
			})

			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{
					"status":  "error",
					"message": fmt.Sprintf("%v", err),
				})
				return
			}

			// now get credentials
			creds, err := svc.GetCredentialsForIdentity(&cognitoidentity.GetCredentialsForIdentityInput{
				IdentityId: ids.IdentityId,
				Logins:     logins,
			})

			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{
					"status":  "error",
					"message": fmt.Sprintf("%v", err),
				})
				return
			}

			context.Set(r, "aws_credentials", creds)

			fmt.Printf("%v\n", creds)
		}
	}

	h.handler.ServeHTTP(w, r)
}

func authHandler(h http.Handler) http.Handler {
	return _authHandler{handler: h}
}

func tokenSignIn(w http.ResponseWriter, r *http.Request) {
	req := &tokenRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	writer := json.NewEncoder(w)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writer.Encode(map[string]string{
			"status":  "error",
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		endpoint := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%v", req.IDToken)

		client := http.Client{
			Timeout: 5 * time.Second,
		}

		repr, err := client.Get(endpoint)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writer.Encode(map[string]string{
				"status":  "error",
				"message": fmt.Sprintf("%v", err),
			})
		} else {
			w.WriteHeader(repr.StatusCode)
			io.Copy(w, repr.Body)
		}
	}
}
