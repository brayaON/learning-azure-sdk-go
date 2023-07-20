package account

import (
	"context"
	"log"
	"os"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const (
    containerName = "containergo"
)

func UploadBlobs() {
    ctx := context.Background()

    accountName := os.Getenv("AZURE_SA_NAME")
    accountSAS := os.Getenv("AZURE_SA_SAS")
    accountURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, accountSAS)
    wd, err := os.Getwd()
    if err != nil {
	log.Fatal(err)
    }
    pathToFiles := wd + "/data"


    // AzBlob Client
    blobClient, err := azblob.NewClientWithNoCredential(accountURL, nil)
    if err != nil {
	log.Fatal(err)
    }

    // Create container
    _, err = blobClient.CreateContainer(ctx, containerName, nil)
    if err != nil {
	log.Fatal(err)
    }

    // Upload files under /data directory
    dirHandler, err := os.ReadDir(pathToFiles)
    for _, entry := range dirHandler {
	filePath := pathToFiles + "/" + entry.Name()

	fileHandler, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
	    log.Fatal(err)
	}

	defer fileHandler.Close()

	_, err = blobClient.UploadFile(ctx, containerName, "root"+filePath, fileHandler, nil)
	if err != nil {
	    log.Fatal(err)
	}
    }
}
