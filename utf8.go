package main

import "fmt"

func encode(utf32 []rune) []byte {
	ans := []byte{}
	for _, b := range utf32 {
		if b < 128 {
			ans = append(ans, byte(b))
		} else if b < 2048 {
			ans = append(ans, byte((b>>6)+128+64), byte(b%64+128))
		} else if b < 65536 {
			ans = append(ans, byte((b>>12)+128+64+32), byte((b>>6)%64+128), byte(b%64+128))
		} else {
			ans = append(ans, byte((b>>18)+128+64+32+16), byte((b>>12)%64+128), byte((b>>6)%64+128), byte(b%64+128))
		}
	}
	return ans
}

func decode(utf8 []byte) []rune {
	ans := []rune{}
	i := 0
	for i < len(utf8) {
		if utf8[i]>>4 == 15 {
			ans = append(ans, rune(rune((utf8[i]-128-64-32-16))<<18+rune((utf8[i+1]-128))<<12+rune((utf8[i+2]-128))<<6+rune(utf8[i+3]-128)))
			i += 4
		} else if utf8[i]>>5 == 7 {
			fmt.Println("3 bytes")
			ans = append(ans, rune(rune((utf8[i]-128-64-32))<<12+rune((utf8[i+1]-128))<<6+rune(utf8[i+2]-128)))
			i += 3
		} else if utf8[i]>>6 == 3 {
			ans = append(ans, rune((utf8[i]-128-64)<<6+utf8[i+1]-128))
			i += 2
		} else {
			ans = append(ans, rune(utf8[i]))
			i += 1
		}
	}
	return ans
}

func main() {
	// s := "€"
	s := "𐍈"
	fmt.Println(([]rune)(s))
	fmt.Println(decode(([]byte)(s)))
	// decode()
}
