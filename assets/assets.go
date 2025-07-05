package assets

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed fonts/FiraSans-Regular.ttf
var firaSansRegular []byte

// SharedTitleFace is the font face used for titles.
var SharedTitleFace *text.GoTextFace

// FontsLoaded indicates whether the fonts were loaded successfully.
var FontsLoaded bool

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(firaSansRegular))
	if err != nil {
		FontsLoaded = false
		log.Printf("Error loading fonts: %v", err)
		return
	}

	SharedTitleFace = &text.GoTextFace{Source: s, Size: 40}
	FontsLoaded = true
}
