package lolfun_cookies

import (
	"LoLFun/gopacs/data_utils"
	"errors"
	"net/http"
)

func StringifyCookie(cookie *http.Cookie, quote_type string) string {
	repr := "Cookie{"
	repr += data_utils.GetFormattedKeyValue("Name", cookie.Name, quote_type)
	repr += data_utils.GetFormattedKeyValue("Value", cookie.Value, quote_type)
	repr += "}"
	return repr
}

func NewUserCookieValueArray() [Number_of_caracters]rune {
	var array [Number_of_caracters]rune
	for i := 0; i < int(Number_of_caracters); i++ {
		array[i] = 0
	}
	return array
}

func NextCookie() error {
	for i := 0; i < int(Number_of_caracters); i++ {
		if Ascii_repr_caracs[i] < rune(Ascii_limit) {
			Ascii_repr_caracs[i]++
			return nil
		}
	}
	return errors.New("limit of cookies reached")
}
