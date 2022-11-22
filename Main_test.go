package main

import "testing"

func TestConvert(t *testing.T) {
	s := [][]string{{"EUR 1.203,09", "EUR:1203.09"},
		{"$30.65", "$:30.65"},
		{"$30.65 ($13.74 / Pound)", "$:30.65"},
		{" $1,000,010.46", "$:1000010.46"},
		{"+&nbsp;EUR 18,00&nbsp;di&nbsp;spedizione", "EUR:18.00"},
		{"MP3", ""},
		{"P3", ""},
		{"3", ""},
		{"CDN$ 59.71", "CDN$:59.71"},
		{"CDN$&nbsp;59.71", "CDN$:59.71"},
		{"CDN$59.71", "CDN$:59.71"},
		{"AUD 59.71", "AUD:59.71"},
		{"CAD 59.71", "CAD:59.71"},
		{"GBP 59.71", "GBP:59.71"},
		{"39,99 €", "€:39.99"}}

	for _, v := range s {
		if rez := convert(v[0]); rez != v[1] {
			t.Errorf("Expected '%s', but got '%s'", v[1], rez)
		}
	}
}
