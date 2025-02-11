package credential

type user struct {
	userName    string
	domainName  string
	workstation string
}

// User name.
func (n *user) UserName() string {
	if n != nil {
		return n.userName
	}
	return ""
}

// Domain name.
func (n *user) DomainName() string {
	if n != nil {
		return n.domainName
	}
	return ""
}

// Workstation.
func (p *user) Workstation() string {
	if p != nil {
		return p.workstation
	}
	return ""
}
