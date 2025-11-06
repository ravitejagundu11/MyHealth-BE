package models

import "time"

// User represents a system user (patient or doctor)
type user struct {
    userID         string    `json:"userId"`
    name           string    `json:"name"`
    dateOfBirth    time.Time `json:"dateOfBirth"`
    age            int       `json:"age"`
    contactInfo    string    `json:"contactInfo"`
    medicalHistory string    `json:"medicalHistory"`
}

// updateProfile updates the user's contact information
func (u *user) updateProfile(newContact string) {
    u.contactInfo = newContact
}

// viewMedicalRecords returns a list of medical records associated with the user
func (u *user) viewMedicalRecords() []medicalRecord {
    // This is just a placeholder — actual implementation may query a database
    return []medicalRecord{}
}

