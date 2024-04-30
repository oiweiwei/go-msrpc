package win32

var (
	ErrorSuccess                                           = &Error{0x00000000, "ERROR_SUCCESS", "The operation completed successfully."}
	NerrSuccess                                            = &Error{0x00000000, "NERR_Success", "The operation completed successfully."}
	ErrorInvalidFunction                                   = &Error{0x00000001, "ERROR_INVALID_FUNCTION", "Incorrect function."}
	ErrorFileNotFound                                      = &Error{0x00000002, "ERROR_FILE_NOT_FOUND", "The system cannot find the file specified."}
	ErrorPathNotFound                                      = &Error{0x00000003, "ERROR_PATH_NOT_FOUND", "The system cannot find the path specified."}
	ErrorTooManyOpenFiles                                  = &Error{0x00000004, "ERROR_TOO_MANY_OPEN_FILES", "The system cannot open the file."}
	ErrorAccessDenied                                      = &Error{0x00000005, "ERROR_ACCESS_DENIED", "Access is denied."}
	ErrorInvalidHandle                                     = &Error{0x00000006, "ERROR_INVALID_HANDLE", "The handle is invalid."}
	ErrorArenaTrashed                                      = &Error{0x00000007, "ERROR_ARENA_TRASHED", "The storage control blocks were destroyed."}
	ErrorNotEnoughMemory                                   = &Error{0x00000008, "ERROR_NOT_ENOUGH_MEMORY", "Not enough storage is available to process this command."}
	ErrorInvalidBlock                                      = &Error{0x00000009, "ERROR_INVALID_BLOCK", "The storage control block address is invalid."}
	ErrorBadEnvironment                                    = &Error{0x0000000A, "ERROR_BAD_ENVIRONMENT", "The environment is incorrect."}
	ErrorBadFormat                                         = &Error{0x0000000B, "ERROR_BAD_FORMAT", "An attempt was made to load a program with an incorrect format."}
	ErrorInvalidAccess                                     = &Error{0x0000000C, "ERROR_INVALID_ACCESS", "The access code is invalid."}
	ErrorInvalidData                                       = &Error{0x0000000D, "ERROR_INVALID_DATA", "The data is invalid."}
	ErrorOutofmemory                                       = &Error{0x0000000E, "ERROR_OUTOFMEMORY", "Not enough storage is available to complete this operation."}
	ErrorInvalidDrive                                      = &Error{0x0000000F, "ERROR_INVALID_DRIVE", "The system cannot find the drive specified."}
	ErrorCurrentDirectory                                  = &Error{0x00000010, "ERROR_CURRENT_DIRECTORY", "The directory cannot be removed."}
	ErrorNotSameDevice                                     = &Error{0x00000011, "ERROR_NOT_SAME_DEVICE", "The system cannot move the file to a different disk drive."}
	ErrorNoMoreFiles                                       = &Error{0x00000012, "ERROR_NO_MORE_FILES", "There are no more files."}
	ErrorWriteProtect                                      = &Error{0x00000013, "ERROR_WRITE_PROTECT", "The media is write-protected."}
	ErrorBadUnit                                           = &Error{0x00000014, "ERROR_BAD_UNIT", "The system cannot find the device specified."}
	ErrorNotReady                                          = &Error{0x00000015, "ERROR_NOT_READY", "The device is not ready."}
	ErrorBadCommand                                        = &Error{0x00000016, "ERROR_BAD_COMMAND", "The device does not recognize the command."}
	ErrorCrc                                               = &Error{0x00000017, "ERROR_CRC", "Data error (cyclic redundancy check)."}
	ErrorBadLength                                         = &Error{0x00000018, "ERROR_BAD_LENGTH", "The program issued a command but the command length is incorrect."}
	ErrorSeek                                              = &Error{0x00000019, "ERROR_SEEK", "The drive cannot locate a specific area or track on the disk."}
	ErrorNotDosDisk                                        = &Error{0x0000001A, "ERROR_NOT_DOS_DISK", "The specified disk cannot be accessed."}
	ErrorSectorNotFound                                    = &Error{0x0000001B, "ERROR_SECTOR_NOT_FOUND", "The drive cannot find the sector requested."}
	ErrorOutOfPaper                                        = &Error{0x0000001C, "ERROR_OUT_OF_PAPER", "The printer is out of paper."}
	ErrorWriteFault                                        = &Error{0x0000001D, "ERROR_WRITE_FAULT", "The system cannot write to the specified device."}
	ErrorReadFault                                         = &Error{0x0000001E, "ERROR_READ_FAULT", "The system cannot read from the specified device."}
	ErrorGenFailure                                        = &Error{0x0000001F, "ERROR_GEN_FAILURE", "A device attached to the system is not functioning."}
	ErrorSharingViolation                                  = &Error{0x00000020, "ERROR_SHARING_VIOLATION", "The process cannot access the file because it is being used by another process."}
	ErrorLockViolation                                     = &Error{0x00000021, "ERROR_LOCK_VIOLATION", "The process cannot access the file because another process has locked a portion of the file."}
	ErrorWrongDisk                                         = &Error{0x00000022, "ERROR_WRONG_DISK", "The wrong disk is in the drive. Insert %2 (Volume Serial Number: %3) into drive %1."}
	ErrorSharingBufferExceeded                             = &Error{0x00000024, "ERROR_SHARING_BUFFER_EXCEEDED", "Too many files opened for sharing."}
	ErrorHandleEof                                         = &Error{0x00000026, "ERROR_HANDLE_EOF", "Reached the end of the file."}
	ErrorHandleDiskFull                                    = &Error{0x00000027, "ERROR_HANDLE_DISK_FULL", "The disk is full."}
	ErrorNotSupported                                      = &Error{0x00000032, "ERROR_NOT_SUPPORTED", "The request is not supported."}
	ErrorRemNotList                                        = &Error{0x00000033, "ERROR_REM_NOT_LIST", "Windows cannot find the network path. Verify that the network path is correct and the destination computer is not busy or turned off. If Windows still cannot find the network path, contact your network administrator."}
	ErrorDupName                                           = &Error{0x00000034, "ERROR_DUP_NAME", "You were not connected because a duplicate name exists on the network. Go to System in Control Panel to change the computer name, and then try again."}
	ErrorBadNetpath                                        = &Error{0x00000035, "ERROR_BAD_NETPATH", "The network path was not found."}
	ErrorNetworkBusy                                       = &Error{0x00000036, "ERROR_NETWORK_BUSY", "The network is busy."}
	ErrorDevNotExist                                       = &Error{0x00000037, "ERROR_DEV_NOT_EXIST", "The specified network resource or device is no longer available."}
	ErrorTooManyCmds                                       = &Error{0x00000038, "ERROR_TOO_MANY_CMDS", "The network BIOS command limit has been reached."}
	ErrorAdapHdwErr                                        = &Error{0x00000039, "ERROR_ADAP_HDW_ERR", "A network adapter hardware error occurred."}
	ErrorBadNetResp                                        = &Error{0x0000003A, "ERROR_BAD_NET_RESP", "The specified server cannot perform the requested operation."}
	ErrorUnexpNetErr                                       = &Error{0x0000003B, "ERROR_UNEXP_NET_ERR", "An unexpected network error occurred."}
	ErrorBadRemAdap                                        = &Error{0x0000003C, "ERROR_BAD_REM_ADAP", "The remote adapter is not compatible."}
	ErrorPrintqFull                                        = &Error{0x0000003D, "ERROR_PRINTQ_FULL", "The print queue is full."}
	ErrorNoSpoolSpace                                      = &Error{0x0000003E, "ERROR_NO_SPOOL_SPACE", "Space to store the file waiting to be printed is not available on the server."}
	ErrorPrintCancelled                                    = &Error{0x0000003F, "ERROR_PRINT_CANCELLED", "Your file waiting to be printed was deleted."}
	ErrorNetnameDeleted                                    = &Error{0x00000040, "ERROR_NETNAME_DELETED", "The specified network name is no longer available."}
	ErrorNetworkAccessDenied                               = &Error{0x00000041, "ERROR_NETWORK_ACCESS_DENIED", "Network access is denied."}
	ErrorBadDevType                                        = &Error{0x00000042, "ERROR_BAD_DEV_TYPE", "The network resource type is not correct."}
	ErrorBadNetName                                        = &Error{0x00000043, "ERROR_BAD_NET_NAME", "The network name cannot be found."}
	ErrorTooManyNames                                      = &Error{0x00000044, "ERROR_TOO_MANY_NAMES", "The name limit for the local computer network adapter card was exceeded."}
	ErrorTooManySess                                       = &Error{0x00000045, "ERROR_TOO_MANY_SESS", "The network BIOS session limit was exceeded."}
	ErrorSharingPaused                                     = &Error{0x00000046, "ERROR_SHARING_PAUSED", "The remote server has been paused or is in the process of being started."}
	ErrorReqNotAccep                                       = &Error{0x00000047, "ERROR_REQ_NOT_ACCEP", "No more connections can be made to this remote computer at this time because the computer has accepted the maximum number of connections."}
	ErrorRedirPaused                                       = &Error{0x00000048, "ERROR_REDIR_PAUSED", "The specified printer or disk device has been paused."}
	ErrorFileExists                                        = &Error{0x00000050, "ERROR_FILE_EXISTS", "The file exists."}
	ErrorCannotMake                                        = &Error{0x00000052, "ERROR_CANNOT_MAKE", "The directory or file cannot be created."}
	ErrorFailI24                                           = &Error{0x00000053, "ERROR_FAIL_I24", "Fail on INT 24."}
	ErrorOutOfStructures                                   = &Error{0x00000054, "ERROR_OUT_OF_STRUCTURES", "Storage to process this request is not available."}
	ErrorAlreadyAssigned                                   = &Error{0x00000055, "ERROR_ALREADY_ASSIGNED", "The local device name is already in use."}
	ErrorInvalidPassword                                   = &Error{0x00000056, "ERROR_INVALID_PASSWORD", "The specified network password is not correct."}
	ErrorInvalidParameter                                  = &Error{0x00000057, "ERROR_INVALID_PARAMETER", "The parameter is incorrect."}
	ErrorNetWriteFault                                     = &Error{0x00000058, "ERROR_NET_WRITE_FAULT", "A write fault occurred on the network."}
	ErrorNoProcSlots                                       = &Error{0x00000059, "ERROR_NO_PROC_SLOTS", "The system cannot start another process at this time."}
	ErrorTooManySemaphores                                 = &Error{0x00000064, "ERROR_TOO_MANY_SEMAPHORES", "Cannot create another system semaphore."}
	ErrorExclSemAlreadyOwned                               = &Error{0x00000065, "ERROR_EXCL_SEM_ALREADY_OWNED", "The exclusive semaphore is owned by another process."}
	ErrorSemIsSet                                          = &Error{0x00000066, "ERROR_SEM_IS_SET", "The semaphore is set and cannot be closed."}
	ErrorTooManySemRequests                                = &Error{0x00000067, "ERROR_TOO_MANY_SEM_REQUESTS", "The semaphore cannot be set again."}
	ErrorInvalidAtInterruptTime                            = &Error{0x00000068, "ERROR_INVALID_AT_INTERRUPT_TIME", "Cannot request exclusive semaphores at interrupt time."}
	ErrorSemOwnerDied                                      = &Error{0x00000069, "ERROR_SEM_OWNER_DIED", "The previous ownership of this semaphore has ended."}
	ErrorSemUserLimit                                      = &Error{0x0000006A, "ERROR_SEM_USER_LIMIT", "Insert the disk for drive %1."}
	ErrorDiskChange                                        = &Error{0x0000006B, "ERROR_DISK_CHANGE", "The program stopped because an alternate disk was not inserted."}
	ErrorDriveLocked                                       = &Error{0x0000006C, "ERROR_DRIVE_LOCKED", "The disk is in use or locked by another process."}
	ErrorBrokenPipe                                        = &Error{0x0000006D, "ERROR_BROKEN_PIPE", "The pipe has been ended."}
	ErrorOpenFailed                                        = &Error{0x0000006E, "ERROR_OPEN_FAILED", "The system cannot open the device or file specified."}
	ErrorBufferOverflow                                    = &Error{0x0000006F, "ERROR_BUFFER_OVERFLOW", "The file name is too long."}
	ErrorDiskFull                                          = &Error{0x00000070, "ERROR_DISK_FULL", "There is not enough space on the disk."}
	ErrorNoMoreSearchHandles                               = &Error{0x00000071, "ERROR_NO_MORE_SEARCH_HANDLES", "No more internal file identifiers are available."}
	ErrorInvalidTargetHandle                               = &Error{0x00000072, "ERROR_INVALID_TARGET_HANDLE", "The target internal file identifier is incorrect."}
	ErrorInvalidCategory                                   = &Error{0x00000075, "ERROR_INVALID_CATEGORY", "The Input Output Control (IOCTL) call made by the application program is not correct."}
	ErrorInvalidVerifySwitch                               = &Error{0x00000076, "ERROR_INVALID_VERIFY_SWITCH", "The verify-on-write switch parameter value is not correct."}
	ErrorBadDriverLevel                                    = &Error{0x00000077, "ERROR_BAD_DRIVER_LEVEL", "The system does not support the command requested."}
	ErrorCallNotImplemented                                = &Error{0x00000078, "ERROR_CALL_NOT_IMPLEMENTED", "This function is not supported on this system."}
	ErrorSemTimeout                                        = &Error{0x00000079, "ERROR_SEM_TIMEOUT", "The semaphore time-out period has expired."}
	ErrorInsufficientBuffer                                = &Error{0x0000007A, "ERROR_INSUFFICIENT_BUFFER", "The data area passed to a system call is too small."}
	ErrorInvalidName                                       = &Error{0x0000007B, "ERROR_INVALID_NAME", "The file name, directory name, or volume label syntax is incorrect."}
	ErrorInvalidLevel                                      = &Error{0x0000007C, "ERROR_INVALID_LEVEL", "The system call level is not correct."}
	ErrorNoVolumeLabel                                     = &Error{0x0000007D, "ERROR_NO_VOLUME_LABEL", "The disk has no volume label."}
	ErrorModNotFound                                       = &Error{0x0000007E, "ERROR_MOD_NOT_FOUND", "The specified module could not be found."}
	ErrorProcNotFound                                      = &Error{0x0000007F, "ERROR_PROC_NOT_FOUND", "The specified procedure could not be found."}
	ErrorWaitNoChildren                                    = &Error{0x00000080, "ERROR_WAIT_NO_CHILDREN", "There are no child processes to wait for."}
	ErrorChildNotComplete                                  = &Error{0x00000081, "ERROR_CHILD_NOT_COMPLETE", "The %1 application cannot be run in Win32 mode."}
	ErrorDirectAccessHandle                                = &Error{0x00000082, "ERROR_DIRECT_ACCESS_HANDLE", "Attempt to use a file handle to an open disk partition for an operation other than raw disk I/O."}
	ErrorNegativeSeek                                      = &Error{0x00000083, "ERROR_NEGATIVE_SEEK", "An attempt was made to move the file pointer before the beginning of the file."}
	ErrorSeekOnDevice                                      = &Error{0x00000084, "ERROR_SEEK_ON_DEVICE", "The file pointer cannot be set on the specified device or file."}
	ErrorIsJoinTarget                                      = &Error{0x00000085, "ERROR_IS_JOIN_TARGET", "A JOIN or SUBST command cannot be used for a drive that contains previously joined drives."}
	ErrorIsJoined                                          = &Error{0x00000086, "ERROR_IS_JOINED", "An attempt was made to use a JOIN or SUBST command on a drive that has already been joined."}
	ErrorIsSubsted                                         = &Error{0x00000087, "ERROR_IS_SUBSTED", "An attempt was made to use a JOIN or SUBST command on a drive that has already been substituted."}
	ErrorNotJoined                                         = &Error{0x00000088, "ERROR_NOT_JOINED", "The system tried to delete the JOIN of a drive that is not joined."}
	ErrorNotSubsted                                        = &Error{0x00000089, "ERROR_NOT_SUBSTED", "The system tried to delete the substitution of a drive that is not substituted."}
	ErrorJoinToJoin                                        = &Error{0x0000008A, "ERROR_JOIN_TO_JOIN", "The system tried to join a drive to a directory on a joined drive."}
	ErrorSubstToSubst                                      = &Error{0x0000008B, "ERROR_SUBST_TO_SUBST", "The system tried to substitute a drive to a directory on a substituted drive."}
	ErrorJoinToSubst                                       = &Error{0x0000008C, "ERROR_JOIN_TO_SUBST", "The system tried to join a drive to a directory on a substituted drive."}
	ErrorSubstToJoin                                       = &Error{0x0000008D, "ERROR_SUBST_TO_JOIN", "The system tried to SUBST a drive to a directory on a joined drive."}
	ErrorBusyDrive                                         = &Error{0x0000008E, "ERROR_BUSY_DRIVE", "The system cannot perform a JOIN or SUBST at this time."}
	ErrorSameDrive                                         = &Error{0x0000008F, "ERROR_SAME_DRIVE", "The system cannot join or substitute a drive to or for a directory on the same drive."}
	ErrorDirNotRoot                                        = &Error{0x00000090, "ERROR_DIR_NOT_ROOT", "The directory is not a subdirectory of the root directory."}
	ErrorDirNotEmpty                                       = &Error{0x00000091, "ERROR_DIR_NOT_EMPTY", "The directory is not empty."}
	ErrorIsSubstPath                                       = &Error{0x00000092, "ERROR_IS_SUBST_PATH", "The path specified is being used in a substitute."}
	ErrorIsJoinPath                                        = &Error{0x00000093, "ERROR_IS_JOIN_PATH", "Not enough resources are available to process this command."}
	ErrorPathBusy                                          = &Error{0x00000094, "ERROR_PATH_BUSY", "The path specified cannot be used at this time."}
	ErrorIsSubstTarget                                     = &Error{0x00000095, "ERROR_IS_SUBST_TARGET", "An attempt was made to join or substitute a drive for which a directory on the drive is the target of a previous substitute."}
	ErrorSystemTrace                                       = &Error{0x00000096, "ERROR_SYSTEM_TRACE", "System trace information was not specified in your CONFIG.SYS file, or tracing is disallowed."}
	ErrorInvalidEventCount                                 = &Error{0x00000097, "ERROR_INVALID_EVENT_COUNT", "The number of specified semaphore events for DosMuxSemWait is not correct."}
	ErrorTooManyMuxwaiters                                 = &Error{0x00000098, "ERROR_TOO_MANY_MUXWAITERS", "DosMuxSemWait did not execute; too many semaphores are already set."}
	ErrorInvalidListFormat                                 = &Error{0x00000099, "ERROR_INVALID_LIST_FORMAT", "The DosMuxSemWait list is not correct."}
	ErrorLabelTooLong                                      = &Error{0x0000009A, "ERROR_LABEL_TOO_LONG", "The volume label you entered exceeds the label character limit of the destination file system."}
	ErrorTooManyTcbs                                       = &Error{0x0000009B, "ERROR_TOO_MANY_TCBS", "Cannot create another thread."}
	ErrorSignalRefused                                     = &Error{0x0000009C, "ERROR_SIGNAL_REFUSED", "The recipient process has refused the signal."}
	ErrorDiscarded                                         = &Error{0x0000009D, "ERROR_DISCARDED", "The segment is already discarded and cannot be locked."}
	ErrorNotLocked                                         = &Error{0x0000009E, "ERROR_NOT_LOCKED", "The segment is already unlocked."}
	ErrorBadThreadidAddr                                   = &Error{0x0000009F, "ERROR_BAD_THREADID_ADDR", "The address for the thread ID is not correct."}
	ErrorBadArguments                                      = &Error{0x000000A0, "ERROR_BAD_ARGUMENTS", "One or more arguments are not correct."}
	ErrorBadPathname                                       = &Error{0x000000A1, "ERROR_BAD_PATHNAME", "The specified path is invalid."}
	ErrorSignalPending                                     = &Error{0x000000A2, "ERROR_SIGNAL_PENDING", "A signal is already pending."}
	ErrorMaxThrdsReached                                   = &Error{0x000000A4, "ERROR_MAX_THRDS_REACHED", "No more threads can be created in the system."}
	ErrorLockFailed                                        = &Error{0x000000A7, "ERROR_LOCK_FAILED", "Unable to lock a region of a file."}
	ErrorBusy                                              = &Error{0x000000AA, "ERROR_BUSY", "The requested resource is in use."}
	ErrorCancelViolation                                   = &Error{0x000000AD, "ERROR_CANCEL_VIOLATION", "A lock request was not outstanding for the supplied cancel region."}
	ErrorAtomicLocksNotSupported                           = &Error{0x000000AE, "ERROR_ATOMIC_LOCKS_NOT_SUPPORTED", "The file system does not support atomic changes to the lock type."}
	ErrorInvalidSegmentNumber                              = &Error{0x000000B4, "ERROR_INVALID_SEGMENT_NUMBER", "The system detected a segment number that was not correct."}
	ErrorInvalidOrdinal                                    = &Error{0x000000B6, "ERROR_INVALID_ORDINAL", "The operating system cannot run %1."}
	ErrorAlreadyExists                                     = &Error{0x000000B7, "ERROR_ALREADY_EXISTS", "Cannot create a file when that file already exists."}
	ErrorInvalidFlagNumber                                 = &Error{0x000000BA, "ERROR_INVALID_FLAG_NUMBER", "The flag passed is not correct."}
	ErrorSemNotFound                                       = &Error{0x000000BB, "ERROR_SEM_NOT_FOUND", "The specified system semaphore name was not found."}
	ErrorInvalidStartingCodeseg                            = &Error{0x000000BC, "ERROR_INVALID_STARTING_CODESEG", "The operating system cannot run %1."}
	ErrorInvalidStackseg                                   = &Error{0x000000BD, "ERROR_INVALID_STACKSEG", "The operating system cannot run %1."}
	ErrorInvalidModuletype                                 = &Error{0x000000BE, "ERROR_INVALID_MODULETYPE", "The operating system cannot run %1."}
	ErrorInvalidExeSignature                               = &Error{0x000000BF, "ERROR_INVALID_EXE_SIGNATURE", "Cannot run %1 in Win32 mode."}
	ErrorExeMarkedInvalid                                  = &Error{0x000000C0, "ERROR_EXE_MARKED_INVALID", "The operating system cannot run %1."}
	ErrorBadExeFormat                                      = &Error{0x000000C1, "ERROR_BAD_EXE_FORMAT", "%1 is not a valid Win32 application."}
	ErrorIteratedDataExceeds64k                            = &Error{0x000000C2, "ERROR_ITERATED_DATA_EXCEEDS_64k", "The operating system cannot run %1."}
	ErrorInvalidMinallocsize                               = &Error{0x000000C3, "ERROR_INVALID_MINALLOCSIZE", "The operating system cannot run %1."}
	ErrorDynlinkFromInvalidRing                            = &Error{0x000000C4, "ERROR_DYNLINK_FROM_INVALID_RING", "The operating system cannot run this application program."}
	ErrorIoplNotEnabled                                    = &Error{0x000000C5, "ERROR_IOPL_NOT_ENABLED", "The operating system is not presently configured to run this application."}
	ErrorInvalidSegdpl                                     = &Error{0x000000C6, "ERROR_INVALID_SEGDPL", "The operating system cannot run %1."}
	ErrorAutodatasegExceeds64k                             = &Error{0x000000C7, "ERROR_AUTODATASEG_EXCEEDS_64k", "The operating system cannot run this application program."}
	ErrorRing2segMustBeMovable                             = &Error{0x000000C8, "ERROR_RING2SEG_MUST_BE_MOVABLE", "The code segment cannot be greater than or equal to 64 KB."}
	ErrorRelocChainXeedsSeglim                             = &Error{0x000000C9, "ERROR_RELOC_CHAIN_XEEDS_SEGLIM", "The operating system cannot run %1."}
	ErrorInfloopInRelocChain                               = &Error{0x000000CA, "ERROR_INFLOOP_IN_RELOC_CHAIN", "The operating system cannot run %1."}
	ErrorEnvvarNotFound                                    = &Error{0x000000CB, "ERROR_ENVVAR_NOT_FOUND", "The system could not find the environment option that was entered."}
	ErrorNoSignalSent                                      = &Error{0x000000CD, "ERROR_NO_SIGNAL_SENT", "No process in the command subtree has a signal handler."}
	ErrorFilenameExcedRange                                = &Error{0x000000CE, "ERROR_FILENAME_EXCED_RANGE", "The file name or extension is too long."}
	ErrorRing2StackInUse                                   = &Error{0x000000CF, "ERROR_RING2_STACK_IN_USE", "The ring 2 stack is in use."}
	ErrorMetaExpansionTooLong                              = &Error{0x000000D0, "ERROR_META_EXPANSION_TOO_LONG", "The asterisk (*) or question mark (?) global file name characters are entered incorrectly, or too many global file name characters are specified."}
	ErrorInvalidSignalNumber                               = &Error{0x000000D1, "ERROR_INVALID_SIGNAL_NUMBER", "The signal being posted is not correct."}
	ErrorThread1Inactive                                   = &Error{0x000000D2, "ERROR_THREAD_1_INACTIVE", "The signal handler cannot be set."}
	ErrorLocked                                            = &Error{0x000000D4, "ERROR_LOCKED", "The segment is locked and cannot be reallocated."}
	ErrorTooManyModules                                    = &Error{0x000000D6, "ERROR_TOO_MANY_MODULES", "Too many dynamic-link modules are attached to this program or dynamic-link module."}
	ErrorNestingNotAllowed                                 = &Error{0x000000D7, "ERROR_NESTING_NOT_ALLOWED", "Cannot nest calls to LoadModule."}
	ErrorExeMachineTypeMismatch                            = &Error{0x000000D8, "ERROR_EXE_MACHINE_TYPE_MISMATCH", "This version of %1 is not compatible with the version of Windows you're running. Check your computer's system information to see whether you need an x86 (32-bit) or x64 (64-bit) version of the program, and then contact the software publisher."}
	ErrorExeCannotModifySignedBinary                       = &Error{0x000000D9, "ERROR_EXE_CANNOT_MODIFY_SIGNED_BINARY", "The image file %1 is signed, unable to modify."}
	ErrorExeCannotModifyStrongSignedBinary                 = &Error{0x000000DA, "ERROR_EXE_CANNOT_MODIFY_STRONG_SIGNED_BINARY", "The image file %1 is strong signed, unable to modify."}
	ErrorFileCheckedOut                                    = &Error{0x000000DC, "ERROR_FILE_CHECKED_OUT", "This file is checked out or locked for editing by another user."}
	ErrorCheckoutRequired                                  = &Error{0x000000DD, "ERROR_CHECKOUT_REQUIRED", "The file must be checked out before saving changes."}
	ErrorBadFileType                                       = &Error{0x000000DE, "ERROR_BAD_FILE_TYPE", "The file type being saved or retrieved has been blocked."}
	ErrorFileTooLarge                                      = &Error{0x000000DF, "ERROR_FILE_TOO_LARGE", "The file size exceeds the limit allowed and cannot be saved."}
	ErrorFormsAuthRequired                                 = &Error{0x000000E0, "ERROR_FORMS_AUTH_REQUIRED", "Access denied. Before opening files in this location, you must first browse to the website and select the option to sign in automatically."}
	ErrorVirusInfected                                     = &Error{0x000000E1, "ERROR_VIRUS_INFECTED", "Operation did not complete successfully because the file contains a virus."}
	ErrorVirusDeleted                                      = &Error{0x000000E2, "ERROR_VIRUS_DELETED", "This file contains a virus and cannot be opened. Due to the nature of this virus, the file has been removed from this location."}
	ErrorPipeLocal                                         = &Error{0x000000E5, "ERROR_PIPE_LOCAL", "The pipe is local."}
	ErrorBadPipe                                           = &Error{0x000000E6, "ERROR_BAD_PIPE", "The pipe state is invalid."}
	ErrorPipeBusy                                          = &Error{0x000000E7, "ERROR_PIPE_BUSY", "All pipe instances are busy."}
	ErrorNoData                                            = &Error{0x000000E8, "ERROR_NO_DATA", "The pipe is being closed."}
	ErrorPipeNotConnected                                  = &Error{0x000000E9, "ERROR_PIPE_NOT_CONNECTED", "No process is on the other end of the pipe."}
	ErrorMoreData                                          = &Error{0x000000EA, "ERROR_MORE_DATA", "More data is available."}
	ErrorVcDisconnected                                    = &Error{0x000000F0, "ERROR_VC_DISCONNECTED", "The session was canceled."}
	ErrorInvalidEaName                                     = &Error{0x000000FE, "ERROR_INVALID_EA_NAME", "The specified extended attribute name was invalid."}
	ErrorEaListInconsistent                                = &Error{0x000000FF, "ERROR_EA_LIST_INCONSISTENT", "The extended attributes are inconsistent."}
	WaitTimeout                                            = &Error{0x00000102, "WAIT_TIMEOUT", "The wait operation timed out."}
	ErrorNoMoreItems                                       = &Error{0x00000103, "ERROR_NO_MORE_ITEMS", "No more data is available."}
	ErrorCannotCopy                                        = &Error{0x0000010A, "ERROR_CANNOT_COPY", "The copy functions cannot be used."}
	ErrorDirectory                                         = &Error{0x0000010B, "ERROR_DIRECTORY", "The directory name is invalid."}
	ErrorEasDidntFit                                       = &Error{0x00000113, "ERROR_EAS_DIDNT_FIT", "The extended attributes did not fit in the buffer."}
	ErrorEaFileCorrupt                                     = &Error{0x00000114, "ERROR_EA_FILE_CORRUPT", "The extended attribute file on the mounted file system is corrupt."}
	ErrorEaTableFull                                       = &Error{0x00000115, "ERROR_EA_TABLE_FULL", "The extended attribute table file is full."}
	ErrorInvalidEaHandle                                   = &Error{0x00000116, "ERROR_INVALID_EA_HANDLE", "The specified extended attribute handle is invalid."}
	ErrorEasNotSupported                                   = &Error{0x0000011A, "ERROR_EAS_NOT_SUPPORTED", "The mounted file system does not support extended attributes."}
	ErrorNotOwner                                          = &Error{0x00000120, "ERROR_NOT_OWNER", "Attempt to release mutex not owned by caller."}
	ErrorTooManyPosts                                      = &Error{0x0000012A, "ERROR_TOO_MANY_POSTS", "Too many posts were made to a semaphore."}
	ErrorPartialCopy                                       = &Error{0x0000012B, "ERROR_PARTIAL_COPY", "Only part of a ReadProcessMemory or WriteProcessMemory request was completed."}
	ErrorOplockNotGranted                                  = &Error{0x0000012C, "ERROR_OPLOCK_NOT_GRANTED", "The oplock request is denied."}
	ErrorInvalidOplockProtocol                             = &Error{0x0000012D, "ERROR_INVALID_OPLOCK_PROTOCOL", "An invalid oplock acknowledgment was received by the system."}
	ErrorDiskTooFragmented                                 = &Error{0x0000012E, "ERROR_DISK_TOO_FRAGMENTED", "The volume is too fragmented to complete this operation."}
	ErrorDeletePending                                     = &Error{0x0000012F, "ERROR_DELETE_PENDING", "The file cannot be opened because it is in the process of being deleted."}
	ErrorMrMidNotFound                                     = &Error{0x0000013D, "ERROR_MR_MID_NOT_FOUND", "The system cannot find message text for message number 0x%1 in the message file for %2."}
	ErrorScopeNotFound                                     = &Error{0x0000013E, "ERROR_SCOPE_NOT_FOUND", "The scope specified was not found."}
	ErrorFailNoactionReboot                                = &Error{0x0000015E, "ERROR_FAIL_NOACTION_REBOOT", "No action was taken because a system reboot is required."}
	ErrorFailShutdown                                      = &Error{0x0000015F, "ERROR_FAIL_SHUTDOWN", "The shutdown operation failed."}
	ErrorFailRestart                                       = &Error{0x00000160, "ERROR_FAIL_RESTART", "The restart operation failed."}
	ErrorMaxSessionsReached                                = &Error{0x00000161, "ERROR_MAX_SESSIONS_REACHED", "The maximum number of sessions has been reached."}
	ErrorThreadModeAlreadyBackground                       = &Error{0x00000190, "ERROR_THREAD_MODE_ALREADY_BACKGROUND", "The thread is already in background processing mode."}
	ErrorThreadModeNotBackground                           = &Error{0x00000191, "ERROR_THREAD_MODE_NOT_BACKGROUND", "The thread is not in background processing mode."}
	ErrorProcessModeAlreadyBackground                      = &Error{0x00000192, "ERROR_PROCESS_MODE_ALREADY_BACKGROUND", "The process is already in background processing mode."}
	ErrorProcessModeNotBackground                          = &Error{0x00000193, "ERROR_PROCESS_MODE_NOT_BACKGROUND", "The process is not in background processing mode."}
	ErrorInvalidAddress                                    = &Error{0x000001E7, "ERROR_INVALID_ADDRESS", "Attempt to access invalid address."}
	ErrorUserProfileLoad                                   = &Error{0x000001F4, "ERROR_USER_PROFILE_LOAD", "User profile cannot be loaded."}
	ErrorArithmeticOverflow                                = &Error{0x00000216, "ERROR_ARITHMETIC_OVERFLOW", "Arithmetic result exceeded 32 bits."}
	ErrorPipeConnected                                     = &Error{0x00000217, "ERROR_PIPE_CONNECTED", "There is a process on the other end of the pipe."}
	ErrorPipeListening                                     = &Error{0x00000218, "ERROR_PIPE_LISTENING", "Waiting for a process to open the other end of the pipe."}
	ErrorVerifierStop                                      = &Error{0x00000219, "ERROR_VERIFIER_STOP", "Application verifier has found an error in the current process."}
	ErrorAbiosError                                        = &Error{0x0000021A, "ERROR_ABIOS_ERROR", "An error occurred in the ABIOS subsystem."}
	ErrorWx86Warning                                       = &Error{0x0000021B, "ERROR_WX86_WARNING", "A warning occurred in the WX86 subsystem."}
	ErrorWx86Error                                         = &Error{0x0000021C, "ERROR_WX86_ERROR", "An error occurred in the WX86 subsystem."}
	ErrorTimerNotCanceled                                  = &Error{0x0000021D, "ERROR_TIMER_NOT_CANCELED", "An attempt was made to cancel or set a timer that has an associated asynchronous procedure call (APC) and the subject thread is not the thread that originally set the timer with an associated APC routine."}
	ErrorUnwind                                            = &Error{0x0000021E, "ERROR_UNWIND", "Unwind exception code."}
	ErrorBadStack                                          = &Error{0x0000021F, "ERROR_BAD_STACK", "An invalid or unaligned stack was encountered during an unwind operation."}
	ErrorInvalidUnwindTarget                               = &Error{0x00000220, "ERROR_INVALID_UNWIND_TARGET", "An invalid unwind target was encountered during an unwind operation."}
	ErrorInvalidPortAttributes                             = &Error{0x00000221, "ERROR_INVALID_PORT_ATTRIBUTES", "Invalid object attributes specified to NtCreatePort or invalid port attributes specified to NtConnectPort."}
	ErrorPortMessageTooLong                                = &Error{0x00000222, "ERROR_PORT_MESSAGE_TOO_LONG", "Length of message passed to NtRequestPort or NtRequestWaitReplyPort was longer than the maximum message allowed by the port."}
	ErrorInvalidQuotaLower                                 = &Error{0x00000223, "ERROR_INVALID_QUOTA_LOWER", "An attempt was made to lower a quota limit below the current usage."}
	ErrorDeviceAlreadyAttached                             = &Error{0x00000224, "ERROR_DEVICE_ALREADY_ATTACHED", "An attempt was made to attach to a device that was already attached to another device."}
	ErrorInstructionMisalignment                           = &Error{0x00000225, "ERROR_INSTRUCTION_MISALIGNMENT", "An attempt was made to execute an instruction at an unaligned address, and the host system does not support unaligned instruction references."}
	ErrorProfilingNotStarted                               = &Error{0x00000226, "ERROR_PROFILING_NOT_STARTED", "Profiling not started."}
	ErrorProfilingNotStopped                               = &Error{0x00000227, "ERROR_PROFILING_NOT_STOPPED", "Profiling not stopped."}
	ErrorCouldNotInterpret                                 = &Error{0x00000228, "ERROR_COULD_NOT_INTERPRET", "The passed ACL did not contain the minimum required information."}
	ErrorProfilingAtLimit                                  = &Error{0x00000229, "ERROR_PROFILING_AT_LIMIT", "The number of active profiling objects is at the maximum and no more can be started."}
	ErrorCantWait                                          = &Error{0x0000022A, "ERROR_CANT_WAIT", "Used to indicate that an operation cannot continue without blocking for I/O."}
	ErrorCantTerminateSelf                                 = &Error{0x0000022B, "ERROR_CANT_TERMINATE_SELF", "Indicates that a thread attempted to terminate itself by default (called NtTerminateThread with NULL) and it was the last thread in the current process."}
	ErrorUnexpectedMmCreateErr                             = &Error{0x0000022C, "ERROR_UNEXPECTED_MM_CREATE_ERR", "If an MM error is returned that is not defined in the standard FsRtl filter, it is converted to one of the following errors that is guaranteed to be in the filter. In this case, information is lost; however, the filter correctly handles the exception."}
	ErrorUnexpectedMmMapError                              = &Error{0x0000022D, "ERROR_UNEXPECTED_MM_MAP_ERROR", "If an MM error is returned that is not defined in the standard FsRtl filter, it is converted to one of the following errors that is guaranteed to be in the filter. In this case, information is lost; however, the filter correctly handles the exception."}
	ErrorUnexpectedMmExtendErr                             = &Error{0x0000022E, "ERROR_UNEXPECTED_MM_EXTEND_ERR", "If an MM error is returned that is not defined in the standard FsRtl filter, it is converted to one of the following errors that is guaranteed to be in the filter. In this case, information is lost; however, the filter correctly handles the exception."}
	ErrorBadFunctionTable                                  = &Error{0x0000022F, "ERROR_BAD_FUNCTION_TABLE", "A malformed function table was encountered during an unwind operation."}
	ErrorNoGuidTranslation                                 = &Error{0x00000230, "ERROR_NO_GUID_TRANSLATION", "Indicates that an attempt was made to assign protection to a file system file or directory and one of the SIDs in the security descriptor could not be translated into a GUID that could be stored by the file system. This causes the protection attempt to fail, which might cause a file creation attempt to fail."}
	ErrorInvalidLdtSize                                    = &Error{0x00000231, "ERROR_INVALID_LDT_SIZE", "Indicates that an attempt was made to grow a local domain table (LDT) by setting its size, or that the size was not an even number of selectors."}
	ErrorInvalidLdtOffset                                  = &Error{0x00000233, "ERROR_INVALID_LDT_OFFSET", "Indicates that the starting value for the LDT information was not an integral multiple of the selector size."}
	ErrorInvalidLdtDescriptor                              = &Error{0x00000234, "ERROR_INVALID_LDT_DESCRIPTOR", "Indicates that the user supplied an invalid descriptor when trying to set up LDT descriptors."}
	ErrorTooManyThreads                                    = &Error{0x00000235, "ERROR_TOO_MANY_THREADS", "Indicates a process has too many threads to perform the requested action. For example, assignment of a primary token can be performed only when a process has zero or one threads."}
	ErrorThreadNotInProcess                                = &Error{0x00000236, "ERROR_THREAD_NOT_IN_PROCESS", "An attempt was made to operate on a thread within a specific process, but the thread specified is not in the process specified."}
	ErrorPagefileQuotaExceeded                             = &Error{0x00000237, "ERROR_PAGEFILE_QUOTA_EXCEEDED", "Page file quota was exceeded."}
	ErrorLogonServerConflict                               = &Error{0x00000238, "ERROR_LOGON_SERVER_CONFLICT", "The Netlogon service cannot start because another Netlogon service running in the domain conflicts with the specified role."}
	ErrorSynchronizationRequired                           = &Error{0x00000239, "ERROR_SYNCHRONIZATION_REQUIRED", "On applicable Windows Server releases, the Security Accounts Manager (SAM) database is significantly out of synchronization with the copy on the domain controller. A complete synchronization is required."}
	ErrorNetOpenFailed                                     = &Error{0x0000023A, "ERROR_NET_OPEN_FAILED", "The NtCreateFile API failed. This error should never be returned to an application, it is a place holder for the Windows LAN Manager Redirector to use in its internal error mapping routines."}
	ErrorIoPrivilegeFailed                                 = &Error{0x0000023B, "ERROR_IO_PRIVILEGE_FAILED", "{Privilege Failed} The I/O permissions for the process could not be changed."}
	ErrorControlCExit                                      = &Error{0x0000023C, "ERROR_CONTROL_C_EXIT", "{Application Exit by CTRL+C} The application terminated as a result of a CTRL+C."}
	ErrorMissingSystemfile                                 = &Error{0x0000023D, "ERROR_MISSING_SYSTEMFILE", "{Missing System File} The required system file %hs is bad or missing."}
	ErrorUnhandledException                                = &Error{0x0000023E, "ERROR_UNHANDLED_EXCEPTION", "{Application Error} The exception %s (0x%08lx) occurred in the application at location 0x%08lx."}
	ErrorAppInitFailure                                    = &Error{0x0000023F, "ERROR_APP_INIT_FAILURE", "{Application Error} The application failed to initialize properly (0x%lx). Click OK to terminate the application."}
	ErrorPagefileCreateFailed                              = &Error{0x00000240, "ERROR_PAGEFILE_CREATE_FAILED", "{Unable to Create Paging File} The creation of the paging file %hs failed (%lx). The requested size was %ld."}
	ErrorInvalidImageHash                                  = &Error{0x00000241, "ERROR_INVALID_IMAGE_HASH", "The hash for the image cannot be found in the system catalogs. The image is likely corrupt or the victim of tampering."}
	ErrorNoPagefile                                        = &Error{0x00000242, "ERROR_NO_PAGEFILE", "{No Paging File Specified} No paging file was specified in the system configuration."}
	ErrorIllegalFloatContext                               = &Error{0x00000243, "ERROR_ILLEGAL_FLOAT_CONTEXT", "{EXCEPTION} A real-mode application issued a floating-point instruction, and floating-point hardware is not present."}
	ErrorNoEventPair                                       = &Error{0x00000244, "ERROR_NO_EVENT_PAIR", "An event pair synchronization operation was performed using the thread-specific client/server event pair object, but no event pair object was associated with the thread."}
	ErrorDomainCtrlrConfigError                            = &Error{0x00000245, "ERROR_DOMAIN_CTRLR_CONFIG_ERROR", "A domain server has an incorrect configuration."}
	ErrorIllegalCharacter                                  = &Error{0x00000246, "ERROR_ILLEGAL_CHARACTER", "An illegal character was encountered. For a multibyte character set, this includes a lead byte without a succeeding trail byte. For the Unicode character set, this includes the characters 0xFFFF and 0xFFFE."}
	ErrorUndefinedCharacter                                = &Error{0x00000247, "ERROR_UNDEFINED_CHARACTER", "The Unicode character is not defined in the Unicode character set installed on the system."}
	ErrorFloppyVolume                                      = &Error{0x00000248, "ERROR_FLOPPY_VOLUME", "The paging file cannot be created on a floppy disk."}
	ErrorBiosFailedToConnectInterrupt                      = &Error{0x00000249, "ERROR_BIOS_FAILED_TO_CONNECT_INTERRUPT", "The system bios failed to connect a system interrupt to the device or bus for which the device is connected."}
	ErrorBackupController                                  = &Error{0x0000024A, "ERROR_BACKUP_CONTROLLER", "This operation is only allowed for the primary domain controller (PDC) of the domain."}
	ErrorMutantLimitExceeded                               = &Error{0x0000024B, "ERROR_MUTANT_LIMIT_EXCEEDED", "An attempt was made to acquire a mutant such that its maximum count would have been exceeded."}
	ErrorFsDriverRequired                                  = &Error{0x0000024C, "ERROR_FS_DRIVER_REQUIRED", "A volume has been accessed for which a file system driver is required that has not yet been loaded."}
	ErrorCannotLoadRegistryFile                            = &Error{0x0000024D, "ERROR_CANNOT_LOAD_REGISTRY_FILE", "{Registry File Failure} The registry cannot load the hive (file): %hs or its log or alternate. It is corrupt, absent, or not writable."}
	ErrorDebugAttachFailed                                 = &Error{0x0000024E, "ERROR_DEBUG_ATTACH_FAILED", "{Unexpected Failure in DebugActiveProcess} An unexpected failure occurred while processing a DebugActiveProcess API request. Choosing OK will terminate the process, and choosing Cancel will ignore the error."}
	ErrorSystemProcessTerminated                           = &Error{0x0000024F, "ERROR_SYSTEM_PROCESS_TERMINATED", "{Fatal System Error} The %hs system process terminated unexpectedly with a status of 0x%08x (0x%08x 0x%08x). The system has been shut down."}
	ErrorDataNotAccepted                                   = &Error{0x00000250, "ERROR_DATA_NOT_ACCEPTED", "{Data Not Accepted} The transport driver interface (TDI) client could not handle the data received during an indication."}
	ErrorVdmHardError                                      = &Error{0x00000251, "ERROR_VDM_HARD_ERROR", "The NT Virtual DOS Machine (NTVDM) encountered a hard error."}
	ErrorDriverCancelTimeout                               = &Error{0x00000252, "ERROR_DRIVER_CANCEL_TIMEOUT", "{Cancel Timeout} The driver %hs failed to complete a canceled I/O request in the allotted time."}
	ErrorReplyMessageMismatch                              = &Error{0x00000253, "ERROR_REPLY_MESSAGE_MISMATCH", "{Reply Message Mismatch} An attempt was made to reply to a local procedure call (LPC) message, but the thread specified by the client ID in the message was not waiting on that message."}
	ErrorLostWritebehindData                               = &Error{0x00000254, "ERROR_LOST_WRITEBEHIND_DATA", "{Delayed Write Failed} Windows was unable to save all the data for the file %hs. The data has been lost. This error might be caused by a failure of your computer hardware or network connection. Try to save this file elsewhere."}
	ErrorClientServerParametersInvalid                     = &Error{0x00000255, "ERROR_CLIENT_SERVER_PARAMETERS_INVALID", "The parameters passed to the server in the client/server shared memory window were invalid. Too much data might have been put in the shared memory window."}
	ErrorNotTinyStream                                     = &Error{0x00000256, "ERROR_NOT_TINY_STREAM", "The stream is not a tiny stream."}
	ErrorStackOverflowRead                                 = &Error{0x00000257, "ERROR_STACK_OVERFLOW_READ", "The request must be handled by the stack overflow code."}
	ErrorConvertToLarge                                    = &Error{0x00000258, "ERROR_CONVERT_TO_LARGE", "Internal OFS status codes indicating how an allocation operation is handled. Either it is retried after the containing onode is moved or the extent stream is converted to a large stream."}
	ErrorFoundOutOfScope                                   = &Error{0x00000259, "ERROR_FOUND_OUT_OF_SCOPE", "The attempt to find the object found an object matching by ID on the volume but it is out of the scope of the handle used for the operation."}
	ErrorAllocateBucket                                    = &Error{0x0000025A, "ERROR_ALLOCATE_BUCKET", "The bucket array must be grown. Retry transaction after doing so."}
	ErrorMarshallOverflow                                  = &Error{0x0000025B, "ERROR_MARSHALL_OVERFLOW", "The user/kernel marshaling buffer has overflowed."}
	ErrorInvalidVariant                                    = &Error{0x0000025C, "ERROR_INVALID_VARIANT", "The supplied variant structure contains invalid data."}
	ErrorBadCompressionBuffer                              = &Error{0x0000025D, "ERROR_BAD_COMPRESSION_BUFFER", "The specified buffer contains ill-formed data."}
	ErrorAuditFailed                                       = &Error{0x0000025E, "ERROR_AUDIT_FAILED", "{Audit Failed} An attempt to generate a security audit failed."}
	ErrorTimerResolutionNotSet                             = &Error{0x0000025F, "ERROR_TIMER_RESOLUTION_NOT_SET", "The timer resolution was not previously set by the current process."}
	ErrorInsufficientLogonInfo                             = &Error{0x00000260, "ERROR_INSUFFICIENT_LOGON_INFO", "There is insufficient account information to log you on."}
	ErrorBadDllEntrypoint                                  = &Error{0x00000261, "ERROR_BAD_DLL_ENTRYPOINT", "{Invalid DLL Entrypoint} The dynamic link library %hs is not written correctly. The stack pointer has been left in an inconsistent state. The entry point should be declared as WINAPI or STDCALL. Select YES to fail the DLL load. Select NO to continue execution. Selecting NO can cause the application to operate incorrectly."}
	ErrorBadServiceEntrypoint                              = &Error{0x00000262, "ERROR_BAD_SERVICE_ENTRYPOINT", "{Invalid Service Callback Entrypoint} The %hs service is not written correctly. The stack pointer has been left in an inconsistent state. The callback entry point should be declared as WINAPI or STDCALL. Selecting OK will cause the service to continue operation. However, the service process might operate incorrectly."}
	ErrorIpAddressConflict1                                = &Error{0x00000263, "ERROR_IP_ADDRESS_CONFLICT1", "There is an IP address conflict with another system on the network."}
	ErrorIpAddressConflict2                                = &Error{0x00000264, "ERROR_IP_ADDRESS_CONFLICT2", "There is an IP address conflict with another system on the network."}
	ErrorRegistryQuotaLimit                                = &Error{0x00000265, "ERROR_REGISTRY_QUOTA_LIMIT", "{Low On Registry Space} The system has reached the maximum size allowed for the system part of the registry. Additional storage requests will be ignored."}
	ErrorNoCallbackActive                                  = &Error{0x00000266, "ERROR_NO_CALLBACK_ACTIVE", "A callback return system service cannot be executed when no callback is active."}
	ErrorPwdTooShort                                       = &Error{0x00000267, "ERROR_PWD_TOO_SHORT", "The password provided is too short to meet the policy of your user account. Choose a longer password."}
	ErrorPwdTooRecent                                      = &Error{0x00000268, "ERROR_PWD_TOO_RECENT", "The policy of your user account does not allow you to change passwords too frequently. This is done to prevent users from changing back to a familiar, but potentially discovered, password. If you feel your password has been compromised, contact your administrator immediately to have a new one assigned."}
	ErrorPwdHistoryConflict                                = &Error{0x00000269, "ERROR_PWD_HISTORY_CONFLICT", "You have attempted to change your password to one that you have used in the past. The policy of your user account does not allow this. Select a password that you have not previously used."}
	ErrorUnsupportedCompression                            = &Error{0x0000026A, "ERROR_UNSUPPORTED_COMPRESSION", "The specified compression format is unsupported."}
	ErrorInvalidHwProfile                                  = &Error{0x0000026B, "ERROR_INVALID_HW_PROFILE", "The specified hardware profile configuration is invalid."}
	ErrorInvalidPlugplayDevicePath                         = &Error{0x0000026C, "ERROR_INVALID_PLUGPLAY_DEVICE_PATH", "The specified Plug and Play registry device path is invalid."}
	ErrorQuotaListInconsistent                             = &Error{0x0000026D, "ERROR_QUOTA_LIST_INCONSISTENT", "The specified quota list is internally inconsistent with its descriptor."}
	ErrorEvaluationExpiration                              = &Error{0x0000026E, "ERROR_EVALUATION_EXPIRATION", "{Windows Evaluation Notification} The evaluation period for this installation of Windows has expired. This system will shut down in 1 hour. To restore access to this installation of Windows, upgrade this installation using a licensed distribution of this product."}
	ErrorIllegalDllRelocation                              = &Error{0x0000026F, "ERROR_ILLEGAL_DLL_RELOCATION", "{Illegal System DLL Relocation} The system DLL %hs was relocated in memory. The application will not run properly. The relocation occurred because the DLL %hs occupied an address range reserved for Windows system DLLs. The vendor supplying the DLL should be contacted for a new DLL."}
	ErrorDllInitFailedLogoff                               = &Error{0x00000270, "ERROR_DLL_INIT_FAILED_LOGOFF", "{DLL Initialization Failed} The application failed to initialize because the window station is shutting down."}
	ErrorValidateContinue                                  = &Error{0x00000271, "ERROR_VALIDATE_CONTINUE", "The validation process needs to continue on to the next step."}
	ErrorNoMoreMatches                                     = &Error{0x00000272, "ERROR_NO_MORE_MATCHES", "There are no more matches for the current index enumeration."}
	ErrorRangeListConflict                                 = &Error{0x00000273, "ERROR_RANGE_LIST_CONFLICT", "The range could not be added to the range list because of a conflict."}
	ErrorServerSidMismatch                                 = &Error{0x00000274, "ERROR_SERVER_SID_MISMATCH", "The server process is running under a SID different than that required by the client."}
	ErrorCantEnableDenyOnly                                = &Error{0x00000275, "ERROR_CANT_ENABLE_DENY_ONLY", "A group marked use for deny only cannot be enabled."}
	ErrorFloatMultipleFaults                               = &Error{0x00000276, "ERROR_FLOAT_MULTIPLE_FAULTS", "{EXCEPTION} Multiple floating point faults."}
	ErrorFloatMultipleTraps                                = &Error{0x00000277, "ERROR_FLOAT_MULTIPLE_TRAPS", "{EXCEPTION} Multiple floating point traps."}
	ErrorNointerface                                       = &Error{0x00000278, "ERROR_NOINTERFACE", "The requested interface is not supported."}
	ErrorDriverFailedSleep                                 = &Error{0x00000279, "ERROR_DRIVER_FAILED_SLEEP", "{System Standby Failed} The driver %hs does not support standby mode. Updating this driver might allow the system to go to standby mode."}
	ErrorCorruptSystemFile                                 = &Error{0x0000027A, "ERROR_CORRUPT_SYSTEM_FILE", "The system file %1 has become corrupt and has been replaced."}
	ErrorCommitmentMinimum                                 = &Error{0x0000027B, "ERROR_COMMITMENT_MINIMUM", "{Virtual Memory Minimum Too Low} Your system is low on virtual memory. Windows is increasing the size of your virtual memory paging file. During this process, memory requests for some applications might be denied. For more information, see Help."}
	ErrorPnpRestartEnumeration                             = &Error{0x0000027C, "ERROR_PNP_RESTART_ENUMERATION", "A device was removed so enumeration must be restarted."}
	ErrorSystemImageBadSignature                           = &Error{0x0000027D, "ERROR_SYSTEM_IMAGE_BAD_SIGNATURE", "{Fatal System Error} The system image %s is not properly signed. The file has been replaced with the signed file. The system has been shut down."}
	ErrorPnpRebootRequired                                 = &Error{0x0000027E, "ERROR_PNP_REBOOT_REQUIRED", "Device will not start without a reboot."}
	ErrorInsufficientPower                                 = &Error{0x0000027F, "ERROR_INSUFFICIENT_POWER", "There is not enough power to complete the requested operation."}
	ErrorSystemShutdown                                    = &Error{0x00000281, "ERROR_SYSTEM_SHUTDOWN", "The system is in the process of shutting down."}
	ErrorPortNotSet                                        = &Error{0x00000282, "ERROR_PORT_NOT_SET", "An attempt to remove a process DebugPort was made, but a port was not already associated with the process."}
	ErrorDsVersionCheckFailure                             = &Error{0x00000283, "ERROR_DS_VERSION_CHECK_FAILURE", "This version of Windows is not compatible with the behavior version of directory forest, domain, or domain controller."}
	ErrorRangeNotFound                                     = &Error{0x00000284, "ERROR_RANGE_NOT_FOUND", "The specified range could not be found in the range list."}
	ErrorNotSafeModeDriver                                 = &Error{0x00000286, "ERROR_NOT_SAFE_MODE_DRIVER", "The driver was not loaded because the system is booting into safe mode."}
	ErrorFailedDriverEntry                                 = &Error{0x00000287, "ERROR_FAILED_DRIVER_ENTRY", "The driver was not loaded because it failed its initialization call."}
	ErrorDeviceEnumerationError                            = &Error{0x00000288, "ERROR_DEVICE_ENUMERATION_ERROR", "The device encountered an error while applying power or reading the device configuration. This might be caused by a failure of your hardware or by a poor connection."}
	ErrorMountPointNotResolved                             = &Error{0x00000289, "ERROR_MOUNT_POINT_NOT_RESOLVED", "The create operation failed because the name contained at least one mount point that resolves to a volume to which the specified device object is not attached."}
	ErrorInvalidDeviceObjectParameter                      = &Error{0x0000028A, "ERROR_INVALID_DEVICE_OBJECT_PARAMETER", "The device object parameter is either not a valid device object or is not attached to the volume specified by the file name."}
	ErrorMcaOccured                                        = &Error{0x0000028B, "ERROR_MCA_OCCURED", "A machine check error has occurred. Check the system event log for additional information."}
	ErrorDriverDatabaseError                               = &Error{0x0000028C, "ERROR_DRIVER_DATABASE_ERROR", "There was an error [%2] processing the driver database."}
	ErrorSystemHiveTooLarge                                = &Error{0x0000028D, "ERROR_SYSTEM_HIVE_TOO_LARGE", "The system hive size has exceeded its limit."}
	ErrorDriverFailedPriorUnload                           = &Error{0x0000028E, "ERROR_DRIVER_FAILED_PRIOR_UNLOAD", "The driver could not be loaded because a previous version of the driver is still in memory."}
	ErrorVolsnapPrepareHibernate                           = &Error{0x0000028F, "ERROR_VOLSNAP_PREPARE_HIBERNATE", "{Volume Shadow Copy Service} Wait while the Volume Shadow Copy Service prepares volume %hs for hibernation."}
	ErrorHibernationFailure                                = &Error{0x00000290, "ERROR_HIBERNATION_FAILURE", "The system has failed to hibernate (the error code is %hs). Hibernation will be disabled until the system is restarted."}
	ErrorFileSystemLimitation                              = &Error{0x00000299, "ERROR_FILE_SYSTEM_LIMITATION", "The requested operation could not be completed due to a file system limitation."}
	ErrorAssertionFailure                                  = &Error{0x0000029C, "ERROR_ASSERTION_FAILURE", "An assertion failure has occurred."}
	ErrorAcpiError                                         = &Error{0x0000029D, "ERROR_ACPI_ERROR", "An error occurred in the Advanced Configuration and Power Interface (ACPI) subsystem."}
	ErrorWowAssertion                                      = &Error{0x0000029E, "ERROR_WOW_ASSERTION", "WOW assertion error."}
	ErrorPnpBadMpsTable                                    = &Error{0x0000029F, "ERROR_PNP_BAD_MPS_TABLE", "A device is missing in the system BIOS MultiProcessor Specification (MPS) table. This device will not be used. Contact your system vendor for system BIOS update."}
	ErrorPnpTranslationFailed                              = &Error{0x000002A0, "ERROR_PNP_TRANSLATION_FAILED", "A translator failed to translate resources."}
	ErrorPnpIrqTranslationFailed                           = &Error{0x000002A1, "ERROR_PNP_IRQ_TRANSLATION_FAILED", "An interrupt request (IRQ) translator failed to translate resources."}
	ErrorPnpInvalidId                                      = &Error{0x000002A2, "ERROR_PNP_INVALID_ID", "Driver %2 returned invalid ID for a child device (%3)."}
	ErrorWakeSystemDebugger                                = &Error{0x000002A3, "ERROR_WAKE_SYSTEM_DEBUGGER", "{Kernel Debugger Awakened} the system debugger was awakened by an interrupt."}
	ErrorHandlesClosed                                     = &Error{0x000002A4, "ERROR_HANDLES_CLOSED", "{Handles Closed} Handles to objects have been automatically closed because of the requested operation."}
	ErrorExtraneousInformation                             = &Error{0x000002A5, "ERROR_EXTRANEOUS_INFORMATION", "{Too Much Information} The specified ACL contained more information than was expected."}
	ErrorRxactCommitNecessary                              = &Error{0x000002A6, "ERROR_RXACT_COMMIT_NECESSARY", "This warning level status indicates that the transaction state already exists for the registry subtree, but that a transaction commit was previously aborted. The commit has NOT been completed, but it has not been rolled back either (so it can still be committed if desired)."}
	ErrorMediaCheck                                        = &Error{0x000002A7, "ERROR_MEDIA_CHECK", "{Media Changed} The media might have changed."}
	ErrorGuidSubstitutionMade                              = &Error{0x000002A8, "ERROR_GUID_SUBSTITUTION_MADE", "{GUID Substitution} During the translation of a GUID to a Windows SID, no administratively defined GUID prefix was found. A substitute prefix was used, which will not compromise system security. However, this might provide more restrictive access than intended."}
	ErrorStoppedOnSymlink                                  = &Error{0x000002A9, "ERROR_STOPPED_ON_SYMLINK", "The create operation stopped after reaching a symbolic link."}
	ErrorLongjump                                          = &Error{0x000002AA, "ERROR_LONGJUMP", "A long jump has been executed."}
	ErrorPlugplayQueryVetoed                               = &Error{0x000002AB, "ERROR_PLUGPLAY_QUERY_VETOED", "The Plug and Play query operation was not successful."}
	ErrorUnwindConsolidate                                 = &Error{0x000002AC, "ERROR_UNWIND_CONSOLIDATE", "A frame consolidation has been executed."}
	ErrorRegistryHiveRecovered                             = &Error{0x000002AD, "ERROR_REGISTRY_HIVE_RECOVERED", "{Registry Hive Recovered} Registry hive (file): %hs was corrupted and it has been recovered. Some data might have been lost."}
	ErrorDllMightBeInsecure                                = &Error{0x000002AE, "ERROR_DLL_MIGHT_BE_INSECURE", "The application is attempting to run executable code from the module %hs. This might be insecure. An alternative, %hs, is available. Should the application use the secure module %hs?"}
	ErrorDllMightBeIncompatible                            = &Error{0x000002AF, "ERROR_DLL_MIGHT_BE_INCOMPATIBLE", "The application is loading executable code from the module %hs. This is secure, but might be incompatible with previous releases of the operating system. An alternative, %hs, is available. Should the application use the secure module %hs?"}
	ErrorDbgExceptionNotHandled                            = &Error{0x000002B0, "ERROR_DBG_EXCEPTION_NOT_HANDLED", "Debugger did not handle the exception."}
	ErrorDbgReplyLater                                     = &Error{0x000002B1, "ERROR_DBG_REPLY_LATER", "Debugger will reply later."}
	ErrorDbgUnableToProvideHandle                          = &Error{0x000002B2, "ERROR_DBG_UNABLE_TO_PROVIDE_HANDLE", "Debugger cannot provide handle."}
	ErrorDbgTerminateThread                                = &Error{0x000002B3, "ERROR_DBG_TERMINATE_THREAD", "Debugger terminated thread."}
	ErrorDbgTerminateProcess                               = &Error{0x000002B4, "ERROR_DBG_TERMINATE_PROCESS", "Debugger terminated process."}
	ErrorDbgControlC                                       = &Error{0x000002B5, "ERROR_DBG_CONTROL_C", "Debugger got control C."}
	ErrorDbgPrintexceptionC                                = &Error{0x000002B6, "ERROR_DBG_PRINTEXCEPTION_C", "Debugger printed exception on control C."}
	ErrorDbgRipexception                                   = &Error{0x000002B7, "ERROR_DBG_RIPEXCEPTION", "Debugger received Routing Information Protocol (RIP) exception."}
	ErrorDbgControlBreak                                   = &Error{0x000002B8, "ERROR_DBG_CONTROL_BREAK", "Debugger received control break."}
	ErrorDbgCommandException                               = &Error{0x000002B9, "ERROR_DBG_COMMAND_EXCEPTION", "Debugger command communication exception."}
	ErrorObjectNameExists                                  = &Error{0x000002BA, "ERROR_OBJECT_NAME_EXISTS", "{Object Exists} An attempt was made to create an object and the object name already existed."}
	ErrorThreadWasSuspended                                = &Error{0x000002BB, "ERROR_THREAD_WAS_SUSPENDED", "{Thread Suspended} A thread termination occurred while the thread was suspended. The thread was resumed and termination proceeded."}
	ErrorImageNotAtBase                                    = &Error{0x000002BC, "ERROR_IMAGE_NOT_AT_BASE", "{Image Relocated} An image file could not be mapped at the address specified in the image file. Local fixes must be performed on this image."}
	ErrorRxactStateCreated                                 = &Error{0x000002BD, "ERROR_RXACT_STATE_CREATED", "This informational level status indicates that a specified registry subtree transaction state did not yet exist and had to be created."}
	ErrorSegmentNotification                               = &Error{0x000002BE, "ERROR_SEGMENT_NOTIFICATION", "{Segment Load} A virtual DOS machine (VDM) is loading, unloading, or moving an MS-DOS or Win16 program segment image. An exception is raised so a debugger can load, unload, or track symbols and breakpoints within these 16-bit segments."}
	ErrorBadCurrentDirectory                               = &Error{0x000002BF, "ERROR_BAD_CURRENT_DIRECTORY", "{Invalid Current Directory} The process cannot switch to the startup current directory %hs. Select OK to set current directory to %hs, or select CANCEL to exit."}
	ErrorFtReadRecoveryFromBackup                          = &Error{0x000002C0, "ERROR_FT_READ_RECOVERY_FROM_BACKUP", "{Redundant Read} To satisfy a read request, the NT fault-tolerant file system successfully read the requested data from a redundant copy. This was done because the file system encountered a failure on a member of the fault-tolerant volume, but it was unable to reassign the failing area of the device."}
	ErrorFtWriteRecovery                                   = &Error{0x000002C1, "ERROR_FT_WRITE_RECOVERY", "{Redundant Write} To satisfy a write request, the Windows NT operating system fault-tolerant file system successfully wrote a redundant copy of the information. This was done because the file system encountered a failure on a member of the fault-tolerant volume, but it was not able to reassign the failing area of the device."}
	ErrorImageMachineTypeMismatch                          = &Error{0x000002C2, "ERROR_IMAGE_MACHINE_TYPE_MISMATCH", "{Machine Type Mismatch} The image file %hs is valid, but is for a machine type other than the current machine. Select OK to continue, or CANCEL to fail the DLL load."}
	ErrorReceivePartial                                    = &Error{0x000002C3, "ERROR_RECEIVE_PARTIAL", "{Partial Data Received} The network transport returned partial data to its client. The remaining data will be sent later."}
	ErrorReceiveExpedited                                  = &Error{0x000002C4, "ERROR_RECEIVE_EXPEDITED", "{Expedited Data Received} The network transport returned data to its client that was marked as expedited by the remote system."}
	ErrorReceivePartialExpedited                           = &Error{0x000002C5, "ERROR_RECEIVE_PARTIAL_EXPEDITED", "{Partial Expedited Data Received} The network transport returned partial data to its client and this data was marked as expedited by the remote system. The remaining data will be sent later."}
	ErrorEventDone                                         = &Error{0x000002C6, "ERROR_EVENT_DONE", "{TDI Event Done} The TDI indication has completed successfully."}
	ErrorEventPending                                      = &Error{0x000002C7, "ERROR_EVENT_PENDING", "{TDI Event Pending} The TDI indication has entered the pending state."}
	ErrorCheckingFileSystem                                = &Error{0x000002C8, "ERROR_CHECKING_FILE_SYSTEM", "Checking file system on %wZ."}
	ErrorFatalAppExit                                      = &Error{0x000002C9, "ERROR_FATAL_APP_EXIT", "{Fatal Application Exit} %hs."}
	ErrorPredefinedHandle                                  = &Error{0x000002CA, "ERROR_PREDEFINED_HANDLE", "The specified registry key is referenced by a predefined handle."}
	ErrorWasUnlocked                                       = &Error{0x000002CB, "ERROR_WAS_UNLOCKED", "{Page Unlocked} The page protection of a locked page was changed to 'No Access' and the page was unlocked from memory and from the process."}
	ErrorWasLocked                                         = &Error{0x000002CD, "ERROR_WAS_LOCKED", "{Page Locked} One of the pages to lock was already locked."}
	ErrorAlreadyWin32                                      = &Error{0x000002CF, "ERROR_ALREADY_WIN32", "The value already corresponds with a Win 32 error code."}
	ErrorImageMachineTypeMismatchExe                       = &Error{0x000002D0, "ERROR_IMAGE_MACHINE_TYPE_MISMATCH_EXE", "{Machine Type Mismatch} The image file %hs is valid, but is for a machine type other than the current machine."}
	ErrorNoYieldPerformed                                  = &Error{0x000002D1, "ERROR_NO_YIELD_PERFORMED", "A yield execution was performed and no thread was available to run."}
	ErrorTimerResumeIgnored                                = &Error{0x000002D2, "ERROR_TIMER_RESUME_IGNORED", "The resume flag to a timer API was ignored."}
	ErrorArbitrationUnhandled                              = &Error{0x000002D3, "ERROR_ARBITRATION_UNHANDLED", "The arbiter has deferred arbitration of these resources to its parent."}
	ErrorCardbusNotSupported                               = &Error{0x000002D4, "ERROR_CARDBUS_NOT_SUPPORTED", "The inserted CardBus device cannot be started because of a configuration error on %hs\".\""}
	ErrorMpProcessorMismatch                               = &Error{0x000002D5, "ERROR_MP_PROCESSOR_MISMATCH", "The CPUs in this multiprocessor system are not all the same revision level. To use all processors the operating system restricts itself to the features of the least capable processor in the system. If problems occur with this system, contact the CPU manufacturer to see if this mix of processors is supported."}
	ErrorHibernated                                        = &Error{0x000002D6, "ERROR_HIBERNATED", "The system was put into hibernation."}
	ErrorResumeHibernation                                 = &Error{0x000002D7, "ERROR_RESUME_HIBERNATION", "The system was resumed from hibernation."}
	ErrorFirmwareUpdated                                   = &Error{0x000002D8, "ERROR_FIRMWARE_UPDATED", "Windows has detected that the system firmware (BIOS) was updated (previous firmware date = %2, current firmware date %3)."}
	ErrorDriversLeakingLockedPages                         = &Error{0x000002D9, "ERROR_DRIVERS_LEAKING_LOCKED_PAGES", "A device driver is leaking locked I/O pages, causing system degradation. The system has automatically enabled a tracking code to try and catch the culprit."}
	ErrorWakeSystem                                        = &Error{0x000002DA, "ERROR_WAKE_SYSTEM", "The system has awoken."}
	ErrorAbandonedWait0                                    = &Error{0x000002DF, "ERROR_ABANDONED_WAIT_0", "The call failed because the handle associated with it was closed."}
	ErrorElevationRequired                                 = &Error{0x000002E4, "ERROR_ELEVATION_REQUIRED", "The requested operation requires elevation."}
	ErrorReparse                                           = &Error{0x000002E5, "ERROR_REPARSE", "A reparse should be performed by the object manager because the name of the file resulted in a symbolic link."}
	ErrorOplockBreakInProgress                             = &Error{0x000002E6, "ERROR_OPLOCK_BREAK_IN_PROGRESS", "An open/create operation completed while an oplock break is underway."}
	ErrorVolumeMounted                                     = &Error{0x000002E7, "ERROR_VOLUME_MOUNTED", "A new volume has been mounted by a file system."}
	ErrorRxactCommitted                                    = &Error{0x000002E8, "ERROR_RXACT_COMMITTED", "This success level status indicates that the transaction state already exists for the registry subtree, but that a transaction commit was previously aborted. The commit has now been completed."}
	ErrorNotifyCleanup                                     = &Error{0x000002E9, "ERROR_NOTIFY_CLEANUP", "This indicates that a notify change request has been completed due to closing the handle which made the notify change request."}
	ErrorPrimaryTransportConnectFailed                     = &Error{0x000002EA, "ERROR_PRIMARY_TRANSPORT_CONNECT_FAILED", "{Connect Failure on Primary Transport} An attempt was made to connect to the remote server %hs on the primary transport, but the connection failed. The computer was able to connect on a secondary transport."}
	ErrorPageFaultTransition                               = &Error{0x000002EB, "ERROR_PAGE_FAULT_TRANSITION", "Page fault was a transition fault."}
	ErrorPageFaultDemandZero                               = &Error{0x000002EC, "ERROR_PAGE_FAULT_DEMAND_ZERO", "Page fault was a demand zero fault."}
	ErrorPageFaultCopyOnWrite                              = &Error{0x000002ED, "ERROR_PAGE_FAULT_COPY_ON_WRITE", "Page fault was a demand zero fault."}
	ErrorPageFaultGuardPage                                = &Error{0x000002EE, "ERROR_PAGE_FAULT_GUARD_PAGE", "Page fault was a demand zero fault."}
	ErrorPageFaultPagingFile                               = &Error{0x000002EF, "ERROR_PAGE_FAULT_PAGING_FILE", "Page fault was satisfied by reading from a secondary storage device."}
	ErrorCachePageLocked                                   = &Error{0x000002F0, "ERROR_CACHE_PAGE_LOCKED", "Cached page was locked during operation."}
	ErrorCrashDump                                         = &Error{0x000002F1, "ERROR_CRASH_DUMP", "Crash dump exists in paging file."}
	ErrorBufferAllZeros                                    = &Error{0x000002F2, "ERROR_BUFFER_ALL_ZEROS", "Specified buffer contains all zeros."}
	ErrorReparseObject                                     = &Error{0x000002F3, "ERROR_REPARSE_OBJECT", "A reparse should be performed by the object manager because the name of the file resulted in a symbolic link."}
	ErrorResourceRequirementsChanged                       = &Error{0x000002F4, "ERROR_RESOURCE_REQUIREMENTS_CHANGED", "The device has succeeded a query-stop and its resource requirements have changed."}
	ErrorTranslationComplete                               = &Error{0x000002F5, "ERROR_TRANSLATION_COMPLETE", "The translator has translated these resources into the global space and no further translations should be performed."}
	ErrorNothingToTerminate                                = &Error{0x000002F6, "ERROR_NOTHING_TO_TERMINATE", "A process being terminated has no threads to terminate."}
	ErrorProcessNotInJob                                   = &Error{0x000002F7, "ERROR_PROCESS_NOT_IN_JOB", "The specified process is not part of a job."}
	ErrorProcessInJob                                      = &Error{0x000002F8, "ERROR_PROCESS_IN_JOB", "The specified process is part of a job."}
	ErrorVolsnapHibernateReady                             = &Error{0x000002F9, "ERROR_VOLSNAP_HIBERNATE_READY", "{Volume Shadow Copy Service} The system is now ready for hibernation."}
	ErrorFsfilterOpCompletedSuccessfully                   = &Error{0x000002FA, "ERROR_FSFILTER_OP_COMPLETED_SUCCESSFULLY", "A file system or file system filter driver has successfully completed an FsFilter operation."}
	ErrorInterruptVectorAlreadyConnected                   = &Error{0x000002FB, "ERROR_INTERRUPT_VECTOR_ALREADY_CONNECTED", "The specified interrupt vector was already connected."}
	ErrorInterruptStillConnected                           = &Error{0x000002FC, "ERROR_INTERRUPT_STILL_CONNECTED", "The specified interrupt vector is still connected."}
	ErrorWaitForOplock                                     = &Error{0x000002FD, "ERROR_WAIT_FOR_OPLOCK", "An operation is blocked waiting for an oplock."}
	ErrorDbgExceptionHandled                               = &Error{0x000002FE, "ERROR_DBG_EXCEPTION_HANDLED", "Debugger handled exception."}
	ErrorDbgContinue                                       = &Error{0x000002FF, "ERROR_DBG_CONTINUE", "Debugger continued."}
	ErrorCallbackPopStack                                  = &Error{0x00000300, "ERROR_CALLBACK_POP_STACK", "An exception occurred in a user mode callback and the kernel callback frame should be removed."}
	ErrorCompressionDisabled                               = &Error{0x00000301, "ERROR_COMPRESSION_DISABLED", "Compression is disabled for this volume."}
	ErrorCantfetchbackwards                                = &Error{0x00000302, "ERROR_CANTFETCHBACKWARDS", "The data provider cannot fetch backward through a result set."}
	ErrorCantscrollbackwards                               = &Error{0x00000303, "ERROR_CANTSCROLLBACKWARDS", "The data provider cannot scroll backward through a result set."}
	ErrorRowsnotreleased                                   = &Error{0x00000304, "ERROR_ROWSNOTRELEASED", "The data provider requires that previously fetched data is released before asking for more data."}
	ErrorBadAccessorFlags                                  = &Error{0x00000305, "ERROR_BAD_ACCESSOR_FLAGS", "The data provider was not able to interpret the flags set for a column binding in an accessor."}
	ErrorErrorsEncountered                                 = &Error{0x00000306, "ERROR_ERRORS_ENCOUNTERED", "One or more errors occurred while processing the request."}
	ErrorNotCapable                                        = &Error{0x00000307, "ERROR_NOT_CAPABLE", "The implementation is not capable of performing the request."}
	ErrorRequestOutOfSequence                              = &Error{0x00000308, "ERROR_REQUEST_OUT_OF_SEQUENCE", "The client of a component requested an operation that is not valid given the state of the component instance."}
	ErrorVersionParseError                                 = &Error{0x00000309, "ERROR_VERSION_PARSE_ERROR", "A version number could not be parsed."}
	ErrorBadstartposition                                  = &Error{0x0000030A, "ERROR_BADSTARTPOSITION", "The iterator's start position is invalid."}
	ErrorMemoryHardware                                    = &Error{0x0000030B, "ERROR_MEMORY_HARDWARE", "The hardware has reported an uncorrectable memory error."}
	ErrorDiskRepairDisabled                                = &Error{0x0000030C, "ERROR_DISK_REPAIR_DISABLED", "The attempted operation required self-healing to be enabled."}
	ErrorInsufficientResourceForSpecifiedSharedSectionSize = &Error{0x0000030D, "ERROR_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE", "The Desktop heap encountered an error while allocating session memory. There is more information in the system event log."}
	ErrorSystemPowerstateTransition                        = &Error{0x0000030E, "ERROR_SYSTEM_POWERSTATE_TRANSITION", "The system power state is transitioning from %2 to %3."}
	ErrorSystemPowerstateComplexTransition                 = &Error{0x0000030F, "ERROR_SYSTEM_POWERSTATE_COMPLEX_TRANSITION", "The system power state is transitioning from %2 to %3 but could enter %4."}
	ErrorMcaException                                      = &Error{0x00000310, "ERROR_MCA_EXCEPTION", "A thread is getting dispatched with MCA EXCEPTION because of MCA."}
	ErrorAccessAuditByPolicy                               = &Error{0x00000311, "ERROR_ACCESS_AUDIT_BY_POLICY", "Access to %1 is monitored by policy rule %2."}
	ErrorAccessDisabledNoSaferUiByPolicy                   = &Error{0x00000312, "ERROR_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY", "Access to %1 has been restricted by your administrator by policy rule %2."}
	ErrorAbandonHiberfile                                  = &Error{0x00000313, "ERROR_ABANDON_HIBERFILE", "A valid hibernation file has been invalidated and should be abandoned."}
	ErrorLostWritebehindDataNetworkDisconnected            = &Error{0x00000314, "ERROR_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED", "{Delayed Write Failed} Windows was unable to save all the data for the file %hs; the data has been lost. This error can be caused by network connectivity issues. Try to save this file elsewhere."}
	ErrorLostWritebehindDataNetworkServerError             = &Error{0x00000315, "ERROR_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR", "{Delayed Write Failed} Windows was unable to save all the data for the file %hs; the data has been lost. This error was returned by the server on which the file exists. Try to save this file elsewhere."}
	ErrorLostWritebehindDataLocalDiskError                 = &Error{0x00000316, "ERROR_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR", "{Delayed Write Failed} Windows was unable to save all the data for the file %hs; the data has been lost. This error can be caused if the device has been removed or the media is write-protected."}
	ErrorEaAccessDenied                                    = &Error{0x000003E2, "ERROR_EA_ACCESS_DENIED", "Access to the extended attribute was denied."}
	ErrorOperationAborted                                  = &Error{0x000003E3, "ERROR_OPERATION_ABORTED", "The I/O operation has been aborted because of either a thread exit or an application request."}
	ErrorIoIncomplete                                      = &Error{0x000003E4, "ERROR_IO_INCOMPLETE", "Overlapped I/O event is not in a signaled state."}
	ErrorIoPending                                         = &Error{0x000003E5, "ERROR_IO_PENDING", "Overlapped I/O operation is in progress."}
	ErrorNoaccess                                          = &Error{0x000003E6, "ERROR_NOACCESS", "Invalid access to memory location."}
	ErrorSwaperror                                         = &Error{0x000003E7, "ERROR_SWAPERROR", "Error performing in-page operation."}
	ErrorStackOverflow                                     = &Error{0x000003E9, "ERROR_STACK_OVERFLOW", "Recursion too deep; the stack overflowed."}
	ErrorInvalidMessage                                    = &Error{0x000003EA, "ERROR_INVALID_MESSAGE", "The window cannot act on the sent message."}
	ErrorCanNotComplete                                    = &Error{0x000003EB, "ERROR_CAN_NOT_COMPLETE", "Cannot complete this function."}
	ErrorInvalidFlags                                      = &Error{0x000003EC, "ERROR_INVALID_FLAGS", "Invalid flags."}
	ErrorUnrecognizedVolume                                = &Error{0x000003ED, "ERROR_UNRECOGNIZED_VOLUME", "The volume does not contain a recognized file system. Be sure that all required file system drivers are loaded and that the volume is not corrupted."}
	ErrorFileInvalid                                       = &Error{0x000003EE, "ERROR_FILE_INVALID", "The volume for a file has been externally altered so that the opened file is no longer valid."}
	ErrorFullscreenMode                                    = &Error{0x000003EF, "ERROR_FULLSCREEN_MODE", "The requested operation cannot be performed in full-screen mode."}
	ErrorNoToken                                           = &Error{0x000003F0, "ERROR_NO_TOKEN", "An attempt was made to reference a token that does not exist."}
	ErrorBaddb                                             = &Error{0x000003F1, "ERROR_BADDB", "The configuration registry database is corrupt."}
	ErrorBadkey                                            = &Error{0x000003F2, "ERROR_BADKEY", "The configuration registry key is invalid."}
	ErrorCantopen                                          = &Error{0x000003F3, "ERROR_CANTOPEN", "The configuration registry key could not be opened."}
	ErrorCantread                                          = &Error{0x000003F4, "ERROR_CANTREAD", "The configuration registry key could not be read."}
	ErrorCantwrite                                         = &Error{0x000003F5, "ERROR_CANTWRITE", "The configuration registry key could not be written."}
	ErrorRegistryRecovered                                 = &Error{0x000003F6, "ERROR_REGISTRY_RECOVERED", "One of the files in the registry database had to be recovered by use of a log or alternate copy. The recovery was successful."}
	ErrorRegistryCorrupt                                   = &Error{0x000003F7, "ERROR_REGISTRY_CORRUPT", "The registry is corrupted. The structure of one of the files containing registry data is corrupted, or the system's memory image of the file is corrupted, or the file could not be recovered because the alternate copy or log was absent or corrupted."}
	ErrorRegistryIoFailed                                  = &Error{0x000003F8, "ERROR_REGISTRY_IO_FAILED", "An I/O operation initiated by the registry failed and cannot be recovered. The registry could not read in, write out, or flush one of the files that contain the system's image of the registry."}
	ErrorNotRegistryFile                                   = &Error{0x000003F9, "ERROR_NOT_REGISTRY_FILE", "The system attempted to load or restore a file into the registry, but the specified file is not in a registry file format."}
	ErrorKeyDeleted                                        = &Error{0x000003FA, "ERROR_KEY_DELETED", "Illegal operation attempted on a registry key that has been marked for deletion."}
	ErrorNoLogSpace                                        = &Error{0x000003FB, "ERROR_NO_LOG_SPACE", "System could not allocate the required space in a registry log."}
	ErrorKeyHasChildren                                    = &Error{0x000003FC, "ERROR_KEY_HAS_CHILDREN", "Cannot create a symbolic link in a registry key that already has subkeys or values."}
	ErrorChildMustBeVolatile                               = &Error{0x000003FD, "ERROR_CHILD_MUST_BE_VOLATILE", "Cannot create a stable subkey under a volatile parent key."}
	ErrorNotifyEnumDir                                     = &Error{0x000003FE, "ERROR_NOTIFY_ENUM_DIR", "A notify change request is being completed and the information is not being returned in the caller's buffer. The caller now needs to enumerate the files to find the changes."}
	ErrorDependentServicesRunning                          = &Error{0x0000041B, "ERROR_DEPENDENT_SERVICES_RUNNING", "A stop control has been sent to a service that other running services are dependent on."}
	ErrorInvalidServiceControl                             = &Error{0x0000041C, "ERROR_INVALID_SERVICE_CONTROL", "The requested control is not valid for this service."}
	ErrorServiceRequestTimeout                             = &Error{0x0000041D, "ERROR_SERVICE_REQUEST_TIMEOUT", "The service did not respond to the start or control request in a timely fashion."}
	ErrorServiceNoThread                                   = &Error{0x0000041E, "ERROR_SERVICE_NO_THREAD", "A thread could not be created for the service."}
	ErrorServiceDatabaseLocked                             = &Error{0x0000041F, "ERROR_SERVICE_DATABASE_LOCKED", "The service database is locked."}
	ErrorServiceAlreadyRunning                             = &Error{0x00000420, "ERROR_SERVICE_ALREADY_RUNNING", "An instance of the service is already running."}
	ErrorInvalidServiceAccount                             = &Error{0x00000421, "ERROR_INVALID_SERVICE_ACCOUNT", "The account name is invalid or does not exist, or the password is invalid for the account name specified."}
	ErrorServiceDisabled                                   = &Error{0x00000422, "ERROR_SERVICE_DISABLED", "The service cannot be started, either because it is disabled or because it has no enabled devices associated with it."}
	ErrorCircularDependency                                = &Error{0x00000423, "ERROR_CIRCULAR_DEPENDENCY", "Circular service dependency was specified."}
	ErrorServiceDoesNotExist                               = &Error{0x00000424, "ERROR_SERVICE_DOES_NOT_EXIST", "The specified service does not exist as an installed service."}
	ErrorServiceCannotAcceptCtrl                           = &Error{0x00000425, "ERROR_SERVICE_CANNOT_ACCEPT_CTRL", "The service cannot accept control messages at this time."}
	ErrorServiceNotActive                                  = &Error{0x00000426, "ERROR_SERVICE_NOT_ACTIVE", "The service has not been started."}
	ErrorFailedServiceControllerConnect                    = &Error{0x00000427, "ERROR_FAILED_SERVICE_CONTROLLER_CONNECT", "The service process could not connect to the service controller."}
	ErrorExceptionInService                                = &Error{0x00000428, "ERROR_EXCEPTION_IN_SERVICE", "An exception occurred in the service when handling the control request."}
	ErrorDatabaseDoesNotExist                              = &Error{0x00000429, "ERROR_DATABASE_DOES_NOT_EXIST", "The database specified does not exist."}
	ErrorServiceSpecificError                              = &Error{0x0000042A, "ERROR_SERVICE_SPECIFIC_ERROR", "The service has returned a service-specific error code."}
	ErrorProcessAborted                                    = &Error{0x0000042B, "ERROR_PROCESS_ABORTED", "The process terminated unexpectedly."}
	ErrorServiceDependencyFail                             = &Error{0x0000042C, "ERROR_SERVICE_DEPENDENCY_FAIL", "The dependency service or group failed to start."}
	ErrorServiceLogonFailed                                = &Error{0x0000042D, "ERROR_SERVICE_LOGON_FAILED", "The service did not start due to a logon failure."}
	ErrorServiceStartHang                                  = &Error{0x0000042E, "ERROR_SERVICE_START_HANG", "After starting, the service stopped responding in a start-pending state."}
	ErrorInvalidServiceLock                                = &Error{0x0000042F, "ERROR_INVALID_SERVICE_LOCK", "The specified service database lock is invalid."}
	ErrorServiceMarkedForDelete                            = &Error{0x00000430, "ERROR_SERVICE_MARKED_FOR_DELETE", "The specified service has been marked for deletion."}
	ErrorServiceExists                                     = &Error{0x00000431, "ERROR_SERVICE_EXISTS", "The specified service already exists."}
	ErrorAlreadyRunningLkg                                 = &Error{0x00000432, "ERROR_ALREADY_RUNNING_LKG", "The system is currently running with the last-known-good configuration."}
	ErrorServiceDependencyDeleted                          = &Error{0x00000433, "ERROR_SERVICE_DEPENDENCY_DELETED", "The dependency service does not exist or has been marked for deletion."}
	ErrorBootAlreadyAccepted                               = &Error{0x00000434, "ERROR_BOOT_ALREADY_ACCEPTED", "The current boot has already been accepted for use as the last-known-good control set."}
	ErrorServiceNeverStarted                               = &Error{0x00000435, "ERROR_SERVICE_NEVER_STARTED", "No attempts to start the service have been made since the last boot."}
	ErrorDuplicateServiceName                              = &Error{0x00000436, "ERROR_DUPLICATE_SERVICE_NAME", "The name is already in use as either a service name or a service display name."}
	ErrorDifferentServiceAccount                           = &Error{0x00000437, "ERROR_DIFFERENT_SERVICE_ACCOUNT", "The account specified for this service is different from the account specified for other services running in the same process."}
	ErrorCannotDetectDriverFailure                         = &Error{0x00000438, "ERROR_CANNOT_DETECT_DRIVER_FAILURE", "Failure actions can only be set for Win32 services, not for drivers."}
	ErrorCannotDetectProcessAbort                          = &Error{0x00000439, "ERROR_CANNOT_DETECT_PROCESS_ABORT", "This service runs in the same process as the service control manager. Therefore, the service control manager cannot take action if this service's process terminates unexpectedly."}
	ErrorNoRecoveryProgram                                 = &Error{0x0000043A, "ERROR_NO_RECOVERY_PROGRAM", "No recovery program has been configured for this service."}
	ErrorServiceNotInExe                                   = &Error{0x0000043B, "ERROR_SERVICE_NOT_IN_EXE", "The executable program that this service is configured to run in does not implement the service."}
	ErrorNotSafebootService                                = &Error{0x0000043C, "ERROR_NOT_SAFEBOOT_SERVICE", "This service cannot be started in Safe Mode."}
	ErrorEndOfMedia                                        = &Error{0x0000044C, "ERROR_END_OF_MEDIA", "The physical end of the tape has been reached."}
	ErrorFilemarkDetected                                  = &Error{0x0000044D, "ERROR_FILEMARK_DETECTED", "A tape access reached a filemark."}
	ErrorBeginningOfMedia                                  = &Error{0x0000044E, "ERROR_BEGINNING_OF_MEDIA", "The beginning of the tape or a partition was encountered."}
	ErrorSetmarkDetected                                   = &Error{0x0000044F, "ERROR_SETMARK_DETECTED", "A tape access reached the end of a set of files."}
	ErrorNoDataDetected                                    = &Error{0x00000450, "ERROR_NO_DATA_DETECTED", "No more data is on the tape."}
	ErrorPartitionFailure                                  = &Error{0x00000451, "ERROR_PARTITION_FAILURE", "Tape could not be partitioned."}
	ErrorInvalidBlockLength                                = &Error{0x00000452, "ERROR_INVALID_BLOCK_LENGTH", "When accessing a new tape of a multivolume partition, the current block size is incorrect."}
	ErrorDeviceNotPartitioned                              = &Error{0x00000453, "ERROR_DEVICE_NOT_PARTITIONED", "Tape partition information could not be found when loading a tape."}
	ErrorUnableToLockMedia                                 = &Error{0x00000454, "ERROR_UNABLE_TO_LOCK_MEDIA", "Unable to lock the media eject mechanism."}
	ErrorUnableToUnloadMedia                               = &Error{0x00000455, "ERROR_UNABLE_TO_UNLOAD_MEDIA", "Unable to unload the media."}
	ErrorMediaChanged                                      = &Error{0x00000456, "ERROR_MEDIA_CHANGED", "The media in the drive might have changed."}
	ErrorBusReset                                          = &Error{0x00000457, "ERROR_BUS_RESET", "The I/O bus was reset."}
	ErrorNoMediaInDrive                                    = &Error{0x00000458, "ERROR_NO_MEDIA_IN_DRIVE", "No media in drive."}
	ErrorNoUnicodeTranslation                              = &Error{0x00000459, "ERROR_NO_UNICODE_TRANSLATION", "No mapping for the Unicode character exists in the target multibyte code page."}
	ErrorDllInitFailed                                     = &Error{0x0000045A, "ERROR_DLL_INIT_FAILED", "A DLL initialization routine failed."}
	ErrorShutdownInProgress                                = &Error{0x0000045B, "ERROR_SHUTDOWN_IN_PROGRESS", "A system shutdown is in progress."}
	ErrorNoShutdownInProgress                              = &Error{0x0000045C, "ERROR_NO_SHUTDOWN_IN_PROGRESS", "Unable to abort the system shutdown because no shutdown was in progress."}
	ErrorIoDevice                                          = &Error{0x0000045D, "ERROR_IO_DEVICE", "The request could not be performed because of an I/O device error."}
	ErrorSerialNoDevice                                    = &Error{0x0000045E, "ERROR_SERIAL_NO_DEVICE", "No serial device was successfully initialized. The serial driver will unload."}
	ErrorIrqBusy                                           = &Error{0x0000045F, "ERROR_IRQ_BUSY", "Unable to open a device that was sharing an IRQ with other devices. At least one other device that uses that IRQ was already opened."}
	ErrorMoreWrites                                        = &Error{0x00000460, "ERROR_MORE_WRITES", "A serial I/O operation was completed by another write to the serial port. (The IOCTL_SERIAL_XOFF_COUNTER reached zero.)"}
	ErrorCounterTimeout                                    = &Error{0x00000461, "ERROR_COUNTER_TIMEOUT", "A serial I/O operation completed because the time-out period expired. (The IOCTL_SERIAL_XOFF_COUNTER did not reach zero.)"}
	ErrorFloppyIdMarkNotFound                              = &Error{0x00000462, "ERROR_FLOPPY_ID_MARK_NOT_FOUND", "No ID address mark was found on the floppy disk."}
	ErrorFloppyWrongCylinder                               = &Error{0x00000463, "ERROR_FLOPPY_WRONG_CYLINDER", "Mismatch between the floppy disk sector ID field and the floppy disk controller track address."}
	ErrorFloppyUnknownError                                = &Error{0x00000464, "ERROR_FLOPPY_UNKNOWN_ERROR", "The floppy disk controller reported an error that is not recognized by the floppy disk driver."}
	ErrorFloppyBadRegisters                                = &Error{0x00000465, "ERROR_FLOPPY_BAD_REGISTERS", "The floppy disk controller returned inconsistent results in its registers."}
	ErrorDiskRecalibrateFailed                             = &Error{0x00000466, "ERROR_DISK_RECALIBRATE_FAILED", "While accessing the hard disk, a recalibrate operation failed, even after retries."}
	ErrorDiskOperationFailed                               = &Error{0x00000467, "ERROR_DISK_OPERATION_FAILED", "While accessing the hard disk, a disk operation failed even after retries."}
	ErrorDiskResetFailed                                   = &Error{0x00000468, "ERROR_DISK_RESET_FAILED", "While accessing the hard disk, a disk controller reset was needed, but that also failed."}
	ErrorEomOverflow                                       = &Error{0x00000469, "ERROR_EOM_OVERFLOW", "Physical end of tape encountered."}
	ErrorNotEnoughServerMemory                             = &Error{0x0000046A, "ERROR_NOT_ENOUGH_SERVER_MEMORY", "Not enough server storage is available to process this command."}
	ErrorPossibleDeadlock                                  = &Error{0x0000046B, "ERROR_POSSIBLE_DEADLOCK", "A potential deadlock condition has been detected."}
	ErrorMappedAlignment                                   = &Error{0x0000046C, "ERROR_MAPPED_ALIGNMENT", "The base address or the file offset specified does not have the proper alignment."}
	ErrorSetPowerStateVetoed                               = &Error{0x00000474, "ERROR_SET_POWER_STATE_VETOED", "An attempt to change the system power state was vetoed by another application or driver."}
	ErrorSetPowerStateFailed                               = &Error{0x00000475, "ERROR_SET_POWER_STATE_FAILED", "The system BIOS failed an attempt to change the system power state."}
	ErrorTooManyLinks                                      = &Error{0x00000476, "ERROR_TOO_MANY_LINKS", "An attempt was made to create more links on a file than the file system supports."}
	ErrorOldWinVersion                                     = &Error{0x0000047E, "ERROR_OLD_WIN_VERSION", "The specified program requires a newer version of Windows."}
	ErrorAppWrongOs                                        = &Error{0x0000047F, "ERROR_APP_WRONG_OS", "The specified program is not a Windows or MS-DOS program."}
	ErrorSingleInstanceApp                                 = &Error{0x00000480, "ERROR_SINGLE_INSTANCE_APP", "Cannot start more than one instance of the specified program."}
	ErrorRmodeApp                                          = &Error{0x00000481, "ERROR_RMODE_APP", "The specified program was written for an earlier version of Windows."}
	ErrorInvalidDll                                        = &Error{0x00000482, "ERROR_INVALID_DLL", "One of the library files needed to run this application is damaged."}
	ErrorNoAssociation                                     = &Error{0x00000483, "ERROR_NO_ASSOCIATION", "No application is associated with the specified file for this operation."}
	ErrorDdeFail                                           = &Error{0x00000484, "ERROR_DDE_FAIL", "An error occurred in sending the command to the application."}
	ErrorDllNotFound                                       = &Error{0x00000485, "ERROR_DLL_NOT_FOUND", "One of the library files needed to run this application cannot be found."}
	ErrorNoMoreUserHandles                                 = &Error{0x00000486, "ERROR_NO_MORE_USER_HANDLES", "The current process has used all of its system allowance of handles for Windows manager objects."}
	ErrorMessageSyncOnly                                   = &Error{0x00000487, "ERROR_MESSAGE_SYNC_ONLY", "The message can be used only with synchronous operations."}
	ErrorSourceElementEmpty                                = &Error{0x00000488, "ERROR_SOURCE_ELEMENT_EMPTY", "The indicated source element has no media."}
	ErrorDestinationElementFull                            = &Error{0x00000489, "ERROR_DESTINATION_ELEMENT_FULL", "The indicated destination element already contains media."}
	ErrorIllegalElementAddress                             = &Error{0x0000048A, "ERROR_ILLEGAL_ELEMENT_ADDRESS", "The indicated element does not exist."}
	ErrorMagazineNotPresent                                = &Error{0x0000048B, "ERROR_MAGAZINE_NOT_PRESENT", "The indicated element is part of a magazine that is not present."}
	ErrorDeviceReinitializationNeeded                      = &Error{0x0000048C, "ERROR_DEVICE_REINITIALIZATION_NEEDED", "The indicated device requires re-initialization due to hardware errors."}
	ErrorDeviceRequiresCleaning                            = &Error{0x0000048D, "ERROR_DEVICE_REQUIRES_CLEANING", "The device has indicated that cleaning is required before further operations are attempted."}
	ErrorDeviceDoorOpen                                    = &Error{0x0000048E, "ERROR_DEVICE_DOOR_OPEN", "The device has indicated that its door is open."}
	ErrorDeviceNotConnected                                = &Error{0x0000048F, "ERROR_DEVICE_NOT_CONNECTED", "The device is not connected."}
	ErrorNotFound                                          = &Error{0x00000490, "ERROR_NOT_FOUND", "Element not found."}
	ErrorNoMatch                                           = &Error{0x00000491, "ERROR_NO_MATCH", "There was no match for the specified key in the index."}
	ErrorSetNotFound                                       = &Error{0x00000492, "ERROR_SET_NOT_FOUND", "The property set specified does not exist on the object."}
	ErrorPointNotFound                                     = &Error{0x00000493, "ERROR_POINT_NOT_FOUND", "The point passed to GetMouseMovePoints is not in the buffer."}
	ErrorNoTrackingService                                 = &Error{0x00000494, "ERROR_NO_TRACKING_SERVICE", "The tracking (workstation) service is not running."}
	ErrorNoVolumeId                                        = &Error{0x00000495, "ERROR_NO_VOLUME_ID", "The volume ID could not be found."}
	ErrorUnableToRemoveReplaced                            = &Error{0x00000497, "ERROR_UNABLE_TO_REMOVE_REPLACED", "Unable to remove the file to be replaced."}
	ErrorUnableToMoveReplacement                           = &Error{0x00000498, "ERROR_UNABLE_TO_MOVE_REPLACEMENT", "Unable to move the replacement file to the file to be replaced. The file to be replaced has retained its original name."}
	ErrorUnableToMoveReplacement2                          = &Error{0x00000499, "ERROR_UNABLE_TO_MOVE_REPLACEMENT_2", "Unable to move the replacement file to the file to be replaced. The file to be replaced has been renamed using the backup name."}
	ErrorJournalDeleteInProgress                           = &Error{0x0000049A, "ERROR_JOURNAL_DELETE_IN_PROGRESS", "The volume change journal is being deleted."}
	ErrorJournalNotActive                                  = &Error{0x0000049B, "ERROR_JOURNAL_NOT_ACTIVE", "The volume change journal is not active."}
	ErrorPotentialFileFound                                = &Error{0x0000049C, "ERROR_POTENTIAL_FILE_FOUND", "A file was found, but it might not be the correct file."}
	ErrorJournalEntryDeleted                               = &Error{0x0000049D, "ERROR_JOURNAL_ENTRY_DELETED", "The journal entry has been deleted from the journal."}
	ErrorShutdownIsScheduled                               = &Error{0x000004A6, "ERROR_SHUTDOWN_IS_SCHEDULED", "A system shutdown has already been scheduled."}
	ErrorShutdownUsersLoggedOn                             = &Error{0x000004A7, "ERROR_SHUTDOWN_USERS_LOGGED_ON", "The system shutdown cannot be initiated because there are other users logged on to the computer."}
	ErrorBadDevice                                         = &Error{0x000004B0, "ERROR_BAD_DEVICE", "The specified device name is invalid."}
	ErrorConnectionUnavail                                 = &Error{0x000004B1, "ERROR_CONNECTION_UNAVAIL", "The device is not currently connected but it is a remembered connection."}
	ErrorDeviceAlreadyRemembered                           = &Error{0x000004B2, "ERROR_DEVICE_ALREADY_REMEMBERED", "The local device name has a remembered connection to another network resource."}
	ErrorNoNetOrBadPath                                    = &Error{0x000004B3, "ERROR_NO_NET_OR_BAD_PATH", "The network path was either typed incorrectly, does not exist, or the network provider is not currently available. Try retyping the path or contact your network administrator."}
	ErrorBadProvider                                       = &Error{0x000004B4, "ERROR_BAD_PROVIDER", "The specified network provider name is invalid."}
	ErrorCannotOpenProfile                                 = &Error{0x000004B5, "ERROR_CANNOT_OPEN_PROFILE", "Unable to open the network connection profile."}
	ErrorBadProfile                                        = &Error{0x000004B6, "ERROR_BAD_PROFILE", "The network connection profile is corrupted."}
	ErrorNotContainer                                      = &Error{0x000004B7, "ERROR_NOT_CONTAINER", "Cannot enumerate a noncontainer."}
	ErrorExtendedError                                     = &Error{0x000004B8, "ERROR_EXTENDED_ERROR", "An extended error has occurred."}
	ErrorInvalidGroupname                                  = &Error{0x000004B9, "ERROR_INVALID_GROUPNAME", "The format of the specified group name is invalid."}
	ErrorInvalidComputername                               = &Error{0x000004BA, "ERROR_INVALID_COMPUTERNAME", "The format of the specified computer name is invalid."}
	ErrorInvalidEventname                                  = &Error{0x000004BB, "ERROR_INVALID_EVENTNAME", "The format of the specified event name is invalid."}
	ErrorInvalidDomainname                                 = &Error{0x000004BC, "ERROR_INVALID_DOMAINNAME", "The format of the specified domain name is invalid."}
	ErrorInvalidServicename                                = &Error{0x000004BD, "ERROR_INVALID_SERVICENAME", "The format of the specified service name is invalid."}
	ErrorInvalidNetname                                    = &Error{0x000004BE, "ERROR_INVALID_NETNAME", "The format of the specified network name is invalid."}
	ErrorInvalidSharename                                  = &Error{0x000004BF, "ERROR_INVALID_SHARENAME", "The format of the specified share name is invalid."}
	ErrorInvalidPasswordname                               = &Error{0x000004C0, "ERROR_INVALID_PASSWORDNAME", "The format of the specified password is invalid."}
	ErrorInvalidMessagename                                = &Error{0x000004C1, "ERROR_INVALID_MESSAGENAME", "The format of the specified message name is invalid."}
	ErrorInvalidMessagedest                                = &Error{0x000004C2, "ERROR_INVALID_MESSAGEDEST", "The format of the specified message destination is invalid."}
	ErrorSessionCredentialConflict                         = &Error{0x000004C3, "ERROR_SESSION_CREDENTIAL_CONFLICT", "Multiple connections to a server or shared resource by the same user, using more than one user name, are not allowed. Disconnect all previous connections to the server or shared resource and try again."}
	ErrorRemoteSessionLimitExceeded                        = &Error{0x000004C4, "ERROR_REMOTE_SESSION_LIMIT_EXCEEDED", "An attempt was made to establish a session to a network server, but there are already too many sessions established to that server."}
	ErrorDupDomainname                                     = &Error{0x000004C5, "ERROR_DUP_DOMAINNAME", "The workgroup or domain name is already in use by another computer on the network."}
	ErrorNoNetwork                                         = &Error{0x000004C6, "ERROR_NO_NETWORK", "The network is not present or not started."}
	ErrorCancelled                                         = &Error{0x000004C7, "ERROR_CANCELLED", "The operation was canceled by the user."}
	ErrorUserMappedFile                                    = &Error{0x000004C8, "ERROR_USER_MAPPED_FILE", "The requested operation cannot be performed on a file with a user-mapped section open."}
	ErrorConnectionRefused                                 = &Error{0x000004C9, "ERROR_CONNECTION_REFUSED", "The remote system refused the network connection."}
	ErrorGracefulDisconnect                                = &Error{0x000004CA, "ERROR_GRACEFUL_DISCONNECT", "The network connection was gracefully closed."}
	ErrorAddressAlreadyAssociated                          = &Error{0x000004CB, "ERROR_ADDRESS_ALREADY_ASSOCIATED", "The network transport endpoint already has an address associated with it."}
	ErrorAddressNotAssociated                              = &Error{0x000004CC, "ERROR_ADDRESS_NOT_ASSOCIATED", "An address has not yet been associated with the network endpoint."}
	ErrorConnectionInvalid                                 = &Error{0x000004CD, "ERROR_CONNECTION_INVALID", "An operation was attempted on a nonexistent network connection."}
	ErrorConnectionActive                                  = &Error{0x000004CE, "ERROR_CONNECTION_ACTIVE", "An invalid operation was attempted on an active network connection."}
	ErrorNetworkUnreachable                                = &Error{0x000004CF, "ERROR_NETWORK_UNREACHABLE", "The network location cannot be reached. For information about network troubleshooting, see Windows Help."}
	ErrorHostUnreachable                                   = &Error{0x000004D0, "ERROR_HOST_UNREACHABLE", "The network location cannot be reached. For information about network troubleshooting, see Windows Help."}
	ErrorProtocolUnreachable                               = &Error{0x000004D1, "ERROR_PROTOCOL_UNREACHABLE", "The network location cannot be reached. For information about network troubleshooting, see Windows Help."}
	ErrorPortUnreachable                                   = &Error{0x000004D2, "ERROR_PORT_UNREACHABLE", "No service is operating at the destination network endpoint on the remote system."}
	ErrorRequestAborted                                    = &Error{0x000004D3, "ERROR_REQUEST_ABORTED", "The request was aborted."}
	ErrorConnectionAborted                                 = &Error{0x000004D4, "ERROR_CONNECTION_ABORTED", "The network connection was aborted by the local system."}
	ErrorRetry                                             = &Error{0x000004D5, "ERROR_RETRY", "The operation could not be completed. A retry should be performed."}
	ErrorConnectionCountLimit                              = &Error{0x000004D6, "ERROR_CONNECTION_COUNT_LIMIT", "A connection to the server could not be made because the limit on the number of concurrent connections for this account has been reached."}
	ErrorLoginTimeRestriction                              = &Error{0x000004D7, "ERROR_LOGIN_TIME_RESTRICTION", "Attempting to log on during an unauthorized time of day for this account."}
	ErrorLoginWkstaRestriction                             = &Error{0x000004D8, "ERROR_LOGIN_WKSTA_RESTRICTION", "The account is not authorized to log on from this station."}
	ErrorIncorrectAddress                                  = &Error{0x000004D9, "ERROR_INCORRECT_ADDRESS", "The network address could not be used for the operation requested."}
	ErrorAlreadyRegistered                                 = &Error{0x000004DA, "ERROR_ALREADY_REGISTERED", "The service is already registered."}
	ErrorServiceNotFound                                   = &Error{0x000004DB, "ERROR_SERVICE_NOT_FOUND", "The specified service does not exist."}
	ErrorNotAuthenticated                                  = &Error{0x000004DC, "ERROR_NOT_AUTHENTICATED", "The operation being requested was not performed because the user has not been authenticated."}
	ErrorNotLoggedOn                                       = &Error{0x000004DD, "ERROR_NOT_LOGGED_ON", "The operation being requested was not performed because the user has not logged on to the network. The specified service does not exist."}
	ErrorContinue                                          = &Error{0x000004DE, "ERROR_CONTINUE", "Continue with work in progress."}
	ErrorAlreadyInitialized                                = &Error{0x000004DF, "ERROR_ALREADY_INITIALIZED", "An attempt was made to perform an initialization operation when initialization has already been completed."}
	ErrorNoMoreDevices                                     = &Error{0x000004E0, "ERROR_NO_MORE_DEVICES", "No more local devices."}
	ErrorNoSuchSite                                        = &Error{0x000004E1, "ERROR_NO_SUCH_SITE", "The specified site does not exist."}
	ErrorDomainControllerExists                            = &Error{0x000004E2, "ERROR_DOMAIN_CONTROLLER_EXISTS", "A domain controller with the specified name already exists."}
	ErrorOnlyIfConnected                                   = &Error{0x000004E3, "ERROR_ONLY_IF_CONNECTED", "This operation is supported only when you are connected to the server."}
	ErrorOverrideNochanges                                 = &Error{0x000004E4, "ERROR_OVERRIDE_NOCHANGES", "The group policy framework should call the extension even if there are no changes."}
	ErrorBadUserProfile                                    = &Error{0x000004E5, "ERROR_BAD_USER_PROFILE", "The specified user does not have a valid profile."}
	ErrorNotSupportedOnSbs                                 = &Error{0x000004E6, "ERROR_NOT_SUPPORTED_ON_SBS", "This operation is not supported on a computer running Windows Server 2003 operating system for Small Business Server."}
	ErrorServerShutdownInProgress                          = &Error{0x000004E7, "ERROR_SERVER_SHUTDOWN_IN_PROGRESS", "The server machine is shutting down."}
	ErrorHostDown                                          = &Error{0x000004E8, "ERROR_HOST_DOWN", "The remote system is not available. For information about network troubleshooting, see Windows Help."}
	ErrorNonAccountSid                                     = &Error{0x000004E9, "ERROR_NON_ACCOUNT_SID", "The security identifier provided is not from an account domain."}
	ErrorNonDomainSid                                      = &Error{0x000004EA, "ERROR_NON_DOMAIN_SID", "The security identifier provided does not have a domain component."}
	ErrorApphelpBlock                                      = &Error{0x000004EB, "ERROR_APPHELP_BLOCK", "AppHelp dialog canceled, thus preventing the application from starting."}
	ErrorAccessDisabledByPolicy                            = &Error{0x000004EC, "ERROR_ACCESS_DISABLED_BY_POLICY", "This program is blocked by Group Policy. For more information, contact your system administrator."}
	ErrorRegNatConsumption                                 = &Error{0x000004ED, "ERROR_REG_NAT_CONSUMPTION", "A program attempt to use an invalid register value. Normally caused by an uninitialized register. This error is Itanium specific."}
	ErrorCscshareOffline                                   = &Error{0x000004EE, "ERROR_CSCSHARE_OFFLINE", "The share is currently offline or does not exist."}
	ErrorPkinitFailure                                     = &Error{0x000004EF, "ERROR_PKINIT_FAILURE", "The Kerberos protocol encountered an error while validating the KDC certificate during smartcard logon. There is more information in the system event log."}
	ErrorSmartcardSubsystemFailure                         = &Error{0x000004F0, "ERROR_SMARTCARD_SUBSYSTEM_FAILURE", "The Kerberos protocol encountered an error while attempting to utilize the smartcard subsystem."}
	ErrorDowngradeDetected                                 = &Error{0x000004F1, "ERROR_DOWNGRADE_DETECTED", "The system detected a possible attempt to compromise security. Ensure that you can contact the server that authenticated you."}
	ErrorMachineLocked                                     = &Error{0x000004F7, "ERROR_MACHINE_LOCKED", "The machine is locked and cannot be shut down without the force option."}
	ErrorCallbackSuppliedInvalidData                       = &Error{0x000004F9, "ERROR_CALLBACK_SUPPLIED_INVALID_DATA", "An application-defined callback gave invalid data when called."}
	ErrorSyncForegroundRefreshRequired                     = &Error{0x000004FA, "ERROR_SYNC_FOREGROUND_REFRESH_REQUIRED", "The Group Policy framework should call the extension in the synchronous foreground policy refresh."}
	ErrorDriverBlocked                                     = &Error{0x000004FB, "ERROR_DRIVER_BLOCKED", "This driver has been blocked from loading."}
	ErrorInvalidImportOfNonDll                             = &Error{0x000004FC, "ERROR_INVALID_IMPORT_OF_NON_DLL", "A DLL referenced a module that was neither a DLL nor the process's executable image."}
	ErrorAccessDisabledWebblade                            = &Error{0x000004FD, "ERROR_ACCESS_DISABLED_WEBBLADE", "Windows cannot open this program because it has been disabled."}
	ErrorAccessDisabledWebbladeTamper                      = &Error{0x000004FE, "ERROR_ACCESS_DISABLED_WEBBLADE_TAMPER", "Windows cannot open this program because the license enforcement system has been tampered with or become corrupted."}
	ErrorRecoveryFailure                                   = &Error{0x000004FF, "ERROR_RECOVERY_FAILURE", "A transaction recover failed."}
	ErrorAlreadyFiber                                      = &Error{0x00000500, "ERROR_ALREADY_FIBER", "The current thread has already been converted to a fiber."}
	ErrorAlreadyThread                                     = &Error{0x00000501, "ERROR_ALREADY_THREAD", "The current thread has already been converted from a fiber."}
	ErrorStackBufferOverrun                                = &Error{0x00000502, "ERROR_STACK_BUFFER_OVERRUN", "The system detected an overrun of a stack-based buffer in this application. This overrun could potentially allow a malicious user to gain control of this application."}
	ErrorParameterQuotaExceeded                            = &Error{0x00000503, "ERROR_PARAMETER_QUOTA_EXCEEDED", "Data present in one of the parameters is more than the function can operate on."}
	ErrorDebuggerInactive                                  = &Error{0x00000504, "ERROR_DEBUGGER_INACTIVE", "An attempt to perform an operation on a debug object failed because the object is in the process of being deleted."}
	ErrorDelayLoadFailed                                   = &Error{0x00000505, "ERROR_DELAY_LOAD_FAILED", "An attempt to delay-load a .dll or get a function address in a delay-loaded .dll failed."}
	ErrorVdmDisallowed                                     = &Error{0x00000506, "ERROR_VDM_DISALLOWED", "%1 is a 16-bit application. You do not have permissions to execute 16-bit applications. Check your permissions with your system administrator."}
	ErrorUnidentifiedError                                 = &Error{0x00000507, "ERROR_UNIDENTIFIED_ERROR", "Insufficient information exists to identify the cause of failure."}
	ErrorInvalidCruntimeParameter                          = &Error{0x00000508, "ERROR_INVALID_CRUNTIME_PARAMETER", "The parameter passed to a C runtime function is incorrect."}
	ErrorBeyondVdl                                         = &Error{0x00000509, "ERROR_BEYOND_VDL", "The operation occurred beyond the valid data length of the file."}
	ErrorIncompatibleServiceSidType                        = &Error{0x0000050A, "ERROR_INCOMPATIBLE_SERVICE_SID_TYPE", "The service start failed because one or more services in the same process have an incompatible service SID type setting. A service with a restricted service SID type can only coexist in the same process with other services with a restricted SID type."}
	ErrorDriverProcessTerminated                           = &Error{0x0000050B, "ERROR_DRIVER_PROCESS_TERMINATED", "The process hosting the driver for this device has been terminated."}
	ErrorImplementationLimit                               = &Error{0x0000050C, "ERROR_IMPLEMENTATION_LIMIT", "An operation attempted to exceed an implementation-defined limit."}
	ErrorProcessIsProtected                                = &Error{0x0000050D, "ERROR_PROCESS_IS_PROTECTED", "Either the target process, or the target thread's containing process, is a protected process."}
	ErrorServiceNotifyClientLagging                        = &Error{0x0000050E, "ERROR_SERVICE_NOTIFY_CLIENT_LAGGING", "The service notification client is lagging too far behind the current state of services in the machine."}
	ErrorDiskQuotaExceeded                                 = &Error{0x0000050F, "ERROR_DISK_QUOTA_EXCEEDED", "An operation failed because the storage quota was exceeded."}
	ErrorContentBlocked                                    = &Error{0x00000510, "ERROR_CONTENT_BLOCKED", "An operation failed because the content was blocked."}
	ErrorIncompatibleServicePrivilege                      = &Error{0x00000511, "ERROR_INCOMPATIBLE_SERVICE_PRIVILEGE", "A privilege that the service requires to function properly does not exist in the service account configuration. The Services Microsoft Management Console (MMC) snap-in (Services.msc) and the Local Security Settings MMC snap-in (Secpol.msc) can be used to view the service configuration and the account configuration."}
	ErrorInvalidLabel                                      = &Error{0x00000513, "ERROR_INVALID_LABEL", "Indicates a particular SID cannot be assigned as the label of an object."}
	ErrorNotAllAssigned                                    = &Error{0x00000514, "ERROR_NOT_ALL_ASSIGNED", "Not all privileges or groups referenced are assigned to the caller."}
	ErrorSomeNotMapped                                     = &Error{0x00000515, "ERROR_SOME_NOT_MAPPED", "Some mapping between account names and SIDs was not done."}
	ErrorNoQuotasForAccount                                = &Error{0x00000516, "ERROR_NO_QUOTAS_FOR_ACCOUNT", "No system quota limits are specifically set for this account."}
	ErrorLocalUserSessionKey                               = &Error{0x00000517, "ERROR_LOCAL_USER_SESSION_KEY", "No encryption key is available. A well-known encryption key was returned."}
	ErrorNullLmPassword                                    = &Error{0x00000518, "ERROR_NULL_LM_PASSWORD", "The password is too complex to be converted to a LAN Manager password. The LAN Manager password returned is a null string."}
	ErrorUnknownRevision                                   = &Error{0x00000519, "ERROR_UNKNOWN_REVISION", "The revision level is unknown."}
	ErrorRevisionMismatch                                  = &Error{0x0000051A, "ERROR_REVISION_MISMATCH", "Indicates two revision levels are incompatible."}
	ErrorInvalidOwner                                      = &Error{0x0000051B, "ERROR_INVALID_OWNER", "This SID cannot be assigned as the owner of this object."}
	ErrorInvalidPrimaryGroup                               = &Error{0x0000051C, "ERROR_INVALID_PRIMARY_GROUP", "This SID cannot be assigned as the primary group of an object."}
	ErrorNoImpersonationToken                              = &Error{0x0000051D, "ERROR_NO_IMPERSONATION_TOKEN", "An attempt has been made to operate on an impersonation token by a thread that is not currently impersonating a client."}
	ErrorCantDisableMandatory                              = &Error{0x0000051E, "ERROR_CANT_DISABLE_MANDATORY", "The group cannot be disabled."}
	ErrorNoLogonServers                                    = &Error{0x0000051F, "ERROR_NO_LOGON_SERVERS", "There are currently no logon servers available to service the logon request."}
	ErrorNoSuchLogonSession                                = &Error{0x00000520, "ERROR_NO_SUCH_LOGON_SESSION", "A specified logon session does not exist. It might already have been terminated."}
	ErrorNoSuchPrivilege                                   = &Error{0x00000521, "ERROR_NO_SUCH_PRIVILEGE", "A specified privilege does not exist."}
	ErrorPrivilegeNotHeld                                  = &Error{0x00000522, "ERROR_PRIVILEGE_NOT_HELD", "A required privilege is not held by the client."}
	ErrorInvalidAccountName                                = &Error{0x00000523, "ERROR_INVALID_ACCOUNT_NAME", "The name provided is not a properly formed account name."}
	ErrorUserExists                                        = &Error{0x00000524, "ERROR_USER_EXISTS", "The specified account already exists."}
	ErrorNoSuchUser                                        = &Error{0x00000525, "ERROR_NO_SUCH_USER", "The specified account does not exist."}
	ErrorGroupExists                                       = &Error{0x00000526, "ERROR_GROUP_EXISTS", "The specified group already exists."}
	ErrorNoSuchGroup                                       = &Error{0x00000527, "ERROR_NO_SUCH_GROUP", "The specified group does not exist."}
	ErrorMemberInGroup                                     = &Error{0x00000528, "ERROR_MEMBER_IN_GROUP", "Either the specified user account is already a member of the specified group, or the specified group cannot be deleted because it contains a member."}
	ErrorMemberNotInGroup                                  = &Error{0x00000529, "ERROR_MEMBER_NOT_IN_GROUP", "The specified user account is not a member of the specified group account."}
	ErrorLastAdmin                                         = &Error{0x0000052A, "ERROR_LAST_ADMIN", "The last remaining administration account cannot be disabled or deleted."}
	ErrorWrongPassword                                     = &Error{0x0000052B, "ERROR_WRONG_PASSWORD", "Unable to update the password. The value provided as the current password is incorrect."}
	ErrorIllFormedPassword                                 = &Error{0x0000052C, "ERROR_ILL_FORMED_PASSWORD", "Unable to update the password. The value provided for the new password contains values that are not allowed in passwords."}
	ErrorPasswordRestriction                               = &Error{0x0000052D, "ERROR_PASSWORD_RESTRICTION", "Unable to update the password. The value provided for the new password does not meet the length, complexity, or history requirements of the domain."}
	ErrorLogonFailure                                      = &Error{0x0000052E, "ERROR_LOGON_FAILURE", "Logon failure: Unknown user name or bad password."}
	ErrorAccountRestriction                                = &Error{0x0000052F, "ERROR_ACCOUNT_RESTRICTION", "Logon failure: User account restriction. Possible reasons are blank passwords not allowed, logon hour restrictions, or a policy restriction has been enforced."}
	ErrorInvalidLogonHours                                 = &Error{0x00000530, "ERROR_INVALID_LOGON_HOURS", "Logon failure: Account logon time restriction violation."}
	ErrorInvalidWorkstation                                = &Error{0x00000531, "ERROR_INVALID_WORKSTATION", "Logon failure: User not allowed to log on to this computer."}
	ErrorPasswordExpired                                   = &Error{0x00000532, "ERROR_PASSWORD_EXPIRED", "Logon failure: The specified account password has expired."}
	ErrorAccountDisabled                                   = &Error{0x00000533, "ERROR_ACCOUNT_DISABLED", "Logon failure: Account currently disabled."}
	ErrorNoneMapped                                        = &Error{0x00000534, "ERROR_NONE_MAPPED", "No mapping between account names and SIDs was done."}
	ErrorTooManyLuidsRequested                             = &Error{0x00000535, "ERROR_TOO_MANY_LUIDS_REQUESTED", "Too many local user identifiers (LUIDs) were requested at one time."}
	ErrorLuidsExhausted                                    = &Error{0x00000536, "ERROR_LUIDS_EXHAUSTED", "No more LUIDs are available."}
	ErrorInvalidSubAuthority                               = &Error{0x00000537, "ERROR_INVALID_SUB_AUTHORITY", "The sub-authority part of an SID is invalid for this particular use."}
	ErrorInvalidAcl                                        = &Error{0x00000538, "ERROR_INVALID_ACL", "The ACL structure is invalid."}
	ErrorInvalidSid                                        = &Error{0x00000539, "ERROR_INVALID_SID", "The SID structure is invalid."}
	ErrorInvalidSecurityDescr                              = &Error{0x0000053A, "ERROR_INVALID_SECURITY_DESCR", "The security descriptor structure is invalid."}
	ErrorBadInheritanceAcl                                 = &Error{0x0000053C, "ERROR_BAD_INHERITANCE_ACL", "The inherited ACL or ACE could not be built."}
	ErrorServerDisabled                                    = &Error{0x0000053D, "ERROR_SERVER_DISABLED", "The server is currently disabled."}
	ErrorServerNotDisabled                                 = &Error{0x0000053E, "ERROR_SERVER_NOT_DISABLED", "The server is currently enabled."}
	ErrorInvalidIdAuthority                                = &Error{0x0000053F, "ERROR_INVALID_ID_AUTHORITY", "The value provided was an invalid value for an identifier authority."}
	ErrorAllottedSpaceExceeded                             = &Error{0x00000540, "ERROR_ALLOTTED_SPACE_EXCEEDED", "No more memory is available for security information updates."}
	ErrorInvalidGroupAttributes                            = &Error{0x00000541, "ERROR_INVALID_GROUP_ATTRIBUTES", "The specified attributes are invalid, or incompatible with the attributes for the group as a whole."}
	ErrorBadImpersonationLevel                             = &Error{0x00000542, "ERROR_BAD_IMPERSONATION_LEVEL", "Either a required impersonation level was not provided, or the provided impersonation level is invalid."}
	ErrorCantOpenAnonymous                                 = &Error{0x00000543, "ERROR_CANT_OPEN_ANONYMOUS", "Cannot open an anonymous level security token."}
	ErrorBadValidationClass                                = &Error{0x00000544, "ERROR_BAD_VALIDATION_CLASS", "The validation information class requested was invalid."}
	ErrorBadTokenType                                      = &Error{0x00000545, "ERROR_BAD_TOKEN_TYPE", "The type of the token is inappropriate for its attempted use."}
	ErrorNoSecurityOnObject                                = &Error{0x00000546, "ERROR_NO_SECURITY_ON_OBJECT", "Unable to perform a security operation on an object that has no associated security."}
	ErrorCantAccessDomainInfo                              = &Error{0x00000547, "ERROR_CANT_ACCESS_DOMAIN_INFO", "Configuration information could not be read from the domain controller, either because the machine is unavailable, or access has been denied."}
	ErrorInvalidServerState                                = &Error{0x00000548, "ERROR_INVALID_SERVER_STATE", "The SAM or local security authority (LSA) server was in the wrong state to perform the security operation."}
	ErrorInvalidDomainState                                = &Error{0x00000549, "ERROR_INVALID_DOMAIN_STATE", "The domain was in the wrong state to perform the security operation."}
	ErrorInvalidDomainRole                                 = &Error{0x0000054A, "ERROR_INVALID_DOMAIN_ROLE", "This operation is only allowed for the PDC of the domain."}
	ErrorNoSuchDomain                                      = &Error{0x0000054B, "ERROR_NO_SUCH_DOMAIN", "The specified domain either does not exist or could not be contacted."}
	ErrorDomainExists                                      = &Error{0x0000054C, "ERROR_DOMAIN_EXISTS", "The specified domain already exists."}
	ErrorDomainLimitExceeded                               = &Error{0x0000054D, "ERROR_DOMAIN_LIMIT_EXCEEDED", "An attempt was made to exceed the limit on the number of domains per server."}
	ErrorInternalDbCorruption                              = &Error{0x0000054E, "ERROR_INTERNAL_DB_CORRUPTION", "Unable to complete the requested operation because of either a catastrophic media failure or a data structure corruption on the disk."}
	ErrorInternalError                                     = &Error{0x0000054F, "ERROR_INTERNAL_ERROR", "An internal error occurred."}
	ErrorGenericNotMapped                                  = &Error{0x00000550, "ERROR_GENERIC_NOT_MAPPED", "Generic access types were contained in an access mask that should already be mapped to nongeneric types."}
	ErrorBadDescriptorFormat                               = &Error{0x00000551, "ERROR_BAD_DESCRIPTOR_FORMAT", "A security descriptor is not in the right format (absolute or self-relative)."}
	ErrorNotLogonProcess                                   = &Error{0x00000552, "ERROR_NOT_LOGON_PROCESS", "The requested action is restricted for use by logon processes only. The calling process has not registered as a logon process."}
	ErrorLogonSessionExists                                = &Error{0x00000553, "ERROR_LOGON_SESSION_EXISTS", "Cannot start a new logon session with an ID that is already in use."}
	ErrorNoSuchPackage                                     = &Error{0x00000554, "ERROR_NO_SUCH_PACKAGE", "A specified authentication package is unknown."}
	ErrorBadLogonSessionState                              = &Error{0x00000555, "ERROR_BAD_LOGON_SESSION_STATE", "The logon session is not in a state that is consistent with the requested operation."}
	ErrorLogonSessionCollision                             = &Error{0x00000556, "ERROR_LOGON_SESSION_COLLISION", "The logon session ID is already in use."}
	ErrorInvalidLogonType                                  = &Error{0x00000557, "ERROR_INVALID_LOGON_TYPE", "A logon request contained an invalid logon type value."}
	ErrorCannotImpersonate                                 = &Error{0x00000558, "ERROR_CANNOT_IMPERSONATE", "Unable to impersonate using a named pipe until data has been read from that pipe."}
	ErrorRxactInvalidState                                 = &Error{0x00000559, "ERROR_RXACT_INVALID_STATE", "The transaction state of a registry subtree is incompatible with the requested operation."}
	ErrorRxactCommitFailure                                = &Error{0x0000055A, "ERROR_RXACT_COMMIT_FAILURE", "An internal security database corruption has been encountered."}
	ErrorSpecialAccount                                    = &Error{0x0000055B, "ERROR_SPECIAL_ACCOUNT", "Cannot perform this operation on built-in accounts."}
	ErrorSpecialGroup                                      = &Error{0x0000055C, "ERROR_SPECIAL_GROUP", "Cannot perform this operation on this built-in special group."}
	ErrorSpecialUser                                       = &Error{0x0000055D, "ERROR_SPECIAL_USER", "Cannot perform this operation on this built-in special user."}
	ErrorMembersPrimaryGroup                               = &Error{0x0000055E, "ERROR_MEMBERS_PRIMARY_GROUP", "The user cannot be removed from a group because the group is currently the user's primary group."}
	ErrorTokenAlreadyInUse                                 = &Error{0x0000055F, "ERROR_TOKEN_ALREADY_IN_USE", "The token is already in use as a primary token."}
	ErrorNoSuchAlias                                       = &Error{0x00000560, "ERROR_NO_SUCH_ALIAS", "The specified local group does not exist."}
	ErrorMemberNotInAlias                                  = &Error{0x00000561, "ERROR_MEMBER_NOT_IN_ALIAS", "The specified account name is not a member of the group."}
	ErrorMemberInAlias                                     = &Error{0x00000562, "ERROR_MEMBER_IN_ALIAS", "The specified account name is already a member of the group."}
	ErrorAliasExists                                       = &Error{0x00000563, "ERROR_ALIAS_EXISTS", "The specified local group already exists."}
	ErrorLogonNotGranted                                   = &Error{0x00000564, "ERROR_LOGON_NOT_GRANTED", "Logon failure: The user has not been granted the requested logon type at this computer."}
	ErrorTooManySecrets                                    = &Error{0x00000565, "ERROR_TOO_MANY_SECRETS", "The maximum number of secrets that can be stored in a single system has been exceeded."}
	ErrorSecretTooLong                                     = &Error{0x00000566, "ERROR_SECRET_TOO_LONG", "The length of a secret exceeds the maximum length allowed."}
	ErrorInternalDbError                                   = &Error{0x00000567, "ERROR_INTERNAL_DB_ERROR", "The local security authority database contains an internal inconsistency."}
	ErrorTooManyContextIds                                 = &Error{0x00000568, "ERROR_TOO_MANY_CONTEXT_IDS", "During a logon attempt, the user's security context accumulated too many SIDs."}
	ErrorLogonTypeNotGranted                               = &Error{0x00000569, "ERROR_LOGON_TYPE_NOT_GRANTED", "Logon failure: The user has not been granted the requested logon type at this computer."}
	ErrorNtCrossEncryptionRequired                         = &Error{0x0000056A, "ERROR_NT_CROSS_ENCRYPTION_REQUIRED", "A cross-encrypted password is necessary to change a user password."}
	ErrorNoSuchMember                                      = &Error{0x0000056B, "ERROR_NO_SUCH_MEMBER", "A member could not be added to or removed from the local group because the member does not exist."}
	ErrorInvalidMember                                     = &Error{0x0000056C, "ERROR_INVALID_MEMBER", "A new member could not be added to a local group because the member has the wrong account type."}
	ErrorTooManySids                                       = &Error{0x0000056D, "ERROR_TOO_MANY_SIDS", "Too many SIDs have been specified."}
	ErrorLmCrossEncryptionRequired                         = &Error{0x0000056E, "ERROR_LM_CROSS_ENCRYPTION_REQUIRED", "A cross-encrypted password is necessary to change this user password."}
	ErrorNoInheritance                                     = &Error{0x0000056F, "ERROR_NO_INHERITANCE", "Indicates an ACL contains no inheritable components."}
	ErrorFileCorrupt                                       = &Error{0x00000570, "ERROR_FILE_CORRUPT", "The file or directory is corrupted and unreadable."}
	ErrorDiskCorrupt                                       = &Error{0x00000571, "ERROR_DISK_CORRUPT", "The disk structure is corrupted and unreadable."}
	ErrorNoUserSessionKey                                  = &Error{0x00000572, "ERROR_NO_USER_SESSION_KEY", "There is no user session key for the specified logon session."}
	ErrorLicenseQuotaExceeded                              = &Error{0x00000573, "ERROR_LICENSE_QUOTA_EXCEEDED", "The service being accessed is licensed for a particular number of connections. No more connections can be made to the service at this time because the service has accepted the maximum number of connections."}
	ErrorWrongTargetName                                   = &Error{0x00000574, "ERROR_WRONG_TARGET_NAME", "Logon failure: The target account name is incorrect."}
	ErrorMutualAuthFailed                                  = &Error{0x00000575, "ERROR_MUTUAL_AUTH_FAILED", "Mutual authentication failed. The server's password is out of date at the domain controller."}
	ErrorTimeSkew                                          = &Error{0x00000576, "ERROR_TIME_SKEW", "There is a time and/or date difference between the client and server."}
	ErrorCurrentDomainNotAllowed                           = &Error{0x00000577, "ERROR_CURRENT_DOMAIN_NOT_ALLOWED", "This operation cannot be performed on the current domain."}
	ErrorInvalidWindowHandle                               = &Error{0x00000578, "ERROR_INVALID_WINDOW_HANDLE", "Invalid window handle."}
	ErrorInvalidMenuHandle                                 = &Error{0x00000579, "ERROR_INVALID_MENU_HANDLE", "Invalid menu handle."}
	ErrorInvalidCursorHandle                               = &Error{0x0000057A, "ERROR_INVALID_CURSOR_HANDLE", "Invalid cursor handle."}
	ErrorInvalidAccelHandle                                = &Error{0x0000057B, "ERROR_INVALID_ACCEL_HANDLE", "Invalid accelerator table handle."}
	ErrorInvalidHookHandle                                 = &Error{0x0000057C, "ERROR_INVALID_HOOK_HANDLE", "Invalid hook handle."}
	ErrorInvalidDwpHandle                                  = &Error{0x0000057D, "ERROR_INVALID_DWP_HANDLE", "Invalid handle to a multiple-window position structure."}
	ErrorTlwWithWschild                                    = &Error{0x0000057E, "ERROR_TLW_WITH_WSCHILD", "Cannot create a top-level child window."}
	ErrorCannotFindWndClass                                = &Error{0x0000057F, "ERROR_CANNOT_FIND_WND_CLASS", "Cannot find window class."}
	ErrorWindowOfOtherThread                               = &Error{0x00000580, "ERROR_WINDOW_OF_OTHER_THREAD", "Invalid window; it belongs to other thread."}
	ErrorHotkeyAlreadyRegistered                           = &Error{0x00000581, "ERROR_HOTKEY_ALREADY_REGISTERED", "Hot key is already registered."}
	ErrorClassAlreadyExists                                = &Error{0x00000582, "ERROR_CLASS_ALREADY_EXISTS", "Class already exists."}
	ErrorClassDoesNotExist                                 = &Error{0x00000583, "ERROR_CLASS_DOES_NOT_EXIST", "Class does not exist."}
	ErrorClassHasWindows                                   = &Error{0x00000584, "ERROR_CLASS_HAS_WINDOWS", "Class still has open windows."}
	ErrorInvalidIndex                                      = &Error{0x00000585, "ERROR_INVALID_INDEX", "Invalid index."}
	ErrorInvalidIconHandle                                 = &Error{0x00000586, "ERROR_INVALID_ICON_HANDLE", "Invalid icon handle."}
	ErrorPrivateDialogIndex                                = &Error{0x00000587, "ERROR_PRIVATE_DIALOG_INDEX", "Using private DIALOG window words."}
	ErrorListboxIdNotFound                                 = &Error{0x00000588, "ERROR_LISTBOX_ID_NOT_FOUND", "The list box identifier was not found."}
	ErrorNoWildcardCharacters                              = &Error{0x00000589, "ERROR_NO_WILDCARD_CHARACTERS", "No wildcards were found."}
	ErrorClipboardNotOpen                                  = &Error{0x0000058A, "ERROR_CLIPBOARD_NOT_OPEN", "Thread does not have a clipboard open."}
	ErrorHotkeyNotRegistered                               = &Error{0x0000058B, "ERROR_HOTKEY_NOT_REGISTERED", "Hot key is not registered."}
	ErrorWindowNotDialog                                   = &Error{0x0000058C, "ERROR_WINDOW_NOT_DIALOG", "The window is not a valid dialog window."}
	ErrorControlIdNotFound                                 = &Error{0x0000058D, "ERROR_CONTROL_ID_NOT_FOUND", "Control ID not found."}
	ErrorInvalidComboboxMessage                            = &Error{0x0000058E, "ERROR_INVALID_COMBOBOX_MESSAGE", "Invalid message for a combo box because it does not have an edit control."}
	ErrorWindowNotCombobox                                 = &Error{0x0000058F, "ERROR_WINDOW_NOT_COMBOBOX", "The window is not a combo box."}
	ErrorInvalidEditHeight                                 = &Error{0x00000590, "ERROR_INVALID_EDIT_HEIGHT", "Height must be less than 256."}
	ErrorDcNotFound                                        = &Error{0x00000591, "ERROR_DC_NOT_FOUND", "Invalid device context (DC) handle."}
	ErrorInvalidHookFilter                                 = &Error{0x00000592, "ERROR_INVALID_HOOK_FILTER", "Invalid hook procedure type."}
	ErrorInvalidFilterProc                                 = &Error{0x00000593, "ERROR_INVALID_FILTER_PROC", "Invalid hook procedure."}
	ErrorHookNeedsHmod                                     = &Error{0x00000594, "ERROR_HOOK_NEEDS_HMOD", "Cannot set nonlocal hook without a module handle."}
	ErrorGlobalOnlyHook                                    = &Error{0x00000595, "ERROR_GLOBAL_ONLY_HOOK", "This hook procedure can only be set globally."}
	ErrorJournalHookSet                                    = &Error{0x00000596, "ERROR_JOURNAL_HOOK_SET", "The journal hook procedure is already installed."}
	ErrorHookNotInstalled                                  = &Error{0x00000597, "ERROR_HOOK_NOT_INSTALLED", "The hook procedure is not installed."}
	ErrorInvalidLbMessage                                  = &Error{0x00000598, "ERROR_INVALID_LB_MESSAGE", "Invalid message for single-selection list box."}
	ErrorSetcountOnBadLb                                   = &Error{0x00000599, "ERROR_SETCOUNT_ON_BAD_LB", "LB_SETCOUNT sent to non-lazy list box."}
	ErrorLbWithoutTabstops                                 = &Error{0x0000059A, "ERROR_LB_WITHOUT_TABSTOPS", "This list box does not support tab stops."}
	ErrorDestroyObjectOfOtherThread                        = &Error{0x0000059B, "ERROR_DESTROY_OBJECT_OF_OTHER_THREAD", "Cannot destroy object created by another thread."}
	ErrorChildWindowMenu                                   = &Error{0x0000059C, "ERROR_CHILD_WINDOW_MENU", "Child windows cannot have menus."}
	ErrorNoSystemMenu                                      = &Error{0x0000059D, "ERROR_NO_SYSTEM_MENU", "The window does not have a system menu."}
	ErrorInvalidMsgboxStyle                                = &Error{0x0000059E, "ERROR_INVALID_MSGBOX_STYLE", "Invalid message box style."}
	ErrorInvalidSpiValue                                   = &Error{0x0000059F, "ERROR_INVALID_SPI_VALUE", "Invalid system-wide (SPI_*) parameter."}
	ErrorScreenAlreadyLocked                               = &Error{0x000005A0, "ERROR_SCREEN_ALREADY_LOCKED", "Screen already locked."}
	ErrorHwndsHaveDiffParent                               = &Error{0x000005A1, "ERROR_HWNDS_HAVE_DIFF_PARENT", "All handles to windows in a multiple-window position structure must have the same parent."}
	ErrorNotChildWindow                                    = &Error{0x000005A2, "ERROR_NOT_CHILD_WINDOW", "The window is not a child window."}
	ErrorInvalidGwCommand                                  = &Error{0x000005A3, "ERROR_INVALID_GW_COMMAND", "Invalid GW_* command."}
	ErrorInvalidThreadId                                   = &Error{0x000005A4, "ERROR_INVALID_THREAD_ID", "Invalid thread identifier."}
	ErrorNonMdichildWindow                                 = &Error{0x000005A5, "ERROR_NON_MDICHILD_WINDOW", "Cannot process a message from a window that is not a multiple document interface (MDI) window."}
	ErrorPopupAlreadyActive                                = &Error{0x000005A6, "ERROR_POPUP_ALREADY_ACTIVE", "Pop-up menu already active."}
	ErrorNoScrollbars                                      = &Error{0x000005A7, "ERROR_NO_SCROLLBARS", "The window does not have scroll bars."}
	ErrorInvalidScrollbarRange                             = &Error{0x000005A8, "ERROR_INVALID_SCROLLBAR_RANGE", "Scroll bar range cannot be greater than MAXLONG."}
	ErrorInvalidShowwinCommand                             = &Error{0x000005A9, "ERROR_INVALID_SHOWWIN_COMMAND", "Cannot show or remove the window in the way specified."}
	ErrorNoSystemResources                                 = &Error{0x000005AA, "ERROR_NO_SYSTEM_RESOURCES", "Insufficient system resources exist to complete the requested service."}
	ErrorNonpagedSystemResources                           = &Error{0x000005AB, "ERROR_NONPAGED_SYSTEM_RESOURCES", "Insufficient system resources exist to complete the requested service."}
	ErrorPagedSystemResources                              = &Error{0x000005AC, "ERROR_PAGED_SYSTEM_RESOURCES", "Insufficient system resources exist to complete the requested service."}
	ErrorWorkingSetQuota                                   = &Error{0x000005AD, "ERROR_WORKING_SET_QUOTA", "Insufficient quota to complete the requested service."}
	ErrorPagefileQuota                                     = &Error{0x000005AE, "ERROR_PAGEFILE_QUOTA", "Insufficient quota to complete the requested service."}
	ErrorCommitmentLimit                                   = &Error{0x000005AF, "ERROR_COMMITMENT_LIMIT", "The paging file is too small for this operation to complete."}
	ErrorMenuItemNotFound                                  = &Error{0x000005B0, "ERROR_MENU_ITEM_NOT_FOUND", "A menu item was not found."}
	ErrorInvalidKeyboardHandle                             = &Error{0x000005B1, "ERROR_INVALID_KEYBOARD_HANDLE", "Invalid keyboard layout handle."}
	ErrorHookTypeNotAllowed                                = &Error{0x000005B2, "ERROR_HOOK_TYPE_NOT_ALLOWED", "Hook type not allowed."}
	ErrorRequiresInteractiveWindowstation                  = &Error{0x000005B3, "ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION", "This operation requires an interactive window station."}
	ErrorTimeout                                           = &Error{0x000005B4, "ERROR_TIMEOUT", "This operation returned because the time-out period expired."}
	ErrorInvalidMonitorHandle                              = &Error{0x000005B5, "ERROR_INVALID_MONITOR_HANDLE", "Invalid monitor handle."}
	ErrorIncorrectSize                                     = &Error{0x000005B6, "ERROR_INCORRECT_SIZE", "Incorrect size argument."}
	ErrorSymlinkClassDisabled                              = &Error{0x000005B7, "ERROR_SYMLINK_CLASS_DISABLED", "The symbolic link cannot be followed because its type is disabled."}
	ErrorSymlinkNotSupported                               = &Error{0x000005B8, "ERROR_SYMLINK_NOT_SUPPORTED", "This application does not support the current operation on symbolic links."}
	ErrorEventlogFileCorrupt                               = &Error{0x000005DC, "ERROR_EVENTLOG_FILE_CORRUPT", "The event log file is corrupted."}
	ErrorEventlogCantStart                                 = &Error{0x000005DD, "ERROR_EVENTLOG_CANT_START", "No event log file could be opened, so the event logging service did not start."}
	ErrorLogFileFull                                       = &Error{0x000005DE, "ERROR_LOG_FILE_FULL", "The event log file is full."}
	ErrorEventlogFileChanged                               = &Error{0x000005DF, "ERROR_EVENTLOG_FILE_CHANGED", "The event log file has changed between read operations."}
	ErrorInvalidTaskName                                   = &Error{0x0000060E, "ERROR_INVALID_TASK_NAME", "The specified task name is invalid."}
	ErrorInvalidTaskIndex                                  = &Error{0x0000060F, "ERROR_INVALID_TASK_INDEX", "The specified task index is invalid."}
	ErrorThreadAlreadyInTask                               = &Error{0x00000610, "ERROR_THREAD_ALREADY_IN_TASK", "The specified thread is already joining a task."}
	ErrorInstallServiceFailure                             = &Error{0x00000641, "ERROR_INSTALL_SERVICE_FAILURE", "The Windows Installer service could not be accessed. This can occur if the Windows Installer is not correctly installed. Contact your support personnel for assistance."}
	ErrorInstallUserexit                                   = &Error{0x00000642, "ERROR_INSTALL_USEREXIT", "User canceled installation."}
	ErrorInstallFailure                                    = &Error{0x00000643, "ERROR_INSTALL_FAILURE", "Fatal error during installation."}
	ErrorInstallSuspend                                    = &Error{0x00000644, "ERROR_INSTALL_SUSPEND", "Installation suspended, incomplete."}
	ErrorUnknownProduct                                    = &Error{0x00000645, "ERROR_UNKNOWN_PRODUCT", "This action is valid only for products that are currently installed."}
	ErrorUnknownFeature                                    = &Error{0x00000646, "ERROR_UNKNOWN_FEATURE", "Feature ID not registered."}
	ErrorUnknownComponent                                  = &Error{0x00000647, "ERROR_UNKNOWN_COMPONENT", "Component ID not registered."}
	ErrorUnknownProperty                                   = &Error{0x00000648, "ERROR_UNKNOWN_PROPERTY", "Unknown property."}
	ErrorInvalidHandleState                                = &Error{0x00000649, "ERROR_INVALID_HANDLE_STATE", "Handle is in an invalid state."}
	ErrorBadConfiguration                                  = &Error{0x0000064A, "ERROR_BAD_CONFIGURATION", "The configuration data for this product is corrupt. Contact your support personnel."}
	ErrorIndexAbsent                                       = &Error{0x0000064B, "ERROR_INDEX_ABSENT", "Component qualifier not present."}
	ErrorInstallSourceAbsent                               = &Error{0x0000064C, "ERROR_INSTALL_SOURCE_ABSENT", "The installation source for this product is not available. Verify that the source exists and that you can access it."}
	ErrorInstallPackageVersion                             = &Error{0x0000064D, "ERROR_INSTALL_PACKAGE_VERSION", "This installation package cannot be installed by the Windows Installer service. You must install a Windows service pack that contains a newer version of the Windows Installer service."}
	ErrorProductUninstalled                                = &Error{0x0000064E, "ERROR_PRODUCT_UNINSTALLED", "Product is uninstalled."}
	ErrorBadQuerySyntax                                    = &Error{0x0000064F, "ERROR_BAD_QUERY_SYNTAX", "SQL query syntax invalid or unsupported."}
	ErrorInvalidField                                      = &Error{0x00000650, "ERROR_INVALID_FIELD", "Record field does not exist."}
	ErrorDeviceRemoved                                     = &Error{0x00000651, "ERROR_DEVICE_REMOVED", "The device has been removed."}
	ErrorInstallAlreadyRunning                             = &Error{0x00000652, "ERROR_INSTALL_ALREADY_RUNNING", "Another installation is already in progress. Complete that installation before proceeding with this install."}
	ErrorInstallPackageOpenFailed                          = &Error{0x00000653, "ERROR_INSTALL_PACKAGE_OPEN_FAILED", "This installation package could not be opened. Verify that the package exists and that you can access it, or contact the application vendor to verify that this is a valid Windows Installer package."}
	ErrorInstallPackageInvalid                             = &Error{0x00000654, "ERROR_INSTALL_PACKAGE_INVALID", "This installation package could not be opened. Contact the application vendor to verify that this is a valid Windows Installer package."}
	ErrorInstallUiFailure                                  = &Error{0x00000655, "ERROR_INSTALL_UI_FAILURE", "There was an error starting the Windows Installer service user interface. Contact your support personnel."}
	ErrorInstallLogFailure                                 = &Error{0x00000656, "ERROR_INSTALL_LOG_FAILURE", "Error opening installation log file. Verify that the specified log file location exists and that you can write to it."}
	ErrorInstallLanguageUnsupported                        = &Error{0x00000657, "ERROR_INSTALL_LANGUAGE_UNSUPPORTED", "The language of this installation package is not supported by your system."}
	ErrorInstallTransformFailure                           = &Error{0x00000658, "ERROR_INSTALL_TRANSFORM_FAILURE", "Error applying transforms. Verify that the specified transform paths are valid."}
	ErrorInstallPackageRejected                            = &Error{0x00000659, "ERROR_INSTALL_PACKAGE_REJECTED", "This installation is forbidden by system policy. Contact your system administrator."}
	ErrorFunctionNotCalled                                 = &Error{0x0000065A, "ERROR_FUNCTION_NOT_CALLED", "Function could not be executed."}
	ErrorFunctionFailed                                    = &Error{0x0000065B, "ERROR_FUNCTION_FAILED", "Function failed during execution."}
	ErrorInvalidTable                                      = &Error{0x0000065C, "ERROR_INVALID_TABLE", "Invalid or unknown table specified."}
	ErrorDatatypeMismatch                                  = &Error{0x0000065D, "ERROR_DATATYPE_MISMATCH", "Data supplied is of wrong type."}
	ErrorUnsupportedType                                   = &Error{0x0000065E, "ERROR_UNSUPPORTED_TYPE", "Data of this type is not supported."}
	ErrorCreateFailed                                      = &Error{0x0000065F, "ERROR_CREATE_FAILED", "The Windows Installer service failed to start. Contact your support personnel."}
	ErrorInstallTempUnwritable                             = &Error{0x00000660, "ERROR_INSTALL_TEMP_UNWRITABLE", "The Temp folder is on a drive that is full or is inaccessible. Free up space on the drive or verify that you have write permission on the Temp folder."}
	ErrorInstallPlatformUnsupported                        = &Error{0x00000661, "ERROR_INSTALL_PLATFORM_UNSUPPORTED", "This installation package is not supported by this processor type. Contact your product vendor."}
	ErrorInstallNotused                                    = &Error{0x00000662, "ERROR_INSTALL_NOTUSED", "Component not used on this computer."}
	ErrorPatchPackageOpenFailed                            = &Error{0x00000663, "ERROR_PATCH_PACKAGE_OPEN_FAILED", "This update package could not be opened. Verify that the update package exists and that you can access it, or contact the application vendor to verify that this is a valid Windows Installer update package."}
	ErrorPatchPackageInvalid                               = &Error{0x00000664, "ERROR_PATCH_PACKAGE_INVALID", "This update package could not be opened. Contact the application vendor to verify that this is a valid Windows Installer update package."}
	ErrorPatchPackageUnsupported                           = &Error{0x00000665, "ERROR_PATCH_PACKAGE_UNSUPPORTED", "This update package cannot be processed by the Windows Installer service. You must install a Windows service pack that contains a newer version of the Windows Installer service."}
	ErrorProductVersion                                    = &Error{0x00000666, "ERROR_PRODUCT_VERSION", "Another version of this product is already installed. Installation of this version cannot continue. To configure or remove the existing version of this product, use Add/Remove Programs in Control Panel."}
	ErrorInvalidCommandLine                                = &Error{0x00000667, "ERROR_INVALID_COMMAND_LINE", "Invalid command-line argument. Consult the Windows Installer SDK for detailed command line help."}
	ErrorInstallRemoteDisallowed                           = &Error{0x00000668, "ERROR_INSTALL_REMOTE_DISALLOWED", "Only administrators have permission to add, remove, or configure server software during a Terminal Services remote session. If you want to install or configure software on the server, contact your network administrator."}
	ErrorSuccessRebootInitiated                            = &Error{0x00000669, "ERROR_SUCCESS_REBOOT_INITIATED", "The requested operation completed successfully. The system will be restarted so the changes can take effect."}
	ErrorPatchTargetNotFound                               = &Error{0x0000066A, "ERROR_PATCH_TARGET_NOT_FOUND", "The upgrade cannot be installed by the Windows Installer service because the program to be upgraded might be missing, or the upgrade might update a different version of the program. Verify that the program to be upgraded exists on your computer and that you have the correct upgrade."}
	ErrorPatchPackageRejected                              = &Error{0x0000066B, "ERROR_PATCH_PACKAGE_REJECTED", "The update package is not permitted by a software restriction policy."}
	ErrorInstallTransformRejected                          = &Error{0x0000066C, "ERROR_INSTALL_TRANSFORM_REJECTED", "One or more customizations are not permitted by a software restriction policy."}
	ErrorInstallRemoteProhibited                           = &Error{0x0000066D, "ERROR_INSTALL_REMOTE_PROHIBITED", "The Windows Installer does not permit installation from a Remote Desktop Connection."}
	ErrorPatchRemovalUnsupported                           = &Error{0x0000066E, "ERROR_PATCH_REMOVAL_UNSUPPORTED", "Uninstallation of the update package is not supported."}
	ErrorUnknownPatch                                      = &Error{0x0000066F, "ERROR_UNKNOWN_PATCH", "The update is not applied to this product."}
	ErrorPatchNoSequence                                   = &Error{0x00000670, "ERROR_PATCH_NO_SEQUENCE", "No valid sequence could be found for the set of updates."}
	ErrorPatchRemovalDisallowed                            = &Error{0x00000671, "ERROR_PATCH_REMOVAL_DISALLOWED", "Update removal was disallowed by policy."}
	ErrorInvalidPatchXml                                   = &Error{0x00000672, "ERROR_INVALID_PATCH_XML", "The XML update data is invalid."}
	ErrorPatchManagedAdvertisedProduct                     = &Error{0x00000673, "ERROR_PATCH_MANAGED_ADVERTISED_PRODUCT", "Windows Installer does not permit updating of managed advertised products. At least one feature of the product must be installed before applying the update."}
	ErrorInstallServiceSafeboot                            = &Error{0x00000674, "ERROR_INSTALL_SERVICE_SAFEBOOT", "The Windows Installer service is not accessible in Safe Mode. Try again when your computer is not in Safe Mode or you can use System Restore to return your machine to a previous good state."}
	RpcSInvalidStringBinding                               = &Error{0x000006A4, "RPC_S_INVALID_STRING_BINDING", "The string binding is invalid."}
	RpcSWrongKindOfBinding                                 = &Error{0x000006A5, "RPC_S_WRONG_KIND_OF_BINDING", "The binding handle is not the correct type."}
	RpcSInvalidBinding                                     = &Error{0x000006A6, "RPC_S_INVALID_BINDING", "The binding handle is invalid."}
	RpcSProtseqNotSupported                                = &Error{0x000006A7, "RPC_S_PROTSEQ_NOT_SUPPORTED", "The RPC protocol sequence is not supported."}
	RpcSInvalidRpcProtseq                                  = &Error{0x000006A8, "RPC_S_INVALID_RPC_PROTSEQ", "The RPC protocol sequence is invalid."}
	RpcSInvalidStringUuid                                  = &Error{0x000006A9, "RPC_S_INVALID_STRING_UUID", "The string UUID is invalid."}
	RpcSInvalidEndpointFormat                              = &Error{0x000006AA, "RPC_S_INVALID_ENDPOINT_FORMAT", "The endpoint format is invalid."}
	RpcSInvalidNetAddr                                     = &Error{0x000006AB, "RPC_S_INVALID_NET_ADDR", "The network address is invalid."}
	RpcSNoEndpointFound                                    = &Error{0x000006AC, "RPC_S_NO_ENDPOINT_FOUND", "No endpoint was found."}
	RpcSInvalidTimeout                                     = &Error{0x000006AD, "RPC_S_INVALID_TIMEOUT", "The time-out value is invalid."}
	RpcSObjectNotFound                                     = &Error{0x000006AE, "RPC_S_OBJECT_NOT_FOUND", "The object UUID) was not found."}
	RpcSAlreadyRegistered                                  = &Error{0x000006AF, "RPC_S_ALREADY_REGISTERED", "The object UUID) has already been registered."}
	RpcSTypeAlreadyRegistered                              = &Error{0x000006B0, "RPC_S_TYPE_ALREADY_REGISTERED", "The type UUID has already been registered."}
	RpcSAlreadyListening                                   = &Error{0x000006B1, "RPC_S_ALREADY_LISTENING", "The RPC server is already listening."}
	RpcSNoProtseqsRegistered                               = &Error{0x000006B2, "RPC_S_NO_PROTSEQS_REGISTERED", "No protocol sequences have been registered."}
	RpcSNotListening                                       = &Error{0x000006B3, "RPC_S_NOT_LISTENING", "The RPC server is not listening."}
	RpcSUnknownMgrType                                     = &Error{0x000006B4, "RPC_S_UNKNOWN_MGR_TYPE", "The manager type is unknown."}
	RpcSUnknownIf                                          = &Error{0x000006B5, "RPC_S_UNKNOWN_IF", "The interface is unknown."}
	RpcSNoBindings                                         = &Error{0x000006B6, "RPC_S_NO_BINDINGS", "There are no bindings."}
	RpcSNoProtseqs                                         = &Error{0x000006B7, "RPC_S_NO_PROTSEQS", "There are no protocol sequences."}
	RpcSCantCreateEndpoint                                 = &Error{0x000006B8, "RPC_S_CANT_CREATE_ENDPOINT", "The endpoint cannot be created."}
	RpcSOutOfResources                                     = &Error{0x000006B9, "RPC_S_OUT_OF_RESOURCES", "Not enough resources are available to complete this operation."}
	RpcSServerUnavailable                                  = &Error{0x000006BA, "RPC_S_SERVER_UNAVAILABLE", "The RPC server is unavailable."}
	RpcSServerTooBusy                                      = &Error{0x000006BB, "RPC_S_SERVER_TOO_BUSY", "The RPC server is too busy to complete this operation."}
	RpcSInvalidNetworkOptions                              = &Error{0x000006BC, "RPC_S_INVALID_NETWORK_OPTIONS", "The network options are invalid."}
	RpcSNoCallActive                                       = &Error{0x000006BD, "RPC_S_NO_CALL_ACTIVE", "There are no RPCs active on this thread."}
	RpcSCallFailed                                         = &Error{0x000006BE, "RPC_S_CALL_FAILED", "The RPC failed."}
	RpcSCallFailedDne                                      = &Error{0x000006BF, "RPC_S_CALL_FAILED_DNE", "The RPC failed and did not execute."}
	RpcSProtocolError                                      = &Error{0x000006C0, "RPC_S_PROTOCOL_ERROR", "An RPC protocol error occurred."}
	RpcSProxyAccessDenied                                  = &Error{0x000006C1, "RPC_S_PROXY_ACCESS_DENIED", "Access to the HTTP proxy is denied."}
	RpcSUnsupportedTransSyn                                = &Error{0x000006C2, "RPC_S_UNSUPPORTED_TRANS_SYN", "The transfer syntax is not supported by the RPC server."}
	RpcSUnsupportedType                                    = &Error{0x000006C4, "RPC_S_UNSUPPORTED_TYPE", "The UUID type is not supported."}
	RpcSInvalidTag                                         = &Error{0x000006C5, "RPC_S_INVALID_TAG", "The tag is invalid."}
	RpcSInvalidBound                                       = &Error{0x000006C6, "RPC_S_INVALID_BOUND", "The array bounds are invalid."}
	RpcSNoEntryName                                        = &Error{0x000006C7, "RPC_S_NO_ENTRY_NAME", "The binding does not contain an entry name."}
	RpcSInvalidNameSyntax                                  = &Error{0x000006C8, "RPC_S_INVALID_NAME_SYNTAX", "The name syntax is invalid."}
	RpcSUnsupportedNameSyntax                              = &Error{0x000006C9, "RPC_S_UNSUPPORTED_NAME_SYNTAX", "The name syntax is not supported."}
	RpcSUuidNoAddress                                      = &Error{0x000006CB, "RPC_S_UUID_NO_ADDRESS", "No network address is available to use to construct a UUID."}
	RpcSDuplicateEndpoint                                  = &Error{0x000006CC, "RPC_S_DUPLICATE_ENDPOINT", "The endpoint is a duplicate."}
	RpcSUnknownAuthnType                                   = &Error{0x000006CD, "RPC_S_UNKNOWN_AUTHN_TYPE", "The authentication type is unknown."}
	RpcSMaxCallsTooSmall                                   = &Error{0x000006CE, "RPC_S_MAX_CALLS_TOO_SMALL", "The maximum number of calls is too small."}
	RpcSStringTooLong                                      = &Error{0x000006CF, "RPC_S_STRING_TOO_LONG", "The string is too long."}
	RpcSProtseqNotFound                                    = &Error{0x000006D0, "RPC_S_PROTSEQ_NOT_FOUND", "The RPC protocol sequence was not found."}
	RpcSProcnumOutOfRange                                  = &Error{0x000006D1, "RPC_S_PROCNUM_OUT_OF_RANGE", "The procedure number is out of range."}
	RpcSBindingHasNoAuth                                   = &Error{0x000006D2, "RPC_S_BINDING_HAS_NO_AUTH", "The binding does not contain any authentication information."}
	RpcSUnknownAuthnService                                = &Error{0x000006D3, "RPC_S_UNKNOWN_AUTHN_SERVICE", "The authentication service is unknown."}
	RpcSUnknownAuthnLevel                                  = &Error{0x000006D4, "RPC_S_UNKNOWN_AUTHN_LEVEL", "The authentication level is unknown."}
	RpcSInvalidAuthIdentity                                = &Error{0x000006D5, "RPC_S_INVALID_AUTH_IDENTITY", "The security context is invalid."}
	RpcSUnknownAuthzService                                = &Error{0x000006D6, "RPC_S_UNKNOWN_AUTHZ_SERVICE", "The authorization service is unknown."}
	EptSInvalidEntry                                       = &Error{0x000006D7, "EPT_S_INVALID_ENTRY", "The entry is invalid."}
	EptSCantPerformOp                                      = &Error{0x000006D8, "EPT_S_CANT_PERFORM_OP", "The server endpoint cannot perform the operation."}
	EptSNotRegistered                                      = &Error{0x000006D9, "EPT_S_NOT_REGISTERED", "There are no more endpoints available from the endpoint mapper."}
	RpcSNothingToExport                                    = &Error{0x000006DA, "RPC_S_NOTHING_TO_EXPORT", "No interfaces have been exported."}
	RpcSIncompleteName                                     = &Error{0x000006DB, "RPC_S_INCOMPLETE_NAME", "The entry name is incomplete."}
	RpcSInvalidVersOption                                  = &Error{0x000006DC, "RPC_S_INVALID_VERS_OPTION", "The version option is invalid."}
	RpcSNoMoreMembers                                      = &Error{0x000006DD, "RPC_S_NO_MORE_MEMBERS", "There are no more members."}
	RpcSNotAllObjsUnexported                               = &Error{0x000006DE, "RPC_S_NOT_ALL_OBJS_UNEXPORTED", "There is nothing to unexport."}
	RpcSInterfaceNotFound                                  = &Error{0x000006DF, "RPC_S_INTERFACE_NOT_FOUND", "The interface was not found."}
	RpcSEntryAlreadyExists                                 = &Error{0x000006E0, "RPC_S_ENTRY_ALREADY_EXISTS", "The entry already exists."}
	RpcSEntryNotFound                                      = &Error{0x000006E1, "RPC_S_ENTRY_NOT_FOUND", "The entry is not found."}
	RpcSNameServiceUnavailable                             = &Error{0x000006E2, "RPC_S_NAME_SERVICE_UNAVAILABLE", "The name service is unavailable."}
	RpcSInvalidNafId                                       = &Error{0x000006E3, "RPC_S_INVALID_NAF_ID", "The network address family is invalid."}
	RpcSCannotSupport                                      = &Error{0x000006E4, "RPC_S_CANNOT_SUPPORT", "The requested operation is not supported."}
	RpcSNoContextAvailable                                 = &Error{0x000006E5, "RPC_S_NO_CONTEXT_AVAILABLE", "No security context is available to allow impersonation."}
	RpcSInternalError                                      = &Error{0x000006E6, "RPC_S_INTERNAL_ERROR", "An internal error occurred in an RPC."}
	RpcSZeroDivide                                         = &Error{0x000006E7, "RPC_S_ZERO_DIVIDE", "The RPC server attempted an integer division by zero."}
	RpcSAddressError                                       = &Error{0x000006E8, "RPC_S_ADDRESS_ERROR", "An addressing error occurred in the RPC server."}
	RpcSFpDivZero                                          = &Error{0x000006E9, "RPC_S_FP_DIV_ZERO", "A floating-point operation at the RPC server caused a division by zero."}
	RpcSFpUnderflow                                        = &Error{0x000006EA, "RPC_S_FP_UNDERFLOW", "A floating-point underflow occurred at the RPC server."}
	RpcSFpOverflow                                         = &Error{0x000006EB, "RPC_S_FP_OVERFLOW", "A floating-point overflow occurred at the RPC server."}
	RpcXNoMoreEntries                                      = &Error{0x000006EC, "RPC_X_NO_MORE_ENTRIES", "The list of RPC servers available for the binding of auto handles has been exhausted."}
	RpcXSsCharTransOpenFail                                = &Error{0x000006ED, "RPC_X_SS_CHAR_TRANS_OPEN_FAIL", "Unable to open the character translation table file."}
	RpcXSsCharTransShortFile                               = &Error{0x000006EE, "RPC_X_SS_CHAR_TRANS_SHORT_FILE", "The file containing the character translation table has fewer than 512 bytes."}
	RpcXSsInNullContext                                    = &Error{0x000006EF, "RPC_X_SS_IN_NULL_CONTEXT", "A null context handle was passed from the client to the host during an RPC."}
	RpcXSsContextDamaged                                   = &Error{0x000006F1, "RPC_X_SS_CONTEXT_DAMAGED", "The context handle changed during an RPC."}
	RpcXSsHandlesMismatch                                  = &Error{0x000006F2, "RPC_X_SS_HANDLES_MISMATCH", "The binding handles passed to an RPC do not match."}
	RpcXSsCannotGetCallHandle                              = &Error{0x000006F3, "RPC_X_SS_CANNOT_GET_CALL_HANDLE", "The stub is unable to get the RPC handle."}
	RpcXNullRefPointer                                     = &Error{0x000006F4, "RPC_X_NULL_REF_POINTER", "A null reference pointer was passed to the stub."}
	RpcXEnumValueOutOfRange                                = &Error{0x000006F5, "RPC_X_ENUM_VALUE_OUT_OF_RANGE", "The enumeration value is out of range."}
	RpcXByteCountTooSmall                                  = &Error{0x000006F6, "RPC_X_BYTE_COUNT_TOO_SMALL", "The byte count is too small."}
	RpcXBadStubData                                        = &Error{0x000006F7, "RPC_X_BAD_STUB_DATA", "The stub received bad data."}
	ErrorInvalidUserBuffer                                 = &Error{0x000006F8, "ERROR_INVALID_USER_BUFFER", "The supplied user buffer is not valid for the requested operation."}
	ErrorUnrecognizedMedia                                 = &Error{0x000006F9, "ERROR_UNRECOGNIZED_MEDIA", "The disk media is not recognized. It might not be formatted."}
	ErrorNoTrustLsaSecret                                  = &Error{0x000006FA, "ERROR_NO_TRUST_LSA_SECRET", "The workstation does not have a trust secret."}
	ErrorNoTrustSamAccount                                 = &Error{0x000006FB, "ERROR_NO_TRUST_SAM_ACCOUNT", "The security database on the server does not have a computer account for this workstation trust relationship."}
	ErrorTrustedDomainFailure                              = &Error{0x000006FC, "ERROR_TRUSTED_DOMAIN_FAILURE", "The trust relationship between the primary domain and the trusted domain failed."}
	ErrorTrustedRelationshipFailure                        = &Error{0x000006FD, "ERROR_TRUSTED_RELATIONSHIP_FAILURE", "The trust relationship between this workstation and the primary domain failed."}
	ErrorTrustFailure                                      = &Error{0x000006FE, "ERROR_TRUST_FAILURE", "The network logon failed."}
	RpcSCallInProgress                                     = &Error{0x000006FF, "RPC_S_CALL_IN_PROGRESS", "An RPC is already in progress for this thread."}
	ErrorNetlogonNotStarted                                = &Error{0x00000700, "ERROR_NETLOGON_NOT_STARTED", "An attempt was made to log on, but the network logon service was not started."}
	ErrorAccountExpired                                    = &Error{0x00000701, "ERROR_ACCOUNT_EXPIRED", "The user's account has expired."}
	ErrorRedirectorHasOpenHandles                          = &Error{0x00000702, "ERROR_REDIRECTOR_HAS_OPEN_HANDLES", "The redirector is in use and cannot be unloaded."}
	ErrorPrinterDriverAlreadyInstalled                     = &Error{0x00000703, "ERROR_PRINTER_DRIVER_ALREADY_INSTALLED", "The specified printer driver is already installed."}
	ErrorUnknownPort                                       = &Error{0x00000704, "ERROR_UNKNOWN_PORT", "The specified port is unknown."}
	ErrorUnknownPrinterDriver                              = &Error{0x00000705, "ERROR_UNKNOWN_PRINTER_DRIVER", "The printer driver is unknown."}
	ErrorUnknownPrintprocessor                             = &Error{0x00000706, "ERROR_UNKNOWN_PRINTPROCESSOR", "The print processor is unknown."}
	ErrorInvalidSeparatorFile                              = &Error{0x00000707, "ERROR_INVALID_SEPARATOR_FILE", "The specified separator file is invalid."}
	ErrorInvalidPriority                                   = &Error{0x00000708, "ERROR_INVALID_PRIORITY", "The specified priority is invalid."}
	ErrorInvalidPrinterName                                = &Error{0x00000709, "ERROR_INVALID_PRINTER_NAME", "The printer name is invalid."}
	ErrorPrinterAlreadyExists                              = &Error{0x0000070A, "ERROR_PRINTER_ALREADY_EXISTS", "The printer already exists."}
	ErrorInvalidPrinterCommand                             = &Error{0x0000070B, "ERROR_INVALID_PRINTER_COMMAND", "The printer command is invalid."}
	ErrorInvalidDatatype                                   = &Error{0x0000070C, "ERROR_INVALID_DATATYPE", "The specified data type is invalid."}
	ErrorInvalidEnvironment                                = &Error{0x0000070D, "ERROR_INVALID_ENVIRONMENT", "The environment specified is invalid."}
	RpcSNoMoreBindings                                     = &Error{0x0000070E, "RPC_S_NO_MORE_BINDINGS", "There are no more bindings."}
	ErrorNologonInterdomainTrustAccount                    = &Error{0x0000070F, "ERROR_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT", "The account used is an interdomain trust account. Use your global user account or local user account to access this server."}
	ErrorNologonWorkstationTrustAccount                    = &Error{0x00000710, "ERROR_NOLOGON_WORKSTATION_TRUST_ACCOUNT", "The account used is a computer account. Use your global user account or local user account to access this server."}
	ErrorNologonServerTrustAccount                         = &Error{0x00000711, "ERROR_NOLOGON_SERVER_TRUST_ACCOUNT", "The account used is a server trust account. Use your global user account or local user account to access this server."}
	ErrorDomainTrustInconsistent                           = &Error{0x00000712, "ERROR_DOMAIN_TRUST_INCONSISTENT", "The name or SID of the domain specified is inconsistent with the trust information for that domain."}
	ErrorServerHasOpenHandles                              = &Error{0x00000713, "ERROR_SERVER_HAS_OPEN_HANDLES", "The server is in use and cannot be unloaded."}
	ErrorResourceDataNotFound                              = &Error{0x00000714, "ERROR_RESOURCE_DATA_NOT_FOUND", "The specified image file did not contain a resource section."}
	ErrorResourceTypeNotFound                              = &Error{0x00000715, "ERROR_RESOURCE_TYPE_NOT_FOUND", "The specified resource type cannot be found in the image file."}
	ErrorResourceNameNotFound                              = &Error{0x00000716, "ERROR_RESOURCE_NAME_NOT_FOUND", "The specified resource name cannot be found in the image file."}
	ErrorResourceLangNotFound                              = &Error{0x00000717, "ERROR_RESOURCE_LANG_NOT_FOUND", "The specified resource language ID cannot be found in the image file."}
	ErrorNotEnoughQuota                                    = &Error{0x00000718, "ERROR_NOT_ENOUGH_QUOTA", "Not enough quota is available to process this command."}
	RpcSNoInterfaces                                       = &Error{0x00000719, "RPC_S_NO_INTERFACES", "No interfaces have been registered."}
	RpcSCallCancelled                                      = &Error{0x0000071A, "RPC_S_CALL_CANCELLED", "The RPC was canceled."}
	RpcSBindingIncomplete                                  = &Error{0x0000071B, "RPC_S_BINDING_INCOMPLETE", "The binding handle does not contain all the required information."}
	RpcSCommFailure                                        = &Error{0x0000071C, "RPC_S_COMM_FAILURE", "A communications failure occurred during an RPC."}
	RpcSUnsupportedAuthnLevel                              = &Error{0x0000071D, "RPC_S_UNSUPPORTED_AUTHN_LEVEL", "The requested authentication level is not supported."}
	RpcSNoPrincName                                        = &Error{0x0000071E, "RPC_S_NO_PRINC_NAME", "No principal name is registered."}
	RpcSNotRpcError                                        = &Error{0x0000071F, "RPC_S_NOT_RPC_ERROR", "The error specified is not a valid Windows RPC error code."}
	RpcSUuidLocalOnly                                      = &Error{0x00000720, "RPC_S_UUID_LOCAL_ONLY", "A UUID that is valid only on this computer has been allocated."}
	RpcSSecPkgError                                        = &Error{0x00000721, "RPC_S_SEC_PKG_ERROR", "A security package-specific error occurred."}
	RpcSNotCancelled                                       = &Error{0x00000722, "RPC_S_NOT_CANCELLED", "The thread is not canceled."}
	RpcXInvalidEsAction                                    = &Error{0x00000723, "RPC_X_INVALID_ES_ACTION", "Invalid operation on the encoding/decoding handle."}
	RpcXWrongEsVersion                                     = &Error{0x00000724, "RPC_X_WRONG_ES_VERSION", "Incompatible version of the serializing package."}
	RpcXWrongStubVersion                                   = &Error{0x00000725, "RPC_X_WRONG_STUB_VERSION", "Incompatible version of the RPC stub."}
	RpcXInvalidPipeObject                                  = &Error{0x00000726, "RPC_X_INVALID_PIPE_OBJECT", "The RPC pipe object is invalid or corrupted."}
	RpcXWrongPipeOrder                                     = &Error{0x00000727, "RPC_X_WRONG_PIPE_ORDER", "An invalid operation was attempted on an RPC pipe object."}
	RpcXWrongPipeVersion                                   = &Error{0x00000728, "RPC_X_WRONG_PIPE_VERSION", "Unsupported RPC pipe version."}
	RpcSGroupMemberNotFound                                = &Error{0x0000076A, "RPC_S_GROUP_MEMBER_NOT_FOUND", "The group member was not found."}
	EptSCantCreate                                         = &Error{0x0000076B, "EPT_S_CANT_CREATE", "The endpoint mapper database entry could not be created."}
	RpcSInvalidObject                                      = &Error{0x0000076C, "RPC_S_INVALID_OBJECT", "The object UUID is the nil UUID."}
	ErrorInvalidTime                                       = &Error{0x0000076D, "ERROR_INVALID_TIME", "The specified time is invalid."}
	ErrorInvalidFormName                                   = &Error{0x0000076E, "ERROR_INVALID_FORM_NAME", "The specified form name is invalid."}
	ErrorInvalidFormSize                                   = &Error{0x0000076F, "ERROR_INVALID_FORM_SIZE", "The specified form size is invalid."}
	ErrorAlreadyWaiting                                    = &Error{0x00000770, "ERROR_ALREADY_WAITING", "The specified printer handle is already being waited on."}
	ErrorPrinterDeleted                                    = &Error{0x00000771, "ERROR_PRINTER_DELETED", "The specified printer has been deleted."}
	ErrorInvalidPrinterState                               = &Error{0x00000772, "ERROR_INVALID_PRINTER_STATE", "The state of the printer is invalid."}
	ErrorPasswordMustChange                                = &Error{0x00000773, "ERROR_PASSWORD_MUST_CHANGE", "The user's password must be changed before logging on the first time."}
	ErrorDomainControllerNotFound                          = &Error{0x00000774, "ERROR_DOMAIN_CONTROLLER_NOT_FOUND", "Could not find the domain controller for this domain."}
	ErrorAccountLockedOut                                  = &Error{0x00000775, "ERROR_ACCOUNT_LOCKED_OUT", "The referenced account is currently locked out and cannot be logged on to."}
	OrInvalidOxid                                          = &Error{0x00000776, "OR_INVALID_OXID", "The object exporter specified was not found."}
	OrInvalidOid                                           = &Error{0x00000777, "OR_INVALID_OID", "The object specified was not found."}
	OrInvalidSet                                           = &Error{0x00000778, "OR_INVALID_SET", "The object set specified was not found."}
	RpcSSendIncomplete                                     = &Error{0x00000779, "RPC_S_SEND_INCOMPLETE", "Some data remains to be sent in the request buffer."}
	RpcSInvalidAsyncHandle                                 = &Error{0x0000077A, "RPC_S_INVALID_ASYNC_HANDLE", "Invalid asynchronous RPC handle."}
	RpcSInvalidAsyncCall                                   = &Error{0x0000077B, "RPC_S_INVALID_ASYNC_CALL", "Invalid asynchronous RPC call handle for this operation."}
	RpcXPipeClosed                                         = &Error{0x0000077C, "RPC_X_PIPE_CLOSED", "The RPC pipe object has already been closed."}
	RpcXPipeDisciplineError                                = &Error{0x0000077D, "RPC_X_PIPE_DISCIPLINE_ERROR", "The RPC call completed before all pipes were processed."}
	RpcXPipeEmpty                                          = &Error{0x0000077E, "RPC_X_PIPE_EMPTY", "No more data is available from the RPC pipe."}
	ErrorNoSitename                                        = &Error{0x0000077F, "ERROR_NO_SITENAME", "No site name is available for this machine."}
	ErrorCantAccessFile                                    = &Error{0x00000780, "ERROR_CANT_ACCESS_FILE", "The file cannot be accessed by the system."}
	ErrorCantResolveFilename                               = &Error{0x00000781, "ERROR_CANT_RESOLVE_FILENAME", "The name of the file cannot be resolved by the system."}
	RpcSEntryTypeMismatch                                  = &Error{0x00000782, "RPC_S_ENTRY_TYPE_MISMATCH", "The entry is not of the expected type."}
	RpcSNotAllObjsExported                                 = &Error{0x00000783, "RPC_S_NOT_ALL_OBJS_EXPORTED", "Not all object UUIDs could be exported to the specified entry."}
	RpcSInterfaceNotExported                               = &Error{0x00000784, "RPC_S_INTERFACE_NOT_EXPORTED", "The interface could not be exported to the specified entry."}
	RpcSProfileNotAdded                                    = &Error{0x00000785, "RPC_S_PROFILE_NOT_ADDED", "The specified profile entry could not be added."}
	RpcSPrfEltNotAdded                                     = &Error{0x00000786, "RPC_S_PRF_ELT_NOT_ADDED", "The specified profile element could not be added."}
	RpcSPrfEltNotRemoved                                   = &Error{0x00000787, "RPC_S_PRF_ELT_NOT_REMOVED", "The specified profile element could not be removed."}
	RpcSGrpEltNotAdded                                     = &Error{0x00000788, "RPC_S_GRP_ELT_NOT_ADDED", "The group element could not be added."}
	RpcSGrpEltNotRemoved                                   = &Error{0x00000789, "RPC_S_GRP_ELT_NOT_REMOVED", "The group element could not be removed."}
	ErrorKmDriverBlocked                                   = &Error{0x0000078A, "ERROR_KM_DRIVER_BLOCKED", "The printer driver is not compatible with a policy enabled on your computer that blocks Windows NT 4.0 operating system drivers."}
	ErrorContextExpired                                    = &Error{0x0000078B, "ERROR_CONTEXT_EXPIRED", "The context has expired and can no longer be used."}
	ErrorPerUserTrustQuotaExceeded                         = &Error{0x0000078C, "ERROR_PER_USER_TRUST_QUOTA_EXCEEDED", "The current user's delegated trust creation quota has been exceeded."}
	ErrorAllUserTrustQuotaExceeded                         = &Error{0x0000078D, "ERROR_ALL_USER_TRUST_QUOTA_EXCEEDED", "The total delegated trust creation quota has been exceeded."}
	ErrorUserDeleteTrustQuotaExceeded                      = &Error{0x0000078E, "ERROR_USER_DELETE_TRUST_QUOTA_EXCEEDED", "The current user's delegated trust deletion quota has been exceeded."}
	ErrorAuthenticationFirewallFailed                      = &Error{0x0000078F, "ERROR_AUTHENTICATION_FIREWALL_FAILED", "Logon failure: The machine you are logging on to is protected by an authentication firewall. The specified account is not allowed to authenticate to the machine."}
	ErrorRemotePrintConnectionsBlocked                     = &Error{0x00000790, "ERROR_REMOTE_PRINT_CONNECTIONS_BLOCKED", "Remote connections to the Print Spooler are blocked by a policy set on your machine."}
	ErrorInvalidPixelFormat                                = &Error{0x000007D0, "ERROR_INVALID_PIXEL_FORMAT", "The pixel format is invalid."}
	ErrorBadDriver                                         = &Error{0x000007D1, "ERROR_BAD_DRIVER", "The specified driver is invalid."}
	ErrorInvalidWindowStyle                                = &Error{0x000007D2, "ERROR_INVALID_WINDOW_STYLE", "The window style or class attribute is invalid for this operation."}
	ErrorMetafileNotSupported                              = &Error{0x000007D3, "ERROR_METAFILE_NOT_SUPPORTED", "The requested metafile operation is not supported."}
	ErrorTransformNotSupported                             = &Error{0x000007D4, "ERROR_TRANSFORM_NOT_SUPPORTED", "The requested transformation operation is not supported."}
	ErrorClippingNotSupported                              = &Error{0x000007D5, "ERROR_CLIPPING_NOT_SUPPORTED", "The requested clipping operation is not supported."}
	ErrorInvalidCmm                                        = &Error{0x000007DA, "ERROR_INVALID_CMM", "The specified color management module is invalid."}
	ErrorInvalidProfile                                    = &Error{0x000007DB, "ERROR_INVALID_PROFILE", "The specified color profile is invalid."}
	ErrorTagNotFound                                       = &Error{0x000007DC, "ERROR_TAG_NOT_FOUND", "The specified tag was not found."}
	ErrorTagNotPresent                                     = &Error{0x000007DD, "ERROR_TAG_NOT_PRESENT", "A required tag is not present."}
	ErrorDuplicateTag                                      = &Error{0x000007DE, "ERROR_DUPLICATE_TAG", "The specified tag is already present."}
	ErrorProfileNotAssociatedWithDevice                    = &Error{0x000007DF, "ERROR_PROFILE_NOT_ASSOCIATED_WITH_DEVICE", "The specified color profile is not associated with any device."}
	ErrorProfileNotFound                                   = &Error{0x000007E0, "ERROR_PROFILE_NOT_FOUND", "The specified color profile was not found."}
	ErrorInvalidColorspace                                 = &Error{0x000007E1, "ERROR_INVALID_COLORSPACE", "The specified color space is invalid."}
	ErrorIcmNotEnabled                                     = &Error{0x000007E2, "ERROR_ICM_NOT_ENABLED", "Image Color Management is not enabled."}
	ErrorDeletingIcmXform                                  = &Error{0x000007E3, "ERROR_DELETING_ICM_XFORM", "There was an error while deleting the color transform."}
	ErrorInvalidTransform                                  = &Error{0x000007E4, "ERROR_INVALID_TRANSFORM", "The specified color transform is invalid."}
	ErrorColorspaceMismatch                                = &Error{0x000007E5, "ERROR_COLORSPACE_MISMATCH", "The specified transform does not match the bitmap's color space."}
	ErrorInvalidColorindex                                 = &Error{0x000007E6, "ERROR_INVALID_COLORINDEX", "The specified named color index is not present in the profile."}
	ErrorProfileDoesNotMatchDevice                         = &Error{0x000007E7, "ERROR_PROFILE_DOES_NOT_MATCH_DEVICE", "The specified profile is intended for a device of a different type than the specified device."}
	NerrNetnotstarted                                      = &Error{0x00000836, "NERR_NetNotStarted", "The workstation driver is not installed."}
	NerrUnknownserver                                      = &Error{0x00000837, "NERR_UnknownServer", "The server could not be located."}
	NerrSharemem                                           = &Error{0x00000838, "NERR_ShareMem", "An internal error occurred. The network cannot access a shared memory segment."}
	NerrNonetworkresource                                  = &Error{0x00000839, "NERR_NoNetworkResource", "A network resource shortage occurred."}
	NerrRemoteonly                                         = &Error{0x0000083A, "NERR_RemoteOnly", "This operation is not supported on workstations."}
	NerrDevnotredirected                                   = &Error{0x0000083B, "NERR_DevNotRedirected", "The device is not connected."}
	ErrorConnectedOtherPassword                            = &Error{0x0000083C, "ERROR_CONNECTED_OTHER_PASSWORD", "The network connection was made successfully, but the user had to be prompted for a password other than the one originally specified."}
	ErrorConnectedOtherPasswordDefault                     = &Error{0x0000083D, "ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT", "The network connection was made successfully using default credentials."}
	NerrServernotstarted                                   = &Error{0x00000842, "NERR_ServerNotStarted", "The Server service is not started."}
	NerrItemnotfound                                       = &Error{0x00000843, "NERR_ItemNotFound", "The queue is empty."}
	NerrUnknowndevdir                                      = &Error{0x00000844, "NERR_UnknownDevDir", "The device or directory does not exist."}
	NerrRedirectedpath                                     = &Error{0x00000845, "NERR_RedirectedPath", "The operation is invalid on a redirected resource."}
	NerrDuplicateshare                                     = &Error{0x00000846, "NERR_DuplicateShare", "The name has already been shared."}
	NerrNoroom                                             = &Error{0x00000847, "NERR_NoRoom", "The server is currently out of the requested resource."}
	NerrToomanyitems                                       = &Error{0x00000849, "NERR_TooManyItems", "Requested addition of items exceeds the maximum allowed."}
	NerrInvalidmaxusers                                    = &Error{0x0000084A, "NERR_InvalidMaxUsers", "The Peer service supports only two simultaneous users."}
	NerrBuftoosmall                                        = &Error{0x0000084B, "NERR_BufTooSmall", "The API return buffer is too small."}
	NerrRemoteerr                                          = &Error{0x0000084F, "NERR_RemoteErr", "A remote API error occurred."}
	NerrLanmaninierror                                     = &Error{0x00000853, "NERR_LanmanIniError", "An error occurred when opening or reading the configuration file."}
	NerrNetworkerror                                       = &Error{0x00000858, "NERR_NetworkError", "A general network error occurred."}
	NerrWkstainconsistentstate                             = &Error{0x00000859, "NERR_WkstaInconsistentState", "The Workstation service is in an inconsistent state. Restart the computer before restarting the Workstation service."}
	NerrWkstanotstarted                                    = &Error{0x0000085A, "NERR_WkstaNotStarted", "The Workstation service has not been started."}
	NerrBrowsernotstarted                                  = &Error{0x0000085B, "NERR_BrowserNotStarted", "The requested information is not available."}
	NerrInternalerror                                      = &Error{0x0000085C, "NERR_InternalError", "An internal error occurred."}
	NerrBadtransactconfig                                  = &Error{0x0000085D, "NERR_BadTransactConfig", "The server is not configured for transactions."}
	NerrInvalidapi                                         = &Error{0x0000085E, "NERR_InvalidAPI", "The requested API is not supported on the remote server."}
	NerrBadeventname                                       = &Error{0x0000085F, "NERR_BadEventName", "The event name is invalid."}
	NerrDupnamereboot                                      = &Error{0x00000860, "NERR_DupNameReboot", "The computer name already exists on the network. Change it and reboot the computer."}
	NerrCfgcompnotfound                                    = &Error{0x00000862, "NERR_CfgCompNotFound", "The specified component could not be found in the configuration information."}
	NerrCfgparamnotfound                                   = &Error{0x00000863, "NERR_CfgParamNotFound", "The specified parameter could not be found in the configuration information."}
	NerrLinetoolong                                        = &Error{0x00000865, "NERR_LineTooLong", "A line in the configuration file is too long."}
	NerrQnotfound                                          = &Error{0x00000866, "NERR_QNotFound", "The printer does not exist."}
	NerrJobnotfound                                        = &Error{0x00000867, "NERR_JobNotFound", "The print job does not exist."}
	NerrDestnotfound                                       = &Error{0x00000868, "NERR_DestNotFound", "The printer destination cannot be found."}
	NerrDestexists                                         = &Error{0x00000869, "NERR_DestExists", "The printer destination already exists."}
	NerrQexists                                            = &Error{0x0000086A, "NERR_QExists", "The print queue already exists."}
	NerrQnoroom                                            = &Error{0x0000086B, "NERR_QNoRoom", "No more printers can be added."}
	NerrJobnoroom                                          = &Error{0x0000086C, "NERR_JobNoRoom", "No more print jobs can be added."}
	NerrDestnoroom                                         = &Error{0x0000086D, "NERR_DestNoRoom", "No more printer destinations can be added."}
	NerrDestidle                                           = &Error{0x0000086E, "NERR_DestIdle", "This printer destination is idle and cannot accept control operations."}
	NerrDestinvalidop                                      = &Error{0x0000086F, "NERR_DestInvalidOp", "This printer destination request contains an invalid control function."}
	NerrProcnorespond                                      = &Error{0x00000870, "NERR_ProcNoRespond", "The print processor is not responding."}
	NerrSpoolernotloaded                                   = &Error{0x00000871, "NERR_SpoolerNotLoaded", "The spooler is not running."}
	NerrDestinvalidstate                                   = &Error{0x00000872, "NERR_DestInvalidState", "This operation cannot be performed on the print destination in its current state."}
	NerrQinvalidstate                                      = &Error{0x00000873, "NERR_QinvalidState", "This operation cannot be performed on the print queue in its current state."}
	NerrJobinvalidstate                                    = &Error{0x00000874, "NERR_JobInvalidState", "This operation cannot be performed on the print job in its current state."}
	NerrSpoolnomemory                                      = &Error{0x00000875, "NERR_SpoolNoMemory", "A spooler memory allocation failure occurred."}
	NerrDrivernotfound                                     = &Error{0x00000876, "NERR_DriverNotFound", "The device driver does not exist."}
	NerrDatatypeinvalid                                    = &Error{0x00000877, "NERR_DataTypeInvalid", "The data type is not supported by the print processor."}
	NerrProcnotfound                                       = &Error{0x00000878, "NERR_ProcNotFound", "The print processor is not installed."}
	NerrServicetablelocked                                 = &Error{0x00000884, "NERR_ServiceTableLocked", "The service database is locked."}
	NerrServicetablefull                                   = &Error{0x00000885, "NERR_ServiceTableFull", "The service table is full."}
	NerrServiceinstalled                                   = &Error{0x00000886, "NERR_ServiceInstalled", "The requested service has already been started."}
	NerrServiceentrylocked                                 = &Error{0x00000887, "NERR_ServiceEntryLocked", "The service does not respond to control actions."}
	NerrServicenotinstalled                                = &Error{0x00000888, "NERR_ServiceNotInstalled", "The service has not been started."}
	NerrBadservicename                                     = &Error{0x00000889, "NERR_BadServiceName", "The service name is invalid."}
	NerrServicectltimeout                                  = &Error{0x0000088A, "NERR_ServiceCtlTimeout", "The service is not responding to the control function."}
	NerrServicectlbusy                                     = &Error{0x0000088B, "NERR_ServiceCtlBusy", "The service control is busy."}
	NerrBadserviceprogname                                 = &Error{0x0000088C, "NERR_BadServiceProgName", "The configuration file contains an invalid service program name."}
	NerrServicenotctrl                                     = &Error{0x0000088D, "NERR_ServiceNotCtrl", "The service could not be controlled in its present state."}
	NerrServicekillproc                                    = &Error{0x0000088E, "NERR_ServiceKillProc", "The service ended abnormally."}
	NerrServicectlnotvalid                                 = &Error{0x0000088F, "NERR_ServiceCtlNotValid", "The requested pause or stop is not valid for this service."}
	NerrNotindispatchtbl                                   = &Error{0x00000890, "NERR_NotInDispatchTbl", "The service control dispatcher could not find the service name in the dispatch table."}
	NerrBadcontrolrecv                                     = &Error{0x00000891, "NERR_BadControlRecv", "The service control dispatcher pipe read failed."}
	NerrServicenotstarting                                 = &Error{0x00000892, "NERR_ServiceNotStarting", "A thread for the new service could not be created."}
	NerrAlreadyloggedon                                    = &Error{0x00000898, "NERR_AlreadyLoggedOn", "This workstation is already logged on to the LAN."}
	NerrNotloggedon                                        = &Error{0x00000899, "NERR_NotLoggedOn", "The workstation is not logged on to the LAN."}
	NerrBadusername                                        = &Error{0x0000089A, "NERR_BadUsername", "The user name or group name parameter is invalid."}
	NerrBadpassword                                        = &Error{0x0000089B, "NERR_BadPassword", "The password parameter is invalid."}
	NerrUnabletoaddnameW                                   = &Error{0x0000089C, "NERR_UnableToAddName_W", "The logon processor did not add the message alias."}
	NerrUnabletoaddnameF                                   = &Error{0x0000089D, "NERR_UnableToAddName_F", "The logon processor did not add the message alias."}
	NerrUnabletodelnameW                                   = &Error{0x0000089E, "NERR_UnableToDelName_W", "The logoff processor did not delete the message alias."}
	NerrUnabletodelnameF                                   = &Error{0x0000089F, "NERR_UnableToDelName_F", "The logoff processor did not delete the message alias."}
	NerrLogonspaused                                       = &Error{0x000008A1, "NERR_LogonsPaused", "Network logons are paused."}
	NerrLogonserverconflict                                = &Error{0x000008A2, "NERR_LogonServerConflict", "A centralized logon server conflict occurred."}
	NerrLogonnouserpath                                    = &Error{0x000008A3, "NERR_LogonNoUserPath", "The server is configured without a valid user path."}
	NerrLogonscripterror                                   = &Error{0x000008A4, "NERR_LogonScriptError", "An error occurred while loading or running the logon script."}
	NerrStandalonelogon                                    = &Error{0x000008A6, "NERR_StandaloneLogon", "The logon server was not specified. The computer will be logged on as STANDALONE."}
	NerrLogonservernotfound                                = &Error{0x000008A7, "NERR_LogonServerNotFound", "The logon server could not be found."}
	NerrLogondomainexists                                  = &Error{0x000008A8, "NERR_LogonDomainExists", "There is already a logon domain for this computer."}
	NerrNonvalidatedlogon                                  = &Error{0x000008A9, "NERR_NonValidatedLogon", "The logon server could not validate the logon."}
	NerrAcfnotfound                                        = &Error{0x000008AB, "NERR_ACFNotFound", "The security database could not be found."}
	NerrGroupnotfound                                      = &Error{0x000008AC, "NERR_GroupNotFound", "The group name could not be found."}
	NerrUsernotfound                                       = &Error{0x000008AD, "NERR_UserNotFound", "The user name could not be found."}
	NerrResourcenotfound                                   = &Error{0x000008AE, "NERR_ResourceNotFound", "The resource name could not be found."}
	NerrGroupexists                                        = &Error{0x000008AF, "NERR_GroupExists", "The group already exists."}
	NerrUserexists                                         = &Error{0x000008B0, "NERR_UserExists", "The user account already exists."}
	NerrResourceexists                                     = &Error{0x000008B1, "NERR_ResourceExists", "The resource permission list already exists."}
	NerrNotprimary                                         = &Error{0x000008B2, "NERR_NotPrimary", "This operation is allowed only on the PDC of the domain."}
	NerrAcfnotloaded                                       = &Error{0x000008B3, "NERR_ACFNotLoaded", "The security database has not been started."}
	NerrAcfnoroom                                          = &Error{0x000008B4, "NERR_ACFNoRoom", "There are too many names in the user accounts database."}
	NerrAcffileiofail                                      = &Error{0x000008B5, "NERR_ACFFileIOFail", "A disk I/O failure occurred."}
	NerrAcftoomanylists                                    = &Error{0x000008B6, "NERR_ACFTooManyLists", "The limit of 64 entries per resource was exceeded."}
	NerrUserlogon                                          = &Error{0x000008B7, "NERR_UserLogon", "Deleting a user with a session is not allowed."}
	NerrAcfnoparent                                        = &Error{0x000008B8, "NERR_ACFNoParent", "The parent directory could not be located."}
	NerrCannotgrowsegment                                  = &Error{0x000008B9, "NERR_CanNotGrowSegment", "Unable to add to the security database session cache segment."}
	NerrSpegroupop                                         = &Error{0x000008BA, "NERR_SpeGroupOp", "This operation is not allowed on this special group."}
	NerrNotincache                                         = &Error{0x000008BB, "NERR_NotInCache", "This user is not cached in the user accounts database session cache."}
	NerrUseringroup                                        = &Error{0x000008BC, "NERR_UserInGroup", "The user already belongs to this group."}
	NerrUsernotingroup                                     = &Error{0x000008BD, "NERR_UserNotInGroup", "The user does not belong to this group."}
	NerrAccountundefined                                   = &Error{0x000008BE, "NERR_AccountUndefined", "This user account is undefined."}
	NerrAccountexpired                                     = &Error{0x000008BF, "NERR_AccountExpired", "This user account has expired."}
	NerrInvalidworkstation                                 = &Error{0x000008C0, "NERR_InvalidWorkstation", "The user is not allowed to log on from this workstation."}
	NerrInvalidlogonhours                                  = &Error{0x000008C1, "NERR_InvalidLogonHours", "The user is not allowed to log on at this time."}
	NerrPasswordexpired                                    = &Error{0x000008C2, "NERR_PasswordExpired", "The password of this user has expired."}
	NerrPasswordcantchange                                 = &Error{0x000008C3, "NERR_PasswordCantChange", "The password of this user cannot change."}
	NerrPasswordhistconflict                               = &Error{0x000008C4, "NERR_PasswordHistConflict", "This password cannot be used now."}
	NerrPasswordtooshort                                   = &Error{0x000008C5, "NERR_PasswordTooShort", "The password does not meet the password policy requirements. Check the minimum password length, password complexity, and password history requirements."}
	NerrPasswordtoorecent                                  = &Error{0x000008C6, "NERR_PasswordTooRecent", "The password of this user is too recent to change."}
	NerrInvaliddatabase                                    = &Error{0x000008C7, "NERR_InvalidDatabase", "The security database is corrupted."}
	NerrDatabaseuptodate                                   = &Error{0x000008C8, "NERR_DatabaseUpToDate", "No updates are necessary to this replicant network or local security database."}
	NerrSyncrequired                                       = &Error{0x000008C9, "NERR_SyncRequired", "This replicant database is outdated; synchronization is required."}
	NerrUsenotfound                                        = &Error{0x000008CA, "NERR_UseNotFound", "The network connection could not be found."}
	NerrBadasgtype                                         = &Error{0x000008CB, "NERR_BadAsgType", "This asg_type is invalid."}
	NerrDeviceisshared                                     = &Error{0x000008CC, "NERR_DeviceIsShared", "This device is currently being shared."}
	NerrNocomputername                                     = &Error{0x000008DE, "NERR_NoComputerName", "The computer name could not be added as a message alias. The name might already exist on the network."}
	NerrMsgalreadystarted                                  = &Error{0x000008DF, "NERR_MsgAlreadyStarted", "The Messenger service is already started."}
	NerrMsginitfailed                                      = &Error{0x000008E0, "NERR_MsgInitFailed", "The Messenger service failed to start."}
	NerrNamenotfound                                       = &Error{0x000008E1, "NERR_NameNotFound", "The message alias could not be found on the network."}
	NerrAlreadyforwarded                                   = &Error{0x000008E2, "NERR_AlreadyForwarded", "This message alias has already been forwarded."}
	NerrAddforwarded                                       = &Error{0x000008E3, "NERR_AddForwarded", "This message alias has been added but is still forwarded."}
	NerrAlreadyexists                                      = &Error{0x000008E4, "NERR_AlreadyExists", "This message alias already exists locally."}
	NerrToomanynames                                       = &Error{0x000008E5, "NERR_TooManyNames", "The maximum number of added message aliases has been exceeded."}
	NerrDelcomputername                                    = &Error{0x000008E6, "NERR_DelComputerName", "The computer name could not be deleted."}
	NerrLocalforward                                       = &Error{0x000008E7, "NERR_LocalForward", "Messages cannot be forwarded back to the same workstation."}
	NerrGrpmsgprocessor                                    = &Error{0x000008E8, "NERR_GrpMsgProcessor", "An error occurred in the domain message processor."}
	NerrPausedremote                                       = &Error{0x000008E9, "NERR_PausedRemote", "The message was sent, but the recipient has paused the Messenger service."}
	NerrBadreceive                                         = &Error{0x000008EA, "NERR_BadReceive", "The message was sent but not received."}
	NerrNameinuse                                          = &Error{0x000008EB, "NERR_NameInUse", "The message alias is currently in use. Try again later."}
	NerrMsgnotstarted                                      = &Error{0x000008EC, "NERR_MsgNotStarted", "The Messenger service has not been started."}
	NerrNotlocalname                                       = &Error{0x000008ED, "NERR_NotLocalName", "The name is not on the local computer."}
	NerrNoforwardname                                      = &Error{0x000008EE, "NERR_NoForwardName", "The forwarded message alias could not be found on the network."}
	NerrRemotefull                                         = &Error{0x000008EF, "NERR_RemoteFull", "The message alias table on the remote station is full."}
	NerrNamenotforwarded                                   = &Error{0x000008F0, "NERR_NameNotForwarded", "Messages for this alias are not currently being forwarded."}
	NerrTruncatedbroadcast                                 = &Error{0x000008F1, "NERR_TruncatedBroadcast", "The broadcast message was truncated."}
	NerrInvaliddevice                                      = &Error{0x000008F6, "NERR_InvalidDevice", "This is an invalid device name."}
	NerrWritefault                                         = &Error{0x000008F7, "NERR_WriteFault", "A write fault occurred."}
	NerrDuplicatename                                      = &Error{0x000008F9, "NERR_DuplicateName", "A duplicate message alias exists on the network."}
	NerrDeletelater                                        = &Error{0x000008FA, "NERR_DeleteLater", "This message alias will be deleted later."}
	NerrIncompletedel                                      = &Error{0x000008FB, "NERR_IncompleteDel", "The message alias was not successfully deleted from all networks."}
	NerrMultiplenets                                       = &Error{0x000008FC, "NERR_MultipleNets", "This operation is not supported on computers with multiple networks."}
	NerrNetnamenotfound                                    = &Error{0x00000906, "NERR_NetNameNotFound", "This shared resource does not exist."}
	NerrDevicenotshared                                    = &Error{0x00000907, "NERR_DeviceNotShared", "This device is not shared."}
	NerrClientnamenotfound                                 = &Error{0x00000908, "NERR_ClientNameNotFound", "A session does not exist with that computer name."}
	NerrFileidnotfound                                     = &Error{0x0000090A, "NERR_FileIdNotFound", "There is not an open file with that identification number."}
	NerrExecfailure                                        = &Error{0x0000090B, "NERR_ExecFailure", "A failure occurred when executing a remote administration command."}
	NerrTmpfile                                            = &Error{0x0000090C, "NERR_TmpFile", "A failure occurred when opening a remote temporary file."}
	NerrToomuchdata                                        = &Error{0x0000090D, "NERR_TooMuchData", "The data returned from a remote administration command has been truncated to 64 KB."}
	NerrDeviceshareconflict                                = &Error{0x0000090E, "NERR_DeviceShareConflict", "This device cannot be shared as both a spooled and a nonspooled resource."}
	NerrBrowsertableincomplete                             = &Error{0x0000090F, "NERR_BrowserTableIncomplete", "The information in the list of servers might be incorrect."}
	NerrNotlocaldomain                                     = &Error{0x00000910, "NERR_NotLocalDomain", "The computer is not active in this domain."}
	NerrIsdfsshare                                         = &Error{0x00000911, "NERR_IsDfsShare", "The share must be removed from the Distributed File System (DFS) before it can be deleted."}
	NerrDevinvalidopcode                                   = &Error{0x0000091B, "NERR_DevInvalidOpCode", "The operation is invalid for this device."}
	NerrDevnotfound                                        = &Error{0x0000091C, "NERR_DevNotFound", "This device cannot be shared."}
	NerrDevnotopen                                         = &Error{0x0000091D, "NERR_DevNotOpen", "This device was not open."}
	NerrBadqueuedevstring                                  = &Error{0x0000091E, "NERR_BadQueueDevString", "This device name list is invalid."}
	NerrBadqueuepriority                                   = &Error{0x0000091F, "NERR_BadQueuePriority", "The queue priority is invalid."}
	NerrNocommdevs                                         = &Error{0x00000921, "NERR_NoCommDevs", "There are no shared communication devices."}
	NerrQueuenotfound                                      = &Error{0x00000922, "NERR_QueueNotFound", "The queue you specified does not exist."}
	NerrBaddevstring                                       = &Error{0x00000924, "NERR_BadDevString", "This list of devices is invalid."}
	NerrBaddev                                             = &Error{0x00000925, "NERR_BadDev", "The requested device is invalid."}
	NerrInusebyspooler                                     = &Error{0x00000926, "NERR_InUseBySpooler", "This device is already in use by the spooler."}
	NerrCommdevinuse                                       = &Error{0x00000927, "NERR_CommDevInUse", "This device is already in use as a communication device."}
	NerrInvalidcomputer                                    = &Error{0x0000092F, "NERR_InvalidComputer", "This computer name is invalid."}
	NerrMaxlenexceeded                                     = &Error{0x00000932, "NERR_MaxLenExceeded", "The string and prefix specified are too long."}
	NerrBadcomponent                                       = &Error{0x00000934, "NERR_BadComponent", "This path component is invalid."}
	NerrCanttype                                           = &Error{0x00000935, "NERR_CantType", "Could not determine the type of input."}
	NerrToomanyentries                                     = &Error{0x0000093A, "NERR_TooManyEntries", "The buffer for types is not big enough."}
	NerrProfilefiletoobig                                  = &Error{0x00000942, "NERR_ProfileFileTooBig", "Profile files cannot exceed 64 KB."}
	NerrProfileoffset                                      = &Error{0x00000943, "NERR_ProfileOffset", "The start offset is out of range."}
	NerrProfilecleanup                                     = &Error{0x00000944, "NERR_ProfileCleanup", "The system cannot delete current connections to network resources."}
	NerrProfileunknowncmd                                  = &Error{0x00000945, "NERR_ProfileUnknownCmd", "The system was unable to parse the command line in this file."}
	NerrProfileloaderr                                     = &Error{0x00000946, "NERR_ProfileLoadErr", "An error occurred while loading the profile file."}
	NerrProfilesaveerr                                     = &Error{0x00000947, "NERR_ProfileSaveErr", "Errors occurred while saving the profile file. The profile was partially saved."}
	NerrLogoverflow                                        = &Error{0x00000949, "NERR_LogOverflow", "Log file %1 is full."}
	NerrLogfilechanged                                     = &Error{0x0000094A, "NERR_LogFileChanged", "This log file has changed between reads."}
	NerrLogfilecorrupt                                     = &Error{0x0000094B, "NERR_LogFileCorrupt", "Log file %1 is corrupt."}
	NerrSourceisdir                                        = &Error{0x0000094C, "NERR_SourceIsDir", "The source path cannot be a directory."}
	NerrBadsource                                          = &Error{0x0000094D, "NERR_BadSource", "The source path is illegal."}
	NerrBaddest                                            = &Error{0x0000094E, "NERR_BadDest", "The destination path is illegal."}
	NerrDifferentservers                                   = &Error{0x0000094F, "NERR_DifferentServers", "The source and destination paths are on different servers."}
	NerrRunsrvpaused                                       = &Error{0x00000951, "NERR_RunSrvPaused", "The Run server you requested is paused."}
	NerrErrcommrunsrv                                      = &Error{0x00000955, "NERR_ErrCommRunSrv", "An error occurred when communicating with a Run server."}
	NerrErrorexecingghost                                  = &Error{0x00000957, "NERR_ErrorExecingGhost", "An error occurred when starting a background process."}
	NerrSharenotfound                                      = &Error{0x00000958, "NERR_ShareNotFound", "The shared resource you are connected to could not be found."}
	NerrInvalidlana                                        = &Error{0x00000960, "NERR_InvalidLana", "The LAN adapter number is invalid."}
	NerrOpenfiles                                          = &Error{0x00000961, "NERR_OpenFiles", "There are open files on the connection."}
	NerrActiveconns                                        = &Error{0x00000962, "NERR_ActiveConns", "Active connections still exist."}
	NerrBadpasswordcore                                    = &Error{0x00000963, "NERR_BadPasswordCore", "This share name or password is invalid."}
	NerrDevinuse                                           = &Error{0x00000964, "NERR_DevInUse", "The device is being accessed by an active process."}
	NerrLocaldrive                                         = &Error{0x00000965, "NERR_LocalDrive", "The drive letter is in use locally."}
	NerrAlertexists                                        = &Error{0x0000097E, "NERR_AlertExists", "The specified client is already registered for the specified event."}
	NerrToomanyalerts                                      = &Error{0x0000097F, "NERR_TooManyAlerts", "The alert table is full."}
	NerrNosuchalert                                        = &Error{0x00000980, "NERR_NoSuchAlert", "An invalid or nonexistent alert name was raised."}
	NerrBadrecipient                                       = &Error{0x00000981, "NERR_BadRecipient", "The alert recipient is invalid."}
	NerrAcctlimitexceeded                                  = &Error{0x00000982, "NERR_AcctLimitExceeded", "A user's session with this server has been deleted."}
	NerrInvalidlogseek                                     = &Error{0x00000988, "NERR_InvalidLogSeek", "The log file does not contain the requested record number."}
	NerrBaduasconfig                                       = &Error{0x00000992, "NERR_BadUasConfig", "The user accounts database is not configured correctly."}
	NerrInvaliduasop                                       = &Error{0x00000993, "NERR_InvalidUASOp", "This operation is not permitted when the Net Logon service is running."}
	NerrLastadmin                                          = &Error{0x00000994, "NERR_LastAdmin", "This operation is not allowed on the last administrative account."}
	NerrDcnotfound                                         = &Error{0x00000995, "NERR_DCNotFound", "Could not find the domain controller for this domain."}
	NerrLogontrackingerror                                 = &Error{0x00000996, "NERR_LogonTrackingError", "Could not set logon information for this user."}
	NerrNetlogonnotstarted                                 = &Error{0x00000997, "NERR_NetlogonNotStarted", "The Net Logon service has not been started."}
	NerrCannotgrowuasfile                                  = &Error{0x00000998, "NERR_CanNotGrowUASFile", "Unable to add to the user accounts database."}
	NerrTimediffatdc                                       = &Error{0x00000999, "NERR_TimeDiffAtDC", "This server's clock is not synchronized with the PDC's clock."}
	NerrPasswordmismatch                                   = &Error{0x0000099A, "NERR_PasswordMismatch", "A password mismatch has been detected."}
	NerrNosuchserver                                       = &Error{0x0000099C, "NERR_NoSuchServer", "The server identification does not specify a valid server."}
	NerrNosuchsession                                      = &Error{0x0000099D, "NERR_NoSuchSession", "The session identification does not specify a valid session."}
	NerrNosuchconnection                                   = &Error{0x0000099E, "NERR_NoSuchConnection", "The connection identification does not specify a valid connection."}
	NerrToomanyservers                                     = &Error{0x0000099F, "NERR_TooManyServers", "There is no space for another entry in the table of available servers."}
	NerrToomanysessions                                    = &Error{0x000009A0, "NERR_TooManySessions", "The server has reached the maximum number of sessions it supports."}
	NerrToomanyconnections                                 = &Error{0x000009A1, "NERR_TooManyConnections", "The server has reached the maximum number of connections it supports."}
	NerrToomanyfiles                                       = &Error{0x000009A2, "NERR_TooManyFiles", "The server cannot open more files because it has reached its maximum number."}
	NerrNoalternateservers                                 = &Error{0x000009A3, "NERR_NoAlternateServers", "There are no alternate servers registered on this server."}
	NerrTrydownlevel                                       = &Error{0x000009A6, "NERR_TryDownLevel", "Try the down-level (remote admin protocol) version of API instead."}
	NerrUpsdrivernotstarted                                = &Error{0x000009B0, "NERR_UPSDriverNotStarted", "The uninterruptible power supply (UPS) driver could not be accessed by the UPS service."}
	NerrUpsinvalidconfig                                   = &Error{0x000009B1, "NERR_UPSInvalidConfig", "The UPS service is not configured correctly."}
	NerrUpsinvalidcommport                                 = &Error{0x000009B2, "NERR_UPSInvalidCommPort", "The UPS service could not access the specified Comm Port."}
	NerrUpssignalasserted                                  = &Error{0x000009B3, "NERR_UPSSignalAsserted", "The UPS indicated a line fail or low battery situation. Service not started."}
	NerrUpsshutdownfailed                                  = &Error{0x000009B4, "NERR_UPSShutdownFailed", "The UPS service failed to perform a system shut down."}
	NerrBaddosretcode                                      = &Error{0x000009C4, "NERR_BadDosRetCode", "The program below returned an MS-DOS error code."}
	NerrProgneedsextramem                                  = &Error{0x000009C5, "NERR_ProgNeedsExtraMem", "The program below needs more memory."}
	NerrBaddosfunction                                     = &Error{0x000009C6, "NERR_BadDosFunction", "The program below called an unsupported MS-DOS function."}
	NerrRemotebootfailed                                   = &Error{0x000009C7, "NERR_RemoteBootFailed", "The workstation failed to boot."}
	NerrBadfilechecksum                                    = &Error{0x000009C8, "NERR_BadFileCheckSum", "The file below is corrupt."}
	NerrNorplbootsystem                                    = &Error{0x000009C9, "NERR_NoRplBootSystem", "No loader is specified in the boot-block definition file."}
	NerrRplloadrnetbioserr                                 = &Error{0x000009CA, "NERR_RplLoadrNetBiosErr", "NetBIOS returned an error: The network control blocks (NCBs) and Server Message Block (SMB) are dumped above."}
	NerrRplloadrdiskerr                                    = &Error{0x000009CB, "NERR_RplLoadrDiskErr", "A disk I/O error occurred."}
	NerrImageparamerr                                      = &Error{0x000009CC, "NERR_ImageParamErr", "Image parameter substitution failed."}
	NerrToomanyimageparams                                 = &Error{0x000009CD, "NERR_TooManyImageParams", "Too many image parameters cross disk sector boundaries."}
	NerrNondosfloppyused                                   = &Error{0x000009CE, "NERR_NonDosFloppyUsed", "The image was not generated from an MS-DOS disk formatted with /S."}
	NerrRplbootrestart                                     = &Error{0x000009CF, "NERR_RplBootRestart", "Remote boot will be restarted later."}
	NerrRplsrvrcallfailed                                  = &Error{0x000009D0, "NERR_RplSrvrCallFailed", "The call to the Remoteboot server failed."}
	NerrCantconnectrplsrvr                                 = &Error{0x000009D1, "NERR_CantConnectRplSrvr", "Cannot connect to the Remoteboot server."}
	NerrCantopenimagefile                                  = &Error{0x000009D2, "NERR_CantOpenImageFile", "Cannot open image file on the Remoteboot server."}
	NerrCallingrplsrvr                                     = &Error{0x000009D3, "NERR_CallingRplSrvr", "Connecting to the Remoteboot server."}
	NerrStartingrplboot                                    = &Error{0x000009D4, "NERR_StartingRplBoot", "Connecting to the Remoteboot server."}
	NerrRplbootserviceterm                                 = &Error{0x000009D5, "NERR_RplBootServiceTerm", "Remote boot service was stopped, check the error log for the cause of the problem."}
	NerrRplbootstartfailed                                 = &Error{0x000009D6, "NERR_RplBootStartFailed", "Remote boot startup failed; check the error log for the cause of the problem."}
	NerrRplConnected                                       = &Error{0x000009D7, "NERR_RPL_CONNECTED", "A second connection to a Remoteboot resource is not allowed."}
	NerrBrowserconfiguredtonotrun                          = &Error{0x000009F6, "NERR_BrowserConfiguredToNotRun", "The browser service was configured with MaintainServerList=No."}
	NerrRplnoadaptersstarted                               = &Error{0x00000A32, "NERR_RplNoAdaptersStarted", "Service failed to start because none of the network adapters started with this service."}
	NerrRplbadregistry                                     = &Error{0x00000A33, "NERR_RplBadRegistry", "Service failed to start due to bad startup information in the registry."}
	NerrRplbaddatabase                                     = &Error{0x00000A34, "NERR_RplBadDatabase", "Service failed to start because its database is absent or corrupt."}
	NerrRplrplfilesshare                                   = &Error{0x00000A35, "NERR_RplRplfilesShare", "Service failed to start because the RPLFILES share is absent."}
	NerrRplnotrplserver                                    = &Error{0x00000A36, "NERR_RplNotRplServer", "Service failed to start because the RPLUSER group is absent."}
	NerrRplcannotenum                                      = &Error{0x00000A37, "NERR_RplCannotEnum", "Cannot enumerate service records."}
	NerrRplwkstainfocorrupted                              = &Error{0x00000A38, "NERR_RplWkstaInfoCorrupted", "Workstation record information has been corrupted."}
	NerrRplwkstanotfound                                   = &Error{0x00000A39, "NERR_RplWkstaNotFound", "Workstation record was not found."}
	NerrRplwkstanameunavailable                            = &Error{0x00000A3A, "NERR_RplWkstaNameUnavailable", "Workstation name is in use by some other workstation."}
	NerrRplprofileinfocorrupted                            = &Error{0x00000A3B, "NERR_RplProfileInfoCorrupted", "Profile record information has been corrupted."}
	NerrRplprofilenotfound                                 = &Error{0x00000A3C, "NERR_RplProfileNotFound", "Profile record was not found."}
	NerrRplprofilenameunavailable                          = &Error{0x00000A3D, "NERR_RplProfileNameUnavailable", "Profile name is in use by some other profile."}
	NerrRplprofilenotempty                                 = &Error{0x00000A3E, "NERR_RplProfileNotEmpty", "There are workstations using this profile."}
	NerrRplconfiginfocorrupted                             = &Error{0x00000A3F, "NERR_RplConfigInfoCorrupted", "Configuration record information has been corrupted."}
	NerrRplconfignotfound                                  = &Error{0x00000A40, "NERR_RplConfigNotFound", "Configuration record was not found."}
	NerrRpladapterinfocorrupted                            = &Error{0x00000A41, "NERR_RplAdapterInfoCorrupted", "Adapter ID record information has been corrupted."}
	NerrRplinternal                                        = &Error{0x00000A42, "NERR_RplInternal", "An internal service error has occurred."}
	NerrRplvendorinfocorrupted                             = &Error{0x00000A43, "NERR_RplVendorInfoCorrupted", "Vendor ID record information has been corrupted."}
	NerrRplbootinfocorrupted                               = &Error{0x00000A44, "NERR_RplBootInfoCorrupted", "Boot block record information has been corrupted."}
	NerrRplwkstaneedsuseracct                              = &Error{0x00000A45, "NERR_RplWkstaNeedsUserAcct", "The user account for this workstation record is missing."}
	NerrRplneedsrpluseracct                                = &Error{0x00000A46, "NERR_RplNeedsRPLUSERAcct", "The RPLUSER local group could not be found."}
	NerrRplbootnotfound                                    = &Error{0x00000A47, "NERR_RplBootNotFound", "Boot block record was not found."}
	NerrRplincompatibleprofile                             = &Error{0x00000A48, "NERR_RplIncompatibleProfile", "Chosen profile is incompatible with this workstation."}
	NerrRpladapternameunavailable                          = &Error{0x00000A49, "NERR_RplAdapterNameUnavailable", "Chosen network adapter ID is in use by some other workstation."}
	NerrRplconfignotempty                                  = &Error{0x00000A4A, "NERR_RplConfigNotEmpty", "There are profiles using this configuration."}
	NerrRplbootinuse                                       = &Error{0x00000A4B, "NERR_RplBootInUse", "There are workstations, profiles, or configurations using this boot block."}
	NerrRplbackupdatabase                                  = &Error{0x00000A4C, "NERR_RplBackupDatabase", "Service failed to back up the Remoteboot database."}
	NerrRpladapternotfound                                 = &Error{0x00000A4D, "NERR_RplAdapterNotFound", "Adapter record was not found."}
	NerrRplvendornotfound                                  = &Error{0x00000A4E, "NERR_RplVendorNotFound", "Vendor record was not found."}
	NerrRplvendornameunavailable                           = &Error{0x00000A4F, "NERR_RplVendorNameUnavailable", "Vendor name is in use by some other vendor record."}
	NerrRplbootnameunavailable                             = &Error{0x00000A50, "NERR_RplBootNameUnavailable", "The boot name or vendor ID is in use by some other boot block record."}
	NerrRplconfignameunavailable                           = &Error{0x00000A51, "NERR_RplConfigNameUnavailable", "The configuration name is in use by some other configuration."}
	NerrDfsinternalcorruption                              = &Error{0x00000A64, "NERR_DfsInternalCorruption", "The internal database maintained by the DFS service is corrupt."}
	NerrDfsvolumedatacorrupt                               = &Error{0x00000A65, "NERR_DfsVolumeDataCorrupt", "One of the records in the internal DFS database is corrupt."}
	NerrDfsnosuchvolume                                    = &Error{0x00000A66, "NERR_DfsNoSuchVolume", "There is no DFS name whose entry path matches the input entry path."}
	NerrDfsvolumealreadyexists                             = &Error{0x00000A67, "NERR_DfsVolumeAlreadyExists", "A root or link with the given name already exists."}
	NerrDfsalreadyshared                                   = &Error{0x00000A68, "NERR_DfsAlreadyShared", "The server share specified is already shared in the DFS."}
	NerrDfsnosuchshare                                     = &Error{0x00000A69, "NERR_DfsNoSuchShare", "The indicated server share does not support the indicated DFS namespace."}
	NerrDfsnotaleafvolume                                  = &Error{0x00000A6A, "NERR_DfsNotALeafVolume", "The operation is not valid in this portion of the namespace."}
	NerrDfsleafvolume                                      = &Error{0x00000A6B, "NERR_DfsLeafVolume", "The operation is not valid in this portion of the namespace."}
	NerrDfsvolumehasmultipleservers                        = &Error{0x00000A6C, "NERR_DfsVolumeHasMultipleServers", "The operation is ambiguous because the link has multiple servers."}
	NerrDfscantcreatejunctionpoint                         = &Error{0x00000A6D, "NERR_DfsCantCreateJunctionPoint", "Unable to create a link."}
	NerrDfsservernotdfsaware                               = &Error{0x00000A6E, "NERR_DfsServerNotDfsAware", "The server is not DFS-aware."}
	NerrDfsbadrenamepath                                   = &Error{0x00000A6F, "NERR_DfsBadRenamePath", "The specified rename target path is invalid."}
	NerrDfsvolumeisoffline                                 = &Error{0x00000A70, "NERR_DfsVolumeIsOffline", "The specified DFS link is offline."}
	NerrDfsnosuchserver                                    = &Error{0x00000A71, "NERR_DfsNoSuchServer", "The specified server is not a server for this link."}
	NerrDfscyclicalname                                    = &Error{0x00000A72, "NERR_DfsCyclicalName", "A cycle in the DFS name was detected."}
	NerrDfsnotsupportedinserverdfs                         = &Error{0x00000A73, "NERR_DfsNotSupportedInServerDfs", "The operation is not supported on a server-based DFS."}
	NerrDfsduplicateservice                                = &Error{0x00000A74, "NERR_DfsDuplicateService", "This link is already supported by the specified server share."}
	NerrDfscantremovelastservershare                       = &Error{0x00000A75, "NERR_DfsCantRemoveLastServerShare", "Cannot remove the last server share supporting this root or link."}
	NerrDfsvolumeisinterdfs                                = &Error{0x00000A76, "NERR_DfsVolumeIsInterDfs", "The operation is not supported for an inter-DFS link."}
	NerrDfsinconsistent                                    = &Error{0x00000A77, "NERR_DfsInconsistent", "The internal state of the DFS Service has become inconsistent."}
	NerrDfsserverupgraded                                  = &Error{0x00000A78, "NERR_DfsServerUpgraded", "The DFS Service has been installed on the specified server."}
	NerrDfsdataisidentical                                 = &Error{0x00000A79, "NERR_DfsDataIsIdentical", "The DFS data being reconciled is identical."}
	NerrDfscantremovedfsroot                               = &Error{0x00000A7A, "NERR_DfsCantRemoveDfsRoot", "The DFS root cannot be deleted. Uninstall DFS if required."}
	NerrDfschildorparentindfs                              = &Error{0x00000A7B, "NERR_DfsChildOrParentInDfs", "A child or parent directory of the share is already in a DFS."}
	NerrDfsinternalerror                                   = &Error{0x00000A82, "NERR_DfsInternalError", "DFS internal error."}
	NerrSetupalreadyjoined                                 = &Error{0x00000A83, "NERR_SetupAlreadyJoined", "This machine is already joined to a domain."}
	NerrSetupnotjoined                                     = &Error{0x00000A84, "NERR_SetupNotJoined", "This machine is not currently joined to a domain."}
	NerrSetupdomaincontroller                              = &Error{0x00000A85, "NERR_SetupDomainController", "This machine is a domain controller and cannot be unjoined from a domain."}
	NerrDefaultjoinrequired                                = &Error{0x00000A86, "NERR_DefaultJoinRequired", "The destination domain controller does not support creating machine accounts in organizational units (OUs)."}
	NerrInvalidworkgroupname                               = &Error{0x00000A87, "NERR_InvalidWorkgroupName", "The specified workgroup name is invalid."}
	NerrNameusesincompatiblecodepage                       = &Error{0x00000A88, "NERR_NameUsesIncompatibleCodePage", "The specified computer name is incompatible with the default language used on the domain controller."}
	NerrComputeraccountnotfound                            = &Error{0x00000A89, "NERR_ComputerAccountNotFound", "The specified computer account could not be found."}
	NerrPersonalsku                                        = &Error{0x00000A8A, "NERR_PersonalSku", "This version of Windows cannot be joined to a domain."}
	NerrPasswordmustchange                                 = &Error{0x00000A8D, "NERR_PasswordMustChange", "The password must change at the next logon."}
	NerrAccountlockedout                                   = &Error{0x00000A8E, "NERR_AccountLockedOut", "The account is locked out."}
	NerrPasswordtoolong                                    = &Error{0x00000A8F, "NERR_PasswordTooLong", "The password is too long."}
	NerrPasswordnotcomplexenough                           = &Error{0x00000A90, "NERR_PasswordNotComplexEnough", "The password does not meet the complexity policy."}
	NerrPasswordfiltererror                                = &Error{0x00000A91, "NERR_PasswordFilterError", "The password does not meet the requirements of the password filter DLLs."}
	ErrorUnknownPrintMonitor                               = &Error{0x00000BB8, "ERROR_UNKNOWN_PRINT_MONITOR", "The specified print monitor is unknown."}
	ErrorPrinterDriverInUse                                = &Error{0x00000BB9, "ERROR_PRINTER_DRIVER_IN_USE", "The specified printer driver is currently in use."}
	ErrorSpoolFileNotFound                                 = &Error{0x00000BBA, "ERROR_SPOOL_FILE_NOT_FOUND", "The spool file was not found."}
	ErrorSplNoStartdoc                                     = &Error{0x00000BBB, "ERROR_SPL_NO_STARTDOC", "A StartDocPrinter call was not issued."}
	ErrorSplNoAddjob                                       = &Error{0x00000BBC, "ERROR_SPL_NO_ADDJOB", "An AddJob call was not issued."}
	ErrorPrintProcessorAlreadyInstalled                    = &Error{0x00000BBD, "ERROR_PRINT_PROCESSOR_ALREADY_INSTALLED", "The specified print processor has already been installed."}
	ErrorPrintMonitorAlreadyInstalled                      = &Error{0x00000BBE, "ERROR_PRINT_MONITOR_ALREADY_INSTALLED", "The specified print monitor has already been installed."}
	ErrorInvalidPrintMonitor                               = &Error{0x00000BBF, "ERROR_INVALID_PRINT_MONITOR", "The specified print monitor does not have the required functions."}
	ErrorPrintMonitorInUse                                 = &Error{0x00000BC0, "ERROR_PRINT_MONITOR_IN_USE", "The specified print monitor is currently in use."}
	ErrorPrinterHasJobsQueued                              = &Error{0x00000BC1, "ERROR_PRINTER_HAS_JOBS_QUEUED", "The requested operation is not allowed when there are jobs queued to the printer."}
	ErrorSuccessRebootRequired                             = &Error{0x00000BC2, "ERROR_SUCCESS_REBOOT_REQUIRED", "The requested operation is successful. Changes will not be effective until the system is rebooted."}
	ErrorSuccessRestartRequired                            = &Error{0x00000BC3, "ERROR_SUCCESS_RESTART_REQUIRED", "The requested operation is successful. Changes will not be effective until the service is restarted."}
	ErrorPrinterNotFound                                   = &Error{0x00000BC4, "ERROR_PRINTER_NOT_FOUND", "No printers were found."}
	ErrorPrinterDriverWarned                               = &Error{0x00000BC5, "ERROR_PRINTER_DRIVER_WARNED", "The printer driver is known to be unreliable."}
	ErrorPrinterDriverBlocked                              = &Error{0x00000BC6, "ERROR_PRINTER_DRIVER_BLOCKED", "The printer driver is known to harm the system."}
	ErrorPrinterDriverPackageInUse                         = &Error{0x00000BC7, "ERROR_PRINTER_DRIVER_PACKAGE_IN_USE", "The specified printer driver package is currently in use."}
	ErrorCoreDriverPackageNotFound                         = &Error{0x00000BC8, "ERROR_CORE_DRIVER_PACKAGE_NOT_FOUND", "Unable to find a core driver package that is required by the printer driver package."}
	ErrorFailRebootRequired                                = &Error{0x00000BC9, "ERROR_FAIL_REBOOT_REQUIRED", "The requested operation failed. A system reboot is required to roll back changes made."}
	ErrorFailRebootInitiated                               = &Error{0x00000BCA, "ERROR_FAIL_REBOOT_INITIATED", "The requested operation failed. A system reboot has been initiated to roll back changes made."}
	ErrorPrinterDriverDownloadNeeded                       = &Error{0x00000BCB, "ERROR_PRINTER_DRIVER_DOWNLOAD_NEEDED", "The specified printer driver was not found on the system and needs to be downloaded."}
	ErrorPrinterNotShareable                               = &Error{0x00000BCE, "ERROR_PRINTER_NOT_SHAREABLE", "The specified printer cannot be shared."}
	ErrorIoReissueAsCached                                 = &Error{0x00000F6E, "ERROR_IO_REISSUE_AS_CACHED", "Reissue the given operation as a cached I/O operation."}
	ErrorWinsInternal                                      = &Error{0x00000FA0, "ERROR_WINS_INTERNAL", "Windows Internet Name Service (WINS) encountered an error while processing the command."}
	ErrorCanNotDelLocalWins                                = &Error{0x00000FA1, "ERROR_CAN_NOT_DEL_LOCAL_WINS", "The local WINS cannot be deleted."}
	ErrorStaticInit                                        = &Error{0x00000FA2, "ERROR_STATIC_INIT", "The importation from the file failed."}
	ErrorIncBackup                                         = &Error{0x00000FA3, "ERROR_INC_BACKUP", "The backup failed. Was a full backup done before?"}
	ErrorFullBackup                                        = &Error{0x00000FA4, "ERROR_FULL_BACKUP", "The backup failed. Check the directory to which you are backing the database."}
	ErrorRecNonExistent                                    = &Error{0x00000FA5, "ERROR_REC_NON_EXISTENT", "The name does not exist in the WINS database."}
	ErrorRplNotAllowed                                     = &Error{0x00000FA6, "ERROR_RPL_NOT_ALLOWED", "Replication with a nonconfigured partner is not allowed."}
	PeerdistErrorContentinfoVersionUnsupported             = &Error{0x00000FD2, "PEERDIST_ERROR_CONTENTINFO_VERSION_UNSUPPORTED", "The version of the supplied content information is not supported."}
	PeerdistErrorCannotParseContentinfo                    = &Error{0x00000FD3, "PEERDIST_ERROR_CANNOT_PARSE_CONTENTINFO", "The supplied content information is malformed."}
	PeerdistErrorMissingData                               = &Error{0x00000FD4, "PEERDIST_ERROR_MISSING_DATA", "The requested data cannot be found in local or peer caches."}
	PeerdistErrorNoMore                                    = &Error{0x00000FD5, "PEERDIST_ERROR_NO_MORE", "No more data is available or required."}
	PeerdistErrorNotInitialized                            = &Error{0x00000FD6, "PEERDIST_ERROR_NOT_INITIALIZED", "The supplied object has not been initialized."}
	PeerdistErrorAlreadyInitialized                        = &Error{0x00000FD7, "PEERDIST_ERROR_ALREADY_INITIALIZED", "The supplied object has already been initialized."}
	PeerdistErrorShutdownInProgress                        = &Error{0x00000FD8, "PEERDIST_ERROR_SHUTDOWN_IN_PROGRESS", "A shutdown operation is already in progress."}
	PeerdistErrorInvalidated                               = &Error{0x00000FD9, "PEERDIST_ERROR_INVALIDATED", "The supplied object has already been invalidated."}
	PeerdistErrorAlreadyExists                             = &Error{0x00000FDA, "PEERDIST_ERROR_ALREADY_EXISTS", "An element already exists and was not replaced."}
	PeerdistErrorOperationNotfound                         = &Error{0x00000FDB, "PEERDIST_ERROR_OPERATION_NOTFOUND", "Cannot cancel the requested operation as it has already been completed."}
	PeerdistErrorAlreadyCompleted                          = &Error{0x00000FDC, "PEERDIST_ERROR_ALREADY_COMPLETED", "Cannot perform the requested operation because it has already been carried out."}
	PeerdistErrorOutOfBounds                               = &Error{0x00000FDD, "PEERDIST_ERROR_OUT_OF_BOUNDS", "An operation accessed data beyond the bounds of valid data."}
	PeerdistErrorVersionUnsupported                        = &Error{0x00000FDE, "PEERDIST_ERROR_VERSION_UNSUPPORTED", "The requested version is not supported."}
	PeerdistErrorInvalidConfiguration                      = &Error{0x00000FDF, "PEERDIST_ERROR_INVALID_CONFIGURATION", "A configuration value is invalid."}
	PeerdistErrorNotLicensed                               = &Error{0x00000FE0, "PEERDIST_ERROR_NOT_LICENSED", "The SKU is not licensed."}
	PeerdistErrorServiceUnavailable                        = &Error{0x00000FE1, "PEERDIST_ERROR_SERVICE_UNAVAILABLE", "PeerDist Service is still initializing and will be available shortly."}
	ErrorDhcpAddressConflict                               = &Error{0x00001004, "ERROR_DHCP_ADDRESS_CONFLICT", "The Dynamic Host Configuration Protocol (DHCP) client has obtained an IP address that is already in use on the network. The local interface will be disabled until the DHCP client can obtain a new address."}
	ErrorWmiGuidNotFound                                   = &Error{0x00001068, "ERROR_WMI_GUID_NOT_FOUND", "The GUID passed was not recognized as valid by a WMI data provider."}
	ErrorWmiInstanceNotFound                               = &Error{0x00001069, "ERROR_WMI_INSTANCE_NOT_FOUND", "The instance name passed was not recognized as valid by a WMI data provider."}
	ErrorWmiItemidNotFound                                 = &Error{0x0000106A, "ERROR_WMI_ITEMID_NOT_FOUND", "The data item ID passed was not recognized as valid by a WMI data provider."}
	ErrorWmiTryAgain                                       = &Error{0x0000106B, "ERROR_WMI_TRY_AGAIN", "The WMI request could not be completed and should be retried."}
	ErrorWmiDpNotFound                                     = &Error{0x0000106C, "ERROR_WMI_DP_NOT_FOUND", "The WMI data provider could not be located."}
	ErrorWmiUnresolvedInstanceRef                          = &Error{0x0000106D, "ERROR_WMI_UNRESOLVED_INSTANCE_REF", "The WMI data provider references an instance set that has not been registered."}
	ErrorWmiAlreadyEnabled                                 = &Error{0x0000106E, "ERROR_WMI_ALREADY_ENABLED", "The WMI data block or event notification has already been enabled."}
	ErrorWmiGuidDisconnected                               = &Error{0x0000106F, "ERROR_WMI_GUID_DISCONNECTED", "The WMI data block is no longer available."}
	ErrorWmiServerUnavailable                              = &Error{0x00001070, "ERROR_WMI_SERVER_UNAVAILABLE", "The WMI data service is not available."}
	ErrorWmiDpFailed                                       = &Error{0x00001071, "ERROR_WMI_DP_FAILED", "The WMI data provider failed to carry out the request."}
	ErrorWmiInvalidMof                                     = &Error{0x00001072, "ERROR_WMI_INVALID_MOF", "The WMI Managed Object Format (MOF) information is not valid."}
	ErrorWmiInvalidReginfo                                 = &Error{0x00001073, "ERROR_WMI_INVALID_REGINFO", "The WMI registration information is not valid."}
	ErrorWmiAlreadyDisabled                                = &Error{0x00001074, "ERROR_WMI_ALREADY_DISABLED", "The WMI data block or event notification has already been disabled."}
	ErrorWmiReadOnly                                       = &Error{0x00001075, "ERROR_WMI_READ_ONLY", "The WMI data item or data block is read-only."}
	ErrorWmiSetFailure                                     = &Error{0x00001076, "ERROR_WMI_SET_FAILURE", "The WMI data item or data block could not be changed."}
	ErrorInvalidMedia                                      = &Error{0x000010CC, "ERROR_INVALID_MEDIA", "The media identifier does not represent a valid medium."}
	ErrorInvalidLibrary                                    = &Error{0x000010CD, "ERROR_INVALID_LIBRARY", "The library identifier does not represent a valid library."}
	ErrorInvalidMediaPool                                  = &Error{0x000010CE, "ERROR_INVALID_MEDIA_POOL", "The media pool identifier does not represent a valid media pool."}
	ErrorDriveMediaMismatch                                = &Error{0x000010CF, "ERROR_DRIVE_MEDIA_MISMATCH", "The drive and medium are not compatible, or they exist in different libraries."}
	ErrorMediaOffline                                      = &Error{0x000010D0, "ERROR_MEDIA_OFFLINE", "The medium currently exists in an offline library and must be online to perform this operation."}
	ErrorLibraryOffline                                    = &Error{0x000010D1, "ERROR_LIBRARY_OFFLINE", "The operation cannot be performed on an offline library."}
	ErrorEmpty                                             = &Error{0x000010D2, "ERROR_EMPTY", "The library, drive, or media pool is empty."}
	ErrorNotEmpty                                          = &Error{0x000010D3, "ERROR_NOT_EMPTY", "The library, drive, or media pool must be empty to perform this operation."}
	ErrorMediaUnavailable                                  = &Error{0x000010D4, "ERROR_MEDIA_UNAVAILABLE", "No media is currently available in this media pool or library."}
	ErrorResourceDisabled                                  = &Error{0x000010D5, "ERROR_RESOURCE_DISABLED", "A resource required for this operation is disabled."}
	ErrorInvalidCleaner                                    = &Error{0x000010D6, "ERROR_INVALID_CLEANER", "The media identifier does not represent a valid cleaner."}
	ErrorUnableToClean                                     = &Error{0x000010D7, "ERROR_UNABLE_TO_CLEAN", "The drive cannot be cleaned or does not support cleaning."}
	ErrorObjectNotFound                                    = &Error{0x000010D8, "ERROR_OBJECT_NOT_FOUND", "The object identifier does not represent a valid object."}
	ErrorDatabaseFailure                                   = &Error{0x000010D9, "ERROR_DATABASE_FAILURE", "Unable to read from or write to the database."}
	ErrorDatabaseFull                                      = &Error{0x000010DA, "ERROR_DATABASE_FULL", "The database is full."}
	ErrorMediaIncompatible                                 = &Error{0x000010DB, "ERROR_MEDIA_INCOMPATIBLE", "The medium is not compatible with the device or media pool."}
	ErrorResourceNotPresent                                = &Error{0x000010DC, "ERROR_RESOURCE_NOT_PRESENT", "The resource required for this operation does not exist."}
	ErrorInvalidOperation                                  = &Error{0x000010DD, "ERROR_INVALID_OPERATION", "The operation identifier is not valid."}
	ErrorMediaNotAvailable                                 = &Error{0x000010DE, "ERROR_MEDIA_NOT_AVAILABLE", "The media is not mounted or ready for use."}
	ErrorDeviceNotAvailable                                = &Error{0x000010DF, "ERROR_DEVICE_NOT_AVAILABLE", "The device is not ready for use."}
	ErrorRequestRefused                                    = &Error{0x000010E0, "ERROR_REQUEST_REFUSED", "The operator or administrator has refused the request."}
	ErrorInvalidDriveObject                                = &Error{0x000010E1, "ERROR_INVALID_DRIVE_OBJECT", "The drive identifier does not represent a valid drive."}
	ErrorLibraryFull                                       = &Error{0x000010E2, "ERROR_LIBRARY_FULL", "Library is full. No slot is available for use."}
	ErrorMediumNotAccessible                               = &Error{0x000010E3, "ERROR_MEDIUM_NOT_ACCESSIBLE", "The transport cannot access the medium."}
	ErrorUnableToLoadMedium                                = &Error{0x000010E4, "ERROR_UNABLE_TO_LOAD_MEDIUM", "Unable to load the medium into the drive."}
	ErrorUnableToInventoryDrive                            = &Error{0x000010E5, "ERROR_UNABLE_TO_INVENTORY_DRIVE", "Unable to retrieve the drive status."}
	ErrorUnableToInventorySlot                             = &Error{0x000010E6, "ERROR_UNABLE_TO_INVENTORY_SLOT", "Unable to retrieve the slot status."}
	ErrorUnableToInventoryTransport                        = &Error{0x000010E7, "ERROR_UNABLE_TO_INVENTORY_TRANSPORT", "Unable to retrieve status about the transport."}
	ErrorTransportFull                                     = &Error{0x000010E8, "ERROR_TRANSPORT_FULL", "Cannot use the transport because it is already in use."}
	ErrorControllingIeport                                 = &Error{0x000010E9, "ERROR_CONTROLLING_IEPORT", "Unable to open or close the inject/eject port."}
	ErrorUnableToEjectMountedMedia                         = &Error{0x000010EA, "ERROR_UNABLE_TO_EJECT_MOUNTED_MEDIA", "Unable to eject the medium because it is in a drive."}
	ErrorCleanerSlotSet                                    = &Error{0x000010EB, "ERROR_CLEANER_SLOT_SET", "A cleaner slot is already reserved."}
	ErrorCleanerSlotNotSet                                 = &Error{0x000010EC, "ERROR_CLEANER_SLOT_NOT_SET", "A cleaner slot is not reserved."}
	ErrorCleanerCartridgeSpent                             = &Error{0x000010ED, "ERROR_CLEANER_CARTRIDGE_SPENT", "The cleaner cartridge has performed the maximum number of drive cleanings."}
	ErrorUnexpectedOmid                                    = &Error{0x000010EE, "ERROR_UNEXPECTED_OMID", "Unexpected on-medium identifier."}
	ErrorCantDeleteLastItem                                = &Error{0x000010EF, "ERROR_CANT_DELETE_LAST_ITEM", "The last remaining item in this group or resource cannot be deleted."}
	ErrorMessageExceedsMaxSize                             = &Error{0x000010F0, "ERROR_MESSAGE_EXCEEDS_MAX_SIZE", "The message provided exceeds the maximum size allowed for this parameter."}
	ErrorVolumeContainsSysFiles                            = &Error{0x000010F1, "ERROR_VOLUME_CONTAINS_SYS_FILES", "The volume contains system or paging files."}
	ErrorIndigenousType                                    = &Error{0x000010F2, "ERROR_INDIGENOUS_TYPE", "The media type cannot be removed from this library because at least one drive in the library reports it can support this media type."}
	ErrorNoSupportingDrives                                = &Error{0x000010F3, "ERROR_NO_SUPPORTING_DRIVES", "This offline media cannot be mounted on this system because no enabled drives are present that can be used."}
	ErrorCleanerCartridgeInstalled                         = &Error{0x000010F4, "ERROR_CLEANER_CARTRIDGE_INSTALLED", "A cleaner cartridge is present in the tape library."}
	ErrorIeportFull                                        = &Error{0x000010F5, "ERROR_IEPORT_FULL", "Cannot use the IEport because it is not empty."}
	ErrorFileOffline                                       = &Error{0x000010FE, "ERROR_FILE_OFFLINE", "The remote storage service was not able to recall the file."}
	ErrorRemoteStorageNotActive                            = &Error{0x000010FF, "ERROR_REMOTE_STORAGE_NOT_ACTIVE", "The remote storage service is not operational at this time."}
	ErrorRemoteStorageMediaError                           = &Error{0x00001100, "ERROR_REMOTE_STORAGE_MEDIA_ERROR", "The remote storage service encountered a media error."}
	ErrorNotAReparsePoint                                  = &Error{0x00001126, "ERROR_NOT_A_REPARSE_POINT", "The file or directory is not a reparse point."}
	ErrorReparseAttributeConflict                          = &Error{0x00001127, "ERROR_REPARSE_ATTRIBUTE_CONFLICT", "The reparse point attribute cannot be set because it conflicts with an existing attribute."}
	ErrorInvalidReparseData                                = &Error{0x00001128, "ERROR_INVALID_REPARSE_DATA", "The data present in the reparse point buffer is invalid."}
	ErrorReparseTagInvalid                                 = &Error{0x00001129, "ERROR_REPARSE_TAG_INVALID", "The tag present in the reparse point buffer is invalid."}
	ErrorReparseTagMismatch                                = &Error{0x0000112A, "ERROR_REPARSE_TAG_MISMATCH", "There is a mismatch between the tag specified in the request and the tag present in the reparse point."}
	ErrorVolumeNotSisEnabled                               = &Error{0x00001194, "ERROR_VOLUME_NOT_SIS_ENABLED", "Single Instance Storage (SIS) is not available on this volume."}
	ErrorDependentResourceExists                           = &Error{0x00001389, "ERROR_DEPENDENT_RESOURCE_EXISTS", "The operation cannot be completed because other resources depend on this resource."}
	ErrorDependencyNotFound                                = &Error{0x0000138A, "ERROR_DEPENDENCY_NOT_FOUND", "The cluster resource dependency cannot be found."}
	ErrorDependencyAlreadyExists                           = &Error{0x0000138B, "ERROR_DEPENDENCY_ALREADY_EXISTS", "The cluster resource cannot be made dependent on the specified resource because it is already dependent."}
	ErrorResourceNotOnline                                 = &Error{0x0000138C, "ERROR_RESOURCE_NOT_ONLINE", "The cluster resource is not online."}
	ErrorHostNodeNotAvailable                              = &Error{0x0000138D, "ERROR_HOST_NODE_NOT_AVAILABLE", "A cluster node is not available for this operation."}
	ErrorResourceNotAvailable                              = &Error{0x0000138E, "ERROR_RESOURCE_NOT_AVAILABLE", "The cluster resource is not available."}
	ErrorResourceNotFound                                  = &Error{0x0000138F, "ERROR_RESOURCE_NOT_FOUND", "The cluster resource could not be found."}
	ErrorShutdownCluster                                   = &Error{0x00001390, "ERROR_SHUTDOWN_CLUSTER", "The cluster is being shut down."}
	ErrorCantEvictActiveNode                               = &Error{0x00001391, "ERROR_CANT_EVICT_ACTIVE_NODE", "A cluster node cannot be evicted from the cluster unless the node is down or it is the last node."}
	ErrorObjectAlreadyExists                               = &Error{0x00001392, "ERROR_OBJECT_ALREADY_EXISTS", "The object already exists."}
	ErrorObjectInList                                      = &Error{0x00001393, "ERROR_OBJECT_IN_LIST", "The object is already in the list."}
	ErrorGroupNotAvailable                                 = &Error{0x00001394, "ERROR_GROUP_NOT_AVAILABLE", "The cluster group is not available for any new requests."}
	ErrorGroupNotFound                                     = &Error{0x00001395, "ERROR_GROUP_NOT_FOUND", "The cluster group could not be found."}
	ErrorGroupNotOnline                                    = &Error{0x00001396, "ERROR_GROUP_NOT_ONLINE", "The operation could not be completed because the cluster group is not online."}
	ErrorHostNodeNotResourceOwner                          = &Error{0x00001397, "ERROR_HOST_NODE_NOT_RESOURCE_OWNER", "The operation failed because either the specified cluster node is not the owner of the resource, or the node is not a possible owner of the resource."}
	ErrorHostNodeNotGroupOwner                             = &Error{0x00001398, "ERROR_HOST_NODE_NOT_GROUP_OWNER", "The operation failed because either the specified cluster node is not the owner of the group, or the node is not a possible owner of the group."}
	ErrorResmonCreateFailed                                = &Error{0x00001399, "ERROR_RESMON_CREATE_FAILED", "The cluster resource could not be created in the specified resource monitor."}
	ErrorResmonOnlineFailed                                = &Error{0x0000139A, "ERROR_RESMON_ONLINE_FAILED", "The cluster resource could not be brought online by the resource monitor."}
	ErrorResourceOnline                                    = &Error{0x0000139B, "ERROR_RESOURCE_ONLINE", "The operation could not be completed because the cluster resource is online."}
	ErrorQuorumResource                                    = &Error{0x0000139C, "ERROR_QUORUM_RESOURCE", "The cluster resource could not be deleted or brought offline because it is the quorum resource."}
	ErrorNotQuorumCapable                                  = &Error{0x0000139D, "ERROR_NOT_QUORUM_CAPABLE", "The cluster could not make the specified resource a quorum resource because it is not capable of being a quorum resource."}
	ErrorClusterShuttingDown                               = &Error{0x0000139E, "ERROR_CLUSTER_SHUTTING_DOWN", "The cluster software is shutting down."}
	ErrorInvalidState                                      = &Error{0x0000139F, "ERROR_INVALID_STATE", "The group or resource is not in the correct state to perform the requested operation."}
	ErrorResourcePropertiesStored                          = &Error{0x000013A0, "ERROR_RESOURCE_PROPERTIES_STORED", "The properties were stored but not all changes will take effect until the next time the resource is brought online."}
	ErrorNotQuorumClass                                    = &Error{0x000013A1, "ERROR_NOT_QUORUM_CLASS", "The cluster could not make the specified resource a quorum resource because it does not belong to a shared storage class."}
	ErrorCoreResource                                      = &Error{0x000013A2, "ERROR_CORE_RESOURCE", "The cluster resource could not be deleted because it is a core resource."}
	ErrorQuorumResourceOnlineFailed                        = &Error{0x000013A3, "ERROR_QUORUM_RESOURCE_ONLINE_FAILED", "The quorum resource failed to come online."}
	ErrorQuorumlogOpenFailed                               = &Error{0x000013A4, "ERROR_QUORUMLOG_OPEN_FAILED", "The quorum log could not be created or mounted successfully."}
	ErrorClusterlogCorrupt                                 = &Error{0x000013A5, "ERROR_CLUSTERLOG_CORRUPT", "The cluster log is corrupt."}
	ErrorClusterlogRecordExceedsMaxsize                    = &Error{0x000013A6, "ERROR_CLUSTERLOG_RECORD_EXCEEDS_MAXSIZE", "The record could not be written to the cluster log because it exceeds the maximum size."}
	ErrorClusterlogExceedsMaxsize                          = &Error{0x000013A7, "ERROR_CLUSTERLOG_EXCEEDS_MAXSIZE", "The cluster log exceeds its maximum size."}
	ErrorClusterlogChkpointNotFound                        = &Error{0x000013A8, "ERROR_CLUSTERLOG_CHKPOINT_NOT_FOUND", "No checkpoint record was found in the cluster log."}
	ErrorClusterlogNotEnoughSpace                          = &Error{0x000013A9, "ERROR_CLUSTERLOG_NOT_ENOUGH_SPACE", "The minimum required disk space needed for logging is not available."}
	ErrorQuorumOwnerAlive                                  = &Error{0x000013AA, "ERROR_QUORUM_OWNER_ALIVE", "The cluster node failed to take control of the quorum resource because the resource is owned by another active node."}
	ErrorNetworkNotAvailable                               = &Error{0x000013AB, "ERROR_NETWORK_NOT_AVAILABLE", "A cluster network is not available for this operation."}
	ErrorNodeNotAvailable                                  = &Error{0x000013AC, "ERROR_NODE_NOT_AVAILABLE", "A cluster node is not available for this operation."}
	ErrorAllNodesNotAvailable                              = &Error{0x000013AD, "ERROR_ALL_NODES_NOT_AVAILABLE", "All cluster nodes must be running to perform this operation."}
	ErrorResourceFailed                                    = &Error{0x000013AE, "ERROR_RESOURCE_FAILED", "A cluster resource failed."}
	ErrorClusterInvalidNode                                = &Error{0x000013AF, "ERROR_CLUSTER_INVALID_NODE", "The cluster node is not valid."}
	ErrorClusterNodeExists                                 = &Error{0x000013B0, "ERROR_CLUSTER_NODE_EXISTS", "The cluster node already exists."}
	ErrorClusterJoinInProgress                             = &Error{0x000013B1, "ERROR_CLUSTER_JOIN_IN_PROGRESS", "A node is in the process of joining the cluster."}
	ErrorClusterNodeNotFound                               = &Error{0x000013B2, "ERROR_CLUSTER_NODE_NOT_FOUND", "The cluster node was not found."}
	ErrorClusterLocalNodeNotFound                          = &Error{0x000013B3, "ERROR_CLUSTER_LOCAL_NODE_NOT_FOUND", "The cluster local node information was not found."}
	ErrorClusterNetworkExists                              = &Error{0x000013B4, "ERROR_CLUSTER_NETWORK_EXISTS", "The cluster network already exists."}
	ErrorClusterNetworkNotFound                            = &Error{0x000013B5, "ERROR_CLUSTER_NETWORK_NOT_FOUND", "The cluster network was not found."}
	ErrorClusterNetinterfaceExists                         = &Error{0x000013B6, "ERROR_CLUSTER_NETINTERFACE_EXISTS", "The cluster network interface already exists."}
	ErrorClusterNetinterfaceNotFound                       = &Error{0x000013B7, "ERROR_CLUSTER_NETINTERFACE_NOT_FOUND", "The cluster network interface was not found."}
	ErrorClusterInvalidRequest                             = &Error{0x000013B8, "ERROR_CLUSTER_INVALID_REQUEST", "The cluster request is not valid for this object."}
	ErrorClusterInvalidNetworkProvider                     = &Error{0x000013B9, "ERROR_CLUSTER_INVALID_NETWORK_PROVIDER", "The cluster network provider is not valid."}
	ErrorClusterNodeDown                                   = &Error{0x000013BA, "ERROR_CLUSTER_NODE_DOWN", "The cluster node is down."}
	ErrorClusterNodeUnreachable                            = &Error{0x000013BB, "ERROR_CLUSTER_NODE_UNREACHABLE", "The cluster node is not reachable."}
	ErrorClusterNodeNotMember                              = &Error{0x000013BC, "ERROR_CLUSTER_NODE_NOT_MEMBER", "The cluster node is not a member of the cluster."}
	ErrorClusterJoinNotInProgress                          = &Error{0x000013BD, "ERROR_CLUSTER_JOIN_NOT_IN_PROGRESS", "A cluster join operation is not in progress."}
	ErrorClusterInvalidNetwork                             = &Error{0x000013BE, "ERROR_CLUSTER_INVALID_NETWORK", "The cluster network is not valid."}
	ErrorClusterNodeUp                                     = &Error{0x000013C0, "ERROR_CLUSTER_NODE_UP", "The cluster node is up."}
	ErrorClusterIpaddrInUse                                = &Error{0x000013C1, "ERROR_CLUSTER_IPADDR_IN_USE", "The cluster IP address is already in use."}
	ErrorClusterNodeNotPaused                              = &Error{0x000013C2, "ERROR_CLUSTER_NODE_NOT_PAUSED", "The cluster node is not paused."}
	ErrorClusterNoSecurityContext                          = &Error{0x000013C3, "ERROR_CLUSTER_NO_SECURITY_CONTEXT", "No cluster security context is available."}
	ErrorClusterNetworkNotInternal                         = &Error{0x000013C4, "ERROR_CLUSTER_NETWORK_NOT_INTERNAL", "The cluster network is not configured for internal cluster communication."}
	ErrorClusterNodeAlreadyUp                              = &Error{0x000013C5, "ERROR_CLUSTER_NODE_ALREADY_UP", "The cluster node is already up."}
	ErrorClusterNodeAlreadyDown                            = &Error{0x000013C6, "ERROR_CLUSTER_NODE_ALREADY_DOWN", "The cluster node is already down."}
	ErrorClusterNetworkAlreadyOnline                       = &Error{0x000013C7, "ERROR_CLUSTER_NETWORK_ALREADY_ONLINE", "The cluster network is already online."}
	ErrorClusterNetworkAlreadyOffline                      = &Error{0x000013C8, "ERROR_CLUSTER_NETWORK_ALREADY_OFFLINE", "The cluster network is already offline."}
	ErrorClusterNodeAlreadyMember                          = &Error{0x000013C9, "ERROR_CLUSTER_NODE_ALREADY_MEMBER", "The cluster node is already a member of the cluster."}
	ErrorClusterLastInternalNetwork                        = &Error{0x000013CA, "ERROR_CLUSTER_LAST_INTERNAL_NETWORK", "The cluster network is the only one configured for internal cluster communication between two or more active cluster nodes. The internal communication capability cannot be removed from the network."}
	ErrorClusterNetworkHasDependents                       = &Error{0x000013CB, "ERROR_CLUSTER_NETWORK_HAS_DEPENDENTS", "One or more cluster resources depend on the network to provide service to clients. The client access capability cannot be removed from the network."}
	ErrorInvalidOperationOnQuorum                          = &Error{0x000013CC, "ERROR_INVALID_OPERATION_ON_QUORUM", "This operation cannot be performed on the cluster resource because it is the quorum resource. This quorum resource cannot be brought offline and its possible owners list cannot be modified."}
	ErrorDependencyNotAllowed                              = &Error{0x000013CD, "ERROR_DEPENDENCY_NOT_ALLOWED", "The cluster quorum resource is not allowed to have any dependencies."}
	ErrorClusterNodePaused                                 = &Error{0x000013CE, "ERROR_CLUSTER_NODE_PAUSED", "The cluster node is paused."}
	ErrorNodeCantHostResource                              = &Error{0x000013CF, "ERROR_NODE_CANT_HOST_RESOURCE", "The cluster resource cannot be brought online. The owner node cannot run this resource."}
	ErrorClusterNodeNotReady                               = &Error{0x000013D0, "ERROR_CLUSTER_NODE_NOT_READY", "The cluster node is not ready to perform the requested operation."}
	ErrorClusterNodeShuttingDown                           = &Error{0x000013D1, "ERROR_CLUSTER_NODE_SHUTTING_DOWN", "The cluster node is shutting down."}
	ErrorClusterJoinAborted                                = &Error{0x000013D2, "ERROR_CLUSTER_JOIN_ABORTED", "The cluster join operation was aborted."}
	ErrorClusterIncompatibleVersions                       = &Error{0x000013D3, "ERROR_CLUSTER_INCOMPATIBLE_VERSIONS", "The cluster join operation failed due to incompatible software versions between the joining node and its sponsor."}
	ErrorClusterMaxnumOfResourcesExceeded                  = &Error{0x000013D4, "ERROR_CLUSTER_MAXNUM_OF_RESOURCES_EXCEEDED", "This resource cannot be created because the cluster has reached the limit on the number of resources it can monitor."}
	ErrorClusterSystemConfigChanged                        = &Error{0x000013D5, "ERROR_CLUSTER_SYSTEM_CONFIG_CHANGED", "The system configuration changed during the cluster join or form operation. The join or form operation was aborted."}
	ErrorClusterResourceTypeNotFound                       = &Error{0x000013D6, "ERROR_CLUSTER_RESOURCE_TYPE_NOT_FOUND", "The specified resource type was not found."}
	ErrorClusterRestypeNotSupported                        = &Error{0x000013D7, "ERROR_CLUSTER_RESTYPE_NOT_SUPPORTED", "The specified node does not support a resource of this type. This might be due to version inconsistencies or due to the absence of the resource DLL on this node."}
	ErrorClusterResnameNotFound                            = &Error{0x000013D8, "ERROR_CLUSTER_RESNAME_NOT_FOUND", "The specified resource name is not supported by this resource DLL. This might be due to a bad (or changed) name supplied to the resource DLL."}
	ErrorClusterNoRpcPackagesRegistered                    = &Error{0x000013D9, "ERROR_CLUSTER_NO_RPC_PACKAGES_REGISTERED", "No authentication package could be registered with the RPC server."}
	ErrorClusterOwnerNotInPreflist                         = &Error{0x000013DA, "ERROR_CLUSTER_OWNER_NOT_IN_PREFLIST", "You cannot bring the group online because the owner of the group is not in the preferred list for the group. To change the owner node for the group, move the group."}
	ErrorClusterDatabaseSeqmismatch                        = &Error{0x000013DB, "ERROR_CLUSTER_DATABASE_SEQMISMATCH", "The join operation failed because the cluster database sequence number has changed or is incompatible with the locker node. This can happen during a join operation if the cluster database was changing during the join."}
	ErrorResmonInvalidState                                = &Error{0x000013DC, "ERROR_RESMON_INVALID_STATE", "The resource monitor will not allow the fail operation to be performed while the resource is in its current state. This can happen if the resource is in a pending state."}
	ErrorClusterGumNotLocker                               = &Error{0x000013DD, "ERROR_CLUSTER_GUM_NOT_LOCKER", "A non-locker code received a request to reserve the lock for making global updates."}
	ErrorQuorumDiskNotFound                                = &Error{0x000013DE, "ERROR_QUORUM_DISK_NOT_FOUND", "The quorum disk could not be located by the cluster service."}
	ErrorDatabaseBackupCorrupt                             = &Error{0x000013DF, "ERROR_DATABASE_BACKUP_CORRUPT", "The backed-up cluster database is possibly corrupt."}
	ErrorClusterNodeAlreadyHasDfsRoot                      = &Error{0x000013E0, "ERROR_CLUSTER_NODE_ALREADY_HAS_DFS_ROOT", "A DFS root already exists in this cluster node."}
	ErrorResourcePropertyUnchangeable                      = &Error{0x000013E1, "ERROR_RESOURCE_PROPERTY_UNCHANGEABLE", "An attempt to modify a resource property failed because it conflicts with another existing property."}
	ErrorClusterMembershipInvalidState                     = &Error{0x00001702, "ERROR_CLUSTER_MEMBERSHIP_INVALID_STATE", "An operation was attempted that is incompatible with the current membership state of the node."}
	ErrorClusterQuorumlogNotFound                          = &Error{0x00001703, "ERROR_CLUSTER_QUORUMLOG_NOT_FOUND", "The quorum resource does not contain the quorum log."}
	ErrorClusterMembershipHalt                             = &Error{0x00001704, "ERROR_CLUSTER_MEMBERSHIP_HALT", "The membership engine requested shutdown of the cluster service on this node."}
	ErrorClusterInstanceIdMismatch                         = &Error{0x00001705, "ERROR_CLUSTER_INSTANCE_ID_MISMATCH", "The join operation failed because the cluster instance ID of the joining node does not match the cluster instance ID of the sponsor node."}
	ErrorClusterNetworkNotFoundForIp                       = &Error{0x00001706, "ERROR_CLUSTER_NETWORK_NOT_FOUND_FOR_IP", "A matching cluster network for the specified IP address could not be found."}
	ErrorClusterPropertyDataTypeMismatch                   = &Error{0x00001707, "ERROR_CLUSTER_PROPERTY_DATA_TYPE_MISMATCH", "The actual data type of the property did not match the expected data type of the property."}
	ErrorClusterEvictWithoutCleanup                        = &Error{0x00001708, "ERROR_CLUSTER_EVICT_WITHOUT_CLEANUP", "The cluster node was evicted from the cluster successfully, but the node was not cleaned up. To determine what clean-up steps failed and how to recover, see the Failover Clustering application event log using Event Viewer."}
	ErrorClusterParameterMismatch                          = &Error{0x00001709, "ERROR_CLUSTER_PARAMETER_MISMATCH", "Two or more parameter values specified for a resource's properties are in conflict."}
	ErrorNodeCannotBeClustered                             = &Error{0x0000170A, "ERROR_NODE_CANNOT_BE_CLUSTERED", "This computer cannot be made a member of a cluster."}
	ErrorClusterWrongOsVersion                             = &Error{0x0000170B, "ERROR_CLUSTER_WRONG_OS_VERSION", "This computer cannot be made a member of a cluster because it does not have the correct version of Windows installed."}
	ErrorClusterCantCreateDupClusterName                   = &Error{0x0000170C, "ERROR_CLUSTER_CANT_CREATE_DUP_CLUSTER_NAME", "A cluster cannot be created with the specified cluster name because that cluster name is already in use. Specify a different name for the cluster."}
	ErrorCluscfgAlreadyCommitted                           = &Error{0x0000170D, "ERROR_CLUSCFG_ALREADY_COMMITTED", "The cluster configuration action has already been committed."}
	ErrorCluscfgRollbackFailed                             = &Error{0x0000170E, "ERROR_CLUSCFG_ROLLBACK_FAILED", "The cluster configuration action could not be rolled back."}
	ErrorCluscfgSystemDiskDriveLetterConflict              = &Error{0x0000170F, "ERROR_CLUSCFG_SYSTEM_DISK_DRIVE_LETTER_CONFLICT", "The drive letter assigned to a system disk on one node conflicted with the drive letter assigned to a disk on another node."}
	ErrorClusterOldVersion                                 = &Error{0x00001710, "ERROR_CLUSTER_OLD_VERSION", "One or more nodes in the cluster are running a version of Windows that does not support this operation."}
	ErrorClusterMismatchedComputerAcctName                 = &Error{0x00001711, "ERROR_CLUSTER_MISMATCHED_COMPUTER_ACCT_NAME", "The name of the corresponding computer account does not match the network name for this resource."}
	ErrorClusterNoNetAdapters                              = &Error{0x00001712, "ERROR_CLUSTER_NO_NET_ADAPTERS", "No network adapters are available."}
	ErrorClusterPoisoned                                   = &Error{0x00001713, "ERROR_CLUSTER_POISONED", "The cluster node has been poisoned."}
	ErrorClusterGroupMoving                                = &Error{0x00001714, "ERROR_CLUSTER_GROUP_MOVING", "The group is unable to accept the request because it is moving to another node."}
	ErrorClusterResourceTypeBusy                           = &Error{0x00001715, "ERROR_CLUSTER_RESOURCE_TYPE_BUSY", "The resource type cannot accept the request because it is too busy performing another operation."}
	ErrorResourceCallTimedOut                              = &Error{0x00001716, "ERROR_RESOURCE_CALL_TIMED_OUT", "The call to the cluster resource DLL timed out."}
	ErrorInvalidClusterIpv6Address                         = &Error{0x00001717, "ERROR_INVALID_CLUSTER_IPV6_ADDRESS", "The address is not valid for an IPv6 Address resource. A global IPv6 address is required, and it must match a cluster network. Compatibility addresses are not permitted."}
	ErrorClusterInternalInvalidFunction                    = &Error{0x00001718, "ERROR_CLUSTER_INTERNAL_INVALID_FUNCTION", "An internal cluster error occurred. A call to an invalid function was attempted."}
	ErrorClusterParameterOutOfBounds                       = &Error{0x00001719, "ERROR_CLUSTER_PARAMETER_OUT_OF_BOUNDS", "A parameter value is out of acceptable range."}
	ErrorClusterPartialSend                                = &Error{0x0000171A, "ERROR_CLUSTER_PARTIAL_SEND", "A network error occurred while sending data to another node in the cluster. The number of bytes transmitted was less than required."}
	ErrorClusterRegistryInvalidFunction                    = &Error{0x0000171B, "ERROR_CLUSTER_REGISTRY_INVALID_FUNCTION", "An invalid cluster registry operation was attempted."}
	ErrorClusterInvalidStringTermination                   = &Error{0x0000171C, "ERROR_CLUSTER_INVALID_STRING_TERMINATION", "An input string of characters is not properly terminated."}
	ErrorClusterInvalidStringFormat                        = &Error{0x0000171D, "ERROR_CLUSTER_INVALID_STRING_FORMAT", "An input string of characters is not in a valid format for the data it represents."}
	ErrorClusterDatabaseTransactionInProgress              = &Error{0x0000171E, "ERROR_CLUSTER_DATABASE_TRANSACTION_IN_PROGRESS", "An internal cluster error occurred. A cluster database transaction was attempted while a transaction was already in progress."}
	ErrorClusterDatabaseTransactionNotInProgress           = &Error{0x0000171F, "ERROR_CLUSTER_DATABASE_TRANSACTION_NOT_IN_PROGRESS", "An internal cluster error occurred. There was an attempt to commit a cluster database transaction while no transaction was in progress."}
	ErrorClusterNullData                                   = &Error{0x00001720, "ERROR_CLUSTER_NULL_DATA", "An internal cluster error occurred. Data was not properly initialized."}
	ErrorClusterPartialRead                                = &Error{0x00001721, "ERROR_CLUSTER_PARTIAL_READ", "An error occurred while reading from a stream of data. An unexpected number of bytes was returned."}
	ErrorClusterPartialWrite                               = &Error{0x00001722, "ERROR_CLUSTER_PARTIAL_WRITE", "An error occurred while writing to a stream of data. The required number of bytes could not be written."}
	ErrorClusterCantDeserializeData                        = &Error{0x00001723, "ERROR_CLUSTER_CANT_DESERIALIZE_DATA", "An error occurred while deserializing a stream of cluster data."}
	ErrorDependentResourcePropertyConflict                 = &Error{0x00001724, "ERROR_DEPENDENT_RESOURCE_PROPERTY_CONFLICT", "One or more property values for this resource are in conflict with one or more property values associated with its dependent resources."}
	ErrorClusterNoQuorum                                   = &Error{0x00001725, "ERROR_CLUSTER_NO_QUORUM", "A quorum of cluster nodes was not present to form a cluster."}
	ErrorClusterInvalidIpv6Network                         = &Error{0x00001726, "ERROR_CLUSTER_INVALID_IPV6_NETWORK", "The cluster network is not valid for an IPv6 address resource, or it does not match the configured address."}
	ErrorClusterInvalidIpv6TunnelNetwork                   = &Error{0x00001727, "ERROR_CLUSTER_INVALID_IPV6_TUNNEL_NETWORK", "The cluster network is not valid for an IPv6 tunnel resource. Check the configuration of the IP Address resource on which the IPv6 tunnel resource depends."}
	ErrorQuorumNotAllowedInThisGroup                       = &Error{0x00001728, "ERROR_QUORUM_NOT_ALLOWED_IN_THIS_GROUP", "Quorum resource cannot reside in the available storage group."}
	ErrorEncryptionFailed                                  = &Error{0x00001770, "ERROR_ENCRYPTION_FAILED", "The specified file could not be encrypted."}
	ErrorDecryptionFailed                                  = &Error{0x00001771, "ERROR_DECRYPTION_FAILED", "The specified file could not be decrypted."}
	ErrorFileEncrypted                                     = &Error{0x00001772, "ERROR_FILE_ENCRYPTED", "The specified file is encrypted and the user does not have the ability to decrypt it."}
	ErrorNoRecoveryPolicy                                  = &Error{0x00001773, "ERROR_NO_RECOVERY_POLICY", "There is no valid encryption recovery policy configured for this system."}
	ErrorNoEfs                                             = &Error{0x00001774, "ERROR_NO_EFS", "The required encryption driver is not loaded for this system."}
	ErrorWrongEfs                                          = &Error{0x00001775, "ERROR_WRONG_EFS", "The file was encrypted with a different encryption driver than is currently loaded."}
	ErrorNoUserKeys                                        = &Error{0x00001776, "ERROR_NO_USER_KEYS", "There are no Encrypting File System (EFS) keys defined for the user."}
	ErrorFileNotEncrypted                                  = &Error{0x00001777, "ERROR_FILE_NOT_ENCRYPTED", "The specified file is not encrypted."}
	ErrorNotExportFormat                                   = &Error{0x00001778, "ERROR_NOT_EXPORT_FORMAT", "The specified file is not in the defined EFS export format."}
	ErrorFileReadOnly                                      = &Error{0x00001779, "ERROR_FILE_READ_ONLY", "The specified file is read-only."}
	ErrorDirEfsDisallowed                                  = &Error{0x0000177A, "ERROR_DIR_EFS_DISALLOWED", "The directory has been disabled for encryption."}
	ErrorEfsServerNotTrusted                               = &Error{0x0000177B, "ERROR_EFS_SERVER_NOT_TRUSTED", "The server is not trusted for remote encryption operation."}
	ErrorBadRecoveryPolicy                                 = &Error{0x0000177C, "ERROR_BAD_RECOVERY_POLICY", "Recovery policy configured for this system contains invalid recovery certificate."}
	ErrorEfsAlgBlobTooBig                                  = &Error{0x0000177D, "ERROR_EFS_ALG_BLOB_TOO_BIG", "The encryption algorithm used on the source file needs a bigger key buffer than the one on the destination file."}
	ErrorVolumeNotSupportEfs                               = &Error{0x0000177E, "ERROR_VOLUME_NOT_SUPPORT_EFS", "The disk partition does not support file encryption."}
	ErrorEfsDisabled                                       = &Error{0x0000177F, "ERROR_EFS_DISABLED", "This machine is disabled for file encryption."}
	ErrorEfsVersionNotSupport                              = &Error{0x00001780, "ERROR_EFS_VERSION_NOT_SUPPORT", "A newer system is required to decrypt this encrypted file."}
	ErrorCsEncryptionInvalidServerResponse                 = &Error{0x00001781, "ERROR_CS_ENCRYPTION_INVALID_SERVER_RESPONSE", "The remote server sent an invalid response for a file being opened with client-side encryption."}
	ErrorCsEncryptionUnsupportedServer                     = &Error{0x00001782, "ERROR_CS_ENCRYPTION_UNSUPPORTED_SERVER", "Client-side encryption is not supported by the remote server even though it claims to support it."}
	ErrorCsEncryptionExistingEncryptedFile                 = &Error{0x00001783, "ERROR_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE", "File is encrypted and should be opened in client-side encryption mode."}
	ErrorCsEncryptionNewEncryptedFile                      = &Error{0x00001784, "ERROR_CS_ENCRYPTION_NEW_ENCRYPTED_FILE", "A new encrypted file is being created and a $EFS needs to be provided."}
	ErrorCsEncryptionFileNotCse                            = &Error{0x00001785, "ERROR_CS_ENCRYPTION_FILE_NOT_CSE", "The SMB client requested a client-side extension (CSE) file system control (FSCTL) on a non-CSE file."}
	ErrorNoBrowserServersFound                             = &Error{0x000017E6, "ERROR_NO_BROWSER_SERVERS_FOUND", "The list of servers for this workgroup is not currently available"}
	SchedEServiceNotLocalsystem                            = &Error{0x00001838, "SCHED_E_SERVICE_NOT_LOCALSYSTEM", "The Task Scheduler service must be configured to run in the System account to function properly. Individual tasks can be configured to run in other accounts."}
	ErrorLogSectorInvalid                                  = &Error{0x000019C8, "ERROR_LOG_SECTOR_INVALID", "The log service encountered an invalid log sector."}
	ErrorLogSectorParityInvalid                            = &Error{0x000019C9, "ERROR_LOG_SECTOR_PARITY_INVALID", "The log service encountered a log sector with invalid block parity."}
	ErrorLogSectorRemapped                                 = &Error{0x000019CA, "ERROR_LOG_SECTOR_REMAPPED", "The log service encountered a remapped log sector."}
	ErrorLogBlockIncomplete                                = &Error{0x000019CB, "ERROR_LOG_BLOCK_INCOMPLETE", "The log service encountered a partial or incomplete log block."}
	ErrorLogInvalidRange                                   = &Error{0x000019CC, "ERROR_LOG_INVALID_RANGE", "The log service encountered an attempt to access data outside the active log range."}
	ErrorLogBlocksExhausted                                = &Error{0x000019CD, "ERROR_LOG_BLOCKS_EXHAUSTED", "The log service user marshaling buffers are exhausted."}
	ErrorLogReadContextInvalid                             = &Error{0x000019CE, "ERROR_LOG_READ_CONTEXT_INVALID", "The log service encountered an attempt to read from a marshaling area with an invalid read context."}
	ErrorLogRestartInvalid                                 = &Error{0x000019CF, "ERROR_LOG_RESTART_INVALID", "The log service encountered an invalid log restart area."}
	ErrorLogBlockVersion                                   = &Error{0x000019D0, "ERROR_LOG_BLOCK_VERSION", "The log service encountered an invalid log block version."}
	ErrorLogBlockInvalid                                   = &Error{0x000019D1, "ERROR_LOG_BLOCK_INVALID", "The log service encountered an invalid log block."}
	ErrorLogReadModeInvalid                                = &Error{0x000019D2, "ERROR_LOG_READ_MODE_INVALID", "The log service encountered an attempt to read the log with an invalid read mode."}
	ErrorLogNoRestart                                      = &Error{0x000019D3, "ERROR_LOG_NO_RESTART", "The log service encountered a log stream with no restart area."}
	ErrorLogMetadataCorrupt                                = &Error{0x000019D4, "ERROR_LOG_METADATA_CORRUPT", "The log service encountered a corrupted metadata file."}
	ErrorLogMetadataInvalid                                = &Error{0x000019D5, "ERROR_LOG_METADATA_INVALID", "The log service encountered a metadata file that could not be created by the log file system."}
	ErrorLogMetadataInconsistent                           = &Error{0x000019D6, "ERROR_LOG_METADATA_INCONSISTENT", "The log service encountered a metadata file with inconsistent data."}
	ErrorLogReservationInvalid                             = &Error{0x000019D7, "ERROR_LOG_RESERVATION_INVALID", "The log service encountered an attempt to erroneous allocate or dispose reservation space."}
	ErrorLogCantDelete                                     = &Error{0x000019D8, "ERROR_LOG_CANT_DELETE", "The log service cannot delete a log file or file system container."}
	ErrorLogContainerLimitExceeded                         = &Error{0x000019D9, "ERROR_LOG_CONTAINER_LIMIT_EXCEEDED", "The log service has reached the maximum allowable containers allocated to a log file."}
	ErrorLogStartOfLog                                     = &Error{0x000019DA, "ERROR_LOG_START_OF_LOG", "The log service has attempted to read or write backward past the start of the log."}
	ErrorLogPolicyAlreadyInstalled                         = &Error{0x000019DB, "ERROR_LOG_POLICY_ALREADY_INSTALLED", "The log policy could not be installed because a policy of the same type is already present."}
	ErrorLogPolicyNotInstalled                             = &Error{0x000019DC, "ERROR_LOG_POLICY_NOT_INSTALLED", "The log policy in question was not installed at the time of the request."}
	ErrorLogPolicyInvalid                                  = &Error{0x000019DD, "ERROR_LOG_POLICY_INVALID", "The installed set of policies on the log is invalid."}
	ErrorLogPolicyConflict                                 = &Error{0x000019DE, "ERROR_LOG_POLICY_CONFLICT", "A policy on the log in question prevented the operation from completing."}
	ErrorLogPinnedArchiveTail                              = &Error{0x000019DF, "ERROR_LOG_PINNED_ARCHIVE_TAIL", "Log space cannot be reclaimed because the log is pinned by the archive tail."}
	ErrorLogRecordNonexistent                              = &Error{0x000019E0, "ERROR_LOG_RECORD_NONEXISTENT", "The log record is not a record in the log file."}
	ErrorLogRecordsReservedInvalid                         = &Error{0x000019E1, "ERROR_LOG_RECORDS_RESERVED_INVALID", "The number of reserved log records or the adjustment of the number of reserved log records is invalid."}
	ErrorLogSpaceReservedInvalid                           = &Error{0x000019E2, "ERROR_LOG_SPACE_RESERVED_INVALID", "The reserved log space or the adjustment of the log space is invalid."}
	ErrorLogTailInvalid                                    = &Error{0x000019E3, "ERROR_LOG_TAIL_INVALID", "A new or existing archive tail or base of the active log is invalid."}
	ErrorLogFull                                           = &Error{0x000019E4, "ERROR_LOG_FULL", "The log space is exhausted."}
	ErrorCouldNotResizeLog                                 = &Error{0x000019E5, "ERROR_COULD_NOT_RESIZE_LOG", "The log could not be set to the requested size."}
	ErrorLogMultiplexed                                    = &Error{0x000019E6, "ERROR_LOG_MULTIPLEXED", "The log is multiplexed; no direct writes to the physical log are allowed."}
	ErrorLogDedicated                                      = &Error{0x000019E7, "ERROR_LOG_DEDICATED", "The operation failed because the log is a dedicated log."}
	ErrorLogArchiveNotInProgress                           = &Error{0x000019E8, "ERROR_LOG_ARCHIVE_NOT_IN_PROGRESS", "The operation requires an archive context."}
	ErrorLogArchiveInProgress                              = &Error{0x000019E9, "ERROR_LOG_ARCHIVE_IN_PROGRESS", "Log archival is in progress."}
	ErrorLogEphemeral                                      = &Error{0x000019EA, "ERROR_LOG_EPHEMERAL", "The operation requires a non-ephemeral log, but the log is ephemeral."}
	ErrorLogNotEnoughContainers                            = &Error{0x000019EB, "ERROR_LOG_NOT_ENOUGH_CONTAINERS", "The log must have at least two containers before it can be read from or written to."}
	ErrorLogClientAlreadyRegistered                        = &Error{0x000019EC, "ERROR_LOG_CLIENT_ALREADY_REGISTERED", "A log client has already registered on the stream."}
	ErrorLogClientNotRegistered                            = &Error{0x000019ED, "ERROR_LOG_CLIENT_NOT_REGISTERED", "A log client has not been registered on the stream."}
	ErrorLogFullHandlerInProgress                          = &Error{0x000019EE, "ERROR_LOG_FULL_HANDLER_IN_PROGRESS", "A request has already been made to handle the log full condition."}
	ErrorLogContainerReadFailed                            = &Error{0x000019EF, "ERROR_LOG_CONTAINER_READ_FAILED", "The log service encountered an error when attempting to read from a log container."}
	ErrorLogContainerWriteFailed                           = &Error{0x000019F0, "ERROR_LOG_CONTAINER_WRITE_FAILED", "The log service encountered an error when attempting to write to a log container."}
	ErrorLogContainerOpenFailed                            = &Error{0x000019F1, "ERROR_LOG_CONTAINER_OPEN_FAILED", "The log service encountered an error when attempting to open a log container."}
	ErrorLogContainerStateInvalid                          = &Error{0x000019F2, "ERROR_LOG_CONTAINER_STATE_INVALID", "The log service encountered an invalid container state when attempting a requested action."}
	ErrorLogStateInvalid                                   = &Error{0x000019F3, "ERROR_LOG_STATE_INVALID", "The log service is not in the correct state to perform a requested action."}
	ErrorLogPinned                                         = &Error{0x000019F4, "ERROR_LOG_PINNED", "The log space cannot be reclaimed because the log is pinned."}
	ErrorLogMetadataFlushFailed                            = &Error{0x000019F5, "ERROR_LOG_METADATA_FLUSH_FAILED", "The log metadata flush failed."}
	ErrorLogInconsistentSecurity                           = &Error{0x000019F6, "ERROR_LOG_INCONSISTENT_SECURITY", "Security on the log and its containers is inconsistent."}
	ErrorLogAppendedFlushFailed                            = &Error{0x000019F7, "ERROR_LOG_APPENDED_FLUSH_FAILED", "Records were appended to the log or reservation changes were made, but the log could not be flushed."}
	ErrorLogPinnedReservation                              = &Error{0x000019F8, "ERROR_LOG_PINNED_RESERVATION", "The log is pinned due to reservation consuming most of the log space. Free some reserved records to make space available."}
	ErrorInvalidTransaction                                = &Error{0x00001A2C, "ERROR_INVALID_TRANSACTION", "The transaction handle associated with this operation is not valid."}
	ErrorTransactionNotActive                              = &Error{0x00001A2D, "ERROR_TRANSACTION_NOT_ACTIVE", "The requested operation was made in the context of a transaction that is no longer active."}
	ErrorTransactionRequestNotValid                        = &Error{0x00001A2E, "ERROR_TRANSACTION_REQUEST_NOT_VALID", "The requested operation is not valid on the transaction object in its current state."}
	ErrorTransactionNotRequested                           = &Error{0x00001A2F, "ERROR_TRANSACTION_NOT_REQUESTED", "The caller has called a response API, but the response is not expected because the transaction manager did not issue the corresponding request to the caller."}
	ErrorTransactionAlreadyAborted                         = &Error{0x00001A30, "ERROR_TRANSACTION_ALREADY_ABORTED", "It is too late to perform the requested operation because the transaction has already been aborted."}
	ErrorTransactionAlreadyCommitted                       = &Error{0x00001A31, "ERROR_TRANSACTION_ALREADY_COMMITTED", "It is too late to perform the requested operation because the transaction has already been committed."}
	ErrorTmInitializationFailed                            = &Error{0x00001A32, "ERROR_TM_INITIALIZATION_FAILED", "The transaction manager was unable to be successfully initialized. Transacted operations are not supported."}
	ErrorResourcemanagerReadOnly                           = &Error{0x00001A33, "ERROR_RESOURCEMANAGER_READ_ONLY", "The specified resource manager made no changes or updates to the resource under this transaction."}
	ErrorTransactionNotJoined                              = &Error{0x00001A34, "ERROR_TRANSACTION_NOT_JOINED", "The resource manager has attempted to prepare a transaction that it has not successfully joined."}
	ErrorTransactionSuperiorExists                         = &Error{0x00001A35, "ERROR_TRANSACTION_SUPERIOR_EXISTS", "The transaction object already has a superior enlistment, and the caller attempted an operation that would have created a new superior. Only a single superior enlistment is allowed."}
	ErrorCrmProtocolAlreadyExists                          = &Error{0x00001A36, "ERROR_CRM_PROTOCOL_ALREADY_EXISTS", "The resource manager tried to register a protocol that already exists."}
	ErrorTransactionPropagationFailed                      = &Error{0x00001A37, "ERROR_TRANSACTION_PROPAGATION_FAILED", "The attempt to propagate the transaction failed."}
	ErrorCrmProtocolNotFound                               = &Error{0x00001A38, "ERROR_CRM_PROTOCOL_NOT_FOUND", "The requested propagation protocol was not registered as a CRM."}
	ErrorTransactionInvalidMarshallBuffer                  = &Error{0x00001A39, "ERROR_TRANSACTION_INVALID_MARSHALL_BUFFER", "The buffer passed in to PushTransaction or PullTransaction is not in a valid format."}
	ErrorCurrentTransactionNotValid                        = &Error{0x00001A3A, "ERROR_CURRENT_TRANSACTION_NOT_VALID", "The current transaction context associated with the thread is not a valid handle to a transaction object."}
	ErrorTransactionNotFound                               = &Error{0x00001A3B, "ERROR_TRANSACTION_NOT_FOUND", "The specified transaction object could not be opened because it was not found."}
	ErrorResourcemanagerNotFound                           = &Error{0x00001A3C, "ERROR_RESOURCEMANAGER_NOT_FOUND", "The specified resource manager object could not be opened because it was not found."}
	ErrorEnlistmentNotFound                                = &Error{0x00001A3D, "ERROR_ENLISTMENT_NOT_FOUND", "The specified enlistment object could not be opened because it was not found."}
	ErrorTransactionmanagerNotFound                        = &Error{0x00001A3E, "ERROR_TRANSACTIONMANAGER_NOT_FOUND", "The specified transaction manager object could not be opened because it was not found."}
	ErrorTransactionmanagerNotOnline                       = &Error{0x00001A3F, "ERROR_TRANSACTIONMANAGER_NOT_ONLINE", "The specified resource manager was unable to create an enlistment because its associated transaction manager is not online."}
	ErrorTransactionmanagerRecoveryNameCollision           = &Error{0x00001A40, "ERROR_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION", "The specified transaction manager was unable to create the objects contained in its log file in the ObjectB namespace. Therefore, the transaction manager was unable to recover."}
	ErrorTransactionalConflict                             = &Error{0x00001A90, "ERROR_TRANSACTIONAL_CONFLICT", "The function attempted to use a name that is reserved for use by another transaction."}
	ErrorRmNotActive                                       = &Error{0x00001A91, "ERROR_RM_NOT_ACTIVE", "Transaction support within the specified file system resource manager is not started or was shut down due to an error."}
	ErrorRmMetadataCorrupt                                 = &Error{0x00001A92, "ERROR_RM_METADATA_CORRUPT", "The metadata of the resource manager has been corrupted. The resource manager will not function."}
	ErrorDirectoryNotRm                                    = &Error{0x00001A93, "ERROR_DIRECTORY_NOT_RM", "The specified directory does not contain a resource manager."}
	ErrorTransactionsUnsupportedRemote                     = &Error{0x00001A95, "ERROR_TRANSACTIONS_UNSUPPORTED_REMOTE", "The remote server or share does not support transacted file operations."}
	ErrorLogResizeInvalidSize                              = &Error{0x00001A96, "ERROR_LOG_RESIZE_INVALID_SIZE", "The requested log size is invalid."}
	ErrorObjectNoLongerExists                              = &Error{0x00001A97, "ERROR_OBJECT_NO_LONGER_EXISTS", "The object (file, stream, link) corresponding to the handle has been deleted by a transaction savepoint rollback."}
	ErrorStreamMiniversionNotFound                         = &Error{0x00001A98, "ERROR_STREAM_MINIVERSION_NOT_FOUND", "The specified file miniversion was not found for this transacted file open."}
	ErrorStreamMiniversionNotValid                         = &Error{0x00001A99, "ERROR_STREAM_MINIVERSION_NOT_VALID", "The specified file miniversion was found but has been invalidated. The most likely cause is a transaction savepoint rollback."}
	ErrorMiniversionInaccessibleFromSpecifiedTransaction   = &Error{0x00001A9A, "ERROR_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION", "A miniversion can only be opened in the context of the transaction that created it."}
	ErrorCantOpenMiniversionWithModifyIntent               = &Error{0x00001A9B, "ERROR_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT", "It is not possible to open a miniversion with modify access."}
	ErrorCantCreateMoreStreamMiniversions                  = &Error{0x00001A9C, "ERROR_CANT_CREATE_MORE_STREAM_MINIVERSIONS", "It is not possible to create any more miniversions for this stream."}
	ErrorRemoteFileVersionMismatch                         = &Error{0x00001A9E, "ERROR_REMOTE_FILE_VERSION_MISMATCH", "The remote server sent mismatching version numbers or FID for a file opened with transactions."}
	ErrorHandleNoLongerValid                               = &Error{0x00001A9F, "ERROR_HANDLE_NO_LONGER_VALID", "The handle has been invalidated by a transaction. The most likely cause is the presence of memory mapping on a file, or an open handle when the transaction ended or rolled back to savepoint."}
	ErrorNoTxfMetadata                                     = &Error{0x00001AA0, "ERROR_NO_TXF_METADATA", "There is no transaction metadata on the file."}
	ErrorLogCorruptionDetected                             = &Error{0x00001AA1, "ERROR_LOG_CORRUPTION_DETECTED", "The log data is corrupt."}
	ErrorCantRecoverWithHandleOpen                         = &Error{0x00001AA2, "ERROR_CANT_RECOVER_WITH_HANDLE_OPEN", "The file cannot be recovered because a handle is still open on it."}
	ErrorRmDisconnected                                    = &Error{0x00001AA3, "ERROR_RM_DISCONNECTED", "The transaction outcome is unavailable because the resource manager responsible for it is disconnected."}
	ErrorEnlistmentNotSuperior                             = &Error{0x00001AA4, "ERROR_ENLISTMENT_NOT_SUPERIOR", "The request was rejected because the enlistment in question is not a superior enlistment."}
	ErrorRecoveryNotNeeded                                 = &Error{0x00001AA5, "ERROR_RECOVERY_NOT_NEEDED", "The transactional resource manager is already consistent. Recovery is not needed."}
	ErrorRmAlreadyStarted                                  = &Error{0x00001AA6, "ERROR_RM_ALREADY_STARTED", "The transactional resource manager has already been started."}
	ErrorFileIdentityNotPersistent                         = &Error{0x00001AA7, "ERROR_FILE_IDENTITY_NOT_PERSISTENT", "The file cannot be opened in a transaction because its identity depends on the outcome of an unresolved transaction."}
	ErrorCantBreakTransactionalDependency                  = &Error{0x00001AA8, "ERROR_CANT_BREAK_TRANSACTIONAL_DEPENDENCY", "The operation cannot be performed because another transaction is depending on the fact that this property will not change."}
	ErrorCantCrossRmBoundary                               = &Error{0x00001AA9, "ERROR_CANT_CROSS_RM_BOUNDARY", "The operation would involve a single file with two transactional resource managers and is therefore not allowed."}
	ErrorTxfDirNotEmpty                                    = &Error{0x00001AAA, "ERROR_TXF_DIR_NOT_EMPTY", "The $Txf directory must be empty for this operation to succeed."}
	ErrorIndoubtTransactionsExist                          = &Error{0x00001AAB, "ERROR_INDOUBT_TRANSACTIONS_EXIST", "The operation would leave a transactional resource manager in an inconsistent state and is, therefore, not allowed."}
	ErrorTmVolatile                                        = &Error{0x00001AAC, "ERROR_TM_VOLATILE", "The operation could not be completed because the transaction manager does not have a log."}
	ErrorRollbackTimerExpired                              = &Error{0x00001AAD, "ERROR_ROLLBACK_TIMER_EXPIRED", "A rollback could not be scheduled because a previously scheduled rollback has already been executed or is queued for execution."}
	ErrorTxfAttributeCorrupt                               = &Error{0x00001AAE, "ERROR_TXF_ATTRIBUTE_CORRUPT", "The transactional metadata attribute on the file or directory is corrupt and unreadable."}
	ErrorEfsNotAllowedInTransaction                        = &Error{0x00001AAF, "ERROR_EFS_NOT_ALLOWED_IN_TRANSACTION", "The encryption operation could not be completed because a transaction is active."}
	ErrorTransactionalOpenNotAllowed                       = &Error{0x00001AB0, "ERROR_TRANSACTIONAL_OPEN_NOT_ALLOWED", "This object is not allowed to be opened in a transaction."}
	ErrorLogGrowthFailed                                   = &Error{0x00001AB1, "ERROR_LOG_GROWTH_FAILED", "An attempt to create space in the transactional resource manager's log failed. The failure status has been recorded in the event log."}
	ErrorTransactedMappingUnsupportedRemote                = &Error{0x00001AB2, "ERROR_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE", "Memory mapping (creating a mapped section) to a remote file under a transaction is not supported."}
	ErrorTxfMetadataAlreadyPresent                         = &Error{0x00001AB3, "ERROR_TXF_METADATA_ALREADY_PRESENT", "Transaction metadata is already present on this file and cannot be superseded."}
	ErrorTransactionScopeCallbacksNotSet                   = &Error{0x00001AB4, "ERROR_TRANSACTION_SCOPE_CALLBACKS_NOT_SET", "A transaction scope could not be entered because the scope handler has not been initialized."}
	ErrorTransactionRequiredPromotion                      = &Error{0x00001AB5, "ERROR_TRANSACTION_REQUIRED_PROMOTION", "Promotion was required to allow the resource manager to enlist, but the transaction was set to disallow it."}
	ErrorCannotExecuteFileInTransaction                    = &Error{0x00001AB6, "ERROR_CANNOT_EXECUTE_FILE_IN_TRANSACTION", "This file is open for modification in an unresolved transaction and can be opened for execution only by a transacted reader."}
	ErrorTransactionsNotFrozen                             = &Error{0x00001AB7, "ERROR_TRANSACTIONS_NOT_FROZEN", "The request to thaw frozen transactions was ignored because transactions were not previously frozen."}
	ErrorTransactionFreezeInProgress                       = &Error{0x00001AB8, "ERROR_TRANSACTION_FREEZE_IN_PROGRESS", "Transactions cannot be frozen because a freeze is already in progress."}
	ErrorNotSnapshotVolume                                 = &Error{0x00001AB9, "ERROR_NOT_SNAPSHOT_VOLUME", "The target volume is not a snapshot volume. This operation is only valid on a volume mounted as a snapshot."}
	ErrorNoSavepointWithOpenFiles                          = &Error{0x00001ABA, "ERROR_NO_SAVEPOINT_WITH_OPEN_FILES", "The savepoint operation failed because files are open on the transaction. This is not permitted."}
	ErrorDataLostRepair                                    = &Error{0x00001ABB, "ERROR_DATA_LOST_REPAIR", "Windows has discovered corruption in a file, and that file has since been repaired. Data loss might have occurred."}
	ErrorSparseNotAllowedInTransaction                     = &Error{0x00001ABC, "ERROR_SPARSE_NOT_ALLOWED_IN_TRANSACTION", "The sparse operation could not be completed because a transaction is active on the file."}
	ErrorTmIdentityMismatch                                = &Error{0x00001ABD, "ERROR_TM_IDENTITY_MISMATCH", "The call to create a transaction manager object failed because the Tm Identity stored in the logfile does not match the Tm Identity that was passed in as an argument."}
	ErrorFloatedSection                                    = &Error{0x00001ABE, "ERROR_FLOATED_SECTION", "I/O was attempted on a section object that has been floated as a result of a transaction ending. There is no valid data."}
	ErrorCannotAcceptTransactedWork                        = &Error{0x00001ABF, "ERROR_CANNOT_ACCEPT_TRANSACTED_WORK", "The transactional resource manager cannot currently accept transacted work due to a transient condition, such as low resources."}
	ErrorCannotAbortTransactions                           = &Error{0x00001AC0, "ERROR_CANNOT_ABORT_TRANSACTIONS", "The transactional resource manager had too many transactions outstanding that could not be aborted. The transactional resource manager has been shut down."}
	ErrorCtxWinstationNameInvalid                          = &Error{0x00001B59, "ERROR_CTX_WINSTATION_NAME_INVALID", "The specified session name is invalid."}
	ErrorCtxInvalidPd                                      = &Error{0x00001B5A, "ERROR_CTX_INVALID_PD", "The specified protocol driver is invalid."}
	ErrorCtxPdNotFound                                     = &Error{0x00001B5B, "ERROR_CTX_PD_NOT_FOUND", "The specified protocol driver was not found in the system path."}
	ErrorCtxWdNotFound                                     = &Error{0x00001B5C, "ERROR_CTX_WD_NOT_FOUND", "The specified terminal connection driver was not found in the system path."}
	ErrorCtxCannotMakeEventlogEntry                        = &Error{0x00001B5D, "ERROR_CTX_CANNOT_MAKE_EVENTLOG_ENTRY", "A registry key for event logging could not be created for this session."}
	ErrorCtxServiceNameCollision                           = &Error{0x00001B5E, "ERROR_CTX_SERVICE_NAME_COLLISION", "A service with the same name already exists on the system."}
	ErrorCtxClosePending                                   = &Error{0x00001B5F, "ERROR_CTX_CLOSE_PENDING", "A close operation is pending on the session."}
	ErrorCtxNoOutbuf                                       = &Error{0x00001B60, "ERROR_CTX_NO_OUTBUF", "There are no free output buffers available."}
	ErrorCtxModemInfNotFound                               = &Error{0x00001B61, "ERROR_CTX_MODEM_INF_NOT_FOUND", "The MODEM.INF file was not found."}
	ErrorCtxInvalidModemname                               = &Error{0x00001B62, "ERROR_CTX_INVALID_MODEMNAME", "The modem name was not found in the MODEM.INF file."}
	ErrorCtxModemResponseError                             = &Error{0x00001B63, "ERROR_CTX_MODEM_RESPONSE_ERROR", "The modem did not accept the command sent to it. Verify that the configured modem name matches the attached modem."}
	ErrorCtxModemResponseTimeout                           = &Error{0x00001B64, "ERROR_CTX_MODEM_RESPONSE_TIMEOUT", "The modem did not respond to the command sent to it. Verify that the modem is properly cabled and turned on."}
	ErrorCtxModemResponseNoCarrier                         = &Error{0x00001B65, "ERROR_CTX_MODEM_RESPONSE_NO_CARRIER", "Carrier detect has failed or carrier has been dropped due to disconnect."}
	ErrorCtxModemResponseNoDialtone                        = &Error{0x00001B66, "ERROR_CTX_MODEM_RESPONSE_NO_DIALTONE", "Dial tone not detected within the required time. Verify that the phone cable is properly attached and functional."}
	ErrorCtxModemResponseBusy                              = &Error{0x00001B67, "ERROR_CTX_MODEM_RESPONSE_BUSY", "Busy signal detected at remote site on callback."}
	ErrorCtxModemResponseVoice                             = &Error{0x00001B68, "ERROR_CTX_MODEM_RESPONSE_VOICE", "Voice detected at remote site on callback."}
	ErrorCtxTdError                                        = &Error{0x00001B69, "ERROR_CTX_TD_ERROR", "Transport driver error."}
	ErrorCtxWinstationNotFound                             = &Error{0x00001B6E, "ERROR_CTX_WINSTATION_NOT_FOUND", "The specified session cannot be found."}
	ErrorCtxWinstationAlreadyExists                        = &Error{0x00001B6F, "ERROR_CTX_WINSTATION_ALREADY_EXISTS", "The specified session name is already in use."}
	ErrorCtxWinstationBusy                                 = &Error{0x00001B70, "ERROR_CTX_WINSTATION_BUSY", "The requested operation cannot be completed because the terminal connection is currently busy processing a connect, disconnect, reset, or delete operation."}
	ErrorCtxBadVideoMode                                   = &Error{0x00001B71, "ERROR_CTX_BAD_VIDEO_MODE", "An attempt has been made to connect to a session whose video mode is not supported by the current client."}
	ErrorCtxGraphicsInvalid                                = &Error{0x00001B7B, "ERROR_CTX_GRAPHICS_INVALID", "The application attempted to enable DOS graphics mode. DOS graphics mode is not supported."}
	ErrorCtxLogonDisabled                                  = &Error{0x00001B7D, "ERROR_CTX_LOGON_DISABLED", "Your interactive logon privilege has been disabled. Contact your administrator."}
	ErrorCtxNotConsole                                     = &Error{0x00001B7E, "ERROR_CTX_NOT_CONSOLE", "The requested operation can be performed only on the system console. This is most often the result of a driver or system DLL requiring direct console access."}
	ErrorCtxClientQueryTimeout                             = &Error{0x00001B80, "ERROR_CTX_CLIENT_QUERY_TIMEOUT", "The client failed to respond to the server connect message."}
	ErrorCtxConsoleDisconnect                              = &Error{0x00001B81, "ERROR_CTX_CONSOLE_DISCONNECT", "Disconnecting the console session is not supported."}
	ErrorCtxConsoleConnect                                 = &Error{0x00001B82, "ERROR_CTX_CONSOLE_CONNECT", "Reconnecting a disconnected session to the console is not supported."}
	ErrorCtxShadowDenied                                   = &Error{0x00001B84, "ERROR_CTX_SHADOW_DENIED", "The request to control another session remotely was denied."}
	ErrorCtxWinstationAccessDenied                         = &Error{0x00001B85, "ERROR_CTX_WINSTATION_ACCESS_DENIED", "The requested session access is denied."}
	ErrorCtxInvalidWd                                      = &Error{0x00001B89, "ERROR_CTX_INVALID_WD", "The specified terminal connection driver is invalid."}
	ErrorCtxShadowInvalid                                  = &Error{0x00001B8A, "ERROR_CTX_SHADOW_INVALID", "The requested session cannot be controlled remotely. This might be because the session is disconnected or does not currently have a user logged on."}
	ErrorCtxShadowDisabled                                 = &Error{0x00001B8B, "ERROR_CTX_SHADOW_DISABLED", "The requested session is not configured to allow remote control."}
	ErrorCtxClientLicenseInUse                             = &Error{0x00001B8C, "ERROR_CTX_CLIENT_LICENSE_IN_USE", "Your request to connect to this terminal server has been rejected. Your terminal server client license number is currently being used by another user. Call your system administrator to obtain a unique license number."}
	ErrorCtxClientLicenseNotSet                            = &Error{0x00001B8D, "ERROR_CTX_CLIENT_LICENSE_NOT_SET", "Your request to connect to this terminal server has been rejected. Your terminal server client license number has not been entered for this copy of the terminal server client. Contact your system administrator."}
	ErrorCtxLicenseNotAvailable                            = &Error{0x00001B8E, "ERROR_CTX_LICENSE_NOT_AVAILABLE", "The number of connections to this computer is limited and all connections are in use right now. Try connecting later or contact your system administrator."}
	ErrorCtxLicenseClientInvalid                           = &Error{0x00001B8F, "ERROR_CTX_LICENSE_CLIENT_INVALID", "The client you are using is not licensed to use this system. Your logon request is denied."}
	ErrorCtxLicenseExpired                                 = &Error{0x00001B90, "ERROR_CTX_LICENSE_EXPIRED", "The system license has expired. Your logon request is denied."}
	ErrorCtxShadowNotRunning                               = &Error{0x00001B91, "ERROR_CTX_SHADOW_NOT_RUNNING", "Remote control could not be terminated because the specified session is not currently being remotely controlled."}
	ErrorCtxShadowEndedByModeChange                        = &Error{0x00001B92, "ERROR_CTX_SHADOW_ENDED_BY_MODE_CHANGE", "The remote control of the console was terminated because the display mode was changed. Changing the display mode in a remote control session is not supported."}
	ErrorActivationCountExceeded                           = &Error{0x00001B93, "ERROR_ACTIVATION_COUNT_EXCEEDED", "Activation has already been reset the maximum number of times for this installation. Your activation timer will not be cleared."}
	ErrorCtxWinstationsDisabled                            = &Error{0x00001B94, "ERROR_CTX_WINSTATIONS_DISABLED", "Remote logons are currently disabled."}
	ErrorCtxEncryptionLevelRequired                        = &Error{0x00001B95, "ERROR_CTX_ENCRYPTION_LEVEL_REQUIRED", "You do not have the proper encryption level to access this session."}
	ErrorCtxSessionInUse                                   = &Error{0x00001B96, "ERROR_CTX_SESSION_IN_USE", "The user %s\\\\%s is currently logged on to this computer. Only the current user or an administrator can log on to this computer."}
	ErrorCtxNoForceLogoff                                  = &Error{0x00001B97, "ERROR_CTX_NO_FORCE_LOGOFF", "The user %s\\\\%s is already logged on to the console of this computer. You do not have permission to log in at this time. To resolve this issue, contact %s\\\\%s and have them log off."}
	ErrorCtxAccountRestriction                             = &Error{0x00001B98, "ERROR_CTX_ACCOUNT_RESTRICTION", "Unable to log you on because of an account restriction."}
	ErrorRdpProtocolError                                  = &Error{0x00001B99, "ERROR_RDP_PROTOCOL_ERROR", "The RDP component %2 detected an error in the protocol stream and has disconnected the client."}
	ErrorCtxCdmConnect                                     = &Error{0x00001B9A, "ERROR_CTX_CDM_CONNECT", "The Client Drive Mapping Service has connected on terminal connection."}
	ErrorCtxCdmDisconnect                                  = &Error{0x00001B9B, "ERROR_CTX_CDM_DISCONNECT", "The Client Drive Mapping Service has disconnected on terminal connection."}
	ErrorCtxSecurityLayerError                             = &Error{0x00001B9C, "ERROR_CTX_SECURITY_LAYER_ERROR", "The terminal server security layer detected an error in the protocol stream and has disconnected the client."}
	ErrorTsIncompatibleSessions                            = &Error{0x00001B9D, "ERROR_TS_INCOMPATIBLE_SESSIONS", "The target session is incompatible with the current session."}
	FrsErrInvalidApiSequence                               = &Error{0x00001F41, "FRS_ERR_INVALID_API_SEQUENCE", "The file replication service API was called incorrectly."}
	FrsErrStartingService                                  = &Error{0x00001F42, "FRS_ERR_STARTING_SERVICE", "The file replication service cannot be started."}
	FrsErrStoppingService                                  = &Error{0x00001F43, "FRS_ERR_STOPPING_SERVICE", "The file replication service cannot be stopped."}
	FrsErrInternalApi                                      = &Error{0x00001F44, "FRS_ERR_INTERNAL_API", "The file replication service API terminated the request. The event log might contain more information."}
	FrsErrInternal                                         = &Error{0x00001F45, "FRS_ERR_INTERNAL", "The file replication service terminated the request. The event log might contain more information."}
	FrsErrServiceComm                                      = &Error{0x00001F46, "FRS_ERR_SERVICE_COMM", "The file replication service cannot be contacted. The event log might contain more information."}
	FrsErrInsufficientPriv                                 = &Error{0x00001F47, "FRS_ERR_INSUFFICIENT_PRIV", "The file replication service cannot satisfy the request because the user has insufficient privileges. The event log might contain more information."}
	FrsErrAuthentication                                   = &Error{0x00001F48, "FRS_ERR_AUTHENTICATION", "The file replication service cannot satisfy the request because authenticated RPC is not available. The event log might contain more information."}
	FrsErrParentInsufficientPriv                           = &Error{0x00001F49, "FRS_ERR_PARENT_INSUFFICIENT_PRIV", "The file replication service cannot satisfy the request because the user has insufficient privileges on the domain controller. The event log might contain more information."}
	FrsErrParentAuthentication                             = &Error{0x00001F4A, "FRS_ERR_PARENT_AUTHENTICATION", "The file replication service cannot satisfy the request because authenticated RPC is not available on the domain controller. The event log might contain more information."}
	FrsErrChildToParentComm                                = &Error{0x00001F4B, "FRS_ERR_CHILD_TO_PARENT_COMM", "The file replication service cannot communicate with the file replication service on the domain controller. The event log might contain more information."}
	FrsErrParentToChildComm                                = &Error{0x00001F4C, "FRS_ERR_PARENT_TO_CHILD_COMM", "The file replication service on the domain controller cannot communicate with the file replication service on this computer. The event log might contain more information."}
	FrsErrSysvolPopulate                                   = &Error{0x00001F4D, "FRS_ERR_SYSVOL_POPULATE", "The file replication service cannot populate the system volume because of an internal error. The event log might contain more information."}
	FrsErrSysvolPopulateTimeout                            = &Error{0x00001F4E, "FRS_ERR_SYSVOL_POPULATE_TIMEOUT", "The file replication service cannot populate the system volume because of an internal time-out. The event log might contain more information."}
	FrsErrSysvolIsBusy                                     = &Error{0x00001F4F, "FRS_ERR_SYSVOL_IS_BUSY", "The file replication service cannot process the request. The system volume is busy with a previous request."}
	FrsErrSysvolDemote                                     = &Error{0x00001F50, "FRS_ERR_SYSVOL_DEMOTE", "The file replication service cannot stop replicating the system volume because of an internal error. The event log might contain more information."}
	FrsErrInvalidServiceParameter                          = &Error{0x00001F51, "FRS_ERR_INVALID_SERVICE_PARAMETER", "The file replication service detected an invalid parameter."}
	ErrorDsNotInstalled                                    = &Error{0x00002008, "ERROR_DS_NOT_INSTALLED", "An error occurred while installing the directory service. For more information, see the event log."}
	ErrorDsMembershipEvaluatedLocally                      = &Error{0x00002009, "ERROR_DS_MEMBERSHIP_EVALUATED_LOCALLY", "The directory service evaluated group memberships locally."}
	ErrorDsNoAttributeOrValue                              = &Error{0x0000200A, "ERROR_DS_NO_ATTRIBUTE_OR_VALUE", "The specified directory service attribute or value does not exist."}
	ErrorDsInvalidAttributeSyntax                          = &Error{0x0000200B, "ERROR_DS_INVALID_ATTRIBUTE_SYNTAX", "The attribute syntax specified to the directory service is invalid."}
	ErrorDsAttributeTypeUndefined                          = &Error{0x0000200C, "ERROR_DS_ATTRIBUTE_TYPE_UNDEFINED", "The attribute type specified to the directory service is not defined."}
	ErrorDsAttributeOrValueExists                          = &Error{0x0000200D, "ERROR_DS_ATTRIBUTE_OR_VALUE_EXISTS", "The specified directory service attribute or value already exists."}
	ErrorDsBusy                                            = &Error{0x0000200E, "ERROR_DS_BUSY", "The directory service is busy."}
	ErrorDsUnavailable                                     = &Error{0x0000200F, "ERROR_DS_UNAVAILABLE", "The directory service is unavailable."}
	ErrorDsNoRidsAllocated                                 = &Error{0x00002010, "ERROR_DS_NO_RIDS_ALLOCATED", "The directory service was unable to allocate a relative identifier."}
	ErrorDsNoMoreRids                                      = &Error{0x00002011, "ERROR_DS_NO_MORE_RIDS", "The directory service has exhausted the pool of relative identifiers."}
	ErrorDsIncorrectRoleOwner                              = &Error{0x00002012, "ERROR_DS_INCORRECT_ROLE_OWNER", "The requested operation could not be performed because the directory service is not the master for that type of operation."}
	ErrorDsRidmgrInitError                                 = &Error{0x00002013, "ERROR_DS_RIDMGR_INIT_ERROR", "The directory service was unable to initialize the subsystem that allocates relative identifiers."}
	ErrorDsObjClassViolation                               = &Error{0x00002014, "ERROR_DS_OBJ_CLASS_VIOLATION", "The requested operation did not satisfy one or more constraints associated with the class of the object."}
	ErrorDsCantOnNonLeaf                                   = &Error{0x00002015, "ERROR_DS_CANT_ON_NON_LEAF", "The directory service can perform the requested operation only on a leaf object."}
	ErrorDsCantOnRdn                                       = &Error{0x00002016, "ERROR_DS_CANT_ON_RDN", "The directory service cannot perform the requested operation on the relative distinguished name (RDN) attribute of an object."}
	ErrorDsCantModObjClass                                 = &Error{0x00002017, "ERROR_DS_CANT_MOD_OBJ_CLASS", "The directory service detected an attempt to modify the object class of an object."}
	ErrorDsCrossDomMoveError                               = &Error{0x00002018, "ERROR_DS_CROSS_DOM_MOVE_ERROR", "The requested cross-domain move operation could not be performed."}
	ErrorDsGcNotAvailable                                  = &Error{0x00002019, "ERROR_DS_GC_NOT_AVAILABLE", "Unable to contact the global catalog (GC) server."}
	ErrorSharedPolicy                                      = &Error{0x0000201A, "ERROR_SHARED_POLICY", "The policy object is shared and can only be modified at the root."}
	ErrorPolicyObjectNotFound                              = &Error{0x0000201B, "ERROR_POLICY_OBJECT_NOT_FOUND", "The policy object does not exist."}
	ErrorPolicyOnlyInDs                                    = &Error{0x0000201C, "ERROR_POLICY_ONLY_IN_DS", "The requested policy information is only in the directory service."}
	ErrorPromotionActive                                   = &Error{0x0000201D, "ERROR_PROMOTION_ACTIVE", "A domain controller promotion is currently active."}
	ErrorNoPromotionActive                                 = &Error{0x0000201E, "ERROR_NO_PROMOTION_ACTIVE", "A domain controller promotion is not currently active."}
	ErrorDsOperationsError                                 = &Error{0x00002020, "ERROR_DS_OPERATIONS_ERROR", "An operations error occurred."}
	ErrorDsProtocolError                                   = &Error{0x00002021, "ERROR_DS_PROTOCOL_ERROR", "A protocol error occurred."}
	ErrorDsTimelimitExceeded                               = &Error{0x00002022, "ERROR_DS_TIMELIMIT_EXCEEDED", "The time limit for this request was exceeded."}
	ErrorDsSizelimitExceeded                               = &Error{0x00002023, "ERROR_DS_SIZELIMIT_EXCEEDED", "The size limit for this request was exceeded."}
	ErrorDsAdminLimitExceeded                              = &Error{0x00002024, "ERROR_DS_ADMIN_LIMIT_EXCEEDED", "The administrative limit for this request was exceeded."}
	ErrorDsCompareFalse                                    = &Error{0x00002025, "ERROR_DS_COMPARE_FALSE", "The compare response was false."}
	ErrorDsCompareTrue                                     = &Error{0x00002026, "ERROR_DS_COMPARE_TRUE", "The compare response was true."}
	ErrorDsAuthMethodNotSupported                          = &Error{0x00002027, "ERROR_DS_AUTH_METHOD_NOT_SUPPORTED", "The requested authentication method is not supported by the server."}
	ErrorDsStrongAuthRequired                              = &Error{0x00002028, "ERROR_DS_STRONG_AUTH_REQUIRED", "A more secure authentication method is required for this server."}
	ErrorDsInappropriateAuth                               = &Error{0x00002029, "ERROR_DS_INAPPROPRIATE_AUTH", "Inappropriate authentication."}
	ErrorDsAuthUnknown                                     = &Error{0x0000202A, "ERROR_DS_AUTH_UNKNOWN", "The authentication mechanism is unknown."}
	ErrorDsReferral                                        = &Error{0x0000202B, "ERROR_DS_REFERRAL", "A referral was returned from the server."}
	ErrorDsUnavailableCritExtension                        = &Error{0x0000202C, "ERROR_DS_UNAVAILABLE_CRIT_EXTENSION", "The server does not support the requested critical extension."}
	ErrorDsConfidentialityRequired                         = &Error{0x0000202D, "ERROR_DS_CONFIDENTIALITY_REQUIRED", "This request requires a secure connection."}
	ErrorDsInappropriateMatching                           = &Error{0x0000202E, "ERROR_DS_INAPPROPRIATE_MATCHING", "Inappropriate matching."}
	ErrorDsConstraintViolation                             = &Error{0x0000202F, "ERROR_DS_CONSTRAINT_VIOLATION", "A constraint violation occurred."}
	ErrorDsNoSuchObject                                    = &Error{0x00002030, "ERROR_DS_NO_SUCH_OBJECT", "There is no such object on the server."}
	ErrorDsAliasProblem                                    = &Error{0x00002031, "ERROR_DS_ALIAS_PROBLEM", "There is an alias problem."}
	ErrorDsInvalidDnSyntax                                 = &Error{0x00002032, "ERROR_DS_INVALID_DN_SYNTAX", "An invalid dn syntax has been specified."}
	ErrorDsIsLeaf                                          = &Error{0x00002033, "ERROR_DS_IS_LEAF", "The object is a leaf object."}
	ErrorDsAliasDerefProblem                               = &Error{0x00002034, "ERROR_DS_ALIAS_DEREF_PROBLEM", "There is an alias dereferencing problem."}
	ErrorDsUnwillingToPerform                              = &Error{0x00002035, "ERROR_DS_UNWILLING_TO_PERFORM", "The server is unwilling to process the request."}
	ErrorDsLoopDetect                                      = &Error{0x00002036, "ERROR_DS_LOOP_DETECT", "A loop has been detected."}
	ErrorDsNamingViolation                                 = &Error{0x00002037, "ERROR_DS_NAMING_VIOLATION", "There is a naming violation."}
	ErrorDsObjectResultsTooLarge                           = &Error{0x00002038, "ERROR_DS_OBJECT_RESULTS_TOO_LARGE", "The result set is too large."}
	ErrorDsAffectsMultipleDsas                             = &Error{0x00002039, "ERROR_DS_AFFECTS_MULTIPLE_DSAS", "The operation affects multiple DSAs."}
	ErrorDsServerDown                                      = &Error{0x0000203A, "ERROR_DS_SERVER_DOWN", "The server is not operational."}
	ErrorDsLocalError                                      = &Error{0x0000203B, "ERROR_DS_LOCAL_ERROR", "A local error has occurred."}
	ErrorDsEncodingError                                   = &Error{0x0000203C, "ERROR_DS_ENCODING_ERROR", "An encoding error has occurred."}
	ErrorDsDecodingError                                   = &Error{0x0000203D, "ERROR_DS_DECODING_ERROR", "A decoding error has occurred."}
	ErrorDsFilterUnknown                                   = &Error{0x0000203E, "ERROR_DS_FILTER_UNKNOWN", "The search filter cannot be recognized."}
	ErrorDsParamError                                      = &Error{0x0000203F, "ERROR_DS_PARAM_ERROR", "One or more parameters are illegal."}
	ErrorDsNotSupported                                    = &Error{0x00002040, "ERROR_DS_NOT_SUPPORTED", "The specified method is not supported."}
	ErrorDsNoResultsReturned                               = &Error{0x00002041, "ERROR_DS_NO_RESULTS_RETURNED", "No results were returned."}
	ErrorDsControlNotFound                                 = &Error{0x00002042, "ERROR_DS_CONTROL_NOT_FOUND", "The specified control is not supported by the server."}
	ErrorDsClientLoop                                      = &Error{0x00002043, "ERROR_DS_CLIENT_LOOP", "A referral loop was detected by the client."}
	ErrorDsReferralLimitExceeded                           = &Error{0x00002044, "ERROR_DS_REFERRAL_LIMIT_EXCEEDED", "The preset referral limit was exceeded."}
	ErrorDsSortControlMissing                              = &Error{0x00002045, "ERROR_DS_SORT_CONTROL_MISSING", "The search requires a SORT control."}
	ErrorDsOffsetRangeError                                = &Error{0x00002046, "ERROR_DS_OFFSET_RANGE_ERROR", "The search results exceed the offset range specified."}
	ErrorDsRootMustBeNc                                    = &Error{0x0000206D, "ERROR_DS_ROOT_MUST_BE_NC", "The root object must be the head of a naming context. The root object cannot have an instantiated parent."}
	ErrorDsAddReplicaInhibited                             = &Error{0x0000206E, "ERROR_DS_ADD_REPLICA_INHIBITED", "The add replica operation cannot be performed. The naming context must be writable to create the replica."}
	ErrorDsAttNotDefInSchema                               = &Error{0x0000206F, "ERROR_DS_ATT_NOT_DEF_IN_SCHEMA", "A reference to an attribute that is not defined in the schema occurred."}
	ErrorDsMaxObjSizeExceeded                              = &Error{0x00002070, "ERROR_DS_MAX_OBJ_SIZE_EXCEEDED", "The maximum size of an object has been exceeded."}
	ErrorDsObjStringNameExists                             = &Error{0x00002071, "ERROR_DS_OBJ_STRING_NAME_EXISTS", "An attempt was made to add an object to the directory with a name that is already in use."}
	ErrorDsNoRdnDefinedInSchema                            = &Error{0x00002072, "ERROR_DS_NO_RDN_DEFINED_IN_SCHEMA", "An attempt was made to add an object of a class that does not have an RDN defined in the schema."}
	ErrorDsRdnDoesntMatchSchema                            = &Error{0x00002073, "ERROR_DS_RDN_DOESNT_MATCH_SCHEMA", "An attempt was made to add an object using an RDN that is not the RDN defined in the schema."}
	ErrorDsNoRequestedAttsFound                            = &Error{0x00002074, "ERROR_DS_NO_REQUESTED_ATTS_FOUND", "None of the requested attributes were found on the objects."}
	ErrorDsUserBufferToSmall                               = &Error{0x00002075, "ERROR_DS_USER_BUFFER_TO_SMALL", "The user buffer is too small."}
	ErrorDsAttIsNotOnObj                                   = &Error{0x00002076, "ERROR_DS_ATT_IS_NOT_ON_OBJ", "The attribute specified in the operation is not present on the object."}
	ErrorDsIllegalModOperation                             = &Error{0x00002077, "ERROR_DS_ILLEGAL_MOD_OPERATION", "Illegal modify operation. Some aspect of the modification is not permitted."}
	ErrorDsObjTooLarge                                     = &Error{0x00002078, "ERROR_DS_OBJ_TOO_LARGE", "The specified object is too large."}
	ErrorDsBadInstanceType                                 = &Error{0x00002079, "ERROR_DS_BAD_INSTANCE_TYPE", "The specified instance type is not valid."}
	ErrorDsMasterdsaRequired                               = &Error{0x0000207A, "ERROR_DS_MASTERDSA_REQUIRED", "The operation must be performed at a master DSA."}
	ErrorDsObjectClassRequired                             = &Error{0x0000207B, "ERROR_DS_OBJECT_CLASS_REQUIRED", "The object class attribute must be specified."}
	ErrorDsMissingRequiredAtt                              = &Error{0x0000207C, "ERROR_DS_MISSING_REQUIRED_ATT", "A required attribute is missing."}
	ErrorDsAttNotDefForClass                               = &Error{0x0000207D, "ERROR_DS_ATT_NOT_DEF_FOR_CLASS", "An attempt was made to modify an object to include an attribute that is not legal for its class."}
	ErrorDsAttAlreadyExists                                = &Error{0x0000207E, "ERROR_DS_ATT_ALREADY_EXISTS", "The specified attribute is already present on the object."}
	ErrorDsCantAddAttValues                                = &Error{0x00002080, "ERROR_DS_CANT_ADD_ATT_VALUES", "The specified attribute is not present, or has no values."}
	ErrorDsSingleValueConstraint                           = &Error{0x00002081, "ERROR_DS_SINGLE_VALUE_CONSTRAINT", "Multiple values were specified for an attribute that can have only one value."}
	ErrorDsRangeConstraint                                 = &Error{0x00002082, "ERROR_DS_RANGE_CONSTRAINT", "A value for the attribute was not in the acceptable range of values."}
	ErrorDsAttValAlreadyExists                             = &Error{0x00002083, "ERROR_DS_ATT_VAL_ALREADY_EXISTS", "The specified value already exists."}
	ErrorDsCantRemMissingAtt                               = &Error{0x00002084, "ERROR_DS_CANT_REM_MISSING_ATT", "The attribute cannot be removed because it is not present on the object."}
	ErrorDsCantRemMissingAttVal                            = &Error{0x00002085, "ERROR_DS_CANT_REM_MISSING_ATT_VAL", "The attribute value cannot be removed because it is not present on the object."}
	ErrorDsRootCantBeSubref                                = &Error{0x00002086, "ERROR_DS_ROOT_CANT_BE_SUBREF", "The specified root object cannot be a subreference."}
	ErrorDsNoChaining                                      = &Error{0x00002087, "ERROR_DS_NO_CHAINING", "Chaining is not permitted."}
	ErrorDsNoChainedEval                                   = &Error{0x00002088, "ERROR_DS_NO_CHAINED_EVAL", "Chained evaluation is not permitted."}
	ErrorDsNoParentObject                                  = &Error{0x00002089, "ERROR_DS_NO_PARENT_OBJECT", "The operation could not be performed because the object's parent is either uninstantiated or deleted."}
	ErrorDsParentIsAnAlias                                 = &Error{0x0000208A, "ERROR_DS_PARENT_IS_AN_ALIAS", "Having a parent that is an alias is not permitted. Aliases are leaf objects."}
	ErrorDsCantMixMasterAndReps                            = &Error{0x0000208B, "ERROR_DS_CANT_MIX_MASTER_AND_REPS", "The object and parent must be of the same type, either both masters or both replicas."}
	ErrorDsChildrenExist                                   = &Error{0x0000208C, "ERROR_DS_CHILDREN_EXIST", "The operation cannot be performed because child objects exist. This operation can only be performed on a leaf object."}
	ErrorDsObjNotFound                                     = &Error{0x0000208D, "ERROR_DS_OBJ_NOT_FOUND", "Directory object not found."}
	ErrorDsAliasedObjMissing                               = &Error{0x0000208E, "ERROR_DS_ALIASED_OBJ_MISSING", "The aliased object is missing."}
	ErrorDsBadNameSyntax                                   = &Error{0x0000208F, "ERROR_DS_BAD_NAME_SYNTAX", "The object name has bad syntax."}
	ErrorDsAliasPointsToAlias                              = &Error{0x00002090, "ERROR_DS_ALIAS_POINTS_TO_ALIAS", "An alias is not permitted to refer to another alias."}
	ErrorDsCantDerefAlias                                  = &Error{0x00002091, "ERROR_DS_CANT_DEREF_ALIAS", "The alias cannot be dereferenced."}
	ErrorDsOutOfScope                                      = &Error{0x00002092, "ERROR_DS_OUT_OF_SCOPE", "The operation is out of scope."}
	ErrorDsObjectBeingRemoved                              = &Error{0x00002093, "ERROR_DS_OBJECT_BEING_REMOVED", "The operation cannot continue because the object is in the process of being removed."}
	ErrorDsCantDeleteDsaObj                                = &Error{0x00002094, "ERROR_DS_CANT_DELETE_DSA_OBJ", "The DSA object cannot be deleted."}
	ErrorDsGenericError                                    = &Error{0x00002095, "ERROR_DS_GENERIC_ERROR", "A directory service error has occurred."}
	ErrorDsDsaMustBeIntMaster                              = &Error{0x00002096, "ERROR_DS_DSA_MUST_BE_INT_MASTER", "The operation can only be performed on an internal master DSA object."}
	ErrorDsClassNotDsa                                     = &Error{0x00002097, "ERROR_DS_CLASS_NOT_DSA", "The object must be of class DSA."}
	ErrorDsInsuffAccessRights                              = &Error{0x00002098, "ERROR_DS_INSUFF_ACCESS_RIGHTS", "Insufficient access rights to perform the operation."}
	ErrorDsIllegalSuperior                                 = &Error{0x00002099, "ERROR_DS_ILLEGAL_SUPERIOR", "The object cannot be added because the parent is not on the list of possible superiors."}
	ErrorDsAttributeOwnedBySam                             = &Error{0x0000209A, "ERROR_DS_ATTRIBUTE_OWNED_BY_SAM", "Access to the attribute is not permitted because the attribute is owned by the SAM."}
	ErrorDsNameTooManyParts                                = &Error{0x0000209B, "ERROR_DS_NAME_TOO_MANY_PARTS", "The name has too many parts."}
	ErrorDsNameTooLong                                     = &Error{0x0000209C, "ERROR_DS_NAME_TOO_LONG", "The name is too long."}
	ErrorDsNameValueTooLong                                = &Error{0x0000209D, "ERROR_DS_NAME_VALUE_TOO_LONG", "The name value is too long."}
	ErrorDsNameUnparseable                                 = &Error{0x0000209E, "ERROR_DS_NAME_UNPARSEABLE", "The directory service encountered an error parsing a name."}
	ErrorDsNameTypeUnknown                                 = &Error{0x0000209F, "ERROR_DS_NAME_TYPE_UNKNOWN", "The directory service cannot get the attribute type for a name."}
	ErrorDsNotAnObject                                     = &Error{0x000020A0, "ERROR_DS_NOT_AN_OBJECT", "The name does not identify an object; the name identifies a phantom."}
	ErrorDsSecDescTooShort                                 = &Error{0x000020A1, "ERROR_DS_SEC_DESC_TOO_SHORT", "The security descriptor is too short."}
	ErrorDsSecDescInvalid                                  = &Error{0x000020A2, "ERROR_DS_SEC_DESC_INVALID", "The security descriptor is invalid."}
	ErrorDsNoDeletedName                                   = &Error{0x000020A3, "ERROR_DS_NO_DELETED_NAME", "Failed to create name for deleted object."}
	ErrorDsSubrefMustHaveParent                            = &Error{0x000020A4, "ERROR_DS_SUBREF_MUST_HAVE_PARENT", "The parent of a new subreference must exist."}
	ErrorDsNcnameMustBeNc                                  = &Error{0x000020A5, "ERROR_DS_NCNAME_MUST_BE_NC", "The object must be a naming context."}
	ErrorDsCantAddSystemOnly                               = &Error{0x000020A6, "ERROR_DS_CANT_ADD_SYSTEM_ONLY", "It is not permitted to add an attribute that is owned by the system."}
	ErrorDsClassMustBeConcrete                             = &Error{0x000020A7, "ERROR_DS_CLASS_MUST_BE_CONCRETE", "The class of the object must be structural; you cannot instantiate an abstract class."}
	ErrorDsInvalidDmd                                      = &Error{0x000020A8, "ERROR_DS_INVALID_DMD", "The schema object could not be found."}
	ErrorDsObjGuidExists                                   = &Error{0x000020A9, "ERROR_DS_OBJ_GUID_EXISTS", "A local object with this GUID (dead or alive) already exists."}
	ErrorDsNotOnBacklink                                   = &Error{0x000020AA, "ERROR_DS_NOT_ON_BACKLINK", "The operation cannot be performed on a back link."}
	ErrorDsNoCrossrefForNc                                 = &Error{0x000020AB, "ERROR_DS_NO_CROSSREF_FOR_NC", "The cross-reference for the specified naming context could not be found."}
	ErrorDsShuttingDown                                    = &Error{0x000020AC, "ERROR_DS_SHUTTING_DOWN", "The operation could not be performed because the directory service is shutting down."}
	ErrorDsUnknownOperation                                = &Error{0x000020AD, "ERROR_DS_UNKNOWN_OPERATION", "The directory service request is invalid."}
	ErrorDsInvalidRoleOwner                                = &Error{0x000020AE, "ERROR_DS_INVALID_ROLE_OWNER", "The role owner attribute could not be read."}
	ErrorDsCouldntContactFsmo                              = &Error{0x000020AF, "ERROR_DS_COULDNT_CONTACT_FSMO", "The requested Flexible Single Master Operations (FSMO) operation failed. The current FSMO holder could not be contacted."}
	ErrorDsCrossNcDnRename                                 = &Error{0x000020B0, "ERROR_DS_CROSS_NC_DN_RENAME", "Modification of a distinguished name across a naming context is not permitted."}
	ErrorDsCantModSystemOnly                               = &Error{0x000020B1, "ERROR_DS_CANT_MOD_SYSTEM_ONLY", "The attribute cannot be modified because it is owned by the system."}
	ErrorDsReplicatorOnly                                  = &Error{0x000020B2, "ERROR_DS_REPLICATOR_ONLY", "Only the replicator can perform this function."}
	ErrorDsObjClassNotDefined                              = &Error{0x000020B3, "ERROR_DS_OBJ_CLASS_NOT_DEFINED", "The specified class is not defined."}
	ErrorDsObjClassNotSubclass                             = &Error{0x000020B4, "ERROR_DS_OBJ_CLASS_NOT_SUBCLASS", "The specified class is not a subclass."}
	ErrorDsNameReferenceInvalid                            = &Error{0x000020B5, "ERROR_DS_NAME_REFERENCE_INVALID", "The name reference is invalid."}
	ErrorDsCrossRefExists                                  = &Error{0x000020B6, "ERROR_DS_CROSS_REF_EXISTS", "A cross-reference already exists."}
	ErrorDsCantDelMasterCrossref                           = &Error{0x000020B7, "ERROR_DS_CANT_DEL_MASTER_CROSSREF", "It is not permitted to delete a master cross-reference."}
	ErrorDsSubtreeNotifyNotNcHead                          = &Error{0x000020B8, "ERROR_DS_SUBTREE_NOTIFY_NOT_NC_HEAD", "Subtree notifications are only supported on naming context (NC) heads."}
	ErrorDsNotifyFilterTooComplex                          = &Error{0x000020B9, "ERROR_DS_NOTIFY_FILTER_TOO_COMPLEX", "Notification filter is too complex."}
	ErrorDsDupRdn                                          = &Error{0x000020BA, "ERROR_DS_DUP_RDN", "Schema update failed: Duplicate RDN."}
	ErrorDsDupOid                                          = &Error{0x000020BB, "ERROR_DS_DUP_OID", "Schema update failed: Duplicate OID."}
	ErrorDsDupMapiId                                       = &Error{0x000020BC, "ERROR_DS_DUP_MAPI_ID", "Schema update failed: Duplicate Message Application Programming Interface (MAPI) identifier."}
	ErrorDsDupSchemaIdGuid                                 = &Error{0x000020BD, "ERROR_DS_DUP_SCHEMA_ID_GUID", "Schema update failed: Duplicate schema ID GUID."}
	ErrorDsDupLdapDisplayName                              = &Error{0x000020BE, "ERROR_DS_DUP_LDAP_DISPLAY_NAME", "Schema update failed: Duplicate LDAP display name."}
	ErrorDsSemanticAttTest                                 = &Error{0x000020BF, "ERROR_DS_SEMANTIC_ATT_TEST", "Schema update failed: Range-Lower less than Range-Upper."}
	ErrorDsSyntaxMismatch                                  = &Error{0x000020C0, "ERROR_DS_SYNTAX_MISMATCH", "Schema update failed: Syntax mismatch."}
	ErrorDsExistsInMustHave                                = &Error{0x000020C1, "ERROR_DS_EXISTS_IN_MUST_HAVE", "Schema deletion failed: Attribute is used in the Must-Contain list."}
	ErrorDsExistsInMayHave                                 = &Error{0x000020C2, "ERROR_DS_EXISTS_IN_MAY_HAVE", "Schema deletion failed: Attribute is used in the May-Contain list."}
	ErrorDsNonexistentMayHave                              = &Error{0x000020C3, "ERROR_DS_NONEXISTENT_MAY_HAVE", "Schema update failed: Attribute in May-Contain list does not exist."}
	ErrorDsNonexistentMustHave                             = &Error{0x000020C4, "ERROR_DS_NONEXISTENT_MUST_HAVE", "Schema update failed: Attribute in the Must-Contain list does not exist."}
	ErrorDsAuxClsTestFail                                  = &Error{0x000020C5, "ERROR_DS_AUX_CLS_TEST_FAIL", "Schema update failed: Class in the Aux Class list does not exist or is not an auxiliary class."}
	ErrorDsNonexistentPossSup                              = &Error{0x000020C6, "ERROR_DS_NONEXISTENT_POSS_SUP", "Schema update failed: Class in the Poss-Superiors list does not exist."}
	ErrorDsSubClsTestFail                                  = &Error{0x000020C7, "ERROR_DS_SUB_CLS_TEST_FAIL", "Schema update failed: Class in the subclass of the list does not exist or does not satisfy hierarchy rules."}
	ErrorDsBadRdnAttIdSyntax                               = &Error{0x000020C8, "ERROR_DS_BAD_RDN_ATT_ID_SYNTAX", "Schema update failed: Rdn-Att-Id has wrong syntax."}
	ErrorDsExistsInAuxCls                                  = &Error{0x000020C9, "ERROR_DS_EXISTS_IN_AUX_CLS", "Schema deletion failed: Class is used as an auxiliary class."}
	ErrorDsExistsInSubCls                                  = &Error{0x000020CA, "ERROR_DS_EXISTS_IN_SUB_CLS", "Schema deletion failed: Class is used as a subclass."}
	ErrorDsExistsInPossSup                                 = &Error{0x000020CB, "ERROR_DS_EXISTS_IN_POSS_SUP", "Schema deletion failed: Class is used as a Poss-Superior."}
	ErrorDsRecalcschemaFailed                              = &Error{0x000020CC, "ERROR_DS_RECALCSCHEMA_FAILED", "Schema update failed in recalculating validation cache."}
	ErrorDsTreeDeleteNotFinished                           = &Error{0x000020CD, "ERROR_DS_TREE_DELETE_NOT_FINISHED", "The tree deletion is not finished. The request must be made again to continue deleting the tree."}
	ErrorDsCantDelete                                      = &Error{0x000020CE, "ERROR_DS_CANT_DELETE", "The requested delete operation could not be performed."}
	ErrorDsAttSchemaReqId                                  = &Error{0x000020CF, "ERROR_DS_ATT_SCHEMA_REQ_ID", "Cannot read the governs class identifier for the schema record."}
	ErrorDsBadAttSchemaSyntax                              = &Error{0x000020D0, "ERROR_DS_BAD_ATT_SCHEMA_SYNTAX", "The attribute schema has bad syntax."}
	ErrorDsCantCacheAtt                                    = &Error{0x000020D1, "ERROR_DS_CANT_CACHE_ATT", "The attribute could not be cached."}
	ErrorDsCantCacheClass                                  = &Error{0x000020D2, "ERROR_DS_CANT_CACHE_CLASS", "The class could not be cached."}
	ErrorDsCantRemoveAttCache                              = &Error{0x000020D3, "ERROR_DS_CANT_REMOVE_ATT_CACHE", "The attribute could not be removed from the cache."}
	ErrorDsCantRemoveClassCache                            = &Error{0x000020D4, "ERROR_DS_CANT_REMOVE_CLASS_CACHE", "The class could not be removed from the cache."}
	ErrorDsCantRetrieveDn                                  = &Error{0x000020D5, "ERROR_DS_CANT_RETRIEVE_DN", "The distinguished name attribute could not be read."}
	ErrorDsMissingSupref                                   = &Error{0x000020D6, "ERROR_DS_MISSING_SUPREF", "No superior reference has been configured for the directory service. The directory service is, therefore, unable to issue referrals to objects outside this forest."}
	ErrorDsCantRetrieveInstance                            = &Error{0x000020D7, "ERROR_DS_CANT_RETRIEVE_INSTANCE", "The instance type attribute could not be retrieved."}
	ErrorDsCodeInconsistency                               = &Error{0x000020D8, "ERROR_DS_CODE_INCONSISTENCY", "An internal error has occurred."}
	ErrorDsDatabaseError                                   = &Error{0x000020D9, "ERROR_DS_DATABASE_ERROR", "A database error has occurred."}
	ErrorDsGovernsidMissing                                = &Error{0x000020DA, "ERROR_DS_GOVERNSID_MISSING", "The governsID attribute is missing."}
	ErrorDsMissingExpectedAtt                              = &Error{0x000020DB, "ERROR_DS_MISSING_EXPECTED_ATT", "An expected attribute is missing."}
	ErrorDsNcnameMissingCrRef                              = &Error{0x000020DC, "ERROR_DS_NCNAME_MISSING_CR_REF", "The specified naming context is missing a cross-reference."}
	ErrorDsSecurityCheckingError                           = &Error{0x000020DD, "ERROR_DS_SECURITY_CHECKING_ERROR", "A security checking error has occurred."}
	ErrorDsSchemaNotLoaded                                 = &Error{0x000020DE, "ERROR_DS_SCHEMA_NOT_LOADED", "The schema is not loaded."}
	ErrorDsSchemaAllocFailed                               = &Error{0x000020DF, "ERROR_DS_SCHEMA_ALLOC_FAILED", "Schema allocation failed. Check if the machine is running low on memory."}
	ErrorDsAttSchemaReqSyntax                              = &Error{0x000020E0, "ERROR_DS_ATT_SCHEMA_REQ_SYNTAX", "Failed to obtain the required syntax for the attribute schema."}
	ErrorDsGcverifyError                                   = &Error{0x000020E1, "ERROR_DS_GCVERIFY_ERROR", "The GC verification failed. The GC is not available or does not support the operation. Some part of the directory is currently not available."}
	ErrorDsDraSchemaMismatch                               = &Error{0x000020E2, "ERROR_DS_DRA_SCHEMA_MISMATCH", "The replication operation failed because of a schema mismatch between the servers involved."}
	ErrorDsCantFindDsaObj                                  = &Error{0x000020E3, "ERROR_DS_CANT_FIND_DSA_OBJ", "The DSA object could not be found."}
	ErrorDsCantFindExpectedNc                              = &Error{0x000020E4, "ERROR_DS_CANT_FIND_EXPECTED_NC", "The naming context could not be found."}
	ErrorDsCantFindNcInCache                               = &Error{0x000020E5, "ERROR_DS_CANT_FIND_NC_IN_CACHE", "The naming context could not be found in the cache."}
	ErrorDsCantRetrieveChild                               = &Error{0x000020E6, "ERROR_DS_CANT_RETRIEVE_CHILD", "The child object could not be retrieved."}
	ErrorDsSecurityIllegalModify                           = &Error{0x000020E7, "ERROR_DS_SECURITY_ILLEGAL_MODIFY", "The modification was not permitted for security reasons."}
	ErrorDsCantReplaceHiddenRec                            = &Error{0x000020E8, "ERROR_DS_CANT_REPLACE_HIDDEN_REC", "The operation cannot replace the hidden record."}
	ErrorDsBadHierarchyFile                                = &Error{0x000020E9, "ERROR_DS_BAD_HIERARCHY_FILE", "The hierarchy file is invalid."}
	ErrorDsBuildHierarchyTableFailed                       = &Error{0x000020EA, "ERROR_DS_BUILD_HIERARCHY_TABLE_FAILED", "The attempt to build the hierarchy table failed."}
	ErrorDsConfigParamMissing                              = &Error{0x000020EB, "ERROR_DS_CONFIG_PARAM_MISSING", "The directory configuration parameter is missing from the registry."}
	ErrorDsCountingAbIndicesFailed                         = &Error{0x000020EC, "ERROR_DS_COUNTING_AB_INDICES_FAILED", "The attempt to count the address book indices failed."}
	ErrorDsHierarchyTableMallocFailed                      = &Error{0x000020ED, "ERROR_DS_HIERARCHY_TABLE_MALLOC_FAILED", "The allocation of the hierarchy table failed."}
	ErrorDsInternalFailure                                 = &Error{0x000020EE, "ERROR_DS_INTERNAL_FAILURE", "The directory service encountered an internal failure."}
	ErrorDsUnknownError                                    = &Error{0x000020EF, "ERROR_DS_UNKNOWN_ERROR", "The directory service encountered an unknown failure."}
	ErrorDsRootRequiresClassTop                            = &Error{0x000020F0, "ERROR_DS_ROOT_REQUIRES_CLASS_TOP", "A root object requires a class of \"top\"."}
	ErrorDsRefusingFsmoRoles                               = &Error{0x000020F1, "ERROR_DS_REFUSING_FSMO_ROLES", "This directory server is shutting down, and cannot take ownership of new floating single-master operation roles."}
	ErrorDsMissingFsmoSettings                             = &Error{0x000020F2, "ERROR_DS_MISSING_FSMO_SETTINGS", "The directory service is missing mandatory configuration information and is unable to determine the ownership of floating single-master operation roles."}
	ErrorDsUnableToSurrenderRoles                          = &Error{0x000020F3, "ERROR_DS_UNABLE_TO_SURRENDER_ROLES", "The directory service was unable to transfer ownership of one or more floating single-master operation roles to other servers."}
	ErrorDsDraGeneric                                      = &Error{0x000020F4, "ERROR_DS_DRA_GENERIC", "The replication operation failed."}
	ErrorDsDraInvalidParameter                             = &Error{0x000020F5, "ERROR_DS_DRA_INVALID_PARAMETER", "An invalid parameter was specified for this replication operation."}
	ErrorDsDraBusy                                         = &Error{0x000020F6, "ERROR_DS_DRA_BUSY", "The directory service is too busy to complete the replication operation at this time."}
	ErrorDsDraBadDn                                        = &Error{0x000020F7, "ERROR_DS_DRA_BAD_DN", "The DN specified for this replication operation is invalid."}
	ErrorDsDraBadNc                                        = &Error{0x000020F8, "ERROR_DS_DRA_BAD_NC", "The naming context specified for this replication operation is invalid."}
	ErrorDsDraDnExists                                     = &Error{0x000020F9, "ERROR_DS_DRA_DN_EXISTS", "The DN specified for this replication operation already exists."}
	ErrorDsDraInternalError                                = &Error{0x000020FA, "ERROR_DS_DRA_INTERNAL_ERROR", "The replication system encountered an internal error."}
	ErrorDsDraInconsistentDit                              = &Error{0x000020FB, "ERROR_DS_DRA_INCONSISTENT_DIT", "The replication operation encountered a database inconsistency."}
	ErrorDsDraConnectionFailed                             = &Error{0x000020FC, "ERROR_DS_DRA_CONNECTION_FAILED", "The server specified for this replication operation could not be contacted."}
	ErrorDsDraBadInstanceType                              = &Error{0x000020FD, "ERROR_DS_DRA_BAD_INSTANCE_TYPE", "The replication operation encountered an object with an invalid instance type."}
	ErrorDsDraOutOfMem                                     = &Error{0x000020FE, "ERROR_DS_DRA_OUT_OF_MEM", "The replication operation failed to allocate memory."}
	ErrorDsDraMailProblem                                  = &Error{0x000020FF, "ERROR_DS_DRA_MAIL_PROBLEM", "The replication operation encountered an error with the mail system."}
	ErrorDsDraRefAlreadyExists                             = &Error{0x00002100, "ERROR_DS_DRA_REF_ALREADY_EXISTS", "The replication reference information for the target server already exists."}
	ErrorDsDraRefNotFound                                  = &Error{0x00002101, "ERROR_DS_DRA_REF_NOT_FOUND", "The replication reference information for the target server does not exist."}
	ErrorDsDraObjIsRepSource                               = &Error{0x00002102, "ERROR_DS_DRA_OBJ_IS_REP_SOURCE", "The naming context cannot be removed because it is replicated to another server."}
	ErrorDsDraDbError                                      = &Error{0x00002103, "ERROR_DS_DRA_DB_ERROR", "The replication operation encountered a database error."}
	ErrorDsDraNoReplica                                    = &Error{0x00002104, "ERROR_DS_DRA_NO_REPLICA", "The naming context is in the process of being removed or is not replicated from the specified server."}
	ErrorDsDraAccessDenied                                 = &Error{0x00002105, "ERROR_DS_DRA_ACCESS_DENIED", "Replication access was denied."}
	ErrorDsDraNotSupported                                 = &Error{0x00002106, "ERROR_DS_DRA_NOT_SUPPORTED", "The requested operation is not supported by this version of the directory service."}
	ErrorDsDraRpcCancelled                                 = &Error{0x00002107, "ERROR_DS_DRA_RPC_CANCELLED", "The replication RPC was canceled."}
	ErrorDsDraSourceDisabled                               = &Error{0x00002108, "ERROR_DS_DRA_SOURCE_DISABLED", "The source server is currently rejecting replication requests."}
	ErrorDsDraSinkDisabled                                 = &Error{0x00002109, "ERROR_DS_DRA_SINK_DISABLED", "The destination server is currently rejecting replication requests."}
	ErrorDsDraNameCollision                                = &Error{0x0000210A, "ERROR_DS_DRA_NAME_COLLISION", "The replication operation failed due to a collision of object names."}
	ErrorDsDraSourceReinstalled                            = &Error{0x0000210B, "ERROR_DS_DRA_SOURCE_REINSTALLED", "The replication source has been reinstalled."}
	ErrorDsDraMissingParent                                = &Error{0x0000210C, "ERROR_DS_DRA_MISSING_PARENT", "The replication operation failed because a required parent object is missing."}
	ErrorDsDraPreempted                                    = &Error{0x0000210D, "ERROR_DS_DRA_PREEMPTED", "The replication operation was preempted."}
	ErrorDsDraAbandonSync                                  = &Error{0x0000210E, "ERROR_DS_DRA_ABANDON_SYNC", "The replication synchronization attempt was abandoned because of a lack of updates."}
	ErrorDsDraShutdown                                     = &Error{0x0000210F, "ERROR_DS_DRA_SHUTDOWN", "The replication operation was terminated because the system is shutting down."}
	ErrorDsDraIncompatiblePartialSet                       = &Error{0x00002110, "ERROR_DS_DRA_INCOMPATIBLE_PARTIAL_SET", "A synchronization attempt failed because the destination DC is currently waiting to synchronize new partial attributes from the source. This condition is normal if a recent schema change modified the partial attribute set. The destination partial attribute set is not a subset of the source partial attribute set."}
	ErrorDsDraSourceIsPartialReplica                       = &Error{0x00002111, "ERROR_DS_DRA_SOURCE_IS_PARTIAL_REPLICA", "The replication synchronization attempt failed because a master replica attempted to sync from a partial replica."}
	ErrorDsDraExtnConnectionFailed                         = &Error{0x00002112, "ERROR_DS_DRA_EXTN_CONNECTION_FAILED", "The server specified for this replication operation was contacted, but that server was unable to contact an additional server needed to complete the operation."}
	ErrorDsInstallSchemaMismatch                           = &Error{0x00002113, "ERROR_DS_INSTALL_SCHEMA_MISMATCH", "The version of the directory service schema of the source forest is not compatible with the version of the directory service on this computer."}
	ErrorDsDupLinkId                                       = &Error{0x00002114, "ERROR_DS_DUP_LINK_ID", "Schema update failed: An attribute with the same link identifier already exists."}
	ErrorDsNameErrorResolving                              = &Error{0x00002115, "ERROR_DS_NAME_ERROR_RESOLVING", "Name translation: Generic processing error."}
	ErrorDsNameErrorNotFound                               = &Error{0x00002116, "ERROR_DS_NAME_ERROR_NOT_FOUND", "Name translation: Could not find the name or insufficient right to see name."}
	ErrorDsNameErrorNotUnique                              = &Error{0x00002117, "ERROR_DS_NAME_ERROR_NOT_UNIQUE", "Name translation: Input name mapped to more than one output name."}
	ErrorDsNameErrorNoMapping                              = &Error{0x00002118, "ERROR_DS_NAME_ERROR_NO_MAPPING", "Name translation: The input name was found but not the associated output format."}
	ErrorDsNameErrorDomainOnly                             = &Error{0x00002119, "ERROR_DS_NAME_ERROR_DOMAIN_ONLY", "Name translation: Unable to resolve completely, only the domain was found."}
	ErrorDsNameErrorNoSyntacticalMapping                   = &Error{0x0000211A, "ERROR_DS_NAME_ERROR_NO_SYNTACTICAL_MAPPING", "Name translation: Unable to perform purely syntactical mapping at the client without going out to the wire."}
	ErrorDsConstructedAttMod                               = &Error{0x0000211B, "ERROR_DS_CONSTRUCTED_ATT_MOD", "Modification of a constructed attribute is not allowed."}
	ErrorDsWrongOmObjClass                                 = &Error{0x0000211C, "ERROR_DS_WRONG_OM_OBJ_CLASS", "The OM-Object-Class specified is incorrect for an attribute with the specified syntax."}
	ErrorDsDraReplPending                                  = &Error{0x0000211D, "ERROR_DS_DRA_REPL_PENDING", "The replication request has been posted; waiting for a reply."}
	ErrorDsDsRequired                                      = &Error{0x0000211E, "ERROR_DS_DS_REQUIRED", "The requested operation requires a directory service, and none was available."}
	ErrorDsInvalidLdapDisplayName                          = &Error{0x0000211F, "ERROR_DS_INVALID_LDAP_DISPLAY_NAME", "The LDAP display name of the class or attribute contains non-ASCII characters."}
	ErrorDsNonBaseSearch                                   = &Error{0x00002120, "ERROR_DS_NON_BASE_SEARCH", "The requested search operation is only supported for base searches."}
	ErrorDsCantRetrieveAtts                                = &Error{0x00002121, "ERROR_DS_CANT_RETRIEVE_ATTS", "The search failed to retrieve attributes from the database."}
	ErrorDsBacklinkWithoutLink                             = &Error{0x00002122, "ERROR_DS_BACKLINK_WITHOUT_LINK", "The schema update operation tried to add a backward link attribute that has no corresponding forward link."}
	ErrorDsEpochMismatch                                   = &Error{0x00002123, "ERROR_DS_EPOCH_MISMATCH", "The source and destination of a cross-domain move do not agree on the object's epoch number. Either the source or the destination does not have the latest version of the object."}
	ErrorDsSrcNameMismatch                                 = &Error{0x00002124, "ERROR_DS_SRC_NAME_MISMATCH", "The source and destination of a cross-domain move do not agree on the object's current name. Either the source or the destination does not have the latest version of the object."}
	ErrorDsSrcAndDstNcIdentical                            = &Error{0x00002125, "ERROR_DS_SRC_AND_DST_NC_IDENTICAL", "The source and destination for the cross-domain move operation are identical. The caller should use a local move operation instead of a cross-domain move operation."}
	ErrorDsDstNcMismatch                                   = &Error{0x00002126, "ERROR_DS_DST_NC_MISMATCH", "The source and destination for a cross-domain move do not agree on the naming contexts in the forest. Either the source or the destination does not have the latest version of the Partitions container."}
	ErrorDsNotAuthoritiveForDstNc                          = &Error{0x00002127, "ERROR_DS_NOT_AUTHORITIVE_FOR_DST_NC", "The destination of a cross-domain move is not authoritative for the destination naming context."}
	ErrorDsSrcGuidMismatch                                 = &Error{0x00002128, "ERROR_DS_SRC_GUID_MISMATCH", "The source and destination of a cross-domain move do not agree on the identity of the source object. Either the source or the destination does not have the latest version of the source object."}
	ErrorDsCantMoveDeletedObject                           = &Error{0x00002129, "ERROR_DS_CANT_MOVE_DELETED_OBJECT", "The object being moved across domains is already known to be deleted by the destination server. The source server does not have the latest version of the source object."}
	ErrorDsPdcOperationInProgress                          = &Error{0x0000212A, "ERROR_DS_PDC_OPERATION_IN_PROGRESS", "Another operation that requires exclusive access to the PDC FSMO is already in progress."}
	ErrorDsCrossDomainCleanupReqd                          = &Error{0x0000212B, "ERROR_DS_CROSS_DOMAIN_CLEANUP_REQD", "A cross-domain move operation failed because two versions of the moved object exist—one each in the source and destination domains. The destination object needs to be removed to restore the system to a consistent state."}
	ErrorDsIllegalXdomMoveOperation                        = &Error{0x0000212C, "ERROR_DS_ILLEGAL_XDOM_MOVE_OPERATION", "This object cannot be moved across domain boundaries either because cross-domain moves for this class are not allowed, or the object has some special characteristics, for example, a trust account or a restricted relative identifier (RID), that prevent its move."}
	ErrorDsCantWithAcctGroupMembershps                     = &Error{0x0000212D, "ERROR_DS_CANT_WITH_ACCT_GROUP_MEMBERSHPS", "Cannot move objects with memberships across domain boundaries because, once moved, this violates the membership conditions of the account group. Remove the object from any account group memberships and retry."}
	ErrorDsNcMustHaveNcParent                              = &Error{0x0000212E, "ERROR_DS_NC_MUST_HAVE_NC_PARENT", "A naming context head must be the immediate child of another naming context head, not of an interior node."}
	ErrorDsCrImpossibleToValidate                          = &Error{0x0000212F, "ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE", "The directory cannot validate the proposed naming context name because it does not hold a replica of the naming context above the proposed naming context. Ensure that the domain naming master role is held by a server that is configured as a GC server, and that the server is up-to-date with its replication partners. (Applies only to Windows 2000 operating system domain naming masters.)"}
	ErrorDsDstDomainNotNative                              = &Error{0x00002130, "ERROR_DS_DST_DOMAIN_NOT_NATIVE", "Destination domain must be in native mode."}
	ErrorDsMissingInfrastructureContainer                  = &Error{0x00002131, "ERROR_DS_MISSING_INFRASTRUCTURE_CONTAINER", "The operation cannot be performed because the server does not have an infrastructure container in the domain of interest."}
	ErrorDsCantMoveAccountGroup                            = &Error{0x00002132, "ERROR_DS_CANT_MOVE_ACCOUNT_GROUP", "Cross-domain moves of nonempty account groups is not allowed."}
	ErrorDsCantMoveResourceGroup                           = &Error{0x00002133, "ERROR_DS_CANT_MOVE_RESOURCE_GROUP", "Cross-domain moves of nonempty resource groups is not allowed."}
	ErrorDsInvalidSearchFlag                               = &Error{0x00002134, "ERROR_DS_INVALID_SEARCH_FLAG", "The search flags for the attribute are invalid. The ambiguous name resolution (ANR) bit is valid only on attributes of Unicode or Teletex strings."}
	ErrorDsNoTreeDeleteAboveNc                             = &Error{0x00002135, "ERROR_DS_NO_TREE_DELETE_ABOVE_NC", "Tree deletions starting at an object that has an NC head as a descendant are not allowed."}
	ErrorDsCouldntLockTreeForDelete                        = &Error{0x00002136, "ERROR_DS_COULDNT_LOCK_TREE_FOR_DELETE", "The directory service failed to lock a tree in preparation for a tree deletion because the tree was in use."}
	ErrorDsCouldntIdentifyObjectsForTreeDelete             = &Error{0x00002137, "ERROR_DS_COULDNT_IDENTIFY_OBJECTS_FOR_TREE_DELETE", "The directory service failed to identify the list of objects to delete while attempting a tree deletion."}
	ErrorDsSamInitFailure                                  = &Error{0x00002138, "ERROR_DS_SAM_INIT_FAILURE", "SAM initialization failed because of the following error: %1. Error Status: 0x%2. Click OK to shut down the system and reboot into Directory Services Restore Mode. Check the event log for detailed information."}
	ErrorDsSensitiveGroupViolation                         = &Error{0x00002139, "ERROR_DS_SENSITIVE_GROUP_VIOLATION", "Only an administrator can modify the membership list of an administrative group."}
	ErrorDsCantModPrimarygroupid                           = &Error{0x0000213A, "ERROR_DS_CANT_MOD_PRIMARYGROUPID", "Cannot change the primary group ID of a domain controller account."}
	ErrorDsIllegalBaseSchemaMod                            = &Error{0x0000213B, "ERROR_DS_ILLEGAL_BASE_SCHEMA_MOD", "An attempt was made to modify the base schema."}
	ErrorDsNonsafeSchemaChange                             = &Error{0x0000213C, "ERROR_DS_NONSAFE_SCHEMA_CHANGE", "Adding a new mandatory attribute to an existing class, deleting a mandatory attribute from an existing class, or adding an optional attribute to the special class Top that is not a backlink attribute (directly or through inheritance, for example, by adding or deleting an auxiliary class) is not allowed."}
	ErrorDsSchemaUpdateDisallowed                          = &Error{0x0000213D, "ERROR_DS_SCHEMA_UPDATE_DISALLOWED", "Schema update is not allowed on this DC because the DC is not the schema FSMO role owner."}
	ErrorDsCantCreateUnderSchema                           = &Error{0x0000213E, "ERROR_DS_CANT_CREATE_UNDER_SCHEMA", "An object of this class cannot be created under the schema container. You can only create Attribute-Schema and Class-Schema objects under the schema container."}
	ErrorDsInstallNoSrcSchVersion                          = &Error{0x0000213F, "ERROR_DS_INSTALL_NO_SRC_SCH_VERSION", "The replica or child install failed to get the objectVersion attribute on the schema container on the source DC. Either the attribute is missing on the schema container or the credentials supplied do not have permission to read it."}
	ErrorDsInstallNoSchVersionInInifile                    = &Error{0x00002140, "ERROR_DS_INSTALL_NO_SCH_VERSION_IN_INIFILE", "The replica or child install failed to read the objectVersion attribute in the SCHEMA section of the file schema.ini in the System32 directory."}
	ErrorDsInvalidGroupType                                = &Error{0x00002141, "ERROR_DS_INVALID_GROUP_TYPE", "The specified group type is invalid."}
	ErrorDsNoNestGlobalgroupInMixeddomain                  = &Error{0x00002142, "ERROR_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN", "You cannot nest global groups in a mixed domain if the group is security-enabled."}
	ErrorDsNoNestLocalgroupInMixeddomain                   = &Error{0x00002143, "ERROR_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN", "You cannot nest local groups in a mixed domain if the group is security-enabled."}
	ErrorDsGlobalCantHaveLocalMember                       = &Error{0x00002144, "ERROR_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER", "A global group cannot have a local group as a member."}
	ErrorDsGlobalCantHaveUniversalMember                   = &Error{0x00002145, "ERROR_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER", "A global group cannot have a universal group as a member."}
	ErrorDsUniversalCantHaveLocalMember                    = &Error{0x00002146, "ERROR_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER", "A universal group cannot have a local group as a member."}
	ErrorDsGlobalCantHaveCrossdomainMember                 = &Error{0x00002147, "ERROR_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER", "A global group cannot have a cross-domain member."}
	ErrorDsLocalCantHaveCrossdomainLocalMember             = &Error{0x00002148, "ERROR_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER", "A local group cannot have another cross domain local group as a member."}
	ErrorDsHavePrimaryMembers                              = &Error{0x00002149, "ERROR_DS_HAVE_PRIMARY_MEMBERS", "A group with primary members cannot change to a security-disabled group."}
	ErrorDsStringSdConversionFailed                        = &Error{0x0000214A, "ERROR_DS_STRING_SD_CONVERSION_FAILED", "The schema cache load failed to convert the string default security descriptor (SD) on a class-schema object."}
	ErrorDsNamingMasterGc                                  = &Error{0x0000214B, "ERROR_DS_NAMING_MASTER_GC", "Only DSAs configured to be GC servers should be allowed to hold the domain naming master FSMO role. (Applies only to Windows 2000 servers.)"}
	ErrorDsDnsLookupFailure                                = &Error{0x0000214C, "ERROR_DS_DNS_LOOKUP_FAILURE", "The DSA operation is unable to proceed because of a DNS lookup failure."}
	ErrorDsCouldntUpdateSpns                               = &Error{0x0000214D, "ERROR_DS_COULDNT_UPDATE_SPNS", "While processing a change to the DNS host name for an object, the SPN values could not be kept in sync."}
	ErrorDsCantRetrieveSd                                  = &Error{0x0000214E, "ERROR_DS_CANT_RETRIEVE_SD", "The Security Descriptor attribute could not be read."}
	ErrorDsKeyNotUnique                                    = &Error{0x0000214F, "ERROR_DS_KEY_NOT_UNIQUE", "The object requested was not found, but an object with that key was found."}
	ErrorDsWrongLinkedAttSyntax                            = &Error{0x00002150, "ERROR_DS_WRONG_LINKED_ATT_SYNTAX", "The syntax of the linked attribute being added is incorrect. Forward links can only have syntax 2.5.5.1, 2.5.5.7, and 2.5.5.14, and backlinks can only have syntax 2.5.5.1."}
	ErrorDsSamNeedBootkeyPassword                          = &Error{0x00002151, "ERROR_DS_SAM_NEED_BOOTKEY_PASSWORD", "SAM needs to get the boot password."}
	ErrorDsSamNeedBootkeyFloppy                            = &Error{0x00002152, "ERROR_DS_SAM_NEED_BOOTKEY_FLOPPY", "SAM needs to get the boot key from the floppy disk."}
	ErrorDsCantStart                                       = &Error{0x00002153, "ERROR_DS_CANT_START", "Directory Service cannot start."}
	ErrorDsInitFailure                                     = &Error{0x00002154, "ERROR_DS_INIT_FAILURE", "Directory Services could not start."}
	ErrorDsNoPktPrivacyOnConnection                        = &Error{0x00002155, "ERROR_DS_NO_PKT_PRIVACY_ON_CONNECTION", "The connection between client and server requires packet privacy or better."}
	ErrorDsSourceDomainInForest                            = &Error{0x00002156, "ERROR_DS_SOURCE_DOMAIN_IN_FOREST", "The source domain cannot be in the same forest as the destination."}
	ErrorDsDestinationDomainNotInForest                    = &Error{0x00002157, "ERROR_DS_DESTINATION_DOMAIN_NOT_IN_FOREST", "The destination domain MUST be in the forest."}
	ErrorDsDestinationAuditingNotEnabled                   = &Error{0x00002158, "ERROR_DS_DESTINATION_AUDITING_NOT_ENABLED", "The operation requires that destination domain auditing be enabled."}
	ErrorDsCantFindDcForSrcDomain                          = &Error{0x00002159, "ERROR_DS_CANT_FIND_DC_FOR_SRC_DOMAIN", "The operation could not locate a DC for the source domain."}
	ErrorDsSrcObjNotGroupOrUser                            = &Error{0x0000215A, "ERROR_DS_SRC_OBJ_NOT_GROUP_OR_USER", "The source object must be a group or user."}
	ErrorDsSrcSidExistsInForest                            = &Error{0x0000215B, "ERROR_DS_SRC_SID_EXISTS_IN_FOREST", "The source object's SID already exists in the destination forest."}
	ErrorDsSrcAndDstObjectClassMismatch                    = &Error{0x0000215C, "ERROR_DS_SRC_AND_DST_OBJECT_CLASS_MISMATCH", "The source and destination object must be of the same type."}
	ErrorSamInitFailure                                    = &Error{0x0000215D, "ERROR_SAM_INIT_FAILURE", "SAM initialization failed because of the following error: %1. Error Status: 0x%2. Click OK to shut down the system and reboot into Safe Mode. Check the event log for detailed information."}
	ErrorDsDraSchemaInfoShip                               = &Error{0x0000215E, "ERROR_DS_DRA_SCHEMA_INFO_SHIP", "Schema information could not be included in the replication request."}
	ErrorDsDraSchemaConflict                               = &Error{0x0000215F, "ERROR_DS_DRA_SCHEMA_CONFLICT", "The replication operation could not be completed due to a schema incompatibility."}
	ErrorDsDraEarlierSchemaConflict                        = &Error{0x00002160, "ERROR_DS_DRA_EARLIER_SCHEMA_CONFLICT", "The replication operation could not be completed due to a previous schema incompatibility."}
	ErrorDsDraObjNcMismatch                                = &Error{0x00002161, "ERROR_DS_DRA_OBJ_NC_MISMATCH", "The replication update could not be applied because either the source or the destination has not yet received information regarding a recent cross-domain move operation."}
	ErrorDsNcStillHasDsas                                  = &Error{0x00002162, "ERROR_DS_NC_STILL_HAS_DSAS", "The requested domain could not be deleted because there exist domain controllers that still host this domain."}
	ErrorDsGcRequired                                      = &Error{0x00002163, "ERROR_DS_GC_REQUIRED", "The requested operation can be performed only on a GC server."}
	ErrorDsLocalMemberOfLocalOnly                          = &Error{0x00002164, "ERROR_DS_LOCAL_MEMBER_OF_LOCAL_ONLY", "A local group can only be a member of other local groups in the same domain."}
	ErrorDsNoFpoInUniversalGroups                          = &Error{0x00002165, "ERROR_DS_NO_FPO_IN_UNIVERSAL_GROUPS", "Foreign security principals cannot be members of universal groups."}
	ErrorDsCantAddToGc                                     = &Error{0x00002166, "ERROR_DS_CANT_ADD_TO_GC", "The attribute is not allowed to be replicated to the GC because of security reasons."}
	ErrorDsNoCheckpointWithPdc                             = &Error{0x00002167, "ERROR_DS_NO_CHECKPOINT_WITH_PDC", "The checkpoint with the PDC could not be taken because too many modifications are currently being processed."}
	ErrorDsSourceAuditingNotEnabled                        = &Error{0x00002168, "ERROR_DS_SOURCE_AUDITING_NOT_ENABLED", "The operation requires that source domain auditing be enabled."}
	ErrorDsCantCreateInNondomainNc                         = &Error{0x00002169, "ERROR_DS_CANT_CREATE_IN_NONDOMAIN_NC", "Security principal objects can only be created inside domain naming contexts."}
	ErrorDsInvalidNameForSpn                               = &Error{0x0000216A, "ERROR_DS_INVALID_NAME_FOR_SPN", "An SPN could not be constructed because the provided host name is not in the necessary format."}
	ErrorDsFilterUsesContructedAttrs                       = &Error{0x0000216B, "ERROR_DS_FILTER_USES_CONTRUCTED_ATTRS", "A filter was passed that uses constructed attributes."}
	ErrorDsUnicodepwdNotInQuotes                           = &Error{0x0000216C, "ERROR_DS_UNICODEPWD_NOT_IN_QUOTES", "The unicodePwd attribute value must be enclosed in quotation marks."}
	ErrorDsMachineAccountQuotaExceeded                     = &Error{0x0000216D, "ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED", "Your computer could not be joined to the domain. You have exceeded the maximum number of computer accounts you are allowed to create in this domain. Contact your system administrator to have this limit reset or increased."}
	ErrorDsMustBeRunOnDstDc                                = &Error{0x0000216E, "ERROR_DS_MUST_BE_RUN_ON_DST_DC", "For security reasons, the operation must be run on the destination DC."}
	ErrorDsSrcDcMustBeSp4OrGreater                         = &Error{0x0000216F, "ERROR_DS_SRC_DC_MUST_BE_SP4_OR_GREATER", "For security reasons, the source DC must be NT4SP4 or greater."}
	ErrorDsCantTreeDeleteCriticalObj                       = &Error{0x00002170, "ERROR_DS_CANT_TREE_DELETE_CRITICAL_OBJ", "Critical directory service system objects cannot be deleted during tree deletion operations. The tree deletion might have been partially performed."}
	ErrorDsInitFailureConsole                              = &Error{0x00002171, "ERROR_DS_INIT_FAILURE_CONSOLE", "Directory Services could not start because of the following error: %1. Error Status: 0x%2. Click OK to shut down the system. You can use the Recovery Console to further diagnose the system."}
	ErrorDsSamInitFailureConsole                           = &Error{0x00002172, "ERROR_DS_SAM_INIT_FAILURE_CONSOLE", "SAM initialization failed because of the following error: %1. Error Status: 0x%2. Click OK to shut down the system. You can use the Recovery Console to further diagnose the system."}
	ErrorDsForestVersionTooHigh                            = &Error{0x00002173, "ERROR_DS_FOREST_VERSION_TOO_HIGH", "The version of the operating system installed is incompatible with the current forest functional level. You must upgrade to a new version of the operating system before this server can become a domain controller in this forest."}
	ErrorDsDomainVersionTooHigh                            = &Error{0x00002174, "ERROR_DS_DOMAIN_VERSION_TOO_HIGH", "The version of the operating system installed is incompatible with the current domain functional level. You must upgrade to a new version of the operating system before this server can become a domain controller in this domain."}
	ErrorDsForestVersionTooLow                             = &Error{0x00002175, "ERROR_DS_FOREST_VERSION_TOO_LOW", "The version of the operating system installed on this server no longer supports the current forest functional level. You must raise the forest functional level before this server can become a domain controller in this forest."}
	ErrorDsDomainVersionTooLow                             = &Error{0x00002176, "ERROR_DS_DOMAIN_VERSION_TOO_LOW", "The version of the operating system installed on this server no longer supports the current domain functional level. You must raise the domain functional level before this server can become a domain controller in this domain."}
	ErrorDsIncompatibleVersion                             = &Error{0x00002177, "ERROR_DS_INCOMPATIBLE_VERSION", "The version of the operating system installed on this server is incompatible with the functional level of the domain or forest."}
	ErrorDsLowDsaVersion                                   = &Error{0x00002178, "ERROR_DS_LOW_DSA_VERSION", "The functional level of the domain (or forest) cannot be raised to the requested value because one or more domain controllers in the domain (or forest) are at a lower, incompatible functional level."}
	ErrorDsNoBehaviorVersionInMixeddomain                  = &Error{0x00002179, "ERROR_DS_NO_BEHAVIOR_VERSION_IN_MIXEDDOMAIN", "The forest functional level cannot be raised to the requested value because one or more domains are still in mixed-domain mode. All domains in the forest must be in native mode for you to raise the forest functional level."}
	ErrorDsNotSupportedSortOrder                           = &Error{0x0000217A, "ERROR_DS_NOT_SUPPORTED_SORT_ORDER", "The sort order requested is not supported."}
	ErrorDsNameNotUnique                                   = &Error{0x0000217B, "ERROR_DS_NAME_NOT_UNIQUE", "The requested name already exists as a unique identifier."}
	ErrorDsMachineAccountCreatedPrent4                     = &Error{0x0000217C, "ERROR_DS_MACHINE_ACCOUNT_CREATED_PRENT4", "The machine account was created before Windows NT 4.0. The account needs to be re-created."}
	ErrorDsOutOfVersionStore                               = &Error{0x0000217D, "ERROR_DS_OUT_OF_VERSION_STORE", "The database is out of version store."}
	ErrorDsIncompatibleControlsUsed                        = &Error{0x0000217E, "ERROR_DS_INCOMPATIBLE_CONTROLS_USED", "Unable to continue operation because multiple conflicting controls were used."}
	ErrorDsNoRefDomain                                     = &Error{0x0000217F, "ERROR_DS_NO_REF_DOMAIN", "Unable to find a valid security descriptor reference domain for this partition."}
	ErrorDsReservedLinkId                                  = &Error{0x00002180, "ERROR_DS_RESERVED_LINK_ID", "Schema update failed: The link identifier is reserved."}
	ErrorDsLinkIdNotAvailable                              = &Error{0x00002181, "ERROR_DS_LINK_ID_NOT_AVAILABLE", "Schema update failed: There are no link identifiers available."}
	ErrorDsAgCantHaveUniversalMember                       = &Error{0x00002182, "ERROR_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER", "An account group cannot have a universal group as a member."}
	ErrorDsModifydnDisallowedByInstanceType                = &Error{0x00002183, "ERROR_DS_MODIFYDN_DISALLOWED_BY_INSTANCE_TYPE", "Rename or move operations on naming context heads or read-only objects are not allowed."}
	ErrorDsNoObjectMoveInSchemaNc                          = &Error{0x00002184, "ERROR_DS_NO_OBJECT_MOVE_IN_SCHEMA_NC", "Move operations on objects in the schema naming context are not allowed."}
	ErrorDsModifydnDisallowedByFlag                        = &Error{0x00002185, "ERROR_DS_MODIFYDN_DISALLOWED_BY_FLAG", "A system flag has been set on the object that does not allow the object to be moved or renamed."}
	ErrorDsModifydnWrongGrandparent                        = &Error{0x00002186, "ERROR_DS_MODIFYDN_WRONG_GRANDPARENT", "This object is not allowed to change its grandparent container. Moves are not forbidden on this object, but are restricted to sibling containers."}
	ErrorDsNameErrorTrustReferral                          = &Error{0x00002187, "ERROR_DS_NAME_ERROR_TRUST_REFERRAL", "Unable to resolve completely; a referral to another forest was generated."}
	ErrorNotSupportedOnStandardServer                      = &Error{0x00002188, "ERROR_NOT_SUPPORTED_ON_STANDARD_SERVER", "The requested action is not supported on a standard server."}
	ErrorDsCantAccessRemotePartOfAd                        = &Error{0x00002189, "ERROR_DS_CANT_ACCESS_REMOTE_PART_OF_AD", "Could not access a partition of the directory service located on a remote server. Make sure at least one server is running for the partition in question."}
	ErrorDsCrImpossibleToValidateV2                        = &Error{0x0000218A, "ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE_V2", "The directory cannot validate the proposed naming context (or partition) name because it does not hold a replica, nor can it contact a replica of the naming context above the proposed naming context. Ensure that the parent naming context is properly registered in the DNS, and at least one replica of this naming context is reachable by the domain naming master."}
	ErrorDsThreadLimitExceeded                             = &Error{0x0000218B, "ERROR_DS_THREAD_LIMIT_EXCEEDED", "The thread limit for this request was exceeded."}
	ErrorDsNotClosest                                      = &Error{0x0000218C, "ERROR_DS_NOT_CLOSEST", "The GC server is not in the closest site."}
	ErrorDsCantDeriveSpnWithoutServerRef                   = &Error{0x0000218D, "ERROR_DS_CANT_DERIVE_SPN_WITHOUT_SERVER_REF", "The directory service cannot derive an SPN with which to mutually authenticate the target server because the corresponding server object in the local DS database has no serverReference attribute."}
	ErrorDsSingleUserModeFailed                            = &Error{0x0000218E, "ERROR_DS_SINGLE_USER_MODE_FAILED", "The directory service failed to enter single-user mode."}
	ErrorDsNtdscriptSyntaxError                            = &Error{0x0000218F, "ERROR_DS_NTDSCRIPT_SYNTAX_ERROR", "The directory service cannot parse the script because of a syntax error."}
	ErrorDsNtdscriptProcessError                           = &Error{0x00002190, "ERROR_DS_NTDSCRIPT_PROCESS_ERROR", "The directory service cannot process the script because of an error."}
	ErrorDsDifferentReplEpochs                             = &Error{0x00002191, "ERROR_DS_DIFFERENT_REPL_EPOCHS", "The directory service cannot perform the requested operation because the servers involved are of different replication epochs (which is usually related to a domain rename that is in progress)."}
	ErrorDsDrsExtensionsChanged                            = &Error{0x00002192, "ERROR_DS_DRS_EXTENSIONS_CHANGED", "The directory service binding must be renegotiated due to a change in the server extensions information."}
	ErrorDsReplicaSetChangeNotAllowedOnDisabledCr          = &Error{0x00002193, "ERROR_DS_REPLICA_SET_CHANGE_NOT_ALLOWED_ON_DISABLED_CR", "The operation is not allowed on a disabled cross-reference."}
	ErrorDsNoMsdsIntid                                     = &Error{0x00002194, "ERROR_DS_NO_MSDS_INTID", "Schema update failed: No values for msDS-IntId are available."}
	ErrorDsDupMsdsIntid                                    = &Error{0x00002195, "ERROR_DS_DUP_MSDS_INTID", "Schema update failed: Duplicate msDS-IntId. Retry the operation."}
	ErrorDsExistsInRdnattid                                = &Error{0x00002196, "ERROR_DS_EXISTS_IN_RDNATTID", "Schema deletion failed: Attribute is used in rDNAttID."}
	ErrorDsAuthorizationFailed                             = &Error{0x00002197, "ERROR_DS_AUTHORIZATION_FAILED", "The directory service failed to authorize the request."}
	ErrorDsInvalidScript                                   = &Error{0x00002198, "ERROR_DS_INVALID_SCRIPT", "The directory service cannot process the script because it is invalid."}
	ErrorDsRemoteCrossrefOpFailed                          = &Error{0x00002199, "ERROR_DS_REMOTE_CROSSREF_OP_FAILED", "The remote create cross-reference operation failed on the domain naming master FSMO. The operation's error is in the extended data."}
	ErrorDsCrossRefBusy                                    = &Error{0x0000219A, "ERROR_DS_CROSS_REF_BUSY", "A cross-reference is in use locally with the same name."}
	ErrorDsCantDeriveSpnForDeletedDomain                   = &Error{0x0000219B, "ERROR_DS_CANT_DERIVE_SPN_FOR_DELETED_DOMAIN", "The directory service cannot derive an SPN with which to mutually authenticate the target server because the server's domain has been deleted from the forest."}
	ErrorDsCantDemoteWithWriteableNc                       = &Error{0x0000219C, "ERROR_DS_CANT_DEMOTE_WITH_WRITEABLE_NC", "Writable NCs prevent this DC from demoting."}
	ErrorDsDuplicateIdFound                                = &Error{0x0000219D, "ERROR_DS_DUPLICATE_ID_FOUND", "The requested object has a nonunique identifier and cannot be retrieved."}
	ErrorDsInsufficientAttrToCreateObject                  = &Error{0x0000219E, "ERROR_DS_INSUFFICIENT_ATTR_TO_CREATE_OBJECT", "Insufficient attributes were given to create an object. This object might not exist because it might have been deleted and the garbage already collected."}
	ErrorDsGroupConversionError                            = &Error{0x0000219F, "ERROR_DS_GROUP_CONVERSION_ERROR", "The group cannot be converted due to attribute restrictions on the requested group type."}
	ErrorDsCantMoveAppBasicGroup                           = &Error{0x000021A0, "ERROR_DS_CANT_MOVE_APP_BASIC_GROUP", "Cross-domain moves of nonempty basic application groups is not allowed."}
	ErrorDsCantMoveAppQueryGroup                           = &Error{0x000021A1, "ERROR_DS_CANT_MOVE_APP_QUERY_GROUP", "Cross-domain moves of nonempty query-based application groups is not allowed."}
	ErrorDsRoleNotVerified                                 = &Error{0x000021A2, "ERROR_DS_ROLE_NOT_VERIFIED", "The FSMO role ownership could not be verified because its directory partition did not replicate successfully with at least one replication partner."}
	ErrorDsWkoContainerCannotBeSpecial                     = &Error{0x000021A3, "ERROR_DS_WKO_CONTAINER_CANNOT_BE_SPECIAL", "The target container for a redirection of a well-known object container cannot already be a special container."}
	ErrorDsDomainRenameInProgress                          = &Error{0x000021A4, "ERROR_DS_DOMAIN_RENAME_IN_PROGRESS", "The directory service cannot perform the requested operation because a domain rename operation is in progress."}
	ErrorDsExistingAdChildNc                               = &Error{0x000021A5, "ERROR_DS_EXISTING_AD_CHILD_NC", "The directory service detected a child partition below the requested partition name. The partition hierarchy must be created in a top down method."}
	ErrorDsReplLifetimeExceeded                            = &Error{0x000021A6, "ERROR_DS_REPL_LIFETIME_EXCEEDED", "The directory service cannot replicate with this server because the time since the last replication with this server has exceeded the tombstone lifetime."}
	ErrorDsDisallowedInSystemContainer                     = &Error{0x000021A7, "ERROR_DS_DISALLOWED_IN_SYSTEM_CONTAINER", "The requested operation is not allowed on an object under the system container."}
	ErrorDsLdapSendQueueFull                               = &Error{0x000021A8, "ERROR_DS_LDAP_SEND_QUEUE_FULL", "The LDAP server's network send queue has filled up because the client is not processing the results of its requests fast enough. No more requests will be processed until the client catches up. If the client does not catch up then it will be disconnected."}
	ErrorDsDraOutScheduleWindow                            = &Error{0x000021A9, "ERROR_DS_DRA_OUT_SCHEDULE_WINDOW", "The scheduled replication did not take place because the system was too busy to execute the request within the schedule window. The replication queue is overloaded. Consider reducing the number of partners or decreasing the scheduled replication frequency."}
	ErrorDsPolicyNotKnown                                  = &Error{0x000021AA, "ERROR_DS_POLICY_NOT_KNOWN", "At this time, it cannot be determined if the branch replication policy is available on the hub domain controller. Retry at a later time to account for replication latencies."}
	ErrorNoSiteSettingsObject                              = &Error{0x000021AB, "ERROR_NO_SITE_SETTINGS_OBJECT", "The site settings object for the specified site does not exist."}
	ErrorNoSecrets                                         = &Error{0x000021AC, "ERROR_NO_SECRETS", "The local account store does not contain secret material for the specified account."}
	ErrorNoWritableDcFound                                 = &Error{0x000021AD, "ERROR_NO_WRITABLE_DC_FOUND", "Could not find a writable domain controller in the domain."}
	ErrorDsNoServerObject                                  = &Error{0x000021AE, "ERROR_DS_NO_SERVER_OBJECT", "The server object for the domain controller does not exist."}
	ErrorDsNoNtdsaObject                                   = &Error{0x000021AF, "ERROR_DS_NO_NTDSA_OBJECT", "The NTDS Settings object for the domain controller does not exist."}
	ErrorDsNonAsqSearch                                    = &Error{0x000021B0, "ERROR_DS_NON_ASQ_SEARCH", "The requested search operation is not supported for attribute scoped query (ASQ) searches."}
	ErrorDsAuditFailure                                    = &Error{0x000021B1, "ERROR_DS_AUDIT_FAILURE", "A required audit event could not be generated for the operation."}
	ErrorDsInvalidSearchFlagSubtree                        = &Error{0x000021B2, "ERROR_DS_INVALID_SEARCH_FLAG_SUBTREE", "The search flags for the attribute are invalid. The subtree index bit is valid only on single-valued attributes."}
	ErrorDsInvalidSearchFlagTuple                          = &Error{0x000021B3, "ERROR_DS_INVALID_SEARCH_FLAG_TUPLE", "The search flags for the attribute are invalid. The tuple index bit is valid only on attributes of Unicode strings."}
	ErrorDsDraRecycledTarget                               = &Error{0x000021BF, "ERROR_DS_DRA_RECYCLED_TARGET", "The replication operation failed because the target object referenced by a link value is recycled."}
	ErrorDsHighDsaVersion                                  = &Error{0x000021C2, "ERROR_DS_HIGH_DSA_VERSION", "The functional level of the domain (or forest) cannot be lowered to the requested value."}
	ErrorDsSpnValueNotUniqueInForest                       = &Error{0x000021C7, "ERROR_DS_SPN_VALUE_NOT_UNIQUE_IN_FOREST", "The operation failed because the SPN value provided for addition/modification is not unique forest-wide."}
	ErrorDsUpnValueNotUniqueInForest                       = &Error{0x000021C8, "ERROR_DS_UPN_VALUE_NOT_UNIQUE_IN_FOREST", "The operation failed because the UPN value provided for addition/modification is not unique forest-wide."}
	DnsErrorRcodeFormatError                               = &Error{0x00002329, "DNS_ERROR_RCODE_FORMAT_ERROR", "DNS server unable to interpret format."}
	DnsErrorRcodeServerFailure                             = &Error{0x0000232A, "DNS_ERROR_RCODE_SERVER_FAILURE", "DNS server failure."}
	DnsErrorRcodeNameError                                 = &Error{0x0000232B, "DNS_ERROR_RCODE_NAME_ERROR", "DNS name does not exist."}
	DnsErrorRcodeNotImplemented                            = &Error{0x0000232C, "DNS_ERROR_RCODE_NOT_IMPLEMENTED", "DNS request not supported by name server."}
	DnsErrorRcodeRefused                                   = &Error{0x0000232D, "DNS_ERROR_RCODE_REFUSED", "DNS operation refused."}
	DnsErrorRcodeYxdomain                                  = &Error{0x0000232E, "DNS_ERROR_RCODE_YXDOMAIN", "DNS name that should not exist, does exist."}
	DnsErrorRcodeYxrrset                                   = &Error{0x0000232F, "DNS_ERROR_RCODE_YXRRSET", "DNS resource record (RR) set that should not exist, does exist."}
	DnsErrorRcodeNxrrset                                   = &Error{0x00002330, "DNS_ERROR_RCODE_NXRRSET", "DNS RR set that should to exist, does not exist."}
	DnsErrorRcodeNotauth                                   = &Error{0x00002331, "DNS_ERROR_RCODE_NOTAUTH", "DNS server not authoritative for zone."}
	DnsErrorRcodeNotzone                                   = &Error{0x00002332, "DNS_ERROR_RCODE_NOTZONE", "DNS name in update or prereq is not in zone."}
	DnsErrorRcodeBadsig                                    = &Error{0x00002338, "DNS_ERROR_RCODE_BADSIG", "DNS signature failed to verify."}
	DnsErrorRcodeBadkey                                    = &Error{0x00002339, "DNS_ERROR_RCODE_BADKEY", "DNS bad key."}
	DnsErrorRcodeBadtime                                   = &Error{0x0000233A, "DNS_ERROR_RCODE_BADTIME", "DNS signature validity expired."}
	DnsInfoNoRecords                                       = &Error{0x0000251D, "DNS_INFO_NO_RECORDS", "No records found for given DNS query."}
	DnsErrorBadPacket                                      = &Error{0x0000251E, "DNS_ERROR_BAD_PACKET", "Bad DNS packet."}
	DnsErrorNoPacket                                       = &Error{0x0000251F, "DNS_ERROR_NO_PACKET", "No DNS packet."}
	DnsErrorRcode                                          = &Error{0x00002520, "DNS_ERROR_RCODE", "DNS error, check rcode."}
	DnsErrorUnsecurePacket                                 = &Error{0x00002521, "DNS_ERROR_UNSECURE_PACKET", "Unsecured DNS packet."}
	DnsErrorInvalidType                                    = &Error{0x0000254F, "DNS_ERROR_INVALID_TYPE", "Invalid DNS type."}
	DnsErrorInvalidIpAddress                               = &Error{0x00002550, "DNS_ERROR_INVALID_IP_ADDRESS", "Invalid IP address."}
	DnsErrorInvalidProperty                                = &Error{0x00002551, "DNS_ERROR_INVALID_PROPERTY", "Invalid property."}
	DnsErrorTryAgainLater                                  = &Error{0x00002552, "DNS_ERROR_TRY_AGAIN_LATER", "Try DNS operation again later."}
	DnsErrorNotUnique                                      = &Error{0x00002553, "DNS_ERROR_NOT_UNIQUE", "Record for given name and type is not unique."}
	DnsErrorNonRfcName                                     = &Error{0x00002554, "DNS_ERROR_NON_RFC_NAME", "DNS name does not comply with RFC specifications."}
	DnsStatusFqdn                                          = &Error{0x00002555, "DNS_STATUS_FQDN", "DNS name is a fully qualified DNS name."}
	DnsStatusDottedName                                    = &Error{0x00002556, "DNS_STATUS_DOTTED_NAME", "DNS name is dotted (multilabel)."}
	DnsStatusSinglePartName                                = &Error{0x00002557, "DNS_STATUS_SINGLE_PART_NAME", "DNS name is a single-part name."}
	DnsErrorInvalidNameChar                                = &Error{0x00002558, "DNS_ERROR_INVALID_NAME_CHAR", "DNS name contains an invalid character."}
	DnsErrorNumericName                                    = &Error{0x00002559, "DNS_ERROR_NUMERIC_NAME", "DNS name is entirely numeric."}
	DnsErrorNotAllowedOnRootServer                         = &Error{0x0000255A, "DNS_ERROR_NOT_ALLOWED_ON_ROOT_SERVER", "The operation requested is not permitted on a DNS root server."}
	DnsErrorNotAllowedUnderDelegation                      = &Error{0x0000255B, "DNS_ERROR_NOT_ALLOWED_UNDER_DELEGATION", "The record could not be created because this part of the DNS namespace has been delegated to another server."}
	DnsErrorCannotFindRootHints                            = &Error{0x0000255C, "DNS_ERROR_CANNOT_FIND_ROOT_HINTS", "The DNS server could not find a set of root hints."}
	DnsErrorInconsistentRootHints                          = &Error{0x0000255D, "DNS_ERROR_INCONSISTENT_ROOT_HINTS", "The DNS server found root hints but they were not consistent across all adapters."}
	DnsErrorDwordValueTooSmall                             = &Error{0x0000255E, "DNS_ERROR_DWORD_VALUE_TOO_SMALL", "The specified value is too small for this parameter."}
	DnsErrorDwordValueTooLarge                             = &Error{0x0000255F, "DNS_ERROR_DWORD_VALUE_TOO_LARGE", "The specified value is too large for this parameter."}
	DnsErrorBackgroundLoading                              = &Error{0x00002560, "DNS_ERROR_BACKGROUND_LOADING", "This operation is not allowed while the DNS server is loading zones in the background. Try again later."}
	DnsErrorNotAllowedOnRodc                               = &Error{0x00002561, "DNS_ERROR_NOT_ALLOWED_ON_RODC", "The operation requested is not permitted on against a DNS server running on a read-only DC."}
	DnsErrorZoneDoesNotExist                               = &Error{0x00002581, "DNS_ERROR_ZONE_DOES_NOT_EXIST", "DNS zone does not exist."}
	DnsErrorNoZoneInfo                                     = &Error{0x00002582, "DNS_ERROR_NO_ZONE_INFO", "DNS zone information not available."}
	DnsErrorInvalidZoneOperation                           = &Error{0x00002583, "DNS_ERROR_INVALID_ZONE_OPERATION", "Invalid operation for DNS zone."}
	DnsErrorZoneConfigurationError                         = &Error{0x00002584, "DNS_ERROR_ZONE_CONFIGURATION_ERROR", "Invalid DNS zone configuration."}
	DnsErrorZoneHasNoSoaRecord                             = &Error{0x00002585, "DNS_ERROR_ZONE_HAS_NO_SOA_RECORD", "DNS zone has no start of authority (SOA) record."}
	DnsErrorZoneHasNoNsRecords                             = &Error{0x00002586, "DNS_ERROR_ZONE_HAS_NO_NS_RECORDS", "DNS zone has no Name Server (NS) record."}
	DnsErrorZoneLocked                                     = &Error{0x00002587, "DNS_ERROR_ZONE_LOCKED", "DNS zone is locked."}
	DnsErrorZoneCreationFailed                             = &Error{0x00002588, "DNS_ERROR_ZONE_CREATION_FAILED", "DNS zone creation failed."}
	DnsErrorZoneAlreadyExists                              = &Error{0x00002589, "DNS_ERROR_ZONE_ALREADY_EXISTS", "DNS zone already exists."}
	DnsErrorAutozoneAlreadyExists                          = &Error{0x0000258A, "DNS_ERROR_AUTOZONE_ALREADY_EXISTS", "DNS automatic zone already exists."}
	DnsErrorInvalidZoneType                                = &Error{0x0000258B, "DNS_ERROR_INVALID_ZONE_TYPE", "Invalid DNS zone type."}
	DnsErrorSecondaryRequiresMasterIp                      = &Error{0x0000258C, "DNS_ERROR_SECONDARY_REQUIRES_MASTER_IP", "Secondary DNS zone requires master IP address."}
	DnsErrorZoneNotSecondary                               = &Error{0x0000258D, "DNS_ERROR_ZONE_NOT_SECONDARY", "DNS zone not secondary."}
	DnsErrorNeedSecondaryAddresses                         = &Error{0x0000258E, "DNS_ERROR_NEED_SECONDARY_ADDRESSES", "Need secondary IP address."}
	DnsErrorWinsInitFailed                                 = &Error{0x0000258F, "DNS_ERROR_WINS_INIT_FAILED", "WINS initialization failed."}
	DnsErrorNeedWinsServers                                = &Error{0x00002590, "DNS_ERROR_NEED_WINS_SERVERS", "Need WINS servers."}
	DnsErrorNbstatInitFailed                               = &Error{0x00002591, "DNS_ERROR_NBSTAT_INIT_FAILED", "NBTSTAT initialization call failed."}
	DnsErrorSoaDeleteInvalid                               = &Error{0x00002592, "DNS_ERROR_SOA_DELETE_INVALID", "Invalid delete of SOA."}
	DnsErrorForwarderAlreadyExists                         = &Error{0x00002593, "DNS_ERROR_FORWARDER_ALREADY_EXISTS", "A conditional forwarding zone already exists for that name."}
	DnsErrorZoneRequiresMasterIp                           = &Error{0x00002594, "DNS_ERROR_ZONE_REQUIRES_MASTER_IP", "This zone must be configured with one or more master DNS server IP addresses."}
	DnsErrorZoneIsShutdown                                 = &Error{0x00002595, "DNS_ERROR_ZONE_IS_SHUTDOWN", "The operation cannot be performed because this zone is shut down."}
	DnsErrorPrimaryRequiresDatafile                        = &Error{0x000025B3, "DNS_ERROR_PRIMARY_REQUIRES_DATAFILE", "The primary DNS zone requires a data file."}
	DnsErrorInvalidDatafileName                            = &Error{0x000025B4, "DNS_ERROR_INVALID_DATAFILE_NAME", "Invalid data file name for the DNS zone."}
	DnsErrorDatafileOpenFailure                            = &Error{0x000025B5, "DNS_ERROR_DATAFILE_OPEN_FAILURE", "Failed to open the data file for the DNS zone."}
	DnsErrorFileWritebackFailed                            = &Error{0x000025B6, "DNS_ERROR_FILE_WRITEBACK_FAILED", "Failed to write the data file for the DNS zone."}
	DnsErrorDatafileParsing                                = &Error{0x000025B7, "DNS_ERROR_DATAFILE_PARSING", "Failure while reading datafile for DNS zone."}
	DnsErrorRecordDoesNotExist                             = &Error{0x000025E5, "DNS_ERROR_RECORD_DOES_NOT_EXIST", "DNS record does not exist."}
	DnsErrorRecordFormat                                   = &Error{0x000025E6, "DNS_ERROR_RECORD_FORMAT", "DNS record format error."}
	DnsErrorNodeCreationFailed                             = &Error{0x000025E7, "DNS_ERROR_NODE_CREATION_FAILED", "Node creation failure in DNS."}
	DnsErrorUnknownRecordType                              = &Error{0x000025E8, "DNS_ERROR_UNKNOWN_RECORD_TYPE", "Unknown DNS record type."}
	DnsErrorRecordTimedOut                                 = &Error{0x000025E9, "DNS_ERROR_RECORD_TIMED_OUT", "DNS record timed out."}
	DnsErrorNameNotInZone                                  = &Error{0x000025EA, "DNS_ERROR_NAME_NOT_IN_ZONE", "Name not in DNS zone."}
	DnsErrorCnameLoop                                      = &Error{0x000025EB, "DNS_ERROR_CNAME_LOOP", "CNAME loop detected."}
	DnsErrorNodeIsCname                                    = &Error{0x000025EC, "DNS_ERROR_NODE_IS_CNAME", "Node is a CNAME DNS record."}
	DnsErrorCnameCollision                                 = &Error{0x000025ED, "DNS_ERROR_CNAME_COLLISION", "A CNAME record already exists for the given name."}
	DnsErrorRecordOnlyAtZoneRoot                           = &Error{0x000025EE, "DNS_ERROR_RECORD_ONLY_AT_ZONE_ROOT", "Record is only at DNS zone root."}
	DnsErrorRecordAlreadyExists                            = &Error{0x000025EF, "DNS_ERROR_RECORD_ALREADY_EXISTS", "DNS record already exists."}
	DnsErrorSecondaryData                                  = &Error{0x000025F0, "DNS_ERROR_SECONDARY_DATA", "Secondary DNS zone data error."}
	DnsErrorNoCreateCacheData                              = &Error{0x000025F1, "DNS_ERROR_NO_CREATE_CACHE_DATA", "Could not create DNS cache data."}
	DnsErrorNameDoesNotExist                               = &Error{0x000025F2, "DNS_ERROR_NAME_DOES_NOT_EXIST", "DNS name does not exist."}
	DnsWarningPtrCreateFailed                              = &Error{0x000025F3, "DNS_WARNING_PTR_CREATE_FAILED", "Could not create pointer (PTR) record."}
	DnsWarningDomainUndeleted                              = &Error{0x000025F4, "DNS_WARNING_DOMAIN_UNDELETED", "DNS domain was undeleted."}
	DnsErrorDsUnavailable                                  = &Error{0x000025F5, "DNS_ERROR_DS_UNAVAILABLE", "The directory service is unavailable."}
	DnsErrorDsZoneAlreadyExists                            = &Error{0x000025F6, "DNS_ERROR_DS_ZONE_ALREADY_EXISTS", "DNS zone already exists in the directory service."}
	DnsErrorNoBootfileIfDsZone                             = &Error{0x000025F7, "DNS_ERROR_NO_BOOTFILE_IF_DS_ZONE", "DNS server not creating or reading the boot file for the directory service integrated DNS zone."}
	DnsInfoAxfrComplete                                    = &Error{0x00002617, "DNS_INFO_AXFR_COMPLETE", "DNS AXFR (zone transfer) complete."}
	DnsErrorAxfr                                           = &Error{0x00002618, "DNS_ERROR_AXFR", "DNS zone transfer failed."}
	DnsInfoAddedLocalWins                                  = &Error{0x00002619, "DNS_INFO_ADDED_LOCAL_WINS", "Added local WINS server."}
	DnsStatusContinueNeeded                                = &Error{0x00002649, "DNS_STATUS_CONTINUE_NEEDED", "Secure update call needs to continue update request."}
	DnsErrorNoTcpip                                        = &Error{0x0000267B, "DNS_ERROR_NO_TCPIP", "TCP/IP network protocol not installed."}
	DnsErrorNoDnsServers                                   = &Error{0x0000267C, "DNS_ERROR_NO_DNS_SERVERS", "No DNS servers configured for local system."}
	DnsErrorDpDoesNotExist                                 = &Error{0x000026AD, "DNS_ERROR_DP_DOES_NOT_EXIST", "The specified directory partition does not exist."}
	DnsErrorDpAlreadyExists                                = &Error{0x000026AE, "DNS_ERROR_DP_ALREADY_EXISTS", "The specified directory partition already exists."}
	DnsErrorDpNotEnlisted                                  = &Error{0x000026AF, "DNS_ERROR_DP_NOT_ENLISTED", "This DNS server is not enlisted in the specified directory partition."}
	DnsErrorDpAlreadyEnlisted                              = &Error{0x000026B0, "DNS_ERROR_DP_ALREADY_ENLISTED", "This DNS server is already enlisted in the specified directory partition."}
	DnsErrorDpNotAvailable                                 = &Error{0x000026B1, "DNS_ERROR_DP_NOT_AVAILABLE", "The directory partition is not available at this time. Wait a few minutes and try again."}
	DnsErrorDpFsmoError                                    = &Error{0x000026B2, "DNS_ERROR_DP_FSMO_ERROR", "The application directory partition operation failed. The domain controller holding the domain naming master role is down or unable to service the request or is not running Windows Server 2003."}
	Wsaeintr                                               = &Error{0x00002714, "WSAEINTR", "A blocking operation was interrupted by a call to WSACancelBlockingCall."}
	Wsaebadf                                               = &Error{0x00002719, "WSAEBADF", "The file handle supplied is not valid."}
	Wsaeacces                                              = &Error{0x0000271D, "WSAEACCES", "An attempt was made to access a socket in a way forbidden by its access permissions."}
	Wsaefault                                              = &Error{0x0000271E, "WSAEFAULT", "The system detected an invalid pointer address in attempting to use a pointer argument in a call."}
	Wsaeinval                                              = &Error{0x00002726, "WSAEINVAL", "An invalid argument was supplied."}
	Wsaemfile                                              = &Error{0x00002728, "WSAEMFILE", "Too many open sockets."}
	Wsaewouldblock                                         = &Error{0x00002733, "WSAEWOULDBLOCK", "A nonblocking socket operation could not be completed immediately."}
	Wsaeinprogress                                         = &Error{0x00002734, "WSAEINPROGRESS", "A blocking operation is currently executing."}
	Wsaealready                                            = &Error{0x00002735, "WSAEALREADY", "An operation was attempted on a nonblocking socket that already had an operation in progress."}
	Wsaenotsock                                            = &Error{0x00002736, "WSAENOTSOCK", "An operation was attempted on something that is not a socket."}
	Wsaedestaddrreq                                        = &Error{0x00002737, "WSAEDESTADDRREQ", "A required address was omitted from an operation on a socket."}
	Wsaemsgsize                                            = &Error{0x00002738, "WSAEMSGSIZE", "A message sent on a datagram socket was larger than the internal message buffer or some other network limit, or the buffer used to receive a datagram into was smaller than the datagram itself."}
	Wsaeprototype                                          = &Error{0x00002739, "WSAEPROTOTYPE", "A protocol was specified in the socket function call that does not support the semantics of the socket type requested."}
	Wsaenoprotoopt                                         = &Error{0x0000273A, "WSAENOPROTOOPT", "An unknown, invalid, or unsupported option or level was specified in a getsockopt or setsockopt call."}
	Wsaeprotonosupport                                     = &Error{0x0000273B, "WSAEPROTONOSUPPORT", "The requested protocol has not been configured into the system, or no implementation for it exists."}
	Wsaesocktnosupport                                     = &Error{0x0000273C, "WSAESOCKTNOSUPPORT", "The support for the specified socket type does not exist in this address family."}
	Wsaeopnotsupp                                          = &Error{0x0000273D, "WSAEOPNOTSUPP", "The attempted operation is not supported for the type of object referenced."}
	Wsaepfnosupport                                        = &Error{0x0000273E, "WSAEPFNOSUPPORT", "The protocol family has not been configured into the system or no implementation for it exists."}
	Wsaeafnosupport                                        = &Error{0x0000273F, "WSAEAFNOSUPPORT", "An address incompatible with the requested protocol was used."}
	Wsaeaddrinuse                                          = &Error{0x00002740, "WSAEADDRINUSE", "Only one usage of each socket address (protocol/network address/port) is normally permitted."}
	Wsaeaddrnotavail                                       = &Error{0x00002741, "WSAEADDRNOTAVAIL", "The requested address is not valid in its context."}
	Wsaenetdown                                            = &Error{0x00002742, "WSAENETDOWN", "A socket operation encountered a dead network."}
	Wsaenetunreach                                         = &Error{0x00002743, "WSAENETUNREACH", "A socket operation was attempted to an unreachable network."}
	Wsaenetreset                                           = &Error{0x00002744, "WSAENETRESET", "The connection has been broken due to keep-alive activity detecting a failure while the operation was in progress."}
	Wsaeconnaborted                                        = &Error{0x00002745, "WSAECONNABORTED", "An established connection was aborted by the software in your host machine."}
	Wsaeconnreset                                          = &Error{0x00002746, "WSAECONNRESET", "An existing connection was forcibly closed by the remote host."}
	Wsaenobufs                                             = &Error{0x00002747, "WSAENOBUFS", "An operation on a socket could not be performed because the system lacked sufficient buffer space or because a queue was full."}
	Wsaeisconn                                             = &Error{0x00002748, "WSAEISCONN", "A connect request was made on an already connected socket."}
	Wsaenotconn                                            = &Error{0x00002749, "WSAENOTCONN", "A request to send or receive data was disallowed because the socket is not connected and (when sending on a datagram socket using a sendto call) no address was supplied."}
	Wsaeshutdown                                           = &Error{0x0000274A, "WSAESHUTDOWN", "A request to send or receive data was disallowed because the socket had already been shut down in that direction with a previous shutdown call."}
	Wsaetoomanyrefs                                        = &Error{0x0000274B, "WSAETOOMANYREFS", "Too many references to a kernel object."}
	Wsaetimedout                                           = &Error{0x0000274C, "WSAETIMEDOUT", "A connection attempt failed because the connected party did not properly respond after a period of time, or the established connection failed because the connected host failed to respond."}
	Wsaeconnrefused                                        = &Error{0x0000274D, "WSAECONNREFUSED", "No connection could be made because the target machine actively refused it."}
	Wsaeloop                                               = &Error{0x0000274E, "WSAELOOP", "Cannot translate name."}
	Wsaenametoolong                                        = &Error{0x0000274F, "WSAENAMETOOLONG", "Name or name component was too long."}
	Wsaehostdown                                           = &Error{0x00002750, "WSAEHOSTDOWN", "A socket operation failed because the destination host was down."}
	Wsaehostunreach                                        = &Error{0x00002751, "WSAEHOSTUNREACH", "A socket operation was attempted to an unreachable host."}
	Wsaenotempty                                           = &Error{0x00002752, "WSAENOTEMPTY", "Cannot remove a directory that is not empty."}
	Wsaeproclim                                            = &Error{0x00002753, "WSAEPROCLIM", "A Windows Sockets implementation might have a limit on the number of applications that can use it simultaneously."}
	Wsaeusers                                              = &Error{0x00002754, "WSAEUSERS", "Ran out of quota."}
	Wsaedquot                                              = &Error{0x00002755, "WSAEDQUOT", "Ran out of disk quota."}
	Wsaestale                                              = &Error{0x00002756, "WSAESTALE", "File handle reference is no longer available."}
	Wsaeremote                                             = &Error{0x00002757, "WSAEREMOTE", "Item is not available locally."}
	Wsasysnotready                                         = &Error{0x0000276B, "WSASYSNOTREADY", "WSAStartup cannot function at this time because the underlying system it uses to provide network services is currently unavailable."}
	Wsavernotsupported                                     = &Error{0x0000276C, "WSAVERNOTSUPPORTED", "The Windows Sockets version requested is not supported."}
	Wsanotinitialised                                      = &Error{0x0000276D, "WSANOTINITIALISED", "Either the application has not called WSAStartup, or WSAStartup failed."}
	Wsaediscon                                             = &Error{0x00002775, "WSAEDISCON", "Returned by WSARecv or WSARecvFrom to indicate that the remote party has initiated a graceful shutdown sequence."}
	Wsaenomore                                             = &Error{0x00002776, "WSAENOMORE", "No more results can be returned by WSALookupServiceNext."}
	Wsaecancelled                                          = &Error{0x00002777, "WSAECANCELLED", "A call to WSALookupServiceEnd was made while this call was still processing. The call has been canceled."}
	Wsaeinvalidproctable                                   = &Error{0x00002778, "WSAEINVALIDPROCTABLE", "The procedure call table is invalid."}
	Wsaeinvalidprovider                                    = &Error{0x00002779, "WSAEINVALIDPROVIDER", "The requested service provider is invalid."}
	Wsaeproviderfailedinit                                 = &Error{0x0000277A, "WSAEPROVIDERFAILEDINIT", "The requested service provider could not be loaded or initialized."}
	Wsasyscallfailure                                      = &Error{0x0000277B, "WSASYSCALLFAILURE", "A system call that should never fail has failed."}
	WsaserviceNotFound                                     = &Error{0x0000277C, "WSASERVICE_NOT_FOUND", "No such service is known. The service cannot be found in the specified namespace."}
	WsatypeNotFound                                        = &Error{0x0000277D, "WSATYPE_NOT_FOUND", "The specified class was not found."}
	WsaENoMore                                             = &Error{0x0000277E, "WSA_E_NO_MORE", "No more results can be returned by WSALookupServiceNext."}
	WsaECancelled                                          = &Error{0x0000277F, "WSA_E_CANCELLED", "A call to WSALookupServiceEnd was made while this call was still processing. The call has been canceled."}
	Wsaerefused                                            = &Error{0x00002780, "WSAEREFUSED", "A database query failed because it was actively refused."}
	WsahostNotFound                                        = &Error{0x00002AF9, "WSAHOST_NOT_FOUND", "No such host is known."}
	WsatryAgain                                            = &Error{0x00002AFA, "WSATRY_AGAIN", "This is usually a temporary error during host name resolution and means that the local server did not receive a response from an authoritative server."}
	WsanoRecovery                                          = &Error{0x00002AFB, "WSANO_RECOVERY", "A nonrecoverable error occurred during a database lookup."}
	WsanoData                                              = &Error{0x00002AFC, "WSANO_DATA", "The requested name is valid, but no data of the requested type was found."}
	WsaQosReceivers                                        = &Error{0x00002AFD, "WSA_QOS_RECEIVERS", "At least one reserve has arrived."}
	WsaQosSenders                                          = &Error{0x00002AFE, "WSA_QOS_SENDERS", "At least one path has arrived."}
	WsaQosNoSenders                                        = &Error{0x00002AFF, "WSA_QOS_NO_SENDERS", "There are no senders."}
	WsaQosNoReceivers                                      = &Error{0x00002B00, "WSA_QOS_NO_RECEIVERS", "There are no receivers."}
	WsaQosRequestConfirmed                                 = &Error{0x00002B01, "WSA_QOS_REQUEST_CONFIRMED", "Reserve has been confirmed."}
	WsaQosAdmissionFailure                                 = &Error{0x00002B02, "WSA_QOS_ADMISSION_FAILURE", "Error due to lack of resources."}
	WsaQosPolicyFailure                                    = &Error{0x00002B03, "WSA_QOS_POLICY_FAILURE", "Rejected for administrative reasons—bad credentials."}
	WsaQosBadStyle                                         = &Error{0x00002B04, "WSA_QOS_BAD_STYLE", "Unknown or conflicting style."}
	WsaQosBadObject                                        = &Error{0x00002B05, "WSA_QOS_BAD_OBJECT", "There is a problem with some part of the filterspec or provider-specific buffer in general."}
	WsaQosTrafficCtrlError                                 = &Error{0x00002B06, "WSA_QOS_TRAFFIC_CTRL_ERROR", "There is a problem with some part of the flowspec."}
	WsaQosGenericError                                     = &Error{0x00002B07, "WSA_QOS_GENERIC_ERROR", "General quality of serve (QOS) error."}
	WsaQosEservicetype                                     = &Error{0x00002B08, "WSA_QOS_ESERVICETYPE", "An invalid or unrecognized service type was found in the flowspec."}
	WsaQosEflowspec                                        = &Error{0x00002B09, "WSA_QOS_EFLOWSPEC", "An invalid or inconsistent flowspec was found in the QOS structure."}
	WsaQosEprovspecbuf                                     = &Error{0x00002B0A, "WSA_QOS_EPROVSPECBUF", "Invalid QOS provider-specific buffer."}
	WsaQosEfilterstyle                                     = &Error{0x00002B0B, "WSA_QOS_EFILTERSTYLE", "An invalid QOS filter style was used."}
	WsaQosEfiltertype                                      = &Error{0x00002B0C, "WSA_QOS_EFILTERTYPE", "An invalid QOS filter type was used."}
	WsaQosEfiltercount                                     = &Error{0x00002B0D, "WSA_QOS_EFILTERCOUNT", "An incorrect number of QOS FILTERSPECs were specified in the FLOWDESCRIPTOR."}
	WsaQosEobjlength                                       = &Error{0x00002B0E, "WSA_QOS_EOBJLENGTH", "An object with an invalid ObjectLength field was specified in the QOS provider-specific buffer."}
	WsaQosEflowcount                                       = &Error{0x00002B0F, "WSA_QOS_EFLOWCOUNT", "An incorrect number of flow descriptors was specified in the QOS structure."}
	WsaQosEunkownpsobj                                     = &Error{0x00002B10, "WSA_QOS_EUNKOWNPSOBJ", "An unrecognized object was found in the QOS provider-specific buffer."}
	WsaQosEpolicyobj                                       = &Error{0x00002B11, "WSA_QOS_EPOLICYOBJ", "An invalid policy object was found in the QOS provider-specific buffer."}
	WsaQosEflowdesc                                        = &Error{0x00002B12, "WSA_QOS_EFLOWDESC", "An invalid QOS flow descriptor was found in the flow descriptor list."}
	WsaQosEpsflowspec                                      = &Error{0x00002B13, "WSA_QOS_EPSFLOWSPEC", "An invalid or inconsistent flowspec was found in the QOS provider-specific buffer."}
	WsaQosEpsfilterspec                                    = &Error{0x00002B14, "WSA_QOS_EPSFILTERSPEC", "An invalid FILTERSPEC was found in the QOS provider-specific buffer."}
	WsaQosEsdmodeobj                                       = &Error{0x00002B15, "WSA_QOS_ESDMODEOBJ", "An invalid shape discard mode object was found in the QOS provider-specific buffer."}
	WsaQosEshaperateobj                                    = &Error{0x00002B16, "WSA_QOS_ESHAPERATEOBJ", "An invalid shaping rate object was found in the QOS provider-specific buffer."}
	WsaQosReservedPetype                                   = &Error{0x00002B17, "WSA_QOS_RESERVED_PETYPE", "A reserved policy element was found in the QOS provider-specific buffer."}
	ErrorIpsecQmPolicyExists                               = &Error{0x000032C8, "ERROR_IPSEC_QM_POLICY_EXISTS", "The specified quick mode policy already exists."}
	ErrorIpsecQmPolicyNotFound                             = &Error{0x000032C9, "ERROR_IPSEC_QM_POLICY_NOT_FOUND", "The specified quick mode policy was not found."}
	ErrorIpsecQmPolicyInUse                                = &Error{0x000032CA, "ERROR_IPSEC_QM_POLICY_IN_USE", "The specified quick mode policy is being used."}
	ErrorIpsecMmPolicyExists                               = &Error{0x000032CB, "ERROR_IPSEC_MM_POLICY_EXISTS", "The specified main mode policy already exists."}
	ErrorIpsecMmPolicyNotFound                             = &Error{0x000032CC, "ERROR_IPSEC_MM_POLICY_NOT_FOUND", "The specified main mode policy was not found."}
	ErrorIpsecMmPolicyInUse                                = &Error{0x000032CD, "ERROR_IPSEC_MM_POLICY_IN_USE", "The specified main mode policy is being used."}
	ErrorIpsecMmFilterExists                               = &Error{0x000032CE, "ERROR_IPSEC_MM_FILTER_EXISTS", "The specified main mode filter already exists."}
	ErrorIpsecMmFilterNotFound                             = &Error{0x000032CF, "ERROR_IPSEC_MM_FILTER_NOT_FOUND", "The specified main mode filter was not found."}
	ErrorIpsecTransportFilterExists                        = &Error{0x000032D0, "ERROR_IPSEC_TRANSPORT_FILTER_EXISTS", "The specified transport mode filter already exists."}
	ErrorIpsecTransportFilterNotFound                      = &Error{0x000032D1, "ERROR_IPSEC_TRANSPORT_FILTER_NOT_FOUND", "The specified transport mode filter does not exist."}
	ErrorIpsecMmAuthExists                                 = &Error{0x000032D2, "ERROR_IPSEC_MM_AUTH_EXISTS", "The specified main mode authentication list exists."}
	ErrorIpsecMmAuthNotFound                               = &Error{0x000032D3, "ERROR_IPSEC_MM_AUTH_NOT_FOUND", "The specified main mode authentication list was not found."}
	ErrorIpsecMmAuthInUse                                  = &Error{0x000032D4, "ERROR_IPSEC_MM_AUTH_IN_USE", "The specified main mode authentication list is being used."}
	ErrorIpsecDefaultMmPolicyNotFound                      = &Error{0x000032D5, "ERROR_IPSEC_DEFAULT_MM_POLICY_NOT_FOUND", "The specified default main mode policy was not found."}
	ErrorIpsecDefaultMmAuthNotFound                        = &Error{0x000032D6, "ERROR_IPSEC_DEFAULT_MM_AUTH_NOT_FOUND", "The specified default main mode authentication list was not found."}
	ErrorIpsecDefaultQmPolicyNotFound                      = &Error{0x000032D7, "ERROR_IPSEC_DEFAULT_QM_POLICY_NOT_FOUND", "The specified default quick mode policy was not found."}
	ErrorIpsecTunnelFilterExists                           = &Error{0x000032D8, "ERROR_IPSEC_TUNNEL_FILTER_EXISTS", "The specified tunnel mode filter exists."}
	ErrorIpsecTunnelFilterNotFound                         = &Error{0x000032D9, "ERROR_IPSEC_TUNNEL_FILTER_NOT_FOUND", "The specified tunnel mode filter was not found."}
	ErrorIpsecMmFilterPendingDeletion                      = &Error{0x000032DA, "ERROR_IPSEC_MM_FILTER_PENDING_DELETION", "The main mode filter is pending deletion."}
	ErrorIpsecTransportFilterEndingDeletion                = &Error{0x000032DB, "ERROR_IPSEC_TRANSPORT_FILTER_ENDING_DELETION", "The transport filter is pending deletion."}
	ErrorIpsecTunnelFilterPendingDeletion                  = &Error{0x000032DC, "ERROR_IPSEC_TUNNEL_FILTER_PENDING_DELETION", "The tunnel filter is pending deletion."}
	ErrorIpsecMmPolicyPendingEletion                       = &Error{0x000032DD, "ERROR_IPSEC_MM_POLICY_PENDING_ELETION", "The main mode policy is pending deletion."}
	ErrorIpsecMmAuthPendingDeletion                        = &Error{0x000032DE, "ERROR_IPSEC_MM_AUTH_PENDING_DELETION", "The main mode authentication bundle is pending deletion."}
	ErrorIpsecQmPolicyPendingDeletion                      = &Error{0x000032DF, "ERROR_IPSEC_QM_POLICY_PENDING_DELETION", "The quick mode policy is pending deletion."}
	WarningIpsecMmPolicyPruned                             = &Error{0x000032E0, "WARNING_IPSEC_MM_POLICY_PRUNED", "The main mode policy was successfully added, but some of the requested offers are not supported."}
	WarningIpsecQmPolicyPruned                             = &Error{0x000032E1, "WARNING_IPSEC_QM_POLICY_PRUNED", "The quick mode policy was successfully added, but some of the requested offers are not supported."}
	ErrorIpsecIkeNegStatusBegin                            = &Error{0x000035E8, "ERROR_IPSEC_IKE_NEG_STATUS_BEGIN", "Starts the list of frequencies of various IKE Win32 error codes encountered during negotiations."}
	ErrorIpsecIkeAuthFail                                  = &Error{0x000035E9, "ERROR_IPSEC_IKE_AUTH_FAIL", "The IKE authentication credentials are unacceptable."}
	ErrorIpsecIkeAttribFail                                = &Error{0x000035EA, "ERROR_IPSEC_IKE_ATTRIB_FAIL", "The IKE security attributes are unacceptable."}
	ErrorIpsecIkeNegotiationPending                        = &Error{0x000035EB, "ERROR_IPSEC_IKE_NEGOTIATION_PENDING", "The IKE negotiation is in progress."}
	ErrorIpsecIkeGeneralProcessingError                    = &Error{0x000035EC, "ERROR_IPSEC_IKE_GENERAL_PROCESSING_ERROR", "General processing error."}
	ErrorIpsecIkeTimedOut                                  = &Error{0x000035ED, "ERROR_IPSEC_IKE_TIMED_OUT", "Negotiation timed out."}
	ErrorIpsecIkeNoCert                                    = &Error{0x000035EE, "ERROR_IPSEC_IKE_NO_CERT", "The IKE failed to find a valid machine certificate. Contact your network security administrator about installing a valid certificate in the appropriate certificate store."}
	ErrorIpsecIkeSaDeleted                                 = &Error{0x000035EF, "ERROR_IPSEC_IKE_SA_DELETED", "The IKE security association (SA) was deleted by a peer before it was completely established."}
	ErrorIpsecIkeSaReaped                                  = &Error{0x000035F0, "ERROR_IPSEC_IKE_SA_REAPED", "The IKE SA was deleted before it was completely established."}
	ErrorIpsecIkeMmAcquireDrop                             = &Error{0x000035F1, "ERROR_IPSEC_IKE_MM_ACQUIRE_DROP", "The negotiation request sat in the queue too long."}
	ErrorIpsecIkeQmAcquireDrop                             = &Error{0x000035F2, "ERROR_IPSEC_IKE_QM_ACQUIRE_DROP", "The negotiation request sat in the queue too long."}
	ErrorIpsecIkeQueueDropMm                               = &Error{0x000035F3, "ERROR_IPSEC_IKE_QUEUE_DROP_MM", "The negotiation request sat in the queue too long."}
	ErrorIpsecIkeQueueDropNoMm                             = &Error{0x000035F4, "ERROR_IPSEC_IKE_QUEUE_DROP_NO_MM", "The negotiation request sat in the queue too long."}
	ErrorIpsecIkeDropNoResponse                            = &Error{0x000035F5, "ERROR_IPSEC_IKE_DROP_NO_RESPONSE", "There was no response from a peer."}
	ErrorIpsecIkeMmDelayDrop                               = &Error{0x000035F6, "ERROR_IPSEC_IKE_MM_DELAY_DROP", "The negotiation took too long."}
	ErrorIpsecIkeQmDelayDrop                               = &Error{0x000035F7, "ERROR_IPSEC_IKE_QM_DELAY_DROP", "The negotiation took too long."}
	ErrorIpsecIkeError                                     = &Error{0x000035F8, "ERROR_IPSEC_IKE_ERROR", "An unknown error occurred."}
	ErrorIpsecIkeCrlFailed                                 = &Error{0x000035F9, "ERROR_IPSEC_IKE_CRL_FAILED", "The certificate revocation check failed."}
	ErrorIpsecIkeInvalidKeyUsage                           = &Error{0x000035FA, "ERROR_IPSEC_IKE_INVALID_KEY_USAGE", "Invalid certificate key usage."}
	ErrorIpsecIkeInvalidCertType                           = &Error{0x000035FB, "ERROR_IPSEC_IKE_INVALID_CERT_TYPE", "Invalid certificate type."}
	ErrorIpsecIkeNoPrivateKey                              = &Error{0x000035FC, "ERROR_IPSEC_IKE_NO_PRIVATE_KEY", "The IKE negotiation failed because the machine certificate used does not have a private key. IPsec certificates require a private key. Contact your network security administrator about a certificate that has a private key."}
	ErrorIpsecIkeDhFail                                    = &Error{0x000035FE, "ERROR_IPSEC_IKE_DH_FAIL", "There was a failure in the Diffie-Hellman computation."}
	ErrorIpsecIkeInvalidHeader                             = &Error{0x00003600, "ERROR_IPSEC_IKE_INVALID_HEADER", "Invalid header."}
	ErrorIpsecIkeNoPolicy                                  = &Error{0x00003601, "ERROR_IPSEC_IKE_NO_POLICY", "No policy configured."}
	ErrorIpsecIkeInvalidSignature                          = &Error{0x00003602, "ERROR_IPSEC_IKE_INVALID_SIGNATURE", "Failed to verify signature."}
	ErrorIpsecIkeKerberosError                             = &Error{0x00003603, "ERROR_IPSEC_IKE_KERBEROS_ERROR", "Failed to authenticate using Kerberos."}
	ErrorIpsecIkeNoPublicKey                               = &Error{0x00003604, "ERROR_IPSEC_IKE_NO_PUBLIC_KEY", "The peer's certificate did not have a public key."}
	ErrorIpsecIkeProcessErr                                = &Error{0x00003605, "ERROR_IPSEC_IKE_PROCESS_ERR", "Error processing the error payload."}
	ErrorIpsecIkeProcessErrSa                              = &Error{0x00003606, "ERROR_IPSEC_IKE_PROCESS_ERR_SA", "Error processing the SA payload."}
	ErrorIpsecIkeProcessErrProp                            = &Error{0x00003607, "ERROR_IPSEC_IKE_PROCESS_ERR_PROP", "Error processing the proposal payload."}
	ErrorIpsecIkeProcessErrTrans                           = &Error{0x00003608, "ERROR_IPSEC_IKE_PROCESS_ERR_TRANS", "Error processing the transform payload."}
	ErrorIpsecIkeProcessErrKe                              = &Error{0x00003609, "ERROR_IPSEC_IKE_PROCESS_ERR_KE", "Error processing the key exchange payload."}
	ErrorIpsecIkeProcessErrId                              = &Error{0x0000360A, "ERROR_IPSEC_IKE_PROCESS_ERR_ID", "Error processing the ID payload."}
	ErrorIpsecIkeProcessErrCert                            = &Error{0x0000360B, "ERROR_IPSEC_IKE_PROCESS_ERR_CERT", "Error processing the certification payload."}
	ErrorIpsecIkeProcessErrCertReq                         = &Error{0x0000360C, "ERROR_IPSEC_IKE_PROCESS_ERR_CERT_REQ", "Error processing the certificate request payload."}
	ErrorIpsecIkeProcessErrHash                            = &Error{0x0000360D, "ERROR_IPSEC_IKE_PROCESS_ERR_HASH", "Error processing the hash payload."}
	ErrorIpsecIkeProcessErrSig                             = &Error{0x0000360E, "ERROR_IPSEC_IKE_PROCESS_ERR_SIG", "Error processing the signature payload."}
	ErrorIpsecIkeProcessErrNonce                           = &Error{0x0000360F, "ERROR_IPSEC_IKE_PROCESS_ERR_NONCE", "Error processing the nonce payload."}
	ErrorIpsecIkeProcessErrNotify                          = &Error{0x00003610, "ERROR_IPSEC_IKE_PROCESS_ERR_NOTIFY", "Error processing the notify payload."}
	ErrorIpsecIkeProcessErrDelete                          = &Error{0x00003611, "ERROR_IPSEC_IKE_PROCESS_ERR_DELETE", "Error processing the delete payload."}
	ErrorIpsecIkeProcessErrVendor                          = &Error{0x00003612, "ERROR_IPSEC_IKE_PROCESS_ERR_VENDOR", "Error processing the VendorId payload."}
	ErrorIpsecIkeInvalidPayload                            = &Error{0x00003613, "ERROR_IPSEC_IKE_INVALID_PAYLOAD", "Invalid payload received."}
	ErrorIpsecIkeLoadSoftSa                                = &Error{0x00003614, "ERROR_IPSEC_IKE_LOAD_SOFT_SA", "Soft SA loaded."}
	ErrorIpsecIkeSoftSaTornDown                            = &Error{0x00003615, "ERROR_IPSEC_IKE_SOFT_SA_TORN_DOWN", "Soft SA torn down."}
	ErrorIpsecIkeInvalidCookie                             = &Error{0x00003616, "ERROR_IPSEC_IKE_INVALID_COOKIE", "Invalid cookie received."}
	ErrorIpsecIkeNoPeerCert                                = &Error{0x00003617, "ERROR_IPSEC_IKE_NO_PEER_CERT", "Peer failed to send valid machine certificate."}
	ErrorIpsecIkePeerCrlFailed                             = &Error{0x00003618, "ERROR_IPSEC_IKE_PEER_CRL_FAILED", "Certification revocation check of peer's certificate failed."}
	ErrorIpsecIkePolicyChange                              = &Error{0x00003619, "ERROR_IPSEC_IKE_POLICY_CHANGE", "New policy invalidated SAs formed with the old policy."}
	ErrorIpsecIkeNoMmPolicy                                = &Error{0x0000361A, "ERROR_IPSEC_IKE_NO_MM_POLICY", "There is no available main mode IKE policy."}
	ErrorIpsecIkeNotcbpriv                                 = &Error{0x0000361B, "ERROR_IPSEC_IKE_NOTCBPRIV", "Failed to enabled trusted computer base (TCB) privilege."}
	ErrorIpsecIkeSecloadfail                               = &Error{0x0000361C, "ERROR_IPSEC_IKE_SECLOADFAIL", "Failed to load SECURITY.DLL."}
	ErrorIpsecIkeFailsspinit                               = &Error{0x0000361D, "ERROR_IPSEC_IKE_FAILSSPINIT", "Failed to obtain the security function table dispatch address from the SSPI."}
	ErrorIpsecIkeFailqueryssp                              = &Error{0x0000361E, "ERROR_IPSEC_IKE_FAILQUERYSSP", "Failed to query the Kerberos package to obtain the max token size."}
	ErrorIpsecIkeSrvacqfail                                = &Error{0x0000361F, "ERROR_IPSEC_IKE_SRVACQFAIL", "Failed to obtain the Kerberos server credentials for the Internet Security Association and Key Management Protocol (ISAKMP)/ERROR_IPSEC_IKE service. Kerberos authentication will not function. The most likely reason for this is lack of domain membership. This is normal if your computer is a member of a workgroup."}
	ErrorIpsecIkeSrvquerycred                              = &Error{0x00003620, "ERROR_IPSEC_IKE_SRVQUERYCRED", "Failed to determine the SSPI principal name for ISAKMP/ERROR_IPSEC_IKE service (QueryCredentialsAttributes)."}
	ErrorIpsecIkeGetspifail                                = &Error{0x00003621, "ERROR_IPSEC_IKE_GETSPIFAIL", "Failed to obtain a new service provider interface (SPI) for the inbound SA from the IPsec driver. The most common cause for this is that the driver does not have the correct filter. Check your policy to verify the filters."}
	ErrorIpsecIkeInvalidFilter                             = &Error{0x00003622, "ERROR_IPSEC_IKE_INVALID_FILTER", "Given filter is invalid."}
	ErrorIpsecIkeOutOfMemory                               = &Error{0x00003623, "ERROR_IPSEC_IKE_OUT_OF_MEMORY", "Memory allocation failed."}
	ErrorIpsecIkeAddUpdateKeyFailed                        = &Error{0x00003624, "ERROR_IPSEC_IKE_ADD_UPDATE_KEY_FAILED", "Failed to add an SA to the IPSec driver. The most common cause for this is if the IKE negotiation took too long to complete. If the problem persists, reduce the load on the faulting machine."}
	ErrorIpsecIkeInvalidPolicy                             = &Error{0x00003625, "ERROR_IPSEC_IKE_INVALID_POLICY", "Invalid policy."}
	ErrorIpsecIkeUnknownDoi                                = &Error{0x00003626, "ERROR_IPSEC_IKE_UNKNOWN_DOI", "Invalid digital object identifier (DOI)."}
	ErrorIpsecIkeInvalidSituation                          = &Error{0x00003627, "ERROR_IPSEC_IKE_INVALID_SITUATION", "Invalid situation."}
	ErrorIpsecIkeDhFailure                                 = &Error{0x00003628, "ERROR_IPSEC_IKE_DH_FAILURE", "Diffie-Hellman failure."}
	ErrorIpsecIkeInvalidGroup                              = &Error{0x00003629, "ERROR_IPSEC_IKE_INVALID_GROUP", "Invalid Diffie-Hellman group."}
	ErrorIpsecIkeEncrypt                                   = &Error{0x0000362A, "ERROR_IPSEC_IKE_ENCRYPT", "Error encrypting payload."}
	ErrorIpsecIkeDecrypt                                   = &Error{0x0000362B, "ERROR_IPSEC_IKE_DECRYPT", "Error decrypting payload."}
	ErrorIpsecIkePolicyMatch                               = &Error{0x0000362C, "ERROR_IPSEC_IKE_POLICY_MATCH", "Policy match error."}
	ErrorIpsecIkeUnsupportedId                             = &Error{0x0000362D, "ERROR_IPSEC_IKE_UNSUPPORTED_ID", "Unsupported ID."}
	ErrorIpsecIkeInvalidHash                               = &Error{0x0000362E, "ERROR_IPSEC_IKE_INVALID_HASH", "Hash verification failed."}
	ErrorIpsecIkeInvalidHashAlg                            = &Error{0x0000362F, "ERROR_IPSEC_IKE_INVALID_HASH_ALG", "Invalid hash algorithm."}
	ErrorIpsecIkeInvalidHashSize                           = &Error{0x00003630, "ERROR_IPSEC_IKE_INVALID_HASH_SIZE", "Invalid hash size."}
	ErrorIpsecIkeInvalidEncryptAlg                         = &Error{0x00003631, "ERROR_IPSEC_IKE_INVALID_ENCRYPT_ALG", "Invalid encryption algorithm."}
	ErrorIpsecIkeInvalidAuthAlg                            = &Error{0x00003632, "ERROR_IPSEC_IKE_INVALID_AUTH_ALG", "Invalid authentication algorithm."}
	ErrorIpsecIkeInvalidSig                                = &Error{0x00003633, "ERROR_IPSEC_IKE_INVALID_SIG", "Invalid certificate signature."}
	ErrorIpsecIkeLoadFailed                                = &Error{0x00003634, "ERROR_IPSEC_IKE_LOAD_FAILED", "Load failed."}
	ErrorIpsecIkeRpcDelete                                 = &Error{0x00003635, "ERROR_IPSEC_IKE_RPC_DELETE", "Deleted by using an RPC call."}
	ErrorIpsecIkeBenignReinit                              = &Error{0x00003636, "ERROR_IPSEC_IKE_BENIGN_REINIT", "A temporary state was created to perform reinitialization. This is not a real failure."}
	ErrorIpsecIkeInvalidResponderLifetimeNotify            = &Error{0x00003637, "ERROR_IPSEC_IKE_INVALID_RESPONDER_LIFETIME_NOTIFY", "The lifetime value received in the Responder Lifetime Notify is below the Windows 2000 configured minimum value. Fix the policy on the peer machine."}
	ErrorIpsecIkeInvalidCertKeylen                         = &Error{0x00003639, "ERROR_IPSEC_IKE_INVALID_CERT_KEYLEN", "Key length in the certificate is too small for configured security requirements."}
	ErrorIpsecIkeMmLimit                                   = &Error{0x0000363A, "ERROR_IPSEC_IKE_MM_LIMIT", "Maximum number of established MM SAs to peer exceeded."}
	ErrorIpsecIkeNegotiationDisabled                       = &Error{0x0000363B, "ERROR_IPSEC_IKE_NEGOTIATION_DISABLED", "The IKE received a policy that disables negotiation."}
	ErrorIpsecIkeQmLimit                                   = &Error{0x0000363C, "ERROR_IPSEC_IKE_QM_LIMIT", "Reached maximum quick mode limit for the main mode. New main mode will be started."}
	ErrorIpsecIkeMmExpired                                 = &Error{0x0000363D, "ERROR_IPSEC_IKE_MM_EXPIRED", "Main mode SA lifetime expired or the peer sent a main mode delete."}
	ErrorIpsecIkePeerMmAssumedInvalid                      = &Error{0x0000363E, "ERROR_IPSEC_IKE_PEER_MM_ASSUMED_INVALID", "Main mode SA assumed to be invalid because peer stopped responding."}
	ErrorIpsecIkeCertChainPolicyMismatch                   = &Error{0x0000363F, "ERROR_IPSEC_IKE_CERT_CHAIN_POLICY_MISMATCH", "Certificate does not chain to a trusted root in IPsec policy."}
	ErrorIpsecIkeUnexpectedMessageId                       = &Error{0x00003640, "ERROR_IPSEC_IKE_UNEXPECTED_MESSAGE_ID", "Received unexpected message ID."}
	ErrorIpsecIkeInvalidUmatts                             = &Error{0x00003641, "ERROR_IPSEC_IKE_INVALID_UMATTS", "Received invalid AuthIP user mode attributes."}
	ErrorIpsecIkeDosCookieSent                             = &Error{0x00003642, "ERROR_IPSEC_IKE_DOS_COOKIE_SENT", "Sent DOS cookie notify to initiator."}
	ErrorIpsecIkeShuttingDown                              = &Error{0x00003643, "ERROR_IPSEC_IKE_SHUTTING_DOWN", "The IKE service is shutting down."}
	ErrorIpsecIkeCgaAuthFailed                             = &Error{0x00003644, "ERROR_IPSEC_IKE_CGA_AUTH_FAILED", "Could not verify the binding between the color graphics adapter (CGA) address and the certificate."}
	ErrorIpsecIkeProcessErrNatoa                           = &Error{0x00003645, "ERROR_IPSEC_IKE_PROCESS_ERR_NATOA", "Error processing the NatOA payload."}
	ErrorIpsecIkeInvalidMmForQm                            = &Error{0x00003646, "ERROR_IPSEC_IKE_INVALID_MM_FOR_QM", "The parameters of the main mode are invalid for this quick mode."}
	ErrorIpsecIkeQmExpired                                 = &Error{0x00003647, "ERROR_IPSEC_IKE_QM_EXPIRED", "The quick mode SA was expired by the IPsec driver."}
	ErrorIpsecIkeTooManyFilters                            = &Error{0x00003648, "ERROR_IPSEC_IKE_TOO_MANY_FILTERS", "Too many dynamically added IKEEXT filters were detected."}
	ErrorIpsecIkeNegStatusEnd                              = &Error{0x00003649, "ERROR_IPSEC_IKE_NEG_STATUS_END", "Ends the list of frequencies of various IKE Win32 error codes encountered during negotiations."}
	ErrorSxsSectionNotFound                                = &Error{0x000036B0, "ERROR_SXS_SECTION_NOT_FOUND", "The requested section was not present in the activation context."}
	ErrorSxsCantGenActctx                                  = &Error{0x000036B1, "ERROR_SXS_CANT_GEN_ACTCTX", "The application has failed to start because its side-by-side configuration is incorrect. See the application event log for more detail."}
	ErrorSxsInvalidActctxdataFormat                        = &Error{0x000036B2, "ERROR_SXS_INVALID_ACTCTXDATA_FORMAT", "The application binding data format is invalid."}
	ErrorSxsAssemblyNotFound                               = &Error{0x000036B3, "ERROR_SXS_ASSEMBLY_NOT_FOUND", "The referenced assembly is not installed on your system."}
	ErrorSxsManifestFormatError                            = &Error{0x000036B4, "ERROR_SXS_MANIFEST_FORMAT_ERROR", "The manifest file does not begin with the required tag and format information."}
	ErrorSxsManifestParseError                             = &Error{0x000036B5, "ERROR_SXS_MANIFEST_PARSE_ERROR", "The manifest file contains one or more syntax errors."}
	ErrorSxsActivationContextDisabled                      = &Error{0x000036B6, "ERROR_SXS_ACTIVATION_CONTEXT_DISABLED", "The application attempted to activate a disabled activation context."}
	ErrorSxsKeyNotFound                                    = &Error{0x000036B7, "ERROR_SXS_KEY_NOT_FOUND", "The requested lookup key was not found in any active activation context."}
	ErrorSxsVersionConflict                                = &Error{0x000036B8, "ERROR_SXS_VERSION_CONFLICT", "A component version required by the application conflicts with another active component version."}
	ErrorSxsWrongSectionType                               = &Error{0x000036B9, "ERROR_SXS_WRONG_SECTION_TYPE", "The type requested activation context section does not match the query API used."}
	ErrorSxsThreadQueriesDisabled                          = &Error{0x000036BA, "ERROR_SXS_THREAD_QUERIES_DISABLED", "Lack of system resources has required isolated activation to be disabled for the current thread of execution."}
	ErrorSxsProcessDefaultAlreadySet                       = &Error{0x000036BB, "ERROR_SXS_PROCESS_DEFAULT_ALREADY_SET", "An attempt to set the process default activation context failed because the process default activation context was already set."}
	ErrorSxsUnknownEncodingGroup                           = &Error{0x000036BC, "ERROR_SXS_UNKNOWN_ENCODING_GROUP", "The encoding group identifier specified is not recognized."}
	ErrorSxsUnknownEncoding                                = &Error{0x000036BD, "ERROR_SXS_UNKNOWN_ENCODING", "The encoding requested is not recognized."}
	ErrorSxsInvalidXmlNamespaceUri                         = &Error{0x000036BE, "ERROR_SXS_INVALID_XML_NAMESPACE_URI", "The manifest contains a reference to an invalid URI."}
	ErrorSxsRootManifestDependencyOtInstalled              = &Error{0x000036BF, "ERROR_SXS_ROOT_MANIFEST_DEPENDENCY_OT_INSTALLED", "The application manifest contains a reference to a dependent assembly that is not installed."}
	ErrorSxsLeafManifestDependencyNotInstalled             = &Error{0x000036C0, "ERROR_SXS_LEAF_MANIFEST_DEPENDENCY_NOT_INSTALLED", "The manifest for an assembly used by the application has a reference to a dependent assembly that is not installed."}
	ErrorSxsInvalidAssemblyIdentityAttribute               = &Error{0x000036C1, "ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE", "The manifest contains an attribute for the assembly identity that is not valid."}
	ErrorSxsManifestMissingRequiredDefaultNamespace        = &Error{0x000036C2, "ERROR_SXS_MANIFEST_MISSING_REQUIRED_DEFAULT_NAMESPACE", "The manifest is missing the required default namespace specification on the assembly element."}
	ErrorSxsManifestInvalidRequiredDefaultNamespace        = &Error{0x000036C3, "ERROR_SXS_MANIFEST_INVALID_REQUIRED_DEFAULT_NAMESPACE", "The manifest has a default namespace specified on the assembly element but its value is not urn:schemas-microsoft-com:asm.v1\".\""}
	ErrorSxsPrivateManifestCrossPathWithReparsePoint       = &Error{0x000036C4, "ERROR_SXS_PRIVATE_MANIFEST_CROSS_PATH_WITH_REPARSE_POINT", "The private manifest probed has crossed the reparse-point-associated path."}
	ErrorSxsDuplicateDllName                               = &Error{0x000036C5, "ERROR_SXS_DUPLICATE_DLL_NAME", "Two or more components referenced directly or indirectly by the application manifest have files by the same name."}
	ErrorSxsDuplicateWindowclassName                       = &Error{0x000036C6, "ERROR_SXS_DUPLICATE_WINDOWCLASS_NAME", "Two or more components referenced directly or indirectly by the application manifest have window classes with the same name."}
	ErrorSxsDuplicateClsid                                 = &Error{0x000036C7, "ERROR_SXS_DUPLICATE_CLSID", "Two or more components referenced directly or indirectly by the application manifest have the same COM server CLSIDs."}
	ErrorSxsDuplicateIid                                   = &Error{0x000036C8, "ERROR_SXS_DUPLICATE_IID", "Two or more components referenced directly or indirectly by the application manifest have proxies for the same COM interface IIDs."}
	ErrorSxsDuplicateTlbid                                 = &Error{0x000036C9, "ERROR_SXS_DUPLICATE_TLBID", "Two or more components referenced directly or indirectly by the application manifest have the same COM type library TLBIDs."}
	ErrorSxsDuplicateProgid                                = &Error{0x000036CA, "ERROR_SXS_DUPLICATE_PROGID", "Two or more components referenced directly or indirectly by the application manifest have the same COM ProgIDs."}
	ErrorSxsDuplicateAssemblyName                          = &Error{0x000036CB, "ERROR_SXS_DUPLICATE_ASSEMBLY_NAME", "Two or more components referenced directly or indirectly by the application manifest are different versions of the same component, which is not permitted."}
	ErrorSxsFileHashMismatch                               = &Error{0x000036CC, "ERROR_SXS_FILE_HASH_MISMATCH", "A component's file does not match the verification information present in the component manifest."}
	ErrorSxsPolicyParseError                               = &Error{0x000036CD, "ERROR_SXS_POLICY_PARSE_ERROR", "The policy manifest contains one or more syntax errors."}
	ErrorSxsXmlEMissingquote                               = &Error{0x000036CE, "ERROR_SXS_XML_E_MISSINGQUOTE", "Manifest Parse Error: A string literal was expected, but no opening quotation mark was found."}
	ErrorSxsXmlECommentsyntax                              = &Error{0x000036CF, "ERROR_SXS_XML_E_COMMENTSYNTAX", "Manifest Parse Error: Incorrect syntax was used in a comment."}
	ErrorSxsXmlEBadstartnamechar                           = &Error{0x000036D0, "ERROR_SXS_XML_E_BADSTARTNAMECHAR", "Manifest Parse Error: A name started with an invalid character."}
	ErrorSxsXmlEBadnamechar                                = &Error{0x000036D1, "ERROR_SXS_XML_E_BADNAMECHAR", "Manifest Parse Error: A name contained an invalid character."}
	ErrorSxsXmlEBadcharinstring                            = &Error{0x000036D2, "ERROR_SXS_XML_E_BADCHARINSTRING", "Manifest Parse Error: A string literal contained an invalid character."}
	ErrorSxsXmlEXmldeclsyntax                              = &Error{0x000036D3, "ERROR_SXS_XML_E_XMLDECLSYNTAX", "Manifest Parse Error: Invalid syntax for an XML declaration."}
	ErrorSxsXmlEBadchardata                                = &Error{0x000036D4, "ERROR_SXS_XML_E_BADCHARDATA", "Manifest Parse Error: An Invalid character was found in text content."}
	ErrorSxsXmlEMissingwhitespace                          = &Error{0x000036D5, "ERROR_SXS_XML_E_MISSINGWHITESPACE", "Manifest Parse Error: Required white space was missing."}
	ErrorSxsXmlEExpectingtagend                            = &Error{0x000036D6, "ERROR_SXS_XML_E_EXPECTINGTAGEND", "Manifest Parse Error: The angle bracket (>) character was expected."}
	ErrorSxsXmlEMissingsemicolon                           = &Error{0x000036D7, "ERROR_SXS_XML_E_MISSINGSEMICOLON", "Manifest Parse Error: A semicolon (;) was expected."}
	ErrorSxsXmlEUnbalancedparen                            = &Error{0x000036D8, "ERROR_SXS_XML_E_UNBALANCEDPAREN", "Manifest Parse Error: Unbalanced parentheses."}
	ErrorSxsXmlEInternalerror                              = &Error{0x000036D9, "ERROR_SXS_XML_E_INTERNALERROR", "Manifest Parse Error: Internal error."}
	ErrorSxsXmlEUnexpectedWhitespace                       = &Error{0x000036DA, "ERROR_SXS_XML_E_UNEXPECTED_WHITESPACE", "Manifest Parse Error: Whitespace is not allowed at this location."}
	ErrorSxsXmlEIncompleteEncoding                         = &Error{0x000036DB, "ERROR_SXS_XML_E_INCOMPLETE_ENCODING", "Manifest Parse Error: End of file reached in invalid state for current encoding."}
	ErrorSxsXmlEMissingParen                               = &Error{0x000036DC, "ERROR_SXS_XML_E_MISSING_PAREN", "Manifest Parse Error: Missing parenthesis."}
	ErrorSxsXmlEExpectingclosequote                        = &Error{0x000036DD, "ERROR_SXS_XML_E_EXPECTINGCLOSEQUOTE", "Manifest Parse Error: A single (') or double (\") quotation mark is missing."}
	ErrorSxsXmlEMultipleColons                             = &Error{0x000036DE, "ERROR_SXS_XML_E_MULTIPLE_COLONS", "Manifest Parse Error: Multiple colons are not allowed in a name."}
	ErrorSxsXmlEInvalidDecimal                             = &Error{0x000036DF, "ERROR_SXS_XML_E_INVALID_DECIMAL", "Manifest Parse Error: Invalid character for decimal digit."}
	ErrorSxsXmlEInvalidHexidecimal                         = &Error{0x000036E0, "ERROR_SXS_XML_E_INVALID_HEXIDECIMAL", "Manifest Parse Error: Invalid character for hexadecimal digit."}
	ErrorSxsXmlEInvalidUnicode                             = &Error{0x000036E1, "ERROR_SXS_XML_E_INVALID_UNICODE", "Manifest Parse Error: Invalid Unicode character value for this platform."}
	ErrorSxsXmlEWhitespaceorquestionmark                   = &Error{0x000036E2, "ERROR_SXS_XML_E_WHITESPACEORQUESTIONMARK", "Manifest Parse Error: Expecting whitespace or question mark (?)."}
	ErrorSxsXmlEUnexpectedendtag                           = &Error{0x000036E3, "ERROR_SXS_XML_E_UNEXPECTEDENDTAG", "Manifest Parse Error: End tag was not expected at this location."}
	ErrorSxsXmlEUnclosedtag                                = &Error{0x000036E4, "ERROR_SXS_XML_E_UNCLOSEDTAG", "Manifest Parse Error: The following tags were not closed: %1."}
	ErrorSxsXmlEDuplicateattribute                         = &Error{0x000036E5, "ERROR_SXS_XML_E_DUPLICATEATTRIBUTE", "Manifest Parse Error: Duplicate attribute."}
	ErrorSxsXmlEMultipleroots                              = &Error{0x000036E6, "ERROR_SXS_XML_E_MULTIPLEROOTS", "Manifest Parse Error: Only one top-level element is allowed in an XML document."}
	ErrorSxsXmlEInvalidatrootlevel                         = &Error{0x000036E7, "ERROR_SXS_XML_E_INVALIDATROOTLEVEL", "Manifest Parse Error: Invalid at the top level of the document."}
	ErrorSxsXmlEBadxmldecl                                 = &Error{0x000036E8, "ERROR_SXS_XML_E_BADXMLDECL", "Manifest Parse Error: Invalid XML declaration."}
	ErrorSxsXmlEMissingroot                                = &Error{0x000036E9, "ERROR_SXS_XML_E_MISSINGROOT", "Manifest Parse Error: XML document must have a top-level element."}
	ErrorSxsXmlEUnexpectedeof                              = &Error{0x000036EA, "ERROR_SXS_XML_E_UNEXPECTEDEOF", "Manifest Parse Error: Unexpected end of file."}
	ErrorSxsXmlEBadperefinsubset                           = &Error{0x000036EB, "ERROR_SXS_XML_E_BADPEREFINSUBSET", "Manifest Parse Error: Parameter entities cannot be used inside markup declarations in an internal subset."}
	ErrorSxsXmlEUnclosedstarttag                           = &Error{0x000036EC, "ERROR_SXS_XML_E_UNCLOSEDSTARTTAG", "Manifest Parse Error: Element was not closed."}
	ErrorSxsXmlEUnclosedendtag                             = &Error{0x000036ED, "ERROR_SXS_XML_E_UNCLOSEDENDTAG", "Manifest Parse Error: End element was missing the angle bracket (>) character."}
	ErrorSxsXmlEUnclosedstring                             = &Error{0x000036EE, "ERROR_SXS_XML_E_UNCLOSEDSTRING", "Manifest Parse Error: A string literal was not closed."}
	ErrorSxsXmlEUnclosedcomment                            = &Error{0x000036EF, "ERROR_SXS_XML_E_UNCLOSEDCOMMENT", "Manifest Parse Error: A comment was not closed."}
	ErrorSxsXmlEUncloseddecl                               = &Error{0x000036F0, "ERROR_SXS_XML_E_UNCLOSEDDECL", "Manifest Parse Error: A declaration was not closed."}
	ErrorSxsXmlEUnclosedcdata                              = &Error{0x000036F1, "ERROR_SXS_XML_E_UNCLOSEDCDATA", "Manifest Parse Error: A CDATA section was not closed."}
	ErrorSxsXmlEReservednamespace                          = &Error{0x000036F2, "ERROR_SXS_XML_E_RESERVEDNAMESPACE", "Manifest Parse Error: The namespace prefix is not allowed to start with the reserved string xml\".\""}
	ErrorSxsXmlEInvalidencoding                            = &Error{0x000036F3, "ERROR_SXS_XML_E_INVALIDENCODING", "Manifest Parse Error: System does not support the specified encoding."}
	ErrorSxsXmlEInvalidswitch                              = &Error{0x000036F4, "ERROR_SXS_XML_E_INVALIDSWITCH", "Manifest Parse Error: Switch from current encoding to specified encoding not supported."}
	ErrorSxsXmlEBadxmlcase                                 = &Error{0x000036F5, "ERROR_SXS_XML_E_BADXMLCASE", "Manifest Parse Error: The name \"xml\" is reserved and must be lowercase."}
	ErrorSxsXmlEInvalidStandalone                          = &Error{0x000036F6, "ERROR_SXS_XML_E_INVALID_STANDALONE", "Manifest Parse Error: The stand-alone attribute must have the value \"yes\" or \"no\"."}
	ErrorSxsXmlEUnexpectedStandalone                       = &Error{0x000036F7, "ERROR_SXS_XML_E_UNEXPECTED_STANDALONE", "Manifest Parse Error: The stand-alone attribute cannot be used in external entities."}
	ErrorSxsXmlEInvalidVersion                             = &Error{0x000036F8, "ERROR_SXS_XML_E_INVALID_VERSION", "Manifest Parse Error: Invalid version number."}
	ErrorSxsXmlEMissingequals                              = &Error{0x000036F9, "ERROR_SXS_XML_E_MISSINGEQUALS", "Manifest Parse Error: Missing equal sign (=) between the attribute and the attribute value."}
	ErrorSxsProtectionRecoveryFailed                       = &Error{0x000036FA, "ERROR_SXS_PROTECTION_RECOVERY_FAILED", "Assembly Protection Error: Unable to recover the specified assembly."}
	ErrorSxsProtectionPublicKeyOoShort                     = &Error{0x000036FB, "ERROR_SXS_PROTECTION_PUBLIC_KEY_OO_SHORT", "Assembly Protection Error: The public key for an assembly was too short to be allowed."}
	ErrorSxsProtectionCatalogNotValid                      = &Error{0x000036FC, "ERROR_SXS_PROTECTION_CATALOG_NOT_VALID", "Assembly Protection Error: The catalog for an assembly is not valid, or does not match the assembly's manifest."}
	ErrorSxsUntranslatableHresult                          = &Error{0x000036FD, "ERROR_SXS_UNTRANSLATABLE_HRESULT", "An HRESULT could not be translated to a corresponding Win32 error code."}
	ErrorSxsProtectionCatalogFileMissing                   = &Error{0x000036FE, "ERROR_SXS_PROTECTION_CATALOG_FILE_MISSING", "Assembly Protection Error: The catalog for an assembly is missing."}
	ErrorSxsMissingAssemblyIdentityAttribute               = &Error{0x000036FF, "ERROR_SXS_MISSING_ASSEMBLY_IDENTITY_ATTRIBUTE", "The supplied assembly identity is missing one or more attributes that must be present in this context."}
	ErrorSxsInvalidAssemblyIdentityAttributeName           = &Error{0x00003700, "ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE_NAME", "The supplied assembly identity has one or more attribute names that contain characters not permitted in XML names."}
	ErrorSxsAssemblyMissing                                = &Error{0x00003701, "ERROR_SXS_ASSEMBLY_MISSING", "The referenced assembly could not be found."}
	ErrorSxsCorruptActivationStack                         = &Error{0x00003702, "ERROR_SXS_CORRUPT_ACTIVATION_STACK", "The activation context activation stack for the running thread of execution is corrupt."}
	ErrorSxsCorruption                                     = &Error{0x00003703, "ERROR_SXS_CORRUPTION", "The application isolation metadata for this process or thread has become corrupt."}
	ErrorSxsEarlyDeactivation                              = &Error{0x00003704, "ERROR_SXS_EARLY_DEACTIVATION", "The activation context being deactivated is not the most recently activated one."}
	ErrorSxsInvalidDeactivation                            = &Error{0x00003705, "ERROR_SXS_INVALID_DEACTIVATION", "The activation context being deactivated is not active for the current thread of execution."}
	ErrorSxsMultipleDeactivation                           = &Error{0x00003706, "ERROR_SXS_MULTIPLE_DEACTIVATION", "The activation context being deactivated has already been deactivated."}
	ErrorSxsProcessTerminationRequested                    = &Error{0x00003707, "ERROR_SXS_PROCESS_TERMINATION_REQUESTED", "A component used by the isolation facility has requested to terminate the process."}
	ErrorSxsReleaseActivationOntext                        = &Error{0x00003708, "ERROR_SXS_RELEASE_ACTIVATION_ONTEXT", "A kernel mode component is releasing a reference on an activation context."}
	ErrorSxsSystemDefaultActivationContextEmpty            = &Error{0x00003709, "ERROR_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY", "The activation context of the system default assembly could not be generated."}
	ErrorSxsInvalidIdentityAttributeValue                  = &Error{0x0000370A, "ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE", "The value of an attribute in an identity is not within the legal range."}
	ErrorSxsInvalidIdentityAttributeName                   = &Error{0x0000370B, "ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME", "The name of an attribute in an identity is not within the legal range."}
	ErrorSxsIdentityDuplicateAttribute                     = &Error{0x0000370C, "ERROR_SXS_IDENTITY_DUPLICATE_ATTRIBUTE", "An identity contains two definitions for the same attribute."}
	ErrorSxsIdentityParseError                             = &Error{0x0000370D, "ERROR_SXS_IDENTITY_PARSE_ERROR", "The identity string is malformed. This might be due to a trailing comma, more than two unnamed attributes, a missing attribute name, or a missing attribute value."}
	ErrorMalformedSubstitutionString                       = &Error{0x0000370E, "ERROR_MALFORMED_SUBSTITUTION_STRING", "A string containing localized substitutable content was malformed. Either a dollar sign ($) was followed by something other than a left parenthesis or another dollar sign, or a substitution's right parenthesis was not found."}
	ErrorSxsIncorrectPublicKeyOken                         = &Error{0x0000370F, "ERROR_SXS_INCORRECT_PUBLIC_KEY_OKEN", "The public key token does not correspond to the public key specified."}
	ErrorUnmappedSubstitutionString                        = &Error{0x00003710, "ERROR_UNMAPPED_SUBSTITUTION_STRING", "A substitution string had no mapping."}
	ErrorSxsAssemblyNotLocked                              = &Error{0x00003711, "ERROR_SXS_ASSEMBLY_NOT_LOCKED", "The component must be locked before making the request."}
	ErrorSxsComponentStoreCorrupt                          = &Error{0x00003712, "ERROR_SXS_COMPONENT_STORE_CORRUPT", "The component store has been corrupted."}
	ErrorAdvancedInstallerFailed                           = &Error{0x00003713, "ERROR_ADVANCED_INSTALLER_FAILED", "An advanced installer failed during setup or servicing."}
	ErrorXmlEncodingMismatch                               = &Error{0x00003714, "ERROR_XML_ENCODING_MISMATCH", "The character encoding in the XML declaration did not match the encoding used in the document."}
	ErrorSxsManifestIdentitySameButContentsDifferent       = &Error{0x00003715, "ERROR_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT", "The identities of the manifests are identical, but the contents are different."}
	ErrorSxsIdentitiesDifferent                            = &Error{0x00003716, "ERROR_SXS_IDENTITIES_DIFFERENT", "The component identities are different."}
	ErrorSxsAssemblyIsNotADeployment                       = &Error{0x00003717, "ERROR_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT", "The assembly is not a deployment."}
	ErrorSxsFileNotPartOfAssembly                          = &Error{0x00003718, "ERROR_SXS_FILE_NOT_PART_OF_ASSEMBLY", "The file is not a part of the assembly."}
	ErrorSxsManifestTooBig                                 = &Error{0x00003719, "ERROR_SXS_MANIFEST_TOO_BIG", "The size of the manifest exceeds the maximum allowed."}
	ErrorSxsSettingNotRegistered                           = &Error{0x0000371A, "ERROR_SXS_SETTING_NOT_REGISTERED", "The setting is not registered."}
	ErrorSxsTransactionClosureIncomplete                   = &Error{0x0000371B, "ERROR_SXS_TRANSACTION_CLOSURE_INCOMPLETE", "One or more required members of the transaction are not present."}
	ErrorEvtInvalidChannelPath                             = &Error{0x00003A98, "ERROR_EVT_INVALID_CHANNEL_PATH", "The specified channel path is invalid."}
	ErrorEvtInvalidQuery                                   = &Error{0x00003A99, "ERROR_EVT_INVALID_QUERY", "The specified query is invalid."}
	ErrorEvtPublisherMetadataNotFound                      = &Error{0x00003A9A, "ERROR_EVT_PUBLISHER_METADATA_NOT_FOUND", "The publisher metadata cannot be found in the resource."}
	ErrorEvtEventTemplateNotFound                          = &Error{0x00003A9B, "ERROR_EVT_EVENT_TEMPLATE_NOT_FOUND", "The template for an event definition cannot be found in the resource (error = %1)."}
	ErrorEvtInvalidPublisherName                           = &Error{0x00003A9C, "ERROR_EVT_INVALID_PUBLISHER_NAME", "The specified publisher name is invalid."}
	ErrorEvtInvalidEventData                               = &Error{0x00003A9D, "ERROR_EVT_INVALID_EVENT_DATA", "The event data raised by the publisher is not compatible with the event template definition in the publisher's manifest."}
	ErrorEvtChannelNotFound                                = &Error{0x00003A9F, "ERROR_EVT_CHANNEL_NOT_FOUND", "The specified channel could not be found. Check channel configuration."}
	ErrorEvtMalformedXmlText                               = &Error{0x00003AA0, "ERROR_EVT_MALFORMED_XML_TEXT", "The specified XML text was not well-formed. See extended error for more details."}
	ErrorEvtSubscriptionToDirectChannel                    = &Error{0x00003AA1, "ERROR_EVT_SUBSCRIPTION_TO_DIRECT_CHANNEL", "The caller is trying to subscribe to a direct channel which is not allowed. The events for a direct channel go directly to a log file and cannot be subscribed to."}
	ErrorEvtConfigurationError                             = &Error{0x00003AA2, "ERROR_EVT_CONFIGURATION_ERROR", "Configuration error."}
	ErrorEvtQueryResultStale                               = &Error{0x00003AA3, "ERROR_EVT_QUERY_RESULT_STALE", "The query result is stale or invalid. This might be due to the log being cleared or rolling over after the query result was created. Users should handle this code by releasing the query result object and reissuing the query."}
	ErrorEvtQueryResultInvalidPosition                     = &Error{0x00003AA4, "ERROR_EVT_QUERY_RESULT_INVALID_POSITION", "Query result is currently at an invalid position."}
	ErrorEvtNonValidatingMsxml                             = &Error{0x00003AA5, "ERROR_EVT_NON_VALIDATING_MSXML", "Registered Microsoft XML (MSXML) does not support validation."}
	ErrorEvtFilterAlreadyscoped                            = &Error{0x00003AA6, "ERROR_EVT_FILTER_ALREADYSCOPED", "An expression can only be followed by a change-of-scope operation if it itself evaluates to a node set and is not already part of some other change-of-scope operation."}
	ErrorEvtFilterNoteltset                                = &Error{0x00003AA7, "ERROR_EVT_FILTER_NOTELTSET", "Cannot perform a step operation from a term that does not represent an element set."}
	ErrorEvtFilterInvarg                                   = &Error{0x00003AA8, "ERROR_EVT_FILTER_INVARG", "Left side arguments to binary operators must be either attributes, nodes, or variables and right side arguments must be constants."}
	ErrorEvtFilterInvtest                                  = &Error{0x00003AA9, "ERROR_EVT_FILTER_INVTEST", "A step operation must involve either a node test or, in the case of a predicate, an algebraic expression against which to test each node in the node set identified by the preceding node set can be evaluated."}
	ErrorEvtFilterInvtype                                  = &Error{0x00003AAA, "ERROR_EVT_FILTER_INVTYPE", "This data type is currently unsupported."}
	ErrorEvtFilterParseerr                                 = &Error{0x00003AAB, "ERROR_EVT_FILTER_PARSEERR", "A syntax error occurred at position %1!d!"}
	ErrorEvtFilterUnsupportedop                            = &Error{0x00003AAC, "ERROR_EVT_FILTER_UNSUPPORTEDOP", "This operator is unsupported by this implementation of the filter."}
	ErrorEvtFilterUnexpectedtoken                          = &Error{0x00003AAD, "ERROR_EVT_FILTER_UNEXPECTEDTOKEN", "The token encountered was unexpected."}
	ErrorEvtInvalidOperationOverEnabledDirectChannel       = &Error{0x00003AAE, "ERROR_EVT_INVALID_OPERATION_OVER_ENABLED_DIRECT_CHANNEL", "The requested operation cannot be performed over an enabled direct channel. The channel must first be disabled before performing the requested operation."}
	ErrorEvtInvalidChannelPropertyValue                    = &Error{0x00003AAF, "ERROR_EVT_INVALID_CHANNEL_PROPERTY_VALUE", "Channel property %1!s! contains an invalid value. The value has an invalid type, is outside the valid range, cannot be updated, or is not supported by this type of channel."}
	ErrorEvtInvalidPublisherPropertyValue                  = &Error{0x00003AB0, "ERROR_EVT_INVALID_PUBLISHER_PROPERTY_VALUE", "Publisher property %1!s! contains an invalid value. The value has an invalid type, is outside the valid range, cannot be updated, or is not supported by this type of publisher."}
	ErrorEvtChannelCannotActivate                          = &Error{0x00003AB1, "ERROR_EVT_CHANNEL_CANNOT_ACTIVATE", "The channel fails to activate."}
	ErrorEvtFilterTooComplex                               = &Error{0x00003AB2, "ERROR_EVT_FILTER_TOO_COMPLEX", "The xpath expression exceeded supported complexity. Simplify it or split it into two or more simple expressions."}
	ErrorEvtMessageNotFound                                = &Error{0x00003AB3, "ERROR_EVT_MESSAGE_NOT_FOUND", "The message resource is present but the message is not found in the string or message table."}
	ErrorEvtMessageIdNotFound                              = &Error{0x00003AB4, "ERROR_EVT_MESSAGE_ID_NOT_FOUND", "The message ID for the desired message could not be found."}
	ErrorEvtUnresolvedValueInsert                          = &Error{0x00003AB5, "ERROR_EVT_UNRESOLVED_VALUE_INSERT", "The substitution string for the insert index (%1) could not be found."}
	ErrorEvtUnresolvedParameterInsert                      = &Error{0x00003AB6, "ERROR_EVT_UNRESOLVED_PARAMETER_INSERT", "The description string for the parameter reference (%1) could not be found."}
	ErrorEvtMaxInsertsReached                              = &Error{0x00003AB7, "ERROR_EVT_MAX_INSERTS_REACHED", "The maximum number of replacements has been reached."}
	ErrorEvtEventDefinitionNotOund                         = &Error{0x00003AB8, "ERROR_EVT_EVENT_DEFINITION_NOT_OUND", "The event definition could not be found for the event ID (%1)."}
	ErrorEvtMessageLocaleNotFound                          = &Error{0x00003AB9, "ERROR_EVT_MESSAGE_LOCALE_NOT_FOUND", "The locale-specific resource for the desired message is not present."}
	ErrorEvtVersionTooOld                                  = &Error{0x00003ABA, "ERROR_EVT_VERSION_TOO_OLD", "The resource is too old to be compatible."}
	ErrorEvtVersionTooNew                                  = &Error{0x00003ABB, "ERROR_EVT_VERSION_TOO_NEW", "The resource is too new to be compatible."}
	ErrorEvtCannotOpenChannelOfQuery                       = &Error{0x00003ABC, "ERROR_EVT_CANNOT_OPEN_CHANNEL_OF_QUERY", "The channel at index %1 of the query cannot be opened."}
	ErrorEvtPublisherDisabled                              = &Error{0x00003ABD, "ERROR_EVT_PUBLISHER_DISABLED", "The publisher has been disabled and its resource is not available. This usually occurs when the publisher is in the process of being uninstalled or upgraded."}
	ErrorEcSubscriptionCannotActivate                      = &Error{0x00003AE8, "ERROR_EC_SUBSCRIPTION_CANNOT_ACTIVATE", "The subscription fails to activate."}
	ErrorEcLogDisabled                                     = &Error{0x00003AE9, "ERROR_EC_LOG_DISABLED", "The log of the subscription is in a disabled state and events cannot be forwarded to it. The log must first be enabled before the subscription can be activated."}
	ErrorMuiFileNotFound                                   = &Error{0x00003AFC, "ERROR_MUI_FILE_NOT_FOUND", "The resource loader failed to find the Multilingual User Interface (MUI) file."}
	ErrorMuiInvalidFile                                    = &Error{0x00003AFD, "ERROR_MUI_INVALID_FILE", "The resource loader failed to load the MUI file because the file failed to pass validation."}
	ErrorMuiInvalidRcConfig                                = &Error{0x00003AFE, "ERROR_MUI_INVALID_RC_CONFIG", "The release candidate (RC) manifest is corrupted with garbage data, is an unsupported version, or is missing a required item."}
	ErrorMuiInvalidLocaleName                              = &Error{0x00003AFF, "ERROR_MUI_INVALID_LOCALE_NAME", "The RC manifest has an invalid culture name."}
	ErrorMuiInvalidUltimatefallbackName                    = &Error{0x00003B00, "ERROR_MUI_INVALID_ULTIMATEFALLBACK_NAME", "The RC Manifest has an invalid ultimate fallback name."}
	ErrorMuiFileNotLoaded                                  = &Error{0x00003B01, "ERROR_MUI_FILE_NOT_LOADED", "The resource loader cache does not have a loaded MUI entry."}
	ErrorResourceEnumUserStop                              = &Error{0x00003B02, "ERROR_RESOURCE_ENUM_USER_STOP", "The user stopped resource enumeration."}
	ErrorMuiIntlsettingsUilangNotInstalled                 = &Error{0x00003B03, "ERROR_MUI_INTLSETTINGS_UILANG_NOT_INSTALLED", "User interface language installation failed."}
	ErrorMuiIntlsettingsInvalidLocaleName                  = &Error{0x00003B04, "ERROR_MUI_INTLSETTINGS_INVALID_LOCALE_NAME", "Locale installation failed."}
	ErrorMcaInvalidCapabilitiesString                      = &Error{0x00003B60, "ERROR_MCA_INVALID_CAPABILITIES_STRING", "The monitor returned a DDC/CI capabilities string that did not comply with the ACCESS.bus 3.0, DDC/CI 1.1, or MCCS 2 Revision 1 specification."}
	ErrorMcaInvalidVcpVersion                              = &Error{0x00003B61, "ERROR_MCA_INVALID_VCP_VERSION", "The monitor's VCP version (0xDF) VCP code returned an invalid version value."}
	ErrorMcaMonitorViolatesMccsSpecification               = &Error{0x00003B62, "ERROR_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION", "The monitor does not comply with the MCCS specification it claims to support."}
	ErrorMcaMccsVersionMismatch                            = &Error{0x00003B63, "ERROR_MCA_MCCS_VERSION_MISMATCH", "The MCCS version in a monitor's mccs_ver capability does not match the MCCS version the monitor reports when the VCP version (0xDF) VCP code is used."}
	ErrorMcaUnsupportedMccsVersion                         = &Error{0x00003B64, "ERROR_MCA_UNSUPPORTED_MCCS_VERSION", "The monitor configuration API works only with monitors that support the MCCS 1.0, MCCS 2.0, or MCCS 2.0 Revision 1 specifications."}
	ErrorMcaInternalError                                  = &Error{0x00003B65, "ERROR_MCA_INTERNAL_ERROR", "An internal monitor configuration API error occurred."}
	ErrorMcaInvalidTechnologyTypeReturned                  = &Error{0x00003B66, "ERROR_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED", "The monitor returned an invalid monitor technology type. CRT, plasma, and LCD (TFT) are examples of monitor technology types. This error implies that the monitor violated the MCCS 2.0 or MCCS 2.0 Revision 1 specification."}
	ErrorMcaUnsupportedColorTemperature                    = &Error{0x00003B67, "ERROR_MCA_UNSUPPORTED_COLOR_TEMPERATURE", "The SetMonitorColorTemperature() caller passed a color temperature to it that the current monitor did not support. CRT, plasma, and LCD (TFT) are examples of monitor technology types. This error implies that the monitor violated the MCCS 2.0 or MCCS 2.0 Revision 1 specification."}
	ErrorAmbiguousSystemDevice                             = &Error{0x00003B92, "ERROR_AMBIGUOUS_SYSTEM_DEVICE", "The requested system device cannot be identified due to multiple indistinguishable devices potentially matching the identification criteria."}
	ErrorSystemDeviceNotFound                              = &Error{0x00003BC3, "ERROR_SYSTEM_DEVICE_NOT_FOUND", "The requested system device cannot be found."}
)

func FromCode(code uint32) error {

	if code == 0 {
		return nil
	}

	switch code {
	case 0x00000000:
		return ErrorSuccess
	case 0x00000001:
		return ErrorInvalidFunction
	case 0x00000002:
		return ErrorFileNotFound
	case 0x00000003:
		return ErrorPathNotFound
	case 0x00000004:
		return ErrorTooManyOpenFiles
	case 0x00000005:
		return ErrorAccessDenied
	case 0x00000006:
		return ErrorInvalidHandle
	case 0x00000007:
		return ErrorArenaTrashed
	case 0x00000008:
		return ErrorNotEnoughMemory
	case 0x00000009:
		return ErrorInvalidBlock
	case 0x0000000A:
		return ErrorBadEnvironment
	case 0x0000000B:
		return ErrorBadFormat
	case 0x0000000C:
		return ErrorInvalidAccess
	case 0x0000000D:
		return ErrorInvalidData
	case 0x0000000E:
		return ErrorOutofmemory
	case 0x0000000F:
		return ErrorInvalidDrive
	case 0x00000010:
		return ErrorCurrentDirectory
	case 0x00000011:
		return ErrorNotSameDevice
	case 0x00000012:
		return ErrorNoMoreFiles
	case 0x00000013:
		return ErrorWriteProtect
	case 0x00000014:
		return ErrorBadUnit
	case 0x00000015:
		return ErrorNotReady
	case 0x00000016:
		return ErrorBadCommand
	case 0x00000017:
		return ErrorCrc
	case 0x00000018:
		return ErrorBadLength
	case 0x00000019:
		return ErrorSeek
	case 0x0000001A:
		return ErrorNotDosDisk
	case 0x0000001B:
		return ErrorSectorNotFound
	case 0x0000001C:
		return ErrorOutOfPaper
	case 0x0000001D:
		return ErrorWriteFault
	case 0x0000001E:
		return ErrorReadFault
	case 0x0000001F:
		return ErrorGenFailure
	case 0x00000020:
		return ErrorSharingViolation
	case 0x00000021:
		return ErrorLockViolation
	case 0x00000022:
		return ErrorWrongDisk
	case 0x00000024:
		return ErrorSharingBufferExceeded
	case 0x00000026:
		return ErrorHandleEof
	case 0x00000027:
		return ErrorHandleDiskFull
	case 0x00000032:
		return ErrorNotSupported
	case 0x00000033:
		return ErrorRemNotList
	case 0x00000034:
		return ErrorDupName
	case 0x00000035:
		return ErrorBadNetpath
	case 0x00000036:
		return ErrorNetworkBusy
	case 0x00000037:
		return ErrorDevNotExist
	case 0x00000038:
		return ErrorTooManyCmds
	case 0x00000039:
		return ErrorAdapHdwErr
	case 0x0000003A:
		return ErrorBadNetResp
	case 0x0000003B:
		return ErrorUnexpNetErr
	case 0x0000003C:
		return ErrorBadRemAdap
	case 0x0000003D:
		return ErrorPrintqFull
	case 0x0000003E:
		return ErrorNoSpoolSpace
	case 0x0000003F:
		return ErrorPrintCancelled
	case 0x00000040:
		return ErrorNetnameDeleted
	case 0x00000041:
		return ErrorNetworkAccessDenied
	case 0x00000042:
		return ErrorBadDevType
	case 0x00000043:
		return ErrorBadNetName
	case 0x00000044:
		return ErrorTooManyNames
	case 0x00000045:
		return ErrorTooManySess
	case 0x00000046:
		return ErrorSharingPaused
	case 0x00000047:
		return ErrorReqNotAccep
	case 0x00000048:
		return ErrorRedirPaused
	case 0x00000050:
		return ErrorFileExists
	case 0x00000052:
		return ErrorCannotMake
	case 0x00000053:
		return ErrorFailI24
	case 0x00000054:
		return ErrorOutOfStructures
	case 0x00000055:
		return ErrorAlreadyAssigned
	case 0x00000056:
		return ErrorInvalidPassword
	case 0x00000057:
		return ErrorInvalidParameter
	case 0x00000058:
		return ErrorNetWriteFault
	case 0x00000059:
		return ErrorNoProcSlots
	case 0x00000064:
		return ErrorTooManySemaphores
	case 0x00000065:
		return ErrorExclSemAlreadyOwned
	case 0x00000066:
		return ErrorSemIsSet
	case 0x00000067:
		return ErrorTooManySemRequests
	case 0x00000068:
		return ErrorInvalidAtInterruptTime
	case 0x00000069:
		return ErrorSemOwnerDied
	case 0x0000006A:
		return ErrorSemUserLimit
	case 0x0000006B:
		return ErrorDiskChange
	case 0x0000006C:
		return ErrorDriveLocked
	case 0x0000006D:
		return ErrorBrokenPipe
	case 0x0000006E:
		return ErrorOpenFailed
	case 0x0000006F:
		return ErrorBufferOverflow
	case 0x00000070:
		return ErrorDiskFull
	case 0x00000071:
		return ErrorNoMoreSearchHandles
	case 0x00000072:
		return ErrorInvalidTargetHandle
	case 0x00000075:
		return ErrorInvalidCategory
	case 0x00000076:
		return ErrorInvalidVerifySwitch
	case 0x00000077:
		return ErrorBadDriverLevel
	case 0x00000078:
		return ErrorCallNotImplemented
	case 0x00000079:
		return ErrorSemTimeout
	case 0x0000007A:
		return ErrorInsufficientBuffer
	case 0x0000007B:
		return ErrorInvalidName
	case 0x0000007C:
		return ErrorInvalidLevel
	case 0x0000007D:
		return ErrorNoVolumeLabel
	case 0x0000007E:
		return ErrorModNotFound
	case 0x0000007F:
		return ErrorProcNotFound
	case 0x00000080:
		return ErrorWaitNoChildren
	case 0x00000081:
		return ErrorChildNotComplete
	case 0x00000082:
		return ErrorDirectAccessHandle
	case 0x00000083:
		return ErrorNegativeSeek
	case 0x00000084:
		return ErrorSeekOnDevice
	case 0x00000085:
		return ErrorIsJoinTarget
	case 0x00000086:
		return ErrorIsJoined
	case 0x00000087:
		return ErrorIsSubsted
	case 0x00000088:
		return ErrorNotJoined
	case 0x00000089:
		return ErrorNotSubsted
	case 0x0000008A:
		return ErrorJoinToJoin
	case 0x0000008B:
		return ErrorSubstToSubst
	case 0x0000008C:
		return ErrorJoinToSubst
	case 0x0000008D:
		return ErrorSubstToJoin
	case 0x0000008E:
		return ErrorBusyDrive
	case 0x0000008F:
		return ErrorSameDrive
	case 0x00000090:
		return ErrorDirNotRoot
	case 0x00000091:
		return ErrorDirNotEmpty
	case 0x00000092:
		return ErrorIsSubstPath
	case 0x00000093:
		return ErrorIsJoinPath
	case 0x00000094:
		return ErrorPathBusy
	case 0x00000095:
		return ErrorIsSubstTarget
	case 0x00000096:
		return ErrorSystemTrace
	case 0x00000097:
		return ErrorInvalidEventCount
	case 0x00000098:
		return ErrorTooManyMuxwaiters
	case 0x00000099:
		return ErrorInvalidListFormat
	case 0x0000009A:
		return ErrorLabelTooLong
	case 0x0000009B:
		return ErrorTooManyTcbs
	case 0x0000009C:
		return ErrorSignalRefused
	case 0x0000009D:
		return ErrorDiscarded
	case 0x0000009E:
		return ErrorNotLocked
	case 0x0000009F:
		return ErrorBadThreadidAddr
	case 0x000000A0:
		return ErrorBadArguments
	case 0x000000A1:
		return ErrorBadPathname
	case 0x000000A2:
		return ErrorSignalPending
	case 0x000000A4:
		return ErrorMaxThrdsReached
	case 0x000000A7:
		return ErrorLockFailed
	case 0x000000AA:
		return ErrorBusy
	case 0x000000AD:
		return ErrorCancelViolation
	case 0x000000AE:
		return ErrorAtomicLocksNotSupported
	case 0x000000B4:
		return ErrorInvalidSegmentNumber
	case 0x000000B6:
		return ErrorInvalidOrdinal
	case 0x000000B7:
		return ErrorAlreadyExists
	case 0x000000BA:
		return ErrorInvalidFlagNumber
	case 0x000000BB:
		return ErrorSemNotFound
	case 0x000000BC:
		return ErrorInvalidStartingCodeseg
	case 0x000000BD:
		return ErrorInvalidStackseg
	case 0x000000BE:
		return ErrorInvalidModuletype
	case 0x000000BF:
		return ErrorInvalidExeSignature
	case 0x000000C0:
		return ErrorExeMarkedInvalid
	case 0x000000C1:
		return ErrorBadExeFormat
	case 0x000000C2:
		return ErrorIteratedDataExceeds64k
	case 0x000000C3:
		return ErrorInvalidMinallocsize
	case 0x000000C4:
		return ErrorDynlinkFromInvalidRing
	case 0x000000C5:
		return ErrorIoplNotEnabled
	case 0x000000C6:
		return ErrorInvalidSegdpl
	case 0x000000C7:
		return ErrorAutodatasegExceeds64k
	case 0x000000C8:
		return ErrorRing2segMustBeMovable
	case 0x000000C9:
		return ErrorRelocChainXeedsSeglim
	case 0x000000CA:
		return ErrorInfloopInRelocChain
	case 0x000000CB:
		return ErrorEnvvarNotFound
	case 0x000000CD:
		return ErrorNoSignalSent
	case 0x000000CE:
		return ErrorFilenameExcedRange
	case 0x000000CF:
		return ErrorRing2StackInUse
	case 0x000000D0:
		return ErrorMetaExpansionTooLong
	case 0x000000D1:
		return ErrorInvalidSignalNumber
	case 0x000000D2:
		return ErrorThread1Inactive
	case 0x000000D4:
		return ErrorLocked
	case 0x000000D6:
		return ErrorTooManyModules
	case 0x000000D7:
		return ErrorNestingNotAllowed
	case 0x000000D8:
		return ErrorExeMachineTypeMismatch
	case 0x000000D9:
		return ErrorExeCannotModifySignedBinary
	case 0x000000DA:
		return ErrorExeCannotModifyStrongSignedBinary
	case 0x000000DC:
		return ErrorFileCheckedOut
	case 0x000000DD:
		return ErrorCheckoutRequired
	case 0x000000DE:
		return ErrorBadFileType
	case 0x000000DF:
		return ErrorFileTooLarge
	case 0x000000E0:
		return ErrorFormsAuthRequired
	case 0x000000E1:
		return ErrorVirusInfected
	case 0x000000E2:
		return ErrorVirusDeleted
	case 0x000000E5:
		return ErrorPipeLocal
	case 0x000000E6:
		return ErrorBadPipe
	case 0x000000E7:
		return ErrorPipeBusy
	case 0x000000E8:
		return ErrorNoData
	case 0x000000E9:
		return ErrorPipeNotConnected
	case 0x000000EA:
		return ErrorMoreData
	case 0x000000F0:
		return ErrorVcDisconnected
	case 0x000000FE:
		return ErrorInvalidEaName
	case 0x000000FF:
		return ErrorEaListInconsistent
	case 0x00000102:
		return WaitTimeout
	case 0x00000103:
		return ErrorNoMoreItems
	case 0x0000010A:
		return ErrorCannotCopy
	case 0x0000010B:
		return ErrorDirectory
	case 0x00000113:
		return ErrorEasDidntFit
	case 0x00000114:
		return ErrorEaFileCorrupt
	case 0x00000115:
		return ErrorEaTableFull
	case 0x00000116:
		return ErrorInvalidEaHandle
	case 0x0000011A:
		return ErrorEasNotSupported
	case 0x00000120:
		return ErrorNotOwner
	case 0x0000012A:
		return ErrorTooManyPosts
	case 0x0000012B:
		return ErrorPartialCopy
	case 0x0000012C:
		return ErrorOplockNotGranted
	case 0x0000012D:
		return ErrorInvalidOplockProtocol
	case 0x0000012E:
		return ErrorDiskTooFragmented
	case 0x0000012F:
		return ErrorDeletePending
	case 0x0000013D:
		return ErrorMrMidNotFound
	case 0x0000013E:
		return ErrorScopeNotFound
	case 0x0000015E:
		return ErrorFailNoactionReboot
	case 0x0000015F:
		return ErrorFailShutdown
	case 0x00000160:
		return ErrorFailRestart
	case 0x00000161:
		return ErrorMaxSessionsReached
	case 0x00000190:
		return ErrorThreadModeAlreadyBackground
	case 0x00000191:
		return ErrorThreadModeNotBackground
	case 0x00000192:
		return ErrorProcessModeAlreadyBackground
	case 0x00000193:
		return ErrorProcessModeNotBackground
	case 0x000001E7:
		return ErrorInvalidAddress
	case 0x000001F4:
		return ErrorUserProfileLoad
	case 0x00000216:
		return ErrorArithmeticOverflow
	case 0x00000217:
		return ErrorPipeConnected
	case 0x00000218:
		return ErrorPipeListening
	case 0x00000219:
		return ErrorVerifierStop
	case 0x0000021A:
		return ErrorAbiosError
	case 0x0000021B:
		return ErrorWx86Warning
	case 0x0000021C:
		return ErrorWx86Error
	case 0x0000021D:
		return ErrorTimerNotCanceled
	case 0x0000021E:
		return ErrorUnwind
	case 0x0000021F:
		return ErrorBadStack
	case 0x00000220:
		return ErrorInvalidUnwindTarget
	case 0x00000221:
		return ErrorInvalidPortAttributes
	case 0x00000222:
		return ErrorPortMessageTooLong
	case 0x00000223:
		return ErrorInvalidQuotaLower
	case 0x00000224:
		return ErrorDeviceAlreadyAttached
	case 0x00000225:
		return ErrorInstructionMisalignment
	case 0x00000226:
		return ErrorProfilingNotStarted
	case 0x00000227:
		return ErrorProfilingNotStopped
	case 0x00000228:
		return ErrorCouldNotInterpret
	case 0x00000229:
		return ErrorProfilingAtLimit
	case 0x0000022A:
		return ErrorCantWait
	case 0x0000022B:
		return ErrorCantTerminateSelf
	case 0x0000022C:
		return ErrorUnexpectedMmCreateErr
	case 0x0000022D:
		return ErrorUnexpectedMmMapError
	case 0x0000022E:
		return ErrorUnexpectedMmExtendErr
	case 0x0000022F:
		return ErrorBadFunctionTable
	case 0x00000230:
		return ErrorNoGuidTranslation
	case 0x00000231:
		return ErrorInvalidLdtSize
	case 0x00000233:
		return ErrorInvalidLdtOffset
	case 0x00000234:
		return ErrorInvalidLdtDescriptor
	case 0x00000235:
		return ErrorTooManyThreads
	case 0x00000236:
		return ErrorThreadNotInProcess
	case 0x00000237:
		return ErrorPagefileQuotaExceeded
	case 0x00000238:
		return ErrorLogonServerConflict
	case 0x00000239:
		return ErrorSynchronizationRequired
	case 0x0000023A:
		return ErrorNetOpenFailed
	case 0x0000023B:
		return ErrorIoPrivilegeFailed
	case 0x0000023C:
		return ErrorControlCExit
	case 0x0000023D:
		return ErrorMissingSystemfile
	case 0x0000023E:
		return ErrorUnhandledException
	case 0x0000023F:
		return ErrorAppInitFailure
	case 0x00000240:
		return ErrorPagefileCreateFailed
	case 0x00000241:
		return ErrorInvalidImageHash
	case 0x00000242:
		return ErrorNoPagefile
	case 0x00000243:
		return ErrorIllegalFloatContext
	case 0x00000244:
		return ErrorNoEventPair
	case 0x00000245:
		return ErrorDomainCtrlrConfigError
	case 0x00000246:
		return ErrorIllegalCharacter
	case 0x00000247:
		return ErrorUndefinedCharacter
	case 0x00000248:
		return ErrorFloppyVolume
	case 0x00000249:
		return ErrorBiosFailedToConnectInterrupt
	case 0x0000024A:
		return ErrorBackupController
	case 0x0000024B:
		return ErrorMutantLimitExceeded
	case 0x0000024C:
		return ErrorFsDriverRequired
	case 0x0000024D:
		return ErrorCannotLoadRegistryFile
	case 0x0000024E:
		return ErrorDebugAttachFailed
	case 0x0000024F:
		return ErrorSystemProcessTerminated
	case 0x00000250:
		return ErrorDataNotAccepted
	case 0x00000251:
		return ErrorVdmHardError
	case 0x00000252:
		return ErrorDriverCancelTimeout
	case 0x00000253:
		return ErrorReplyMessageMismatch
	case 0x00000254:
		return ErrorLostWritebehindData
	case 0x00000255:
		return ErrorClientServerParametersInvalid
	case 0x00000256:
		return ErrorNotTinyStream
	case 0x00000257:
		return ErrorStackOverflowRead
	case 0x00000258:
		return ErrorConvertToLarge
	case 0x00000259:
		return ErrorFoundOutOfScope
	case 0x0000025A:
		return ErrorAllocateBucket
	case 0x0000025B:
		return ErrorMarshallOverflow
	case 0x0000025C:
		return ErrorInvalidVariant
	case 0x0000025D:
		return ErrorBadCompressionBuffer
	case 0x0000025E:
		return ErrorAuditFailed
	case 0x0000025F:
		return ErrorTimerResolutionNotSet
	case 0x00000260:
		return ErrorInsufficientLogonInfo
	case 0x00000261:
		return ErrorBadDllEntrypoint
	case 0x00000262:
		return ErrorBadServiceEntrypoint
	case 0x00000263:
		return ErrorIpAddressConflict1
	case 0x00000264:
		return ErrorIpAddressConflict2
	case 0x00000265:
		return ErrorRegistryQuotaLimit
	case 0x00000266:
		return ErrorNoCallbackActive
	case 0x00000267:
		return ErrorPwdTooShort
	case 0x00000268:
		return ErrorPwdTooRecent
	case 0x00000269:
		return ErrorPwdHistoryConflict
	case 0x0000026A:
		return ErrorUnsupportedCompression
	case 0x0000026B:
		return ErrorInvalidHwProfile
	case 0x0000026C:
		return ErrorInvalidPlugplayDevicePath
	case 0x0000026D:
		return ErrorQuotaListInconsistent
	case 0x0000026E:
		return ErrorEvaluationExpiration
	case 0x0000026F:
		return ErrorIllegalDllRelocation
	case 0x00000270:
		return ErrorDllInitFailedLogoff
	case 0x00000271:
		return ErrorValidateContinue
	case 0x00000272:
		return ErrorNoMoreMatches
	case 0x00000273:
		return ErrorRangeListConflict
	case 0x00000274:
		return ErrorServerSidMismatch
	case 0x00000275:
		return ErrorCantEnableDenyOnly
	case 0x00000276:
		return ErrorFloatMultipleFaults
	case 0x00000277:
		return ErrorFloatMultipleTraps
	case 0x00000278:
		return ErrorNointerface
	case 0x00000279:
		return ErrorDriverFailedSleep
	case 0x0000027A:
		return ErrorCorruptSystemFile
	case 0x0000027B:
		return ErrorCommitmentMinimum
	case 0x0000027C:
		return ErrorPnpRestartEnumeration
	case 0x0000027D:
		return ErrorSystemImageBadSignature
	case 0x0000027E:
		return ErrorPnpRebootRequired
	case 0x0000027F:
		return ErrorInsufficientPower
	case 0x00000281:
		return ErrorSystemShutdown
	case 0x00000282:
		return ErrorPortNotSet
	case 0x00000283:
		return ErrorDsVersionCheckFailure
	case 0x00000284:
		return ErrorRangeNotFound
	case 0x00000286:
		return ErrorNotSafeModeDriver
	case 0x00000287:
		return ErrorFailedDriverEntry
	case 0x00000288:
		return ErrorDeviceEnumerationError
	case 0x00000289:
		return ErrorMountPointNotResolved
	case 0x0000028A:
		return ErrorInvalidDeviceObjectParameter
	case 0x0000028B:
		return ErrorMcaOccured
	case 0x0000028C:
		return ErrorDriverDatabaseError
	case 0x0000028D:
		return ErrorSystemHiveTooLarge
	case 0x0000028E:
		return ErrorDriverFailedPriorUnload
	case 0x0000028F:
		return ErrorVolsnapPrepareHibernate
	case 0x00000290:
		return ErrorHibernationFailure
	case 0x00000299:
		return ErrorFileSystemLimitation
	case 0x0000029C:
		return ErrorAssertionFailure
	case 0x0000029D:
		return ErrorAcpiError
	case 0x0000029E:
		return ErrorWowAssertion
	case 0x0000029F:
		return ErrorPnpBadMpsTable
	case 0x000002A0:
		return ErrorPnpTranslationFailed
	case 0x000002A1:
		return ErrorPnpIrqTranslationFailed
	case 0x000002A2:
		return ErrorPnpInvalidId
	case 0x000002A3:
		return ErrorWakeSystemDebugger
	case 0x000002A4:
		return ErrorHandlesClosed
	case 0x000002A5:
		return ErrorExtraneousInformation
	case 0x000002A6:
		return ErrorRxactCommitNecessary
	case 0x000002A7:
		return ErrorMediaCheck
	case 0x000002A8:
		return ErrorGuidSubstitutionMade
	case 0x000002A9:
		return ErrorStoppedOnSymlink
	case 0x000002AA:
		return ErrorLongjump
	case 0x000002AB:
		return ErrorPlugplayQueryVetoed
	case 0x000002AC:
		return ErrorUnwindConsolidate
	case 0x000002AD:
		return ErrorRegistryHiveRecovered
	case 0x000002AE:
		return ErrorDllMightBeInsecure
	case 0x000002AF:
		return ErrorDllMightBeIncompatible
	case 0x000002B0:
		return ErrorDbgExceptionNotHandled
	case 0x000002B1:
		return ErrorDbgReplyLater
	case 0x000002B2:
		return ErrorDbgUnableToProvideHandle
	case 0x000002B3:
		return ErrorDbgTerminateThread
	case 0x000002B4:
		return ErrorDbgTerminateProcess
	case 0x000002B5:
		return ErrorDbgControlC
	case 0x000002B6:
		return ErrorDbgPrintexceptionC
	case 0x000002B7:
		return ErrorDbgRipexception
	case 0x000002B8:
		return ErrorDbgControlBreak
	case 0x000002B9:
		return ErrorDbgCommandException
	case 0x000002BA:
		return ErrorObjectNameExists
	case 0x000002BB:
		return ErrorThreadWasSuspended
	case 0x000002BC:
		return ErrorImageNotAtBase
	case 0x000002BD:
		return ErrorRxactStateCreated
	case 0x000002BE:
		return ErrorSegmentNotification
	case 0x000002BF:
		return ErrorBadCurrentDirectory
	case 0x000002C0:
		return ErrorFtReadRecoveryFromBackup
	case 0x000002C1:
		return ErrorFtWriteRecovery
	case 0x000002C2:
		return ErrorImageMachineTypeMismatch
	case 0x000002C3:
		return ErrorReceivePartial
	case 0x000002C4:
		return ErrorReceiveExpedited
	case 0x000002C5:
		return ErrorReceivePartialExpedited
	case 0x000002C6:
		return ErrorEventDone
	case 0x000002C7:
		return ErrorEventPending
	case 0x000002C8:
		return ErrorCheckingFileSystem
	case 0x000002C9:
		return ErrorFatalAppExit
	case 0x000002CA:
		return ErrorPredefinedHandle
	case 0x000002CB:
		return ErrorWasUnlocked
	case 0x000002CD:
		return ErrorWasLocked
	case 0x000002CF:
		return ErrorAlreadyWin32
	case 0x000002D0:
		return ErrorImageMachineTypeMismatchExe
	case 0x000002D1:
		return ErrorNoYieldPerformed
	case 0x000002D2:
		return ErrorTimerResumeIgnored
	case 0x000002D3:
		return ErrorArbitrationUnhandled
	case 0x000002D4:
		return ErrorCardbusNotSupported
	case 0x000002D5:
		return ErrorMpProcessorMismatch
	case 0x000002D6:
		return ErrorHibernated
	case 0x000002D7:
		return ErrorResumeHibernation
	case 0x000002D8:
		return ErrorFirmwareUpdated
	case 0x000002D9:
		return ErrorDriversLeakingLockedPages
	case 0x000002DA:
		return ErrorWakeSystem
	case 0x000002DF:
		return ErrorAbandonedWait0
	case 0x000002E4:
		return ErrorElevationRequired
	case 0x000002E5:
		return ErrorReparse
	case 0x000002E6:
		return ErrorOplockBreakInProgress
	case 0x000002E7:
		return ErrorVolumeMounted
	case 0x000002E8:
		return ErrorRxactCommitted
	case 0x000002E9:
		return ErrorNotifyCleanup
	case 0x000002EA:
		return ErrorPrimaryTransportConnectFailed
	case 0x000002EB:
		return ErrorPageFaultTransition
	case 0x000002EC:
		return ErrorPageFaultDemandZero
	case 0x000002ED:
		return ErrorPageFaultCopyOnWrite
	case 0x000002EE:
		return ErrorPageFaultGuardPage
	case 0x000002EF:
		return ErrorPageFaultPagingFile
	case 0x000002F0:
		return ErrorCachePageLocked
	case 0x000002F1:
		return ErrorCrashDump
	case 0x000002F2:
		return ErrorBufferAllZeros
	case 0x000002F3:
		return ErrorReparseObject
	case 0x000002F4:
		return ErrorResourceRequirementsChanged
	case 0x000002F5:
		return ErrorTranslationComplete
	case 0x000002F6:
		return ErrorNothingToTerminate
	case 0x000002F7:
		return ErrorProcessNotInJob
	case 0x000002F8:
		return ErrorProcessInJob
	case 0x000002F9:
		return ErrorVolsnapHibernateReady
	case 0x000002FA:
		return ErrorFsfilterOpCompletedSuccessfully
	case 0x000002FB:
		return ErrorInterruptVectorAlreadyConnected
	case 0x000002FC:
		return ErrorInterruptStillConnected
	case 0x000002FD:
		return ErrorWaitForOplock
	case 0x000002FE:
		return ErrorDbgExceptionHandled
	case 0x000002FF:
		return ErrorDbgContinue
	case 0x00000300:
		return ErrorCallbackPopStack
	case 0x00000301:
		return ErrorCompressionDisabled
	case 0x00000302:
		return ErrorCantfetchbackwards
	case 0x00000303:
		return ErrorCantscrollbackwards
	case 0x00000304:
		return ErrorRowsnotreleased
	case 0x00000305:
		return ErrorBadAccessorFlags
	case 0x00000306:
		return ErrorErrorsEncountered
	case 0x00000307:
		return ErrorNotCapable
	case 0x00000308:
		return ErrorRequestOutOfSequence
	case 0x00000309:
		return ErrorVersionParseError
	case 0x0000030A:
		return ErrorBadstartposition
	case 0x0000030B:
		return ErrorMemoryHardware
	case 0x0000030C:
		return ErrorDiskRepairDisabled
	case 0x0000030D:
		return ErrorInsufficientResourceForSpecifiedSharedSectionSize
	case 0x0000030E:
		return ErrorSystemPowerstateTransition
	case 0x0000030F:
		return ErrorSystemPowerstateComplexTransition
	case 0x00000310:
		return ErrorMcaException
	case 0x00000311:
		return ErrorAccessAuditByPolicy
	case 0x00000312:
		return ErrorAccessDisabledNoSaferUiByPolicy
	case 0x00000313:
		return ErrorAbandonHiberfile
	case 0x00000314:
		return ErrorLostWritebehindDataNetworkDisconnected
	case 0x00000315:
		return ErrorLostWritebehindDataNetworkServerError
	case 0x00000316:
		return ErrorLostWritebehindDataLocalDiskError
	case 0x000003E2:
		return ErrorEaAccessDenied
	case 0x000003E3:
		return ErrorOperationAborted
	case 0x000003E4:
		return ErrorIoIncomplete
	case 0x000003E5:
		return ErrorIoPending
	case 0x000003E6:
		return ErrorNoaccess
	case 0x000003E7:
		return ErrorSwaperror
	case 0x000003E9:
		return ErrorStackOverflow
	case 0x000003EA:
		return ErrorInvalidMessage
	case 0x000003EB:
		return ErrorCanNotComplete
	case 0x000003EC:
		return ErrorInvalidFlags
	case 0x000003ED:
		return ErrorUnrecognizedVolume
	case 0x000003EE:
		return ErrorFileInvalid
	case 0x000003EF:
		return ErrorFullscreenMode
	case 0x000003F0:
		return ErrorNoToken
	case 0x000003F1:
		return ErrorBaddb
	case 0x000003F2:
		return ErrorBadkey
	case 0x000003F3:
		return ErrorCantopen
	case 0x000003F4:
		return ErrorCantread
	case 0x000003F5:
		return ErrorCantwrite
	case 0x000003F6:
		return ErrorRegistryRecovered
	case 0x000003F7:
		return ErrorRegistryCorrupt
	case 0x000003F8:
		return ErrorRegistryIoFailed
	case 0x000003F9:
		return ErrorNotRegistryFile
	case 0x000003FA:
		return ErrorKeyDeleted
	case 0x000003FB:
		return ErrorNoLogSpace
	case 0x000003FC:
		return ErrorKeyHasChildren
	case 0x000003FD:
		return ErrorChildMustBeVolatile
	case 0x000003FE:
		return ErrorNotifyEnumDir
	case 0x0000041B:
		return ErrorDependentServicesRunning
	case 0x0000041C:
		return ErrorInvalidServiceControl
	case 0x0000041D:
		return ErrorServiceRequestTimeout
	case 0x0000041E:
		return ErrorServiceNoThread
	case 0x0000041F:
		return ErrorServiceDatabaseLocked
	case 0x00000420:
		return ErrorServiceAlreadyRunning
	case 0x00000421:
		return ErrorInvalidServiceAccount
	case 0x00000422:
		return ErrorServiceDisabled
	case 0x00000423:
		return ErrorCircularDependency
	case 0x00000424:
		return ErrorServiceDoesNotExist
	case 0x00000425:
		return ErrorServiceCannotAcceptCtrl
	case 0x00000426:
		return ErrorServiceNotActive
	case 0x00000427:
		return ErrorFailedServiceControllerConnect
	case 0x00000428:
		return ErrorExceptionInService
	case 0x00000429:
		return ErrorDatabaseDoesNotExist
	case 0x0000042A:
		return ErrorServiceSpecificError
	case 0x0000042B:
		return ErrorProcessAborted
	case 0x0000042C:
		return ErrorServiceDependencyFail
	case 0x0000042D:
		return ErrorServiceLogonFailed
	case 0x0000042E:
		return ErrorServiceStartHang
	case 0x0000042F:
		return ErrorInvalidServiceLock
	case 0x00000430:
		return ErrorServiceMarkedForDelete
	case 0x00000431:
		return ErrorServiceExists
	case 0x00000432:
		return ErrorAlreadyRunningLkg
	case 0x00000433:
		return ErrorServiceDependencyDeleted
	case 0x00000434:
		return ErrorBootAlreadyAccepted
	case 0x00000435:
		return ErrorServiceNeverStarted
	case 0x00000436:
		return ErrorDuplicateServiceName
	case 0x00000437:
		return ErrorDifferentServiceAccount
	case 0x00000438:
		return ErrorCannotDetectDriverFailure
	case 0x00000439:
		return ErrorCannotDetectProcessAbort
	case 0x0000043A:
		return ErrorNoRecoveryProgram
	case 0x0000043B:
		return ErrorServiceNotInExe
	case 0x0000043C:
		return ErrorNotSafebootService
	case 0x0000044C:
		return ErrorEndOfMedia
	case 0x0000044D:
		return ErrorFilemarkDetected
	case 0x0000044E:
		return ErrorBeginningOfMedia
	case 0x0000044F:
		return ErrorSetmarkDetected
	case 0x00000450:
		return ErrorNoDataDetected
	case 0x00000451:
		return ErrorPartitionFailure
	case 0x00000452:
		return ErrorInvalidBlockLength
	case 0x00000453:
		return ErrorDeviceNotPartitioned
	case 0x00000454:
		return ErrorUnableToLockMedia
	case 0x00000455:
		return ErrorUnableToUnloadMedia
	case 0x00000456:
		return ErrorMediaChanged
	case 0x00000457:
		return ErrorBusReset
	case 0x00000458:
		return ErrorNoMediaInDrive
	case 0x00000459:
		return ErrorNoUnicodeTranslation
	case 0x0000045A:
		return ErrorDllInitFailed
	case 0x0000045B:
		return ErrorShutdownInProgress
	case 0x0000045C:
		return ErrorNoShutdownInProgress
	case 0x0000045D:
		return ErrorIoDevice
	case 0x0000045E:
		return ErrorSerialNoDevice
	case 0x0000045F:
		return ErrorIrqBusy
	case 0x00000460:
		return ErrorMoreWrites
	case 0x00000461:
		return ErrorCounterTimeout
	case 0x00000462:
		return ErrorFloppyIdMarkNotFound
	case 0x00000463:
		return ErrorFloppyWrongCylinder
	case 0x00000464:
		return ErrorFloppyUnknownError
	case 0x00000465:
		return ErrorFloppyBadRegisters
	case 0x00000466:
		return ErrorDiskRecalibrateFailed
	case 0x00000467:
		return ErrorDiskOperationFailed
	case 0x00000468:
		return ErrorDiskResetFailed
	case 0x00000469:
		return ErrorEomOverflow
	case 0x0000046A:
		return ErrorNotEnoughServerMemory
	case 0x0000046B:
		return ErrorPossibleDeadlock
	case 0x0000046C:
		return ErrorMappedAlignment
	case 0x00000474:
		return ErrorSetPowerStateVetoed
	case 0x00000475:
		return ErrorSetPowerStateFailed
	case 0x00000476:
		return ErrorTooManyLinks
	case 0x0000047E:
		return ErrorOldWinVersion
	case 0x0000047F:
		return ErrorAppWrongOs
	case 0x00000480:
		return ErrorSingleInstanceApp
	case 0x00000481:
		return ErrorRmodeApp
	case 0x00000482:
		return ErrorInvalidDll
	case 0x00000483:
		return ErrorNoAssociation
	case 0x00000484:
		return ErrorDdeFail
	case 0x00000485:
		return ErrorDllNotFound
	case 0x00000486:
		return ErrorNoMoreUserHandles
	case 0x00000487:
		return ErrorMessageSyncOnly
	case 0x00000488:
		return ErrorSourceElementEmpty
	case 0x00000489:
		return ErrorDestinationElementFull
	case 0x0000048A:
		return ErrorIllegalElementAddress
	case 0x0000048B:
		return ErrorMagazineNotPresent
	case 0x0000048C:
		return ErrorDeviceReinitializationNeeded
	case 0x0000048D:
		return ErrorDeviceRequiresCleaning
	case 0x0000048E:
		return ErrorDeviceDoorOpen
	case 0x0000048F:
		return ErrorDeviceNotConnected
	case 0x00000490:
		return ErrorNotFound
	case 0x00000491:
		return ErrorNoMatch
	case 0x00000492:
		return ErrorSetNotFound
	case 0x00000493:
		return ErrorPointNotFound
	case 0x00000494:
		return ErrorNoTrackingService
	case 0x00000495:
		return ErrorNoVolumeId
	case 0x00000497:
		return ErrorUnableToRemoveReplaced
	case 0x00000498:
		return ErrorUnableToMoveReplacement
	case 0x00000499:
		return ErrorUnableToMoveReplacement2
	case 0x0000049A:
		return ErrorJournalDeleteInProgress
	case 0x0000049B:
		return ErrorJournalNotActive
	case 0x0000049C:
		return ErrorPotentialFileFound
	case 0x0000049D:
		return ErrorJournalEntryDeleted
	case 0x000004A6:
		return ErrorShutdownIsScheduled
	case 0x000004A7:
		return ErrorShutdownUsersLoggedOn
	case 0x000004B0:
		return ErrorBadDevice
	case 0x000004B1:
		return ErrorConnectionUnavail
	case 0x000004B2:
		return ErrorDeviceAlreadyRemembered
	case 0x000004B3:
		return ErrorNoNetOrBadPath
	case 0x000004B4:
		return ErrorBadProvider
	case 0x000004B5:
		return ErrorCannotOpenProfile
	case 0x000004B6:
		return ErrorBadProfile
	case 0x000004B7:
		return ErrorNotContainer
	case 0x000004B8:
		return ErrorExtendedError
	case 0x000004B9:
		return ErrorInvalidGroupname
	case 0x000004BA:
		return ErrorInvalidComputername
	case 0x000004BB:
		return ErrorInvalidEventname
	case 0x000004BC:
		return ErrorInvalidDomainname
	case 0x000004BD:
		return ErrorInvalidServicename
	case 0x000004BE:
		return ErrorInvalidNetname
	case 0x000004BF:
		return ErrorInvalidSharename
	case 0x000004C0:
		return ErrorInvalidPasswordname
	case 0x000004C1:
		return ErrorInvalidMessagename
	case 0x000004C2:
		return ErrorInvalidMessagedest
	case 0x000004C3:
		return ErrorSessionCredentialConflict
	case 0x000004C4:
		return ErrorRemoteSessionLimitExceeded
	case 0x000004C5:
		return ErrorDupDomainname
	case 0x000004C6:
		return ErrorNoNetwork
	case 0x000004C7:
		return ErrorCancelled
	case 0x000004C8:
		return ErrorUserMappedFile
	case 0x000004C9:
		return ErrorConnectionRefused
	case 0x000004CA:
		return ErrorGracefulDisconnect
	case 0x000004CB:
		return ErrorAddressAlreadyAssociated
	case 0x000004CC:
		return ErrorAddressNotAssociated
	case 0x000004CD:
		return ErrorConnectionInvalid
	case 0x000004CE:
		return ErrorConnectionActive
	case 0x000004CF:
		return ErrorNetworkUnreachable
	case 0x000004D0:
		return ErrorHostUnreachable
	case 0x000004D1:
		return ErrorProtocolUnreachable
	case 0x000004D2:
		return ErrorPortUnreachable
	case 0x000004D3:
		return ErrorRequestAborted
	case 0x000004D4:
		return ErrorConnectionAborted
	case 0x000004D5:
		return ErrorRetry
	case 0x000004D6:
		return ErrorConnectionCountLimit
	case 0x000004D7:
		return ErrorLoginTimeRestriction
	case 0x000004D8:
		return ErrorLoginWkstaRestriction
	case 0x000004D9:
		return ErrorIncorrectAddress
	case 0x000004DA:
		return ErrorAlreadyRegistered
	case 0x000004DB:
		return ErrorServiceNotFound
	case 0x000004DC:
		return ErrorNotAuthenticated
	case 0x000004DD:
		return ErrorNotLoggedOn
	case 0x000004DE:
		return ErrorContinue
	case 0x000004DF:
		return ErrorAlreadyInitialized
	case 0x000004E0:
		return ErrorNoMoreDevices
	case 0x000004E1:
		return ErrorNoSuchSite
	case 0x000004E2:
		return ErrorDomainControllerExists
	case 0x000004E3:
		return ErrorOnlyIfConnected
	case 0x000004E4:
		return ErrorOverrideNochanges
	case 0x000004E5:
		return ErrorBadUserProfile
	case 0x000004E6:
		return ErrorNotSupportedOnSbs
	case 0x000004E7:
		return ErrorServerShutdownInProgress
	case 0x000004E8:
		return ErrorHostDown
	case 0x000004E9:
		return ErrorNonAccountSid
	case 0x000004EA:
		return ErrorNonDomainSid
	case 0x000004EB:
		return ErrorApphelpBlock
	case 0x000004EC:
		return ErrorAccessDisabledByPolicy
	case 0x000004ED:
		return ErrorRegNatConsumption
	case 0x000004EE:
		return ErrorCscshareOffline
	case 0x000004EF:
		return ErrorPkinitFailure
	case 0x000004F0:
		return ErrorSmartcardSubsystemFailure
	case 0x000004F1:
		return ErrorDowngradeDetected
	case 0x000004F7:
		return ErrorMachineLocked
	case 0x000004F9:
		return ErrorCallbackSuppliedInvalidData
	case 0x000004FA:
		return ErrorSyncForegroundRefreshRequired
	case 0x000004FB:
		return ErrorDriverBlocked
	case 0x000004FC:
		return ErrorInvalidImportOfNonDll
	case 0x000004FD:
		return ErrorAccessDisabledWebblade
	case 0x000004FE:
		return ErrorAccessDisabledWebbladeTamper
	case 0x000004FF:
		return ErrorRecoveryFailure
	case 0x00000500:
		return ErrorAlreadyFiber
	case 0x00000501:
		return ErrorAlreadyThread
	case 0x00000502:
		return ErrorStackBufferOverrun
	case 0x00000503:
		return ErrorParameterQuotaExceeded
	case 0x00000504:
		return ErrorDebuggerInactive
	case 0x00000505:
		return ErrorDelayLoadFailed
	case 0x00000506:
		return ErrorVdmDisallowed
	case 0x00000507:
		return ErrorUnidentifiedError
	case 0x00000508:
		return ErrorInvalidCruntimeParameter
	case 0x00000509:
		return ErrorBeyondVdl
	case 0x0000050A:
		return ErrorIncompatibleServiceSidType
	case 0x0000050B:
		return ErrorDriverProcessTerminated
	case 0x0000050C:
		return ErrorImplementationLimit
	case 0x0000050D:
		return ErrorProcessIsProtected
	case 0x0000050E:
		return ErrorServiceNotifyClientLagging
	case 0x0000050F:
		return ErrorDiskQuotaExceeded
	case 0x00000510:
		return ErrorContentBlocked
	case 0x00000511:
		return ErrorIncompatibleServicePrivilege
	case 0x00000513:
		return ErrorInvalidLabel
	case 0x00000514:
		return ErrorNotAllAssigned
	case 0x00000515:
		return ErrorSomeNotMapped
	case 0x00000516:
		return ErrorNoQuotasForAccount
	case 0x00000517:
		return ErrorLocalUserSessionKey
	case 0x00000518:
		return ErrorNullLmPassword
	case 0x00000519:
		return ErrorUnknownRevision
	case 0x0000051A:
		return ErrorRevisionMismatch
	case 0x0000051B:
		return ErrorInvalidOwner
	case 0x0000051C:
		return ErrorInvalidPrimaryGroup
	case 0x0000051D:
		return ErrorNoImpersonationToken
	case 0x0000051E:
		return ErrorCantDisableMandatory
	case 0x0000051F:
		return ErrorNoLogonServers
	case 0x00000520:
		return ErrorNoSuchLogonSession
	case 0x00000521:
		return ErrorNoSuchPrivilege
	case 0x00000522:
		return ErrorPrivilegeNotHeld
	case 0x00000523:
		return ErrorInvalidAccountName
	case 0x00000524:
		return ErrorUserExists
	case 0x00000525:
		return ErrorNoSuchUser
	case 0x00000526:
		return ErrorGroupExists
	case 0x00000527:
		return ErrorNoSuchGroup
	case 0x00000528:
		return ErrorMemberInGroup
	case 0x00000529:
		return ErrorMemberNotInGroup
	case 0x0000052A:
		return ErrorLastAdmin
	case 0x0000052B:
		return ErrorWrongPassword
	case 0x0000052C:
		return ErrorIllFormedPassword
	case 0x0000052D:
		return ErrorPasswordRestriction
	case 0x0000052E:
		return ErrorLogonFailure
	case 0x0000052F:
		return ErrorAccountRestriction
	case 0x00000530:
		return ErrorInvalidLogonHours
	case 0x00000531:
		return ErrorInvalidWorkstation
	case 0x00000532:
		return ErrorPasswordExpired
	case 0x00000533:
		return ErrorAccountDisabled
	case 0x00000534:
		return ErrorNoneMapped
	case 0x00000535:
		return ErrorTooManyLuidsRequested
	case 0x00000536:
		return ErrorLuidsExhausted
	case 0x00000537:
		return ErrorInvalidSubAuthority
	case 0x00000538:
		return ErrorInvalidAcl
	case 0x00000539:
		return ErrorInvalidSid
	case 0x0000053A:
		return ErrorInvalidSecurityDescr
	case 0x0000053C:
		return ErrorBadInheritanceAcl
	case 0x0000053D:
		return ErrorServerDisabled
	case 0x0000053E:
		return ErrorServerNotDisabled
	case 0x0000053F:
		return ErrorInvalidIdAuthority
	case 0x00000540:
		return ErrorAllottedSpaceExceeded
	case 0x00000541:
		return ErrorInvalidGroupAttributes
	case 0x00000542:
		return ErrorBadImpersonationLevel
	case 0x00000543:
		return ErrorCantOpenAnonymous
	case 0x00000544:
		return ErrorBadValidationClass
	case 0x00000545:
		return ErrorBadTokenType
	case 0x00000546:
		return ErrorNoSecurityOnObject
	case 0x00000547:
		return ErrorCantAccessDomainInfo
	case 0x00000548:
		return ErrorInvalidServerState
	case 0x00000549:
		return ErrorInvalidDomainState
	case 0x0000054A:
		return ErrorInvalidDomainRole
	case 0x0000054B:
		return ErrorNoSuchDomain
	case 0x0000054C:
		return ErrorDomainExists
	case 0x0000054D:
		return ErrorDomainLimitExceeded
	case 0x0000054E:
		return ErrorInternalDbCorruption
	case 0x0000054F:
		return ErrorInternalError
	case 0x00000550:
		return ErrorGenericNotMapped
	case 0x00000551:
		return ErrorBadDescriptorFormat
	case 0x00000552:
		return ErrorNotLogonProcess
	case 0x00000553:
		return ErrorLogonSessionExists
	case 0x00000554:
		return ErrorNoSuchPackage
	case 0x00000555:
		return ErrorBadLogonSessionState
	case 0x00000556:
		return ErrorLogonSessionCollision
	case 0x00000557:
		return ErrorInvalidLogonType
	case 0x00000558:
		return ErrorCannotImpersonate
	case 0x00000559:
		return ErrorRxactInvalidState
	case 0x0000055A:
		return ErrorRxactCommitFailure
	case 0x0000055B:
		return ErrorSpecialAccount
	case 0x0000055C:
		return ErrorSpecialGroup
	case 0x0000055D:
		return ErrorSpecialUser
	case 0x0000055E:
		return ErrorMembersPrimaryGroup
	case 0x0000055F:
		return ErrorTokenAlreadyInUse
	case 0x00000560:
		return ErrorNoSuchAlias
	case 0x00000561:
		return ErrorMemberNotInAlias
	case 0x00000562:
		return ErrorMemberInAlias
	case 0x00000563:
		return ErrorAliasExists
	case 0x00000564:
		return ErrorLogonNotGranted
	case 0x00000565:
		return ErrorTooManySecrets
	case 0x00000566:
		return ErrorSecretTooLong
	case 0x00000567:
		return ErrorInternalDbError
	case 0x00000568:
		return ErrorTooManyContextIds
	case 0x00000569:
		return ErrorLogonTypeNotGranted
	case 0x0000056A:
		return ErrorNtCrossEncryptionRequired
	case 0x0000056B:
		return ErrorNoSuchMember
	case 0x0000056C:
		return ErrorInvalidMember
	case 0x0000056D:
		return ErrorTooManySids
	case 0x0000056E:
		return ErrorLmCrossEncryptionRequired
	case 0x0000056F:
		return ErrorNoInheritance
	case 0x00000570:
		return ErrorFileCorrupt
	case 0x00000571:
		return ErrorDiskCorrupt
	case 0x00000572:
		return ErrorNoUserSessionKey
	case 0x00000573:
		return ErrorLicenseQuotaExceeded
	case 0x00000574:
		return ErrorWrongTargetName
	case 0x00000575:
		return ErrorMutualAuthFailed
	case 0x00000576:
		return ErrorTimeSkew
	case 0x00000577:
		return ErrorCurrentDomainNotAllowed
	case 0x00000578:
		return ErrorInvalidWindowHandle
	case 0x00000579:
		return ErrorInvalidMenuHandle
	case 0x0000057A:
		return ErrorInvalidCursorHandle
	case 0x0000057B:
		return ErrorInvalidAccelHandle
	case 0x0000057C:
		return ErrorInvalidHookHandle
	case 0x0000057D:
		return ErrorInvalidDwpHandle
	case 0x0000057E:
		return ErrorTlwWithWschild
	case 0x0000057F:
		return ErrorCannotFindWndClass
	case 0x00000580:
		return ErrorWindowOfOtherThread
	case 0x00000581:
		return ErrorHotkeyAlreadyRegistered
	case 0x00000582:
		return ErrorClassAlreadyExists
	case 0x00000583:
		return ErrorClassDoesNotExist
	case 0x00000584:
		return ErrorClassHasWindows
	case 0x00000585:
		return ErrorInvalidIndex
	case 0x00000586:
		return ErrorInvalidIconHandle
	case 0x00000587:
		return ErrorPrivateDialogIndex
	case 0x00000588:
		return ErrorListboxIdNotFound
	case 0x00000589:
		return ErrorNoWildcardCharacters
	case 0x0000058A:
		return ErrorClipboardNotOpen
	case 0x0000058B:
		return ErrorHotkeyNotRegistered
	case 0x0000058C:
		return ErrorWindowNotDialog
	case 0x0000058D:
		return ErrorControlIdNotFound
	case 0x0000058E:
		return ErrorInvalidComboboxMessage
	case 0x0000058F:
		return ErrorWindowNotCombobox
	case 0x00000590:
		return ErrorInvalidEditHeight
	case 0x00000591:
		return ErrorDcNotFound
	case 0x00000592:
		return ErrorInvalidHookFilter
	case 0x00000593:
		return ErrorInvalidFilterProc
	case 0x00000594:
		return ErrorHookNeedsHmod
	case 0x00000595:
		return ErrorGlobalOnlyHook
	case 0x00000596:
		return ErrorJournalHookSet
	case 0x00000597:
		return ErrorHookNotInstalled
	case 0x00000598:
		return ErrorInvalidLbMessage
	case 0x00000599:
		return ErrorSetcountOnBadLb
	case 0x0000059A:
		return ErrorLbWithoutTabstops
	case 0x0000059B:
		return ErrorDestroyObjectOfOtherThread
	case 0x0000059C:
		return ErrorChildWindowMenu
	case 0x0000059D:
		return ErrorNoSystemMenu
	case 0x0000059E:
		return ErrorInvalidMsgboxStyle
	case 0x0000059F:
		return ErrorInvalidSpiValue
	case 0x000005A0:
		return ErrorScreenAlreadyLocked
	case 0x000005A1:
		return ErrorHwndsHaveDiffParent
	case 0x000005A2:
		return ErrorNotChildWindow
	case 0x000005A3:
		return ErrorInvalidGwCommand
	case 0x000005A4:
		return ErrorInvalidThreadId
	case 0x000005A5:
		return ErrorNonMdichildWindow
	case 0x000005A6:
		return ErrorPopupAlreadyActive
	case 0x000005A7:
		return ErrorNoScrollbars
	case 0x000005A8:
		return ErrorInvalidScrollbarRange
	case 0x000005A9:
		return ErrorInvalidShowwinCommand
	case 0x000005AA:
		return ErrorNoSystemResources
	case 0x000005AB:
		return ErrorNonpagedSystemResources
	case 0x000005AC:
		return ErrorPagedSystemResources
	case 0x000005AD:
		return ErrorWorkingSetQuota
	case 0x000005AE:
		return ErrorPagefileQuota
	case 0x000005AF:
		return ErrorCommitmentLimit
	case 0x000005B0:
		return ErrorMenuItemNotFound
	case 0x000005B1:
		return ErrorInvalidKeyboardHandle
	case 0x000005B2:
		return ErrorHookTypeNotAllowed
	case 0x000005B3:
		return ErrorRequiresInteractiveWindowstation
	case 0x000005B4:
		return ErrorTimeout
	case 0x000005B5:
		return ErrorInvalidMonitorHandle
	case 0x000005B6:
		return ErrorIncorrectSize
	case 0x000005B7:
		return ErrorSymlinkClassDisabled
	case 0x000005B8:
		return ErrorSymlinkNotSupported
	case 0x000005DC:
		return ErrorEventlogFileCorrupt
	case 0x000005DD:
		return ErrorEventlogCantStart
	case 0x000005DE:
		return ErrorLogFileFull
	case 0x000005DF:
		return ErrorEventlogFileChanged
	case 0x0000060E:
		return ErrorInvalidTaskName
	case 0x0000060F:
		return ErrorInvalidTaskIndex
	case 0x00000610:
		return ErrorThreadAlreadyInTask
	case 0x00000641:
		return ErrorInstallServiceFailure
	case 0x00000642:
		return ErrorInstallUserexit
	case 0x00000643:
		return ErrorInstallFailure
	case 0x00000644:
		return ErrorInstallSuspend
	case 0x00000645:
		return ErrorUnknownProduct
	case 0x00000646:
		return ErrorUnknownFeature
	case 0x00000647:
		return ErrorUnknownComponent
	case 0x00000648:
		return ErrorUnknownProperty
	case 0x00000649:
		return ErrorInvalidHandleState
	case 0x0000064A:
		return ErrorBadConfiguration
	case 0x0000064B:
		return ErrorIndexAbsent
	case 0x0000064C:
		return ErrorInstallSourceAbsent
	case 0x0000064D:
		return ErrorInstallPackageVersion
	case 0x0000064E:
		return ErrorProductUninstalled
	case 0x0000064F:
		return ErrorBadQuerySyntax
	case 0x00000650:
		return ErrorInvalidField
	case 0x00000651:
		return ErrorDeviceRemoved
	case 0x00000652:
		return ErrorInstallAlreadyRunning
	case 0x00000653:
		return ErrorInstallPackageOpenFailed
	case 0x00000654:
		return ErrorInstallPackageInvalid
	case 0x00000655:
		return ErrorInstallUiFailure
	case 0x00000656:
		return ErrorInstallLogFailure
	case 0x00000657:
		return ErrorInstallLanguageUnsupported
	case 0x00000658:
		return ErrorInstallTransformFailure
	case 0x00000659:
		return ErrorInstallPackageRejected
	case 0x0000065A:
		return ErrorFunctionNotCalled
	case 0x0000065B:
		return ErrorFunctionFailed
	case 0x0000065C:
		return ErrorInvalidTable
	case 0x0000065D:
		return ErrorDatatypeMismatch
	case 0x0000065E:
		return ErrorUnsupportedType
	case 0x0000065F:
		return ErrorCreateFailed
	case 0x00000660:
		return ErrorInstallTempUnwritable
	case 0x00000661:
		return ErrorInstallPlatformUnsupported
	case 0x00000662:
		return ErrorInstallNotused
	case 0x00000663:
		return ErrorPatchPackageOpenFailed
	case 0x00000664:
		return ErrorPatchPackageInvalid
	case 0x00000665:
		return ErrorPatchPackageUnsupported
	case 0x00000666:
		return ErrorProductVersion
	case 0x00000667:
		return ErrorInvalidCommandLine
	case 0x00000668:
		return ErrorInstallRemoteDisallowed
	case 0x00000669:
		return ErrorSuccessRebootInitiated
	case 0x0000066A:
		return ErrorPatchTargetNotFound
	case 0x0000066B:
		return ErrorPatchPackageRejected
	case 0x0000066C:
		return ErrorInstallTransformRejected
	case 0x0000066D:
		return ErrorInstallRemoteProhibited
	case 0x0000066E:
		return ErrorPatchRemovalUnsupported
	case 0x0000066F:
		return ErrorUnknownPatch
	case 0x00000670:
		return ErrorPatchNoSequence
	case 0x00000671:
		return ErrorPatchRemovalDisallowed
	case 0x00000672:
		return ErrorInvalidPatchXml
	case 0x00000673:
		return ErrorPatchManagedAdvertisedProduct
	case 0x00000674:
		return ErrorInstallServiceSafeboot
	case 0x000006A4:
		return RpcSInvalidStringBinding
	case 0x000006A5:
		return RpcSWrongKindOfBinding
	case 0x000006A6:
		return RpcSInvalidBinding
	case 0x000006A7:
		return RpcSProtseqNotSupported
	case 0x000006A8:
		return RpcSInvalidRpcProtseq
	case 0x000006A9:
		return RpcSInvalidStringUuid
	case 0x000006AA:
		return RpcSInvalidEndpointFormat
	case 0x000006AB:
		return RpcSInvalidNetAddr
	case 0x000006AC:
		return RpcSNoEndpointFound
	case 0x000006AD:
		return RpcSInvalidTimeout
	case 0x000006AE:
		return RpcSObjectNotFound
	case 0x000006AF:
		return RpcSAlreadyRegistered
	case 0x000006B0:
		return RpcSTypeAlreadyRegistered
	case 0x000006B1:
		return RpcSAlreadyListening
	case 0x000006B2:
		return RpcSNoProtseqsRegistered
	case 0x000006B3:
		return RpcSNotListening
	case 0x000006B4:
		return RpcSUnknownMgrType
	case 0x000006B5:
		return RpcSUnknownIf
	case 0x000006B6:
		return RpcSNoBindings
	case 0x000006B7:
		return RpcSNoProtseqs
	case 0x000006B8:
		return RpcSCantCreateEndpoint
	case 0x000006B9:
		return RpcSOutOfResources
	case 0x000006BA:
		return RpcSServerUnavailable
	case 0x000006BB:
		return RpcSServerTooBusy
	case 0x000006BC:
		return RpcSInvalidNetworkOptions
	case 0x000006BD:
		return RpcSNoCallActive
	case 0x000006BE:
		return RpcSCallFailed
	case 0x000006BF:
		return RpcSCallFailedDne
	case 0x000006C0:
		return RpcSProtocolError
	case 0x000006C1:
		return RpcSProxyAccessDenied
	case 0x000006C2:
		return RpcSUnsupportedTransSyn
	case 0x000006C4:
		return RpcSUnsupportedType
	case 0x000006C5:
		return RpcSInvalidTag
	case 0x000006C6:
		return RpcSInvalidBound
	case 0x000006C7:
		return RpcSNoEntryName
	case 0x000006C8:
		return RpcSInvalidNameSyntax
	case 0x000006C9:
		return RpcSUnsupportedNameSyntax
	case 0x000006CB:
		return RpcSUuidNoAddress
	case 0x000006CC:
		return RpcSDuplicateEndpoint
	case 0x000006CD:
		return RpcSUnknownAuthnType
	case 0x000006CE:
		return RpcSMaxCallsTooSmall
	case 0x000006CF:
		return RpcSStringTooLong
	case 0x000006D0:
		return RpcSProtseqNotFound
	case 0x000006D1:
		return RpcSProcnumOutOfRange
	case 0x000006D2:
		return RpcSBindingHasNoAuth
	case 0x000006D3:
		return RpcSUnknownAuthnService
	case 0x000006D4:
		return RpcSUnknownAuthnLevel
	case 0x000006D5:
		return RpcSInvalidAuthIdentity
	case 0x000006D6:
		return RpcSUnknownAuthzService
	case 0x000006D7:
		return EptSInvalidEntry
	case 0x000006D8:
		return EptSCantPerformOp
	case 0x000006D9:
		return EptSNotRegistered
	case 0x000006DA:
		return RpcSNothingToExport
	case 0x000006DB:
		return RpcSIncompleteName
	case 0x000006DC:
		return RpcSInvalidVersOption
	case 0x000006DD:
		return RpcSNoMoreMembers
	case 0x000006DE:
		return RpcSNotAllObjsUnexported
	case 0x000006DF:
		return RpcSInterfaceNotFound
	case 0x000006E0:
		return RpcSEntryAlreadyExists
	case 0x000006E1:
		return RpcSEntryNotFound
	case 0x000006E2:
		return RpcSNameServiceUnavailable
	case 0x000006E3:
		return RpcSInvalidNafId
	case 0x000006E4:
		return RpcSCannotSupport
	case 0x000006E5:
		return RpcSNoContextAvailable
	case 0x000006E6:
		return RpcSInternalError
	case 0x000006E7:
		return RpcSZeroDivide
	case 0x000006E8:
		return RpcSAddressError
	case 0x000006E9:
		return RpcSFpDivZero
	case 0x000006EA:
		return RpcSFpUnderflow
	case 0x000006EB:
		return RpcSFpOverflow
	case 0x000006EC:
		return RpcXNoMoreEntries
	case 0x000006ED:
		return RpcXSsCharTransOpenFail
	case 0x000006EE:
		return RpcXSsCharTransShortFile
	case 0x000006EF:
		return RpcXSsInNullContext
	case 0x000006F1:
		return RpcXSsContextDamaged
	case 0x000006F2:
		return RpcXSsHandlesMismatch
	case 0x000006F3:
		return RpcXSsCannotGetCallHandle
	case 0x000006F4:
		return RpcXNullRefPointer
	case 0x000006F5:
		return RpcXEnumValueOutOfRange
	case 0x000006F6:
		return RpcXByteCountTooSmall
	case 0x000006F7:
		return RpcXBadStubData
	case 0x000006F8:
		return ErrorInvalidUserBuffer
	case 0x000006F9:
		return ErrorUnrecognizedMedia
	case 0x000006FA:
		return ErrorNoTrustLsaSecret
	case 0x000006FB:
		return ErrorNoTrustSamAccount
	case 0x000006FC:
		return ErrorTrustedDomainFailure
	case 0x000006FD:
		return ErrorTrustedRelationshipFailure
	case 0x000006FE:
		return ErrorTrustFailure
	case 0x000006FF:
		return RpcSCallInProgress
	case 0x00000700:
		return ErrorNetlogonNotStarted
	case 0x00000701:
		return ErrorAccountExpired
	case 0x00000702:
		return ErrorRedirectorHasOpenHandles
	case 0x00000703:
		return ErrorPrinterDriverAlreadyInstalled
	case 0x00000704:
		return ErrorUnknownPort
	case 0x00000705:
		return ErrorUnknownPrinterDriver
	case 0x00000706:
		return ErrorUnknownPrintprocessor
	case 0x00000707:
		return ErrorInvalidSeparatorFile
	case 0x00000708:
		return ErrorInvalidPriority
	case 0x00000709:
		return ErrorInvalidPrinterName
	case 0x0000070A:
		return ErrorPrinterAlreadyExists
	case 0x0000070B:
		return ErrorInvalidPrinterCommand
	case 0x0000070C:
		return ErrorInvalidDatatype
	case 0x0000070D:
		return ErrorInvalidEnvironment
	case 0x0000070E:
		return RpcSNoMoreBindings
	case 0x0000070F:
		return ErrorNologonInterdomainTrustAccount
	case 0x00000710:
		return ErrorNologonWorkstationTrustAccount
	case 0x00000711:
		return ErrorNologonServerTrustAccount
	case 0x00000712:
		return ErrorDomainTrustInconsistent
	case 0x00000713:
		return ErrorServerHasOpenHandles
	case 0x00000714:
		return ErrorResourceDataNotFound
	case 0x00000715:
		return ErrorResourceTypeNotFound
	case 0x00000716:
		return ErrorResourceNameNotFound
	case 0x00000717:
		return ErrorResourceLangNotFound
	case 0x00000718:
		return ErrorNotEnoughQuota
	case 0x00000719:
		return RpcSNoInterfaces
	case 0x0000071A:
		return RpcSCallCancelled
	case 0x0000071B:
		return RpcSBindingIncomplete
	case 0x0000071C:
		return RpcSCommFailure
	case 0x0000071D:
		return RpcSUnsupportedAuthnLevel
	case 0x0000071E:
		return RpcSNoPrincName
	case 0x0000071F:
		return RpcSNotRpcError
	case 0x00000720:
		return RpcSUuidLocalOnly
	case 0x00000721:
		return RpcSSecPkgError
	case 0x00000722:
		return RpcSNotCancelled
	case 0x00000723:
		return RpcXInvalidEsAction
	case 0x00000724:
		return RpcXWrongEsVersion
	case 0x00000725:
		return RpcXWrongStubVersion
	case 0x00000726:
		return RpcXInvalidPipeObject
	case 0x00000727:
		return RpcXWrongPipeOrder
	case 0x00000728:
		return RpcXWrongPipeVersion
	case 0x0000076A:
		return RpcSGroupMemberNotFound
	case 0x0000076B:
		return EptSCantCreate
	case 0x0000076C:
		return RpcSInvalidObject
	case 0x0000076D:
		return ErrorInvalidTime
	case 0x0000076E:
		return ErrorInvalidFormName
	case 0x0000076F:
		return ErrorInvalidFormSize
	case 0x00000770:
		return ErrorAlreadyWaiting
	case 0x00000771:
		return ErrorPrinterDeleted
	case 0x00000772:
		return ErrorInvalidPrinterState
	case 0x00000773:
		return ErrorPasswordMustChange
	case 0x00000774:
		return ErrorDomainControllerNotFound
	case 0x00000775:
		return ErrorAccountLockedOut
	case 0x00000776:
		return OrInvalidOxid
	case 0x00000777:
		return OrInvalidOid
	case 0x00000778:
		return OrInvalidSet
	case 0x00000779:
		return RpcSSendIncomplete
	case 0x0000077A:
		return RpcSInvalidAsyncHandle
	case 0x0000077B:
		return RpcSInvalidAsyncCall
	case 0x0000077C:
		return RpcXPipeClosed
	case 0x0000077D:
		return RpcXPipeDisciplineError
	case 0x0000077E:
		return RpcXPipeEmpty
	case 0x0000077F:
		return ErrorNoSitename
	case 0x00000780:
		return ErrorCantAccessFile
	case 0x00000781:
		return ErrorCantResolveFilename
	case 0x00000782:
		return RpcSEntryTypeMismatch
	case 0x00000783:
		return RpcSNotAllObjsExported
	case 0x00000784:
		return RpcSInterfaceNotExported
	case 0x00000785:
		return RpcSProfileNotAdded
	case 0x00000786:
		return RpcSPrfEltNotAdded
	case 0x00000787:
		return RpcSPrfEltNotRemoved
	case 0x00000788:
		return RpcSGrpEltNotAdded
	case 0x00000789:
		return RpcSGrpEltNotRemoved
	case 0x0000078A:
		return ErrorKmDriverBlocked
	case 0x0000078B:
		return ErrorContextExpired
	case 0x0000078C:
		return ErrorPerUserTrustQuotaExceeded
	case 0x0000078D:
		return ErrorAllUserTrustQuotaExceeded
	case 0x0000078E:
		return ErrorUserDeleteTrustQuotaExceeded
	case 0x0000078F:
		return ErrorAuthenticationFirewallFailed
	case 0x00000790:
		return ErrorRemotePrintConnectionsBlocked
	case 0x000007D0:
		return ErrorInvalidPixelFormat
	case 0x000007D1:
		return ErrorBadDriver
	case 0x000007D2:
		return ErrorInvalidWindowStyle
	case 0x000007D3:
		return ErrorMetafileNotSupported
	case 0x000007D4:
		return ErrorTransformNotSupported
	case 0x000007D5:
		return ErrorClippingNotSupported
	case 0x000007DA:
		return ErrorInvalidCmm
	case 0x000007DB:
		return ErrorInvalidProfile
	case 0x000007DC:
		return ErrorTagNotFound
	case 0x000007DD:
		return ErrorTagNotPresent
	case 0x000007DE:
		return ErrorDuplicateTag
	case 0x000007DF:
		return ErrorProfileNotAssociatedWithDevice
	case 0x000007E0:
		return ErrorProfileNotFound
	case 0x000007E1:
		return ErrorInvalidColorspace
	case 0x000007E2:
		return ErrorIcmNotEnabled
	case 0x000007E3:
		return ErrorDeletingIcmXform
	case 0x000007E4:
		return ErrorInvalidTransform
	case 0x000007E5:
		return ErrorColorspaceMismatch
	case 0x000007E6:
		return ErrorInvalidColorindex
	case 0x000007E7:
		return ErrorProfileDoesNotMatchDevice
	case 0x00000836:
		return NerrNetnotstarted
	case 0x00000837:
		return NerrUnknownserver
	case 0x00000838:
		return NerrSharemem
	case 0x00000839:
		return NerrNonetworkresource
	case 0x0000083A:
		return NerrRemoteonly
	case 0x0000083B:
		return NerrDevnotredirected
	case 0x0000083C:
		return ErrorConnectedOtherPassword
	case 0x0000083D:
		return ErrorConnectedOtherPasswordDefault
	case 0x00000842:
		return NerrServernotstarted
	case 0x00000843:
		return NerrItemnotfound
	case 0x00000844:
		return NerrUnknowndevdir
	case 0x00000845:
		return NerrRedirectedpath
	case 0x00000846:
		return NerrDuplicateshare
	case 0x00000847:
		return NerrNoroom
	case 0x00000849:
		return NerrToomanyitems
	case 0x0000084A:
		return NerrInvalidmaxusers
	case 0x0000084B:
		return NerrBuftoosmall
	case 0x0000084F:
		return NerrRemoteerr
	case 0x00000853:
		return NerrLanmaninierror
	case 0x00000858:
		return NerrNetworkerror
	case 0x00000859:
		return NerrWkstainconsistentstate
	case 0x0000085A:
		return NerrWkstanotstarted
	case 0x0000085B:
		return NerrBrowsernotstarted
	case 0x0000085C:
		return NerrInternalerror
	case 0x0000085D:
		return NerrBadtransactconfig
	case 0x0000085E:
		return NerrInvalidapi
	case 0x0000085F:
		return NerrBadeventname
	case 0x00000860:
		return NerrDupnamereboot
	case 0x00000862:
		return NerrCfgcompnotfound
	case 0x00000863:
		return NerrCfgparamnotfound
	case 0x00000865:
		return NerrLinetoolong
	case 0x00000866:
		return NerrQnotfound
	case 0x00000867:
		return NerrJobnotfound
	case 0x00000868:
		return NerrDestnotfound
	case 0x00000869:
		return NerrDestexists
	case 0x0000086A:
		return NerrQexists
	case 0x0000086B:
		return NerrQnoroom
	case 0x0000086C:
		return NerrJobnoroom
	case 0x0000086D:
		return NerrDestnoroom
	case 0x0000086E:
		return NerrDestidle
	case 0x0000086F:
		return NerrDestinvalidop
	case 0x00000870:
		return NerrProcnorespond
	case 0x00000871:
		return NerrSpoolernotloaded
	case 0x00000872:
		return NerrDestinvalidstate
	case 0x00000873:
		return NerrQinvalidstate
	case 0x00000874:
		return NerrJobinvalidstate
	case 0x00000875:
		return NerrSpoolnomemory
	case 0x00000876:
		return NerrDrivernotfound
	case 0x00000877:
		return NerrDatatypeinvalid
	case 0x00000878:
		return NerrProcnotfound
	case 0x00000884:
		return NerrServicetablelocked
	case 0x00000885:
		return NerrServicetablefull
	case 0x00000886:
		return NerrServiceinstalled
	case 0x00000887:
		return NerrServiceentrylocked
	case 0x00000888:
		return NerrServicenotinstalled
	case 0x00000889:
		return NerrBadservicename
	case 0x0000088A:
		return NerrServicectltimeout
	case 0x0000088B:
		return NerrServicectlbusy
	case 0x0000088C:
		return NerrBadserviceprogname
	case 0x0000088D:
		return NerrServicenotctrl
	case 0x0000088E:
		return NerrServicekillproc
	case 0x0000088F:
		return NerrServicectlnotvalid
	case 0x00000890:
		return NerrNotindispatchtbl
	case 0x00000891:
		return NerrBadcontrolrecv
	case 0x00000892:
		return NerrServicenotstarting
	case 0x00000898:
		return NerrAlreadyloggedon
	case 0x00000899:
		return NerrNotloggedon
	case 0x0000089A:
		return NerrBadusername
	case 0x0000089B:
		return NerrBadpassword
	case 0x0000089C:
		return NerrUnabletoaddnameW
	case 0x0000089D:
		return NerrUnabletoaddnameF
	case 0x0000089E:
		return NerrUnabletodelnameW
	case 0x0000089F:
		return NerrUnabletodelnameF
	case 0x000008A1:
		return NerrLogonspaused
	case 0x000008A2:
		return NerrLogonserverconflict
	case 0x000008A3:
		return NerrLogonnouserpath
	case 0x000008A4:
		return NerrLogonscripterror
	case 0x000008A6:
		return NerrStandalonelogon
	case 0x000008A7:
		return NerrLogonservernotfound
	case 0x000008A8:
		return NerrLogondomainexists
	case 0x000008A9:
		return NerrNonvalidatedlogon
	case 0x000008AB:
		return NerrAcfnotfound
	case 0x000008AC:
		return NerrGroupnotfound
	case 0x000008AD:
		return NerrUsernotfound
	case 0x000008AE:
		return NerrResourcenotfound
	case 0x000008AF:
		return NerrGroupexists
	case 0x000008B0:
		return NerrUserexists
	case 0x000008B1:
		return NerrResourceexists
	case 0x000008B2:
		return NerrNotprimary
	case 0x000008B3:
		return NerrAcfnotloaded
	case 0x000008B4:
		return NerrAcfnoroom
	case 0x000008B5:
		return NerrAcffileiofail
	case 0x000008B6:
		return NerrAcftoomanylists
	case 0x000008B7:
		return NerrUserlogon
	case 0x000008B8:
		return NerrAcfnoparent
	case 0x000008B9:
		return NerrCannotgrowsegment
	case 0x000008BA:
		return NerrSpegroupop
	case 0x000008BB:
		return NerrNotincache
	case 0x000008BC:
		return NerrUseringroup
	case 0x000008BD:
		return NerrUsernotingroup
	case 0x000008BE:
		return NerrAccountundefined
	case 0x000008BF:
		return NerrAccountexpired
	case 0x000008C0:
		return NerrInvalidworkstation
	case 0x000008C1:
		return NerrInvalidlogonhours
	case 0x000008C2:
		return NerrPasswordexpired
	case 0x000008C3:
		return NerrPasswordcantchange
	case 0x000008C4:
		return NerrPasswordhistconflict
	case 0x000008C5:
		return NerrPasswordtooshort
	case 0x000008C6:
		return NerrPasswordtoorecent
	case 0x000008C7:
		return NerrInvaliddatabase
	case 0x000008C8:
		return NerrDatabaseuptodate
	case 0x000008C9:
		return NerrSyncrequired
	case 0x000008CA:
		return NerrUsenotfound
	case 0x000008CB:
		return NerrBadasgtype
	case 0x000008CC:
		return NerrDeviceisshared
	case 0x000008DE:
		return NerrNocomputername
	case 0x000008DF:
		return NerrMsgalreadystarted
	case 0x000008E0:
		return NerrMsginitfailed
	case 0x000008E1:
		return NerrNamenotfound
	case 0x000008E2:
		return NerrAlreadyforwarded
	case 0x000008E3:
		return NerrAddforwarded
	case 0x000008E4:
		return NerrAlreadyexists
	case 0x000008E5:
		return NerrToomanynames
	case 0x000008E6:
		return NerrDelcomputername
	case 0x000008E7:
		return NerrLocalforward
	case 0x000008E8:
		return NerrGrpmsgprocessor
	case 0x000008E9:
		return NerrPausedremote
	case 0x000008EA:
		return NerrBadreceive
	case 0x000008EB:
		return NerrNameinuse
	case 0x000008EC:
		return NerrMsgnotstarted
	case 0x000008ED:
		return NerrNotlocalname
	case 0x000008EE:
		return NerrNoforwardname
	case 0x000008EF:
		return NerrRemotefull
	case 0x000008F0:
		return NerrNamenotforwarded
	case 0x000008F1:
		return NerrTruncatedbroadcast
	case 0x000008F6:
		return NerrInvaliddevice
	case 0x000008F7:
		return NerrWritefault
	case 0x000008F9:
		return NerrDuplicatename
	case 0x000008FA:
		return NerrDeletelater
	case 0x000008FB:
		return NerrIncompletedel
	case 0x000008FC:
		return NerrMultiplenets
	case 0x00000906:
		return NerrNetnamenotfound
	case 0x00000907:
		return NerrDevicenotshared
	case 0x00000908:
		return NerrClientnamenotfound
	case 0x0000090A:
		return NerrFileidnotfound
	case 0x0000090B:
		return NerrExecfailure
	case 0x0000090C:
		return NerrTmpfile
	case 0x0000090D:
		return NerrToomuchdata
	case 0x0000090E:
		return NerrDeviceshareconflict
	case 0x0000090F:
		return NerrBrowsertableincomplete
	case 0x00000910:
		return NerrNotlocaldomain
	case 0x00000911:
		return NerrIsdfsshare
	case 0x0000091B:
		return NerrDevinvalidopcode
	case 0x0000091C:
		return NerrDevnotfound
	case 0x0000091D:
		return NerrDevnotopen
	case 0x0000091E:
		return NerrBadqueuedevstring
	case 0x0000091F:
		return NerrBadqueuepriority
	case 0x00000921:
		return NerrNocommdevs
	case 0x00000922:
		return NerrQueuenotfound
	case 0x00000924:
		return NerrBaddevstring
	case 0x00000925:
		return NerrBaddev
	case 0x00000926:
		return NerrInusebyspooler
	case 0x00000927:
		return NerrCommdevinuse
	case 0x0000092F:
		return NerrInvalidcomputer
	case 0x00000932:
		return NerrMaxlenexceeded
	case 0x00000934:
		return NerrBadcomponent
	case 0x00000935:
		return NerrCanttype
	case 0x0000093A:
		return NerrToomanyentries
	case 0x00000942:
		return NerrProfilefiletoobig
	case 0x00000943:
		return NerrProfileoffset
	case 0x00000944:
		return NerrProfilecleanup
	case 0x00000945:
		return NerrProfileunknowncmd
	case 0x00000946:
		return NerrProfileloaderr
	case 0x00000947:
		return NerrProfilesaveerr
	case 0x00000949:
		return NerrLogoverflow
	case 0x0000094A:
		return NerrLogfilechanged
	case 0x0000094B:
		return NerrLogfilecorrupt
	case 0x0000094C:
		return NerrSourceisdir
	case 0x0000094D:
		return NerrBadsource
	case 0x0000094E:
		return NerrBaddest
	case 0x0000094F:
		return NerrDifferentservers
	case 0x00000951:
		return NerrRunsrvpaused
	case 0x00000955:
		return NerrErrcommrunsrv
	case 0x00000957:
		return NerrErrorexecingghost
	case 0x00000958:
		return NerrSharenotfound
	case 0x00000960:
		return NerrInvalidlana
	case 0x00000961:
		return NerrOpenfiles
	case 0x00000962:
		return NerrActiveconns
	case 0x00000963:
		return NerrBadpasswordcore
	case 0x00000964:
		return NerrDevinuse
	case 0x00000965:
		return NerrLocaldrive
	case 0x0000097E:
		return NerrAlertexists
	case 0x0000097F:
		return NerrToomanyalerts
	case 0x00000980:
		return NerrNosuchalert
	case 0x00000981:
		return NerrBadrecipient
	case 0x00000982:
		return NerrAcctlimitexceeded
	case 0x00000988:
		return NerrInvalidlogseek
	case 0x00000992:
		return NerrBaduasconfig
	case 0x00000993:
		return NerrInvaliduasop
	case 0x00000994:
		return NerrLastadmin
	case 0x00000995:
		return NerrDcnotfound
	case 0x00000996:
		return NerrLogontrackingerror
	case 0x00000997:
		return NerrNetlogonnotstarted
	case 0x00000998:
		return NerrCannotgrowuasfile
	case 0x00000999:
		return NerrTimediffatdc
	case 0x0000099A:
		return NerrPasswordmismatch
	case 0x0000099C:
		return NerrNosuchserver
	case 0x0000099D:
		return NerrNosuchsession
	case 0x0000099E:
		return NerrNosuchconnection
	case 0x0000099F:
		return NerrToomanyservers
	case 0x000009A0:
		return NerrToomanysessions
	case 0x000009A1:
		return NerrToomanyconnections
	case 0x000009A2:
		return NerrToomanyfiles
	case 0x000009A3:
		return NerrNoalternateservers
	case 0x000009A6:
		return NerrTrydownlevel
	case 0x000009B0:
		return NerrUpsdrivernotstarted
	case 0x000009B1:
		return NerrUpsinvalidconfig
	case 0x000009B2:
		return NerrUpsinvalidcommport
	case 0x000009B3:
		return NerrUpssignalasserted
	case 0x000009B4:
		return NerrUpsshutdownfailed
	case 0x000009C4:
		return NerrBaddosretcode
	case 0x000009C5:
		return NerrProgneedsextramem
	case 0x000009C6:
		return NerrBaddosfunction
	case 0x000009C7:
		return NerrRemotebootfailed
	case 0x000009C8:
		return NerrBadfilechecksum
	case 0x000009C9:
		return NerrNorplbootsystem
	case 0x000009CA:
		return NerrRplloadrnetbioserr
	case 0x000009CB:
		return NerrRplloadrdiskerr
	case 0x000009CC:
		return NerrImageparamerr
	case 0x000009CD:
		return NerrToomanyimageparams
	case 0x000009CE:
		return NerrNondosfloppyused
	case 0x000009CF:
		return NerrRplbootrestart
	case 0x000009D0:
		return NerrRplsrvrcallfailed
	case 0x000009D1:
		return NerrCantconnectrplsrvr
	case 0x000009D2:
		return NerrCantopenimagefile
	case 0x000009D3:
		return NerrCallingrplsrvr
	case 0x000009D4:
		return NerrStartingrplboot
	case 0x000009D5:
		return NerrRplbootserviceterm
	case 0x000009D6:
		return NerrRplbootstartfailed
	case 0x000009D7:
		return NerrRplConnected
	case 0x000009F6:
		return NerrBrowserconfiguredtonotrun
	case 0x00000A32:
		return NerrRplnoadaptersstarted
	case 0x00000A33:
		return NerrRplbadregistry
	case 0x00000A34:
		return NerrRplbaddatabase
	case 0x00000A35:
		return NerrRplrplfilesshare
	case 0x00000A36:
		return NerrRplnotrplserver
	case 0x00000A37:
		return NerrRplcannotenum
	case 0x00000A38:
		return NerrRplwkstainfocorrupted
	case 0x00000A39:
		return NerrRplwkstanotfound
	case 0x00000A3A:
		return NerrRplwkstanameunavailable
	case 0x00000A3B:
		return NerrRplprofileinfocorrupted
	case 0x00000A3C:
		return NerrRplprofilenotfound
	case 0x00000A3D:
		return NerrRplprofilenameunavailable
	case 0x00000A3E:
		return NerrRplprofilenotempty
	case 0x00000A3F:
		return NerrRplconfiginfocorrupted
	case 0x00000A40:
		return NerrRplconfignotfound
	case 0x00000A41:
		return NerrRpladapterinfocorrupted
	case 0x00000A42:
		return NerrRplinternal
	case 0x00000A43:
		return NerrRplvendorinfocorrupted
	case 0x00000A44:
		return NerrRplbootinfocorrupted
	case 0x00000A45:
		return NerrRplwkstaneedsuseracct
	case 0x00000A46:
		return NerrRplneedsrpluseracct
	case 0x00000A47:
		return NerrRplbootnotfound
	case 0x00000A48:
		return NerrRplincompatibleprofile
	case 0x00000A49:
		return NerrRpladapternameunavailable
	case 0x00000A4A:
		return NerrRplconfignotempty
	case 0x00000A4B:
		return NerrRplbootinuse
	case 0x00000A4C:
		return NerrRplbackupdatabase
	case 0x00000A4D:
		return NerrRpladapternotfound
	case 0x00000A4E:
		return NerrRplvendornotfound
	case 0x00000A4F:
		return NerrRplvendornameunavailable
	case 0x00000A50:
		return NerrRplbootnameunavailable
	case 0x00000A51:
		return NerrRplconfignameunavailable
	case 0x00000A64:
		return NerrDfsinternalcorruption
	case 0x00000A65:
		return NerrDfsvolumedatacorrupt
	case 0x00000A66:
		return NerrDfsnosuchvolume
	case 0x00000A67:
		return NerrDfsvolumealreadyexists
	case 0x00000A68:
		return NerrDfsalreadyshared
	case 0x00000A69:
		return NerrDfsnosuchshare
	case 0x00000A6A:
		return NerrDfsnotaleafvolume
	case 0x00000A6B:
		return NerrDfsleafvolume
	case 0x00000A6C:
		return NerrDfsvolumehasmultipleservers
	case 0x00000A6D:
		return NerrDfscantcreatejunctionpoint
	case 0x00000A6E:
		return NerrDfsservernotdfsaware
	case 0x00000A6F:
		return NerrDfsbadrenamepath
	case 0x00000A70:
		return NerrDfsvolumeisoffline
	case 0x00000A71:
		return NerrDfsnosuchserver
	case 0x00000A72:
		return NerrDfscyclicalname
	case 0x00000A73:
		return NerrDfsnotsupportedinserverdfs
	case 0x00000A74:
		return NerrDfsduplicateservice
	case 0x00000A75:
		return NerrDfscantremovelastservershare
	case 0x00000A76:
		return NerrDfsvolumeisinterdfs
	case 0x00000A77:
		return NerrDfsinconsistent
	case 0x00000A78:
		return NerrDfsserverupgraded
	case 0x00000A79:
		return NerrDfsdataisidentical
	case 0x00000A7A:
		return NerrDfscantremovedfsroot
	case 0x00000A7B:
		return NerrDfschildorparentindfs
	case 0x00000A82:
		return NerrDfsinternalerror
	case 0x00000A83:
		return NerrSetupalreadyjoined
	case 0x00000A84:
		return NerrSetupnotjoined
	case 0x00000A85:
		return NerrSetupdomaincontroller
	case 0x00000A86:
		return NerrDefaultjoinrequired
	case 0x00000A87:
		return NerrInvalidworkgroupname
	case 0x00000A88:
		return NerrNameusesincompatiblecodepage
	case 0x00000A89:
		return NerrComputeraccountnotfound
	case 0x00000A8A:
		return NerrPersonalsku
	case 0x00000A8D:
		return NerrPasswordmustchange
	case 0x00000A8E:
		return NerrAccountlockedout
	case 0x00000A8F:
		return NerrPasswordtoolong
	case 0x00000A90:
		return NerrPasswordnotcomplexenough
	case 0x00000A91:
		return NerrPasswordfiltererror
	case 0x00000BB8:
		return ErrorUnknownPrintMonitor
	case 0x00000BB9:
		return ErrorPrinterDriverInUse
	case 0x00000BBA:
		return ErrorSpoolFileNotFound
	case 0x00000BBB:
		return ErrorSplNoStartdoc
	case 0x00000BBC:
		return ErrorSplNoAddjob
	case 0x00000BBD:
		return ErrorPrintProcessorAlreadyInstalled
	case 0x00000BBE:
		return ErrorPrintMonitorAlreadyInstalled
	case 0x00000BBF:
		return ErrorInvalidPrintMonitor
	case 0x00000BC0:
		return ErrorPrintMonitorInUse
	case 0x00000BC1:
		return ErrorPrinterHasJobsQueued
	case 0x00000BC2:
		return ErrorSuccessRebootRequired
	case 0x00000BC3:
		return ErrorSuccessRestartRequired
	case 0x00000BC4:
		return ErrorPrinterNotFound
	case 0x00000BC5:
		return ErrorPrinterDriverWarned
	case 0x00000BC6:
		return ErrorPrinterDriverBlocked
	case 0x00000BC7:
		return ErrorPrinterDriverPackageInUse
	case 0x00000BC8:
		return ErrorCoreDriverPackageNotFound
	case 0x00000BC9:
		return ErrorFailRebootRequired
	case 0x00000BCA:
		return ErrorFailRebootInitiated
	case 0x00000BCB:
		return ErrorPrinterDriverDownloadNeeded
	case 0x00000BCE:
		return ErrorPrinterNotShareable
	case 0x00000F6E:
		return ErrorIoReissueAsCached
	case 0x00000FA0:
		return ErrorWinsInternal
	case 0x00000FA1:
		return ErrorCanNotDelLocalWins
	case 0x00000FA2:
		return ErrorStaticInit
	case 0x00000FA3:
		return ErrorIncBackup
	case 0x00000FA4:
		return ErrorFullBackup
	case 0x00000FA5:
		return ErrorRecNonExistent
	case 0x00000FA6:
		return ErrorRplNotAllowed
	case 0x00000FD2:
		return PeerdistErrorContentinfoVersionUnsupported
	case 0x00000FD3:
		return PeerdistErrorCannotParseContentinfo
	case 0x00000FD4:
		return PeerdistErrorMissingData
	case 0x00000FD5:
		return PeerdistErrorNoMore
	case 0x00000FD6:
		return PeerdistErrorNotInitialized
	case 0x00000FD7:
		return PeerdistErrorAlreadyInitialized
	case 0x00000FD8:
		return PeerdistErrorShutdownInProgress
	case 0x00000FD9:
		return PeerdistErrorInvalidated
	case 0x00000FDA:
		return PeerdistErrorAlreadyExists
	case 0x00000FDB:
		return PeerdistErrorOperationNotfound
	case 0x00000FDC:
		return PeerdistErrorAlreadyCompleted
	case 0x00000FDD:
		return PeerdistErrorOutOfBounds
	case 0x00000FDE:
		return PeerdistErrorVersionUnsupported
	case 0x00000FDF:
		return PeerdistErrorInvalidConfiguration
	case 0x00000FE0:
		return PeerdistErrorNotLicensed
	case 0x00000FE1:
		return PeerdistErrorServiceUnavailable
	case 0x00001004:
		return ErrorDhcpAddressConflict
	case 0x00001068:
		return ErrorWmiGuidNotFound
	case 0x00001069:
		return ErrorWmiInstanceNotFound
	case 0x0000106A:
		return ErrorWmiItemidNotFound
	case 0x0000106B:
		return ErrorWmiTryAgain
	case 0x0000106C:
		return ErrorWmiDpNotFound
	case 0x0000106D:
		return ErrorWmiUnresolvedInstanceRef
	case 0x0000106E:
		return ErrorWmiAlreadyEnabled
	case 0x0000106F:
		return ErrorWmiGuidDisconnected
	case 0x00001070:
		return ErrorWmiServerUnavailable
	case 0x00001071:
		return ErrorWmiDpFailed
	case 0x00001072:
		return ErrorWmiInvalidMof
	case 0x00001073:
		return ErrorWmiInvalidReginfo
	case 0x00001074:
		return ErrorWmiAlreadyDisabled
	case 0x00001075:
		return ErrorWmiReadOnly
	case 0x00001076:
		return ErrorWmiSetFailure
	case 0x000010CC:
		return ErrorInvalidMedia
	case 0x000010CD:
		return ErrorInvalidLibrary
	case 0x000010CE:
		return ErrorInvalidMediaPool
	case 0x000010CF:
		return ErrorDriveMediaMismatch
	case 0x000010D0:
		return ErrorMediaOffline
	case 0x000010D1:
		return ErrorLibraryOffline
	case 0x000010D2:
		return ErrorEmpty
	case 0x000010D3:
		return ErrorNotEmpty
	case 0x000010D4:
		return ErrorMediaUnavailable
	case 0x000010D5:
		return ErrorResourceDisabled
	case 0x000010D6:
		return ErrorInvalidCleaner
	case 0x000010D7:
		return ErrorUnableToClean
	case 0x000010D8:
		return ErrorObjectNotFound
	case 0x000010D9:
		return ErrorDatabaseFailure
	case 0x000010DA:
		return ErrorDatabaseFull
	case 0x000010DB:
		return ErrorMediaIncompatible
	case 0x000010DC:
		return ErrorResourceNotPresent
	case 0x000010DD:
		return ErrorInvalidOperation
	case 0x000010DE:
		return ErrorMediaNotAvailable
	case 0x000010DF:
		return ErrorDeviceNotAvailable
	case 0x000010E0:
		return ErrorRequestRefused
	case 0x000010E1:
		return ErrorInvalidDriveObject
	case 0x000010E2:
		return ErrorLibraryFull
	case 0x000010E3:
		return ErrorMediumNotAccessible
	case 0x000010E4:
		return ErrorUnableToLoadMedium
	case 0x000010E5:
		return ErrorUnableToInventoryDrive
	case 0x000010E6:
		return ErrorUnableToInventorySlot
	case 0x000010E7:
		return ErrorUnableToInventoryTransport
	case 0x000010E8:
		return ErrorTransportFull
	case 0x000010E9:
		return ErrorControllingIeport
	case 0x000010EA:
		return ErrorUnableToEjectMountedMedia
	case 0x000010EB:
		return ErrorCleanerSlotSet
	case 0x000010EC:
		return ErrorCleanerSlotNotSet
	case 0x000010ED:
		return ErrorCleanerCartridgeSpent
	case 0x000010EE:
		return ErrorUnexpectedOmid
	case 0x000010EF:
		return ErrorCantDeleteLastItem
	case 0x000010F0:
		return ErrorMessageExceedsMaxSize
	case 0x000010F1:
		return ErrorVolumeContainsSysFiles
	case 0x000010F2:
		return ErrorIndigenousType
	case 0x000010F3:
		return ErrorNoSupportingDrives
	case 0x000010F4:
		return ErrorCleanerCartridgeInstalled
	case 0x000010F5:
		return ErrorIeportFull
	case 0x000010FE:
		return ErrorFileOffline
	case 0x000010FF:
		return ErrorRemoteStorageNotActive
	case 0x00001100:
		return ErrorRemoteStorageMediaError
	case 0x00001126:
		return ErrorNotAReparsePoint
	case 0x00001127:
		return ErrorReparseAttributeConflict
	case 0x00001128:
		return ErrorInvalidReparseData
	case 0x00001129:
		return ErrorReparseTagInvalid
	case 0x0000112A:
		return ErrorReparseTagMismatch
	case 0x00001194:
		return ErrorVolumeNotSisEnabled
	case 0x00001389:
		return ErrorDependentResourceExists
	case 0x0000138A:
		return ErrorDependencyNotFound
	case 0x0000138B:
		return ErrorDependencyAlreadyExists
	case 0x0000138C:
		return ErrorResourceNotOnline
	case 0x0000138D:
		return ErrorHostNodeNotAvailable
	case 0x0000138E:
		return ErrorResourceNotAvailable
	case 0x0000138F:
		return ErrorResourceNotFound
	case 0x00001390:
		return ErrorShutdownCluster
	case 0x00001391:
		return ErrorCantEvictActiveNode
	case 0x00001392:
		return ErrorObjectAlreadyExists
	case 0x00001393:
		return ErrorObjectInList
	case 0x00001394:
		return ErrorGroupNotAvailable
	case 0x00001395:
		return ErrorGroupNotFound
	case 0x00001396:
		return ErrorGroupNotOnline
	case 0x00001397:
		return ErrorHostNodeNotResourceOwner
	case 0x00001398:
		return ErrorHostNodeNotGroupOwner
	case 0x00001399:
		return ErrorResmonCreateFailed
	case 0x0000139A:
		return ErrorResmonOnlineFailed
	case 0x0000139B:
		return ErrorResourceOnline
	case 0x0000139C:
		return ErrorQuorumResource
	case 0x0000139D:
		return ErrorNotQuorumCapable
	case 0x0000139E:
		return ErrorClusterShuttingDown
	case 0x0000139F:
		return ErrorInvalidState
	case 0x000013A0:
		return ErrorResourcePropertiesStored
	case 0x000013A1:
		return ErrorNotQuorumClass
	case 0x000013A2:
		return ErrorCoreResource
	case 0x000013A3:
		return ErrorQuorumResourceOnlineFailed
	case 0x000013A4:
		return ErrorQuorumlogOpenFailed
	case 0x000013A5:
		return ErrorClusterlogCorrupt
	case 0x000013A6:
		return ErrorClusterlogRecordExceedsMaxsize
	case 0x000013A7:
		return ErrorClusterlogExceedsMaxsize
	case 0x000013A8:
		return ErrorClusterlogChkpointNotFound
	case 0x000013A9:
		return ErrorClusterlogNotEnoughSpace
	case 0x000013AA:
		return ErrorQuorumOwnerAlive
	case 0x000013AB:
		return ErrorNetworkNotAvailable
	case 0x000013AC:
		return ErrorNodeNotAvailable
	case 0x000013AD:
		return ErrorAllNodesNotAvailable
	case 0x000013AE:
		return ErrorResourceFailed
	case 0x000013AF:
		return ErrorClusterInvalidNode
	case 0x000013B0:
		return ErrorClusterNodeExists
	case 0x000013B1:
		return ErrorClusterJoinInProgress
	case 0x000013B2:
		return ErrorClusterNodeNotFound
	case 0x000013B3:
		return ErrorClusterLocalNodeNotFound
	case 0x000013B4:
		return ErrorClusterNetworkExists
	case 0x000013B5:
		return ErrorClusterNetworkNotFound
	case 0x000013B6:
		return ErrorClusterNetinterfaceExists
	case 0x000013B7:
		return ErrorClusterNetinterfaceNotFound
	case 0x000013B8:
		return ErrorClusterInvalidRequest
	case 0x000013B9:
		return ErrorClusterInvalidNetworkProvider
	case 0x000013BA:
		return ErrorClusterNodeDown
	case 0x000013BB:
		return ErrorClusterNodeUnreachable
	case 0x000013BC:
		return ErrorClusterNodeNotMember
	case 0x000013BD:
		return ErrorClusterJoinNotInProgress
	case 0x000013BE:
		return ErrorClusterInvalidNetwork
	case 0x000013C0:
		return ErrorClusterNodeUp
	case 0x000013C1:
		return ErrorClusterIpaddrInUse
	case 0x000013C2:
		return ErrorClusterNodeNotPaused
	case 0x000013C3:
		return ErrorClusterNoSecurityContext
	case 0x000013C4:
		return ErrorClusterNetworkNotInternal
	case 0x000013C5:
		return ErrorClusterNodeAlreadyUp
	case 0x000013C6:
		return ErrorClusterNodeAlreadyDown
	case 0x000013C7:
		return ErrorClusterNetworkAlreadyOnline
	case 0x000013C8:
		return ErrorClusterNetworkAlreadyOffline
	case 0x000013C9:
		return ErrorClusterNodeAlreadyMember
	case 0x000013CA:
		return ErrorClusterLastInternalNetwork
	case 0x000013CB:
		return ErrorClusterNetworkHasDependents
	case 0x000013CC:
		return ErrorInvalidOperationOnQuorum
	case 0x000013CD:
		return ErrorDependencyNotAllowed
	case 0x000013CE:
		return ErrorClusterNodePaused
	case 0x000013CF:
		return ErrorNodeCantHostResource
	case 0x000013D0:
		return ErrorClusterNodeNotReady
	case 0x000013D1:
		return ErrorClusterNodeShuttingDown
	case 0x000013D2:
		return ErrorClusterJoinAborted
	case 0x000013D3:
		return ErrorClusterIncompatibleVersions
	case 0x000013D4:
		return ErrorClusterMaxnumOfResourcesExceeded
	case 0x000013D5:
		return ErrorClusterSystemConfigChanged
	case 0x000013D6:
		return ErrorClusterResourceTypeNotFound
	case 0x000013D7:
		return ErrorClusterRestypeNotSupported
	case 0x000013D8:
		return ErrorClusterResnameNotFound
	case 0x000013D9:
		return ErrorClusterNoRpcPackagesRegistered
	case 0x000013DA:
		return ErrorClusterOwnerNotInPreflist
	case 0x000013DB:
		return ErrorClusterDatabaseSeqmismatch
	case 0x000013DC:
		return ErrorResmonInvalidState
	case 0x000013DD:
		return ErrorClusterGumNotLocker
	case 0x000013DE:
		return ErrorQuorumDiskNotFound
	case 0x000013DF:
		return ErrorDatabaseBackupCorrupt
	case 0x000013E0:
		return ErrorClusterNodeAlreadyHasDfsRoot
	case 0x000013E1:
		return ErrorResourcePropertyUnchangeable
	case 0x00001702:
		return ErrorClusterMembershipInvalidState
	case 0x00001703:
		return ErrorClusterQuorumlogNotFound
	case 0x00001704:
		return ErrorClusterMembershipHalt
	case 0x00001705:
		return ErrorClusterInstanceIdMismatch
	case 0x00001706:
		return ErrorClusterNetworkNotFoundForIp
	case 0x00001707:
		return ErrorClusterPropertyDataTypeMismatch
	case 0x00001708:
		return ErrorClusterEvictWithoutCleanup
	case 0x00001709:
		return ErrorClusterParameterMismatch
	case 0x0000170A:
		return ErrorNodeCannotBeClustered
	case 0x0000170B:
		return ErrorClusterWrongOsVersion
	case 0x0000170C:
		return ErrorClusterCantCreateDupClusterName
	case 0x0000170D:
		return ErrorCluscfgAlreadyCommitted
	case 0x0000170E:
		return ErrorCluscfgRollbackFailed
	case 0x0000170F:
		return ErrorCluscfgSystemDiskDriveLetterConflict
	case 0x00001710:
		return ErrorClusterOldVersion
	case 0x00001711:
		return ErrorClusterMismatchedComputerAcctName
	case 0x00001712:
		return ErrorClusterNoNetAdapters
	case 0x00001713:
		return ErrorClusterPoisoned
	case 0x00001714:
		return ErrorClusterGroupMoving
	case 0x00001715:
		return ErrorClusterResourceTypeBusy
	case 0x00001716:
		return ErrorResourceCallTimedOut
	case 0x00001717:
		return ErrorInvalidClusterIpv6Address
	case 0x00001718:
		return ErrorClusterInternalInvalidFunction
	case 0x00001719:
		return ErrorClusterParameterOutOfBounds
	case 0x0000171A:
		return ErrorClusterPartialSend
	case 0x0000171B:
		return ErrorClusterRegistryInvalidFunction
	case 0x0000171C:
		return ErrorClusterInvalidStringTermination
	case 0x0000171D:
		return ErrorClusterInvalidStringFormat
	case 0x0000171E:
		return ErrorClusterDatabaseTransactionInProgress
	case 0x0000171F:
		return ErrorClusterDatabaseTransactionNotInProgress
	case 0x00001720:
		return ErrorClusterNullData
	case 0x00001721:
		return ErrorClusterPartialRead
	case 0x00001722:
		return ErrorClusterPartialWrite
	case 0x00001723:
		return ErrorClusterCantDeserializeData
	case 0x00001724:
		return ErrorDependentResourcePropertyConflict
	case 0x00001725:
		return ErrorClusterNoQuorum
	case 0x00001726:
		return ErrorClusterInvalidIpv6Network
	case 0x00001727:
		return ErrorClusterInvalidIpv6TunnelNetwork
	case 0x00001728:
		return ErrorQuorumNotAllowedInThisGroup
	case 0x00001770:
		return ErrorEncryptionFailed
	case 0x00001771:
		return ErrorDecryptionFailed
	case 0x00001772:
		return ErrorFileEncrypted
	case 0x00001773:
		return ErrorNoRecoveryPolicy
	case 0x00001774:
		return ErrorNoEfs
	case 0x00001775:
		return ErrorWrongEfs
	case 0x00001776:
		return ErrorNoUserKeys
	case 0x00001777:
		return ErrorFileNotEncrypted
	case 0x00001778:
		return ErrorNotExportFormat
	case 0x00001779:
		return ErrorFileReadOnly
	case 0x0000177A:
		return ErrorDirEfsDisallowed
	case 0x0000177B:
		return ErrorEfsServerNotTrusted
	case 0x0000177C:
		return ErrorBadRecoveryPolicy
	case 0x0000177D:
		return ErrorEfsAlgBlobTooBig
	case 0x0000177E:
		return ErrorVolumeNotSupportEfs
	case 0x0000177F:
		return ErrorEfsDisabled
	case 0x00001780:
		return ErrorEfsVersionNotSupport
	case 0x00001781:
		return ErrorCsEncryptionInvalidServerResponse
	case 0x00001782:
		return ErrorCsEncryptionUnsupportedServer
	case 0x00001783:
		return ErrorCsEncryptionExistingEncryptedFile
	case 0x00001784:
		return ErrorCsEncryptionNewEncryptedFile
	case 0x00001785:
		return ErrorCsEncryptionFileNotCse
	case 0x000017E6:
		return ErrorNoBrowserServersFound
	case 0x00001838:
		return SchedEServiceNotLocalsystem
	case 0x000019C8:
		return ErrorLogSectorInvalid
	case 0x000019C9:
		return ErrorLogSectorParityInvalid
	case 0x000019CA:
		return ErrorLogSectorRemapped
	case 0x000019CB:
		return ErrorLogBlockIncomplete
	case 0x000019CC:
		return ErrorLogInvalidRange
	case 0x000019CD:
		return ErrorLogBlocksExhausted
	case 0x000019CE:
		return ErrorLogReadContextInvalid
	case 0x000019CF:
		return ErrorLogRestartInvalid
	case 0x000019D0:
		return ErrorLogBlockVersion
	case 0x000019D1:
		return ErrorLogBlockInvalid
	case 0x000019D2:
		return ErrorLogReadModeInvalid
	case 0x000019D3:
		return ErrorLogNoRestart
	case 0x000019D4:
		return ErrorLogMetadataCorrupt
	case 0x000019D5:
		return ErrorLogMetadataInvalid
	case 0x000019D6:
		return ErrorLogMetadataInconsistent
	case 0x000019D7:
		return ErrorLogReservationInvalid
	case 0x000019D8:
		return ErrorLogCantDelete
	case 0x000019D9:
		return ErrorLogContainerLimitExceeded
	case 0x000019DA:
		return ErrorLogStartOfLog
	case 0x000019DB:
		return ErrorLogPolicyAlreadyInstalled
	case 0x000019DC:
		return ErrorLogPolicyNotInstalled
	case 0x000019DD:
		return ErrorLogPolicyInvalid
	case 0x000019DE:
		return ErrorLogPolicyConflict
	case 0x000019DF:
		return ErrorLogPinnedArchiveTail
	case 0x000019E0:
		return ErrorLogRecordNonexistent
	case 0x000019E1:
		return ErrorLogRecordsReservedInvalid
	case 0x000019E2:
		return ErrorLogSpaceReservedInvalid
	case 0x000019E3:
		return ErrorLogTailInvalid
	case 0x000019E4:
		return ErrorLogFull
	case 0x000019E5:
		return ErrorCouldNotResizeLog
	case 0x000019E6:
		return ErrorLogMultiplexed
	case 0x000019E7:
		return ErrorLogDedicated
	case 0x000019E8:
		return ErrorLogArchiveNotInProgress
	case 0x000019E9:
		return ErrorLogArchiveInProgress
	case 0x000019EA:
		return ErrorLogEphemeral
	case 0x000019EB:
		return ErrorLogNotEnoughContainers
	case 0x000019EC:
		return ErrorLogClientAlreadyRegistered
	case 0x000019ED:
		return ErrorLogClientNotRegistered
	case 0x000019EE:
		return ErrorLogFullHandlerInProgress
	case 0x000019EF:
		return ErrorLogContainerReadFailed
	case 0x000019F0:
		return ErrorLogContainerWriteFailed
	case 0x000019F1:
		return ErrorLogContainerOpenFailed
	case 0x000019F2:
		return ErrorLogContainerStateInvalid
	case 0x000019F3:
		return ErrorLogStateInvalid
	case 0x000019F4:
		return ErrorLogPinned
	case 0x000019F5:
		return ErrorLogMetadataFlushFailed
	case 0x000019F6:
		return ErrorLogInconsistentSecurity
	case 0x000019F7:
		return ErrorLogAppendedFlushFailed
	case 0x000019F8:
		return ErrorLogPinnedReservation
	case 0x00001A2C:
		return ErrorInvalidTransaction
	case 0x00001A2D:
		return ErrorTransactionNotActive
	case 0x00001A2E:
		return ErrorTransactionRequestNotValid
	case 0x00001A2F:
		return ErrorTransactionNotRequested
	case 0x00001A30:
		return ErrorTransactionAlreadyAborted
	case 0x00001A31:
		return ErrorTransactionAlreadyCommitted
	case 0x00001A32:
		return ErrorTmInitializationFailed
	case 0x00001A33:
		return ErrorResourcemanagerReadOnly
	case 0x00001A34:
		return ErrorTransactionNotJoined
	case 0x00001A35:
		return ErrorTransactionSuperiorExists
	case 0x00001A36:
		return ErrorCrmProtocolAlreadyExists
	case 0x00001A37:
		return ErrorTransactionPropagationFailed
	case 0x00001A38:
		return ErrorCrmProtocolNotFound
	case 0x00001A39:
		return ErrorTransactionInvalidMarshallBuffer
	case 0x00001A3A:
		return ErrorCurrentTransactionNotValid
	case 0x00001A3B:
		return ErrorTransactionNotFound
	case 0x00001A3C:
		return ErrorResourcemanagerNotFound
	case 0x00001A3D:
		return ErrorEnlistmentNotFound
	case 0x00001A3E:
		return ErrorTransactionmanagerNotFound
	case 0x00001A3F:
		return ErrorTransactionmanagerNotOnline
	case 0x00001A40:
		return ErrorTransactionmanagerRecoveryNameCollision
	case 0x00001A90:
		return ErrorTransactionalConflict
	case 0x00001A91:
		return ErrorRmNotActive
	case 0x00001A92:
		return ErrorRmMetadataCorrupt
	case 0x00001A93:
		return ErrorDirectoryNotRm
	case 0x00001A95:
		return ErrorTransactionsUnsupportedRemote
	case 0x00001A96:
		return ErrorLogResizeInvalidSize
	case 0x00001A97:
		return ErrorObjectNoLongerExists
	case 0x00001A98:
		return ErrorStreamMiniversionNotFound
	case 0x00001A99:
		return ErrorStreamMiniversionNotValid
	case 0x00001A9A:
		return ErrorMiniversionInaccessibleFromSpecifiedTransaction
	case 0x00001A9B:
		return ErrorCantOpenMiniversionWithModifyIntent
	case 0x00001A9C:
		return ErrorCantCreateMoreStreamMiniversions
	case 0x00001A9E:
		return ErrorRemoteFileVersionMismatch
	case 0x00001A9F:
		return ErrorHandleNoLongerValid
	case 0x00001AA0:
		return ErrorNoTxfMetadata
	case 0x00001AA1:
		return ErrorLogCorruptionDetected
	case 0x00001AA2:
		return ErrorCantRecoverWithHandleOpen
	case 0x00001AA3:
		return ErrorRmDisconnected
	case 0x00001AA4:
		return ErrorEnlistmentNotSuperior
	case 0x00001AA5:
		return ErrorRecoveryNotNeeded
	case 0x00001AA6:
		return ErrorRmAlreadyStarted
	case 0x00001AA7:
		return ErrorFileIdentityNotPersistent
	case 0x00001AA8:
		return ErrorCantBreakTransactionalDependency
	case 0x00001AA9:
		return ErrorCantCrossRmBoundary
	case 0x00001AAA:
		return ErrorTxfDirNotEmpty
	case 0x00001AAB:
		return ErrorIndoubtTransactionsExist
	case 0x00001AAC:
		return ErrorTmVolatile
	case 0x00001AAD:
		return ErrorRollbackTimerExpired
	case 0x00001AAE:
		return ErrorTxfAttributeCorrupt
	case 0x00001AAF:
		return ErrorEfsNotAllowedInTransaction
	case 0x00001AB0:
		return ErrorTransactionalOpenNotAllowed
	case 0x00001AB1:
		return ErrorLogGrowthFailed
	case 0x00001AB2:
		return ErrorTransactedMappingUnsupportedRemote
	case 0x00001AB3:
		return ErrorTxfMetadataAlreadyPresent
	case 0x00001AB4:
		return ErrorTransactionScopeCallbacksNotSet
	case 0x00001AB5:
		return ErrorTransactionRequiredPromotion
	case 0x00001AB6:
		return ErrorCannotExecuteFileInTransaction
	case 0x00001AB7:
		return ErrorTransactionsNotFrozen
	case 0x00001AB8:
		return ErrorTransactionFreezeInProgress
	case 0x00001AB9:
		return ErrorNotSnapshotVolume
	case 0x00001ABA:
		return ErrorNoSavepointWithOpenFiles
	case 0x00001ABB:
		return ErrorDataLostRepair
	case 0x00001ABC:
		return ErrorSparseNotAllowedInTransaction
	case 0x00001ABD:
		return ErrorTmIdentityMismatch
	case 0x00001ABE:
		return ErrorFloatedSection
	case 0x00001ABF:
		return ErrorCannotAcceptTransactedWork
	case 0x00001AC0:
		return ErrorCannotAbortTransactions
	case 0x00001B59:
		return ErrorCtxWinstationNameInvalid
	case 0x00001B5A:
		return ErrorCtxInvalidPd
	case 0x00001B5B:
		return ErrorCtxPdNotFound
	case 0x00001B5C:
		return ErrorCtxWdNotFound
	case 0x00001B5D:
		return ErrorCtxCannotMakeEventlogEntry
	case 0x00001B5E:
		return ErrorCtxServiceNameCollision
	case 0x00001B5F:
		return ErrorCtxClosePending
	case 0x00001B60:
		return ErrorCtxNoOutbuf
	case 0x00001B61:
		return ErrorCtxModemInfNotFound
	case 0x00001B62:
		return ErrorCtxInvalidModemname
	case 0x00001B63:
		return ErrorCtxModemResponseError
	case 0x00001B64:
		return ErrorCtxModemResponseTimeout
	case 0x00001B65:
		return ErrorCtxModemResponseNoCarrier
	case 0x00001B66:
		return ErrorCtxModemResponseNoDialtone
	case 0x00001B67:
		return ErrorCtxModemResponseBusy
	case 0x00001B68:
		return ErrorCtxModemResponseVoice
	case 0x00001B69:
		return ErrorCtxTdError
	case 0x00001B6E:
		return ErrorCtxWinstationNotFound
	case 0x00001B6F:
		return ErrorCtxWinstationAlreadyExists
	case 0x00001B70:
		return ErrorCtxWinstationBusy
	case 0x00001B71:
		return ErrorCtxBadVideoMode
	case 0x00001B7B:
		return ErrorCtxGraphicsInvalid
	case 0x00001B7D:
		return ErrorCtxLogonDisabled
	case 0x00001B7E:
		return ErrorCtxNotConsole
	case 0x00001B80:
		return ErrorCtxClientQueryTimeout
	case 0x00001B81:
		return ErrorCtxConsoleDisconnect
	case 0x00001B82:
		return ErrorCtxConsoleConnect
	case 0x00001B84:
		return ErrorCtxShadowDenied
	case 0x00001B85:
		return ErrorCtxWinstationAccessDenied
	case 0x00001B89:
		return ErrorCtxInvalidWd
	case 0x00001B8A:
		return ErrorCtxShadowInvalid
	case 0x00001B8B:
		return ErrorCtxShadowDisabled
	case 0x00001B8C:
		return ErrorCtxClientLicenseInUse
	case 0x00001B8D:
		return ErrorCtxClientLicenseNotSet
	case 0x00001B8E:
		return ErrorCtxLicenseNotAvailable
	case 0x00001B8F:
		return ErrorCtxLicenseClientInvalid
	case 0x00001B90:
		return ErrorCtxLicenseExpired
	case 0x00001B91:
		return ErrorCtxShadowNotRunning
	case 0x00001B92:
		return ErrorCtxShadowEndedByModeChange
	case 0x00001B93:
		return ErrorActivationCountExceeded
	case 0x00001B94:
		return ErrorCtxWinstationsDisabled
	case 0x00001B95:
		return ErrorCtxEncryptionLevelRequired
	case 0x00001B96:
		return ErrorCtxSessionInUse
	case 0x00001B97:
		return ErrorCtxNoForceLogoff
	case 0x00001B98:
		return ErrorCtxAccountRestriction
	case 0x00001B99:
		return ErrorRdpProtocolError
	case 0x00001B9A:
		return ErrorCtxCdmConnect
	case 0x00001B9B:
		return ErrorCtxCdmDisconnect
	case 0x00001B9C:
		return ErrorCtxSecurityLayerError
	case 0x00001B9D:
		return ErrorTsIncompatibleSessions
	case 0x00001F41:
		return FrsErrInvalidApiSequence
	case 0x00001F42:
		return FrsErrStartingService
	case 0x00001F43:
		return FrsErrStoppingService
	case 0x00001F44:
		return FrsErrInternalApi
	case 0x00001F45:
		return FrsErrInternal
	case 0x00001F46:
		return FrsErrServiceComm
	case 0x00001F47:
		return FrsErrInsufficientPriv
	case 0x00001F48:
		return FrsErrAuthentication
	case 0x00001F49:
		return FrsErrParentInsufficientPriv
	case 0x00001F4A:
		return FrsErrParentAuthentication
	case 0x00001F4B:
		return FrsErrChildToParentComm
	case 0x00001F4C:
		return FrsErrParentToChildComm
	case 0x00001F4D:
		return FrsErrSysvolPopulate
	case 0x00001F4E:
		return FrsErrSysvolPopulateTimeout
	case 0x00001F4F:
		return FrsErrSysvolIsBusy
	case 0x00001F50:
		return FrsErrSysvolDemote
	case 0x00001F51:
		return FrsErrInvalidServiceParameter
	case 0x00002008:
		return ErrorDsNotInstalled
	case 0x00002009:
		return ErrorDsMembershipEvaluatedLocally
	case 0x0000200A:
		return ErrorDsNoAttributeOrValue
	case 0x0000200B:
		return ErrorDsInvalidAttributeSyntax
	case 0x0000200C:
		return ErrorDsAttributeTypeUndefined
	case 0x0000200D:
		return ErrorDsAttributeOrValueExists
	case 0x0000200E:
		return ErrorDsBusy
	case 0x0000200F:
		return ErrorDsUnavailable
	case 0x00002010:
		return ErrorDsNoRidsAllocated
	case 0x00002011:
		return ErrorDsNoMoreRids
	case 0x00002012:
		return ErrorDsIncorrectRoleOwner
	case 0x00002013:
		return ErrorDsRidmgrInitError
	case 0x00002014:
		return ErrorDsObjClassViolation
	case 0x00002015:
		return ErrorDsCantOnNonLeaf
	case 0x00002016:
		return ErrorDsCantOnRdn
	case 0x00002017:
		return ErrorDsCantModObjClass
	case 0x00002018:
		return ErrorDsCrossDomMoveError
	case 0x00002019:
		return ErrorDsGcNotAvailable
	case 0x0000201A:
		return ErrorSharedPolicy
	case 0x0000201B:
		return ErrorPolicyObjectNotFound
	case 0x0000201C:
		return ErrorPolicyOnlyInDs
	case 0x0000201D:
		return ErrorPromotionActive
	case 0x0000201E:
		return ErrorNoPromotionActive
	case 0x00002020:
		return ErrorDsOperationsError
	case 0x00002021:
		return ErrorDsProtocolError
	case 0x00002022:
		return ErrorDsTimelimitExceeded
	case 0x00002023:
		return ErrorDsSizelimitExceeded
	case 0x00002024:
		return ErrorDsAdminLimitExceeded
	case 0x00002025:
		return ErrorDsCompareFalse
	case 0x00002026:
		return ErrorDsCompareTrue
	case 0x00002027:
		return ErrorDsAuthMethodNotSupported
	case 0x00002028:
		return ErrorDsStrongAuthRequired
	case 0x00002029:
		return ErrorDsInappropriateAuth
	case 0x0000202A:
		return ErrorDsAuthUnknown
	case 0x0000202B:
		return ErrorDsReferral
	case 0x0000202C:
		return ErrorDsUnavailableCritExtension
	case 0x0000202D:
		return ErrorDsConfidentialityRequired
	case 0x0000202E:
		return ErrorDsInappropriateMatching
	case 0x0000202F:
		return ErrorDsConstraintViolation
	case 0x00002030:
		return ErrorDsNoSuchObject
	case 0x00002031:
		return ErrorDsAliasProblem
	case 0x00002032:
		return ErrorDsInvalidDnSyntax
	case 0x00002033:
		return ErrorDsIsLeaf
	case 0x00002034:
		return ErrorDsAliasDerefProblem
	case 0x00002035:
		return ErrorDsUnwillingToPerform
	case 0x00002036:
		return ErrorDsLoopDetect
	case 0x00002037:
		return ErrorDsNamingViolation
	case 0x00002038:
		return ErrorDsObjectResultsTooLarge
	case 0x00002039:
		return ErrorDsAffectsMultipleDsas
	case 0x0000203A:
		return ErrorDsServerDown
	case 0x0000203B:
		return ErrorDsLocalError
	case 0x0000203C:
		return ErrorDsEncodingError
	case 0x0000203D:
		return ErrorDsDecodingError
	case 0x0000203E:
		return ErrorDsFilterUnknown
	case 0x0000203F:
		return ErrorDsParamError
	case 0x00002040:
		return ErrorDsNotSupported
	case 0x00002041:
		return ErrorDsNoResultsReturned
	case 0x00002042:
		return ErrorDsControlNotFound
	case 0x00002043:
		return ErrorDsClientLoop
	case 0x00002044:
		return ErrorDsReferralLimitExceeded
	case 0x00002045:
		return ErrorDsSortControlMissing
	case 0x00002046:
		return ErrorDsOffsetRangeError
	case 0x0000206D:
		return ErrorDsRootMustBeNc
	case 0x0000206E:
		return ErrorDsAddReplicaInhibited
	case 0x0000206F:
		return ErrorDsAttNotDefInSchema
	case 0x00002070:
		return ErrorDsMaxObjSizeExceeded
	case 0x00002071:
		return ErrorDsObjStringNameExists
	case 0x00002072:
		return ErrorDsNoRdnDefinedInSchema
	case 0x00002073:
		return ErrorDsRdnDoesntMatchSchema
	case 0x00002074:
		return ErrorDsNoRequestedAttsFound
	case 0x00002075:
		return ErrorDsUserBufferToSmall
	case 0x00002076:
		return ErrorDsAttIsNotOnObj
	case 0x00002077:
		return ErrorDsIllegalModOperation
	case 0x00002078:
		return ErrorDsObjTooLarge
	case 0x00002079:
		return ErrorDsBadInstanceType
	case 0x0000207A:
		return ErrorDsMasterdsaRequired
	case 0x0000207B:
		return ErrorDsObjectClassRequired
	case 0x0000207C:
		return ErrorDsMissingRequiredAtt
	case 0x0000207D:
		return ErrorDsAttNotDefForClass
	case 0x0000207E:
		return ErrorDsAttAlreadyExists
	case 0x00002080:
		return ErrorDsCantAddAttValues
	case 0x00002081:
		return ErrorDsSingleValueConstraint
	case 0x00002082:
		return ErrorDsRangeConstraint
	case 0x00002083:
		return ErrorDsAttValAlreadyExists
	case 0x00002084:
		return ErrorDsCantRemMissingAtt
	case 0x00002085:
		return ErrorDsCantRemMissingAttVal
	case 0x00002086:
		return ErrorDsRootCantBeSubref
	case 0x00002087:
		return ErrorDsNoChaining
	case 0x00002088:
		return ErrorDsNoChainedEval
	case 0x00002089:
		return ErrorDsNoParentObject
	case 0x0000208A:
		return ErrorDsParentIsAnAlias
	case 0x0000208B:
		return ErrorDsCantMixMasterAndReps
	case 0x0000208C:
		return ErrorDsChildrenExist
	case 0x0000208D:
		return ErrorDsObjNotFound
	case 0x0000208E:
		return ErrorDsAliasedObjMissing
	case 0x0000208F:
		return ErrorDsBadNameSyntax
	case 0x00002090:
		return ErrorDsAliasPointsToAlias
	case 0x00002091:
		return ErrorDsCantDerefAlias
	case 0x00002092:
		return ErrorDsOutOfScope
	case 0x00002093:
		return ErrorDsObjectBeingRemoved
	case 0x00002094:
		return ErrorDsCantDeleteDsaObj
	case 0x00002095:
		return ErrorDsGenericError
	case 0x00002096:
		return ErrorDsDsaMustBeIntMaster
	case 0x00002097:
		return ErrorDsClassNotDsa
	case 0x00002098:
		return ErrorDsInsuffAccessRights
	case 0x00002099:
		return ErrorDsIllegalSuperior
	case 0x0000209A:
		return ErrorDsAttributeOwnedBySam
	case 0x0000209B:
		return ErrorDsNameTooManyParts
	case 0x0000209C:
		return ErrorDsNameTooLong
	case 0x0000209D:
		return ErrorDsNameValueTooLong
	case 0x0000209E:
		return ErrorDsNameUnparseable
	case 0x0000209F:
		return ErrorDsNameTypeUnknown
	case 0x000020A0:
		return ErrorDsNotAnObject
	case 0x000020A1:
		return ErrorDsSecDescTooShort
	case 0x000020A2:
		return ErrorDsSecDescInvalid
	case 0x000020A3:
		return ErrorDsNoDeletedName
	case 0x000020A4:
		return ErrorDsSubrefMustHaveParent
	case 0x000020A5:
		return ErrorDsNcnameMustBeNc
	case 0x000020A6:
		return ErrorDsCantAddSystemOnly
	case 0x000020A7:
		return ErrorDsClassMustBeConcrete
	case 0x000020A8:
		return ErrorDsInvalidDmd
	case 0x000020A9:
		return ErrorDsObjGuidExists
	case 0x000020AA:
		return ErrorDsNotOnBacklink
	case 0x000020AB:
		return ErrorDsNoCrossrefForNc
	case 0x000020AC:
		return ErrorDsShuttingDown
	case 0x000020AD:
		return ErrorDsUnknownOperation
	case 0x000020AE:
		return ErrorDsInvalidRoleOwner
	case 0x000020AF:
		return ErrorDsCouldntContactFsmo
	case 0x000020B0:
		return ErrorDsCrossNcDnRename
	case 0x000020B1:
		return ErrorDsCantModSystemOnly
	case 0x000020B2:
		return ErrorDsReplicatorOnly
	case 0x000020B3:
		return ErrorDsObjClassNotDefined
	case 0x000020B4:
		return ErrorDsObjClassNotSubclass
	case 0x000020B5:
		return ErrorDsNameReferenceInvalid
	case 0x000020B6:
		return ErrorDsCrossRefExists
	case 0x000020B7:
		return ErrorDsCantDelMasterCrossref
	case 0x000020B8:
		return ErrorDsSubtreeNotifyNotNcHead
	case 0x000020B9:
		return ErrorDsNotifyFilterTooComplex
	case 0x000020BA:
		return ErrorDsDupRdn
	case 0x000020BB:
		return ErrorDsDupOid
	case 0x000020BC:
		return ErrorDsDupMapiId
	case 0x000020BD:
		return ErrorDsDupSchemaIdGuid
	case 0x000020BE:
		return ErrorDsDupLdapDisplayName
	case 0x000020BF:
		return ErrorDsSemanticAttTest
	case 0x000020C0:
		return ErrorDsSyntaxMismatch
	case 0x000020C1:
		return ErrorDsExistsInMustHave
	case 0x000020C2:
		return ErrorDsExistsInMayHave
	case 0x000020C3:
		return ErrorDsNonexistentMayHave
	case 0x000020C4:
		return ErrorDsNonexistentMustHave
	case 0x000020C5:
		return ErrorDsAuxClsTestFail
	case 0x000020C6:
		return ErrorDsNonexistentPossSup
	case 0x000020C7:
		return ErrorDsSubClsTestFail
	case 0x000020C8:
		return ErrorDsBadRdnAttIdSyntax
	case 0x000020C9:
		return ErrorDsExistsInAuxCls
	case 0x000020CA:
		return ErrorDsExistsInSubCls
	case 0x000020CB:
		return ErrorDsExistsInPossSup
	case 0x000020CC:
		return ErrorDsRecalcschemaFailed
	case 0x000020CD:
		return ErrorDsTreeDeleteNotFinished
	case 0x000020CE:
		return ErrorDsCantDelete
	case 0x000020CF:
		return ErrorDsAttSchemaReqId
	case 0x000020D0:
		return ErrorDsBadAttSchemaSyntax
	case 0x000020D1:
		return ErrorDsCantCacheAtt
	case 0x000020D2:
		return ErrorDsCantCacheClass
	case 0x000020D3:
		return ErrorDsCantRemoveAttCache
	case 0x000020D4:
		return ErrorDsCantRemoveClassCache
	case 0x000020D5:
		return ErrorDsCantRetrieveDn
	case 0x000020D6:
		return ErrorDsMissingSupref
	case 0x000020D7:
		return ErrorDsCantRetrieveInstance
	case 0x000020D8:
		return ErrorDsCodeInconsistency
	case 0x000020D9:
		return ErrorDsDatabaseError
	case 0x000020DA:
		return ErrorDsGovernsidMissing
	case 0x000020DB:
		return ErrorDsMissingExpectedAtt
	case 0x000020DC:
		return ErrorDsNcnameMissingCrRef
	case 0x000020DD:
		return ErrorDsSecurityCheckingError
	case 0x000020DE:
		return ErrorDsSchemaNotLoaded
	case 0x000020DF:
		return ErrorDsSchemaAllocFailed
	case 0x000020E0:
		return ErrorDsAttSchemaReqSyntax
	case 0x000020E1:
		return ErrorDsGcverifyError
	case 0x000020E2:
		return ErrorDsDraSchemaMismatch
	case 0x000020E3:
		return ErrorDsCantFindDsaObj
	case 0x000020E4:
		return ErrorDsCantFindExpectedNc
	case 0x000020E5:
		return ErrorDsCantFindNcInCache
	case 0x000020E6:
		return ErrorDsCantRetrieveChild
	case 0x000020E7:
		return ErrorDsSecurityIllegalModify
	case 0x000020E8:
		return ErrorDsCantReplaceHiddenRec
	case 0x000020E9:
		return ErrorDsBadHierarchyFile
	case 0x000020EA:
		return ErrorDsBuildHierarchyTableFailed
	case 0x000020EB:
		return ErrorDsConfigParamMissing
	case 0x000020EC:
		return ErrorDsCountingAbIndicesFailed
	case 0x000020ED:
		return ErrorDsHierarchyTableMallocFailed
	case 0x000020EE:
		return ErrorDsInternalFailure
	case 0x000020EF:
		return ErrorDsUnknownError
	case 0x000020F0:
		return ErrorDsRootRequiresClassTop
	case 0x000020F1:
		return ErrorDsRefusingFsmoRoles
	case 0x000020F2:
		return ErrorDsMissingFsmoSettings
	case 0x000020F3:
		return ErrorDsUnableToSurrenderRoles
	case 0x000020F4:
		return ErrorDsDraGeneric
	case 0x000020F5:
		return ErrorDsDraInvalidParameter
	case 0x000020F6:
		return ErrorDsDraBusy
	case 0x000020F7:
		return ErrorDsDraBadDn
	case 0x000020F8:
		return ErrorDsDraBadNc
	case 0x000020F9:
		return ErrorDsDraDnExists
	case 0x000020FA:
		return ErrorDsDraInternalError
	case 0x000020FB:
		return ErrorDsDraInconsistentDit
	case 0x000020FC:
		return ErrorDsDraConnectionFailed
	case 0x000020FD:
		return ErrorDsDraBadInstanceType
	case 0x000020FE:
		return ErrorDsDraOutOfMem
	case 0x000020FF:
		return ErrorDsDraMailProblem
	case 0x00002100:
		return ErrorDsDraRefAlreadyExists
	case 0x00002101:
		return ErrorDsDraRefNotFound
	case 0x00002102:
		return ErrorDsDraObjIsRepSource
	case 0x00002103:
		return ErrorDsDraDbError
	case 0x00002104:
		return ErrorDsDraNoReplica
	case 0x00002105:
		return ErrorDsDraAccessDenied
	case 0x00002106:
		return ErrorDsDraNotSupported
	case 0x00002107:
		return ErrorDsDraRpcCancelled
	case 0x00002108:
		return ErrorDsDraSourceDisabled
	case 0x00002109:
		return ErrorDsDraSinkDisabled
	case 0x0000210A:
		return ErrorDsDraNameCollision
	case 0x0000210B:
		return ErrorDsDraSourceReinstalled
	case 0x0000210C:
		return ErrorDsDraMissingParent
	case 0x0000210D:
		return ErrorDsDraPreempted
	case 0x0000210E:
		return ErrorDsDraAbandonSync
	case 0x0000210F:
		return ErrorDsDraShutdown
	case 0x00002110:
		return ErrorDsDraIncompatiblePartialSet
	case 0x00002111:
		return ErrorDsDraSourceIsPartialReplica
	case 0x00002112:
		return ErrorDsDraExtnConnectionFailed
	case 0x00002113:
		return ErrorDsInstallSchemaMismatch
	case 0x00002114:
		return ErrorDsDupLinkId
	case 0x00002115:
		return ErrorDsNameErrorResolving
	case 0x00002116:
		return ErrorDsNameErrorNotFound
	case 0x00002117:
		return ErrorDsNameErrorNotUnique
	case 0x00002118:
		return ErrorDsNameErrorNoMapping
	case 0x00002119:
		return ErrorDsNameErrorDomainOnly
	case 0x0000211A:
		return ErrorDsNameErrorNoSyntacticalMapping
	case 0x0000211B:
		return ErrorDsConstructedAttMod
	case 0x0000211C:
		return ErrorDsWrongOmObjClass
	case 0x0000211D:
		return ErrorDsDraReplPending
	case 0x0000211E:
		return ErrorDsDsRequired
	case 0x0000211F:
		return ErrorDsInvalidLdapDisplayName
	case 0x00002120:
		return ErrorDsNonBaseSearch
	case 0x00002121:
		return ErrorDsCantRetrieveAtts
	case 0x00002122:
		return ErrorDsBacklinkWithoutLink
	case 0x00002123:
		return ErrorDsEpochMismatch
	case 0x00002124:
		return ErrorDsSrcNameMismatch
	case 0x00002125:
		return ErrorDsSrcAndDstNcIdentical
	case 0x00002126:
		return ErrorDsDstNcMismatch
	case 0x00002127:
		return ErrorDsNotAuthoritiveForDstNc
	case 0x00002128:
		return ErrorDsSrcGuidMismatch
	case 0x00002129:
		return ErrorDsCantMoveDeletedObject
	case 0x0000212A:
		return ErrorDsPdcOperationInProgress
	case 0x0000212B:
		return ErrorDsCrossDomainCleanupReqd
	case 0x0000212C:
		return ErrorDsIllegalXdomMoveOperation
	case 0x0000212D:
		return ErrorDsCantWithAcctGroupMembershps
	case 0x0000212E:
		return ErrorDsNcMustHaveNcParent
	case 0x0000212F:
		return ErrorDsCrImpossibleToValidate
	case 0x00002130:
		return ErrorDsDstDomainNotNative
	case 0x00002131:
		return ErrorDsMissingInfrastructureContainer
	case 0x00002132:
		return ErrorDsCantMoveAccountGroup
	case 0x00002133:
		return ErrorDsCantMoveResourceGroup
	case 0x00002134:
		return ErrorDsInvalidSearchFlag
	case 0x00002135:
		return ErrorDsNoTreeDeleteAboveNc
	case 0x00002136:
		return ErrorDsCouldntLockTreeForDelete
	case 0x00002137:
		return ErrorDsCouldntIdentifyObjectsForTreeDelete
	case 0x00002138:
		return ErrorDsSamInitFailure
	case 0x00002139:
		return ErrorDsSensitiveGroupViolation
	case 0x0000213A:
		return ErrorDsCantModPrimarygroupid
	case 0x0000213B:
		return ErrorDsIllegalBaseSchemaMod
	case 0x0000213C:
		return ErrorDsNonsafeSchemaChange
	case 0x0000213D:
		return ErrorDsSchemaUpdateDisallowed
	case 0x0000213E:
		return ErrorDsCantCreateUnderSchema
	case 0x0000213F:
		return ErrorDsInstallNoSrcSchVersion
	case 0x00002140:
		return ErrorDsInstallNoSchVersionInInifile
	case 0x00002141:
		return ErrorDsInvalidGroupType
	case 0x00002142:
		return ErrorDsNoNestGlobalgroupInMixeddomain
	case 0x00002143:
		return ErrorDsNoNestLocalgroupInMixeddomain
	case 0x00002144:
		return ErrorDsGlobalCantHaveLocalMember
	case 0x00002145:
		return ErrorDsGlobalCantHaveUniversalMember
	case 0x00002146:
		return ErrorDsUniversalCantHaveLocalMember
	case 0x00002147:
		return ErrorDsGlobalCantHaveCrossdomainMember
	case 0x00002148:
		return ErrorDsLocalCantHaveCrossdomainLocalMember
	case 0x00002149:
		return ErrorDsHavePrimaryMembers
	case 0x0000214A:
		return ErrorDsStringSdConversionFailed
	case 0x0000214B:
		return ErrorDsNamingMasterGc
	case 0x0000214C:
		return ErrorDsDnsLookupFailure
	case 0x0000214D:
		return ErrorDsCouldntUpdateSpns
	case 0x0000214E:
		return ErrorDsCantRetrieveSd
	case 0x0000214F:
		return ErrorDsKeyNotUnique
	case 0x00002150:
		return ErrorDsWrongLinkedAttSyntax
	case 0x00002151:
		return ErrorDsSamNeedBootkeyPassword
	case 0x00002152:
		return ErrorDsSamNeedBootkeyFloppy
	case 0x00002153:
		return ErrorDsCantStart
	case 0x00002154:
		return ErrorDsInitFailure
	case 0x00002155:
		return ErrorDsNoPktPrivacyOnConnection
	case 0x00002156:
		return ErrorDsSourceDomainInForest
	case 0x00002157:
		return ErrorDsDestinationDomainNotInForest
	case 0x00002158:
		return ErrorDsDestinationAuditingNotEnabled
	case 0x00002159:
		return ErrorDsCantFindDcForSrcDomain
	case 0x0000215A:
		return ErrorDsSrcObjNotGroupOrUser
	case 0x0000215B:
		return ErrorDsSrcSidExistsInForest
	case 0x0000215C:
		return ErrorDsSrcAndDstObjectClassMismatch
	case 0x0000215D:
		return ErrorSamInitFailure
	case 0x0000215E:
		return ErrorDsDraSchemaInfoShip
	case 0x0000215F:
		return ErrorDsDraSchemaConflict
	case 0x00002160:
		return ErrorDsDraEarlierSchemaConflict
	case 0x00002161:
		return ErrorDsDraObjNcMismatch
	case 0x00002162:
		return ErrorDsNcStillHasDsas
	case 0x00002163:
		return ErrorDsGcRequired
	case 0x00002164:
		return ErrorDsLocalMemberOfLocalOnly
	case 0x00002165:
		return ErrorDsNoFpoInUniversalGroups
	case 0x00002166:
		return ErrorDsCantAddToGc
	case 0x00002167:
		return ErrorDsNoCheckpointWithPdc
	case 0x00002168:
		return ErrorDsSourceAuditingNotEnabled
	case 0x00002169:
		return ErrorDsCantCreateInNondomainNc
	case 0x0000216A:
		return ErrorDsInvalidNameForSpn
	case 0x0000216B:
		return ErrorDsFilterUsesContructedAttrs
	case 0x0000216C:
		return ErrorDsUnicodepwdNotInQuotes
	case 0x0000216D:
		return ErrorDsMachineAccountQuotaExceeded
	case 0x0000216E:
		return ErrorDsMustBeRunOnDstDc
	case 0x0000216F:
		return ErrorDsSrcDcMustBeSp4OrGreater
	case 0x00002170:
		return ErrorDsCantTreeDeleteCriticalObj
	case 0x00002171:
		return ErrorDsInitFailureConsole
	case 0x00002172:
		return ErrorDsSamInitFailureConsole
	case 0x00002173:
		return ErrorDsForestVersionTooHigh
	case 0x00002174:
		return ErrorDsDomainVersionTooHigh
	case 0x00002175:
		return ErrorDsForestVersionTooLow
	case 0x00002176:
		return ErrorDsDomainVersionTooLow
	case 0x00002177:
		return ErrorDsIncompatibleVersion
	case 0x00002178:
		return ErrorDsLowDsaVersion
	case 0x00002179:
		return ErrorDsNoBehaviorVersionInMixeddomain
	case 0x0000217A:
		return ErrorDsNotSupportedSortOrder
	case 0x0000217B:
		return ErrorDsNameNotUnique
	case 0x0000217C:
		return ErrorDsMachineAccountCreatedPrent4
	case 0x0000217D:
		return ErrorDsOutOfVersionStore
	case 0x0000217E:
		return ErrorDsIncompatibleControlsUsed
	case 0x0000217F:
		return ErrorDsNoRefDomain
	case 0x00002180:
		return ErrorDsReservedLinkId
	case 0x00002181:
		return ErrorDsLinkIdNotAvailable
	case 0x00002182:
		return ErrorDsAgCantHaveUniversalMember
	case 0x00002183:
		return ErrorDsModifydnDisallowedByInstanceType
	case 0x00002184:
		return ErrorDsNoObjectMoveInSchemaNc
	case 0x00002185:
		return ErrorDsModifydnDisallowedByFlag
	case 0x00002186:
		return ErrorDsModifydnWrongGrandparent
	case 0x00002187:
		return ErrorDsNameErrorTrustReferral
	case 0x00002188:
		return ErrorNotSupportedOnStandardServer
	case 0x00002189:
		return ErrorDsCantAccessRemotePartOfAd
	case 0x0000218A:
		return ErrorDsCrImpossibleToValidateV2
	case 0x0000218B:
		return ErrorDsThreadLimitExceeded
	case 0x0000218C:
		return ErrorDsNotClosest
	case 0x0000218D:
		return ErrorDsCantDeriveSpnWithoutServerRef
	case 0x0000218E:
		return ErrorDsSingleUserModeFailed
	case 0x0000218F:
		return ErrorDsNtdscriptSyntaxError
	case 0x00002190:
		return ErrorDsNtdscriptProcessError
	case 0x00002191:
		return ErrorDsDifferentReplEpochs
	case 0x00002192:
		return ErrorDsDrsExtensionsChanged
	case 0x00002193:
		return ErrorDsReplicaSetChangeNotAllowedOnDisabledCr
	case 0x00002194:
		return ErrorDsNoMsdsIntid
	case 0x00002195:
		return ErrorDsDupMsdsIntid
	case 0x00002196:
		return ErrorDsExistsInRdnattid
	case 0x00002197:
		return ErrorDsAuthorizationFailed
	case 0x00002198:
		return ErrorDsInvalidScript
	case 0x00002199:
		return ErrorDsRemoteCrossrefOpFailed
	case 0x0000219A:
		return ErrorDsCrossRefBusy
	case 0x0000219B:
		return ErrorDsCantDeriveSpnForDeletedDomain
	case 0x0000219C:
		return ErrorDsCantDemoteWithWriteableNc
	case 0x0000219D:
		return ErrorDsDuplicateIdFound
	case 0x0000219E:
		return ErrorDsInsufficientAttrToCreateObject
	case 0x0000219F:
		return ErrorDsGroupConversionError
	case 0x000021A0:
		return ErrorDsCantMoveAppBasicGroup
	case 0x000021A1:
		return ErrorDsCantMoveAppQueryGroup
	case 0x000021A2:
		return ErrorDsRoleNotVerified
	case 0x000021A3:
		return ErrorDsWkoContainerCannotBeSpecial
	case 0x000021A4:
		return ErrorDsDomainRenameInProgress
	case 0x000021A5:
		return ErrorDsExistingAdChildNc
	case 0x000021A6:
		return ErrorDsReplLifetimeExceeded
	case 0x000021A7:
		return ErrorDsDisallowedInSystemContainer
	case 0x000021A8:
		return ErrorDsLdapSendQueueFull
	case 0x000021A9:
		return ErrorDsDraOutScheduleWindow
	case 0x000021AA:
		return ErrorDsPolicyNotKnown
	case 0x000021AB:
		return ErrorNoSiteSettingsObject
	case 0x000021AC:
		return ErrorNoSecrets
	case 0x000021AD:
		return ErrorNoWritableDcFound
	case 0x000021AE:
		return ErrorDsNoServerObject
	case 0x000021AF:
		return ErrorDsNoNtdsaObject
	case 0x000021B0:
		return ErrorDsNonAsqSearch
	case 0x000021B1:
		return ErrorDsAuditFailure
	case 0x000021B2:
		return ErrorDsInvalidSearchFlagSubtree
	case 0x000021B3:
		return ErrorDsInvalidSearchFlagTuple
	case 0x000021BF:
		return ErrorDsDraRecycledTarget
	case 0x000021C2:
		return ErrorDsHighDsaVersion
	case 0x000021C7:
		return ErrorDsSpnValueNotUniqueInForest
	case 0x000021C8:
		return ErrorDsUpnValueNotUniqueInForest
	case 0x00002329:
		return DnsErrorRcodeFormatError
	case 0x0000232A:
		return DnsErrorRcodeServerFailure
	case 0x0000232B:
		return DnsErrorRcodeNameError
	case 0x0000232C:
		return DnsErrorRcodeNotImplemented
	case 0x0000232D:
		return DnsErrorRcodeRefused
	case 0x0000232E:
		return DnsErrorRcodeYxdomain
	case 0x0000232F:
		return DnsErrorRcodeYxrrset
	case 0x00002330:
		return DnsErrorRcodeNxrrset
	case 0x00002331:
		return DnsErrorRcodeNotauth
	case 0x00002332:
		return DnsErrorRcodeNotzone
	case 0x00002338:
		return DnsErrorRcodeBadsig
	case 0x00002339:
		return DnsErrorRcodeBadkey
	case 0x0000233A:
		return DnsErrorRcodeBadtime
	case 0x0000251D:
		return DnsInfoNoRecords
	case 0x0000251E:
		return DnsErrorBadPacket
	case 0x0000251F:
		return DnsErrorNoPacket
	case 0x00002520:
		return DnsErrorRcode
	case 0x00002521:
		return DnsErrorUnsecurePacket
	case 0x0000254F:
		return DnsErrorInvalidType
	case 0x00002550:
		return DnsErrorInvalidIpAddress
	case 0x00002551:
		return DnsErrorInvalidProperty
	case 0x00002552:
		return DnsErrorTryAgainLater
	case 0x00002553:
		return DnsErrorNotUnique
	case 0x00002554:
		return DnsErrorNonRfcName
	case 0x00002555:
		return DnsStatusFqdn
	case 0x00002556:
		return DnsStatusDottedName
	case 0x00002557:
		return DnsStatusSinglePartName
	case 0x00002558:
		return DnsErrorInvalidNameChar
	case 0x00002559:
		return DnsErrorNumericName
	case 0x0000255A:
		return DnsErrorNotAllowedOnRootServer
	case 0x0000255B:
		return DnsErrorNotAllowedUnderDelegation
	case 0x0000255C:
		return DnsErrorCannotFindRootHints
	case 0x0000255D:
		return DnsErrorInconsistentRootHints
	case 0x0000255E:
		return DnsErrorDwordValueTooSmall
	case 0x0000255F:
		return DnsErrorDwordValueTooLarge
	case 0x00002560:
		return DnsErrorBackgroundLoading
	case 0x00002561:
		return DnsErrorNotAllowedOnRodc
	case 0x00002581:
		return DnsErrorZoneDoesNotExist
	case 0x00002582:
		return DnsErrorNoZoneInfo
	case 0x00002583:
		return DnsErrorInvalidZoneOperation
	case 0x00002584:
		return DnsErrorZoneConfigurationError
	case 0x00002585:
		return DnsErrorZoneHasNoSoaRecord
	case 0x00002586:
		return DnsErrorZoneHasNoNsRecords
	case 0x00002587:
		return DnsErrorZoneLocked
	case 0x00002588:
		return DnsErrorZoneCreationFailed
	case 0x00002589:
		return DnsErrorZoneAlreadyExists
	case 0x0000258A:
		return DnsErrorAutozoneAlreadyExists
	case 0x0000258B:
		return DnsErrorInvalidZoneType
	case 0x0000258C:
		return DnsErrorSecondaryRequiresMasterIp
	case 0x0000258D:
		return DnsErrorZoneNotSecondary
	case 0x0000258E:
		return DnsErrorNeedSecondaryAddresses
	case 0x0000258F:
		return DnsErrorWinsInitFailed
	case 0x00002590:
		return DnsErrorNeedWinsServers
	case 0x00002591:
		return DnsErrorNbstatInitFailed
	case 0x00002592:
		return DnsErrorSoaDeleteInvalid
	case 0x00002593:
		return DnsErrorForwarderAlreadyExists
	case 0x00002594:
		return DnsErrorZoneRequiresMasterIp
	case 0x00002595:
		return DnsErrorZoneIsShutdown
	case 0x000025B3:
		return DnsErrorPrimaryRequiresDatafile
	case 0x000025B4:
		return DnsErrorInvalidDatafileName
	case 0x000025B5:
		return DnsErrorDatafileOpenFailure
	case 0x000025B6:
		return DnsErrorFileWritebackFailed
	case 0x000025B7:
		return DnsErrorDatafileParsing
	case 0x000025E5:
		return DnsErrorRecordDoesNotExist
	case 0x000025E6:
		return DnsErrorRecordFormat
	case 0x000025E7:
		return DnsErrorNodeCreationFailed
	case 0x000025E8:
		return DnsErrorUnknownRecordType
	case 0x000025E9:
		return DnsErrorRecordTimedOut
	case 0x000025EA:
		return DnsErrorNameNotInZone
	case 0x000025EB:
		return DnsErrorCnameLoop
	case 0x000025EC:
		return DnsErrorNodeIsCname
	case 0x000025ED:
		return DnsErrorCnameCollision
	case 0x000025EE:
		return DnsErrorRecordOnlyAtZoneRoot
	case 0x000025EF:
		return DnsErrorRecordAlreadyExists
	case 0x000025F0:
		return DnsErrorSecondaryData
	case 0x000025F1:
		return DnsErrorNoCreateCacheData
	case 0x000025F2:
		return DnsErrorNameDoesNotExist
	case 0x000025F3:
		return DnsWarningPtrCreateFailed
	case 0x000025F4:
		return DnsWarningDomainUndeleted
	case 0x000025F5:
		return DnsErrorDsUnavailable
	case 0x000025F6:
		return DnsErrorDsZoneAlreadyExists
	case 0x000025F7:
		return DnsErrorNoBootfileIfDsZone
	case 0x00002617:
		return DnsInfoAxfrComplete
	case 0x00002618:
		return DnsErrorAxfr
	case 0x00002619:
		return DnsInfoAddedLocalWins
	case 0x00002649:
		return DnsStatusContinueNeeded
	case 0x0000267B:
		return DnsErrorNoTcpip
	case 0x0000267C:
		return DnsErrorNoDnsServers
	case 0x000026AD:
		return DnsErrorDpDoesNotExist
	case 0x000026AE:
		return DnsErrorDpAlreadyExists
	case 0x000026AF:
		return DnsErrorDpNotEnlisted
	case 0x000026B0:
		return DnsErrorDpAlreadyEnlisted
	case 0x000026B1:
		return DnsErrorDpNotAvailable
	case 0x000026B2:
		return DnsErrorDpFsmoError
	case 0x00002714:
		return Wsaeintr
	case 0x00002719:
		return Wsaebadf
	case 0x0000271D:
		return Wsaeacces
	case 0x0000271E:
		return Wsaefault
	case 0x00002726:
		return Wsaeinval
	case 0x00002728:
		return Wsaemfile
	case 0x00002733:
		return Wsaewouldblock
	case 0x00002734:
		return Wsaeinprogress
	case 0x00002735:
		return Wsaealready
	case 0x00002736:
		return Wsaenotsock
	case 0x00002737:
		return Wsaedestaddrreq
	case 0x00002738:
		return Wsaemsgsize
	case 0x00002739:
		return Wsaeprototype
	case 0x0000273A:
		return Wsaenoprotoopt
	case 0x0000273B:
		return Wsaeprotonosupport
	case 0x0000273C:
		return Wsaesocktnosupport
	case 0x0000273D:
		return Wsaeopnotsupp
	case 0x0000273E:
		return Wsaepfnosupport
	case 0x0000273F:
		return Wsaeafnosupport
	case 0x00002740:
		return Wsaeaddrinuse
	case 0x00002741:
		return Wsaeaddrnotavail
	case 0x00002742:
		return Wsaenetdown
	case 0x00002743:
		return Wsaenetunreach
	case 0x00002744:
		return Wsaenetreset
	case 0x00002745:
		return Wsaeconnaborted
	case 0x00002746:
		return Wsaeconnreset
	case 0x00002747:
		return Wsaenobufs
	case 0x00002748:
		return Wsaeisconn
	case 0x00002749:
		return Wsaenotconn
	case 0x0000274A:
		return Wsaeshutdown
	case 0x0000274B:
		return Wsaetoomanyrefs
	case 0x0000274C:
		return Wsaetimedout
	case 0x0000274D:
		return Wsaeconnrefused
	case 0x0000274E:
		return Wsaeloop
	case 0x0000274F:
		return Wsaenametoolong
	case 0x00002750:
		return Wsaehostdown
	case 0x00002751:
		return Wsaehostunreach
	case 0x00002752:
		return Wsaenotempty
	case 0x00002753:
		return Wsaeproclim
	case 0x00002754:
		return Wsaeusers
	case 0x00002755:
		return Wsaedquot
	case 0x00002756:
		return Wsaestale
	case 0x00002757:
		return Wsaeremote
	case 0x0000276B:
		return Wsasysnotready
	case 0x0000276C:
		return Wsavernotsupported
	case 0x0000276D:
		return Wsanotinitialised
	case 0x00002775:
		return Wsaediscon
	case 0x00002776:
		return Wsaenomore
	case 0x00002777:
		return Wsaecancelled
	case 0x00002778:
		return Wsaeinvalidproctable
	case 0x00002779:
		return Wsaeinvalidprovider
	case 0x0000277A:
		return Wsaeproviderfailedinit
	case 0x0000277B:
		return Wsasyscallfailure
	case 0x0000277C:
		return WsaserviceNotFound
	case 0x0000277D:
		return WsatypeNotFound
	case 0x0000277E:
		return WsaENoMore
	case 0x0000277F:
		return WsaECancelled
	case 0x00002780:
		return Wsaerefused
	case 0x00002AF9:
		return WsahostNotFound
	case 0x00002AFA:
		return WsatryAgain
	case 0x00002AFB:
		return WsanoRecovery
	case 0x00002AFC:
		return WsanoData
	case 0x00002AFD:
		return WsaQosReceivers
	case 0x00002AFE:
		return WsaQosSenders
	case 0x00002AFF:
		return WsaQosNoSenders
	case 0x00002B00:
		return WsaQosNoReceivers
	case 0x00002B01:
		return WsaQosRequestConfirmed
	case 0x00002B02:
		return WsaQosAdmissionFailure
	case 0x00002B03:
		return WsaQosPolicyFailure
	case 0x00002B04:
		return WsaQosBadStyle
	case 0x00002B05:
		return WsaQosBadObject
	case 0x00002B06:
		return WsaQosTrafficCtrlError
	case 0x00002B07:
		return WsaQosGenericError
	case 0x00002B08:
		return WsaQosEservicetype
	case 0x00002B09:
		return WsaQosEflowspec
	case 0x00002B0A:
		return WsaQosEprovspecbuf
	case 0x00002B0B:
		return WsaQosEfilterstyle
	case 0x00002B0C:
		return WsaQosEfiltertype
	case 0x00002B0D:
		return WsaQosEfiltercount
	case 0x00002B0E:
		return WsaQosEobjlength
	case 0x00002B0F:
		return WsaQosEflowcount
	case 0x00002B10:
		return WsaQosEunkownpsobj
	case 0x00002B11:
		return WsaQosEpolicyobj
	case 0x00002B12:
		return WsaQosEflowdesc
	case 0x00002B13:
		return WsaQosEpsflowspec
	case 0x00002B14:
		return WsaQosEpsfilterspec
	case 0x00002B15:
		return WsaQosEsdmodeobj
	case 0x00002B16:
		return WsaQosEshaperateobj
	case 0x00002B17:
		return WsaQosReservedPetype
	case 0x000032C8:
		return ErrorIpsecQmPolicyExists
	case 0x000032C9:
		return ErrorIpsecQmPolicyNotFound
	case 0x000032CA:
		return ErrorIpsecQmPolicyInUse
	case 0x000032CB:
		return ErrorIpsecMmPolicyExists
	case 0x000032CC:
		return ErrorIpsecMmPolicyNotFound
	case 0x000032CD:
		return ErrorIpsecMmPolicyInUse
	case 0x000032CE:
		return ErrorIpsecMmFilterExists
	case 0x000032CF:
		return ErrorIpsecMmFilterNotFound
	case 0x000032D0:
		return ErrorIpsecTransportFilterExists
	case 0x000032D1:
		return ErrorIpsecTransportFilterNotFound
	case 0x000032D2:
		return ErrorIpsecMmAuthExists
	case 0x000032D3:
		return ErrorIpsecMmAuthNotFound
	case 0x000032D4:
		return ErrorIpsecMmAuthInUse
	case 0x000032D5:
		return ErrorIpsecDefaultMmPolicyNotFound
	case 0x000032D6:
		return ErrorIpsecDefaultMmAuthNotFound
	case 0x000032D7:
		return ErrorIpsecDefaultQmPolicyNotFound
	case 0x000032D8:
		return ErrorIpsecTunnelFilterExists
	case 0x000032D9:
		return ErrorIpsecTunnelFilterNotFound
	case 0x000032DA:
		return ErrorIpsecMmFilterPendingDeletion
	case 0x000032DB:
		return ErrorIpsecTransportFilterEndingDeletion
	case 0x000032DC:
		return ErrorIpsecTunnelFilterPendingDeletion
	case 0x000032DD:
		return ErrorIpsecMmPolicyPendingEletion
	case 0x000032DE:
		return ErrorIpsecMmAuthPendingDeletion
	case 0x000032DF:
		return ErrorIpsecQmPolicyPendingDeletion
	case 0x000032E0:
		return WarningIpsecMmPolicyPruned
	case 0x000032E1:
		return WarningIpsecQmPolicyPruned
	case 0x000035E8:
		return ErrorIpsecIkeNegStatusBegin
	case 0x000035E9:
		return ErrorIpsecIkeAuthFail
	case 0x000035EA:
		return ErrorIpsecIkeAttribFail
	case 0x000035EB:
		return ErrorIpsecIkeNegotiationPending
	case 0x000035EC:
		return ErrorIpsecIkeGeneralProcessingError
	case 0x000035ED:
		return ErrorIpsecIkeTimedOut
	case 0x000035EE:
		return ErrorIpsecIkeNoCert
	case 0x000035EF:
		return ErrorIpsecIkeSaDeleted
	case 0x000035F0:
		return ErrorIpsecIkeSaReaped
	case 0x000035F1:
		return ErrorIpsecIkeMmAcquireDrop
	case 0x000035F2:
		return ErrorIpsecIkeQmAcquireDrop
	case 0x000035F3:
		return ErrorIpsecIkeQueueDropMm
	case 0x000035F4:
		return ErrorIpsecIkeQueueDropNoMm
	case 0x000035F5:
		return ErrorIpsecIkeDropNoResponse
	case 0x000035F6:
		return ErrorIpsecIkeMmDelayDrop
	case 0x000035F7:
		return ErrorIpsecIkeQmDelayDrop
	case 0x000035F8:
		return ErrorIpsecIkeError
	case 0x000035F9:
		return ErrorIpsecIkeCrlFailed
	case 0x000035FA:
		return ErrorIpsecIkeInvalidKeyUsage
	case 0x000035FB:
		return ErrorIpsecIkeInvalidCertType
	case 0x000035FC:
		return ErrorIpsecIkeNoPrivateKey
	case 0x000035FE:
		return ErrorIpsecIkeDhFail
	case 0x00003600:
		return ErrorIpsecIkeInvalidHeader
	case 0x00003601:
		return ErrorIpsecIkeNoPolicy
	case 0x00003602:
		return ErrorIpsecIkeInvalidSignature
	case 0x00003603:
		return ErrorIpsecIkeKerberosError
	case 0x00003604:
		return ErrorIpsecIkeNoPublicKey
	case 0x00003605:
		return ErrorIpsecIkeProcessErr
	case 0x00003606:
		return ErrorIpsecIkeProcessErrSa
	case 0x00003607:
		return ErrorIpsecIkeProcessErrProp
	case 0x00003608:
		return ErrorIpsecIkeProcessErrTrans
	case 0x00003609:
		return ErrorIpsecIkeProcessErrKe
	case 0x0000360A:
		return ErrorIpsecIkeProcessErrId
	case 0x0000360B:
		return ErrorIpsecIkeProcessErrCert
	case 0x0000360C:
		return ErrorIpsecIkeProcessErrCertReq
	case 0x0000360D:
		return ErrorIpsecIkeProcessErrHash
	case 0x0000360E:
		return ErrorIpsecIkeProcessErrSig
	case 0x0000360F:
		return ErrorIpsecIkeProcessErrNonce
	case 0x00003610:
		return ErrorIpsecIkeProcessErrNotify
	case 0x00003611:
		return ErrorIpsecIkeProcessErrDelete
	case 0x00003612:
		return ErrorIpsecIkeProcessErrVendor
	case 0x00003613:
		return ErrorIpsecIkeInvalidPayload
	case 0x00003614:
		return ErrorIpsecIkeLoadSoftSa
	case 0x00003615:
		return ErrorIpsecIkeSoftSaTornDown
	case 0x00003616:
		return ErrorIpsecIkeInvalidCookie
	case 0x00003617:
		return ErrorIpsecIkeNoPeerCert
	case 0x00003618:
		return ErrorIpsecIkePeerCrlFailed
	case 0x00003619:
		return ErrorIpsecIkePolicyChange
	case 0x0000361A:
		return ErrorIpsecIkeNoMmPolicy
	case 0x0000361B:
		return ErrorIpsecIkeNotcbpriv
	case 0x0000361C:
		return ErrorIpsecIkeSecloadfail
	case 0x0000361D:
		return ErrorIpsecIkeFailsspinit
	case 0x0000361E:
		return ErrorIpsecIkeFailqueryssp
	case 0x0000361F:
		return ErrorIpsecIkeSrvacqfail
	case 0x00003620:
		return ErrorIpsecIkeSrvquerycred
	case 0x00003621:
		return ErrorIpsecIkeGetspifail
	case 0x00003622:
		return ErrorIpsecIkeInvalidFilter
	case 0x00003623:
		return ErrorIpsecIkeOutOfMemory
	case 0x00003624:
		return ErrorIpsecIkeAddUpdateKeyFailed
	case 0x00003625:
		return ErrorIpsecIkeInvalidPolicy
	case 0x00003626:
		return ErrorIpsecIkeUnknownDoi
	case 0x00003627:
		return ErrorIpsecIkeInvalidSituation
	case 0x00003628:
		return ErrorIpsecIkeDhFailure
	case 0x00003629:
		return ErrorIpsecIkeInvalidGroup
	case 0x0000362A:
		return ErrorIpsecIkeEncrypt
	case 0x0000362B:
		return ErrorIpsecIkeDecrypt
	case 0x0000362C:
		return ErrorIpsecIkePolicyMatch
	case 0x0000362D:
		return ErrorIpsecIkeUnsupportedId
	case 0x0000362E:
		return ErrorIpsecIkeInvalidHash
	case 0x0000362F:
		return ErrorIpsecIkeInvalidHashAlg
	case 0x00003630:
		return ErrorIpsecIkeInvalidHashSize
	case 0x00003631:
		return ErrorIpsecIkeInvalidEncryptAlg
	case 0x00003632:
		return ErrorIpsecIkeInvalidAuthAlg
	case 0x00003633:
		return ErrorIpsecIkeInvalidSig
	case 0x00003634:
		return ErrorIpsecIkeLoadFailed
	case 0x00003635:
		return ErrorIpsecIkeRpcDelete
	case 0x00003636:
		return ErrorIpsecIkeBenignReinit
	case 0x00003637:
		return ErrorIpsecIkeInvalidResponderLifetimeNotify
	case 0x00003639:
		return ErrorIpsecIkeInvalidCertKeylen
	case 0x0000363A:
		return ErrorIpsecIkeMmLimit
	case 0x0000363B:
		return ErrorIpsecIkeNegotiationDisabled
	case 0x0000363C:
		return ErrorIpsecIkeQmLimit
	case 0x0000363D:
		return ErrorIpsecIkeMmExpired
	case 0x0000363E:
		return ErrorIpsecIkePeerMmAssumedInvalid
	case 0x0000363F:
		return ErrorIpsecIkeCertChainPolicyMismatch
	case 0x00003640:
		return ErrorIpsecIkeUnexpectedMessageId
	case 0x00003641:
		return ErrorIpsecIkeInvalidUmatts
	case 0x00003642:
		return ErrorIpsecIkeDosCookieSent
	case 0x00003643:
		return ErrorIpsecIkeShuttingDown
	case 0x00003644:
		return ErrorIpsecIkeCgaAuthFailed
	case 0x00003645:
		return ErrorIpsecIkeProcessErrNatoa
	case 0x00003646:
		return ErrorIpsecIkeInvalidMmForQm
	case 0x00003647:
		return ErrorIpsecIkeQmExpired
	case 0x00003648:
		return ErrorIpsecIkeTooManyFilters
	case 0x00003649:
		return ErrorIpsecIkeNegStatusEnd
	case 0x000036B0:
		return ErrorSxsSectionNotFound
	case 0x000036B1:
		return ErrorSxsCantGenActctx
	case 0x000036B2:
		return ErrorSxsInvalidActctxdataFormat
	case 0x000036B3:
		return ErrorSxsAssemblyNotFound
	case 0x000036B4:
		return ErrorSxsManifestFormatError
	case 0x000036B5:
		return ErrorSxsManifestParseError
	case 0x000036B6:
		return ErrorSxsActivationContextDisabled
	case 0x000036B7:
		return ErrorSxsKeyNotFound
	case 0x000036B8:
		return ErrorSxsVersionConflict
	case 0x000036B9:
		return ErrorSxsWrongSectionType
	case 0x000036BA:
		return ErrorSxsThreadQueriesDisabled
	case 0x000036BB:
		return ErrorSxsProcessDefaultAlreadySet
	case 0x000036BC:
		return ErrorSxsUnknownEncodingGroup
	case 0x000036BD:
		return ErrorSxsUnknownEncoding
	case 0x000036BE:
		return ErrorSxsInvalidXmlNamespaceUri
	case 0x000036BF:
		return ErrorSxsRootManifestDependencyOtInstalled
	case 0x000036C0:
		return ErrorSxsLeafManifestDependencyNotInstalled
	case 0x000036C1:
		return ErrorSxsInvalidAssemblyIdentityAttribute
	case 0x000036C2:
		return ErrorSxsManifestMissingRequiredDefaultNamespace
	case 0x000036C3:
		return ErrorSxsManifestInvalidRequiredDefaultNamespace
	case 0x000036C4:
		return ErrorSxsPrivateManifestCrossPathWithReparsePoint
	case 0x000036C5:
		return ErrorSxsDuplicateDllName
	case 0x000036C6:
		return ErrorSxsDuplicateWindowclassName
	case 0x000036C7:
		return ErrorSxsDuplicateClsid
	case 0x000036C8:
		return ErrorSxsDuplicateIid
	case 0x000036C9:
		return ErrorSxsDuplicateTlbid
	case 0x000036CA:
		return ErrorSxsDuplicateProgid
	case 0x000036CB:
		return ErrorSxsDuplicateAssemblyName
	case 0x000036CC:
		return ErrorSxsFileHashMismatch
	case 0x000036CD:
		return ErrorSxsPolicyParseError
	case 0x000036CE:
		return ErrorSxsXmlEMissingquote
	case 0x000036CF:
		return ErrorSxsXmlECommentsyntax
	case 0x000036D0:
		return ErrorSxsXmlEBadstartnamechar
	case 0x000036D1:
		return ErrorSxsXmlEBadnamechar
	case 0x000036D2:
		return ErrorSxsXmlEBadcharinstring
	case 0x000036D3:
		return ErrorSxsXmlEXmldeclsyntax
	case 0x000036D4:
		return ErrorSxsXmlEBadchardata
	case 0x000036D5:
		return ErrorSxsXmlEMissingwhitespace
	case 0x000036D6:
		return ErrorSxsXmlEExpectingtagend
	case 0x000036D7:
		return ErrorSxsXmlEMissingsemicolon
	case 0x000036D8:
		return ErrorSxsXmlEUnbalancedparen
	case 0x000036D9:
		return ErrorSxsXmlEInternalerror
	case 0x000036DA:
		return ErrorSxsXmlEUnexpectedWhitespace
	case 0x000036DB:
		return ErrorSxsXmlEIncompleteEncoding
	case 0x000036DC:
		return ErrorSxsXmlEMissingParen
	case 0x000036DD:
		return ErrorSxsXmlEExpectingclosequote
	case 0x000036DE:
		return ErrorSxsXmlEMultipleColons
	case 0x000036DF:
		return ErrorSxsXmlEInvalidDecimal
	case 0x000036E0:
		return ErrorSxsXmlEInvalidHexidecimal
	case 0x000036E1:
		return ErrorSxsXmlEInvalidUnicode
	case 0x000036E2:
		return ErrorSxsXmlEWhitespaceorquestionmark
	case 0x000036E3:
		return ErrorSxsXmlEUnexpectedendtag
	case 0x000036E4:
		return ErrorSxsXmlEUnclosedtag
	case 0x000036E5:
		return ErrorSxsXmlEDuplicateattribute
	case 0x000036E6:
		return ErrorSxsXmlEMultipleroots
	case 0x000036E7:
		return ErrorSxsXmlEInvalidatrootlevel
	case 0x000036E8:
		return ErrorSxsXmlEBadxmldecl
	case 0x000036E9:
		return ErrorSxsXmlEMissingroot
	case 0x000036EA:
		return ErrorSxsXmlEUnexpectedeof
	case 0x000036EB:
		return ErrorSxsXmlEBadperefinsubset
	case 0x000036EC:
		return ErrorSxsXmlEUnclosedstarttag
	case 0x000036ED:
		return ErrorSxsXmlEUnclosedendtag
	case 0x000036EE:
		return ErrorSxsXmlEUnclosedstring
	case 0x000036EF:
		return ErrorSxsXmlEUnclosedcomment
	case 0x000036F0:
		return ErrorSxsXmlEUncloseddecl
	case 0x000036F1:
		return ErrorSxsXmlEUnclosedcdata
	case 0x000036F2:
		return ErrorSxsXmlEReservednamespace
	case 0x000036F3:
		return ErrorSxsXmlEInvalidencoding
	case 0x000036F4:
		return ErrorSxsXmlEInvalidswitch
	case 0x000036F5:
		return ErrorSxsXmlEBadxmlcase
	case 0x000036F6:
		return ErrorSxsXmlEInvalidStandalone
	case 0x000036F7:
		return ErrorSxsXmlEUnexpectedStandalone
	case 0x000036F8:
		return ErrorSxsXmlEInvalidVersion
	case 0x000036F9:
		return ErrorSxsXmlEMissingequals
	case 0x000036FA:
		return ErrorSxsProtectionRecoveryFailed
	case 0x000036FB:
		return ErrorSxsProtectionPublicKeyOoShort
	case 0x000036FC:
		return ErrorSxsProtectionCatalogNotValid
	case 0x000036FD:
		return ErrorSxsUntranslatableHresult
	case 0x000036FE:
		return ErrorSxsProtectionCatalogFileMissing
	case 0x000036FF:
		return ErrorSxsMissingAssemblyIdentityAttribute
	case 0x00003700:
		return ErrorSxsInvalidAssemblyIdentityAttributeName
	case 0x00003701:
		return ErrorSxsAssemblyMissing
	case 0x00003702:
		return ErrorSxsCorruptActivationStack
	case 0x00003703:
		return ErrorSxsCorruption
	case 0x00003704:
		return ErrorSxsEarlyDeactivation
	case 0x00003705:
		return ErrorSxsInvalidDeactivation
	case 0x00003706:
		return ErrorSxsMultipleDeactivation
	case 0x00003707:
		return ErrorSxsProcessTerminationRequested
	case 0x00003708:
		return ErrorSxsReleaseActivationOntext
	case 0x00003709:
		return ErrorSxsSystemDefaultActivationContextEmpty
	case 0x0000370A:
		return ErrorSxsInvalidIdentityAttributeValue
	case 0x0000370B:
		return ErrorSxsInvalidIdentityAttributeName
	case 0x0000370C:
		return ErrorSxsIdentityDuplicateAttribute
	case 0x0000370D:
		return ErrorSxsIdentityParseError
	case 0x0000370E:
		return ErrorMalformedSubstitutionString
	case 0x0000370F:
		return ErrorSxsIncorrectPublicKeyOken
	case 0x00003710:
		return ErrorUnmappedSubstitutionString
	case 0x00003711:
		return ErrorSxsAssemblyNotLocked
	case 0x00003712:
		return ErrorSxsComponentStoreCorrupt
	case 0x00003713:
		return ErrorAdvancedInstallerFailed
	case 0x00003714:
		return ErrorXmlEncodingMismatch
	case 0x00003715:
		return ErrorSxsManifestIdentitySameButContentsDifferent
	case 0x00003716:
		return ErrorSxsIdentitiesDifferent
	case 0x00003717:
		return ErrorSxsAssemblyIsNotADeployment
	case 0x00003718:
		return ErrorSxsFileNotPartOfAssembly
	case 0x00003719:
		return ErrorSxsManifestTooBig
	case 0x0000371A:
		return ErrorSxsSettingNotRegistered
	case 0x0000371B:
		return ErrorSxsTransactionClosureIncomplete
	case 0x00003A98:
		return ErrorEvtInvalidChannelPath
	case 0x00003A99:
		return ErrorEvtInvalidQuery
	case 0x00003A9A:
		return ErrorEvtPublisherMetadataNotFound
	case 0x00003A9B:
		return ErrorEvtEventTemplateNotFound
	case 0x00003A9C:
		return ErrorEvtInvalidPublisherName
	case 0x00003A9D:
		return ErrorEvtInvalidEventData
	case 0x00003A9F:
		return ErrorEvtChannelNotFound
	case 0x00003AA0:
		return ErrorEvtMalformedXmlText
	case 0x00003AA1:
		return ErrorEvtSubscriptionToDirectChannel
	case 0x00003AA2:
		return ErrorEvtConfigurationError
	case 0x00003AA3:
		return ErrorEvtQueryResultStale
	case 0x00003AA4:
		return ErrorEvtQueryResultInvalidPosition
	case 0x00003AA5:
		return ErrorEvtNonValidatingMsxml
	case 0x00003AA6:
		return ErrorEvtFilterAlreadyscoped
	case 0x00003AA7:
		return ErrorEvtFilterNoteltset
	case 0x00003AA8:
		return ErrorEvtFilterInvarg
	case 0x00003AA9:
		return ErrorEvtFilterInvtest
	case 0x00003AAA:
		return ErrorEvtFilterInvtype
	case 0x00003AAB:
		return ErrorEvtFilterParseerr
	case 0x00003AAC:
		return ErrorEvtFilterUnsupportedop
	case 0x00003AAD:
		return ErrorEvtFilterUnexpectedtoken
	case 0x00003AAE:
		return ErrorEvtInvalidOperationOverEnabledDirectChannel
	case 0x00003AAF:
		return ErrorEvtInvalidChannelPropertyValue
	case 0x00003AB0:
		return ErrorEvtInvalidPublisherPropertyValue
	case 0x00003AB1:
		return ErrorEvtChannelCannotActivate
	case 0x00003AB2:
		return ErrorEvtFilterTooComplex
	case 0x00003AB3:
		return ErrorEvtMessageNotFound
	case 0x00003AB4:
		return ErrorEvtMessageIdNotFound
	case 0x00003AB5:
		return ErrorEvtUnresolvedValueInsert
	case 0x00003AB6:
		return ErrorEvtUnresolvedParameterInsert
	case 0x00003AB7:
		return ErrorEvtMaxInsertsReached
	case 0x00003AB8:
		return ErrorEvtEventDefinitionNotOund
	case 0x00003AB9:
		return ErrorEvtMessageLocaleNotFound
	case 0x00003ABA:
		return ErrorEvtVersionTooOld
	case 0x00003ABB:
		return ErrorEvtVersionTooNew
	case 0x00003ABC:
		return ErrorEvtCannotOpenChannelOfQuery
	case 0x00003ABD:
		return ErrorEvtPublisherDisabled
	case 0x00003AE8:
		return ErrorEcSubscriptionCannotActivate
	case 0x00003AE9:
		return ErrorEcLogDisabled
	case 0x00003AFC:
		return ErrorMuiFileNotFound
	case 0x00003AFD:
		return ErrorMuiInvalidFile
	case 0x00003AFE:
		return ErrorMuiInvalidRcConfig
	case 0x00003AFF:
		return ErrorMuiInvalidLocaleName
	case 0x00003B00:
		return ErrorMuiInvalidUltimatefallbackName
	case 0x00003B01:
		return ErrorMuiFileNotLoaded
	case 0x00003B02:
		return ErrorResourceEnumUserStop
	case 0x00003B03:
		return ErrorMuiIntlsettingsUilangNotInstalled
	case 0x00003B04:
		return ErrorMuiIntlsettingsInvalidLocaleName
	case 0x00003B60:
		return ErrorMcaInvalidCapabilitiesString
	case 0x00003B61:
		return ErrorMcaInvalidVcpVersion
	case 0x00003B62:
		return ErrorMcaMonitorViolatesMccsSpecification
	case 0x00003B63:
		return ErrorMcaMccsVersionMismatch
	case 0x00003B64:
		return ErrorMcaUnsupportedMccsVersion
	case 0x00003B65:
		return ErrorMcaInternalError
	case 0x00003B66:
		return ErrorMcaInvalidTechnologyTypeReturned
	case 0x00003B67:
		return ErrorMcaUnsupportedColorTemperature
	case 0x00003B92:
		return ErrorAmbiguousSystemDevice
	case 0x00003BC3:
		return ErrorSystemDeviceNotFound
	}
	return nil
}

