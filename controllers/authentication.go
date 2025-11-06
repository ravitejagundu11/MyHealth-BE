package models

// Authentication stores login credentials and user role
type Authentication struct {
    userId       string
    userName     string
    passwordHash string
    role         string
}

// NewAuthentication is a constructor function that creates and returns
// a pointer to an Authentication instance.
func newAuthentication(userId, userName, passwordHash, role) *Authentication {
	return &Authentication{
		userId:       userId,
		userName:     userName,
		passwordHash: passwordHash,
		role:         role,
	}
}

// Example functions that could be implemented later
func (a *Authentication) Login(userName, password string) bool {
    // Placeholder logic for authentication
    return userName == a.userName && password == a.passwordHash
}

func (a *Authentication) Register(userName, password, role string) bool {
    // Placeholder for registration logic
    a.userName = userName
    a.passwordHash = password
    a.role = role
    return true
}

func (a *Authentication) Logout() {
    // Example logout logic
}
