package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boljen/go-bitmap"
)

func main() {
	// log.Printf(get("https://jsonplaceholder.typicode.com/posts/1"))
	width := 3
	height := 3
	bitmapBytes := createBitmap(width, height)

	ioutil.WriteFile("./mybitmap2.bmp", bitmapBytes, 0644)
	// readBitmap()
}

func get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	bodyString := string(body)
	return bodyString
}

func constructVoyagerUrl(path string) string {
	baseUrl := "https://www.linkedin.com/voyager/api/"
	return baseUrl + path
}

func toByteArray(i int32) (arr [4]byte) {
	binary.LittleEndian.PutUint32(arr[0:4], uint32(i))
	return
}

func modifyByteArray(bytesToModify []byte, position int, bytesToAdd [4]byte) {
	length := len(bytesToAdd)
	for i := 0; i < length; i++ {
		bytesToModify[i+position] = bytesToAdd[i]
	}
}

func createBitmap(width int, height int) []byte {
	metadataSize := 54
	sizeOfBitmapFile := height * width * 3
	totalSize := metadataSize + sizeOfBitmapFile

	image := bitmap.New(totalSize * 8)
	// BITMAP FILE HEADERS
	// Set first two chars:
	image[0] = 0x42 //'B'
	image[1] = 0x4D //'M'
	// Add size of file
	totalSizeAsBytes := toByteArray(int32(totalSize))
	modifyByteArray(image, 2, totalSizeAsBytes)
	// Add offset of pixel data from start of file
	metadataSizeAsBytes := toByteArray(int32(metadataSize))
	modifyByteArray(image, 10, metadataSizeAsBytes)

	// DIB HEADERS: 14 - 53
	image[14] = 40                                         // size of this header in bytes
	modifyByteArray(image, 18, toByteArray(int32(width)))  // bitmap width in pixels
	modifyByteArray(image, 22, toByteArray(int32(height))) // bitmap width in pixels
	image[26] = 1                                          // number of color planes - must be 1
	image[28] = 1                                          // number of bits per pixel
	// 30: implicitly no image compression
	image[38] = 18
	image[39] = 11
	image[42] = 18
	image[43] = 11
	// COLOR TABLE - two 4-byte colours: 54-61
	// Implicitly - 54 - 56 are the colour white, 57 is the alpha channel
	image[58] = 255
	image[59] = 255
	image[60] = 255
	// 61 is alpha channel

	// Image data - 62+
	// Each row is rounded up to the nearest multiple of 4 bytes, with
	// colours as bits in a row
	image[62] = 160
	image[66] = 64
	image[70] = 160
	// bits := [8]bool{true, true, true, false, false, false, false, false}
	// var b byte
	// b = 0
	// for index, element := range bits {
	// 	fmt.Println(index)
	// 	fmt.Println(element)

	// 	bitmap.SetBit(b, index, element)
	// 	fmt.Println(b)

	// }
	// fmt.Println(b)
	// bitwise or |= operator can set bits.
	// XOR can be used to toggle bits from one value to another.
	for i, s := range image {
		if i < 100 {
			fmt.Println(i, s)
		}
	}
	return image.Data(true)
}

func readBitmap() {
	image, err := ioutil.ReadFile("./example5.bmp")
	if err != nil {
		log.Fatalln(err)
	}
	for i, s := range image {
		fmt.Println(i, s)
	}
}

func bitsToIntArray(bits []bool) []int {
	// start with num 1. Double for each bit. for each true, add the num to an int.
	nums := (make([]int, 0))
	for i, bit := range bits {
		if bit {
			nums[(i+1)/8] += 2 ^ i%8
		}
	}
	return nums
	// Could return an array of ints. The consumer can round length up to nearest multiple of 4.
}

// TODO:
// - Pass in an array of pixels? 2d array?
// - Add function that can translate a smaller 'drawing' (2d array)
//   to a particular coordinate on a bigger 2d array
// - Add a function that can translate a 2d array to a bunch of bytes.
