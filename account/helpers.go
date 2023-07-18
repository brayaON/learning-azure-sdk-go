package account

import (
    "os"
    "log"
)


func GetAccountDetails() (string, string) {
    accountName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME")

    if len(accountName) == 0 {
	log.Fatal("AZURE_STORAGE_ACCOUNT_NAME is not set.")
    }

    accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT_KEY")
    if len(accountKey) == 0 {
	log.Fatal("AZURE_STORAGE_ACCOUNT_KEY is not set.")
    }

    return accountName, accountKey
}
