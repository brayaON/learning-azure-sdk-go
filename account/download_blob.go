package account

import (
    "fmt"
    "os"
    "log"
    "context"
    "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func DownloadBlobs() {
    accountName := os.Getenv("AZURE_SA_NAME")
    accountKey := os.Getenv("AZURE_SA_KEY")
    accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)
    cw, err := os.Getwd()
    if err != nil {
	log.Fatal(err)
    }

    ctx := context.Background()
    
    // Authentication
    cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
    if err != nil {
	log.Fatal(err)
    }

    // Azblob client
    blobClient, err := azblob.NewClientWithSharedKeyCredential(accountURL, cred, nil)
    if err != nil {
	log.Fatal(err)
    }

    err = os.Mkdir("./data2", 0755)
    if err != nil {
	log.Fatal(err)
    }

    dirHandler, err := os.ReadDir("./data")
    if err != nil {
	log.Fatal(err)
    }

    for _, entry := range dirHandler {
	newPath := "./data2/"+entry.Name()
	newFile, err := os.Create(newPath)
	blobPath := "containergo/root" + cw + "/data"
	if err != nil {
	    log.Fatal(err)
	}

	defer newFile.Close()

	_, err = blobClient.DownloadFile(ctx, blobPath, entry.Name(), newFile, nil)
	if err != nil {
	    log.Fatal(err)
	}
    }
}
