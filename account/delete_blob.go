package account

import (
    "context"
    "log"
    "os"
    "fmt"
    "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func DeleteBlobs() {
    accountName := os.Getenv("AZURE_SA_NAME")
    accountSAS := os.Getenv("AZURE_SA_SAS")
    accountURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, accountSAS)
    containerName := "containergo"

    ctx := context.Background()

    wd, err := os.Getwd()
    if err != nil {
	log.Fatal(err)
    }
    filesPath := wd + "/data"

    // AzBlob Client
    blobClient, err := azblob.NewClientWithNoCredential(accountURL, nil)
    if err != nil {
	log.Fatal(err)
    }

    dirHandler, err := os.ReadDir(filesPath)
    if err != nil {
	log.Fatal(err)
    }

    for _, entry := range dirHandler {
	blobPath := "root" + filesPath + "/" + entry.Name()

	_, err := blobClient.DeleteBlob(ctx, containerName, blobPath, nil)
	if err != nil {
	    log.Fatal(err)
	}
    }
} 
