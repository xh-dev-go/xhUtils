package tgUtils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PreProcessMarkup(msg string) string{
	var replace = msg
	replace = strings.ReplaceAll(replace, "-","\\-")
	replace = strings.ReplaceAll(replace, "=","\\=")
	replace = strings.ReplaceAll(replace, "[","\\[")
	replace = strings.ReplaceAll(replace, "]","\\]")
	replace = strings.ReplaceAll(replace, ".","\\.")
	return replace
}

func Send(to string, msg string, token string) error {
	if to == "" {
		return errors.New("No chat_id provided!")
	} else if msg == "" {
		return errors.New("No message provided!")
	}
	//println("sending...")
	//println(msg)
	data := url.Values{
		"parse_mode": {"MarkdownV2"},
		"chat_id":    {to},
		"text":       {msg},
	}

	if resp, err := http.PostForm(fmt.Sprintf(`https://api.telegram.org/bot%s/sendMessage`, token), data); err != nil {
		panic(err)
	} else if byte, err := ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	} else {
		println(string(byte))
	}
	return nil
}

