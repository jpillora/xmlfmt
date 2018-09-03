package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jpillora/opts"
)

//Version of this tool, set at build time
var Version = "0.0.0-src"

var config = struct {
	File     string   `type:"arg" help:"path to xml [file]"`
	Write    bool     `type:"opt" help:"write over xml file with formatted"`
	Settings settings `type:"embedded"`
}{
	File: "-",
}

type settings struct {
	MaxWidth int `type:"opt" help:"max width of the file in characters (default unlimited)"`
}

func main() {
	//parse cli
	opts.
		New(&config).
		Repo("github.com/jpillora/xmlfmt").
		Version(Version).
		Parse()
	//run program
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var r io.Reader
	//choose input
	if config.File != "" && config.File != "-" {
		f, err := os.Open(config.File)
		if err != nil {
			return err
		}
		defer f.Close()
		r = f
	} else {
		r = os.Stdin
	}
	//validate
	if config.File == "" && config.Write {
		return errors.New("file missing")
	}
	var w io.Writer
	var b bytes.Buffer
	//choose output
	if config.Write {
		w = &b
	} else {
		w = os.Stdout
	}
	//format
	if err := xmlfmt(r, w, config.Settings); err != nil {
		return err
	}
	//write results to file?
	if config.Write {
		if err := ioutil.WriteFile(config.File, b.Bytes(), 0x777); err != nil {
			return err
		}
		log.Printf("wrote '%s'", config.File)
	}
	//done
	return nil
}

func xmlfmt(r io.Reader, w io.Writer, s settings) error {
	//internal state
	indent := 0
	inner := false
	var last xml.Token
	//local best-effort escape function
	esbuff := bytes.Buffer{}
	escape := func(s string) string {
		if err := xml.EscapeText(&esbuff, []byte(s)); err == nil {
			s = esbuff.String()
		}
		esbuff.Reset()
		return s
	}
	//decode input xml
	d := xml.NewDecoder(r)
	for {
		t, err := d.RawToken()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		switch v := t.(type) {
		case xml.StartElement:
			//two starts? new line!
			if _, ok := last.(xml.StartElement); ok {
				fmt.Fprintf(w, "\n")
			}
			//render
			currWidth := 0
			var n int
			n, _ = fmt.Fprintf(w, "%s<%s", spaces(indent), v.Name.Local)
			currWidth += n
			for _, attr := range v.Attr {
				//insert a newline if this line has surpassed max width
				if s.MaxWidth > 0 && currWidth > s.MaxWidth {
					fmt.Fprintf(w, "\n")
					n, _ = fmt.Fprint(w, spaces(indent))
					currWidth = n //reset current width
				}
				//write attribute key-val
				n, _ = fmt.Fprintf(w, ` %s="%s"`, attr.Name.Local, escape(attr.Value))
				currWidth += n
			}
			fmt.Fprintf(w, ">")
			//add one indent level
			indent++
			inner = true
			last = t
		case xml.EndElement:
			//remove one indent level
			indent--
			//render
			if _, ok := last.(xml.EndElement); ok {
				fmt.Fprintf(w, "%s", spaces(indent))
			}
			fmt.Fprintf(w, "</%s>", v.Name.Local)
			fmt.Fprintf(w, "\n")
			inner = false
			last = t
		case xml.CharData:
			//render
			if inner {
				fmt.Fprintf(w, "%s", strings.TrimSpace(string(v)))
			}
		case xml.Comment:
			//render
			fmt.Fprintf(w, "%s<!--%s-->\n", spaces(indent), v)
		case xml.ProcInst:
			//render
			fmt.Fprintf(w, "<?%s %s?>\n", v.Target, v.Inst)
		case xml.Directive:
			//render
			fmt.Fprintf(w, "<!%s>", v)
		}
	}
	return nil
}

func spaces(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n*2; i++ {
		sb.WriteRune(' ')
	}
	return sb.String()
}
