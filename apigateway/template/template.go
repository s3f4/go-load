package template

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"text/template"
)

// InfraBuilderService is used to build new
// terraform files to create new digitalocean droplets
type InfraBuilderService interface {
	Parse(path string) (*bytes.Buffer, error)
	Write() error
}

type infraBuilder struct {
	Instances []string
}

// NewInfraBuilder returns a new infraBuilder instance
func NewInfraBuilder(instances []string) InfraBuilderService {
	return &infraBuilder{
		Instances: instances,
	}
}

// Parse template file
func (ib *infraBuilder) Parse(path string) (*bytes.Buffer, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return nil, nil
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, *ib)
	if err != nil {
		log.Print("execute: ", err)
		return nil, err
	}
	return &tpl, nil
}

// Write to template file
func (ib *infraBuilder) Write() error {
	f, err := os.Create("infra/workers.tf")
	if err != nil {
		return err
	}
	defer f.Close()

	tpl, err := ib.Parse("template/workers.tpl")
	if err != nil {
		return err
	}

	reader := bufio.NewReader(tpl)
	io.Copy(f, reader)

	return nil
}
