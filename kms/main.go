package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {
	sess := session.Must(session.NewSession(
		&aws.Config{Region: aws.String("us-east-1")},
	))

	c := kms.New(sess)

	aliases, err := c.ListAliases(&kms.ListAliasesInput{})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", aliases)

	encReq := &kms.EncryptInput{
		KeyId:     aws.String("alias/vault"),
		Plaintext: []byte("hello world"),
	}

	enc, err := c.Encrypt(encReq)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", enc.CiphertextBlob)

	decReq := &kms.DecryptInput{
		CiphertextBlob: enc.CiphertextBlob,
	}

	dec, err := c.Decrypt(decReq)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(dec.Plaintext))
}
