package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

//go:embed wh
//go:embed wo
//go:embed nums/0.png
var f embed.FS

func TaskFunc(image_file io.Reader) string {
	// Netwrok Structure: input: 784, hidden_layer: 200, output: 10
	// Load Input Image
	img, _ := png.Decode(image_file)
	pixels := make([]float64, 28*28)
	counter := 0
	for y := 0; y < 28; y++ {
		for x := 0; x < 28; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y
			pixels[counter] = (float64(255-(gray)) / 255.0 * 0.999) + 0.001
			counter += 1
		}
	}

	var hiddenWeights, outputWeights, hiddenInputs, hiddenOutputs, finalInputs, output mat.Dense

	// Load Model Parameters
	h, _ := f.ReadFile("wh")
	o, _ := f.ReadFile("wo")
	hiddenWeights.UnmarshalBinaryFrom(bytes.NewReader(h))
	outputWeights.UnmarshalBinaryFrom(bytes.NewReader(o))

	// Predict
	inputs := mat.NewDense(784, 1, pixels)
	hiddenInputs.Mul(&hiddenWeights, inputs)
	hiddenOutputs.Apply(sigmoid, &hiddenInputs)
	finalInputs.Mul(&outputWeights, &hiddenOutputs)
	output.Apply(sigmoid, &finalInputs)

	// Find Best Prediction
	best := 0
	highest := 0.0
	for i := 0; i < 10; i++ {
		if output.At(i, 0) > highest {
			best = i
			highest = output.At(i, 0)
		}
	}

	return strconv.Itoa(best)
}

func sigmoid(r, c int, z float64) float64 {
	return 1.0 / (1 + math.Exp(-1*z))
}

func main() {
	img, err := f.ReadFile("nums/0.png")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	_ = TaskFunc(bytes.NewReader(img))
	// resMsg := fmt.Sprintf("识别到 %s 的数字", ret)
}
