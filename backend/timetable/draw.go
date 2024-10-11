package timetable

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var dictTimetable = map[string]string{
	"1": "8.30 - 9.10",
	"2": "9.20 - 10.00",
	"3": "10.20 - 11.00",
	"4": "11.10 - 11.50",
	"5": "12.00 - 12.40",
	"6": "13.00 - 13.40",
	"7": "14.00 - 14.40",
	"8": "14.50 - 15.30",
	"9": "15.40 - 16.20",
}

var FontFace font.Face

func init() {
	fontBytes, err := os.ReadFile("backend/fonts/ofont.ru_GOST type B.ttf") // Замените на путь к вашему файлу шрифта
	if err != nil {
		log.Fatal("загрузка шрифта", err)
	}
	fontFace, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal("парсинг шрифта", err)
	}
	FontFace = truetype.NewFace(fontFace, &truetype.Options{Size: 20})
}

// addLabel adds a label to the image at the specified coordinates using the provided font, label, and color.
//
// Parameters:
// - img: the image to draw on
// - x, y: the coordinates where the label will be placed
// - fontFace: the font to use for the label
// - label: the text to be drawn
// - color: the color of the label
func addLabel(img *image.RGBA, x, y int, fontFace font.Face, label string, color color.Color) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: fontFace,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)},
	}
	d.DrawString(label)
}

// DrawTimetable draws a timetable image based on the given lessons, data, and teacher flag.
//
// Parameters:
// - lessons: a 2D slice of strings representing the lessons.
// - data: a string representing additional data to be displayed on the image.
// - teacher: a boolean flag indicating whether the timetable is for a teacher.
//
// Returns:
// - []byte: the encoded image in PNG format.
// - error: an error if encoding the image fails.
func DrawTimetable(lessons [][]string, data string, teacher bool, colors ...[]uint8) ([]byte, error) {
	color1 := color.RGBA{131, 236, 156, 255}
	color2 := color.RGBA{255, 255, 255, 255}
	if len(colors) == 2 {
		if len(colors[0]) == 3 {
			color1 = color.RGBA{colors[0][0], colors[0][1], colors[0][2], 255}
		}
		if len(colors[1]) == 3 {
			color2 = color.RGBA{colors[1][0], colors[1][1], colors[1][2], 255}
		}
	}
	
	var newLessons [][]string
	count := 0
	if teacher {
		for i := len(lessons) - 1; i >= 0; i-- {
			lesson := lessons[i]
			if len(lesson) > 0 && lesson[0] != "None" && lesson[0] != "" {
				newLessons = append([][]string{{lesson[0], dictTimetable[lesson[1]]}}, newLessons...)
				count = 1
			} else {
				if count > 0 {
					newLessons = append([][]string{{"Окно"}}, newLessons...)
				}
			}
		}
		lessons = newLessons
	}

	imageWidth, imageHeight := 500, (len(lessons)+1)*50

	for _, value := range lessons {
		value := value[0]
		if lenn(value) > 20 && lenn(value) < 30 && imageWidth < 600 {
			imageWidth = 600
		} else if lenn(value) > 30 && lenn(value) < 40 && imageWidth < 700 {
			imageWidth = 700
		} else if lenn(value) > 40 && lenn(value) < 65 && imageWidth < 800 {
			imageWidth = 800
		} else if lenn(value) > 65 {
			imageWidth = 900
			break
		}
	}

	// Создание изображения
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	// рисование прямоугольника белого цвета для полосы данных
	for x := 0; x < imageWidth; x++ {
		for yy := 0; yy < 50; yy++ {
			img.Set(x, yy, color.White)
		}
	}

	// Рисование расписания
	y := 50
	count = 0
	addLabel(img, 10, 30, FontFace, data, color.Black)
	for i, lesson := range lessons {
		var color_ color.RGBA
		if count < 2 {
			color_ = color1
			count++
		} else if count < 4 {
			color_ = color2
			count++
		} else {
			color_ = color1
			count = 1
		}
		for x := 0; x < imageWidth; x++ {
			for yy := y; yy < y+50; yy++ {
				img.Set(x, yy, color_)
			}
		}
		addLabel(img, 10, y+30, FontFace, fmt.Sprintf("%d. %s", i+1, lesson[0]), color.Black)
		if len(lesson) > 1 {
			if teacher {
				addLabel(img, imageWidth-150, y+30, FontFace, lesson[1], color.Black)
			} else {
				addLabel(img, imageWidth-100, y+30, FontFace, lesson[1], color.Black)
			}
		}
		y += 50
	}

	// Кодирование изображения в PNG и запись в буфер
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func lenn(str string) int {
	var count int
	for range str {
		count++
	}
	return count
}