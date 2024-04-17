package calc

import (
	"unicode"
)

func (v *View) updateNumberLabel(text string) {
	v.counter = len(text)
	v.numberLabel.SetText(text)
}

func (v *View) updateXLabel(text string) {
	if !checkInput(text) {
		v.updateNumberLabel(text)
	} else {
		v.xLabel.SetText(text)
		v.updateNumberLabel("0")
	}
}

func (v *View) printOperator(text string) {
	if v.counter < 256 {
		s := v.numberLabel.Text
		if !checkInput(s) {
			s = ""
		}
		v.updateNumberLabel(s + text)
	}
}

func checkInput(s string) bool {
	if s == "error" || s == "-Inf" ||
		s == "+Inf" || s == "NaN" ||
		s == "-NaN" {
		return false
	}

	return true
}

func (v *View) printButton(text string) {
	if v.counter < 256 {
		s := v.numberLabel.Text
		if (text != "-" && text != "(" &&
			text != ")") && s[v.counter-1] == 'x' {
			return
		}

		if s == "0" || !checkInput(s) {
			s = ""
		}
		v.updateNumberLabel(s + text)
	}
}

func (v *View) acButton() {
	v.updateNumberLabel("0")
}

func (v *View) pointerButton() {
	c := v.counter - 1
	r := []rune(v.numberLabel.Text)

	for unicode.IsDigit(r[c]) && c > 0 {
		c--
	}

	if (c != v.counter-1 || c == 0) && r[c] != '.' {
		v.updateNumberLabel(string(r) + ".")
	}
}

func (v *View) delButton() {
	s := v.numberLabel.Text

	r := []rune(s)
	for (unicode.IsLetter(r[v.counter-1]) ||
		r[v.counter-1] == '(') && v.counter > 1 {
		v.counter--
	}

	if v.counter > 1 {
		v.updateNumberLabel(s[:v.counter-1])
	} else {
		v.updateNumberLabel("0")
	}
}

func (v *View) xButton() {
	if v.counter < 256 {
		s := v.numberLabel.Text
		if s[v.counter-1] != 'x' && !unicode.IsDigit(rune(s[v.counter-1])) {
			v.updateNumberLabel(s + "x")
		}
		if s == "0" {
			v.updateNumberLabel("x")
		}
	}
}

func (v *View) xInitButton() {
	s := v.numberLabel.Text
	if checkInput(s) {
		if s != "0" {
			v.writeHistory()
		}
		v.presenter.evaluateX(&s, v.xLabel.Text)
	} else {
		v.updateNumberLabel("0")
	}
}

func (v *View) evaluateButton() {
	s := v.numberLabel.Text
	if !checkInput(s) {
		s = "0"
	}
	if s != "0" {
		v.writeHistory()
	}
	v.presenter.evaluate(&s, v.xLabel.Text)
}

func (v *View) inverseButton() {
	s := v.numberLabel.Text
	if !checkInput(s) {
		s = "0"
	}

	if s != "0" {
		s = "(" + s + ")*(-1)"
	}
	if s != "0" {
		v.writeHistory()
	}
	v.presenter.evaluate(&s, v.xLabel.Text)
}

func (v *View) mrcButton() {
	if v.memory.flag {
		v.updateNumberLabel(v.memory.value)
		v.memory.flag = false
	} else {
		s := v.numberLabel.Text
		res := v.presenter.calc(&s, v.xLabel.Text)
		if checkInput(res) {
			v.memory.value = res
			v.memory.flag = true
		}
	}
}

func (v *View) mMinusButton() {
	if v.memory.flag {
		s := v.numberLabel.Text
		if checkInput(s) {
			s := v.memory.value + "-(" + s + ")"
			res := v.presenter.calc(&s, v.xLabel.Text)
			if checkInput(res) {
				v.memory.value = res
			}
		}
	}
}

func (v *View) mPlusButton() {
	if v.memory.flag {
		s := v.numberLabel.Text
		if checkInput(s) {
			s = v.memory.value + "+(" + s + ")"
			res := v.presenter.calc(&s, v.xLabel.Text)
			if checkInput(res) {
				v.memory.value = res
			}
		}
	}
}
