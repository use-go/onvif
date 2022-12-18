package main

import (
	"context"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
	"text/template"
)

type CodegenDumpEnv struct {
	Package      string
	PackageUpper string
	Source       string

	Methods []string

	// The device package requires a special management because it is both a
	// main entry point of the onvif SDK but also a core internal API
	IsNotDevicePackage bool
}

func codegenDump(ctx context.Context, pkg, source string) error {
	const mainTemplate = `// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"context"{{if .IsNotDevicePackage}}
	"github.com/use-go/onvif/device"{{end}}
	"github.com/use-go/onvif/{{.Package}}"
)

type {{.PackageUpper}}Output struct { {{range $val := .Methods}}
	{{ $val | printf "%-40s" }} *{{$.Package}}.Get{{$val}}Response{{end}}
}

func detail{{.PackageUpper}}(ctx context.Context, dev *device.Device) {{.PackageUpper}}Output {
	var out {{.PackageUpper}}Output
	calls := make([]func(c context.Context), 0)
{{range $val := .Methods}}
	calls = append(calls, func(c context.Context) {
		if p, err := {{$.Package}}.Call_Get{{$val}}(c, dev, {{$.Package}}.Get{{$val}} {}); err == nil {
			out.{{$val}} = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "{{$val}}").Msg("{{$.Package}}")
		}
	})
{{end}}
	runAll(ctx, calls...)
	return out
}
`

	env := CodegenDumpEnv{
		Package:            pkg,
		IsNotDevicePackage: pkg != "device",
		PackageUpper:       cases.Title(language.BritishEnglish, cases.Compact).String(pkg),
		Source:             source,
		Methods:            make([]string, 0),
	}
	Logger.Info().Str("wd", getwd()).Str("file", env.Source).Interface("env", env).Msg("rendering")

	body, err := template.New("file").Parse(mainTemplate)
	if err != nil {
		Logger.Fatal().Err(err).Msg("BUG: invalid template")
	}

	for _, method := range getMethods(env.Source) {
		if !strings.HasPrefix(method.Name, "Get") || !method.NoArg {
			continue
		}
		env.Methods = append(env.Methods, method.Name[3:])
	}

	if fout, err := os.Create("dump_" + env.Package + "_auto.go"); err != nil {
		Logger.Fatal().Err(err).Str("wd", getwd()).Str("file", env.Source).Msg("file creation failure")
	} else {
		if err = body.Execute(fout, env); err != nil {
			Logger.Fatal().Err(err).Str("wd", getwd()).Str("file", env.Source).Msg("file generation failure")
		}
		fout.Close()
	}

	return nil
}
