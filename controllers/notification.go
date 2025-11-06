
type Notification struct {
    notificationId string
    userId string
    message string
    timestamp string
    isRead bool
}

// NewNotification is constructor for Notification.
func NewNotification(notificationId, userId, message, timestamp string, isRead bool) *Notification{
    return &Notification{
        notificationId,
        userId,
        message,
        timestamp,
        isRead
    }
}

// SendNotification will sends a notification to given user with given message.
func(c *Notification) SendNotification(userId , message string) {

}

// MarksAsRead marks messages are read.
func(c *Notification) MarksAsRead() {
}
