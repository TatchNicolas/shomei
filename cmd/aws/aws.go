package aws

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
)

var (
	payload string
	service string

	Cmd = &cobra.Command{
		Use:   "aws",
		Short: "Print HTTP request header for AWS",
		Long:  "Print Authorization header for AWS using default credentials",
		Run:   AWS,
	}
)

func AWS(cmd *cobra.Command, args []string) {
	method := args[0]
	url := args[1]

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("ERROR in LoadDefaultConfig: %s", err)
	}
	credentials, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		fmt.Printf("ERROR in Credentials.Retrieve: %s", err)
	}
	req, err := http.NewRequest(method, url, nil)
	signer := v4.NewSigner()
	signer.SignHTTP(context.TODO(), credentials, req, getPayloadHash(payload), service, cfg.Region, time.Now())

	// TODO compose string with different format for curl and httpie
	// for k, v := range req.Header {
	// }

	// as JSON
	// headerBytes, err := json.Marshal(req.Header)
	// fmt.Println(string(headerBytes))

	// DEBUG
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR in client.Do: %s", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR in ReadAll: %s", err)
	}
	fmt.Println(string(bodyBytes))
}

func getPayloadHash(payload string) string {
	hashed := sha256.Sum256([]byte(payload))
	return hex.EncodeToString(hashed[:])
}

func init() {
	Cmd.Flags().StringVarP(&payload, "payload", "p", "", "Request payload")
	Cmd.Flags().StringVarP(&service, "service", "s", "", "Request payload")
}
