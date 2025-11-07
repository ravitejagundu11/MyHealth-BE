package main

import "time"
import "fmt"

type appointment struct {
    appointmentId   int
    patientId       int
    doctorId        int
    appointmentTime time.Time 
    appointmentStatus string
}

// newAppointment is a constructor function that creates and returns
// a pointer to an appointment instance.
func newAppointment(
    appointmentId int, 
    patientId int, 
    doctorId int, 
    appointmentTime time.Time, 
    appointmentStatus string) *appointment {
        
    return &appointment{
        appointmentId:   appointmentId,
        patientId:       patientId,
        doctorId:        doctorId,
        appointmentTime: appointmentTime, 
        appointmentStatus: appointmentStatus,
    }
}

// Example functions that could be implemented later

// cancelAppointment checks the status and cancels the appointment.
func (c *appointment) cancelAppointment() bool {
    // Logic: Check c.appointmentStatus, update it to "Canceled" or similar.
    if c.appointmentStatus != "Canceled" {
        c.appointmentStatus = "Canceled"
        fmt.Printf("Appointment ID %d for Patient %d has been CANCELED.\n", c.appointmentId, c.patientId)
        return true
    }
    return false
}

// rescheduleAppointment updates the time and status of the appointment.
func (c *appointment) rescheduleAppointment(newAppointmentTime time.Time) bool {
    // Logic: Update the time and maybe set status to "Rescheduled".
    c.appointmentTime = newAppointmentTime
    c.appointmentStatus = "Rescheduled"
    fmt.Printf("Appointment ID %d RESCHEDULED to %v.\n", c.appointmentId, newAppointmentTime.Format("2025-11-15 13:30"))
    return true
}