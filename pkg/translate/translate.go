package translate

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
	"github.com/drpaneas/astrobot/pkg/earthsky"
	"github.com/drpaneas/astrobot/pkg/space"
	"github.com/drpaneas/astrobot/pkg/universetoday"
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

// NewsDBgo for UniverseToday
func NewsDBgo() {
	for i, db := range universetoday.NewsDBUniverseToday {
		log.Println(db.Title)
		universetoday.NewsDBUniverseToday[i].GreekTitle = TextToGreek(db.Title)
		log.Println(db.Description)
		universetoday.NewsDBUniverseToday[i].GreekDesc = TextToGreek(db.Description)
	}
}

// NewsEarthSkygo for UniverseToday
func NewsEarthSkygo() {
	for i, db := range earthsky.NewsDBEarthSky {
		earthsky.NewsDBEarthSky[i].GreekTitle = TextToGreek(db.Title)
		earthsky.NewsDBEarthSky[i].GreekDesc = TextToGreek(db.Description)
	}
}

// NewsSpacego for space.com
func NewsSpacego() {
	for i, db := range space.NewsDBSpace {
		space.NewsDBSpace[i].GreekTitle = TextToGreek(db.Title)
		space.NewsDBSpace[i].GreekDesc = TextToGreek(db.Description)
	}
}
