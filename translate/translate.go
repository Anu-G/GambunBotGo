package translate

import gt "github.com/bas24/googletranslatefree"

func TranslateJPtoEN(text string) string {
	result, _ := gt.Translate(text, "ja", "en")
	strRes := text + "\n\n" + result

	return strRes
}

func TranslateEntoJP(text string) string {
	result, _ := gt.Translate(text, "en", "ja")
	strRes := text + "\n\n" + result

	return strRes
}
