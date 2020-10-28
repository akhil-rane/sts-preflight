package jwt

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sjenning/sts-preflight/pkg/cmd/create"
	"github.com/sjenning/sts-preflight/pkg/cmd/token"
)

func New(config token.Config) {
	var state create.State
	state.Read()

	privateKey, err := ioutil.ReadFile("_output/sa-signer")
	if err != nil {
		log.Fatal(err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "openshift-install",
		"aud": "sts.amazonaws.com",
		"iss": fmt.Sprintf("https://s3-%s.amazonaws.com/%s-installer", state.Region, state.InfraName),
		"exp": time.Now().Unix() + config.ExpireSeconds,
		"iat": time.Now().Unix(),
	})

	token.Header["kid"] = state.Kid
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}

	tokenFile := "_output/token"
	f, err := os.Create(tokenFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(tokenString)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Token written to ", tokenFile)
}
