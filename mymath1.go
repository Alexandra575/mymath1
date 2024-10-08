package mymath1

import "math"

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Acos(x float64) float64 {
	return math.Acos(x)
}

func Acosh(x float64) float64 {
	return math.Acosh(x)
}

func Asin(x float64) float64 {
	return math.Asin(x)
}

func Atan(x float64) float64 {
	return math.Atan(x)
}

func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

func Atanh(x float64) float64 {
	return math.Atanh(x)
}

func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Cosh(x float64) float64 {
	return math.Cosh(x)
}

func Exp(x float64) float64 {
	return math.Exp(x)
}

func Exp2(x float64) float64 {
	return math.Exp2(x)
}

func Expm1(x float64) float64 {
	return math.Expm1(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Gamma(x float64) float64 {
	return math.Gamma(x)
}

func Hypot(p, q float64) float64 {
	return math.Hypot(p, q)
}

func Inf() bool {
	return math.IsInf(1, 0) 
}

func IsInf(x float64, sign int) bool {
	return math.IsInf(x, sign)
}

func IsNaN(x float64) bool {
	return math.IsNaN(x)
}

func J0(x float64) float64 {
	return math.J0(x)
}

func J1(x float64) float64 {
	return math.J1(x)
}

func Ldexp(frac float64, exp int) float64 {
	return math.Ldexp(frac, exp)
}

func Log(x float64) float64 {
	return math.Log(x)
}

func Log10(x float64) float64 {
	return math.Log10(x)
}

func Log1p(x float64) float64 {
	return math.Log1p(x)
}

func Log2(x float64) float64 {
	return math.Log2(x)
}

func Logb(x float64) float64 {
	return math.Logb(x)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Mod(x, y float64) float64 {
	return math.Mod(x, y)
}

func Modf(f float64) (frac, int float64) {
	return math.Modf(f)
}

func NaN() bool {
	return math.IsNaN(1)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Pow10(x int) float64 {
	return math.Pow10(x)
}

func Remainder(x, y float64) float64 {
	return math.Remainder(x, y)
}

func Round(x float64) float64 {
	return math.Round(x)
}

func RoundToEven(x float64) float64 {
	return math.RoundToEven(x)
}

func Signbit(x float64) bool {
	return math.Signbit(x)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Sincos(x float64) (sin, cos float64) {
	return math.Sincos(x)
}

func Sinh(x float64) float64 {
	return math.Sinh(x)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Tan(x float64) float64 {
	return math.Tan(x)
}

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func Trunc(x float64) float64 {
	return math.Trunc(x)
}

func Y0(x float64) float64 {
	return math.Y0(x)
}

func Y1(x float64) float64 {
	return math.Y1(x)
}
