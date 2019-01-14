package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"

	rz "github.com/adrianpk/rezerw/core"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}
	// Read
	accounts, err := readAccounts("accounts.json")
	if err != nil {
		log.Fatal(err)
	}
	// Loop
	for _, account := range accounts {
		fmt.Println("Creating account...")
		err = insertAccount(cfg, account)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Account ", account.Name, "created.")
	}
}

func readAccounts(fileName string) ([]rz.Account, error) {
	accounts := make([]rz.Account, 0)
	// Read file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return accounts, err
	}
	// Unmarshal
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return accounts, err
	}
	// Result
	return accounts, nil
}

func insertAccount(cfg aws.Config, account rz.Account) error {
	fmt.Printf("%+v", account)
	item, err := dynamodbattribute.MarshalMap(account)
	if err != nil {
		return err
	}
	// DynamoDB instance
	svc := dynamodb.New(cfg)
	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String("rezerw-accounts"),
		Item:      item,
	})
	_, err = req.Send()
	if err != nil {
		return err
	}
	return nil
}
