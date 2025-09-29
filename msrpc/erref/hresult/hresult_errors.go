package hresult

var (
	StgSConverted                                                = &Error{0x00030200, "STG_S_CONVERTED", "The underlying file was converted to compound file format."}
	StgSBlock                                                    = &Error{0x00030201, "STG_S_BLOCK", "The storage operation should block until more data is available."}
	StgSRetrynow                                                 = &Error{0x00030202, "STG_S_RETRYNOW", "The storage operation should retry immediately."}
	StgSMonitoring                                               = &Error{0x00030203, "STG_S_MONITORING", "The notified event sink will not influence the storage operation."}
	StgSMultipleopens                                            = &Error{0x00030204, "STG_S_MULTIPLEOPENS", "Multiple opens prevent consolidated (commit succeeded)."}
	StgSConsolidationfailed                                      = &Error{0x00030205, "STG_S_CONSOLIDATIONFAILED", "Consolidation of the storage file failed (commit succeeded)."}
	StgSCannotconsolidate                                        = &Error{0x00030206, "STG_S_CANNOTCONSOLIDATE", "Consolidation of the storage file is inappropriate (commit succeeded)."}
	OleSUsereg                                                   = &Error{0x00040000, "OLE_S_USEREG", "Use the registry database to provide the requested information."}
	OleSStatic                                                   = &Error{0x00040001, "OLE_S_STATIC", "Success, but static."}
	OleSMacClipformat                                            = &Error{0x00040002, "OLE_S_MAC_CLIPFORMAT", "Macintosh clipboard format."}
	DragdropSDrop                                                = &Error{0x00040100, "DRAGDROP_S_DROP", "Successful drop took place."}
	DragdropSCancel                                              = &Error{0x00040101, "DRAGDROP_S_CANCEL", "Drag-drop operation canceled."}
	DragdropSUsedefaultcursors                                   = &Error{0x00040102, "DRAGDROP_S_USEDEFAULTCURSORS", "Use the default cursor."}
	DataSSameformatetc                                           = &Error{0x00040130, "DATA_S_SAMEFORMATETC", "Data has same FORMATETC."}
	ViewSAlreadyFrozen                                           = &Error{0x00040140, "VIEW_S_ALREADY_FROZEN", "View is already frozen."}
	CacheSFormatetcNotsupported                                  = &Error{0x00040170, "CACHE_S_FORMATETC_NOTSUPPORTED", "FORMATETC not supported."}
	CacheSSamecache                                              = &Error{0x00040171, "CACHE_S_SAMECACHE", "Same cache."}
	CacheSSomecachesNotupdated                                   = &Error{0x00040172, "CACHE_S_SOMECACHES_NOTUPDATED", "Some caches are not updated."}
	OleobjSInvalidverb                                           = &Error{0x00040180, "OLEOBJ_S_INVALIDVERB", "Invalid verb for OLE object."}
	OleobjSCannotDoverbNow                                       = &Error{0x00040181, "OLEOBJ_S_CANNOT_DOVERB_NOW", "Verb number is valid but verb cannot be done now."}
	OleobjSInvalidhwnd                                           = &Error{0x00040182, "OLEOBJ_S_INVALIDHWND", "Invalid window handle passed."}
	InplaceSTruncated                                            = &Error{0x000401A0, "INPLACE_S_TRUNCATED", "Message is too long; some of it had to be truncated before displaying."}
	Convert10SNoPresentation                                     = &Error{0x000401C0, "CONVERT10_S_NO_PRESENTATION", "Unable to convert OLESTREAM to IStorage."}
	MkSReducedToSelf                                             = &Error{0x000401E2, "MK_S_REDUCED_TO_SELF", "Moniker reduced to itself."}
	MkSMe                                                        = &Error{0x000401E4, "MK_S_ME", "Common prefix is this moniker."}
	MkSHim                                                       = &Error{0x000401E5, "MK_S_HIM", "Common prefix is input moniker."}
	MkSUs                                                        = &Error{0x000401E6, "MK_S_US", "Common prefix is both monikers."}
	MkSMonikeralreadyregistered                                  = &Error{0x000401E7, "MK_S_MONIKERALREADYREGISTERED", "Moniker is already registered in running object table."}
	EventSSomeSubscribersFailed                                  = &Error{0x00040200, "EVENT_S_SOME_SUBSCRIBERS_FAILED", "An event was able to invoke some, but not all, of the subscribers."}
	EventSNosubscribers                                          = &Error{0x00040202, "EVENT_S_NOSUBSCRIBERS", "An event was delivered, but there were no subscribers."}
	SchedSTaskReady                                              = &Error{0x00041300, "SCHED_S_TASK_READY", "The task is ready to run at its next scheduled time."}
	SchedSTaskRunning                                            = &Error{0x00041301, "SCHED_S_TASK_RUNNING", "The task is currently running."}
	SchedSTaskDisabled                                           = &Error{0x00041302, "SCHED_S_TASK_DISABLED", "The task will not run at the scheduled times because it has been disabled."}
	SchedSTaskHasNotRun                                          = &Error{0x00041303, "SCHED_S_TASK_HAS_NOT_RUN", "The task has not yet run."}
	SchedSTaskNoMoreRuns                                         = &Error{0x00041304, "SCHED_S_TASK_NO_MORE_RUNS", "There are no more runs scheduled for this task."}
	SchedSTaskNotScheduled                                       = &Error{0x00041305, "SCHED_S_TASK_NOT_SCHEDULED", "One or more of the properties that are needed to run this task on a schedule have not been set."}
	SchedSTaskTerminated                                         = &Error{0x00041306, "SCHED_S_TASK_TERMINATED", "The last run of the task was terminated by the user."}
	SchedSTaskNoValidTriggers                                    = &Error{0x00041307, "SCHED_S_TASK_NO_VALID_TRIGGERS", "Either the task has no triggers, or the existing triggers are disabled or not set."}
	SchedSEventTrigger                                           = &Error{0x00041308, "SCHED_S_EVENT_TRIGGER", "Event triggers do not have set run times."}
	SchedSSomeTriggersFailed                                     = &Error{0x0004131B, "SCHED_S_SOME_TRIGGERS_FAILED", "The task is registered, but not all specified triggers will start the task."}
	SchedSBatchLogonProblem                                      = &Error{0x0004131C, "SCHED_S_BATCH_LOGON_PROBLEM", "The task is registered, but it might fail to start. Batch logon privilege needs to be enabled for the task principal."}
	XactSAsync                                                   = &Error{0x0004D000, "XACT_S_ASYNC", "An asynchronous operation was specified. The operation has begun, but its outcome is not known yet."}
	XactSReadonly                                                = &Error{0x0004D002, "XACT_S_READONLY", "The method call succeeded because the transaction was read-only."}
	XactSSomenoretain                                            = &Error{0x0004D003, "XACT_S_SOMENORETAIN", "The transaction was successfully aborted. However, this is a coordinated transaction, and a number of enlisted resources were aborted outright because they could not support abort-retaining semantics."}
	XactSOkinform                                                = &Error{0x0004D004, "XACT_S_OKINFORM", "No changes were made during this call, but the sink wants another chance to look if any other sinks make further changes."}
	XactSMadechangescontent                                      = &Error{0x0004D005, "XACT_S_MADECHANGESCONTENT", "The sink is content and wants the transaction to proceed. Changes were made to one or more resources during this call."}
	XactSMadechangesinform                                       = &Error{0x0004D006, "XACT_S_MADECHANGESINFORM", "The sink is for the moment and wants the transaction to proceed, but if other changes are made following this return by other event sinks, this sink wants another chance to look."}
	XactSAllnoretain                                             = &Error{0x0004D007, "XACT_S_ALLNORETAIN", "The transaction was successfully aborted. However, the abort was nonretaining."}
	XactSAborting                                                = &Error{0x0004D008, "XACT_S_ABORTING", "An abort operation was already in progress."}
	XactSSinglephase                                             = &Error{0x0004D009, "XACT_S_SINGLEPHASE", "The resource manager has performed a single-phase commit of the transaction."}
	XactSLocallyOk                                               = &Error{0x0004D00A, "XACT_S_LOCALLY_OK", "The local transaction has not aborted."}
	XactSLastresourcemanager                                     = &Error{0x0004D010, "XACT_S_LASTRESOURCEMANAGER", "The resource manager has requested to be the coordinator (last resource manager) for the transaction."}
	CoSNotallinterfaces                                          = &Error{0x00080012, "CO_S_NOTALLINTERFACES", "Not all the requested interfaces were available."}
	CoSMachinenamenotfound                                       = &Error{0x00080013, "CO_S_MACHINENAMENOTFOUND", "The specified machine name was not found in the cache."}
	SecIContinueNeeded                                           = &Error{0x00090312, "SEC_I_CONTINUE_NEEDED", "The function completed successfully, but it must be called again to complete the context."}
	SecICompleteNeeded                                           = &Error{0x00090313, "SEC_I_COMPLETE_NEEDED", "The function completed successfully, but CompleteToken must be called."}
	SecICompleteAndContinue                                      = &Error{0x00090314, "SEC_I_COMPLETE_AND_CONTINUE", "The function completed successfully, but both CompleteToken and this function must be called to complete the context."}
	SecILocalLogon                                               = &Error{0x00090315, "SEC_I_LOCAL_LOGON", "The logon was completed, but no network authority was available. The logon was made using locally known information."}
	SecIContextExpired                                           = &Error{0x00090317, "SEC_I_CONTEXT_EXPIRED", "The context has expired and can no longer be used."}
	SecIIncompleteCredentials                                    = &Error{0x00090320, "SEC_I_INCOMPLETE_CREDENTIALS", "The credentials supplied were not complete and could not be verified. Additional information can be returned from the context."}
	SecIRenegotiate                                              = &Error{0x00090321, "SEC_I_RENEGOTIATE", "The context data must be renegotiated with the peer."}
	SecINoLsaContext                                             = &Error{0x00090323, "SEC_I_NO_LSA_CONTEXT", "There is no LSA mode context associated with this context."}
	SecISignatureNeeded                                          = &Error{0x0009035C, "SEC_I_SIGNATURE_NEEDED", "A signature operation must be performed before the user can authenticate."}
	CryptINewProtectionRequired                                  = &Error{0x00091012, "CRYPT_I_NEW_PROTECTION_REQUIRED", "The protected data needs to be reprotected."}
	NsSCallpending                                               = &Error{0x000D0000, "NS_S_CALLPENDING", "The requested operation is pending completion."}
	NsSCallaborted                                               = &Error{0x000D0001, "NS_S_CALLABORTED", "The requested operation was aborted by the client."}
	NsSStreamTruncated                                           = &Error{0x000D0002, "NS_S_STREAM_TRUNCATED", "The stream was purposefully stopped before completion."}
	NsSRebuffering                                               = &Error{0x000D0BC8, "NS_S_REBUFFERING", "The requested operation has caused the source to rebuffer."}
	NsSDegradingQuality                                          = &Error{0x000D0BC9, "NS_S_DEGRADING_QUALITY", "The requested operation has caused the source to degrade codec quality."}
	NsSTranscryptorEof                                           = &Error{0x000D0BDB, "NS_S_TRANSCRYPTOR_EOF", "The transcryptor object has reached end of file."}
	NsSWmpUiVersionmismatch                                      = &Error{0x000D0FE8, "NS_S_WMP_UI_VERSIONMISMATCH", "An upgrade is needed for the theme manager to correctly show this skin. Skin reports version: %.1f."}
	NsSWmpException                                              = &Error{0x000D0FE9, "NS_S_WMP_EXCEPTION", "An error occurred in one of the UI components."}
	NsSWmpLoadedGifImage                                         = &Error{0x000D1040, "NS_S_WMP_LOADED_GIF_IMAGE", "Successfully loaded a GIF file."}
	NsSWmpLoadedPngImage                                         = &Error{0x000D1041, "NS_S_WMP_LOADED_PNG_IMAGE", "Successfully loaded a PNG file."}
	NsSWmpLoadedBmpImage                                         = &Error{0x000D1042, "NS_S_WMP_LOADED_BMP_IMAGE", "Successfully loaded a BMP file."}
	NsSWmpLoadedJpgImage                                         = &Error{0x000D1043, "NS_S_WMP_LOADED_JPG_IMAGE", "Successfully loaded a JPG file."}
	NsSWmgForceDropFrame                                         = &Error{0x000D104F, "NS_S_WMG_FORCE_DROP_FRAME", "Drop this frame."}
	NsSWmrAlreadyrendered                                        = &Error{0x000D105F, "NS_S_WMR_ALREADYRENDERED", "The specified stream has already been rendered."}
	NsSWmrPintypepartialmatch                                    = &Error{0x000D1060, "NS_S_WMR_PINTYPEPARTIALMATCH", "The specified type partially matches this pin type."}
	NsSWmrPintypefullmatch                                       = &Error{0x000D1061, "NS_S_WMR_PINTYPEFULLMATCH", "The specified type fully matches this pin type."}
	NsSWmgAdviseDropFrame                                        = &Error{0x000D1066, "NS_S_WMG_ADVISE_DROP_FRAME", "The timestamp is late compared to the current render position. Advise dropping this frame."}
	NsSWmgAdviseDropToKeyframe                                   = &Error{0x000D1067, "NS_S_WMG_ADVISE_DROP_TO_KEYFRAME", "The timestamp is severely late compared to the current render position. Advise dropping everything up to the next key frame."}
	NsSNeedToBuyBurnRights                                       = &Error{0x000D10DB, "NS_S_NEED_TO_BUY_BURN_RIGHTS", "No burn rights. You will be prompted to buy burn rights when you try to burn this file to an audio CD."}
	NsSWmpcorePlaylistclearabort                                 = &Error{0x000D10FE, "NS_S_WMPCORE_PLAYLISTCLEARABORT", "Failed to clear playlist because it was aborted by user."}
	NsSWmpcorePlaylistremoveitemabort                            = &Error{0x000D10FF, "NS_S_WMPCORE_PLAYLISTREMOVEITEMABORT", "Failed to remove item in the playlist since it was aborted by user."}
	NsSWmpcorePlaylistCreationPending                            = &Error{0x000D1102, "NS_S_WMPCORE_PLAYLIST_CREATION_PENDING", "Playlist is being generated asynchronously."}
	NsSWmpcoreMediaValidationPending                             = &Error{0x000D1103, "NS_S_WMPCORE_MEDIA_VALIDATION_PENDING", "Validation of the media is pending."}
	NsSWmpcorePlaylistRepeatSecondarySegmentsIgnored             = &Error{0x000D1104, "NS_S_WMPCORE_PLAYLIST_REPEAT_SECONDARY_SEGMENTS_IGNORED", "Encountered more than one Repeat block during ASX processing."}
	NsSWmpcoreCommandNotAvailable                                = &Error{0x000D1105, "NS_S_WMPCORE_COMMAND_NOT_AVAILABLE", "Current state of WMP disallows calling this method or property."}
	NsSWmpcorePlaylistNameAutoGenerated                          = &Error{0x000D1106, "NS_S_WMPCORE_PLAYLIST_NAME_AUTO_GENERATED", "Name for the playlist has been auto generated."}
	NsSWmpcorePlaylistImportMissingItems                         = &Error{0x000D1107, "NS_S_WMPCORE_PLAYLIST_IMPORT_MISSING_ITEMS", "The imported playlist does not contain all items from the original."}
	NsSWmpcorePlaylistCollapsedToSingleMedia                     = &Error{0x000D1108, "NS_S_WMPCORE_PLAYLIST_COLLAPSED_TO_SINGLE_MEDIA", "The M3U playlist has been ignored because it only contains one item."}
	NsSWmpcoreMediaChildPlaylistOpenPending                      = &Error{0x000D1109, "NS_S_WMPCORE_MEDIA_CHILD_PLAYLIST_OPEN_PENDING", "The open for the child playlist associated with this media is pending."}
	NsSWmpcoreMoreNodesAvaiable                                  = &Error{0x000D110A, "NS_S_WMPCORE_MORE_NODES_AVAIABLE", "More nodes support the interface requested, but the array for returning them is full."}
	NsSWmpbrSuccess                                              = &Error{0x000D1135, "NS_S_WMPBR_SUCCESS", "Backup or Restore successful!."}
	NsSWmpbrPartialsuccess                                       = &Error{0x000D1136, "NS_S_WMPBR_PARTIALSUCCESS", "Transfer complete with limitations."}
	NsSWmpeffectTransparent                                      = &Error{0x000D1144, "NS_S_WMPEFFECT_TRANSPARENT", "Request to the effects control to change transparency status to transparent."}
	NsSWmpeffectOpaque                                           = &Error{0x000D1145, "NS_S_WMPEFFECT_OPAQUE", "Request to the effects control to change transparency status to opaque."}
	NsSOperationPending                                          = &Error{0x000D114E, "NS_S_OPERATION_PENDING", "The requested application pane is performing an operation and will not be released."}
	NsSTrackBuyRequiresAlbumPurchase                             = &Error{0x000D1359, "NS_S_TRACK_BUY_REQUIRES_ALBUM_PURCHASE", "The file is only available for purchase when you buy the entire album."}
	NsSNavigationCompleteWithErrors                              = &Error{0x000D135E, "NS_S_NAVIGATION_COMPLETE_WITH_ERRORS", "There were problems completing the requested navigation. There are identifiers missing in the catalog."}
	NsSTrackAlreadyDownloaded                                    = &Error{0x000D1361, "NS_S_TRACK_ALREADY_DOWNLOADED", "Track already downloaded."}
	NsSPublishingPointStartedWithFailedSinks                     = &Error{0x000D1519, "NS_S_PUBLISHING_POINT_STARTED_WITH_FAILED_SINKS", "The publishing point successfully started, but one or more of the requested data writer plug-ins failed."}
	NsSDrmLicenseAcquired                                        = &Error{0x000D2726, "NS_S_DRM_LICENSE_ACQUIRED", "Status message: The license was acquired."}
	NsSDrmIndividualized                                         = &Error{0x000D2727, "NS_S_DRM_INDIVIDUALIZED", "Status message: The security upgrade has been completed."}
	NsSDrmMonitorCancelled                                       = &Error{0x000D2746, "NS_S_DRM_MONITOR_CANCELLED", "Status message: License monitoring has been canceled."}
	NsSDrmAcquireCancelled                                       = &Error{0x000D2747, "NS_S_DRM_ACQUIRE_CANCELLED", "Status message: License acquisition has been canceled."}
	NsSDrmBurnableTrack                                          = &Error{0x000D276E, "NS_S_DRM_BURNABLE_TRACK", "The track is burnable and had no playlist burn limit."}
	NsSDrmBurnableTrackWithPlaylistRestriction                   = &Error{0x000D276F, "NS_S_DRM_BURNABLE_TRACK_WITH_PLAYLIST_RESTRICTION", "The track is burnable but has a playlist burn limit."}
	NsSDrmNeedsIndividualization                                 = &Error{0x000D27DE, "NS_S_DRM_NEEDS_INDIVIDUALIZATION", "A security upgrade is required to perform the operation on this media file."}
	NsSRebootRecommended                                         = &Error{0x000D2AF8, "NS_S_REBOOT_RECOMMENDED", "Installation was successful; however, some file cleanup is not complete. For best results, restart your computer."}
	NsSRebootRequired                                            = &Error{0x000D2AF9, "NS_S_REBOOT_REQUIRED", "Installation was successful; however, some file cleanup is not complete. To continue, you must restart your computer."}
	NsSEosreceding                                               = &Error{0x000D2F09, "NS_S_EOSRECEDING", "EOS hit during rewinding."}
	NsSChangenotice                                              = &Error{0x000D2F0D, "NS_S_CHANGENOTICE", "Internal."}
	ErrorFltIoComplete                                           = &Error{0x001F0001, "ERROR_FLT_IO_COMPLETE", "The IO was completed by a filter."}
	ErrorGraphicsModeNotPinned                                   = &Error{0x00262307, "ERROR_GRAPHICS_MODE_NOT_PINNED", "No mode is pinned on the specified VidPN source or target."}
	ErrorGraphicsNoPreferredMode                                 = &Error{0x0026231E, "ERROR_GRAPHICS_NO_PREFERRED_MODE", "Specified mode set does not specify preference for one of its modes."}
	ErrorGraphicsDatasetIsEmpty                                  = &Error{0x0026234B, "ERROR_GRAPHICS_DATASET_IS_EMPTY", "Specified data set (for example, mode set, frequency range set, descriptor set, and topology) is empty."}
	ErrorGraphicsNoMoreElementsInDataset                         = &Error{0x0026234C, "ERROR_GRAPHICS_NO_MORE_ELEMENTS_IN_DATASET", "Specified data set (for example, mode set, frequency range set, descriptor set, and topology) does not contain any more elements."}
	ErrorGraphicsPathContentGeometryTransformationNotPinned      = &Error{0x00262351, "ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_PINNED", "Specified content transformation is not pinned on the specified VidPN present path."}
	PlaSPropertyIgnored                                          = &Error{0x00300100, "PLA_S_PROPERTY_IGNORED", "Property value will be ignored."}
	ErrorNdisIndicationRequired                                  = &Error{0x00340001, "ERROR_NDIS_INDICATION_REQUIRED", "The request will be completed later by a Network Driver Interface Specification (NDIS) status indication."}
	TrkSOutOfSync                                                = &Error{0x0DEAD100, "TRK_S_OUT_OF_SYNC", "The VolumeSequenceNumber of a MOVE_NOTIFICATION request is incorrect."}
	TrkVolumeNotFound                                            = &Error{0x0DEAD102, "TRK_VOLUME_NOT_FOUND", "The VolumeID in a request was not found in the server's ServerVolumeTable."}
	TrkVolumeNotOwned                                            = &Error{0x0DEAD103, "TRK_VOLUME_NOT_OWNED", "A notification was sent to the LnkSvrMessage method, but the RequestMachine for the request was not the VolumeOwner for a VolumeID in the request."}
	TrkSNotificationQuotaExceeded                                = &Error{0x0DEAD107, "TRK_S_NOTIFICATION_QUOTA_EXCEEDED", "The server received a MOVE_NOTIFICATION request, but the FileTable size limit has already been reached."}
	NsITigerStart                                                = &Error{0x400D004F, "NS_I_TIGER_START", "The Title Server %1 is running."}
	NsICubStart                                                  = &Error{0x400D0051, "NS_I_CUB_START", "Content Server %1 (%2) is starting."}
	NsICubRunning                                                = &Error{0x400D0052, "NS_I_CUB_RUNNING", "Content Server %1 (%2) is running."}
	NsIDiskStart                                                 = &Error{0x400D0054, "NS_I_DISK_START", "Disk %1 ( %2 ) on Content Server %3, is running."}
	NsIDiskRebuildStarted                                        = &Error{0x400D0056, "NS_I_DISK_REBUILD_STARTED", "Started rebuilding disk %1 ( %2 ) on Content Server %3."}
	NsIDiskRebuildFinished                                       = &Error{0x400D0057, "NS_I_DISK_REBUILD_FINISHED", "Finished rebuilding disk %1 ( %2 ) on Content Server %3."}
	NsIDiskRebuildAborted                                        = &Error{0x400D0058, "NS_I_DISK_REBUILD_ABORTED", "Aborted rebuilding disk %1 ( %2 ) on Content Server %3."}
	NsILimitFunnels                                              = &Error{0x400D0059, "NS_I_LIMIT_FUNNELS", "A NetShow administrator at network location %1 set the data stream limit to %2 streams."}
	NsIStartDisk                                                 = &Error{0x400D005A, "NS_I_START_DISK", "A NetShow administrator at network location %1 started disk %2."}
	NsIStopDisk                                                  = &Error{0x400D005B, "NS_I_STOP_DISK", "A NetShow administrator at network location %1 stopped disk %2."}
	NsIStopCub                                                   = &Error{0x400D005C, "NS_I_STOP_CUB", "A NetShow administrator at network location %1 stopped Content Server %2."}
	NsIKillUsersession                                           = &Error{0x400D005D, "NS_I_KILL_USERSESSION", "A NetShow administrator at network location %1 aborted user session %2 from the system."}
	NsIKillConnection                                            = &Error{0x400D005E, "NS_I_KILL_CONNECTION", "A NetShow administrator at network location %1 aborted obsolete connection %2 from the system."}
	NsIRebuildDisk                                               = &Error{0x400D005F, "NS_I_REBUILD_DISK", "A NetShow administrator at network location %1 started rebuilding disk %2."}
	McmadmINoEvents                                              = &Error{0x400D0069, "MCMADM_I_NO_EVENTS", "Event initialization failed, there will be no MCM events."}
	NsILoggingFailed                                             = &Error{0x400D006E, "NS_I_LOGGING_FAILED", "The logging operation failed."}
	NsILimitBandwidth                                            = &Error{0x400D0070, "NS_I_LIMIT_BANDWIDTH", "A NetShow administrator at network location %1 set the maximum bandwidth limit to %2 bps."}
	NsICubUnfailLink                                             = &Error{0x400D0191, "NS_I_CUB_UNFAIL_LINK", "Content Server %1 (%2) has established its link to Content Server %3."}
	NsIRestripeStart                                             = &Error{0x400D0193, "NS_I_RESTRIPE_START", "Restripe operation has started."}
	NsIRestripeDone                                              = &Error{0x400D0194, "NS_I_RESTRIPE_DONE", "Restripe operation has completed."}
	NsIRestripeDiskOut                                           = &Error{0x400D0196, "NS_I_RESTRIPE_DISK_OUT", "Content disk %1 (%2) on Content Server %3 has been restriped out."}
	NsIRestripeCubOut                                            = &Error{0x400D0197, "NS_I_RESTRIPE_CUB_OUT", "Content server %1 (%2) has been restriped out."}
	NsIDiskStop                                                  = &Error{0x400D0198, "NS_I_DISK_STOP", "Disk %1 ( %2 ) on Content Server %3, has been offlined."}
	NsIPlaylistChangeReceding                                    = &Error{0x400D14BE, "NS_I_PLAYLIST_CHANGE_RECEDING", "The playlist change occurred while receding."}
	NsIReconnected                                               = &Error{0x400D2EFF, "NS_I_RECONNECTED", "The client is reconnected."}
	NsINologStop                                                 = &Error{0x400D2F01, "NS_I_NOLOG_STOP", "Forcing a switch to a pending header on start."}
	NsIExistingPacketizer                                        = &Error{0x400D2F03, "NS_I_EXISTING_PACKETIZER", "There is already an existing packetizer plugin for the stream."}
	NsIManualProxy                                               = &Error{0x400D2F04, "NS_I_MANUAL_PROXY", "The proxy setting is manual."}
	ErrorGraphicsDriverMismatch                                  = &Error{0x40262009, "ERROR_GRAPHICS_DRIVER_MISMATCH", "The kernel driver detected a version mismatch between it and the user mode driver."}
	ErrorGraphicsUnknownChildStatus                              = &Error{0x4026242F, "ERROR_GRAPHICS_UNKNOWN_CHILD_STATUS", "Child device presence was not reliably detected."}
	ErrorGraphicsLeadlinkStartDeferred                           = &Error{0x40262437, "ERROR_GRAPHICS_LEADLINK_START_DEFERRED", "Starting the lead-link adapter has been deferred temporarily."}
	ErrorGraphicsPollingTooFrequently                            = &Error{0x40262439, "ERROR_GRAPHICS_POLLING_TOO_FREQUENTLY", "The display adapter is being polled for children too frequently at the same polling level."}
	ErrorGraphicsStartDeferred                                   = &Error{0x4026243A, "ERROR_GRAPHICS_START_DEFERRED", "Starting the adapter has been deferred temporarily."}
	EPending                                                     = &Error{0x8000000A, "E_PENDING", "The data necessary to complete this operation is not yet available."}
	ENotimpl                                                     = &Error{0x80004001, "E_NOTIMPL", "Not implemented."}
	ENointerface                                                 = &Error{0x80004002, "E_NOINTERFACE", "No such interface supported."}
	EPointer                                                     = &Error{0x80004003, "E_POINTER", "Invalid pointer."}
	EAbort                                                       = &Error{0x80004004, "E_ABORT", "Operation aborted."}
	EFail                                                        = &Error{0x80004005, "E_FAIL", "Unspecified error."}
	CoEInitTls                                                   = &Error{0x80004006, "CO_E_INIT_TLS", "Thread local storage failure."}
	CoEInitSharedAllocator                                       = &Error{0x80004007, "CO_E_INIT_SHARED_ALLOCATOR", "Get shared memory allocator failure."}
	CoEInitMemoryAllocator                                       = &Error{0x80004008, "CO_E_INIT_MEMORY_ALLOCATOR", "Get memory allocator failure."}
	CoEInitClassCache                                            = &Error{0x80004009, "CO_E_INIT_CLASS_CACHE", "Unable to initialize class cache."}
	CoEInitRpcChannel                                            = &Error{0x8000400A, "CO_E_INIT_RPC_CHANNEL", "Unable to initialize remote procedure call (RPC) services."}
	CoEInitTlsSetChannelControl                                  = &Error{0x8000400B, "CO_E_INIT_TLS_SET_CHANNEL_CONTROL", "Cannot set thread local storage channel control."}
	CoEInitTlsChannelControl                                     = &Error{0x8000400C, "CO_E_INIT_TLS_CHANNEL_CONTROL", "Could not allocate thread local storage channel control."}
	CoEInitUnacceptedUserAllocator                               = &Error{0x8000400D, "CO_E_INIT_UNACCEPTED_USER_ALLOCATOR", "The user-supplied memory allocator is unacceptable."}
	CoEInitScmMutexExists                                        = &Error{0x8000400E, "CO_E_INIT_SCM_MUTEX_EXISTS", "The OLE service mutex already exists."}
	CoEInitScmFileMappingExists                                  = &Error{0x8000400F, "CO_E_INIT_SCM_FILE_MAPPING_EXISTS", "The OLE service file mapping already exists."}
	CoEInitScmMapViewOfFile                                      = &Error{0x80004010, "CO_E_INIT_SCM_MAP_VIEW_OF_FILE", "Unable to map view of file for OLE service."}
	CoEInitScmExecFailure                                        = &Error{0x80004011, "CO_E_INIT_SCM_EXEC_FAILURE", "Failure attempting to launch OLE service."}
	CoEInitOnlySingleThreaded                                    = &Error{0x80004012, "CO_E_INIT_ONLY_SINGLE_THREADED", "There was an attempt to call CoInitialize a second time while single-threaded."}
	CoECantRemote                                                = &Error{0x80004013, "CO_E_CANT_REMOTE", "A Remote activation was necessary but was not allowed."}
	CoEBadServerName                                             = &Error{0x80004014, "CO_E_BAD_SERVER_NAME", "A Remote activation was necessary, but the server name provided was invalid."}
	CoEWrongServerIdentity                                       = &Error{0x80004015, "CO_E_WRONG_SERVER_IDENTITY", "The class is configured to run as a security ID different from the caller."}
	CoEOle1ddeDisabled                                           = &Error{0x80004016, "CO_E_OLE1DDE_DISABLED", "Use of OLE1 services requiring Dynamic Data Exchange (DDE) Windows is disabled."}
	CoERunasSyntax                                               = &Error{0x80004017, "CO_E_RUNAS_SYNTAX", "A RunAs specification must be <domain name>\\<user name> or simply <user name>."}
	CoECreateprocessFailure                                      = &Error{0x80004018, "CO_E_CREATEPROCESS_FAILURE", "The server process could not be started. The path name might be incorrect."}
	CoERunasCreateprocessFailure                                 = &Error{0x80004019, "CO_E_RUNAS_CREATEPROCESS_FAILURE", "The server process could not be started as the configured identity. The path name might be incorrect or unavailable."}
	CoERunasLogonFailure                                         = &Error{0x8000401A, "CO_E_RUNAS_LOGON_FAILURE", "The server process could not be started because the configured identity is incorrect. Check the user name and password."}
	CoELaunchPermssionDenied                                     = &Error{0x8000401B, "CO_E_LAUNCH_PERMSSION_DENIED", "The client is not allowed to launch this server."}
	CoEStartServiceFailure                                       = &Error{0x8000401C, "CO_E_START_SERVICE_FAILURE", "The service providing this server could not be started."}
	CoERemoteCommunicationFailure                                = &Error{0x8000401D, "CO_E_REMOTE_COMMUNICATION_FAILURE", "This computer was unable to communicate with the computer providing the server."}
	CoEServerStartTimeout                                        = &Error{0x8000401E, "CO_E_SERVER_START_TIMEOUT", "The server did not respond after being launched."}
	CoEClsregInconsistent                                        = &Error{0x8000401F, "CO_E_CLSREG_INCONSISTENT", "The registration information for this server is inconsistent or incomplete."}
	CoEIidregInconsistent                                        = &Error{0x80004020, "CO_E_IIDREG_INCONSISTENT", "The registration information for this interface is inconsistent or incomplete."}
	CoENotSupported                                              = &Error{0x80004021, "CO_E_NOT_SUPPORTED", "The operation attempted is not supported."}
	CoEReloadDll                                                 = &Error{0x80004022, "CO_E_RELOAD_DLL", "A DLL must be loaded."}
	CoEMsiError                                                  = &Error{0x80004023, "CO_E_MSI_ERROR", "A Microsoft Software Installer error was encountered."}
	CoEAttemptToCreateOutsideClientContext                       = &Error{0x80004024, "CO_E_ATTEMPT_TO_CREATE_OUTSIDE_CLIENT_CONTEXT", "The specified activation could not occur in the client context as specified."}
	CoEServerPaused                                              = &Error{0x80004025, "CO_E_SERVER_PAUSED", "Activations on the server are paused."}
	CoEServerNotPaused                                           = &Error{0x80004026, "CO_E_SERVER_NOT_PAUSED", "Activations on the server are not paused."}
	CoEClassDisabled                                             = &Error{0x80004027, "CO_E_CLASS_DISABLED", "The component or application containing the component has been disabled."}
	CoEClrnotavailable                                           = &Error{0x80004028, "CO_E_CLRNOTAVAILABLE", "The common language runtime is not available."}
	CoEAsyncWorkRejected                                         = &Error{0x80004029, "CO_E_ASYNC_WORK_REJECTED", "The thread-pool rejected the submitted asynchronous work."}
	CoEServerInitTimeout                                         = &Error{0x8000402A, "CO_E_SERVER_INIT_TIMEOUT", "The server started, but it did not finish initializing in a timely fashion."}
	CoENoSecctxInActivate                                        = &Error{0x8000402B, "CO_E_NO_SECCTX_IN_ACTIVATE", "Unable to complete the call because there is no COM+ security context inside IObjectControl.Activate."}
	CoETrackerConfig                                             = &Error{0x80004030, "CO_E_TRACKER_CONFIG", "The provided tracker configuration is invalid."}
	CoEThreadpoolConfig                                          = &Error{0x80004031, "CO_E_THREADPOOL_CONFIG", "The provided thread pool configuration is invalid."}
	CoESxsConfig                                                 = &Error{0x80004032, "CO_E_SXS_CONFIG", "The provided side-by-side configuration is invalid."}
	CoEMalformedSpn                                              = &Error{0x80004033, "CO_E_MALFORMED_SPN", "The server principal name (SPN) obtained during security negotiation is malformed."}
	EUnexpected                                                  = &Error{0x8000FFFF, "E_UNEXPECTED", "Catastrophic failure."}
	RpcECallRejected                                             = &Error{0x80010001, "RPC_E_CALL_REJECTED", "Call was rejected by callee."}
	RpcECallCanceled                                             = &Error{0x80010002, "RPC_E_CALL_CANCELED", "Call was canceled by the message filter."}
	RpcECantpostInsendcall                                       = &Error{0x80010003, "RPC_E_CANTPOST_INSENDCALL", "The caller is dispatching an intertask SendMessage call and cannot call out via PostMessage."}
	RpcECantcalloutInasynccall                                   = &Error{0x80010004, "RPC_E_CANTCALLOUT_INASYNCCALL", "The caller is dispatching an asynchronous call and cannot make an outgoing call on behalf of this call."}
	RpcECantcalloutInexternalcall                                = &Error{0x80010005, "RPC_E_CANTCALLOUT_INEXTERNALCALL", "It is illegal to call out while inside message filter."}
	RpcEConnectionTerminated                                     = &Error{0x80010006, "RPC_E_CONNECTION_TERMINATED", "The connection terminated or is in a bogus state and can no longer be used. Other connections are still valid."}
	RpcEServerDied                                               = &Error{0x80010007, "RPC_E_SERVER_DIED", "The callee (the server, not the server application) is not available and disappeared; all connections are invalid. The call might have executed."}
	RpcEClientDied                                               = &Error{0x80010008, "RPC_E_CLIENT_DIED", "The caller (client) disappeared while the callee (server) was processing a call."}
	RpcEInvalidDatapacket                                        = &Error{0x80010009, "RPC_E_INVALID_DATAPACKET", "The data packet with the marshaled parameter data is incorrect."}
	RpcECanttransmitCall                                         = &Error{0x8001000A, "RPC_E_CANTTRANSMIT_CALL", "The call was not transmitted properly; the message queue was full and was not emptied after yielding."}
	RpcEClientCantmarshalData                                    = &Error{0x8001000B, "RPC_E_CLIENT_CANTMARSHAL_DATA", "The client RPC caller cannot marshal the parameter data due to errors (such as low memory)."}
	RpcEClientCantunmarshalData                                  = &Error{0x8001000C, "RPC_E_CLIENT_CANTUNMARSHAL_DATA", "The client RPC caller cannot unmarshal the return data due to errors (such as low memory)."}
	RpcEServerCantmarshalData                                    = &Error{0x8001000D, "RPC_E_SERVER_CANTMARSHAL_DATA", "The server RPC callee cannot marshal the return data due to errors (such as low memory)."}
	RpcEServerCantunmarshalData                                  = &Error{0x8001000E, "RPC_E_SERVER_CANTUNMARSHAL_DATA", "The server RPC callee cannot unmarshal the parameter data due to errors (such as low memory)."}
	RpcEInvalidData                                              = &Error{0x8001000F, "RPC_E_INVALID_DATA", "Received data is invalid. The data might be server or client data."}
	RpcEInvalidParameter                                         = &Error{0x80010010, "RPC_E_INVALID_PARAMETER", "A particular parameter is invalid and cannot be (un)marshaled."}
	RpcECantcalloutAgain                                         = &Error{0x80010011, "RPC_E_CANTCALLOUT_AGAIN", "There is no second outgoing call on same channel in DDE conversation."}
	RpcEServerDiedDne                                            = &Error{0x80010012, "RPC_E_SERVER_DIED_DNE", "The callee (the server, not the server application) is not available and disappeared; all connections are invalid. The call did not execute."}
	RpcESysCallFailed                                            = &Error{0x80010100, "RPC_E_SYS_CALL_FAILED", "System call failed."}
	RpcEOutOfResources                                           = &Error{0x80010101, "RPC_E_OUT_OF_RESOURCES", "Could not allocate some required resource (such as memory or events)"}
	RpcEAttemptedMultithread                                     = &Error{0x80010102, "RPC_E_ATTEMPTED_MULTITHREAD", "Attempted to make calls on more than one thread in single-threaded mode."}
	RpcENotRegistered                                            = &Error{0x80010103, "RPC_E_NOT_REGISTERED", "The requested interface is not registered on the server object."}
	RpcEFault                                                    = &Error{0x80010104, "RPC_E_FAULT", "RPC could not call the server or could not return the results of calling the server."}
	RpcEServerfault                                              = &Error{0x80010105, "RPC_E_SERVERFAULT", "The server threw an exception."}
	RpcEChangedMode                                              = &Error{0x80010106, "RPC_E_CHANGED_MODE", "Cannot change thread mode after it is set."}
	RpcEInvalidmethod                                            = &Error{0x80010107, "RPC_E_INVALIDMETHOD", "The method called does not exist on the server."}
	RpcEDisconnected                                             = &Error{0x80010108, "RPC_E_DISCONNECTED", "The object invoked has disconnected from its clients."}
	RpcERetry                                                    = &Error{0x80010109, "RPC_E_RETRY", "The object invoked chose not to process the call now. Try again later."}
	RpcEServercallRetrylater                                     = &Error{0x8001010A, "RPC_E_SERVERCALL_RETRYLATER", "The message filter indicated that the application is busy."}
	RpcEServercallRejected                                       = &Error{0x8001010B, "RPC_E_SERVERCALL_REJECTED", "The message filter rejected the call."}
	RpcEInvalidCalldata                                          = &Error{0x8001010C, "RPC_E_INVALID_CALLDATA", "A call control interface was called with invalid data."}
	RpcECantcalloutIninputsynccall                               = &Error{0x8001010D, "RPC_E_CANTCALLOUT_ININPUTSYNCCALL", "An outgoing call cannot be made because the application is dispatching an input-synchronous call."}
	RpcEWrongThread                                              = &Error{0x8001010E, "RPC_E_WRONG_THREAD", "The application called an interface that was marshaled for a different thread."}
	RpcEThreadNotInit                                            = &Error{0x8001010F, "RPC_E_THREAD_NOT_INIT", "CoInitialize has not been called on the current thread."}
	RpcEVersionMismatch                                          = &Error{0x80010110, "RPC_E_VERSION_MISMATCH", "The version of OLE on the client and server machines does not match."}
	RpcEInvalidHeader                                            = &Error{0x80010111, "RPC_E_INVALID_HEADER", "OLE received a packet with an invalid header."}
	RpcEInvalidExtension                                         = &Error{0x80010112, "RPC_E_INVALID_EXTENSION", "OLE received a packet with an invalid extension."}
	RpcEInvalidIpid                                              = &Error{0x80010113, "RPC_E_INVALID_IPID", "The requested object or interface does not exist."}
	RpcEInvalidObject                                            = &Error{0x80010114, "RPC_E_INVALID_OBJECT", "The requested object does not exist."}
	RpcSCallpending                                              = &Error{0x80010115, "RPC_S_CALLPENDING", "OLE has sent a request and is waiting for a reply."}
	RpcSWaitontimer                                              = &Error{0x80010116, "RPC_S_WAITONTIMER", "OLE is waiting before retrying a request."}
	RpcECallComplete                                             = &Error{0x80010117, "RPC_E_CALL_COMPLETE", "Call context cannot be accessed after call completed."}
	RpcEUnsecureCall                                             = &Error{0x80010118, "RPC_E_UNSECURE_CALL", "Impersonate on unsecure calls is not supported."}
	RpcETooLate                                                  = &Error{0x80010119, "RPC_E_TOO_LATE", "Security must be initialized before any interfaces are marshaled or unmarshaled. It cannot be changed after initialized."}
	RpcENoGoodSecurityPackages                                   = &Error{0x8001011A, "RPC_E_NO_GOOD_SECURITY_PACKAGES", "No security packages are installed on this machine, the user is not logged on, or there are no compatible security packages between the client and server."}
	RpcEAccessDenied                                             = &Error{0x8001011B, "RPC_E_ACCESS_DENIED", "Access is denied."}
	RpcERemoteDisabled                                           = &Error{0x8001011C, "RPC_E_REMOTE_DISABLED", "Remote calls are not allowed for this process."}
	RpcEInvalidObjref                                            = &Error{0x8001011D, "RPC_E_INVALID_OBJREF", "The marshaled interface data packet (OBJREF) has an invalid or unknown format."}
	RpcENoContext                                                = &Error{0x8001011E, "RPC_E_NO_CONTEXT", "No context is associated with this call. This happens for some custom marshaled calls and on the client side of the call."}
	RpcETimeout                                                  = &Error{0x8001011F, "RPC_E_TIMEOUT", "This operation returned because the time-out period expired."}
	RpcENoSync                                                   = &Error{0x80010120, "RPC_E_NO_SYNC", "There are no synchronize objects to wait on."}
	RpcEFullsicRequired                                          = &Error{0x80010121, "RPC_E_FULLSIC_REQUIRED", "Full subject issuer chain Secure Sockets Layer (SSL) principal name expected from the server."}
	RpcEInvalidStdName                                           = &Error{0x80010122, "RPC_E_INVALID_STD_NAME", "Principal name is not a valid Microsoft standard (msstd) name."}
	CoEFailedtoimpersonate                                       = &Error{0x80010123, "CO_E_FAILEDTOIMPERSONATE", "Unable to impersonate DCOM client."}
	CoEFailedtogetsecctx                                         = &Error{0x80010124, "CO_E_FAILEDTOGETSECCTX", "Unable to obtain server's security context."}
	CoEFailedtoopenthreadtoken                                   = &Error{0x80010125, "CO_E_FAILEDTOOPENTHREADTOKEN", "Unable to open the access token of the current thread."}
	CoEFailedtogettokeninfo                                      = &Error{0x80010126, "CO_E_FAILEDTOGETTOKENINFO", "Unable to obtain user information from an access token."}
	CoETrusteedoesntmatchclient                                  = &Error{0x80010127, "CO_E_TRUSTEEDOESNTMATCHCLIENT", "The client who called IAccessControl::IsAccessPermitted was not the trustee provided to the method."}
	CoEFailedtoqueryclientblanket                                = &Error{0x80010128, "CO_E_FAILEDTOQUERYCLIENTBLANKET", "Unable to obtain the client's security blanket."}
	CoEFailedtosetdacl                                           = &Error{0x80010129, "CO_E_FAILEDTOSETDACL", "Unable to set a discretionary access control list (ACL) into a security descriptor."}
	CoEAccesscheckfailed                                         = &Error{0x8001012A, "CO_E_ACCESSCHECKFAILED", "The system function AccessCheck returned false."}
	CoENetaccessapifailed                                        = &Error{0x8001012B, "CO_E_NETACCESSAPIFAILED", "Either NetAccessDel or NetAccessAdd returned an error code."}
	CoEWrongtrusteenamesyntax                                    = &Error{0x8001012C, "CO_E_WRONGTRUSTEENAMESYNTAX", "One of the trustee strings provided by the user did not conform to the <Domain>\\<Name> syntax and it was not the *\" string\"."}
	CoEInvalidsid                                                = &Error{0x8001012D, "CO_E_INVALIDSID", "One of the security identifiers provided by the user was invalid."}
	CoEConversionfailed                                          = &Error{0x8001012E, "CO_E_CONVERSIONFAILED", "Unable to convert a wide character trustee string to a multiple-byte trustee string."}
	CoENomatchingsidfound                                        = &Error{0x8001012F, "CO_E_NOMATCHINGSIDFOUND", "Unable to find a security identifier that corresponds to a trustee string provided by the user."}
	CoELookupaccsidfailed                                        = &Error{0x80010130, "CO_E_LOOKUPACCSIDFAILED", "The system function LookupAccountSID failed."}
	CoENomatchingnamefound                                       = &Error{0x80010131, "CO_E_NOMATCHINGNAMEFOUND", "Unable to find a trustee name that corresponds to a security identifier provided by the user."}
	CoELookupaccnamefailed                                       = &Error{0x80010132, "CO_E_LOOKUPACCNAMEFAILED", "The system function LookupAccountName failed."}
	CoESetserlhndlfailed                                         = &Error{0x80010133, "CO_E_SETSERLHNDLFAILED", "Unable to set or reset a serialization handle."}
	CoEFailedtogetwindir                                         = &Error{0x80010134, "CO_E_FAILEDTOGETWINDIR", "Unable to obtain the Windows directory."}
	CoEPathtoolong                                               = &Error{0x80010135, "CO_E_PATHTOOLONG", "Path too long."}
	CoEFailedtogenuuid                                           = &Error{0x80010136, "CO_E_FAILEDTOGENUUID", "Unable to generate a UUID."}
	CoEFailedtocreatefile                                        = &Error{0x80010137, "CO_E_FAILEDTOCREATEFILE", "Unable to create file."}
	CoEFailedtoclosehandle                                       = &Error{0x80010138, "CO_E_FAILEDTOCLOSEHANDLE", "Unable to close a serialization handle or a file handle."}
	CoEExceedsysacllimit                                         = &Error{0x80010139, "CO_E_EXCEEDSYSACLLIMIT", "The number of access control entries (ACEs) in an ACL exceeds the system limit."}
	CoEAcesinwrongorder                                          = &Error{0x8001013A, "CO_E_ACESINWRONGORDER", "Not all the DENY_ACCESS ACEs are arranged in front of the GRANT_ACCESS ACEs in the stream."}
	CoEIncompatiblestreamversion                                 = &Error{0x8001013B, "CO_E_INCOMPATIBLESTREAMVERSION", "The version of ACL format in the stream is not supported by this implementation of IAccessControl."}
	CoEFailedtoopenprocesstoken                                  = &Error{0x8001013C, "CO_E_FAILEDTOOPENPROCESSTOKEN", "Unable to open the access token of the server process."}
	CoEDecodefailed                                              = &Error{0x8001013D, "CO_E_DECODEFAILED", "Unable to decode the ACL in the stream provided by the user."}
	CoEAcnotinitialized                                          = &Error{0x8001013F, "CO_E_ACNOTINITIALIZED", "The COM IAccessControl object is not initialized."}
	CoECancelDisabled                                            = &Error{0x80010140, "CO_E_CANCEL_DISABLED", "Call Cancellation is disabled."}
	RpcEUnexpected                                               = &Error{0x8001FFFF, "RPC_E_UNEXPECTED", "An internal error occurred."}
	DispEUnknowninterface                                        = &Error{0x80020001, "DISP_E_UNKNOWNINTERFACE", "Unknown interface."}
	DispEMembernotfound                                          = &Error{0x80020003, "DISP_E_MEMBERNOTFOUND", "Member not found."}
	DispEParamnotfound                                           = &Error{0x80020004, "DISP_E_PARAMNOTFOUND", "Parameter not found."}
	DispETypemismatch                                            = &Error{0x80020005, "DISP_E_TYPEMISMATCH", "Type mismatch."}
	DispEUnknownname                                             = &Error{0x80020006, "DISP_E_UNKNOWNNAME", "Unknown name."}
	DispENonamedargs                                             = &Error{0x80020007, "DISP_E_NONAMEDARGS", "No named arguments."}
	DispEBadvartype                                              = &Error{0x80020008, "DISP_E_BADVARTYPE", "Bad variable type."}
	DispEException                                               = &Error{0x80020009, "DISP_E_EXCEPTION", "Exception occurred."}
	DispEOverflow                                                = &Error{0x8002000A, "DISP_E_OVERFLOW", "Out of present range."}
	DispEBadindex                                                = &Error{0x8002000B, "DISP_E_BADINDEX", "Invalid index."}
	DispEUnknownlcid                                             = &Error{0x8002000C, "DISP_E_UNKNOWNLCID", "Unknown language."}
	DispEArrayislocked                                           = &Error{0x8002000D, "DISP_E_ARRAYISLOCKED", "Memory is locked."}
	DispEBadparamcount                                           = &Error{0x8002000E, "DISP_E_BADPARAMCOUNT", "Invalid number of parameters."}
	DispEParamnotoptional                                        = &Error{0x8002000F, "DISP_E_PARAMNOTOPTIONAL", "Parameter not optional."}
	DispEBadcallee                                               = &Error{0x80020010, "DISP_E_BADCALLEE", "Invalid callee."}
	DispENotacollection                                          = &Error{0x80020011, "DISP_E_NOTACOLLECTION", "Does not support a collection."}
	DispEDivbyzero                                               = &Error{0x80020012, "DISP_E_DIVBYZERO", "Division by zero."}
	DispEBuffertoosmall                                          = &Error{0x80020013, "DISP_E_BUFFERTOOSMALL", "Buffer too small."}
	TypeEBuffertoosmall                                          = &Error{0x80028016, "TYPE_E_BUFFERTOOSMALL", "Buffer too small."}
	TypeEFieldnotfound                                           = &Error{0x80028017, "TYPE_E_FIELDNOTFOUND", "Field name not defined in the record."}
	TypeEInvdataread                                             = &Error{0x80028018, "TYPE_E_INVDATAREAD", "Old format or invalid type library."}
	TypeEUnsupformat                                             = &Error{0x80028019, "TYPE_E_UNSUPFORMAT", "Old format or invalid type library."}
	TypeERegistryaccess                                          = &Error{0x8002801C, "TYPE_E_REGISTRYACCESS", "Error accessing the OLE registry."}
	TypeELibnotregistered                                        = &Error{0x8002801D, "TYPE_E_LIBNOTREGISTERED", "Library not registered."}
	TypeEUndefinedtype                                           = &Error{0x80028027, "TYPE_E_UNDEFINEDTYPE", "Bound to unknown type."}
	TypeEQualifiednamedisallowed                                 = &Error{0x80028028, "TYPE_E_QUALIFIEDNAMEDISALLOWED", "Qualified name disallowed."}
	TypeEInvalidstate                                            = &Error{0x80028029, "TYPE_E_INVALIDSTATE", "Invalid forward reference, or reference to uncompiled type."}
	TypeEWrongtypekind                                           = &Error{0x8002802A, "TYPE_E_WRONGTYPEKIND", "Type mismatch."}
	TypeEElementnotfound                                         = &Error{0x8002802B, "TYPE_E_ELEMENTNOTFOUND", "Element not found."}
	TypeEAmbiguousname                                           = &Error{0x8002802C, "TYPE_E_AMBIGUOUSNAME", "Ambiguous name."}
	TypeENameconflict                                            = &Error{0x8002802D, "TYPE_E_NAMECONFLICT", "Name already exists in the library."}
	TypeEUnknownlcid                                             = &Error{0x8002802E, "TYPE_E_UNKNOWNLCID", "Unknown language code identifier (LCID)."}
	TypeEDllfunctionnotfound                                     = &Error{0x8002802F, "TYPE_E_DLLFUNCTIONNOTFOUND", "Function not defined in specified DLL."}
	TypeEBadmodulekind                                           = &Error{0x800288BD, "TYPE_E_BADMODULEKIND", "Wrong module kind for the operation."}
	TypeESizetoobig                                              = &Error{0x800288C5, "TYPE_E_SIZETOOBIG", "Size cannot exceed 64 KB."}
	TypeEDuplicateid                                             = &Error{0x800288C6, "TYPE_E_DUPLICATEID", "Duplicate ID in inheritance hierarchy."}
	TypeEInvalidid                                               = &Error{0x800288CF, "TYPE_E_INVALIDID", "Incorrect inheritance depth in standard OLE hmember."}
	TypeETypemismatch                                            = &Error{0x80028CA0, "TYPE_E_TYPEMISMATCH", "Type mismatch."}
	TypeEOutofbounds                                             = &Error{0x80028CA1, "TYPE_E_OUTOFBOUNDS", "Invalid number of arguments."}
	TypeEIoerror                                                 = &Error{0x80028CA2, "TYPE_E_IOERROR", "I/O error."}
	TypeECantcreatetmpfile                                       = &Error{0x80028CA3, "TYPE_E_CANTCREATETMPFILE", "Error creating unique .tmp file."}
	TypeECantloadlibrary                                         = &Error{0x80029C4A, "TYPE_E_CANTLOADLIBRARY", "Error loading type library or DLL."}
	TypeEInconsistentpropfuncs                                   = &Error{0x80029C83, "TYPE_E_INCONSISTENTPROPFUNCS", "Inconsistent property functions."}
	TypeECirculartype                                            = &Error{0x80029C84, "TYPE_E_CIRCULARTYPE", "Circular dependency between types and modules."}
	StgEInvalidfunction                                          = &Error{0x80030001, "STG_E_INVALIDFUNCTION", "Unable to perform requested operation."}
	StgEFilenotfound                                             = &Error{0x80030002, "STG_E_FILENOTFOUND", "%1 could not be found."}
	StgEPathnotfound                                             = &Error{0x80030003, "STG_E_PATHNOTFOUND", "The path %1 could not be found."}
	StgEToomanyopenfiles                                         = &Error{0x80030004, "STG_E_TOOMANYOPENFILES", "There are insufficient resources to open another file."}
	StgEAccessdenied                                             = &Error{0x80030005, "STG_E_ACCESSDENIED", "Access denied."}
	StgEInvalidhandle                                            = &Error{0x80030006, "STG_E_INVALIDHANDLE", "Attempted an operation on an invalid object."}
	StgEInsufficientmemory                                       = &Error{0x80030008, "STG_E_INSUFFICIENTMEMORY", "There is insufficient memory available to complete operation."}
	StgEInvalidpointer                                           = &Error{0x80030009, "STG_E_INVALIDPOINTER", "Invalid pointer error."}
	StgENomorefiles                                              = &Error{0x80030012, "STG_E_NOMOREFILES", "There are no more entries to return."}
	StgEDiskiswriteprotected                                     = &Error{0x80030013, "STG_E_DISKISWRITEPROTECTED", "Disk is write-protected."}
	StgESeekerror                                                = &Error{0x80030019, "STG_E_SEEKERROR", "An error occurred during a seek operation."}
	StgEWritefault                                               = &Error{0x8003001D, "STG_E_WRITEFAULT", "A disk error occurred during a write operation."}
	StgEReadfault                                                = &Error{0x8003001E, "STG_E_READFAULT", "A disk error occurred during a read operation."}
	StgEShareviolation                                           = &Error{0x80030020, "STG_E_SHAREVIOLATION", "A share violation has occurred."}
	StgELockviolation                                            = &Error{0x80030021, "STG_E_LOCKVIOLATION", "A lock violation has occurred."}
	StgEFilealreadyexists                                        = &Error{0x80030050, "STG_E_FILEALREADYEXISTS", "%1 already exists."}
	StgEInvalidparameter                                         = &Error{0x80030057, "STG_E_INVALIDPARAMETER", "Invalid parameter error."}
	StgEMediumfull                                               = &Error{0x80030070, "STG_E_MEDIUMFULL", "There is insufficient disk space to complete operation."}
	StgEPropsetmismatched                                        = &Error{0x800300F0, "STG_E_PROPSETMISMATCHED", "Illegal write of non-simple property to simple property set."}
	StgEAbnormalapiexit                                          = &Error{0x800300FA, "STG_E_ABNORMALAPIEXIT", "An application programming interface (API) call exited abnormally."}
	StgEInvalidheader                                            = &Error{0x800300FB, "STG_E_INVALIDHEADER", "The file %1 is not a valid compound file."}
	StgEInvalidname                                              = &Error{0x800300FC, "STG_E_INVALIDNAME", "The name %1 is not valid."}
	StgEUnknown                                                  = &Error{0x800300FD, "STG_E_UNKNOWN", "An unexpected error occurred."}
	StgEUnimplementedfunction                                    = &Error{0x800300FE, "STG_E_UNIMPLEMENTEDFUNCTION", "That function is not implemented."}
	StgEInvalidflag                                              = &Error{0x800300FF, "STG_E_INVALIDFLAG", "Invalid flag error."}
	StgEInuse                                                    = &Error{0x80030100, "STG_E_INUSE", "Attempted to use an object that is busy."}
	StgENotcurrent                                               = &Error{0x80030101, "STG_E_NOTCURRENT", "The storage has been changed since the last commit."}
	StgEReverted                                                 = &Error{0x80030102, "STG_E_REVERTED", "Attempted to use an object that has ceased to exist."}
	StgECantsave                                                 = &Error{0x80030103, "STG_E_CANTSAVE", "Cannot save."}
	StgEOldformat                                                = &Error{0x80030104, "STG_E_OLDFORMAT", "The compound file %1 was produced with an incompatible version of storage."}
	StgEOlddll                                                   = &Error{0x80030105, "STG_E_OLDDLL", "The compound file %1 was produced with a newer version of storage."}
	StgESharerequired                                            = &Error{0x80030106, "STG_E_SHAREREQUIRED", "Share.exe or equivalent is required for operation."}
	StgENotfilebasedstorage                                      = &Error{0x80030107, "STG_E_NOTFILEBASEDSTORAGE", "Illegal operation called on non-file based storage."}
	StgEExtantmarshallings                                       = &Error{0x80030108, "STG_E_EXTANTMARSHALLINGS", "Illegal operation called on object with extant marshalings."}
	StgEDocfilecorrupt                                           = &Error{0x80030109, "STG_E_DOCFILECORRUPT", "The docfile has been corrupted."}
	StgEBadbaseaddress                                           = &Error{0x80030110, "STG_E_BADBASEADDRESS", "OLE32.DLL has been loaded at the wrong address."}
	StgEDocfiletoolarge                                          = &Error{0x80030111, "STG_E_DOCFILETOOLARGE", "The compound file is too large for the current implementation."}
	StgENotsimpleformat                                          = &Error{0x80030112, "STG_E_NOTSIMPLEFORMAT", "The compound file was not created with the STGM_SIMPLE flag."}
	StgEIncomplete                                               = &Error{0x80030201, "STG_E_INCOMPLETE", "The file download was aborted abnormally. The file is incomplete."}
	StgETerminated                                               = &Error{0x80030202, "STG_E_TERMINATED", "The file download has been terminated."}
	StgEStatusCopyProtectionFailure                              = &Error{0x80030305, "STG_E_STATUS_COPY_PROTECTION_FAILURE", "Generic Copy Protection Error."}
	StgECssAuthenticationFailure                                 = &Error{0x80030306, "STG_E_CSS_AUTHENTICATION_FAILURE", "Copy Protection Error—DVD CSS Authentication failed."}
	StgECssKeyNotPresent                                         = &Error{0x80030307, "STG_E_CSS_KEY_NOT_PRESENT", "Copy Protection Error—The given sector does not have a valid CSS key."}
	StgECssKeyNotEstablished                                     = &Error{0x80030308, "STG_E_CSS_KEY_NOT_ESTABLISHED", "Copy Protection Error—DVD session key not established."}
	StgECssScrambledSector                                       = &Error{0x80030309, "STG_E_CSS_SCRAMBLED_SECTOR", "Copy Protection Error—The read failed because the sector is encrypted."}
	StgECssRegionMismatch                                        = &Error{0x8003030A, "STG_E_CSS_REGION_MISMATCH", "Copy Protection Error—The current DVD's region does not correspond to the region setting of the drive."}
	StgEResetsExhausted                                          = &Error{0x8003030B, "STG_E_RESETS_EXHAUSTED", "Copy Protection Error—The drive's region setting might be permanent or the number of user resets has been exhausted."}
	OleEOleverb                                                  = &Error{0x80040000, "OLE_E_OLEVERB", "Invalid OLEVERB structure."}
	OleEAdvf                                                     = &Error{0x80040001, "OLE_E_ADVF", "Invalid advise flags."}
	OleEEnumNomore                                               = &Error{0x80040002, "OLE_E_ENUM_NOMORE", "Cannot enumerate any more because the associated data is missing."}
	OleEAdvisenotsupported                                       = &Error{0x80040003, "OLE_E_ADVISENOTSUPPORTED", "This implementation does not take advises."}
	OleENoconnection                                             = &Error{0x80040004, "OLE_E_NOCONNECTION", "There is no connection for this connection ID."}
	OleENotrunning                                               = &Error{0x80040005, "OLE_E_NOTRUNNING", "Need to run the object to perform this operation."}
	OleENocache                                                  = &Error{0x80040006, "OLE_E_NOCACHE", "There is no cache to operate on."}
	OleEBlank                                                    = &Error{0x80040007, "OLE_E_BLANK", "Uninitialized object."}
	OleEClassdiff                                                = &Error{0x80040008, "OLE_E_CLASSDIFF", "Linked object's source class has changed."}
	OleECantGetmoniker                                           = &Error{0x80040009, "OLE_E_CANT_GETMONIKER", "Not able to get the moniker of the object."}
	OleECantBindtosource                                         = &Error{0x8004000A, "OLE_E_CANT_BINDTOSOURCE", "Not able to bind to the source."}
	OleEStatic                                                   = &Error{0x8004000B, "OLE_E_STATIC", "Object is static; operation not allowed."}
	OleEPromptsavecancelled                                      = &Error{0x8004000C, "OLE_E_PROMPTSAVECANCELLED", "User canceled out of the Save dialog box."}
	OleEInvalidrect                                              = &Error{0x8004000D, "OLE_E_INVALIDRECT", "Invalid rectangle."}
	OleEWrongcompobj                                             = &Error{0x8004000E, "OLE_E_WRONGCOMPOBJ", "compobj.dll is too old for the ole2.dll initialized."}
	OleEInvalidhwnd                                              = &Error{0x8004000F, "OLE_E_INVALIDHWND", "Invalid window handle."}
	OleENotInplaceactive                                         = &Error{0x80040010, "OLE_E_NOT_INPLACEACTIVE", "Object is not in any of the inplace active states."}
	OleECantconvert                                              = &Error{0x80040011, "OLE_E_CANTCONVERT", "Not able to convert object."}
	OleENostorage                                                = &Error{0x80040012, "OLE_E_NOSTORAGE", "Not able to perform the operation because object is not given storage yet."}
	DvEFormatetc                                                 = &Error{0x80040064, "DV_E_FORMATETC", "Invalid FORMATETC structure."}
	DvEDvtargetdevice                                            = &Error{0x80040065, "DV_E_DVTARGETDEVICE", "Invalid DVTARGETDEVICE structure."}
	DvEStgmedium                                                 = &Error{0x80040066, "DV_E_STGMEDIUM", "Invalid STDGMEDIUM structure."}
	DvEStatdata                                                  = &Error{0x80040067, "DV_E_STATDATA", "Invalid STATDATA structure."}
	DvELindex                                                    = &Error{0x80040068, "DV_E_LINDEX", "Invalid lindex."}
	DvETymed                                                     = &Error{0x80040069, "DV_E_TYMED", "Invalid TYMED structure."}
	DvEClipformat                                                = &Error{0x8004006A, "DV_E_CLIPFORMAT", "Invalid clipboard format."}
	DvEDvaspect                                                  = &Error{0x8004006B, "DV_E_DVASPECT", "Invalid aspects."}
	DvEDvtargetdeviceSize                                        = &Error{0x8004006C, "DV_E_DVTARGETDEVICE_SIZE", "The tdSize parameter of the DVTARGETDEVICE structure is invalid."}
	DvENoiviewobject                                             = &Error{0x8004006D, "DV_E_NOIVIEWOBJECT", "Object does not support IViewObject interface."}
	DragdropENotregistered                                       = &Error{0x80040100, "DRAGDROP_E_NOTREGISTERED", "Trying to revoke a drop target that has not been registered."}
	DragdropEAlreadyregistered                                   = &Error{0x80040101, "DRAGDROP_E_ALREADYREGISTERED", "This window has already been registered as a drop target."}
	DragdropEInvalidhwnd                                         = &Error{0x80040102, "DRAGDROP_E_INVALIDHWND", "Invalid window handle."}
	ClassENoaggregation                                          = &Error{0x80040110, "CLASS_E_NOAGGREGATION", "Class does not support aggregation (or class object is remote)."}
	ClassEClassnotavailable                                      = &Error{0x80040111, "CLASS_E_CLASSNOTAVAILABLE", "ClassFactory cannot supply requested class."}
	ClassENotlicensed                                            = &Error{0x80040112, "CLASS_E_NOTLICENSED", "Class is not licensed for use."}
	ViewEDraw                                                    = &Error{0x80040140, "VIEW_E_DRAW", "Error drawing view."}
	RegdbEReadregdb                                              = &Error{0x80040150, "REGDB_E_READREGDB", "Could not read key from registry."}
	RegdbEWriteregdb                                             = &Error{0x80040151, "REGDB_E_WRITEREGDB", "Could not write key to registry."}
	RegdbEKeymissing                                             = &Error{0x80040152, "REGDB_E_KEYMISSING", "Could not find the key in the registry."}
	RegdbEInvalidvalue                                           = &Error{0x80040153, "REGDB_E_INVALIDVALUE", "Invalid value for registry."}
	RegdbEClassnotreg                                            = &Error{0x80040154, "REGDB_E_CLASSNOTREG", "Class not registered."}
	RegdbEIidnotreg                                              = &Error{0x80040155, "REGDB_E_IIDNOTREG", "Interface not registered."}
	RegdbEBadthreadingmodel                                      = &Error{0x80040156, "REGDB_E_BADTHREADINGMODEL", "Threading model entry is not valid."}
	CatECatidnoexist                                             = &Error{0x80040160, "CAT_E_CATIDNOEXIST", "CATID does not exist."}
	CatENodescription                                            = &Error{0x80040161, "CAT_E_NODESCRIPTION", "Description not found."}
	CsEPackageNotfound                                           = &Error{0x80040164, "CS_E_PACKAGE_NOTFOUND", "No package in the software installation data in Active Directory meets this criteria."}
	CsENotDeletable                                              = &Error{0x80040165, "CS_E_NOT_DELETABLE", "Deleting this will break the referential integrity of the software installation data in Active Directory."}
	CsEClassNotfound                                             = &Error{0x80040166, "CS_E_CLASS_NOTFOUND", "The CLSID was not found in the software installation data in Active Directory."}
	CsEInvalidVersion                                            = &Error{0x80040167, "CS_E_INVALID_VERSION", "The software installation data in Active Directory is corrupt."}
	CsENoClassstore                                              = &Error{0x80040168, "CS_E_NO_CLASSSTORE", "There is no software installation data in Active Directory."}
	CsEObjectNotfound                                            = &Error{0x80040169, "CS_E_OBJECT_NOTFOUND", "There is no software installation data object in Active Directory."}
	CsEObjectAlreadyExists                                       = &Error{0x8004016A, "CS_E_OBJECT_ALREADY_EXISTS", "The software installation data object in Active Directory already exists."}
	CsEInvalidPath                                               = &Error{0x8004016B, "CS_E_INVALID_PATH", "The path to the software installation data in Active Directory is not correct."}
	CsENetworkError                                              = &Error{0x8004016C, "CS_E_NETWORK_ERROR", "A network error interrupted the operation."}
	CsEAdminLimitExceeded                                        = &Error{0x8004016D, "CS_E_ADMIN_LIMIT_EXCEEDED", "The size of this object exceeds the maximum size set by the administrator."}
	CsESchemaMismatch                                            = &Error{0x8004016E, "CS_E_SCHEMA_MISMATCH", "The schema for the software installation data in Active Directory does not match the required schema."}
	CsEInternalError                                             = &Error{0x8004016F, "CS_E_INTERNAL_ERROR", "An error occurred in the software installation data in Active Directory."}
	CacheENocacheUpdated                                         = &Error{0x80040170, "CACHE_E_NOCACHE_UPDATED", "Cache not updated."}
	OleobjENoverbs                                               = &Error{0x80040180, "OLEOBJ_E_NOVERBS", "No verbs for OLE object."}
	OleobjEInvalidverb                                           = &Error{0x80040181, "OLEOBJ_E_INVALIDVERB", "Invalid verb for OLE object."}
	InplaceENotundoable                                          = &Error{0x800401A0, "INPLACE_E_NOTUNDOABLE", "Undo is not available."}
	InplaceENotoolspace                                          = &Error{0x800401A1, "INPLACE_E_NOTOOLSPACE", "Space for tools is not available."}
	Convert10EOlestreamGet                                       = &Error{0x800401C0, "CONVERT10_E_OLESTREAM_GET", "OLESTREAM Get method failed."}
	Convert10EOlestreamPut                                       = &Error{0x800401C1, "CONVERT10_E_OLESTREAM_PUT", "OLESTREAM Put method failed."}
	Convert10EOlestreamFmt                                       = &Error{0x800401C2, "CONVERT10_E_OLESTREAM_FMT", "Contents of the OLESTREAM not in correct format."}
	Convert10EOlestreamBitmapToDib                               = &Error{0x800401C3, "CONVERT10_E_OLESTREAM_BITMAP_TO_DIB", "There was an error in a Windows GDI call while converting the bitmap to a device-independent bitmap (DIB)."}
	Convert10EStgFmt                                             = &Error{0x800401C4, "CONVERT10_E_STG_FMT", "Contents of the IStorage not in correct format."}
	Convert10EStgNoStdStream                                     = &Error{0x800401C5, "CONVERT10_E_STG_NO_STD_STREAM", "Contents of IStorage is missing one of the standard streams."}
	Convert10EStgDibToBitmap                                     = &Error{0x800401C6, "CONVERT10_E_STG_DIB_TO_BITMAP", "There was an error in a Windows Graphics Device Interface (GDI) call while converting the DIB to a bitmap."}
	ClipbrdECantOpen                                             = &Error{0x800401D0, "CLIPBRD_E_CANT_OPEN", "OpenClipboard failed."}
	ClipbrdECantEmpty                                            = &Error{0x800401D1, "CLIPBRD_E_CANT_EMPTY", "EmptyClipboard failed."}
	ClipbrdECantSet                                              = &Error{0x800401D2, "CLIPBRD_E_CANT_SET", "SetClipboard failed."}
	ClipbrdEBadData                                              = &Error{0x800401D3, "CLIPBRD_E_BAD_DATA", "Data on clipboard is invalid."}
	ClipbrdECantClose                                            = &Error{0x800401D4, "CLIPBRD_E_CANT_CLOSE", "CloseClipboard failed."}
	MkEConnectmanually                                           = &Error{0x800401E0, "MK_E_CONNECTMANUALLY", "Moniker needs to be connected manually."}
	MkEExceededdeadline                                          = &Error{0x800401E1, "MK_E_EXCEEDEDDEADLINE", "Operation exceeded deadline."}
	MkENeedgeneric                                               = &Error{0x800401E2, "MK_E_NEEDGENERIC", "Moniker needs to be generic."}
	MkEUnavailable                                               = &Error{0x800401E3, "MK_E_UNAVAILABLE", "Operation unavailable."}
	MkESyntax                                                    = &Error{0x800401E4, "MK_E_SYNTAX", "Invalid syntax."}
	MkENoobject                                                  = &Error{0x800401E5, "MK_E_NOOBJECT", "No object for moniker."}
	MkEInvalidextension                                          = &Error{0x800401E6, "MK_E_INVALIDEXTENSION", "Bad extension for file."}
	MkEIntermediateinterfacenotsupported                         = &Error{0x800401E7, "MK_E_INTERMEDIATEINTERFACENOTSUPPORTED", "Intermediate operation failed."}
	MkENotbindable                                               = &Error{0x800401E8, "MK_E_NOTBINDABLE", "Moniker is not bindable."}
	MkENotbound                                                  = &Error{0x800401E9, "MK_E_NOTBOUND", "Moniker is not bound."}
	MkECantopenfile                                              = &Error{0x800401EA, "MK_E_CANTOPENFILE", "Moniker cannot open file."}
	MkEMustbotheruser                                            = &Error{0x800401EB, "MK_E_MUSTBOTHERUSER", "User input required for operation to succeed."}
	MkENoinverse                                                 = &Error{0x800401EC, "MK_E_NOINVERSE", "Moniker class has no inverse."}
	MkENostorage                                                 = &Error{0x800401ED, "MK_E_NOSTORAGE", "Moniker does not refer to storage."}
	MkENoprefix                                                  = &Error{0x800401EE, "MK_E_NOPREFIX", "No common prefix."}
	MkEEnumerationFailed                                         = &Error{0x800401EF, "MK_E_ENUMERATION_FAILED", "Moniker could not be enumerated."}
	CoENotinitialized                                            = &Error{0x800401F0, "CO_E_NOTINITIALIZED", "CoInitialize has not been called."}
	CoEAlreadyinitialized                                        = &Error{0x800401F1, "CO_E_ALREADYINITIALIZED", "CoInitialize has already been called."}
	CoECantdetermineclass                                        = &Error{0x800401F2, "CO_E_CANTDETERMINECLASS", "Class of object cannot be determined."}
	CoEClassstring                                               = &Error{0x800401F3, "CO_E_CLASSSTRING", "Invalid class string."}
	CoEIidstring                                                 = &Error{0x800401F4, "CO_E_IIDSTRING", "Invalid interface string."}
	CoEAppnotfound                                               = &Error{0x800401F5, "CO_E_APPNOTFOUND", "Application not found."}
	CoEAppsingleuse                                              = &Error{0x800401F6, "CO_E_APPSINGLEUSE", "Application cannot be run more than once."}
	CoEErrorinapp                                                = &Error{0x800401F7, "CO_E_ERRORINAPP", "Some error in application."}
	CoEDllnotfound                                               = &Error{0x800401F8, "CO_E_DLLNOTFOUND", "DLL for class not found."}
	CoEErrorindll                                                = &Error{0x800401F9, "CO_E_ERRORINDLL", "Error in the DLL."}
	CoEWrongosforapp                                             = &Error{0x800401FA, "CO_E_WRONGOSFORAPP", "Wrong operating system or operating system version for application."}
	CoEObjnotreg                                                 = &Error{0x800401FB, "CO_E_OBJNOTREG", "Object is not registered."}
	CoEObjisreg                                                  = &Error{0x800401FC, "CO_E_OBJISREG", "Object is already registered."}
	CoEObjnotconnected                                           = &Error{0x800401FD, "CO_E_OBJNOTCONNECTED", "Object is not connected to server."}
	CoEAppdidntreg                                               = &Error{0x800401FE, "CO_E_APPDIDNTREG", "Application was launched, but it did not register a class factory."}
	CoEReleased                                                  = &Error{0x800401FF, "CO_E_RELEASED", "Object has been released."}
	EventEAllSubscribersFailed                                   = &Error{0x80040201, "EVENT_E_ALL_SUBSCRIBERS_FAILED", "An event was unable to invoke any of the subscribers."}
	EventEQuerysyntax                                            = &Error{0x80040203, "EVENT_E_QUERYSYNTAX", "A syntax error occurred trying to evaluate a query string."}
	EventEQueryfield                                             = &Error{0x80040204, "EVENT_E_QUERYFIELD", "An invalid field name was used in a query string."}
	EventEInternalexception                                      = &Error{0x80040205, "EVENT_E_INTERNALEXCEPTION", "An unexpected exception was raised."}
	EventEInternalerror                                          = &Error{0x80040206, "EVENT_E_INTERNALERROR", "An unexpected internal error was detected."}
	EventEInvalidPerUserSid                                      = &Error{0x80040207, "EVENT_E_INVALID_PER_USER_SID", "The owner security identifier (SID) on a per-user subscription does not exist."}
	EventEUserException                                          = &Error{0x80040208, "EVENT_E_USER_EXCEPTION", "A user-supplied component or subscriber raised an exception."}
	EventETooManyMethods                                         = &Error{0x80040209, "EVENT_E_TOO_MANY_METHODS", "An interface has too many methods to fire events from."}
	EventEMissingEventclass                                      = &Error{0x8004020A, "EVENT_E_MISSING_EVENTCLASS", "A subscription cannot be stored unless its event class already exists."}
	EventENotAllRemoved                                          = &Error{0x8004020B, "EVENT_E_NOT_ALL_REMOVED", "Not all the objects requested could be removed."}
	EventEComplusNotInstalled                                    = &Error{0x8004020C, "EVENT_E_COMPLUS_NOT_INSTALLED", "COM+ is required for this operation, but it is not installed."}
	EventECantModifyOrDeleteUnconfiguredObject                   = &Error{0x8004020D, "EVENT_E_CANT_MODIFY_OR_DELETE_UNCONFIGURED_OBJECT", "Cannot modify or delete an object that was not added using the COM+ Administrative SDK."}
	EventECantModifyOrDeleteConfiguredObject                     = &Error{0x8004020E, "EVENT_E_CANT_MODIFY_OR_DELETE_CONFIGURED_OBJECT", "Cannot modify or delete an object that was added using the COM+ Administrative SDK."}
	EventEInvalidEventClassPartition                             = &Error{0x8004020F, "EVENT_E_INVALID_EVENT_CLASS_PARTITION", "The event class for this subscription is in an invalid partition."}
	EventEPerUserSidNotLoggedOn                                  = &Error{0x80040210, "EVENT_E_PER_USER_SID_NOT_LOGGED_ON", "The owner of the PerUser subscription is not logged on to the system specified."}
	SchedETriggerNotFound                                        = &Error{0x80041309, "SCHED_E_TRIGGER_NOT_FOUND", "Trigger not found."}
	SchedETaskNotReady                                           = &Error{0x8004130A, "SCHED_E_TASK_NOT_READY", "One or more of the properties that are needed to run this task have not been set."}
	SchedETaskNotRunning                                         = &Error{0x8004130B, "SCHED_E_TASK_NOT_RUNNING", "There is no running instance of the task."}
	SchedEServiceNotInstalled                                    = &Error{0x8004130C, "SCHED_E_SERVICE_NOT_INSTALLED", "The Task Scheduler service is not installed on this computer."}
	SchedECannotOpenTask                                         = &Error{0x8004130D, "SCHED_E_CANNOT_OPEN_TASK", "The task object could not be opened."}
	SchedEInvalidTask                                            = &Error{0x8004130E, "SCHED_E_INVALID_TASK", "The object is either an invalid task object or is not a task object."}
	SchedEAccountInformationNotSet                               = &Error{0x8004130F, "SCHED_E_ACCOUNT_INFORMATION_NOT_SET", "No account information could be found in the Task Scheduler security database for the task indicated."}
	SchedEAccountNameNotFound                                    = &Error{0x80041310, "SCHED_E_ACCOUNT_NAME_NOT_FOUND", "Unable to establish existence of the account specified."}
	SchedEAccountDbaseCorrupt                                    = &Error{0x80041311, "SCHED_E_ACCOUNT_DBASE_CORRUPT", "Corruption was detected in the Task Scheduler security database; the database has been reset."}
	SchedENoSecurityServices                                     = &Error{0x80041312, "SCHED_E_NO_SECURITY_SERVICES", "Task Scheduler security services are available only on Windows NT operating system."}
	SchedEUnknownObjectVersion                                   = &Error{0x80041313, "SCHED_E_UNKNOWN_OBJECT_VERSION", "The task object version is either unsupported or invalid."}
	SchedEUnsupportedAccountOption                               = &Error{0x80041314, "SCHED_E_UNSUPPORTED_ACCOUNT_OPTION", "The task has been configured with an unsupported combination of account settings and run-time options."}
	SchedEServiceNotRunning                                      = &Error{0x80041315, "SCHED_E_SERVICE_NOT_RUNNING", "The Task Scheduler service is not running."}
	SchedEUnexpectednode                                         = &Error{0x80041316, "SCHED_E_UNEXPECTEDNODE", "The task XML contains an unexpected node."}
	SchedENamespace                                              = &Error{0x80041317, "SCHED_E_NAMESPACE", "The task XML contains an element or attribute from an unexpected namespace."}
	SchedEInvalidvalue                                           = &Error{0x80041318, "SCHED_E_INVALIDVALUE", "The task XML contains a value that is incorrectly formatted or out of range."}
	SchedEMissingnode                                            = &Error{0x80041319, "SCHED_E_MISSINGNODE", "The task XML is missing a required element or attribute."}
	SchedEMalformedxml                                           = &Error{0x8004131A, "SCHED_E_MALFORMEDXML", "The task XML is malformed."}
	SchedETooManyNodes                                           = &Error{0x8004131D, "SCHED_E_TOO_MANY_NODES", "The task XML contains too many nodes of the same type."}
	SchedEPastEndBoundary                                        = &Error{0x8004131E, "SCHED_E_PAST_END_BOUNDARY", "The task cannot be started after the trigger's end boundary."}
	SchedEAlreadyRunning                                         = &Error{0x8004131F, "SCHED_E_ALREADY_RUNNING", "An instance of this task is already running."}
	SchedEUserNotLoggedOn                                        = &Error{0x80041320, "SCHED_E_USER_NOT_LOGGED_ON", "The task will not run because the user is not logged on."}
	SchedEInvalidTaskHash                                        = &Error{0x80041321, "SCHED_E_INVALID_TASK_HASH", "The task image is corrupt or has been tampered with."}
	SchedEServiceNotAvailable                                    = &Error{0x80041322, "SCHED_E_SERVICE_NOT_AVAILABLE", "The Task Scheduler service is not available."}
	SchedEServiceTooBusy                                         = &Error{0x80041323, "SCHED_E_SERVICE_TOO_BUSY", "The Task Scheduler service is too busy to handle your request. Try again later."}
	SchedETaskAttempted                                          = &Error{0x80041324, "SCHED_E_TASK_ATTEMPTED", "The Task Scheduler service attempted to run the task, but the task did not run due to one of the constraints in the task definition."}
	XactEAlreadyothersinglephase                                 = &Error{0x8004D000, "XACT_E_ALREADYOTHERSINGLEPHASE", "Another single phase resource manager has already been enlisted in this transaction."}
	XactECantretain                                              = &Error{0x8004D001, "XACT_E_CANTRETAIN", "A retaining commit or abort is not supported."}
	XactECommitfailed                                            = &Error{0x8004D002, "XACT_E_COMMITFAILED", "The transaction failed to commit for an unknown reason. The transaction was aborted."}
	XactECommitprevented                                         = &Error{0x8004D003, "XACT_E_COMMITPREVENTED", "Cannot call commit on this transaction object because the calling application did not initiate the transaction."}
	XactEHeuristicabort                                          = &Error{0x8004D004, "XACT_E_HEURISTICABORT", "Instead of committing, the resource heuristically aborted."}
	XactEHeuristiccommit                                         = &Error{0x8004D005, "XACT_E_HEURISTICCOMMIT", "Instead of aborting, the resource heuristically committed."}
	XactEHeuristicdamage                                         = &Error{0x8004D006, "XACT_E_HEURISTICDAMAGE", "Some of the states of the resource were committed while others were aborted, likely because of heuristic decisions."}
	XactEHeuristicdanger                                         = &Error{0x8004D007, "XACT_E_HEURISTICDANGER", "Some of the states of the resource might have been committed while others were aborted, likely because of heuristic decisions."}
	XactEIsolationlevel                                          = &Error{0x8004D008, "XACT_E_ISOLATIONLEVEL", "The requested isolation level is not valid or supported."}
	XactENoasync                                                 = &Error{0x8004D009, "XACT_E_NOASYNC", "The transaction manager does not support an asynchronous operation for this method."}
	XactENoenlist                                                = &Error{0x8004D00A, "XACT_E_NOENLIST", "Unable to enlist in the transaction."}
	XactENoisoretain                                             = &Error{0x8004D00B, "XACT_E_NOISORETAIN", "The requested semantics of retention of isolation across retaining commit and abort boundaries cannot be supported by this transaction implementation, or isoFlags was not equal to 0."}
	XactENoresource                                              = &Error{0x8004D00C, "XACT_E_NORESOURCE", "There is no resource presently associated with this enlistment."}
	XactENotcurrent                                              = &Error{0x8004D00D, "XACT_E_NOTCURRENT", "The transaction failed to commit due to the failure of optimistic concurrency control in at least one of the resource managers."}
	XactENotransaction                                           = &Error{0x8004D00E, "XACT_E_NOTRANSACTION", "The transaction has already been implicitly or explicitly committed or aborted."}
	XactENotsupported                                            = &Error{0x8004D00F, "XACT_E_NOTSUPPORTED", "An invalid combination of flags was specified."}
	XactEUnknownrmgrid                                           = &Error{0x8004D010, "XACT_E_UNKNOWNRMGRID", "The resource manager ID is not associated with this transaction or the transaction manager."}
	XactEWrongstate                                              = &Error{0x8004D011, "XACT_E_WRONGSTATE", "This method was called in the wrong state."}
	XactEWronguow                                                = &Error{0x8004D012, "XACT_E_WRONGUOW", "The indicated unit of work does not match the unit of work expected by the resource manager."}
	XactEXtionexists                                             = &Error{0x8004D013, "XACT_E_XTIONEXISTS", "An enlistment in a transaction already exists."}
	XactENoimportobject                                          = &Error{0x8004D014, "XACT_E_NOIMPORTOBJECT", "An import object for the transaction could not be found."}
	XactEInvalidcookie                                           = &Error{0x8004D015, "XACT_E_INVALIDCOOKIE", "The transaction cookie is invalid."}
	XactEIndoubt                                                 = &Error{0x8004D016, "XACT_E_INDOUBT", "The transaction status is in doubt. A communication failure occurred, or a transaction manager or resource manager has failed."}
	XactENotimeout                                               = &Error{0x8004D017, "XACT_E_NOTIMEOUT", "A time-out was specified, but time-outs are not supported."}
	XactEAlreadyinprogress                                       = &Error{0x8004D018, "XACT_E_ALREADYINPROGRESS", "The requested operation is already in progress for the transaction."}
	XactEAborted                                                 = &Error{0x8004D019, "XACT_E_ABORTED", "The transaction has already been aborted."}
	XactELogfull                                                 = &Error{0x8004D01A, "XACT_E_LOGFULL", "The Transaction Manager returned a log full error."}
	XactETmnotavailable                                          = &Error{0x8004D01B, "XACT_E_TMNOTAVAILABLE", "The transaction manager is not available."}
	XactEConnectionDown                                          = &Error{0x8004D01C, "XACT_E_CONNECTION_DOWN", "A connection with the transaction manager was lost."}
	XactEConnectionDenied                                        = &Error{0x8004D01D, "XACT_E_CONNECTION_DENIED", "A request to establish a connection with the transaction manager was denied."}
	XactEReenlisttimeout                                         = &Error{0x8004D01E, "XACT_E_REENLISTTIMEOUT", "Resource manager reenlistment to determine transaction status timed out."}
	XactETipConnectFailed                                        = &Error{0x8004D01F, "XACT_E_TIP_CONNECT_FAILED", "The transaction manager failed to establish a connection with another Transaction Internet Protocol (TIP) transaction manager."}
	XactETipProtocolError                                        = &Error{0x8004D020, "XACT_E_TIP_PROTOCOL_ERROR", "The transaction manager encountered a protocol error with another TIP transaction manager."}
	XactETipPullFailed                                           = &Error{0x8004D021, "XACT_E_TIP_PULL_FAILED", "The transaction manager could not propagate a transaction from another TIP transaction manager."}
	XactEDestTmnotavailable                                      = &Error{0x8004D022, "XACT_E_DEST_TMNOTAVAILABLE", "The transaction manager on the destination machine is not available."}
	XactETipDisabled                                             = &Error{0x8004D023, "XACT_E_TIP_DISABLED", "The transaction manager has disabled its support for TIP."}
	XactENetworkTxDisabled                                       = &Error{0x8004D024, "XACT_E_NETWORK_TX_DISABLED", "The transaction manager has disabled its support for remote or network transactions."}
	XactEPartnerNetworkTxDisabled                                = &Error{0x8004D025, "XACT_E_PARTNER_NETWORK_TX_DISABLED", "The partner transaction manager has disabled its support for remote or network transactions."}
	XactEXaTxDisabled                                            = &Error{0x8004D026, "XACT_E_XA_TX_DISABLED", "The transaction manager has disabled its support for XA transactions."}
	XactEUnableToReadDtcConfig                                   = &Error{0x8004D027, "XACT_E_UNABLE_TO_READ_DTC_CONFIG", "Microsoft Distributed Transaction Coordinator (MSDTC) was unable to read its configuration information."}
	XactEUnableToLoadDtcProxy                                    = &Error{0x8004D028, "XACT_E_UNABLE_TO_LOAD_DTC_PROXY", "MSDTC was unable to load the DTC proxy DLL."}
	XactEAborting                                                = &Error{0x8004D029, "XACT_E_ABORTING", "The local transaction has aborted."}
	XactEClerknotfound                                           = &Error{0x8004D080, "XACT_E_CLERKNOTFOUND", "The specified CRM clerk was not found. It might have completed before it could be held."}
	XactEClerkexists                                             = &Error{0x8004D081, "XACT_E_CLERKEXISTS", "The specified CRM clerk does not exist."}
	XactERecoveryinprogress                                      = &Error{0x8004D082, "XACT_E_RECOVERYINPROGRESS", "Recovery of the CRM log file is still in progress."}
	XactETransactionclosed                                       = &Error{0x8004D083, "XACT_E_TRANSACTIONCLOSED", "The transaction has completed, and the log records have been discarded from the log file. They are no longer available."}
	XactEInvalidlsn                                              = &Error{0x8004D084, "XACT_E_INVALIDLSN", "lsnToRead is outside of the current limits of the log"}
	XactEReplayrequest                                           = &Error{0x8004D085, "XACT_E_REPLAYREQUEST", "The COM+ Compensating Resource Manager has records it wishes to replay."}
	XactEConnectionRequestDenied                                 = &Error{0x8004D100, "XACT_E_CONNECTION_REQUEST_DENIED", "The request to connect to the specified transaction coordinator was denied."}
	XactEToomanyEnlistments                                      = &Error{0x8004D101, "XACT_E_TOOMANY_ENLISTMENTS", "The maximum number of enlistments for the specified transaction has been reached."}
	XactEDuplicateGuid                                           = &Error{0x8004D102, "XACT_E_DUPLICATE_GUID", "A resource manager with the same identifier is already registered with the specified transaction coordinator."}
	XactENotsinglephase                                          = &Error{0x8004D103, "XACT_E_NOTSINGLEPHASE", "The prepare request given was not eligible for single-phase optimizations."}
	XactERecoveryalreadydone                                     = &Error{0x8004D104, "XACT_E_RECOVERYALREADYDONE", "RecoveryComplete has already been called for the given resource manager."}
	XactEProtocol                                                = &Error{0x8004D105, "XACT_E_PROTOCOL", "The interface call made was incorrect for the current state of the protocol."}
	XactERmFailure                                               = &Error{0x8004D106, "XACT_E_RM_FAILURE", "The xa_open call failed for the XA resource."}
	XactERecoveryFailed                                          = &Error{0x8004D107, "XACT_E_RECOVERY_FAILED", "The xa_recover call failed for the XA resource."}
	XactELuNotFound                                              = &Error{0x8004D108, "XACT_E_LU_NOT_FOUND", "The logical unit of work specified cannot be found."}
	XactEDuplicateLu                                             = &Error{0x8004D109, "XACT_E_DUPLICATE_LU", "The specified logical unit of work already exists."}
	XactELuNotConnected                                          = &Error{0x8004D10A, "XACT_E_LU_NOT_CONNECTED", "Subordinate creation failed. The specified logical unit of work was not connected."}
	XactEDuplicateTransid                                        = &Error{0x8004D10B, "XACT_E_DUPLICATE_TRANSID", "A transaction with the given identifier already exists."}
	XactELuBusy                                                  = &Error{0x8004D10C, "XACT_E_LU_BUSY", "The resource is in use."}
	XactELuNoRecoveryProcess                                     = &Error{0x8004D10D, "XACT_E_LU_NO_RECOVERY_PROCESS", "The LU Recovery process is down."}
	XactELuDown                                                  = &Error{0x8004D10E, "XACT_E_LU_DOWN", "The remote session was lost."}
	XactELuRecovering                                            = &Error{0x8004D10F, "XACT_E_LU_RECOVERING", "The resource is currently recovering."}
	XactELuRecoveryMismatch                                      = &Error{0x8004D110, "XACT_E_LU_RECOVERY_MISMATCH", "There was a mismatch in driving recovery."}
	XactERmUnavailable                                           = &Error{0x8004D111, "XACT_E_RM_UNAVAILABLE", "An error occurred with the XA resource."}
	ContextEAborted                                              = &Error{0x8004E002, "CONTEXT_E_ABORTED", "The root transaction wanted to commit, but the transaction aborted."}
	ContextEAborting                                             = &Error{0x8004E003, "CONTEXT_E_ABORTING", "The COM+ component on which the method call was made has a transaction that has already aborted or is in the process of aborting."}
	ContextENocontext                                            = &Error{0x8004E004, "CONTEXT_E_NOCONTEXT", "There is no Microsoft Transaction Server (MTS) object context."}
	ContextEWouldDeadlock                                        = &Error{0x8004E005, "CONTEXT_E_WOULD_DEADLOCK", "The component is configured to use synchronization, and this method call would cause a deadlock to occur."}
	ContextESynchTimeout                                         = &Error{0x8004E006, "CONTEXT_E_SYNCH_TIMEOUT", "The component is configured to use synchronization, and a thread has timed out waiting to enter the context."}
	ContextEOldref                                               = &Error{0x8004E007, "CONTEXT_E_OLDREF", "You made a method call on a COM+ component that has a transaction that has already committed or aborted."}
	ContextERolenotfound                                         = &Error{0x8004E00C, "CONTEXT_E_ROLENOTFOUND", "The specified role was not configured for the application."}
	ContextETmnotavailable                                       = &Error{0x8004E00F, "CONTEXT_E_TMNOTAVAILABLE", "COM+ was unable to talk to the MSDTC."}
	CoEActivationfailed                                          = &Error{0x8004E021, "CO_E_ACTIVATIONFAILED", "An unexpected error occurred during COM+ activation."}
	CoEActivationfailedEventlogged                               = &Error{0x8004E022, "CO_E_ACTIVATIONFAILED_EVENTLOGGED", "COM+ activation failed. Check the event log for more information."}
	CoEActivationfailedCatalogerror                              = &Error{0x8004E023, "CO_E_ACTIVATIONFAILED_CATALOGERROR", "COM+ activation failed due to a catalog or configuration error."}
	CoEActivationfailedTimeout                                   = &Error{0x8004E024, "CO_E_ACTIVATIONFAILED_TIMEOUT", "COM+ activation failed because the activation could not be completed in the specified amount of time."}
	CoEInitializationfailed                                      = &Error{0x8004E025, "CO_E_INITIALIZATIONFAILED", "COM+ activation failed because an initialization function failed. Check the event log for more information."}
	ContextENojit                                                = &Error{0x8004E026, "CONTEXT_E_NOJIT", "The requested operation requires that just-in-time (JIT) be in the current context, and it is not."}
	ContextENotransaction                                        = &Error{0x8004E027, "CONTEXT_E_NOTRANSACTION", "The requested operation requires that the current context have a transaction, and it does not."}
	CoEThreadingmodelChanged                                     = &Error{0x8004E028, "CO_E_THREADINGMODEL_CHANGED", "The components threading model has changed after install into a COM+ application. Re-install component."}
	CoENoiisintrinsics                                           = &Error{0x8004E029, "CO_E_NOIISINTRINSICS", "Internet Information Services (IIS) intrinsics not available. Start your work with IIS."}
	CoENocookies                                                 = &Error{0x8004E02A, "CO_E_NOCOOKIES", "An attempt to write a cookie failed."}
	CoEDberror                                                   = &Error{0x8004E02B, "CO_E_DBERROR", "An attempt to use a database generated a database-specific error."}
	CoENotpooled                                                 = &Error{0x8004E02C, "CO_E_NOTPOOLED", "The COM+ component you created must use object pooling to work."}
	CoENotconstructed                                            = &Error{0x8004E02D, "CO_E_NOTCONSTRUCTED", "The COM+ component you created must use object construction to work correctly."}
	CoENosynchronization                                         = &Error{0x8004E02E, "CO_E_NOSYNCHRONIZATION", "The COM+ component requires synchronization, and it is not configured for it."}
	CoEIsolevelmismatch                                          = &Error{0x8004E02F, "CO_E_ISOLEVELMISMATCH", "The TxIsolation Level property for the COM+ component being created is stronger than the TxIsolationLevel for the root."}
	CoECallOutOfTxScopeNotAllowed                                = &Error{0x8004E030, "CO_E_CALL_OUT_OF_TX_SCOPE_NOT_ALLOWED", "The component attempted to make a cross-context call between invocations of EnterTransactionScope and ExitTransactionScope. This is not allowed. Cross-context calls cannot be made while inside a transaction scope."}
	CoEExitTransactionScopeNotCalled                             = &Error{0x8004E031, "CO_E_EXIT_TRANSACTION_SCOPE_NOT_CALLED", "The component made a call to EnterTransactionScope, but did not make a corresponding call to ExitTransactionScope before returning."}
	EAccessdenied                                                = &Error{0x80070005, "E_ACCESSDENIED", "General access denied error."}
	EOutofmemory                                                 = &Error{0x8007000E, "E_OUTOFMEMORY", "The server does not have enough memory for the new channel."}
	ErrorNotSupported                                            = &Error{0x80070032, "ERROR_NOT_SUPPORTED", "The server cannot support a client request for a dynamic virtual channel."}
	EInvalidarg                                                  = &Error{0x80070057, "E_INVALIDARG", "One or more arguments are invalid."}
	ErrorDiskFull                                                = &Error{0x80070070, "ERROR_DISK_FULL", "There is not enough space on the disk."}
	CoEClassCreateFailed                                         = &Error{0x80080001, "CO_E_CLASS_CREATE_FAILED", "Attempt to create a class object failed."}
	CoEScmError                                                  = &Error{0x80080002, "CO_E_SCM_ERROR", "OLE service could not bind object."}
	CoEScmRpcFailure                                             = &Error{0x80080003, "CO_E_SCM_RPC_FAILURE", "RPC communication failed with OLE service."}
	CoEBadPath                                                   = &Error{0x80080004, "CO_E_BAD_PATH", "Bad path to object."}
	CoEServerExecFailure                                         = &Error{0x80080005, "CO_E_SERVER_EXEC_FAILURE", "Server execution failed."}
	CoEObjsrvRpcFailure                                          = &Error{0x80080006, "CO_E_OBJSRV_RPC_FAILURE", "OLE service could not communicate with the object server."}
	MkENoNormalized                                              = &Error{0x80080007, "MK_E_NO_NORMALIZED", "Moniker path could not be normalized."}
	CoEServerStopping                                            = &Error{0x80080008, "CO_E_SERVER_STOPPING", "Object server is stopping when OLE service contacts it."}
	MemEInvalidRoot                                              = &Error{0x80080009, "MEM_E_INVALID_ROOT", "An invalid root block pointer was specified."}
	MemEInvalidLink                                              = &Error{0x80080010, "MEM_E_INVALID_LINK", "An allocation chain contained an invalid link pointer."}
	MemEInvalidSize                                              = &Error{0x80080011, "MEM_E_INVALID_SIZE", "The requested allocation size was too large."}
	CoEMissingDisplayname                                        = &Error{0x80080015, "CO_E_MISSING_DISPLAYNAME", "The activation requires a display name to be present under the class identifier (CLSID) key."}
	CoERunasValueMustBeAaa                                       = &Error{0x80080016, "CO_E_RUNAS_VALUE_MUST_BE_AAA", "The activation requires that the RunAs value for the application is Activate As Activator."}
	CoEElevationDisabled                                         = &Error{0x80080017, "CO_E_ELEVATION_DISABLED", "The class is not configured to support elevated activation."}
	NteBadUid                                                    = &Error{0x80090001, "NTE_BAD_UID", "Bad UID."}
	NteBadHash                                                   = &Error{0x80090002, "NTE_BAD_HASH", "Bad hash."}
	NteBadKey                                                    = &Error{0x80090003, "NTE_BAD_KEY", "Bad key."}
	NteBadLen                                                    = &Error{0x80090004, "NTE_BAD_LEN", "Bad length."}
	NteBadData                                                   = &Error{0x80090005, "NTE_BAD_DATA", "Bad data."}
	NteBadSignature                                              = &Error{0x80090006, "NTE_BAD_SIGNATURE", "Invalid signature."}
	NteBadVer                                                    = &Error{0x80090007, "NTE_BAD_VER", "Bad version of provider."}
	NteBadAlgid                                                  = &Error{0x80090008, "NTE_BAD_ALGID", "Invalid algorithm specified."}
	NteBadFlags                                                  = &Error{0x80090009, "NTE_BAD_FLAGS", "Invalid flags specified."}
	NteBadType                                                   = &Error{0x8009000A, "NTE_BAD_TYPE", "Invalid type specified."}
	NteBadKeyState                                               = &Error{0x8009000B, "NTE_BAD_KEY_STATE", "Key not valid for use in specified state."}
	NteBadHashState                                              = &Error{0x8009000C, "NTE_BAD_HASH_STATE", "Hash not valid for use in specified state."}
	NteNoKey                                                     = &Error{0x8009000D, "NTE_NO_KEY", "Key does not exist."}
	NteNoMemory                                                  = &Error{0x8009000E, "NTE_NO_MEMORY", "Insufficient memory available for the operation."}
	NteExists                                                    = &Error{0x8009000F, "NTE_EXISTS", "Object already exists."}
	NtePerm                                                      = &Error{0x80090010, "NTE_PERM", "Access denied."}
	NteNotFound                                                  = &Error{0x80090011, "NTE_NOT_FOUND", "Object was not found."}
	NteDoubleEncrypt                                             = &Error{0x80090012, "NTE_DOUBLE_ENCRYPT", "Data already encrypted."}
	NteBadProvider                                               = &Error{0x80090013, "NTE_BAD_PROVIDER", "Invalid provider specified."}
	NteBadProvType                                               = &Error{0x80090014, "NTE_BAD_PROV_TYPE", "Invalid provider type specified."}
	NteBadPublicKey                                              = &Error{0x80090015, "NTE_BAD_PUBLIC_KEY", "Provider's public key is invalid."}
	NteBadKeyset                                                 = &Error{0x80090016, "NTE_BAD_KEYSET", "Key set does not exist."}
	NteProvTypeNotDef                                            = &Error{0x80090017, "NTE_PROV_TYPE_NOT_DEF", "Provider type not defined."}
	NteProvTypeEntryBad                                          = &Error{0x80090018, "NTE_PROV_TYPE_ENTRY_BAD", "The provider type, as registered, is invalid."}
	NteKeysetNotDef                                              = &Error{0x80090019, "NTE_KEYSET_NOT_DEF", "The key set is not defined."}
	NteKeysetEntryBad                                            = &Error{0x8009001A, "NTE_KEYSET_ENTRY_BAD", "The key set, as registered, is invalid."}
	NteProvTypeNoMatch                                           = &Error{0x8009001B, "NTE_PROV_TYPE_NO_MATCH", "Provider type does not match registered value."}
	NteSignatureFileBad                                          = &Error{0x8009001C, "NTE_SIGNATURE_FILE_BAD", "The digital signature file is corrupt."}
	NteProviderDllFail                                           = &Error{0x8009001D, "NTE_PROVIDER_DLL_FAIL", "Provider DLL failed to initialize correctly."}
	NteProvDllNotFound                                           = &Error{0x8009001E, "NTE_PROV_DLL_NOT_FOUND", "Provider DLL could not be found."}
	NteBadKeysetParam                                            = &Error{0x8009001F, "NTE_BAD_KEYSET_PARAM", "The keyset parameter is invalid."}
	NteFail                                                      = &Error{0x80090020, "NTE_FAIL", "An internal error occurred."}
	NteSysErr                                                    = &Error{0x80090021, "NTE_SYS_ERR", "A base error occurred."}
	NteSilentContext                                             = &Error{0x80090022, "NTE_SILENT_CONTEXT", "Provider could not perform the action because the context was acquired as silent."}
	NteTokenKeysetStorageFull                                    = &Error{0x80090023, "NTE_TOKEN_KEYSET_STORAGE_FULL", "The security token does not have storage space available for an additional container."}
	NteTemporaryProfile                                          = &Error{0x80090024, "NTE_TEMPORARY_PROFILE", "The profile for the user is a temporary profile."}
	NteFixedparameter                                            = &Error{0x80090025, "NTE_FIXEDPARAMETER", "The key parameters could not be set because the configuration service provider (CSP) uses fixed parameters."}
	NteInvalidHandle                                             = &Error{0x80090026, "NTE_INVALID_HANDLE", "The supplied handle is invalid."}
	NteInvalidParameter                                          = &Error{0x80090027, "NTE_INVALID_PARAMETER", "The parameter is incorrect."}
	NteBufferTooSmall                                            = &Error{0x80090028, "NTE_BUFFER_TOO_SMALL", "The buffer supplied to a function was too small."}
	NteNotSupported                                              = &Error{0x80090029, "NTE_NOT_SUPPORTED", "The requested operation is not supported."}
	NteNoMoreItems                                               = &Error{0x8009002A, "NTE_NO_MORE_ITEMS", "No more data is available."}
	NteBuffersOverlap                                            = &Error{0x8009002B, "NTE_BUFFERS_OVERLAP", "The supplied buffers overlap incorrectly."}
	NteDecryptionFailure                                         = &Error{0x8009002C, "NTE_DECRYPTION_FAILURE", "The specified data could not be decrypted."}
	NteInternalError                                             = &Error{0x8009002D, "NTE_INTERNAL_ERROR", "An internal consistency check failed."}
	NteUiRequired                                                = &Error{0x8009002E, "NTE_UI_REQUIRED", "This operation requires input from the user."}
	NteHmacNotSupported                                          = &Error{0x8009002F, "NTE_HMAC_NOT_SUPPORTED", "The cryptographic provider does not support Hash Message Authentication Code (HMAC)."}
	SecEInsufficientMemory                                       = &Error{0x80090300, "SEC_E_INSUFFICIENT_MEMORY", "Not enough memory is available to complete this request."}
	SecEInvalidHandle                                            = &Error{0x80090301, "SEC_E_INVALID_HANDLE", "The handle specified is invalid."}
	SecEUnsupportedFunction                                      = &Error{0x80090302, "SEC_E_UNSUPPORTED_FUNCTION", "The function requested is not supported."}
	SecETargetUnknown                                            = &Error{0x80090303, "SEC_E_TARGET_UNKNOWN", "The specified target is unknown or unreachable."}
	SecEInternalError                                            = &Error{0x80090304, "SEC_E_INTERNAL_ERROR", "The Local Security Authority (LSA) cannot be contacted."}
	SecESecpkgNotFound                                           = &Error{0x80090305, "SEC_E_SECPKG_NOT_FOUND", "The requested security package does not exist."}
	SecENotOwner                                                 = &Error{0x80090306, "SEC_E_NOT_OWNER", "The caller is not the owner of the desired credentials."}
	SecECannotInstall                                            = &Error{0x80090307, "SEC_E_CANNOT_INSTALL", "The security package failed to initialize and cannot be installed."}
	SecEInvalidToken                                             = &Error{0x80090308, "SEC_E_INVALID_TOKEN", "The token supplied to the function is invalid."}
	SecECannotPack                                               = &Error{0x80090309, "SEC_E_CANNOT_PACK", "The security package is not able to marshal the logon buffer, so the logon attempt has failed."}
	SecEQopNotSupported                                          = &Error{0x8009030A, "SEC_E_QOP_NOT_SUPPORTED", "The per-message quality of protection is not supported by the security package."}
	SecENoImpersonation                                          = &Error{0x8009030B, "SEC_E_NO_IMPERSONATION", "The security context does not allow impersonation of the client."}
	SecELogonDenied                                              = &Error{0x8009030C, "SEC_E_LOGON_DENIED", "The logon attempt failed."}
	SecEUnknownCredentials                                       = &Error{0x8009030D, "SEC_E_UNKNOWN_CREDENTIALS", "The credentials supplied to the package were not recognized."}
	SecENoCredentials                                            = &Error{0x8009030E, "SEC_E_NO_CREDENTIALS", "No credentials are available in the security package."}
	SecEMessageAltered                                           = &Error{0x8009030F, "SEC_E_MESSAGE_ALTERED", "The message or signature supplied for verification has been altered."}
	SecEOutOfSequence                                            = &Error{0x80090310, "SEC_E_OUT_OF_SEQUENCE", "The message supplied for verification is out of sequence."}
	SecENoAuthenticatingAuthority                                = &Error{0x80090311, "SEC_E_NO_AUTHENTICATING_AUTHORITY", "No authority could be contacted for authentication."}
	SecEBadPkgid                                                 = &Error{0x80090316, "SEC_E_BAD_PKGID", "The requested security package does not exist."}
	SecEContextExpired                                           = &Error{0x80090317, "SEC_E_CONTEXT_EXPIRED", "The context has expired and can no longer be used."}
	SecEIncompleteMessage                                        = &Error{0x80090318, "SEC_E_INCOMPLETE_MESSAGE", "The supplied message is incomplete. The signature was not verified."}
	SecEIncompleteCredentials                                    = &Error{0x80090320, "SEC_E_INCOMPLETE_CREDENTIALS", "The credentials supplied were not complete and could not be verified. The context could not be initialized."}
	SecEBufferTooSmall                                           = &Error{0x80090321, "SEC_E_BUFFER_TOO_SMALL", "The buffers supplied to a function was too small."}
	SecEWrongPrincipal                                           = &Error{0x80090322, "SEC_E_WRONG_PRINCIPAL", "The target principal name is incorrect."}
	SecETimeSkew                                                 = &Error{0x80090324, "SEC_E_TIME_SKEW", "The clocks on the client and server machines are skewed."}
	SecEUntrustedRoot                                            = &Error{0x80090325, "SEC_E_UNTRUSTED_ROOT", "The certificate chain was issued by an authority that is not trusted."}
	SecEIllegalMessage                                           = &Error{0x80090326, "SEC_E_ILLEGAL_MESSAGE", "The message received was unexpected or badly formatted."}
	SecECertUnknown                                              = &Error{0x80090327, "SEC_E_CERT_UNKNOWN", "An unknown error occurred while processing the certificate."}
	SecECertExpired                                              = &Error{0x80090328, "SEC_E_CERT_EXPIRED", "The received certificate has expired."}
	SecEEncryptFailure                                           = &Error{0x80090329, "SEC_E_ENCRYPT_FAILURE", "The specified data could not be encrypted."}
	SecEDecryptFailure                                           = &Error{0x80090330, "SEC_E_DECRYPT_FAILURE", "The specified data could not be decrypted."}
	SecEAlgorithmMismatch                                        = &Error{0x80090331, "SEC_E_ALGORITHM_MISMATCH", "The client and server cannot communicate because they do not possess a common algorithm."}
	SecESecurityQosFailed                                        = &Error{0x80090332, "SEC_E_SECURITY_QOS_FAILED", "The security context could not be established due to a failure in the requested quality of service (for example, mutual authentication or delegation)."}
	SecEUnfinishedContextDeleted                                 = &Error{0x80090333, "SEC_E_UNFINISHED_CONTEXT_DELETED", "A security context was deleted before the context was completed. This is considered a logon failure."}
	SecENoTgtReply                                               = &Error{0x80090334, "SEC_E_NO_TGT_REPLY", "The client is trying to negotiate a context and the server requires user-to-user but did not send a ticket granting ticket (TGT) reply."}
	SecENoIpAddresses                                            = &Error{0x80090335, "SEC_E_NO_IP_ADDRESSES", "Unable to accomplish the requested task because the local machine does not have an IP addresses."}
	SecEWrongCredentialHandle                                    = &Error{0x80090336, "SEC_E_WRONG_CREDENTIAL_HANDLE", "The supplied credential handle does not match the credential associated with the security context."}
	SecECryptoSystemInvalid                                      = &Error{0x80090337, "SEC_E_CRYPTO_SYSTEM_INVALID", "The cryptographic system or checksum function is invalid because a required function is unavailable."}
	SecEMaxReferralsExceeded                                     = &Error{0x80090338, "SEC_E_MAX_REFERRALS_EXCEEDED", "The number of maximum ticket referrals has been exceeded."}
	SecEMustBeKdc                                                = &Error{0x80090339, "SEC_E_MUST_BE_KDC", "The local machine must be a Kerberos domain controller (KDC), and it is not."}
	SecEStrongCryptoNotSupported                                 = &Error{0x8009033A, "SEC_E_STRONG_CRYPTO_NOT_SUPPORTED", "The other end of the security negotiation requires strong cryptographics, but it is not supported on the local machine."}
	SecETooManyPrincipals                                        = &Error{0x8009033B, "SEC_E_TOO_MANY_PRINCIPALS", "The KDC reply contained more than one principal name."}
	SecENoPaData                                                 = &Error{0x8009033C, "SEC_E_NO_PA_DATA", "Expected to find PA data for a hint of what etype to use, but it was not found."}
	SecEPkinitNameMismatch                                       = &Error{0x8009033D, "SEC_E_PKINIT_NAME_MISMATCH", "The client certificate does not contain a valid user principal name (UPN), or does not match the client name in the logon request. Contact your administrator."}
	SecESmartcardLogonRequired                                   = &Error{0x8009033E, "SEC_E_SMARTCARD_LOGON_REQUIRED", "Smart card logon is required and was not used."}
	SecEShutdownInProgress                                       = &Error{0x8009033F, "SEC_E_SHUTDOWN_IN_PROGRESS", "A system shutdown is in progress."}
	SecEKdcInvalidRequest                                        = &Error{0x80090340, "SEC_E_KDC_INVALID_REQUEST", "An invalid request was sent to the KDC."}
	SecEKdcUnableToRefer                                         = &Error{0x80090341, "SEC_E_KDC_UNABLE_TO_REFER", "The KDC was unable to generate a referral for the service requested."}
	SecEKdcUnknownEtype                                          = &Error{0x80090342, "SEC_E_KDC_UNKNOWN_ETYPE", "The encryption type requested is not supported by the KDC."}
	SecEUnsupportedPreauth                                       = &Error{0x80090343, "SEC_E_UNSUPPORTED_PREAUTH", "An unsupported pre-authentication mechanism was presented to the Kerberos package."}
	SecEDelegationRequired                                       = &Error{0x80090345, "SEC_E_DELEGATION_REQUIRED", "The requested operation cannot be completed. The computer must be trusted for delegation, and the current user account must be configured to allow delegation."}
	SecEBadBindings                                              = &Error{0x80090346, "SEC_E_BAD_BINDINGS", "Client's supplied Security Support Provider Interface (SSPI) channel bindings were incorrect."}
	SecEMultipleAccounts                                         = &Error{0x80090347, "SEC_E_MULTIPLE_ACCOUNTS", "The received certificate was mapped to multiple accounts."}
	SecENoKerbKey                                                = &Error{0x80090348, "SEC_E_NO_KERB_KEY", "No Kerberos key was found."}
	SecECertWrongUsage                                           = &Error{0x80090349, "SEC_E_CERT_WRONG_USAGE", "The certificate is not valid for the requested usage."}
	SecEDowngradeDetected                                        = &Error{0x80090350, "SEC_E_DOWNGRADE_DETECTED", "The system detected a possible attempt to compromise security. Ensure that you can contact the server that authenticated you."}
	SecESmartcardCertRevoked                                     = &Error{0x80090351, "SEC_E_SMARTCARD_CERT_REVOKED", "The smart card certificate used for authentication has been revoked. Contact your system administrator. The event log might contain additional information."}
	SecEIssuingCaUntrusted                                       = &Error{0x80090352, "SEC_E_ISSUING_CA_UNTRUSTED", "An untrusted certification authority (CA) was detected while processing the smart card certificate used for authentication. Contact your system administrator."}
	SecERevocationOfflineC                                       = &Error{0x80090353, "SEC_E_REVOCATION_OFFLINE_C", "The revocation status of the smart card certificate used for authentication could not be determined. Contact your system administrator."}
	SecEPkinitClientFailure                                      = &Error{0x80090354, "SEC_E_PKINIT_CLIENT_FAILURE", "The smart card certificate used for authentication was not trusted. Contact your system administrator."}
	SecESmartcardCertExpired                                     = &Error{0x80090355, "SEC_E_SMARTCARD_CERT_EXPIRED", "The smart card certificate used for authentication has expired. Contact your system administrator."}
	SecENoS4uProtSupport                                         = &Error{0x80090356, "SEC_E_NO_S4U_PROT_SUPPORT", "The Kerberos subsystem encountered an error. A service for user protocol requests was made against a domain controller that does not support services for users."}
	SecECrossrealmDelegationFailure                              = &Error{0x80090357, "SEC_E_CROSSREALM_DELEGATION_FAILURE", "An attempt was made by this server to make a Kerberos-constrained delegation request for a target outside the server's realm. This is not supported and indicates a misconfiguration on this server's allowed-to-delegate-to list. Contact your administrator."}
	SecERevocationOfflineKdc                                     = &Error{0x80090358, "SEC_E_REVOCATION_OFFLINE_KDC", "The revocation status of the domain controller certificate used for smart card authentication could not be determined. The system event log contains additional information. Contact your system administrator."}
	SecEIssuingCaUntrustedKdc                                    = &Error{0x80090359, "SEC_E_ISSUING_CA_UNTRUSTED_KDC", "An untrusted CA was detected while processing the domain controller certificate used for authentication. The system event log contains additional information. Contact your system administrator."}
	SecEKdcCertExpired                                           = &Error{0x8009035A, "SEC_E_KDC_CERT_EXPIRED", "The domain controller certificate used for smart card logon has expired. Contact your system administrator with the contents of your system event log."}
	SecEKdcCertRevoked                                           = &Error{0x8009035B, "SEC_E_KDC_CERT_REVOKED", "The domain controller certificate used for smart card logon has been revoked. Contact your system administrator with the contents of your system event log."}
	SecEInvalidParameter                                         = &Error{0x8009035D, "SEC_E_INVALID_PARAMETER", "One or more of the parameters passed to the function were invalid."}
	SecEDelegationPolicy                                         = &Error{0x8009035E, "SEC_E_DELEGATION_POLICY", "The client policy does not allow credential delegation to the target server."}
	SecEPolicyNltmOnly                                           = &Error{0x8009035F, "SEC_E_POLICY_NLTM_ONLY", "The client policy does not allow credential delegation to the target server with NLTM only authentication."}
	CryptEMsgError                                               = &Error{0x80091001, "CRYPT_E_MSG_ERROR", "An error occurred while performing an operation on a cryptographic message."}
	CryptEUnknownAlgo                                            = &Error{0x80091002, "CRYPT_E_UNKNOWN_ALGO", "Unknown cryptographic algorithm."}
	CryptEOidFormat                                              = &Error{0x80091003, "CRYPT_E_OID_FORMAT", "The object identifier is poorly formatted."}
	CryptEInvalidMsgType                                         = &Error{0x80091004, "CRYPT_E_INVALID_MSG_TYPE", "Invalid cryptographic message type."}
	CryptEUnexpectedEncoding                                     = &Error{0x80091005, "CRYPT_E_UNEXPECTED_ENCODING", "Unexpected cryptographic message encoding."}
	CryptEAuthAttrMissing                                        = &Error{0x80091006, "CRYPT_E_AUTH_ATTR_MISSING", "The cryptographic message does not contain an expected authenticated attribute."}
	CryptEHashValue                                              = &Error{0x80091007, "CRYPT_E_HASH_VALUE", "The hash value is not correct."}
	CryptEInvalidIndex                                           = &Error{0x80091008, "CRYPT_E_INVALID_INDEX", "The index value is not valid."}
	CryptEAlreadyDecrypted                                       = &Error{0x80091009, "CRYPT_E_ALREADY_DECRYPTED", "The content of the cryptographic message has already been decrypted."}
	CryptENotDecrypted                                           = &Error{0x8009100A, "CRYPT_E_NOT_DECRYPTED", "The content of the cryptographic message has not been decrypted yet."}
	CryptERecipientNotFound                                      = &Error{0x8009100B, "CRYPT_E_RECIPIENT_NOT_FOUND", "The enveloped-data message does not contain the specified recipient."}
	CryptEControlType                                            = &Error{0x8009100C, "CRYPT_E_CONTROL_TYPE", "Invalid control type."}
	CryptEIssuerSerialnumber                                     = &Error{0x8009100D, "CRYPT_E_ISSUER_SERIALNUMBER", "Invalid issuer or serial number."}
	CryptESignerNotFound                                         = &Error{0x8009100E, "CRYPT_E_SIGNER_NOT_FOUND", "Cannot find the original signer."}
	CryptEAttributesMissing                                      = &Error{0x8009100F, "CRYPT_E_ATTRIBUTES_MISSING", "The cryptographic message does not contain all of the requested attributes."}
	CryptEStreamMsgNotReady                                      = &Error{0x80091010, "CRYPT_E_STREAM_MSG_NOT_READY", "The streamed cryptographic message is not ready to return data."}
	CryptEStreamInsufficientData                                 = &Error{0x80091011, "CRYPT_E_STREAM_INSUFFICIENT_DATA", "The streamed cryptographic message requires more data to complete the decode operation."}
	CryptEBadLen                                                 = &Error{0x80092001, "CRYPT_E_BAD_LEN", "The length specified for the output data was insufficient."}
	CryptEBadEncode                                              = &Error{0x80092002, "CRYPT_E_BAD_ENCODE", "An error occurred during the encode or decode operation."}
	CryptEFileError                                              = &Error{0x80092003, "CRYPT_E_FILE_ERROR", "An error occurred while reading or writing to a file."}
	CryptENotFound                                               = &Error{0x80092004, "CRYPT_E_NOT_FOUND", "Cannot find object or property."}
	CryptEExists                                                 = &Error{0x80092005, "CRYPT_E_EXISTS", "The object or property already exists."}
	CryptENoProvider                                             = &Error{0x80092006, "CRYPT_E_NO_PROVIDER", "No provider was specified for the store or object."}
	CryptESelfSigned                                             = &Error{0x80092007, "CRYPT_E_SELF_SIGNED", "The specified certificate is self-signed."}
	CryptEDeletedPrev                                            = &Error{0x80092008, "CRYPT_E_DELETED_PREV", "The previous certificate or certificate revocation list (CRL) context was deleted."}
	CryptENoMatch                                                = &Error{0x80092009, "CRYPT_E_NO_MATCH", "Cannot find the requested object."}
	CryptEUnexpectedMsgType                                      = &Error{0x8009200A, "CRYPT_E_UNEXPECTED_MSG_TYPE", "The certificate does not have a property that references a private key."}
	CryptENoKeyProperty                                          = &Error{0x8009200B, "CRYPT_E_NO_KEY_PROPERTY", "Cannot find the certificate and private key for decryption."}
	CryptENoDecryptCert                                          = &Error{0x8009200C, "CRYPT_E_NO_DECRYPT_CERT", "Cannot find the certificate and private key to use for decryption."}
	CryptEBadMsg                                                 = &Error{0x8009200D, "CRYPT_E_BAD_MSG", "Not a cryptographic message or the cryptographic message is not formatted correctly."}
	CryptENoSigner                                               = &Error{0x8009200E, "CRYPT_E_NO_SIGNER", "The signed cryptographic message does not have a signer for the specified signer index."}
	CryptEPendingClose                                           = &Error{0x8009200F, "CRYPT_E_PENDING_CLOSE", "Final closure is pending until additional frees or closes."}
	CryptERevoked                                                = &Error{0x80092010, "CRYPT_E_REVOKED", "The certificate is revoked."}
	CryptENoRevocationDll                                        = &Error{0x80092011, "CRYPT_E_NO_REVOCATION_DLL", "No DLL or exported function was found to verify revocation."}
	CryptENoRevocationCheck                                      = &Error{0x80092012, "CRYPT_E_NO_REVOCATION_CHECK", "The revocation function was unable to check revocation for the certificate."}
	CryptERevocationOffline                                      = &Error{0x80092013, "CRYPT_E_REVOCATION_OFFLINE", "The revocation function was unable to check revocation because the revocation server was offline."}
	CryptENotInRevocationDatabase                                = &Error{0x80092014, "CRYPT_E_NOT_IN_REVOCATION_DATABASE", "The certificate is not in the revocation server's database."}
	CryptEInvalidNumericString                                   = &Error{0x80092020, "CRYPT_E_INVALID_NUMERIC_STRING", "The string contains a non-numeric character."}
	CryptEInvalidPrintableString                                 = &Error{0x80092021, "CRYPT_E_INVALID_PRINTABLE_STRING", "The string contains a nonprintable character."}
	CryptEInvalidIa5String                                       = &Error{0x80092022, "CRYPT_E_INVALID_IA5_STRING", "The string contains a character not in the 7-bit ASCII character set."}
	CryptEInvalidX500String                                      = &Error{0x80092023, "CRYPT_E_INVALID_X500_STRING", "The string contains an invalid X500 name attribute key, object identifier (OID), value, or delimiter."}
	CryptENotCharString                                          = &Error{0x80092024, "CRYPT_E_NOT_CHAR_STRING", "The dwValueType for the CERT_NAME_VALUE is not one of the character strings. Most likely it is either a CERT_RDN_ENCODED_BLOB or CERT_TDN_OCTED_STRING."}
	CryptEFileresized                                            = &Error{0x80092025, "CRYPT_E_FILERESIZED", "The Put operation cannot continue. The file needs to be resized. However, there is already a signature present. A complete signing operation must be done."}
	CryptESecuritySettings                                       = &Error{0x80092026, "CRYPT_E_SECURITY_SETTINGS", "The cryptographic operation failed due to a local security option setting."}
	CryptENoVerifyUsageDll                                       = &Error{0x80092027, "CRYPT_E_NO_VERIFY_USAGE_DLL", "No DLL or exported function was found to verify subject usage."}
	CryptENoVerifyUsageCheck                                     = &Error{0x80092028, "CRYPT_E_NO_VERIFY_USAGE_CHECK", "The called function was unable to perform a usage check on the subject."}
	CryptEVerifyUsageOffline                                     = &Error{0x80092029, "CRYPT_E_VERIFY_USAGE_OFFLINE", "The called function was unable to complete the usage check because the server was offline."}
	CryptENotInCtl                                               = &Error{0x8009202A, "CRYPT_E_NOT_IN_CTL", "The subject was not found in a certificate trust list (CTL)."}
	CryptENoTrustedSigner                                        = &Error{0x8009202B, "CRYPT_E_NO_TRUSTED_SIGNER", "None of the signers of the cryptographic message or certificate trust list is trusted."}
	CryptEMissingPubkeyPara                                      = &Error{0x8009202C, "CRYPT_E_MISSING_PUBKEY_PARA", "The public key's algorithm parameters are missing."}
	CryptEOssError                                               = &Error{0x80093000, "CRYPT_E_OSS_ERROR", "OSS Certificate encode/decode error code base."}
	OssMoreBuf                                                   = &Error{0x80093001, "OSS_MORE_BUF", "OSS ASN.1 Error: Output Buffer is too small."}
	OssNegativeUinteger                                          = &Error{0x80093002, "OSS_NEGATIVE_UINTEGER", "OSS ASN.1 Error: Signed integer is encoded as a unsigned integer."}
	OssPduRange                                                  = &Error{0x80093003, "OSS_PDU_RANGE", "OSS ASN.1 Error: Unknown ASN.1 data type."}
	OssMoreInput                                                 = &Error{0x80093004, "OSS_MORE_INPUT", "OSS ASN.1 Error: Output buffer is too small; the decoded data has been truncated."}
	OssDataError                                                 = &Error{0x80093005, "OSS_DATA_ERROR", "OSS ASN.1 Error: Invalid data."}
	OssBadArg                                                    = &Error{0x80093006, "OSS_BAD_ARG", "OSS ASN.1 Error: Invalid argument."}
	OssBadVersion                                                = &Error{0x80093007, "OSS_BAD_VERSION", "OSS ASN.1 Error: Encode/Decode version mismatch."}
	OssOutMemory                                                 = &Error{0x80093008, "OSS_OUT_MEMORY", "OSS ASN.1 Error: Out of memory."}
	OssPduMismatch                                               = &Error{0x80093009, "OSS_PDU_MISMATCH", "OSS ASN.1 Error: Encode/Decode error."}
	OssLimited                                                   = &Error{0x8009300A, "OSS_LIMITED", "OSS ASN.1 Error: Internal error."}
	OssBadPtr                                                    = &Error{0x8009300B, "OSS_BAD_PTR", "OSS ASN.1 Error: Invalid data."}
	OssBadTime                                                   = &Error{0x8009300C, "OSS_BAD_TIME", "OSS ASN.1 Error: Invalid data."}
	OssIndefiniteNotSupported                                    = &Error{0x8009300D, "OSS_INDEFINITE_NOT_SUPPORTED", "OSS ASN.1 Error: Unsupported BER indefinite-length encoding."}
	OssMemError                                                  = &Error{0x8009300E, "OSS_MEM_ERROR", "OSS ASN.1 Error: Access violation."}
	OssBadTable                                                  = &Error{0x8009300F, "OSS_BAD_TABLE", "OSS ASN.1 Error: Invalid data."}
	OssTooLong                                                   = &Error{0x80093010, "OSS_TOO_LONG", "OSS ASN.1 Error: Invalid data."}
	OssConstraintViolated                                        = &Error{0x80093011, "OSS_CONSTRAINT_VIOLATED", "OSS ASN.1 Error: Invalid data."}
	OssFatalError                                                = &Error{0x80093012, "OSS_FATAL_ERROR", "OSS ASN.1 Error: Internal error."}
	OssAccessSerializationError                                  = &Error{0x80093013, "OSS_ACCESS_SERIALIZATION_ERROR", "OSS ASN.1 Error: Multithreading conflict."}
	OssNullTbl                                                   = &Error{0x80093014, "OSS_NULL_TBL", "OSS ASN.1 Error: Invalid data."}
	OssNullFcn                                                   = &Error{0x80093015, "OSS_NULL_FCN", "OSS ASN.1 Error: Invalid data."}
	OssBadEncrules                                               = &Error{0x80093016, "OSS_BAD_ENCRULES", "OSS ASN.1 Error: Invalid data."}
	OssUnavailEncrules                                           = &Error{0x80093017, "OSS_UNAVAIL_ENCRULES", "OSS ASN.1 Error: Encode/Decode function not implemented."}
	OssCantOpenTraceWindow                                       = &Error{0x80093018, "OSS_CANT_OPEN_TRACE_WINDOW", "OSS ASN.1 Error: Trace file error."}
	OssUnimplemented                                             = &Error{0x80093019, "OSS_UNIMPLEMENTED", "OSS ASN.1 Error: Function not implemented."}
	OssOidDllNotLinked                                           = &Error{0x8009301A, "OSS_OID_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssCantOpenTraceFile                                         = &Error{0x8009301B, "OSS_CANT_OPEN_TRACE_FILE", "OSS ASN.1 Error: Trace file error."}
	OssTraceFileAlreadyOpen                                      = &Error{0x8009301C, "OSS_TRACE_FILE_ALREADY_OPEN", "OSS ASN.1 Error: Trace file error."}
	OssTableMismatch                                             = &Error{0x8009301D, "OSS_TABLE_MISMATCH", "OSS ASN.1 Error: Invalid data."}
	OssTypeNotSupported                                          = &Error{0x8009301E, "OSS_TYPE_NOT_SUPPORTED", "OSS ASN.1 Error: Invalid data."}
	OssRealDllNotLinked                                          = &Error{0x8009301F, "OSS_REAL_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssRealCodeNotLinked                                         = &Error{0x80093020, "OSS_REAL_CODE_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssOutOfRange                                                = &Error{0x80093021, "OSS_OUT_OF_RANGE", "OSS ASN.1 Error: Program link error."}
	OssCopierDllNotLinked                                        = &Error{0x80093022, "OSS_COPIER_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssConstraintDllNotLinked                                    = &Error{0x80093023, "OSS_CONSTRAINT_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssComparatorDllNotLinked                                    = &Error{0x80093024, "OSS_COMPARATOR_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssComparatorCodeNotLinked                                   = &Error{0x80093025, "OSS_COMPARATOR_CODE_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssMemMgrDllNotLinked                                        = &Error{0x80093026, "OSS_MEM_MGR_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssPdvDllNotLinked                                           = &Error{0x80093027, "OSS_PDV_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssPdvCodeNotLinked                                          = &Error{0x80093028, "OSS_PDV_CODE_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssApiDllNotLinked                                           = &Error{0x80093029, "OSS_API_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssBerderDllNotLinked                                        = &Error{0x8009302A, "OSS_BERDER_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssPerDllNotLinked                                           = &Error{0x8009302B, "OSS_PER_DLL_NOT_LINKED", "OSS ASN.1 Error: Program link error."}
	OssOpenTypeError                                             = &Error{0x8009302C, "OSS_OPEN_TYPE_ERROR", "OSS ASN.1 Error: Program link error."}
	OssMutexNotCreated                                           = &Error{0x8009302D, "OSS_MUTEX_NOT_CREATED", "OSS ASN.1 Error: System resource error."}
	OssCantCloseTraceFile                                        = &Error{0x8009302E, "OSS_CANT_CLOSE_TRACE_FILE", "OSS ASN.1 Error: Trace file error."}
	CryptEAsn1Error                                              = &Error{0x80093100, "CRYPT_E_ASN1_ERROR", "ASN1 Certificate encode/decode error code base."}
	CryptEAsn1Internal                                           = &Error{0x80093101, "CRYPT_E_ASN1_INTERNAL", "ASN1 internal encode or decode error."}
	CryptEAsn1Eod                                                = &Error{0x80093102, "CRYPT_E_ASN1_EOD", "ASN1 unexpected end of data."}
	CryptEAsn1Corrupt                                            = &Error{0x80093103, "CRYPT_E_ASN1_CORRUPT", "ASN1 corrupted data."}
	CryptEAsn1Large                                              = &Error{0x80093104, "CRYPT_E_ASN1_LARGE", "ASN1 value too large."}
	CryptEAsn1Constraint                                         = &Error{0x80093105, "CRYPT_E_ASN1_CONSTRAINT", "ASN1 constraint violated."}
	CryptEAsn1Memory                                             = &Error{0x80093106, "CRYPT_E_ASN1_MEMORY", "ASN1 out of memory."}
	CryptEAsn1Overflow                                           = &Error{0x80093107, "CRYPT_E_ASN1_OVERFLOW", "ASN1 buffer overflow."}
	CryptEAsn1Badpdu                                             = &Error{0x80093108, "CRYPT_E_ASN1_BADPDU", "ASN1 function not supported for this protocol data unit (PDU)."}
	CryptEAsn1Badargs                                            = &Error{0x80093109, "CRYPT_E_ASN1_BADARGS", "ASN1 bad arguments to function call."}
	CryptEAsn1Badreal                                            = &Error{0x8009310A, "CRYPT_E_ASN1_BADREAL", "ASN1 bad real value."}
	CryptEAsn1Badtag                                             = &Error{0x8009310B, "CRYPT_E_ASN1_BADTAG", "ASN1 bad tag value met."}
	CryptEAsn1Choice                                             = &Error{0x8009310C, "CRYPT_E_ASN1_CHOICE", "ASN1 bad choice value."}
	CryptEAsn1Rule                                               = &Error{0x8009310D, "CRYPT_E_ASN1_RULE", "ASN1 bad encoding rule."}
	CryptEAsn1Utf8                                               = &Error{0x8009310E, "CRYPT_E_ASN1_UTF8", "ASN1 bad Unicode (UTF8)."}
	CryptEAsn1PduType                                            = &Error{0x80093133, "CRYPT_E_ASN1_PDU_TYPE", "ASN1 bad PDU type."}
	CryptEAsn1Nyi                                                = &Error{0x80093134, "CRYPT_E_ASN1_NYI", "ASN1 not yet implemented."}
	CryptEAsn1Extended                                           = &Error{0x80093201, "CRYPT_E_ASN1_EXTENDED", "ASN1 skipped unknown extensions."}
	CryptEAsn1Noeod                                              = &Error{0x80093202, "CRYPT_E_ASN1_NOEOD", "ASN1 end of data expected."}
	CertsrvEBadRequestsubject                                    = &Error{0x80094001, "CERTSRV_E_BAD_REQUESTSUBJECT", "The request subject name is invalid or too long."}
	CertsrvENoRequest                                            = &Error{0x80094002, "CERTSRV_E_NO_REQUEST", "The request does not exist."}
	CertsrvEBadRequeststatus                                     = &Error{0x80094003, "CERTSRV_E_BAD_REQUESTSTATUS", "The request's current status does not allow this operation."}
	CertsrvEPropertyEmpty                                        = &Error{0x80094004, "CERTSRV_E_PROPERTY_EMPTY", "The requested property value is empty."}
	CertsrvEInvalidCaCertificate                                 = &Error{0x80094005, "CERTSRV_E_INVALID_CA_CERTIFICATE", "The CA's certificate contains invalid data."}
	CertsrvEServerSuspended                                      = &Error{0x80094006, "CERTSRV_E_SERVER_SUSPENDED", "Certificate service has been suspended for a database restore operation."}
	CertsrvEEncodingLength                                       = &Error{0x80094007, "CERTSRV_E_ENCODING_LENGTH", "The certificate contains an encoded length that is potentially incompatible with older enrollment software."}
	CertsrvERoleconflict                                         = &Error{0x80094008, "CERTSRV_E_ROLECONFLICT", "The operation is denied. The user has multiple roles assigned, and the CA is configured to enforce role separation."}
	CertsrvERestrictedofficer                                    = &Error{0x80094009, "CERTSRV_E_RESTRICTEDOFFICER", "The operation is denied. It can only be performed by a certificate manager that is allowed to manage certificates for the current requester."}
	CertsrvEKeyArchivalNotConfigured                             = &Error{0x8009400A, "CERTSRV_E_KEY_ARCHIVAL_NOT_CONFIGURED", "Cannot archive private key. The CA is not configured for key archival."}
	CertsrvENoValidKra                                           = &Error{0x8009400B, "CERTSRV_E_NO_VALID_KRA", "Cannot archive private key. The CA could not verify one or more key recovery certificates."}
	CertsrvEBadRequestKeyArchival                                = &Error{0x8009400C, "CERTSRV_E_BAD_REQUEST_KEY_ARCHIVAL", "The request is incorrectly formatted. The encrypted private key must be in an unauthenticated attribute in an outermost signature."}
	CertsrvENoCaadminDefined                                     = &Error{0x8009400D, "CERTSRV_E_NO_CAADMIN_DEFINED", "At least one security principal must have the permission to manage this CA."}
	CertsrvEBadRenewalCertAttribute                              = &Error{0x8009400E, "CERTSRV_E_BAD_RENEWAL_CERT_ATTRIBUTE", "The request contains an invalid renewal certificate attribute."}
	CertsrvENoDbSessions                                         = &Error{0x8009400F, "CERTSRV_E_NO_DB_SESSIONS", "An attempt was made to open a CA database session, but there are already too many active sessions. The server needs to be configured to allow additional sessions."}
	CertsrvEAlignmentFault                                       = &Error{0x80094010, "CERTSRV_E_ALIGNMENT_FAULT", "A memory reference caused a data alignment fault."}
	CertsrvEEnrollDenied                                         = &Error{0x80094011, "CERTSRV_E_ENROLL_DENIED", "The permissions on this CA do not allow the current user to enroll for certificates."}
	CertsrvETemplateDenied                                       = &Error{0x80094012, "CERTSRV_E_TEMPLATE_DENIED", "The permissions on the certificate template do not allow the current user to enroll for this type of certificate."}
	CertsrvEDownlevelDcSslOrUpgrade                              = &Error{0x80094013, "CERTSRV_E_DOWNLEVEL_DC_SSL_OR_UPGRADE", "The contacted domain controller cannot support signed Lightweight Directory Access Protocol (LDAP) traffic. Update the domain controller or configure Certificate Services to use SSL for Active Directory access."}
	CertsrvEUnsupportedCertType                                  = &Error{0x80094800, "CERTSRV_E_UNSUPPORTED_CERT_TYPE", "The requested certificate template is not supported by this CA."}
	CertsrvENoCertType                                           = &Error{0x80094801, "CERTSRV_E_NO_CERT_TYPE", "The request contains no certificate template information."}
	CertsrvETemplateConflict                                     = &Error{0x80094802, "CERTSRV_E_TEMPLATE_CONFLICT", "The request contains conflicting template information."}
	CertsrvESubjectAltNameRequired                               = &Error{0x80094803, "CERTSRV_E_SUBJECT_ALT_NAME_REQUIRED", "The request is missing a required Subject Alternate name extension."}
	CertsrvEArchivedKeyRequired                                  = &Error{0x80094804, "CERTSRV_E_ARCHIVED_KEY_REQUIRED", "The request is missing a required private key for archival by the server."}
	CertsrvESmimeRequired                                        = &Error{0x80094805, "CERTSRV_E_SMIME_REQUIRED", "The request is missing a required SMIME capabilities extension."}
	CertsrvEBadRenewalSubject                                    = &Error{0x80094806, "CERTSRV_E_BAD_RENEWAL_SUBJECT", "The request was made on behalf of a subject other than the caller. The certificate template must be configured to require at least one signature to authorize the request."}
	CertsrvEBadTemplateVersion                                   = &Error{0x80094807, "CERTSRV_E_BAD_TEMPLATE_VERSION", "The request template version is newer than the supported template version."}
	CertsrvETemplatePolicyRequired                               = &Error{0x80094808, "CERTSRV_E_TEMPLATE_POLICY_REQUIRED", "The template is missing a required signature policy attribute."}
	CertsrvESignaturePolicyRequired                              = &Error{0x80094809, "CERTSRV_E_SIGNATURE_POLICY_REQUIRED", "The request is missing required signature policy information."}
	CertsrvESignatureCount                                       = &Error{0x8009480A, "CERTSRV_E_SIGNATURE_COUNT", "The request is missing one or more required signatures."}
	CertsrvESignatureRejected                                    = &Error{0x8009480B, "CERTSRV_E_SIGNATURE_REJECTED", "One or more signatures did not include the required application or issuance policies. The request is missing one or more required valid signatures."}
	CertsrvEIssuancePolicyRequired                               = &Error{0x8009480C, "CERTSRV_E_ISSUANCE_POLICY_REQUIRED", "The request is missing one or more required signature issuance policies."}
	CertsrvESubjectUpnRequired                                   = &Error{0x8009480D, "CERTSRV_E_SUBJECT_UPN_REQUIRED", "The UPN is unavailable and cannot be added to the Subject Alternate name."}
	CertsrvESubjectDirectoryGuidRequired                         = &Error{0x8009480E, "CERTSRV_E_SUBJECT_DIRECTORY_GUID_REQUIRED", "The Active Directory GUID is unavailable and cannot be added to the Subject Alternate name."}
	CertsrvESubjectDnsRequired                                   = &Error{0x8009480F, "CERTSRV_E_SUBJECT_DNS_REQUIRED", "The Domain Name System (DNS) name is unavailable and cannot be added to the Subject Alternate name."}
	CertsrvEArchivedKeyUnexpected                                = &Error{0x80094810, "CERTSRV_E_ARCHIVED_KEY_UNEXPECTED", "The request includes a private key for archival by the server, but key archival is not enabled for the specified certificate template."}
	CertsrvEKeyLength                                            = &Error{0x80094811, "CERTSRV_E_KEY_LENGTH", "The public key does not meet the minimum size required by the specified certificate template."}
	CertsrvESubjectEmailRequired                                 = &Error{0x80094812, "CERTSRV_E_SUBJECT_EMAIL_REQUIRED", "The email name is unavailable and cannot be added to the Subject or Subject Alternate name."}
	CertsrvEUnknownCertType                                      = &Error{0x80094813, "CERTSRV_E_UNKNOWN_CERT_TYPE", "One or more certificate templates to be enabled on this CA could not be found."}
	CertsrvECertTypeOverlap                                      = &Error{0x80094814, "CERTSRV_E_CERT_TYPE_OVERLAP", "The certificate template renewal period is longer than the certificate validity period. The template should be reconfigured or the CA certificate renewed."}
	CertsrvETooManySignatures                                    = &Error{0x80094815, "CERTSRV_E_TOO_MANY_SIGNATURES", "The certificate template requires too many return authorization (RA) signatures. Only one RA signature is allowed."}
	CertsrvERenewalBadPublicKey                                  = &Error{0x80094816, "CERTSRV_E_RENEWAL_BAD_PUBLIC_KEY", "The key used in a renewal request does not match one of the certificates being renewed."}
	CertsrvEInvalidEk                                            = &Error{0x80094817, "CERTSRV_E_INVALID_EK", "The endorsement key certificate is not valid."}
	CertsrvEKeyAttestation                                       = &Error{0x8009481A, "CERTSRV_E_KEY_ATTESTATION", "Key attestation did not succeed."}
	XenrollEKeyNotExportable                                     = &Error{0x80095000, "XENROLL_E_KEY_NOT_EXPORTABLE", "The key is not exportable."}
	XenrollECannotAddRootCert                                    = &Error{0x80095001, "XENROLL_E_CANNOT_ADD_ROOT_CERT", "You cannot add the root CA certificate into your local store."}
	XenrollEResponseKaHashNotFound                               = &Error{0x80095002, "XENROLL_E_RESPONSE_KA_HASH_NOT_FOUND", "The key archival hash attribute was not found in the response."}
	XenrollEResponseUnexpectedKaHash                             = &Error{0x80095003, "XENROLL_E_RESPONSE_UNEXPECTED_KA_HASH", "An unexpected key archival hash attribute was found in the response."}
	XenrollEResponseKaHashMismatch                               = &Error{0x80095004, "XENROLL_E_RESPONSE_KA_HASH_MISMATCH", "There is a key archival hash mismatch between the request and the response."}
	XenrollEKeyspecSmimeMismatch                                 = &Error{0x80095005, "XENROLL_E_KEYSPEC_SMIME_MISMATCH", "Signing certificate cannot include SMIME extension."}
	TrustESystemError                                            = &Error{0x80096001, "TRUST_E_SYSTEM_ERROR", "A system-level error occurred while verifying trust."}
	TrustENoSignerCert                                           = &Error{0x80096002, "TRUST_E_NO_SIGNER_CERT", "The certificate for the signer of the message is invalid or not found."}
	TrustECounterSigner                                          = &Error{0x80096003, "TRUST_E_COUNTER_SIGNER", "One of the counter signatures was invalid."}
	TrustECertSignature                                          = &Error{0x80096004, "TRUST_E_CERT_SIGNATURE", "The signature of the certificate cannot be verified."}
	TrustETimeStamp                                              = &Error{0x80096005, "TRUST_E_TIME_STAMP", "The time-stamp signature or certificate could not be verified or is malformed."}
	TrustEBadDigest                                              = &Error{0x80096010, "TRUST_E_BAD_DIGEST", "The digital signature of the object did not verify."}
	TrustEBasicConstraints                                       = &Error{0x80096019, "TRUST_E_BASIC_CONSTRAINTS", "A certificate's basic constraint extension has not been observed."}
	TrustEFinancialCriteria                                      = &Error{0x8009601E, "TRUST_E_FINANCIAL_CRITERIA", "The certificate does not meet or contain the Authenticode financial extensions."}
	MssipotfEOutofmemrange                                       = &Error{0x80097001, "MSSIPOTF_E_OUTOFMEMRANGE", "Tried to reference a part of the file outside the proper range."}
	MssipotfECantgetobject                                       = &Error{0x80097002, "MSSIPOTF_E_CANTGETOBJECT", "Could not retrieve an object from the file."}
	MssipotfENoheadtable                                         = &Error{0x80097003, "MSSIPOTF_E_NOHEADTABLE", "Could not find the head table in the file."}
	MssipotfEBadMagicnumber                                      = &Error{0x80097004, "MSSIPOTF_E_BAD_MAGICNUMBER", "The magic number in the head table is incorrect."}
	MssipotfEBadOffsetTable                                      = &Error{0x80097005, "MSSIPOTF_E_BAD_OFFSET_TABLE", "The offset table has incorrect values."}
	MssipotfETableTagorder                                       = &Error{0x80097006, "MSSIPOTF_E_TABLE_TAGORDER", "Duplicate table tags or the tags are out of alphabetical order."}
	MssipotfETableLongword                                       = &Error{0x80097007, "MSSIPOTF_E_TABLE_LONGWORD", "A table does not start on a long word boundary."}
	MssipotfEBadFirstTablePlacement                              = &Error{0x80097008, "MSSIPOTF_E_BAD_FIRST_TABLE_PLACEMENT", "First table does not appear after header information."}
	MssipotfETablesOverlap                                       = &Error{0x80097009, "MSSIPOTF_E_TABLES_OVERLAP", "Two or more tables overlap."}
	MssipotfETablePadbytes                                       = &Error{0x8009700A, "MSSIPOTF_E_TABLE_PADBYTES", "Too many pad bytes between tables, or pad bytes are not 0."}
	MssipotfEFiletoosmall                                        = &Error{0x8009700B, "MSSIPOTF_E_FILETOOSMALL", "File is too small to contain the last table."}
	MssipotfETableChecksum                                       = &Error{0x8009700C, "MSSIPOTF_E_TABLE_CHECKSUM", "A table checksum is incorrect."}
	MssipotfEFileChecksum                                        = &Error{0x8009700D, "MSSIPOTF_E_FILE_CHECKSUM", "The file checksum is incorrect."}
	MssipotfEFailedPolicy                                        = &Error{0x80097010, "MSSIPOTF_E_FAILED_POLICY", "The signature does not have the correct attributes for the policy."}
	MssipotfEFailedHintsCheck                                    = &Error{0x80097011, "MSSIPOTF_E_FAILED_HINTS_CHECK", "The file did not pass the hints check."}
	MssipotfENotOpentype                                         = &Error{0x80097012, "MSSIPOTF_E_NOT_OPENTYPE", "The file is not an OpenType file."}
	MssipotfEFile                                                = &Error{0x80097013, "MSSIPOTF_E_FILE", "Failed on a file operation (such as open, map, read, or write)."}
	MssipotfECrypt                                               = &Error{0x80097014, "MSSIPOTF_E_CRYPT", "A call to a CryptoAPI function failed."}
	MssipotfEBadversion                                          = &Error{0x80097015, "MSSIPOTF_E_BADVERSION", "There is a bad version number in the file."}
	MssipotfEDsigStructure                                       = &Error{0x80097016, "MSSIPOTF_E_DSIG_STRUCTURE", "The structure of the DSIG table is incorrect."}
	MssipotfEPconstCheck                                         = &Error{0x80097017, "MSSIPOTF_E_PCONST_CHECK", "A check failed in a partially constant table."}
	MssipotfEStructure                                           = &Error{0x80097018, "MSSIPOTF_E_STRUCTURE", "Some kind of structural error."}
	ErrorCredRequiresConfirmation                                = &Error{0x80097019, "ERROR_CRED_REQUIRES_CONFIRMATION", "The requested credential requires confirmation."}
	TrustEProviderUnknown                                        = &Error{0x800B0001, "TRUST_E_PROVIDER_UNKNOWN", "Unknown trust provider."}
	TrustEActionUnknown                                          = &Error{0x800B0002, "TRUST_E_ACTION_UNKNOWN", "The trust verification action specified is not supported by the specified trust provider."}
	TrustESubjectFormUnknown                                     = &Error{0x800B0003, "TRUST_E_SUBJECT_FORM_UNKNOWN", "The form specified for the subject is not one supported or known by the specified trust provider."}
	TrustESubjectNotTrusted                                      = &Error{0x800B0004, "TRUST_E_SUBJECT_NOT_TRUSTED", "The subject is not trusted for the specified action."}
	DigsigEEncode                                                = &Error{0x800B0005, "DIGSIG_E_ENCODE", "Error due to problem in ASN.1 encoding process."}
	DigsigEDecode                                                = &Error{0x800B0006, "DIGSIG_E_DECODE", "Error due to problem in ASN.1 decoding process."}
	DigsigEExtensibility                                         = &Error{0x800B0007, "DIGSIG_E_EXTENSIBILITY", "Reading/writing extensions where attributes are appropriate, and vice versa."}
	DigsigECrypto                                                = &Error{0x800B0008, "DIGSIG_E_CRYPTO", "Unspecified cryptographic failure."}
	PersistESizedefinite                                         = &Error{0x800B0009, "PERSIST_E_SIZEDEFINITE", "The size of the data could not be determined."}
	PersistESizeindefinite                                       = &Error{0x800B000A, "PERSIST_E_SIZEINDEFINITE", "The size of the indefinite-sized data could not be determined."}
	PersistENotselfsizing                                        = &Error{0x800B000B, "PERSIST_E_NOTSELFSIZING", "This object does not read and write self-sizing data."}
	TrustENosignature                                            = &Error{0x800B0100, "TRUST_E_NOSIGNATURE", "No signature was present in the subject."}
	CertEExpired                                                 = &Error{0x800B0101, "CERT_E_EXPIRED", "A required certificate is not within its validity period when verifying against the current system clock or the time stamp in the signed file."}
	CertEValidityperiodnesting                                   = &Error{0x800B0102, "CERT_E_VALIDITYPERIODNESTING", "The validity periods of the certification chain do not nest correctly."}
	CertERole                                                    = &Error{0x800B0103, "CERT_E_ROLE", "A certificate that can only be used as an end entity is being used as a CA or vice versa."}
	CertEPathlenconst                                            = &Error{0x800B0104, "CERT_E_PATHLENCONST", "A path length constraint in the certification chain has been violated."}
	CertECritical                                                = &Error{0x800B0105, "CERT_E_CRITICAL", "A certificate contains an unknown extension that is marked \"critical\"."}
	CertEPurpose                                                 = &Error{0x800B0106, "CERT_E_PURPOSE", "A certificate is being used for a purpose other than the ones specified by its CA."}
	CertEIssuerchaining                                          = &Error{0x800B0107, "CERT_E_ISSUERCHAINING", "A parent of a given certificate did not issue that child certificate."}
	CertEMalformed                                               = &Error{0x800B0108, "CERT_E_MALFORMED", "A certificate is missing or has an empty value for an important field, such as a subject or issuer name."}
	CertEUntrustedroot                                           = &Error{0x800B0109, "CERT_E_UNTRUSTEDROOT", "A certificate chain processed, but terminated in a root certificate that is not trusted by the trust provider."}
	CertEChaining                                                = &Error{0x800B010A, "CERT_E_CHAINING", "A certificate chain could not be built to a trusted root authority."}
	TrustEFail                                                   = &Error{0x800B010B, "TRUST_E_FAIL", "Generic trust failure."}
	CertERevoked                                                 = &Error{0x800B010C, "CERT_E_REVOKED", "A certificate was explicitly revoked by its issuer. If the certificate is Microsoft Windows PCA 2010, then the driver was signed by a certificate no longer recognized by Windows.<3>"}
	CertEUntrustedtestroot                                       = &Error{0x800B010D, "CERT_E_UNTRUSTEDTESTROOT", "The certification path terminates with the test root that is not trusted with the current policy settings."}
	CertERevocationFailure                                       = &Error{0x800B010E, "CERT_E_REVOCATION_FAILURE", "The revocation process could not continue—the certificates could not be checked."}
	CertECnNoMatch                                               = &Error{0x800B010F, "CERT_E_CN_NO_MATCH", "The certificate's CN name does not match the passed value."}
	CertEWrongUsage                                              = &Error{0x800B0110, "CERT_E_WRONG_USAGE", "The certificate is not valid for the requested usage."}
	TrustEExplicitDistrust                                       = &Error{0x800B0111, "TRUST_E_EXPLICIT_DISTRUST", "The certificate was explicitly marked as untrusted by the user."}
	CertEUntrustedca                                             = &Error{0x800B0112, "CERT_E_UNTRUSTEDCA", "A certification chain processed correctly, but one of the CA certificates is not trusted by the policy provider."}
	CertEInvalidPolicy                                           = &Error{0x800B0113, "CERT_E_INVALID_POLICY", "The certificate has invalid policy."}
	CertEInvalidName                                             = &Error{0x800B0114, "CERT_E_INVALID_NAME", "The certificate has an invalid name. The name is not included in the permitted list or is explicitly excluded."}
	NsWServerBandwidthLimit                                      = &Error{0x800D0003, "NS_W_SERVER_BANDWIDTH_LIMIT", "The maximum filebitrate value specified is greater than the server's configured maximum bandwidth."}
	NsWFileBandwidthLimit                                        = &Error{0x800D0004, "NS_W_FILE_BANDWIDTH_LIMIT", "The maximum bandwidth value specified is less than the maximum filebitrate."}
	NsWUnknownEvent                                              = &Error{0x800D0060, "NS_W_UNKNOWN_EVENT", "Unknown %1 event encountered."}
	NsICatatonicFailure                                          = &Error{0x800D0199, "NS_I_CATATONIC_FAILURE", "Disk %1 ( %2 ) on Content Server %3, will be failed because it is catatonic."}
	NsICatatonicAutoUnfail                                       = &Error{0x800D019A, "NS_I_CATATONIC_AUTO_UNFAIL", "Disk %1 ( %2 ) on Content Server %3, auto online from catatonic state."}
	SpapiEExpectedSectionName                                    = &Error{0x800F0000, "SPAPI_E_EXPECTED_SECTION_NAME", "A non-empty line was encountered in the INF before the start of a section."}
	SpapiEBadSectionNameLine                                     = &Error{0x800F0001, "SPAPI_E_BAD_SECTION_NAME_LINE", "A section name marker in the information file (INF) is not complete or does not exist on a line by itself."}
	SpapiESectionNameTooLong                                     = &Error{0x800F0002, "SPAPI_E_SECTION_NAME_TOO_LONG", "An INF section was encountered whose name exceeds the maximum section name length."}
	SpapiEGeneralSyntax                                          = &Error{0x800F0003, "SPAPI_E_GENERAL_SYNTAX", "The syntax of the INF is invalid."}
	SpapiEWrongInfStyle                                          = &Error{0x800F0100, "SPAPI_E_WRONG_INF_STYLE", "The style of the INF is different than what was requested."}
	SpapiESectionNotFound                                        = &Error{0x800F0101, "SPAPI_E_SECTION_NOT_FOUND", "The required section was not found in the INF."}
	SpapiELineNotFound                                           = &Error{0x800F0102, "SPAPI_E_LINE_NOT_FOUND", "The required line was not found in the INF."}
	SpapiENoBackup                                               = &Error{0x800F0103, "SPAPI_E_NO_BACKUP", "The files affected by the installation of this file queue have not been backed up for uninstall."}
	SpapiENoAssociatedClass                                      = &Error{0x800F0200, "SPAPI_E_NO_ASSOCIATED_CLASS", "The INF or the device information set or element does not have an associated install class."}
	SpapiEClassMismatch                                          = &Error{0x800F0201, "SPAPI_E_CLASS_MISMATCH", "The INF or the device information set or element does not match the specified install class."}
	SpapiEDuplicateFound                                         = &Error{0x800F0202, "SPAPI_E_DUPLICATE_FOUND", "An existing device was found that is a duplicate of the device being manually installed."}
	SpapiENoDriverSelected                                       = &Error{0x800F0203, "SPAPI_E_NO_DRIVER_SELECTED", "There is no driver selected for the device information set or element."}
	SpapiEKeyDoesNotExist                                        = &Error{0x800F0204, "SPAPI_E_KEY_DOES_NOT_EXIST", "The requested device registry key does not exist."}
	SpapiEInvalidDevinstName                                     = &Error{0x800F0205, "SPAPI_E_INVALID_DEVINST_NAME", "The device instance name is invalid."}
	SpapiEInvalidClass                                           = &Error{0x800F0206, "SPAPI_E_INVALID_CLASS", "The install class is not present or is invalid."}
	SpapiEDevinstAlreadyExists                                   = &Error{0x800F0207, "SPAPI_E_DEVINST_ALREADY_EXISTS", "The device instance cannot be created because it already exists."}
	SpapiEDevinfoNotRegistered                                   = &Error{0x800F0208, "SPAPI_E_DEVINFO_NOT_REGISTERED", "The operation cannot be performed on a device information element that has not been registered."}
	SpapiEInvalidRegProperty                                     = &Error{0x800F0209, "SPAPI_E_INVALID_REG_PROPERTY", "The device property code is invalid."}
	SpapiENoInf                                                  = &Error{0x800F020A, "SPAPI_E_NO_INF", "The INF from which a driver list is to be built does not exist."}
	SpapiENoSuchDevinst                                          = &Error{0x800F020B, "SPAPI_E_NO_SUCH_DEVINST", "The device instance does not exist in the hardware tree."}
	SpapiECantLoadClassIcon                                      = &Error{0x800F020C, "SPAPI_E_CANT_LOAD_CLASS_ICON", "The icon representing this install class cannot be loaded."}
	SpapiEInvalidClassInstaller                                  = &Error{0x800F020D, "SPAPI_E_INVALID_CLASS_INSTALLER", "The class installer registry entry is invalid."}
	SpapiEDiDoDefault                                            = &Error{0x800F020E, "SPAPI_E_DI_DO_DEFAULT", "The class installer has indicated that the default action should be performed for this installation request."}
	SpapiEDiNofilecopy                                           = &Error{0x800F020F, "SPAPI_E_DI_NOFILECOPY", "The operation does not require any files to be copied."}
	SpapiEInvalidHwprofile                                       = &Error{0x800F0210, "SPAPI_E_INVALID_HWPROFILE", "The specified hardware profile does not exist."}
	SpapiENoDeviceSelected                                       = &Error{0x800F0211, "SPAPI_E_NO_DEVICE_SELECTED", "There is no device information element currently selected for this device information set."}
	SpapiEDevinfoListLocked                                      = &Error{0x800F0212, "SPAPI_E_DEVINFO_LIST_LOCKED", "The operation cannot be performed because the device information set is locked."}
	SpapiEDevinfoDataLocked                                      = &Error{0x800F0213, "SPAPI_E_DEVINFO_DATA_LOCKED", "The operation cannot be performed because the device information element is locked."}
	SpapiEDiBadPath                                              = &Error{0x800F0214, "SPAPI_E_DI_BAD_PATH", "The specified path does not contain any applicable device INFs."}
	SpapiENoClassinstallParams                                   = &Error{0x800F0215, "SPAPI_E_NO_CLASSINSTALL_PARAMS", "No class installer parameters have been set for the device information set or element."}
	SpapiEFilequeueLocked                                        = &Error{0x800F0216, "SPAPI_E_FILEQUEUE_LOCKED", "The operation cannot be performed because the file queue is locked."}
	SpapiEBadServiceInstallsect                                  = &Error{0x800F0217, "SPAPI_E_BAD_SERVICE_INSTALLSECT", "A service installation section in this INF is invalid."}
	SpapiENoClassDriverList                                      = &Error{0x800F0218, "SPAPI_E_NO_CLASS_DRIVER_LIST", "There is no class driver list for the device information element."}
	SpapiENoAssociatedService                                    = &Error{0x800F0219, "SPAPI_E_NO_ASSOCIATED_SERVICE", "The installation failed because a function driver was not specified for this device instance."}
	SpapiENoDefaultDeviceInterface                               = &Error{0x800F021A, "SPAPI_E_NO_DEFAULT_DEVICE_INTERFACE", "There is presently no default device interface designated for this interface class."}
	SpapiEDeviceInterfaceActive                                  = &Error{0x800F021B, "SPAPI_E_DEVICE_INTERFACE_ACTIVE", "The operation cannot be performed because the device interface is currently active."}
	SpapiEDeviceInterfaceRemoved                                 = &Error{0x800F021C, "SPAPI_E_DEVICE_INTERFACE_REMOVED", "The operation cannot be performed because the device interface has been removed from the system."}
	SpapiEBadInterfaceInstallsect                                = &Error{0x800F021D, "SPAPI_E_BAD_INTERFACE_INSTALLSECT", "An interface installation section in this INF is invalid."}
	SpapiENoSuchInterfaceClass                                   = &Error{0x800F021E, "SPAPI_E_NO_SUCH_INTERFACE_CLASS", "This interface class does not exist in the system."}
	SpapiEInvalidReferenceString                                 = &Error{0x800F021F, "SPAPI_E_INVALID_REFERENCE_STRING", "The reference string supplied for this interface device is invalid."}
	SpapiEInvalidMachinename                                     = &Error{0x800F0220, "SPAPI_E_INVALID_MACHINENAME", "The specified machine name does not conform to Universal Naming Convention (UNCs)."}
	SpapiERemoteCommFailure                                      = &Error{0x800F0221, "SPAPI_E_REMOTE_COMM_FAILURE", "A general remote communication error occurred."}
	SpapiEMachineUnavailable                                     = &Error{0x800F0222, "SPAPI_E_MACHINE_UNAVAILABLE", "The machine selected for remote communication is not available at this time."}
	SpapiENoConfigmgrServices                                    = &Error{0x800F0223, "SPAPI_E_NO_CONFIGMGR_SERVICES", "The Plug and Play service is not available on the remote machine."}
	SpapiEInvalidProppageProvider                                = &Error{0x800F0224, "SPAPI_E_INVALID_PROPPAGE_PROVIDER", "The property page provider registry entry is invalid."}
	SpapiENoSuchDeviceInterface                                  = &Error{0x800F0225, "SPAPI_E_NO_SUCH_DEVICE_INTERFACE", "The requested device interface is not present in the system."}
	SpapiEDiPostprocessingRequired                               = &Error{0x800F0226, "SPAPI_E_DI_POSTPROCESSING_REQUIRED", "The device's co-installer has additional work to perform after installation is complete."}
	SpapiEInvalidCoinstaller                                     = &Error{0x800F0227, "SPAPI_E_INVALID_COINSTALLER", "The device's co-installer is invalid."}
	SpapiENoCompatDrivers                                        = &Error{0x800F0228, "SPAPI_E_NO_COMPAT_DRIVERS", "There are no compatible drivers for this device."}
	SpapiENoDeviceIcon                                           = &Error{0x800F0229, "SPAPI_E_NO_DEVICE_ICON", "There is no icon that represents this device or device type."}
	SpapiEInvalidInfLogconfig                                    = &Error{0x800F022A, "SPAPI_E_INVALID_INF_LOGCONFIG", "A logical configuration specified in this INF is invalid."}
	SpapiEDiDontInstall                                          = &Error{0x800F022B, "SPAPI_E_DI_DONT_INSTALL", "The class installer has denied the request to install or upgrade this device."}
	SpapiEInvalidFilterDriver                                    = &Error{0x800F022C, "SPAPI_E_INVALID_FILTER_DRIVER", "One of the filter drivers installed for this device is invalid."}
	SpapiENonWindowsNtDriver                                     = &Error{0x800F022D, "SPAPI_E_NON_WINDOWS_NT_DRIVER", "The driver selected for this device does not support Windows XP operating system."}
	SpapiENonWindowsDriver                                       = &Error{0x800F022E, "SPAPI_E_NON_WINDOWS_DRIVER", "The driver selected for this device does not support Windows."}
	SpapiENoCatalogForOemInf                                     = &Error{0x800F022F, "SPAPI_E_NO_CATALOG_FOR_OEM_INF", "The third-party INF does not contain digital signature information."}
	SpapiEDevinstallQueueNonnative                               = &Error{0x800F0230, "SPAPI_E_DEVINSTALL_QUEUE_NONNATIVE", "An invalid attempt was made to use a device installation file queue for verification of digital signatures relative to other platforms."}
	SpapiENotDisableable                                         = &Error{0x800F0231, "SPAPI_E_NOT_DISABLEABLE", "The device cannot be disabled."}
	SpapiECantRemoveDevinst                                      = &Error{0x800F0232, "SPAPI_E_CANT_REMOVE_DEVINST", "The device could not be dynamically removed."}
	SpapiEInvalidTarget                                          = &Error{0x800F0233, "SPAPI_E_INVALID_TARGET", "Cannot copy to specified target."}
	SpapiEDriverNonnative                                        = &Error{0x800F0234, "SPAPI_E_DRIVER_NONNATIVE", "Driver is not intended for this platform."}
	SpapiEInWow64                                                = &Error{0x800F0235, "SPAPI_E_IN_WOW64", "Operation not allowed in WOW64."}
	SpapiESetSystemRestorePoint                                  = &Error{0x800F0236, "SPAPI_E_SET_SYSTEM_RESTORE_POINT", "The operation involving unsigned file copying was rolled back, so that a system restore point could be set."}
	SpapiEIncorrectlyCopiedInf                                   = &Error{0x800F0237, "SPAPI_E_INCORRECTLY_COPIED_INF", "An INF was copied into the Windows INF directory in an improper manner."}
	SpapiESceDisabled                                            = &Error{0x800F0238, "SPAPI_E_SCE_DISABLED", "The Security Configuration Editor (SCE) APIs have been disabled on this embedded product."}
	SpapiEUnknownException                                       = &Error{0x800F0239, "SPAPI_E_UNKNOWN_EXCEPTION", "An unknown exception was encountered."}
	SpapiEPnpRegistryError                                       = &Error{0x800F023A, "SPAPI_E_PNP_REGISTRY_ERROR", "A problem was encountered when accessing the Plug and Play registry database."}
	SpapiERemoteRequestUnsupported                               = &Error{0x800F023B, "SPAPI_E_REMOTE_REQUEST_UNSUPPORTED", "The requested operation is not supported for a remote machine."}
	SpapiENotAnInstalledOemInf                                   = &Error{0x800F023C, "SPAPI_E_NOT_AN_INSTALLED_OEM_INF", "The specified file is not an installed original equipment manufacturer (OEM) INF."}
	SpapiEInfInUseByDevices                                      = &Error{0x800F023D, "SPAPI_E_INF_IN_USE_BY_DEVICES", "One or more devices are presently installed using the specified INF."}
	SpapiEDiFunctionObsolete                                     = &Error{0x800F023E, "SPAPI_E_DI_FUNCTION_OBSOLETE", "The requested device install operation is obsolete."}
	SpapiENoAuthenticodeCatalog                                  = &Error{0x800F023F, "SPAPI_E_NO_AUTHENTICODE_CATALOG", "A file could not be verified because it does not have an associated catalog signed via Authenticode."}
	SpapiEAuthenticodeDisallowed                                 = &Error{0x800F0240, "SPAPI_E_AUTHENTICODE_DISALLOWED", "Authenticode signature verification is not supported for the specified INF."}
	SpapiEAuthenticodeTrustedPublisher                           = &Error{0x800F0241, "SPAPI_E_AUTHENTICODE_TRUSTED_PUBLISHER", "The INF was signed with an Authenticode catalog from a trusted publisher."}
	SpapiEAuthenticodeTrustNotEstablished                        = &Error{0x800F0242, "SPAPI_E_AUTHENTICODE_TRUST_NOT_ESTABLISHED", "The publisher of an Authenticode-signed catalog has not yet been established as trusted."}
	SpapiEAuthenticodePublisherNotTrusted                        = &Error{0x800F0243, "SPAPI_E_AUTHENTICODE_PUBLISHER_NOT_TRUSTED", "The publisher of an Authenticode-signed catalog was not established as trusted."}
	SpapiESignatureOsattributeMismatch                           = &Error{0x800F0244, "SPAPI_E_SIGNATURE_OSATTRIBUTE_MISMATCH", "The software was tested for compliance with Windows logo requirements on a different version of Windows and might not be compatible with this version."}
	SpapiEOnlyValidateViaAuthenticode                            = &Error{0x800F0245, "SPAPI_E_ONLY_VALIDATE_VIA_AUTHENTICODE", "The file can be validated only by a catalog signed via Authenticode."}
	SpapiEDeviceInstallerNotReady                                = &Error{0x800F0246, "SPAPI_E_DEVICE_INSTALLER_NOT_READY", "One of the installers for this device cannot perform the installation at this time."}
	SpapiEDriverStoreAddFailed                                   = &Error{0x800F0247, "SPAPI_E_DRIVER_STORE_ADD_FAILED", "A problem was encountered while attempting to add the driver to the store."}
	SpapiEDeviceInstallBlocked                                   = &Error{0x800F0248, "SPAPI_E_DEVICE_INSTALL_BLOCKED", "The installation of this device is forbidden by system policy. Contact your system administrator."}
	SpapiEDriverInstallBlocked                                   = &Error{0x800F0249, "SPAPI_E_DRIVER_INSTALL_BLOCKED", "The installation of this driver is forbidden by system policy. Contact your system administrator."}
	SpapiEWrongInfType                                           = &Error{0x800F024A, "SPAPI_E_WRONG_INF_TYPE", "The specified INF is the wrong type for this operation."}
	SpapiEFileHashNotInCatalog                                   = &Error{0x800F024B, "SPAPI_E_FILE_HASH_NOT_IN_CATALOG", "The hash for the file is not present in the specified catalog file. The file is likely corrupt or the victim of tampering."}
	SpapiEDriverStoreDeleteFailed                                = &Error{0x800F024C, "SPAPI_E_DRIVER_STORE_DELETE_FAILED", "A problem was encountered while attempting to delete the driver from the store."}
	SpapiEUnrecoverableStackOverflow                             = &Error{0x800F0300, "SPAPI_E_UNRECOVERABLE_STACK_OVERFLOW", "An unrecoverable stack overflow was encountered."}
	SpapiEErrorNotInstalled                                      = &Error{0x800F1000, "SPAPI_E_ERROR_NOT_INSTALLED", "No installed components were detected."}
	ScardFInternalError                                          = &Error{0x80100001, "SCARD_F_INTERNAL_ERROR", "An internal consistency check failed."}
	ScardECancelled                                              = &Error{0x80100002, "SCARD_E_CANCELLED", "The action was canceled by an SCardCancel request."}
	ScardEInvalidHandle                                          = &Error{0x80100003, "SCARD_E_INVALID_HANDLE", "The supplied handle was invalid."}
	ScardEInvalidParameter                                       = &Error{0x80100004, "SCARD_E_INVALID_PARAMETER", "One or more of the supplied parameters could not be properly interpreted."}
	ScardEInvalidTarget                                          = &Error{0x80100005, "SCARD_E_INVALID_TARGET", "Registry startup information is missing or invalid."}
	ScardENoMemory                                               = &Error{0x80100006, "SCARD_E_NO_MEMORY", "Not enough memory available to complete this command."}
	ScardFWaitedTooLong                                          = &Error{0x80100007, "SCARD_F_WAITED_TOO_LONG", "An internal consistency timer has expired."}
	ScardEInsufficientBuffer                                     = &Error{0x80100008, "SCARD_E_INSUFFICIENT_BUFFER", "The data buffer to receive returned data is too small for the returned data."}
	ScardEUnknownReader                                          = &Error{0x80100009, "SCARD_E_UNKNOWN_READER", "The specified reader name is not recognized."}
	ScardETimeout                                                = &Error{0x8010000A, "SCARD_E_TIMEOUT", "The user-specified time-out value has expired."}
	ScardESharingViolation                                       = &Error{0x8010000B, "SCARD_E_SHARING_VIOLATION", "The smart card cannot be accessed because of other connections outstanding."}
	ScardENoSmartcard                                            = &Error{0x8010000C, "SCARD_E_NO_SMARTCARD", "The operation requires a smart card, but no smart card is currently in the device."}
	ScardEUnknownCard                                            = &Error{0x8010000D, "SCARD_E_UNKNOWN_CARD", "The specified smart card name is not recognized."}
	ScardECantDispose                                            = &Error{0x8010000E, "SCARD_E_CANT_DISPOSE", "The system could not dispose of the media in the requested manner."}
	ScardEProtoMismatch                                          = &Error{0x8010000F, "SCARD_E_PROTO_MISMATCH", "The requested protocols are incompatible with the protocol currently in use with the smart card."}
	ScardENotReady                                               = &Error{0x80100010, "SCARD_E_NOT_READY", "The reader or smart card is not ready to accept commands."}
	ScardEInvalidValue                                           = &Error{0x80100011, "SCARD_E_INVALID_VALUE", "One or more of the supplied parameters values could not be properly interpreted."}
	ScardESystemCancelled                                        = &Error{0x80100012, "SCARD_E_SYSTEM_CANCELLED", "The action was canceled by the system, presumably to log off or shut down."}
	ScardFCommError                                              = &Error{0x80100013, "SCARD_F_COMM_ERROR", "An internal communications error has been detected."}
	ScardFUnknownError                                           = &Error{0x80100014, "SCARD_F_UNKNOWN_ERROR", "An internal error has been detected, but the source is unknown."}
	ScardEInvalidAtr                                             = &Error{0x80100015, "SCARD_E_INVALID_ATR", "An automatic terminal recognition (ATR) obtained from the registry is not a valid ATR string."}
	ScardENotTransacted                                          = &Error{0x80100016, "SCARD_E_NOT_TRANSACTED", "An attempt was made to end a nonexistent transaction."}
	ScardEReaderUnavailable                                      = &Error{0x80100017, "SCARD_E_READER_UNAVAILABLE", "The specified reader is not currently available for use."}
	ScardPShutdown                                               = &Error{0x80100018, "SCARD_P_SHUTDOWN", "The operation has been aborted to allow the server application to exit."}
	ScardEPciTooSmall                                            = &Error{0x80100019, "SCARD_E_PCI_TOO_SMALL", "The peripheral component interconnect (PCI) Receive buffer was too small."}
	ScardEReaderUnsupported                                      = &Error{0x8010001A, "SCARD_E_READER_UNSUPPORTED", "The reader driver does not meet minimal requirements for support."}
	ScardEDuplicateReader                                        = &Error{0x8010001B, "SCARD_E_DUPLICATE_READER", "The reader driver did not produce a unique reader name."}
	ScardECardUnsupported                                        = &Error{0x8010001C, "SCARD_E_CARD_UNSUPPORTED", "The smart card does not meet minimal requirements for support."}
	ScardENoService                                              = &Error{0x8010001D, "SCARD_E_NO_SERVICE", "The smart card resource manager is not running."}
	ScardEServiceStopped                                         = &Error{0x8010001E, "SCARD_E_SERVICE_STOPPED", "The smart card resource manager has shut down."}
	ScardEUnexpected                                             = &Error{0x8010001F, "SCARD_E_UNEXPECTED", "An unexpected card error has occurred."}
	ScardEIccInstallation                                        = &Error{0x80100020, "SCARD_E_ICC_INSTALLATION", "No primary provider can be found for the smart card."}
	ScardEIccCreateorder                                         = &Error{0x80100021, "SCARD_E_ICC_CREATEORDER", "The requested order of object creation is not supported."}
	ScardEUnsupportedFeature                                     = &Error{0x80100022, "SCARD_E_UNSUPPORTED_FEATURE", "This smart card does not support the requested feature."}
	ScardEDirNotFound                                            = &Error{0x80100023, "SCARD_E_DIR_NOT_FOUND", "The identified directory does not exist in the smart card."}
	ScardEFileNotFound                                           = &Error{0x80100024, "SCARD_E_FILE_NOT_FOUND", "The identified file does not exist in the smart card."}
	ScardENoDir                                                  = &Error{0x80100025, "SCARD_E_NO_DIR", "The supplied path does not represent a smart card directory."}
	ScardENoFile                                                 = &Error{0x80100026, "SCARD_E_NO_FILE", "The supplied path does not represent a smart card file."}
	ScardENoAccess                                               = &Error{0x80100027, "SCARD_E_NO_ACCESS", "Access is denied to this file."}
	ScardEWriteTooMany                                           = &Error{0x80100028, "SCARD_E_WRITE_TOO_MANY", "The smart card does not have enough memory to store the information."}
	ScardEBadSeek                                                = &Error{0x80100029, "SCARD_E_BAD_SEEK", "There was an error trying to set the smart card file object pointer."}
	ScardEInvalidChv                                             = &Error{0x8010002A, "SCARD_E_INVALID_CHV", "The supplied PIN is incorrect."}
	ScardEUnknownResMng                                          = &Error{0x8010002B, "SCARD_E_UNKNOWN_RES_MNG", "An unrecognized error code was returned from a layered component."}
	ScardENoSuchCertificate                                      = &Error{0x8010002C, "SCARD_E_NO_SUCH_CERTIFICATE", "The requested certificate does not exist."}
	ScardECertificateUnavailable                                 = &Error{0x8010002D, "SCARD_E_CERTIFICATE_UNAVAILABLE", "The requested certificate could not be obtained."}
	ScardENoReadersAvailable                                     = &Error{0x8010002E, "SCARD_E_NO_READERS_AVAILABLE", "Cannot find a smart card reader."}
	ScardECommDataLost                                           = &Error{0x8010002F, "SCARD_E_COMM_DATA_LOST", "A communications error with the smart card has been detected. Retry the operation."}
	ScardENoKeyContainer                                         = &Error{0x80100030, "SCARD_E_NO_KEY_CONTAINER", "The requested key container does not exist on the smart card."}
	ScardEServerTooBusy                                          = &Error{0x80100031, "SCARD_E_SERVER_TOO_BUSY", "The smart card resource manager is too busy to complete this operation."}
	ScardWUnsupportedCard                                        = &Error{0x80100065, "SCARD_W_UNSUPPORTED_CARD", "The reader cannot communicate with the smart card, due to ATR configuration conflicts."}
	ScardWUnresponsiveCard                                       = &Error{0x80100066, "SCARD_W_UNRESPONSIVE_CARD", "The smart card is not responding to a reset."}
	ScardWUnpoweredCard                                          = &Error{0x80100067, "SCARD_W_UNPOWERED_CARD", "Power has been removed from the smart card, so that further communication is not possible."}
	ScardWResetCard                                              = &Error{0x80100068, "SCARD_W_RESET_CARD", "The smart card has been reset, so any shared state information is invalid."}
	ScardWRemovedCard                                            = &Error{0x80100069, "SCARD_W_REMOVED_CARD", "The smart card has been removed, so that further communication is not possible."}
	ScardWSecurityViolation                                      = &Error{0x8010006A, "SCARD_W_SECURITY_VIOLATION", "Access was denied because of a security violation."}
	ScardWWrongChv                                               = &Error{0x8010006B, "SCARD_W_WRONG_CHV", "The card cannot be accessed because the wrong PIN was presented."}
	ScardWChvBlocked                                             = &Error{0x8010006C, "SCARD_W_CHV_BLOCKED", "The card cannot be accessed because the maximum number of PIN entry attempts has been reached."}
	ScardWEof                                                    = &Error{0x8010006D, "SCARD_W_EOF", "The end of the smart card file has been reached."}
	ScardWCancelledByUser                                        = &Error{0x8010006E, "SCARD_W_CANCELLED_BY_USER", "The action was canceled by the user."}
	ScardWCardNotAuthenticated                                   = &Error{0x8010006F, "SCARD_W_CARD_NOT_AUTHENTICATED", "No PIN was presented to the smart card."}
	ComadminEObjecterrors                                        = &Error{0x80110401, "COMADMIN_E_OBJECTERRORS", "Errors occurred accessing one or more objects—the ErrorInfo collection contains more detail."}
	ComadminEObjectinvalid                                       = &Error{0x80110402, "COMADMIN_E_OBJECTINVALID", "One or more of the object's properties are missing or invalid."}
	ComadminEKeymissing                                          = &Error{0x80110403, "COMADMIN_E_KEYMISSING", "The object was not found in the catalog."}
	ComadminEAlreadyinstalled                                    = &Error{0x80110404, "COMADMIN_E_ALREADYINSTALLED", "The object is already registered."}
	ComadminEAppFileWritefail                                    = &Error{0x80110407, "COMADMIN_E_APP_FILE_WRITEFAIL", "An error occurred writing to the application file."}
	ComadminEAppFileReadfail                                     = &Error{0x80110408, "COMADMIN_E_APP_FILE_READFAIL", "An error occurred reading the application file."}
	ComadminEAppFileVersion                                      = &Error{0x80110409, "COMADMIN_E_APP_FILE_VERSION", "Invalid version number in application file."}
	ComadminEBadpath                                             = &Error{0x8011040A, "COMADMIN_E_BADPATH", "The file path is invalid."}
	ComadminEApplicationexists                                   = &Error{0x8011040B, "COMADMIN_E_APPLICATIONEXISTS", "The application is already installed."}
	ComadminERoleexists                                          = &Error{0x8011040C, "COMADMIN_E_ROLEEXISTS", "The role already exists."}
	ComadminECantcopyfile                                        = &Error{0x8011040D, "COMADMIN_E_CANTCOPYFILE", "An error occurred copying the file."}
	ComadminENouser                                              = &Error{0x8011040F, "COMADMIN_E_NOUSER", "One or more users are not valid."}
	ComadminEInvaliduserids                                      = &Error{0x80110410, "COMADMIN_E_INVALIDUSERIDS", "One or more users in the application file are not valid."}
	ComadminENoregistryclsid                                     = &Error{0x80110411, "COMADMIN_E_NOREGISTRYCLSID", "The component's CLSID is missing or corrupt."}
	ComadminEBadregistryprogid                                   = &Error{0x80110412, "COMADMIN_E_BADREGISTRYPROGID", "The component's programmatic ID is missing or corrupt."}
	ComadminEAuthenticationlevel                                 = &Error{0x80110413, "COMADMIN_E_AUTHENTICATIONLEVEL", "Unable to set required authentication level for update request."}
	ComadminEUserpasswdnotvalid                                  = &Error{0x80110414, "COMADMIN_E_USERPASSWDNOTVALID", "The identity or password set on the application is not valid."}
	ComadminEClsidoriidmismatch                                  = &Error{0x80110418, "COMADMIN_E_CLSIDORIIDMISMATCH", "Application file CLSIDs or instance identifiers (IIDs) do not match corresponding DLLs."}
	ComadminERemoteinterface                                     = &Error{0x80110419, "COMADMIN_E_REMOTEINTERFACE", "Interface information is either missing or changed."}
	ComadminEDllregisterserver                                   = &Error{0x8011041A, "COMADMIN_E_DLLREGISTERSERVER", "DllRegisterServer failed on component install."}
	ComadminENoservershare                                       = &Error{0x8011041B, "COMADMIN_E_NOSERVERSHARE", "No server file share available."}
	ComadminEDllloadfailed                                       = &Error{0x8011041D, "COMADMIN_E_DLLLOADFAILED", "DLL could not be loaded."}
	ComadminEBadregistrylibid                                    = &Error{0x8011041E, "COMADMIN_E_BADREGISTRYLIBID", "The registered TypeLib ID is not valid."}
	ComadminEAppdirnotfound                                      = &Error{0x8011041F, "COMADMIN_E_APPDIRNOTFOUND", "Application install directory not found."}
	ComadminERegistrarfailed                                     = &Error{0x80110423, "COMADMIN_E_REGISTRARFAILED", "Errors occurred while in the component registrar."}
	ComadminECompfileDoesnotexist                                = &Error{0x80110424, "COMADMIN_E_COMPFILE_DOESNOTEXIST", "The file does not exist."}
	ComadminECompfileLoaddllfail                                 = &Error{0x80110425, "COMADMIN_E_COMPFILE_LOADDLLFAIL", "The DLL could not be loaded."}
	ComadminECompfileGetclassobj                                 = &Error{0x80110426, "COMADMIN_E_COMPFILE_GETCLASSOBJ", "GetClassObject failed in the DLL."}
	ComadminECompfileClassnotavail                               = &Error{0x80110427, "COMADMIN_E_COMPFILE_CLASSNOTAVAIL", "The DLL does not support the components listed in the TypeLib."}
	ComadminECompfileBadtlb                                      = &Error{0x80110428, "COMADMIN_E_COMPFILE_BADTLB", "The TypeLib could not be loaded."}
	ComadminECompfileNotinstallable                              = &Error{0x80110429, "COMADMIN_E_COMPFILE_NOTINSTALLABLE", "The file does not contain components or component information."}
	ComadminENotchangeable                                       = &Error{0x8011042A, "COMADMIN_E_NOTCHANGEABLE", "Changes to this object and its subobjects have been disabled."}
	ComadminENotdeleteable                                       = &Error{0x8011042B, "COMADMIN_E_NOTDELETEABLE", "The delete function has been disabled for this object."}
	ComadminESession                                             = &Error{0x8011042C, "COMADMIN_E_SESSION", "The server catalog version is not supported."}
	ComadminECompMoveLocked                                      = &Error{0x8011042D, "COMADMIN_E_COMP_MOVE_LOCKED", "The component move was disallowed because the source or destination application is either a system application or currently locked against changes."}
	ComadminECompMoveBadDest                                     = &Error{0x8011042E, "COMADMIN_E_COMP_MOVE_BAD_DEST", "The component move failed because the destination application no longer exists."}
	ComadminERegistertlb                                         = &Error{0x80110430, "COMADMIN_E_REGISTERTLB", "The system was unable to register the TypeLib."}
	ComadminESystemapp                                           = &Error{0x80110433, "COMADMIN_E_SYSTEMAPP", "This operation cannot be performed on the system application."}
	ComadminECompfileNoregistrar                                 = &Error{0x80110434, "COMADMIN_E_COMPFILE_NOREGISTRAR", "The component registrar referenced in this file is not available."}
	ComadminECoreqcompinstalled                                  = &Error{0x80110435, "COMADMIN_E_COREQCOMPINSTALLED", "A component in the same DLL is already installed."}
	ComadminEServicenotinstalled                                 = &Error{0x80110436, "COMADMIN_E_SERVICENOTINSTALLED", "The service is not installed."}
	ComadminEPropertysavefailed                                  = &Error{0x80110437, "COMADMIN_E_PROPERTYSAVEFAILED", "One or more property settings are either invalid or in conflict with each other."}
	ComadminEObjectexists                                        = &Error{0x80110438, "COMADMIN_E_OBJECTEXISTS", "The object you are attempting to add or rename already exists."}
	ComadminEComponentexists                                     = &Error{0x80110439, "COMADMIN_E_COMPONENTEXISTS", "The component already exists."}
	ComadminERegfileCorrupt                                      = &Error{0x8011043B, "COMADMIN_E_REGFILE_CORRUPT", "The registration file is corrupt."}
	ComadminEPropertyOverflow                                    = &Error{0x8011043C, "COMADMIN_E_PROPERTY_OVERFLOW", "The property value is too large."}
	ComadminENotinregistry                                       = &Error{0x8011043E, "COMADMIN_E_NOTINREGISTRY", "Object was not found in registry."}
	ComadminEObjectnotpoolable                                   = &Error{0x8011043F, "COMADMIN_E_OBJECTNOTPOOLABLE", "This object cannot be pooled."}
	ComadminEApplidMatchesClsid                                  = &Error{0x80110446, "COMADMIN_E_APPLID_MATCHES_CLSID", "A CLSID with the same GUID as the new application ID is already installed on this machine."}
	ComadminERoleDoesNotExist                                    = &Error{0x80110447, "COMADMIN_E_ROLE_DOES_NOT_EXIST", "A role assigned to a component, interface, or method did not exist in the application."}
	ComadminEStartAppNeedsComponents                             = &Error{0x80110448, "COMADMIN_E_START_APP_NEEDS_COMPONENTS", "You must have components in an application to start the application."}
	ComadminERequiresDifferentPlatform                           = &Error{0x80110449, "COMADMIN_E_REQUIRES_DIFFERENT_PLATFORM", "This operation is not enabled on this platform."}
	ComadminECanNotExportAppProxy                                = &Error{0x8011044A, "COMADMIN_E_CAN_NOT_EXPORT_APP_PROXY", "Application proxy is not exportable."}
	ComadminECanNotStartApp                                      = &Error{0x8011044B, "COMADMIN_E_CAN_NOT_START_APP", "Failed to start application because it is either a library application or an application proxy."}
	ComadminECanNotExportSysApp                                  = &Error{0x8011044C, "COMADMIN_E_CAN_NOT_EXPORT_SYS_APP", "System application is not exportable."}
	ComadminECantSubscribeToComponent                            = &Error{0x8011044D, "COMADMIN_E_CANT_SUBSCRIBE_TO_COMPONENT", "Cannot subscribe to this component (the component might have been imported)."}
	ComadminEEventclassCantBeSubscriber                          = &Error{0x8011044E, "COMADMIN_E_EVENTCLASS_CANT_BE_SUBSCRIBER", "An event class cannot also be a subscriber component."}
	ComadminELibAppProxyIncompatible                             = &Error{0x8011044F, "COMADMIN_E_LIB_APP_PROXY_INCOMPATIBLE", "Library applications and application proxies are incompatible."}
	ComadminEBasePartitionOnly                                   = &Error{0x80110450, "COMADMIN_E_BASE_PARTITION_ONLY", "This function is valid for the base partition only."}
	ComadminEStartAppDisabled                                    = &Error{0x80110451, "COMADMIN_E_START_APP_DISABLED", "You cannot start an application that has been disabled."}
	ComadminECatDuplicatePartitionName                           = &Error{0x80110457, "COMADMIN_E_CAT_DUPLICATE_PARTITION_NAME", "The specified partition name is already in use on this computer."}
	ComadminECatInvalidPartitionName                             = &Error{0x80110458, "COMADMIN_E_CAT_INVALID_PARTITION_NAME", "The specified partition name is invalid. Check that the name contains at least one visible character."}
	ComadminECatPartitionInUse                                   = &Error{0x80110459, "COMADMIN_E_CAT_PARTITION_IN_USE", "The partition cannot be deleted because it is the default partition for one or more users."}
	ComadminEFilePartitionDuplicateFiles                         = &Error{0x8011045A, "COMADMIN_E_FILE_PARTITION_DUPLICATE_FILES", "The partition cannot be exported because one or more components in the partition have the same file name."}
	ComadminECatImportedComponentsNotAllowed                     = &Error{0x8011045B, "COMADMIN_E_CAT_IMPORTED_COMPONENTS_NOT_ALLOWED", "Applications that contain one or more imported components cannot be installed into a nonbase partition."}
	ComadminEAmbiguousApplicationName                            = &Error{0x8011045C, "COMADMIN_E_AMBIGUOUS_APPLICATION_NAME", "The application name is not unique and cannot be resolved to an application ID."}
	ComadminEAmbiguousPartitionName                              = &Error{0x8011045D, "COMADMIN_E_AMBIGUOUS_PARTITION_NAME", "The partition name is not unique and cannot be resolved to a partition ID."}
	ComadminERegdbNotinitialized                                 = &Error{0x80110472, "COMADMIN_E_REGDB_NOTINITIALIZED", "The COM+ registry database has not been initialized."}
	ComadminERegdbNotopen                                        = &Error{0x80110473, "COMADMIN_E_REGDB_NOTOPEN", "The COM+ registry database is not open."}
	ComadminERegdbSystemerr                                      = &Error{0x80110474, "COMADMIN_E_REGDB_SYSTEMERR", "The COM+ registry database detected a system error."}
	ComadminERegdbAlreadyrunning                                 = &Error{0x80110475, "COMADMIN_E_REGDB_ALREADYRUNNING", "The COM+ registry database is already running."}
	ComadminEMigVersionnotsupported                              = &Error{0x80110480, "COMADMIN_E_MIG_VERSIONNOTSUPPORTED", "This version of the COM+ registry database cannot be migrated."}
	ComadminEMigSchemanotfound                                   = &Error{0x80110481, "COMADMIN_E_MIG_SCHEMANOTFOUND", "The schema version to be migrated could not be found in the COM+ registry database."}
	ComadminECatBitnessmismatch                                  = &Error{0x80110482, "COMADMIN_E_CAT_BITNESSMISMATCH", "There was a type mismatch between binaries."}
	ComadminECatUnacceptablebitness                              = &Error{0x80110483, "COMADMIN_E_CAT_UNACCEPTABLEBITNESS", "A binary of unknown or invalid type was provided."}
	ComadminECatWrongappbitness                                  = &Error{0x80110484, "COMADMIN_E_CAT_WRONGAPPBITNESS", "There was a type mismatch between a binary and an application."}
	ComadminECatPauseResumeNotSupported                          = &Error{0x80110485, "COMADMIN_E_CAT_PAUSE_RESUME_NOT_SUPPORTED", "The application cannot be paused or resumed."}
	ComadminECatServerfault                                      = &Error{0x80110486, "COMADMIN_E_CAT_SERVERFAULT", "The COM+ catalog server threw an exception during execution."}
	ComqcEApplicationNotQueued                                   = &Error{0x80110600, "COMQC_E_APPLICATION_NOT_QUEUED", "Only COM+ applications marked \"queued\" can be invoked using the \"queue\" moniker."}
	ComqcENoQueueableInterfaces                                  = &Error{0x80110601, "COMQC_E_NO_QUEUEABLE_INTERFACES", "At least one interface must be marked \"queued\" to create a queued component instance with the \"queue\" moniker."}
	ComqcEQueuingServiceNotAvailable                             = &Error{0x80110602, "COMQC_E_QUEUING_SERVICE_NOT_AVAILABLE", "Message Queuing is required for the requested operation and is not installed."}
	ComqcENoIpersiststream                                       = &Error{0x80110603, "COMQC_E_NO_IPERSISTSTREAM", "Unable to marshal an interface that does not support IPersistStream."}
	ComqcEBadMessage                                             = &Error{0x80110604, "COMQC_E_BAD_MESSAGE", "The message is improperly formatted or was damaged in transit."}
	ComqcEUnauthenticated                                        = &Error{0x80110605, "COMQC_E_UNAUTHENTICATED", "An unauthenticated message was received by an application that accepts only authenticated messages."}
	ComqcEUntrustedEnqueuer                                      = &Error{0x80110606, "COMQC_E_UNTRUSTED_ENQUEUER", "The message was requeued or moved by a user not in the QC Trusted User \"role\"."}
	MsdtcEDuplicateResource                                      = &Error{0x80110701, "MSDTC_E_DUPLICATE_RESOURCE", "Cannot create a duplicate resource of type Distributed Transaction Coordinator."}
	ComadminEObjectParentMissing                                 = &Error{0x80110808, "COMADMIN_E_OBJECT_PARENT_MISSING", "One of the objects being inserted or updated does not belong to a valid parent collection."}
	ComadminEObjectDoesNotExist                                  = &Error{0x80110809, "COMADMIN_E_OBJECT_DOES_NOT_EXIST", "One of the specified objects cannot be found."}
	ComadminEAppNotRunning                                       = &Error{0x8011080A, "COMADMIN_E_APP_NOT_RUNNING", "The specified application is not currently running."}
	ComadminEInvalidPartition                                    = &Error{0x8011080B, "COMADMIN_E_INVALID_PARTITION", "The partitions specified are not valid."}
	ComadminESvcappNotPoolableOrRecyclable                       = &Error{0x8011080D, "COMADMIN_E_SVCAPP_NOT_POOLABLE_OR_RECYCLABLE", "COM+ applications that run as Windows NT service cannot be pooled or recycled."}
	ComadminEUserInSet                                           = &Error{0x8011080E, "COMADMIN_E_USER_IN_SET", "One or more users are already assigned to a local partition set."}
	ComadminECantrecyclelibraryapps                              = &Error{0x8011080F, "COMADMIN_E_CANTRECYCLELIBRARYAPPS", "Library applications cannot be recycled."}
	ComadminECantrecycleserviceapps                              = &Error{0x80110811, "COMADMIN_E_CANTRECYCLESERVICEAPPS", "Applications running as Windows NT services cannot be recycled."}
	ComadminEProcessalreadyrecycled                              = &Error{0x80110812, "COMADMIN_E_PROCESSALREADYRECYCLED", "The process has already been recycled."}
	ComadminEPausedprocessmaynotberecycled                       = &Error{0x80110813, "COMADMIN_E_PAUSEDPROCESSMAYNOTBERECYCLED", "A paused process cannot be recycled."}
	ComadminECantmakeinprocservice                               = &Error{0x80110814, "COMADMIN_E_CANTMAKEINPROCSERVICE", "Library applications cannot be Windows NT services."}
	ComadminEProgidinusebyclsid                                  = &Error{0x80110815, "COMADMIN_E_PROGIDINUSEBYCLSID", "The ProgID provided to the copy operation is invalid. The ProgID is in use by another registered CLSID."}
	ComadminEDefaultPartitionNotInSet                            = &Error{0x80110816, "COMADMIN_E_DEFAULT_PARTITION_NOT_IN_SET", "The partition specified as the default is not a member of the partition set."}
	ComadminERecycledprocessmaynotbepaused                       = &Error{0x80110817, "COMADMIN_E_RECYCLEDPROCESSMAYNOTBEPAUSED", "A recycled process cannot be paused."}
	ComadminEPartitionAccessdenied                               = &Error{0x80110818, "COMADMIN_E_PARTITION_ACCESSDENIED", "Access to the specified partition is denied."}
	ComadminEPartitionMsiOnly                                    = &Error{0x80110819, "COMADMIN_E_PARTITION_MSI_ONLY", "Only application files (*.msi files) can be installed into partitions."}
	ComadminELegacycompsNotAllowedIn10Format                     = &Error{0x8011081A, "COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_1_0_FORMAT", "Applications containing one or more legacy components cannot be exported to 1.0 format."}
	ComadminELegacycompsNotAllowedInNonbasePartitions            = &Error{0x8011081B, "COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_NONBASE_PARTITIONS", "Legacy components cannot exist in nonbase partitions."}
	ComadminECompMoveSource                                      = &Error{0x8011081C, "COMADMIN_E_COMP_MOVE_SOURCE", "A component cannot be moved (or copied) from the System Application, an application proxy, or a nonchangeable application."}
	ComadminECompMoveDest                                        = &Error{0x8011081D, "COMADMIN_E_COMP_MOVE_DEST", "A component cannot be moved (or copied) to the System Application, an application proxy or a nonchangeable application."}
	ComadminECompMovePrivate                                     = &Error{0x8011081E, "COMADMIN_E_COMP_MOVE_PRIVATE", "A private component cannot be moved (or copied) to a library application or to the base partition."}
	ComadminEBasepartitionRequiredInSet                          = &Error{0x8011081F, "COMADMIN_E_BASEPARTITION_REQUIRED_IN_SET", "The Base Application Partition exists in all partition sets and cannot be removed."}
	ComadminECannotAliasEventclass                               = &Error{0x80110820, "COMADMIN_E_CANNOT_ALIAS_EVENTCLASS", "Alas, Event Class components cannot be aliased."}
	ComadminEPrivateAccessdenied                                 = &Error{0x80110821, "COMADMIN_E_PRIVATE_ACCESSDENIED", "Access is denied because the component is private."}
	ComadminESaferinvalid                                        = &Error{0x80110822, "COMADMIN_E_SAFERINVALID", "The specified SAFER level is invalid."}
	ComadminERegistryAccessdenied                                = &Error{0x80110823, "COMADMIN_E_REGISTRY_ACCESSDENIED", "The specified user cannot write to the system registry."}
	ComadminEPartitionsDisabled                                  = &Error{0x80110824, "COMADMIN_E_PARTITIONS_DISABLED", "COM+ partitions are currently disabled."}
	ErrorFltNoHandlerDefined                                     = &Error{0x801F0001, "ERROR_FLT_NO_HANDLER_DEFINED", "A handler was not defined by the filter for this operation."}
	ErrorFltContextAlreadyDefined                                = &Error{0x801F0002, "ERROR_FLT_CONTEXT_ALREADY_DEFINED", "A context is already defined for this object."}
	ErrorFltInvalidAsynchronousRequest                           = &Error{0x801F0003, "ERROR_FLT_INVALID_ASYNCHRONOUS_REQUEST", "Asynchronous requests are not valid for this operation."}
	ErrorFltDisallowFastIo                                       = &Error{0x801F0004, "ERROR_FLT_DISALLOW_FAST_IO", "Disallow the Fast IO path for this operation."}
	ErrorFltInvalidNameRequest                                   = &Error{0x801F0005, "ERROR_FLT_INVALID_NAME_REQUEST", "An invalid name request was made. The name requested cannot be retrieved at this time."}
	ErrorFltNotSafeToPostOperation                               = &Error{0x801F0006, "ERROR_FLT_NOT_SAFE_TO_POST_OPERATION", "Posting this operation to a worker thread for further processing is not safe at this time because it could lead to a system deadlock."}
	ErrorFltNotInitialized                                       = &Error{0x801F0007, "ERROR_FLT_NOT_INITIALIZED", "The Filter Manager was not initialized when a filter tried to register. Be sure that the Filter Manager is being loaded as a driver."}
	ErrorFltFilterNotReady                                       = &Error{0x801F0008, "ERROR_FLT_FILTER_NOT_READY", "The filter is not ready for attachment to volumes because it has not finished initializing (FltStartFiltering has not been called)."}
	ErrorFltPostOperationCleanup                                 = &Error{0x801F0009, "ERROR_FLT_POST_OPERATION_CLEANUP", "The filter must clean up any operation-specific context at this time because it is being removed from the system before the operation is completed by the lower drivers."}
	ErrorFltInternalError                                        = &Error{0x801F000A, "ERROR_FLT_INTERNAL_ERROR", "The Filter Manager had an internal error from which it cannot recover; therefore, the operation has been failed. This is usually the result of a filter returning an invalid value from a preoperation callback."}
	ErrorFltDeletingObject                                       = &Error{0x801F000B, "ERROR_FLT_DELETING_OBJECT", "The object specified for this action is in the process of being deleted; therefore, the action requested cannot be completed at this time."}
	ErrorFltMustBeNonpagedPool                                   = &Error{0x801F000C, "ERROR_FLT_MUST_BE_NONPAGED_POOL", "Nonpaged pool must be used for this type of context."}
	ErrorFltDuplicateEntry                                       = &Error{0x801F000D, "ERROR_FLT_DUPLICATE_ENTRY", "A duplicate handler definition has been provided for an operation."}
	ErrorFltCbdqDisabled                                         = &Error{0x801F000E, "ERROR_FLT_CBDQ_DISABLED", "The callback data queue has been disabled."}
	ErrorFltDoNotAttach                                          = &Error{0x801F000F, "ERROR_FLT_DO_NOT_ATTACH", "Do not attach the filter to the volume at this time."}
	ErrorFltDoNotDetach                                          = &Error{0x801F0010, "ERROR_FLT_DO_NOT_DETACH", "Do not detach the filter from the volume at this time."}
	ErrorFltInstanceAltitudeCollision                            = &Error{0x801F0011, "ERROR_FLT_INSTANCE_ALTITUDE_COLLISION", "An instance already exists at this altitude on the volume specified."}
	ErrorFltInstanceNameCollision                                = &Error{0x801F0012, "ERROR_FLT_INSTANCE_NAME_COLLISION", "An instance already exists with this name on the volume specified."}
	ErrorFltFilterNotFound                                       = &Error{0x801F0013, "ERROR_FLT_FILTER_NOT_FOUND", "The system could not find the filter specified."}
	ErrorFltVolumeNotFound                                       = &Error{0x801F0014, "ERROR_FLT_VOLUME_NOT_FOUND", "The system could not find the volume specified."}
	ErrorFltInstanceNotFound                                     = &Error{0x801F0015, "ERROR_FLT_INSTANCE_NOT_FOUND", "The system could not find the instance specified."}
	ErrorFltContextAllocationNotFound                            = &Error{0x801F0016, "ERROR_FLT_CONTEXT_ALLOCATION_NOT_FOUND", "No registered context allocation definition was found for the given request."}
	ErrorFltInvalidContextRegistration                           = &Error{0x801F0017, "ERROR_FLT_INVALID_CONTEXT_REGISTRATION", "An invalid parameter was specified during context registration."}
	ErrorFltNameCacheMiss                                        = &Error{0x801F0018, "ERROR_FLT_NAME_CACHE_MISS", "The name requested was not found in the Filter Manager name cache and could not be retrieved from the file system."}
	ErrorFltNoDeviceObject                                       = &Error{0x801F0019, "ERROR_FLT_NO_DEVICE_OBJECT", "The requested device object does not exist for the given volume."}
	ErrorFltVolumeAlreadyMounted                                 = &Error{0x801F001A, "ERROR_FLT_VOLUME_ALREADY_MOUNTED", "The specified volume is already mounted."}
	ErrorFltAlreadyEnlisted                                      = &Error{0x801F001B, "ERROR_FLT_ALREADY_ENLISTED", "The specified Transaction Context is already enlisted in a transaction."}
	ErrorFltContextAlreadyLinked                                 = &Error{0x801F001C, "ERROR_FLT_CONTEXT_ALREADY_LINKED", "The specified context is already attached to another object."}
	ErrorFltNoWaiterForReply                                     = &Error{0x801F0020, "ERROR_FLT_NO_WAITER_FOR_REPLY", "No waiter is present for the filter's reply to this message."}
	ErrorHungDisplayDriverThread                                 = &Error{0x80260001, "ERROR_HUNG_DISPLAY_DRIVER_THREAD", "{Display Driver Stopped Responding} The %hs display driver has stopped working normally. Save your work and reboot the system to restore full display functionality. The next time you reboot the machine a dialog will be displayed giving you a chance to report this failure to Microsoft."}
	ErrorMonitorNoDescriptor                                     = &Error{0x80261001, "ERROR_MONITOR_NO_DESCRIPTOR", "Monitor descriptor could not be obtained."}
	ErrorMonitorUnknownDescriptorFormat                          = &Error{0x80261002, "ERROR_MONITOR_UNKNOWN_DESCRIPTOR_FORMAT", "Format of the obtained monitor descriptor is not supported by this release."}
	DwmECompositiondisabled                                      = &Error{0x80263001, "DWM_E_COMPOSITIONDISABLED", "{Desktop Composition is Disabled} The operation could not be completed because desktop composition is disabled."}
	DwmERemotingNotSupported                                     = &Error{0x80263002, "DWM_E_REMOTING_NOT_SUPPORTED", "{Some Desktop Composition APIs Are Not Supported While Remoting} Some desktop composition APIs are not supported while remoting. The operation is not supported while running in a remote session."}
	DwmENoRedirectionSurfaceAvailable                            = &Error{0x80263003, "DWM_E_NO_REDIRECTION_SURFACE_AVAILABLE", "{No DWM Redirection Surface is Available} The Desktop Window Manager (DWM) was unable to provide a redirection surface to complete the DirectX present."}
	DwmENotQueuingPresents                                       = &Error{0x80263004, "DWM_E_NOT_QUEUING_PRESENTS", "{DWM Is Not Queuing Presents for the Specified Window} The window specified is not currently using queued presents."}
	TpmEErrorMask                                                = &Error{0x80280000, "TPM_E_ERROR_MASK", "This is an error mask to convert Trusted Platform Module (TPM) hardware errors to Win32 errors."}
	TpmEAuthfail                                                 = &Error{0x80280001, "TPM_E_AUTHFAIL", "Authentication failed."}
	TpmEBadindex                                                 = &Error{0x80280002, "TPM_E_BADINDEX", "The index to a Platform Configuration Register (PCR), DIR, or other register is incorrect."}
	TpmEBadParameter                                             = &Error{0x80280003, "TPM_E_BAD_PARAMETER", "One or more parameters are bad."}
	TpmEAuditfailure                                             = &Error{0x80280004, "TPM_E_AUDITFAILURE", "An operation completed successfully but the auditing of that operation failed."}
	TpmEClearDisabled                                            = &Error{0x80280005, "TPM_E_CLEAR_DISABLED", "The clear disable flag is set and all clear operations now require physical access."}
	TpmEDeactivated                                              = &Error{0x80280006, "TPM_E_DEACTIVATED", "The TPM is deactivated."}
	TpmEDisabled                                                 = &Error{0x80280007, "TPM_E_DISABLED", "The TPM is disabled."}
	TpmEDisabledCmd                                              = &Error{0x80280008, "TPM_E_DISABLED_CMD", "The target command has been disabled."}
	TpmEFail                                                     = &Error{0x80280009, "TPM_E_FAIL", "The operation failed."}
	TpmEBadOrdinal                                               = &Error{0x8028000A, "TPM_E_BAD_ORDINAL", "The ordinal was unknown or inconsistent."}
	TpmEInstallDisabled                                          = &Error{0x8028000B, "TPM_E_INSTALL_DISABLED", "The ability to install an owner is disabled."}
	TpmEInvalidKeyhandle                                         = &Error{0x8028000C, "TPM_E_INVALID_KEYHANDLE", "The key handle cannot be interpreted."}
	TpmEKeynotfound                                              = &Error{0x8028000D, "TPM_E_KEYNOTFOUND", "The key handle points to an invalid key."}
	TpmEInappropriateEnc                                         = &Error{0x8028000E, "TPM_E_INAPPROPRIATE_ENC", "Unacceptable encryption scheme."}
	TpmEMigratefail                                              = &Error{0x8028000F, "TPM_E_MIGRATEFAIL", "Migration authorization failed."}
	TpmEInvalidPcrInfo                                           = &Error{0x80280010, "TPM_E_INVALID_PCR_INFO", "PCR information could not be interpreted."}
	TpmENospace                                                  = &Error{0x80280011, "TPM_E_NOSPACE", "No room to load key."}
	TpmENosrk                                                    = &Error{0x80280012, "TPM_E_NOSRK", "There is no storage root key (SRK) set."}
	TpmENotsealedBlob                                            = &Error{0x80280013, "TPM_E_NOTSEALED_BLOB", "An encrypted blob is invalid or was not created by this TPM."}
	TpmEOwnerSet                                                 = &Error{0x80280014, "TPM_E_OWNER_SET", "There is already an owner."}
	TpmEResources                                                = &Error{0x80280015, "TPM_E_RESOURCES", "The TPM has insufficient internal resources to perform the requested action."}
	TpmEShortrandom                                              = &Error{0x80280016, "TPM_E_SHORTRANDOM", "A random string was too short."}
	TpmESize                                                     = &Error{0x80280017, "TPM_E_SIZE", "The TPM does not have the space to perform the operation."}
	TpmEWrongpcrval                                              = &Error{0x80280018, "TPM_E_WRONGPCRVAL", "The named PCR value does not match the current PCR value."}
	TpmEBadParamSize                                             = &Error{0x80280019, "TPM_E_BAD_PARAM_SIZE", "The paramSize argument to the command has the incorrect value."}
	TpmEShaThread                                                = &Error{0x8028001A, "TPM_E_SHA_THREAD", "There is no existing SHA-1 thread."}
	TpmEShaError                                                 = &Error{0x8028001B, "TPM_E_SHA_ERROR", "The calculation is unable to proceed because the existing SHA-1 thread has already encountered an error."}
	TpmEFailedselftest                                           = &Error{0x8028001C, "TPM_E_FAILEDSELFTEST", "Self-test has failed and the TPM has shut down."}
	TpmEAuth2fail                                                = &Error{0x8028001D, "TPM_E_AUTH2FAIL", "The authorization for the second key in a two-key function failed authorization."}
	TpmEBadtag                                                   = &Error{0x8028001E, "TPM_E_BADTAG", "The tag value sent to for a command is invalid."}
	TpmEIoerror                                                  = &Error{0x8028001F, "TPM_E_IOERROR", "An I/O error occurred transmitting information to the TPM."}
	TpmEEncryptError                                             = &Error{0x80280020, "TPM_E_ENCRYPT_ERROR", "The encryption process had a problem."}
	TpmEDecryptError                                             = &Error{0x80280021, "TPM_E_DECRYPT_ERROR", "The decryption process did not complete."}
	TpmEInvalidAuthhandle                                        = &Error{0x80280022, "TPM_E_INVALID_AUTHHANDLE", "An invalid handle was used."}
	TpmENoEndorsement                                            = &Error{0x80280023, "TPM_E_NO_ENDORSEMENT", "The TPM does not have an endorsement key (EK) installed."}
	TpmEInvalidKeyusage                                          = &Error{0x80280024, "TPM_E_INVALID_KEYUSAGE", "The usage of a key is not allowed."}
	TpmEWrongEntitytype                                          = &Error{0x80280025, "TPM_E_WRONG_ENTITYTYPE", "The submitted entity type is not allowed."}
	TpmEInvalidPostinit                                          = &Error{0x80280026, "TPM_E_INVALID_POSTINIT", "The command was received in the wrong sequence relative to TPM_Init and a subsequent TPM_Startup."}
	TpmEInappropriateSig                                         = &Error{0x80280027, "TPM_E_INAPPROPRIATE_SIG", "Signed data cannot include additional DER information."}
	TpmEBadKeyProperty                                           = &Error{0x80280028, "TPM_E_BAD_KEY_PROPERTY", "The key properties in TPM_KEY_PARMs are not supported by this TPM."}
	TpmEBadMigration                                             = &Error{0x80280029, "TPM_E_BAD_MIGRATION", "The migration properties of this key are incorrect."}
	TpmEBadScheme                                                = &Error{0x8028002A, "TPM_E_BAD_SCHEME", "The signature or encryption scheme for this key is incorrect or not permitted in this situation."}
	TpmEBadDatasize                                              = &Error{0x8028002B, "TPM_E_BAD_DATASIZE", "The size of the data (or blob) parameter is bad or inconsistent with the referenced key."}
	TpmEBadMode                                                  = &Error{0x8028002C, "TPM_E_BAD_MODE", "A mode parameter is bad, such as capArea or subCapArea for TPM_GetCapability, physicalPresence parameter for TPM_PhysicalPresence, or migrationType for TPM_CreateMigrationBlob."}
	TpmEBadPresence                                              = &Error{0x8028002D, "TPM_E_BAD_PRESENCE", "Either the physicalPresence or physicalPresenceLock bits have the wrong value."}
	TpmEBadVersion                                               = &Error{0x8028002E, "TPM_E_BAD_VERSION", "The TPM cannot perform this version of the capability."}
	TpmENoWrapTransport                                          = &Error{0x8028002F, "TPM_E_NO_WRAP_TRANSPORT", "The TPM does not allow for wrapped transport sessions."}
	TpmEAuditfailUnsuccessful                                    = &Error{0x80280030, "TPM_E_AUDITFAIL_UNSUCCESSFUL", "TPM audit construction failed and the underlying command was returning a failure code also."}
	TpmEAuditfailSuccessful                                      = &Error{0x80280031, "TPM_E_AUDITFAIL_SUCCESSFUL", "TPM audit construction failed and the underlying command was returning success."}
	TpmENotresetable                                             = &Error{0x80280032, "TPM_E_NOTRESETABLE", "Attempt to reset a PCR that does not have the resettable attribute."}
	TpmENotlocal                                                 = &Error{0x80280033, "TPM_E_NOTLOCAL", "Attempt to reset a PCR register that requires locality and the locality modifier not part of command transport."}
	TpmEBadType                                                  = &Error{0x80280034, "TPM_E_BAD_TYPE", "Make identity blob not properly typed."}
	TpmEInvalidResource                                          = &Error{0x80280035, "TPM_E_INVALID_RESOURCE", "When saving context identified resource type does not match actual resource."}
	TpmENotfips                                                  = &Error{0x80280036, "TPM_E_NOTFIPS", "The TPM is attempting to execute a command only available when in Federal Information Processing Standards (FIPS) mode."}
	TpmEInvalidFamily                                            = &Error{0x80280037, "TPM_E_INVALID_FAMILY", "The command is attempting to use an invalid family ID."}
	TpmENoNvPermission                                           = &Error{0x80280038, "TPM_E_NO_NV_PERMISSION", "The permission to manipulate the NV storage is not available."}
	TpmERequiresSign                                             = &Error{0x80280039, "TPM_E_REQUIRES_SIGN", "The operation requires a signed command."}
	TpmEKeyNotsupported                                          = &Error{0x8028003A, "TPM_E_KEY_NOTSUPPORTED", "Wrong operation to load an NV key."}
	TpmEAuthConflict                                             = &Error{0x8028003B, "TPM_E_AUTH_CONFLICT", "NV_LoadKey blob requires both owner and blob authorization."}
	TpmEAreaLocked                                               = &Error{0x8028003C, "TPM_E_AREA_LOCKED", "The NV area is locked and not writable."}
	TpmEBadLocality                                              = &Error{0x8028003D, "TPM_E_BAD_LOCALITY", "The locality is incorrect for the attempted operation."}
	TpmEReadOnly                                                 = &Error{0x8028003E, "TPM_E_READ_ONLY", "The NV area is read-only and cannot be written to."}
	TpmEPerNowrite                                               = &Error{0x8028003F, "TPM_E_PER_NOWRITE", "There is no protection on the write to the NV area."}
	TpmEFamilycount                                              = &Error{0x80280040, "TPM_E_FAMILYCOUNT", "The family count value does not match."}
	TpmEWriteLocked                                              = &Error{0x80280041, "TPM_E_WRITE_LOCKED", "The NV area has already been written to."}
	TpmEBadAttributes                                            = &Error{0x80280042, "TPM_E_BAD_ATTRIBUTES", "The NV area attributes conflict."}
	TpmEInvalidStructure                                         = &Error{0x80280043, "TPM_E_INVALID_STRUCTURE", "The structure tag and version are invalid or inconsistent."}
	TpmEKeyOwnerControl                                          = &Error{0x80280044, "TPM_E_KEY_OWNER_CONTROL", "The key is under control of the TPM owner and can only be evicted by the TPM owner."}
	TpmEBadCounter                                               = &Error{0x80280045, "TPM_E_BAD_COUNTER", "The counter handle is incorrect."}
	TpmENotFullwrite                                             = &Error{0x80280046, "TPM_E_NOT_FULLWRITE", "The write is not a complete write of the area."}
	TpmEContextGap                                               = &Error{0x80280047, "TPM_E_CONTEXT_GAP", "The gap between saved context counts is too large."}
	TpmEMaxnvwrites                                              = &Error{0x80280048, "TPM_E_MAXNVWRITES", "The maximum number of NV writes without an owner has been exceeded."}
	TpmENooperator                                               = &Error{0x80280049, "TPM_E_NOOPERATOR", "No operator AuthData value is set."}
	TpmEResourcemissing                                          = &Error{0x8028004A, "TPM_E_RESOURCEMISSING", "The resource pointed to by context is not loaded."}
	TpmEDelegateLock                                             = &Error{0x8028004B, "TPM_E_DELEGATE_LOCK", "The delegate administration is locked."}
	TpmEDelegateFamily                                           = &Error{0x8028004C, "TPM_E_DELEGATE_FAMILY", "Attempt to manage a family other then the delegated family."}
	TpmEDelegateAdmin                                            = &Error{0x8028004D, "TPM_E_DELEGATE_ADMIN", "Delegation table management not enabled."}
	TpmETransportNotexclusive                                    = &Error{0x8028004E, "TPM_E_TRANSPORT_NOTEXCLUSIVE", "There was a command executed outside an exclusive transport session."}
	TpmEOwnerControl                                             = &Error{0x8028004F, "TPM_E_OWNER_CONTROL", "Attempt to context save an owner evict controlled key."}
	TpmEDaaResources                                             = &Error{0x80280050, "TPM_E_DAA_RESOURCES", "The DAA command has no resources available to execute the command."}
	TpmEDaaInputData0                                            = &Error{0x80280051, "TPM_E_DAA_INPUT_DATA0", "The consistency check on DAA parameter inputData0 has failed."}
	TpmEDaaInputData1                                            = &Error{0x80280052, "TPM_E_DAA_INPUT_DATA1", "The consistency check on DAA parameter inputData1 has failed."}
	TpmEDaaIssuerSettings                                        = &Error{0x80280053, "TPM_E_DAA_ISSUER_SETTINGS", "The consistency check on DAA_issuerSettings has failed."}
	TpmEDaaTpmSettings                                           = &Error{0x80280054, "TPM_E_DAA_TPM_SETTINGS", "The consistency check on DAA_tpmSpecific has failed."}
	TpmEDaaStage                                                 = &Error{0x80280055, "TPM_E_DAA_STAGE", "The atomic process indicated by the submitted DAA command is not the expected process."}
	TpmEDaaIssuerValidity                                        = &Error{0x80280056, "TPM_E_DAA_ISSUER_VALIDITY", "The issuer's validity check has detected an inconsistency."}
	TpmEDaaWrongW                                                = &Error{0x80280057, "TPM_E_DAA_WRONG_W", "The consistency check on w has failed."}
	TpmEBadHandle                                                = &Error{0x80280058, "TPM_E_BAD_HANDLE", "The handle is incorrect."}
	TpmEBadDelegate                                              = &Error{0x80280059, "TPM_E_BAD_DELEGATE", "Delegation is not correct."}
	TpmEBadcontext                                               = &Error{0x8028005A, "TPM_E_BADCONTEXT", "The context blob is invalid."}
	TpmEToomanycontexts                                          = &Error{0x8028005B, "TPM_E_TOOMANYCONTEXTS", "Too many contexts held by the TPM."}
	TpmEMaTicketSignature                                        = &Error{0x8028005C, "TPM_E_MA_TICKET_SIGNATURE", "Migration authority signature validation failure."}
	TpmEMaDestination                                            = &Error{0x8028005D, "TPM_E_MA_DESTINATION", "Migration destination not authenticated."}
	TpmEMaSource                                                 = &Error{0x8028005E, "TPM_E_MA_SOURCE", "Migration source incorrect."}
	TpmEMaAuthority                                              = &Error{0x8028005F, "TPM_E_MA_AUTHORITY", "Incorrect migration authority."}
	TpmEPermanentek                                              = &Error{0x80280061, "TPM_E_PERMANENTEK", "Attempt to revoke the EK and the EK is not revocable."}
	TpmEBadSignature                                             = &Error{0x80280062, "TPM_E_BAD_SIGNATURE", "Bad signature of CMK ticket."}
	TpmENocontextspace                                           = &Error{0x80280063, "TPM_E_NOCONTEXTSPACE", "There is no room in the context list for additional contexts."}
	TpmECommandBlocked                                           = &Error{0x80280400, "TPM_E_COMMAND_BLOCKED", "The command was blocked."}
	TpmEInvalidHandle                                            = &Error{0x80280401, "TPM_E_INVALID_HANDLE", "The specified handle was not found."}
	TpmEDuplicateVhandle                                         = &Error{0x80280402, "TPM_E_DUPLICATE_VHANDLE", "The TPM returned a duplicate handle and the command needs to be resubmitted."}
	TpmEEmbeddedCommandBlocked                                   = &Error{0x80280403, "TPM_E_EMBEDDED_COMMAND_BLOCKED", "The command within the transport was blocked."}
	TpmEEmbeddedCommandUnsupported                               = &Error{0x80280404, "TPM_E_EMBEDDED_COMMAND_UNSUPPORTED", "The command within the transport is not supported."}
	TpmERetry                                                    = &Error{0x80280800, "TPM_E_RETRY", "The TPM is too busy to respond to the command immediately, but the command could be resubmitted at a later time."}
	TpmENeedsSelftest                                            = &Error{0x80280801, "TPM_E_NEEDS_SELFTEST", "SelfTestFull has not been run."}
	TpmEDoingSelftest                                            = &Error{0x80280802, "TPM_E_DOING_SELFTEST", "The TPM is currently executing a full self-test."}
	TpmEDefendLockRunning                                        = &Error{0x80280803, "TPM_E_DEFEND_LOCK_RUNNING", "The TPM is defending against dictionary attacks and is in a time-out period."}
	TbsEInternalError                                            = &Error{0x80284001, "TBS_E_INTERNAL_ERROR", "An internal software error has been detected."}
	TbsEBadParameter                                             = &Error{0x80284002, "TBS_E_BAD_PARAMETER", "One or more input parameters are bad."}
	TbsEInvalidOutputPointer                                     = &Error{0x80284003, "TBS_E_INVALID_OUTPUT_POINTER", "A specified output pointer is bad."}
	TbsEInvalidContext                                           = &Error{0x80284004, "TBS_E_INVALID_CONTEXT", "The specified context handle does not refer to a valid context."}
	TbsEInsufficientBuffer                                       = &Error{0x80284005, "TBS_E_INSUFFICIENT_BUFFER", "A specified output buffer is too small."}
	TbsEIoerror                                                  = &Error{0x80284006, "TBS_E_IOERROR", "An error occurred while communicating with the TPM."}
	TbsEInvalidContextParam                                      = &Error{0x80284007, "TBS_E_INVALID_CONTEXT_PARAM", "One or more context parameters are invalid."}
	TbsEServiceNotRunning                                        = &Error{0x80284008, "TBS_E_SERVICE_NOT_RUNNING", "The TPM Base Services (TBS) is not running and could not be started."}
	TbsETooManyTbsContexts                                       = &Error{0x80284009, "TBS_E_TOO_MANY_TBS_CONTEXTS", "A new context could not be created because there are too many open contexts."}
	TbsETooManyResources                                         = &Error{0x8028400A, "TBS_E_TOO_MANY_RESOURCES", "A new virtual resource could not be created because there are too many open virtual resources."}
	TbsEServiceStartPending                                      = &Error{0x8028400B, "TBS_E_SERVICE_START_PENDING", "The TBS service has been started but is not yet running."}
	TbsEPpiNotSupported                                          = &Error{0x8028400C, "TBS_E_PPI_NOT_SUPPORTED", "The physical presence interface is not supported."}
	TbsECommandCanceled                                          = &Error{0x8028400D, "TBS_E_COMMAND_CANCELED", "The command was canceled."}
	TbsEBufferTooLarge                                           = &Error{0x8028400E, "TBS_E_BUFFER_TOO_LARGE", "The input or output buffer is too large."}
	TpmapiEInvalidState                                          = &Error{0x80290100, "TPMAPI_E_INVALID_STATE", "The command buffer is not in the correct state."}
	TpmapiENotEnoughData                                         = &Error{0x80290101, "TPMAPI_E_NOT_ENOUGH_DATA", "The command buffer does not contain enough data to satisfy the request."}
	TpmapiETooMuchData                                           = &Error{0x80290102, "TPMAPI_E_TOO_MUCH_DATA", "The command buffer cannot contain any more data."}
	TpmapiEInvalidOutputPointer                                  = &Error{0x80290103, "TPMAPI_E_INVALID_OUTPUT_POINTER", "One or more output parameters was null or invalid."}
	TpmapiEInvalidParameter                                      = &Error{0x80290104, "TPMAPI_E_INVALID_PARAMETER", "One or more input parameters are invalid."}
	TpmapiEOutOfMemory                                           = &Error{0x80290105, "TPMAPI_E_OUT_OF_MEMORY", "Not enough memory was available to satisfy the request."}
	TpmapiEBufferTooSmall                                        = &Error{0x80290106, "TPMAPI_E_BUFFER_TOO_SMALL", "The specified buffer was too small."}
	TpmapiEInternalError                                         = &Error{0x80290107, "TPMAPI_E_INTERNAL_ERROR", "An internal error was detected."}
	TpmapiEAccessDenied                                          = &Error{0x80290108, "TPMAPI_E_ACCESS_DENIED", "The caller does not have the appropriate rights to perform the requested operation."}
	TpmapiEAuthorizationFailed                                   = &Error{0x80290109, "TPMAPI_E_AUTHORIZATION_FAILED", "The specified authorization information was invalid."}
	TpmapiEInvalidContextHandle                                  = &Error{0x8029010A, "TPMAPI_E_INVALID_CONTEXT_HANDLE", "The specified context handle was not valid."}
	TpmapiETbsCommunicationError                                 = &Error{0x8029010B, "TPMAPI_E_TBS_COMMUNICATION_ERROR", "An error occurred while communicating with the TBS."}
	TpmapiETpmCommandError                                       = &Error{0x8029010C, "TPMAPI_E_TPM_COMMAND_ERROR", "The TPM returned an unexpected result."}
	TpmapiEMessageTooLarge                                       = &Error{0x8029010D, "TPMAPI_E_MESSAGE_TOO_LARGE", "The message was too large for the encoding scheme."}
	TpmapiEInvalidEncoding                                       = &Error{0x8029010E, "TPMAPI_E_INVALID_ENCODING", "The encoding in the binary large object (BLOB) was not recognized."}
	TpmapiEInvalidKeySize                                        = &Error{0x8029010F, "TPMAPI_E_INVALID_KEY_SIZE", "The key size is not valid."}
	TpmapiEEncryptionFailed                                      = &Error{0x80290110, "TPMAPI_E_ENCRYPTION_FAILED", "The encryption operation failed."}
	TpmapiEInvalidKeyParams                                      = &Error{0x80290111, "TPMAPI_E_INVALID_KEY_PARAMS", "The key parameters structure was not valid."}
	TpmapiEInvalidMigrationAuthorizationBlob                     = &Error{0x80290112, "TPMAPI_E_INVALID_MIGRATION_AUTHORIZATION_BLOB", "The requested supplied data does not appear to be a valid migration authorization BLOB."}
	TpmapiEInvalidPcrIndex                                       = &Error{0x80290113, "TPMAPI_E_INVALID_PCR_INDEX", "The specified PCR index was invalid."}
	TpmapiEInvalidDelegateBlob                                   = &Error{0x80290114, "TPMAPI_E_INVALID_DELEGATE_BLOB", "The data given does not appear to be a valid delegate BLOB."}
	TpmapiEInvalidContextParams                                  = &Error{0x80290115, "TPMAPI_E_INVALID_CONTEXT_PARAMS", "One or more of the specified context parameters was not valid."}
	TpmapiEInvalidKeyBlob                                        = &Error{0x80290116, "TPMAPI_E_INVALID_KEY_BLOB", "The data given does not appear to be a valid key BLOB."}
	TpmapiEInvalidPcrData                                        = &Error{0x80290117, "TPMAPI_E_INVALID_PCR_DATA", "The specified PCR data was invalid."}
	TpmapiEInvalidOwnerAuth                                      = &Error{0x80290118, "TPMAPI_E_INVALID_OWNER_AUTH", "The format of the owner authorization data was invalid."}
	TbsimpEBufferTooSmall                                        = &Error{0x80290200, "TBSIMP_E_BUFFER_TOO_SMALL", "The specified buffer was too small."}
	TbsimpECleanupFailed                                         = &Error{0x80290201, "TBSIMP_E_CLEANUP_FAILED", "The context could not be cleaned up."}
	TbsimpEInvalidContextHandle                                  = &Error{0x80290202, "TBSIMP_E_INVALID_CONTEXT_HANDLE", "The specified context handle is invalid."}
	TbsimpEInvalidContextParam                                   = &Error{0x80290203, "TBSIMP_E_INVALID_CONTEXT_PARAM", "An invalid context parameter was specified."}
	TbsimpETpmError                                              = &Error{0x80290204, "TBSIMP_E_TPM_ERROR", "An error occurred while communicating with the TPM."}
	TbsimpEHashBadKey                                            = &Error{0x80290205, "TBSIMP_E_HASH_BAD_KEY", "No entry with the specified key was found."}
	TbsimpEDuplicateVhandle                                      = &Error{0x80290206, "TBSIMP_E_DUPLICATE_VHANDLE", "The specified virtual handle matches a virtual handle already in use."}
	TbsimpEInvalidOutputPointer                                  = &Error{0x80290207, "TBSIMP_E_INVALID_OUTPUT_POINTER", "The pointer to the returned handle location was null or invalid."}
	TbsimpEInvalidParameter                                      = &Error{0x80290208, "TBSIMP_E_INVALID_PARAMETER", "One or more parameters are invalid."}
	TbsimpERpcInitFailed                                         = &Error{0x80290209, "TBSIMP_E_RPC_INIT_FAILED", "The RPC subsystem could not be initialized."}
	TbsimpESchedulerNotRunning                                   = &Error{0x8029020A, "TBSIMP_E_SCHEDULER_NOT_RUNNING", "The TBS scheduler is not running."}
	TbsimpECommandCanceled                                       = &Error{0x8029020B, "TBSIMP_E_COMMAND_CANCELED", "The command was canceled."}
	TbsimpEOutOfMemory                                           = &Error{0x8029020C, "TBSIMP_E_OUT_OF_MEMORY", "There was not enough memory to fulfill the request."}
	TbsimpEListNoMoreItems                                       = &Error{0x8029020D, "TBSIMP_E_LIST_NO_MORE_ITEMS", "The specified list is empty, or the iteration has reached the end of the list."}
	TbsimpEListNotFound                                          = &Error{0x8029020E, "TBSIMP_E_LIST_NOT_FOUND", "The specified item was not found in the list."}
	TbsimpENotEnoughSpace                                        = &Error{0x8029020F, "TBSIMP_E_NOT_ENOUGH_SPACE", "The TPM does not have enough space to load the requested resource."}
	TbsimpENotEnoughTpmContexts                                  = &Error{0x80290210, "TBSIMP_E_NOT_ENOUGH_TPM_CONTEXTS", "There are too many TPM contexts in use."}
	TbsimpECommandFailed                                         = &Error{0x80290211, "TBSIMP_E_COMMAND_FAILED", "The TPM command failed."}
	TbsimpEUnknownOrdinal                                        = &Error{0x80290212, "TBSIMP_E_UNKNOWN_ORDINAL", "The TBS does not recognize the specified ordinal."}
	TbsimpEResourceExpired                                       = &Error{0x80290213, "TBSIMP_E_RESOURCE_EXPIRED", "The requested resource is no longer available."}
	TbsimpEInvalidResource                                       = &Error{0x80290214, "TBSIMP_E_INVALID_RESOURCE", "The resource type did not match."}
	TbsimpENothingToUnload                                       = &Error{0x80290215, "TBSIMP_E_NOTHING_TO_UNLOAD", "No resources can be unloaded."}
	TbsimpEHashTableFull                                         = &Error{0x80290216, "TBSIMP_E_HASH_TABLE_FULL", "No new entries can be added to the hash table."}
	TbsimpETooManyTbsContexts                                    = &Error{0x80290217, "TBSIMP_E_TOO_MANY_TBS_CONTEXTS", "A new TBS context could not be created because there are too many open contexts."}
	TbsimpETooManyResources                                      = &Error{0x80290218, "TBSIMP_E_TOO_MANY_RESOURCES", "A new virtual resource could not be created because there are too many open virtual resources."}
	TbsimpEPpiNotSupported                                       = &Error{0x80290219, "TBSIMP_E_PPI_NOT_SUPPORTED", "The physical presence interface is not supported."}
	TbsimpETpmIncompatible                                       = &Error{0x8029021A, "TBSIMP_E_TPM_INCOMPATIBLE", "TBS is not compatible with the version of TPM found on the system."}
	TpmEPpiAcpiFailure                                           = &Error{0x80290300, "TPM_E_PPI_ACPI_FAILURE", "A general error was detected when attempting to acquire the BIOS response to a physical presence command."}
	TpmEPpiUserAbort                                             = &Error{0x80290301, "TPM_E_PPI_USER_ABORT", "The user failed to confirm the TPM operation request."}
	TpmEPpiBiosFailure                                           = &Error{0x80290302, "TPM_E_PPI_BIOS_FAILURE", "The BIOS failure prevented the successful execution of the requested TPM operation (for example, invalid TPM operation request, BIOS communication error with the TPM)."}
	TpmEPpiNotSupported                                          = &Error{0x80290303, "TPM_E_PPI_NOT_SUPPORTED", "The BIOS does not support the physical presence interface."}
	PlaEDcsNotFound                                              = &Error{0x80300002, "PLA_E_DCS_NOT_FOUND", "A Data Collector Set was not found."}
	PlaETooManyFolders                                           = &Error{0x80300045, "PLA_E_TOO_MANY_FOLDERS", "Unable to start Data Collector Set because there are too many folders."}
	PlaENoMinDisk                                                = &Error{0x80300070, "PLA_E_NO_MIN_DISK", "Not enough free disk space to start Data Collector Set."}
	PlaEDcsInUse                                                 = &Error{0x803000AA, "PLA_E_DCS_IN_USE", "Data Collector Set is in use."}
	PlaEDcsAlreadyExists                                         = &Error{0x803000B7, "PLA_E_DCS_ALREADY_EXISTS", "Data Collector Set already exists."}
	PlaEPropertyConflict                                         = &Error{0x80300101, "PLA_E_PROPERTY_CONFLICT", "Property value conflict."}
	PlaEDcsSingletonRequired                                     = &Error{0x80300102, "PLA_E_DCS_SINGLETON_REQUIRED", "The current configuration for this Data Collector Set requires that it contain exactly one Data Collector."}
	PlaECredentialsRequired                                      = &Error{0x80300103, "PLA_E_CREDENTIALS_REQUIRED", "A user account is required to commit the current Data Collector Set properties."}
	PlaEDcsNotRunning                                            = &Error{0x80300104, "PLA_E_DCS_NOT_RUNNING", "Data Collector Set is not running."}
	PlaEConflictInclExclApi                                      = &Error{0x80300105, "PLA_E_CONFLICT_INCL_EXCL_API", "A conflict was detected in the list of include and exclude APIs. Do not specify the same API in both the include list and the exclude list."}
	PlaENetworkExeNotValid                                       = &Error{0x80300106, "PLA_E_NETWORK_EXE_NOT_VALID", "The executable path specified refers to a network share or UNC path."}
	PlaEExeAlreadyConfigured                                     = &Error{0x80300107, "PLA_E_EXE_ALREADY_CONFIGURED", "The executable path specified is already configured for API tracing."}
	PlaEExePathNotValid                                          = &Error{0x80300108, "PLA_E_EXE_PATH_NOT_VALID", "The executable path specified does not exist. Verify that the specified path is correct."}
	PlaEDcAlreadyExists                                          = &Error{0x80300109, "PLA_E_DC_ALREADY_EXISTS", "Data Collector already exists."}
	PlaEDcsStartWaitTimeout                                      = &Error{0x8030010A, "PLA_E_DCS_START_WAIT_TIMEOUT", "The wait for the Data Collector Set start notification has timed out."}
	PlaEDcStartWaitTimeout                                       = &Error{0x8030010B, "PLA_E_DC_START_WAIT_TIMEOUT", "The wait for the Data Collector to start has timed out."}
	PlaEReportWaitTimeout                                        = &Error{0x8030010C, "PLA_E_REPORT_WAIT_TIMEOUT", "The wait for the report generation tool to finish has timed out."}
	PlaENoDuplicates                                             = &Error{0x8030010D, "PLA_E_NO_DUPLICATES", "Duplicate items are not allowed."}
	PlaEExeFullPathRequired                                      = &Error{0x8030010E, "PLA_E_EXE_FULL_PATH_REQUIRED", "When specifying the executable to trace, you must specify a full path to the executable and not just a file name."}
	PlaEInvalidSessionName                                       = &Error{0x8030010F, "PLA_E_INVALID_SESSION_NAME", "The session name provided is invalid."}
	PlaEPlaChannelNotEnabled                                     = &Error{0x80300110, "PLA_E_PLA_CHANNEL_NOT_ENABLED", "The Event Log channel Microsoft-Windows-Diagnosis-PLA/Operational must be enabled to perform this operation."}
	PlaETaskschedChannelNotEnabled                               = &Error{0x80300111, "PLA_E_TASKSCHED_CHANNEL_NOT_ENABLED", "The Event Log channel Microsoft-Windows-TaskScheduler must be enabled to perform this operation."}
	FveELockedVolume                                             = &Error{0x80310000, "FVE_E_LOCKED_VOLUME", "The volume must be unlocked before it can be used."}
	FveENotEncrypted                                             = &Error{0x80310001, "FVE_E_NOT_ENCRYPTED", "The volume is fully decrypted and no key is available."}
	FveENoTpmBios                                                = &Error{0x80310002, "FVE_E_NO_TPM_BIOS", "The firmware does not support using a TPM during boot."}
	FveENoMbrMetric                                              = &Error{0x80310003, "FVE_E_NO_MBR_METRIC", "The firmware does not use a TPM to perform initial program load (IPL) measurement."}
	FveENoBootsectorMetric                                       = &Error{0x80310004, "FVE_E_NO_BOOTSECTOR_METRIC", "The master boot record (MBR) is not TPM-aware."}
	FveENoBootmgrMetric                                          = &Error{0x80310005, "FVE_E_NO_BOOTMGR_METRIC", "The BOOTMGR is not being measured by the TPM."}
	FveEWrongBootmgr                                             = &Error{0x80310006, "FVE_E_WRONG_BOOTMGR", "The BOOTMGR component does not perform expected TPM measurements."}
	FveESecureKeyRequired                                        = &Error{0x80310007, "FVE_E_SECURE_KEY_REQUIRED", "No secure key protection mechanism has been defined."}
	FveENotActivated                                             = &Error{0x80310008, "FVE_E_NOT_ACTIVATED", "This volume has not been provisioned for encryption."}
	FveEActionNotAllowed                                         = &Error{0x80310009, "FVE_E_ACTION_NOT_ALLOWED", "Requested action was denied by the full-volume encryption (FVE) control engine."}
	FveEAdSchemaNotInstalled                                     = &Error{0x8031000A, "FVE_E_AD_SCHEMA_NOT_INSTALLED", "The Active Directory forest does not contain the required attributes and classes to host FVE or TPM information."}
	FveEAdInvalidDatatype                                        = &Error{0x8031000B, "FVE_E_AD_INVALID_DATATYPE", "The type of data obtained from Active Directory was not expected."}
	FveEAdInvalidDatasize                                        = &Error{0x8031000C, "FVE_E_AD_INVALID_DATASIZE", "The size of the data obtained from Active Directory was not expected."}
	FveEAdNoValues                                               = &Error{0x8031000D, "FVE_E_AD_NO_VALUES", "The attribute read from Active Directory has no (zero) values."}
	FveEAdAttrNotSet                                             = &Error{0x8031000E, "FVE_E_AD_ATTR_NOT_SET", "The attribute was not set."}
	FveEAdGuidNotFound                                           = &Error{0x8031000F, "FVE_E_AD_GUID_NOT_FOUND", "The specified GUID could not be found."}
	FveEBadInformation                                           = &Error{0x80310010, "FVE_E_BAD_INFORMATION", "The control block for the encrypted volume is not valid."}
	FveETooSmall                                                 = &Error{0x80310011, "FVE_E_TOO_SMALL", "Not enough free space remaining on volume to allow encryption."}
	FveESystemVolume                                             = &Error{0x80310012, "FVE_E_SYSTEM_VOLUME", "The volume cannot be encrypted because it is required to boot the operating system."}
	FveEFailedWrongFs                                            = &Error{0x80310013, "FVE_E_FAILED_WRONG_FS", "The volume cannot be encrypted because the file system is not supported."}
	FveEFailedBadFs                                              = &Error{0x80310014, "FVE_E_FAILED_BAD_FS", "The file system is inconsistent. Run CHKDSK."}
	FveENotSupported                                             = &Error{0x80310015, "FVE_E_NOT_SUPPORTED", "This volume cannot be encrypted."}
	FveEBadData                                                  = &Error{0x80310016, "FVE_E_BAD_DATA", "Data supplied is malformed."}
	FveEVolumeNotBound                                           = &Error{0x80310017, "FVE_E_VOLUME_NOT_BOUND", "Volume is not bound to the system."}
	FveETpmNotOwned                                              = &Error{0x80310018, "FVE_E_TPM_NOT_OWNED", "TPM must be owned before a volume can be bound to it."}
	FveENotDataVolume                                            = &Error{0x80310019, "FVE_E_NOT_DATA_VOLUME", "The volume specified is not a data volume."}
	FveEAdInsufficientBuffer                                     = &Error{0x8031001A, "FVE_E_AD_INSUFFICIENT_BUFFER", "The buffer supplied to a function was insufficient to contain the returned data."}
	FveEConvRead                                                 = &Error{0x8031001B, "FVE_E_CONV_READ", "A read operation failed while converting the volume."}
	FveEConvWrite                                                = &Error{0x8031001C, "FVE_E_CONV_WRITE", "A write operation failed while converting the volume."}
	FveEKeyRequired                                              = &Error{0x8031001D, "FVE_E_KEY_REQUIRED", "One or more key protection mechanisms are required for this volume."}
	FveEClusteringNotSupported                                   = &Error{0x8031001E, "FVE_E_CLUSTERING_NOT_SUPPORTED", "Cluster configurations are not supported."}
	FveEVolumeBoundAlready                                       = &Error{0x8031001F, "FVE_E_VOLUME_BOUND_ALREADY", "The volume is already bound to the system."}
	FveEOsNotProtected                                           = &Error{0x80310020, "FVE_E_OS_NOT_PROTECTED", "The boot OS volume is not being protected via FVE."}
	FveEProtectionDisabled                                       = &Error{0x80310021, "FVE_E_PROTECTION_DISABLED", "All protection mechanisms are effectively disabled (clear key exists)."}
	FveERecoveryKeyRequired                                      = &Error{0x80310022, "FVE_E_RECOVERY_KEY_REQUIRED", "A recovery key protection mechanism is required."}
	FveEForeignVolume                                            = &Error{0x80310023, "FVE_E_FOREIGN_VOLUME", "This volume cannot be bound to a TPM."}
	FveEOverlappedUpdate                                         = &Error{0x80310024, "FVE_E_OVERLAPPED_UPDATE", "The control block for the encrypted volume was updated by another thread. Try again."}
	FveETpmSrkAuthNotZero                                        = &Error{0x80310025, "FVE_E_TPM_SRK_AUTH_NOT_ZERO", "The SRK authentication of the TPM is not zero and, therefore, is not compatible."}
	FveEFailedSectorSize                                         = &Error{0x80310026, "FVE_E_FAILED_SECTOR_SIZE", "The volume encryption algorithm cannot be used on this sector size."}
	FveEFailedAuthentication                                     = &Error{0x80310027, "FVE_E_FAILED_AUTHENTICATION", "BitLocker recovery authentication failed."}
	FveENotOsVolume                                              = &Error{0x80310028, "FVE_E_NOT_OS_VOLUME", "The volume specified is not the boot OS volume."}
	FveEAutounlockEnabled                                        = &Error{0x80310029, "FVE_E_AUTOUNLOCK_ENABLED", "Auto-unlock information for data volumes is present on the boot OS volume."}
	FveEWrongBootsector                                          = &Error{0x8031002A, "FVE_E_WRONG_BOOTSECTOR", "The system partition boot sector does not perform TPM measurements."}
	FveEWrongSystemFs                                            = &Error{0x8031002B, "FVE_E_WRONG_SYSTEM_FS", "The system partition file system must be NTFS."}
	FveEPolicyPasswordRequired                                   = &Error{0x8031002C, "FVE_E_POLICY_PASSWORD_REQUIRED", "Group policy requires a recovery password before encryption can begin."}
	FveECannotSetFvekEncrypted                                   = &Error{0x8031002D, "FVE_E_CANNOT_SET_FVEK_ENCRYPTED", "The volume encryption algorithm and key cannot be set on an encrypted volume."}
	FveECannotEncryptNoKey                                       = &Error{0x8031002E, "FVE_E_CANNOT_ENCRYPT_NO_KEY", "A key must be specified before encryption can begin."}
	FveEBootableCddvd                                            = &Error{0x80310030, "FVE_E_BOOTABLE_CDDVD", "A bootable CD/DVD is in the system. Remove the CD/DVD and reboot the system."}
	FveEProtectorExists                                          = &Error{0x80310031, "FVE_E_PROTECTOR_EXISTS", "An instance of this key protector already exists on the volume."}
	FveERelativePath                                             = &Error{0x80310032, "FVE_E_RELATIVE_PATH", "The file cannot be saved to a relative path."}
	FwpECalloutNotFound                                          = &Error{0x80320001, "FWP_E_CALLOUT_NOT_FOUND", "The callout does not exist."}
	FwpEConditionNotFound                                        = &Error{0x80320002, "FWP_E_CONDITION_NOT_FOUND", "The filter condition does not exist."}
	FwpEFilterNotFound                                           = &Error{0x80320003, "FWP_E_FILTER_NOT_FOUND", "The filter does not exist."}
	FwpELayerNotFound                                            = &Error{0x80320004, "FWP_E_LAYER_NOT_FOUND", "The layer does not exist."}
	FwpEProviderNotFound                                         = &Error{0x80320005, "FWP_E_PROVIDER_NOT_FOUND", "The provider does not exist."}
	FwpEProviderContextNotFound                                  = &Error{0x80320006, "FWP_E_PROVIDER_CONTEXT_NOT_FOUND", "The provider context does not exist."}
	FwpESublayerNotFound                                         = &Error{0x80320007, "FWP_E_SUBLAYER_NOT_FOUND", "The sublayer does not exist."}
	FwpENotFound                                                 = &Error{0x80320008, "FWP_E_NOT_FOUND", "The object does not exist."}
	FwpEAlreadyExists                                            = &Error{0x80320009, "FWP_E_ALREADY_EXISTS", "An object with that GUID or LUID already exists."}
	FwpEInUse                                                    = &Error{0x8032000A, "FWP_E_IN_USE", "The object is referenced by other objects and, therefore, cannot be deleted."}
	FwpEDynamicSessionInProgress                                 = &Error{0x8032000B, "FWP_E_DYNAMIC_SESSION_IN_PROGRESS", "The call is not allowed from within a dynamic session."}
	FwpEWrongSession                                             = &Error{0x8032000C, "FWP_E_WRONG_SESSION", "The call was made from the wrong session and, therefore, cannot be completed."}
	FwpENoTxnInProgress                                          = &Error{0x8032000D, "FWP_E_NO_TXN_IN_PROGRESS", "The call must be made from within an explicit transaction."}
	FwpETxnInProgress                                            = &Error{0x8032000E, "FWP_E_TXN_IN_PROGRESS", "The call is not allowed from within an explicit transaction."}
	FwpETxnAborted                                               = &Error{0x8032000F, "FWP_E_TXN_ABORTED", "The explicit transaction has been forcibly canceled."}
	FwpESessionAborted                                           = &Error{0x80320010, "FWP_E_SESSION_ABORTED", "The session has been canceled."}
	FwpEIncompatibleTxn                                          = &Error{0x80320011, "FWP_E_INCOMPATIBLE_TXN", "The call is not allowed from within a read-only transaction."}
	FwpETimeout                                                  = &Error{0x80320012, "FWP_E_TIMEOUT", "The call timed out while waiting to acquire the transaction lock."}
	FwpENetEventsDisabled                                        = &Error{0x80320013, "FWP_E_NET_EVENTS_DISABLED", "Collection of network diagnostic events is disabled."}
	FwpEIncompatibleLayer                                        = &Error{0x80320014, "FWP_E_INCOMPATIBLE_LAYER", "The operation is not supported by the specified layer."}
	FwpEKmClientsOnly                                            = &Error{0x80320015, "FWP_E_KM_CLIENTS_ONLY", "The call is allowed for kernel-mode callers only."}
	FwpELifetimeMismatch                                         = &Error{0x80320016, "FWP_E_LIFETIME_MISMATCH", "The call tried to associate two objects with incompatible lifetimes."}
	FwpEBuiltinObject                                            = &Error{0x80320017, "FWP_E_BUILTIN_OBJECT", "The object is built in and, therefore, cannot be deleted."}
	FwpETooManyBoottimeFilters                                   = &Error{0x80320018, "FWP_E_TOO_MANY_BOOTTIME_FILTERS", "The maximum number of boot-time filters has been reached."}
	FwpENotificationDropped                                      = &Error{0x80320019, "FWP_E_NOTIFICATION_DROPPED", "A notification could not be delivered because a message queue is at its maximum capacity."}
	FwpETrafficMismatch                                          = &Error{0x8032001A, "FWP_E_TRAFFIC_MISMATCH", "The traffic parameters do not match those for the security association context."}
	FwpEIncompatibleSaState                                      = &Error{0x8032001B, "FWP_E_INCOMPATIBLE_SA_STATE", "The call is not allowed for the current security association state."}
	FwpENullPointer                                              = &Error{0x8032001C, "FWP_E_NULL_POINTER", "A required pointer is null."}
	FwpEInvalidEnumerator                                        = &Error{0x8032001D, "FWP_E_INVALID_ENUMERATOR", "An enumerator is not valid."}
	FwpEInvalidFlags                                             = &Error{0x8032001E, "FWP_E_INVALID_FLAGS", "The flags field contains an invalid value."}
	FwpEInvalidNetMask                                           = &Error{0x8032001F, "FWP_E_INVALID_NET_MASK", "A network mask is not valid."}
	FwpEInvalidRange                                             = &Error{0x80320020, "FWP_E_INVALID_RANGE", "An FWP_RANGE is not valid."}
	FwpEInvalidInterval                                          = &Error{0x80320021, "FWP_E_INVALID_INTERVAL", "The time interval is not valid."}
	FwpEZeroLengthArray                                          = &Error{0x80320022, "FWP_E_ZERO_LENGTH_ARRAY", "An array that must contain at least one element that is zero-length."}
	FwpENullDisplayName                                          = &Error{0x80320023, "FWP_E_NULL_DISPLAY_NAME", "The displayData.name field cannot be null."}
	FwpEInvalidActionType                                        = &Error{0x80320024, "FWP_E_INVALID_ACTION_TYPE", "The action type is not one of the allowed action types for a filter."}
	FwpEInvalidWeight                                            = &Error{0x80320025, "FWP_E_INVALID_WEIGHT", "The filter weight is not valid."}
	FwpEMatchTypeMismatch                                        = &Error{0x80320026, "FWP_E_MATCH_TYPE_MISMATCH", "A filter condition contains a match type that is not compatible with the operands."}
	FwpETypeMismatch                                             = &Error{0x80320027, "FWP_E_TYPE_MISMATCH", "An FWP_VALUE or FWPM_CONDITION_VALUE is of the wrong type."}
	FwpEOutOfBounds                                              = &Error{0x80320028, "FWP_E_OUT_OF_BOUNDS", "An integer value is outside the allowed range."}
	FwpEReserved                                                 = &Error{0x80320029, "FWP_E_RESERVED", "A reserved field is nonzero."}
	FwpEDuplicateCondition                                       = &Error{0x8032002A, "FWP_E_DUPLICATE_CONDITION", "A filter cannot contain multiple conditions operating on a single field."}
	FwpEDuplicateKeymod                                          = &Error{0x8032002B, "FWP_E_DUPLICATE_KEYMOD", "A policy cannot contain the same keying module more than once."}
	FwpEActionIncompatibleWithLayer                              = &Error{0x8032002C, "FWP_E_ACTION_INCOMPATIBLE_WITH_LAYER", "The action type is not compatible with the layer."}
	FwpEActionIncompatibleWithSublayer                           = &Error{0x8032002D, "FWP_E_ACTION_INCOMPATIBLE_WITH_SUBLAYER", "The action type is not compatible with the sublayer."}
	FwpEContextIncompatibleWithLayer                             = &Error{0x8032002E, "FWP_E_CONTEXT_INCOMPATIBLE_WITH_LAYER", "The raw context or the provider context is not compatible with the layer."}
	FwpEContextIncompatibleWithCallout                           = &Error{0x8032002F, "FWP_E_CONTEXT_INCOMPATIBLE_WITH_CALLOUT", "The raw context or the provider context is not compatible with the callout."}
	FwpEIncompatibleAuthMethod                                   = &Error{0x80320030, "FWP_E_INCOMPATIBLE_AUTH_METHOD", "The authentication method is not compatible with the policy type."}
	FwpEIncompatibleDhGroup                                      = &Error{0x80320031, "FWP_E_INCOMPATIBLE_DH_GROUP", "The Diffie-Hellman group is not compatible with the policy type."}
	FwpEEmNotSupported                                           = &Error{0x80320032, "FWP_E_EM_NOT_SUPPORTED", "An Internet Key Exchange (IKE) policy cannot contain an Extended Mode policy."}
	FwpENeverMatch                                               = &Error{0x80320033, "FWP_E_NEVER_MATCH", "The enumeration template or subscription will never match any objects."}
	FwpEProviderContextMismatch                                  = &Error{0x80320034, "FWP_E_PROVIDER_CONTEXT_MISMATCH", "The provider context is of the wrong type."}
	FwpEInvalidParameter                                         = &Error{0x80320035, "FWP_E_INVALID_PARAMETER", "The parameter is incorrect."}
	FwpETooManySublayers                                         = &Error{0x80320036, "FWP_E_TOO_MANY_SUBLAYERS", "The maximum number of sublayers has been reached."}
	FwpECalloutNotificationFailed                                = &Error{0x80320037, "FWP_E_CALLOUT_NOTIFICATION_FAILED", "The notification function for a callout returned an error."}
	FwpEIncompatibleAuthConfig                                   = &Error{0x80320038, "FWP_E_INCOMPATIBLE_AUTH_CONFIG", "The IPsec authentication configuration is not compatible with the authentication type."}
	FwpEIncompatibleCipherConfig                                 = &Error{0x80320039, "FWP_E_INCOMPATIBLE_CIPHER_CONFIG", "The IPsec cipher configuration is not compatible with the cipher type."}
	ErrorNdisInterfaceClosing                                    = &Error{0x80340002, "ERROR_NDIS_INTERFACE_CLOSING", "The binding to the network interface is being closed."}
	ErrorNdisBadVersion                                          = &Error{0x80340004, "ERROR_NDIS_BAD_VERSION", "An invalid version was specified."}
	ErrorNdisBadCharacteristics                                  = &Error{0x80340005, "ERROR_NDIS_BAD_CHARACTERISTICS", "An invalid characteristics table was used."}
	ErrorNdisAdapterNotFound                                     = &Error{0x80340006, "ERROR_NDIS_ADAPTER_NOT_FOUND", "Failed to find the network interface, or the network interface is not ready."}
	ErrorNdisOpenFailed                                          = &Error{0x80340007, "ERROR_NDIS_OPEN_FAILED", "Failed to open the network interface."}
	ErrorNdisDeviceFailed                                        = &Error{0x80340008, "ERROR_NDIS_DEVICE_FAILED", "The network interface has encountered an internal unrecoverable failure."}
	ErrorNdisMulticastFull                                       = &Error{0x80340009, "ERROR_NDIS_MULTICAST_FULL", "The multicast list on the network interface is full."}
	ErrorNdisMulticastExists                                     = &Error{0x8034000A, "ERROR_NDIS_MULTICAST_EXISTS", "An attempt was made to add a duplicate multicast address to the list."}
	ErrorNdisMulticastNotFound                                   = &Error{0x8034000B, "ERROR_NDIS_MULTICAST_NOT_FOUND", "At attempt was made to remove a multicast address that was never added."}
	ErrorNdisRequestAborted                                      = &Error{0x8034000C, "ERROR_NDIS_REQUEST_ABORTED", "The network interface aborted the request."}
	ErrorNdisResetInProgress                                     = &Error{0x8034000D, "ERROR_NDIS_RESET_IN_PROGRESS", "The network interface cannot process the request because it is being reset."}
	ErrorNdisInvalidPacket                                       = &Error{0x8034000F, "ERROR_NDIS_INVALID_PACKET", "An attempt was made to send an invalid packet on a network interface."}
	ErrorNdisInvalidDeviceRequest                                = &Error{0x80340010, "ERROR_NDIS_INVALID_DEVICE_REQUEST", "The specified request is not a valid operation for the target device."}
	ErrorNdisAdapterNotReady                                     = &Error{0x80340011, "ERROR_NDIS_ADAPTER_NOT_READY", "The network interface is not ready to complete this operation."}
	ErrorNdisInvalidLength                                       = &Error{0x80340014, "ERROR_NDIS_INVALID_LENGTH", "The length of the buffer submitted for this operation is not valid."}
	ErrorNdisInvalidData                                         = &Error{0x80340015, "ERROR_NDIS_INVALID_DATA", "The data used for this operation is not valid."}
	ErrorNdisBufferTooShort                                      = &Error{0x80340016, "ERROR_NDIS_BUFFER_TOO_SHORT", "The length of the buffer submitted for this operation is too small."}
	ErrorNdisInvalidOid                                          = &Error{0x80340017, "ERROR_NDIS_INVALID_OID", "The network interface does not support this OID."}
	ErrorNdisAdapterRemoved                                      = &Error{0x80340018, "ERROR_NDIS_ADAPTER_REMOVED", "The network interface has been removed."}
	ErrorNdisUnsupportedMedia                                    = &Error{0x80340019, "ERROR_NDIS_UNSUPPORTED_MEDIA", "The network interface does not support this media type."}
	ErrorNdisGroupAddressInUse                                   = &Error{0x8034001A, "ERROR_NDIS_GROUP_ADDRESS_IN_USE", "An attempt was made to remove a token ring group address that is in use by other components."}
	ErrorNdisFileNotFound                                        = &Error{0x8034001B, "ERROR_NDIS_FILE_NOT_FOUND", "An attempt was made to map a file that cannot be found."}
	ErrorNdisErrorReadingFile                                    = &Error{0x8034001C, "ERROR_NDIS_ERROR_READING_FILE", "An error occurred while the NDIS tried to map the file."}
	ErrorNdisAlreadyMapped                                       = &Error{0x8034001D, "ERROR_NDIS_ALREADY_MAPPED", "An attempt was made to map a file that is already mapped."}
	ErrorNdisResourceConflict                                    = &Error{0x8034001E, "ERROR_NDIS_RESOURCE_CONFLICT", "An attempt to allocate a hardware resource failed because the resource is used by another component."}
	ErrorNdisMediaDisconnected                                   = &Error{0x8034001F, "ERROR_NDIS_MEDIA_DISCONNECTED", "The I/O operation failed because network media is disconnected or the wireless access point is out of range."}
	ErrorNdisInvalidAddress                                      = &Error{0x80340022, "ERROR_NDIS_INVALID_ADDRESS", "The network address used in the request is invalid."}
	ErrorNdisPaused                                              = &Error{0x8034002A, "ERROR_NDIS_PAUSED", "The offload operation on the network interface has been paused."}
	ErrorNdisInterfaceNotFound                                   = &Error{0x8034002B, "ERROR_NDIS_INTERFACE_NOT_FOUND", "The network interface was not found."}
	ErrorNdisUnsupportedRevision                                 = &Error{0x8034002C, "ERROR_NDIS_UNSUPPORTED_REVISION", "The revision number specified in the structure is not supported."}
	ErrorNdisInvalidPort                                         = &Error{0x8034002D, "ERROR_NDIS_INVALID_PORT", "The specified port does not exist on this network interface."}
	ErrorNdisInvalidPortState                                    = &Error{0x8034002E, "ERROR_NDIS_INVALID_PORT_STATE", "The current state of the specified port on this network interface does not support the requested operation."}
	ErrorNdisNotSupported                                        = &Error{0x803400BB, "ERROR_NDIS_NOT_SUPPORTED", "The network interface does not support this request."}
	ErrorNdisDot11AutoConfigEnabled                              = &Error{0x80342000, "ERROR_NDIS_DOT11_AUTO_CONFIG_ENABLED", "The wireless local area network (LAN) interface is in auto-configuration mode and does not support the requested parameter change operation."}
	ErrorNdisDot11MediaInUse                                     = &Error{0x80342001, "ERROR_NDIS_DOT11_MEDIA_IN_USE", "The wireless LAN interface is busy and cannot perform the requested operation."}
	ErrorNdisDot11PowerStateInvalid                              = &Error{0x80342002, "ERROR_NDIS_DOT11_POWER_STATE_INVALID", "The wireless LAN interface is shutting down and does not support the requested operation."}
	TrkENotFound                                                 = &Error{0x8DEAD01B, "TRK_E_NOT_FOUND", "A requested object was not found."}
	TrkEVolumeQuotaExceeded                                      = &Error{0x8DEAD01C, "TRK_E_VOLUME_QUOTA_EXCEEDED", "The server received a CREATE_VOLUME subrequest of a SYNC_VOLUMES request, but the ServerVolumeTable size limit for the RequestMachine has already been reached."}
	TrkServerTooBusy                                             = &Error{0x8DEAD01E, "TRK_SERVER_TOO_BUSY", "The server is busy, and the client should retry the request at a later time."}
	ErrorAuditingDisabled                                        = &Error{0xC0090001, "ERROR_AUDITING_DISABLED", "The specified event is currently not being audited."}
	ErrorAllSidsFiltered                                         = &Error{0xC0090002, "ERROR_ALL_SIDS_FILTERED", "The SID filtering operation removed all SIDs."}
	ErrorBizrulesNotEnabled                                      = &Error{0xC0090003, "ERROR_BIZRULES_NOT_ENABLED", "Business rule scripts are disabled for the calling application."}
	NsENoconnection                                              = &Error{0xC00D0005, "NS_E_NOCONNECTION", "There is no connection established with the Windows Media server. The operation failed."}
	NsECannotconnect                                             = &Error{0xC00D0006, "NS_E_CANNOTCONNECT", "Unable to establish a connection to the server."}
	NsECannotdestroytitle                                        = &Error{0xC00D0007, "NS_E_CANNOTDESTROYTITLE", "Unable to destroy the title."}
	NsECannotrenametitle                                         = &Error{0xC00D0008, "NS_E_CANNOTRENAMETITLE", "Unable to rename the title."}
	NsECannotofflinedisk                                         = &Error{0xC00D0009, "NS_E_CANNOTOFFLINEDISK", "Unable to offline disk."}
	NsECannotonlinedisk                                          = &Error{0xC00D000A, "NS_E_CANNOTONLINEDISK", "Unable to online disk."}
	NsENoregisteredwalker                                        = &Error{0xC00D000B, "NS_E_NOREGISTEREDWALKER", "There is no file parser registered for this type of file."}
	NsENofunnel                                                  = &Error{0xC00D000C, "NS_E_NOFUNNEL", "There is no data connection established."}
	NsENoLocalplay                                               = &Error{0xC00D000D, "NS_E_NO_LOCALPLAY", "Failed to load the local play DLL."}
	NsENetworkBusy                                               = &Error{0xC00D000E, "NS_E_NETWORK_BUSY", "The network is busy."}
	NsETooManySess                                               = &Error{0xC00D000F, "NS_E_TOO_MANY_SESS", "The server session limit was exceeded."}
	NsEAlreadyConnected                                          = &Error{0xC00D0010, "NS_E_ALREADY_CONNECTED", "The network connection already exists."}
	NsEInvalidIndex                                              = &Error{0xC00D0011, "NS_E_INVALID_INDEX", "Index %1 is invalid."}
	NsEProtocolMismatch                                          = &Error{0xC00D0012, "NS_E_PROTOCOL_MISMATCH", "There is no protocol or protocol version supported by both the client and the server."}
	NsETimeout                                                   = &Error{0xC00D0013, "NS_E_TIMEOUT", "The server, a computer set up to offer multimedia content to other computers, could not handle your request for multimedia content in a timely manner. Please try again later."}
	NsENetWrite                                                  = &Error{0xC00D0014, "NS_E_NET_WRITE", "Error writing to the network."}
	NsENetRead                                                   = &Error{0xC00D0015, "NS_E_NET_READ", "Error reading from the network."}
	NsEDiskWrite                                                 = &Error{0xC00D0016, "NS_E_DISK_WRITE", "Error writing to a disk."}
	NsEDiskRead                                                  = &Error{0xC00D0017, "NS_E_DISK_READ", "Error reading from a disk."}
	NsEFileWrite                                                 = &Error{0xC00D0018, "NS_E_FILE_WRITE", "Error writing to a file."}
	NsEFileRead                                                  = &Error{0xC00D0019, "NS_E_FILE_READ", "Error reading from a file."}
	NsEFileNotFound                                              = &Error{0xC00D001A, "NS_E_FILE_NOT_FOUND", "The system cannot find the file specified."}
	NsEFileExists                                                = &Error{0xC00D001B, "NS_E_FILE_EXISTS", "The file already exists."}
	NsEInvalidName                                               = &Error{0xC00D001C, "NS_E_INVALID_NAME", "The file name, directory name, or volume label syntax is incorrect."}
	NsEFileOpenFailed                                            = &Error{0xC00D001D, "NS_E_FILE_OPEN_FAILED", "Failed to open a file."}
	NsEFileAllocationFailed                                      = &Error{0xC00D001E, "NS_E_FILE_ALLOCATION_FAILED", "Unable to allocate a file."}
	NsEFileInitFailed                                            = &Error{0xC00D001F, "NS_E_FILE_INIT_FAILED", "Unable to initialize a file."}
	NsEFilePlayFailed                                            = &Error{0xC00D0020, "NS_E_FILE_PLAY_FAILED", "Unable to play a file."}
	NsESetDiskUidFailed                                          = &Error{0xC00D0021, "NS_E_SET_DISK_UID_FAILED", "Could not set the disk UID."}
	NsEInduced                                                   = &Error{0xC00D0022, "NS_E_INDUCED", "An error was induced for testing purposes."}
	NsECclinkDown                                                = &Error{0xC00D0023, "NS_E_CCLINK_DOWN", "Two Content Servers failed to communicate."}
	NsEInternal                                                  = &Error{0xC00D0024, "NS_E_INTERNAL", "An unknown error occurred."}
	NsEBusy                                                      = &Error{0xC00D0025, "NS_E_BUSY", "The requested resource is in use."}
	NsEUnrecognizedStreamType                                    = &Error{0xC00D0026, "NS_E_UNRECOGNIZED_STREAM_TYPE", "The specified protocol is not recognized. Be sure that the file name and syntax, such as slashes, are correct for the protocol."}
	NsENetworkServiceFailure                                     = &Error{0xC00D0027, "NS_E_NETWORK_SERVICE_FAILURE", "The network service provider failed."}
	NsENetworkResourceFailure                                    = &Error{0xC00D0028, "NS_E_NETWORK_RESOURCE_FAILURE", "An attempt to acquire a network resource failed."}
	NsEConnectionFailure                                         = &Error{0xC00D0029, "NS_E_CONNECTION_FAILURE", "The network connection has failed."}
	NsEShutdown                                                  = &Error{0xC00D002A, "NS_E_SHUTDOWN", "The session is being terminated locally."}
	NsEInvalidRequest                                            = &Error{0xC00D002B, "NS_E_INVALID_REQUEST", "The request is invalid in the current state."}
	NsEInsufficientBandwidth                                     = &Error{0xC00D002C, "NS_E_INSUFFICIENT_BANDWIDTH", "There is insufficient bandwidth available to fulfill the request."}
	NsENotRebuilding                                             = &Error{0xC00D002D, "NS_E_NOT_REBUILDING", "The disk is not rebuilding."}
	NsELateOperation                                             = &Error{0xC00D002E, "NS_E_LATE_OPERATION", "An operation requested for a particular time could not be carried out on schedule."}
	NsEInvalidData                                               = &Error{0xC00D002F, "NS_E_INVALID_DATA", "Invalid or corrupt data was encountered."}
	NsEFileBandwidthLimit                                        = &Error{0xC00D0030, "NS_E_FILE_BANDWIDTH_LIMIT", "The bandwidth required to stream a file is higher than the maximum file bandwidth allowed on the server."}
	NsEOpenFileLimit                                             = &Error{0xC00D0031, "NS_E_OPEN_FILE_LIMIT", "The client cannot have any more files open simultaneously."}
	NsEBadControlData                                            = &Error{0xC00D0032, "NS_E_BAD_CONTROL_DATA", "The server received invalid data from the client on the control connection."}
	NsENoStream                                                  = &Error{0xC00D0033, "NS_E_NO_STREAM", "There is no stream available."}
	NsEStreamEnd                                                 = &Error{0xC00D0034, "NS_E_STREAM_END", "There is no more data in the stream."}
	NsEServerNotFound                                            = &Error{0xC00D0035, "NS_E_SERVER_NOT_FOUND", "The specified server could not be found."}
	NsEDuplicateName                                             = &Error{0xC00D0036, "NS_E_DUPLICATE_NAME", "The specified name is already in use."}
	NsEDuplicateAddress                                          = &Error{0xC00D0037, "NS_E_DUPLICATE_ADDRESS", "The specified address is already in use."}
	NsEBadMulticastAddress                                       = &Error{0xC00D0038, "NS_E_BAD_MULTICAST_ADDRESS", "The specified address is not a valid multicast address."}
	NsEBadAdapterAddress                                         = &Error{0xC00D0039, "NS_E_BAD_ADAPTER_ADDRESS", "The specified adapter address is invalid."}
	NsEBadDeliveryMode                                           = &Error{0xC00D003A, "NS_E_BAD_DELIVERY_MODE", "The specified delivery mode is invalid."}
	NsEInvalidChannel                                            = &Error{0xC00D003B, "NS_E_INVALID_CHANNEL", "The specified station does not exist."}
	NsEInvalidStream                                             = &Error{0xC00D003C, "NS_E_INVALID_STREAM", "The specified stream does not exist."}
	NsEInvalidArchive                                            = &Error{0xC00D003D, "NS_E_INVALID_ARCHIVE", "The specified archive could not be opened."}
	NsENotitles                                                  = &Error{0xC00D003E, "NS_E_NOTITLES", "The system cannot find any titles on the server."}
	NsEInvalidClient                                             = &Error{0xC00D003F, "NS_E_INVALID_CLIENT", "The system cannot find the client specified."}
	NsEInvalidBlackholeAddress                                   = &Error{0xC00D0040, "NS_E_INVALID_BLACKHOLE_ADDRESS", "The Blackhole Address is not initialized."}
	NsEIncompatibleFormat                                        = &Error{0xC00D0041, "NS_E_INCOMPATIBLE_FORMAT", "The station does not support the stream format."}
	NsEInvalidKey                                                = &Error{0xC00D0042, "NS_E_INVALID_KEY", "The specified key is not valid."}
	NsEInvalidPort                                               = &Error{0xC00D0043, "NS_E_INVALID_PORT", "The specified port is not valid."}
	NsEInvalidTtl                                                = &Error{0xC00D0044, "NS_E_INVALID_TTL", "The specified TTL is not valid."}
	NsEStrideRefused                                             = &Error{0xC00D0045, "NS_E_STRIDE_REFUSED", "The request to fast forward or rewind could not be fulfilled."}
	NsEMmsautoserverCantfindwalker                               = &Error{0xC00D0046, "NS_E_MMSAUTOSERVER_CANTFINDWALKER", "Unable to load the appropriate file parser."}
	NsEMaxBitrate                                                = &Error{0xC00D0047, "NS_E_MAX_BITRATE", "Cannot exceed the maximum bandwidth limit."}
	NsELogfileperiod                                             = &Error{0xC00D0048, "NS_E_LOGFILEPERIOD", "Invalid value for LogFilePeriod."}
	NsEMaxClients                                                = &Error{0xC00D0049, "NS_E_MAX_CLIENTS", "Cannot exceed the maximum client limit."}
	NsELogFileSize                                               = &Error{0xC00D004A, "NS_E_LOG_FILE_SIZE", "The maximum log file size has been reached."}
	NsEMaxFilerate                                               = &Error{0xC00D004B, "NS_E_MAX_FILERATE", "Cannot exceed the maximum file rate."}
	NsEWalkerUnknown                                             = &Error{0xC00D004C, "NS_E_WALKER_UNKNOWN", "Unknown file type."}
	NsEWalkerServer                                              = &Error{0xC00D004D, "NS_E_WALKER_SERVER", "The specified file, %1, cannot be loaded onto the specified server, %2."}
	NsEWalkerUsage                                               = &Error{0xC00D004E, "NS_E_WALKER_USAGE", "There was a usage error with file parser."}
	NsETigerFail                                                 = &Error{0xC00D0050, "NS_E_TIGER_FAIL", "The Title Server %1 has failed."}
	NsECubFail                                                   = &Error{0xC00D0053, "NS_E_CUB_FAIL", "Content Server %1 (%2) has failed."}
	NsEDiskFail                                                  = &Error{0xC00D0055, "NS_E_DISK_FAIL", "Disk %1 ( %2 ) on Content Server %3, has failed."}
	NsEMaxFunnelsAlert                                           = &Error{0xC00D0060, "NS_E_MAX_FUNNELS_ALERT", "The NetShow data stream limit of %1 streams was reached."}
	NsEAllocateFileFail                                          = &Error{0xC00D0061, "NS_E_ALLOCATE_FILE_FAIL", "The NetShow Video Server was unable to allocate a %1 block file named %2."}
	NsEPagingError                                               = &Error{0xC00D0062, "NS_E_PAGING_ERROR", "A Content Server was unable to page a block."}
	NsEBadBlock0Version                                          = &Error{0xC00D0063, "NS_E_BAD_BLOCK0_VERSION", "Disk %1 has unrecognized control block version %2."}
	NsEBadDiskUid                                                = &Error{0xC00D0064, "NS_E_BAD_DISK_UID", "Disk %1 has incorrect uid %2."}
	NsEBadFsmajorVersion                                         = &Error{0xC00D0065, "NS_E_BAD_FSMAJOR_VERSION", "Disk %1 has unsupported file system major version %2."}
	NsEBadStampnumber                                            = &Error{0xC00D0066, "NS_E_BAD_STAMPNUMBER", "Disk %1 has bad stamp number in control block."}
	NsEPartiallyRebuiltDisk                                      = &Error{0xC00D0067, "NS_E_PARTIALLY_REBUILT_DISK", "Disk %1 is partially reconstructed."}
	NsEEnactplanGiveup                                           = &Error{0xC00D0068, "NS_E_ENACTPLAN_GIVEUP", "EnactPlan gives up."}
	McmadmERegkeyNotFound                                        = &Error{0xC00D006A, "MCMADM_E_REGKEY_NOT_FOUND", "The key was not found in the registry."}
	NsENoFormats                                                 = &Error{0xC00D006B, "NS_E_NO_FORMATS", "The publishing point cannot be started because the server does not have the appropriate stream formats. Use the Multicast Announcement Wizard to create a new announcement for this publishing point."}
	NsENoReferences                                              = &Error{0xC00D006C, "NS_E_NO_REFERENCES", "No reference URLs were found in an ASX file."}
	NsEWaveOpen                                                  = &Error{0xC00D006D, "NS_E_WAVE_OPEN", "Error opening wave device, the device might be in use."}
	NsECannotconnectevents                                       = &Error{0xC00D006F, "NS_E_CANNOTCONNECTEVENTS", "Unable to establish a connection to the NetShow event monitor service."}
	NsENoDevice                                                  = &Error{0xC00D0071, "NS_E_NO_DEVICE", "No device driver is present on the system."}
	NsENoSpecifiedDevice                                         = &Error{0xC00D0072, "NS_E_NO_SPECIFIED_DEVICE", "No specified device driver is present."}
	NsEMonitorGiveup                                             = &Error{0xC00D00C8, "NS_E_MONITOR_GIVEUP", "Netshow Events Monitor is not operational and has been disconnected."}
	NsERemirroredDisk                                            = &Error{0xC00D00C9, "NS_E_REMIRRORED_DISK", "Disk %1 is remirrored."}
	NsEInsufficientData                                          = &Error{0xC00D00CA, "NS_E_INSUFFICIENT_DATA", "Insufficient data found."}
	NsEAssert                                                    = &Error{0xC00D00CB, "NS_E_ASSERT", "1 failed in file %2 line %3."}
	NsEBadAdapterName                                            = &Error{0xC00D00CC, "NS_E_BAD_ADAPTER_NAME", "The specified adapter name is invalid."}
	NsENotLicensed                                               = &Error{0xC00D00CD, "NS_E_NOT_LICENSED", "The application is not licensed for this feature."}
	NsENoServerContact                                           = &Error{0xC00D00CE, "NS_E_NO_SERVER_CONTACT", "Unable to contact the server."}
	NsETooManyTitles                                             = &Error{0xC00D00CF, "NS_E_TOO_MANY_TITLES", "Maximum number of titles exceeded."}
	NsETitleSizeExceeded                                         = &Error{0xC00D00D0, "NS_E_TITLE_SIZE_EXCEEDED", "Maximum size of a title exceeded."}
	NsEUdpDisabled                                               = &Error{0xC00D00D1, "NS_E_UDP_DISABLED", "UDP protocol not enabled. Not trying %1!ls!."}
	NsETcpDisabled                                               = &Error{0xC00D00D2, "NS_E_TCP_DISABLED", "TCP protocol not enabled. Not trying %1!ls!."}
	NsEHttpDisabled                                              = &Error{0xC00D00D3, "NS_E_HTTP_DISABLED", "HTTP protocol not enabled. Not trying %1!ls!."}
	NsELicenseExpired                                            = &Error{0xC00D00D4, "NS_E_LICENSE_EXPIRED", "The product license has expired."}
	NsETitleBitrate                                              = &Error{0xC00D00D5, "NS_E_TITLE_BITRATE", "Source file exceeds the per title maximum bitrate. See NetShow Theater documentation for more information."}
	NsEEmptyProgramName                                          = &Error{0xC00D00D6, "NS_E_EMPTY_PROGRAM_NAME", "The program name cannot be empty."}
	NsEMissingChannel                                            = &Error{0xC00D00D7, "NS_E_MISSING_CHANNEL", "Station %1 does not exist."}
	NsENoChannels                                                = &Error{0xC00D00D8, "NS_E_NO_CHANNELS", "You need to define at least one station before this operation can complete."}
	NsEInvalidIndex2                                             = &Error{0xC00D00D9, "NS_E_INVALID_INDEX2", "The index specified is invalid."}
	NsECubFailLink                                               = &Error{0xC00D0190, "NS_E_CUB_FAIL_LINK", "Content Server %1 (%2) has failed its link to Content Server %3."}
	NsEBadCubUid                                                 = &Error{0xC00D0192, "NS_E_BAD_CUB_UID", "Content Server %1 (%2) has incorrect uid %3."}
	NsEGlitchMode                                                = &Error{0xC00D0195, "NS_E_GLITCH_MODE", "Server unreliable because multiple components failed."}
	NsENoMediaProtocol                                           = &Error{0xC00D019B, "NS_E_NO_MEDIA_PROTOCOL", "Content Server %1 (%2) is unable to communicate with the Media System Network Protocol."}
	NsENothingToDo                                               = &Error{0xC00D07F1, "NS_E_NOTHING_TO_DO", "Nothing to do."}
	NsENoMulticast                                               = &Error{0xC00D07F2, "NS_E_NO_MULTICAST", "Not receiving data from the server."}
	NsEInvalidInputFormat                                        = &Error{0xC00D0BB8, "NS_E_INVALID_INPUT_FORMAT", "The input media format is invalid."}
	NsEMsaudioNotInstalled                                       = &Error{0xC00D0BB9, "NS_E_MSAUDIO_NOT_INSTALLED", "The MSAudio codec is not installed on this system."}
	NsEUnexpectedMsaudioError                                    = &Error{0xC00D0BBA, "NS_E_UNEXPECTED_MSAUDIO_ERROR", "An unexpected error occurred with the MSAudio codec."}
	NsEInvalidOutputFormat                                       = &Error{0xC00D0BBB, "NS_E_INVALID_OUTPUT_FORMAT", "The output media format is invalid."}
	NsENotConfigured                                             = &Error{0xC00D0BBC, "NS_E_NOT_CONFIGURED", "The object must be fully configured before audio samples can be processed."}
	NsEProtectedContent                                          = &Error{0xC00D0BBD, "NS_E_PROTECTED_CONTENT", "You need a license to perform the requested operation on this media file."}
	NsELicenseRequired                                           = &Error{0xC00D0BBE, "NS_E_LICENSE_REQUIRED", "You need a license to perform the requested operation on this media file."}
	NsETamperedContent                                           = &Error{0xC00D0BBF, "NS_E_TAMPERED_CONTENT", "This media file is corrupted or invalid. Contact the content provider for a new file."}
	NsELicenseOutofdate                                          = &Error{0xC00D0BC0, "NS_E_LICENSE_OUTOFDATE", "The license for this media file has expired. Get a new license or contact the content provider for further assistance."}
	NsELicenseIncorrectRights                                    = &Error{0xC00D0BC1, "NS_E_LICENSE_INCORRECT_RIGHTS", "You are not allowed to open this file. Contact the content provider for further assistance."}
	NsEAudioCodecNotInstalled                                    = &Error{0xC00D0BC2, "NS_E_AUDIO_CODEC_NOT_INSTALLED", "The requested audio codec is not installed on this system."}
	NsEAudioCodecError                                           = &Error{0xC00D0BC3, "NS_E_AUDIO_CODEC_ERROR", "An unexpected error occurred with the audio codec."}
	NsEVideoCodecNotInstalled                                    = &Error{0xC00D0BC4, "NS_E_VIDEO_CODEC_NOT_INSTALLED", "The requested video codec is not installed on this system."}
	NsEVideoCodecError                                           = &Error{0xC00D0BC5, "NS_E_VIDEO_CODEC_ERROR", "An unexpected error occurred with the video codec."}
	NsEInvalidprofile                                            = &Error{0xC00D0BC6, "NS_E_INVALIDPROFILE", "The Profile is invalid."}
	NsEIncompatibleVersion                                       = &Error{0xC00D0BC7, "NS_E_INCOMPATIBLE_VERSION", "A new version of the SDK is needed to play the requested content."}
	NsEOfflineMode                                               = &Error{0xC00D0BCA, "NS_E_OFFLINE_MODE", "The requested URL is not available in offline mode."}
	NsENotConnected                                              = &Error{0xC00D0BCB, "NS_E_NOT_CONNECTED", "The requested URL cannot be accessed because there is no network connection."}
	NsETooMuchData                                               = &Error{0xC00D0BCC, "NS_E_TOO_MUCH_DATA", "The encoding process was unable to keep up with the amount of supplied data."}
	NsEUnsupportedProperty                                       = &Error{0xC00D0BCD, "NS_E_UNSUPPORTED_PROPERTY", "The given property is not supported."}
	NsE8bitWaveUnsupported                                       = &Error{0xC00D0BCE, "NS_E_8BIT_WAVE_UNSUPPORTED", "Windows Media Player cannot copy the files to the CD because they are 8-bit. Convert the files to 16-bit, 44-kHz stereo files by using Sound Recorder or another audio-processing program, and then try again."}
	NsENoMoreSamples                                             = &Error{0xC00D0BCF, "NS_E_NO_MORE_SAMPLES", "There are no more samples in the current range."}
	NsEInvalidSamplingRate                                       = &Error{0xC00D0BD0, "NS_E_INVALID_SAMPLING_RATE", "The given sampling rate is invalid."}
	NsEMaxPacketSizeTooSmall                                     = &Error{0xC00D0BD1, "NS_E_MAX_PACKET_SIZE_TOO_SMALL", "The given maximum packet size is too small to accommodate this profile.)"}
	NsELatePacket                                                = &Error{0xC00D0BD2, "NS_E_LATE_PACKET", "The packet arrived too late to be of use."}
	NsEDuplicatePacket                                           = &Error{0xC00D0BD3, "NS_E_DUPLICATE_PACKET", "The packet is a duplicate of one received before."}
	NsESdkBuffertoosmall                                         = &Error{0xC00D0BD4, "NS_E_SDK_BUFFERTOOSMALL", "Supplied buffer is too small."}
	NsEInvalidNumPasses                                          = &Error{0xC00D0BD5, "NS_E_INVALID_NUM_PASSES", "The wrong number of preprocessing passes was used for the stream's output type."}
	NsEAttributeReadOnly                                         = &Error{0xC00D0BD6, "NS_E_ATTRIBUTE_READ_ONLY", "An attempt was made to add, modify, or delete a read only attribute."}
	NsEAttributeNotAllowed                                       = &Error{0xC00D0BD7, "NS_E_ATTRIBUTE_NOT_ALLOWED", "An attempt was made to add attribute that is not allowed for the given media type."}
	NsEInvalidEdl                                                = &Error{0xC00D0BD8, "NS_E_INVALID_EDL", "The EDL provided is invalid."}
	NsEDataUnitExtensionTooLarge                                 = &Error{0xC00D0BD9, "NS_E_DATA_UNIT_EXTENSION_TOO_LARGE", "The Data Unit Extension data was too large to be used."}
	NsECodecDmoError                                             = &Error{0xC00D0BDA, "NS_E_CODEC_DMO_ERROR", "An unexpected error occurred with a DMO codec."}
	NsEFeatureDisabledByGroupPolicy                              = &Error{0xC00D0BDC, "NS_E_FEATURE_DISABLED_BY_GROUP_POLICY", "This feature has been disabled by group policy."}
	NsEFeatureDisabledInSku                                      = &Error{0xC00D0BDD, "NS_E_FEATURE_DISABLED_IN_SKU", "This feature is disabled in this SKU."}
	NsENoCd                                                      = &Error{0xC00D0FA0, "NS_E_NO_CD", "There is no CD in the CD drive. Insert a CD, and then try again."}
	NsECantReadDigital                                           = &Error{0xC00D0FA1, "NS_E_CANT_READ_DIGITAL", "Windows Media Player could not use digital playback to play the CD. To switch to analog playback, on the Tools menu, click Options, and then click the Devices tab. Double-click the CD drive, and then in the Playback area, click Analog. For additional assistance, click Web Help."}
	NsEDeviceDisconnected                                        = &Error{0xC00D0FA2, "NS_E_DEVICE_DISCONNECTED", "Windows Media Player no longer detects a connected portable device. Reconnect your portable device, and then try synchronizing the file again."}
	NsEDeviceNotSupportFormat                                    = &Error{0xC00D0FA3, "NS_E_DEVICE_NOT_SUPPORT_FORMAT", "Windows Media Player cannot play the file. The portable device does not support the specified file type."}
	NsESlowReadDigital                                           = &Error{0xC00D0FA4, "NS_E_SLOW_READ_DIGITAL", "Windows Media Player could not use digital playback to play the CD. The Player has automatically switched the CD drive to analog playback. To switch back to digital CD playback, use the Devices tab. For additional assistance, click Web Help."}
	NsEMixerInvalidLine                                          = &Error{0xC00D0FA5, "NS_E_MIXER_INVALID_LINE", "An invalid line error occurred in the mixer."}
	NsEMixerInvalidControl                                       = &Error{0xC00D0FA6, "NS_E_MIXER_INVALID_CONTROL", "An invalid control error occurred in the mixer."}
	NsEMixerInvalidValue                                         = &Error{0xC00D0FA7, "NS_E_MIXER_INVALID_VALUE", "An invalid value error occurred in the mixer."}
	NsEMixerUnknownMmresult                                      = &Error{0xC00D0FA8, "NS_E_MIXER_UNKNOWN_MMRESULT", "An unrecognized MMRESULT occurred in the mixer."}
	NsEUserStop                                                  = &Error{0xC00D0FA9, "NS_E_USER_STOP", "User has stopped the operation."}
	NsEMp3FormatNotFound                                         = &Error{0xC00D0FAA, "NS_E_MP3_FORMAT_NOT_FOUND", "Windows Media Player cannot rip the track because a compatible MP3 encoder is not installed on your computer. Install a compatible MP3 encoder or choose a different format to rip to (such as Windows Media Audio)."}
	NsECdReadErrorNoCorrection                                   = &Error{0xC00D0FAB, "NS_E_CD_READ_ERROR_NO_CORRECTION", "Windows Media Player cannot read the CD. The disc might be dirty or damaged. Turn on error correction, and then try again."}
	NsECdReadError                                               = &Error{0xC00D0FAC, "NS_E_CD_READ_ERROR", "Windows Media Player cannot read the CD. The disc might be dirty or damaged or the CD drive might be malfunctioning."}
	NsECdSlowCopy                                                = &Error{0xC00D0FAD, "NS_E_CD_SLOW_COPY", "For best performance, do not play CD tracks while ripping them."}
	NsECdCopytoCd                                                = &Error{0xC00D0FAE, "NS_E_CD_COPYTO_CD", "It is not possible to directly burn tracks from one CD to another CD. You must first rip the tracks from the CD to your computer, and then burn the files to a blank CD."}
	NsEMixerNodriver                                             = &Error{0xC00D0FAF, "NS_E_MIXER_NODRIVER", "Could not open a sound mixer driver."}
	NsERedbookEnabledWhileCopying                                = &Error{0xC00D0FB0, "NS_E_REDBOOK_ENABLED_WHILE_COPYING", "Windows Media Player cannot rip tracks from the CD correctly because the CD drive settings in Device Manager do not match the CD drive settings in the Player."}
	NsECdRefresh                                                 = &Error{0xC00D0FB1, "NS_E_CD_REFRESH", "Windows Media Player is busy reading the CD."}
	NsECdDriverProblem                                           = &Error{0xC00D0FB2, "NS_E_CD_DRIVER_PROBLEM", "Windows Media Player could not use digital playback to play the CD. The Player has automatically switched the CD drive to analog playback. To switch back to digital CD playback, use the Devices tab. For additional assistance, click Web Help."}
	NsEWontDoDigital                                             = &Error{0xC00D0FB3, "NS_E_WONT_DO_DIGITAL", "Windows Media Player could not use digital playback to play the CD. The Player has automatically switched the CD drive to analog playback. To switch back to digital CD playback, use the Devices tab. For additional assistance, click Web Help."}
	NsEWmpxmlNoerror                                             = &Error{0xC00D0FB4, "NS_E_WMPXML_NOERROR", "A call was made to GetParseError on the XML parser but there was no error to retrieve."}
	NsEWmpxmlEndofdata                                           = &Error{0xC00D0FB5, "NS_E_WMPXML_ENDOFDATA", "The XML Parser ran out of data while parsing."}
	NsEWmpxmlParseerror                                          = &Error{0xC00D0FB6, "NS_E_WMPXML_PARSEERROR", "A generic parse error occurred in the XML parser but no information is available."}
	NsEWmpxmlAttributenotfound                                   = &Error{0xC00D0FB7, "NS_E_WMPXML_ATTRIBUTENOTFOUND", "A call get GetNamedAttribute or GetNamedAttributeIndex on the XML parser resulted in the index not being found."}
	NsEWmpxmlPinotfound                                          = &Error{0xC00D0FB8, "NS_E_WMPXML_PINOTFOUND", "A call was made go GetNamedPI on the XML parser, but the requested Processing Instruction was not found."}
	NsEWmpxmlEmptydoc                                            = &Error{0xC00D0FB9, "NS_E_WMPXML_EMPTYDOC", "Persist was called on the XML parser, but the parser has no data to persist."}
	NsEWmpPathAlreadyInLibrary                                   = &Error{0xC00D0FBA, "NS_E_WMP_PATH_ALREADY_IN_LIBRARY", "This file path is already in the library."}
	NsEWmpFilescanalreadystarted                                 = &Error{0xC00D0FBE, "NS_E_WMP_FILESCANALREADYSTARTED", "Windows Media Player is already searching for files to add to your library. Wait for the current process to finish before attempting to search again."}
	NsEWmpHmeInvalidobjectid                                     = &Error{0xC00D0FBF, "NS_E_WMP_HME_INVALIDOBJECTID", "Windows Media Player is unable to find the media you are looking for."}
	NsEWmpMfCodeExpired                                          = &Error{0xC00D0FC0, "NS_E_WMP_MF_CODE_EXPIRED", "A component of Windows Media Player is out-of-date. If you are running a pre-release version of Windows, try upgrading to a more recent version."}
	NsEWmpHmeNotsearchableforitems                               = &Error{0xC00D0FC1, "NS_E_WMP_HME_NOTSEARCHABLEFORITEMS", "This container does not support search on items."}
	NsEWmpAddtolibraryFailed                                     = &Error{0xC00D0FC7, "NS_E_WMP_ADDTOLIBRARY_FAILED", "Windows Media Player encountered a problem while adding one or more files to the library. For additional assistance, click Web Help."}
	NsEWmpWindowsapifailure                                      = &Error{0xC00D0FC8, "NS_E_WMP_WINDOWSAPIFAILURE", "A Windows API call failed but no error information was available."}
	NsEWmpRecordingNotAllowed                                    = &Error{0xC00D0FC9, "NS_E_WMP_RECORDING_NOT_ALLOWED", "This file does not have burn rights. If you obtained this file from an online store, go to the online store to get burn rights."}
	NsEDeviceNotReady                                            = &Error{0xC00D0FCA, "NS_E_DEVICE_NOT_READY", "Windows Media Player no longer detects a connected portable device. Reconnect your portable device, and then try to sync the file again."}
	NsEDamagedFile                                               = &Error{0xC00D0FCB, "NS_E_DAMAGED_FILE", "Windows Media Player cannot play the file because it is corrupted."}
	NsEMpdbGeneric                                               = &Error{0xC00D0FCC, "NS_E_MPDB_GENERIC", "Windows Media Player encountered an error while attempting to access information in the library. Try restarting the Player."}
	NsEFileFailedChecks                                          = &Error{0xC00D0FCD, "NS_E_FILE_FAILED_CHECKS", "The file cannot be added to the library because it is smaller than the \"Skip files smaller than\" setting. To add the file, change the setting on the Library tab. For additional assistance, click Web Help."}
	NsEMediaLibraryFailed                                        = &Error{0xC00D0FCE, "NS_E_MEDIA_LIBRARY_FAILED", "Windows Media Player cannot create the library. You must be logged on as an administrator or a member of the Administrators group to install the Player. For more information, contact your system administrator."}
	NsESharingViolation                                          = &Error{0xC00D0FCF, "NS_E_SHARING_VIOLATION", "The file is already in use. Close other programs that might be using the file, or stop playing the file, and then try again."}
	NsENoErrorStringFound                                        = &Error{0xC00D0FD0, "NS_E_NO_ERROR_STRING_FOUND", "Windows Media Player has encountered an unknown error."}
	NsEWmpocxNoRemoteCore                                        = &Error{0xC00D0FD1, "NS_E_WMPOCX_NO_REMOTE_CORE", "The Windows Media Player ActiveX control cannot connect to remote media services, but will continue with local media services."}
	NsEWmpocxNoActiveCore                                        = &Error{0xC00D0FD2, "NS_E_WMPOCX_NO_ACTIVE_CORE", "The requested method or property is not available because the Windows Media Player ActiveX control has not been properly activated."}
	NsEWmpocxNotRunningRemotely                                  = &Error{0xC00D0FD3, "NS_E_WMPOCX_NOT_RUNNING_REMOTELY", "The Windows Media Player ActiveX control is not running in remote mode."}
	NsEWmpocxNoRemoteWindow                                      = &Error{0xC00D0FD4, "NS_E_WMPOCX_NO_REMOTE_WINDOW", "An error occurred while trying to get the remote Windows Media Player window."}
	NsEWmpocxErrormanagernotavailable                            = &Error{0xC00D0FD5, "NS_E_WMPOCX_ERRORMANAGERNOTAVAILABLE", "Windows Media Player has encountered an unknown error."}
	NsEPluginNotshutdown                                         = &Error{0xC00D0FD6, "NS_E_PLUGIN_NOTSHUTDOWN", "Windows Media Player was not closed properly. A damaged or incompatible plug-in might have caused the problem to occur. As a precaution, all optional plug-ins have been disabled."}
	NsEWmpCannotFindFolder                                       = &Error{0xC00D0FD7, "NS_E_WMP_CANNOT_FIND_FOLDER", "Windows Media Player cannot find the specified path. Verify that the path is typed correctly. If it is, the path does not exist in the specified location, or the computer where the path is located is not available."}
	NsEWmpStreamingRecordingNotAllowed                           = &Error{0xC00D0FD8, "NS_E_WMP_STREAMING_RECORDING_NOT_ALLOWED", "Windows Media Player cannot save a file that is being streamed."}
	NsEWmpPlugindllNotfound                                      = &Error{0xC00D0FD9, "NS_E_WMP_PLUGINDLL_NOTFOUND", "Windows Media Player cannot find the selected plug-in. The Player will try to remove it from the menu. To use this plug-in, install it again."}
	NsENeedToAskUser                                             = &Error{0xC00D0FDA, "NS_E_NEED_TO_ASK_USER", "Action requires input from the user."}
	NsEWmpocxPlayerNotDocked                                     = &Error{0xC00D0FDB, "NS_E_WMPOCX_PLAYER_NOT_DOCKED", "The Windows Media Player ActiveX control must be in a docked state for this action to be performed."}
	NsEWmpExternalNotready                                       = &Error{0xC00D0FDC, "NS_E_WMP_EXTERNAL_NOTREADY", "The Windows Media Player external object is not ready."}
	NsEWmpMlsStaleData                                           = &Error{0xC00D0FDD, "NS_E_WMP_MLS_STALE_DATA", "Windows Media Player cannot perform the requested action. Your computer's time and date might not be set correctly."}
	NsEWmpUiSubcontrolsnotsupported                              = &Error{0xC00D0FDE, "NS_E_WMP_UI_SUBCONTROLSNOTSUPPORTED", "The control (%s) does not support creation of sub-controls, yet (%d) sub-controls have been specified."}
	NsEWmpUiVersionmismatch                                      = &Error{0xC00D0FDF, "NS_E_WMP_UI_VERSIONMISMATCH", "Version mismatch: (%.1f required, %.1f found)."}
	NsEWmpUiNotathemefile                                        = &Error{0xC00D0FE0, "NS_E_WMP_UI_NOTATHEMEFILE", "The layout manager was given valid XML that wasn't a theme file."}
	NsEWmpUiSubelementnotfound                                   = &Error{0xC00D0FE1, "NS_E_WMP_UI_SUBELEMENTNOTFOUND", "The %s subelement could not be found on the %s object."}
	NsEWmpUiVersionparse                                         = &Error{0xC00D0FE2, "NS_E_WMP_UI_VERSIONPARSE", "An error occurred parsing the version tag. Valid version tags are of the form: <?wmp version='1.0'?>."}
	NsEWmpUiViewidnotfound                                       = &Error{0xC00D0FE3, "NS_E_WMP_UI_VIEWIDNOTFOUND", "The view specified in for the 'currentViewID' property (%s) was not found in this theme file."}
	NsEWmpUiPassthrough                                          = &Error{0xC00D0FE4, "NS_E_WMP_UI_PASSTHROUGH", "This error used internally for hit testing."}
	NsEWmpUiObjectnotfound                                       = &Error{0xC00D0FE5, "NS_E_WMP_UI_OBJECTNOTFOUND", "Attributes were specified for the %s object, but the object was not available to send them to."}
	NsEWmpUiSecondhandler                                        = &Error{0xC00D0FE6, "NS_E_WMP_UI_SECONDHANDLER", "The %s event already has a handler, the second handler was ignored."}
	NsEWmpUiNoskininzip                                          = &Error{0xC00D0FE7, "NS_E_WMP_UI_NOSKININZIP", "No .wms file found in skin archive."}
	NsEWmpUrldownloadfailed                                      = &Error{0xC00D0FEA, "NS_E_WMP_URLDOWNLOADFAILED", "Windows Media Player encountered a problem while downloading the file. For additional assistance, click Web Help."}
	NsEWmpocxUnableToLoadSkin                                    = &Error{0xC00D0FEB, "NS_E_WMPOCX_UNABLE_TO_LOAD_SKIN", "The Windows Media Player ActiveX control cannot load the requested uiMode and cannot roll back to the existing uiMode."}
	NsEWmpInvalidSkin                                            = &Error{0xC00D0FEC, "NS_E_WMP_INVALID_SKIN", "Windows Media Player encountered a problem with the skin file. The skin file might not be valid."}
	NsEWmpSendmailfailed                                         = &Error{0xC00D0FED, "NS_E_WMP_SENDMAILFAILED", "Windows Media Player cannot send the link because your email program is not responding. Verify that your email program is configured properly, and then try again. For more information about email, see Windows Help."}
	NsEWmpLockedinskinmode                                       = &Error{0xC00D0FEE, "NS_E_WMP_LOCKEDINSKINMODE", "Windows Media Player cannot switch to full mode because your computer administrator has locked this skin."}
	NsEWmpFailedToSaveFile                                       = &Error{0xC00D0FEF, "NS_E_WMP_FAILED_TO_SAVE_FILE", "Windows Media Player encountered a problem while saving the file. For additional assistance, click Web Help."}
	NsEWmpSaveasReadonly                                         = &Error{0xC00D0FF0, "NS_E_WMP_SAVEAS_READONLY", "Windows Media Player cannot overwrite a read-only file. Try using a different file name."}
	NsEWmpFailedToSavePlaylist                                   = &Error{0xC00D0FF1, "NS_E_WMP_FAILED_TO_SAVE_PLAYLIST", "Windows Media Player encountered a problem while creating or saving the playlist. For additional assistance, click Web Help."}
	NsEWmpFailedToOpenWmd                                        = &Error{0xC00D0FF2, "NS_E_WMP_FAILED_TO_OPEN_WMD", "Windows Media Player cannot open the Windows Media Download file. The file might be damaged."}
	NsEWmpCantPlayProtected                                      = &Error{0xC00D0FF3, "NS_E_WMP_CANT_PLAY_PROTECTED", "The file cannot be added to the library because it is a protected DVR-MS file. This content cannot be played back by Windows Media Player."}
	NsESharingStateOutOfSync                                     = &Error{0xC00D0FF4, "NS_E_SHARING_STATE_OUT_OF_SYNC", "Media sharing has been turned off because a required Windows setting or component has changed. For additional assistance, click Web Help."}
	NsEWmpocxRemotePlayerAlreadyRunning                          = &Error{0xC00D0FFA, "NS_E_WMPOCX_REMOTE_PLAYER_ALREADY_RUNNING", "Exclusive Services launch failed because the Windows Media Player is already running."}
	NsEWmpRbcJpgmappingimage                                     = &Error{0xC00D1004, "NS_E_WMP_RBC_JPGMAPPINGIMAGE", "JPG Images are not recommended for use as a mappingImage."}
	NsEWmpJpgtransparency                                        = &Error{0xC00D1005, "NS_E_WMP_JPGTRANSPARENCY", "JPG Images are not recommended when using a transparencyColor."}
	NsEWmpInvalidMaxVal                                          = &Error{0xC00D1009, "NS_E_WMP_INVALID_MAX_VAL", "The Max property cannot be less than Min property."}
	NsEWmpInvalidMinVal                                          = &Error{0xC00D100A, "NS_E_WMP_INVALID_MIN_VAL", "The Min property cannot be greater than Max property."}
	NsEWmpCsJpgpositionimage                                     = &Error{0xC00D100E, "NS_E_WMP_CS_JPGPOSITIONIMAGE", "JPG Images are not recommended for use as a positionImage."}
	NsEWmpCsNotevenlydivisible                                   = &Error{0xC00D100F, "NS_E_WMP_CS_NOTEVENLYDIVISIBLE", "The (%s) image's size is not evenly divisible by the positionImage's size."}
	NsEWmpzipNotazipfile                                         = &Error{0xC00D1018, "NS_E_WMPZIP_NOTAZIPFILE", "The ZIP reader opened a file and its signature did not match that of the ZIP files."}
	NsEWmpzipCorrupt                                             = &Error{0xC00D1019, "NS_E_WMPZIP_CORRUPT", "The ZIP reader has detected that the file is corrupted."}
	NsEWmpzipFilenotfound                                        = &Error{0xC00D101A, "NS_E_WMPZIP_FILENOTFOUND", "GetFileStream, SaveToFile, or SaveTemp file was called on the ZIP reader with a file name that was not found in the ZIP file."}
	NsEWmpImageFiletypeUnsupported                               = &Error{0xC00D1022, "NS_E_WMP_IMAGE_FILETYPE_UNSUPPORTED", "Image type not supported."}
	NsEWmpImageInvalidFormat                                     = &Error{0xC00D1023, "NS_E_WMP_IMAGE_INVALID_FORMAT", "Image file might be corrupt."}
	NsEWmpGifUnexpectedEndoffile                                 = &Error{0xC00D1024, "NS_E_WMP_GIF_UNEXPECTED_ENDOFFILE", "Unexpected end of file. GIF file might be corrupt."}
	NsEWmpGifInvalidFormat                                       = &Error{0xC00D1025, "NS_E_WMP_GIF_INVALID_FORMAT", "Invalid GIF file."}
	NsEWmpGifBadVersionNumber                                    = &Error{0xC00D1026, "NS_E_WMP_GIF_BAD_VERSION_NUMBER", "Invalid GIF version. Only 87a or 89a supported."}
	NsEWmpGifNoImageInFile                                       = &Error{0xC00D1027, "NS_E_WMP_GIF_NO_IMAGE_IN_FILE", "No images found in GIF file."}
	NsEWmpPngInvalidformat                                       = &Error{0xC00D1028, "NS_E_WMP_PNG_INVALIDFORMAT", "Invalid PNG image file format."}
	NsEWmpPngUnsupportedBitdepth                                 = &Error{0xC00D1029, "NS_E_WMP_PNG_UNSUPPORTED_BITDEPTH", "PNG bitdepth not supported."}
	NsEWmpPngUnsupportedCompression                              = &Error{0xC00D102A, "NS_E_WMP_PNG_UNSUPPORTED_COMPRESSION", "Compression format defined in PNG file not supported,"}
	NsEWmpPngUnsupportedFilter                                   = &Error{0xC00D102B, "NS_E_WMP_PNG_UNSUPPORTED_FILTER", "Filter method defined in PNG file not supported."}
	NsEWmpPngUnsupportedInterlace                                = &Error{0xC00D102C, "NS_E_WMP_PNG_UNSUPPORTED_INTERLACE", "Interlace method defined in PNG file not supported."}
	NsEWmpPngUnsupportedBadCrc                                   = &Error{0xC00D102D, "NS_E_WMP_PNG_UNSUPPORTED_BAD_CRC", "Bad CRC in PNG file."}
	NsEWmpBmpInvalidBitmask                                      = &Error{0xC00D102E, "NS_E_WMP_BMP_INVALID_BITMASK", "Invalid bitmask in BMP file."}
	NsEWmpBmpTopdownDibUnsupported                               = &Error{0xC00D102F, "NS_E_WMP_BMP_TOPDOWN_DIB_UNSUPPORTED", "Topdown DIB not supported."}
	NsEWmpBmpBitmapNotCreated                                    = &Error{0xC00D1030, "NS_E_WMP_BMP_BITMAP_NOT_CREATED", "Bitmap could not be created."}
	NsEWmpBmpCompressionUnsupported                              = &Error{0xC00D1031, "NS_E_WMP_BMP_COMPRESSION_UNSUPPORTED", "Compression format defined in BMP not supported."}
	NsEWmpBmpInvalidFormat                                       = &Error{0xC00D1032, "NS_E_WMP_BMP_INVALID_FORMAT", "Invalid Bitmap format."}
	NsEWmpJpgJerrArithcodingNotimpl                              = &Error{0xC00D1033, "NS_E_WMP_JPG_JERR_ARITHCODING_NOTIMPL", "JPEG Arithmetic coding not supported."}
	NsEWmpJpgInvalidFormat                                       = &Error{0xC00D1034, "NS_E_WMP_JPG_INVALID_FORMAT", "Invalid JPEG format."}
	NsEWmpJpgBadDctsize                                          = &Error{0xC00D1035, "NS_E_WMP_JPG_BAD_DCTSIZE", "Invalid JPEG format."}
	NsEWmpJpgBadVersionNumber                                    = &Error{0xC00D1036, "NS_E_WMP_JPG_BAD_VERSION_NUMBER", "Internal version error. Unexpected JPEG library version."}
	NsEWmpJpgBadPrecision                                        = &Error{0xC00D1037, "NS_E_WMP_JPG_BAD_PRECISION", "Internal JPEG Library error. Unsupported JPEG data precision."}
	NsEWmpJpgCcir601Notimpl                                      = &Error{0xC00D1038, "NS_E_WMP_JPG_CCIR601_NOTIMPL", "JPEG CCIR601 not supported."}
	NsEWmpJpgNoImageInFile                                       = &Error{0xC00D1039, "NS_E_WMP_JPG_NO_IMAGE_IN_FILE", "No image found in JPEG file."}
	NsEWmpJpgReadError                                           = &Error{0xC00D103A, "NS_E_WMP_JPG_READ_ERROR", "Could not read JPEG file."}
	NsEWmpJpgFractSampleNotimpl                                  = &Error{0xC00D103B, "NS_E_WMP_JPG_FRACT_SAMPLE_NOTIMPL", "JPEG Fractional sampling not supported."}
	NsEWmpJpgImageTooBig                                         = &Error{0xC00D103C, "NS_E_WMP_JPG_IMAGE_TOO_BIG", "JPEG image too large. Maximum image size supported is 65500 X 65500."}
	NsEWmpJpgUnexpectedEndoffile                                 = &Error{0xC00D103D, "NS_E_WMP_JPG_UNEXPECTED_ENDOFFILE", "Unexpected end of file reached in JPEG file."}
	NsEWmpJpgSofUnsupported                                      = &Error{0xC00D103E, "NS_E_WMP_JPG_SOF_UNSUPPORTED", "Unsupported JPEG SOF marker found."}
	NsEWmpJpgUnknownMarker                                       = &Error{0xC00D103F, "NS_E_WMP_JPG_UNKNOWN_MARKER", "Unknown JPEG marker found."}
	NsEWmpFailedToOpenImage                                      = &Error{0xC00D1044, "NS_E_WMP_FAILED_TO_OPEN_IMAGE", "Windows Media Player cannot display the picture file. The player either does not support the picture type or the picture is corrupted."}
	NsEWmpDaiSongtooshort                                        = &Error{0xC00D1049, "NS_E_WMP_DAI_SONGTOOSHORT", "Windows Media Player cannot compute a Digital Audio Id for the song. It is too short."}
	NsEWmgRateunavailable                                        = &Error{0xC00D104A, "NS_E_WMG_RATEUNAVAILABLE", "Windows Media Player cannot play the file at the requested speed."}
	NsEWmgPluginunavailable                                      = &Error{0xC00D104B, "NS_E_WMG_PLUGINUNAVAILABLE", "The rendering or digital signal processing plug-in cannot be instantiated."}
	NsEWmgCannotqueue                                            = &Error{0xC00D104C, "NS_E_WMG_CANNOTQUEUE", "The file cannot be queued for seamless playback."}
	NsEWmgPrerolllicenseacquisitionnotallowed                    = &Error{0xC00D104D, "NS_E_WMG_PREROLLLICENSEACQUISITIONNOTALLOWED", "Windows Media Player cannot download media usage rights for a file in the playlist."}
	NsEWmgUnexpectedprerollstatus                                = &Error{0xC00D104E, "NS_E_WMG_UNEXPECTEDPREROLLSTATUS", "Windows Media Player encountered an error while trying to queue a file."}
	NsEWmgInvalidCoppCertificate                                 = &Error{0xC00D1051, "NS_E_WMG_INVALID_COPP_CERTIFICATE", "Windows Media Player cannot play the protected file. The Player cannot verify that the connection to your video card is secure. Try installing an updated device driver for your video card."}
	NsEWmgCoppSecurityInvalid                                    = &Error{0xC00D1052, "NS_E_WMG_COPP_SECURITY_INVALID", "Windows Media Player cannot play the protected file. The Player detected that the connection to your hardware might not be secure."}
	NsEWmgCoppUnsupported                                        = &Error{0xC00D1053, "NS_E_WMG_COPP_UNSUPPORTED", "Windows Media Player output link protection is unsupported on this system."}
	NsEWmgInvalidstate                                           = &Error{0xC00D1054, "NS_E_WMG_INVALIDSTATE", "Operation attempted in an invalid graph state."}
	NsEWmgSinkalreadyexists                                      = &Error{0xC00D1055, "NS_E_WMG_SINKALREADYEXISTS", "A renderer cannot be inserted in a stream while one already exists."}
	NsEWmgNosdkinterface                                         = &Error{0xC00D1056, "NS_E_WMG_NOSDKINTERFACE", "The Windows Media SDK interface needed to complete the operation does not exist at this time."}
	NsEWmgNotalloutputsrendered                                  = &Error{0xC00D1057, "NS_E_WMG_NOTALLOUTPUTSRENDERED", "Windows Media Player cannot play a portion of the file because it requires a codec that either could not be downloaded or that is not supported by the Player."}
	NsEWmgFiletransfernotallowed                                 = &Error{0xC00D1058, "NS_E_WMG_FILETRANSFERNOTALLOWED", "File transfer streams are not allowed in the standalone Player."}
	NsEWmrUnsupportedstream                                      = &Error{0xC00D1059, "NS_E_WMR_UNSUPPORTEDSTREAM", "Windows Media Player cannot play the file. The Player does not support the format you are trying to play."}
	NsEWmrPinnotfound                                            = &Error{0xC00D105A, "NS_E_WMR_PINNOTFOUND", "An operation was attempted on a pin that does not exist in the DirectShow filter graph."}
	NsEWmrWaitingonformatswitch                                  = &Error{0xC00D105B, "NS_E_WMR_WAITINGONFORMATSWITCH", "Specified operation cannot be completed while waiting for a media format change from the SDK."}
	NsEWmrNosourcefilter                                         = &Error{0xC00D105C, "NS_E_WMR_NOSOURCEFILTER", "Specified operation cannot be completed because the source filter does not exist."}
	NsEWmrPintypenomatch                                         = &Error{0xC00D105D, "NS_E_WMR_PINTYPENOMATCH", "The specified type does not match this pin."}
	NsEWmrNocallbackavailable                                    = &Error{0xC00D105E, "NS_E_WMR_NOCALLBACKAVAILABLE", "The WMR Source Filter does not have a callback available."}
	NsEWmrSamplepropertynotset                                   = &Error{0xC00D1062, "NS_E_WMR_SAMPLEPROPERTYNOTSET", "The specified property has not been set on this sample."}
	NsEWmrCannotRenderBinaryStream                               = &Error{0xC00D1063, "NS_E_WMR_CANNOT_RENDER_BINARY_STREAM", "A plug-in is required to correctly play the file. To determine if the plug-in is available to download, click Web Help."}
	NsEWmgLicenseTampered                                        = &Error{0xC00D1064, "NS_E_WMG_LICENSE_TAMPERED", "Windows Media Player cannot play the file because your media usage rights are corrupted. If you previously backed up your media usage rights, try restoring them."}
	NsEWmrWillnotRenderBinaryStream                              = &Error{0xC00D1065, "NS_E_WMR_WILLNOT_RENDER_BINARY_STREAM", "Windows Media Player cannot play protected files that contain binary streams."}
	NsEWmxUnrecognizedPlaylistFormat                             = &Error{0xC00D1068, "NS_E_WMX_UNRECOGNIZED_PLAYLIST_FORMAT", "Windows Media Player cannot play the playlist because it is not valid."}
	NsEAsxInvalidformat                                          = &Error{0xC00D1069, "NS_E_ASX_INVALIDFORMAT", "Windows Media Player cannot play the playlist because it is not valid."}
	NsEAsxInvalidversion                                         = &Error{0xC00D106A, "NS_E_ASX_INVALIDVERSION", "A later version of Windows Media Player might be required to play this playlist."}
	NsEAsxInvalidRepeatBlock                                     = &Error{0xC00D106B, "NS_E_ASX_INVALID_REPEAT_BLOCK", "The format of a REPEAT loop within the current playlist file is not valid."}
	NsEAsxNothingToWrite                                         = &Error{0xC00D106C, "NS_E_ASX_NOTHING_TO_WRITE", "Windows Media Player cannot save the playlist because it does not contain any items."}
	NsEUrllistInvalidformat                                      = &Error{0xC00D106D, "NS_E_URLLIST_INVALIDFORMAT", "Windows Media Player cannot play the playlist because it is not valid."}
	NsEWmxAttributeDoesNotExist                                  = &Error{0xC00D106E, "NS_E_WMX_ATTRIBUTE_DOES_NOT_EXIST", "The specified attribute does not exist."}
	NsEWmxAttributeAlreadyExists                                 = &Error{0xC00D106F, "NS_E_WMX_ATTRIBUTE_ALREADY_EXISTS", "The specified attribute already exists."}
	NsEWmxAttributeUnretrievable                                 = &Error{0xC00D1070, "NS_E_WMX_ATTRIBUTE_UNRETRIEVABLE", "Cannot retrieve the specified attribute."}
	NsEWmxItemDoesNotExist                                       = &Error{0xC00D1071, "NS_E_WMX_ITEM_DOES_NOT_EXIST", "The specified item does not exist in the current playlist."}
	NsEWmxItemTypeIllegal                                        = &Error{0xC00D1072, "NS_E_WMX_ITEM_TYPE_ILLEGAL", "Items of the specified type cannot be created within the current playlist."}
	NsEWmxItemUnsettable                                         = &Error{0xC00D1073, "NS_E_WMX_ITEM_UNSETTABLE", "The specified item cannot be set in the current playlist."}
	NsEWmxPlaylistEmpty                                          = &Error{0xC00D1074, "NS_E_WMX_PLAYLIST_EMPTY", "Windows Media Player cannot perform the requested action because the playlist does not contain any items."}
	NsEMlsSmartplaylistFilterNotRegistered                       = &Error{0xC00D1075, "NS_E_MLS_SMARTPLAYLIST_FILTER_NOT_REGISTERED", "The specified auto playlist contains a filter type that is either not valid or is not installed on this computer."}
	NsEWmxInvalidFormatOverNesting                               = &Error{0xC00D1076, "NS_E_WMX_INVALID_FORMAT_OVER_NESTING", "Windows Media Player cannot play the file because the associated playlist contains too many nested playlists."}
	NsEWmpcoreNosourceurlstring                                  = &Error{0xC00D107C, "NS_E_WMPCORE_NOSOURCEURLSTRING", "Windows Media Player cannot find the file. Verify that the path is typed correctly. If it is, the file might not exist in the specified location, or the computer where the file is stored might not be available."}
	NsEWmpcoreCocreatefailedforgitobject                         = &Error{0xC00D107D, "NS_E_WMPCORE_COCREATEFAILEDFORGITOBJECT", "Failed to create the Global Interface Table."}
	NsEWmpcoreFailedtogetmarshalledeventhandlerinterface         = &Error{0xC00D107E, "NS_E_WMPCORE_FAILEDTOGETMARSHALLEDEVENTHANDLERINTERFACE", "Failed to get the marshaled graph event handler interface."}
	NsEWmpcoreBuffertoosmall                                     = &Error{0xC00D107F, "NS_E_WMPCORE_BUFFERTOOSMALL", "Buffer is too small for copying media type."}
	NsEWmpcoreUnavailable                                        = &Error{0xC00D1080, "NS_E_WMPCORE_UNAVAILABLE", "The current state of the Player does not allow this operation."}
	NsEWmpcoreInvalidplaylistmode                                = &Error{0xC00D1081, "NS_E_WMPCORE_INVALIDPLAYLISTMODE", "The playlist manager does not understand the current play mode (for example, shuffle or normal)."}
	NsEWmpcoreItemnotinplaylist                                  = &Error{0xC00D1086, "NS_E_WMPCORE_ITEMNOTINPLAYLIST", "Windows Media Player cannot play the file because it is not in the current playlist."}
	NsEWmpcorePlaylistempty                                      = &Error{0xC00D1087, "NS_E_WMPCORE_PLAYLISTEMPTY", "There are no items in the playlist. Add items to the playlist, and then try again."}
	NsEWmpcoreNobrowser                                          = &Error{0xC00D1088, "NS_E_WMPCORE_NOBROWSER", "The web page cannot be displayed because no web browser is installed on your computer."}
	NsEWmpcoreUnrecognizedMediaUrl                               = &Error{0xC00D1089, "NS_E_WMPCORE_UNRECOGNIZED_MEDIA_URL", "Windows Media Player cannot find the specified file. Verify the path is typed correctly. If it is, the file does not exist in the specified location, or the computer where the file is stored is not available."}
	NsEWmpcoreGraphNotInList                                     = &Error{0xC00D108A, "NS_E_WMPCORE_GRAPH_NOT_IN_LIST", "Graph with the specified URL was not found in the prerolled graph list."}
	NsEWmpcorePlaylistEmptyOrSingleMedia                         = &Error{0xC00D108B, "NS_E_WMPCORE_PLAYLIST_EMPTY_OR_SINGLE_MEDIA", "Windows Media Player cannot perform the requested operation because there is only one item in the playlist."}
	NsEWmpcoreErrorsinknotregistered                             = &Error{0xC00D108C, "NS_E_WMPCORE_ERRORSINKNOTREGISTERED", "An error sink was never registered for the calling object."}
	NsEWmpcoreErrormanagernotavailable                           = &Error{0xC00D108D, "NS_E_WMPCORE_ERRORMANAGERNOTAVAILABLE", "The error manager is not available to respond to errors."}
	NsEWmpcoreWebhelpfailed                                      = &Error{0xC00D108E, "NS_E_WMPCORE_WEBHELPFAILED", "The Web Help URL cannot be opened."}
	NsEWmpcoreMediaErrorResumeFailed                             = &Error{0xC00D108F, "NS_E_WMPCORE_MEDIA_ERROR_RESUME_FAILED", "Could not resume playing next item in playlist."}
	NsEWmpcoreNoRefInEntry                                       = &Error{0xC00D1090, "NS_E_WMPCORE_NO_REF_IN_ENTRY", "Windows Media Player cannot play the file because the associated playlist does not contain any items or the playlist is not valid."}
	NsEWmpcoreWmxListAttributeNameEmpty                          = &Error{0xC00D1091, "NS_E_WMPCORE_WMX_LIST_ATTRIBUTE_NAME_EMPTY", "An empty string for playlist attribute name was found."}
	NsEWmpcoreWmxListAttributeNameIllegal                        = &Error{0xC00D1092, "NS_E_WMPCORE_WMX_LIST_ATTRIBUTE_NAME_ILLEGAL", "A playlist attribute name that is not valid was found."}
	NsEWmpcoreWmxListAttributeValueEmpty                         = &Error{0xC00D1093, "NS_E_WMPCORE_WMX_LIST_ATTRIBUTE_VALUE_EMPTY", "An empty string for a playlist attribute value was found."}
	NsEWmpcoreWmxListAttributeValueIllegal                       = &Error{0xC00D1094, "NS_E_WMPCORE_WMX_LIST_ATTRIBUTE_VALUE_ILLEGAL", "An illegal value for a playlist attribute was found."}
	NsEWmpcoreWmxListItemAttributeNameEmpty                      = &Error{0xC00D1095, "NS_E_WMPCORE_WMX_LIST_ITEM_ATTRIBUTE_NAME_EMPTY", "An empty string for a playlist item attribute name was found."}
	NsEWmpcoreWmxListItemAttributeNameIllegal                    = &Error{0xC00D1096, "NS_E_WMPCORE_WMX_LIST_ITEM_ATTRIBUTE_NAME_ILLEGAL", "An illegal value for a playlist item attribute name was found."}
	NsEWmpcoreWmxListItemAttributeValueEmpty                     = &Error{0xC00D1097, "NS_E_WMPCORE_WMX_LIST_ITEM_ATTRIBUTE_VALUE_EMPTY", "An illegal value for a playlist item attribute was found."}
	NsEWmpcoreListEntryNoRef                                     = &Error{0xC00D1098, "NS_E_WMPCORE_LIST_ENTRY_NO_REF", "The playlist does not contain any items."}
	NsEWmpcoreMisnamedFile                                       = &Error{0xC00D1099, "NS_E_WMPCORE_MISNAMED_FILE", "Windows Media Player cannot play the file. The file is either corrupted or the Player does not support the format you are trying to play."}
	NsEWmpcoreCodecNotTrusted                                    = &Error{0xC00D109A, "NS_E_WMPCORE_CODEC_NOT_TRUSTED", "The codec downloaded for this file does not appear to be properly signed, so it cannot be installed."}
	NsEWmpcoreCodecNotFound                                      = &Error{0xC00D109B, "NS_E_WMPCORE_CODEC_NOT_FOUND", "Windows Media Player cannot play the file. One or more codecs required to play the file could not be found."}
	NsEWmpcoreCodecDownloadNotAllowed                            = &Error{0xC00D109C, "NS_E_WMPCORE_CODEC_DOWNLOAD_NOT_ALLOWED", "Windows Media Player cannot play the file because a required codec is not installed on your computer. To try downloading the codec, turn on the \"Download codecs automatically\" option."}
	NsEWmpcoreErrorDownloadingPlaylist                           = &Error{0xC00D109D, "NS_E_WMPCORE_ERROR_DOWNLOADING_PLAYLIST", "Windows Media Player encountered a problem while downloading the playlist. For additional assistance, click Web Help."}
	NsEWmpcoreFailedToBuildPlaylist                              = &Error{0xC00D109E, "NS_E_WMPCORE_FAILED_TO_BUILD_PLAYLIST", "Failed to build the playlist."}
	NsEWmpcorePlaylistItemAlternateNone                          = &Error{0xC00D109F, "NS_E_WMPCORE_PLAYLIST_ITEM_ALTERNATE_NONE", "Playlist has no alternates to switch into."}
	NsEWmpcorePlaylistItemAlternateExhausted                     = &Error{0xC00D10A0, "NS_E_WMPCORE_PLAYLIST_ITEM_ALTERNATE_EXHAUSTED", "No more playlist alternates available to switch to."}
	NsEWmpcorePlaylistItemAlternateNameNotFound                  = &Error{0xC00D10A1, "NS_E_WMPCORE_PLAYLIST_ITEM_ALTERNATE_NAME_NOT_FOUND", "Could not find the name of the alternate playlist to switch into."}
	NsEWmpcorePlaylistItemAlternateMorphFailed                   = &Error{0xC00D10A2, "NS_E_WMPCORE_PLAYLIST_ITEM_ALTERNATE_MORPH_FAILED", "Failed to switch to an alternate for this media."}
	NsEWmpcorePlaylistItemAlternateInitFailed                    = &Error{0xC00D10A3, "NS_E_WMPCORE_PLAYLIST_ITEM_ALTERNATE_INIT_FAILED", "Failed to initialize an alternate for the media."}
	NsEWmpcoreMediaAlternateRefEmpty                             = &Error{0xC00D10A4, "NS_E_WMPCORE_MEDIA_ALTERNATE_REF_EMPTY", "No URL specified for the roll over Refs in the playlist file."}
	NsEWmpcorePlaylistNoEventName                                = &Error{0xC00D10A5, "NS_E_WMPCORE_PLAYLIST_NO_EVENT_NAME", "Encountered a playlist with no name."}
	NsEWmpcorePlaylistEventAttributeAbsent                       = &Error{0xC00D10A6, "NS_E_WMPCORE_PLAYLIST_EVENT_ATTRIBUTE_ABSENT", "A required attribute in the event block of the playlist was not found."}
	NsEWmpcorePlaylistEventEmpty                                 = &Error{0xC00D10A7, "NS_E_WMPCORE_PLAYLIST_EVENT_EMPTY", "No items were found in the event block of the playlist."}
	NsEWmpcorePlaylistStackEmpty                                 = &Error{0xC00D10A8, "NS_E_WMPCORE_PLAYLIST_STACK_EMPTY", "No playlist was found while returning from a nested playlist."}
	NsEWmpcoreCurrentMediaNotActive                              = &Error{0xC00D10A9, "NS_E_WMPCORE_CURRENT_MEDIA_NOT_ACTIVE", "The media item is not active currently."}
	NsEWmpcoreUserCancel                                         = &Error{0xC00D10AB, "NS_E_WMPCORE_USER_CANCEL", "Windows Media Player cannot perform the requested action because you chose to cancel it."}
	NsEWmpcorePlaylistRepeatEmpty                                = &Error{0xC00D10AC, "NS_E_WMPCORE_PLAYLIST_REPEAT_EMPTY", "Windows Media Player encountered a problem with the playlist. The format of the playlist is not valid."}
	NsEWmpcorePlaylistRepeatStartMediaNone                       = &Error{0xC00D10AD, "NS_E_WMPCORE_PLAYLIST_REPEAT_START_MEDIA_NONE", "Media object corresponding to start of a playlist repeat block was not found."}
	NsEWmpcorePlaylistRepeatEndMediaNone                         = &Error{0xC00D10AE, "NS_E_WMPCORE_PLAYLIST_REPEAT_END_MEDIA_NONE", "Media object corresponding to the end of a playlist repeat block was not found."}
	NsEWmpcoreInvalidPlaylistUrl                                 = &Error{0xC00D10AF, "NS_E_WMPCORE_INVALID_PLAYLIST_URL", "The playlist URL supplied to the playlist manager is not valid."}
	NsEWmpcoreMismatchedRuntime                                  = &Error{0xC00D10B0, "NS_E_WMPCORE_MISMATCHED_RUNTIME", "Windows Media Player cannot play the file because it is corrupted."}
	NsEWmpcorePlaylistImportFailedNoItems                        = &Error{0xC00D10B1, "NS_E_WMPCORE_PLAYLIST_IMPORT_FAILED_NO_ITEMS", "Windows Media Player cannot add the playlist to the library because the playlist does not contain any items."}
	NsEWmpcoreVideoTransformFilterInsertion                      = &Error{0xC00D10B2, "NS_E_WMPCORE_VIDEO_TRANSFORM_FILTER_INSERTION", "An error has occurred that could prevent the changing of the video contrast on this media."}
	NsEWmpcoreMediaUnavailable                                   = &Error{0xC00D10B3, "NS_E_WMPCORE_MEDIA_UNAVAILABLE", "Windows Media Player cannot play the file. If the file is located on the Internet, connect to the Internet. If the file is located on a removable storage card, insert the storage card."}
	NsEWmpcoreWmxEntryrefNoRef                                   = &Error{0xC00D10B4, "NS_E_WMPCORE_WMX_ENTRYREF_NO_REF", "The playlist contains an ENTRYREF for which no href was parsed. Check the syntax of playlist file."}
	NsEWmpcoreNoPlayableMediaInPlaylist                          = &Error{0xC00D10B5, "NS_E_WMPCORE_NO_PLAYABLE_MEDIA_IN_PLAYLIST", "Windows Media Player cannot play any items in the playlist. To find information about the problem, click the Now Playing tab, and then click the icon next to each file in the List pane."}
	NsEWmpcorePlaylistEmptyNestedPlaylistSkippedItems            = &Error{0xC00D10B6, "NS_E_WMPCORE_PLAYLIST_EMPTY_NESTED_PLAYLIST_SKIPPED_ITEMS", "Windows Media Player cannot play some or all of the items in the playlist because the playlist is nested."}
	NsEWmpcoreBusy                                               = &Error{0xC00D10B7, "NS_E_WMPCORE_BUSY", "Windows Media Player cannot play the file at this time. Try again later."}
	NsEWmpcoreMediaChildPlaylistUnavailable                      = &Error{0xC00D10B8, "NS_E_WMPCORE_MEDIA_CHILD_PLAYLIST_UNAVAILABLE", "There is no child playlist available for this media item at this time."}
	NsEWmpcoreMediaNoChildPlaylist                               = &Error{0xC00D10B9, "NS_E_WMPCORE_MEDIA_NO_CHILD_PLAYLIST", "There is no child playlist for this media item."}
	NsEWmpcoreFileNotFound                                       = &Error{0xC00D10BA, "NS_E_WMPCORE_FILE_NOT_FOUND", "Windows Media Player cannot find the file. The link from the item in the library to its associated digital media file might be broken. To fix the problem, try repairing the link or removing the item from the library."}
	NsEWmpcoreTempFileNotFound                                   = &Error{0xC00D10BB, "NS_E_WMPCORE_TEMP_FILE_NOT_FOUND", "The temporary file was not found."}
	NsEWmdmRevoked                                               = &Error{0xC00D10BC, "NS_E_WMDM_REVOKED", "Windows Media Player cannot sync the file because the device needs to be updated."}
	NsEDdrawGeneric                                              = &Error{0xC00D10BD, "NS_E_DDRAW_GENERIC", "Windows Media Player cannot play the video because there is a problem with your video card."}
	NsEDisplayModeChangeFailed                                   = &Error{0xC00D10BE, "NS_E_DISPLAY_MODE_CHANGE_FAILED", "Windows Media Player failed to change the screen mode for full-screen video playback."}
	NsEPlaylistContainsErrors                                    = &Error{0xC00D10BF, "NS_E_PLAYLIST_CONTAINS_ERRORS", "Windows Media Player cannot play one or more files. For additional information, right-click an item that cannot be played, and then click Error Details."}
	NsEChangingProxyName                                         = &Error{0xC00D10C0, "NS_E_CHANGING_PROXY_NAME", "Cannot change the proxy name if the proxy setting is not set to custom."}
	NsEChangingProxyPort                                         = &Error{0xC00D10C1, "NS_E_CHANGING_PROXY_PORT", "Cannot change the proxy port if the proxy setting is not set to custom."}
	NsEChangingProxyExceptionlist                                = &Error{0xC00D10C2, "NS_E_CHANGING_PROXY_EXCEPTIONLIST", "Cannot change the proxy exception list if the proxy setting is not set to custom."}
	NsEChangingProxybypass                                       = &Error{0xC00D10C3, "NS_E_CHANGING_PROXYBYPASS", "Cannot change the proxy bypass flag if the proxy setting is not set to custom."}
	NsEChangingProxyProtocolNotFound                             = &Error{0xC00D10C4, "NS_E_CHANGING_PROXY_PROTOCOL_NOT_FOUND", "Cannot find the specified protocol."}
	NsEGraphNoaudiolanguage                                      = &Error{0xC00D10C5, "NS_E_GRAPH_NOAUDIOLANGUAGE", "Cannot change the language settings. Either the graph has no audio or the audio only supports one language."}
	NsEGraphNoaudiolanguageselected                              = &Error{0xC00D10C6, "NS_E_GRAPH_NOAUDIOLANGUAGESELECTED", "The graph has no audio language selected."}
	NsECorecdNotamediacd                                         = &Error{0xC00D10C7, "NS_E_CORECD_NOTAMEDIACD", "This is not a media CD."}
	NsEWmpcoreMediaUrlTooLong                                    = &Error{0xC00D10C8, "NS_E_WMPCORE_MEDIA_URL_TOO_LONG", "Windows Media Player cannot play the file because the URL is too long."}
	NsEWmpflashCantFindComServer                                 = &Error{0xC00D10C9, "NS_E_WMPFLASH_CANT_FIND_COM_SERVER", "To play the selected item, you must install the Macromedia Flash Player. To download the Macromedia Flash Player, go to the Adobe website."}
	NsEWmpflashIncompatibleversion                               = &Error{0xC00D10CA, "NS_E_WMPFLASH_INCOMPATIBLEVERSION", "To play the selected item, you must install a later version of the Macromedia Flash Player. To download the Macromedia Flash Player, go to the Adobe website."}
	NsEWmpocxgraphIeDisallowsActivexControls                     = &Error{0xC00D10CB, "NS_E_WMPOCXGRAPH_IE_DISALLOWS_ACTIVEX_CONTROLS", "Windows Media Player cannot play the file because your Internet security settings prohibit the use of ActiveX controls."}
	NsENeedCoreReference                                         = &Error{0xC00D10CC, "NS_E_NEED_CORE_REFERENCE", "The use of this method requires an existing reference to the Player object."}
	NsEMediacdReadError                                          = &Error{0xC00D10CD, "NS_E_MEDIACD_READ_ERROR", "Windows Media Player cannot play the CD. The disc might be dirty or damaged."}
	NsEIeDisallowsActivexControls                                = &Error{0xC00D10CE, "NS_E_IE_DISALLOWS_ACTIVEX_CONTROLS", "Windows Media Player cannot play the file because your Internet security settings prohibit the use of ActiveX controls."}
	NsEFlashPlaybackNotAllowed                                   = &Error{0xC00D10CF, "NS_E_FLASH_PLAYBACK_NOT_ALLOWED", "Flash playback has been turned off in Windows Media Player."}
	NsEUnableToCreateRipLocation                                 = &Error{0xC00D10D0, "NS_E_UNABLE_TO_CREATE_RIP_LOCATION", "Windows Media Player cannot rip the CD because a valid rip location cannot be created."}
	NsEWmpcoreSomeCodecsMissing                                  = &Error{0xC00D10D1, "NS_E_WMPCORE_SOME_CODECS_MISSING", "Windows Media Player cannot play the file because a required codec is not installed on your computer."}
	NsEWmpRipFailed                                              = &Error{0xC00D10D2, "NS_E_WMP_RIP_FAILED", "Windows Media Player cannot rip one or more tracks from the CD."}
	NsEWmpFailedToRipTrack                                       = &Error{0xC00D10D3, "NS_E_WMP_FAILED_TO_RIP_TRACK", "Windows Media Player encountered a problem while ripping the track from the CD. For additional assistance, click Web Help."}
	NsEWmpEraseFailed                                            = &Error{0xC00D10D4, "NS_E_WMP_ERASE_FAILED", "Windows Media Player encountered a problem while erasing the disc. For additional assistance, click Web Help."}
	NsEWmpFormatFailed                                           = &Error{0xC00D10D5, "NS_E_WMP_FORMAT_FAILED", "Windows Media Player encountered a problem while formatting the device. For additional assistance, click Web Help."}
	NsEWmpCannotBurnNonLocalFile                                 = &Error{0xC00D10D6, "NS_E_WMP_CANNOT_BURN_NON_LOCAL_FILE", "This file cannot be burned to a CD because it is not located on your computer."}
	NsEWmpFileTypeCannotBurnToAudioCd                            = &Error{0xC00D10D7, "NS_E_WMP_FILE_TYPE_CANNOT_BURN_TO_AUDIO_CD", "It is not possible to burn this file type to an audio CD. Windows Media Player can burn the following file types to an audio CD: WMA, MP3, or WAV."}
	NsEWmpFileDoesNotFitOnCd                                     = &Error{0xC00D10D8, "NS_E_WMP_FILE_DOES_NOT_FIT_ON_CD", "This file is too large to fit on a disc."}
	NsEWmpFileNoDuration                                         = &Error{0xC00D10D9, "NS_E_WMP_FILE_NO_DURATION", "It is not possible to determine if this file can fit on a disc because Windows Media Player cannot detect the length of the file. Playing the file before burning might enable the Player to detect the file length."}
	NsEPdaFailedToBurn                                           = &Error{0xC00D10DA, "NS_E_PDA_FAILED_TO_BURN", "Windows Media Player encountered a problem while burning the file to the disc. For additional assistance, click Web Help."}
	NsEFailedDownloadAbortBurn                                   = &Error{0xC00D10DC, "NS_E_FAILED_DOWNLOAD_ABORT_BURN", "Windows Media Player cannot burn the audio CD because some items in the list that you chose to buy could not be downloaded from the online store."}
	NsEWmpcoreDeviceDriversMissing                               = &Error{0xC00D10DD, "NS_E_WMPCORE_DEVICE_DRIVERS_MISSING", "Windows Media Player cannot play the file. Try using Windows Update or Device Manager to update the device drivers for your audio and video cards. For information about using Windows Update or Device Manager, see Windows Help."}
	NsEWmpimUseroffline                                          = &Error{0xC00D1126, "NS_E_WMPIM_USEROFFLINE", "Windows Media Player has detected that you are not connected to the Internet. Connect to the Internet, and then try again."}
	NsEWmpimUsercanceled                                         = &Error{0xC00D1127, "NS_E_WMPIM_USERCANCELED", "The attempt to connect to the Internet was canceled."}
	NsEWmpimDialupfailed                                         = &Error{0xC00D1128, "NS_E_WMPIM_DIALUPFAILED", "The attempt to connect to the Internet failed."}
	NsEWinsockErrorString                                        = &Error{0xC00D1129, "NS_E_WINSOCK_ERROR_STRING", "Windows Media Player has encountered an unknown network error."}
	NsEWmpbrNolistener                                           = &Error{0xC00D1130, "NS_E_WMPBR_NOLISTENER", "No window is currently listening to Backup and Restore events."}
	NsEWmpbrBackupcancel                                         = &Error{0xC00D1131, "NS_E_WMPBR_BACKUPCANCEL", "Your media usage rights were not backed up because the backup was canceled."}
	NsEWmpbrRestorecancel                                        = &Error{0xC00D1132, "NS_E_WMPBR_RESTORECANCEL", "Your media usage rights were not restored because the restoration was canceled."}
	NsEWmpbrErrorwithurl                                         = &Error{0xC00D1133, "NS_E_WMPBR_ERRORWITHURL", "An error occurred while backing up or restoring your media usage rights. A required web page cannot be displayed."}
	NsEWmpbrNamecollision                                        = &Error{0xC00D1134, "NS_E_WMPBR_NAMECOLLISION", "Your media usage rights were not backed up because the backup was canceled."}
	NsEWmpbrDriveInvalid                                         = &Error{0xC00D1137, "NS_E_WMPBR_DRIVE_INVALID", "Windows Media Player cannot restore your media usage rights from the specified location. Choose another location, and then try again."}
	NsEWmpbrBackuprestorefailed                                  = &Error{0xC00D1138, "NS_E_WMPBR_BACKUPRESTOREFAILED", "Windows Media Player cannot backup or restore your media usage rights."}
	NsEWmpConvertFileFailed                                      = &Error{0xC00D1158, "NS_E_WMP_CONVERT_FILE_FAILED", "Windows Media Player cannot add the file to the library."}
	NsEWmpConvertNoRightsErrorurl                                = &Error{0xC00D1159, "NS_E_WMP_CONVERT_NO_RIGHTS_ERRORURL", "Windows Media Player cannot add the file to the library because the content provider prohibits it. For assistance, contact the company that provided the file."}
	NsEWmpConvertNoRightsNoerrorurl                              = &Error{0xC00D115A, "NS_E_WMP_CONVERT_NO_RIGHTS_NOERRORURL", "Windows Media Player cannot add the file to the library because the content provider prohibits it. For assistance, contact the company that provided the file."}
	NsEWmpConvertFileCorrupt                                     = &Error{0xC00D115B, "NS_E_WMP_CONVERT_FILE_CORRUPT", "Windows Media Player cannot add the file to the library. The file might not be valid."}
	NsEWmpConvertPluginUnavailableErrorurl                       = &Error{0xC00D115C, "NS_E_WMP_CONVERT_PLUGIN_UNAVAILABLE_ERRORURL", "Windows Media Player cannot add the file to the library. The plug-in required to add the file is not installed properly. For assistance, click Web Help to display the website of the company that provided the file."}
	NsEWmpConvertPluginUnavailableNoerrorurl                     = &Error{0xC00D115D, "NS_E_WMP_CONVERT_PLUGIN_UNAVAILABLE_NOERRORURL", "Windows Media Player cannot add the file to the library. The plug-in required to add the file is not installed properly. For assistance, contact the company that provided the file."}
	NsEWmpConvertPluginUnknownFileOwner                          = &Error{0xC00D115E, "NS_E_WMP_CONVERT_PLUGIN_UNKNOWN_FILE_OWNER", "Windows Media Player cannot add the file to the library. The plug-in required to add the file is not installed properly. For assistance, contact the company that provided the file."}
	NsEDvdDiscCopyProtectOutputNs                                = &Error{0xC00D1160, "NS_E_DVD_DISC_COPY_PROTECT_OUTPUT_NS", "Windows Media Player cannot play this DVD. Try installing an updated driver for your video card or obtaining a newer video card."}
	NsEDvdDiscCopyProtectOutputFailed                            = &Error{0xC00D1161, "NS_E_DVD_DISC_COPY_PROTECT_OUTPUT_FAILED", "This DVD's resolution exceeds the maximum allowed by your component video outputs. Try reducing your screen resolution to 640 x 480, or turn off analog component outputs and use a VGA connection to your monitor."}
	NsEDvdNoSubpictureStream                                     = &Error{0xC00D1162, "NS_E_DVD_NO_SUBPICTURE_STREAM", "Windows Media Player cannot display subtitles or highlights in DVD menus. Reinstall the DVD decoder or contact the DVD drive manufacturer to obtain an updated decoder."}
	NsEDvdCopyProtect                                            = &Error{0xC00D1163, "NS_E_DVD_COPY_PROTECT", "Windows Media Player cannot play this DVD because there is a problem with digital copy protection between your DVD drive, decoder, and video card. Try installing an updated driver for your video card."}
	NsEDvdAuthoringProblem                                       = &Error{0xC00D1164, "NS_E_DVD_AUTHORING_PROBLEM", "Windows Media Player cannot play the DVD. The disc was created in a manner that the Player does not support."}
	NsEDvdInvalidDiscRegion                                      = &Error{0xC00D1165, "NS_E_DVD_INVALID_DISC_REGION", "Windows Media Player cannot play the DVD because the disc prohibits playback in your region of the world. You must obtain a disc that is intended for your geographic region."}
	NsEDvdCompatibleVideoCard                                    = &Error{0xC00D1166, "NS_E_DVD_COMPATIBLE_VIDEO_CARD", "Windows Media Player cannot play the DVD because your video card does not support DVD playback."}
	NsEDvdMacrovision                                            = &Error{0xC00D1167, "NS_E_DVD_MACROVISION", "Windows Media Player cannot play this DVD because it is not possible to turn on analog copy protection on the output display. Try installing an updated driver for your video card."}
	NsEDvdSystemDecoderRegion                                    = &Error{0xC00D1168, "NS_E_DVD_SYSTEM_DECODER_REGION", "Windows Media Player cannot play the DVD because the region assigned to your DVD drive does not match the region assigned to your DVD decoder."}
	NsEDvdDiscDecoderRegion                                      = &Error{0xC00D1169, "NS_E_DVD_DISC_DECODER_REGION", "Windows Media Player cannot play the DVD because the disc prohibits playback in your region of the world. You must obtain a disc that is intended for your geographic region."}
	NsEDvdNoVideoStream                                          = &Error{0xC00D116A, "NS_E_DVD_NO_VIDEO_STREAM", "Windows Media Player cannot play DVD video. You might need to adjust your Windows display settings. Open display settings in Control Panel, and then try lowering your screen resolution and color quality settings."}
	NsEDvdNoAudioStream                                          = &Error{0xC00D116B, "NS_E_DVD_NO_AUDIO_STREAM", "Windows Media Player cannot play DVD audio. Verify that your sound card is set up correctly, and then try again."}
	NsEDvdGraphBuilding                                          = &Error{0xC00D116C, "NS_E_DVD_GRAPH_BUILDING", "Windows Media Player cannot play DVD video. Close any open files and quit any other programs, and then try again. If the problem persists, restart your computer."}
	NsEDvdNoDecoder                                              = &Error{0xC00D116D, "NS_E_DVD_NO_DECODER", "Windows Media Player cannot play the DVD because a compatible DVD decoder is not installed on your computer."}
	NsEDvdParental                                               = &Error{0xC00D116E, "NS_E_DVD_PARENTAL", "Windows Media Player cannot play the scene because it has a parental rating higher than the rating that you are authorized to view."}
	NsEDvdCannotJump                                             = &Error{0xC00D116F, "NS_E_DVD_CANNOT_JUMP", "Windows Media Player cannot skip to the requested location on the DVD."}
	NsEDvdDeviceContention                                       = &Error{0xC00D1170, "NS_E_DVD_DEVICE_CONTENTION", "Windows Media Player cannot play the DVD because it is currently in use by another program. Quit the other program that is using the DVD, and then try again."}
	NsEDvdNoVideoMemory                                          = &Error{0xC00D1171, "NS_E_DVD_NO_VIDEO_MEMORY", "Windows Media Player cannot play DVD video. You might need to adjust your Windows display settings. Open display settings in Control Panel, and then try lowering your screen resolution and color quality settings."}
	NsEDvdCannotCopyProtected                                    = &Error{0xC00D1172, "NS_E_DVD_CANNOT_COPY_PROTECTED", "Windows Media Player cannot rip the DVD because it is copy protected."}
	NsEDvdRequiredPropertyNotSet                                 = &Error{0xC00D1173, "NS_E_DVD_REQUIRED_PROPERTY_NOT_SET", "One of more of the required properties has not been set."}
	NsEDvdInvalidTitleChapter                                    = &Error{0xC00D1174, "NS_E_DVD_INVALID_TITLE_CHAPTER", "The specified title and/or chapter number does not exist on this DVD."}
	NsENoCdBurner                                                = &Error{0xC00D1176, "NS_E_NO_CD_BURNER", "Windows Media Player cannot burn the files because the Player cannot find a burner. If the burner is connected properly, try using Windows Update to install the latest device driver."}
	NsEDeviceIsNotReady                                          = &Error{0xC00D1177, "NS_E_DEVICE_IS_NOT_READY", "Windows Media Player does not detect storage media in the selected device. Insert storage media into the device, and then try again."}
	NsEPdaUnsupportedFormat                                      = &Error{0xC00D1178, "NS_E_PDA_UNSUPPORTED_FORMAT", "Windows Media Player cannot sync this file. The Player might not support the file type."}
	NsENoPda                                                     = &Error{0xC00D1179, "NS_E_NO_PDA", "Windows Media Player does not detect a portable device. Connect your portable device, and then try again."}
	NsEPdaUnspecifiedError                                       = &Error{0xC00D117A, "NS_E_PDA_UNSPECIFIED_ERROR", "Windows Media Player encountered an error while communicating with the device. The storage card on the device might be full, the device might be turned off, or the device might not allow playlists or folders to be created on it."}
	NsEMemstorageBadData                                         = &Error{0xC00D117B, "NS_E_MEMSTORAGE_BAD_DATA", "Windows Media Player encountered an error while burning a CD."}
	NsEPdaFailSelectDevice                                       = &Error{0xC00D117C, "NS_E_PDA_FAIL_SELECT_DEVICE", "Windows Media Player encountered an error while communicating with a portable device or CD drive."}
	NsEPdaFailReadWaveFile                                       = &Error{0xC00D117D, "NS_E_PDA_FAIL_READ_WAVE_FILE", "Windows Media Player cannot open the WAV file."}
	NsEImapiLossofstreaming                                      = &Error{0xC00D117E, "NS_E_IMAPI_LOSSOFSTREAMING", "Windows Media Player failed to burn all the files to the CD. Select a slower recording speed, and then try again."}
	NsEPdaDeviceFull                                             = &Error{0xC00D117F, "NS_E_PDA_DEVICE_FULL", "There is not enough storage space on the portable device to complete this operation. Delete some unneeded files on the portable device, and then try again."}
	NsEFailLaunchRoxioPlugin                                     = &Error{0xC00D1180, "NS_E_FAIL_LAUNCH_ROXIO_PLUGIN", "Windows Media Player cannot burn the files. Verify that your burner is connected properly, and then try again. If the problem persists, reinstall the Player."}
	NsEPdaDeviceFullInSession                                    = &Error{0xC00D1181, "NS_E_PDA_DEVICE_FULL_IN_SESSION", "Windows Media Player did not sync some files to the device because there is not enough storage space on the device."}
	NsEImapiMediumInvalidtype                                    = &Error{0xC00D1182, "NS_E_IMAPI_MEDIUM_INVALIDTYPE", "The disc in the burner is not valid. Insert a blank disc into the burner, and then try again."}
	NsEPdaManualdevice                                           = &Error{0xC00D1183, "NS_E_PDA_MANUALDEVICE", "Windows Media Player cannot perform the requested action because the device does not support sync."}
	NsEPdaPartnershipnotexist                                    = &Error{0xC00D1184, "NS_E_PDA_PARTNERSHIPNOTEXIST", "To perform the requested action, you must first set up sync with the device."}
	NsEPdaCannotCreateAdditionalSyncRelationship                 = &Error{0xC00D1185, "NS_E_PDA_CANNOT_CREATE_ADDITIONAL_SYNC_RELATIONSHIP", "You have already created sync partnerships with 16 devices. To create a new sync partnership, you must first end an existing partnership."}
	NsEPdaNoTranscodeOfDrm                                       = &Error{0xC00D1186, "NS_E_PDA_NO_TRANSCODE_OF_DRM", "Windows Media Player cannot sync the file because protected files cannot be converted to the required quality level or file format."}
	NsEPdaTranscodecachefull                                     = &Error{0xC00D1187, "NS_E_PDA_TRANSCODECACHEFULL", "The folder that stores converted files is full. Either empty the folder or increase its size, and then try again."}
	NsEPdaTooManyFileCollisions                                  = &Error{0xC00D1188, "NS_E_PDA_TOO_MANY_FILE_COLLISIONS", "There are too many files with the same name in the folder on the device. Change the file name or sync to a different folder."}
	NsEPdaCannotTranscode                                        = &Error{0xC00D1189, "NS_E_PDA_CANNOT_TRANSCODE", "Windows Media Player cannot convert the file to the format required by the device."}
	NsEPdaTooManyFilesInDirectory                                = &Error{0xC00D118A, "NS_E_PDA_TOO_MANY_FILES_IN_DIRECTORY", "You have reached the maximum number of files your device allows in a folder. If your device supports playback from subfolders, try creating subfolders on the device and storing some files in them."}
	NsEProcessingshowsyncwizard                                  = &Error{0xC00D118B, "NS_E_PROCESSINGSHOWSYNCWIZARD", "Windows Media Player is already trying to start the Device Setup Wizard."}
	NsEPdaTranscodeNotPermitted                                  = &Error{0xC00D118C, "NS_E_PDA_TRANSCODE_NOT_PERMITTED", "Windows Media Player cannot convert this file format. If an updated version of the codec used to compress this file is available, install it and then try to sync the file again."}
	NsEPdaInitializingdevices                                    = &Error{0xC00D118D, "NS_E_PDA_INITIALIZINGDEVICES", "Windows Media Player is busy setting up devices. Try again later."}
	NsEPdaObsoleteSp                                             = &Error{0xC00D118E, "NS_E_PDA_OBSOLETE_SP", "Your device is using an outdated driver that is no longer supported by Windows Media Player. For additional assistance, click Web Help."}
	NsEPdaTitleCollision                                         = &Error{0xC00D118F, "NS_E_PDA_TITLE_COLLISION", "Windows Media Player cannot sync the file because a file with the same name already exists on the device. Change the file name or try to sync the file to a different folder."}
	NsEPdaDevicesupportdisabled                                  = &Error{0xC00D1190, "NS_E_PDA_DEVICESUPPORTDISABLED", "Automatic and manual sync have been turned off temporarily. To sync to a device, restart Windows Media Player."}
	NsEPdaNoLongerAvailable                                      = &Error{0xC00D1191, "NS_E_PDA_NO_LONGER_AVAILABLE", "This device is not available. Connect the device to the computer, and then try again."}
	NsEPdaEncoderNotResponding                                   = &Error{0xC00D1192, "NS_E_PDA_ENCODER_NOT_RESPONDING", "Windows Media Player cannot sync the file because an error occurred while converting the file to another quality level or format. If the problem persists, remove the file from the list of files to sync."}
	NsEPdaCannotSyncFromLocation                                 = &Error{0xC00D1193, "NS_E_PDA_CANNOT_SYNC_FROM_LOCATION", "Windows Media Player cannot sync the file to your device. The file might be stored in a location that is not supported. Copy the file from its current location to your hard disk, add it to your library, and then try to sync the file again."}
	NsEWmpProtocolProblem                                        = &Error{0xC00D1194, "NS_E_WMP_PROTOCOL_PROBLEM", "Windows Media Player cannot open the specified URL. Verify that the Player is configured to use all available protocols, and then try again."}
	NsEWmpNoDiskSpace                                            = &Error{0xC00D1195, "NS_E_WMP_NO_DISK_SPACE", "Windows Media Player cannot perform the requested action because there is not enough storage space on your computer. Delete some unneeded files on your hard disk, and then try again."}
	NsEWmpLogonFailure                                           = &Error{0xC00D1196, "NS_E_WMP_LOGON_FAILURE", "The server denied access to the file. Verify that you are using the correct user name and password."}
	NsEWmpCannotFindFile                                         = &Error{0xC00D1197, "NS_E_WMP_CANNOT_FIND_FILE", "Windows Media Player cannot find the file. If you are trying to play, burn, or sync an item that is in your library, the item might point to a file that has been moved, renamed, or deleted."}
	NsEWmpServerInaccessible                                     = &Error{0xC00D1198, "NS_E_WMP_SERVER_INACCESSIBLE", "Windows Media Player cannot connect to the server. The server name might not be correct, the server might not be available, or your proxy settings might not be correct."}
	NsEWmpUnsupportedFormat                                      = &Error{0xC00D1199, "NS_E_WMP_UNSUPPORTED_FORMAT", "Windows Media Player cannot play the file. The Player might not support the file type or might not support the codec that was used to compress the file."}
	NsEWmpDshowUnsupportedFormat                                 = &Error{0xC00D119A, "NS_E_WMP_DSHOW_UNSUPPORTED_FORMAT", "Windows Media Player cannot play the file. The Player might not support the file type or a required codec might not be installed on your computer."}
	NsEWmpPlaylistExists                                         = &Error{0xC00D119B, "NS_E_WMP_PLAYLIST_EXISTS", "Windows Media Player cannot create the playlist because the name already exists. Type a different playlist name."}
	NsEWmpNonmediaFiles                                          = &Error{0xC00D119C, "NS_E_WMP_NONMEDIA_FILES", "Windows Media Player cannot delete the playlist because it contains items that are not digital media files. Any digital media files in the playlist were deleted."}
	NsEWmpInvalidAsx                                             = &Error{0xC00D119D, "NS_E_WMP_INVALID_ASX", "The playlist cannot be opened because it is stored in a shared folder on another computer. If possible, move the playlist to the playlists folder on your computer."}
	NsEWmpAlreadyInUse                                           = &Error{0xC00D119E, "NS_E_WMP_ALREADY_IN_USE", "Windows Media Player is already in use. Stop playing any items, close all Player dialog boxes, and then try again."}
	NsEWmpImapiFailure                                           = &Error{0xC00D119F, "NS_E_WMP_IMAPI_FAILURE", "Windows Media Player encountered an error while burning. Verify that the burner is connected properly and that the disc is clean and not damaged."}
	NsEWmpWmdmFailure                                            = &Error{0xC00D11A0, "NS_E_WMP_WMDM_FAILURE", "Windows Media Player has encountered an unknown error with your portable device. Reconnect your portable device, and then try again."}
	NsEWmpCodecNeededWith4cc                                     = &Error{0xC00D11A1, "NS_E_WMP_CODEC_NEEDED_WITH_4CC", "A codec is required to play this file. To determine if this codec is available to download from the web, click Web Help."}
	NsEWmpCodecNeededWithFormattag                               = &Error{0xC00D11A2, "NS_E_WMP_CODEC_NEEDED_WITH_FORMATTAG", "An audio codec is needed to play this file. To determine if this codec is available to download from the web, click Web Help."}
	NsEWmpMssapNotAvailable                                      = &Error{0xC00D11A3, "NS_E_WMP_MSSAP_NOT_AVAILABLE", "To play the file, you must install the latest Windows service pack. To install the service pack from the Windows Update website, click Web Help."}
	NsEWmpWmdmInterfacedead                                      = &Error{0xC00D11A4, "NS_E_WMP_WMDM_INTERFACEDEAD", "Windows Media Player no longer detects a portable device. Reconnect your portable device, and then try again."}
	NsEWmpWmdmNotcertified                                       = &Error{0xC00D11A5, "NS_E_WMP_WMDM_NOTCERTIFIED", "Windows Media Player cannot sync the file because the portable device does not support protected files."}
	NsEWmpWmdmLicenseNotexist                                    = &Error{0xC00D11A6, "NS_E_WMP_WMDM_LICENSE_NOTEXIST", "This file does not have sync rights. If you obtained this file from an online store, go to the online store to get sync rights."}
	NsEWmpWmdmLicenseExpired                                     = &Error{0xC00D11A7, "NS_E_WMP_WMDM_LICENSE_EXPIRED", "Windows Media Player cannot sync the file because the sync rights have expired. Go to the content provider's online store to get new sync rights."}
	NsEWmpWmdmBusy                                               = &Error{0xC00D11A8, "NS_E_WMP_WMDM_BUSY", "The portable device is already in use. Wait until the current task finishes or quit other programs that might be using the portable device, and then try again."}
	NsEWmpWmdmNorights                                           = &Error{0xC00D11A9, "NS_E_WMP_WMDM_NORIGHTS", "Windows Media Player cannot sync the file because the content provider or device prohibits it. You might be able to resolve this problem by going to the content provider's online store to get sync rights."}
	NsEWmpWmdmIncorrectRights                                    = &Error{0xC00D11AA, "NS_E_WMP_WMDM_INCORRECT_RIGHTS", "The content provider has not granted you the right to sync this file. Go to the content provider's online store to get sync rights."}
	NsEWmpImapiGeneric                                           = &Error{0xC00D11AB, "NS_E_WMP_IMAPI_GENERIC", "Windows Media Player cannot burn the files to the CD. Verify that the disc is clean and not damaged. If necessary, select a slower recording speed or try a different brand of blank discs."}
	NsEWmpImapiDeviceNotpresent                                  = &Error{0xC00D11AD, "NS_E_WMP_IMAPI_DEVICE_NOTPRESENT", "Windows Media Player cannot burn the files. Verify that the burner is connected properly, and then try again."}
	NsEWmpImapiDeviceBusy                                        = &Error{0xC00D11AE, "NS_E_WMP_IMAPI_DEVICE_BUSY", "Windows Media Player cannot burn the files. Verify that the burner is connected properly and that the disc is clean and not damaged. If the burner is already in use, wait until the current task finishes or quit other programs that might be using the burner."}
	NsEWmpImapiLossOfStreaming                                   = &Error{0xC00D11AF, "NS_E_WMP_IMAPI_LOSS_OF_STREAMING", "Windows Media Player cannot burn the files to the CD."}
	NsEWmpServerUnavailable                                      = &Error{0xC00D11B0, "NS_E_WMP_SERVER_UNAVAILABLE", "Windows Media Player cannot play the file. The server might not be available or there might be a problem with your network or firewall settings."}
	NsEWmpFileOpenFailed                                         = &Error{0xC00D11B1, "NS_E_WMP_FILE_OPEN_FAILED", "Windows Media Player encountered a problem while playing the file. For additional assistance, click Web Help."}
	NsEWmpVerifyOnline                                           = &Error{0xC00D11B2, "NS_E_WMP_VERIFY_ONLINE", "Windows Media Player must connect to the Internet to verify the file's media usage rights. Connect to the Internet, and then try again."}
	NsEWmpServerNotResponding                                    = &Error{0xC00D11B3, "NS_E_WMP_SERVER_NOT_RESPONDING", "Windows Media Player cannot play the file because a network error occurred. The server might not be available. Verify that you are connected to the network and that your proxy settings are correct."}
	NsEWmpDrmCorruptBackup                                       = &Error{0xC00D11B4, "NS_E_WMP_DRM_CORRUPT_BACKUP", "Windows Media Player cannot restore your media usage rights because it could not find any backed up rights on your computer."}
	NsEWmpDrmLicenseServerUnavailable                            = &Error{0xC00D11B5, "NS_E_WMP_DRM_LICENSE_SERVER_UNAVAILABLE", "Windows Media Player cannot download media usage rights because the server is not available (for example, the server might be busy or not online)."}
	NsEWmpNetworkFirewall                                        = &Error{0xC00D11B6, "NS_E_WMP_NETWORK_FIREWALL", "Windows Media Player cannot play the file. A network firewall might be preventing the Player from opening the file by using the UDP transport protocol. If you typed a URL in the Open URL dialog box, try using a different transport protocol (for example, \"http:\")."}
	NsEWmpNoRemovableMedia                                       = &Error{0xC00D11B7, "NS_E_WMP_NO_REMOVABLE_MEDIA", "Insert the removable media, and then try again."}
	NsEWmpProxyConnectTimeout                                    = &Error{0xC00D11B8, "NS_E_WMP_PROXY_CONNECT_TIMEOUT", "Windows Media Player cannot play the file because the proxy server is not responding. The proxy server might be temporarily unavailable or your Player proxy settings might not be valid."}
	NsEWmpNeedUpgrade                                            = &Error{0xC00D11B9, "NS_E_WMP_NEED_UPGRADE", "To play the file, you might need to install a later version of Windows Media Player. On the Help menu, click Check for Updates, and then follow the instructions. For additional assistance, click Web Help."}
	NsEWmpAudioHwProblem                                         = &Error{0xC00D11BA, "NS_E_WMP_AUDIO_HW_PROBLEM", "Windows Media Player cannot play the file because there is a problem with your sound device. There might not be a sound device installed on your computer, it might be in use by another program, or it might not be functioning properly."}
	NsEWmpInvalidProtocol                                        = &Error{0xC00D11BB, "NS_E_WMP_INVALID_PROTOCOL", "Windows Media Player cannot play the file because the specified protocol is not supported. If you typed a URL in the Open URL dialog box, try using a different transport protocol (for example, \"http:\" or \"rtsp:\")."}
	NsEWmpInvalidLibraryAdd                                      = &Error{0xC00D11BC, "NS_E_WMP_INVALID_LIBRARY_ADD", "Windows Media Player cannot add the file to the library because the file format is not supported."}
	NsEWmpMmsNotSupported                                        = &Error{0xC00D11BD, "NS_E_WMP_MMS_NOT_SUPPORTED", "Windows Media Player cannot play the file because the specified protocol is not supported. If you typed a URL in the Open URL dialog box, try using a different transport protocol (for example, \"mms:\")."}
	NsEWmpNoProtocolsSelected                                    = &Error{0xC00D11BE, "NS_E_WMP_NO_PROTOCOLS_SELECTED", "Windows Media Player cannot play the file because there are no streaming protocols selected. Select one or more protocols, and then try again."}
	NsEWmpGofullscreenFailed                                     = &Error{0xC00D11BF, "NS_E_WMP_GOFULLSCREEN_FAILED", "Windows Media Player cannot switch to Full Screen. You might need to adjust your Windows display settings. Open display settings in Control Panel, and then try setting Hardware acceleration to Full."}
	NsEWmpNetworkError                                           = &Error{0xC00D11C0, "NS_E_WMP_NETWORK_ERROR", "Windows Media Player cannot play the file because a network error occurred. The server might not be available (for example, the server is busy or not online) or you might not be connected to the network."}
	NsEWmpConnectTimeout                                         = &Error{0xC00D11C1, "NS_E_WMP_CONNECT_TIMEOUT", "Windows Media Player cannot play the file because the server is not responding. Verify that you are connected to the network, and then try again later."}
	NsEWmpMulticastDisabled                                      = &Error{0xC00D11C2, "NS_E_WMP_MULTICAST_DISABLED", "Windows Media Player cannot play the file because the multicast protocol is not enabled. On the Tools menu, click Options, click the Network tab, and then select the Multicast check box. For additional assistance, click Web Help."}
	NsEWmpServerDnsTimeout                                       = &Error{0xC00D11C3, "NS_E_WMP_SERVER_DNS_TIMEOUT", "Windows Media Player cannot play the file because a network problem occurred. Verify that you are connected to the network, and then try again later."}
	NsEWmpProxyNotFound                                          = &Error{0xC00D11C4, "NS_E_WMP_PROXY_NOT_FOUND", "Windows Media Player cannot play the file because the network proxy server cannot be found. Verify that your proxy settings are correct, and then try again."}
	NsEWmpTamperedContent                                        = &Error{0xC00D11C5, "NS_E_WMP_TAMPERED_CONTENT", "Windows Media Player cannot play the file because it is corrupted."}
	NsEWmpOutofmemory                                            = &Error{0xC00D11C6, "NS_E_WMP_OUTOFMEMORY", "Your computer is running low on memory. Quit other programs, and then try again."}
	NsEWmpAudioCodecNotInstalled                                 = &Error{0xC00D11C7, "NS_E_WMP_AUDIO_CODEC_NOT_INSTALLED", "Windows Media Player cannot play, burn, rip, or sync the file because a required audio codec is not installed on your computer."}
	NsEWmpVideoCodecNotInstalled                                 = &Error{0xC00D11C8, "NS_E_WMP_VIDEO_CODEC_NOT_INSTALLED", "Windows Media Player cannot play the file because the required video codec is not installed on your computer."}
	NsEWmpImapiDeviceInvalidtype                                 = &Error{0xC00D11C9, "NS_E_WMP_IMAPI_DEVICE_INVALIDTYPE", "Windows Media Player cannot burn the files. If the burner is busy, wait for the current task to finish. If necessary, verify that the burner is connected properly and that you have installed the latest device driver."}
	NsEWmpDrmDriverAuthFailure                                   = &Error{0xC00D11CA, "NS_E_WMP_DRM_DRIVER_AUTH_FAILURE", "Windows Media Player cannot play the protected file because there is a problem with your sound device. Try installing a new device driver or use a different sound device."}
	NsEWmpNetworkResourceFailure                                 = &Error{0xC00D11CB, "NS_E_WMP_NETWORK_RESOURCE_FAILURE", "Windows Media Player encountered a network error. Restart the Player."}
	NsEWmpUpgradeApplication                                     = &Error{0xC00D11CC, "NS_E_WMP_UPGRADE_APPLICATION", "Windows Media Player is not installed properly. Reinstall the Player."}
	NsEWmpUnknownError                                           = &Error{0xC00D11CD, "NS_E_WMP_UNKNOWN_ERROR", "Windows Media Player encountered an unknown error. For additional assistance, click Web Help."}
	NsEWmpInvalidKey                                             = &Error{0xC00D11CE, "NS_E_WMP_INVALID_KEY", "Windows Media Player cannot play the file because the required codec is not valid."}
	NsEWmpCdAnotherUser                                          = &Error{0xC00D11CF, "NS_E_WMP_CD_ANOTHER_USER", "The CD drive is in use by another user. Wait for the task to complete, and then try again."}
	NsEWmpDrmNeedsAuthorization                                  = &Error{0xC00D11D0, "NS_E_WMP_DRM_NEEDS_AUTHORIZATION", "Windows Media Player cannot play, sync, or burn the protected file because a problem occurred with the Windows Media Digital Rights Management (DRM) system. You might need to connect to the Internet to update your DRM components. For additional assistance, click Web Help."}
	NsEWmpBadDriver                                              = &Error{0xC00D11D1, "NS_E_WMP_BAD_DRIVER", "Windows Media Player cannot play the file because there might be a problem with your sound or video device. Try installing an updated device driver."}
	NsEWmpAccessDenied                                           = &Error{0xC00D11D2, "NS_E_WMP_ACCESS_DENIED", "Windows Media Player cannot access the file. The file might be in use, you might not have access to the computer where the file is stored, or your proxy settings might not be correct."}
	NsEWmpLicenseRestricts                                       = &Error{0xC00D11D3, "NS_E_WMP_LICENSE_RESTRICTS", "The content provider prohibits this action. Go to the content provider's online store to get new media usage rights."}
	NsEWmpInvalidRequest                                         = &Error{0xC00D11D4, "NS_E_WMP_INVALID_REQUEST", "Windows Media Player cannot perform the requested action at this time."}
	NsEWmpCdStashNoSpace                                         = &Error{0xC00D11D5, "NS_E_WMP_CD_STASH_NO_SPACE", "Windows Media Player cannot burn the files because there is not enough free disk space to store the temporary files. Delete some unneeded files on your hard disk, and then try again."}
	NsEWmpDrmNewHardware                                         = &Error{0xC00D11D6, "NS_E_WMP_DRM_NEW_HARDWARE", "Your media usage rights have become corrupted or are no longer valid. This might happen if you have replaced hardware components in your computer."}
	NsEWmpDrmInvalidSig                                          = &Error{0xC00D11D7, "NS_E_WMP_DRM_INVALID_SIG", "The required Windows Media Digital Rights Management (DRM) component cannot be validated. You might be able resolve the problem by reinstalling the Player."}
	NsEWmpDrmCannotRestore                                       = &Error{0xC00D11D8, "NS_E_WMP_DRM_CANNOT_RESTORE", "You have exceeded your restore limit for the day. Try restoring your media usage rights tomorrow."}
	NsEWmpBurnDiscOverflow                                       = &Error{0xC00D11D9, "NS_E_WMP_BURN_DISC_OVERFLOW", "Some files might not fit on the CD. The required space cannot be calculated accurately because some files might be missing duration information. To ensure the calculation is accurate, play the files that are missing duration information."}
	NsEWmpDrmGenericLicenseFailure                               = &Error{0xC00D11DA, "NS_E_WMP_DRM_GENERIC_LICENSE_FAILURE", "Windows Media Player cannot verify the file's media usage rights. If you obtained this file from an online store, go to the online store to get the necessary rights."}
	NsEWmpDrmNoSecureClock                                       = &Error{0xC00D11DB, "NS_E_WMP_DRM_NO_SECURE_CLOCK", "It is not possible to sync because this device's internal clock is not set correctly. To set the clock, select the option to set the device clock on the Privacy tab of the Options dialog box, connect to the Internet, and then sync the device again. For additional assistance, click Web Help."}
	NsEWmpDrmNoRights                                            = &Error{0xC00D11DC, "NS_E_WMP_DRM_NO_RIGHTS", "Windows Media Player cannot play, burn, rip, or sync the protected file because you do not have the appropriate rights."}
	NsEWmpDrmIndivFailed                                         = &Error{0xC00D11DD, "NS_E_WMP_DRM_INDIV_FAILED", "Windows Media Player encountered an error during upgrade."}
	NsEWmpServerNonewconnections                                 = &Error{0xC00D11DE, "NS_E_WMP_SERVER_NONEWCONNECTIONS", "Windows Media Player cannot connect to the server because it is not accepting any new connections. This could be because it has reached its maximum connection limit. Please try again later."}
	NsEWmpMultipleErrorInPlaylist                                = &Error{0xC00D11DF, "NS_E_WMP_MULTIPLE_ERROR_IN_PLAYLIST", "A number of queued files cannot be played. To find information about the problem, click the Now Playing tab, and then click the icon next to each file in the List pane."}
	NsEWmpImapi2EraseFail                                        = &Error{0xC00D11E0, "NS_E_WMP_IMAPI2_ERASE_FAIL", "Windows Media Player encountered an error while erasing the rewritable CD or DVD. Verify that the CD or DVD burner is connected properly and that the disc is clean and not damaged."}
	NsEWmpImapi2EraseDeviceBusy                                  = &Error{0xC00D11E1, "NS_E_WMP_IMAPI2_ERASE_DEVICE_BUSY", "Windows Media Player cannot erase the rewritable CD or DVD. Verify that the CD or DVD burner is connected properly and that the disc is clean and not damaged. If the burner is already in use, wait until the current task finishes or quit other programs that might be using the burner."}
	NsEWmpDrmComponentFailure                                    = &Error{0xC00D11E2, "NS_E_WMP_DRM_COMPONENT_FAILURE", "A Windows Media Digital Rights Management (DRM) component encountered a problem. If you are trying to use a file that you obtained from an online store, try going to the online store and getting the appropriate usage rights."}
	NsEWmpDrmNoDeviceCert                                        = &Error{0xC00D11E3, "NS_E_WMP_DRM_NO_DEVICE_CERT", "It is not possible to obtain device's certificate. Please contact the device manufacturer for a firmware update or for other steps to resolve this problem."}
	NsEWmpServerSecurityError                                    = &Error{0xC00D11E4, "NS_E_WMP_SERVER_SECURITY_ERROR", "Windows Media Player encountered an error when connecting to the server. The security information from the server could not be validated."}
	NsEWmpAudioDeviceLost                                        = &Error{0xC00D11E5, "NS_E_WMP_AUDIO_DEVICE_LOST", "An audio device was disconnected or reconfigured. Verify that the audio device is connected, and then try to play the item again."}
	NsEWmpImapiMediaIncompatible                                 = &Error{0xC00D11E6, "NS_E_WMP_IMAPI_MEDIA_INCOMPATIBLE", "Windows Media Player could not complete burning because the disc is not compatible with your drive. Try inserting a different kind of recordable media or use a disc that supports a write speed that is compatible with your drive."}
	NsESyncwizDeviceFull                                         = &Error{0xC00D11EE, "NS_E_SYNCWIZ_DEVICE_FULL", "Windows Media Player cannot save the sync settings because your device is full. Delete some unneeded files on your device and then try again."}
	NsESyncwizCannotChangeSettings                               = &Error{0xC00D11EF, "NS_E_SYNCWIZ_CANNOT_CHANGE_SETTINGS", "It is not possible to change sync settings at this time. Try again later."}
	NsETranscodeDeletecacheerror                                 = &Error{0xC00D11F0, "NS_E_TRANSCODE_DELETECACHEERROR", "Windows Media Player cannot delete these files currently. If the Player is synchronizing, wait until it is complete and then try again."}
	NsECdNoBuffersRead                                           = &Error{0xC00D11F8, "NS_E_CD_NO_BUFFERS_READ", "Windows Media Player could not use digital mode to read the CD. The Player has automatically switched the CD drive to analog mode. To switch back to digital mode, use the Devices tab. For additional assistance, click Web Help."}
	NsECdEmptyTrackQueue                                         = &Error{0xC00D11F9, "NS_E_CD_EMPTY_TRACK_QUEUE", "No CD track was specified for playback."}
	NsECdNoReader                                                = &Error{0xC00D11FA, "NS_E_CD_NO_READER", "The CD filter was not able to create the CD reader."}
	NsECdIsrcInvalid                                             = &Error{0xC00D11FB, "NS_E_CD_ISRC_INVALID", "Invalid ISRC code."}
	NsECdMediaCatalogNumberInvalid                               = &Error{0xC00D11FC, "NS_E_CD_MEDIA_CATALOG_NUMBER_INVALID", "Invalid Media Catalog Number."}
	NsESlowReadDigitalWithErrorcorrection                        = &Error{0xC00D11FD, "NS_E_SLOW_READ_DIGITAL_WITH_ERRORCORRECTION", "Windows Media Player cannot play audio CDs correctly because the CD drive is slow and error correction is turned on. To increase performance, turn off playback error correction for this drive."}
	NsECdSpeeddetectNotEnoughReads                               = &Error{0xC00D11FE, "NS_E_CD_SPEEDDETECT_NOT_ENOUGH_READS", "Windows Media Player cannot estimate the CD drive's playback speed because the CD track is too short."}
	NsECdQueueingDisabled                                        = &Error{0xC00D11FF, "NS_E_CD_QUEUEING_DISABLED", "Cannot queue the CD track because queuing is not enabled."}
	NsEWmpDrmAcquiringLicense                                    = &Error{0xC00D1202, "NS_E_WMP_DRM_ACQUIRING_LICENSE", "Windows Media Player cannot download additional media usage rights until the current download is complete."}
	NsEWmpDrmLicenseExpired                                      = &Error{0xC00D1203, "NS_E_WMP_DRM_LICENSE_EXPIRED", "The media usage rights for this file have expired or are no longer valid. If you obtained the file from an online store, sign in to the store, and then try again."}
	NsEWmpDrmLicenseNotacquired                                  = &Error{0xC00D1204, "NS_E_WMP_DRM_LICENSE_NOTACQUIRED", "Windows Media Player cannot download the media usage rights for the file. If you obtained the file from an online store, sign in to the store, and then try again."}
	NsEWmpDrmLicenseNotenabled                                   = &Error{0xC00D1205, "NS_E_WMP_DRM_LICENSE_NOTENABLED", "The media usage rights for this file are not yet valid. To see when they will become valid, right-click the file in the library, click Properties, and then click the Media Usage Rights tab."}
	NsEWmpDrmLicenseUnusable                                     = &Error{0xC00D1206, "NS_E_WMP_DRM_LICENSE_UNUSABLE", "The media usage rights for this file are not valid. If you obtained this file from an online store, contact the store for assistance."}
	NsEWmpDrmLicenseContentRevoked                               = &Error{0xC00D1207, "NS_E_WMP_DRM_LICENSE_CONTENT_REVOKED", "The content provider has revoked the media usage rights for this file. If you obtained this file from an online store, ask the store if a new version of the file is available."}
	NsEWmpDrmLicenseNosap                                        = &Error{0xC00D1208, "NS_E_WMP_DRM_LICENSE_NOSAP", "The media usage rights for this file require a feature that is not supported in your current version of Windows Media Player or your current version of Windows. Try installing the latest version of the Player. If you obtained this file from an online store, contact the store for further assistance."}
	NsEWmpDrmUnableToAcquireLicense                              = &Error{0xC00D1209, "NS_E_WMP_DRM_UNABLE_TO_ACQUIRE_LICENSE", "Windows Media Player cannot download media usage rights at this time. Try again later."}
	NsEWmpLicenseRequired                                        = &Error{0xC00D120A, "NS_E_WMP_LICENSE_REQUIRED", "Windows Media Player cannot play, burn, or sync the file because the media usage rights are missing. If you obtained the file from an online store, sign in to the store, and then try again."}
	NsEWmpProtectedContent                                       = &Error{0xC00D120B, "NS_E_WMP_PROTECTED_CONTENT", "Windows Media Player cannot play, burn, or sync the file because the media usage rights are missing. If you obtained the file from an online store, sign in to the store, and then try again."}
	NsEWmpPolicyValueNotConfigured                               = &Error{0xC00D122A, "NS_E_WMP_POLICY_VALUE_NOT_CONFIGURED", "Windows Media Player cannot read a policy. This can occur when the policy does not exist in the registry or when the registry cannot be read."}
	NsEPdaCannotSyncFromInternet                                 = &Error{0xC00D1234, "NS_E_PDA_CANNOT_SYNC_FROM_INTERNET", "Windows Media Player cannot sync content streamed directly from the Internet. If possible, download the file to your computer, and then try to sync the file."}
	NsEPdaCannotSyncInvalidPlaylist                              = &Error{0xC00D1235, "NS_E_PDA_CANNOT_SYNC_INVALID_PLAYLIST", "This playlist is not valid or is corrupted. Create a new playlist using Windows Media Player, then sync the new playlist instead."}
	NsEPdaFailedToSynchronizeFile                                = &Error{0xC00D1236, "NS_E_PDA_FAILED_TO_SYNCHRONIZE_FILE", "Windows Media Player encountered a problem while synchronizing the file to the device. For additional assistance, click Web Help."}
	NsEPdaSyncFailed                                             = &Error{0xC00D1237, "NS_E_PDA_SYNC_FAILED", "Windows Media Player encountered an error while synchronizing to the device."}
	NsEPdaDeleteFailed                                           = &Error{0xC00D1238, "NS_E_PDA_DELETE_FAILED", "Windows Media Player cannot delete a file from the device."}
	NsEPdaFailedToRetrieveFile                                   = &Error{0xC00D1239, "NS_E_PDA_FAILED_TO_RETRIEVE_FILE", "Windows Media Player cannot copy a file from the device to your library."}
	NsEPdaDeviceNotResponding                                    = &Error{0xC00D123A, "NS_E_PDA_DEVICE_NOT_RESPONDING", "Windows Media Player cannot communicate with the device because the device is not responding. Try reconnecting the device, resetting the device, or contacting the device manufacturer for updated firmware."}
	NsEPdaFailedToTranscodePhoto                                 = &Error{0xC00D123B, "NS_E_PDA_FAILED_TO_TRANSCODE_PHOTO", "Windows Media Player cannot sync the picture to the device because a problem occurred while converting the file to another quality level or format. The original file might be damaged or corrupted."}
	NsEPdaFailedToEncryptTranscodedFile                          = &Error{0xC00D123C, "NS_E_PDA_FAILED_TO_ENCRYPT_TRANSCODED_FILE", "Windows Media Player cannot convert the file. The file might have been encrypted by the Encrypted File System (EFS). Try decrypting the file first and then synchronizing it. For information about how to decrypt a file, see Windows Help and Support."}
	NsEPdaCannotTranscodeToAudio                                 = &Error{0xC00D123D, "NS_E_PDA_CANNOT_TRANSCODE_TO_AUDIO", "Your device requires that this file be converted in order to play on the device. However, the device either does not support playing audio, or Windows Media Player cannot convert the file to an audio format that is supported by the device."}
	NsEPdaCannotTranscodeToVideo                                 = &Error{0xC00D123E, "NS_E_PDA_CANNOT_TRANSCODE_TO_VIDEO", "Your device requires that this file be converted in order to play on the device. However, the device either does not support playing video, or Windows Media Player cannot convert the file to a video format that is supported by the device."}
	NsEPdaCannotTranscodeToImage                                 = &Error{0xC00D123F, "NS_E_PDA_CANNOT_TRANSCODE_TO_IMAGE", "Your device requires that this file be converted in order to play on the device. However, the device either does not support displaying pictures, or Windows Media Player cannot convert the file to a picture format that is supported by the device."}
	NsEPdaRetrievedFileFilenameTooLong                           = &Error{0xC00D1240, "NS_E_PDA_RETRIEVED_FILE_FILENAME_TOO_LONG", "Windows Media Player cannot sync the file to your computer because the file name is too long. Try renaming the file on the device."}
	NsEPdaCewmdmDrmError                                         = &Error{0xC00D1241, "NS_E_PDA_CEWMDM_DRM_ERROR", "Windows Media Player cannot sync the file because the device is not responding. This typically occurs when there is a problem with the device firmware. For additional assistance, click Web Help."}
	NsEIncompletePlaylist                                        = &Error{0xC00D1242, "NS_E_INCOMPLETE_PLAYLIST", "Incomplete playlist."}
	NsEPdaSyncRunning                                            = &Error{0xC00D1243, "NS_E_PDA_SYNC_RUNNING", "It is not possible to perform the requested action because sync is in progress. You can either stop sync or wait for it to complete, and then try again."}
	NsEPdaSyncLoginError                                         = &Error{0xC00D1244, "NS_E_PDA_SYNC_LOGIN_ERROR", "Windows Media Player cannot sync the subscription content because you are not signed in to the online store that provided it. Sign in to the online store, and then try again."}
	NsEPdaTranscodeCodecNotFound                                 = &Error{0xC00D1245, "NS_E_PDA_TRANSCODE_CODEC_NOT_FOUND", "Windows Media Player cannot convert the file to the format required by the device. One or more codecs required to convert the file could not be found."}
	NsECannotSyncDrmToNonJanusDevice                             = &Error{0xC00D1246, "NS_E_CANNOT_SYNC_DRM_TO_NON_JANUS_DEVICE", "It is not possible to sync subscription files to this device."}
	NsECannotSyncPreviousSyncRunning                             = &Error{0xC00D1247, "NS_E_CANNOT_SYNC_PREVIOUS_SYNC_RUNNING", "Your device is operating slowly or is not responding. Until the device responds, it is not possible to sync again. To return the device to normal operation, try disconnecting it from the computer or resetting it."}
	NsEWmpHwndNotfound                                           = &Error{0xC00D125C, "NS_E_WMP_HWND_NOTFOUND", "The Windows Media Player download manager cannot function properly because the Player main window cannot be found. Try restarting the Player."}
	NsEBkgdownloadWrongNoFiles                                   = &Error{0xC00D125D, "NS_E_BKGDOWNLOAD_WRONG_NO_FILES", "Windows Media Player encountered a download that has the wrong number of files. This might occur if another program is trying to create jobs with the same signature as the Player."}
	NsEBkgdownloadCompletecancelledjob                           = &Error{0xC00D125E, "NS_E_BKGDOWNLOAD_COMPLETECANCELLEDJOB", "Windows Media Player tried to complete a download that was already canceled. The file will not be available."}
	NsEBkgdownloadCancelcompletedjob                             = &Error{0xC00D125F, "NS_E_BKGDOWNLOAD_CANCELCOMPLETEDJOB", "Windows Media Player tried to cancel a download that was already completed. The file will not be removed."}
	NsEBkgdownloadNojobpointer                                   = &Error{0xC00D1260, "NS_E_BKGDOWNLOAD_NOJOBPOINTER", "Windows Media Player is trying to access a download that is not valid."}
	NsEBkgdownloadInvalidjobsignature                            = &Error{0xC00D1261, "NS_E_BKGDOWNLOAD_INVALIDJOBSIGNATURE", "This download was not created by Windows Media Player."}
	NsEBkgdownloadFailedToCreateTempfile                         = &Error{0xC00D1262, "NS_E_BKGDOWNLOAD_FAILED_TO_CREATE_TEMPFILE", "The Windows Media Player download manager cannot create a temporary file name. This might occur if the path is not valid or if the disk is full."}
	NsEBkgdownloadPluginFailedinitialize                         = &Error{0xC00D1263, "NS_E_BKGDOWNLOAD_PLUGIN_FAILEDINITIALIZE", "The Windows Media Player download manager plug-in cannot start. This might occur if the system is out of resources."}
	NsEBkgdownloadPluginFailedtomovefile                         = &Error{0xC00D1264, "NS_E_BKGDOWNLOAD_PLUGIN_FAILEDTOMOVEFILE", "The Windows Media Player download manager cannot move the file."}
	NsEBkgdownloadCallfuncfailed                                 = &Error{0xC00D1265, "NS_E_BKGDOWNLOAD_CALLFUNCFAILED", "The Windows Media Player download manager cannot perform a task because the system has no resources to allocate."}
	NsEBkgdownloadCallfunctimeout                                = &Error{0xC00D1266, "NS_E_BKGDOWNLOAD_CALLFUNCTIMEOUT", "The Windows Media Player download manager cannot perform a task because the task took too long to run."}
	NsEBkgdownloadCallfuncended                                  = &Error{0xC00D1267, "NS_E_BKGDOWNLOAD_CALLFUNCENDED", "The Windows Media Player download manager cannot perform a task because the Player is terminating the service. The task will be recovered when the Player restarts."}
	NsEBkgdownloadWmdunpackfailed                                = &Error{0xC00D1268, "NS_E_BKGDOWNLOAD_WMDUNPACKFAILED", "The Windows Media Player download manager cannot expand a WMD file. The file will be deleted and the operation will not be completed successfully."}
	NsEBkgdownloadFailedinitialize                               = &Error{0xC00D1269, "NS_E_BKGDOWNLOAD_FAILEDINITIALIZE", "The Windows Media Player download manager cannot start. This might occur if the system is out of resources."}
	NsEInterfaceNotRegisteredInGit                               = &Error{0xC00D126A, "NS_E_INTERFACE_NOT_REGISTERED_IN_GIT", "Windows Media Player cannot access a required functionality. This might occur if the wrong system files or Player DLLs are loaded."}
	NsEBkgdownloadInvalidFileName                                = &Error{0xC00D126B, "NS_E_BKGDOWNLOAD_INVALID_FILE_NAME", "Windows Media Player cannot get the file name of the requested download. The requested download will be canceled."}
	NsEImageDownloadFailed                                       = &Error{0xC00D128E, "NS_E_IMAGE_DOWNLOAD_FAILED", "Windows Media Player encountered an error while downloading an image."}
	NsEWmpUdrmNouserlist                                         = &Error{0xC00D12C0, "NS_E_WMP_UDRM_NOUSERLIST", "Windows Media Player cannot update your media usage rights because the Player cannot verify the list of activated users of this computer."}
	NsEWmpDrmNotAcquiring                                        = &Error{0xC00D12C1, "NS_E_WMP_DRM_NOT_ACQUIRING", "Windows Media Player is trying to acquire media usage rights for a file that is no longer being used. Rights acquisition will stop."}
	NsEWmpBstrTooLong                                            = &Error{0xC00D12F2, "NS_E_WMP_BSTR_TOO_LONG", "The parameter is not valid."}
	NsEWmpAutoplayInvalidState                                   = &Error{0xC00D12FC, "NS_E_WMP_AUTOPLAY_INVALID_STATE", "The state is not valid for this request."}
	NsEWmpComponentRevoked                                       = &Error{0xC00D1306, "NS_E_WMP_COMPONENT_REVOKED", "Windows Media Player cannot play this file until you complete the software component upgrade. After the component has been upgraded, try to play the file again."}
	NsECurlNotsafe                                               = &Error{0xC00D1324, "NS_E_CURL_NOTSAFE", "The URL is not safe for the operation specified."}
	NsECurlInvalidchar                                           = &Error{0xC00D1325, "NS_E_CURL_INVALIDCHAR", "The URL contains one or more characters that are not valid."}
	NsECurlInvalidhostname                                       = &Error{0xC00D1326, "NS_E_CURL_INVALIDHOSTNAME", "The URL contains a host name that is not valid."}
	NsECurlInvalidpath                                           = &Error{0xC00D1327, "NS_E_CURL_INVALIDPATH", "The URL contains a path that is not valid."}
	NsECurlInvalidscheme                                         = &Error{0xC00D1328, "NS_E_CURL_INVALIDSCHEME", "The URL contains a scheme that is not valid."}
	NsECurlInvalidurl                                            = &Error{0xC00D1329, "NS_E_CURL_INVALIDURL", "The URL is not valid."}
	NsECurlCantwalk                                              = &Error{0xC00D132B, "NS_E_CURL_CANTWALK", "Windows Media Player cannot play the file. If you clicked a link on a web page, the link might not be valid."}
	NsECurlInvalidport                                           = &Error{0xC00D132C, "NS_E_CURL_INVALIDPORT", "The URL port is not valid."}
	NsECurlhelperNotadirectory                                   = &Error{0xC00D132D, "NS_E_CURLHELPER_NOTADIRECTORY", "The URL is not a directory."}
	NsECurlhelperNotafile                                        = &Error{0xC00D132E, "NS_E_CURLHELPER_NOTAFILE", "The URL is not a file."}
	NsECurlCantdecode                                            = &Error{0xC00D132F, "NS_E_CURL_CANTDECODE", "The URL contains characters that cannot be decoded. The URL might be truncated or incomplete."}
	NsECurlhelperNotrelative                                     = &Error{0xC00D1330, "NS_E_CURLHELPER_NOTRELATIVE", "The specified URL is not a relative URL."}
	NsECurlInvalidbuffersize                                     = &Error{0xC00D1331, "NS_E_CURL_INVALIDBUFFERSIZE", "The buffer is smaller than the size specified."}
	NsESubscriptionservicePlaybackDisallowed                     = &Error{0xC00D1356, "NS_E_SUBSCRIPTIONSERVICE_PLAYBACK_DISALLOWED", "The content provider has not granted you the right to play this file. Go to the content provider's online store to get play rights."}
	NsECannotBuyOrDownloadFromMultipleServices                   = &Error{0xC00D1357, "NS_E_CANNOT_BUY_OR_DOWNLOAD_FROM_MULTIPLE_SERVICES", "Windows Media Player cannot purchase or download content from multiple online stores."}
	NsECannotBuyOrDownloadContent                                = &Error{0xC00D1358, "NS_E_CANNOT_BUY_OR_DOWNLOAD_CONTENT", "The file cannot be purchased or downloaded. The file might not be available from the online store."}
	NsENotContentPartnerTrack                                    = &Error{0xC00D135A, "NS_E_NOT_CONTENT_PARTNER_TRACK", "The provider of this file cannot be identified."}
	NsETrackDownloadRequiresAlbumPurchase                        = &Error{0xC00D135B, "NS_E_TRACK_DOWNLOAD_REQUIRES_ALBUM_PURCHASE", "The file is only available for download when you buy the entire album."}
	NsETrackDownloadRequiresPurchase                             = &Error{0xC00D135C, "NS_E_TRACK_DOWNLOAD_REQUIRES_PURCHASE", "You must buy the file before you can download it."}
	NsETrackPurchaseMaximumExceeded                              = &Error{0xC00D135D, "NS_E_TRACK_PURCHASE_MAXIMUM_EXCEEDED", "You have exceeded the maximum number of files that can be purchased in a single transaction."}
	NsESubscriptionserviceLoginFailed                            = &Error{0xC00D135F, "NS_E_SUBSCRIPTIONSERVICE_LOGIN_FAILED", "Windows Media Player cannot sign in to the online store. Verify that you are using the correct user name and password. If the problem persists, the store might be temporarily unavailable."}
	NsESubscriptionserviceDownloadTimeout                        = &Error{0xC00D1360, "NS_E_SUBSCRIPTIONSERVICE_DOWNLOAD_TIMEOUT", "Windows Media Player cannot download this item because the server is not responding. The server might be temporarily unavailable or the Internet connection might be lost."}
	NsEContentPartnerStillInitializing                           = &Error{0xC00D1362, "NS_E_CONTENT_PARTNER_STILL_INITIALIZING", "Content Partner still initializing."}
	NsEOpenContainingFolderFailed                                = &Error{0xC00D1363, "NS_E_OPEN_CONTAINING_FOLDER_FAILED", "The folder could not be opened. The folder might have been moved or deleted."}
	NsEAdvancededitTooManyPictures                               = &Error{0xC00D136A, "NS_E_ADVANCEDEDIT_TOO_MANY_PICTURES", "Windows Media Player could not add all of the images to the file because the images exceeded the 7 megabyte (MB) limit."}
	NsERedirect                                                  = &Error{0xC00D1388, "NS_E_REDIRECT", "The client redirected to another server."}
	NsEStalePresentation                                         = &Error{0xC00D1389, "NS_E_STALE_PRESENTATION", "The streaming media description is no longer current."}
	NsENamespaceWrongPersist                                     = &Error{0xC00D138A, "NS_E_NAMESPACE_WRONG_PERSIST", "It is not possible to create a persistent namespace node under a transient parent node."}
	NsENamespaceWrongType                                        = &Error{0xC00D138B, "NS_E_NAMESPACE_WRONG_TYPE", "It is not possible to store a value in a namespace node that has a different value type."}
	NsENamespaceNodeConflict                                     = &Error{0xC00D138C, "NS_E_NAMESPACE_NODE_CONFLICT", "It is not possible to remove the root namespace node."}
	NsENamespaceNodeNotFound                                     = &Error{0xC00D138D, "NS_E_NAMESPACE_NODE_NOT_FOUND", "The specified namespace node could not be found."}
	NsENamespaceBufferTooSmall                                   = &Error{0xC00D138E, "NS_E_NAMESPACE_BUFFER_TOO_SMALL", "The buffer supplied to hold namespace node string is too small."}
	NsENamespaceTooManyCallbacks                                 = &Error{0xC00D138F, "NS_E_NAMESPACE_TOO_MANY_CALLBACKS", "The callback list on a namespace node is at the maximum size."}
	NsENamespaceDuplicateCallback                                = &Error{0xC00D1390, "NS_E_NAMESPACE_DUPLICATE_CALLBACK", "It is not possible to register an already-registered callback on a namespace node."}
	NsENamespaceCallbackNotFound                                 = &Error{0xC00D1391, "NS_E_NAMESPACE_CALLBACK_NOT_FOUND", "Cannot find the callback in the namespace when attempting to remove the callback."}
	NsENamespaceNameTooLong                                      = &Error{0xC00D1392, "NS_E_NAMESPACE_NAME_TOO_LONG", "The namespace node name exceeds the allowed maximum length."}
	NsENamespaceDuplicateName                                    = &Error{0xC00D1393, "NS_E_NAMESPACE_DUPLICATE_NAME", "Cannot create a namespace node that already exists."}
	NsENamespaceEmptyName                                        = &Error{0xC00D1394, "NS_E_NAMESPACE_EMPTY_NAME", "The namespace node name cannot be a null string."}
	NsENamespaceIndexTooLarge                                    = &Error{0xC00D1395, "NS_E_NAMESPACE_INDEX_TOO_LARGE", "Finding a child namespace node by index failed because the index exceeded the number of children."}
	NsENamespaceBadName                                          = &Error{0xC00D1396, "NS_E_NAMESPACE_BAD_NAME", "The namespace node name is invalid."}
	NsENamespaceWrongSecurity                                    = &Error{0xC00D1397, "NS_E_NAMESPACE_WRONG_SECURITY", "It is not possible to store a value in a namespace node that has a different security type."}
	NsECacheArchiveConflict                                      = &Error{0xC00D13EC, "NS_E_CACHE_ARCHIVE_CONFLICT", "The archive request conflicts with other requests in progress."}
	NsECacheOriginServerNotFound                                 = &Error{0xC00D13ED, "NS_E_CACHE_ORIGIN_SERVER_NOT_FOUND", "The specified origin server cannot be found."}
	NsECacheOriginServerTimeout                                  = &Error{0xC00D13EE, "NS_E_CACHE_ORIGIN_SERVER_TIMEOUT", "The specified origin server is not responding."}
	NsECacheNotBroadcast                                         = &Error{0xC00D13EF, "NS_E_CACHE_NOT_BROADCAST", "The internal code for HTTP status code 412 Precondition Failed due to not broadcast type."}
	NsECacheCannotBeCached                                       = &Error{0xC00D13F0, "NS_E_CACHE_CANNOT_BE_CACHED", "The internal code for HTTP status code 403 Forbidden due to not cacheable."}
	NsECacheNotModified                                          = &Error{0xC00D13F1, "NS_E_CACHE_NOT_MODIFIED", "The internal code for HTTP status code 304 Not Modified."}
	NsECannotRemovePublishingPoint                               = &Error{0xC00D1450, "NS_E_CANNOT_REMOVE_PUBLISHING_POINT", "It is not possible to remove a cache or proxy publishing point."}
	NsECannotRemovePlugin                                        = &Error{0xC00D1451, "NS_E_CANNOT_REMOVE_PLUGIN", "It is not possible to remove the last instance of a type of plug-in."}
	NsEWrongPublishingPointType                                  = &Error{0xC00D1452, "NS_E_WRONG_PUBLISHING_POINT_TYPE", "Cache and proxy publishing points do not support this property or method."}
	NsEUnsupportedLoadType                                       = &Error{0xC00D1453, "NS_E_UNSUPPORTED_LOAD_TYPE", "The plug-in does not support the specified load type."}
	NsEInvalidPluginLoadTypeConfiguration                        = &Error{0xC00D1454, "NS_E_INVALID_PLUGIN_LOAD_TYPE_CONFIGURATION", "The plug-in does not support any load types. The plug-in must support at least one load type."}
	NsEInvalidPublishingPointName                                = &Error{0xC00D1455, "NS_E_INVALID_PUBLISHING_POINT_NAME", "The publishing point name is invalid."}
	NsETooManyMulticastSinks                                     = &Error{0xC00D1456, "NS_E_TOO_MANY_MULTICAST_SINKS", "Only one multicast data writer plug-in can be enabled for a publishing point."}
	NsEPublishingPointInvalidRequestWhileStarted                 = &Error{0xC00D1457, "NS_E_PUBLISHING_POINT_INVALID_REQUEST_WHILE_STARTED", "The requested operation cannot be completed while the publishing point is started."}
	NsEMulticastPluginNotEnabled                                 = &Error{0xC00D1458, "NS_E_MULTICAST_PLUGIN_NOT_ENABLED", "A multicast data writer plug-in must be enabled in order for this operation to be completed."}
	NsEInvalidOperatingSystemVersion                             = &Error{0xC00D1459, "NS_E_INVALID_OPERATING_SYSTEM_VERSION", "This feature requires Windows Server 2003, Enterprise Edition."}
	NsEPublishingPointRemoved                                    = &Error{0xC00D145A, "NS_E_PUBLISHING_POINT_REMOVED", "The requested operation cannot be completed because the specified publishing point has been removed."}
	NsEInvalidPushPublishingPointStartRequest                    = &Error{0xC00D145B, "NS_E_INVALID_PUSH_PUBLISHING_POINT_START_REQUEST", "Push publishing points are started when the encoder starts pushing the stream. This publishing point cannot be started by the server administrator."}
	NsEUnsupportedLanguage                                       = &Error{0xC00D145C, "NS_E_UNSUPPORTED_LANGUAGE", "The specified language is not supported."}
	NsEWrongOsVersion                                            = &Error{0xC00D145D, "NS_E_WRONG_OS_VERSION", "Windows Media Services will only run on Windows Server 2003, Standard Edition and Windows Server 2003, Enterprise Edition."}
	NsEPublishingPointStopped                                    = &Error{0xC00D145E, "NS_E_PUBLISHING_POINT_STOPPED", "The operation cannot be completed because the publishing point has been stopped."}
	NsEPlaylistEntryAlreadyPlaying                               = &Error{0xC00D14B4, "NS_E_PLAYLIST_ENTRY_ALREADY_PLAYING", "The playlist entry is already playing."}
	NsEEmptyPlaylist                                             = &Error{0xC00D14B5, "NS_E_EMPTY_PLAYLIST", "The playlist or directory you are requesting does not contain content."}
	NsEPlaylistParseFailure                                      = &Error{0xC00D14B6, "NS_E_PLAYLIST_PARSE_FAILURE", "The server was unable to parse the requested playlist file."}
	NsEPlaylistUnsupportedEntry                                  = &Error{0xC00D14B7, "NS_E_PLAYLIST_UNSUPPORTED_ENTRY", "The requested operation is not supported for this type of playlist entry."}
	NsEPlaylistEntryNotInPlaylist                                = &Error{0xC00D14B8, "NS_E_PLAYLIST_ENTRY_NOT_IN_PLAYLIST", "Cannot jump to a playlist entry that is not inserted in the playlist."}
	NsEPlaylistEntrySeek                                         = &Error{0xC00D14B9, "NS_E_PLAYLIST_ENTRY_SEEK", "Cannot seek to the desired playlist entry."}
	NsEPlaylistRecursivePlaylists                                = &Error{0xC00D14BA, "NS_E_PLAYLIST_RECURSIVE_PLAYLISTS", "Cannot play recursive playlist."}
	NsEPlaylistTooManyNestedPlaylists                            = &Error{0xC00D14BB, "NS_E_PLAYLIST_TOO_MANY_NESTED_PLAYLISTS", "The number of nested playlists exceeded the limit the server can handle."}
	NsEPlaylistShutdown                                          = &Error{0xC00D14BC, "NS_E_PLAYLIST_SHUTDOWN", "Cannot execute the requested operation because the playlist has been shut down by the Media Server."}
	NsEPlaylistEndReceding                                       = &Error{0xC00D14BD, "NS_E_PLAYLIST_END_RECEDING", "The playlist has ended while receding."}
	NsEDatapathNoSink                                            = &Error{0xC00D1518, "NS_E_DATAPATH_NO_SINK", "The data path does not have an associated data writer plug-in."}
	NsEInvalidPushTemplate                                       = &Error{0xC00D151A, "NS_E_INVALID_PUSH_TEMPLATE", "The specified push template is invalid."}
	NsEInvalidPushPublishingPoint                                = &Error{0xC00D151B, "NS_E_INVALID_PUSH_PUBLISHING_POINT", "The specified push publishing point is invalid."}
	NsECriticalError                                             = &Error{0xC00D151C, "NS_E_CRITICAL_ERROR", "The requested operation cannot be performed because the server or publishing point is in a critical error state."}
	NsENoNewConnections                                          = &Error{0xC00D151D, "NS_E_NO_NEW_CONNECTIONS", "The content cannot be played because the server is not currently accepting connections. Try connecting at a later time."}
	NsEWsxInvalidVersion                                         = &Error{0xC00D151E, "NS_E_WSX_INVALID_VERSION", "The version of this playlist is not supported by the server."}
	NsEHeaderMismatch                                            = &Error{0xC00D151F, "NS_E_HEADER_MISMATCH", "The command does not apply to the current media header user by a server component."}
	NsEPushDuplicatePublishingPointName                          = &Error{0xC00D1520, "NS_E_PUSH_DUPLICATE_PUBLISHING_POINT_NAME", "The specified publishing point name is already in use."}
	NsENoScriptEngine                                            = &Error{0xC00D157C, "NS_E_NO_SCRIPT_ENGINE", "There is no script engine available for this file."}
	NsEPluginErrorReported                                       = &Error{0xC00D157D, "NS_E_PLUGIN_ERROR_REPORTED", "The plug-in has reported an error. See the Troubleshooting tab or the NT Application Event Log for details."}
	NsESourcePluginNotFound                                      = &Error{0xC00D157E, "NS_E_SOURCE_PLUGIN_NOT_FOUND", "No enabled data source plug-in is available to access the requested content."}
	NsEPlaylistPluginNotFound                                    = &Error{0xC00D157F, "NS_E_PLAYLIST_PLUGIN_NOT_FOUND", "No enabled playlist parser plug-in is available to access the requested content."}
	NsEDataSourceEnumerationNotSupported                         = &Error{0xC00D1580, "NS_E_DATA_SOURCE_ENUMERATION_NOT_SUPPORTED", "The data source plug-in does not support enumeration."}
	NsEMediaParserInvalidFormat                                  = &Error{0xC00D1581, "NS_E_MEDIA_PARSER_INVALID_FORMAT", "The server cannot stream the selected file because it is either damaged or corrupt. Select a different file."}
	NsEScriptDebuggerNotInstalled                                = &Error{0xC00D1582, "NS_E_SCRIPT_DEBUGGER_NOT_INSTALLED", "The plug-in cannot be enabled because a compatible script debugger is not installed on this system. Install a script debugger, or disable the script debugger option on the general tab of the plug-in's properties page and try again."}
	NsEFeatureRequiresEnterpriseServer                           = &Error{0xC00D1583, "NS_E_FEATURE_REQUIRES_ENTERPRISE_SERVER", "The plug-in cannot be loaded because it requires Windows Server 2003, Enterprise Edition."}
	NsEWizardRunning                                             = &Error{0xC00D1584, "NS_E_WIZARD_RUNNING", "Another wizard is currently running. Please close the other wizard or wait until it finishes before attempting to run this wizard again."}
	NsEInvalidLogUrl                                             = &Error{0xC00D1585, "NS_E_INVALID_LOG_URL", "Invalid log URL. Multicast logging URL must look like \"http://servername/isapibackend.dll\"."}
	NsEInvalidMtuRange                                           = &Error{0xC00D1586, "NS_E_INVALID_MTU_RANGE", "Invalid MTU specified. The valid range for maximum packet size is between 36 and 65507 bytes."}
	NsEInvalidPlayStatistics                                     = &Error{0xC00D1587, "NS_E_INVALID_PLAY_STATISTICS", "Invalid play statistics for logging."}
	NsELogNeedToBeSkipped                                        = &Error{0xC00D1588, "NS_E_LOG_NEED_TO_BE_SKIPPED", "The log needs to be skipped."}
	NsEHttpTextDatacontainerSizeLimitExceeded                    = &Error{0xC00D1589, "NS_E_HTTP_TEXT_DATACONTAINER_SIZE_LIMIT_EXCEEDED", "The size of the data exceeded the limit the WMS HTTP Download Data Source plugin can handle."}
	NsEPortInUse                                                 = &Error{0xC00D158A, "NS_E_PORT_IN_USE", "One usage of each socket address (protocol/network address/port) is permitted. Verify that other services or applications are not attempting to use the same port and then try to enable the plug-in again."}
	NsEPortInUseHttp                                             = &Error{0xC00D158B, "NS_E_PORT_IN_USE_HTTP", "One usage of each socket address (protocol/network address/port) is permitted. Verify that other services (such as IIS) or applications are not attempting to use the same port and then try to enable the plug-in again."}
	NsEHttpTextDatacontainerInvalidServerResponse                = &Error{0xC00D158C, "NS_E_HTTP_TEXT_DATACONTAINER_INVALID_SERVER_RESPONSE", "The WMS HTTP Download Data Source plugin was unable to receive the remote server's response."}
	NsEArchiveReachQuota                                         = &Error{0xC00D158D, "NS_E_ARCHIVE_REACH_QUOTA", "The archive plug-in has reached its quota."}
	NsEArchiveAbortDueToBcast                                    = &Error{0xC00D158E, "NS_E_ARCHIVE_ABORT_DUE_TO_BCAST", "The archive plug-in aborted because the source was from broadcast."}
	NsEArchiveGapDetected                                        = &Error{0xC00D158F, "NS_E_ARCHIVE_GAP_DETECTED", "The archive plug-in detected an interrupt in the source."}
	NsEAuthorizationFileNotFound                                 = &Error{0xC00D1590, "NS_E_AUTHORIZATION_FILE_NOT_FOUND", "The system cannot find the file specified."}
	NsEBadMarkin                                                 = &Error{0xC00D1B58, "NS_E_BAD_MARKIN", "The mark-in time should be greater than 0 and less than the mark-out time."}
	NsEBadMarkout                                                = &Error{0xC00D1B59, "NS_E_BAD_MARKOUT", "The mark-out time should be greater than the mark-in time and less than the file duration."}
	NsENomatchingMediasource                                     = &Error{0xC00D1B5A, "NS_E_NOMATCHING_MEDIASOURCE", "No matching media type is found in the source %1."}
	NsEUnsupportedSourcetype                                     = &Error{0xC00D1B5B, "NS_E_UNSUPPORTED_SOURCETYPE", "The specified source type is not supported."}
	NsETooManyAudio                                              = &Error{0xC00D1B5C, "NS_E_TOO_MANY_AUDIO", "It is not possible to specify more than one audio input."}
	NsETooManyVideo                                              = &Error{0xC00D1B5D, "NS_E_TOO_MANY_VIDEO", "It is not possible to specify more than two video inputs."}
	NsENomatchingElement                                         = &Error{0xC00D1B5E, "NS_E_NOMATCHING_ELEMENT", "No matching element is found in the list."}
	NsEMismatchedMediacontent                                    = &Error{0xC00D1B5F, "NS_E_MISMATCHED_MEDIACONTENT", "The profile's media types must match the media types defined for the session."}
	NsECannotDeleteActiveSourcegroup                             = &Error{0xC00D1B60, "NS_E_CANNOT_DELETE_ACTIVE_SOURCEGROUP", "It is not possible to remove an active source while encoding."}
	NsEAudiodeviceBusy                                           = &Error{0xC00D1B61, "NS_E_AUDIODEVICE_BUSY", "It is not possible to open the specified audio capture device because it is currently in use."}
	NsEAudiodeviceUnexpected                                     = &Error{0xC00D1B62, "NS_E_AUDIODEVICE_UNEXPECTED", "It is not possible to open the specified audio capture device because an unexpected error has occurred."}
	NsEAudiodeviceBadformat                                      = &Error{0xC00D1B63, "NS_E_AUDIODEVICE_BADFORMAT", "The audio capture device does not support the specified audio format."}
	NsEVideodeviceBusy                                           = &Error{0xC00D1B64, "NS_E_VIDEODEVICE_BUSY", "It is not possible to open the specified video capture device because it is currently in use."}
	NsEVideodeviceUnexpected                                     = &Error{0xC00D1B65, "NS_E_VIDEODEVICE_UNEXPECTED", "It is not possible to open the specified video capture device because an unexpected error has occurred."}
	NsEInvalidcallWhileEncoderRunning                            = &Error{0xC00D1B66, "NS_E_INVALIDCALL_WHILE_ENCODER_RUNNING", "This operation is not allowed while encoding."}
	NsENoProfileInSourcegroup                                    = &Error{0xC00D1B67, "NS_E_NO_PROFILE_IN_SOURCEGROUP", "No profile is set for the source."}
	NsEVideodriverUnstable                                       = &Error{0xC00D1B68, "NS_E_VIDEODRIVER_UNSTABLE", "The video capture driver returned an unrecoverable error. It is now in an unstable state."}
	NsEVidcapstartfailed                                         = &Error{0xC00D1B69, "NS_E_VIDCAPSTARTFAILED", "It was not possible to start the video device."}
	NsEVidsourcecompression                                      = &Error{0xC00D1B6A, "NS_E_VIDSOURCECOMPRESSION", "The video source does not support the requested output format or color depth."}
	NsEVidsourcesize                                             = &Error{0xC00D1B6B, "NS_E_VIDSOURCESIZE", "The video source does not support the requested capture size."}
	NsEIcmqueryformat                                            = &Error{0xC00D1B6C, "NS_E_ICMQUERYFORMAT", "It was not possible to obtain output information from the video compressor."}
	NsEVidcapcreatewindow                                        = &Error{0xC00D1B6D, "NS_E_VIDCAPCREATEWINDOW", "It was not possible to create a video capture window."}
	NsEVidcapdrvinuse                                            = &Error{0xC00D1B6E, "NS_E_VIDCAPDRVINUSE", "There is already a stream active on this video device."}
	NsENoMediaformatInSource                                     = &Error{0xC00D1B6F, "NS_E_NO_MEDIAFORMAT_IN_SOURCE", "No media format is set in source."}
	NsENoValidOutputStream                                       = &Error{0xC00D1B70, "NS_E_NO_VALID_OUTPUT_STREAM", "Cannot find a valid output stream from the source."}
	NsENoValidSourcePlugin                                       = &Error{0xC00D1B71, "NS_E_NO_VALID_SOURCE_PLUGIN", "It was not possible to find a valid source plug-in for the specified source."}
	NsENoActiveSourcegroup                                       = &Error{0xC00D1B72, "NS_E_NO_ACTIVE_SOURCEGROUP", "No source is currently active."}
	NsENoScriptStream                                            = &Error{0xC00D1B73, "NS_E_NO_SCRIPT_STREAM", "No script stream is set in the current source."}
	NsEInvalidcallWhileArchivalRunning                           = &Error{0xC00D1B74, "NS_E_INVALIDCALL_WHILE_ARCHIVAL_RUNNING", "This operation is not allowed while archiving."}
	NsEInvalidpacketsize                                         = &Error{0xC00D1B75, "NS_E_INVALIDPACKETSIZE", "The setting for the maximum packet size is not valid."}
	NsEPluginClsidInvalid                                        = &Error{0xC00D1B76, "NS_E_PLUGIN_CLSID_INVALID", "The plug-in CLSID specified is not valid."}
	NsEUnsupportedArchivetype                                    = &Error{0xC00D1B77, "NS_E_UNSUPPORTED_ARCHIVETYPE", "This archive type is not supported."}
	NsEUnsupportedArchiveoperation                               = &Error{0xC00D1B78, "NS_E_UNSUPPORTED_ARCHIVEOPERATION", "This archive operation is not supported."}
	NsEArchiveFilenameNotset                                     = &Error{0xC00D1B79, "NS_E_ARCHIVE_FILENAME_NOTSET", "The local archive file name was not set."}
	NsESourcegroupNotprepared                                    = &Error{0xC00D1B7A, "NS_E_SOURCEGROUP_NOTPREPARED", "The source is not yet prepared."}
	NsEProfileMismatch                                           = &Error{0xC00D1B7B, "NS_E_PROFILE_MISMATCH", "Profiles on the sources do not match."}
	NsEIncorrectclipsettings                                     = &Error{0xC00D1B7C, "NS_E_INCORRECTCLIPSETTINGS", "The specified crop values are not valid."}
	NsENostatsavailable                                          = &Error{0xC00D1B7D, "NS_E_NOSTATSAVAILABLE", "No statistics are available at this time."}
	NsENotarchiving                                              = &Error{0xC00D1B7E, "NS_E_NOTARCHIVING", "The encoder is not archiving."}
	NsEInvalidcallWhileEncoderStopped                            = &Error{0xC00D1B7F, "NS_E_INVALIDCALL_WHILE_ENCODER_STOPPED", "This operation is only allowed during encoding."}
	NsENosourcegroups                                            = &Error{0xC00D1B80, "NS_E_NOSOURCEGROUPS", "This SourceGroupCollection doesn't contain any SourceGroups."}
	NsEInvalidinputfps                                           = &Error{0xC00D1B81, "NS_E_INVALIDINPUTFPS", "This source does not have a frame rate of 30 fps. Therefore, it is not possible to apply the inverse telecine filter to the source."}
	NsENoDataviewSupport                                         = &Error{0xC00D1B82, "NS_E_NO_DATAVIEW_SUPPORT", "It is not possible to display your source or output video in the Video panel."}
	NsECodecUnavailable                                          = &Error{0xC00D1B83, "NS_E_CODEC_UNAVAILABLE", "One or more codecs required to open this content could not be found."}
	NsEArchiveSameAsInput                                        = &Error{0xC00D1B84, "NS_E_ARCHIVE_SAME_AS_INPUT", "The archive file has the same name as an input file. Change one of the names before continuing."}
	NsESourceNotspecified                                        = &Error{0xC00D1B85, "NS_E_SOURCE_NOTSPECIFIED", "The source has not been set up completely."}
	NsENoRealtimeTimecompression                                 = &Error{0xC00D1B86, "NS_E_NO_REALTIME_TIMECOMPRESSION", "It is not possible to apply time compression to a broadcast session."}
	NsEUnsupportedEncoderDevice                                  = &Error{0xC00D1B87, "NS_E_UNSUPPORTED_ENCODER_DEVICE", "It is not possible to open this device."}
	NsEUnexpectedDisplaySettings                                 = &Error{0xC00D1B88, "NS_E_UNEXPECTED_DISPLAY_SETTINGS", "It is not possible to start encoding because the display size or color has changed since the current session was defined. Restore the previous settings or create a new session."}
	NsENoAudiodata                                               = &Error{0xC00D1B89, "NS_E_NO_AUDIODATA", "No audio data has been received for several seconds. Check the audio source and restart the encoder."}
	NsEInputsourceProblem                                        = &Error{0xC00D1B8A, "NS_E_INPUTSOURCE_PROBLEM", "One or all of the specified sources are not working properly. Check that the sources are configured correctly."}
	NsEWmeVersionMismatch                                        = &Error{0xC00D1B8B, "NS_E_WME_VERSION_MISMATCH", "The supplied configuration file is not supported by this version of the encoder."}
	NsENoRealtimePreprocess                                      = &Error{0xC00D1B8C, "NS_E_NO_REALTIME_PREPROCESS", "It is not possible to use image preprocessing with live encoding."}
	NsENoRepeatPreprocess                                        = &Error{0xC00D1B8D, "NS_E_NO_REPEAT_PREPROCESS", "It is not possible to use two-pass encoding when the source is set to loop."}
	NsECannotPauseLivebroadcast                                  = &Error{0xC00D1B8E, "NS_E_CANNOT_PAUSE_LIVEBROADCAST", "It is not possible to pause encoding during a broadcast."}
	NsEDrmProfileNotSet                                          = &Error{0xC00D1B8F, "NS_E_DRM_PROFILE_NOT_SET", "A DRM profile has not been set for the current session."}
	NsEDuplicateDrmprofile                                       = &Error{0xC00D1B90, "NS_E_DUPLICATE_DRMPROFILE", "The profile ID is already used by a DRM profile. Specify a different profile ID."}
	NsEInvalidDevice                                             = &Error{0xC00D1B91, "NS_E_INVALID_DEVICE", "The setting of the selected device does not support control for playing back tapes."}
	NsESpeechedlOnNonMixedmode                                   = &Error{0xC00D1B92, "NS_E_SPEECHEDL_ON_NON_MIXEDMODE", "You must specify a mixed voice and audio mode in order to use an optimization definition file."}
	NsEDrmPasswordTooLong                                        = &Error{0xC00D1B93, "NS_E_DRM_PASSWORD_TOO_LONG", "The specified password is too long. Type a password with fewer than 8 characters."}
	NsEDevcontrolFailedSeek                                      = &Error{0xC00D1B94, "NS_E_DEVCONTROL_FAILED_SEEK", "It is not possible to seek to the specified mark-in point."}
	NsEInterlaceRequireSamesize                                  = &Error{0xC00D1B95, "NS_E_INTERLACE_REQUIRE_SAMESIZE", "When you choose to maintain the interlacing in your video, the output video size must match the input video size."}
	NsETooManyDevicecontrol                                      = &Error{0xC00D1B96, "NS_E_TOO_MANY_DEVICECONTROL", "Only one device control plug-in can control a device."}
	NsENoMultipassForLivedevice                                  = &Error{0xC00D1B97, "NS_E_NO_MULTIPASS_FOR_LIVEDEVICE", "You must also enable storing content to hard disk temporarily in order to use two-pass encoding with the input device."}
	NsEMissingAudience                                           = &Error{0xC00D1B98, "NS_E_MISSING_AUDIENCE", "An audience is missing from the output stream configuration."}
	NsEAudienceContenttypeMismatch                               = &Error{0xC00D1B99, "NS_E_AUDIENCE_CONTENTTYPE_MISMATCH", "All audiences in the output tree must have the same content type."}
	NsEMissingSourceIndex                                        = &Error{0xC00D1B9A, "NS_E_MISSING_SOURCE_INDEX", "A source index is missing from the output stream configuration."}
	NsENumLanguageMismatch                                       = &Error{0xC00D1B9B, "NS_E_NUM_LANGUAGE_MISMATCH", "The same source index in different audiences should have the same number of languages."}
	NsELanguageMismatch                                          = &Error{0xC00D1B9C, "NS_E_LANGUAGE_MISMATCH", "The same source index in different audiences should have the same languages."}
	NsEVbrmodeMismatch                                           = &Error{0xC00D1B9D, "NS_E_VBRMODE_MISMATCH", "The same source index in different audiences should use the same VBR encoding mode."}
	NsEInvalidInputAudienceIndex                                 = &Error{0xC00D1B9E, "NS_E_INVALID_INPUT_AUDIENCE_INDEX", "The bit rate index specified is not valid."}
	NsEInvalidInputLanguage                                      = &Error{0xC00D1B9F, "NS_E_INVALID_INPUT_LANGUAGE", "The specified language is not valid."}
	NsEInvalidInputStream                                        = &Error{0xC00D1BA0, "NS_E_INVALID_INPUT_STREAM", "The specified source type is not valid."}
	NsEExpectMonoWavInput                                        = &Error{0xC00D1BA1, "NS_E_EXPECT_MONO_WAV_INPUT", "The source must be a mono channel .wav file."}
	NsEInputWavformatMismatch                                    = &Error{0xC00D1BA2, "NS_E_INPUT_WAVFORMAT_MISMATCH", "All the source .wav files must have the same format."}
	NsERecordqDiskFull                                           = &Error{0xC00D1BA3, "NS_E_RECORDQ_DISK_FULL", "The hard disk being used for temporary storage of content has reached the minimum allowed disk space. Create more space on the hard disk and restart encoding."}
	NsENoPalInverseTelecine                                      = &Error{0xC00D1BA4, "NS_E_NO_PAL_INVERSE_TELECINE", "It is not possible to apply the inverse telecine feature to PAL content."}
	NsEActiveSgDeviceDisconnected                                = &Error{0xC00D1BA5, "NS_E_ACTIVE_SG_DEVICE_DISCONNECTED", "A capture device in the current active source is no longer available."}
	NsEActiveSgDeviceControlDisconnected                         = &Error{0xC00D1BA6, "NS_E_ACTIVE_SG_DEVICE_CONTROL_DISCONNECTED", "A device used in the current active source for device control is no longer available."}
	NsENoFramesSubmittedToAnalyzer                               = &Error{0xC00D1BA7, "NS_E_NO_FRAMES_SUBMITTED_TO_ANALYZER", "No frames have been submitted to the analyzer for analysis."}
	NsEInputDoesnotSupportSmpte                                  = &Error{0xC00D1BA8, "NS_E_INPUT_DOESNOT_SUPPORT_SMPTE", "The source video does not support time codes."}
	NsENoSmpteWithMultipleSourcegroups                           = &Error{0xC00D1BA9, "NS_E_NO_SMPTE_WITH_MULTIPLE_SOURCEGROUPS", "It is not possible to generate a time code when there are multiple sources in a session."}
	NsEBadContentedl                                             = &Error{0xC00D1BAA, "NS_E_BAD_CONTENTEDL", "The voice codec optimization definition file cannot be found or is corrupted."}
	NsEInterlacemodeMismatch                                     = &Error{0xC00D1BAB, "NS_E_INTERLACEMODE_MISMATCH", "The same source index in different audiences should have the same interlace mode."}
	NsENonsquarepixelmodeMismatch                                = &Error{0xC00D1BAC, "NS_E_NONSQUAREPIXELMODE_MISMATCH", "The same source index in different audiences should have the same nonsquare pixel mode."}
	NsESmptemodeMismatch                                         = &Error{0xC00D1BAD, "NS_E_SMPTEMODE_MISMATCH", "The same source index in different audiences should have the same time code mode."}
	NsEEndOfTape                                                 = &Error{0xC00D1BAE, "NS_E_END_OF_TAPE", "Either the end of the tape has been reached or there is no tape. Check the device and tape."}
	NsENoMediaInAudience                                         = &Error{0xC00D1BAF, "NS_E_NO_MEDIA_IN_AUDIENCE", "No audio or video input has been specified."}
	NsENoAudiences                                               = &Error{0xC00D1BB0, "NS_E_NO_AUDIENCES", "The profile must contain a bit rate."}
	NsENoAudioCompat                                             = &Error{0xC00D1BB1, "NS_E_NO_AUDIO_COMPAT", "You must specify at least one audio stream to be compatible with Windows Media Player 7.1."}
	NsEInvalidVbrCompat                                          = &Error{0xC00D1BB2, "NS_E_INVALID_VBR_COMPAT", "Using a VBR encoding mode is not compatible with Windows Media Player 7.1."}
	NsENoProfileName                                             = &Error{0xC00D1BB3, "NS_E_NO_PROFILE_NAME", "You must specify a profile name."}
	NsEInvalidVbrWithUncomp                                      = &Error{0xC00D1BB4, "NS_E_INVALID_VBR_WITH_UNCOMP", "It is not possible to use a VBR encoding mode with uncompressed audio or video."}
	NsEMultipleVbrAudiences                                      = &Error{0xC00D1BB5, "NS_E_MULTIPLE_VBR_AUDIENCES", "It is not possible to use MBR encoding with VBR encoding."}
	NsEUncompCompCombination                                     = &Error{0xC00D1BB6, "NS_E_UNCOMP_COMP_COMBINATION", "It is not possible to mix uncompressed and compressed content in a session."}
	NsEMultipleAudioCodecs                                       = &Error{0xC00D1BB7, "NS_E_MULTIPLE_AUDIO_CODECS", "All audiences must use the same audio codec."}
	NsEMultipleAudioFormats                                      = &Error{0xC00D1BB8, "NS_E_MULTIPLE_AUDIO_FORMATS", "All audiences should use the same audio format to be compatible with Windows Media Player 7.1."}
	NsEAudioBitrateStepdown                                      = &Error{0xC00D1BB9, "NS_E_AUDIO_BITRATE_STEPDOWN", "The audio bit rate for an audience with a higher total bit rate must be greater than one with a lower total bit rate."}
	NsEInvalidAudioPeakrate                                      = &Error{0xC00D1BBA, "NS_E_INVALID_AUDIO_PEAKRATE", "The audio peak bit rate setting is not valid."}
	NsEInvalidAudioPeakrate2                                     = &Error{0xC00D1BBB, "NS_E_INVALID_AUDIO_PEAKRATE_2", "The audio peak bit rate setting must be greater than the audio bit rate setting."}
	NsEInvalidAudioBuffermax                                     = &Error{0xC00D1BBC, "NS_E_INVALID_AUDIO_BUFFERMAX", "The setting for the maximum buffer size for audio is not valid."}
	NsEMultipleVideoCodecs                                       = &Error{0xC00D1BBD, "NS_E_MULTIPLE_VIDEO_CODECS", "All audiences must use the same video codec."}
	NsEMultipleVideoSizes                                        = &Error{0xC00D1BBE, "NS_E_MULTIPLE_VIDEO_SIZES", "All audiences should use the same video size to be compatible with Windows Media Player 7.1."}
	NsEInvalidVideoBitrate                                       = &Error{0xC00D1BBF, "NS_E_INVALID_VIDEO_BITRATE", "The video bit rate setting is not valid."}
	NsEVideoBitrateStepdown                                      = &Error{0xC00D1BC0, "NS_E_VIDEO_BITRATE_STEPDOWN", "The video bit rate for an audience with a higher total bit rate must be greater than one with a lower total bit rate."}
	NsEInvalidVideoPeakrate                                      = &Error{0xC00D1BC1, "NS_E_INVALID_VIDEO_PEAKRATE", "The video peak bit rate setting is not valid."}
	NsEInvalidVideoPeakrate2                                     = &Error{0xC00D1BC2, "NS_E_INVALID_VIDEO_PEAKRATE_2", "The video peak bit rate setting must be greater than the video bit rate setting."}
	NsEInvalidVideoWidth                                         = &Error{0xC00D1BC3, "NS_E_INVALID_VIDEO_WIDTH", "The video width setting is not valid."}
	NsEInvalidVideoHeight                                        = &Error{0xC00D1BC4, "NS_E_INVALID_VIDEO_HEIGHT", "The video height setting is not valid."}
	NsEInvalidVideoFps                                           = &Error{0xC00D1BC5, "NS_E_INVALID_VIDEO_FPS", "The video frame rate setting is not valid."}
	NsEInvalidVideoKeyframe                                      = &Error{0xC00D1BC6, "NS_E_INVALID_VIDEO_KEYFRAME", "The video key frame setting is not valid."}
	NsEInvalidVideoIquality                                      = &Error{0xC00D1BC7, "NS_E_INVALID_VIDEO_IQUALITY", "The video image quality setting is not valid."}
	NsEInvalidVideoCquality                                      = &Error{0xC00D1BC8, "NS_E_INVALID_VIDEO_CQUALITY", "The video codec quality setting is not valid."}
	NsEInvalidVideoBuffer                                        = &Error{0xC00D1BC9, "NS_E_INVALID_VIDEO_BUFFER", "The video buffer setting is not valid."}
	NsEInvalidVideoBuffermax                                     = &Error{0xC00D1BCA, "NS_E_INVALID_VIDEO_BUFFERMAX", "The setting for the maximum buffer size for video is not valid."}
	NsEInvalidVideoBuffermax2                                    = &Error{0xC00D1BCB, "NS_E_INVALID_VIDEO_BUFFERMAX_2", "The value of the video maximum buffer size setting must be greater than the video buffer size setting."}
	NsEInvalidVideoWidthAlign                                    = &Error{0xC00D1BCC, "NS_E_INVALID_VIDEO_WIDTH_ALIGN", "The alignment of the video width is not valid."}
	NsEInvalidVideoHeightAlign                                   = &Error{0xC00D1BCD, "NS_E_INVALID_VIDEO_HEIGHT_ALIGN", "The alignment of the video height is not valid."}
	NsEMultipleScriptBitrates                                    = &Error{0xC00D1BCE, "NS_E_MULTIPLE_SCRIPT_BITRATES", "All bit rates must have the same script bit rate."}
	NsEInvalidScriptBitrate                                      = &Error{0xC00D1BCF, "NS_E_INVALID_SCRIPT_BITRATE", "The script bit rate specified is not valid."}
	NsEMultipleFileBitrates                                      = &Error{0xC00D1BD0, "NS_E_MULTIPLE_FILE_BITRATES", "All bit rates must have the same file transfer bit rate."}
	NsEInvalidFileBitrate                                        = &Error{0xC00D1BD1, "NS_E_INVALID_FILE_BITRATE", "The file transfer bit rate is not valid."}
	NsESameAsInputCombination                                    = &Error{0xC00D1BD2, "NS_E_SAME_AS_INPUT_COMBINATION", "All audiences in a profile should either be same as input or have video width and height specified."}
	NsESourceCannotLoop                                          = &Error{0xC00D1BD3, "NS_E_SOURCE_CANNOT_LOOP", "This source type does not support looping."}
	NsEInvalidFolddownCoefficients                               = &Error{0xC00D1BD4, "NS_E_INVALID_FOLDDOWN_COEFFICIENTS", "The fold-down value needs to be between -144 and 0."}
	NsEDrmprofileNotfound                                        = &Error{0xC00D1BD5, "NS_E_DRMPROFILE_NOTFOUND", "The specified DRM profile does not exist in the system."}
	NsEInvalidTimecode                                           = &Error{0xC00D1BD6, "NS_E_INVALID_TIMECODE", "The specified time code is not valid."}
	NsENoAudioTimecompression                                    = &Error{0xC00D1BD7, "NS_E_NO_AUDIO_TIMECOMPRESSION", "It is not possible to apply time compression to a video-only session."}
	NsENoTwopassTimecompression                                  = &Error{0xC00D1BD8, "NS_E_NO_TWOPASS_TIMECOMPRESSION", "It is not possible to apply time compression to a session that is using two-pass encoding."}
	NsETimecodeRequiresVideostream                               = &Error{0xC00D1BD9, "NS_E_TIMECODE_REQUIRES_VIDEOSTREAM", "It is not possible to generate a time code for an audio-only session."}
	NsENoMbrWithTimecode                                         = &Error{0xC00D1BDA, "NS_E_NO_MBR_WITH_TIMECODE", "It is not possible to generate a time code when you are encoding content at multiple bit rates."}
	NsEInvalidInterlacemode                                      = &Error{0xC00D1BDB, "NS_E_INVALID_INTERLACEMODE", "The video codec selected does not support maintaining interlacing in video."}
	NsEInvalidInterlaceCompat                                    = &Error{0xC00D1BDC, "NS_E_INVALID_INTERLACE_COMPAT", "Maintaining interlacing in video is not compatible with Windows Media Player 7.1."}
	NsEInvalidNonsquarepixelCompat                               = &Error{0xC00D1BDD, "NS_E_INVALID_NONSQUAREPIXEL_COMPAT", "Allowing nonsquare pixel output is not compatible with Windows Media Player 7.1."}
	NsEInvalidSourceWithDeviceControl                            = &Error{0xC00D1BDE, "NS_E_INVALID_SOURCE_WITH_DEVICE_CONTROL", "Only capture devices can be used with device control."}
	NsECannotGenerateBroadcastInfoForQualityvbr                  = &Error{0xC00D1BDF, "NS_E_CANNOT_GENERATE_BROADCAST_INFO_FOR_QUALITYVBR", "It is not possible to generate the stream format file if you are using quality-based VBR encoding for the audio or video stream. Instead use the Windows Media file generated after encoding to create the announcement file."}
	NsEExceedMaxDrmProfileLimit                                  = &Error{0xC00D1BE0, "NS_E_EXCEED_MAX_DRM_PROFILE_LIMIT", "It is not possible to create a DRM profile because the maximum number of profiles has been reached. You must delete some DRM profiles before creating new ones."}
	NsEDevicecontrolUnstable                                     = &Error{0xC00D1BE1, "NS_E_DEVICECONTROL_UNSTABLE", "The device is in an unstable state. Check that the device is functioning properly and a tape is in place."}
	NsEInvalidPixelAspectRatio                                   = &Error{0xC00D1BE2, "NS_E_INVALID_PIXEL_ASPECT_RATIO", "The pixel aspect ratio value must be between 1 and 255."}
	NsEAudienceLanguageContenttypeMismatch                       = &Error{0xC00D1BE3, "NS_E_AUDIENCE__LANGUAGE_CONTENTTYPE_MISMATCH", "All streams with different languages in the same audience must have same properties."}
	NsEInvalidProfileContenttype                                 = &Error{0xC00D1BE4, "NS_E_INVALID_PROFILE_CONTENTTYPE", "The profile must contain at least one audio or video stream."}
	NsETransformPluginNotFound                                   = &Error{0xC00D1BE5, "NS_E_TRANSFORM_PLUGIN_NOT_FOUND", "The transform plug-in could not be found."}
	NsETransformPluginInvalid                                    = &Error{0xC00D1BE6, "NS_E_TRANSFORM_PLUGIN_INVALID", "The transform plug-in is not valid. It might be damaged or you might not have the required permissions to access the plug-in."}
	NsEEdlRequiredForDeviceMultipass                             = &Error{0xC00D1BE7, "NS_E_EDL_REQUIRED_FOR_DEVICE_MULTIPASS", "To use two-pass encoding, you must enable device control and setup an edit decision list (EDL) that has at least one entry."}
	NsEInvalidVideoWidthForInterlacedEncoding                    = &Error{0xC00D1BE8, "NS_E_INVALID_VIDEO_WIDTH_FOR_INTERLACED_ENCODING", "When you choose to maintain the interlacing in your video, the output video size must be a multiple of 4."}
	NsEMarkinUnsupported                                         = &Error{0xC00D1BE9, "NS_E_MARKIN_UNSUPPORTED", "Markin/Markout is unsupported with this source type."}
	NsEDrmInvalidApplication                                     = &Error{0xC00D2711, "NS_E_DRM_INVALID_APPLICATION", "A problem has occurred in the Digital Rights Management component. Contact product support for this application."}
	NsEDrmLicenseStoreError                                      = &Error{0xC00D2712, "NS_E_DRM_LICENSE_STORE_ERROR", "License storage is not working. Contact Microsoft product support."}
	NsEDrmSecureStoreError                                       = &Error{0xC00D2713, "NS_E_DRM_SECURE_STORE_ERROR", "Secure storage is not working. Contact Microsoft product support."}
	NsEDrmLicenseStoreSaveError                                  = &Error{0xC00D2714, "NS_E_DRM_LICENSE_STORE_SAVE_ERROR", "License acquisition did not work. Acquire a new license or contact the content provider for further assistance."}
	NsEDrmSecureStoreUnlockError                                 = &Error{0xC00D2715, "NS_E_DRM_SECURE_STORE_UNLOCK_ERROR", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmInvalidContent                                         = &Error{0xC00D2716, "NS_E_DRM_INVALID_CONTENT", "The media file is corrupted. Contact the content provider to get a new file."}
	NsEDrmUnableToOpenLicense                                    = &Error{0xC00D2717, "NS_E_DRM_UNABLE_TO_OPEN_LICENSE", "The license is corrupted. Acquire a new license."}
	NsEDrmInvalidLicense                                         = &Error{0xC00D2718, "NS_E_DRM_INVALID_LICENSE", "The license is corrupted or invalid. Acquire a new license"}
	NsEDrmInvalidMachine                                         = &Error{0xC00D2719, "NS_E_DRM_INVALID_MACHINE", "Licenses cannot be copied from one computer to another. Use License Management to transfer licenses, or get a new license for the media file."}
	NsEDrmEnumLicenseFailed                                      = &Error{0xC00D271B, "NS_E_DRM_ENUM_LICENSE_FAILED", "License storage is not working. Contact Microsoft product support."}
	NsEDrmInvalidLicenseRequest                                  = &Error{0xC00D271C, "NS_E_DRM_INVALID_LICENSE_REQUEST", "The media file is corrupted. Contact the content provider to get a new file."}
	NsEDrmUnableToInitialize                                     = &Error{0xC00D271D, "NS_E_DRM_UNABLE_TO_INITIALIZE", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToAcquireLicense                                 = &Error{0xC00D271E, "NS_E_DRM_UNABLE_TO_ACQUIRE_LICENSE", "The license could not be acquired. Try again later."}
	NsEDrmInvalidLicenseAcquired                                 = &Error{0xC00D271F, "NS_E_DRM_INVALID_LICENSE_ACQUIRED", "License acquisition did not work. Acquire a new license or contact the content provider for further assistance."}
	NsEDrmNoRights                                               = &Error{0xC00D2720, "NS_E_DRM_NO_RIGHTS", "The requested operation cannot be performed on this file."}
	NsEDrmKeyError                                               = &Error{0xC00D2721, "NS_E_DRM_KEY_ERROR", "The requested action cannot be performed because a problem occurred with the Windows Media Digital Rights Management (DRM) components on your computer."}
	NsEDrmEncryptError                                           = &Error{0xC00D2722, "NS_E_DRM_ENCRYPT_ERROR", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmDecryptError                                           = &Error{0xC00D2723, "NS_E_DRM_DECRYPT_ERROR", "The media file is corrupted. Contact the content provider to get a new file."}
	NsEDrmLicenseInvalidXml                                      = &Error{0xC00D2725, "NS_E_DRM_LICENSE_INVALID_XML", "The license is corrupted. Acquire a new license."}
	NsEDrmNeedsIndividualization                                 = &Error{0xC00D2728, "NS_E_DRM_NEEDS_INDIVIDUALIZATION", "A security upgrade is required to perform the operation on this media file."}
	NsEDrmAlreadyIndividualized                                  = &Error{0xC00D2729, "NS_E_DRM_ALREADY_INDIVIDUALIZED", "You already have the latest security components. No upgrade is necessary at this time."}
	NsEDrmActionNotQueried                                       = &Error{0xC00D272A, "NS_E_DRM_ACTION_NOT_QUERIED", "The application cannot perform this action. Contact product support for this application."}
	NsEDrmAcquiringLicense                                       = &Error{0xC00D272B, "NS_E_DRM_ACQUIRING_LICENSE", "You cannot begin a new license acquisition process until the current one has been completed."}
	NsEDrmIndividualizing                                        = &Error{0xC00D272C, "NS_E_DRM_INDIVIDUALIZING", "You cannot begin a new security upgrade until the current one has been completed."}
	NsEBackupRestoreFailure                                      = &Error{0xC00D272D, "NS_E_BACKUP_RESTORE_FAILURE", "Failure in Backup-Restore."}
	NsEBackupRestoreBadRequestId                                 = &Error{0xC00D272E, "NS_E_BACKUP_RESTORE_BAD_REQUEST_ID", "Bad Request ID in Backup-Restore."}
	NsEDrmParametersMismatched                                   = &Error{0xC00D272F, "NS_E_DRM_PARAMETERS_MISMATCHED", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateLicenseObject                            = &Error{0xC00D2730, "NS_E_DRM_UNABLE_TO_CREATE_LICENSE_OBJECT", "A license cannot be created for this media file. Reinstall the application."}
	NsEDrmUnableToCreateIndiObject                               = &Error{0xC00D2731, "NS_E_DRM_UNABLE_TO_CREATE_INDI_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateEncryptObject                            = &Error{0xC00D2732, "NS_E_DRM_UNABLE_TO_CREATE_ENCRYPT_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateDecryptObject                            = &Error{0xC00D2733, "NS_E_DRM_UNABLE_TO_CREATE_DECRYPT_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreatePropertiesObject                         = &Error{0xC00D2734, "NS_E_DRM_UNABLE_TO_CREATE_PROPERTIES_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateBackupObject                             = &Error{0xC00D2735, "NS_E_DRM_UNABLE_TO_CREATE_BACKUP_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmIndividualizeError                                     = &Error{0xC00D2736, "NS_E_DRM_INDIVIDUALIZE_ERROR", "The security upgrade failed. Try again later."}
	NsEDrmLicenseOpenError                                       = &Error{0xC00D2737, "NS_E_DRM_LICENSE_OPEN_ERROR", "License storage is not working. Contact Microsoft product support."}
	NsEDrmLicenseCloseError                                      = &Error{0xC00D2738, "NS_E_DRM_LICENSE_CLOSE_ERROR", "License storage is not working. Contact Microsoft product support."}
	NsEDrmGetLicenseError                                        = &Error{0xC00D2739, "NS_E_DRM_GET_LICENSE_ERROR", "License storage is not working. Contact Microsoft product support."}
	NsEDrmQueryError                                             = &Error{0xC00D273A, "NS_E_DRM_QUERY_ERROR", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmReportError                                            = &Error{0xC00D273B, "NS_E_DRM_REPORT_ERROR", "A problem has occurred in the Digital Rights Management component. Contact product support for this application."}
	NsEDrmGetLicensestringError                                  = &Error{0xC00D273C, "NS_E_DRM_GET_LICENSESTRING_ERROR", "License storage is not working. Contact Microsoft product support."}
	NsEDrmGetContentstringError                                  = &Error{0xC00D273D, "NS_E_DRM_GET_CONTENTSTRING_ERROR", "The media file is corrupted. Contact the content provider to get a new file."}
	NsEDrmMonitorError                                           = &Error{0xC00D273E, "NS_E_DRM_MONITOR_ERROR", "A problem has occurred in the Digital Rights Management component. Try again later."}
	NsEDrmUnableToSetParameter                                   = &Error{0xC00D273F, "NS_E_DRM_UNABLE_TO_SET_PARAMETER", "The application has made an invalid call to the Digital Rights Management component. Contact product support for this application."}
	NsEDrmInvalidAppdata                                         = &Error{0xC00D2740, "NS_E_DRM_INVALID_APPDATA", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmInvalidAppdataVersion                                  = &Error{0xC00D2741, "NS_E_DRM_INVALID_APPDATA_VERSION", "A problem has occurred in the Digital Rights Management component. Contact product support for this application."}
	NsEDrmBackupExists                                           = &Error{0xC00D2742, "NS_E_DRM_BACKUP_EXISTS", "Licenses are already backed up in this location."}
	NsEDrmBackupCorrupt                                          = &Error{0xC00D2743, "NS_E_DRM_BACKUP_CORRUPT", "One or more backed-up licenses are missing or corrupt."}
	NsEDrmBackuprestoreBusy                                      = &Error{0xC00D2744, "NS_E_DRM_BACKUPRESTORE_BUSY", "You cannot begin a new backup process until the current process has been completed."}
	NsEBackupRestoreBadData                                      = &Error{0xC00D2745, "NS_E_BACKUP_RESTORE_BAD_DATA", "Bad Data sent to Backup-Restore."}
	NsEDrmLicenseUnusable                                        = &Error{0xC00D2748, "NS_E_DRM_LICENSE_UNUSABLE", "The license is invalid. Contact the content provider for further assistance."}
	NsEDrmInvalidProperty                                        = &Error{0xC00D2749, "NS_E_DRM_INVALID_PROPERTY", "A required property was not set by the application. Contact product support for this application."}
	NsEDrmSecureStoreNotFound                                    = &Error{0xC00D274A, "NS_E_DRM_SECURE_STORE_NOT_FOUND", "A problem has occurred in the Digital Rights Management component of this application. Try to acquire a license again."}
	NsEDrmCachedContentError                                     = &Error{0xC00D274B, "NS_E_DRM_CACHED_CONTENT_ERROR", "A license cannot be found for this media file. Use License Management to transfer a license for this file from the original computer, or acquire a new license."}
	NsEDrmIndividualizationIncomplete                            = &Error{0xC00D274C, "NS_E_DRM_INDIVIDUALIZATION_INCOMPLETE", "A problem occurred during the security upgrade. Try again later."}
	NsEDrmDriverAuthFailure                                      = &Error{0xC00D274D, "NS_E_DRM_DRIVER_AUTH_FAILURE", "Certified driver components are required to play this media file. Contact Windows Update to see whether updated drivers are available for your hardware."}
	NsEDrmNeedUpgradeMssap                                       = &Error{0xC00D274E, "NS_E_DRM_NEED_UPGRADE_MSSAP", "One or more of the Secure Audio Path components were not found or an entry point in those components was not found."}
	NsEDrmReopenContent                                          = &Error{0xC00D274F, "NS_E_DRM_REOPEN_CONTENT", "Status message: Reopen the file."}
	NsEDrmDriverDigioutFailure                                   = &Error{0xC00D2750, "NS_E_DRM_DRIVER_DIGIOUT_FAILURE", "Certain driver functionality is required to play this media file. Contact Windows Update to see whether updated drivers are available for your hardware."}
	NsEDrmInvalidSecurestorePassword                             = &Error{0xC00D2751, "NS_E_DRM_INVALID_SECURESTORE_PASSWORD", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmAppcertRevoked                                         = &Error{0xC00D2752, "NS_E_DRM_APPCERT_REVOKED", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmRestoreFraud                                           = &Error{0xC00D2753, "NS_E_DRM_RESTORE_FRAUD", "You cannot restore your license(s)."}
	NsEDrmHardwareInconsistent                                   = &Error{0xC00D2754, "NS_E_DRM_HARDWARE_INCONSISTENT", "The licenses for your media files are corrupted. Contact Microsoft product support."}
	NsEDrmSdmiTrigger                                            = &Error{0xC00D2755, "NS_E_DRM_SDMI_TRIGGER", "To transfer this media file, you must upgrade the application."}
	NsEDrmSdmiNomorecopies                                       = &Error{0xC00D2756, "NS_E_DRM_SDMI_NOMORECOPIES", "You cannot make any more copies of this media file."}
	NsEDrmUnableToCreateHeaderObject                             = &Error{0xC00D2757, "NS_E_DRM_UNABLE_TO_CREATE_HEADER_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateKeysObject                               = &Error{0xC00D2758, "NS_E_DRM_UNABLE_TO_CREATE_KEYS_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmLicenseNotacquired                                     = &Error{0xC00D2759, "NS_E_DRM_LICENSE_NOTACQUIRED", "Unable to obtain license."}
	NsEDrmUnableToCreateCodingObject                             = &Error{0xC00D275A, "NS_E_DRM_UNABLE_TO_CREATE_CODING_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateStateDataObject                          = &Error{0xC00D275B, "NS_E_DRM_UNABLE_TO_CREATE_STATE_DATA_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmBufferTooSmall                                         = &Error{0xC00D275C, "NS_E_DRM_BUFFER_TOO_SMALL", "The buffer supplied is not sufficient."}
	NsEDrmUnsupportedProperty                                    = &Error{0xC00D275D, "NS_E_DRM_UNSUPPORTED_PROPERTY", "The property requested is not supported."}
	NsEDrmErrorBadNetResp                                        = &Error{0xC00D275E, "NS_E_DRM_ERROR_BAD_NET_RESP", "The specified server cannot perform the requested operation."}
	NsEDrmStoreNotallstored                                      = &Error{0xC00D275F, "NS_E_DRM_STORE_NOTALLSTORED", "Some of the licenses could not be stored."}
	NsEDrmSecurityComponentSignatureInvalid                      = &Error{0xC00D2760, "NS_E_DRM_SECURITY_COMPONENT_SIGNATURE_INVALID", "The Digital Rights Management security upgrade component could not be validated. Contact Microsoft product support."}
	NsEDrmInvalidData                                            = &Error{0xC00D2761, "NS_E_DRM_INVALID_DATA", "Invalid or corrupt data was encountered."}
	NsEDrmPolicyDisableOnline                                    = &Error{0xC00D2762, "NS_E_DRM_POLICY_DISABLE_ONLINE", "The Windows Media Digital Rights Management system cannot perform the requested action because your computer or network administrator has enabled the group policy Prevent Windows Media DRM Internet Access. For assistance, contact your administrator."}
	NsEDrmUnableToCreateAuthenticationObject                     = &Error{0xC00D2763, "NS_E_DRM_UNABLE_TO_CREATE_AUTHENTICATION_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmNotConfigured                                          = &Error{0xC00D2764, "NS_E_DRM_NOT_CONFIGURED", "Not all of the necessary properties for DRM have been set."}
	NsEDrmDeviceActivationCanceled                               = &Error{0xC00D2765, "NS_E_DRM_DEVICE_ACTIVATION_CANCELED", "The portable device does not have the security required to copy protected files to it. To obtain the additional security, try to copy the file to your portable device again. When a message appears, click OK."}
	NsEBackupRestoreTooManyResets                                = &Error{0xC00D2766, "NS_E_BACKUP_RESTORE_TOO_MANY_RESETS", "Too many resets in Backup-Restore."}
	NsEDrmDebuggingNotAllowed                                    = &Error{0xC00D2767, "NS_E_DRM_DEBUGGING_NOT_ALLOWED", "Running this process under a debugger while using DRM content is not allowed."}
	NsEDrmOperationCanceled                                      = &Error{0xC00D2768, "NS_E_DRM_OPERATION_CANCELED", "The user canceled the DRM operation."}
	NsEDrmRestrictionsNotRetrieved                               = &Error{0xC00D2769, "NS_E_DRM_RESTRICTIONS_NOT_RETRIEVED", "The license you are using has assocaited output restrictions. This license is unusable until these restrictions are queried."}
	NsEDrmUnableToCreatePlaylistObject                           = &Error{0xC00D276A, "NS_E_DRM_UNABLE_TO_CREATE_PLAYLIST_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreatePlaylistBurnObject                       = &Error{0xC00D276B, "NS_E_DRM_UNABLE_TO_CREATE_PLAYLIST_BURN_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateDeviceRegistrationObject                 = &Error{0xC00D276C, "NS_E_DRM_UNABLE_TO_CREATE_DEVICE_REGISTRATION_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmUnableToCreateMeteringObject                           = &Error{0xC00D276D, "NS_E_DRM_UNABLE_TO_CREATE_METERING_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmTrackExceededPlaylistRestiction                        = &Error{0xC00D2770, "NS_E_DRM_TRACK_EXCEEDED_PLAYLIST_RESTICTION", "The specified track has exceeded it's specified playlist burn limit in this playlist."}
	NsEDrmTrackExceededTrackburnRestriction                      = &Error{0xC00D2771, "NS_E_DRM_TRACK_EXCEEDED_TRACKBURN_RESTRICTION", "The specified track has exceeded it's track burn limit."}
	NsEDrmUnableToGetDeviceCert                                  = &Error{0xC00D2772, "NS_E_DRM_UNABLE_TO_GET_DEVICE_CERT", "A problem has occurred in obtaining the device's certificate. Contact Microsoft product support."}
	NsEDrmUnableToGetSecureClock                                 = &Error{0xC00D2773, "NS_E_DRM_UNABLE_TO_GET_SECURE_CLOCK", "A problem has occurred in obtaining the device's secure clock. Contact Microsoft product support."}
	NsEDrmUnableToSetSecureClock                                 = &Error{0xC00D2774, "NS_E_DRM_UNABLE_TO_SET_SECURE_CLOCK", "A problem has occurred in setting the device's secure clock. Contact Microsoft product support."}
	NsEDrmUnableToGetSecureClockFromServer                       = &Error{0xC00D2775, "NS_E_DRM_UNABLE_TO_GET_SECURE_CLOCK_FROM_SERVER", "A problem has occurred in obtaining the secure clock from server. Contact Microsoft product support."}
	NsEDrmPolicyMeteringDisabled                                 = &Error{0xC00D2776, "NS_E_DRM_POLICY_METERING_DISABLED", "This content requires the metering policy to be enabled."}
	NsEDrmTransferChainedLicensesUnsupported                     = &Error{0xC00D2777, "NS_E_DRM_TRANSFER_CHAINED_LICENSES_UNSUPPORTED", "Transfer of chained licenses unsupported."}
	NsEDrmSdkVersionmismatch                                     = &Error{0xC00D2778, "NS_E_DRM_SDK_VERSIONMISMATCH", "The Digital Rights Management component is not installed properly. Reinstall the Player."}
	NsEDrmLicNeedsDeviceClockSet                                 = &Error{0xC00D2779, "NS_E_DRM_LIC_NEEDS_DEVICE_CLOCK_SET", "The file could not be transferred because the device clock is not set."}
	NsELicenseHeaderMissingUrl                                   = &Error{0xC00D277A, "NS_E_LICENSE_HEADER_MISSING_URL", "The content header is missing an acquisition URL."}
	NsEDeviceNotWmdrmDevice                                      = &Error{0xC00D277B, "NS_E_DEVICE_NOT_WMDRM_DEVICE", "The current attached device does not support WMDRM."}
	NsEDrmInvalidAppcert                                         = &Error{0xC00D277C, "NS_E_DRM_INVALID_APPCERT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmProtocolForcefulTerminationOnPetition                  = &Error{0xC00D277D, "NS_E_DRM_PROTOCOL_FORCEFUL_TERMINATION_ON_PETITION", "The client application has been forcefully terminated during a DRM petition."}
	NsEDrmProtocolForcefulTerminationOnChallenge                 = &Error{0xC00D277E, "NS_E_DRM_PROTOCOL_FORCEFUL_TERMINATION_ON_CHALLENGE", "The client application has been forcefully terminated during a DRM challenge."}
	NsEDrmCheckpointFailed                                       = &Error{0xC00D277F, "NS_E_DRM_CHECKPOINT_FAILED", "Secure storage protection error. Restore your licenses from a previous backup and try again."}
	NsEDrmBbUnableToInitialize                                   = &Error{0xC00D2780, "NS_E_DRM_BB_UNABLE_TO_INITIALIZE", "A problem has occurred in the Digital Rights Management root of trust. Contact Microsoft product support."}
	NsEDrmUnableToLoadHardwareId                                 = &Error{0xC00D2781, "NS_E_DRM_UNABLE_TO_LOAD_HARDWARE_ID", "A problem has occurred in retrieving the Digital Rights Management machine identification. Contact Microsoft product support."}
	NsEDrmUnableToOpenDataStore                                  = &Error{0xC00D2782, "NS_E_DRM_UNABLE_TO_OPEN_DATA_STORE", "A problem has occurred in opening the Digital Rights Management data storage file. Contact Microsoft product."}
	NsEDrmDatastoreCorrupt                                       = &Error{0xC00D2783, "NS_E_DRM_DATASTORE_CORRUPT", "The Digital Rights Management data storage is not functioning properly. Contact Microsoft product support."}
	NsEDrmUnableToCreateInmemorystoreObject                      = &Error{0xC00D2784, "NS_E_DRM_UNABLE_TO_CREATE_INMEMORYSTORE_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmStublibRequired                                        = &Error{0xC00D2785, "NS_E_DRM_STUBLIB_REQUIRED", "A secured library is required to access the requested functionality."}
	NsEDrmUnableToCreateCertificateObject                        = &Error{0xC00D2786, "NS_E_DRM_UNABLE_TO_CREATE_CERTIFICATE_OBJECT", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmMigrationTargetNotOnline                               = &Error{0xC00D2787, "NS_E_DRM_MIGRATION_TARGET_NOT_ONLINE", "A problem has occurred in the Digital Rights Management component during license migration. Contact Microsoft product support."}
	NsEDrmInvalidMigrationImage                                  = &Error{0xC00D2788, "NS_E_DRM_INVALID_MIGRATION_IMAGE", "A problem has occurred in the Digital Rights Management component during license migration. Contact Microsoft product support."}
	NsEDrmMigrationTargetStatesCorrupted                         = &Error{0xC00D2789, "NS_E_DRM_MIGRATION_TARGET_STATES_CORRUPTED", "A problem has occurred in the Digital Rights Management component during license migration. Contact Microsoft product support."}
	NsEDrmMigrationImporterNotAvailable                          = &Error{0xC00D278A, "NS_E_DRM_MIGRATION_IMPORTER_NOT_AVAILABLE", "A problem has occurred in the Digital Rights Management component during license migration. Contact Microsoft product support."}
	NsDrmEMigrationUpgradeWithDiffSid                            = &Error{0xC00D278B, "NS_DRM_E_MIGRATION_UPGRADE_WITH_DIFF_SID", "A problem has occurred in the Digital Rights Management component during license migration. Contact Microsoft product support."}
	NsDrmEMigrationSourceMachineInUse                            = &Error{0xC00D278C, "NS_DRM_E_MIGRATION_SOURCE_MACHINE_IN_USE", "The Digital Rights Management component is in use during license migration. Contact Microsoft product support."}
	NsDrmEMigrationTargetMachineLessThanLh                       = &Error{0xC00D278D, "NS_DRM_E_MIGRATION_TARGET_MACHINE_LESS_THAN_LH", "Licenses are being migrated to a machine running XP or downlevel OS. This operation can only be performed on Windows Vista or a later OS. Contact Microsoft product support."}
	NsDrmEMigrationImageAlreadyExists                            = &Error{0xC00D278E, "NS_DRM_E_MIGRATION_IMAGE_ALREADY_EXISTS", "Migration Image already exists. Contact Microsoft product support."}
	NsEDrmHardwareidMismatch                                     = &Error{0xC00D278F, "NS_E_DRM_HARDWAREID_MISMATCH", "The requested action cannot be performed because a hardware configuration change has been detected by the Windows Media Digital Rights Management (DRM) components on your computer."}
	NsEInvalidDrmv2cltStublib                                    = &Error{0xC00D2790, "NS_E_INVALID_DRMV2CLT_STUBLIB", "The wrong stublib has been linked to an application or DLL using drmv2clt.dll."}
	NsEDrmMigrationInvalidLegacyv2Data                           = &Error{0xC00D2791, "NS_E_DRM_MIGRATION_INVALID_LEGACYV2_DATA", "The legacy V2 data being imported is invalid."}
	NsEDrmMigrationLicenseAlreadyExists                          = &Error{0xC00D2792, "NS_E_DRM_MIGRATION_LICENSE_ALREADY_EXISTS", "The license being imported already exists."}
	NsEDrmMigrationInvalidLegacyv2SstPassword                    = &Error{0xC00D2793, "NS_E_DRM_MIGRATION_INVALID_LEGACYV2_SST_PASSWORD", "The password of the Legacy V2 SST entry being imported is incorrect."}
	NsEDrmMigrationNotSupported                                  = &Error{0xC00D2794, "NS_E_DRM_MIGRATION_NOT_SUPPORTED", "Migration is not supported by the plugin."}
	NsEDrmUnableToCreateMigrationImporterObject                  = &Error{0xC00D2795, "NS_E_DRM_UNABLE_TO_CREATE_MIGRATION_IMPORTER_OBJECT", "A migration importer cannot be created for this media file. Reinstall the application."}
	NsEDrmCheckpointMismatch                                     = &Error{0xC00D2796, "NS_E_DRM_CHECKPOINT_MISMATCH", "The requested action cannot be performed because a problem occurred with the Windows Media Digital Rights Management (DRM) components on your computer."}
	NsEDrmCheckpointCorrupt                                      = &Error{0xC00D2797, "NS_E_DRM_CHECKPOINT_CORRUPT", "The requested action cannot be performed because a problem occurred with the Windows Media Digital Rights Management (DRM) components on your computer."}
	NsERegFlushFailure                                           = &Error{0xC00D2798, "NS_E_REG_FLUSH_FAILURE", "The requested action cannot be performed because a problem occurred with the Windows Media Digital Rights Management (DRM) components on your computer."}
	NsEHdsKeyMismatch                                            = &Error{0xC00D2799, "NS_E_HDS_KEY_MISMATCH", "The requested action cannot be performed because a problem occurred with the Windows Media Digital Rights Management (DRM) components on your computer."}
	NsEDrmMigrationOperationCancelled                            = &Error{0xC00D279A, "NS_E_DRM_MIGRATION_OPERATION_CANCELLED", "Migration was canceled by the user."}
	NsEDrmMigrationObjectInUse                                   = &Error{0xC00D279B, "NS_E_DRM_MIGRATION_OBJECT_IN_USE", "Migration object is already in use and cannot be called until the current operation completes."}
	NsEDrmMalformedContentHeader                                 = &Error{0xC00D279C, "NS_E_DRM_MALFORMED_CONTENT_HEADER", "The content header does not comply with DRM requirements and cannot be used."}
	NsEDrmLicenseExpired                                         = &Error{0xC00D27D8, "NS_E_DRM_LICENSE_EXPIRED", "The license for this file has expired and is no longer valid. Contact your content provider for further assistance."}
	NsEDrmLicenseNotenabled                                      = &Error{0xC00D27D9, "NS_E_DRM_LICENSE_NOTENABLED", "The license for this file is not valid yet, but will be at a future date."}
	NsEDrmLicenseAppseclow                                       = &Error{0xC00D27DA, "NS_E_DRM_LICENSE_APPSECLOW", "The license for this file requires a higher level of security than the player you are currently using has. Try using a different player or download a newer version of your current player."}
	NsEDrmStoreNeedindi                                          = &Error{0xC00D27DB, "NS_E_DRM_STORE_NEEDINDI", "The license cannot be stored as it requires security upgrade of Digital Rights Management component."}
	NsEDrmStoreNotallowed                                        = &Error{0xC00D27DC, "NS_E_DRM_STORE_NOTALLOWED", "Your machine does not meet the requirements for storing the license."}
	NsEDrmLicenseAppNotallowed                                   = &Error{0xC00D27DD, "NS_E_DRM_LICENSE_APP_NOTALLOWED", "The license for this file requires an upgraded version of your player or a different player."}
	NsEDrmLicenseCertExpired                                     = &Error{0xC00D27DF, "NS_E_DRM_LICENSE_CERT_EXPIRED", "The license server's certificate expired. Make sure your system clock is set correctly. Contact your content provider for further assistance."}
	NsEDrmLicenseSeclow                                          = &Error{0xC00D27E0, "NS_E_DRM_LICENSE_SECLOW", "The license for this file requires a higher level of security than the player you are currently using has. Try using a different player or download a newer version of your current player."}
	NsEDrmLicenseContentRevoked                                  = &Error{0xC00D27E1, "NS_E_DRM_LICENSE_CONTENT_REVOKED", "The content owner for the license you just acquired is no longer supporting their content. Contact the content owner for a newer version of the content."}
	NsEDrmDeviceNotRegistered                                    = &Error{0xC00D27E2, "NS_E_DRM_DEVICE_NOT_REGISTERED", "The content owner for the license you just acquired requires your device to register to the current machine."}
	NsEDrmLicenseNosap                                           = &Error{0xC00D280A, "NS_E_DRM_LICENSE_NOSAP", "The license for this file requires a feature that is not supported in your current player or operating system. You can try with newer version of your current player or contact your content provider for further assistance."}
	NsEDrmLicenseNosvp                                           = &Error{0xC00D280B, "NS_E_DRM_LICENSE_NOSVP", "The license for this file requires a feature that is not supported in your current player or operating system. You can try with newer version of your current player or contact your content provider for further assistance."}
	NsEDrmLicenseNowdm                                           = &Error{0xC00D280C, "NS_E_DRM_LICENSE_NOWDM", "The license for this file requires Windows Driver Model (WDM) audio drivers. Contact your sound card manufacturer for further assistance."}
	NsEDrmLicenseNotrustedcodec                                  = &Error{0xC00D280D, "NS_E_DRM_LICENSE_NOTRUSTEDCODEC", "The license for this file requires a higher level of security than the player you are currently using has. Try using a different player or download a newer version of your current player."}
	NsEDrmSourceidNotSupported                                   = &Error{0xC00D280E, "NS_E_DRM_SOURCEID_NOT_SUPPORTED", "The license for this file is not supported by your current player. You can try with newer version of your current player or contact your content provider for further assistance."}
	NsEDrmNeedsUpgradeTempfile                                   = &Error{0xC00D283D, "NS_E_DRM_NEEDS_UPGRADE_TEMPFILE", "An updated version of your media player is required to play the selected content."}
	NsEDrmNeedUpgradePd                                          = &Error{0xC00D283E, "NS_E_DRM_NEED_UPGRADE_PD", "A new version of the Digital Rights Management component is required. Contact product support for this application to get the latest version."}
	NsEDrmSignatureFailure                                       = &Error{0xC00D283F, "NS_E_DRM_SIGNATURE_FAILURE", "Failed to either create or verify the content header."}
	NsEDrmLicenseServerInfoMissing                               = &Error{0xC00D2840, "NS_E_DRM_LICENSE_SERVER_INFO_MISSING", "Could not read the necessary information from the system registry."}
	NsEDrmBusy                                                   = &Error{0xC00D2841, "NS_E_DRM_BUSY", "The DRM subsystem is currently locked by another application or user. Try again later."}
	NsEDrmPdTooManyDevices                                       = &Error{0xC00D2842, "NS_E_DRM_PD_TOO_MANY_DEVICES", "There are too many target devices registered on the portable media."}
	NsEDrmIndivFraud                                             = &Error{0xC00D2843, "NS_E_DRM_INDIV_FRAUD", "The security upgrade cannot be completed because the allowed number of daily upgrades has been exceeded. Try again tomorrow."}
	NsEDrmIndivNoCabs                                            = &Error{0xC00D2844, "NS_E_DRM_INDIV_NO_CABS", "The security upgrade cannot be completed because the server is unable to perform the operation. Try again later."}
	NsEDrmIndivServiceUnavailable                                = &Error{0xC00D2845, "NS_E_DRM_INDIV_SERVICE_UNAVAILABLE", "The security upgrade cannot be performed because the server is not available. Try again later."}
	NsEDrmRestoreServiceUnavailable                              = &Error{0xC00D2846, "NS_E_DRM_RESTORE_SERVICE_UNAVAILABLE", "Windows Media Player cannot restore your licenses because the server is not available. Try again later."}
	NsEDrmClientCodeExpired                                      = &Error{0xC00D2847, "NS_E_DRM_CLIENT_CODE_EXPIRED", "Windows Media Player cannot play the protected file. Verify that your computer's date is set correctly. If it is correct, on the Help menu, click Check for Player Updates to install the latest version of the Player."}
	NsEDrmNoUplinkLicense                                        = &Error{0xC00D2848, "NS_E_DRM_NO_UPLINK_LICENSE", "The chained license cannot be created because the referenced uplink license does not exist."}
	NsEDrmInvalidKid                                             = &Error{0xC00D2849, "NS_E_DRM_INVALID_KID", "The specified KID is invalid."}
	NsEDrmLicenseInitializationError                             = &Error{0xC00D284A, "NS_E_DRM_LICENSE_INITIALIZATION_ERROR", "License initialization did not work. Contact Microsoft product support."}
	NsEDrmChainTooLong                                           = &Error{0xC00D284C, "NS_E_DRM_CHAIN_TOO_LONG", "The uplink license of a chained license cannot itself be a chained license."}
	NsEDrmUnsupportedAlgorithm                                   = &Error{0xC00D284D, "NS_E_DRM_UNSUPPORTED_ALGORITHM", "The specified encryption algorithm is unsupported."}
	NsEDrmLicenseDeletionError                                   = &Error{0xC00D284E, "NS_E_DRM_LICENSE_DELETION_ERROR", "License deletion did not work. Contact Microsoft product support."}
	NsEDrmInvalidCertificate                                     = &Error{0xC00D28A0, "NS_E_DRM_INVALID_CERTIFICATE", "The client's certificate is corrupted or the signature cannot be verified."}
	NsEDrmCertificateRevoked                                     = &Error{0xC00D28A1, "NS_E_DRM_CERTIFICATE_REVOKED", "The client's certificate has been revoked."}
	NsEDrmLicenseUnavailable                                     = &Error{0xC00D28A2, "NS_E_DRM_LICENSE_UNAVAILABLE", "There is no license available for the requested action."}
	NsEDrmDeviceLimitReached                                     = &Error{0xC00D28A3, "NS_E_DRM_DEVICE_LIMIT_REACHED", "The maximum number of devices in use has been reached. Unable to open additional devices."}
	NsEDrmUnableToVerifyProximity                                = &Error{0xC00D28A4, "NS_E_DRM_UNABLE_TO_VERIFY_PROXIMITY", "The proximity detection procedure could not confirm that the receiver is near the transmitter in the network."}
	NsEDrmMustRegister                                           = &Error{0xC00D28A5, "NS_E_DRM_MUST_REGISTER", "The client must be registered before executing the intended operation."}
	NsEDrmMustApprove                                            = &Error{0xC00D28A6, "NS_E_DRM_MUST_APPROVE", "The client must be approved before executing the intended operation."}
	NsEDrmMustRevalidate                                         = &Error{0xC00D28A7, "NS_E_DRM_MUST_REVALIDATE", "The client must be revalidated before executing the intended operation."}
	NsEDrmInvalidProximityResponse                               = &Error{0xC00D28A8, "NS_E_DRM_INVALID_PROXIMITY_RESPONSE", "The response to the proximity detection challenge is invalid."}
	NsEDrmInvalidSession                                         = &Error{0xC00D28A9, "NS_E_DRM_INVALID_SESSION", "The requested session is invalid."}
	NsEDrmDeviceNotOpen                                          = &Error{0xC00D28AA, "NS_E_DRM_DEVICE_NOT_OPEN", "The device must be opened before it can be used to receive content."}
	NsEDrmDeviceAlreadyRegistered                                = &Error{0xC00D28AB, "NS_E_DRM_DEVICE_ALREADY_REGISTERED", "Device registration failed because the device is already registered."}
	NsEDrmUnsupportedProtocolVersion                             = &Error{0xC00D28AC, "NS_E_DRM_UNSUPPORTED_PROTOCOL_VERSION", "Unsupported WMDRM-ND protocol version."}
	NsEDrmUnsupportedAction                                      = &Error{0xC00D28AD, "NS_E_DRM_UNSUPPORTED_ACTION", "The requested action is not supported."}
	NsEDrmCertificateSecurityLevelInadequate                     = &Error{0xC00D28AE, "NS_E_DRM_CERTIFICATE_SECURITY_LEVEL_INADEQUATE", "The certificate does not have an adequate security level for the requested action."}
	NsEDrmUnableToOpenPort                                       = &Error{0xC00D28AF, "NS_E_DRM_UNABLE_TO_OPEN_PORT", "Unable to open the specified port for receiving Proximity messages."}
	NsEDrmBadRequest                                             = &Error{0xC00D28B0, "NS_E_DRM_BAD_REQUEST", "The message format is invalid."}
	NsEDrmInvalidCrl                                             = &Error{0xC00D28B1, "NS_E_DRM_INVALID_CRL", "The Certificate Revocation List is invalid or corrupted."}
	NsEDrmAttributeTooLong                                       = &Error{0xC00D28B2, "NS_E_DRM_ATTRIBUTE_TOO_LONG", "The length of the attribute name or value is too long."}
	NsEDrmExpiredLicenseblob                                     = &Error{0xC00D28B3, "NS_E_DRM_EXPIRED_LICENSEBLOB", "The license blob passed in the cardea request is expired."}
	NsEDrmInvalidLicenseblob                                     = &Error{0xC00D28B4, "NS_E_DRM_INVALID_LICENSEBLOB", "The license blob passed in the cardea request is invalid. Contact Microsoft product support."}
	NsEDrmInclusionListRequired                                  = &Error{0xC00D28B5, "NS_E_DRM_INCLUSION_LIST_REQUIRED", "The requested operation cannot be performed because the license does not contain an inclusion list."}
	NsEDrmDrmv2cltRevoked                                        = &Error{0xC00D28B6, "NS_E_DRM_DRMV2CLT_REVOKED", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEDrmRivTooSmall                                            = &Error{0xC00D28B7, "NS_E_DRM_RIV_TOO_SMALL", "A problem has occurred in the Digital Rights Management component. Contact Microsoft product support."}
	NsEOutputProtectionLevelUnsupported                          = &Error{0xC00D2904, "NS_E_OUTPUT_PROTECTION_LEVEL_UNSUPPORTED", "Windows Media Player does not support the level of output protection required by the content."}
	NsECompressedDigitalVideoProtectionLevelUnsupported          = &Error{0xC00D2905, "NS_E_COMPRESSED_DIGITAL_VIDEO_PROTECTION_LEVEL_UNSUPPORTED", "Windows Media Player does not support the level of protection required for compressed digital video."}
	NsEUncompressedDigitalVideoProtectionLevelUnsupported        = &Error{0xC00D2906, "NS_E_UNCOMPRESSED_DIGITAL_VIDEO_PROTECTION_LEVEL_UNSUPPORTED", "Windows Media Player does not support the level of protection required for uncompressed digital video."}
	NsEAnalogVideoProtectionLevelUnsupported                     = &Error{0xC00D2907, "NS_E_ANALOG_VIDEO_PROTECTION_LEVEL_UNSUPPORTED", "Windows Media Player does not support the level of protection required for analog video."}
	NsECompressedDigitalAudioProtectionLevelUnsupported          = &Error{0xC00D2908, "NS_E_COMPRESSED_DIGITAL_AUDIO_PROTECTION_LEVEL_UNSUPPORTED", "Windows Media Player does not support the level of protection required for compressed digital audio."}
	NsEUncompressedDigitalAudioProtectionLevelUnsupported        = &Error{0xC00D2909, "NS_E_UNCOMPRESSED_DIGITAL_AUDIO_PROTECTION_LEVEL_UNSUPPORTED", "Windows Media Player does not support the level of protection required for uncompressed digital audio."}
	NsEOutputProtectionSchemeUnsupported                         = &Error{0xC00D290A, "NS_E_OUTPUT_PROTECTION_SCHEME_UNSUPPORTED", "Windows Media Player does not support the scheme of output protection required by the content."}
	NsERebootRecommended                                         = &Error{0xC00D2AFA, "NS_E_REBOOT_RECOMMENDED", "Installation was not successful and some file cleanup is not complete. For best results, restart your computer."}
	NsERebootRequired                                            = &Error{0xC00D2AFB, "NS_E_REBOOT_REQUIRED", "Installation was not successful. To continue, you must restart your computer."}
	NsESetupIncomplete                                           = &Error{0xC00D2AFC, "NS_E_SETUP_INCOMPLETE", "Installation was not successful."}
	NsESetupDrmMigrationFailed                                   = &Error{0xC00D2AFD, "NS_E_SETUP_DRM_MIGRATION_FAILED", "Setup cannot migrate the Windows Media Digital Rights Management (DRM) components."}
	NsESetupIgnorableFailure                                     = &Error{0xC00D2AFE, "NS_E_SETUP_IGNORABLE_FAILURE", "Some skin or playlist components cannot be installed."}
	NsESetupDrmMigrationFailedAndIgnorableFailure                = &Error{0xC00D2AFF, "NS_E_SETUP_DRM_MIGRATION_FAILED_AND_IGNORABLE_FAILURE", "Setup cannot migrate the Windows Media Digital Rights Management (DRM) components. In addition, some skin or playlist components cannot be installed."}
	NsESetupBlocked                                              = &Error{0xC00D2B00, "NS_E_SETUP_BLOCKED", "Installation is blocked because your computer does not meet one or more of the setup requirements."}
	NsEUnknownProtocol                                           = &Error{0xC00D2EE0, "NS_E_UNKNOWN_PROTOCOL", "The specified protocol is not supported."}
	NsERedirectToProxy                                           = &Error{0xC00D2EE1, "NS_E_REDIRECT_TO_PROXY", "The client is redirected to a proxy server."}
	NsEInternalServerError                                       = &Error{0xC00D2EE2, "NS_E_INTERNAL_SERVER_ERROR", "The server encountered an unexpected condition which prevented it from fulfilling the request."}
	NsEBadRequest                                                = &Error{0xC00D2EE3, "NS_E_BAD_REQUEST", "The request could not be understood by the server."}
	NsEErrorFromProxy                                            = &Error{0xC00D2EE4, "NS_E_ERROR_FROM_PROXY", "The proxy experienced an error while attempting to contact the media server."}
	NsEProxyTimeout                                              = &Error{0xC00D2EE5, "NS_E_PROXY_TIMEOUT", "The proxy did not receive a timely response while attempting to contact the media server."}
	NsEServerUnavailable                                         = &Error{0xC00D2EE6, "NS_E_SERVER_UNAVAILABLE", "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}
	NsERefusedByServer                                           = &Error{0xC00D2EE7, "NS_E_REFUSED_BY_SERVER", "The server is refusing to fulfill the requested operation."}
	NsEIncompatibleServer                                        = &Error{0xC00D2EE8, "NS_E_INCOMPATIBLE_SERVER", "The server is not a compatible streaming media server."}
	NsEMulticastDisabled                                         = &Error{0xC00D2EE9, "NS_E_MULTICAST_DISABLED", "The content cannot be streamed because the Multicast protocol has been disabled."}
	NsEInvalidRedirect                                           = &Error{0xC00D2EEA, "NS_E_INVALID_REDIRECT", "The server redirected the player to an invalid location."}
	NsEAllProtocolsDisabled                                      = &Error{0xC00D2EEB, "NS_E_ALL_PROTOCOLS_DISABLED", "The content cannot be streamed because all protocols have been disabled."}
	NsEMsbdNoLongerSupported                                     = &Error{0xC00D2EEC, "NS_E_MSBD_NO_LONGER_SUPPORTED", "The MSBD protocol is no longer supported. Please use HTTP to connect to the Windows Media stream."}
	NsEProxyNotFound                                             = &Error{0xC00D2EED, "NS_E_PROXY_NOT_FOUND", "The proxy server could not be located. Please check your proxy server configuration."}
	NsECannotConnectToProxy                                      = &Error{0xC00D2EEE, "NS_E_CANNOT_CONNECT_TO_PROXY", "Unable to establish a connection to the proxy server. Please check your proxy server configuration."}
	NsEServerDnsTimeout                                          = &Error{0xC00D2EEF, "NS_E_SERVER_DNS_TIMEOUT", "Unable to locate the media server. The operation timed out."}
	NsEProxyDnsTimeout                                           = &Error{0xC00D2EF0, "NS_E_PROXY_DNS_TIMEOUT", "Unable to locate the proxy server. The operation timed out."}
	NsEClosedOnSuspend                                           = &Error{0xC00D2EF1, "NS_E_CLOSED_ON_SUSPEND", "Media closed because Windows was shut down."}
	NsECannotReadPlaylistFromMediaserver                         = &Error{0xC00D2EF2, "NS_E_CANNOT_READ_PLAYLIST_FROM_MEDIASERVER", "Unable to read the contents of a playlist file from a media server."}
	NsESessionNotFound                                           = &Error{0xC00D2EF3, "NS_E_SESSION_NOT_FOUND", "Session not found."}
	NsERequireStreamingClient                                    = &Error{0xC00D2EF4, "NS_E_REQUIRE_STREAMING_CLIENT", "Content requires a streaming media client."}
	NsEPlaylistEntryHasChanged                                   = &Error{0xC00D2EF5, "NS_E_PLAYLIST_ENTRY_HAS_CHANGED", "A command applies to a previous playlist entry."}
	NsEProxyAccessdenied                                         = &Error{0xC00D2EF6, "NS_E_PROXY_ACCESSDENIED", "The proxy server is denying access. The username and/or password might be incorrect."}
	NsEProxySourceAccessdenied                                   = &Error{0xC00D2EF7, "NS_E_PROXY_SOURCE_ACCESSDENIED", "The proxy could not provide valid authentication credentials to the media server."}
	NsENetworkSinkWrite                                          = &Error{0xC00D2EF8, "NS_E_NETWORK_SINK_WRITE", "The network sink failed to write data to the network."}
	NsEFirewall                                                  = &Error{0xC00D2EF9, "NS_E_FIREWALL", "Packets are not being received from the server. The packets might be blocked by a filtering device, such as a network firewall."}
	NsEMmsNotSupported                                           = &Error{0xC00D2EFA, "NS_E_MMS_NOT_SUPPORTED", "The MMS protocol is not supported. Please use HTTP or RTSP to connect to the Windows Media stream."}
	NsEServerAccessdenied                                        = &Error{0xC00D2EFB, "NS_E_SERVER_ACCESSDENIED", "The Windows Media server is denying access. The username and/or password might be incorrect."}
	NsEResourceGone                                              = &Error{0xC00D2EFC, "NS_E_RESOURCE_GONE", "The Publishing Point or file on the Windows Media Server is no longer available."}
	NsENoExistingPacketizer                                      = &Error{0xC00D2EFD, "NS_E_NO_EXISTING_PACKETIZER", "There is no existing packetizer plugin for a stream."}
	NsEBadSyntaxInServerResponse                                 = &Error{0xC00D2EFE, "NS_E_BAD_SYNTAX_IN_SERVER_RESPONSE", "The response from the media server could not be understood. This might be caused by an incompatible proxy server or media server."}
	NsEResetSocketConnection                                     = &Error{0xC00D2F00, "NS_E_RESET_SOCKET_CONNECTION", "The Windows Media Server reset the network connection."}
	NsETooManyHops                                               = &Error{0xC00D2F02, "NS_E_TOO_MANY_HOPS", "The request could not reach the media server (too many hops)."}
	NsETooMuchDataFromServer                                     = &Error{0xC00D2F05, "NS_E_TOO_MUCH_DATA_FROM_SERVER", "The server is sending too much data. The connection has been terminated."}
	NsEConnectTimeout                                            = &Error{0xC00D2F06, "NS_E_CONNECT_TIMEOUT", "It was not possible to establish a connection to the media server in a timely manner. The media server might be down for maintenance, or it might be necessary to use a proxy server to access this media server."}
	NsEProxyConnectTimeout                                       = &Error{0xC00D2F07, "NS_E_PROXY_CONNECT_TIMEOUT", "It was not possible to establish a connection to the proxy server in a timely manner. Please check your proxy server configuration."}
	NsESessionInvalid                                            = &Error{0xC00D2F08, "NS_E_SESSION_INVALID", "Session not found."}
	NsEPacketsinkUnknownFecStream                                = &Error{0xC00D2F0A, "NS_E_PACKETSINK_UNKNOWN_FEC_STREAM", "Unknown packet sink stream."}
	NsEPushCannotconnect                                         = &Error{0xC00D2F0B, "NS_E_PUSH_CANNOTCONNECT", "Unable to establish a connection to the server. Ensure Windows Media Services is started and the HTTP Server control protocol is properly enabled."}
	NsEIncompatiblePushServer                                    = &Error{0xC00D2F0C, "NS_E_INCOMPATIBLE_PUSH_SERVER", "The Server service that received the HTTP push request is not a compatible version of Windows Media Services (WMS). This error might indicate the push request was received by IIS instead of WMS. Ensure WMS is started and has the HTTP Server control protocol properly enabled and try again."}
	NsEEndOfPlaylist                                             = &Error{0xC00D32C8, "NS_E_END_OF_PLAYLIST", "The playlist has reached its end."}
	NsEUseFileSource                                             = &Error{0xC00D32C9, "NS_E_USE_FILE_SOURCE", "Use file source."}
	NsEPropertyNotFound                                          = &Error{0xC00D32CA, "NS_E_PROPERTY_NOT_FOUND", "The property was not found."}
	NsEPropertyReadOnly                                          = &Error{0xC00D32CC, "NS_E_PROPERTY_READ_ONLY", "The property is read only."}
	NsETableKeyNotFound                                          = &Error{0xC00D32CD, "NS_E_TABLE_KEY_NOT_FOUND", "The table key was not found."}
	NsEInvalidQueryOperator                                      = &Error{0xC00D32CF, "NS_E_INVALID_QUERY_OPERATOR", "Invalid query operator."}
	NsEInvalidQueryProperty                                      = &Error{0xC00D32D0, "NS_E_INVALID_QUERY_PROPERTY", "Invalid query property."}
	NsEPropertyNotSupported                                      = &Error{0xC00D32D2, "NS_E_PROPERTY_NOT_SUPPORTED", "The property is not supported."}
	NsESchemaClassifyFailure                                     = &Error{0xC00D32D4, "NS_E_SCHEMA_CLASSIFY_FAILURE", "Schema classification failure."}
	NsEMetadataFormatNotSupported                                = &Error{0xC00D32D5, "NS_E_METADATA_FORMAT_NOT_SUPPORTED", "The metadata format is not supported."}
	NsEMetadataNoEditingCapability                               = &Error{0xC00D32D6, "NS_E_METADATA_NO_EDITING_CAPABILITY", "Cannot edit the metadata."}
	NsEMetadataCannotSetLocale                                   = &Error{0xC00D32D7, "NS_E_METADATA_CANNOT_SET_LOCALE", "Cannot set the locale id."}
	NsEMetadataLanguageNotSuported                               = &Error{0xC00D32D8, "NS_E_METADATA_LANGUAGE_NOT_SUPORTED", "The language is not supported in the format."}
	NsEMetadataNoRfc1766NameForLocale                            = &Error{0xC00D32D9, "NS_E_METADATA_NO_RFC1766_NAME_FOR_LOCALE", "There is no RFC1766 name translation for the supplied locale id."}
	NsEMetadataNotAvailable                                      = &Error{0xC00D32DA, "NS_E_METADATA_NOT_AVAILABLE", "The metadata (or metadata item) is not available."}
	NsEMetadataCacheDataNotAvailable                             = &Error{0xC00D32DB, "NS_E_METADATA_CACHE_DATA_NOT_AVAILABLE", "The cached metadata (or metadata item) is not available."}
	NsEMetadataInvalidDocumentType                               = &Error{0xC00D32DC, "NS_E_METADATA_INVALID_DOCUMENT_TYPE", "The metadata document is invalid."}
	NsEMetadataIdentifierNotAvailable                            = &Error{0xC00D32DD, "NS_E_METADATA_IDENTIFIER_NOT_AVAILABLE", "The metadata content identifier is not available."}
	NsEMetadataCannotRetrieveFromOfflineCache                    = &Error{0xC00D32DE, "NS_E_METADATA_CANNOT_RETRIEVE_FROM_OFFLINE_CACHE", "Cannot retrieve metadata from the offline metadata cache."}
	ErrorMonitorInvalidDescriptorChecksum                        = &Error{0xC0261003, "ERROR_MONITOR_INVALID_DESCRIPTOR_CHECKSUM", "Checksum of the obtained monitor descriptor is invalid."}
	ErrorMonitorInvalidStandardTimingBlock                       = &Error{0xC0261004, "ERROR_MONITOR_INVALID_STANDARD_TIMING_BLOCK", "Monitor descriptor contains an invalid standard timing block."}
	ErrorMonitorWmiDatablockRegistrationFailed                   = &Error{0xC0261005, "ERROR_MONITOR_WMI_DATABLOCK_REGISTRATION_FAILED", "Windows Management Instrumentation (WMI) data block registration failed for one of the MSMonitorClass WMI subclasses."}
	ErrorMonitorInvalidSerialNumberMondscBlock                   = &Error{0xC0261006, "ERROR_MONITOR_INVALID_SERIAL_NUMBER_MONDSC_BLOCK", "Provided monitor descriptor block is either corrupted or does not contain the monitor's detailed serial number."}
	ErrorMonitorInvalidUserFriendlyMondscBlock                   = &Error{0xC0261007, "ERROR_MONITOR_INVALID_USER_FRIENDLY_MONDSC_BLOCK", "Provided monitor descriptor block is either corrupted or does not contain the monitor's user-friendly name."}
	ErrorMonitorNoMoreDescriptorData                             = &Error{0xC0261008, "ERROR_MONITOR_NO_MORE_DESCRIPTOR_DATA", "There is no monitor descriptor data at the specified (offset, size) region."}
	ErrorMonitorInvalidDetailedTimingBlock                       = &Error{0xC0261009, "ERROR_MONITOR_INVALID_DETAILED_TIMING_BLOCK", "Monitor descriptor contains an invalid detailed timing block."}
	ErrorGraphicsNotExclusiveModeOwner                           = &Error{0xC0262000, "ERROR_GRAPHICS_NOT_EXCLUSIVE_MODE_OWNER", "Exclusive mode ownership is needed to create unmanaged primary allocation."}
	ErrorGraphicsInsufficientDmaBuffer                           = &Error{0xC0262001, "ERROR_GRAPHICS_INSUFFICIENT_DMA_BUFFER", "The driver needs more direct memory access (DMA) buffer space to complete the requested operation."}
	ErrorGraphicsInvalidDisplayAdapter                           = &Error{0xC0262002, "ERROR_GRAPHICS_INVALID_DISPLAY_ADAPTER", "Specified display adapter handle is invalid."}
	ErrorGraphicsAdapterWasReset                                 = &Error{0xC0262003, "ERROR_GRAPHICS_ADAPTER_WAS_RESET", "Specified display adapter and all of its state has been reset."}
	ErrorGraphicsInvalidDriverModel                              = &Error{0xC0262004, "ERROR_GRAPHICS_INVALID_DRIVER_MODEL", "The driver stack does not match the expected driver model."}
	ErrorGraphicsPresentModeChanged                              = &Error{0xC0262005, "ERROR_GRAPHICS_PRESENT_MODE_CHANGED", "Present happened but ended up into the changed desktop mode."}
	ErrorGraphicsPresentOccluded                                 = &Error{0xC0262006, "ERROR_GRAPHICS_PRESENT_OCCLUDED", "Nothing to present due to desktop occlusion."}
	ErrorGraphicsPresentDenied                                   = &Error{0xC0262007, "ERROR_GRAPHICS_PRESENT_DENIED", "Not able to present due to denial of desktop access."}
	ErrorGraphicsCannotcolorconvert                              = &Error{0xC0262008, "ERROR_GRAPHICS_CANNOTCOLORCONVERT", "Not able to present with color conversion."}
	ErrorGraphicsNoVideoMemory                                   = &Error{0xC0262100, "ERROR_GRAPHICS_NO_VIDEO_MEMORY", "Not enough video memory available to complete the operation."}
	ErrorGraphicsCantLockMemory                                  = &Error{0xC0262101, "ERROR_GRAPHICS_CANT_LOCK_MEMORY", "Could not probe and lock the underlying memory of an allocation."}
	ErrorGraphicsAllocationBusy                                  = &Error{0xC0262102, "ERROR_GRAPHICS_ALLOCATION_BUSY", "The allocation is currently busy."}
	ErrorGraphicsTooManyReferences                               = &Error{0xC0262103, "ERROR_GRAPHICS_TOO_MANY_REFERENCES", "An object being referenced has reach the maximum reference count already and cannot be referenced further."}
	ErrorGraphicsTryAgainLater                                   = &Error{0xC0262104, "ERROR_GRAPHICS_TRY_AGAIN_LATER", "A problem could not be solved due to some currently existing condition. The problem should be tried again later."}
	ErrorGraphicsTryAgainNow                                     = &Error{0xC0262105, "ERROR_GRAPHICS_TRY_AGAIN_NOW", "A problem could not be solved due to some currently existing condition. The problem should be tried again immediately."}
	ErrorGraphicsAllocationInvalid                               = &Error{0xC0262106, "ERROR_GRAPHICS_ALLOCATION_INVALID", "The allocation is invalid."}
	ErrorGraphicsUnswizzlingApertureUnavailable                  = &Error{0xC0262107, "ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNAVAILABLE", "No more unswizzling apertures are currently available."}
	ErrorGraphicsUnswizzlingApertureUnsupported                  = &Error{0xC0262108, "ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNSUPPORTED", "The current allocation cannot be unswizzled by an aperture."}
	ErrorGraphicsCantEvictPinnedAllocation                       = &Error{0xC0262109, "ERROR_GRAPHICS_CANT_EVICT_PINNED_ALLOCATION", "The request failed because a pinned allocation cannot be evicted."}
	ErrorGraphicsInvalidAllocationUsage                          = &Error{0xC0262110, "ERROR_GRAPHICS_INVALID_ALLOCATION_USAGE", "The allocation cannot be used from its current segment location for the specified operation."}
	ErrorGraphicsCantRenderLockedAllocation                      = &Error{0xC0262111, "ERROR_GRAPHICS_CANT_RENDER_LOCKED_ALLOCATION", "A locked allocation cannot be used in the current command buffer."}
	ErrorGraphicsAllocationClosed                                = &Error{0xC0262112, "ERROR_GRAPHICS_ALLOCATION_CLOSED", "The allocation being referenced has been closed permanently."}
	ErrorGraphicsInvalidAllocationInstance                       = &Error{0xC0262113, "ERROR_GRAPHICS_INVALID_ALLOCATION_INSTANCE", "An invalid allocation instance is being referenced."}
	ErrorGraphicsInvalidAllocationHandle                         = &Error{0xC0262114, "ERROR_GRAPHICS_INVALID_ALLOCATION_HANDLE", "An invalid allocation handle is being referenced."}
	ErrorGraphicsWrongAllocationDevice                           = &Error{0xC0262115, "ERROR_GRAPHICS_WRONG_ALLOCATION_DEVICE", "The allocation being referenced does not belong to the current device."}
	ErrorGraphicsAllocationContentLost                           = &Error{0xC0262116, "ERROR_GRAPHICS_ALLOCATION_CONTENT_LOST", "The specified allocation lost its content."}
	ErrorGraphicsGpuExceptionOnDevice                            = &Error{0xC0262200, "ERROR_GRAPHICS_GPU_EXCEPTION_ON_DEVICE", "Graphics processing unit (GPU) exception is detected on the given device. The device is not able to be scheduled."}
	ErrorGraphicsInvalidVidpnTopology                            = &Error{0xC0262300, "ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY", "Specified video present network (VidPN) topology is invalid."}
	ErrorGraphicsVidpnTopologyNotSupported                       = &Error{0xC0262301, "ERROR_GRAPHICS_VIDPN_TOPOLOGY_NOT_SUPPORTED", "Specified VidPN topology is valid but is not supported by this model of the display adapter."}
	ErrorGraphicsVidpnTopologyCurrentlyNotSupported              = &Error{0xC0262302, "ERROR_GRAPHICS_VIDPN_TOPOLOGY_CURRENTLY_NOT_SUPPORTED", "Specified VidPN topology is valid but is not supported by the display adapter at this time, due to current allocation of its resources."}
	ErrorGraphicsInvalidVidpn                                    = &Error{0xC0262303, "ERROR_GRAPHICS_INVALID_VIDPN", "Specified VidPN handle is invalid."}
	ErrorGraphicsInvalidVideoPresentSource                       = &Error{0xC0262304, "ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE", "Specified video present source is invalid."}
	ErrorGraphicsInvalidVideoPresentTarget                       = &Error{0xC0262305, "ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET", "Specified video present target is invalid."}
	ErrorGraphicsVidpnModalityNotSupported                       = &Error{0xC0262306, "ERROR_GRAPHICS_VIDPN_MODALITY_NOT_SUPPORTED", "Specified VidPN modality is not supported (for example, at least two of the pinned modes are not cofunctional)."}
	ErrorGraphicsInvalidVidpnSourcemodeset                       = &Error{0xC0262308, "ERROR_GRAPHICS_INVALID_VIDPN_SOURCEMODESET", "Specified VidPN source mode set is invalid."}
	ErrorGraphicsInvalidVidpnTargetmodeset                       = &Error{0xC0262309, "ERROR_GRAPHICS_INVALID_VIDPN_TARGETMODESET", "Specified VidPN target mode set is invalid."}
	ErrorGraphicsInvalidFrequency                                = &Error{0xC026230A, "ERROR_GRAPHICS_INVALID_FREQUENCY", "Specified video signal frequency is invalid."}
	ErrorGraphicsInvalidActiveRegion                             = &Error{0xC026230B, "ERROR_GRAPHICS_INVALID_ACTIVE_REGION", "Specified video signal active region is invalid."}
	ErrorGraphicsInvalidTotalRegion                              = &Error{0xC026230C, "ERROR_GRAPHICS_INVALID_TOTAL_REGION", "Specified video signal total region is invalid."}
	ErrorGraphicsInvalidVideoPresentSourceMode                   = &Error{0xC0262310, "ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE_MODE", "Specified video present source mode is invalid."}
	ErrorGraphicsInvalidVideoPresentTargetMode                   = &Error{0xC0262311, "ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET_MODE", "Specified video present target mode is invalid."}
	ErrorGraphicsPinnedModeMustRemainInSet                       = &Error{0xC0262312, "ERROR_GRAPHICS_PINNED_MODE_MUST_REMAIN_IN_SET", "Pinned mode must remain in the set on VidPN's cofunctional modality enumeration."}
	ErrorGraphicsPathAlreadyInTopology                           = &Error{0xC0262313, "ERROR_GRAPHICS_PATH_ALREADY_IN_TOPOLOGY", "Specified video present path is already in the VidPN topology."}
	ErrorGraphicsModeAlreadyInModeset                            = &Error{0xC0262314, "ERROR_GRAPHICS_MODE_ALREADY_IN_MODESET", "Specified mode is already in the mode set."}
	ErrorGraphicsInvalidVideopresentsourceset                    = &Error{0xC0262315, "ERROR_GRAPHICS_INVALID_VIDEOPRESENTSOURCESET", "Specified video present source set is invalid."}
	ErrorGraphicsInvalidVideopresenttargetset                    = &Error{0xC0262316, "ERROR_GRAPHICS_INVALID_VIDEOPRESENTTARGETSET", "Specified video present target set is invalid."}
	ErrorGraphicsSourceAlreadyInSet                              = &Error{0xC0262317, "ERROR_GRAPHICS_SOURCE_ALREADY_IN_SET", "Specified video present source is already in the video present source set."}
	ErrorGraphicsTargetAlreadyInSet                              = &Error{0xC0262318, "ERROR_GRAPHICS_TARGET_ALREADY_IN_SET", "Specified video present target is already in the video present target set."}
	ErrorGraphicsInvalidVidpnPresentPath                         = &Error{0xC0262319, "ERROR_GRAPHICS_INVALID_VIDPN_PRESENT_PATH", "Specified VidPN present path is invalid."}
	ErrorGraphicsNoRecommendedVidpnTopology                      = &Error{0xC026231A, "ERROR_GRAPHICS_NO_RECOMMENDED_VIDPN_TOPOLOGY", "Miniport has no recommendation for augmentation of the specified VidPN topology."}
	ErrorGraphicsInvalidMonitorFrequencyrangeset                 = &Error{0xC026231B, "ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGESET", "Specified monitor frequency range set is invalid."}
	ErrorGraphicsInvalidMonitorFrequencyrange                    = &Error{0xC026231C, "ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE", "Specified monitor frequency range is invalid."}
	ErrorGraphicsFrequencyrangeNotInSet                          = &Error{0xC026231D, "ERROR_GRAPHICS_FREQUENCYRANGE_NOT_IN_SET", "Specified frequency range is not in the specified monitor frequency range set."}
	ErrorGraphicsFrequencyrangeAlreadyInSet                      = &Error{0xC026231F, "ERROR_GRAPHICS_FREQUENCYRANGE_ALREADY_IN_SET", "Specified frequency range is already in the specified monitor frequency range set."}
	ErrorGraphicsStaleModeset                                    = &Error{0xC0262320, "ERROR_GRAPHICS_STALE_MODESET", "Specified mode set is stale. Reacquire the new mode set."}
	ErrorGraphicsInvalidMonitorSourcemodeset                     = &Error{0xC0262321, "ERROR_GRAPHICS_INVALID_MONITOR_SOURCEMODESET", "Specified monitor source mode set is invalid."}
	ErrorGraphicsInvalidMonitorSourceMode                        = &Error{0xC0262322, "ERROR_GRAPHICS_INVALID_MONITOR_SOURCE_MODE", "Specified monitor source mode is invalid."}
	ErrorGraphicsNoRecommendedFunctionalVidpn                    = &Error{0xC0262323, "ERROR_GRAPHICS_NO_RECOMMENDED_FUNCTIONAL_VIDPN", "Miniport does not have any recommendation regarding the request to provide a functional VidPN given the current display adapter configuration."}
	ErrorGraphicsModeIdMustBeUnique                              = &Error{0xC0262324, "ERROR_GRAPHICS_MODE_ID_MUST_BE_UNIQUE", "ID of the specified mode is already used by another mode in the set."}
	ErrorGraphicsEmptyAdapterMonitorModeSupportIntersection      = &Error{0xC0262325, "ERROR_GRAPHICS_EMPTY_ADAPTER_MONITOR_MODE_SUPPORT_INTERSECTION", "System failed to determine a mode that is supported by both the display adapter and the monitor connected to it."}
	ErrorGraphicsVideoPresentTargetsLessThanSources              = &Error{0xC0262326, "ERROR_GRAPHICS_VIDEO_PRESENT_TARGETS_LESS_THAN_SOURCES", "Number of video present targets must be greater than or equal to the number of video present sources."}
	ErrorGraphicsPathNotInTopology                               = &Error{0xC0262327, "ERROR_GRAPHICS_PATH_NOT_IN_TOPOLOGY", "Specified present path is not in the VidPN topology."}
	ErrorGraphicsAdapterMustHaveAtLeastOneSource                 = &Error{0xC0262328, "ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_SOURCE", "Display adapter must have at least one video present source."}
	ErrorGraphicsAdapterMustHaveAtLeastOneTarget                 = &Error{0xC0262329, "ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_TARGET", "Display adapter must have at least one video present target."}
	ErrorGraphicsInvalidMonitordescriptorset                     = &Error{0xC026232A, "ERROR_GRAPHICS_INVALID_MONITORDESCRIPTORSET", "Specified monitor descriptor set is invalid."}
	ErrorGraphicsInvalidMonitordescriptor                        = &Error{0xC026232B, "ERROR_GRAPHICS_INVALID_MONITORDESCRIPTOR", "Specified monitor descriptor is invalid."}
	ErrorGraphicsMonitordescriptorNotInSet                       = &Error{0xC026232C, "ERROR_GRAPHICS_MONITORDESCRIPTOR_NOT_IN_SET", "Specified descriptor is not in the specified monitor descriptor set."}
	ErrorGraphicsMonitordescriptorAlreadyInSet                   = &Error{0xC026232D, "ERROR_GRAPHICS_MONITORDESCRIPTOR_ALREADY_IN_SET", "Specified descriptor is already in the specified monitor descriptor set."}
	ErrorGraphicsMonitordescriptorIdMustBeUnique                 = &Error{0xC026232E, "ERROR_GRAPHICS_MONITORDESCRIPTOR_ID_MUST_BE_UNIQUE", "ID of the specified monitor descriptor is already used by another descriptor in the set."}
	ErrorGraphicsInvalidVidpnTargetSubsetType                    = &Error{0xC026232F, "ERROR_GRAPHICS_INVALID_VIDPN_TARGET_SUBSET_TYPE", "Specified video present target subset type is invalid."}
	ErrorGraphicsResourcesNotRelated                             = &Error{0xC0262330, "ERROR_GRAPHICS_RESOURCES_NOT_RELATED", "Two or more of the specified resources are not related to each other, as defined by the interface semantics."}
	ErrorGraphicsSourceIdMustBeUnique                            = &Error{0xC0262331, "ERROR_GRAPHICS_SOURCE_ID_MUST_BE_UNIQUE", "ID of the specified video present source is already used by another source in the set."}
	ErrorGraphicsTargetIdMustBeUnique                            = &Error{0xC0262332, "ERROR_GRAPHICS_TARGET_ID_MUST_BE_UNIQUE", "ID of the specified video present target is already used by another target in the set."}
	ErrorGraphicsNoAvailableVidpnTarget                          = &Error{0xC0262333, "ERROR_GRAPHICS_NO_AVAILABLE_VIDPN_TARGET", "Specified VidPN source cannot be used because there is no available VidPN target to connect it to."}
	ErrorGraphicsMonitorCouldNotBeAssociatedWithAdapter          = &Error{0xC0262334, "ERROR_GRAPHICS_MONITOR_COULD_NOT_BE_ASSOCIATED_WITH_ADAPTER", "Newly arrived monitor could not be associated with a display adapter."}
	ErrorGraphicsNoVidpnmgr                                      = &Error{0xC0262335, "ERROR_GRAPHICS_NO_VIDPNMGR", "Display adapter in question does not have an associated VidPN manager."}
	ErrorGraphicsNoActiveVidpn                                   = &Error{0xC0262336, "ERROR_GRAPHICS_NO_ACTIVE_VIDPN", "VidPN manager of the display adapter in question does not have an active VidPN."}
	ErrorGraphicsStaleVidpnTopology                              = &Error{0xC0262337, "ERROR_GRAPHICS_STALE_VIDPN_TOPOLOGY", "Specified VidPN topology is stale. Re-acquire the new topology."}
	ErrorGraphicsMonitorNotConnected                             = &Error{0xC0262338, "ERROR_GRAPHICS_MONITOR_NOT_CONNECTED", "There is no monitor connected on the specified video present target."}
	ErrorGraphicsSourceNotInTopology                             = &Error{0xC0262339, "ERROR_GRAPHICS_SOURCE_NOT_IN_TOPOLOGY", "Specified source is not part of the specified VidPN topology."}
	ErrorGraphicsInvalidPrimarysurfaceSize                       = &Error{0xC026233A, "ERROR_GRAPHICS_INVALID_PRIMARYSURFACE_SIZE", "Specified primary surface size is invalid."}
	ErrorGraphicsInvalidVisibleregionSize                        = &Error{0xC026233B, "ERROR_GRAPHICS_INVALID_VISIBLEREGION_SIZE", "Specified visible region size is invalid."}
	ErrorGraphicsInvalidStride                                   = &Error{0xC026233C, "ERROR_GRAPHICS_INVALID_STRIDE", "Specified stride is invalid."}
	ErrorGraphicsInvalidPixelformat                              = &Error{0xC026233D, "ERROR_GRAPHICS_INVALID_PIXELFORMAT", "Specified pixel format is invalid."}
	ErrorGraphicsInvalidColorbasis                               = &Error{0xC026233E, "ERROR_GRAPHICS_INVALID_COLORBASIS", "Specified color basis is invalid."}
	ErrorGraphicsInvalidPixelvalueaccessmode                     = &Error{0xC026233F, "ERROR_GRAPHICS_INVALID_PIXELVALUEACCESSMODE", "Specified pixel value access mode is invalid."}
	ErrorGraphicsTargetNotInTopology                             = &Error{0xC0262340, "ERROR_GRAPHICS_TARGET_NOT_IN_TOPOLOGY", "Specified target is not part of the specified VidPN topology."}
	ErrorGraphicsNoDisplayModeManagementSupport                  = &Error{0xC0262341, "ERROR_GRAPHICS_NO_DISPLAY_MODE_MANAGEMENT_SUPPORT", "Failed to acquire display mode management interface."}
	ErrorGraphicsVidpnSourceInUse                                = &Error{0xC0262342, "ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE", "Specified VidPN source is already owned by a display mode manager (DMM) client and cannot be used until that client releases it."}
	ErrorGraphicsCantAccessActiveVidpn                           = &Error{0xC0262343, "ERROR_GRAPHICS_CANT_ACCESS_ACTIVE_VIDPN", "Specified VidPN is active and cannot be accessed."}
	ErrorGraphicsInvalidPathImportanceOrdinal                    = &Error{0xC0262344, "ERROR_GRAPHICS_INVALID_PATH_IMPORTANCE_ORDINAL", "Specified VidPN present path importance ordinal is invalid."}
	ErrorGraphicsInvalidPathContentGeometryTransformation        = &Error{0xC0262345, "ERROR_GRAPHICS_INVALID_PATH_CONTENT_GEOMETRY_TRANSFORMATION", "Specified VidPN present path content geometry transformation is invalid."}
	ErrorGraphicsPathContentGeometryTransformationNotSupported   = &Error{0xC0262346, "ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_SUPPORTED", "Specified content geometry transformation is not supported on the respective VidPN present path."}
	ErrorGraphicsInvalidGammaRamp                                = &Error{0xC0262347, "ERROR_GRAPHICS_INVALID_GAMMA_RAMP", "Specified gamma ramp is invalid."}
	ErrorGraphicsGammaRampNotSupported                           = &Error{0xC0262348, "ERROR_GRAPHICS_GAMMA_RAMP_NOT_SUPPORTED", "Specified gamma ramp is not supported on the respective VidPN present path."}
	ErrorGraphicsMultisamplingNotSupported                       = &Error{0xC0262349, "ERROR_GRAPHICS_MULTISAMPLING_NOT_SUPPORTED", "Multisampling is not supported on the respective VidPN present path."}
	ErrorGraphicsModeNotInModeset                                = &Error{0xC026234A, "ERROR_GRAPHICS_MODE_NOT_IN_MODESET", "Specified mode is not in the specified mode set."}
	ErrorGraphicsInvalidVidpnTopologyRecommendationReason        = &Error{0xC026234D, "ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY_RECOMMENDATION_REASON", "Specified VidPN topology recommendation reason is invalid."}
	ErrorGraphicsInvalidPathContentType                          = &Error{0xC026234E, "ERROR_GRAPHICS_INVALID_PATH_CONTENT_TYPE", "Specified VidPN present path content type is invalid."}
	ErrorGraphicsInvalidCopyprotectionType                       = &Error{0xC026234F, "ERROR_GRAPHICS_INVALID_COPYPROTECTION_TYPE", "Specified VidPN present path copy protection type is invalid."}
	ErrorGraphicsUnassignedModesetAlreadyExists                  = &Error{0xC0262350, "ERROR_GRAPHICS_UNASSIGNED_MODESET_ALREADY_EXISTS", "No more than one unassigned mode set can exist at any given time for a given VidPN source or target."}
	ErrorGraphicsInvalidScanlineOrdering                         = &Error{0xC0262352, "ERROR_GRAPHICS_INVALID_SCANLINE_ORDERING", "The specified scan line ordering type is invalid."}
	ErrorGraphicsTopologyChangesNotAllowed                       = &Error{0xC0262353, "ERROR_GRAPHICS_TOPOLOGY_CHANGES_NOT_ALLOWED", "Topology changes are not allowed for the specified VidPN."}
	ErrorGraphicsNoAvailableImportanceOrdinals                   = &Error{0xC0262354, "ERROR_GRAPHICS_NO_AVAILABLE_IMPORTANCE_ORDINALS", "All available importance ordinals are already used in the specified topology."}
	ErrorGraphicsIncompatiblePrivateFormat                       = &Error{0xC0262355, "ERROR_GRAPHICS_INCOMPATIBLE_PRIVATE_FORMAT", "Specified primary surface has a different private format attribute than the current primary surface."}
	ErrorGraphicsInvalidModePruningAlgorithm                     = &Error{0xC0262356, "ERROR_GRAPHICS_INVALID_MODE_PRUNING_ALGORITHM", "Specified mode pruning algorithm is invalid."}
	ErrorGraphicsSpecifiedChildAlreadyConnected                  = &Error{0xC0262400, "ERROR_GRAPHICS_SPECIFIED_CHILD_ALREADY_CONNECTED", "Specified display adapter child device already has an external device connected to it."}
	ErrorGraphicsChildDescriptorNotSupported                     = &Error{0xC0262401, "ERROR_GRAPHICS_CHILD_DESCRIPTOR_NOT_SUPPORTED", "The display adapter child device does not support reporting a descriptor."}
	ErrorGraphicsNotALinkedAdapter                               = &Error{0xC0262430, "ERROR_GRAPHICS_NOT_A_LINKED_ADAPTER", "The display adapter is not linked to any other adapters."}
	ErrorGraphicsLeadlinkNotEnumerated                           = &Error{0xC0262431, "ERROR_GRAPHICS_LEADLINK_NOT_ENUMERATED", "Lead adapter in a linked configuration was not enumerated yet."}
	ErrorGraphicsChainlinksNotEnumerated                         = &Error{0xC0262432, "ERROR_GRAPHICS_CHAINLINKS_NOT_ENUMERATED", "Some chain adapters in a linked configuration were not enumerated yet."}
	ErrorGraphicsAdapterChainNotReady                            = &Error{0xC0262433, "ERROR_GRAPHICS_ADAPTER_CHAIN_NOT_READY", "The chain of linked adapters is not ready to start because of an unknown failure."}
	ErrorGraphicsChainlinksNotStarted                            = &Error{0xC0262434, "ERROR_GRAPHICS_CHAINLINKS_NOT_STARTED", "An attempt was made to start a lead link display adapter when the chain links were not started yet."}
	ErrorGraphicsChainlinksNotPoweredOn                          = &Error{0xC0262435, "ERROR_GRAPHICS_CHAINLINKS_NOT_POWERED_ON", "An attempt was made to turn on a lead link display adapter when the chain links were turned off."}
	ErrorGraphicsInconsistentDeviceLinkState                     = &Error{0xC0262436, "ERROR_GRAPHICS_INCONSISTENT_DEVICE_LINK_STATE", "The adapter link was found to be in an inconsistent state. Not all adapters are in an expected PNP or power state."}
	ErrorGraphicsNotPostDeviceDriver                             = &Error{0xC0262438, "ERROR_GRAPHICS_NOT_POST_DEVICE_DRIVER", "The driver trying to start is not the same as the driver for the posted display adapter."}
	ErrorGraphicsOpmNotSupported                                 = &Error{0xC0262500, "ERROR_GRAPHICS_OPM_NOT_SUPPORTED", "The driver does not support Output Protection Manager (OPM)."}
	ErrorGraphicsCoppNotSupported                                = &Error{0xC0262501, "ERROR_GRAPHICS_COPP_NOT_SUPPORTED", "The driver does not support Certified Output Protection Protocol (COPP)."}
	ErrorGraphicsUabNotSupported                                 = &Error{0xC0262502, "ERROR_GRAPHICS_UAB_NOT_SUPPORTED", "The driver does not support a user-accessible bus (UAB)."}
	ErrorGraphicsOpmInvalidEncryptedParameters                   = &Error{0xC0262503, "ERROR_GRAPHICS_OPM_INVALID_ENCRYPTED_PARAMETERS", "The specified encrypted parameters are invalid."}
	ErrorGraphicsOpmParameterArrayTooSmall                       = &Error{0xC0262504, "ERROR_GRAPHICS_OPM_PARAMETER_ARRAY_TOO_SMALL", "An array passed to a function cannot hold all of the data that the function wants to put in it."}
	ErrorGraphicsOpmNoVideoOutputsExist                          = &Error{0xC0262505, "ERROR_GRAPHICS_OPM_NO_VIDEO_OUTPUTS_EXIST", "The GDI display device passed to this function does not have any active video outputs."}
	ErrorGraphicsPvpNoDisplayDeviceCorrespondsToName             = &Error{0xC0262506, "ERROR_GRAPHICS_PVP_NO_DISPLAY_DEVICE_CORRESPONDS_TO_NAME", "The protected video path (PVP) cannot find an actual GDI display device that corresponds to the passed-in GDI display device name."}
	ErrorGraphicsPvpDisplayDeviceNotAttachedToDesktop            = &Error{0xC0262507, "ERROR_GRAPHICS_PVP_DISPLAY_DEVICE_NOT_ATTACHED_TO_DESKTOP", "This function failed because the GDI display device passed to it was not attached to the Windows desktop."}
	ErrorGraphicsPvpMirroringDevicesNotSupported                 = &Error{0xC0262508, "ERROR_GRAPHICS_PVP_MIRRORING_DEVICES_NOT_SUPPORTED", "The PVP does not support mirroring display devices because they do not have video outputs."}
	ErrorGraphicsOpmInvalidPointer                               = &Error{0xC026250A, "ERROR_GRAPHICS_OPM_INVALID_POINTER", "The function failed because an invalid pointer parameter was passed to it. A pointer parameter is invalid if it is null, it points to an invalid address, it points to a kernel mode address, or it is not correctly aligned."}
	ErrorGraphicsOpmInternalError                                = &Error{0xC026250B, "ERROR_GRAPHICS_OPM_INTERNAL_ERROR", "An internal error caused this operation to fail."}
	ErrorGraphicsOpmInvalidHandle                                = &Error{0xC026250C, "ERROR_GRAPHICS_OPM_INVALID_HANDLE", "The function failed because the caller passed in an invalid OPM user mode handle."}
	ErrorGraphicsPvpNoMonitorsCorrespondToDisplayDevice          = &Error{0xC026250D, "ERROR_GRAPHICS_PVP_NO_MONITORS_CORRESPOND_TO_DISPLAY_DEVICE", "This function failed because the GDI device passed to it did not have any monitors associated with it."}
	ErrorGraphicsPvpInvalidCertificateLength                     = &Error{0xC026250E, "ERROR_GRAPHICS_PVP_INVALID_CERTIFICATE_LENGTH", "A certificate could not be returned because the certificate buffer passed to the function was too small."}
	ErrorGraphicsOpmSpanningModeEnabled                          = &Error{0xC026250F, "ERROR_GRAPHICS_OPM_SPANNING_MODE_ENABLED", "A video output could not be created because the frame buffer is in spanning mode."}
	ErrorGraphicsOpmTheaterModeEnabled                           = &Error{0xC0262510, "ERROR_GRAPHICS_OPM_THEATER_MODE_ENABLED", "A video output could not be created because the frame buffer is in theater mode."}
	ErrorGraphicsPvpHfsFailed                                    = &Error{0xC0262511, "ERROR_GRAPHICS_PVP_HFS_FAILED", "The function call failed because the display adapter's hardware functionality scan failed to validate the graphics hardware."}
	ErrorGraphicsOpmInvalidSrm                                   = &Error{0xC0262512, "ERROR_GRAPHICS_OPM_INVALID_SRM", "The High-Bandwidth Digital Content Protection (HDCP) System Renewability Message (SRM) passed to this function did not comply with section 5 of the HDCP 1.1 specification."}
	ErrorGraphicsOpmOutputDoesNotSupportHdcp                     = &Error{0xC0262513, "ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_HDCP", "The video output cannot enable the HDCP system because it does not support it."}
	ErrorGraphicsOpmOutputDoesNotSupportAcp                      = &Error{0xC0262514, "ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_ACP", "The video output cannot enable analog copy protection because it does not support it."}
	ErrorGraphicsOpmOutputDoesNotSupportCgmsa                    = &Error{0xC0262515, "ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_CGMSA", "The video output cannot enable the Content Generation Management System Analog (CGMS-A) protection technology because it does not support it."}
	ErrorGraphicsOpmHdcpSrmNeverSet                              = &Error{0xC0262516, "ERROR_GRAPHICS_OPM_HDCP_SRM_NEVER_SET", "IOPMVideoOutput's GetInformation() method cannot return the version of the SRM being used because the application never successfully passed an SRM to the video output."}
	ErrorGraphicsOpmResolutionTooHigh                            = &Error{0xC0262517, "ERROR_GRAPHICS_OPM_RESOLUTION_TOO_HIGH", "IOPMVideoOutput's Configure() method cannot enable the specified output protection technology because the output's screen resolution is too high."}
	ErrorGraphicsOpmAllHdcpHardwareAlreadyInUse                  = &Error{0xC0262518, "ERROR_GRAPHICS_OPM_ALL_HDCP_HARDWARE_ALREADY_IN_USE", "IOPMVideoOutput's Configure() method cannot enable HDCP because the display adapter's HDCP hardware is already being used by other physical outputs."}
	ErrorGraphicsOpmVideoOutputNoLongerExists                    = &Error{0xC0262519, "ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_NO_LONGER_EXISTS", "The operating system asynchronously destroyed this OPM video output because the operating system's state changed. This error typically occurs because the monitor physical device object (PDO) associated with this video output was removed, the monitor PDO associated with this video output was stopped, the video output's session became a nonconsole session or the video output's desktop became an inactive desktop."}
	ErrorGraphicsOpmSessionTypeChangeInProgress                  = &Error{0xC026251A, "ERROR_GRAPHICS_OPM_SESSION_TYPE_CHANGE_IN_PROGRESS", "IOPMVideoOutput's methods cannot be called when a session is changing its type. There are currently three types of sessions: console, disconnected and remote (remote desktop protocol [RDP] or Independent Computing Architecture [ICA])."}
	ErrorGraphicsI2cNotSupported                                 = &Error{0xC0262580, "ERROR_GRAPHICS_I2C_NOT_SUPPORTED", "The monitor connected to the specified video output does not have an I2C bus."}
	ErrorGraphicsI2cDeviceDoesNotExist                           = &Error{0xC0262581, "ERROR_GRAPHICS_I2C_DEVICE_DOES_NOT_EXIST", "No device on the I2C bus has the specified address."}
	ErrorGraphicsI2cErrorTransmittingData                        = &Error{0xC0262582, "ERROR_GRAPHICS_I2C_ERROR_TRANSMITTING_DATA", "An error occurred while transmitting data to the device on the I2C bus."}
	ErrorGraphicsI2cErrorReceivingData                           = &Error{0xC0262583, "ERROR_GRAPHICS_I2C_ERROR_RECEIVING_DATA", "An error occurred while receiving data from the device on the I2C bus."}
	ErrorGraphicsDdcciVcpNotSupported                            = &Error{0xC0262584, "ERROR_GRAPHICS_DDCCI_VCP_NOT_SUPPORTED", "The monitor does not support the specified Virtual Control Panel (VCP) code."}
	ErrorGraphicsDdcciInvalidData                                = &Error{0xC0262585, "ERROR_GRAPHICS_DDCCI_INVALID_DATA", "The data received from the monitor is invalid."}
	ErrorGraphicsDdcciMonitorReturnedInvalidTimingStatusByte     = &Error{0xC0262586, "ERROR_GRAPHICS_DDCCI_MONITOR_RETURNED_INVALID_TIMING_STATUS_BYTE", "A function call failed because a monitor returned an invalid Timing Status byte when the operating system used the Display Data Channel Command Interface (DDC/CI) Get Timing Report and Timing Message command to get a timing report from a monitor."}
	ErrorGraphicsMcaInvalidCapabilitiesString                    = &Error{0xC0262587, "ERROR_GRAPHICS_MCA_INVALID_CAPABILITIES_STRING", "The monitor returned a DDC/CI capabilities string that did not comply with the ACCESS.bus 3.0, DDC/CI 1.1 or MCCS 2 Revision 1 specification."}
	ErrorGraphicsMcaInternalError                                = &Error{0xC0262588, "ERROR_GRAPHICS_MCA_INTERNAL_ERROR", "An internal Monitor Configuration API error occurred."}
	ErrorGraphicsDdcciInvalidMessageCommand                      = &Error{0xC0262589, "ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_COMMAND", "An operation failed because a DDC/CI message had an invalid value in its command field."}
	ErrorGraphicsDdcciInvalidMessageLength                       = &Error{0xC026258A, "ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_LENGTH", "This error occurred because a DDC/CI message length field contained an invalid value."}
	ErrorGraphicsDdcciInvalidMessageChecksum                     = &Error{0xC026258B, "ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_CHECKSUM", "This error occurred because the value in a DDC/CI message checksum field did not match the message's computed checksum value. This error implies that the data was corrupted while it was being transmitted from a monitor to a computer."}
	ErrorGraphicsPmeaInvalidMonitor                              = &Error{0xC02625D6, "ERROR_GRAPHICS_PMEA_INVALID_MONITOR", "The HMONITOR no longer exists, is not attached to the desktop, or corresponds to a mirroring device."}
	ErrorGraphicsPmeaInvalidD3dDevice                            = &Error{0xC02625D7, "ERROR_GRAPHICS_PMEA_INVALID_D3D_DEVICE", "The Direct3D (D3D) device's GDI display device no longer exists, is not attached to the desktop, or is a mirroring display device."}
	ErrorGraphicsDdcciCurrentCurrentValueGreaterThanMaximumValue = &Error{0xC02625D8, "ERROR_GRAPHICS_DDCCI_CURRENT_CURRENT_VALUE_GREATER_THAN_MAXIMUM_VALUE", "A continuous VCP code's current value is greater than its maximum value. This error code indicates that a monitor returned an invalid value."}
	ErrorGraphicsMcaInvalidVcpVersion                            = &Error{0xC02625D9, "ERROR_GRAPHICS_MCA_INVALID_VCP_VERSION", "The monitor's VCP Version (0xDF) VCP code returned an invalid version value."}
	ErrorGraphicsMcaMonitorViolatesMccsSpecification             = &Error{0xC02625DA, "ERROR_GRAPHICS_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION", "The monitor does not comply with the Monitor Control Command Set (MCCS) specification it claims to support."}
	ErrorGraphicsMcaMccsVersionMismatch                          = &Error{0xC02625DB, "ERROR_GRAPHICS_MCA_MCCS_VERSION_MISMATCH", "The MCCS version in a monitor's mccs_ver capability does not match the MCCS version the monitor reports when the VCP Version (0xDF) VCP code is used."}
	ErrorGraphicsMcaUnsupportedMccsVersion                       = &Error{0xC02625DC, "ERROR_GRAPHICS_MCA_UNSUPPORTED_MCCS_VERSION", "The Monitor Configuration API only works with monitors that support the MCCS 1.0 specification, the MCCS 2.0 specification, or the MCCS 2.0 Revision 1 specification."}
	ErrorGraphicsMcaInvalidTechnologyTypeReturned                = &Error{0xC02625DE, "ERROR_GRAPHICS_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED", "The monitor returned an invalid monitor technology type. CRT, plasma, and LCD (TFT) are examples of monitor technology types. This error implies that the monitor violated the MCCS 2.0 or MCCS 2.0 Revision 1 specification."}
	ErrorGraphicsMcaUnsupportedColorTemperature                  = &Error{0xC02625DF, "ERROR_GRAPHICS_MCA_UNSUPPORTED_COLOR_TEMPERATURE", "The SetMonitorColorTemperature() caller passed a color temperature to it that the current monitor did not support. CRT, plasma, and LCD (TFT) are examples of monitor technology types. This error implies that the monitor violated the MCCS 2.0 or MCCS 2.0 Revision 1 specification."}
	ErrorGraphicsOnlyConsoleSessionSupported                     = &Error{0xC02625E0, "ERROR_GRAPHICS_ONLY_CONSOLE_SESSION_SUPPORTED", "This function can be used only if a program is running in the local console session. It cannot be used if the program is running on a remote desktop session or on a terminal server session."}
	InetEAuthenticationRequired                                  = &Error{0x800C0009, "INET_E_AUTHENTICATION_REQUIRED", "Authentication is needed to access the object."}
	InetEBlockedRedirectXsecurityid                              = &Error{0x800C001B, "INET_E_BLOCKED_REDIRECT_XSECURITYID", "Internet Explorer\u00a08. A redirect request was blocked because SIDs do not match and BINDF2_DISABLE_HTTP_REDIRECT_XSECURITYID is set in bind options."}
	InetECannotConnect                                           = &Error{0x800C0004, "INET_E_CANNOT_CONNECT", "The attempt to connect to the Internet has failed."}
	InetECannotInstantiateObject                                 = &Error{0x800C0010, "INET_E_CANNOT_INSTANTIATE_OBJECT", "CoCreateInstance failed."}
	InetECannotLoadData                                          = &Error{0x800C000F, "INET_E_CANNOT_LOAD_DATA", "The object could not be loaded."}
	InetECannotLockRequest                                       = &Error{0x800C0016, "INET_E_CANNOT_LOCK_REQUEST", "The requested resource could not be locked."}
	InetECannotReplaceSfpFile                                    = &Error{0x800C0300, "INET_E_CANNOT_REPLACE_SFP_FILE", "Cannot replace a file that is protected by SFP."}
	InetECodeDownloadDeclined                                    = &Error{0x800C0100, "INET_E_CODE_DOWNLOAD_DECLINED", "The component download was declined by the user."}
	InetECodeInstallSuppressed                                   = &Error{0x800C0400, "INET_E_CODE_INSTALL_SUPPRESSED", "Internet Explorer\u00a06 for Windows\u00a0XP\u00a0SP2 and later. The Authenticode prompt for installing a ActiveX control was not shown because the page restricts the installation of the ActiveX controls. The usual cause is that the Information Bar is shown instead of the Authenticode prompt."}
	InetECodeInstallBlockedByHashPolicy                          = &Error{0x800C0500, "INET_E_CODE_INSTALL_BLOCKED_BY_HASH_POLICY", "Internet Explorer\u00a06 for Windows\u00a0XP\u00a0SP2 and later. Installation of ActiveX control (as identified by cryptographic file hash) has been disallowed by registry key policy."}
	InetEConnectionTimeout                                       = &Error{0x800C000B, "INET_E_CONNECTION_TIMEOUT", "The Internet connection has timed out."}
	InetEDataNotAvailable                                        = &Error{0x800C0007, "INET_E_DATA_NOT_AVAILABLE", "An Internet connection was established, but the data cannot be retrieved."}
	InetEDefaultAction                                           = &Error{0x800C0011, "INET_E_DEFAULT_ACTION", "Use the default security manager for this action. A custom security manager should only process input that is both valid and specific to itself and return INET_E_DEFAULT_ACTION for all other methods or URL actions."}
	InetEDownloadBlockedByInprivate                              = &Error{0x800C0501, "INET_E_DOWNLOAD_BLOCKED_BY_INPRIVATE", "Internet Explorer\u00a08 and later. The download was not permitted due to a private browsing session. InPrivate Browsing prevents Internet Explorer from storing data about the browsing session, such as cookies, temporary files, and history."}
	InetEDownloadFailure                                         = &Error{0x800C0008, "INET_E_DOWNLOAD_FAILURE", "The download has failed (the connection was interrupted)."}
	InetEErrorFirst                                              = &Error{0x800C0002, "INET_E_ERROR_FIRST", "The lowest allowed value for an URL moniker error code."}
	InetEErrorLast                                               = &Error{0x800C0501, "INET_E_ERROR_LAST", "The highest allowed value for an URL moniker error code."}
	InetEInvalidCertificate                                      = &Error{0x800C0019, "INET_E_INVALID_CERTIFICATE", "The SSL certificate is invalid."}
	InetEInvalidRequest                                          = &Error{0x800C000C, "INET_E_INVALID_REQUEST", "The request was invalid."}
	InetEInvalidUrl                                              = &Error{0x800C0002, "INET_E_INVALID_URL", "The URL could not be parsed."}
	InetENoSession                                               = &Error{0x800C0003, "INET_E_NO_SESSION", "No Internet session was established."}
	InetENoValidMedia                                            = &Error{0x800C000A, "INET_E_NO_VALID_MEDIA", "The object is not in one of the acceptable MIME types."}
	InetEObjectNotFound                                          = &Error{0x800C0006, "INET_E_OBJECT_NOT_FOUND", "The object was not found."}
	InetEQueryoptionUnknown                                      = &Error{0x800C0013, "INET_E_QUERYOPTION_UNKNOWN", "The requested option is unknown. (See IInternetProtocolInfo::QueryInfo.)"}
	InetERedirectFailed                                          = &Error{0x800C0014, "INET_E_REDIRECT_FAILED", "WinInet cannot redirect. This error code might also be returned by a custom protocol handler."}
	InetERedirectToDir                                           = &Error{0x800C0015, "INET_E_REDIRECT_TO_DIR", "The request is being redirected to a directory."}
	InetERedirecting                                             = &Error{0x800C0014, "INET_E_REDIRECTING", "The request is being redirected. (Pass this value to IInternetProtocolSink::ReportResult.)"}
	InetEReserved1                                               = &Error{0x800C001A, "INET_E_RESERVED_1", "Internet Explorer\u00a08. Reserved. Do not use."}
	InetEResourceNotFound                                        = &Error{0x800C0005, "INET_E_RESOURCE_NOT_FOUND", "The server or proxy was not found."}
	InetEResultDispatched                                        = &Error{0x800C0200, "INET_E_RESULT_DISPATCHED", "The binding has already been completed and the result has been dispatched, so your abort call has been canceled."}
	InetESecurityProblem                                         = &Error{0x800C000E, "INET_E_SECURITY_PROBLEM", "A security problem was encountered, related to one of the following Win32\u00a0Error Messages:"}
	InetETerminatedBind                                          = &Error{0x800C0018, "INET_E_TERMINATED_BIND", "Binding was terminated. (See IBinding::GetBindResult.)"}
	InetEUnknownProtocol                                         = &Error{0x800C000D, "INET_E_UNKNOWN_PROTOCOL", "The protocol is not known and no pluggable protocols have been entered that match."}
	InetEUseDefaultProtocolhandler                               = &Error{0x800C0011, "INET_E_USE_DEFAULT_PROTOCOLHANDLER", "Use the default protocol handler. (See IInternetProtocolRoot::Start.)"}
	InetEUseDefaultSetting                                       = &Error{0x800C0012, "INET_E_USE_DEFAULT_SETTING", "Use the default settings. (See IInternetBindInfo::GetBindString.)"}
	InetEUseExtendBinding                                        = &Error{0x800C0017, "INET_E_USE_EXTEND_BINDING", "(Microsoft internal.) Reissue request with extended binding."}
)

func FromCode(code uint32) error {

	if code == 0 {
		return nil
	}

	switch code {
	case 0x00030200:
		return StgSConverted
	case 0x00030201:
		return StgSBlock
	case 0x00030202:
		return StgSRetrynow
	case 0x00030203:
		return StgSMonitoring
	case 0x00030204:
		return StgSMultipleopens
	case 0x00030205:
		return StgSConsolidationfailed
	case 0x00030206:
		return StgSCannotconsolidate
	case 0x00040000:
		return OleSUsereg
	case 0x00040001:
		return OleSStatic
	case 0x00040002:
		return OleSMacClipformat
	case 0x00040100:
		return DragdropSDrop
	case 0x00040101:
		return DragdropSCancel
	case 0x00040102:
		return DragdropSUsedefaultcursors
	case 0x00040130:
		return DataSSameformatetc
	case 0x00040140:
		return ViewSAlreadyFrozen
	case 0x00040170:
		return CacheSFormatetcNotsupported
	case 0x00040171:
		return CacheSSamecache
	case 0x00040172:
		return CacheSSomecachesNotupdated
	case 0x00040180:
		return OleobjSInvalidverb
	case 0x00040181:
		return OleobjSCannotDoverbNow
	case 0x00040182:
		return OleobjSInvalidhwnd
	case 0x000401A0:
		return InplaceSTruncated
	case 0x000401C0:
		return Convert10SNoPresentation
	case 0x000401E2:
		return MkSReducedToSelf
	case 0x000401E4:
		return MkSMe
	case 0x000401E5:
		return MkSHim
	case 0x000401E6:
		return MkSUs
	case 0x000401E7:
		return MkSMonikeralreadyregistered
	case 0x00040200:
		return EventSSomeSubscribersFailed
	case 0x00040202:
		return EventSNosubscribers
	case 0x00041300:
		return SchedSTaskReady
	case 0x00041301:
		return SchedSTaskRunning
	case 0x00041302:
		return SchedSTaskDisabled
	case 0x00041303:
		return SchedSTaskHasNotRun
	case 0x00041304:
		return SchedSTaskNoMoreRuns
	case 0x00041305:
		return SchedSTaskNotScheduled
	case 0x00041306:
		return SchedSTaskTerminated
	case 0x00041307:
		return SchedSTaskNoValidTriggers
	case 0x00041308:
		return SchedSEventTrigger
	case 0x0004131B:
		return SchedSSomeTriggersFailed
	case 0x0004131C:
		return SchedSBatchLogonProblem
	case 0x0004D000:
		return XactSAsync
	case 0x0004D002:
		return XactSReadonly
	case 0x0004D003:
		return XactSSomenoretain
	case 0x0004D004:
		return XactSOkinform
	case 0x0004D005:
		return XactSMadechangescontent
	case 0x0004D006:
		return XactSMadechangesinform
	case 0x0004D007:
		return XactSAllnoretain
	case 0x0004D008:
		return XactSAborting
	case 0x0004D009:
		return XactSSinglephase
	case 0x0004D00A:
		return XactSLocallyOk
	case 0x0004D010:
		return XactSLastresourcemanager
	case 0x00080012:
		return CoSNotallinterfaces
	case 0x00080013:
		return CoSMachinenamenotfound
	case 0x00090312:
		return SecIContinueNeeded
	case 0x00090313:
		return SecICompleteNeeded
	case 0x00090314:
		return SecICompleteAndContinue
	case 0x00090315:
		return SecILocalLogon
	case 0x00090317:
		return SecIContextExpired
	case 0x00090320:
		return SecIIncompleteCredentials
	case 0x00090321:
		return SecIRenegotiate
	case 0x00090323:
		return SecINoLsaContext
	case 0x0009035C:
		return SecISignatureNeeded
	case 0x00091012:
		return CryptINewProtectionRequired
	case 0x000D0000:
		return NsSCallpending
	case 0x000D0001:
		return NsSCallaborted
	case 0x000D0002:
		return NsSStreamTruncated
	case 0x000D0BC8:
		return NsSRebuffering
	case 0x000D0BC9:
		return NsSDegradingQuality
	case 0x000D0BDB:
		return NsSTranscryptorEof
	case 0x000D0FE8:
		return NsSWmpUiVersionmismatch
	case 0x000D0FE9:
		return NsSWmpException
	case 0x000D1040:
		return NsSWmpLoadedGifImage
	case 0x000D1041:
		return NsSWmpLoadedPngImage
	case 0x000D1042:
		return NsSWmpLoadedBmpImage
	case 0x000D1043:
		return NsSWmpLoadedJpgImage
	case 0x000D104F:
		return NsSWmgForceDropFrame
	case 0x000D105F:
		return NsSWmrAlreadyrendered
	case 0x000D1060:
		return NsSWmrPintypepartialmatch
	case 0x000D1061:
		return NsSWmrPintypefullmatch
	case 0x000D1066:
		return NsSWmgAdviseDropFrame
	case 0x000D1067:
		return NsSWmgAdviseDropToKeyframe
	case 0x000D10DB:
		return NsSNeedToBuyBurnRights
	case 0x000D10FE:
		return NsSWmpcorePlaylistclearabort
	case 0x000D10FF:
		return NsSWmpcorePlaylistremoveitemabort
	case 0x000D1102:
		return NsSWmpcorePlaylistCreationPending
	case 0x000D1103:
		return NsSWmpcoreMediaValidationPending
	case 0x000D1104:
		return NsSWmpcorePlaylistRepeatSecondarySegmentsIgnored
	case 0x000D1105:
		return NsSWmpcoreCommandNotAvailable
	case 0x000D1106:
		return NsSWmpcorePlaylistNameAutoGenerated
	case 0x000D1107:
		return NsSWmpcorePlaylistImportMissingItems
	case 0x000D1108:
		return NsSWmpcorePlaylistCollapsedToSingleMedia
	case 0x000D1109:
		return NsSWmpcoreMediaChildPlaylistOpenPending
	case 0x000D110A:
		return NsSWmpcoreMoreNodesAvaiable
	case 0x000D1135:
		return NsSWmpbrSuccess
	case 0x000D1136:
		return NsSWmpbrPartialsuccess
	case 0x000D1144:
		return NsSWmpeffectTransparent
	case 0x000D1145:
		return NsSWmpeffectOpaque
	case 0x000D114E:
		return NsSOperationPending
	case 0x000D1359:
		return NsSTrackBuyRequiresAlbumPurchase
	case 0x000D135E:
		return NsSNavigationCompleteWithErrors
	case 0x000D1361:
		return NsSTrackAlreadyDownloaded
	case 0x000D1519:
		return NsSPublishingPointStartedWithFailedSinks
	case 0x000D2726:
		return NsSDrmLicenseAcquired
	case 0x000D2727:
		return NsSDrmIndividualized
	case 0x000D2746:
		return NsSDrmMonitorCancelled
	case 0x000D2747:
		return NsSDrmAcquireCancelled
	case 0x000D276E:
		return NsSDrmBurnableTrack
	case 0x000D276F:
		return NsSDrmBurnableTrackWithPlaylistRestriction
	case 0x000D27DE:
		return NsSDrmNeedsIndividualization
	case 0x000D2AF8:
		return NsSRebootRecommended
	case 0x000D2AF9:
		return NsSRebootRequired
	case 0x000D2F09:
		return NsSEosreceding
	case 0x000D2F0D:
		return NsSChangenotice
	case 0x001F0001:
		return ErrorFltIoComplete
	case 0x00262307:
		return ErrorGraphicsModeNotPinned
	case 0x0026231E:
		return ErrorGraphicsNoPreferredMode
	case 0x0026234B:
		return ErrorGraphicsDatasetIsEmpty
	case 0x0026234C:
		return ErrorGraphicsNoMoreElementsInDataset
	case 0x00262351:
		return ErrorGraphicsPathContentGeometryTransformationNotPinned
	case 0x00300100:
		return PlaSPropertyIgnored
	case 0x00340001:
		return ErrorNdisIndicationRequired
	case 0x0DEAD100:
		return TrkSOutOfSync
	case 0x0DEAD102:
		return TrkVolumeNotFound
	case 0x0DEAD103:
		return TrkVolumeNotOwned
	case 0x0DEAD107:
		return TrkSNotificationQuotaExceeded
	case 0x400D004F:
		return NsITigerStart
	case 0x400D0051:
		return NsICubStart
	case 0x400D0052:
		return NsICubRunning
	case 0x400D0054:
		return NsIDiskStart
	case 0x400D0056:
		return NsIDiskRebuildStarted
	case 0x400D0057:
		return NsIDiskRebuildFinished
	case 0x400D0058:
		return NsIDiskRebuildAborted
	case 0x400D0059:
		return NsILimitFunnels
	case 0x400D005A:
		return NsIStartDisk
	case 0x400D005B:
		return NsIStopDisk
	case 0x400D005C:
		return NsIStopCub
	case 0x400D005D:
		return NsIKillUsersession
	case 0x400D005E:
		return NsIKillConnection
	case 0x400D005F:
		return NsIRebuildDisk
	case 0x400D0069:
		return McmadmINoEvents
	case 0x400D006E:
		return NsILoggingFailed
	case 0x400D0070:
		return NsILimitBandwidth
	case 0x400D0191:
		return NsICubUnfailLink
	case 0x400D0193:
		return NsIRestripeStart
	case 0x400D0194:
		return NsIRestripeDone
	case 0x400D0196:
		return NsIRestripeDiskOut
	case 0x400D0197:
		return NsIRestripeCubOut
	case 0x400D0198:
		return NsIDiskStop
	case 0x400D14BE:
		return NsIPlaylistChangeReceding
	case 0x400D2EFF:
		return NsIReconnected
	case 0x400D2F01:
		return NsINologStop
	case 0x400D2F03:
		return NsIExistingPacketizer
	case 0x400D2F04:
		return NsIManualProxy
	case 0x40262009:
		return ErrorGraphicsDriverMismatch
	case 0x4026242F:
		return ErrorGraphicsUnknownChildStatus
	case 0x40262437:
		return ErrorGraphicsLeadlinkStartDeferred
	case 0x40262439:
		return ErrorGraphicsPollingTooFrequently
	case 0x4026243A:
		return ErrorGraphicsStartDeferred
	case 0x8000000A:
		return EPending
	case 0x80004001:
		return ENotimpl
	case 0x80004002:
		return ENointerface
	case 0x80004003:
		return EPointer
	case 0x80004004:
		return EAbort
	case 0x80004005:
		return EFail
	case 0x80004006:
		return CoEInitTls
	case 0x80004007:
		return CoEInitSharedAllocator
	case 0x80004008:
		return CoEInitMemoryAllocator
	case 0x80004009:
		return CoEInitClassCache
	case 0x8000400A:
		return CoEInitRpcChannel
	case 0x8000400B:
		return CoEInitTlsSetChannelControl
	case 0x8000400C:
		return CoEInitTlsChannelControl
	case 0x8000400D:
		return CoEInitUnacceptedUserAllocator
	case 0x8000400E:
		return CoEInitScmMutexExists
	case 0x8000400F:
		return CoEInitScmFileMappingExists
	case 0x80004010:
		return CoEInitScmMapViewOfFile
	case 0x80004011:
		return CoEInitScmExecFailure
	case 0x80004012:
		return CoEInitOnlySingleThreaded
	case 0x80004013:
		return CoECantRemote
	case 0x80004014:
		return CoEBadServerName
	case 0x80004015:
		return CoEWrongServerIdentity
	case 0x80004016:
		return CoEOle1ddeDisabled
	case 0x80004017:
		return CoERunasSyntax
	case 0x80004018:
		return CoECreateprocessFailure
	case 0x80004019:
		return CoERunasCreateprocessFailure
	case 0x8000401A:
		return CoERunasLogonFailure
	case 0x8000401B:
		return CoELaunchPermssionDenied
	case 0x8000401C:
		return CoEStartServiceFailure
	case 0x8000401D:
		return CoERemoteCommunicationFailure
	case 0x8000401E:
		return CoEServerStartTimeout
	case 0x8000401F:
		return CoEClsregInconsistent
	case 0x80004020:
		return CoEIidregInconsistent
	case 0x80004021:
		return CoENotSupported
	case 0x80004022:
		return CoEReloadDll
	case 0x80004023:
		return CoEMsiError
	case 0x80004024:
		return CoEAttemptToCreateOutsideClientContext
	case 0x80004025:
		return CoEServerPaused
	case 0x80004026:
		return CoEServerNotPaused
	case 0x80004027:
		return CoEClassDisabled
	case 0x80004028:
		return CoEClrnotavailable
	case 0x80004029:
		return CoEAsyncWorkRejected
	case 0x8000402A:
		return CoEServerInitTimeout
	case 0x8000402B:
		return CoENoSecctxInActivate
	case 0x80004030:
		return CoETrackerConfig
	case 0x80004031:
		return CoEThreadpoolConfig
	case 0x80004032:
		return CoESxsConfig
	case 0x80004033:
		return CoEMalformedSpn
	case 0x8000FFFF:
		return EUnexpected
	case 0x80010001:
		return RpcECallRejected
	case 0x80010002:
		return RpcECallCanceled
	case 0x80010003:
		return RpcECantpostInsendcall
	case 0x80010004:
		return RpcECantcalloutInasynccall
	case 0x80010005:
		return RpcECantcalloutInexternalcall
	case 0x80010006:
		return RpcEConnectionTerminated
	case 0x80010007:
		return RpcEServerDied
	case 0x80010008:
		return RpcEClientDied
	case 0x80010009:
		return RpcEInvalidDatapacket
	case 0x8001000A:
		return RpcECanttransmitCall
	case 0x8001000B:
		return RpcEClientCantmarshalData
	case 0x8001000C:
		return RpcEClientCantunmarshalData
	case 0x8001000D:
		return RpcEServerCantmarshalData
	case 0x8001000E:
		return RpcEServerCantunmarshalData
	case 0x8001000F:
		return RpcEInvalidData
	case 0x80010010:
		return RpcEInvalidParameter
	case 0x80010011:
		return RpcECantcalloutAgain
	case 0x80010012:
		return RpcEServerDiedDne
	case 0x80010100:
		return RpcESysCallFailed
	case 0x80010101:
		return RpcEOutOfResources
	case 0x80010102:
		return RpcEAttemptedMultithread
	case 0x80010103:
		return RpcENotRegistered
	case 0x80010104:
		return RpcEFault
	case 0x80010105:
		return RpcEServerfault
	case 0x80010106:
		return RpcEChangedMode
	case 0x80010107:
		return RpcEInvalidmethod
	case 0x80010108:
		return RpcEDisconnected
	case 0x80010109:
		return RpcERetry
	case 0x8001010A:
		return RpcEServercallRetrylater
	case 0x8001010B:
		return RpcEServercallRejected
	case 0x8001010C:
		return RpcEInvalidCalldata
	case 0x8001010D:
		return RpcECantcalloutIninputsynccall
	case 0x8001010E:
		return RpcEWrongThread
	case 0x8001010F:
		return RpcEThreadNotInit
	case 0x80010110:
		return RpcEVersionMismatch
	case 0x80010111:
		return RpcEInvalidHeader
	case 0x80010112:
		return RpcEInvalidExtension
	case 0x80010113:
		return RpcEInvalidIpid
	case 0x80010114:
		return RpcEInvalidObject
	case 0x80010115:
		return RpcSCallpending
	case 0x80010116:
		return RpcSWaitontimer
	case 0x80010117:
		return RpcECallComplete
	case 0x80010118:
		return RpcEUnsecureCall
	case 0x80010119:
		return RpcETooLate
	case 0x8001011A:
		return RpcENoGoodSecurityPackages
	case 0x8001011B:
		return RpcEAccessDenied
	case 0x8001011C:
		return RpcERemoteDisabled
	case 0x8001011D:
		return RpcEInvalidObjref
	case 0x8001011E:
		return RpcENoContext
	case 0x8001011F:
		return RpcETimeout
	case 0x80010120:
		return RpcENoSync
	case 0x80010121:
		return RpcEFullsicRequired
	case 0x80010122:
		return RpcEInvalidStdName
	case 0x80010123:
		return CoEFailedtoimpersonate
	case 0x80010124:
		return CoEFailedtogetsecctx
	case 0x80010125:
		return CoEFailedtoopenthreadtoken
	case 0x80010126:
		return CoEFailedtogettokeninfo
	case 0x80010127:
		return CoETrusteedoesntmatchclient
	case 0x80010128:
		return CoEFailedtoqueryclientblanket
	case 0x80010129:
		return CoEFailedtosetdacl
	case 0x8001012A:
		return CoEAccesscheckfailed
	case 0x8001012B:
		return CoENetaccessapifailed
	case 0x8001012C:
		return CoEWrongtrusteenamesyntax
	case 0x8001012D:
		return CoEInvalidsid
	case 0x8001012E:
		return CoEConversionfailed
	case 0x8001012F:
		return CoENomatchingsidfound
	case 0x80010130:
		return CoELookupaccsidfailed
	case 0x80010131:
		return CoENomatchingnamefound
	case 0x80010132:
		return CoELookupaccnamefailed
	case 0x80010133:
		return CoESetserlhndlfailed
	case 0x80010134:
		return CoEFailedtogetwindir
	case 0x80010135:
		return CoEPathtoolong
	case 0x80010136:
		return CoEFailedtogenuuid
	case 0x80010137:
		return CoEFailedtocreatefile
	case 0x80010138:
		return CoEFailedtoclosehandle
	case 0x80010139:
		return CoEExceedsysacllimit
	case 0x8001013A:
		return CoEAcesinwrongorder
	case 0x8001013B:
		return CoEIncompatiblestreamversion
	case 0x8001013C:
		return CoEFailedtoopenprocesstoken
	case 0x8001013D:
		return CoEDecodefailed
	case 0x8001013F:
		return CoEAcnotinitialized
	case 0x80010140:
		return CoECancelDisabled
	case 0x8001FFFF:
		return RpcEUnexpected
	case 0x80020001:
		return DispEUnknowninterface
	case 0x80020003:
		return DispEMembernotfound
	case 0x80020004:
		return DispEParamnotfound
	case 0x80020005:
		return DispETypemismatch
	case 0x80020006:
		return DispEUnknownname
	case 0x80020007:
		return DispENonamedargs
	case 0x80020008:
		return DispEBadvartype
	case 0x80020009:
		return DispEException
	case 0x8002000A:
		return DispEOverflow
	case 0x8002000B:
		return DispEBadindex
	case 0x8002000C:
		return DispEUnknownlcid
	case 0x8002000D:
		return DispEArrayislocked
	case 0x8002000E:
		return DispEBadparamcount
	case 0x8002000F:
		return DispEParamnotoptional
	case 0x80020010:
		return DispEBadcallee
	case 0x80020011:
		return DispENotacollection
	case 0x80020012:
		return DispEDivbyzero
	case 0x80020013:
		return DispEBuffertoosmall
	case 0x80028016:
		return TypeEBuffertoosmall
	case 0x80028017:
		return TypeEFieldnotfound
	case 0x80028018:
		return TypeEInvdataread
	case 0x80028019:
		return TypeEUnsupformat
	case 0x8002801C:
		return TypeERegistryaccess
	case 0x8002801D:
		return TypeELibnotregistered
	case 0x80028027:
		return TypeEUndefinedtype
	case 0x80028028:
		return TypeEQualifiednamedisallowed
	case 0x80028029:
		return TypeEInvalidstate
	case 0x8002802A:
		return TypeEWrongtypekind
	case 0x8002802B:
		return TypeEElementnotfound
	case 0x8002802C:
		return TypeEAmbiguousname
	case 0x8002802D:
		return TypeENameconflict
	case 0x8002802E:
		return TypeEUnknownlcid
	case 0x8002802F:
		return TypeEDllfunctionnotfound
	case 0x800288BD:
		return TypeEBadmodulekind
	case 0x800288C5:
		return TypeESizetoobig
	case 0x800288C6:
		return TypeEDuplicateid
	case 0x800288CF:
		return TypeEInvalidid
	case 0x80028CA0:
		return TypeETypemismatch
	case 0x80028CA1:
		return TypeEOutofbounds
	case 0x80028CA2:
		return TypeEIoerror
	case 0x80028CA3:
		return TypeECantcreatetmpfile
	case 0x80029C4A:
		return TypeECantloadlibrary
	case 0x80029C83:
		return TypeEInconsistentpropfuncs
	case 0x80029C84:
		return TypeECirculartype
	case 0x80030001:
		return StgEInvalidfunction
	case 0x80030002:
		return StgEFilenotfound
	case 0x80030003:
		return StgEPathnotfound
	case 0x80030004:
		return StgEToomanyopenfiles
	case 0x80030005:
		return StgEAccessdenied
	case 0x80030006:
		return StgEInvalidhandle
	case 0x80030008:
		return StgEInsufficientmemory
	case 0x80030009:
		return StgEInvalidpointer
	case 0x80030012:
		return StgENomorefiles
	case 0x80030013:
		return StgEDiskiswriteprotected
	case 0x80030019:
		return StgESeekerror
	case 0x8003001D:
		return StgEWritefault
	case 0x8003001E:
		return StgEReadfault
	case 0x80030020:
		return StgEShareviolation
	case 0x80030021:
		return StgELockviolation
	case 0x80030050:
		return StgEFilealreadyexists
	case 0x80030057:
		return StgEInvalidparameter
	case 0x80030070:
		return StgEMediumfull
	case 0x800300F0:
		return StgEPropsetmismatched
	case 0x800300FA:
		return StgEAbnormalapiexit
	case 0x800300FB:
		return StgEInvalidheader
	case 0x800300FC:
		return StgEInvalidname
	case 0x800300FD:
		return StgEUnknown
	case 0x800300FE:
		return StgEUnimplementedfunction
	case 0x800300FF:
		return StgEInvalidflag
	case 0x80030100:
		return StgEInuse
	case 0x80030101:
		return StgENotcurrent
	case 0x80030102:
		return StgEReverted
	case 0x80030103:
		return StgECantsave
	case 0x80030104:
		return StgEOldformat
	case 0x80030105:
		return StgEOlddll
	case 0x80030106:
		return StgESharerequired
	case 0x80030107:
		return StgENotfilebasedstorage
	case 0x80030108:
		return StgEExtantmarshallings
	case 0x80030109:
		return StgEDocfilecorrupt
	case 0x80030110:
		return StgEBadbaseaddress
	case 0x80030111:
		return StgEDocfiletoolarge
	case 0x80030112:
		return StgENotsimpleformat
	case 0x80030201:
		return StgEIncomplete
	case 0x80030202:
		return StgETerminated
	case 0x80030305:
		return StgEStatusCopyProtectionFailure
	case 0x80030306:
		return StgECssAuthenticationFailure
	case 0x80030307:
		return StgECssKeyNotPresent
	case 0x80030308:
		return StgECssKeyNotEstablished
	case 0x80030309:
		return StgECssScrambledSector
	case 0x8003030A:
		return StgECssRegionMismatch
	case 0x8003030B:
		return StgEResetsExhausted
	case 0x80040000:
		return OleEOleverb
	case 0x80040001:
		return OleEAdvf
	case 0x80040002:
		return OleEEnumNomore
	case 0x80040003:
		return OleEAdvisenotsupported
	case 0x80040004:
		return OleENoconnection
	case 0x80040005:
		return OleENotrunning
	case 0x80040006:
		return OleENocache
	case 0x80040007:
		return OleEBlank
	case 0x80040008:
		return OleEClassdiff
	case 0x80040009:
		return OleECantGetmoniker
	case 0x8004000A:
		return OleECantBindtosource
	case 0x8004000B:
		return OleEStatic
	case 0x8004000C:
		return OleEPromptsavecancelled
	case 0x8004000D:
		return OleEInvalidrect
	case 0x8004000E:
		return OleEWrongcompobj
	case 0x8004000F:
		return OleEInvalidhwnd
	case 0x80040010:
		return OleENotInplaceactive
	case 0x80040011:
		return OleECantconvert
	case 0x80040012:
		return OleENostorage
	case 0x80040064:
		return DvEFormatetc
	case 0x80040065:
		return DvEDvtargetdevice
	case 0x80040066:
		return DvEStgmedium
	case 0x80040067:
		return DvEStatdata
	case 0x80040068:
		return DvELindex
	case 0x80040069:
		return DvETymed
	case 0x8004006A:
		return DvEClipformat
	case 0x8004006B:
		return DvEDvaspect
	case 0x8004006C:
		return DvEDvtargetdeviceSize
	case 0x8004006D:
		return DvENoiviewobject
	case 0x80040100:
		return DragdropENotregistered
	case 0x80040101:
		return DragdropEAlreadyregistered
	case 0x80040102:
		return DragdropEInvalidhwnd
	case 0x80040110:
		return ClassENoaggregation
	case 0x80040111:
		return ClassEClassnotavailable
	case 0x80040112:
		return ClassENotlicensed
	case 0x80040140:
		return ViewEDraw
	case 0x80040150:
		return RegdbEReadregdb
	case 0x80040151:
		return RegdbEWriteregdb
	case 0x80040152:
		return RegdbEKeymissing
	case 0x80040153:
		return RegdbEInvalidvalue
	case 0x80040154:
		return RegdbEClassnotreg
	case 0x80040155:
		return RegdbEIidnotreg
	case 0x80040156:
		return RegdbEBadthreadingmodel
	case 0x80040160:
		return CatECatidnoexist
	case 0x80040161:
		return CatENodescription
	case 0x80040164:
		return CsEPackageNotfound
	case 0x80040165:
		return CsENotDeletable
	case 0x80040166:
		return CsEClassNotfound
	case 0x80040167:
		return CsEInvalidVersion
	case 0x80040168:
		return CsENoClassstore
	case 0x80040169:
		return CsEObjectNotfound
	case 0x8004016A:
		return CsEObjectAlreadyExists
	case 0x8004016B:
		return CsEInvalidPath
	case 0x8004016C:
		return CsENetworkError
	case 0x8004016D:
		return CsEAdminLimitExceeded
	case 0x8004016E:
		return CsESchemaMismatch
	case 0x8004016F:
		return CsEInternalError
	case 0x80040170:
		return CacheENocacheUpdated
	case 0x80040180:
		return OleobjENoverbs
	case 0x80040181:
		return OleobjEInvalidverb
	case 0x800401A0:
		return InplaceENotundoable
	case 0x800401A1:
		return InplaceENotoolspace
	case 0x800401C0:
		return Convert10EOlestreamGet
	case 0x800401C1:
		return Convert10EOlestreamPut
	case 0x800401C2:
		return Convert10EOlestreamFmt
	case 0x800401C3:
		return Convert10EOlestreamBitmapToDib
	case 0x800401C4:
		return Convert10EStgFmt
	case 0x800401C5:
		return Convert10EStgNoStdStream
	case 0x800401C6:
		return Convert10EStgDibToBitmap
	case 0x800401D0:
		return ClipbrdECantOpen
	case 0x800401D1:
		return ClipbrdECantEmpty
	case 0x800401D2:
		return ClipbrdECantSet
	case 0x800401D3:
		return ClipbrdEBadData
	case 0x800401D4:
		return ClipbrdECantClose
	case 0x800401E0:
		return MkEConnectmanually
	case 0x800401E1:
		return MkEExceededdeadline
	case 0x800401E2:
		return MkENeedgeneric
	case 0x800401E3:
		return MkEUnavailable
	case 0x800401E4:
		return MkESyntax
	case 0x800401E5:
		return MkENoobject
	case 0x800401E6:
		return MkEInvalidextension
	case 0x800401E7:
		return MkEIntermediateinterfacenotsupported
	case 0x800401E8:
		return MkENotbindable
	case 0x800401E9:
		return MkENotbound
	case 0x800401EA:
		return MkECantopenfile
	case 0x800401EB:
		return MkEMustbotheruser
	case 0x800401EC:
		return MkENoinverse
	case 0x800401ED:
		return MkENostorage
	case 0x800401EE:
		return MkENoprefix
	case 0x800401EF:
		return MkEEnumerationFailed
	case 0x800401F0:
		return CoENotinitialized
	case 0x800401F1:
		return CoEAlreadyinitialized
	case 0x800401F2:
		return CoECantdetermineclass
	case 0x800401F3:
		return CoEClassstring
	case 0x800401F4:
		return CoEIidstring
	case 0x800401F5:
		return CoEAppnotfound
	case 0x800401F6:
		return CoEAppsingleuse
	case 0x800401F7:
		return CoEErrorinapp
	case 0x800401F8:
		return CoEDllnotfound
	case 0x800401F9:
		return CoEErrorindll
	case 0x800401FA:
		return CoEWrongosforapp
	case 0x800401FB:
		return CoEObjnotreg
	case 0x800401FC:
		return CoEObjisreg
	case 0x800401FD:
		return CoEObjnotconnected
	case 0x800401FE:
		return CoEAppdidntreg
	case 0x800401FF:
		return CoEReleased
	case 0x80040201:
		return EventEAllSubscribersFailed
	case 0x80040203:
		return EventEQuerysyntax
	case 0x80040204:
		return EventEQueryfield
	case 0x80040205:
		return EventEInternalexception
	case 0x80040206:
		return EventEInternalerror
	case 0x80040207:
		return EventEInvalidPerUserSid
	case 0x80040208:
		return EventEUserException
	case 0x80040209:
		return EventETooManyMethods
	case 0x8004020A:
		return EventEMissingEventclass
	case 0x8004020B:
		return EventENotAllRemoved
	case 0x8004020C:
		return EventEComplusNotInstalled
	case 0x8004020D:
		return EventECantModifyOrDeleteUnconfiguredObject
	case 0x8004020E:
		return EventECantModifyOrDeleteConfiguredObject
	case 0x8004020F:
		return EventEInvalidEventClassPartition
	case 0x80040210:
		return EventEPerUserSidNotLoggedOn
	case 0x80041309:
		return SchedETriggerNotFound
	case 0x8004130A:
		return SchedETaskNotReady
	case 0x8004130B:
		return SchedETaskNotRunning
	case 0x8004130C:
		return SchedEServiceNotInstalled
	case 0x8004130D:
		return SchedECannotOpenTask
	case 0x8004130E:
		return SchedEInvalidTask
	case 0x8004130F:
		return SchedEAccountInformationNotSet
	case 0x80041310:
		return SchedEAccountNameNotFound
	case 0x80041311:
		return SchedEAccountDbaseCorrupt
	case 0x80041312:
		return SchedENoSecurityServices
	case 0x80041313:
		return SchedEUnknownObjectVersion
	case 0x80041314:
		return SchedEUnsupportedAccountOption
	case 0x80041315:
		return SchedEServiceNotRunning
	case 0x80041316:
		return SchedEUnexpectednode
	case 0x80041317:
		return SchedENamespace
	case 0x80041318:
		return SchedEInvalidvalue
	case 0x80041319:
		return SchedEMissingnode
	case 0x8004131A:
		return SchedEMalformedxml
	case 0x8004131D:
		return SchedETooManyNodes
	case 0x8004131E:
		return SchedEPastEndBoundary
	case 0x8004131F:
		return SchedEAlreadyRunning
	case 0x80041320:
		return SchedEUserNotLoggedOn
	case 0x80041321:
		return SchedEInvalidTaskHash
	case 0x80041322:
		return SchedEServiceNotAvailable
	case 0x80041323:
		return SchedEServiceTooBusy
	case 0x80041324:
		return SchedETaskAttempted
	case 0x8004D000:
		return XactEAlreadyothersinglephase
	case 0x8004D001:
		return XactECantretain
	case 0x8004D002:
		return XactECommitfailed
	case 0x8004D003:
		return XactECommitprevented
	case 0x8004D004:
		return XactEHeuristicabort
	case 0x8004D005:
		return XactEHeuristiccommit
	case 0x8004D006:
		return XactEHeuristicdamage
	case 0x8004D007:
		return XactEHeuristicdanger
	case 0x8004D008:
		return XactEIsolationlevel
	case 0x8004D009:
		return XactENoasync
	case 0x8004D00A:
		return XactENoenlist
	case 0x8004D00B:
		return XactENoisoretain
	case 0x8004D00C:
		return XactENoresource
	case 0x8004D00D:
		return XactENotcurrent
	case 0x8004D00E:
		return XactENotransaction
	case 0x8004D00F:
		return XactENotsupported
	case 0x8004D010:
		return XactEUnknownrmgrid
	case 0x8004D011:
		return XactEWrongstate
	case 0x8004D012:
		return XactEWronguow
	case 0x8004D013:
		return XactEXtionexists
	case 0x8004D014:
		return XactENoimportobject
	case 0x8004D015:
		return XactEInvalidcookie
	case 0x8004D016:
		return XactEIndoubt
	case 0x8004D017:
		return XactENotimeout
	case 0x8004D018:
		return XactEAlreadyinprogress
	case 0x8004D019:
		return XactEAborted
	case 0x8004D01A:
		return XactELogfull
	case 0x8004D01B:
		return XactETmnotavailable
	case 0x8004D01C:
		return XactEConnectionDown
	case 0x8004D01D:
		return XactEConnectionDenied
	case 0x8004D01E:
		return XactEReenlisttimeout
	case 0x8004D01F:
		return XactETipConnectFailed
	case 0x8004D020:
		return XactETipProtocolError
	case 0x8004D021:
		return XactETipPullFailed
	case 0x8004D022:
		return XactEDestTmnotavailable
	case 0x8004D023:
		return XactETipDisabled
	case 0x8004D024:
		return XactENetworkTxDisabled
	case 0x8004D025:
		return XactEPartnerNetworkTxDisabled
	case 0x8004D026:
		return XactEXaTxDisabled
	case 0x8004D027:
		return XactEUnableToReadDtcConfig
	case 0x8004D028:
		return XactEUnableToLoadDtcProxy
	case 0x8004D029:
		return XactEAborting
	case 0x8004D080:
		return XactEClerknotfound
	case 0x8004D081:
		return XactEClerkexists
	case 0x8004D082:
		return XactERecoveryinprogress
	case 0x8004D083:
		return XactETransactionclosed
	case 0x8004D084:
		return XactEInvalidlsn
	case 0x8004D085:
		return XactEReplayrequest
	case 0x8004D100:
		return XactEConnectionRequestDenied
	case 0x8004D101:
		return XactEToomanyEnlistments
	case 0x8004D102:
		return XactEDuplicateGuid
	case 0x8004D103:
		return XactENotsinglephase
	case 0x8004D104:
		return XactERecoveryalreadydone
	case 0x8004D105:
		return XactEProtocol
	case 0x8004D106:
		return XactERmFailure
	case 0x8004D107:
		return XactERecoveryFailed
	case 0x8004D108:
		return XactELuNotFound
	case 0x8004D109:
		return XactEDuplicateLu
	case 0x8004D10A:
		return XactELuNotConnected
	case 0x8004D10B:
		return XactEDuplicateTransid
	case 0x8004D10C:
		return XactELuBusy
	case 0x8004D10D:
		return XactELuNoRecoveryProcess
	case 0x8004D10E:
		return XactELuDown
	case 0x8004D10F:
		return XactELuRecovering
	case 0x8004D110:
		return XactELuRecoveryMismatch
	case 0x8004D111:
		return XactERmUnavailable
	case 0x8004E002:
		return ContextEAborted
	case 0x8004E003:
		return ContextEAborting
	case 0x8004E004:
		return ContextENocontext
	case 0x8004E005:
		return ContextEWouldDeadlock
	case 0x8004E006:
		return ContextESynchTimeout
	case 0x8004E007:
		return ContextEOldref
	case 0x8004E00C:
		return ContextERolenotfound
	case 0x8004E00F:
		return ContextETmnotavailable
	case 0x8004E021:
		return CoEActivationfailed
	case 0x8004E022:
		return CoEActivationfailedEventlogged
	case 0x8004E023:
		return CoEActivationfailedCatalogerror
	case 0x8004E024:
		return CoEActivationfailedTimeout
	case 0x8004E025:
		return CoEInitializationfailed
	case 0x8004E026:
		return ContextENojit
	case 0x8004E027:
		return ContextENotransaction
	case 0x8004E028:
		return CoEThreadingmodelChanged
	case 0x8004E029:
		return CoENoiisintrinsics
	case 0x8004E02A:
		return CoENocookies
	case 0x8004E02B:
		return CoEDberror
	case 0x8004E02C:
		return CoENotpooled
	case 0x8004E02D:
		return CoENotconstructed
	case 0x8004E02E:
		return CoENosynchronization
	case 0x8004E02F:
		return CoEIsolevelmismatch
	case 0x8004E030:
		return CoECallOutOfTxScopeNotAllowed
	case 0x8004E031:
		return CoEExitTransactionScopeNotCalled
	case 0x80070005:
		return EAccessdenied
	case 0x8007000E:
		return EOutofmemory
	case 0x80070032:
		return ErrorNotSupported
	case 0x80070057:
		return EInvalidarg
	case 0x80070070:
		return ErrorDiskFull
	case 0x80080001:
		return CoEClassCreateFailed
	case 0x80080002:
		return CoEScmError
	case 0x80080003:
		return CoEScmRpcFailure
	case 0x80080004:
		return CoEBadPath
	case 0x80080005:
		return CoEServerExecFailure
	case 0x80080006:
		return CoEObjsrvRpcFailure
	case 0x80080007:
		return MkENoNormalized
	case 0x80080008:
		return CoEServerStopping
	case 0x80080009:
		return MemEInvalidRoot
	case 0x80080010:
		return MemEInvalidLink
	case 0x80080011:
		return MemEInvalidSize
	case 0x80080015:
		return CoEMissingDisplayname
	case 0x80080016:
		return CoERunasValueMustBeAaa
	case 0x80080017:
		return CoEElevationDisabled
	case 0x80090001:
		return NteBadUid
	case 0x80090002:
		return NteBadHash
	case 0x80090003:
		return NteBadKey
	case 0x80090004:
		return NteBadLen
	case 0x80090005:
		return NteBadData
	case 0x80090006:
		return NteBadSignature
	case 0x80090007:
		return NteBadVer
	case 0x80090008:
		return NteBadAlgid
	case 0x80090009:
		return NteBadFlags
	case 0x8009000A:
		return NteBadType
	case 0x8009000B:
		return NteBadKeyState
	case 0x8009000C:
		return NteBadHashState
	case 0x8009000D:
		return NteNoKey
	case 0x8009000E:
		return NteNoMemory
	case 0x8009000F:
		return NteExists
	case 0x80090010:
		return NtePerm
	case 0x80090011:
		return NteNotFound
	case 0x80090012:
		return NteDoubleEncrypt
	case 0x80090013:
		return NteBadProvider
	case 0x80090014:
		return NteBadProvType
	case 0x80090015:
		return NteBadPublicKey
	case 0x80090016:
		return NteBadKeyset
	case 0x80090017:
		return NteProvTypeNotDef
	case 0x80090018:
		return NteProvTypeEntryBad
	case 0x80090019:
		return NteKeysetNotDef
	case 0x8009001A:
		return NteKeysetEntryBad
	case 0x8009001B:
		return NteProvTypeNoMatch
	case 0x8009001C:
		return NteSignatureFileBad
	case 0x8009001D:
		return NteProviderDllFail
	case 0x8009001E:
		return NteProvDllNotFound
	case 0x8009001F:
		return NteBadKeysetParam
	case 0x80090020:
		return NteFail
	case 0x80090021:
		return NteSysErr
	case 0x80090022:
		return NteSilentContext
	case 0x80090023:
		return NteTokenKeysetStorageFull
	case 0x80090024:
		return NteTemporaryProfile
	case 0x80090025:
		return NteFixedparameter
	case 0x80090026:
		return NteInvalidHandle
	case 0x80090027:
		return NteInvalidParameter
	case 0x80090028:
		return NteBufferTooSmall
	case 0x80090029:
		return NteNotSupported
	case 0x8009002A:
		return NteNoMoreItems
	case 0x8009002B:
		return NteBuffersOverlap
	case 0x8009002C:
		return NteDecryptionFailure
	case 0x8009002D:
		return NteInternalError
	case 0x8009002E:
		return NteUiRequired
	case 0x8009002F:
		return NteHmacNotSupported
	case 0x80090300:
		return SecEInsufficientMemory
	case 0x80090301:
		return SecEInvalidHandle
	case 0x80090302:
		return SecEUnsupportedFunction
	case 0x80090303:
		return SecETargetUnknown
	case 0x80090304:
		return SecEInternalError
	case 0x80090305:
		return SecESecpkgNotFound
	case 0x80090306:
		return SecENotOwner
	case 0x80090307:
		return SecECannotInstall
	case 0x80090308:
		return SecEInvalidToken
	case 0x80090309:
		return SecECannotPack
	case 0x8009030A:
		return SecEQopNotSupported
	case 0x8009030B:
		return SecENoImpersonation
	case 0x8009030C:
		return SecELogonDenied
	case 0x8009030D:
		return SecEUnknownCredentials
	case 0x8009030E:
		return SecENoCredentials
	case 0x8009030F:
		return SecEMessageAltered
	case 0x80090310:
		return SecEOutOfSequence
	case 0x80090311:
		return SecENoAuthenticatingAuthority
	case 0x80090316:
		return SecEBadPkgid
	case 0x80090317:
		return SecEContextExpired
	case 0x80090318:
		return SecEIncompleteMessage
	case 0x80090320:
		return SecEIncompleteCredentials
	case 0x80090321:
		return SecEBufferTooSmall
	case 0x80090322:
		return SecEWrongPrincipal
	case 0x80090324:
		return SecETimeSkew
	case 0x80090325:
		return SecEUntrustedRoot
	case 0x80090326:
		return SecEIllegalMessage
	case 0x80090327:
		return SecECertUnknown
	case 0x80090328:
		return SecECertExpired
	case 0x80090329:
		return SecEEncryptFailure
	case 0x80090330:
		return SecEDecryptFailure
	case 0x80090331:
		return SecEAlgorithmMismatch
	case 0x80090332:
		return SecESecurityQosFailed
	case 0x80090333:
		return SecEUnfinishedContextDeleted
	case 0x80090334:
		return SecENoTgtReply
	case 0x80090335:
		return SecENoIpAddresses
	case 0x80090336:
		return SecEWrongCredentialHandle
	case 0x80090337:
		return SecECryptoSystemInvalid
	case 0x80090338:
		return SecEMaxReferralsExceeded
	case 0x80090339:
		return SecEMustBeKdc
	case 0x8009033A:
		return SecEStrongCryptoNotSupported
	case 0x8009033B:
		return SecETooManyPrincipals
	case 0x8009033C:
		return SecENoPaData
	case 0x8009033D:
		return SecEPkinitNameMismatch
	case 0x8009033E:
		return SecESmartcardLogonRequired
	case 0x8009033F:
		return SecEShutdownInProgress
	case 0x80090340:
		return SecEKdcInvalidRequest
	case 0x80090341:
		return SecEKdcUnableToRefer
	case 0x80090342:
		return SecEKdcUnknownEtype
	case 0x80090343:
		return SecEUnsupportedPreauth
	case 0x80090345:
		return SecEDelegationRequired
	case 0x80090346:
		return SecEBadBindings
	case 0x80090347:
		return SecEMultipleAccounts
	case 0x80090348:
		return SecENoKerbKey
	case 0x80090349:
		return SecECertWrongUsage
	case 0x80090350:
		return SecEDowngradeDetected
	case 0x80090351:
		return SecESmartcardCertRevoked
	case 0x80090352:
		return SecEIssuingCaUntrusted
	case 0x80090353:
		return SecERevocationOfflineC
	case 0x80090354:
		return SecEPkinitClientFailure
	case 0x80090355:
		return SecESmartcardCertExpired
	case 0x80090356:
		return SecENoS4uProtSupport
	case 0x80090357:
		return SecECrossrealmDelegationFailure
	case 0x80090358:
		return SecERevocationOfflineKdc
	case 0x80090359:
		return SecEIssuingCaUntrustedKdc
	case 0x8009035A:
		return SecEKdcCertExpired
	case 0x8009035B:
		return SecEKdcCertRevoked
	case 0x8009035D:
		return SecEInvalidParameter
	case 0x8009035E:
		return SecEDelegationPolicy
	case 0x8009035F:
		return SecEPolicyNltmOnly
	case 0x80091001:
		return CryptEMsgError
	case 0x80091002:
		return CryptEUnknownAlgo
	case 0x80091003:
		return CryptEOidFormat
	case 0x80091004:
		return CryptEInvalidMsgType
	case 0x80091005:
		return CryptEUnexpectedEncoding
	case 0x80091006:
		return CryptEAuthAttrMissing
	case 0x80091007:
		return CryptEHashValue
	case 0x80091008:
		return CryptEInvalidIndex
	case 0x80091009:
		return CryptEAlreadyDecrypted
	case 0x8009100A:
		return CryptENotDecrypted
	case 0x8009100B:
		return CryptERecipientNotFound
	case 0x8009100C:
		return CryptEControlType
	case 0x8009100D:
		return CryptEIssuerSerialnumber
	case 0x8009100E:
		return CryptESignerNotFound
	case 0x8009100F:
		return CryptEAttributesMissing
	case 0x80091010:
		return CryptEStreamMsgNotReady
	case 0x80091011:
		return CryptEStreamInsufficientData
	case 0x80092001:
		return CryptEBadLen
	case 0x80092002:
		return CryptEBadEncode
	case 0x80092003:
		return CryptEFileError
	case 0x80092004:
		return CryptENotFound
	case 0x80092005:
		return CryptEExists
	case 0x80092006:
		return CryptENoProvider
	case 0x80092007:
		return CryptESelfSigned
	case 0x80092008:
		return CryptEDeletedPrev
	case 0x80092009:
		return CryptENoMatch
	case 0x8009200A:
		return CryptEUnexpectedMsgType
	case 0x8009200B:
		return CryptENoKeyProperty
	case 0x8009200C:
		return CryptENoDecryptCert
	case 0x8009200D:
		return CryptEBadMsg
	case 0x8009200E:
		return CryptENoSigner
	case 0x8009200F:
		return CryptEPendingClose
	case 0x80092010:
		return CryptERevoked
	case 0x80092011:
		return CryptENoRevocationDll
	case 0x80092012:
		return CryptENoRevocationCheck
	case 0x80092013:
		return CryptERevocationOffline
	case 0x80092014:
		return CryptENotInRevocationDatabase
	case 0x80092020:
		return CryptEInvalidNumericString
	case 0x80092021:
		return CryptEInvalidPrintableString
	case 0x80092022:
		return CryptEInvalidIa5String
	case 0x80092023:
		return CryptEInvalidX500String
	case 0x80092024:
		return CryptENotCharString
	case 0x80092025:
		return CryptEFileresized
	case 0x80092026:
		return CryptESecuritySettings
	case 0x80092027:
		return CryptENoVerifyUsageDll
	case 0x80092028:
		return CryptENoVerifyUsageCheck
	case 0x80092029:
		return CryptEVerifyUsageOffline
	case 0x8009202A:
		return CryptENotInCtl
	case 0x8009202B:
		return CryptENoTrustedSigner
	case 0x8009202C:
		return CryptEMissingPubkeyPara
	case 0x80093000:
		return CryptEOssError
	case 0x80093001:
		return OssMoreBuf
	case 0x80093002:
		return OssNegativeUinteger
	case 0x80093003:
		return OssPduRange
	case 0x80093004:
		return OssMoreInput
	case 0x80093005:
		return OssDataError
	case 0x80093006:
		return OssBadArg
	case 0x80093007:
		return OssBadVersion
	case 0x80093008:
		return OssOutMemory
	case 0x80093009:
		return OssPduMismatch
	case 0x8009300A:
		return OssLimited
	case 0x8009300B:
		return OssBadPtr
	case 0x8009300C:
		return OssBadTime
	case 0x8009300D:
		return OssIndefiniteNotSupported
	case 0x8009300E:
		return OssMemError
	case 0x8009300F:
		return OssBadTable
	case 0x80093010:
		return OssTooLong
	case 0x80093011:
		return OssConstraintViolated
	case 0x80093012:
		return OssFatalError
	case 0x80093013:
		return OssAccessSerializationError
	case 0x80093014:
		return OssNullTbl
	case 0x80093015:
		return OssNullFcn
	case 0x80093016:
		return OssBadEncrules
	case 0x80093017:
		return OssUnavailEncrules
	case 0x80093018:
		return OssCantOpenTraceWindow
	case 0x80093019:
		return OssUnimplemented
	case 0x8009301A:
		return OssOidDllNotLinked
	case 0x8009301B:
		return OssCantOpenTraceFile
	case 0x8009301C:
		return OssTraceFileAlreadyOpen
	case 0x8009301D:
		return OssTableMismatch
	case 0x8009301E:
		return OssTypeNotSupported
	case 0x8009301F:
		return OssRealDllNotLinked
	case 0x80093020:
		return OssRealCodeNotLinked
	case 0x80093021:
		return OssOutOfRange
	case 0x80093022:
		return OssCopierDllNotLinked
	case 0x80093023:
		return OssConstraintDllNotLinked
	case 0x80093024:
		return OssComparatorDllNotLinked
	case 0x80093025:
		return OssComparatorCodeNotLinked
	case 0x80093026:
		return OssMemMgrDllNotLinked
	case 0x80093027:
		return OssPdvDllNotLinked
	case 0x80093028:
		return OssPdvCodeNotLinked
	case 0x80093029:
		return OssApiDllNotLinked
	case 0x8009302A:
		return OssBerderDllNotLinked
	case 0x8009302B:
		return OssPerDllNotLinked
	case 0x8009302C:
		return OssOpenTypeError
	case 0x8009302D:
		return OssMutexNotCreated
	case 0x8009302E:
		return OssCantCloseTraceFile
	case 0x80093100:
		return CryptEAsn1Error
	case 0x80093101:
		return CryptEAsn1Internal
	case 0x80093102:
		return CryptEAsn1Eod
	case 0x80093103:
		return CryptEAsn1Corrupt
	case 0x80093104:
		return CryptEAsn1Large
	case 0x80093105:
		return CryptEAsn1Constraint
	case 0x80093106:
		return CryptEAsn1Memory
	case 0x80093107:
		return CryptEAsn1Overflow
	case 0x80093108:
		return CryptEAsn1Badpdu
	case 0x80093109:
		return CryptEAsn1Badargs
	case 0x8009310A:
		return CryptEAsn1Badreal
	case 0x8009310B:
		return CryptEAsn1Badtag
	case 0x8009310C:
		return CryptEAsn1Choice
	case 0x8009310D:
		return CryptEAsn1Rule
	case 0x8009310E:
		return CryptEAsn1Utf8
	case 0x80093133:
		return CryptEAsn1PduType
	case 0x80093134:
		return CryptEAsn1Nyi
	case 0x80093201:
		return CryptEAsn1Extended
	case 0x80093202:
		return CryptEAsn1Noeod
	case 0x80094001:
		return CertsrvEBadRequestsubject
	case 0x80094002:
		return CertsrvENoRequest
	case 0x80094003:
		return CertsrvEBadRequeststatus
	case 0x80094004:
		return CertsrvEPropertyEmpty
	case 0x80094005:
		return CertsrvEInvalidCaCertificate
	case 0x80094006:
		return CertsrvEServerSuspended
	case 0x80094007:
		return CertsrvEEncodingLength
	case 0x80094008:
		return CertsrvERoleconflict
	case 0x80094009:
		return CertsrvERestrictedofficer
	case 0x8009400A:
		return CertsrvEKeyArchivalNotConfigured
	case 0x8009400B:
		return CertsrvENoValidKra
	case 0x8009400C:
		return CertsrvEBadRequestKeyArchival
	case 0x8009400D:
		return CertsrvENoCaadminDefined
	case 0x8009400E:
		return CertsrvEBadRenewalCertAttribute
	case 0x8009400F:
		return CertsrvENoDbSessions
	case 0x80094010:
		return CertsrvEAlignmentFault
	case 0x80094011:
		return CertsrvEEnrollDenied
	case 0x80094012:
		return CertsrvETemplateDenied
	case 0x80094013:
		return CertsrvEDownlevelDcSslOrUpgrade
	case 0x80094800:
		return CertsrvEUnsupportedCertType
	case 0x80094801:
		return CertsrvENoCertType
	case 0x80094802:
		return CertsrvETemplateConflict
	case 0x80094803:
		return CertsrvESubjectAltNameRequired
	case 0x80094804:
		return CertsrvEArchivedKeyRequired
	case 0x80094805:
		return CertsrvESmimeRequired
	case 0x80094806:
		return CertsrvEBadRenewalSubject
	case 0x80094807:
		return CertsrvEBadTemplateVersion
	case 0x80094808:
		return CertsrvETemplatePolicyRequired
	case 0x80094809:
		return CertsrvESignaturePolicyRequired
	case 0x8009480A:
		return CertsrvESignatureCount
	case 0x8009480B:
		return CertsrvESignatureRejected
	case 0x8009480C:
		return CertsrvEIssuancePolicyRequired
	case 0x8009480D:
		return CertsrvESubjectUpnRequired
	case 0x8009480E:
		return CertsrvESubjectDirectoryGuidRequired
	case 0x8009480F:
		return CertsrvESubjectDnsRequired
	case 0x80094810:
		return CertsrvEArchivedKeyUnexpected
	case 0x80094811:
		return CertsrvEKeyLength
	case 0x80094812:
		return CertsrvESubjectEmailRequired
	case 0x80094813:
		return CertsrvEUnknownCertType
	case 0x80094814:
		return CertsrvECertTypeOverlap
	case 0x80094815:
		return CertsrvETooManySignatures
	case 0x80094816:
		return CertsrvERenewalBadPublicKey
	case 0x80094817:
		return CertsrvEInvalidEk
	case 0x8009481A:
		return CertsrvEKeyAttestation
	case 0x80095000:
		return XenrollEKeyNotExportable
	case 0x80095001:
		return XenrollECannotAddRootCert
	case 0x80095002:
		return XenrollEResponseKaHashNotFound
	case 0x80095003:
		return XenrollEResponseUnexpectedKaHash
	case 0x80095004:
		return XenrollEResponseKaHashMismatch
	case 0x80095005:
		return XenrollEKeyspecSmimeMismatch
	case 0x80096001:
		return TrustESystemError
	case 0x80096002:
		return TrustENoSignerCert
	case 0x80096003:
		return TrustECounterSigner
	case 0x80096004:
		return TrustECertSignature
	case 0x80096005:
		return TrustETimeStamp
	case 0x80096010:
		return TrustEBadDigest
	case 0x80096019:
		return TrustEBasicConstraints
	case 0x8009601E:
		return TrustEFinancialCriteria
	case 0x80097001:
		return MssipotfEOutofmemrange
	case 0x80097002:
		return MssipotfECantgetobject
	case 0x80097003:
		return MssipotfENoheadtable
	case 0x80097004:
		return MssipotfEBadMagicnumber
	case 0x80097005:
		return MssipotfEBadOffsetTable
	case 0x80097006:
		return MssipotfETableTagorder
	case 0x80097007:
		return MssipotfETableLongword
	case 0x80097008:
		return MssipotfEBadFirstTablePlacement
	case 0x80097009:
		return MssipotfETablesOverlap
	case 0x8009700A:
		return MssipotfETablePadbytes
	case 0x8009700B:
		return MssipotfEFiletoosmall
	case 0x8009700C:
		return MssipotfETableChecksum
	case 0x8009700D:
		return MssipotfEFileChecksum
	case 0x80097010:
		return MssipotfEFailedPolicy
	case 0x80097011:
		return MssipotfEFailedHintsCheck
	case 0x80097012:
		return MssipotfENotOpentype
	case 0x80097013:
		return MssipotfEFile
	case 0x80097014:
		return MssipotfECrypt
	case 0x80097015:
		return MssipotfEBadversion
	case 0x80097016:
		return MssipotfEDsigStructure
	case 0x80097017:
		return MssipotfEPconstCheck
	case 0x80097018:
		return MssipotfEStructure
	case 0x80097019:
		return ErrorCredRequiresConfirmation
	case 0x800B0001:
		return TrustEProviderUnknown
	case 0x800B0002:
		return TrustEActionUnknown
	case 0x800B0003:
		return TrustESubjectFormUnknown
	case 0x800B0004:
		return TrustESubjectNotTrusted
	case 0x800B0005:
		return DigsigEEncode
	case 0x800B0006:
		return DigsigEDecode
	case 0x800B0007:
		return DigsigEExtensibility
	case 0x800B0008:
		return DigsigECrypto
	case 0x800B0009:
		return PersistESizedefinite
	case 0x800B000A:
		return PersistESizeindefinite
	case 0x800B000B:
		return PersistENotselfsizing
	case 0x800B0100:
		return TrustENosignature
	case 0x800B0101:
		return CertEExpired
	case 0x800B0102:
		return CertEValidityperiodnesting
	case 0x800B0103:
		return CertERole
	case 0x800B0104:
		return CertEPathlenconst
	case 0x800B0105:
		return CertECritical
	case 0x800B0106:
		return CertEPurpose
	case 0x800B0107:
		return CertEIssuerchaining
	case 0x800B0108:
		return CertEMalformed
	case 0x800B0109:
		return CertEUntrustedroot
	case 0x800B010A:
		return CertEChaining
	case 0x800B010B:
		return TrustEFail
	case 0x800B010C:
		return CertERevoked
	case 0x800B010D:
		return CertEUntrustedtestroot
	case 0x800B010E:
		return CertERevocationFailure
	case 0x800B010F:
		return CertECnNoMatch
	case 0x800B0110:
		return CertEWrongUsage
	case 0x800B0111:
		return TrustEExplicitDistrust
	case 0x800B0112:
		return CertEUntrustedca
	case 0x800B0113:
		return CertEInvalidPolicy
	case 0x800B0114:
		return CertEInvalidName
	case 0x800D0003:
		return NsWServerBandwidthLimit
	case 0x800D0004:
		return NsWFileBandwidthLimit
	case 0x800D0060:
		return NsWUnknownEvent
	case 0x800D0199:
		return NsICatatonicFailure
	case 0x800D019A:
		return NsICatatonicAutoUnfail
	case 0x800F0000:
		return SpapiEExpectedSectionName
	case 0x800F0001:
		return SpapiEBadSectionNameLine
	case 0x800F0002:
		return SpapiESectionNameTooLong
	case 0x800F0003:
		return SpapiEGeneralSyntax
	case 0x800F0100:
		return SpapiEWrongInfStyle
	case 0x800F0101:
		return SpapiESectionNotFound
	case 0x800F0102:
		return SpapiELineNotFound
	case 0x800F0103:
		return SpapiENoBackup
	case 0x800F0200:
		return SpapiENoAssociatedClass
	case 0x800F0201:
		return SpapiEClassMismatch
	case 0x800F0202:
		return SpapiEDuplicateFound
	case 0x800F0203:
		return SpapiENoDriverSelected
	case 0x800F0204:
		return SpapiEKeyDoesNotExist
	case 0x800F0205:
		return SpapiEInvalidDevinstName
	case 0x800F0206:
		return SpapiEInvalidClass
	case 0x800F0207:
		return SpapiEDevinstAlreadyExists
	case 0x800F0208:
		return SpapiEDevinfoNotRegistered
	case 0x800F0209:
		return SpapiEInvalidRegProperty
	case 0x800F020A:
		return SpapiENoInf
	case 0x800F020B:
		return SpapiENoSuchDevinst
	case 0x800F020C:
		return SpapiECantLoadClassIcon
	case 0x800F020D:
		return SpapiEInvalidClassInstaller
	case 0x800F020E:
		return SpapiEDiDoDefault
	case 0x800F020F:
		return SpapiEDiNofilecopy
	case 0x800F0210:
		return SpapiEInvalidHwprofile
	case 0x800F0211:
		return SpapiENoDeviceSelected
	case 0x800F0212:
		return SpapiEDevinfoListLocked
	case 0x800F0213:
		return SpapiEDevinfoDataLocked
	case 0x800F0214:
		return SpapiEDiBadPath
	case 0x800F0215:
		return SpapiENoClassinstallParams
	case 0x800F0216:
		return SpapiEFilequeueLocked
	case 0x800F0217:
		return SpapiEBadServiceInstallsect
	case 0x800F0218:
		return SpapiENoClassDriverList
	case 0x800F0219:
		return SpapiENoAssociatedService
	case 0x800F021A:
		return SpapiENoDefaultDeviceInterface
	case 0x800F021B:
		return SpapiEDeviceInterfaceActive
	case 0x800F021C:
		return SpapiEDeviceInterfaceRemoved
	case 0x800F021D:
		return SpapiEBadInterfaceInstallsect
	case 0x800F021E:
		return SpapiENoSuchInterfaceClass
	case 0x800F021F:
		return SpapiEInvalidReferenceString
	case 0x800F0220:
		return SpapiEInvalidMachinename
	case 0x800F0221:
		return SpapiERemoteCommFailure
	case 0x800F0222:
		return SpapiEMachineUnavailable
	case 0x800F0223:
		return SpapiENoConfigmgrServices
	case 0x800F0224:
		return SpapiEInvalidProppageProvider
	case 0x800F0225:
		return SpapiENoSuchDeviceInterface
	case 0x800F0226:
		return SpapiEDiPostprocessingRequired
	case 0x800F0227:
		return SpapiEInvalidCoinstaller
	case 0x800F0228:
		return SpapiENoCompatDrivers
	case 0x800F0229:
		return SpapiENoDeviceIcon
	case 0x800F022A:
		return SpapiEInvalidInfLogconfig
	case 0x800F022B:
		return SpapiEDiDontInstall
	case 0x800F022C:
		return SpapiEInvalidFilterDriver
	case 0x800F022D:
		return SpapiENonWindowsNtDriver
	case 0x800F022E:
		return SpapiENonWindowsDriver
	case 0x800F022F:
		return SpapiENoCatalogForOemInf
	case 0x800F0230:
		return SpapiEDevinstallQueueNonnative
	case 0x800F0231:
		return SpapiENotDisableable
	case 0x800F0232:
		return SpapiECantRemoveDevinst
	case 0x800F0233:
		return SpapiEInvalidTarget
	case 0x800F0234:
		return SpapiEDriverNonnative
	case 0x800F0235:
		return SpapiEInWow64
	case 0x800F0236:
		return SpapiESetSystemRestorePoint
	case 0x800F0237:
		return SpapiEIncorrectlyCopiedInf
	case 0x800F0238:
		return SpapiESceDisabled
	case 0x800F0239:
		return SpapiEUnknownException
	case 0x800F023A:
		return SpapiEPnpRegistryError
	case 0x800F023B:
		return SpapiERemoteRequestUnsupported
	case 0x800F023C:
		return SpapiENotAnInstalledOemInf
	case 0x800F023D:
		return SpapiEInfInUseByDevices
	case 0x800F023E:
		return SpapiEDiFunctionObsolete
	case 0x800F023F:
		return SpapiENoAuthenticodeCatalog
	case 0x800F0240:
		return SpapiEAuthenticodeDisallowed
	case 0x800F0241:
		return SpapiEAuthenticodeTrustedPublisher
	case 0x800F0242:
		return SpapiEAuthenticodeTrustNotEstablished
	case 0x800F0243:
		return SpapiEAuthenticodePublisherNotTrusted
	case 0x800F0244:
		return SpapiESignatureOsattributeMismatch
	case 0x800F0245:
		return SpapiEOnlyValidateViaAuthenticode
	case 0x800F0246:
		return SpapiEDeviceInstallerNotReady
	case 0x800F0247:
		return SpapiEDriverStoreAddFailed
	case 0x800F0248:
		return SpapiEDeviceInstallBlocked
	case 0x800F0249:
		return SpapiEDriverInstallBlocked
	case 0x800F024A:
		return SpapiEWrongInfType
	case 0x800F024B:
		return SpapiEFileHashNotInCatalog
	case 0x800F024C:
		return SpapiEDriverStoreDeleteFailed
	case 0x800F0300:
		return SpapiEUnrecoverableStackOverflow
	case 0x800F1000:
		return SpapiEErrorNotInstalled
	case 0x80100001:
		return ScardFInternalError
	case 0x80100002:
		return ScardECancelled
	case 0x80100003:
		return ScardEInvalidHandle
	case 0x80100004:
		return ScardEInvalidParameter
	case 0x80100005:
		return ScardEInvalidTarget
	case 0x80100006:
		return ScardENoMemory
	case 0x80100007:
		return ScardFWaitedTooLong
	case 0x80100008:
		return ScardEInsufficientBuffer
	case 0x80100009:
		return ScardEUnknownReader
	case 0x8010000A:
		return ScardETimeout
	case 0x8010000B:
		return ScardESharingViolation
	case 0x8010000C:
		return ScardENoSmartcard
	case 0x8010000D:
		return ScardEUnknownCard
	case 0x8010000E:
		return ScardECantDispose
	case 0x8010000F:
		return ScardEProtoMismatch
	case 0x80100010:
		return ScardENotReady
	case 0x80100011:
		return ScardEInvalidValue
	case 0x80100012:
		return ScardESystemCancelled
	case 0x80100013:
		return ScardFCommError
	case 0x80100014:
		return ScardFUnknownError
	case 0x80100015:
		return ScardEInvalidAtr
	case 0x80100016:
		return ScardENotTransacted
	case 0x80100017:
		return ScardEReaderUnavailable
	case 0x80100018:
		return ScardPShutdown
	case 0x80100019:
		return ScardEPciTooSmall
	case 0x8010001A:
		return ScardEReaderUnsupported
	case 0x8010001B:
		return ScardEDuplicateReader
	case 0x8010001C:
		return ScardECardUnsupported
	case 0x8010001D:
		return ScardENoService
	case 0x8010001E:
		return ScardEServiceStopped
	case 0x8010001F:
		return ScardEUnexpected
	case 0x80100020:
		return ScardEIccInstallation
	case 0x80100021:
		return ScardEIccCreateorder
	case 0x80100022:
		return ScardEUnsupportedFeature
	case 0x80100023:
		return ScardEDirNotFound
	case 0x80100024:
		return ScardEFileNotFound
	case 0x80100025:
		return ScardENoDir
	case 0x80100026:
		return ScardENoFile
	case 0x80100027:
		return ScardENoAccess
	case 0x80100028:
		return ScardEWriteTooMany
	case 0x80100029:
		return ScardEBadSeek
	case 0x8010002A:
		return ScardEInvalidChv
	case 0x8010002B:
		return ScardEUnknownResMng
	case 0x8010002C:
		return ScardENoSuchCertificate
	case 0x8010002D:
		return ScardECertificateUnavailable
	case 0x8010002E:
		return ScardENoReadersAvailable
	case 0x8010002F:
		return ScardECommDataLost
	case 0x80100030:
		return ScardENoKeyContainer
	case 0x80100031:
		return ScardEServerTooBusy
	case 0x80100065:
		return ScardWUnsupportedCard
	case 0x80100066:
		return ScardWUnresponsiveCard
	case 0x80100067:
		return ScardWUnpoweredCard
	case 0x80100068:
		return ScardWResetCard
	case 0x80100069:
		return ScardWRemovedCard
	case 0x8010006A:
		return ScardWSecurityViolation
	case 0x8010006B:
		return ScardWWrongChv
	case 0x8010006C:
		return ScardWChvBlocked
	case 0x8010006D:
		return ScardWEof
	case 0x8010006E:
		return ScardWCancelledByUser
	case 0x8010006F:
		return ScardWCardNotAuthenticated
	case 0x80110401:
		return ComadminEObjecterrors
	case 0x80110402:
		return ComadminEObjectinvalid
	case 0x80110403:
		return ComadminEKeymissing
	case 0x80110404:
		return ComadminEAlreadyinstalled
	case 0x80110407:
		return ComadminEAppFileWritefail
	case 0x80110408:
		return ComadminEAppFileReadfail
	case 0x80110409:
		return ComadminEAppFileVersion
	case 0x8011040A:
		return ComadminEBadpath
	case 0x8011040B:
		return ComadminEApplicationexists
	case 0x8011040C:
		return ComadminERoleexists
	case 0x8011040D:
		return ComadminECantcopyfile
	case 0x8011040F:
		return ComadminENouser
	case 0x80110410:
		return ComadminEInvaliduserids
	case 0x80110411:
		return ComadminENoregistryclsid
	case 0x80110412:
		return ComadminEBadregistryprogid
	case 0x80110413:
		return ComadminEAuthenticationlevel
	case 0x80110414:
		return ComadminEUserpasswdnotvalid
	case 0x80110418:
		return ComadminEClsidoriidmismatch
	case 0x80110419:
		return ComadminERemoteinterface
	case 0x8011041A:
		return ComadminEDllregisterserver
	case 0x8011041B:
		return ComadminENoservershare
	case 0x8011041D:
		return ComadminEDllloadfailed
	case 0x8011041E:
		return ComadminEBadregistrylibid
	case 0x8011041F:
		return ComadminEAppdirnotfound
	case 0x80110423:
		return ComadminERegistrarfailed
	case 0x80110424:
		return ComadminECompfileDoesnotexist
	case 0x80110425:
		return ComadminECompfileLoaddllfail
	case 0x80110426:
		return ComadminECompfileGetclassobj
	case 0x80110427:
		return ComadminECompfileClassnotavail
	case 0x80110428:
		return ComadminECompfileBadtlb
	case 0x80110429:
		return ComadminECompfileNotinstallable
	case 0x8011042A:
		return ComadminENotchangeable
	case 0x8011042B:
		return ComadminENotdeleteable
	case 0x8011042C:
		return ComadminESession
	case 0x8011042D:
		return ComadminECompMoveLocked
	case 0x8011042E:
		return ComadminECompMoveBadDest
	case 0x80110430:
		return ComadminERegistertlb
	case 0x80110433:
		return ComadminESystemapp
	case 0x80110434:
		return ComadminECompfileNoregistrar
	case 0x80110435:
		return ComadminECoreqcompinstalled
	case 0x80110436:
		return ComadminEServicenotinstalled
	case 0x80110437:
		return ComadminEPropertysavefailed
	case 0x80110438:
		return ComadminEObjectexists
	case 0x80110439:
		return ComadminEComponentexists
	case 0x8011043B:
		return ComadminERegfileCorrupt
	case 0x8011043C:
		return ComadminEPropertyOverflow
	case 0x8011043E:
		return ComadminENotinregistry
	case 0x8011043F:
		return ComadminEObjectnotpoolable
	case 0x80110446:
		return ComadminEApplidMatchesClsid
	case 0x80110447:
		return ComadminERoleDoesNotExist
	case 0x80110448:
		return ComadminEStartAppNeedsComponents
	case 0x80110449:
		return ComadminERequiresDifferentPlatform
	case 0x8011044A:
		return ComadminECanNotExportAppProxy
	case 0x8011044B:
		return ComadminECanNotStartApp
	case 0x8011044C:
		return ComadminECanNotExportSysApp
	case 0x8011044D:
		return ComadminECantSubscribeToComponent
	case 0x8011044E:
		return ComadminEEventclassCantBeSubscriber
	case 0x8011044F:
		return ComadminELibAppProxyIncompatible
	case 0x80110450:
		return ComadminEBasePartitionOnly
	case 0x80110451:
		return ComadminEStartAppDisabled
	case 0x80110457:
		return ComadminECatDuplicatePartitionName
	case 0x80110458:
		return ComadminECatInvalidPartitionName
	case 0x80110459:
		return ComadminECatPartitionInUse
	case 0x8011045A:
		return ComadminEFilePartitionDuplicateFiles
	case 0x8011045B:
		return ComadminECatImportedComponentsNotAllowed
	case 0x8011045C:
		return ComadminEAmbiguousApplicationName
	case 0x8011045D:
		return ComadminEAmbiguousPartitionName
	case 0x80110472:
		return ComadminERegdbNotinitialized
	case 0x80110473:
		return ComadminERegdbNotopen
	case 0x80110474:
		return ComadminERegdbSystemerr
	case 0x80110475:
		return ComadminERegdbAlreadyrunning
	case 0x80110480:
		return ComadminEMigVersionnotsupported
	case 0x80110481:
		return ComadminEMigSchemanotfound
	case 0x80110482:
		return ComadminECatBitnessmismatch
	case 0x80110483:
		return ComadminECatUnacceptablebitness
	case 0x80110484:
		return ComadminECatWrongappbitness
	case 0x80110485:
		return ComadminECatPauseResumeNotSupported
	case 0x80110486:
		return ComadminECatServerfault
	case 0x80110600:
		return ComqcEApplicationNotQueued
	case 0x80110601:
		return ComqcENoQueueableInterfaces
	case 0x80110602:
		return ComqcEQueuingServiceNotAvailable
	case 0x80110603:
		return ComqcENoIpersiststream
	case 0x80110604:
		return ComqcEBadMessage
	case 0x80110605:
		return ComqcEUnauthenticated
	case 0x80110606:
		return ComqcEUntrustedEnqueuer
	case 0x80110701:
		return MsdtcEDuplicateResource
	case 0x80110808:
		return ComadminEObjectParentMissing
	case 0x80110809:
		return ComadminEObjectDoesNotExist
	case 0x8011080A:
		return ComadminEAppNotRunning
	case 0x8011080B:
		return ComadminEInvalidPartition
	case 0x8011080D:
		return ComadminESvcappNotPoolableOrRecyclable
	case 0x8011080E:
		return ComadminEUserInSet
	case 0x8011080F:
		return ComadminECantrecyclelibraryapps
	case 0x80110811:
		return ComadminECantrecycleserviceapps
	case 0x80110812:
		return ComadminEProcessalreadyrecycled
	case 0x80110813:
		return ComadminEPausedprocessmaynotberecycled
	case 0x80110814:
		return ComadminECantmakeinprocservice
	case 0x80110815:
		return ComadminEProgidinusebyclsid
	case 0x80110816:
		return ComadminEDefaultPartitionNotInSet
	case 0x80110817:
		return ComadminERecycledprocessmaynotbepaused
	case 0x80110818:
		return ComadminEPartitionAccessdenied
	case 0x80110819:
		return ComadminEPartitionMsiOnly
	case 0x8011081A:
		return ComadminELegacycompsNotAllowedIn10Format
	case 0x8011081B:
		return ComadminELegacycompsNotAllowedInNonbasePartitions
	case 0x8011081C:
		return ComadminECompMoveSource
	case 0x8011081D:
		return ComadminECompMoveDest
	case 0x8011081E:
		return ComadminECompMovePrivate
	case 0x8011081F:
		return ComadminEBasepartitionRequiredInSet
	case 0x80110820:
		return ComadminECannotAliasEventclass
	case 0x80110821:
		return ComadminEPrivateAccessdenied
	case 0x80110822:
		return ComadminESaferinvalid
	case 0x80110823:
		return ComadminERegistryAccessdenied
	case 0x80110824:
		return ComadminEPartitionsDisabled
	case 0x801F0001:
		return ErrorFltNoHandlerDefined
	case 0x801F0002:
		return ErrorFltContextAlreadyDefined
	case 0x801F0003:
		return ErrorFltInvalidAsynchronousRequest
	case 0x801F0004:
		return ErrorFltDisallowFastIo
	case 0x801F0005:
		return ErrorFltInvalidNameRequest
	case 0x801F0006:
		return ErrorFltNotSafeToPostOperation
	case 0x801F0007:
		return ErrorFltNotInitialized
	case 0x801F0008:
		return ErrorFltFilterNotReady
	case 0x801F0009:
		return ErrorFltPostOperationCleanup
	case 0x801F000A:
		return ErrorFltInternalError
	case 0x801F000B:
		return ErrorFltDeletingObject
	case 0x801F000C:
		return ErrorFltMustBeNonpagedPool
	case 0x801F000D:
		return ErrorFltDuplicateEntry
	case 0x801F000E:
		return ErrorFltCbdqDisabled
	case 0x801F000F:
		return ErrorFltDoNotAttach
	case 0x801F0010:
		return ErrorFltDoNotDetach
	case 0x801F0011:
		return ErrorFltInstanceAltitudeCollision
	case 0x801F0012:
		return ErrorFltInstanceNameCollision
	case 0x801F0013:
		return ErrorFltFilterNotFound
	case 0x801F0014:
		return ErrorFltVolumeNotFound
	case 0x801F0015:
		return ErrorFltInstanceNotFound
	case 0x801F0016:
		return ErrorFltContextAllocationNotFound
	case 0x801F0017:
		return ErrorFltInvalidContextRegistration
	case 0x801F0018:
		return ErrorFltNameCacheMiss
	case 0x801F0019:
		return ErrorFltNoDeviceObject
	case 0x801F001A:
		return ErrorFltVolumeAlreadyMounted
	case 0x801F001B:
		return ErrorFltAlreadyEnlisted
	case 0x801F001C:
		return ErrorFltContextAlreadyLinked
	case 0x801F0020:
		return ErrorFltNoWaiterForReply
	case 0x80260001:
		return ErrorHungDisplayDriverThread
	case 0x80261001:
		return ErrorMonitorNoDescriptor
	case 0x80261002:
		return ErrorMonitorUnknownDescriptorFormat
	case 0x80263001:
		return DwmECompositiondisabled
	case 0x80263002:
		return DwmERemotingNotSupported
	case 0x80263003:
		return DwmENoRedirectionSurfaceAvailable
	case 0x80263004:
		return DwmENotQueuingPresents
	case 0x80280000:
		return TpmEErrorMask
	case 0x80280001:
		return TpmEAuthfail
	case 0x80280002:
		return TpmEBadindex
	case 0x80280003:
		return TpmEBadParameter
	case 0x80280004:
		return TpmEAuditfailure
	case 0x80280005:
		return TpmEClearDisabled
	case 0x80280006:
		return TpmEDeactivated
	case 0x80280007:
		return TpmEDisabled
	case 0x80280008:
		return TpmEDisabledCmd
	case 0x80280009:
		return TpmEFail
	case 0x8028000A:
		return TpmEBadOrdinal
	case 0x8028000B:
		return TpmEInstallDisabled
	case 0x8028000C:
		return TpmEInvalidKeyhandle
	case 0x8028000D:
		return TpmEKeynotfound
	case 0x8028000E:
		return TpmEInappropriateEnc
	case 0x8028000F:
		return TpmEMigratefail
	case 0x80280010:
		return TpmEInvalidPcrInfo
	case 0x80280011:
		return TpmENospace
	case 0x80280012:
		return TpmENosrk
	case 0x80280013:
		return TpmENotsealedBlob
	case 0x80280014:
		return TpmEOwnerSet
	case 0x80280015:
		return TpmEResources
	case 0x80280016:
		return TpmEShortrandom
	case 0x80280017:
		return TpmESize
	case 0x80280018:
		return TpmEWrongpcrval
	case 0x80280019:
		return TpmEBadParamSize
	case 0x8028001A:
		return TpmEShaThread
	case 0x8028001B:
		return TpmEShaError
	case 0x8028001C:
		return TpmEFailedselftest
	case 0x8028001D:
		return TpmEAuth2fail
	case 0x8028001E:
		return TpmEBadtag
	case 0x8028001F:
		return TpmEIoerror
	case 0x80280020:
		return TpmEEncryptError
	case 0x80280021:
		return TpmEDecryptError
	case 0x80280022:
		return TpmEInvalidAuthhandle
	case 0x80280023:
		return TpmENoEndorsement
	case 0x80280024:
		return TpmEInvalidKeyusage
	case 0x80280025:
		return TpmEWrongEntitytype
	case 0x80280026:
		return TpmEInvalidPostinit
	case 0x80280027:
		return TpmEInappropriateSig
	case 0x80280028:
		return TpmEBadKeyProperty
	case 0x80280029:
		return TpmEBadMigration
	case 0x8028002A:
		return TpmEBadScheme
	case 0x8028002B:
		return TpmEBadDatasize
	case 0x8028002C:
		return TpmEBadMode
	case 0x8028002D:
		return TpmEBadPresence
	case 0x8028002E:
		return TpmEBadVersion
	case 0x8028002F:
		return TpmENoWrapTransport
	case 0x80280030:
		return TpmEAuditfailUnsuccessful
	case 0x80280031:
		return TpmEAuditfailSuccessful
	case 0x80280032:
		return TpmENotresetable
	case 0x80280033:
		return TpmENotlocal
	case 0x80280034:
		return TpmEBadType
	case 0x80280035:
		return TpmEInvalidResource
	case 0x80280036:
		return TpmENotfips
	case 0x80280037:
		return TpmEInvalidFamily
	case 0x80280038:
		return TpmENoNvPermission
	case 0x80280039:
		return TpmERequiresSign
	case 0x8028003A:
		return TpmEKeyNotsupported
	case 0x8028003B:
		return TpmEAuthConflict
	case 0x8028003C:
		return TpmEAreaLocked
	case 0x8028003D:
		return TpmEBadLocality
	case 0x8028003E:
		return TpmEReadOnly
	case 0x8028003F:
		return TpmEPerNowrite
	case 0x80280040:
		return TpmEFamilycount
	case 0x80280041:
		return TpmEWriteLocked
	case 0x80280042:
		return TpmEBadAttributes
	case 0x80280043:
		return TpmEInvalidStructure
	case 0x80280044:
		return TpmEKeyOwnerControl
	case 0x80280045:
		return TpmEBadCounter
	case 0x80280046:
		return TpmENotFullwrite
	case 0x80280047:
		return TpmEContextGap
	case 0x80280048:
		return TpmEMaxnvwrites
	case 0x80280049:
		return TpmENooperator
	case 0x8028004A:
		return TpmEResourcemissing
	case 0x8028004B:
		return TpmEDelegateLock
	case 0x8028004C:
		return TpmEDelegateFamily
	case 0x8028004D:
		return TpmEDelegateAdmin
	case 0x8028004E:
		return TpmETransportNotexclusive
	case 0x8028004F:
		return TpmEOwnerControl
	case 0x80280050:
		return TpmEDaaResources
	case 0x80280051:
		return TpmEDaaInputData0
	case 0x80280052:
		return TpmEDaaInputData1
	case 0x80280053:
		return TpmEDaaIssuerSettings
	case 0x80280054:
		return TpmEDaaTpmSettings
	case 0x80280055:
		return TpmEDaaStage
	case 0x80280056:
		return TpmEDaaIssuerValidity
	case 0x80280057:
		return TpmEDaaWrongW
	case 0x80280058:
		return TpmEBadHandle
	case 0x80280059:
		return TpmEBadDelegate
	case 0x8028005A:
		return TpmEBadcontext
	case 0x8028005B:
		return TpmEToomanycontexts
	case 0x8028005C:
		return TpmEMaTicketSignature
	case 0x8028005D:
		return TpmEMaDestination
	case 0x8028005E:
		return TpmEMaSource
	case 0x8028005F:
		return TpmEMaAuthority
	case 0x80280061:
		return TpmEPermanentek
	case 0x80280062:
		return TpmEBadSignature
	case 0x80280063:
		return TpmENocontextspace
	case 0x80280400:
		return TpmECommandBlocked
	case 0x80280401:
		return TpmEInvalidHandle
	case 0x80280402:
		return TpmEDuplicateVhandle
	case 0x80280403:
		return TpmEEmbeddedCommandBlocked
	case 0x80280404:
		return TpmEEmbeddedCommandUnsupported
	case 0x80280800:
		return TpmERetry
	case 0x80280801:
		return TpmENeedsSelftest
	case 0x80280802:
		return TpmEDoingSelftest
	case 0x80280803:
		return TpmEDefendLockRunning
	case 0x80284001:
		return TbsEInternalError
	case 0x80284002:
		return TbsEBadParameter
	case 0x80284003:
		return TbsEInvalidOutputPointer
	case 0x80284004:
		return TbsEInvalidContext
	case 0x80284005:
		return TbsEInsufficientBuffer
	case 0x80284006:
		return TbsEIoerror
	case 0x80284007:
		return TbsEInvalidContextParam
	case 0x80284008:
		return TbsEServiceNotRunning
	case 0x80284009:
		return TbsETooManyTbsContexts
	case 0x8028400A:
		return TbsETooManyResources
	case 0x8028400B:
		return TbsEServiceStartPending
	case 0x8028400C:
		return TbsEPpiNotSupported
	case 0x8028400D:
		return TbsECommandCanceled
	case 0x8028400E:
		return TbsEBufferTooLarge
	case 0x80290100:
		return TpmapiEInvalidState
	case 0x80290101:
		return TpmapiENotEnoughData
	case 0x80290102:
		return TpmapiETooMuchData
	case 0x80290103:
		return TpmapiEInvalidOutputPointer
	case 0x80290104:
		return TpmapiEInvalidParameter
	case 0x80290105:
		return TpmapiEOutOfMemory
	case 0x80290106:
		return TpmapiEBufferTooSmall
	case 0x80290107:
		return TpmapiEInternalError
	case 0x80290108:
		return TpmapiEAccessDenied
	case 0x80290109:
		return TpmapiEAuthorizationFailed
	case 0x8029010A:
		return TpmapiEInvalidContextHandle
	case 0x8029010B:
		return TpmapiETbsCommunicationError
	case 0x8029010C:
		return TpmapiETpmCommandError
	case 0x8029010D:
		return TpmapiEMessageTooLarge
	case 0x8029010E:
		return TpmapiEInvalidEncoding
	case 0x8029010F:
		return TpmapiEInvalidKeySize
	case 0x80290110:
		return TpmapiEEncryptionFailed
	case 0x80290111:
		return TpmapiEInvalidKeyParams
	case 0x80290112:
		return TpmapiEInvalidMigrationAuthorizationBlob
	case 0x80290113:
		return TpmapiEInvalidPcrIndex
	case 0x80290114:
		return TpmapiEInvalidDelegateBlob
	case 0x80290115:
		return TpmapiEInvalidContextParams
	case 0x80290116:
		return TpmapiEInvalidKeyBlob
	case 0x80290117:
		return TpmapiEInvalidPcrData
	case 0x80290118:
		return TpmapiEInvalidOwnerAuth
	case 0x80290200:
		return TbsimpEBufferTooSmall
	case 0x80290201:
		return TbsimpECleanupFailed
	case 0x80290202:
		return TbsimpEInvalidContextHandle
	case 0x80290203:
		return TbsimpEInvalidContextParam
	case 0x80290204:
		return TbsimpETpmError
	case 0x80290205:
		return TbsimpEHashBadKey
	case 0x80290206:
		return TbsimpEDuplicateVhandle
	case 0x80290207:
		return TbsimpEInvalidOutputPointer
	case 0x80290208:
		return TbsimpEInvalidParameter
	case 0x80290209:
		return TbsimpERpcInitFailed
	case 0x8029020A:
		return TbsimpESchedulerNotRunning
	case 0x8029020B:
		return TbsimpECommandCanceled
	case 0x8029020C:
		return TbsimpEOutOfMemory
	case 0x8029020D:
		return TbsimpEListNoMoreItems
	case 0x8029020E:
		return TbsimpEListNotFound
	case 0x8029020F:
		return TbsimpENotEnoughSpace
	case 0x80290210:
		return TbsimpENotEnoughTpmContexts
	case 0x80290211:
		return TbsimpECommandFailed
	case 0x80290212:
		return TbsimpEUnknownOrdinal
	case 0x80290213:
		return TbsimpEResourceExpired
	case 0x80290214:
		return TbsimpEInvalidResource
	case 0x80290215:
		return TbsimpENothingToUnload
	case 0x80290216:
		return TbsimpEHashTableFull
	case 0x80290217:
		return TbsimpETooManyTbsContexts
	case 0x80290218:
		return TbsimpETooManyResources
	case 0x80290219:
		return TbsimpEPpiNotSupported
	case 0x8029021A:
		return TbsimpETpmIncompatible
	case 0x80290300:
		return TpmEPpiAcpiFailure
	case 0x80290301:
		return TpmEPpiUserAbort
	case 0x80290302:
		return TpmEPpiBiosFailure
	case 0x80290303:
		return TpmEPpiNotSupported
	case 0x80300002:
		return PlaEDcsNotFound
	case 0x80300045:
		return PlaETooManyFolders
	case 0x80300070:
		return PlaENoMinDisk
	case 0x803000AA:
		return PlaEDcsInUse
	case 0x803000B7:
		return PlaEDcsAlreadyExists
	case 0x80300101:
		return PlaEPropertyConflict
	case 0x80300102:
		return PlaEDcsSingletonRequired
	case 0x80300103:
		return PlaECredentialsRequired
	case 0x80300104:
		return PlaEDcsNotRunning
	case 0x80300105:
		return PlaEConflictInclExclApi
	case 0x80300106:
		return PlaENetworkExeNotValid
	case 0x80300107:
		return PlaEExeAlreadyConfigured
	case 0x80300108:
		return PlaEExePathNotValid
	case 0x80300109:
		return PlaEDcAlreadyExists
	case 0x8030010A:
		return PlaEDcsStartWaitTimeout
	case 0x8030010B:
		return PlaEDcStartWaitTimeout
	case 0x8030010C:
		return PlaEReportWaitTimeout
	case 0x8030010D:
		return PlaENoDuplicates
	case 0x8030010E:
		return PlaEExeFullPathRequired
	case 0x8030010F:
		return PlaEInvalidSessionName
	case 0x80300110:
		return PlaEPlaChannelNotEnabled
	case 0x80300111:
		return PlaETaskschedChannelNotEnabled
	case 0x80310000:
		return FveELockedVolume
	case 0x80310001:
		return FveENotEncrypted
	case 0x80310002:
		return FveENoTpmBios
	case 0x80310003:
		return FveENoMbrMetric
	case 0x80310004:
		return FveENoBootsectorMetric
	case 0x80310005:
		return FveENoBootmgrMetric
	case 0x80310006:
		return FveEWrongBootmgr
	case 0x80310007:
		return FveESecureKeyRequired
	case 0x80310008:
		return FveENotActivated
	case 0x80310009:
		return FveEActionNotAllowed
	case 0x8031000A:
		return FveEAdSchemaNotInstalled
	case 0x8031000B:
		return FveEAdInvalidDatatype
	case 0x8031000C:
		return FveEAdInvalidDatasize
	case 0x8031000D:
		return FveEAdNoValues
	case 0x8031000E:
		return FveEAdAttrNotSet
	case 0x8031000F:
		return FveEAdGuidNotFound
	case 0x80310010:
		return FveEBadInformation
	case 0x80310011:
		return FveETooSmall
	case 0x80310012:
		return FveESystemVolume
	case 0x80310013:
		return FveEFailedWrongFs
	case 0x80310014:
		return FveEFailedBadFs
	case 0x80310015:
		return FveENotSupported
	case 0x80310016:
		return FveEBadData
	case 0x80310017:
		return FveEVolumeNotBound
	case 0x80310018:
		return FveETpmNotOwned
	case 0x80310019:
		return FveENotDataVolume
	case 0x8031001A:
		return FveEAdInsufficientBuffer
	case 0x8031001B:
		return FveEConvRead
	case 0x8031001C:
		return FveEConvWrite
	case 0x8031001D:
		return FveEKeyRequired
	case 0x8031001E:
		return FveEClusteringNotSupported
	case 0x8031001F:
		return FveEVolumeBoundAlready
	case 0x80310020:
		return FveEOsNotProtected
	case 0x80310021:
		return FveEProtectionDisabled
	case 0x80310022:
		return FveERecoveryKeyRequired
	case 0x80310023:
		return FveEForeignVolume
	case 0x80310024:
		return FveEOverlappedUpdate
	case 0x80310025:
		return FveETpmSrkAuthNotZero
	case 0x80310026:
		return FveEFailedSectorSize
	case 0x80310027:
		return FveEFailedAuthentication
	case 0x80310028:
		return FveENotOsVolume
	case 0x80310029:
		return FveEAutounlockEnabled
	case 0x8031002A:
		return FveEWrongBootsector
	case 0x8031002B:
		return FveEWrongSystemFs
	case 0x8031002C:
		return FveEPolicyPasswordRequired
	case 0x8031002D:
		return FveECannotSetFvekEncrypted
	case 0x8031002E:
		return FveECannotEncryptNoKey
	case 0x80310030:
		return FveEBootableCddvd
	case 0x80310031:
		return FveEProtectorExists
	case 0x80310032:
		return FveERelativePath
	case 0x80320001:
		return FwpECalloutNotFound
	case 0x80320002:
		return FwpEConditionNotFound
	case 0x80320003:
		return FwpEFilterNotFound
	case 0x80320004:
		return FwpELayerNotFound
	case 0x80320005:
		return FwpEProviderNotFound
	case 0x80320006:
		return FwpEProviderContextNotFound
	case 0x80320007:
		return FwpESublayerNotFound
	case 0x80320008:
		return FwpENotFound
	case 0x80320009:
		return FwpEAlreadyExists
	case 0x8032000A:
		return FwpEInUse
	case 0x8032000B:
		return FwpEDynamicSessionInProgress
	case 0x8032000C:
		return FwpEWrongSession
	case 0x8032000D:
		return FwpENoTxnInProgress
	case 0x8032000E:
		return FwpETxnInProgress
	case 0x8032000F:
		return FwpETxnAborted
	case 0x80320010:
		return FwpESessionAborted
	case 0x80320011:
		return FwpEIncompatibleTxn
	case 0x80320012:
		return FwpETimeout
	case 0x80320013:
		return FwpENetEventsDisabled
	case 0x80320014:
		return FwpEIncompatibleLayer
	case 0x80320015:
		return FwpEKmClientsOnly
	case 0x80320016:
		return FwpELifetimeMismatch
	case 0x80320017:
		return FwpEBuiltinObject
	case 0x80320018:
		return FwpETooManyBoottimeFilters
	case 0x80320019:
		return FwpENotificationDropped
	case 0x8032001A:
		return FwpETrafficMismatch
	case 0x8032001B:
		return FwpEIncompatibleSaState
	case 0x8032001C:
		return FwpENullPointer
	case 0x8032001D:
		return FwpEInvalidEnumerator
	case 0x8032001E:
		return FwpEInvalidFlags
	case 0x8032001F:
		return FwpEInvalidNetMask
	case 0x80320020:
		return FwpEInvalidRange
	case 0x80320021:
		return FwpEInvalidInterval
	case 0x80320022:
		return FwpEZeroLengthArray
	case 0x80320023:
		return FwpENullDisplayName
	case 0x80320024:
		return FwpEInvalidActionType
	case 0x80320025:
		return FwpEInvalidWeight
	case 0x80320026:
		return FwpEMatchTypeMismatch
	case 0x80320027:
		return FwpETypeMismatch
	case 0x80320028:
		return FwpEOutOfBounds
	case 0x80320029:
		return FwpEReserved
	case 0x8032002A:
		return FwpEDuplicateCondition
	case 0x8032002B:
		return FwpEDuplicateKeymod
	case 0x8032002C:
		return FwpEActionIncompatibleWithLayer
	case 0x8032002D:
		return FwpEActionIncompatibleWithSublayer
	case 0x8032002E:
		return FwpEContextIncompatibleWithLayer
	case 0x8032002F:
		return FwpEContextIncompatibleWithCallout
	case 0x80320030:
		return FwpEIncompatibleAuthMethod
	case 0x80320031:
		return FwpEIncompatibleDhGroup
	case 0x80320032:
		return FwpEEmNotSupported
	case 0x80320033:
		return FwpENeverMatch
	case 0x80320034:
		return FwpEProviderContextMismatch
	case 0x80320035:
		return FwpEInvalidParameter
	case 0x80320036:
		return FwpETooManySublayers
	case 0x80320037:
		return FwpECalloutNotificationFailed
	case 0x80320038:
		return FwpEIncompatibleAuthConfig
	case 0x80320039:
		return FwpEIncompatibleCipherConfig
	case 0x80340002:
		return ErrorNdisInterfaceClosing
	case 0x80340004:
		return ErrorNdisBadVersion
	case 0x80340005:
		return ErrorNdisBadCharacteristics
	case 0x80340006:
		return ErrorNdisAdapterNotFound
	case 0x80340007:
		return ErrorNdisOpenFailed
	case 0x80340008:
		return ErrorNdisDeviceFailed
	case 0x80340009:
		return ErrorNdisMulticastFull
	case 0x8034000A:
		return ErrorNdisMulticastExists
	case 0x8034000B:
		return ErrorNdisMulticastNotFound
	case 0x8034000C:
		return ErrorNdisRequestAborted
	case 0x8034000D:
		return ErrorNdisResetInProgress
	case 0x8034000F:
		return ErrorNdisInvalidPacket
	case 0x80340010:
		return ErrorNdisInvalidDeviceRequest
	case 0x80340011:
		return ErrorNdisAdapterNotReady
	case 0x80340014:
		return ErrorNdisInvalidLength
	case 0x80340015:
		return ErrorNdisInvalidData
	case 0x80340016:
		return ErrorNdisBufferTooShort
	case 0x80340017:
		return ErrorNdisInvalidOid
	case 0x80340018:
		return ErrorNdisAdapterRemoved
	case 0x80340019:
		return ErrorNdisUnsupportedMedia
	case 0x8034001A:
		return ErrorNdisGroupAddressInUse
	case 0x8034001B:
		return ErrorNdisFileNotFound
	case 0x8034001C:
		return ErrorNdisErrorReadingFile
	case 0x8034001D:
		return ErrorNdisAlreadyMapped
	case 0x8034001E:
		return ErrorNdisResourceConflict
	case 0x8034001F:
		return ErrorNdisMediaDisconnected
	case 0x80340022:
		return ErrorNdisInvalidAddress
	case 0x8034002A:
		return ErrorNdisPaused
	case 0x8034002B:
		return ErrorNdisInterfaceNotFound
	case 0x8034002C:
		return ErrorNdisUnsupportedRevision
	case 0x8034002D:
		return ErrorNdisInvalidPort
	case 0x8034002E:
		return ErrorNdisInvalidPortState
	case 0x803400BB:
		return ErrorNdisNotSupported
	case 0x80342000:
		return ErrorNdisDot11AutoConfigEnabled
	case 0x80342001:
		return ErrorNdisDot11MediaInUse
	case 0x80342002:
		return ErrorNdisDot11PowerStateInvalid
	case 0x8DEAD01B:
		return TrkENotFound
	case 0x8DEAD01C:
		return TrkEVolumeQuotaExceeded
	case 0x8DEAD01E:
		return TrkServerTooBusy
	case 0xC0090001:
		return ErrorAuditingDisabled
	case 0xC0090002:
		return ErrorAllSidsFiltered
	case 0xC0090003:
		return ErrorBizrulesNotEnabled
	case 0xC00D0005:
		return NsENoconnection
	case 0xC00D0006:
		return NsECannotconnect
	case 0xC00D0007:
		return NsECannotdestroytitle
	case 0xC00D0008:
		return NsECannotrenametitle
	case 0xC00D0009:
		return NsECannotofflinedisk
	case 0xC00D000A:
		return NsECannotonlinedisk
	case 0xC00D000B:
		return NsENoregisteredwalker
	case 0xC00D000C:
		return NsENofunnel
	case 0xC00D000D:
		return NsENoLocalplay
	case 0xC00D000E:
		return NsENetworkBusy
	case 0xC00D000F:
		return NsETooManySess
	case 0xC00D0010:
		return NsEAlreadyConnected
	case 0xC00D0011:
		return NsEInvalidIndex
	case 0xC00D0012:
		return NsEProtocolMismatch
	case 0xC00D0013:
		return NsETimeout
	case 0xC00D0014:
		return NsENetWrite
	case 0xC00D0015:
		return NsENetRead
	case 0xC00D0016:
		return NsEDiskWrite
	case 0xC00D0017:
		return NsEDiskRead
	case 0xC00D0018:
		return NsEFileWrite
	case 0xC00D0019:
		return NsEFileRead
	case 0xC00D001A:
		return NsEFileNotFound
	case 0xC00D001B:
		return NsEFileExists
	case 0xC00D001C:
		return NsEInvalidName
	case 0xC00D001D:
		return NsEFileOpenFailed
	case 0xC00D001E:
		return NsEFileAllocationFailed
	case 0xC00D001F:
		return NsEFileInitFailed
	case 0xC00D0020:
		return NsEFilePlayFailed
	case 0xC00D0021:
		return NsESetDiskUidFailed
	case 0xC00D0022:
		return NsEInduced
	case 0xC00D0023:
		return NsECclinkDown
	case 0xC00D0024:
		return NsEInternal
	case 0xC00D0025:
		return NsEBusy
	case 0xC00D0026:
		return NsEUnrecognizedStreamType
	case 0xC00D0027:
		return NsENetworkServiceFailure
	case 0xC00D0028:
		return NsENetworkResourceFailure
	case 0xC00D0029:
		return NsEConnectionFailure
	case 0xC00D002A:
		return NsEShutdown
	case 0xC00D002B:
		return NsEInvalidRequest
	case 0xC00D002C:
		return NsEInsufficientBandwidth
	case 0xC00D002D:
		return NsENotRebuilding
	case 0xC00D002E:
		return NsELateOperation
	case 0xC00D002F:
		return NsEInvalidData
	case 0xC00D0030:
		return NsEFileBandwidthLimit
	case 0xC00D0031:
		return NsEOpenFileLimit
	case 0xC00D0032:
		return NsEBadControlData
	case 0xC00D0033:
		return NsENoStream
	case 0xC00D0034:
		return NsEStreamEnd
	case 0xC00D0035:
		return NsEServerNotFound
	case 0xC00D0036:
		return NsEDuplicateName
	case 0xC00D0037:
		return NsEDuplicateAddress
	case 0xC00D0038:
		return NsEBadMulticastAddress
	case 0xC00D0039:
		return NsEBadAdapterAddress
	case 0xC00D003A:
		return NsEBadDeliveryMode
	case 0xC00D003B:
		return NsEInvalidChannel
	case 0xC00D003C:
		return NsEInvalidStream
	case 0xC00D003D:
		return NsEInvalidArchive
	case 0xC00D003E:
		return NsENotitles
	case 0xC00D003F:
		return NsEInvalidClient
	case 0xC00D0040:
		return NsEInvalidBlackholeAddress
	case 0xC00D0041:
		return NsEIncompatibleFormat
	case 0xC00D0042:
		return NsEInvalidKey
	case 0xC00D0043:
		return NsEInvalidPort
	case 0xC00D0044:
		return NsEInvalidTtl
	case 0xC00D0045:
		return NsEStrideRefused
	case 0xC00D0046:
		return NsEMmsautoserverCantfindwalker
	case 0xC00D0047:
		return NsEMaxBitrate
	case 0xC00D0048:
		return NsELogfileperiod
	case 0xC00D0049:
		return NsEMaxClients
	case 0xC00D004A:
		return NsELogFileSize
	case 0xC00D004B:
		return NsEMaxFilerate
	case 0xC00D004C:
		return NsEWalkerUnknown
	case 0xC00D004D:
		return NsEWalkerServer
	case 0xC00D004E:
		return NsEWalkerUsage
	case 0xC00D0050:
		return NsETigerFail
	case 0xC00D0053:
		return NsECubFail
	case 0xC00D0055:
		return NsEDiskFail
	case 0xC00D0060:
		return NsEMaxFunnelsAlert
	case 0xC00D0061:
		return NsEAllocateFileFail
	case 0xC00D0062:
		return NsEPagingError
	case 0xC00D0063:
		return NsEBadBlock0Version
	case 0xC00D0064:
		return NsEBadDiskUid
	case 0xC00D0065:
		return NsEBadFsmajorVersion
	case 0xC00D0066:
		return NsEBadStampnumber
	case 0xC00D0067:
		return NsEPartiallyRebuiltDisk
	case 0xC00D0068:
		return NsEEnactplanGiveup
	case 0xC00D006A:
		return McmadmERegkeyNotFound
	case 0xC00D006B:
		return NsENoFormats
	case 0xC00D006C:
		return NsENoReferences
	case 0xC00D006D:
		return NsEWaveOpen
	case 0xC00D006F:
		return NsECannotconnectevents
	case 0xC00D0071:
		return NsENoDevice
	case 0xC00D0072:
		return NsENoSpecifiedDevice
	case 0xC00D00C8:
		return NsEMonitorGiveup
	case 0xC00D00C9:
		return NsERemirroredDisk
	case 0xC00D00CA:
		return NsEInsufficientData
	case 0xC00D00CB:
		return NsEAssert
	case 0xC00D00CC:
		return NsEBadAdapterName
	case 0xC00D00CD:
		return NsENotLicensed
	case 0xC00D00CE:
		return NsENoServerContact
	case 0xC00D00CF:
		return NsETooManyTitles
	case 0xC00D00D0:
		return NsETitleSizeExceeded
	case 0xC00D00D1:
		return NsEUdpDisabled
	case 0xC00D00D2:
		return NsETcpDisabled
	case 0xC00D00D3:
		return NsEHttpDisabled
	case 0xC00D00D4:
		return NsELicenseExpired
	case 0xC00D00D5:
		return NsETitleBitrate
	case 0xC00D00D6:
		return NsEEmptyProgramName
	case 0xC00D00D7:
		return NsEMissingChannel
	case 0xC00D00D8:
		return NsENoChannels
	case 0xC00D00D9:
		return NsEInvalidIndex2
	case 0xC00D0190:
		return NsECubFailLink
	case 0xC00D0192:
		return NsEBadCubUid
	case 0xC00D0195:
		return NsEGlitchMode
	case 0xC00D019B:
		return NsENoMediaProtocol
	case 0xC00D07F1:
		return NsENothingToDo
	case 0xC00D07F2:
		return NsENoMulticast
	case 0xC00D0BB8:
		return NsEInvalidInputFormat
	case 0xC00D0BB9:
		return NsEMsaudioNotInstalled
	case 0xC00D0BBA:
		return NsEUnexpectedMsaudioError
	case 0xC00D0BBB:
		return NsEInvalidOutputFormat
	case 0xC00D0BBC:
		return NsENotConfigured
	case 0xC00D0BBD:
		return NsEProtectedContent
	case 0xC00D0BBE:
		return NsELicenseRequired
	case 0xC00D0BBF:
		return NsETamperedContent
	case 0xC00D0BC0:
		return NsELicenseOutofdate
	case 0xC00D0BC1:
		return NsELicenseIncorrectRights
	case 0xC00D0BC2:
		return NsEAudioCodecNotInstalled
	case 0xC00D0BC3:
		return NsEAudioCodecError
	case 0xC00D0BC4:
		return NsEVideoCodecNotInstalled
	case 0xC00D0BC5:
		return NsEVideoCodecError
	case 0xC00D0BC6:
		return NsEInvalidprofile
	case 0xC00D0BC7:
		return NsEIncompatibleVersion
	case 0xC00D0BCA:
		return NsEOfflineMode
	case 0xC00D0BCB:
		return NsENotConnected
	case 0xC00D0BCC:
		return NsETooMuchData
	case 0xC00D0BCD:
		return NsEUnsupportedProperty
	case 0xC00D0BCE:
		return NsE8bitWaveUnsupported
	case 0xC00D0BCF:
		return NsENoMoreSamples
	case 0xC00D0BD0:
		return NsEInvalidSamplingRate
	case 0xC00D0BD1:
		return NsEMaxPacketSizeTooSmall
	case 0xC00D0BD2:
		return NsELatePacket
	case 0xC00D0BD3:
		return NsEDuplicatePacket
	case 0xC00D0BD4:
		return NsESdkBuffertoosmall
	case 0xC00D0BD5:
		return NsEInvalidNumPasses
	case 0xC00D0BD6:
		return NsEAttributeReadOnly
	case 0xC00D0BD7:
		return NsEAttributeNotAllowed
	case 0xC00D0BD8:
		return NsEInvalidEdl
	case 0xC00D0BD9:
		return NsEDataUnitExtensionTooLarge
	case 0xC00D0BDA:
		return NsECodecDmoError
	case 0xC00D0BDC:
		return NsEFeatureDisabledByGroupPolicy
	case 0xC00D0BDD:
		return NsEFeatureDisabledInSku
	case 0xC00D0FA0:
		return NsENoCd
	case 0xC00D0FA1:
		return NsECantReadDigital
	case 0xC00D0FA2:
		return NsEDeviceDisconnected
	case 0xC00D0FA3:
		return NsEDeviceNotSupportFormat
	case 0xC00D0FA4:
		return NsESlowReadDigital
	case 0xC00D0FA5:
		return NsEMixerInvalidLine
	case 0xC00D0FA6:
		return NsEMixerInvalidControl
	case 0xC00D0FA7:
		return NsEMixerInvalidValue
	case 0xC00D0FA8:
		return NsEMixerUnknownMmresult
	case 0xC00D0FA9:
		return NsEUserStop
	case 0xC00D0FAA:
		return NsEMp3FormatNotFound
	case 0xC00D0FAB:
		return NsECdReadErrorNoCorrection
	case 0xC00D0FAC:
		return NsECdReadError
	case 0xC00D0FAD:
		return NsECdSlowCopy
	case 0xC00D0FAE:
		return NsECdCopytoCd
	case 0xC00D0FAF:
		return NsEMixerNodriver
	case 0xC00D0FB0:
		return NsERedbookEnabledWhileCopying
	case 0xC00D0FB1:
		return NsECdRefresh
	case 0xC00D0FB2:
		return NsECdDriverProblem
	case 0xC00D0FB3:
		return NsEWontDoDigital
	case 0xC00D0FB4:
		return NsEWmpxmlNoerror
	case 0xC00D0FB5:
		return NsEWmpxmlEndofdata
	case 0xC00D0FB6:
		return NsEWmpxmlParseerror
	case 0xC00D0FB7:
		return NsEWmpxmlAttributenotfound
	case 0xC00D0FB8:
		return NsEWmpxmlPinotfound
	case 0xC00D0FB9:
		return NsEWmpxmlEmptydoc
	case 0xC00D0FBA:
		return NsEWmpPathAlreadyInLibrary
	case 0xC00D0FBE:
		return NsEWmpFilescanalreadystarted
	case 0xC00D0FBF:
		return NsEWmpHmeInvalidobjectid
	case 0xC00D0FC0:
		return NsEWmpMfCodeExpired
	case 0xC00D0FC1:
		return NsEWmpHmeNotsearchableforitems
	case 0xC00D0FC7:
		return NsEWmpAddtolibraryFailed
	case 0xC00D0FC8:
		return NsEWmpWindowsapifailure
	case 0xC00D0FC9:
		return NsEWmpRecordingNotAllowed
	case 0xC00D0FCA:
		return NsEDeviceNotReady
	case 0xC00D0FCB:
		return NsEDamagedFile
	case 0xC00D0FCC:
		return NsEMpdbGeneric
	case 0xC00D0FCD:
		return NsEFileFailedChecks
	case 0xC00D0FCE:
		return NsEMediaLibraryFailed
	case 0xC00D0FCF:
		return NsESharingViolation
	case 0xC00D0FD0:
		return NsENoErrorStringFound
	case 0xC00D0FD1:
		return NsEWmpocxNoRemoteCore
	case 0xC00D0FD2:
		return NsEWmpocxNoActiveCore
	case 0xC00D0FD3:
		return NsEWmpocxNotRunningRemotely
	case 0xC00D0FD4:
		return NsEWmpocxNoRemoteWindow
	case 0xC00D0FD5:
		return NsEWmpocxErrormanagernotavailable
	case 0xC00D0FD6:
		return NsEPluginNotshutdown
	case 0xC00D0FD7:
		return NsEWmpCannotFindFolder
	case 0xC00D0FD8:
		return NsEWmpStreamingRecordingNotAllowed
	case 0xC00D0FD9:
		return NsEWmpPlugindllNotfound
	case 0xC00D0FDA:
		return NsENeedToAskUser
	case 0xC00D0FDB:
		return NsEWmpocxPlayerNotDocked
	case 0xC00D0FDC:
		return NsEWmpExternalNotready
	case 0xC00D0FDD:
		return NsEWmpMlsStaleData
	case 0xC00D0FDE:
		return NsEWmpUiSubcontrolsnotsupported
	case 0xC00D0FDF:
		return NsEWmpUiVersionmismatch
	case 0xC00D0FE0:
		return NsEWmpUiNotathemefile
	case 0xC00D0FE1:
		return NsEWmpUiSubelementnotfound
	case 0xC00D0FE2:
		return NsEWmpUiVersionparse
	case 0xC00D0FE3:
		return NsEWmpUiViewidnotfound
	case 0xC00D0FE4:
		return NsEWmpUiPassthrough
	case 0xC00D0FE5:
		return NsEWmpUiObjectnotfound
	case 0xC00D0FE6:
		return NsEWmpUiSecondhandler
	case 0xC00D0FE7:
		return NsEWmpUiNoskininzip
	case 0xC00D0FEA:
		return NsEWmpUrldownloadfailed
	case 0xC00D0FEB:
		return NsEWmpocxUnableToLoadSkin
	case 0xC00D0FEC:
		return NsEWmpInvalidSkin
	case 0xC00D0FED:
		return NsEWmpSendmailfailed
	case 0xC00D0FEE:
		return NsEWmpLockedinskinmode
	case 0xC00D0FEF:
		return NsEWmpFailedToSaveFile
	case 0xC00D0FF0:
		return NsEWmpSaveasReadonly
	case 0xC00D0FF1:
		return NsEWmpFailedToSavePlaylist
	case 0xC00D0FF2:
		return NsEWmpFailedToOpenWmd
	case 0xC00D0FF3:
		return NsEWmpCantPlayProtected
	case 0xC00D0FF4:
		return NsESharingStateOutOfSync
	case 0xC00D0FFA:
		return NsEWmpocxRemotePlayerAlreadyRunning
	case 0xC00D1004:
		return NsEWmpRbcJpgmappingimage
	case 0xC00D1005:
		return NsEWmpJpgtransparency
	case 0xC00D1009:
		return NsEWmpInvalidMaxVal
	case 0xC00D100A:
		return NsEWmpInvalidMinVal
	case 0xC00D100E:
		return NsEWmpCsJpgpositionimage
	case 0xC00D100F:
		return NsEWmpCsNotevenlydivisible
	case 0xC00D1018:
		return NsEWmpzipNotazipfile
	case 0xC00D1019:
		return NsEWmpzipCorrupt
	case 0xC00D101A:
		return NsEWmpzipFilenotfound
	case 0xC00D1022:
		return NsEWmpImageFiletypeUnsupported
	case 0xC00D1023:
		return NsEWmpImageInvalidFormat
	case 0xC00D1024:
		return NsEWmpGifUnexpectedEndoffile
	case 0xC00D1025:
		return NsEWmpGifInvalidFormat
	case 0xC00D1026:
		return NsEWmpGifBadVersionNumber
	case 0xC00D1027:
		return NsEWmpGifNoImageInFile
	case 0xC00D1028:
		return NsEWmpPngInvalidformat
	case 0xC00D1029:
		return NsEWmpPngUnsupportedBitdepth
	case 0xC00D102A:
		return NsEWmpPngUnsupportedCompression
	case 0xC00D102B:
		return NsEWmpPngUnsupportedFilter
	case 0xC00D102C:
		return NsEWmpPngUnsupportedInterlace
	case 0xC00D102D:
		return NsEWmpPngUnsupportedBadCrc
	case 0xC00D102E:
		return NsEWmpBmpInvalidBitmask
	case 0xC00D102F:
		return NsEWmpBmpTopdownDibUnsupported
	case 0xC00D1030:
		return NsEWmpBmpBitmapNotCreated
	case 0xC00D1031:
		return NsEWmpBmpCompressionUnsupported
	case 0xC00D1032:
		return NsEWmpBmpInvalidFormat
	case 0xC00D1033:
		return NsEWmpJpgJerrArithcodingNotimpl
	case 0xC00D1034:
		return NsEWmpJpgInvalidFormat
	case 0xC00D1035:
		return NsEWmpJpgBadDctsize
	case 0xC00D1036:
		return NsEWmpJpgBadVersionNumber
	case 0xC00D1037:
		return NsEWmpJpgBadPrecision
	case 0xC00D1038:
		return NsEWmpJpgCcir601Notimpl
	case 0xC00D1039:
		return NsEWmpJpgNoImageInFile
	case 0xC00D103A:
		return NsEWmpJpgReadError
	case 0xC00D103B:
		return NsEWmpJpgFractSampleNotimpl
	case 0xC00D103C:
		return NsEWmpJpgImageTooBig
	case 0xC00D103D:
		return NsEWmpJpgUnexpectedEndoffile
	case 0xC00D103E:
		return NsEWmpJpgSofUnsupported
	case 0xC00D103F:
		return NsEWmpJpgUnknownMarker
	case 0xC00D1044:
		return NsEWmpFailedToOpenImage
	case 0xC00D1049:
		return NsEWmpDaiSongtooshort
	case 0xC00D104A:
		return NsEWmgRateunavailable
	case 0xC00D104B:
		return NsEWmgPluginunavailable
	case 0xC00D104C:
		return NsEWmgCannotqueue
	case 0xC00D104D:
		return NsEWmgPrerolllicenseacquisitionnotallowed
	case 0xC00D104E:
		return NsEWmgUnexpectedprerollstatus
	case 0xC00D1051:
		return NsEWmgInvalidCoppCertificate
	case 0xC00D1052:
		return NsEWmgCoppSecurityInvalid
	case 0xC00D1053:
		return NsEWmgCoppUnsupported
	case 0xC00D1054:
		return NsEWmgInvalidstate
	case 0xC00D1055:
		return NsEWmgSinkalreadyexists
	case 0xC00D1056:
		return NsEWmgNosdkinterface
	case 0xC00D1057:
		return NsEWmgNotalloutputsrendered
	case 0xC00D1058:
		return NsEWmgFiletransfernotallowed
	case 0xC00D1059:
		return NsEWmrUnsupportedstream
	case 0xC00D105A:
		return NsEWmrPinnotfound
	case 0xC00D105B:
		return NsEWmrWaitingonformatswitch
	case 0xC00D105C:
		return NsEWmrNosourcefilter
	case 0xC00D105D:
		return NsEWmrPintypenomatch
	case 0xC00D105E:
		return NsEWmrNocallbackavailable
	case 0xC00D1062:
		return NsEWmrSamplepropertynotset
	case 0xC00D1063:
		return NsEWmrCannotRenderBinaryStream
	case 0xC00D1064:
		return NsEWmgLicenseTampered
	case 0xC00D1065:
		return NsEWmrWillnotRenderBinaryStream
	case 0xC00D1068:
		return NsEWmxUnrecognizedPlaylistFormat
	case 0xC00D1069:
		return NsEAsxInvalidformat
	case 0xC00D106A:
		return NsEAsxInvalidversion
	case 0xC00D106B:
		return NsEAsxInvalidRepeatBlock
	case 0xC00D106C:
		return NsEAsxNothingToWrite
	case 0xC00D106D:
		return NsEUrllistInvalidformat
	case 0xC00D106E:
		return NsEWmxAttributeDoesNotExist
	case 0xC00D106F:
		return NsEWmxAttributeAlreadyExists
	case 0xC00D1070:
		return NsEWmxAttributeUnretrievable
	case 0xC00D1071:
		return NsEWmxItemDoesNotExist
	case 0xC00D1072:
		return NsEWmxItemTypeIllegal
	case 0xC00D1073:
		return NsEWmxItemUnsettable
	case 0xC00D1074:
		return NsEWmxPlaylistEmpty
	case 0xC00D1075:
		return NsEMlsSmartplaylistFilterNotRegistered
	case 0xC00D1076:
		return NsEWmxInvalidFormatOverNesting
	case 0xC00D107C:
		return NsEWmpcoreNosourceurlstring
	case 0xC00D107D:
		return NsEWmpcoreCocreatefailedforgitobject
	case 0xC00D107E:
		return NsEWmpcoreFailedtogetmarshalledeventhandlerinterface
	case 0xC00D107F:
		return NsEWmpcoreBuffertoosmall
	case 0xC00D1080:
		return NsEWmpcoreUnavailable
	case 0xC00D1081:
		return NsEWmpcoreInvalidplaylistmode
	case 0xC00D1086:
		return NsEWmpcoreItemnotinplaylist
	case 0xC00D1087:
		return NsEWmpcorePlaylistempty
	case 0xC00D1088:
		return NsEWmpcoreNobrowser
	case 0xC00D1089:
		return NsEWmpcoreUnrecognizedMediaUrl
	case 0xC00D108A:
		return NsEWmpcoreGraphNotInList
	case 0xC00D108B:
		return NsEWmpcorePlaylistEmptyOrSingleMedia
	case 0xC00D108C:
		return NsEWmpcoreErrorsinknotregistered
	case 0xC00D108D:
		return NsEWmpcoreErrormanagernotavailable
	case 0xC00D108E:
		return NsEWmpcoreWebhelpfailed
	case 0xC00D108F:
		return NsEWmpcoreMediaErrorResumeFailed
	case 0xC00D1090:
		return NsEWmpcoreNoRefInEntry
	case 0xC00D1091:
		return NsEWmpcoreWmxListAttributeNameEmpty
	case 0xC00D1092:
		return NsEWmpcoreWmxListAttributeNameIllegal
	case 0xC00D1093:
		return NsEWmpcoreWmxListAttributeValueEmpty
	case 0xC00D1094:
		return NsEWmpcoreWmxListAttributeValueIllegal
	case 0xC00D1095:
		return NsEWmpcoreWmxListItemAttributeNameEmpty
	case 0xC00D1096:
		return NsEWmpcoreWmxListItemAttributeNameIllegal
	case 0xC00D1097:
		return NsEWmpcoreWmxListItemAttributeValueEmpty
	case 0xC00D1098:
		return NsEWmpcoreListEntryNoRef
	case 0xC00D1099:
		return NsEWmpcoreMisnamedFile
	case 0xC00D109A:
		return NsEWmpcoreCodecNotTrusted
	case 0xC00D109B:
		return NsEWmpcoreCodecNotFound
	case 0xC00D109C:
		return NsEWmpcoreCodecDownloadNotAllowed
	case 0xC00D109D:
		return NsEWmpcoreErrorDownloadingPlaylist
	case 0xC00D109E:
		return NsEWmpcoreFailedToBuildPlaylist
	case 0xC00D109F:
		return NsEWmpcorePlaylistItemAlternateNone
	case 0xC00D10A0:
		return NsEWmpcorePlaylistItemAlternateExhausted
	case 0xC00D10A1:
		return NsEWmpcorePlaylistItemAlternateNameNotFound
	case 0xC00D10A2:
		return NsEWmpcorePlaylistItemAlternateMorphFailed
	case 0xC00D10A3:
		return NsEWmpcorePlaylistItemAlternateInitFailed
	case 0xC00D10A4:
		return NsEWmpcoreMediaAlternateRefEmpty
	case 0xC00D10A5:
		return NsEWmpcorePlaylistNoEventName
	case 0xC00D10A6:
		return NsEWmpcorePlaylistEventAttributeAbsent
	case 0xC00D10A7:
		return NsEWmpcorePlaylistEventEmpty
	case 0xC00D10A8:
		return NsEWmpcorePlaylistStackEmpty
	case 0xC00D10A9:
		return NsEWmpcoreCurrentMediaNotActive
	case 0xC00D10AB:
		return NsEWmpcoreUserCancel
	case 0xC00D10AC:
		return NsEWmpcorePlaylistRepeatEmpty
	case 0xC00D10AD:
		return NsEWmpcorePlaylistRepeatStartMediaNone
	case 0xC00D10AE:
		return NsEWmpcorePlaylistRepeatEndMediaNone
	case 0xC00D10AF:
		return NsEWmpcoreInvalidPlaylistUrl
	case 0xC00D10B0:
		return NsEWmpcoreMismatchedRuntime
	case 0xC00D10B1:
		return NsEWmpcorePlaylistImportFailedNoItems
	case 0xC00D10B2:
		return NsEWmpcoreVideoTransformFilterInsertion
	case 0xC00D10B3:
		return NsEWmpcoreMediaUnavailable
	case 0xC00D10B4:
		return NsEWmpcoreWmxEntryrefNoRef
	case 0xC00D10B5:
		return NsEWmpcoreNoPlayableMediaInPlaylist
	case 0xC00D10B6:
		return NsEWmpcorePlaylistEmptyNestedPlaylistSkippedItems
	case 0xC00D10B7:
		return NsEWmpcoreBusy
	case 0xC00D10B8:
		return NsEWmpcoreMediaChildPlaylistUnavailable
	case 0xC00D10B9:
		return NsEWmpcoreMediaNoChildPlaylist
	case 0xC00D10BA:
		return NsEWmpcoreFileNotFound
	case 0xC00D10BB:
		return NsEWmpcoreTempFileNotFound
	case 0xC00D10BC:
		return NsEWmdmRevoked
	case 0xC00D10BD:
		return NsEDdrawGeneric
	case 0xC00D10BE:
		return NsEDisplayModeChangeFailed
	case 0xC00D10BF:
		return NsEPlaylistContainsErrors
	case 0xC00D10C0:
		return NsEChangingProxyName
	case 0xC00D10C1:
		return NsEChangingProxyPort
	case 0xC00D10C2:
		return NsEChangingProxyExceptionlist
	case 0xC00D10C3:
		return NsEChangingProxybypass
	case 0xC00D10C4:
		return NsEChangingProxyProtocolNotFound
	case 0xC00D10C5:
		return NsEGraphNoaudiolanguage
	case 0xC00D10C6:
		return NsEGraphNoaudiolanguageselected
	case 0xC00D10C7:
		return NsECorecdNotamediacd
	case 0xC00D10C8:
		return NsEWmpcoreMediaUrlTooLong
	case 0xC00D10C9:
		return NsEWmpflashCantFindComServer
	case 0xC00D10CA:
		return NsEWmpflashIncompatibleversion
	case 0xC00D10CB:
		return NsEWmpocxgraphIeDisallowsActivexControls
	case 0xC00D10CC:
		return NsENeedCoreReference
	case 0xC00D10CD:
		return NsEMediacdReadError
	case 0xC00D10CE:
		return NsEIeDisallowsActivexControls
	case 0xC00D10CF:
		return NsEFlashPlaybackNotAllowed
	case 0xC00D10D0:
		return NsEUnableToCreateRipLocation
	case 0xC00D10D1:
		return NsEWmpcoreSomeCodecsMissing
	case 0xC00D10D2:
		return NsEWmpRipFailed
	case 0xC00D10D3:
		return NsEWmpFailedToRipTrack
	case 0xC00D10D4:
		return NsEWmpEraseFailed
	case 0xC00D10D5:
		return NsEWmpFormatFailed
	case 0xC00D10D6:
		return NsEWmpCannotBurnNonLocalFile
	case 0xC00D10D7:
		return NsEWmpFileTypeCannotBurnToAudioCd
	case 0xC00D10D8:
		return NsEWmpFileDoesNotFitOnCd
	case 0xC00D10D9:
		return NsEWmpFileNoDuration
	case 0xC00D10DA:
		return NsEPdaFailedToBurn
	case 0xC00D10DC:
		return NsEFailedDownloadAbortBurn
	case 0xC00D10DD:
		return NsEWmpcoreDeviceDriversMissing
	case 0xC00D1126:
		return NsEWmpimUseroffline
	case 0xC00D1127:
		return NsEWmpimUsercanceled
	case 0xC00D1128:
		return NsEWmpimDialupfailed
	case 0xC00D1129:
		return NsEWinsockErrorString
	case 0xC00D1130:
		return NsEWmpbrNolistener
	case 0xC00D1131:
		return NsEWmpbrBackupcancel
	case 0xC00D1132:
		return NsEWmpbrRestorecancel
	case 0xC00D1133:
		return NsEWmpbrErrorwithurl
	case 0xC00D1134:
		return NsEWmpbrNamecollision
	case 0xC00D1137:
		return NsEWmpbrDriveInvalid
	case 0xC00D1138:
		return NsEWmpbrBackuprestorefailed
	case 0xC00D1158:
		return NsEWmpConvertFileFailed
	case 0xC00D1159:
		return NsEWmpConvertNoRightsErrorurl
	case 0xC00D115A:
		return NsEWmpConvertNoRightsNoerrorurl
	case 0xC00D115B:
		return NsEWmpConvertFileCorrupt
	case 0xC00D115C:
		return NsEWmpConvertPluginUnavailableErrorurl
	case 0xC00D115D:
		return NsEWmpConvertPluginUnavailableNoerrorurl
	case 0xC00D115E:
		return NsEWmpConvertPluginUnknownFileOwner
	case 0xC00D1160:
		return NsEDvdDiscCopyProtectOutputNs
	case 0xC00D1161:
		return NsEDvdDiscCopyProtectOutputFailed
	case 0xC00D1162:
		return NsEDvdNoSubpictureStream
	case 0xC00D1163:
		return NsEDvdCopyProtect
	case 0xC00D1164:
		return NsEDvdAuthoringProblem
	case 0xC00D1165:
		return NsEDvdInvalidDiscRegion
	case 0xC00D1166:
		return NsEDvdCompatibleVideoCard
	case 0xC00D1167:
		return NsEDvdMacrovision
	case 0xC00D1168:
		return NsEDvdSystemDecoderRegion
	case 0xC00D1169:
		return NsEDvdDiscDecoderRegion
	case 0xC00D116A:
		return NsEDvdNoVideoStream
	case 0xC00D116B:
		return NsEDvdNoAudioStream
	case 0xC00D116C:
		return NsEDvdGraphBuilding
	case 0xC00D116D:
		return NsEDvdNoDecoder
	case 0xC00D116E:
		return NsEDvdParental
	case 0xC00D116F:
		return NsEDvdCannotJump
	case 0xC00D1170:
		return NsEDvdDeviceContention
	case 0xC00D1171:
		return NsEDvdNoVideoMemory
	case 0xC00D1172:
		return NsEDvdCannotCopyProtected
	case 0xC00D1173:
		return NsEDvdRequiredPropertyNotSet
	case 0xC00D1174:
		return NsEDvdInvalidTitleChapter
	case 0xC00D1176:
		return NsENoCdBurner
	case 0xC00D1177:
		return NsEDeviceIsNotReady
	case 0xC00D1178:
		return NsEPdaUnsupportedFormat
	case 0xC00D1179:
		return NsENoPda
	case 0xC00D117A:
		return NsEPdaUnspecifiedError
	case 0xC00D117B:
		return NsEMemstorageBadData
	case 0xC00D117C:
		return NsEPdaFailSelectDevice
	case 0xC00D117D:
		return NsEPdaFailReadWaveFile
	case 0xC00D117E:
		return NsEImapiLossofstreaming
	case 0xC00D117F:
		return NsEPdaDeviceFull
	case 0xC00D1180:
		return NsEFailLaunchRoxioPlugin
	case 0xC00D1181:
		return NsEPdaDeviceFullInSession
	case 0xC00D1182:
		return NsEImapiMediumInvalidtype
	case 0xC00D1183:
		return NsEPdaManualdevice
	case 0xC00D1184:
		return NsEPdaPartnershipnotexist
	case 0xC00D1185:
		return NsEPdaCannotCreateAdditionalSyncRelationship
	case 0xC00D1186:
		return NsEPdaNoTranscodeOfDrm
	case 0xC00D1187:
		return NsEPdaTranscodecachefull
	case 0xC00D1188:
		return NsEPdaTooManyFileCollisions
	case 0xC00D1189:
		return NsEPdaCannotTranscode
	case 0xC00D118A:
		return NsEPdaTooManyFilesInDirectory
	case 0xC00D118B:
		return NsEProcessingshowsyncwizard
	case 0xC00D118C:
		return NsEPdaTranscodeNotPermitted
	case 0xC00D118D:
		return NsEPdaInitializingdevices
	case 0xC00D118E:
		return NsEPdaObsoleteSp
	case 0xC00D118F:
		return NsEPdaTitleCollision
	case 0xC00D1190:
		return NsEPdaDevicesupportdisabled
	case 0xC00D1191:
		return NsEPdaNoLongerAvailable
	case 0xC00D1192:
		return NsEPdaEncoderNotResponding
	case 0xC00D1193:
		return NsEPdaCannotSyncFromLocation
	case 0xC00D1194:
		return NsEWmpProtocolProblem
	case 0xC00D1195:
		return NsEWmpNoDiskSpace
	case 0xC00D1196:
		return NsEWmpLogonFailure
	case 0xC00D1197:
		return NsEWmpCannotFindFile
	case 0xC00D1198:
		return NsEWmpServerInaccessible
	case 0xC00D1199:
		return NsEWmpUnsupportedFormat
	case 0xC00D119A:
		return NsEWmpDshowUnsupportedFormat
	case 0xC00D119B:
		return NsEWmpPlaylistExists
	case 0xC00D119C:
		return NsEWmpNonmediaFiles
	case 0xC00D119D:
		return NsEWmpInvalidAsx
	case 0xC00D119E:
		return NsEWmpAlreadyInUse
	case 0xC00D119F:
		return NsEWmpImapiFailure
	case 0xC00D11A0:
		return NsEWmpWmdmFailure
	case 0xC00D11A1:
		return NsEWmpCodecNeededWith4cc
	case 0xC00D11A2:
		return NsEWmpCodecNeededWithFormattag
	case 0xC00D11A3:
		return NsEWmpMssapNotAvailable
	case 0xC00D11A4:
		return NsEWmpWmdmInterfacedead
	case 0xC00D11A5:
		return NsEWmpWmdmNotcertified
	case 0xC00D11A6:
		return NsEWmpWmdmLicenseNotexist
	case 0xC00D11A7:
		return NsEWmpWmdmLicenseExpired
	case 0xC00D11A8:
		return NsEWmpWmdmBusy
	case 0xC00D11A9:
		return NsEWmpWmdmNorights
	case 0xC00D11AA:
		return NsEWmpWmdmIncorrectRights
	case 0xC00D11AB:
		return NsEWmpImapiGeneric
	case 0xC00D11AD:
		return NsEWmpImapiDeviceNotpresent
	case 0xC00D11AE:
		return NsEWmpImapiDeviceBusy
	case 0xC00D11AF:
		return NsEWmpImapiLossOfStreaming
	case 0xC00D11B0:
		return NsEWmpServerUnavailable
	case 0xC00D11B1:
		return NsEWmpFileOpenFailed
	case 0xC00D11B2:
		return NsEWmpVerifyOnline
	case 0xC00D11B3:
		return NsEWmpServerNotResponding
	case 0xC00D11B4:
		return NsEWmpDrmCorruptBackup
	case 0xC00D11B5:
		return NsEWmpDrmLicenseServerUnavailable
	case 0xC00D11B6:
		return NsEWmpNetworkFirewall
	case 0xC00D11B7:
		return NsEWmpNoRemovableMedia
	case 0xC00D11B8:
		return NsEWmpProxyConnectTimeout
	case 0xC00D11B9:
		return NsEWmpNeedUpgrade
	case 0xC00D11BA:
		return NsEWmpAudioHwProblem
	case 0xC00D11BB:
		return NsEWmpInvalidProtocol
	case 0xC00D11BC:
		return NsEWmpInvalidLibraryAdd
	case 0xC00D11BD:
		return NsEWmpMmsNotSupported
	case 0xC00D11BE:
		return NsEWmpNoProtocolsSelected
	case 0xC00D11BF:
		return NsEWmpGofullscreenFailed
	case 0xC00D11C0:
		return NsEWmpNetworkError
	case 0xC00D11C1:
		return NsEWmpConnectTimeout
	case 0xC00D11C2:
		return NsEWmpMulticastDisabled
	case 0xC00D11C3:
		return NsEWmpServerDnsTimeout
	case 0xC00D11C4:
		return NsEWmpProxyNotFound
	case 0xC00D11C5:
		return NsEWmpTamperedContent
	case 0xC00D11C6:
		return NsEWmpOutofmemory
	case 0xC00D11C7:
		return NsEWmpAudioCodecNotInstalled
	case 0xC00D11C8:
		return NsEWmpVideoCodecNotInstalled
	case 0xC00D11C9:
		return NsEWmpImapiDeviceInvalidtype
	case 0xC00D11CA:
		return NsEWmpDrmDriverAuthFailure
	case 0xC00D11CB:
		return NsEWmpNetworkResourceFailure
	case 0xC00D11CC:
		return NsEWmpUpgradeApplication
	case 0xC00D11CD:
		return NsEWmpUnknownError
	case 0xC00D11CE:
		return NsEWmpInvalidKey
	case 0xC00D11CF:
		return NsEWmpCdAnotherUser
	case 0xC00D11D0:
		return NsEWmpDrmNeedsAuthorization
	case 0xC00D11D1:
		return NsEWmpBadDriver
	case 0xC00D11D2:
		return NsEWmpAccessDenied
	case 0xC00D11D3:
		return NsEWmpLicenseRestricts
	case 0xC00D11D4:
		return NsEWmpInvalidRequest
	case 0xC00D11D5:
		return NsEWmpCdStashNoSpace
	case 0xC00D11D6:
		return NsEWmpDrmNewHardware
	case 0xC00D11D7:
		return NsEWmpDrmInvalidSig
	case 0xC00D11D8:
		return NsEWmpDrmCannotRestore
	case 0xC00D11D9:
		return NsEWmpBurnDiscOverflow
	case 0xC00D11DA:
		return NsEWmpDrmGenericLicenseFailure
	case 0xC00D11DB:
		return NsEWmpDrmNoSecureClock
	case 0xC00D11DC:
		return NsEWmpDrmNoRights
	case 0xC00D11DD:
		return NsEWmpDrmIndivFailed
	case 0xC00D11DE:
		return NsEWmpServerNonewconnections
	case 0xC00D11DF:
		return NsEWmpMultipleErrorInPlaylist
	case 0xC00D11E0:
		return NsEWmpImapi2EraseFail
	case 0xC00D11E1:
		return NsEWmpImapi2EraseDeviceBusy
	case 0xC00D11E2:
		return NsEWmpDrmComponentFailure
	case 0xC00D11E3:
		return NsEWmpDrmNoDeviceCert
	case 0xC00D11E4:
		return NsEWmpServerSecurityError
	case 0xC00D11E5:
		return NsEWmpAudioDeviceLost
	case 0xC00D11E6:
		return NsEWmpImapiMediaIncompatible
	case 0xC00D11EE:
		return NsESyncwizDeviceFull
	case 0xC00D11EF:
		return NsESyncwizCannotChangeSettings
	case 0xC00D11F0:
		return NsETranscodeDeletecacheerror
	case 0xC00D11F8:
		return NsECdNoBuffersRead
	case 0xC00D11F9:
		return NsECdEmptyTrackQueue
	case 0xC00D11FA:
		return NsECdNoReader
	case 0xC00D11FB:
		return NsECdIsrcInvalid
	case 0xC00D11FC:
		return NsECdMediaCatalogNumberInvalid
	case 0xC00D11FD:
		return NsESlowReadDigitalWithErrorcorrection
	case 0xC00D11FE:
		return NsECdSpeeddetectNotEnoughReads
	case 0xC00D11FF:
		return NsECdQueueingDisabled
	case 0xC00D1202:
		return NsEWmpDrmAcquiringLicense
	case 0xC00D1203:
		return NsEWmpDrmLicenseExpired
	case 0xC00D1204:
		return NsEWmpDrmLicenseNotacquired
	case 0xC00D1205:
		return NsEWmpDrmLicenseNotenabled
	case 0xC00D1206:
		return NsEWmpDrmLicenseUnusable
	case 0xC00D1207:
		return NsEWmpDrmLicenseContentRevoked
	case 0xC00D1208:
		return NsEWmpDrmLicenseNosap
	case 0xC00D1209:
		return NsEWmpDrmUnableToAcquireLicense
	case 0xC00D120A:
		return NsEWmpLicenseRequired
	case 0xC00D120B:
		return NsEWmpProtectedContent
	case 0xC00D122A:
		return NsEWmpPolicyValueNotConfigured
	case 0xC00D1234:
		return NsEPdaCannotSyncFromInternet
	case 0xC00D1235:
		return NsEPdaCannotSyncInvalidPlaylist
	case 0xC00D1236:
		return NsEPdaFailedToSynchronizeFile
	case 0xC00D1237:
		return NsEPdaSyncFailed
	case 0xC00D1238:
		return NsEPdaDeleteFailed
	case 0xC00D1239:
		return NsEPdaFailedToRetrieveFile
	case 0xC00D123A:
		return NsEPdaDeviceNotResponding
	case 0xC00D123B:
		return NsEPdaFailedToTranscodePhoto
	case 0xC00D123C:
		return NsEPdaFailedToEncryptTranscodedFile
	case 0xC00D123D:
		return NsEPdaCannotTranscodeToAudio
	case 0xC00D123E:
		return NsEPdaCannotTranscodeToVideo
	case 0xC00D123F:
		return NsEPdaCannotTranscodeToImage
	case 0xC00D1240:
		return NsEPdaRetrievedFileFilenameTooLong
	case 0xC00D1241:
		return NsEPdaCewmdmDrmError
	case 0xC00D1242:
		return NsEIncompletePlaylist
	case 0xC00D1243:
		return NsEPdaSyncRunning
	case 0xC00D1244:
		return NsEPdaSyncLoginError
	case 0xC00D1245:
		return NsEPdaTranscodeCodecNotFound
	case 0xC00D1246:
		return NsECannotSyncDrmToNonJanusDevice
	case 0xC00D1247:
		return NsECannotSyncPreviousSyncRunning
	case 0xC00D125C:
		return NsEWmpHwndNotfound
	case 0xC00D125D:
		return NsEBkgdownloadWrongNoFiles
	case 0xC00D125E:
		return NsEBkgdownloadCompletecancelledjob
	case 0xC00D125F:
		return NsEBkgdownloadCancelcompletedjob
	case 0xC00D1260:
		return NsEBkgdownloadNojobpointer
	case 0xC00D1261:
		return NsEBkgdownloadInvalidjobsignature
	case 0xC00D1262:
		return NsEBkgdownloadFailedToCreateTempfile
	case 0xC00D1263:
		return NsEBkgdownloadPluginFailedinitialize
	case 0xC00D1264:
		return NsEBkgdownloadPluginFailedtomovefile
	case 0xC00D1265:
		return NsEBkgdownloadCallfuncfailed
	case 0xC00D1266:
		return NsEBkgdownloadCallfunctimeout
	case 0xC00D1267:
		return NsEBkgdownloadCallfuncended
	case 0xC00D1268:
		return NsEBkgdownloadWmdunpackfailed
	case 0xC00D1269:
		return NsEBkgdownloadFailedinitialize
	case 0xC00D126A:
		return NsEInterfaceNotRegisteredInGit
	case 0xC00D126B:
		return NsEBkgdownloadInvalidFileName
	case 0xC00D128E:
		return NsEImageDownloadFailed
	case 0xC00D12C0:
		return NsEWmpUdrmNouserlist
	case 0xC00D12C1:
		return NsEWmpDrmNotAcquiring
	case 0xC00D12F2:
		return NsEWmpBstrTooLong
	case 0xC00D12FC:
		return NsEWmpAutoplayInvalidState
	case 0xC00D1306:
		return NsEWmpComponentRevoked
	case 0xC00D1324:
		return NsECurlNotsafe
	case 0xC00D1325:
		return NsECurlInvalidchar
	case 0xC00D1326:
		return NsECurlInvalidhostname
	case 0xC00D1327:
		return NsECurlInvalidpath
	case 0xC00D1328:
		return NsECurlInvalidscheme
	case 0xC00D1329:
		return NsECurlInvalidurl
	case 0xC00D132B:
		return NsECurlCantwalk
	case 0xC00D132C:
		return NsECurlInvalidport
	case 0xC00D132D:
		return NsECurlhelperNotadirectory
	case 0xC00D132E:
		return NsECurlhelperNotafile
	case 0xC00D132F:
		return NsECurlCantdecode
	case 0xC00D1330:
		return NsECurlhelperNotrelative
	case 0xC00D1331:
		return NsECurlInvalidbuffersize
	case 0xC00D1356:
		return NsESubscriptionservicePlaybackDisallowed
	case 0xC00D1357:
		return NsECannotBuyOrDownloadFromMultipleServices
	case 0xC00D1358:
		return NsECannotBuyOrDownloadContent
	case 0xC00D135A:
		return NsENotContentPartnerTrack
	case 0xC00D135B:
		return NsETrackDownloadRequiresAlbumPurchase
	case 0xC00D135C:
		return NsETrackDownloadRequiresPurchase
	case 0xC00D135D:
		return NsETrackPurchaseMaximumExceeded
	case 0xC00D135F:
		return NsESubscriptionserviceLoginFailed
	case 0xC00D1360:
		return NsESubscriptionserviceDownloadTimeout
	case 0xC00D1362:
		return NsEContentPartnerStillInitializing
	case 0xC00D1363:
		return NsEOpenContainingFolderFailed
	case 0xC00D136A:
		return NsEAdvancededitTooManyPictures
	case 0xC00D1388:
		return NsERedirect
	case 0xC00D1389:
		return NsEStalePresentation
	case 0xC00D138A:
		return NsENamespaceWrongPersist
	case 0xC00D138B:
		return NsENamespaceWrongType
	case 0xC00D138C:
		return NsENamespaceNodeConflict
	case 0xC00D138D:
		return NsENamespaceNodeNotFound
	case 0xC00D138E:
		return NsENamespaceBufferTooSmall
	case 0xC00D138F:
		return NsENamespaceTooManyCallbacks
	case 0xC00D1390:
		return NsENamespaceDuplicateCallback
	case 0xC00D1391:
		return NsENamespaceCallbackNotFound
	case 0xC00D1392:
		return NsENamespaceNameTooLong
	case 0xC00D1393:
		return NsENamespaceDuplicateName
	case 0xC00D1394:
		return NsENamespaceEmptyName
	case 0xC00D1395:
		return NsENamespaceIndexTooLarge
	case 0xC00D1396:
		return NsENamespaceBadName
	case 0xC00D1397:
		return NsENamespaceWrongSecurity
	case 0xC00D13EC:
		return NsECacheArchiveConflict
	case 0xC00D13ED:
		return NsECacheOriginServerNotFound
	case 0xC00D13EE:
		return NsECacheOriginServerTimeout
	case 0xC00D13EF:
		return NsECacheNotBroadcast
	case 0xC00D13F0:
		return NsECacheCannotBeCached
	case 0xC00D13F1:
		return NsECacheNotModified
	case 0xC00D1450:
		return NsECannotRemovePublishingPoint
	case 0xC00D1451:
		return NsECannotRemovePlugin
	case 0xC00D1452:
		return NsEWrongPublishingPointType
	case 0xC00D1453:
		return NsEUnsupportedLoadType
	case 0xC00D1454:
		return NsEInvalidPluginLoadTypeConfiguration
	case 0xC00D1455:
		return NsEInvalidPublishingPointName
	case 0xC00D1456:
		return NsETooManyMulticastSinks
	case 0xC00D1457:
		return NsEPublishingPointInvalidRequestWhileStarted
	case 0xC00D1458:
		return NsEMulticastPluginNotEnabled
	case 0xC00D1459:
		return NsEInvalidOperatingSystemVersion
	case 0xC00D145A:
		return NsEPublishingPointRemoved
	case 0xC00D145B:
		return NsEInvalidPushPublishingPointStartRequest
	case 0xC00D145C:
		return NsEUnsupportedLanguage
	case 0xC00D145D:
		return NsEWrongOsVersion
	case 0xC00D145E:
		return NsEPublishingPointStopped
	case 0xC00D14B4:
		return NsEPlaylistEntryAlreadyPlaying
	case 0xC00D14B5:
		return NsEEmptyPlaylist
	case 0xC00D14B6:
		return NsEPlaylistParseFailure
	case 0xC00D14B7:
		return NsEPlaylistUnsupportedEntry
	case 0xC00D14B8:
		return NsEPlaylistEntryNotInPlaylist
	case 0xC00D14B9:
		return NsEPlaylistEntrySeek
	case 0xC00D14BA:
		return NsEPlaylistRecursivePlaylists
	case 0xC00D14BB:
		return NsEPlaylistTooManyNestedPlaylists
	case 0xC00D14BC:
		return NsEPlaylistShutdown
	case 0xC00D14BD:
		return NsEPlaylistEndReceding
	case 0xC00D1518:
		return NsEDatapathNoSink
	case 0xC00D151A:
		return NsEInvalidPushTemplate
	case 0xC00D151B:
		return NsEInvalidPushPublishingPoint
	case 0xC00D151C:
		return NsECriticalError
	case 0xC00D151D:
		return NsENoNewConnections
	case 0xC00D151E:
		return NsEWsxInvalidVersion
	case 0xC00D151F:
		return NsEHeaderMismatch
	case 0xC00D1520:
		return NsEPushDuplicatePublishingPointName
	case 0xC00D157C:
		return NsENoScriptEngine
	case 0xC00D157D:
		return NsEPluginErrorReported
	case 0xC00D157E:
		return NsESourcePluginNotFound
	case 0xC00D157F:
		return NsEPlaylistPluginNotFound
	case 0xC00D1580:
		return NsEDataSourceEnumerationNotSupported
	case 0xC00D1581:
		return NsEMediaParserInvalidFormat
	case 0xC00D1582:
		return NsEScriptDebuggerNotInstalled
	case 0xC00D1583:
		return NsEFeatureRequiresEnterpriseServer
	case 0xC00D1584:
		return NsEWizardRunning
	case 0xC00D1585:
		return NsEInvalidLogUrl
	case 0xC00D1586:
		return NsEInvalidMtuRange
	case 0xC00D1587:
		return NsEInvalidPlayStatistics
	case 0xC00D1588:
		return NsELogNeedToBeSkipped
	case 0xC00D1589:
		return NsEHttpTextDatacontainerSizeLimitExceeded
	case 0xC00D158A:
		return NsEPortInUse
	case 0xC00D158B:
		return NsEPortInUseHttp
	case 0xC00D158C:
		return NsEHttpTextDatacontainerInvalidServerResponse
	case 0xC00D158D:
		return NsEArchiveReachQuota
	case 0xC00D158E:
		return NsEArchiveAbortDueToBcast
	case 0xC00D158F:
		return NsEArchiveGapDetected
	case 0xC00D1590:
		return NsEAuthorizationFileNotFound
	case 0xC00D1B58:
		return NsEBadMarkin
	case 0xC00D1B59:
		return NsEBadMarkout
	case 0xC00D1B5A:
		return NsENomatchingMediasource
	case 0xC00D1B5B:
		return NsEUnsupportedSourcetype
	case 0xC00D1B5C:
		return NsETooManyAudio
	case 0xC00D1B5D:
		return NsETooManyVideo
	case 0xC00D1B5E:
		return NsENomatchingElement
	case 0xC00D1B5F:
		return NsEMismatchedMediacontent
	case 0xC00D1B60:
		return NsECannotDeleteActiveSourcegroup
	case 0xC00D1B61:
		return NsEAudiodeviceBusy
	case 0xC00D1B62:
		return NsEAudiodeviceUnexpected
	case 0xC00D1B63:
		return NsEAudiodeviceBadformat
	case 0xC00D1B64:
		return NsEVideodeviceBusy
	case 0xC00D1B65:
		return NsEVideodeviceUnexpected
	case 0xC00D1B66:
		return NsEInvalidcallWhileEncoderRunning
	case 0xC00D1B67:
		return NsENoProfileInSourcegroup
	case 0xC00D1B68:
		return NsEVideodriverUnstable
	case 0xC00D1B69:
		return NsEVidcapstartfailed
	case 0xC00D1B6A:
		return NsEVidsourcecompression
	case 0xC00D1B6B:
		return NsEVidsourcesize
	case 0xC00D1B6C:
		return NsEIcmqueryformat
	case 0xC00D1B6D:
		return NsEVidcapcreatewindow
	case 0xC00D1B6E:
		return NsEVidcapdrvinuse
	case 0xC00D1B6F:
		return NsENoMediaformatInSource
	case 0xC00D1B70:
		return NsENoValidOutputStream
	case 0xC00D1B71:
		return NsENoValidSourcePlugin
	case 0xC00D1B72:
		return NsENoActiveSourcegroup
	case 0xC00D1B73:
		return NsENoScriptStream
	case 0xC00D1B74:
		return NsEInvalidcallWhileArchivalRunning
	case 0xC00D1B75:
		return NsEInvalidpacketsize
	case 0xC00D1B76:
		return NsEPluginClsidInvalid
	case 0xC00D1B77:
		return NsEUnsupportedArchivetype
	case 0xC00D1B78:
		return NsEUnsupportedArchiveoperation
	case 0xC00D1B79:
		return NsEArchiveFilenameNotset
	case 0xC00D1B7A:
		return NsESourcegroupNotprepared
	case 0xC00D1B7B:
		return NsEProfileMismatch
	case 0xC00D1B7C:
		return NsEIncorrectclipsettings
	case 0xC00D1B7D:
		return NsENostatsavailable
	case 0xC00D1B7E:
		return NsENotarchiving
	case 0xC00D1B7F:
		return NsEInvalidcallWhileEncoderStopped
	case 0xC00D1B80:
		return NsENosourcegroups
	case 0xC00D1B81:
		return NsEInvalidinputfps
	case 0xC00D1B82:
		return NsENoDataviewSupport
	case 0xC00D1B83:
		return NsECodecUnavailable
	case 0xC00D1B84:
		return NsEArchiveSameAsInput
	case 0xC00D1B85:
		return NsESourceNotspecified
	case 0xC00D1B86:
		return NsENoRealtimeTimecompression
	case 0xC00D1B87:
		return NsEUnsupportedEncoderDevice
	case 0xC00D1B88:
		return NsEUnexpectedDisplaySettings
	case 0xC00D1B89:
		return NsENoAudiodata
	case 0xC00D1B8A:
		return NsEInputsourceProblem
	case 0xC00D1B8B:
		return NsEWmeVersionMismatch
	case 0xC00D1B8C:
		return NsENoRealtimePreprocess
	case 0xC00D1B8D:
		return NsENoRepeatPreprocess
	case 0xC00D1B8E:
		return NsECannotPauseLivebroadcast
	case 0xC00D1B8F:
		return NsEDrmProfileNotSet
	case 0xC00D1B90:
		return NsEDuplicateDrmprofile
	case 0xC00D1B91:
		return NsEInvalidDevice
	case 0xC00D1B92:
		return NsESpeechedlOnNonMixedmode
	case 0xC00D1B93:
		return NsEDrmPasswordTooLong
	case 0xC00D1B94:
		return NsEDevcontrolFailedSeek
	case 0xC00D1B95:
		return NsEInterlaceRequireSamesize
	case 0xC00D1B96:
		return NsETooManyDevicecontrol
	case 0xC00D1B97:
		return NsENoMultipassForLivedevice
	case 0xC00D1B98:
		return NsEMissingAudience
	case 0xC00D1B99:
		return NsEAudienceContenttypeMismatch
	case 0xC00D1B9A:
		return NsEMissingSourceIndex
	case 0xC00D1B9B:
		return NsENumLanguageMismatch
	case 0xC00D1B9C:
		return NsELanguageMismatch
	case 0xC00D1B9D:
		return NsEVbrmodeMismatch
	case 0xC00D1B9E:
		return NsEInvalidInputAudienceIndex
	case 0xC00D1B9F:
		return NsEInvalidInputLanguage
	case 0xC00D1BA0:
		return NsEInvalidInputStream
	case 0xC00D1BA1:
		return NsEExpectMonoWavInput
	case 0xC00D1BA2:
		return NsEInputWavformatMismatch
	case 0xC00D1BA3:
		return NsERecordqDiskFull
	case 0xC00D1BA4:
		return NsENoPalInverseTelecine
	case 0xC00D1BA5:
		return NsEActiveSgDeviceDisconnected
	case 0xC00D1BA6:
		return NsEActiveSgDeviceControlDisconnected
	case 0xC00D1BA7:
		return NsENoFramesSubmittedToAnalyzer
	case 0xC00D1BA8:
		return NsEInputDoesnotSupportSmpte
	case 0xC00D1BA9:
		return NsENoSmpteWithMultipleSourcegroups
	case 0xC00D1BAA:
		return NsEBadContentedl
	case 0xC00D1BAB:
		return NsEInterlacemodeMismatch
	case 0xC00D1BAC:
		return NsENonsquarepixelmodeMismatch
	case 0xC00D1BAD:
		return NsESmptemodeMismatch
	case 0xC00D1BAE:
		return NsEEndOfTape
	case 0xC00D1BAF:
		return NsENoMediaInAudience
	case 0xC00D1BB0:
		return NsENoAudiences
	case 0xC00D1BB1:
		return NsENoAudioCompat
	case 0xC00D1BB2:
		return NsEInvalidVbrCompat
	case 0xC00D1BB3:
		return NsENoProfileName
	case 0xC00D1BB4:
		return NsEInvalidVbrWithUncomp
	case 0xC00D1BB5:
		return NsEMultipleVbrAudiences
	case 0xC00D1BB6:
		return NsEUncompCompCombination
	case 0xC00D1BB7:
		return NsEMultipleAudioCodecs
	case 0xC00D1BB8:
		return NsEMultipleAudioFormats
	case 0xC00D1BB9:
		return NsEAudioBitrateStepdown
	case 0xC00D1BBA:
		return NsEInvalidAudioPeakrate
	case 0xC00D1BBB:
		return NsEInvalidAudioPeakrate2
	case 0xC00D1BBC:
		return NsEInvalidAudioBuffermax
	case 0xC00D1BBD:
		return NsEMultipleVideoCodecs
	case 0xC00D1BBE:
		return NsEMultipleVideoSizes
	case 0xC00D1BBF:
		return NsEInvalidVideoBitrate
	case 0xC00D1BC0:
		return NsEVideoBitrateStepdown
	case 0xC00D1BC1:
		return NsEInvalidVideoPeakrate
	case 0xC00D1BC2:
		return NsEInvalidVideoPeakrate2
	case 0xC00D1BC3:
		return NsEInvalidVideoWidth
	case 0xC00D1BC4:
		return NsEInvalidVideoHeight
	case 0xC00D1BC5:
		return NsEInvalidVideoFps
	case 0xC00D1BC6:
		return NsEInvalidVideoKeyframe
	case 0xC00D1BC7:
		return NsEInvalidVideoIquality
	case 0xC00D1BC8:
		return NsEInvalidVideoCquality
	case 0xC00D1BC9:
		return NsEInvalidVideoBuffer
	case 0xC00D1BCA:
		return NsEInvalidVideoBuffermax
	case 0xC00D1BCB:
		return NsEInvalidVideoBuffermax2
	case 0xC00D1BCC:
		return NsEInvalidVideoWidthAlign
	case 0xC00D1BCD:
		return NsEInvalidVideoHeightAlign
	case 0xC00D1BCE:
		return NsEMultipleScriptBitrates
	case 0xC00D1BCF:
		return NsEInvalidScriptBitrate
	case 0xC00D1BD0:
		return NsEMultipleFileBitrates
	case 0xC00D1BD1:
		return NsEInvalidFileBitrate
	case 0xC00D1BD2:
		return NsESameAsInputCombination
	case 0xC00D1BD3:
		return NsESourceCannotLoop
	case 0xC00D1BD4:
		return NsEInvalidFolddownCoefficients
	case 0xC00D1BD5:
		return NsEDrmprofileNotfound
	case 0xC00D1BD6:
		return NsEInvalidTimecode
	case 0xC00D1BD7:
		return NsENoAudioTimecompression
	case 0xC00D1BD8:
		return NsENoTwopassTimecompression
	case 0xC00D1BD9:
		return NsETimecodeRequiresVideostream
	case 0xC00D1BDA:
		return NsENoMbrWithTimecode
	case 0xC00D1BDB:
		return NsEInvalidInterlacemode
	case 0xC00D1BDC:
		return NsEInvalidInterlaceCompat
	case 0xC00D1BDD:
		return NsEInvalidNonsquarepixelCompat
	case 0xC00D1BDE:
		return NsEInvalidSourceWithDeviceControl
	case 0xC00D1BDF:
		return NsECannotGenerateBroadcastInfoForQualityvbr
	case 0xC00D1BE0:
		return NsEExceedMaxDrmProfileLimit
	case 0xC00D1BE1:
		return NsEDevicecontrolUnstable
	case 0xC00D1BE2:
		return NsEInvalidPixelAspectRatio
	case 0xC00D1BE3:
		return NsEAudienceLanguageContenttypeMismatch
	case 0xC00D1BE4:
		return NsEInvalidProfileContenttype
	case 0xC00D1BE5:
		return NsETransformPluginNotFound
	case 0xC00D1BE6:
		return NsETransformPluginInvalid
	case 0xC00D1BE7:
		return NsEEdlRequiredForDeviceMultipass
	case 0xC00D1BE8:
		return NsEInvalidVideoWidthForInterlacedEncoding
	case 0xC00D1BE9:
		return NsEMarkinUnsupported
	case 0xC00D2711:
		return NsEDrmInvalidApplication
	case 0xC00D2712:
		return NsEDrmLicenseStoreError
	case 0xC00D2713:
		return NsEDrmSecureStoreError
	case 0xC00D2714:
		return NsEDrmLicenseStoreSaveError
	case 0xC00D2715:
		return NsEDrmSecureStoreUnlockError
	case 0xC00D2716:
		return NsEDrmInvalidContent
	case 0xC00D2717:
		return NsEDrmUnableToOpenLicense
	case 0xC00D2718:
		return NsEDrmInvalidLicense
	case 0xC00D2719:
		return NsEDrmInvalidMachine
	case 0xC00D271B:
		return NsEDrmEnumLicenseFailed
	case 0xC00D271C:
		return NsEDrmInvalidLicenseRequest
	case 0xC00D271D:
		return NsEDrmUnableToInitialize
	case 0xC00D271E:
		return NsEDrmUnableToAcquireLicense
	case 0xC00D271F:
		return NsEDrmInvalidLicenseAcquired
	case 0xC00D2720:
		return NsEDrmNoRights
	case 0xC00D2721:
		return NsEDrmKeyError
	case 0xC00D2722:
		return NsEDrmEncryptError
	case 0xC00D2723:
		return NsEDrmDecryptError
	case 0xC00D2725:
		return NsEDrmLicenseInvalidXml
	case 0xC00D2728:
		return NsEDrmNeedsIndividualization
	case 0xC00D2729:
		return NsEDrmAlreadyIndividualized
	case 0xC00D272A:
		return NsEDrmActionNotQueried
	case 0xC00D272B:
		return NsEDrmAcquiringLicense
	case 0xC00D272C:
		return NsEDrmIndividualizing
	case 0xC00D272D:
		return NsEBackupRestoreFailure
	case 0xC00D272E:
		return NsEBackupRestoreBadRequestId
	case 0xC00D272F:
		return NsEDrmParametersMismatched
	case 0xC00D2730:
		return NsEDrmUnableToCreateLicenseObject
	case 0xC00D2731:
		return NsEDrmUnableToCreateIndiObject
	case 0xC00D2732:
		return NsEDrmUnableToCreateEncryptObject
	case 0xC00D2733:
		return NsEDrmUnableToCreateDecryptObject
	case 0xC00D2734:
		return NsEDrmUnableToCreatePropertiesObject
	case 0xC00D2735:
		return NsEDrmUnableToCreateBackupObject
	case 0xC00D2736:
		return NsEDrmIndividualizeError
	case 0xC00D2737:
		return NsEDrmLicenseOpenError
	case 0xC00D2738:
		return NsEDrmLicenseCloseError
	case 0xC00D2739:
		return NsEDrmGetLicenseError
	case 0xC00D273A:
		return NsEDrmQueryError
	case 0xC00D273B:
		return NsEDrmReportError
	case 0xC00D273C:
		return NsEDrmGetLicensestringError
	case 0xC00D273D:
		return NsEDrmGetContentstringError
	case 0xC00D273E:
		return NsEDrmMonitorError
	case 0xC00D273F:
		return NsEDrmUnableToSetParameter
	case 0xC00D2740:
		return NsEDrmInvalidAppdata
	case 0xC00D2741:
		return NsEDrmInvalidAppdataVersion
	case 0xC00D2742:
		return NsEDrmBackupExists
	case 0xC00D2743:
		return NsEDrmBackupCorrupt
	case 0xC00D2744:
		return NsEDrmBackuprestoreBusy
	case 0xC00D2745:
		return NsEBackupRestoreBadData
	case 0xC00D2748:
		return NsEDrmLicenseUnusable
	case 0xC00D2749:
		return NsEDrmInvalidProperty
	case 0xC00D274A:
		return NsEDrmSecureStoreNotFound
	case 0xC00D274B:
		return NsEDrmCachedContentError
	case 0xC00D274C:
		return NsEDrmIndividualizationIncomplete
	case 0xC00D274D:
		return NsEDrmDriverAuthFailure
	case 0xC00D274E:
		return NsEDrmNeedUpgradeMssap
	case 0xC00D274F:
		return NsEDrmReopenContent
	case 0xC00D2750:
		return NsEDrmDriverDigioutFailure
	case 0xC00D2751:
		return NsEDrmInvalidSecurestorePassword
	case 0xC00D2752:
		return NsEDrmAppcertRevoked
	case 0xC00D2753:
		return NsEDrmRestoreFraud
	case 0xC00D2754:
		return NsEDrmHardwareInconsistent
	case 0xC00D2755:
		return NsEDrmSdmiTrigger
	case 0xC00D2756:
		return NsEDrmSdmiNomorecopies
	case 0xC00D2757:
		return NsEDrmUnableToCreateHeaderObject
	case 0xC00D2758:
		return NsEDrmUnableToCreateKeysObject
	case 0xC00D2759:
		return NsEDrmLicenseNotacquired
	case 0xC00D275A:
		return NsEDrmUnableToCreateCodingObject
	case 0xC00D275B:
		return NsEDrmUnableToCreateStateDataObject
	case 0xC00D275C:
		return NsEDrmBufferTooSmall
	case 0xC00D275D:
		return NsEDrmUnsupportedProperty
	case 0xC00D275E:
		return NsEDrmErrorBadNetResp
	case 0xC00D275F:
		return NsEDrmStoreNotallstored
	case 0xC00D2760:
		return NsEDrmSecurityComponentSignatureInvalid
	case 0xC00D2761:
		return NsEDrmInvalidData
	case 0xC00D2762:
		return NsEDrmPolicyDisableOnline
	case 0xC00D2763:
		return NsEDrmUnableToCreateAuthenticationObject
	case 0xC00D2764:
		return NsEDrmNotConfigured
	case 0xC00D2765:
		return NsEDrmDeviceActivationCanceled
	case 0xC00D2766:
		return NsEBackupRestoreTooManyResets
	case 0xC00D2767:
		return NsEDrmDebuggingNotAllowed
	case 0xC00D2768:
		return NsEDrmOperationCanceled
	case 0xC00D2769:
		return NsEDrmRestrictionsNotRetrieved
	case 0xC00D276A:
		return NsEDrmUnableToCreatePlaylistObject
	case 0xC00D276B:
		return NsEDrmUnableToCreatePlaylistBurnObject
	case 0xC00D276C:
		return NsEDrmUnableToCreateDeviceRegistrationObject
	case 0xC00D276D:
		return NsEDrmUnableToCreateMeteringObject
	case 0xC00D2770:
		return NsEDrmTrackExceededPlaylistRestiction
	case 0xC00D2771:
		return NsEDrmTrackExceededTrackburnRestriction
	case 0xC00D2772:
		return NsEDrmUnableToGetDeviceCert
	case 0xC00D2773:
		return NsEDrmUnableToGetSecureClock
	case 0xC00D2774:
		return NsEDrmUnableToSetSecureClock
	case 0xC00D2775:
		return NsEDrmUnableToGetSecureClockFromServer
	case 0xC00D2776:
		return NsEDrmPolicyMeteringDisabled
	case 0xC00D2777:
		return NsEDrmTransferChainedLicensesUnsupported
	case 0xC00D2778:
		return NsEDrmSdkVersionmismatch
	case 0xC00D2779:
		return NsEDrmLicNeedsDeviceClockSet
	case 0xC00D277A:
		return NsELicenseHeaderMissingUrl
	case 0xC00D277B:
		return NsEDeviceNotWmdrmDevice
	case 0xC00D277C:
		return NsEDrmInvalidAppcert
	case 0xC00D277D:
		return NsEDrmProtocolForcefulTerminationOnPetition
	case 0xC00D277E:
		return NsEDrmProtocolForcefulTerminationOnChallenge
	case 0xC00D277F:
		return NsEDrmCheckpointFailed
	case 0xC00D2780:
		return NsEDrmBbUnableToInitialize
	case 0xC00D2781:
		return NsEDrmUnableToLoadHardwareId
	case 0xC00D2782:
		return NsEDrmUnableToOpenDataStore
	case 0xC00D2783:
		return NsEDrmDatastoreCorrupt
	case 0xC00D2784:
		return NsEDrmUnableToCreateInmemorystoreObject
	case 0xC00D2785:
		return NsEDrmStublibRequired
	case 0xC00D2786:
		return NsEDrmUnableToCreateCertificateObject
	case 0xC00D2787:
		return NsEDrmMigrationTargetNotOnline
	case 0xC00D2788:
		return NsEDrmInvalidMigrationImage
	case 0xC00D2789:
		return NsEDrmMigrationTargetStatesCorrupted
	case 0xC00D278A:
		return NsEDrmMigrationImporterNotAvailable
	case 0xC00D278B:
		return NsDrmEMigrationUpgradeWithDiffSid
	case 0xC00D278C:
		return NsDrmEMigrationSourceMachineInUse
	case 0xC00D278D:
		return NsDrmEMigrationTargetMachineLessThanLh
	case 0xC00D278E:
		return NsDrmEMigrationImageAlreadyExists
	case 0xC00D278F:
		return NsEDrmHardwareidMismatch
	case 0xC00D2790:
		return NsEInvalidDrmv2cltStublib
	case 0xC00D2791:
		return NsEDrmMigrationInvalidLegacyv2Data
	case 0xC00D2792:
		return NsEDrmMigrationLicenseAlreadyExists
	case 0xC00D2793:
		return NsEDrmMigrationInvalidLegacyv2SstPassword
	case 0xC00D2794:
		return NsEDrmMigrationNotSupported
	case 0xC00D2795:
		return NsEDrmUnableToCreateMigrationImporterObject
	case 0xC00D2796:
		return NsEDrmCheckpointMismatch
	case 0xC00D2797:
		return NsEDrmCheckpointCorrupt
	case 0xC00D2798:
		return NsERegFlushFailure
	case 0xC00D2799:
		return NsEHdsKeyMismatch
	case 0xC00D279A:
		return NsEDrmMigrationOperationCancelled
	case 0xC00D279B:
		return NsEDrmMigrationObjectInUse
	case 0xC00D279C:
		return NsEDrmMalformedContentHeader
	case 0xC00D27D8:
		return NsEDrmLicenseExpired
	case 0xC00D27D9:
		return NsEDrmLicenseNotenabled
	case 0xC00D27DA:
		return NsEDrmLicenseAppseclow
	case 0xC00D27DB:
		return NsEDrmStoreNeedindi
	case 0xC00D27DC:
		return NsEDrmStoreNotallowed
	case 0xC00D27DD:
		return NsEDrmLicenseAppNotallowed
	case 0xC00D27DF:
		return NsEDrmLicenseCertExpired
	case 0xC00D27E0:
		return NsEDrmLicenseSeclow
	case 0xC00D27E1:
		return NsEDrmLicenseContentRevoked
	case 0xC00D27E2:
		return NsEDrmDeviceNotRegistered
	case 0xC00D280A:
		return NsEDrmLicenseNosap
	case 0xC00D280B:
		return NsEDrmLicenseNosvp
	case 0xC00D280C:
		return NsEDrmLicenseNowdm
	case 0xC00D280D:
		return NsEDrmLicenseNotrustedcodec
	case 0xC00D280E:
		return NsEDrmSourceidNotSupported
	case 0xC00D283D:
		return NsEDrmNeedsUpgradeTempfile
	case 0xC00D283E:
		return NsEDrmNeedUpgradePd
	case 0xC00D283F:
		return NsEDrmSignatureFailure
	case 0xC00D2840:
		return NsEDrmLicenseServerInfoMissing
	case 0xC00D2841:
		return NsEDrmBusy
	case 0xC00D2842:
		return NsEDrmPdTooManyDevices
	case 0xC00D2843:
		return NsEDrmIndivFraud
	case 0xC00D2844:
		return NsEDrmIndivNoCabs
	case 0xC00D2845:
		return NsEDrmIndivServiceUnavailable
	case 0xC00D2846:
		return NsEDrmRestoreServiceUnavailable
	case 0xC00D2847:
		return NsEDrmClientCodeExpired
	case 0xC00D2848:
		return NsEDrmNoUplinkLicense
	case 0xC00D2849:
		return NsEDrmInvalidKid
	case 0xC00D284A:
		return NsEDrmLicenseInitializationError
	case 0xC00D284C:
		return NsEDrmChainTooLong
	case 0xC00D284D:
		return NsEDrmUnsupportedAlgorithm
	case 0xC00D284E:
		return NsEDrmLicenseDeletionError
	case 0xC00D28A0:
		return NsEDrmInvalidCertificate
	case 0xC00D28A1:
		return NsEDrmCertificateRevoked
	case 0xC00D28A2:
		return NsEDrmLicenseUnavailable
	case 0xC00D28A3:
		return NsEDrmDeviceLimitReached
	case 0xC00D28A4:
		return NsEDrmUnableToVerifyProximity
	case 0xC00D28A5:
		return NsEDrmMustRegister
	case 0xC00D28A6:
		return NsEDrmMustApprove
	case 0xC00D28A7:
		return NsEDrmMustRevalidate
	case 0xC00D28A8:
		return NsEDrmInvalidProximityResponse
	case 0xC00D28A9:
		return NsEDrmInvalidSession
	case 0xC00D28AA:
		return NsEDrmDeviceNotOpen
	case 0xC00D28AB:
		return NsEDrmDeviceAlreadyRegistered
	case 0xC00D28AC:
		return NsEDrmUnsupportedProtocolVersion
	case 0xC00D28AD:
		return NsEDrmUnsupportedAction
	case 0xC00D28AE:
		return NsEDrmCertificateSecurityLevelInadequate
	case 0xC00D28AF:
		return NsEDrmUnableToOpenPort
	case 0xC00D28B0:
		return NsEDrmBadRequest
	case 0xC00D28B1:
		return NsEDrmInvalidCrl
	case 0xC00D28B2:
		return NsEDrmAttributeTooLong
	case 0xC00D28B3:
		return NsEDrmExpiredLicenseblob
	case 0xC00D28B4:
		return NsEDrmInvalidLicenseblob
	case 0xC00D28B5:
		return NsEDrmInclusionListRequired
	case 0xC00D28B6:
		return NsEDrmDrmv2cltRevoked
	case 0xC00D28B7:
		return NsEDrmRivTooSmall
	case 0xC00D2904:
		return NsEOutputProtectionLevelUnsupported
	case 0xC00D2905:
		return NsECompressedDigitalVideoProtectionLevelUnsupported
	case 0xC00D2906:
		return NsEUncompressedDigitalVideoProtectionLevelUnsupported
	case 0xC00D2907:
		return NsEAnalogVideoProtectionLevelUnsupported
	case 0xC00D2908:
		return NsECompressedDigitalAudioProtectionLevelUnsupported
	case 0xC00D2909:
		return NsEUncompressedDigitalAudioProtectionLevelUnsupported
	case 0xC00D290A:
		return NsEOutputProtectionSchemeUnsupported
	case 0xC00D2AFA:
		return NsERebootRecommended
	case 0xC00D2AFB:
		return NsERebootRequired
	case 0xC00D2AFC:
		return NsESetupIncomplete
	case 0xC00D2AFD:
		return NsESetupDrmMigrationFailed
	case 0xC00D2AFE:
		return NsESetupIgnorableFailure
	case 0xC00D2AFF:
		return NsESetupDrmMigrationFailedAndIgnorableFailure
	case 0xC00D2B00:
		return NsESetupBlocked
	case 0xC00D2EE0:
		return NsEUnknownProtocol
	case 0xC00D2EE1:
		return NsERedirectToProxy
	case 0xC00D2EE2:
		return NsEInternalServerError
	case 0xC00D2EE3:
		return NsEBadRequest
	case 0xC00D2EE4:
		return NsEErrorFromProxy
	case 0xC00D2EE5:
		return NsEProxyTimeout
	case 0xC00D2EE6:
		return NsEServerUnavailable
	case 0xC00D2EE7:
		return NsERefusedByServer
	case 0xC00D2EE8:
		return NsEIncompatibleServer
	case 0xC00D2EE9:
		return NsEMulticastDisabled
	case 0xC00D2EEA:
		return NsEInvalidRedirect
	case 0xC00D2EEB:
		return NsEAllProtocolsDisabled
	case 0xC00D2EEC:
		return NsEMsbdNoLongerSupported
	case 0xC00D2EED:
		return NsEProxyNotFound
	case 0xC00D2EEE:
		return NsECannotConnectToProxy
	case 0xC00D2EEF:
		return NsEServerDnsTimeout
	case 0xC00D2EF0:
		return NsEProxyDnsTimeout
	case 0xC00D2EF1:
		return NsEClosedOnSuspend
	case 0xC00D2EF2:
		return NsECannotReadPlaylistFromMediaserver
	case 0xC00D2EF3:
		return NsESessionNotFound
	case 0xC00D2EF4:
		return NsERequireStreamingClient
	case 0xC00D2EF5:
		return NsEPlaylistEntryHasChanged
	case 0xC00D2EF6:
		return NsEProxyAccessdenied
	case 0xC00D2EF7:
		return NsEProxySourceAccessdenied
	case 0xC00D2EF8:
		return NsENetworkSinkWrite
	case 0xC00D2EF9:
		return NsEFirewall
	case 0xC00D2EFA:
		return NsEMmsNotSupported
	case 0xC00D2EFB:
		return NsEServerAccessdenied
	case 0xC00D2EFC:
		return NsEResourceGone
	case 0xC00D2EFD:
		return NsENoExistingPacketizer
	case 0xC00D2EFE:
		return NsEBadSyntaxInServerResponse
	case 0xC00D2F00:
		return NsEResetSocketConnection
	case 0xC00D2F02:
		return NsETooManyHops
	case 0xC00D2F05:
		return NsETooMuchDataFromServer
	case 0xC00D2F06:
		return NsEConnectTimeout
	case 0xC00D2F07:
		return NsEProxyConnectTimeout
	case 0xC00D2F08:
		return NsESessionInvalid
	case 0xC00D2F0A:
		return NsEPacketsinkUnknownFecStream
	case 0xC00D2F0B:
		return NsEPushCannotconnect
	case 0xC00D2F0C:
		return NsEIncompatiblePushServer
	case 0xC00D32C8:
		return NsEEndOfPlaylist
	case 0xC00D32C9:
		return NsEUseFileSource
	case 0xC00D32CA:
		return NsEPropertyNotFound
	case 0xC00D32CC:
		return NsEPropertyReadOnly
	case 0xC00D32CD:
		return NsETableKeyNotFound
	case 0xC00D32CF:
		return NsEInvalidQueryOperator
	case 0xC00D32D0:
		return NsEInvalidQueryProperty
	case 0xC00D32D2:
		return NsEPropertyNotSupported
	case 0xC00D32D4:
		return NsESchemaClassifyFailure
	case 0xC00D32D5:
		return NsEMetadataFormatNotSupported
	case 0xC00D32D6:
		return NsEMetadataNoEditingCapability
	case 0xC00D32D7:
		return NsEMetadataCannotSetLocale
	case 0xC00D32D8:
		return NsEMetadataLanguageNotSuported
	case 0xC00D32D9:
		return NsEMetadataNoRfc1766NameForLocale
	case 0xC00D32DA:
		return NsEMetadataNotAvailable
	case 0xC00D32DB:
		return NsEMetadataCacheDataNotAvailable
	case 0xC00D32DC:
		return NsEMetadataInvalidDocumentType
	case 0xC00D32DD:
		return NsEMetadataIdentifierNotAvailable
	case 0xC00D32DE:
		return NsEMetadataCannotRetrieveFromOfflineCache
	case 0xC0261003:
		return ErrorMonitorInvalidDescriptorChecksum
	case 0xC0261004:
		return ErrorMonitorInvalidStandardTimingBlock
	case 0xC0261005:
		return ErrorMonitorWmiDatablockRegistrationFailed
	case 0xC0261006:
		return ErrorMonitorInvalidSerialNumberMondscBlock
	case 0xC0261007:
		return ErrorMonitorInvalidUserFriendlyMondscBlock
	case 0xC0261008:
		return ErrorMonitorNoMoreDescriptorData
	case 0xC0261009:
		return ErrorMonitorInvalidDetailedTimingBlock
	case 0xC0262000:
		return ErrorGraphicsNotExclusiveModeOwner
	case 0xC0262001:
		return ErrorGraphicsInsufficientDmaBuffer
	case 0xC0262002:
		return ErrorGraphicsInvalidDisplayAdapter
	case 0xC0262003:
		return ErrorGraphicsAdapterWasReset
	case 0xC0262004:
		return ErrorGraphicsInvalidDriverModel
	case 0xC0262005:
		return ErrorGraphicsPresentModeChanged
	case 0xC0262006:
		return ErrorGraphicsPresentOccluded
	case 0xC0262007:
		return ErrorGraphicsPresentDenied
	case 0xC0262008:
		return ErrorGraphicsCannotcolorconvert
	case 0xC0262100:
		return ErrorGraphicsNoVideoMemory
	case 0xC0262101:
		return ErrorGraphicsCantLockMemory
	case 0xC0262102:
		return ErrorGraphicsAllocationBusy
	case 0xC0262103:
		return ErrorGraphicsTooManyReferences
	case 0xC0262104:
		return ErrorGraphicsTryAgainLater
	case 0xC0262105:
		return ErrorGraphicsTryAgainNow
	case 0xC0262106:
		return ErrorGraphicsAllocationInvalid
	case 0xC0262107:
		return ErrorGraphicsUnswizzlingApertureUnavailable
	case 0xC0262108:
		return ErrorGraphicsUnswizzlingApertureUnsupported
	case 0xC0262109:
		return ErrorGraphicsCantEvictPinnedAllocation
	case 0xC0262110:
		return ErrorGraphicsInvalidAllocationUsage
	case 0xC0262111:
		return ErrorGraphicsCantRenderLockedAllocation
	case 0xC0262112:
		return ErrorGraphicsAllocationClosed
	case 0xC0262113:
		return ErrorGraphicsInvalidAllocationInstance
	case 0xC0262114:
		return ErrorGraphicsInvalidAllocationHandle
	case 0xC0262115:
		return ErrorGraphicsWrongAllocationDevice
	case 0xC0262116:
		return ErrorGraphicsAllocationContentLost
	case 0xC0262200:
		return ErrorGraphicsGpuExceptionOnDevice
	case 0xC0262300:
		return ErrorGraphicsInvalidVidpnTopology
	case 0xC0262301:
		return ErrorGraphicsVidpnTopologyNotSupported
	case 0xC0262302:
		return ErrorGraphicsVidpnTopologyCurrentlyNotSupported
	case 0xC0262303:
		return ErrorGraphicsInvalidVidpn
	case 0xC0262304:
		return ErrorGraphicsInvalidVideoPresentSource
	case 0xC0262305:
		return ErrorGraphicsInvalidVideoPresentTarget
	case 0xC0262306:
		return ErrorGraphicsVidpnModalityNotSupported
	case 0xC0262308:
		return ErrorGraphicsInvalidVidpnSourcemodeset
	case 0xC0262309:
		return ErrorGraphicsInvalidVidpnTargetmodeset
	case 0xC026230A:
		return ErrorGraphicsInvalidFrequency
	case 0xC026230B:
		return ErrorGraphicsInvalidActiveRegion
	case 0xC026230C:
		return ErrorGraphicsInvalidTotalRegion
	case 0xC0262310:
		return ErrorGraphicsInvalidVideoPresentSourceMode
	case 0xC0262311:
		return ErrorGraphicsInvalidVideoPresentTargetMode
	case 0xC0262312:
		return ErrorGraphicsPinnedModeMustRemainInSet
	case 0xC0262313:
		return ErrorGraphicsPathAlreadyInTopology
	case 0xC0262314:
		return ErrorGraphicsModeAlreadyInModeset
	case 0xC0262315:
		return ErrorGraphicsInvalidVideopresentsourceset
	case 0xC0262316:
		return ErrorGraphicsInvalidVideopresenttargetset
	case 0xC0262317:
		return ErrorGraphicsSourceAlreadyInSet
	case 0xC0262318:
		return ErrorGraphicsTargetAlreadyInSet
	case 0xC0262319:
		return ErrorGraphicsInvalidVidpnPresentPath
	case 0xC026231A:
		return ErrorGraphicsNoRecommendedVidpnTopology
	case 0xC026231B:
		return ErrorGraphicsInvalidMonitorFrequencyrangeset
	case 0xC026231C:
		return ErrorGraphicsInvalidMonitorFrequencyrange
	case 0xC026231D:
		return ErrorGraphicsFrequencyrangeNotInSet
	case 0xC026231F:
		return ErrorGraphicsFrequencyrangeAlreadyInSet
	case 0xC0262320:
		return ErrorGraphicsStaleModeset
	case 0xC0262321:
		return ErrorGraphicsInvalidMonitorSourcemodeset
	case 0xC0262322:
		return ErrorGraphicsInvalidMonitorSourceMode
	case 0xC0262323:
		return ErrorGraphicsNoRecommendedFunctionalVidpn
	case 0xC0262324:
		return ErrorGraphicsModeIdMustBeUnique
	case 0xC0262325:
		return ErrorGraphicsEmptyAdapterMonitorModeSupportIntersection
	case 0xC0262326:
		return ErrorGraphicsVideoPresentTargetsLessThanSources
	case 0xC0262327:
		return ErrorGraphicsPathNotInTopology
	case 0xC0262328:
		return ErrorGraphicsAdapterMustHaveAtLeastOneSource
	case 0xC0262329:
		return ErrorGraphicsAdapterMustHaveAtLeastOneTarget
	case 0xC026232A:
		return ErrorGraphicsInvalidMonitordescriptorset
	case 0xC026232B:
		return ErrorGraphicsInvalidMonitordescriptor
	case 0xC026232C:
		return ErrorGraphicsMonitordescriptorNotInSet
	case 0xC026232D:
		return ErrorGraphicsMonitordescriptorAlreadyInSet
	case 0xC026232E:
		return ErrorGraphicsMonitordescriptorIdMustBeUnique
	case 0xC026232F:
		return ErrorGraphicsInvalidVidpnTargetSubsetType
	case 0xC0262330:
		return ErrorGraphicsResourcesNotRelated
	case 0xC0262331:
		return ErrorGraphicsSourceIdMustBeUnique
	case 0xC0262332:
		return ErrorGraphicsTargetIdMustBeUnique
	case 0xC0262333:
		return ErrorGraphicsNoAvailableVidpnTarget
	case 0xC0262334:
		return ErrorGraphicsMonitorCouldNotBeAssociatedWithAdapter
	case 0xC0262335:
		return ErrorGraphicsNoVidpnmgr
	case 0xC0262336:
		return ErrorGraphicsNoActiveVidpn
	case 0xC0262337:
		return ErrorGraphicsStaleVidpnTopology
	case 0xC0262338:
		return ErrorGraphicsMonitorNotConnected
	case 0xC0262339:
		return ErrorGraphicsSourceNotInTopology
	case 0xC026233A:
		return ErrorGraphicsInvalidPrimarysurfaceSize
	case 0xC026233B:
		return ErrorGraphicsInvalidVisibleregionSize
	case 0xC026233C:
		return ErrorGraphicsInvalidStride
	case 0xC026233D:
		return ErrorGraphicsInvalidPixelformat
	case 0xC026233E:
		return ErrorGraphicsInvalidColorbasis
	case 0xC026233F:
		return ErrorGraphicsInvalidPixelvalueaccessmode
	case 0xC0262340:
		return ErrorGraphicsTargetNotInTopology
	case 0xC0262341:
		return ErrorGraphicsNoDisplayModeManagementSupport
	case 0xC0262342:
		return ErrorGraphicsVidpnSourceInUse
	case 0xC0262343:
		return ErrorGraphicsCantAccessActiveVidpn
	case 0xC0262344:
		return ErrorGraphicsInvalidPathImportanceOrdinal
	case 0xC0262345:
		return ErrorGraphicsInvalidPathContentGeometryTransformation
	case 0xC0262346:
		return ErrorGraphicsPathContentGeometryTransformationNotSupported
	case 0xC0262347:
		return ErrorGraphicsInvalidGammaRamp
	case 0xC0262348:
		return ErrorGraphicsGammaRampNotSupported
	case 0xC0262349:
		return ErrorGraphicsMultisamplingNotSupported
	case 0xC026234A:
		return ErrorGraphicsModeNotInModeset
	case 0xC026234D:
		return ErrorGraphicsInvalidVidpnTopologyRecommendationReason
	case 0xC026234E:
		return ErrorGraphicsInvalidPathContentType
	case 0xC026234F:
		return ErrorGraphicsInvalidCopyprotectionType
	case 0xC0262350:
		return ErrorGraphicsUnassignedModesetAlreadyExists
	case 0xC0262352:
		return ErrorGraphicsInvalidScanlineOrdering
	case 0xC0262353:
		return ErrorGraphicsTopologyChangesNotAllowed
	case 0xC0262354:
		return ErrorGraphicsNoAvailableImportanceOrdinals
	case 0xC0262355:
		return ErrorGraphicsIncompatiblePrivateFormat
	case 0xC0262356:
		return ErrorGraphicsInvalidModePruningAlgorithm
	case 0xC0262400:
		return ErrorGraphicsSpecifiedChildAlreadyConnected
	case 0xC0262401:
		return ErrorGraphicsChildDescriptorNotSupported
	case 0xC0262430:
		return ErrorGraphicsNotALinkedAdapter
	case 0xC0262431:
		return ErrorGraphicsLeadlinkNotEnumerated
	case 0xC0262432:
		return ErrorGraphicsChainlinksNotEnumerated
	case 0xC0262433:
		return ErrorGraphicsAdapterChainNotReady
	case 0xC0262434:
		return ErrorGraphicsChainlinksNotStarted
	case 0xC0262435:
		return ErrorGraphicsChainlinksNotPoweredOn
	case 0xC0262436:
		return ErrorGraphicsInconsistentDeviceLinkState
	case 0xC0262438:
		return ErrorGraphicsNotPostDeviceDriver
	case 0xC0262500:
		return ErrorGraphicsOpmNotSupported
	case 0xC0262501:
		return ErrorGraphicsCoppNotSupported
	case 0xC0262502:
		return ErrorGraphicsUabNotSupported
	case 0xC0262503:
		return ErrorGraphicsOpmInvalidEncryptedParameters
	case 0xC0262504:
		return ErrorGraphicsOpmParameterArrayTooSmall
	case 0xC0262505:
		return ErrorGraphicsOpmNoVideoOutputsExist
	case 0xC0262506:
		return ErrorGraphicsPvpNoDisplayDeviceCorrespondsToName
	case 0xC0262507:
		return ErrorGraphicsPvpDisplayDeviceNotAttachedToDesktop
	case 0xC0262508:
		return ErrorGraphicsPvpMirroringDevicesNotSupported
	case 0xC026250A:
		return ErrorGraphicsOpmInvalidPointer
	case 0xC026250B:
		return ErrorGraphicsOpmInternalError
	case 0xC026250C:
		return ErrorGraphicsOpmInvalidHandle
	case 0xC026250D:
		return ErrorGraphicsPvpNoMonitorsCorrespondToDisplayDevice
	case 0xC026250E:
		return ErrorGraphicsPvpInvalidCertificateLength
	case 0xC026250F:
		return ErrorGraphicsOpmSpanningModeEnabled
	case 0xC0262510:
		return ErrorGraphicsOpmTheaterModeEnabled
	case 0xC0262511:
		return ErrorGraphicsPvpHfsFailed
	case 0xC0262512:
		return ErrorGraphicsOpmInvalidSrm
	case 0xC0262513:
		return ErrorGraphicsOpmOutputDoesNotSupportHdcp
	case 0xC0262514:
		return ErrorGraphicsOpmOutputDoesNotSupportAcp
	case 0xC0262515:
		return ErrorGraphicsOpmOutputDoesNotSupportCgmsa
	case 0xC0262516:
		return ErrorGraphicsOpmHdcpSrmNeverSet
	case 0xC0262517:
		return ErrorGraphicsOpmResolutionTooHigh
	case 0xC0262518:
		return ErrorGraphicsOpmAllHdcpHardwareAlreadyInUse
	case 0xC0262519:
		return ErrorGraphicsOpmVideoOutputNoLongerExists
	case 0xC026251A:
		return ErrorGraphicsOpmSessionTypeChangeInProgress
	case 0xC0262580:
		return ErrorGraphicsI2cNotSupported
	case 0xC0262581:
		return ErrorGraphicsI2cDeviceDoesNotExist
	case 0xC0262582:
		return ErrorGraphicsI2cErrorTransmittingData
	case 0xC0262583:
		return ErrorGraphicsI2cErrorReceivingData
	case 0xC0262584:
		return ErrorGraphicsDdcciVcpNotSupported
	case 0xC0262585:
		return ErrorGraphicsDdcciInvalidData
	case 0xC0262586:
		return ErrorGraphicsDdcciMonitorReturnedInvalidTimingStatusByte
	case 0xC0262587:
		return ErrorGraphicsMcaInvalidCapabilitiesString
	case 0xC0262588:
		return ErrorGraphicsMcaInternalError
	case 0xC0262589:
		return ErrorGraphicsDdcciInvalidMessageCommand
	case 0xC026258A:
		return ErrorGraphicsDdcciInvalidMessageLength
	case 0xC026258B:
		return ErrorGraphicsDdcciInvalidMessageChecksum
	case 0xC02625D6:
		return ErrorGraphicsPmeaInvalidMonitor
	case 0xC02625D7:
		return ErrorGraphicsPmeaInvalidD3dDevice
	case 0xC02625D8:
		return ErrorGraphicsDdcciCurrentCurrentValueGreaterThanMaximumValue
	case 0xC02625D9:
		return ErrorGraphicsMcaInvalidVcpVersion
	case 0xC02625DA:
		return ErrorGraphicsMcaMonitorViolatesMccsSpecification
	case 0xC02625DB:
		return ErrorGraphicsMcaMccsVersionMismatch
	case 0xC02625DC:
		return ErrorGraphicsMcaUnsupportedMccsVersion
	case 0xC02625DE:
		return ErrorGraphicsMcaInvalidTechnologyTypeReturned
	case 0xC02625DF:
		return ErrorGraphicsMcaUnsupportedColorTemperature
	case 0xC02625E0:
		return ErrorGraphicsOnlyConsoleSessionSupported
	case 0x800C0009:
		return InetEAuthenticationRequired
	case 0x800C001B:
		return InetEBlockedRedirectXsecurityid
	case 0x800C0004:
		return InetECannotConnect
	case 0x800C0010:
		return InetECannotInstantiateObject
	case 0x800C000F:
		return InetECannotLoadData
	case 0x800C0016:
		return InetECannotLockRequest
	case 0x800C0300:
		return InetECannotReplaceSfpFile
	case 0x800C0100:
		return InetECodeDownloadDeclined
	case 0x800C0400:
		return InetECodeInstallSuppressed
	case 0x800C0500:
		return InetECodeInstallBlockedByHashPolicy
	case 0x800C000B:
		return InetEConnectionTimeout
	case 0x800C0007:
		return InetEDataNotAvailable
	case 0x800C0011:
		return InetEDefaultAction
	case 0x800C0501:
		return InetEDownloadBlockedByInprivate
	case 0x800C0008:
		return InetEDownloadFailure
	case 0x800C0002:
		return InetEErrorFirst
	case 0x800C0019:
		return InetEInvalidCertificate
	case 0x800C000C:
		return InetEInvalidRequest
	case 0x800C0003:
		return InetENoSession
	case 0x800C000A:
		return InetENoValidMedia
	case 0x800C0006:
		return InetEObjectNotFound
	case 0x800C0013:
		return InetEQueryoptionUnknown
	case 0x800C0014:
		return InetERedirectFailed
	case 0x800C0015:
		return InetERedirectToDir
	case 0x800C001A:
		return InetEReserved1
	case 0x800C0005:
		return InetEResourceNotFound
	case 0x800C0200:
		return InetEResultDispatched
	case 0x800C000E:
		return InetESecurityProblem
	case 0x800C0018:
		return InetETerminatedBind
	case 0x800C000D:
		return InetEUnknownProtocol
	case 0x800C0012:
		return InetEUseDefaultSetting
	case 0x800C0017:
		return InetEUseExtendBinding
	}
	return nil
}

