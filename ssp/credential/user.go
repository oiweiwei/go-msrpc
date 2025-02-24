package credential

type userCred struct {
	userName    string
	domainName  string
	workstation string
}

// User name.
func (n *userCred) UserName() string {
	if n != nil {
		return n.userName
	}
	return ""
}

// Domain name.
func (n *userCred) DomainName() string {
	if n != nil {
		return n.domainName
	}
	return ""
}

// Workstation.
func (p *userCred) Workstation() string {
	if p != nil {
		return p.workstation
	}
	return ""
}
