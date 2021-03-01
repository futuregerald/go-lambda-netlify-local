package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// ClientContextShape is the context netlify injects in to lambda functions
type clientContext struct {
	Client struct {
		InstallationID string `json:"installation_id"`
		AppTitle       string `json:"app_title"`
		AppVersionCode string `json:"app_version_code"`
		AppPackageName string `json:"app_package_name"`
	} `json:"Client"`
	Env    interface{} `json:"env"`
	Custom struct {
		Netlify string `json:"netlify"`
	} `json:"custom"`
}

// netlifyData is the Netlify params inside the clientContext.Custom
type netlifyData struct {
	Identity struct {
		URL   string `json:"url"`
		Token string `json:"token"`
	} `json:"identity"`
	User struct {
		AppMetadata struct {
			Roles []string `json:"roles"`
		} `json:"app_metadata"`
		Email        string `json:"email"`
		Exp          int    `json:"exp"`
		Sub          string `json:"sub"`
		UserMetadata struct {
			AvatarURL string `json:"avatar_url"`
			FullName  string `json:"full_name"`
		} `json:"user_metadata"`
	} `json:"user"`
	SiteURL string `json:"site_url"`
}

func main() {
	http.HandleFunc("/", lambdaHandler)
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Fatal(http.ListenAndServe(":3000", nil))
	} else {
		log.Fatal(gateway.ListenAndServe(":3000", nil))
	}
}

func lambdaHandler(w http.ResponseWriter, r *http.Request) {
	var decodedNetlifyInfo netlifyData
	lc, ok := lambdacontext.FromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodedNetlifyInfo := lc.ClientContext.Custom["netlify"]
	decoded, err := base64.StdEncoding.DecodeString(encodedNetlifyInfo)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	err = json.Unmarshal(decoded, &decodedNetlifyInfo)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprintf(w, "Hello there! Your site is %s", decodedNetlifyInfo.SiteURL)
}
