package bootstrap

import (
	"flag"
	"fmt"
	"github.com/hetiansu5/urlquery"
	"github.com/spf13/pflag"
	"testing"
)

type TestData struct {
	V string   `query:"v"`
	A int      `query:"a"`
	C []string `query:"c"`
	D struct {
		A string `query:"a"`
		C string `query:"c"`
	} `query:"d"`
}

func TestUrlQuery(t *testing.T) {
	url := "v=ad&a=1&c[]=a&c[]=s&d[a]=2"
	var data TestData
	err := urlquery.Unmarshal([]byte(url), &data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
}

func TestPFlags(t *testing.T) {
	f := pflag.NewFlagSet("aa", pflag.ContinueOnError)
	f.AddGoFlagSet(flag.CommandLine)
	name := f.StringP("name", "n", "xiaoming", "name")
	err := f.Parse([]string{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(*name)
}
