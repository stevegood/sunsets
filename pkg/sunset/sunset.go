package sunset

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
)

func Random(writer io.Writer, width, height int) {
	// get a random theme, sorted light to dark
	theme := LightToDark(RandomTheme())
	// produce a layer for each color
	numLayers := theme.Len()
	layerHeight := height / numLayers

	canvas := svg.New(writer)
	canvas.Start(width, height)
	canvas.LinearGradient("sky", 0, 0, 0, 100, []svg.Offcolor{
		svg.Offcolor{Offset: 5, Opacity: 100, Color: fmt.Sprintf("#%s", theme[0])},
		svg.Offcolor{Offset: 95, Opacity: 100, Color: fmt.Sprintf("#%s", theme[theme.Len()-1])},
	})
	canvas.Rect(0, 0, width, height, "fill:url(#sky);")
	for i, hex := range theme {
		canvas.Rect(0, i*layerHeight, width/2, layerHeight, fmt.Sprintf("fill: #%s;", hex))
	}
	canvas.End()
}
