package account

import (
    "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
    "os"
    "fmt"
    "context"
    "log"
)

func ListBlobs() {
    accountName := os.Getenv("AZURE_SA_NAME")
    accountKey := os.Getenv("AUZRE_SA_KEY")
    accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)
    containerName := "containergo"

    ctx := context.Background()

    // Auth
    cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
    if err != nil {
	log.Fatal(err)
    }

    // Blob Client
    blobClient, err := azblob.NewClientWithSharedKeyCredential(accountURL, cred, nil)
    if err != nil {
	log.Fatal(err)
    }

    // Get blobs
    pager := blobClient.NewListBlobsFlatPager(containerName, nil)

    for pager.More() {
	page, err := pager.NextPage(ctx)
	if err != nil {
	    log.Fatal(err)
	}

	for _, blob := range page.Segment.BlobItems {
	    fmt.Println(*blob.Name)
	}
    }
}
