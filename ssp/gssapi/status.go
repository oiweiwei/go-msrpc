package gssapi

import (
	"errors"
)

// The security context status.
type Status int

const (
	// GSS_S_NO_CONTEXT: no context. (initial state of the context.)
	NoContext Status = iota
	// GSS_S_COMPLETE: normal completion.
	Complete
	// GSS_S_CONTINUE_NEEDED: continuation call to routine required.
	ContinueNeeded
	// GSS_S_DUPLICATE_TOKEN: duplicate per-message token detected.
	DuplicateToken
	// GSS_S_OLD_TOKEN: timed-out per-message token detected.
	OldToken
	// GSS_S_UNSEQ_TOKEN: reordered (early) per-message token detected.
	UnseqToken
	// GSS_S_GAP_TOKEN: skipped predecessor token(s) detected.
	GapToken
	// GSS_S_BAD_BINDINGS: channel binding mismatch.
	BadBindings
	// GSS_S_BAD_MECH: unsupported mechanism requested.
	BadMech
	// GSS_S_BAD_NAME: invalid name provided.
	BadName
	// GSS_S_BAD_NAMETYPE: name of unsupported type provided.
	BadNameType
	// GSS_S_BAD_STATUS: invalid input status selector.
	BadStatus
	// GSS_S_BAD_SIG: token had invalid integrity check.
	BadSig
	// GSS_S_BAD_MIC: token had invalid integrity check.
	BadMIC
	// GSS_S_CONTEXT_EXPIRED: specified security context expired.
	ContextExpired
	// GSS_S_CREDENTIALS_EXPIRED: expired credentials detected.
	CredentialsExpired
	// GSS_S_DEFECTIVE_CREDENTIAL: defective credential detected.
	DefectiveCredential
	// GSS_S_DEFECTIVE_TOKEN: defective token detected.
	DefectiveToken
	// GSS_S_FAILURE: unknown error.
	Unknown
	// GSS_S_NO_CRED: no valid credentials provided.
	NoCred
	// GSS_S_BAD_QOP: unsupported QoP value.
	BadQoP
	// GSS_S_UNAUTHORIZED: operation unauthorized.
	Unauthorized
	// GSS_S_UNAVAILABLE: operation unavailable.
	Unavailable
	// GSS_S_DUPLICATE_ELEMENT: duplicate credential element requested.
	DuplicateElement
	// GSS_S_NAME_NOT_MN: name contains multi-mechanism elements.
	NameNotMN
	// GSS_S_FAILURE: failure, unspecified at GSS-API level.
	Failure
)

type Error struct {
	Status Status
	Err    error
}

// Error function implements Error interface.
func (e *Error) Error() string {
	return e.Err.Error()
}

func NewError(status Status, err error) error {
	return &Error{status, err}
}

var (
	// GSS_S_BAD_BINDINGS.
	ErrBadBindings = NewError(BadBindings, errors.New("channel binding mismatch"))
	// GSS_S_BAD_MECH.
	ErrBadMech = NewError(BadMech, errors.New("unsupported mechanism requested"))
	// GSS_S_BAD_NAME.
	ErrBadName = NewError(BadName, errors.New("invalid name provided"))
	// GSS_S_BAD_NAMETYPE.
	ErrBadNameType = NewError(BadNameType, errors.New("name of unsupported type provided"))
	// GSS_S_BAD_STATUS.
	ErrBadStatus = NewError(BadStatus, errors.New("invalid input status selector"))
	// GSS_S_BAD_SIG.
	ErrBadSig = NewError(BadSig, errors.New("token had invalid integrity check"))
	// GSS_S_BAD_MIC.
	ErrBadMIC = NewError(BadMIC, errors.New("token had invalid integrity check"))
	// GSS_S_CONTEXT_EXPIRED.
	ErrContextExpired = NewError(ContextExpired, errors.New("specified security context expired"))
	// GSS_S_CREDENTIALS_EXPIRED.
	ErrCredentialsExpired = NewError(CredentialsExpired, errors.New("expired credentials detected"))
	// GSS_S_DEFECTIVE_CREDENTIAL.
	ErrDefectiveCredential = NewError(DefectiveCredential, errors.New("defective credential detected"))
	// GSS_S_DEFECTIVE_TOKEN.
	ErrDefectiveToken = NewError(DefectiveToken, errors.New("defective token detected"))
	// GSS_S_FAILURE.
	ErrUnknown = NewError(Unknown, errors.New("unknown error"))
	// GSS_S_NO_CONTEXT.
	ErrNoContext = NewError(NoContext, errors.New("no valid security context specified"))
	// GSS_S_NO_CRED.
	ErrNoCred = NewError(NoCred, errors.New("no valid credentials provided"))
	// GSS_S_BAD_QOP.
	ErrBadQoP = NewError(BadQoP, errors.New("unsupported QoP value"))
	// GSS_S_UNAUTHORIZED.
	ErrUnauthorized = NewError(Unauthorized, errors.New("operation unauthorized"))
	// GSS_S_UNAVAILABLE.
	ErrUnavailable = NewError(Unavailable, errors.New("operation unavailable"))
	// GSS_S_DUPLICATE_ELEMENT.
	ErrDuplicateElement = NewError(DuplicateElement, errors.New("duplicate credential element requested"))
	// GSS_S_NAME_NOT_MN.
	ErrNameNotMN = NewError(NameNotMN, errors.New("name contains multi-mechanism elements"))
	// GSS_S_FAILURE.
	ErrFailure = NewError(Failure, errors.New("failure, unspecified at GSS-API level"))

	// GSS_S_UNSEQ_TOKEN: reordered (early) per-message token detected.
	ErrUnseqToken = NewError(UnseqToken, errors.New("reordered (early) per-message token detected."))
)
