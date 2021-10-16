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
	for i, hex := range theme {
		canvas.Rect(0, i*layerHeight, width, layerHeight, fmt.Sprintf("fill: #%s", hex))
	}
	canvas.End()
}
