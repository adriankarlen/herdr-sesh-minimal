package picker_test

import (
	"testing"

	"github.com/adriankarlen/herdr-sesh-minimal/internal/model"
	"github.com/adriankarlen/herdr-sesh-minimal/internal/picker"
	"github.com/stretchr/testify/assert"
)

func TestFormatDisplay(t *testing.T) {
	s := model.Session{
		Source: "config",
		Name:   "downloads",
		Path:   "/tmp/downloads",
	}
	out := picker.FormatDisplay(s)
	assert.Contains(t, out, "/tmp/downloads")
	assert.Contains(t, out, "")
}
