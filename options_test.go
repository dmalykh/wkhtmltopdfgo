package wkhtmltopdfgo

import (
	"testing"
)

func TestWkhtmlParam_Set(t *testing.T) {

	var param WkhtmlParam
	param.Set(100)
	if param.value.(int) != 100 {
		t.Errorf("Value must be 100, got %v.", param.value)
	}
}


func Test_defaultOptions(t *testing.T) {
	opt := defaultOptions()
	opt2 := defaultOptions()

	if opt != opt2 {
		t.Errorf("Defaults sets different params.")
	}

	if opt.PageSize.name != "--page-size" {
		t.Errorf("Defaults did not set.")
	}
}