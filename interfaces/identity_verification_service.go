package interfaces

type IdentityVerificationService interface {
	Verify(ip string, expected string) bool
}
