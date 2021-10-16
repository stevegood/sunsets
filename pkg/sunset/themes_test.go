package sunset_test

import (
	"reflect"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/stevegood/sunsets/pkg/sunset"
)

func TestThemes(t *testing.T) {
	themes := sunset.Themes()
	if len(themes) <= 0 {
		t.Log("no themes returned")
		t.Fail()
	}
}

func TestRandomTheme(t *testing.T) {
	randomTheme := sunset.RandomTheme()
	if randomTheme.Len() <= 0 {
		t.Log("theme contains no colors")
		t.Fail()
	}
}

func TestLightToDark(t *testing.T) {
	theme := sunset.Theme{"CCCCCC", "000000", "FFFFFF"}
	expected := sunset.Theme{"FFFFFF", "CCCCCC", "000000"}

	sorted := sunset.LightToDark(theme)

	equal := reflect.DeepEqual(sorted, expected)
	if !equal {
		t.Logf("Themes are not equal.\nExpected\n\t%v\nBut got\n\t%v", expected, sorted)
		t.Fail()
	}
}

func TestDarkToLight(t *testing.T) {
	theme := sunset.Theme{"CCCCCC", "000000", "FFFFFF"}
	expected := sunset.Theme{"000000", "CCCCCC", "FFFFFF"}

	sorted := sunset.DarkToLight(theme)

	equal := reflect.DeepEqual(sorted, expected)
	if !equal {
		t.Logf("Themes are not equal.\nExpected\n\t%v\nBut got\n\t%v", expected, sorted)
		t.Fail()
	}
}

func TestColors(t *testing.T) {
	theme := sunset.Theme{"000000", "CCCCCC", "FFFFFF"}
	expected := []colorful.Color{
		{R: 0, G: 0, B: 0},
		{R: 0.8, G: 0.8, B: 0.8},
		{R: 1, G: 1, B: 1},
	}
	colors := theme.Colors()

	equal := reflect.DeepEqual(colors, expected)
	if !equal {
		t.Logf("Themes are not equal.\nExpected\n\t%v\nBut got\n\t%v", expected, colors)
		t.Fail()
	}
}
