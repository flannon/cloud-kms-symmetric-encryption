#### cloud-kms-symmetric-encryption

<p>
An example of symmetric key encryption using the Goole glient library for Go.
</p>


This example assumes a service account with KMS access, and a service account key at ./account.json.  It also assumes a single KMS key ring in the project and a single symmetic key in the keyring.  For multiple key rings, or multiple keys refer to .projectrc and se the values appropriately.  Running main.go will encrypte the contents of the file "nonesence.json", write the encrypted contents to nonesence.enc, and then decrypt nonsense.enc and print the contents.

    . .projectrc
    go run main.go