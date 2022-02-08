package models

type DbBlock struct {
	Try     func()
	Catch   func(interface{})
	Finally func()
}

func (block *DbBlock) DO() {
	if block.Finally != nil {
		defer block.Finally()
	}
	if block.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				block.Catch(r)
			}
		}()
	}
	block.Try()
}
