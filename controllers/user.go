package models

import "time"

// User represents a system user (patient or doctor)
type User struct {
    userID         string
    name           string    
    dateOfBirth    time.Time
    age            int      
    contactInfo    string    
    medicalHistory string 
}

// NewUser is a constructor function that creates and returns a pointer to a User instance.
func newUser(userId, name, dateOfBirth, age, contactInfo, medicalHistory) *User{
	return &User{
		userID:         userId,
		name:           name,
		dateOfBirth:    dateOfBirth,
		age:            age,
		contactInfo:    contactInfo,
		medicalHistory: medicalHistory,
	}
}

// updateProfile updates the user's contact information
func (u *User) updateProfile(newContact string) {
    u.contactInfo = newContact
}

// viewMedicalRecords returns a list of medical records associated with the user
func (u *User) viewMedicalRecords() []medicalRecord {
    // This is just a placeholder — actual implementation may query a database
    return []medicalRecord{}
}

