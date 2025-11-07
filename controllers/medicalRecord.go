package main

import (
    "fmt"
    "time"
)

type medicalRecord struct {
    recordId      int
    patientId     int
    doctorId      int
    appointmentId int
    diagnosis     string 
    treatment     string 
    medications   string 
    updatedTime   time.Time
}

// newMedicalRecord is a constructor function that creates and returns
// a pointer to a medicalRecord instance.
func newMedicalRecord(
    recordId int,
    patientId int,
    doctorId int,
    appointmentId int,
    diagnosis string,     
    treatment string,     
    medications string,   
    updatedTime time.Time) *medicalRecord {

    return &medicalRecord{
        recordId:       recordId,
        patientId:      patientId,
        doctorId:       doctorId,
        appointmentId:  appointmentId,
        diagnosis:      diagnosis,    
        treatment:      treatment,   
        medications:    medications, 
        updatedTime:    updatedTime,
    }
}

// viewRecord retrieves the medical record details.
func (c *medicalRecord) viewRecord() {
    // In a real app, this might fetch data from a database, but here we just display the current data.
    fmt.Println("--- Medical Record Details ---")
    fmt.Printf("Record ID: %d (Appointment ID: %d)\n", c.recordId, c.appointmentId)
    fmt.Printf("Patient ID: %d | Doctor ID: %d\n", c.patientId, c.doctorId)
    fmt.Printf("Diagnosis: %s\n", c.diagnosis)
    fmt.Printf("Treatment: %s\n", c.treatment)
    fmt.Printf("Medications: %s\n", c.medications)
    fmt.Printf("Last Updated: %s\n", c.updatedTime.Format("2025-11-15 14:00"))
}