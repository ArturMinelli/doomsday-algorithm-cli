package speech

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

func Speak(text string) {
	speech := htgotts.Speech{
		Folder:   "/tmp",
		Language: voices.English,
		Handler:  &handlers.MPlayer{},
	}

	speech.Speak(text)
}
