package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type KubernetesSecret struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Data       map[string]string `yaml:"data"`
}

func createKubernetesSecret(secretName, stringData string) (string, error) {
	encodedData := base64.StdEncoding.EncodeToString([]byte(stringData))

	secret := KubernetesSecret{
		APIVersion: "v1",
		Kind:       "Secret",
		Metadata: map[string]string{
			"name": secretName,
		},
		Data: map[string]string{
			secretName: encodedData,
		},
	}

	yamlData, err := yaml.Marshal(&secret)
	if err != nil {
		return "", err
	}
	return string(yamlData), nil
}

func main() {
	secretName := flag.String("name", "", "The name of the secret in the metadata of the created Kubernetes secret. This is also used as the name for the data element itself.")
	secretData := flag.String("data", "", "The secret value")
	base64Encode := flag.Bool("base64", false, "If specified, the generated result will be encoded in base64")

	flag.Parse()

	if *secretName == "" || *secretData == "" {
		fmt.Println("Both --name and --data are required.")
		flag.Usage()
		os.Exit(1)
	}

	yamlSecret, err := createKubernetesSecret(*secretName, *secretData)
	if err != nil {
		fmt.Println("Error creating Kubernetes secret:", err)
		os.Exit(1)
	}

	if *base64Encode {
		encodedYamlSecret := base64.StdEncoding.EncodeToString([]byte(yamlSecret))
		fmt.Print(encodedYamlSecret)
	} else {
		fmt.Print(yamlSecret)
	}
}
