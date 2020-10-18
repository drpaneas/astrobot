package translate

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// TextToGreek translates given text in Greek and returns an error
func TextToGreek(text string) string {
	targetLanguage := "el" // https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		log.Fatalf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		log.Fatalf("Translate: %v", err)
	}
	if len(resp) == 0 {
		log.Fatalf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text
}
