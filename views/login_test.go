package views

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHeader(t *testing.T) {
	// Pipe the rendered template into goquery.
	r, w := io.Pipe()
	go func() {
		_ = Login().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	// Expect the component to be present.
	if doc.Find(`form`).Length() == 0 {
		t.Error("expected form to be rendered, but it wasn't")
	}

}
