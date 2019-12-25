package main

import (
	"fmt"

	"github.com/99designs/iamy/iamy"
)

type PullCommandInput struct {
	Dir       string
	CanDelete bool
	SkipS3    bool
}

func PullCommand(ui Ui, input PullCommandInput) {
	aws := iamy.AwsFetcher{Debug: ui.Debug, SkipS3: input.SkipS3}
	data, err := aws.Fetch()
	if err != nil {
		ui.Error.Fatal(fmt.Printf("%s", err))
	}

	yaml := iamy.YamlLoadDumper{
		Dir: input.Dir,
	}
	err = yaml.Dump(data, input.CanDelete)
	if err != nil {
		ui.Error.Fatal(err)
	}
}
