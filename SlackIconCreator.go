package main

import (
	"os"
	"io/ioutil"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"flag"
	"strings"
	"log"
)

func main() {
	ttfFontName := flag.String("ttf", "", "利用するTTFファイル")
	fileName := flag.String("out", "", "出力先のファイル名")
	messages := flag.String("mes", "", "出力したいテキスト。'|'で改行（複数行化/エスケープは未対応）になる")
	//bkColorWork := flag.String("bkc", "", "背景色。'0xFFFFFFFF'の書式でRGBAの順に記載する")
	//ftColorWork := flag.String("ftc", "", "文字色。'0xFFFFFFFF'の書式でRGBAの順に記載する")

	flag.Parse()

	if *ttfFontName == "" || *fileName == "" || *messages == "" {
		log.Panic("'-h'を参考にパラメータをすべて指定して実行してください")
	}

	log.Printf("ttf: %s", *ttfFontName)
	log.Printf("out: %s", *fileName)
	log.Printf("mes: %s", *messages)

	m := strings.Split(*messages, "|")

	CreateSlackIcon(*fileName, m, *ttfFontName)
}

func CreateSlackIcon(fileName string, messages []string, ttfFontName string) {
	canvas := GetStartCanvas()

	DrawStringCenter(canvas, messages, ttfFontName)

	f, err := os.Create(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, canvas); err != nil {
		log.Panic(err)
	}
}

func GetStartCanvas() *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, 128, 128))
}

func DrawStringCenter(canvas *image.RGBA, messages []string, ttfFontName string) {
	for size := 150; size > 0; size-- {
		fontFace := GetFontFace(ttfFontName, size)
		m := (*fontFace).Metrics()

		messageParams := make([]MessageParam, len(messages))

		height := m.Height.Ceil() * len(messages)
		width := 0
		for i, v := range messages {
//			width = math.Max(font.MeasureString(*fontFace, v).Ceil(), width)
			lineWidth := font.MeasureString(*fontFace, v).Ceil()
			if lineWidth > width {
				width = lineWidth
			}

			messageParams[i] = MessageParam {
				Message:	v,
				Posx:		(canvas.Rect.Max.X / 2) - (lineWidth / 2),
			}
		}

		posx := (canvas.Rect.Max.X / 2) - (width / 2)
		posy := (canvas.Rect.Max.Y / 2) - (height / 2)

		if posx > 0 && posy > 0 && posx < canvas.Rect.Max.X && posy < canvas.Rect.Max.Y {
			for i, v := range messageParams {
				d := &font.Drawer{
					Dst:  canvas,
					Src:  image.NewUniform(color.Black),
					Face: *fontFace,
					Dot:
						fixed.Point26_6 {
							X: fixed.Int26_6(v.Posx * 64),
							// 正確にはフォントのAscent分、下に下げたいのでHeightを足して下に下げている
							Y: fixed.Int26_6((m.Height.Ceil() + posy + (i * m.Height.Ceil())) * 64),
						},
				}

				d.DrawString(v.Message)
			}
			break
		}
	}
}

func GetFontFace(fontFileName string, size int) *font.Face {
	fontBin, err := ioutil.ReadFile(fontFileName)
	if err != nil {
		log.Panic(err)
	}

	tt, err := truetype.Parse(fontBin)
	if err != nil {
		log.Panic(err)
	}

	face := truetype.NewFace(tt, &truetype.Options{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingNone,
	})

	return &face
}

type MessageParam struct {
	Message string
	Posx int
}
