/**
 * TODO:
 *
 * Make pixel parsing concurrent
 * Make Histogram a type
 * Draw histogram or output in any way
 */

package main;

import (
  "fmt"
  "os"
  "image"
  "log"
  _ "image/jpeg"
)

const IMAGE_FILE = "download.jpg"

func readFile(file string) *os.File {
  reader, err := os.Open(file)
  if err != nil {
    log.Fatal(err)
  }

  return reader
}

func readImage(reader *os.File) image.Image {
  info, _ := reader.Stat()
  fmt.Println("Image file size:", info.Size(), "bytes")

  img, _, err := image.Decode(reader)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Image size:", img.Bounds().Max.X, "x", img.Bounds().Max.Y)

  return img
}

func luminosity(r, g, b uint32) uint32 {
  return uint32(0.2126 * float32(r) + 0.7152 * float32(g) + 0.0722 * float32(b))
}

func Histogram(img image.Image, channels int) []float32 {
  histo := make([]float32, channels)


  max := uint32(1 << 16);
  for x := 0; x < img.Bounds().Max.X; x++ {
    for y := 0; y < img.Bounds().Max.Y; y++ {
      r, g, b, _ := img.At(x, y).RGBA()
      brightness := luminosity(r, g, b);

      value := brightness * uint32(channels) / max

      fmt.Println(x, y, value);

      histo[value]++;
    }
  }

  return histo
}

func main() {
  fmt.Println("Hi there")

  fmt.Println("Opening image")

  reader := readFile(IMAGE_FILE)
  defer reader.Close()

  img := readImage(reader)

  // r, g, b, a := img.At(0, 0).RGBA()
  // fmt.Println("RGBA at 0:0", r >> 8, g >> 8, b >> 8, a >> 8)

  fmt.Println(Histogram(img, 16))
}
