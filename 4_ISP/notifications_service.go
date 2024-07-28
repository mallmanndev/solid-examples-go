package main

import "fmt"

type NotificationsService struct{}

func NewNotificationsService() *NotificationsService {
	return &NotificationsService{}
}

func (ns *NotificationsService) SendNotification(message string) error {
	fmt.Printf("Send notification message: %s...\n", message)
	return nil
}
