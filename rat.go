package rat

import (
	"database/sql/driver"
	"fmt"
	"log/slog"
	"math/big"
	"strconv"
	"strings"
)

var DefaultPrecision = 8

type Rational struct {
	bigrat    big.Rat
	precision int
}

// return a/b*100
func Ratio(a *Rational, b *Rational) *Rational {
	return b.Add(a.Neg()).Quo(b).Mul(100)
}

func RatMin(first *Rational, args ...*Rational) *Rational {
	if len(args) == 0 {
		return first
	}
	out := first.Clone()
	for _, arg := range args {
		if out.IsGreaterThan(arg) {
			out.Set(arg)
		}
	}
	return out
}

func RatMax(first *Rational, args ...*Rational) *Rational {
	if len(args) == 0 {
		return first
	}
	ret := first.Clone()
	for _, arg := range args {
		if ret.IsLessThan(arg) {
			ret.Set(arg)
		}
	}
	return ret
}

func Rat[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float64 | ~float32 | ~string | *Rational](v T) *Rational {
	switch v := any(v).(type) {
	case float32:
		return parseFloat64(float64(v))
	case float64:
		return parseFloat64(float64(v))
	case int:
		return parseInt64(int64(v))
	case int8:
		return parseInt64(int64(v))
	case int16:
		return parseInt64(int64(v))
	case int32:
		return parseInt64(int64(v))
	case int64:
		return parseInt64(int64(v))
	case *Rational:
		return RatClone(v)
	case string:
		return parse(v)
	default:
		panic("rat: invalid type")
	}
	return nil
}

func (r *Rational) UnmarshalJSON(data []byte) error {
	sdata := string(data)
	sdata = strings.ReplaceAll(sdata, "\"", "")
	*r = *Rat(sdata)
	return nil
}

func (r *Rational) FloorInt() int {
	v := r.String()
	if strings.Index(v, ".") != -1 {
		v = v[0:strings.Index(v, ".")]
	}
	out, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	if out < 0 {
		return out - 1
	}
	return out
}

func (r *Rational) IntString() string {
	return fmt.Sprint(r.FloorInt())
}

func (r *Rational) Round() *Rational {
	return Rat(r.Add(Rat("0.5")).FloorInt())
}

func (r *Rational) Ceil() *Rational {
	if r.bigrat.IsInt() {
		return r.Clone()
	} else {
		return Rat(r.FloorInt() + 1)
	}
}

func (r *Rational) Floor() *Rational {
	if r.bigrat.IsInt() {
		return r.Clone()
	} else {
		if r.bigrat.Sign() < -1 {
			return Rat(r.FloorInt() - 1)
		}
		return Rat(r.FloorInt())
	}
}

func (r *Rational) Float64() float64 {
	out, _ := r.bigrat.Float64()
	return out
}

func (r *Rational) Sub(in any) *Rational {
	return r.Neg().Add(in).Neg()
}

func (r *Rational) Add(in any) *Rational {
	out := r.Clone()

	switch v := in.(type) {
	case int:
		out.bigrat.Add(&r.bigrat, &Rat(v).bigrat)
	case int32:
		out.bigrat.Add(&r.bigrat, &Rat(v).bigrat)
	case int64:
		out.bigrat.Add(&r.bigrat, &Rat(v).bigrat)
	case float32:
		out.bigrat.Add(&r.bigrat, &Rat(v).bigrat)
	case float64:
		out.bigrat.Add(&r.bigrat, &Rat(v).bigrat)
	case string:
		out.bigrat.Add(&r.bigrat, &Rat(v).bigrat)
	case *Rational:
		out.bigrat.Add(&r.bigrat, &v.bigrat)
	default:
		panic("rat: add invalid type")
	}
	return out
}

func (r *Rational) Neg() *Rational {
	out := r.Clone()
	out.bigrat.Neg(&r.bigrat)
	return out
}

func (r *Rational) Mul(in any) *Rational {
	out := r.Clone()

	switch v := in.(type) {
	case int:
		out.bigrat.Mul(&r.bigrat, &Rat(v).bigrat)
	case int32:
		out.bigrat.Mul(&r.bigrat, &Rat(v).bigrat)
	case int64:
		out.bigrat.Mul(&r.bigrat, &Rat(v).bigrat)
	case float32:
		out.bigrat.Mul(&r.bigrat, &Rat(v).bigrat)
	case float64:
		out.bigrat.Mul(&r.bigrat, &Rat(v).bigrat)
	case string:
		out.bigrat.Mul(&r.bigrat, &Rat(v).bigrat)
	case *Rational:
		out.bigrat.Mul(&r.bigrat, &v.bigrat)
	default:
		panic("rat: add invalid type")
	}
	return out
}

func (r *Rational) Quo(in any) *Rational {
	out := r.Clone()

	switch v := in.(type) {
	case int:
		out.bigrat.Quo(&r.bigrat, &Rat(v).bigrat)
	case int32:
		out.bigrat.Quo(&r.bigrat, &Rat(v).bigrat)
	case int64:
		out.bigrat.Quo(&r.bigrat, &Rat(v).bigrat)
	case float32:
		out.bigrat.Quo(&r.bigrat, &Rat(v).bigrat)
	case float64:
		out.bigrat.Quo(&r.bigrat, &Rat(v).bigrat)
	case string:
		out.bigrat.Quo(&r.bigrat, &Rat(v).bigrat)
	case *Rational:
		out.bigrat.Quo(&r.bigrat, &v.bigrat)
	default:
		panic("rat: add invalid type")
	}
	return out
}

// func (r *Rational) PowInt(n int) *Rational {
// 	out := r.Clone()
// 	if n == 0 {
// 		out.Set(parseInt(1))
// 		return out
// 	} else if n > 0 {
// 		return out.PowInt(n - 1).Mul(r)
// 	} else {
// 		slog.Debug("rat: not implemented")
// 		return nil
// 	}
// }

func (r *Rational) String() string {
	n, exact := r.bigrat.FloatPrec()
	if exact {
		return r.bigrat.FloatString(min(r.precision, n))
		// return r.bigrat.FloatString(min(DefaultPrecision, n))
	}
	return r.bigrat.FloatString(r.precision)
}

func (r *Rational) Set(v *Rational) *Rational {
	r.bigrat.Set(&v.bigrat)
	return r
}

func (r *Rational) Clone() *Rational {
	newr := new(Rational)
	newr.bigrat.Set(&r.bigrat)
	newr.precision = r.precision
	return newr
}

func (r *Rational) IsLessThan(in any) bool {
	inrat := new(Rational)
	switch v := in.(type) {
	case int:
		inrat = Rat(v)
	case int8:
		inrat = Rat(v)
	case int32:
		inrat = Rat(v)
	case int64:
		inrat = Rat(v)
	case float32:
		inrat = Rat(v)
	case float64:
		inrat = Rat(v)
	case string:
		inrat = Rat(v)
	case *Rational:
		inrat = v
	}

	if r.bigrat.Cmp(&inrat.bigrat) == -1 {
		return true
	}
	return false
}

func (r *Rational) IsGreaterThan(in any) bool {
	inrat := new(Rational)
	switch v := in.(type) {
	case int:
		inrat = Rat(v)
	case int8:
		inrat = Rat(v)
	case int32:
		inrat = Rat(v)
	case int64:
		inrat = Rat(v)
	case float32:
		inrat = Rat(v)
	case float64:
		inrat = Rat(v)
	case string:
		inrat = Rat(v)
	case *Rational:
		inrat = v
	}

	if r.bigrat.Cmp(&inrat.bigrat) == 1 {
		return true
	}
	return false
}

func (r *Rational) IsEqual(b *Rational) bool {
	if r.bigrat.Cmp(&b.bigrat) == 0 {
		return true
	}
	return false
}

func (r *Rational) SetPrecision(v int) *Rational {
	r.precision = v
	return r
}

func RatQuo(a *Rational, b *Rational) *Rational {
	return a.Quo(b)
}

func RatMul(a *Rational, b *Rational) *Rational {
	return a.Mul(b)
}

func RatAdd(a *Rational, b *Rational) *Rational {
	return a.Add(b)
}

func RatNeg(a *Rational) *Rational {
	c := big.Rat{}
	c.Neg(&a.bigrat)
	return &Rational{
		bigrat:    c,
		precision: a.precision,
	}
}

func RatZero() *Rational {
	return &Rational{
		precision: DefaultPrecision,
	}
}

func RatClone(r *Rational) *Rational {
	newr := new(Rational)
	newr.bigrat.Set(&r.bigrat)
	newr.precision = r.precision
	return newr
}

func parseFloat64(a float64) *Rational {
	out := new(Rational)
	out.bigrat.SetFloat64(a)
	out.precision = DefaultPrecision
	return out
}

func parseInt64(a int64) *Rational {
	out := new(Rational)
	out.bigrat.SetInt64(a)
	out.precision = DefaultPrecision
	return out
}

func parse(v string) (out *Rational) {
	defer func() {
		if out == nil {
			slog.Error("rat: parse failed, out is nil")
		}
	}()

	if strings.HasSuffix(v, "%") {
		defer func() {
			if out == nil {
				return
			}
			out = out.Quo(100)
		}()
		v = v[0 : len(v)-1]
	}

	if strings.Contains(v, "/") {
		split := strings.Split(v, "/")
		if len(split) != 2 {
			panic("rat: invalid rat string " + v)
		}

		return Rat(split[0]).Quo(Rat(split[1]))
	}

	out = new(Rational)
	_, ok := out.bigrat.SetString(v)
	if !ok {
		return nil
	}
	out.precision = DefaultPrecision
	return out
}

func (r *Rational) Scan(src any) error {
	r.precision = DefaultPrecision
	switch v := src.(type) {
	case string:
		r.Set(Rat(v))
		return nil
	case []byte:
		r.Set(Rat(string(v)))
		return nil
	case int32:
		r.Set(Rat(v))
		return nil
	case int64:
		r.Set(Rat(v))
		return nil
	case float32:
		r.Set(Rat(v))
		return nil
	case float64:
		r.Set(Rat(v))
		return nil
	default:
		return fmt.Errorf("rat: scan err: invalid type %T", src)
	}
}

func (r *Rational) Value() (driver.Value, error) {
	return r.bigrat.RatString(), nil
}

func (r *Rational) GobEncode() ([]byte, error) {
	return r.bigrat.GobEncode()
}

func (r *Rational) GobDecode(buf []byte) error {
	r.precision = DefaultPrecision
	return r.bigrat.GobDecode(buf)
}
