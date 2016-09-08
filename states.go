package gopwa

type AuthorisationState string

const (
	AuthStatePending  AuthorisationState = "Pending"
	AuthStateOpen     AuthorisationState = "Open"
	AuthStateDeclined AuthorisationState = "Declined"
	AuthStateClosed   AuthorisationState = "Closed"
)

type BillingAgreementState string

const (
	BillAgreeStateDraft     BillingAgreementState = "Draft"
	BillAgreeStateOpen      BillingAgreementState = "Open"
	BillAgreeStateSuspended BillingAgreementState = "Suspended"
	BillAgreeStateCanceled  BillingAgreementState = "Canceled"
	BillAgreeStateClosed    BillingAgreementState = "Closed"
)

type CaptureState string

const (
	CaptureStatePending   CaptureState = "Pending"
	CaptureStateDeclined  CaptureState = "Declined"
	CaptureStateCompleted CaptureState = "Completed"
	CaptureStateClosed    CaptureState = "Closed"
)

type OrderReferenceState string

const (
	OrderRefStateDraft     OrderReferenceState = "Draft"
	OrderRefStateOpen      OrderReferenceState = "Open"
	OrderRefStateSuspended OrderReferenceState = "Suspended"
	OrderRefStateCanceled  OrderReferenceState = "Canceled"
	OrderRefStateClosed    OrderReferenceState = "Closed"
)

type RefundState string

const (
	RefundStatePending   RefundState = "Pending"
	RefundStateDeclined  RefundState = "Declined"
	RefundStateCompleted RefundState = "Completed"
)
