// https://leetcode.com/problems/flipping-an-image/

package flipping_an_image

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		image[i] = inverse(reverse(image[i]))
	}
	return image
}

func reverse(pieceOfImage []int) []int {
	for i := 0; i < len(pieceOfImage)-1; i++ {
		for j := i + 1; j < len(pieceOfImage); j++ {
			pieceOfImage[i], pieceOfImage[j] = pieceOfImage[j], pieceOfImage[i]
		}
	}
	return pieceOfImage
}

func inverse(pieceOfImage []int) []int {
	for i := 0; i < len(pieceOfImage); i++ {
		if pieceOfImage[i] == 0 {
			pieceOfImage[i] = 1
		} else {
			pieceOfImage[i] = 0
		}
	}
	return pieceOfImage
}
