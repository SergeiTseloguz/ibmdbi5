// mksyscall_windows.pl api.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package api

import "unsafe"
import "syscall"
import "os"


var (
	mododbc32 = syscall.NewLazyDLL(GetDllName())
	
	//mododbc32, err = syscall.LoadLibrary("square.dll")
   
	procSQLAllocHandle     = mododbc32.NewProc("SQLAllocHandle")
	procSQLDriverConnect   = mododbc32.NewProc("SQLDriverConnect")
	procSQLDisconnect      = mododbc32.NewProc("SQLDisconnect")
	procSQLSetConnectAttr = mododbc32.NewProc("SQLSetConnectAttr")
	
)

func GetDllName() string {
	if winArch := os.Getenv("PROCESSOR_ARCHITECTURE"); winArch == "x86" {
		return "db2cli.dll"
	} else {
		return "db2cli64.dll"
	}
}

func SQLAllocHandle(handleType SQLSMALLINT, inputHandle SQLHANDLE, outputHandle *SQLHANDLE) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLAllocHandle.Addr(), 3, uintptr(handleType), uintptr(inputHandle), uintptr(unsafe.Pointer(outputHandle)))
	ret = SQLRETURN(r0)
	return
}



func SQLDisconnect(connectionHandle SQLHDBC) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLDisconnect.Addr(), 1, uintptr(connectionHandle), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLDriverConnect(connectionHandle SQLHDBC, windowHandle SQLHWND, inConnectionString *SQLWCHAR, stringLength1 SQLSMALLINT, outConnectionString *SQLWCHAR, bufferLength SQLSMALLINT, stringLength2Ptr *SQLSMALLINT, driverCompletion SQLUSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall9(procSQLDriverConnect.Addr(), 8, uintptr(connectionHandle), uintptr(windowHandle), uintptr(unsafe.Pointer(inConnectionString)), uintptr(stringLength1), uintptr(unsafe.Pointer(outConnectionString)), uintptr(bufferLength), uintptr(unsafe.Pointer(stringLength2Ptr)), uintptr(driverCompletion), 0)
	ret = SQLRETURN(r0)
	return
}

func SQLSetConnectAttr(connectionHandle SQLHDBC, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall6(procSQLSetConnectAttrW.Addr(), 4, uintptr(connectionHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLRETURN(r0)
	return
}

