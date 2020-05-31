package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Parametros para la convolucion de la imagen
var smooth int = 1
var matrixConv = [3][3]int{
	{0, 0, 0},
	{0, 1, 0},
	{0, 0, 0}}

// Levanta el servidor, y enruta las direcciones de los recuros de la aplicacion
func main() {
	fmt.Println("Servidor levantado en la direccion http://localhost:8080")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/uploadImage", secondServer)
	http.HandleFunc("/", firstServer)

	http.ListenAndServe(":8080", nil)
}

// Muestra el error y detiene la aplicacion
func CheckFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Muestra una pagina de error
func CheckServerError(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err.Error())
		t, otherError := template.ParseFiles("html/error.html")
		CheckFatalError(otherError)
		t.Execute(w, nil)
	}
}

// Redirige a la pagina principal
// O inicia el procesamiento de la imagen si esta es mandada en un POST y redirige hacia donde se mostrada el resultado
func firstServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		// Obtiene la imagen y la decodifica
		inputFile, _, err := r.FormFile("inputImg")
		CheckServerError(err, w)
		defer inputFile.Close()
		img, _, err := image.Decode(inputFile)
		CheckServerError(err, w)

		// Aplica la convolucion
		imgConv, err := Convolution(img)
		CheckServerError(err, w)

		// Codifica la imagen a JPEG
		buffer := new(bytes.Buffer)
		jpeg.Encode(buffer, imgConv, nil)
		CheckServerError(err, w)

		// Codifica la imagen a String
		str := base64.StdEncoding.EncodeToString(buffer.Bytes())

		// Envia la imagen en forma de string a la pagina donde sera mostrada
		t, err := template.ParseFiles("html/result.html")
		CheckServerError(err, w)
		data := map[string]interface{}{"Image": str}
		t.Execute(w, data)
	} else {
		t, err := template.ParseFiles("html/uploadData.html")
		CheckServerError(err, w)
		t.Execute(w, nil)
	}
}

// Procesa los datos para aplicar el filtro a la imagen y redirige a subir la imagen
func secondServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		r.ParseForm()

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				value := string(r.Form.Get("matrix" + strconv.Itoa(i) + "-" + strconv.Itoa(j)))
				number, err := strconv.Atoi(value)
				if err != nil {
					matrixConv[i][j] = 1
				} else {
					matrixConv[i][j] = number + 1
				}
			}
		}

		value := string(r.Form.Get("smooth"))
		number, err := strconv.Atoi(value)
		if err != nil {
			smooth = 1
		} else {
			smooth = number
		}
	} else {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				matrixConv[i][j] = 1
			}
		}
		smooth = 1
	}
	t, err := template.ParseFiles("html/uploadImg.html")
	CheckServerError(err, w)
	t.Execute(w, nil)
}

// Devulve la posicion ajustada del pixel
func limitPixel(coord int, max int) int {
	if coord >= max {
		return 0
	} else if coord < 0 {
		return max - 1
	}
	return coord
}

// Devuelve el valor ajustado en un rango de 0 a 255
func limitUnit(number int) int {
	var unitBase int = 0xff
	number = number / unitBase
	// fmt.Printf("%d\n", number)
	if number >= 0xff {
		return 0xff
	} else if number <= 0 {
		return 0
	}
	return number
}

// Aplica la convolocion a la imagen
func Convolution(img image.Image) (*image.NRGBA, error) {
	imageNRGBA := image.NewNRGBA(img.Bounds())

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			var sumR = 0
			var sumG = 0
			var sumB = 0

			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					imageX := limitPixel(x+i, w)
					imageY := limitPixel(y+j, h)

					r, g, b, _ := img.At(imageX, imageY).RGBA()
					sumR = (sumR + (int(r) * matrixConv[i+1][j+1]))
					sumG = (sumG + (int(g) * matrixConv[i+1][j+1]))
					sumB = (sumB + (int(b) * matrixConv[i+1][j+1]))
				}
			}

			imageNRGBA.Set(x, y, color.NRGBA{
				uint8(limitUnit((sumR / 9) / smooth)),
				uint8(limitUnit((sumG / 9) / smooth)),
				uint8(limitUnit((sumB / 9) / smooth)),
				255,
			})
		}
	}

	if imageNRGBA == nil {
		return nil, fmt.Errorf("Imagen no soportada")
	}

	return imageNRGBA, nil
}
