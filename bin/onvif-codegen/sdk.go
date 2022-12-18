package main

import (
	"context"
	"log"
	"os"
	"text/template"
)

type CodegenSdkEnv struct {
	Package     string
	Source      string
	TypeReply   string
	TypeRequest string

	// The device package requires a special management because it is both a
	// main entry point of the onvif SDK but also a core internal API
	IsNotDevicePackage bool
}

func codegenSdk(ctx context.Context, pkg, source string) error {
	const mainTemplate = `// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package {{.Package}}

import (
	"context"
	"github.com/juju/errors"{{if .IsNotDevicePackage}}
	"github.com/use-go/onvif/device"{{end}}
	"github.com/use-go/onvif/networking"
)

// Call_{{.TypeRequest}} forwards the call to dev.CallMethod() then parses the payload of the reply as a {{.TypeReply}}.
func Call_{{.TypeRequest}}(ctx context.Context, dev *{{if .IsNotDevicePackage}}device.{{end}}Device, request {{.TypeRequest}}) ({{.TypeReply}}, error) {
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

	env := CodegenSdkEnv{
		Package:            pkg,
		IsNotDevicePackage: pkg != "device",
		Source:             source,
	}

	body, err := template.New("body").Parse(mainTemplate)
	if err != nil {
		Logger.Fatal().Err(err).Msg("BUG: invalid template")
	}

	for _, method := range getMethods(env.Source) {
		env.TypeRequest = method.Name
		env.TypeReply = method.Name + "Response"
		log.Println(env)

		if fout, err := os.Create(env.TypeRequest + "_auto.go"); err != nil {
			Logger.Fatal().Err(err).Str("wd", getwd()).Str("file", env.Source).Msg("file creation failure")
		} else {
			if err = body.Execute(fout, &env); err != nil {
				Logger.Fatal().Err(err).Str("wd", getwd()).Str("file", env.Source).Msg("file generation failure")
			}
			fout.Close()
		}
	}

	return nil
}
