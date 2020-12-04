//+build windows
package box

import "golang.org/x/sys/windows"

// Get the Windows Versio and Build Number
var (
	winVersion, _, buildNumber = windows.RtlGetNtVersionNumbers()
)
