package binaryFlag

import (
	"testing"
)

func TestSetBinaryFlag(t *testing.T) {
	flag := New()
	flag.SetBit(1)
	flag.SetBit(3)
	flag.SetBit(5)
	flag.SetBit(7)
	flag.SetBit(9)
	flag.SetBit(11)
	flag.SetBit(13)
	flag.SetBit(15)
	flag.SetBit(17)
	flag.SetBit(19)
	flag.SetBit(21)
	flag.SetBit(23)
	flag.SetBit(25)
	flag.SetBit(27)
	flag.SetBit(29)
	flag.SetBit(31)

	flag2 := New()
	flag.SetBit(2)
	flag.SetBit(4)
	flag.SetBit(6)
	flag.SetBit(8)
	flag.SetBit(10)
	flag.SetBit(12)
	flag.SetBit(14)
	flag.SetBit(16)
	flag.SetBit(18)
	flag.SetBit(20)
	flag.SetBit(22)
	flag.SetBit(24)
	flag.SetBit(26)
	flag.SetBit(28)
	flag.SetBit(30)

	if flag.SetBinary(*flag2).DefaultValueStr() != "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Error("Test set binary")
	}
}
func TestSimpleFlag(t *testing.T) {
	flag := New()
	var result = ""
	for i := 0; i < FlagSize; i++ {
		flag.SetBit(i)
		result += flag.DefaultValueStr() + "\n"
	}

	if result != "______________________________X\n_____________________________XX\n____________________________XXX\n___________________________XXXX\n__________________________XXXXX\n_________________________XXXXXX\n________________________XXXXXXX\n_______________________XXXXXXXX\n______________________XXXXXXXXX\n_____________________XXXXXXXXXX\n____________________XXXXXXXXXXX\n___________________XXXXXXXXXXXX\n__________________XXXXXXXXXXXXX\n_________________XXXXXXXXXXXXXX\n________________XXXXXXXXXXXXXXX\n_______________XXXXXXXXXXXXXXXX\n______________XXXXXXXXXXXXXXXXX\n_____________XXXXXXXXXXXXXXXXXX\n____________XXXXXXXXXXXXXXXXXXX\n___________XXXXXXXXXXXXXXXXXXXX\n__________XXXXXXXXXXXXXXXXXXXXX\n_________XXXXXXXXXXXXXXXXXXXXXX\n________XXXXXXXXXXXXXXXXXXXXXXX\n_______XXXXXXXXXXXXXXXXXXXXXXXX\n______XXXXXXXXXXXXXXXXXXXXXXXXX\n_____XXXXXXXXXXXXXXXXXXXXXXXXXX\n____XXXXXXXXXXXXXXXXXXXXXXXXXXX\n___XXXXXXXXXXXXXXXXXXXXXXXXXXXX\n__XXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n" {
		t.Error("Test set failed")
	}
	result = ""
	for i := 0; i < FlagSize; i++ {
		flag.UnSetBit(i)
		result += flag.DefaultValueStr() + "\n"
	}
	if result != "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX_\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXX__\nXXXXXXXXXXXXXXXXXXXXXXXXXXXX___\nXXXXXXXXXXXXXXXXXXXXXXXXXXX____\nXXXXXXXXXXXXXXXXXXXXXXXXXX_____\nXXXXXXXXXXXXXXXXXXXXXXXXX______\nXXXXXXXXXXXXXXXXXXXXXXXX_______\nXXXXXXXXXXXXXXXXXXXXXXX________\nXXXXXXXXXXXXXXXXXXXXXX_________\nXXXXXXXXXXXXXXXXXXXXX__________\nXXXXXXXXXXXXXXXXXXXX___________\nXXXXXXXXXXXXXXXXXXX____________\nXXXXXXXXXXXXXXXXXX_____________\nXXXXXXXXXXXXXXXXX______________\nXXXXXXXXXXXXXXXX_______________\nXXXXXXXXXXXXXXX________________\nXXXXXXXXXXXXXX_________________\nXXXXXXXXXXXXX__________________\nXXXXXXXXXXXX___________________\nXXXXXXXXXXX____________________\nXXXXXXXXXX_____________________\nXXXXXXXXX______________________\nXXXXXXXX_______________________\nXXXXXXX________________________\nXXXXXX_________________________\nXXXXX__________________________\nXXXX___________________________\nXXX____________________________\nXX_____________________________\nX______________________________\n_______________________________\n" {
		t.Error("Test unset failed")
	}

	flag = New()
	flag.SetBit(1)
	flag.SetBit(3)
	flag.SetBit(5)
	flag.SetBit(7)
	flag.SetBit(9)
	flag.SetBit(11)
	flag.SetBit(13)
	flag.SetBit(15)
	flag.SetBit(17)
	flag.SetBit(19)
	flag.SetBit(21)
	flag.SetBit(23)
	flag.SetBit(25)
	flag.SetBit(27)
	flag.SetBit(29)
	flag.SetBit(31)

	if flag.DefaultValueStr() != "X_X_X_X_X_X_X_X_X_X_X_X_X_X_X_X" {
		t.Error("Test singular set")
	}

	flag = New()
	flag.SetBit(2)
	flag.SetBit(4)
	flag.SetBit(6)
	flag.SetBit(8)
	flag.SetBit(10)
	flag.SetBit(12)
	flag.SetBit(14)
	flag.SetBit(16)
	flag.SetBit(18)
	flag.SetBit(20)
	flag.SetBit(22)
	flag.SetBit(24)
	flag.SetBit(26)
	flag.SetBit(28)
	flag.SetBit(30)

	if flag.DefaultValueStr() != "_X_X_X_X_X_X_X_X_X_X_X_X_X_X_X_" {
		t.Error("Test odd set")
	}

	if flag.Toggle().DefaultValueStr() != "X_X_X_X_X_X_X_X_X_X_X_X_X_X_X_X" {
		t.Error("Test singular set")
	}
}