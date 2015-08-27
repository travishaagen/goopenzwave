package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
// #include "notification.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Manager struct {
	manager       C.manager_t
	Notifications chan Notification
}

func CreateManager() *Manager {
	m := &Manager{}
	m.manager = C.manager_create()
	m.Notifications = make(chan Notification, 10)
	return m
}

func Get() *Manager {
	m := &Manager{}
	m.manager = C.manager_get()
	return m
}

func DestroyManager() {
	C.manager_destroy()
}

func GetManagerVersionAsString() string {
	return C.GoString(C.manager_getVersionAsString())
}

func GetManagerVersionLongAsString() string {
	return C.GoString(C.manager_getVersionLongAsString())
}

func (m *Manager) AddDriver(controllerPath string) bool {
	cControllerPath := C.CString(controllerPath)
	result := C.manager_addDriver(m.manager, cControllerPath)
	C.free(unsafe.Pointer(cControllerPath))
	if result {
		return true
	}
	return false
}

func (m *Manager) RemoveDriver(controllerPath string) bool {
	cControllerPath := C.CString(controllerPath)
	result := C.manager_removeDriver(m.manager, cControllerPath)
	C.free(unsafe.Pointer(cControllerPath))
	if result {
		return true
	}
	return false
}

// Notification and callbacks from C:
// http://stackoverflow.com/questions/6125683/call-go-functions-from-c

// // This defines the signature of our user's progress handler.
// type NotificationHandler func(notification *Notification, userdata interface{})
//
// // This is an internal type which will pack the users callback function and
// // userdata. It is an instance of this type that we will actually be sending to
// // the C code.
// type notificationContainer struct {
// 	f NotificationHandler // The user's function pointer.
// 	d interface{}         // The user's userdata.
// }

//export goNotificationCB
func goNotificationCB(notification C.notification_t, userdata unsafe.Pointer) {
	// This is the function called from the C world by the OpenZWave
	// notification system. The userdata value contains an instance of
	// *notificationContainer, We unpack it and use it's values to call the
	// actual function that our user supplied.
	m := (*Manager)(userdata)

	// Convert the C notification_t to Go Notification.
	noti := buildNotification(notification)

	// Send the Notification on the channel.
	m.Notifications <- noti
}

func (m *Manager) StartNotifications() error {
	themanager := unsafe.Pointer(m)
	result := C.manager_addWatcher(m.manager, themanager)
	if result {
		return nil
	}
	return fmt.Errorf("failed to add watcher")
}

func (m *Manager) StopNotifications() error {
	themanager := unsafe.Pointer(m)
	result := C.manager_removeWatcher(m.manager, themanager)
	if result {
		return nil
	}
	return fmt.Errorf("failed to remove watcher")
}

// func (m *Manager) AddWatcher(nh NotificationHandler, userdata interface{}) (unsafe.Pointer, bool) {
// 	watcher := unsafe.Pointer(&notificationContainer{nh, userdata})
// 	result := C.manager_addWatcher(m.manager, watcher)
// 	if result {
// 		return watcher, true
// 	}
// 	return nil, false
// }
//
// func (m *Manager) RemoveWatcher(watcher unsafe.Pointer) bool {
// 	result := C.manager_removeWatcher(m.manager, watcher)
// 	if result {
// 		return true
// 	}
// 	return false
// }