package rat

import (
	"encoding/json"
	"testing"
)

func Test10(t *testing.T) {
	a := Rat(1347)
	b := a.Quo(10).Floor().Mul(10)
	if b.String() != "1340" {
		t.Fatalf("expected 1340, got %s", b.String())
	}

	if a.String() != "1347" {
		t.Fatalf("expected 1347, got %s", a.String())
	}
}

func TestJson(t *testing.T) {
	var msg = `{ "balance": "1386929.37231066771348207123", "currency": "KRW" }`

	type Balance struct {
		Balance  *Rational `json:"balance"`
		Currency string    `json:"currency"`
	}

	v := &Balance{}
	err := json.Unmarshal([]byte(msg), v)
	if err != nil {
		t.Fatal(err)
	}

	v.Balance.precision = 20

	if v.Balance.String() != "1386929.37231066771348207123" {
		t.Fatalf("expected 1386929.37231066771348207123, got %s", v.Balance.String())
	}
}

func TestSub(t *testing.T) {
	a := Rat("2")
	if !a.Sub(-2).IsEqual(Rat(4)) {
		t.Fatalf("expected 4, got %s", a.Sub(-2))
	}

	assertEqual(t, a.String(), "2")
}

func TestParseFloat(t *testing.T) {
	// got := parseFloat64(1 / 100)
	got := parseFloat64(0.01)
	got.precision = 2
	if got.String() != "0.01" {
		t.Fatalf("expected 0.01, got %s", got.String())
	}
}

func TestAddNeg(t *testing.T) {
	got := parse("7.004")
	got = got.Add(Rat("0.001").Neg())
	if got.String() != "7.003" {
		t.Fatalf("expected 7.003, got %s", got.String())
	}
}

func TestCeil(t *testing.T) {
	{
		got := parse("1382.5")
		if got.FloorInt() != 1382 {
			t.Fatalf("expected 1382, got %d", got.FloorInt())
		}
	}
	{
		got := parse("-7.004")
		if !got.Floor().IsEqual(Rat(-8)) {
			t.Fatal("floor failed", got.Floor())
		}
	}
	{
		got := parse("-7.004")
		if !got.Ceil().IsEqual(Rat(-7)) {
			t.Fatalf("ceil failed")
		}
	}

	{
		if !parse("-7").Floor().IsEqual(Rat(-7)) {
			t.Fatalf("floor failed")
		}
		if !parse("-7").Ceil().IsEqual(Rat(-7)) {
			t.Fatalf("floor failed")
		}
	}

	{
		if !parse("7").Floor().IsEqual(Rat(7)) {
			t.Fatalf("floor failed")
		}
		if !parse("7.5").Floor().IsEqual(Rat(7)) {
			t.Fatalf("floor failed")
		}
		if !parse("7.4").Floor().IsEqual(Rat(7)) {
			t.Fatalf("floor failed")
		}
		if !parse("7.6").Floor().IsEqual(Rat(7)) {
			t.Fatalf("floor failed")
		}
	}

	{
		if !parse("-7").Floor().IsEqual(Rat(-7)) {
			t.Fatalf("floor failed")
		}
		if !parse("-7.5").Floor().IsEqual(Rat(-8)) {
			t.Fatalf("floor failed")
		}
		if !parse("-7.4").Floor().IsEqual(Rat(-8)) {
			t.Fatalf("floor failed")
		}
		if !parse("-7.6").Floor().IsEqual(Rat(-8)) {
			t.Fatalf("floor failed")
		}
	}

	{
		if !parse("7").Ceil().IsEqual(Rat(7)) {
			t.Fatalf("ceil failed")
		}
		if !parse("7.5").Ceil().IsEqual(Rat(8)) {
			t.Fatalf("ceil failed")
		}
		if !parse("7.4").Ceil().IsEqual(Rat(8)) {
			t.Fatalf("ceil failed")
		}
		if !parse("7.6").Ceil().IsEqual(Rat(8)) {
			t.Fatalf("ceil failed")
		}
	}

	{
		if !parse("-7").Ceil().IsEqual(Rat(-7)) {
			t.Fatalf("ceil failed")
		}
		if !parse("-7.5").Ceil().IsEqual(Rat(-7)) {
			t.Fatalf("ceil failed")
		}
		if !parse("-7.4").Ceil().IsEqual(Rat(-7)) {
			t.Fatalf("ceil failed")
		}
		if !parse("-7.6").Ceil().IsEqual(Rat(-7)) {
			t.Fatalf("ceil failed")
		}
	}

	{
		got := parse("0.95")
		if got.Ceil().IsEqual(Rat(1)) != true {
			t.Fatalf("expected 1382, got %v", got.Floor())
		}
	}
	{
		got := parse("4")
		if got.Ceil().IsEqual(Rat(4)) != true {
			t.Fatalf("expected -6, got %v", got.Floor())
		}
	}
	{
		got := parse("7.004")
		if got.Ceil().IsEqual(Rat(8)) != true {
			t.Fatalf("expected -6, got %v", got.Floor())
		}
	}
	{
		got := parse("-7.004")
		if got.Ceil().IsEqual(Rat(-7)) != true {
			t.Fatalf("expected -6, got %v", got.Floor())
		}
	}
}

func TestFloor(t *testing.T) {
	{
		got := parse("1382.5")
		if got.FloorInt() != 1382 {
			t.Fatalf("expected 1382, got %d", got.FloorInt())
		}
	}
	{
		got := parse("-5.05")
		if got.FloorInt() != -6 {
			t.Fatalf("expected -6, got %d", got.FloorInt())
		}
	}
	{
		got := parse("1382.5")
		if got.Floor().IsEqual(Rat(1382)) != true {
			t.Fatalf("expected 1382, got %v", got.Floor())
		}
	}
	{
		got := parse("-5.05")
		if got.Floor().IsEqual(Rat(-6)) != true {
			t.Fatalf("expected -6, got %v", got.Floor())
		}
	}
}

func TestNeg(t *testing.T) {
	a := parse("1382")
	if a.Neg().String() != "-1382" {
		t.Fatalf("expected -1382, got %s", a.Neg().String())
	}

	assertEqual(t, a.String(), "1382")
}

func TestMin(t *testing.T) {
	a := parse("1382")
	b := parse("1380")
	c := parse("1381")

	if RatMin(a, b, c).String() != "1380" {
		t.Fatalf("expected 1380, got %s", RatMin(a, b, c).String())
	}
	if a.String() != "1382" {
		t.Fatalf("expected 1380, got %s", a.String())
	}

	if b.String() != "1380" {
		t.Fatalf("expected 1381, got %s", b.String())
	}
	if c.String() != "1381" {
		t.Fatalf("expected 1382, got %s", c.String())
	}
}

func TestParcentage(t *testing.T) {
	a := Rat(10000).Mul(Rat("3%"))
	assertEqual(t, a.String(), "300")
}

func TestBasics(t *testing.T) {
	a := Rat("2")
	b := Rat("3")
	astr := a.String()
	bstr := b.String()

	b.Neg() // no-op
	a.Neg() // no-op

	assertEqual(t, "15", a.Add(b).Mul(b).IntString())
	assertEqual(t, "1.5", b.Quo(a).String())

	assertEqual(t, astr, a.String())
	assertEqual(t, bstr, b.String())
}

func TestQuo(t *testing.T) {
	a := parse("2")
	b := parse("4")
	if RatQuo(a, b).String() != "0.5" {
		t.Fatal("Quo")
	}
}

func TestString(t *testing.T) {
	{
		a := parse("0.5/1")
		if a.String() != "0.5" {
			t.Fatal(a.String())
		}

		if _, exact := a.bigrat.FloatPrec(); !exact {
			t.Fatal("not exact")
		}
	}

	{
		a := parse("1/0.5")
		if a.String() != "2" {
			t.Fatal(a.String())
		}
		if _, exact := a.bigrat.FloatPrec(); !exact {
			t.Fatal("not exact")
		}
	}

	{
		a := parse("1/3")
		a.precision = 8
		if a.String() != "0.33333333" {
			t.Fatal(a.String())
		}
		if _, exact := a.bigrat.FloatPrec(); exact {
			t.Fatal("should not be exact")
		}
	}
}

func TestCmp(t *testing.T) {
	a := parse("1/2")
	b := parse("1/3")
	if a.IsLessThan(b) {
		t.Fatal()
	}
}

func assertEqual[T comparable](t *testing.T, a T, b T) {
	if a != b {
		t.Fatalf("assert fail %v %v", a, b)
	}
}
