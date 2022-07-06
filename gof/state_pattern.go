package gof

import "errors"

type Orders struct {
	state string
}

func (b *Orders) Save() {
	b.state = "draft"
}

func (b *Orders) Audit() error {
	if b.state != "draft" {
		return errors.New("草稿状态才可审核")
	}

	b.state = "audit"
	return nil
}

// todo:以订单流转状态来说明
