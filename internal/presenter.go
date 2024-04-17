package calc

import (
	"errors"
	"fmt"

	"SmartCalc_v3.0/internal/model"
)

type Presenter struct {
	view  *View
	model *model.Model
}

func NewPresenter(v *View, m *model.Model) *Presenter {
	return &Presenter{
		view:  v,
		model: m,
	}
}

func (p *Presenter) evaluate(s *string, x string) {
	p.view.updateNumberLabel(p.calc(s, x))
}

func (p *Presenter) evaluateX(s *string, x string) {
	p.view.updateXLabel(fmt.Sprint(p.calc(s, x)))
}

func (p *Presenter) calc(s *string, x string) string {
	var output string
	if res, err := p.model.Calc(s, x); err != nil {
		output = err.Error()
	} else {
		output = fmt.Sprint(res)
	}

	return output
}

func (p *Presenter) calcPlot(s *string, x string) (float64, error) {
	res, _ := p.model.Calc(s, x)
	if !checkInput(fmt.Sprint(res)) {
		return 0.0, errors.New("error")
	}

	return res, nil
}
