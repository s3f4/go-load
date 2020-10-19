package template

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

// InfraBuilderService is used to build new
// terraform files to create new digitalocean droplets
type InfraBuilderService interface {
	Parse() (*bytes.Buffer, error)
	Write() error
}

type infraBuilder struct {
	TemplatePath string
	TargetPath   string
	Vars         map[string]interface{}
}

// NewInfraBuilder returns a new infraBuilder instance
func NewInfraBuilder(templatePath, targetPath string, vars map[string]interface{}) InfraBuilderService {
	return &infraBuilder{
		TemplatePath: templatePath,
		TargetPath:   targetPath,
		Vars:         vars,
	}
}

// Parse template file
func (ib *infraBuilder) Parse() (*bytes.Buffer, error) {
	t, err := template.ParseFiles(ib.TemplatePath)
	if err != nil {
		log.Print(err)
		return nil, nil
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, ib.Vars)
	if err != nil {
		log.Print("execute: ", err)
		return nil, err
	}
	return &tpl, nil
}

// Write to template file
func (ib *infraBuilder) Write() error {
	f, err := os.Create(ib.TargetPath)
	if err != nil {
		return err
	}
	defer f.Close()

	tpl, err := ib.Parse()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(tpl)
	b, _ := ioutil.ReadAll(reader)
	fmt.Println(string(b))
	io.Copy(f, strings.NewReader(string(b)))

	return nil
}
