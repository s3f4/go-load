package template

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
	Parse(path string) (*bytes.Buffer, error)
	Write() error
}

type infraBuilder struct {
	region string
	size   string
	image  string
	count  int
}

type templateStruct struct {
	Region string
	Size   string
	Image  string
	Count  int
}

// NewInfraBuilder returns a new infraBuilder instance
func NewInfraBuilder(region, size, image string, count int) InfraBuilderService {
	return &infraBuilder{
		region: region,
		size:   size,
		image:  image,
		count:  count,
	}
}

// Parse template file
func (ib *infraBuilder) Parse(path string) (*bytes.Buffer, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return nil, nil
	}


	var ts templateStruct
	ts.Size = ib.size
	ts.Region = ib.region
	ts.Count = ib.count
	ts.Image = ib.image

	var tpl bytes.Buffer
	err = t.Execute(&tpl, ts)
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
