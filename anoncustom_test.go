package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIO(t *testing.T) {
	voc.reset()
	name := "fio"
	regexp := "([A-Z][a-zA-Z\\-]{1,})(\\s+[A-Z][a-zA-Z\\-]{1,})(\\s+[A-Z][a-zA-Z\\-]{1,})?"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace FIO", func(t *testing.T) {
		res, err := f("attribute=John Marshall Doe asdasd=Jehny Marshall")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=NTest0 NTest1 NTest2 asdasd=NTest3 NTest1", res, "should return changed input")

		res, err = f("attribute=Alice Smith")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=NTest4 NTest5", res, "should return changed input")
	})
}

func TestFIOInitials(t *testing.T) {
	voc.reset()
	name := "fio_initials"
	regexp := "([A-Z][a-zA-Z\\-]{1,})(\\s+[A-Z]\\.)(\\s*[A-Z]\\.)?"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace FIO", func(t *testing.T) {
		res, err := f("attribute=John D. D.")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=NTest0 X.X.", res, "should return changed input")

		res, err = f("attribute=Peter G.G.")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=NTest1 X.X.", res, "should return changed input")
	})
}

func TestAnyName(t *testing.T) {
	voc.reset()
	name := "name"
	regexp := "([A-Z][a-zA-Z\\-]{1,})"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace Name", func(t *testing.T) {
		res, err := f("attribute=John Doe")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=NTest0 NTest1", res, "should return changed input")
	})
}

func TestIP(t *testing.T) {
	voc.reset()
	name := "ip"
	regexp := "\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace IP", func(t *testing.T) {
		res, err := f("attribute=127.0.0.1")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=***.***.0.1", res, "should return changed input")
	})
}

func TestPhone(t *testing.T) {
	voc.reset()
	name := "phone"
	regexp := "\\b[1,3]9[0-9]{9,11}\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace Phone", func(t *testing.T) {
		res, err := f("attribute=19251234567")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=1925*******", res, "should return changed input")
	})
}

func TestOTP(t *testing.T) {
	voc.reset()
	name := "numbers"
	regexp := "\\b[0-9]{4}\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace OTP", func(t *testing.T) {
		res, err := f("otp=1234")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "otp=****", res, "should return changed input")
	})
}

func TestBirthdate(t *testing.T) {
	voc.reset()
	name := "birthdate"
	regexp := "\\b[0-9]{4}\\-[0-9]{1,2}\\-[0-9]{1,2}\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace birthdate", func(t *testing.T) {
		res, err := f("bd=1970-12-12")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "bd=****-**-**", res, "should return changed input")
	})
}

func TestBase64(t *testing.T) {
	name := "clear"
	regexp := "\\b[A-Za-z0-9+/]{1000,}(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("clear blob", func(t *testing.T) {
		res, err := f("photo=iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAMAAAD04JH5AAAAA3NCSVQICAjb4U/gAAAACXBIWXMAABuvAAAbrwFeGpEcAAAAGXRFWHRTb2Z0d2FyZQB3d3cuaW5rc2NhcGUub3Jnm+48GgAAArtQTFRF////AAAAHR0bAAAAHR0bAAAAAAAAHR0bHR0bMzMzHR0bKysrHR0bJCQkHR0bICAgHR0bFBQUHR0bJCQSHR0bIiIiICAgHh4eHBwcHR0bGxsbGhoaGBgYHR0bIyMXHR0bISEWICAgHR0bHR0dHBwcHR0bHR0bHR0bICAYHh4eHBwcGxsbICAaHx8ZHh4YHh4eHR0bHR0dHR0bGxsbHR0bGxsbHx8aHx8aHR0bHh4ZHR0dHR0bHR0dHBwcHR0bHBwcHR0bGxsbHR0bGxsbHR0bHR0bHR0bHBwcHR0bGxsbHx8bHh4bHh4aHR0bHh4aHR0bHR0dHR0bHx8bHh4bHh4aHR0aHR0bHR0aHR0aHBwcHBwcHh4bHR0bHBwcHBwcHR0bHBwcHh4bHR0aHBwaHBwcHR0bHh4bHR0bHR0bHR0aHR0aHR0bHR0bHh4cHR0bHR0bHR0bHR0aHh4cHh4cHR0bHh4bHR0bHR0bHR0bHR0bHBwaHR0bHh4cHR0bHR0bHR0bHR0bHR0bHBwaHR0bHh4cHR0cHR0bHR0bHR0bHBwbHBwaHR0bHR0cHR0bHR0bHR0bHBwbHh4aHR0cHR0bHR0bHR0bHBwbHR0bHR0bHR0bHBwbHR0bHR0aHR0cHR0bHR0bHBwbHR0bHBwbHR0bHR0bHh4bHR0bHR0bHR0bHR0bHR0bHR0bHR0bHh4bHR0bHR0bHR0bHR0bHR0cHR0bHR0bHR0bHR0bHR0bHR0aHBwbHR0bHh4bHR0bHR0bHR0aHR0bHh4bHR0bHR0bHR0bHR0bHR0bHR0bHR0cHR0bHBwbHR0bHR0bHR0bHR0bHR0bHR0bHR0aHR0bHBwbHR0bHR0bHR0bHR0bHR0bHR0bHR0bHR0cHh4bHR0bHR0bHR0bHR0bHR0bHR0bHR0bHR0bHR0bHR0bHR0bHR0bHR0bFI+BhQAAAOh0Uk5TAAEBAgIDBAQFBQYGBwcICAoNDQ4PDxAREhITFBUWFhcXGBkaGx0fICAiJCYoKSorLCwuLy8wMTIzMzQ1NTY2Nzc4ODk8Pj9AQEFC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")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "photo=", res, "should return changed input")
	})
}

func TestReplaceLogin(t *testing.T) {
	name := "hash"
	regexp := "\\basd[0-9]{1,6}\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("clear blob", func(t *testing.T) {
		res, err := f("cn=asd12345")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "cn=293662475dbe4235a50a011325b9e878194dfcca", res, "should return changed input")
	})
}

func TestReplaceCard(t *testing.T) {
	name := "card"
	regexp := "\\b\\d{16}\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("clear blob", func(t *testing.T) {
		res, err := f("card: 1234123412341234")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "card: 1234********1234", res, "should return changed input")
	})
}

func TestReplaceEmail(t *testing.T) {
	name := "email"
	regexp := "\\b[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("clear blob", func(t *testing.T) {
		res, err := f("email: test@mail.com")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "email: test@domain.com", res, "should return changed input")
	})
}

func TestComplex(t *testing.T) {
	voc.reset()
	name := "phone"
	regexp := "\\b[7,8]9[0-9]{9,11}\\b"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace Name and Phone", func(t *testing.T) {
		res, err := f("attribute=89251234567")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "attribute=8925*******", res, "should return changed input")
	})
}

func TestHostname(t *testing.T) {
	vocHosts.reset()
	name := "hostname"
	regexp := "([a-zA-Z0-9-.]*google\\.com)|([a-zA-Z0-9-.]*go\\.gl)"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace complex regexp", func(t *testing.T) {
		res, err := f("server=maps.google.com server2=test.google.com")
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, "server=NTesthost0 server2=NTesthost1", res, "should return changed input")
	})
}

func TestReplaceMiddleGroup(t *testing.T) {
	name := "replaceMiddleGroup"
	regexp := "(\"value\"\\:\\s*\")([\\w\\W]*)(\")"
	f, _ := custom([]CustomConfig{CustomConfig{Name: name, Regexp: regexp}})
	t.Run("Replace middle grop regexp", func(t *testing.T) {
		res, err := f(`"value": "login"`)
		assert.NoError(t, err, "should return no error")
		assert.Equal(t, `"value": "********"`, res, "should return changed input")
	})
}
