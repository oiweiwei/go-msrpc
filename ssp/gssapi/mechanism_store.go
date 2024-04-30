package gssapi

import (
	"context"
	"fmt"
	"sync"
)

type MechanismStore struct {
	mu    sync.Mutex
	mechs []MechanismFactory
}

// AddMechanism function appends the mechanism to the mechanism store.
func (m *MechanismStore) AddMechanism(f MechanismFactory) {

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, mech := range m.mechs {
		if f.Type().Equal(mech.Type()) {
			panic(fmt.Sprintf("mechanism %s already exist", f.Type()))
		}
	}

	m.mechs = append(m.mechs, f)
}

// GetMechanism function returns the mechanism for the selected OID.
func (m *MechanismStore) GetMechanism(oid OID) MechanismFactory {

	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.mechs) == 0 {
		return nil
	}

	for _, mech := range m.mechs {
		if mech.Type().Equal(oid) {
			return mech
		}
	}

	return m.mechs[0]
}

func (m *MechanismStore) ListMechanisms() []MechanismFactory {

	m.mu.Lock()
	defer m.mu.Unlock()

	ret := make([]MechanismFactory, len(m.mechs))

	for i := range m.mechs {
		ret[i] = m.mechs[i]
	}

	return ret
}

var (
	defaultMechanismStore = new(MechanismStore)
)

// AddMechanism function appends the mechanism to the mechanism store.
func AddMechanism(f MechanismFactory) {
	defaultMechanismStore.AddMechanism(f)
}

// GetMechanism function returns the mechanism for the selected OID.
func GetMechanism(ctx context.Context, oid OID) MechanismFactory {
	cc := fromContext(ctx)
	if cc != nil {
		if cc.MechanismStore != nil {
			return cc.MechanismStore.GetMechanism(oid)
		}
	}
	return defaultMechanismStore.GetMechanism(oid)
}

func ListMechanisms(ctx context.Context) []MechanismFactory {
	cc := fromContext(ctx)
	if cc != nil && cc.MechanismStore != nil {
		return cc.MechanismStore.ListMechanisms()
	}
	return defaultMechanismStore.ListMechanisms()
}
