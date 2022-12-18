// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

var mainTemplate = `// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package {{.Package}}

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_{{.TypeRequest}} forwards the call to dev.CallMethod() then parses the payload of the reply as a {{.TypeReply}}.
func Call_{{.TypeRequest}}(ctx context.Context, dev *onvif.Device, request {{.TypeRequest}}) ({{.TypeReply}}, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			{{.TypeReply}} {{.TypeReply}}
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.{{.TypeReply}}, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "{{.TypeRequest}}")
		return reply.Body.{{.TypeReply}}, errors.Annotate(err, "reply")
	}
}
`

type parserEnv struct {
	Package     string
	Source      string
	TypeReply   string
	TypeRequest string
}

func main() {
	flag.Parse()
	env := parserEnv{
		Package: flag.Arg(0),
		Source:  flag.Arg(1),
	}

	body, err := template.New("body").Parse(mainTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	fin, err := os.Open(env.Source)
	if err != nil {
		log.Fatalf("Failed to open the configuration file [%s]: %v", env.Source, err)
	}
	defer func() { _ = fin.Close() }()

	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		method := strings.TrimSpace(scanner.Text())
		if method == "" {
			continue
		}
		env.TypeRequest = method
		env.TypeReply = method + "Response"
		log.Println(env)

		if fout, err := os.Create(env.TypeRequest + "_auto.go"); err != nil {
			log.Fatalln(err)
		} else {
			if err = body.Execute(fout, &env); err != nil {
				log.Fatalln(err)
			}
			fout.Close()
		}
	}
}
