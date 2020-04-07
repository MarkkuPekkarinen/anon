package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"regexp"
	"strings"
)

//CustomConfig replaces regexp by defied algorithm
type CustomConfig struct {
	Name   string
	Regexp string
}

type vocabulary struct {
	vocabularyMap map[string]string
	prefix        string
}

func (v *vocabulary) getVocabularyVal(s string) string {
	if s == "" {
		return s
	}
	if v.vocabularyMap[s] == "" {
		v.vocabularyMap[s] = v.prefix + fmt.Sprint(len(v.vocabularyMap))
	}
	return v.vocabularyMap[s]
}

func (v *vocabulary) reset() {
	for k := range v.vocabularyMap {
		delete(v.vocabularyMap, k)
	}
}

var voc = vocabulary{vocabularyMap: make(map[string]string), prefix: "NTest"}
var vocHosts = vocabulary{vocabularyMap: make(map[string]string), prefix: "NTesthost"}

func custom(cc []CustomConfig) (Anonymisation, error) {
	return func(s string) (string, error) {
		for _, c := range cc {
			r := regexp.MustCompile(c.Regexp)
			var f func(s string) string
			switch c.Name {
			case "fio":
				f = fio
			case "fio_initials":
				f = fioInitials
			case "name":
				f = name
			case "ip":
				f = ip
			case "phone":
				f = phone
			case "numbers":
				f = numbers
			case "birthdate":
				f = birthdate
			case "clear":
				f = clear
			case "hash":
				f = calcHash
			case "card":
				f = card
			case "email":
				f = email
			case "hostname":
				f = hostname
			case "replaceMiddleGroup":
				f = func(a string) string {
					a = r.ReplaceAllString(a, "$1********$3")
					return a
				}
			default:
				f = func(a string) string {
					return a
				}
			}
			matches := r.FindAllString(s, -1)
			for _, sub := range matches {
				sub2 := r.ReplaceAllStringFunc(sub, f)
				s = strings.ReplaceAll(s, sub, sub2)
			}
			println(c.Name, s)
		}
		return s, nil
	}, nil
}

func fio(a string) string {
	fioSlice := strings.SplitN(a, " ", 3)
	fio := voc.getVocabularyVal(fioSlice[0]) + " " + voc.getVocabularyVal(fioSlice[1])
	if len(fioSlice) >= 3 {
		fio += " " + voc.getVocabularyVal(fioSlice[2])
	}
	return fio

}

func fioInitials(a string) string {
	fioSlice := strings.SplitN(a, " ", 2)
	fio := voc.getVocabularyVal(fioSlice[0]) + " " + "X.X."
	return fio
}

func name(a string) string {
	name := voc.getVocabularyVal(a)
	return name
}

func ip(a string) string {
	ipSlice := strings.Split(a, ".")
	return "***.***." + ipSlice[2] + "." + ipSlice[3]
}

func phone(a string) string {
	return a[0:4] + strings.Repeat("*", len(a)-4)
}

func numbers(a string) string {
	return strings.Repeat("*", len(a))
}

func birthdate(a string) string {
	return "****-**-**"
}

func clear(a string) string {
	return ""
}

func calcHash(a string) string {
	h := sha1.New()
	io.WriteString(h, a)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func card(a string) string {
	return a[0:4] + "********" + a[12:16]
}

func email(a string) string {
	r := regexp.MustCompile("\\b([a-zA-Z0-9_.+-]+)@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+\\b")
	return r.ReplaceAllString(a, "$1@domain.com")
}

func hostname(a string) string {
	return vocHosts.getVocabularyVal(a)
}
