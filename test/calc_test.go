package test

import (
	"math"
	"testing"

	"github.com/fatih/color"

	"SmartCalc_v3.0/internal/model"
)

func TestCalc(t *testing.T) {
	want := []float64{
		1e+10 - 2 - math.Pow(3, 3)/2 - 6 + 2 + 3 - 2/math.Pow(3, 3) - 6 + 2 + 3 -
			2/math.Pow(3, 3),
		1e-10 + 2 - math.Pow(3, 3)/2 - 6 + 2 + 3 - 2/math.Pow(3, 3) - 6 + 2 + 3 -
			2/math.Pow(3, 3),
		2 - math.Pow(3, 3)/2 - 6 + 2 + 3 - 2/math.Pow(3, 3) - 6 + 2 + 3 -
			2/math.Pow(3, 3),
		5*math.Pow(3, 3)*2 - 6,
		2/math.Pow(3, 3) + 2 - 6,
		2 + math.Pow(3, 3)/2 - 6,
		2/math.Pow(3, 3)*2 - 6,
		2*math.Pow(3, 3)/2 - 6,
		2 + math.Pow(3, 3) + 2 - 6,
		2*math.Pow(3, 3) + 2 - 6,
		2 + math.Pow(3, 3)*2 - 6,
		-(7 + (4 + math.Cos(1))),
		12345678 - 0.12345678,
		-(math.Cos(math.Cos(3))),
		math.Sin(1) + math.Cos(1) - 3*-2*math.Mod(30, 2) + math.Sin(1),
		2 + math.Pow(3, 3)*2,
		1 + math.Cos(2) + (5 - 10),
		math.Cos(4) + math.Sin(5),
		-(2 + 2),
		(1 - 2) + 4,
		2 * (-2 - 2),
		-(9 - 2),
		-1 - 3,
		1 - 3,
		123.3213456,
		math.Sin(1) + math.Cos(1) - 3 + 2*-2 + math.Pow(3, 3) - math.Log(10),
		math.Tan(7),
		math.Asin(1),
		math.Atan(2),
		math.Log10(2),
		math.Sqrt(math.Sqrt(math.Sqrt(1000000000))),
		2 + (2 + (2 + 2)),
		-(2 + (4 - (5 - 6 + (3 + (7 + (9 - 8 - (4))))))),
		0}

	val := []string{
		"1e+10-2-3^3/2-6+2+3-2/3^3-6+2+3-2/3^3",
		"1e-10+2-3^3/2-6+2+3-2/3^3-6+2+3-2/3^3",
		"2-3^3/2-6+2+3-2/3^3-6+2+3-2/3^3",
		"5*3^3*2-6",
		"2/3^3+2-6",
		"2+3^3/2-6",
		"2/3^3*2-6",
		"2*3^3/2-6",
		"2+3^3+2-6",
		"2*3^3+2-6",
		"2+3^3*2-6",
		"-(7+(4+c(1)))",
		"12345678-0.12345678",
		"-(c(c(3)))",
		"s1+c1-3*-2*(30%2)+s1",
		"2+3^3*2",
		"1+c2+(5-10)",
		"c4+s5",
		"-(2+2)",
		"(1-2)+4",
		"2*(-2-2)",
		"-(9-2)",
		"-1-3",
		"1-3",
		"123.3213456",
		"s1+c1-3+2*-2+3^3-l10",
		"t7",
		"S1",
		"T2",
		"L2",
		"q(q(q(1000000000)))",
		"2+(2+(2+2))",
		"-(2+(4-(5-6+(3+(7+(9-8-(4)))))))",
		"0",
	}

	calc := model.NewModel("../internal/model/so/calc.so")
	for i := len(val) - 1; i >= 0; i-- {
		got, err := calc.Calc(&val[i], "")
		if got != want[i] || err != nil {
			errorMsg := color.HiMagentaString(
				"calc.Calc(s *string, x string) (float64, error) = %.8f, %v\n want %.8f, <nil>\n%s",
				got, err, want[i], val[i])
			t.Errorf(errorMsg)
		}
	}
}

func TestCalcError(t *testing.T) {
	val := []string{
		"1a", "(((1a2))", "C3s", "^Cc", "1e*10", "^3",
		"+++4-----", "4+++4", "-(-(-(-(-(3+)))))", "",
	}

	calc := model.NewModel("../internal/model/so/calc.so")
	for i := len(val) - 1; i >= 0; i-- {
		got, err := calc.Calc(&val[i], "")
		if got != 0.0 || err == nil {
			errorMsg := color.HiMagentaString(
				"calc.Calc(s *string, x string) (float64, error) = %.8f, %v\n want 0.0, error\n%s",
				got, err, val[i])
			t.Errorf(errorMsg)
		}
	}
}
