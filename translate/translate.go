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

func TranslateENtoJP(text string) string {
	result, _ := gt.Translate(text, "en", "ja")
	strRes := text + "\n\n" + result

	return strRes
}

func TranslateIDtoJP(text string) string {
	result, _ := gt.Translate(text, "id", "ja")
	strRes := text + "\n\n" + result

	return strRes
}
