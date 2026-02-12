package utils

import (
	"fmt"
	"testing"
)

func TestGoodNo(t *testing.T) {
	no1 := "1234567890"
	no2 := "1323456789"
	no3 := "1233389"
	no4 := "12344666789"
	no5 := "122323456789"
	no6 := "2023156789"

	fmt.Println("IsConsecutiveNumber")

	r1, b1 := IsConsecutiveNumber(no1, 3)
	fmt.Println(no1, r1, b1)
	r2, b2 := IsConsecutiveNumber(no2, 3)
	fmt.Println(no2, r2, b2)
	r3, b3 := IsConsecutiveNumber(no3, 3)
	fmt.Println(no3, r3, b3)
	r4, b4 := IsConsecutiveNumber(no4, 3)
	fmt.Println(no4, r4, b4)
	r5, b5 := IsConsecutiveNumber(no5, 3)
	fmt.Println(no5, r5, b5)
	r6, b6 := IsConsecutiveNumber(no6, 3)
	fmt.Println(no6, r6, b6)

	fmt.Println()
	fmt.Println("IsSequentialNumber")

	r1, b1 = IsSequentialNumber(no1, 3)
	fmt.Println(no1, r1, b1)
	r2, b2 = IsSequentialNumber(no2, 3)
	fmt.Println(no2, r2, b2)
	r3, b3 = IsSequentialNumber(no3, 3)
	fmt.Println(no3, r3, b3)
	r4, b4 = IsSequentialNumber(no4, 3)
	fmt.Println(no4, r4, b4)
	r5, b5 = IsSequentialNumber(no5, 3)
	fmt.Println(no5, r5, b5)
	r6, b6 = IsSequentialNumber(no6, 3)
	fmt.Println(no6, r6, b6)
	fmt.Println()
	fmt.Println("IsAABBNumber")

	r1, b1 = IsAABBNumber(no1)
	fmt.Println(no1, r1, b1)
	r2, b2 = IsAABBNumber(no2)
	fmt.Println(no2, r2, b2)
	r3, b3 = IsAABBNumber(no3)
	fmt.Println(no3, r3, b3)
	r4, b4 = IsAABBNumber(no4)
	fmt.Println(no4, r4, b4)
	r5, b5 = IsAABBNumber(no5)
	fmt.Println(no5, r5, b5)
	r6, b6 = IsAABBNumber(no6)
	fmt.Println(no6, r6, b6)

	fmt.Println()
	fmt.Println("IsABABNumber")

	r1, b1 = IsABABNumber(no1)
	fmt.Println(no1, r1, b1)
	r2, b2 = IsABABNumber(no2)
	fmt.Println(no2, r2, b2)
	r3, b3 = IsABABNumber(no3)
	fmt.Println(no3, r3, b3)
	r4, b4 = IsABABNumber(no4)
	fmt.Println(no4, r4, b4)
	r5, b5 = IsABABNumber(no5)
	fmt.Println(no5, r5, b5)
	r6, b6 = IsABABNumber(no6)
	fmt.Println(no6, r6, b6)

	fmt.Println()
	fmt.Println("IsBirthdayNumber")

	r1, b1 = IsBirthdayNumber(no1)
	fmt.Println(no1, r1, b1)
	r2, b2 = IsBirthdayNumber(no2)
	fmt.Println(no2, r2, b2)
	r3, b3 = IsBirthdayNumber(no3)
	fmt.Println(no3, r3, b3)
	r4, b4 = IsBirthdayNumber(no4)
	fmt.Println(no4, r4, b4)
	r5, b5 = IsBirthdayNumber(no5)
	fmt.Println(no5, r5, b5)
	r6, b6 = IsBirthdayNumber(no6)
	fmt.Println(no6, r6, b6)

}
