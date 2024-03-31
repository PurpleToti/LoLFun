package data_utils

func GetFormattedKeyValue(key string, value string, quote_type string) string {
	return quote_type + key + quote_type + ":" + quote_type + value + quote_type
}

func GetFormattedList(strings []string, quote_type string) string {
	fstring := "["
	for i := 0; i < len(strings); i++ {
		fstring += quote_type + strings[i] + quote_type
		if i < len(strings)-1 {
			fstring += ","
		}
	}
	fstring += "]"
	return fstring
}
