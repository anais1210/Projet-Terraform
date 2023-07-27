package main

import (
	"context"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
)
func rootAccess(){

	config := vault.DefaultConfig()

	config.Address = "http://127.0.0.1:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	// Authenticate
	client.SetToken("myroot")

	secretData := map[string]interface{}{
		"password": "Hashi123",
	}

	// Write a secret
	_, err = client.KVv2("secret").Put(context.Background(), "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	fmt.Println("Secret written successfully by root.")

	// Read a secret from the default mount path for KV v2 in dev mode, "secret"
	secret, err := client.KVv2("secret").Get(context.Background(), "my-secret-password")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	if value != "Hashi123" {
		log.Fatalf("unexpected password value %q retrieved from vault", value)
	}

	fmt.Println("Access granted!")
}

func anaisAccess(){
	// Initialise une nouvelle instance du client Vault
	client, err := vault.NewClient(&vault.Config{
		Address: "http://localhost:8200", // Replace with your Vault server address
	})
	if err != nil {
		log.Fatalf("Erreur lors de la création du client Vault: %v", err)
	}

	// Replace with the username and password for authentication
	username := "anais"
	password := "esgi2023"

	// Authenticate using the userpass method
	// The path "/auth/userpass/login/<username>" is used for userpass authentication
	path := fmt.Sprintf("auth/userpass/login/%s", username)
	data := map[string]interface{}{
		"password": password,
	}
	secret, err := client.Logical().Write(path, data)
	if err != nil {
		log.Fatalf("Erreur lors de l'authentification auprès de Vault: %v", err)
	}

	// Check if authentication was successful and get the token
	if secret != nil && secret.Auth != nil {
		client.SetToken(secret.Auth.ClientToken) // Set the token for subsequent requests
		fmt.Println("User Anais, Authentication successful.")
	} else {
		log.Fatal("Échec de l'authentification.")
	}

	// Now, you can use the authenticated client to access and interact with secrets in Vault.
	secretData, err := client.KVv2("kv").Get(context.Background(), "mysecret")
	if err != nil {
		fmt.Printf("Error reading secret from Vault: %s\n", err)
		return
	}

	// Check if the secret exists
	if secretData == nil || secretData.Data == nil {
		fmt.Printf("Secret not found")
		return
	}

	// Print the secret data
	for key, value := range secretData.Data {
		fmt.Printf("Secret fetch by Anais %s: %v\n", key, value)
	}

	//write secret data
	newData := map[string]interface{}{
		"newdata": "newdata",
	}
	_, err = client.KVv2("secret").Put(context.Background(), "newsecret", newData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	fmt.Println("Secret written successfully by anais.")
	
}
func totoAccess(){
	// Initialise une nouvelle instance du client Vault
	client, err := vault.NewClient(&vault.Config{
		Address: "http://localhost:8200", // Replace with your Vault server address
	})
	if err != nil {
		log.Fatalf("Erreur lors de la création du client Vault: %v", err)
	}

	// Replace with the username and password for authentication
	username := "toto"
	password := "esgi2023"

	// Authenticate using the userpass method
	// The path "/auth/userpass/login/<username>" is used for userpass authentication
	path := fmt.Sprintf("auth/userpass/login/%s", username)
	data := map[string]interface{}{
		"password": password,
	}
	secret, err := client.Logical().Write(path, data)
	if err != nil {
		log.Fatalf("Erreur lors de l'authentification auprès de Vault: %v", err)
	}

	// Check if authentication was successful and get the token
	if secret != nil && secret.Auth != nil {
		client.SetToken(secret.Auth.ClientToken) // Set the token for subsequent requests
		fmt.Println("User Toto, Authentication successful.")
	} else {
		log.Fatal("Échec de l'authentification.")
	}

	// Now, you can use the authenticated client to access and interact with secrets in Vault.
	secretData, err := client.KVv2("kv").Get(context.Background(), "mysecret")
	if err != nil {
		fmt.Printf("Error reading secret from Vault: %s\n", err)
		return
	}

	// Check if the secret exists
	if secretData == nil || secretData.Data == nil {
		fmt.Printf("Secret not found")
		return
	}

	// Print the secret data
	for key, value := range secretData.Data {
		fmt.Printf("Secret fetch by toto %s: %v\n", key, value)
	}

	//write secret data
	newData := map[string]interface{}{
		"newdata": "newdata",
	}
	_, err = client.KVv2("secret").Put(context.Background(), "newsecret", newData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	fmt.Println("Secret written successfully by toto.")
	
}
func main() {
	anaisAccess()
	rootAccess()
	totoAccess()
	
}
