package aws

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
)

var (
	service string
	methods = [...]string{"GET", "POST", "PUT", "DELETE", "HEAD", "CONNECT", "OPTIONS", "PATCH", "TRACE"}
	Cmd     = &cobra.Command{
		Use:   "aws",
		Short: "Print HTTP request header for AWS",
		Long:  "Print Authorization header for AWS using default credentials",
		Run:   AWS,
	}
)

func AWS(cmd *cobra.Command, args []string) {
	method, url, _, _, payload := parseArgs(args)

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

func parseArgs(args []string) (string, string, map[string][]string, map[string][]string, string) {
	// TODO define struct for return value
	var (
		method  string
		url     string
		headers map[string][]string
		queries map[string][]string
		payload string
	)

	var kvList []string

	// check if the first arg is a method
	for _, m := range methods {
		if args[0] == m {
			method = m
			url = args[1]
			kvList = args[2:]
			break
		}
	}

	for _, kv := range kvList {
		byColon := strings.Split(kv, ":")
		if len(byColon) == 2 {
			headers[byColon[0]] = append(headers[byColon[0]], byColon[1])
		}

		// bySingleEqual := strings.Split(kv, "=")
		// if len(byColon) == 2 {
		// 	headers[bySingleEqual[0]] = append(payload[bySingleEqual[0]], bySingleEqual[1])
		// }

		byDoubleEqual := strings.Split(kv, "==")
		if len(byDoubleEqual) == 2 {
			queries[byDoubleEqual[0]] = append(queries[byDoubleEqual[0]], byDoubleEqual[1])
		}
	}

	// set default method if not specified
	if method == "" {
		url = args[0]
		if payload == "" {
			method = "GET"
		} else {
			method = "POST"
		}
	}

	return method, url, headers, queries, payload
}

func getPayloadHash(payload string) string {
	hashed := sha256.Sum256([]byte(payload))
	return hex.EncodeToString(hashed[:])
}

func init() {
	Cmd.Flags().StringVarP(&service, "service", "s", "", "")
	Cmd.MarkFlagRequired("service")
}
