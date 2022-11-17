package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/text/language"
)

func TestLanguageTags(t *testing.T) {
	tags := []language.Tag{
		language.Chinese,
		language.SimplifiedChinese,
		language.TraditionalChinese,
		language.English,
		language.AmericanEnglish,
		language.BritishEnglish,
	}
	for _, tag := range tags {
		s := tag.String()
		base, _ := tag.Base()
		t.Logf("language tag is %s, base is %s", s, base)
	}
}

func TestI18n(t *testing.T) {
	t.Run("zh-CN", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(AcceptLanguage, "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
		w := httptest.NewRecorder()
		newServer().ServeHTTP(w, req)
		t.Logf(w.Body.String())
	})

	t.Run("zh-CN", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(AcceptLanguage, language.SimplifiedChinese.String())
		w := httptest.NewRecorder()
		newServer().ServeHTTP(w, req)
		t.Logf(w.Body.String())
	})

	t.Run("en-US", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(AcceptLanguage, language.English.String())
		w := httptest.NewRecorder()
		newServer().ServeHTTP(w, req)
		t.Logf(w.Body.String())
	})
}
