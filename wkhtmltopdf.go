// This package is a cover around wkhtmltopdf library (https://wkhtmltopdf.org/).
// If you put string and input file to input, input string/[]bytes{} will be oin priority
package wkhtmltopdfgo

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

type Wkhtml struct {
	Options    WkhtmlOptions //Options
	Input      string        //Input filepath or URL
	OutputFile string        //Output file
	Stdin      []byte        //Input stdin
	Stdout     bytes.Buffer  //Output stdout
	Stderr     bytes.Buffer
	ExecBin    string //wkhtmltopdf binary file
}

//Create new Wkhtml object with filled options
func New() Wkhtml {
	return Wkhtml{
		Options: defaultOptions(),
		ExecBin: "wkhtmltopdf",
	}
}

//Set path to file with input html data.
func (w *Wkhtml) InputFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("File %s does not exists.", path)
		}
		return fmt.Errorf("Undefined error with %s.", path)
	}
	w.Input = path
	return nil
}

//Put string to stdin of wkhtmltipdf
func (w *Wkhtml) InputString(html string) {
	w.InputBytes([]byte(html))
}

//Set input url
func (w *Wkhtml) InputUrl(url string) error {
	//Check that URL exists
	resp, err := http.Head(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf(`URL "%s" returns status %d %s.'`, url, resp.StatusCode, resp.Status)
	}
	w.Input = url
	return nil
}

//Put bytes to stdin of wkhtmltopdf
func (w *Wkhtml) InputBytes(i []byte) {
	w.Stdin = i
}

//Return bytes of result pdf.
func (w *Wkhtml) Bytes() []byte {
	return w.Stdout.Bytes()
}

//Return slice with arguments for command line.
func (w *Wkhtml) Args() []string {
	args := []string{}
	r := reflect.Indirect(reflect.ValueOf(w.Options))
	for j := 0; j < r.NumField(); j++ {
		r.Field(j).Interface()
		f, ok := r.Field(j).Interface().(WkhtmlParam)
		if ok {
			if f.value != nil && f.name != "" {
				//If bool, then argument must be without value
				if v, ok := f.value.(bool); ok && v {
					args = append(args, fmt.Sprintf("%s", f.name))
				} else {
					args = append(args, f.name, fmt.Sprintf("%v", f.value))
				}
			}
		}
	}
	//If there is no input use stdin
	if w.Input == "" {
		args = append(args, "-")
	} else {
		args = append(args, w.Input)
	}
	//Print output to file or to stdout
	if w.OutputFile == "" {
		args = append(args, "-")
	} else {
		args = append(args, w.OutputFile)
	}
	return args
}

//Create pdf file
func (w *Wkhtml) Create() error {

	args := w.Args()

	cmd := exec.Command(w.ExecBin, args...)

	cmd.Stderr = &w.Stderr
	cmd.Stdin = bytes.NewReader(w.Stdin)
	cmd.Stdout = &w.Stdout

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Can't run wkhtmltopdf command: %s %s; Error: %s; %s.\n", w.ExecBin, strings.Join(args, " "), err.Error(), w.Stderr.String()) //@TODO: Make error cool!
	}

	return nil
}
