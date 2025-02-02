package dcom

func (o *ClassID) Equal(other *ClassID) bool {
	return o.GUID().Equal(other.GUID())
}
