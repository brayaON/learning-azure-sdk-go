package account

import (
    "github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/share"
    "os"
    "log"
    "context"
)

const (
    shareName = "fsbof234"
    dirName = "dirbof234"
    fileName = "filebof234"
)

func CreateFile() {
    connStr := os.Getenv("AZURE_SA_CONN_STR")
    localFilePath := "./account/file_share.go"

    shareClient, err := share.NewClientFromConnectionString(connStr, shareName, nil)
    if err != nil {
	log.Fatal(err)
    }

    _, err = shareClient.Create(context.TODO(), nil)
    if err != nil {
	log.Fatal(err)
    }

    dirClient := shareClient.NewDirectoryClient(dirName)
    _, err = dirClient.Create(context.TODO(), nil)
    if err != nil {
	log.Fatal(err)
    }

    file, err := os.OpenFile(localFilePath, os.O_RDONLY, 0)
    if err != nil {
	log.Fatal(err)
    }

    defer file.Close()

    fileClient := dirClient.NewFileClient(fileName)
    err = fileClient.UploadFile(context.Background(), file, nil)
    if err != nil {
	log.Fatal(err)
    }

}
