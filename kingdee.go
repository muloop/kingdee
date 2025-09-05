package kingdee

type KingDee struct {
	formId string
	c      Config
}

func New(c Config) *KingDee {
	kingDee := &KingDee{
		c: c,
	}

	return kingDee
}
func (k *KingDee) FormId(formId string) *KingDee {
	k.formId = formId
	return k
}

func (k *KingDee) View() *KingDee {
	return k
}
func (k *KingDee) Draft() {

}
func (k *KingDee) Save() {

}
func (k *KingDee) Submit() {

}
func (k *KingDee) Audit() {

}
func (k *KingDee) Allocate() {

}
func (k *KingDee) CancelAllocate() {

}

func (k *KingDee) ExecuteBillQuery() {

}

func (k *KingDee) Result() {

}
