package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomNumber(len int, to int, data *string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		*data += fmt.Sprintf("%d", rand.Intn(10))
	}
}

func ReformatMDN(mdn string) string {
	mdn = strings.Trim(strings.Trim(mdn, " "), "+")
	if strings.HasPrefix(mdn, "628") {
		return mdn
	}

	if strings.HasPrefix(mdn, "+62") {
		//remove leading +62
		mdn = strings.Replace(mdn, "+62", "62", 1)
	} else if strings.HasPrefix(mdn, "08") {
		//remove leading 62
		mdn = strings.Replace(mdn, "08", "628", 1)
	} else if strings.HasPrefix(mdn, "8") {
		//remove leading 0
		mdn = strings.Replace(mdn, "8", "628", 1)
	}

	return mdn
}

var MDNPrefix62 = 1
var MDNPrefix08 = 2

func ReformatMDNWithPrefiX(mdn string, prefixType int) string {
	mdn = strings.Trim(mdn, " ")
	var prefixOld, prefixNew []string
	if prefixType == MDNPrefix62 {
		prefixOld, prefixNew = []string{"+62", "08", "8"}, []string{"62", "628", "628"}
	} else if prefixType == MDNPrefix08 {
		prefixOld, prefixNew = []string{"+62", "628", "62"}, []string{"0", "08", "0"}
	}

	for i, elm := range prefixOld {
		if strings.HasPrefix(mdn, elm) {
			mdn = strings.Replace(mdn, elm, prefixNew[i], 1)
			break
		}
	}

	return mdn
}

func NewReformatMDN(mdn string, pref string) string { // jika pref ada value maka mdn akan dtambah prefixnya
	mdn = strings.Trim(strings.Trim(mdn, " "), "+") // remove char +
	if strings.HasPrefix(mdn, "0") {                // remove prefix char 0 one digit prefix
		mdn = strings.TrimPrefix(mdn, "0")
	}

	if strings.HasPrefix(mdn, "62") { // remove prefix 62XXX -> XXXX
		//remove leading +62
		mdn = strings.Replace(mdn, "62", "", 1)
	}

	if mdn == "" {
		return mdn
	}

	return pref + mdn
}
