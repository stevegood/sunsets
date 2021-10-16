package sunset

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

// Theme is a sortable slice of colors (hex)
type Theme []string

func (t Theme) Len() int {
	return len(t)
}

func (t Theme) Less(i, j int) bool {
	a, _ := colorful.Hex(fmt.Sprintf("#%s", t[i]))

	b, _ := colorful.Hex(fmt.Sprintf("#%s", t[j]))

	aRGB := a.R + a.G + a.B
	bRGB := b.R + b.G + b.B

	return aRGB < bRGB
}

func (t Theme) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Theme) Colors() []colorful.Color {
	var colors []colorful.Color

	for _, hex := range t {
		color, _ := colorful.Hex(fmt.Sprintf("#%s", hex))
		colors = append(colors, color)
	}

	return colors
}

func LightToDark(t Theme) Theme {
	sort.Sort(sort.Reverse(t))
	return t
}

func DarkToLight(t Theme) Theme {
	sort.Sort(t)
	return t
}

func Themes() []Theme {
	return []Theme{
		Theme{"03071e", "370617", "6a040f", "9d0208", "d00000", "dc2f02", "e85d04", "f48c06", "faa307", "ffba08"},
		Theme{"f72585", "b5179e", "7209b7", "560bad", "480ca8", "3a0ca3", "3f37c9", "4361ee", "4895ef", "4cc9f0"},
		Theme{"f94144", "f3722c", "f8961e", "f9844a", "f9c74f", "90be6d", "43aa8b", "4d908e", "577590", "277da1"},
		Theme{"ff6d00", "ff7900", "ff8500", "ff9100", "ff9e00", "240046", "3c096c", "5a189a", "7b2cbf", "9d4edd"},
		Theme{"ff4800", "ff5400", "ff6000", "ff6d00", "ff7900", "ff8500", "ff9100", "ff9e00", "ffaa00", "ffb600"},
		Theme{"ea698b", "d55d92", "c05299", "ac46a1", "973aa8", "822faf", "6d23b6", "6411ad", "571089", "47126b"},
	}
}

func RandomTheme() Theme {
	themes := Themes()
	numThemes := len(themes)
	rand.Seed(time.Now().Unix())
	i := rand.Intn(numThemes)
	return themes[i]
}
