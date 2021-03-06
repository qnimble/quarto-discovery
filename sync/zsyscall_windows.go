// Code generated by 'go generate'; DO NOT EDIT.

package sync

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	moduser32   = windows.NewLazySystemDLL("user32.dll")

	procGetModuleHandleA             = modkernel32.NewProc("GetModuleHandleA")
	procCreateWindowExA              = moduser32.NewProc("CreateWindowExA")
	procDefWindowProcW               = moduser32.NewProc("DefWindowProcW")
	procDestroyWindow                = moduser32.NewProc("DestroyWindow")
	procDispatchMessageA             = moduser32.NewProc("DispatchMessageA")
	procGetMessageA                  = moduser32.NewProc("GetMessageA")
	procRegisterClassA               = moduser32.NewProc("RegisterClassA")
	procRegisterDeviceNotificationA  = moduser32.NewProc("RegisterDeviceNotificationA")
	procUnregisterClassA             = moduser32.NewProc("UnregisterClassA")
	procUnregisterDeviceNotification = moduser32.NewProc("UnregisterDeviceNotification")
)

func getModuleHandle(moduleName *byte) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleHandleA.Addr(), 1, uintptr(unsafe.Pointer(moduleName)), 0, 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func createWindowEx(exstyle uint32, className *byte, windowText *byte, style uint32, x int32, y int32, width int32, height int32, parent syscall.Handle, menu syscall.Handle, hInstance syscall.Handle, lpParam uintptr) (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExA.Addr(), 12, uintptr(exstyle), uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(windowText)), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(parent), uintptr(menu), uintptr(hInstance), uintptr(lpParam))
	hwnd = syscall.Handle(r0)
	if hwnd == 0 {
		err = errnoErr(e1)
	}
	return
}

func defWindowProc(hwnd syscall.Handle, msg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	r0, _, _ := syscall.Syscall6(procDefWindowProcW.Addr(), 4, uintptr(hwnd), uintptr(msg), uintptr(wParam), uintptr(lParam), 0, 0)
	lResult = uintptr(r0)
	return
}

func destroyWindowEx(hwnd syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func dispatchMessage(msg *msg) (res int32, err error) {
	r0, _, e1 := syscall.Syscall(procDispatchMessageA.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	res = int32(r0)
	if res == 0 {
		err = errnoErr(e1)
	}
	return
}

func getMessage(msg *msg, hwnd syscall.Handle, msgFilterMin uint32, msgFilterMax uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetMessageA.Addr(), 4, uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(msgFilterMin), uintptr(msgFilterMax), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func registerClass(wndClass *wndClass) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassA.Addr(), 1, uintptr(unsafe.Pointer(wndClass)), 0, 0)
	atom = uint16(r0)
	if atom == 0 {
		err = errnoErr(e1)
	}
	return
}

func registerDeviceNotification(recipient syscall.Handle, filter *devBroadcastDeviceInterface, flags uint32) (devHandle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterDeviceNotificationA.Addr(), 3, uintptr(recipient), uintptr(unsafe.Pointer(filter)), uintptr(flags))
	devHandle = syscall.Handle(r0)
	if devHandle == 0 {
		err = errnoErr(e1)
	}
	return
}

func unregisterClass(className *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procUnregisterClassA.Addr(), 1, uintptr(unsafe.Pointer(className)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func unregisterDeviceNotification(deviceHandle syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procUnregisterDeviceNotification.Addr(), 1, uintptr(deviceHandle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
