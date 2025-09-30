package examples

import (
	"math/big"
)

/*
f(x) = (a * x) mod p - прямое преобразование
f⁻¹(x) = (a⁻¹ * x) mod p - обратное преобразование

Здесь:
p - большое простое число
a и p - взаимнопростые
a⁻¹ - мультипликативная инверсия a по модулю p
*/

type transformer struct {
	// произведение n * e.a может быть больше MaxInt64,
	// поэтому вычисляем через big.Int
	p        *big.Int // большое простое число
	a        *big.Int // a взаимнопростое с p
	inv      *big.Int // инверсия a
	maxValue int64    // максимальное значение, для которого может применяться трансформация - ограничивает, чтобы не было коллизий
}

func newTransformer() *transformer {
	m := int64(1<<31 - 1) // большое простое число 2147483647
	// Множитель
	a := big.NewInt(123456789)

	// Находим обратный элемент по модулю для множителя a
	inverse := new(big.Int).ModInverse(
		a,
		big.NewInt(m),
	)
	return &transformer{
		p:        big.NewInt(m),
		a:        a,
		inv:      inverse,
		maxValue: 1_000_000_000_000,
	}
}

func (e *transformer) modularTransform(n int64) int64 {
	if n > e.maxValue {
		return 0
	}
	com := new(big.Int).Mul(big.NewInt(n), e.a)
	return com.Mod(com, e.p).Int64()
}

func (e *transformer) modularInverse(n int64) int64 {
	if n > e.maxValue {
		return 0
	}
	com := new(big.Int).Mul(big.NewInt(n), e.inv)
	return com.Mod(com, e.p).Int64()
}
