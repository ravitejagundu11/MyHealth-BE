
type MainController struct {
}

func NewMainController() *MainController {
    return &MainController{}
}

// GetDoctors gets doctor
func (c *MainController) GetDoctors() string {
    return ""
}

// GetAvailableDates gets all available dates for doctor
func (c *MainController) GetAvailableDates(doctorId string) string {
    return ""
}

// GetAvailableTimes gets all available dates for doctor
func (c *MainController) GetAvailableTimes() string {
    return ""
}

// GetDoctors gets doctor
func (c *MainController) GetPatients() string {
    return ""
}