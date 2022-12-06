package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// max unsigned integer
	// MaxUnit64 = 18446744073709551615
	// MaxUnit64 + 1 = 0
	var aUnsignedInt uint64 = math.MaxUint64
	fmt.Printf("MaxUnit64 = %d\nMaxUnit64 + 1 = %d\n", aUnsignedInt, aUnsignedInt+1)

	// max signed integer
	// MaxInt = 9223372036854775807
	// MaxInt + 1 = -9223372036854775808
	var aSignedInt int64 = math.MaxInt
	fmt.Printf("MaxInt = %d\nMaxInt + 1 = %d\n", aSignedInt, aSignedInt+1)

	// big integer
	// MaxInt in big integer = 9223372036854775807
	// MaxInt + 1 in big integer = 9223372036854775808
	bigInt := big.NewInt(aSignedInt)

	bigIntPlus1 := big.NewInt(bigInt.Int64())
	bigIntPlus1.Add(bigIntPlus1, big.NewInt(1))

	fmt.Printf("MaxInt in big integer = %s\nMaxInt + 1 in big integer = %s\n", bigInt.String(), bigIntPlus1.String())

	// big float
	// MaxFloat64 in big float = 1.797693135e+308
	// MaxFloat64 + 1 in big float = 1.797693135e+308
	aMaxFloat64 := math.MaxFloat64

	bigFloat := big.NewFloat(aMaxFloat64)

	bigFloatPlus1 := big.NewFloat(aMaxFloat64)
	bigFloatPlus1.Add(bigFloatPlus1, big.NewFloat(1.0))

	fmt.Printf("MaxFloat64 in big float = %s\nMaxFloat64 + 1 in big float = %s\n", bigFloat.String(), bigFloatPlus1.String())
}
