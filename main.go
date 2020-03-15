package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	cloudkms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

// name := "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
func encryptSymmetric(name string, plaintext []byte) ([]byte, error) {
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, err
	}

	req := &kmspb.EncryptRequest{
		Name:      name,
		Plaintext: plaintext,
	}
	resp, err := client.Encrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Ciphertext, nil
}

func EncryptFile(name string, secretsFile string) {
	// read plaintext content from file
	content, err := ioutil.ReadFile(secretsFile + ".json")
	if err != nil {
		log.Fatal(err)
	}
	// encrypt the plaintext content
	ciphertext, err := encryptSymmetric(name, content)

	// Write ciphertext to file
	ioutil.WriteFile(secretsFile+".enc", ciphertext, 0400)
}

// name := "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
func decryptSymmetric(name string, ciphertext []byte) ([]byte, error) {
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, err
	}

	// build the request
	req := &kmspb.DecryptRequest{
		Name:       name,
		Ciphertext: ciphertext,
	}
	// Call the api
	resp, err := client.Decrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Plaintext, nil
}

func DecryptFile(name string, secretsFile string) {
	// Decrypte file
	// read encrypted content from file
	//content, err := ioutil.ReadFile("./account.enc")
	content, err := ioutil.ReadFile(secretsFile + ".enc")
	if err != nil {
		log.Fatal(err)
	}

	// decrypt the encoded file
	//plaintext, err := decryptSymmetric("projects/"+projectID+"/locations/"+location+"/keyRings/"+keyring+"/cryptoKeys/"+key, content)
	plaintext, err := decryptSymmetric(name, content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plaintext))
}

func main() {
	// Project environment exported from .projectrc
	projectID := os.Getenv("PROJECT_ID")
	location := os.Getenv("LOCATION")
	keyring := os.Getenv("KMS_KEYRING")
	key := os.Getenv("KMS_KEY")
	secretsFile := os.Getenv("SECRETS_FILE_NAME")
	name := "projects/" + projectID + "/locations/" + location + "/keyRings/" + keyring + "/cryptoKeys/" + key

	EncryptFile(name, secretsFile)

	DecryptFile(name, secretsFile)

}
