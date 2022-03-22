package wide

import (
	"testing"
)

func TestDrawPodLabels(t *testing.T) {
	cfg, err := DrawPodProperties("testdata/labels")
	if err != nil {
		t.Error(err)
	}
	if len(cfg) != 4 {
		t.Error("Keys are missing")
	}
}
