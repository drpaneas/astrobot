package translate

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"github.com/drpaneas/astrobot/pkg/earthsky"
	"github.com/drpaneas/astrobot/pkg/space"
	"github.com/drpaneas/astrobot/pkg/universetoday"
	"golang.org/x/text/language"
)

// TextToGreek translates given text in Greek and returns an error
func TextToGreek(text string) (string, error) {
	targetLanguage := "el" // https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

// NewsDBgo for UniverseToday
func NewsDBgo() {
	for i, db := range universetoday.NewsDBUniverseToday {
		universetoday.NewsDBUniverseToday[i].GreekTitle, _ = TextToGreek(db.Title)
		universetoday.NewsDBUniverseToday[i].GreekDesc, _ = TextToGreek(db.Description)
	}
}

// NewsEarthSkygo for UniverseToday
func NewsEarthSkygo() {
	for i, db := range earthsky.NewsDBEarthSky {
		earthsky.NewsDBEarthSky[i].GreekTitle, _ = TextToGreek(db.Title)
		earthsky.NewsDBEarthSky[i].GreekDesc, _ = TextToGreek(db.Description)
	}
}

// NewsSpacego for space.com
func NewsSpacego() {
	for i, db := range space.NewsDBSpace {
		space.NewsDBSpace[i].GreekTitle, _ = TextToGreek(db.Title)
		space.NewsDBSpace[i].GreekDesc, _ = TextToGreek(db.Description)
	}
}
