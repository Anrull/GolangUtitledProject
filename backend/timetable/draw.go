package timetable

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
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

func DrawTimetable(lessons [][]string, data string, teacher bool, num int) {
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

	imageWidth, imageHeight := 400, (len(lessons)+1)*50

	// Создание изображения
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	// Установка шрифта
	fontBytes, err := os.ReadFile("backend/fonts/ofont.ru_GOST type B.ttf") // Замените на путь к вашему файлу шрифта
	if err != nil {
		log.Fatal(err)
	}
	fontFace, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	f := truetype.NewFace(fontFace, &truetype.Options{Size: 20})

	// рисование прямоугольника белого цвета для полосы данных
	for x := 0; x < imageWidth; x++ {
		for yy := 0; yy < 50; yy++ {
			img.Set(x, yy, color.White)
		}
	}

	// Рисование расписания
	y := 50
	count = 0
	addLabel(img, 10, 30, f, data, color.Black)
	for i, lesson := range lessons {
		var color_ color.RGBA
		if count < 2 {
			color_ = color.RGBA{131, 236, 156, 255}
			count++
		} else if count < 4 {
			color_ = color.RGBA{255, 255, 255, 255}
			count++
		} else {
			color_ = color.RGBA{131, 236, 156, 255}
			count = 1
		}
		for x := 0; x < imageWidth; x++ {
			for yy := y; yy < y+50; yy++ {
				img.Set(x, yy, color_)
			}
		}
		// было y + 10
		addLabel(img, 10, y+30, f, fmt.Sprintf("%d. %s", i+1, lesson[0]), color.Black)
		if len(lesson) > 1 {
			if teacher {
				// было y + 10
				addLabel(img, imageWidth-150, y+30, f, lesson[1], color.Black)
			} else {
				// было y + 10
				addLabel(img, imageWidth-100, y+30, f, lesson[1], color.Black)
			}
		}
		y += 50
	}

	// Сохранение изображения
	var filename string
	if num == 0 {
		filename = "data/temp/images/schedule.png"
	} else {
		filename = fmt.Sprintf("data/temp/images/schedule%d.png", num)
	}
	outfile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}

func addLabel(img *image.RGBA, x, y int, fontFace font.Face, label string, color color.Color) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: fontFace,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)},
	}
	d.DrawString(label)
}

func DrawTimetableTest(lessons [][]string, data string, teacher bool) ([]byte, error) {
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

	imageWidth, imageHeight := 400, (len(lessons)+1)*50

	// Создание изображения
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	// Установка шрифта
	fontBytes, err := os.ReadFile("backend/fonts/ofont.ru_GOST type B.ttf") // Замените на путь к вашему файлу шрифта
	if err != nil {
		return nil, err
	}
	fontFace, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	f, err := opentype.NewFace(fontFace, &opentype.FaceOptions{Size: 20, DPI: 72, Hinting: font.HintingFull})
	if err != nil {
		return nil, err
	}

	// рисование прямоугольника белого цвета для полосы данных
	for x := 0; x < imageWidth; x++ {
		for yy := 0; yy < 50; yy++ {
			img.Set(x, yy, color.White)
		}
	}

	// Рисование расписания
	y := 50
	count = 0
	addLabel(img, 10, 30, f, data, color.Black)
	for i, lesson := range lessons {
		var color_ color.RGBA
		if count < 2 {
			color_ = color.RGBA{131, 236, 156, 255}
			count++
		} else if count < 4 {
			color_ = color.RGBA{255, 255, 255, 255}
			count++
		} else {
			color_ = color.RGBA{131, 236, 156, 255}
			count = 1
		}
		for x := 0; x < imageWidth; x++ {
			for yy := y; yy < y+50; yy++ {
				img.Set(x, yy, color_)
			}
		}
		addLabel(img, 10, y+30, f, fmt.Sprintf("%d. %s", i+1, lesson[0]), color.Black)
		if len(lesson) > 1 {
			if teacher {
				addLabel(img, imageWidth-150, y+30, f, lesson[1], color.Black)
			} else {
				addLabel(img, imageWidth-100, y+30, f, lesson[1], color.Black)
			}
		}
		y += 50
	}

	// Кодирование изображения в PNG и запись в буфер
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
