package dtyp

func (o *ACEGUID) MarshalJSON() ([]byte, error) {

	if guid, ok := o.GetValue().(*GUID); ok {
		return guid.MarshalJSON()
	}

	return []byte("null"), nil
}
