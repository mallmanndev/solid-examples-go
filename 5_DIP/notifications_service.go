package main

import "fmt"

type NotificationsService interface {
	SendNotification(message string) error
}

type NotificationsServiceImpl struct{}

func NewNotificationsService() *NotificationsServiceImpl {
	return &NotificationsServiceImpl{}
}

func (ns *NotificationsServiceImpl) SendNotification(message string) error {
	fmt.Printf("Send notification message: %s...\n", message)
	return nil
}
