package ppfevoucherpackage

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	numRows       = 15
	numColumns    = 10
	beginCharacter = '9'
)

var encryptionArray [numRows][numColumns]byte

func init() {
	var encryptValue byte

	for row := 0; row < numRows; row++ {
		encryptValue = byte(beginCharacter - (row % numColumns))
		if encryptValue < '0' {
			encryptValue = '9'
		}
		for column := 0; column < numColumns; column++ {
			encryptionArray[row][column] = encryptValue
			encryptValue = encryptValue - 1
			if encryptValue < '0' {
				encryptValue = '9'
			}
		}
	}
	fmt.Println("Voucher package v1.0.0")
}

func encryptPin(pinNumber uint64) uint64 {
	var tempPin, scrambledPin, newPin string
	var newVoucherNumber int64

	tempPin = fmt.Sprintf("%015d", pinNumber)

	scrambledPin = tempPin[4:5] + tempPin[2:3] + tempPin[5:6] + tempPin[6:7] +
		tempPin[7:8] + tempPin[8:9] + tempPin[3:4] + tempPin[9:10] +
		tempPin[10:11] + tempPin[11:12] + tempPin[12:13] + tempPin[1:2] +
		tempPin[13:14] + tempPin[0:1] + tempPin[14:15]

	for row := 0; row < numRows; row++ {
		for column := 0; column < numColumns; column++ {
			if scrambledPin[row] == encryptionArray[row][column] {
				newPin += string(rune('0' + column))
				break
			}
		}
	}

	newVoucherNumber, _ = strconv.ParseInt(newPin, 10, 64)
	return uint64(newVoucherNumber)
}

func decryptPin(pinNumber uint64) uint64 {
	var tempPin, scrambledPin string
	decryptedPin := int64(0)

	tempPin = fmt.Sprintf("%015d", pinNumber)

	for row := 0; row < numRows; row++ {
		column, _ := strconv.Atoi(tempPin[row : row+1])
		scrambledPin += string(encryptionArray[row][column])
	}

	tempPin = scrambledPin[13:14] + scrambledPin[11:12] + scrambledPin[1:2] + scrambledPin[6:7] + 
		      scrambledPin[0:1] + scrambledPin[2:3] + scrambledPin[3:4] + scrambledPin[4:5] +
			  scrambledPin[5:6] + scrambledPin[7:8] + scrambledPin[8:9] + scrambledPin[9:10] +
	          scrambledPin[10:11] + scrambledPin[12:13] + scrambledPin[14:15]

    decryptedPin, _ = strconv.ParseInt(tempPin, 10, 64)
	return uint64(decryptedPin)
}

func splitSecurityPinAndSerialNumber(voucherNumber uint64) (uint64, uint32) {
	serialNumber := voucherNumber % 100000000000
	securityPin := uint32(voucherNumber / 100000000000)
	return serialNumber, securityPin
}

func randomizeSerialNumber(oldSerialNumber uint64) (uint64,error) {
	var serialNumber uint64 = 0

	if oldSerialNumber > 68719476735 {
		return 0, errors.New("Input value exceeds the maximum 36-bit unsigned integer (value:68719476735)")
	}

	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 35)) << 0)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 34)) << 1)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 33)) << 2)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 32)) << 3)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 4)) << 4)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 5)) << 5)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 30)) << 6)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 31)) << 7)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 8)) << 8)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 9)) << 9)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 29)) << 10)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 28)) << 11)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 27)) << 12)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 26)) << 13)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 25)) << 14)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 24)) << 15)

	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 23)) << 16)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 22)) << 17)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 21)) << 18)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 20)) << 19)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 19)) << 20)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 18)) << 21)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 17)) << 22)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 16)) << 23)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 15)) << 24)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 14)) << 25)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 13)) << 26)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 12)) << 27)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 11)) << 28)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 10)) << 29)

	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 6)) << 30)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 7)) << 31)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 3)) << 32)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 2)) << 33)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 1)) << 34)
	serialNumber = serialNumber | ((1 & (oldSerialNumber >> 0)) << 35)

	return serialNumber,nil
}

func combineSecurityPinAndSerialNumber(serialNumber uint64, securityPin uint32) uint64 {
	combinedSecurityPinAndSerialNumber := (uint64(securityPin) * 100000000000) + serialNumber
	return combinedSecurityPinAndSerialNumber
}

func unRandomizeSerialNumber(serialNumber uint64) uint64 {
	var newSerialNumber uint64 = 0

	newSerialNumber = newSerialNumber | ((1 & serialNumber) << 35)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 1)) << 34)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 2)) << 33)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 3)) << 32)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 4)) << 4)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 5)) << 5)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 6)) << 30)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 7)) << 31)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 8)) << 8)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 9)) << 9)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 10)) << 29)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 11)) << 28)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 12)) << 27)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 13)) << 26)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 14)) << 25)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 15)) << 24)

    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 16)) << 23)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 17)) << 22)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 18)) << 21)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 19)) << 20)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 20)) << 19)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 21)) << 18)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 22)) << 17)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 23)) << 16)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 24)) << 15)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 25)) << 14)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 26)) << 13)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 27)) << 12)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 28)) << 11)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 29)) << 10)

    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 30)) << 6)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 31)) << 7)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 32)) << 3)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 33)) << 2)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 34)) << 1)
    newSerialNumber = newSerialNumber | ((1 & (serialNumber >> 35)))

	return newSerialNumber
}


func GetSerialNumberAndSecurityPin(voucherNumber uint64) (uint64, uint32) {
	var randomSerialNumber uint64
	var decryptedVoucherNumber, serialNumber uint64

	decryptedVoucherNumber = decryptPin(voucherNumber)
	randomSerialNumber, securityPin := splitSecurityPinAndSerialNumber(decryptedVoucherNumber)
	serialNumber = unRandomizeSerialNumber(randomSerialNumber)

	return serialNumber, securityPin
}


func ManufactureVoucherNumber(serialNumber uint64, securityPin uint32) (uint64, error) {
	randomSerialNumber, err := randomizeSerialNumber(serialNumber)
	if err != nil {
		return 0, err
	}

	serialAndPinNumber := combineSecurityPinAndSerialNumber(randomSerialNumber, securityPin)
	encryptedVoucherNumber := encryptPin(serialAndPinNumber)
	return encryptedVoucherNumber, nil
}

