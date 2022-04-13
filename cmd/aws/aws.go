package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "aws",
	Short: "Print HTTP request header for AWS",
	Long:  "Print Authorization header for AWS using default credentials",
	Run:   AWS,
}

func AWS(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("ERROR in LoadDefaultConfig: %s", err)
	}
	credentials, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		fmt.Printf("ERROR in Credentials.Retrieve: %s", err)
	}
	req, err := http.NewRequest("GET", "", nil)
	signer := v4.NewSigner()
	signer.SignHTTP(context.TODO(), credentials, req, "hash", "lambda", cfg.Region, time.Now())
	headerBytes, err := json.Marshal(req.Header)
	fmt.Println(string(headerBytes))
}
