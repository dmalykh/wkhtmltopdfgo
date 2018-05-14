package wkhtmltopdfgo

import "testing"

func TestNew(t *testing.T) {
	wk := New()
	if wk.ExecBin != "wkhtmltopdf" {
		t.Errorf("New Wkhtml initialized empty.")
	}
}


func TestWkhtml_InputFile(t *testing.T) {
	wk := New()
	t.Run("Set bad path", func(t *testing.T) {
		if err := wk.InputFile("/thc/fjy/gvkuh.bix"); err == nil {
			t.Errorf("Path does not exists, but InputFile returns error.")
		}
	})
	t.Run("Set valid path", func(t *testing.T) {
		path := "options.go"
		if err := wk.InputFile(path); err != nil {
			t.Errorf("%s exists, but InputFile returns error.", path)
		}
		if wk.Input != path {
			t.Errorf("Inout did not set.")
		}
	})
}