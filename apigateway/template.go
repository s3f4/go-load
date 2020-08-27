package apigateway

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"log"
	"os"
)

// InfraBuilderService is used to build new
// terraform files to create new digitalocean droplets
type InfraBuilderService interface {
	Parse(path string) *bytes.Buffer
	Write()
}

type infraBuilder struct {
	region string
	size   string
	count  int
}

// NewInfraBuilder returns a new infraBuilder instance
func NewInfraBuilder() InfraBuilderService {
	return &infraBuilder{}
}

// Parse template file
func Parse(path string) *bytes.Buffer {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	config := map[string]string{
		"region": "AMS3",
		"size":   "1GB",
		"index":  "22",
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, config)
	if err != nil {
		log.Print("execute: ", err)
		return nil
	}
	return &tpl
}

// Write to template file
func Write() {
	f, _ := os.Create("infra/workers.tf")
	tpl := Parse("template/instance.tpl")
	reader := bufio.NewReader(tpl)
	io.Copy(f, reader)
	f.Close()
}
