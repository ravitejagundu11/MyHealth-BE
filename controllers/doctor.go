type doctor struct{
	doctorId string
	specialization string
	availability []string
}

//Newdoctor: a constructor for doctor
func Newdoctor(doctorId string, specialization string, availability string) *doctor{
	return &doctor{}
}

//AddAvailability: To add doctor's availability
func(c *doctor) AddAvailability(availability string){
}

//ViewAppointment: To view Appointments booked
func ViewAppointment(){
	// TODO: return appointmnet struct created in appointments controller
	// Update return type
}

//AddMedicalRecord: To add medical records of a patient
func AddMedicalRecord(patientId string, recordDetails string){
	// TODO: Return medical record struct created in medical records controller
	// Update return type
}