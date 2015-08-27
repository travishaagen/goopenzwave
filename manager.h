#ifndef GOZWAVE_MANAGER
#define GOZWAVE_MANAGER

#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>
#include "notification.h"

#ifdef __cplusplus
extern "C" {
#endif

	// Types.
	typedef void* manager_t;

	// Construction.
	manager_t manager_create();
	manager_t manager_get();
	void manager_destroy();
	const char* manager_getVersionAsString();
	const char* manager_getVersionLongAsString();
	// static ozwversion getVersion();

	// Configuration.

	// Drivers.
	bool manager_addDriver(manager_t m, const char *controllerPath);
	bool manager_removeDriver(manager_t m, const char *controllerPath);
	//...

	// Polling Z-Wave devices.

	// Node information.

	// Values.

	// Climate control schedules.

	// Switch all.

	// Configuration parameters.

	// Groups.

	// Notifications.
	extern void goNotificationCB(notification_t notification, void *userdata);
	bool manager_addWatcher(manager_t m, void *userdata);
	bool manager_removeWatcher(manager_t m, void *userdata);
	//...

	// Controller commands.

	// Network commands.

	// Scene commands.

	// Statistics retreival interface.

#ifdef __cplusplus
}
#endif

#endif // define GOZWAVE_MANAGER