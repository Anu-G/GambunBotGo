package translate

import gt "github.com/bas24/googletranslatefree"

func TranslateJPtoEN(text string) string {
	result, _ := gt.Translate(text, "ja", "en")
	strRes := text + "\n\n" + result

	return strRes
}

func TranslateJPtoID(text string) string {
	result, _ := gt.Translate(text, "ja", "id")
	strRes := text + "\n\n" + result

	return strRes
}
