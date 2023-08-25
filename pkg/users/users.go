// Package users provides a client wrapping the WorkOS User Management API.
package users

import (
	"context"
	"net/http"
)

var (
	// DefaultClient is the client used by User management methods
	DefaultClient = NewClient("")
)

// Client represents a client that fetch User Management data from WorkOS API.
type Client struct {
	// The WorkOS api key. It can be found in
	// https://dashboard.workos.com/api-keys.
	//
	// REQUIRED.
	APIKey string

	// The http.Client that is used to send request to WorkOS.
	//
	// Defaults to http.Client.
	HTTPClient *http.Client

	// The endpoint to WorkOS API.
	//
	// Defaults to https://api.workos.com.
	Endpoint string

	// The function used to encode in JSON. Defaults to json.Marshal.
	JSONEncode func(v interface{}) ([]byte, error)
}

// SetAPIKey configures the default client that is used by the User management methods
// It must be called before using those functions.
func SetAPIKey(apiKey string) {
	DefaultClient.APIKey = apiKey
}

// GetUser gets a User.
func GetUser(
	ctx context.Context,
	opts GetUserOpts,
) (User, error) {
	return DefaultClient.GetUser(ctx, opts)
}

// ListUsers gets a list of Users.
func ListUsers(
	ctx context.Context,
	opts ListUsersOpts,
) (ListUsersResponse, error) {
	return DefaultClient.ListUsers(ctx, opts)
}

// CreateUser creates a User.
func CreateUser(
	ctx context.Context,
	opts CreateUserOpts,
) (User, error) {
	return DefaultClient.CreateUser(ctx, opts)
}

// UpdateUser creates a User.
func UpdateUser(
	ctx context.Context,
	opts UpdateUserOpts,
) (User, error) {
	return DefaultClient.UpdateUser(ctx, opts)
}

// UpdateUserPassword updates a User password.
func UpdateUserPassword(
	ctx context.Context,
	opts UpdateUserPasswordOpts,
) (User, error) {
	return DefaultClient.UpdateUserPassword(ctx, opts)
}

// DeleteUser deletes a existing User.
func DeleteUser(
	ctx context.Context,
	opts DeleteUserOpts,
) error {
	return DefaultClient.DeleteUser(ctx, opts)
}

// AddUserToOrganization adds an unmanaged User as a member of the given Organization.
func AddUserToOrganization(
	ctx context.Context,
	opts AddUserToOrganizationOpts,
) (User, error) {
	return DefaultClient.AddUserToOrganization(ctx, opts)
}

// RemoveUserFromOrganization removes an unmanaged User as a member of the given Organization.
func RemoveUserFromOrganization(
	ctx context.Context,
	opts RemoveUserFromOrganizationOpts,
) (User, error) {
	return DefaultClient.RemoveUserFromOrganization(ctx, opts)
}

// AuthenticateUserWithPassword authenticates a user with email and password and optionally creates a session.
func AuthenticateUserWithPassword(
	ctx context.Context,
	opts AuthenticateUserWithPasswordOpts,
) (AuthenticationResponse, error) {
	return DefaultClient.AuthenticateUserWithPassword(ctx, opts)
}

// AuthenticateUserWithCode authenticates an OAuth user or a managed SSO user that is logging in through SSO, and
// optionally creates a session.
func AuthenticateUserWithCode(
	ctx context.Context,
	opts AuthenticateUserWithCodeOpts,
) (AuthenticationResponse, error) {
	return DefaultClient.AuthenticateUserWithCode(ctx, opts)
}

// AuthenticateUserWithMagicAuth authenticates a user by verifying a one-time code sent to the user's email address by
// the Magic Auth Send Code endpoint.
func AuthenticateUserWithMagicAuth(
	ctx context.Context,
	opts AuthenticateUserWithMagicAuthOpts,
) (AuthenticationResponse, error) {
	return DefaultClient.AuthenticateUserWithMagicAuth(ctx, opts)
}

// CreateEmailVerificationChallenge creates an email verification challenge and emails verification token to user.
func CreateEmailVerificationChallenge(
	ctx context.Context,
	opts CreateEmailVerificationChallengeOpts,
) (ChallengeResponse, error) {
	return DefaultClient.CreateEmailVerificationChallenge(ctx, opts)
}

// CompleteEmailVerification verifies user email using verification token that was sent to the user.
func CompleteEmailVerification(
	ctx context.Context,
	opts CompleteEmailVerificationOpts,
) (User, error) {
	return DefaultClient.CompleteEmailVerification(ctx, opts)
}

// CreatePasswordResetChallenge creates a password reset challenge and emails a password reset link to an unmanaged user.
func CreatePasswordResetChallenge(
	ctx context.Context,
	opts CreatePasswordResetChallengeOpts,
) (ChallengeResponse, error) {
	return DefaultClient.CreatePasswordResetChallenge(ctx, opts)
}

// CompletePasswordReset resets user password using token that was sent to the user.
func CompletePasswordReset(
	ctx context.Context,
	opts CompletePasswordResetOpts,
) (User, error) {
	return DefaultClient.CompletePasswordReset(ctx, opts)
}

// SendMagicAuthCode sends a one-time code to the user's email address.
func SendMagicAuthCode(
	ctx context.Context,
	opts SendMagicAuthCodeOpts,
) (User, error) {
	return DefaultClient.SendMagicAuthCode(ctx, opts)
}
