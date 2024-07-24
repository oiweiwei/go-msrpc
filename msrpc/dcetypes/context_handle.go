package dcetypes

func (o *ContextHandle) IsZero() bool {
	return o.Attributes == 0 && o.UUID.IsZero()
}
