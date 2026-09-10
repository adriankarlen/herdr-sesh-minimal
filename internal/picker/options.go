package picker

import (
	"context"
	"io"
)

type Options struct {
	Context                        context.Context
	Output                         io.Writer
	Prompt                         string
	Placeholder                    string
	ShowIcons                      bool
	SeparatorAware                 bool
	DisableHomePrioritization      bool
	HidePath                       bool
	HidePreview                    bool
	DefaultPreviewCommand          string
	FZFCommand                     string
	GumCommand                     string
}
