package main

import (
	"fmt"
	qr "qrcode"
)

func main() {
	// src := "../test_png.png"
	// fmt.Println(qr_common.PrintFromFile(src))
	// return
	// f, err := os.Open("test3_src.jpg")
	// fmt.Println(err)
	// bts, err := io.ReadAll(f)
	// fmt.Println(err)
	// src2 := base64.StdEncoding.EncodeToString(bts)
	// //fmt.Println(src2)
	// //src2 := `iVBORw0KGgoAAAANSUhEUgAAAZAAAAGQAQAAAACoxAthAAACYklEQVR42u2bO3LCMBCGN+PCpY/AUXw0OBpH4QguKRhv9imbxMCkW01+FYkQ+tSs9i2I/zwICBAgQIAAAQIEyP9BbuRjZr4Pqy7eprsunK8nvseXE5AeEP94mWXnIKuycxlZ2Isdst8DpDqiIrYNvO1cQ/o6oRlIX4hodOwM6ROQThGWnanR9CUazQykMyRtMo0P86ryR40zhZV+Z8aBlEIyUtINZpP3k/fBFZBSyDYWEv214forNvlTygOkEuLSF2dKoshqik2R16eVE5AukClCXEr99WtgIZMZ5/NB/gKkICJDRLwlnhfTXz0kv77QiYH0gXATsRaIrmQTd7h6DV64VyDVkG09U1GdqZ+VybgdAqQHZPB10lRU67eLeVWyqgLzLmQCUhoRZzpEOYglFRW21YU4jPMVSBfIzWsIVuIzP9vcq0VK63H2CqQiMoW2erZih4x5mpWM/BoAKY9w018rvHu11r2ttjhfBFdAKiILZTPF3WtsUHZnroHUR24tsiWXvoa4D5f+SzMOpCgSXjU7m1Hrs0NMx3+XbYEURey/1fp2JaPNJp+vBxcGSEHEd3q1llzo2eKMpicfBldA6iH8SBE/5y8RRB0+PABSEmlVoPvPEh/lfQDSA9JGNMjsDa2lLXPW5A/6YkAKIrs3JBz5C7VYtz28BNIDki+7PPGMtjVngLS+e9AOpBiST2ddbbf8JZ+XHLbSgBRHdk+8rKowfpY+kIJIytqDXm+QuZ/9nfIAKYm0XwxF/hKm+BwFIiDdIC1Syqcj0evcJTInID0gZX/DCwQIECBAgAABAqQf5Bv31XBrT0SVVAAAAABJRU5ErkJggg==`
	// fmt.Println(qr_common.Base64ToFile(src2))

	// src2 := `iVBORw0KGgoAAAANSUhEUgAAAZAAAAGQAQAAAACoxAthAAACYklEQVR42u2bO3LCMBCGN+PCpY/AUXw0OBpH4QguKRhv9imbxMCkW01+FYkQ+tSs9i2I/zwICBAgQIAAAQIEyP9BbuRjZr4Pqy7eprsunK8nvseXE5AeEP94mWXnIKuycxlZ2Isdst8DpDqiIrYNvO1cQ/o6oRlIX4hodOwM6ROQThGWnanR9CUazQykMyRtMo0P86ryR40zhZV+Z8aBlEIyUtINZpP3k/fBFZBSyDYWEv214forNvlTygOkEuLSF2dKoshqik2R16eVE5AukClCXEr99WtgIZMZ5/NB/gKkICJDRLwlnhfTXz0kv77QiYH0gXATsRaIrmQTd7h6DV64VyDVkG09U1GdqZ+VybgdAqQHZPB10lRU67eLeVWyqgLzLmQCUhoRZzpEOYglFRW21YU4jPMVSBfIzWsIVuIzP9vcq0VK63H2CqQiMoW2erZih4x5mpWM/BoAKY9w018rvHu11r2ttjhfBFdAKiILZTPF3WtsUHZnroHUR24tsiWXvoa4D5f+SzMOpCgSXjU7m1Hrs0NMx3+XbYEURey/1fp2JaPNJp+vBxcGSEHEd3q1llzo2eKMpicfBldA6iH8SBE/5y8RRB0+PABSEmlVoPvPEh/lfQDSA9JGNMjsDa2lLXPW5A/6YkAKIrs3JBz5C7VYtz28BNIDki+7PPGMtjVngLS+e9AOpBiST2ddbbf8JZ+XHLbSgBRHdk+8rKowfpY+kIJIytqDXm+QuZ/9nfIAKYm0XwxF/hKm+BwFIiDdIC1Syqcj0evcJTInID0gZX/DCwQIECBAgAABAqQf5Bv31XBrT0SVVAAAAABJRU5ErkJggg==`
	// q, err := qr.NewQrcodeFromBase64(src2)
	qc, err := qr.NewQrcodeFromPath("../test_png.png")
	if err != nil {
		fmt.Println(err)
	}
	qc.PrintSourceMatrix()
	qc.SetBound(1)
	qc.PrintOnWindowsConsole()
	//qc.PrintOnWindowsConsole()
}
