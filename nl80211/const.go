// WARNING: This file has automatically been generated on Wed, 28 Dec 2016 02:32:08 UTC.
// By https://git.io/cgogen. DO NOT EDIT.

package nl80211

const (
	// GenlName as defined in nl80211/nl80211.h:44
	GenlName = "nl80211"
	// MulticastGroupConfig as defined in nl80211/nl80211.h:46
	MulticastGroupConfig = "config"
	// MulticastGroupScan as defined in nl80211/nl80211.h:47
	MulticastGroupScan = "scan"
	// MulticastGroupReg as defined in nl80211/nl80211.h:48
	MulticastGroupReg = "regulatory"
	// MulticastGroupMlme as defined in nl80211/nl80211.h:49
	MulticastGroupMlme = "mlme"
	// MulticastGroupVendor as defined in nl80211/nl80211.h:50
	MulticastGroupVendor = "vendor"
	// MulticastGroupNan as defined in nl80211/nl80211.h:51
	MulticastGroupNan = "nan"
	// MulticastGroupTestmode as defined in nl80211/nl80211.h:52
	MulticastGroupTestmode = "testmode"
	// CmdGetMeshParams as defined in nl80211/nl80211.h:1080
	CmdGetMeshParams = 0
	// CmdSetMeshParams as defined in nl80211/nl80211.h:1081
	CmdSetMeshParams = 0
	// MeshSetupVendorPathSelIe as defined in nl80211/nl80211.h:1082
	MeshSetupVendorPathSelIe = 0
	// AttrScanGeneration as defined in nl80211/nl80211.h:2332
	AttrScanGeneration = 0
	// AttrMeshParams as defined in nl80211/nl80211.h:2333
	AttrMeshParams = 0
	// AttrIfaceSocketOwner as defined in nl80211/nl80211.h:2334
	AttrIfaceSocketOwner = 0
	// MaxSuppRates as defined in nl80211/nl80211.h:2336
	MaxSuppRates = 32
	// MaxSuppHtRates as defined in nl80211/nl80211.h:2337
	MaxSuppHtRates = 77
	// MaxSuppRegRules as defined in nl80211/nl80211.h:2338
	MaxSuppRegRules = 64
	// TkipDataOffsetEncrKey as defined in nl80211/nl80211.h:2339
	TkipDataOffsetEncrKey = 0
	// TkipDataOffsetTxMicKey as defined in nl80211/nl80211.h:2340
	TkipDataOffsetTxMicKey = 16
	// TkipDataOffsetRxMicKey as defined in nl80211/nl80211.h:2341
	TkipDataOffsetRxMicKey = 24
	// HtCapabilityLen as defined in nl80211/nl80211.h:2342
	HtCapabilityLen = 26
	// VhtCapabilityLen as defined in nl80211/nl80211.h:2343
	VhtCapabilityLen = 12
	// MaxNrCipherSuites as defined in nl80211/nl80211.h:2345
	MaxNrCipherSuites = 5
	// MaxNrAkmSuites as defined in nl80211/nl80211.h:2346
	MaxNrAkmSuites = 2
	// MinRemainOnChannelTime as defined in nl80211/nl80211.h:2348
	MinRemainOnChannelTime = 10
	// ScanRssiTholdOff as defined in nl80211/nl80211.h:2351
	ScanRssiTholdOff = -300
	// CqmTxeMaxIntvl as defined in nl80211/nl80211.h:2353
	CqmTxeMaxIntvl = 1800
	// StaFlagMaxOldApi as defined in nl80211/nl80211.h:2457
	StaFlagMaxOldApi = 0
	// FrequencyAttrPassiveScan as defined in nl80211/nl80211.h:2858
	FrequencyAttrPassiveScan = 0
	// FrequencyAttrNoIbss as defined in nl80211/nl80211.h:2859
	FrequencyAttrNoIbss = 0
	// FrequencyAttrGoConcurrent as defined in nl80211/nl80211.h:2860
	FrequencyAttrGoConcurrent = 0
	// AttrSchedScanMatchSsid as defined in nl80211/nl80211.h:3002
	AttrSchedScanMatchSsid = 0
	// RrfPassiveScan as defined in nl80211/nl80211.h:3044
	RrfPassiveScan = 0
	// RrfNoIbss as defined in nl80211/nl80211.h:3045
	RrfNoIbss = 0
	// RrfNoHt40 as defined in nl80211/nl80211.h:3046
	RrfNoHt40 = 0
	// RrfGoConcurrent as defined in nl80211/nl80211.h:3048
	RrfGoConcurrent = 0
	// RrfNoIrAll as defined in nl80211/nl80211.h:3051
	RrfNoIrAll = 0
	// SurveyInfoChannelTime as defined in nl80211/nl80211.h:3137
	SurveyInfoChannelTime = 0
	// SurveyInfoChannelTimeBusy as defined in nl80211/nl80211.h:3138
	SurveyInfoChannelTimeBusy = 0
	// SurveyInfoChannelTimeExtBusy as defined in nl80211/nl80211.h:3139
	SurveyInfoChannelTimeExtBusy = 0
	// SurveyInfoChannelTimeRx as defined in nl80211/nl80211.h:3140
	SurveyInfoChannelTimeRx = 0
	// SurveyInfoChannelTimeTx as defined in nl80211/nl80211.h:3141
	SurveyInfoChannelTimeTx = 0
	// TxqAttrQueue as defined in nl80211/nl80211.h:3449
	TxqAttrQueue = 0
	// TxqQVo as defined in nl80211/nl80211.h:3450
	TxqQVo = 0
	// TxqQVi as defined in nl80211/nl80211.h:3451
	TxqQVi = 0
	// TxqQBe as defined in nl80211/nl80211.h:3452
	TxqQBe = 0
	// TxqQBk as defined in nl80211/nl80211.h:3453
	TxqQBk = 0
	// TxrateMcs as defined in nl80211/nl80211.h:3748
	TxrateMcs = 0
	// VhtNssMax as defined in nl80211/nl80211.h:3749
	VhtNssMax = 8
	// __WowlanPktpatInvalid as defined in nl80211/nl80211.h:3914
	__WowlanPktpatInvalid = 0
	// WowlanPktpatMask as defined in nl80211/nl80211.h:3915
	WowlanPktpatMask = 0
	// WowlanPktpatPattern as defined in nl80211/nl80211.h:3916
	WowlanPktpatPattern = 0
	// WowlanPktpatOffset as defined in nl80211/nl80211.h:3917
	WowlanPktpatOffset = 0
	// Num_WowlanPktpat as defined in nl80211/nl80211.h:3918
	Num_WowlanPktpat = 0
	// Max_WowlanPktpat as defined in nl80211/nl80211.h:3919
	Max_WowlanPktpat = 0
	// WowlanPatternSupport as defined in nl80211/nl80211.h:3920
	WowlanPatternSupport = 0
	// KckLen as defined in nl80211/nl80211.h:4327
	KckLen = 16
	// KekLen as defined in nl80211/nl80211.h:4328
	KekLen = 16
	// ReplayCtrLen as defined in nl80211/nl80211.h:4329
	ReplayCtrLen = 8
	// CritProtoMaxDuration as defined in nl80211/nl80211.h:4781
	CritProtoMaxDuration = 5000
	// VendorIdIsLinux as defined in nl80211/nl80211.h:4799
	VendorIdIsLinux = 0x80000000
	// NanFuncServiceIdLen as defined in nl80211/nl80211.h:4960
	NanFuncServiceIdLen = 6
	// NanFuncServiceSpecInfoMaxLen as defined in nl80211/nl80211.h:4961
	NanFuncServiceSpecInfoMaxLen = 0xff
	// NanFuncSrfMaxLen as defined in nl80211/nl80211.h:4962
	NanFuncSrfMaxLen = 0xff
)

// nl80211Commands as declared in nl80211/nl80211.h:880
type nl80211Commands int32

// nl80211Commands enumeration from nl80211/nl80211.h:880
const (
	CmdUnspec = iota
	CmdGetWiphy
	CmdSetWiphy
	CmdNewWiphy
	CmdDelWiphy
	CmdGetInterface
	CmdSetInterface
	CmdNewInterface
	CmdDelInterface
	CmdGetKey
	CmdSetKey
	CmdNewKey
	CmdDelKey
	CmdGetBeacon
	CmdSetBeacon
	CmdStartAp
	CmdNewBeacon = CmdStartAp
	CmdStopAp
	CmdDelBeacon = CmdStopAp
	CmdGetStation
	CmdSetStation
	CmdNewStation
	CmdDelStation
	CmdGetMpath
	CmdSetMpath
	CmdNewMpath
	CmdDelMpath
	CmdSetBss
	CmdSetReg
	CmdReqSetReg
	CmdGetMeshConfig
	CmdSetMeshConfig
	CmdSetMgmtExtraIe
	CmdGetReg
	CmdGetScan
	CmdTriggerScan
	CmdNewScanResults
	CmdScanAborted
	CmdRegChange
	CmdAuthenticate
	CmdAssociate
	CmdDeauthenticate
	CmdDisassociate
	CmdMichaelMicFailure
	CmdRegBeaconHint
	CmdJoinIbss
	CmdLeaveIbss
	CmdTestmode
	CmdConnect
	CmdRoam
	CmdDisconnect
	CmdSetWiphyNetns
	CmdGetSurvey
	CmdNewSurveyResults
	CmdSetPmksa
	CmdDelPmksa
	CmdFlushPmksa
	CmdRemainOnChannel
	CmdCancelRemainOnChannel
	CmdSetTxBitrateMask
	CmdRegisterFrame
	CmdRegisterAction = CmdRegisterFrame
	CmdFrame
	CmdAction = CmdFrame
	CmdFrameTxStatus
	CmdActionTxStatus = CmdFrameTxStatus
	CmdSetPowerSave
	CmdGetPowerSave
	CmdSetCqm
	CmdNotifyCqm
	CmdSetChannel
	CmdSetWdsPeer
	CmdFrameWaitCancel
	CmdJoinMesh
	CmdLeaveMesh
	CmdUnprotDeauthenticate
	CmdUnprotDisassociate
	CmdNewPeerCandidate
	CmdGetWowlan
	CmdSetWowlan
	CmdStartSchedScan
	CmdStopSchedScan
	CmdSchedScanResults
	CmdSchedScanStopped
	CmdSetRekeyOffload
	CmdPmksaCandidate
	CmdTdlsOper
	CmdTdlsMgmt
	CmdUnexpectedFrame
	CmdProbeClient
	CmdRegisterBeacons
	CmdUnexpected4addrFrame
	CmdSetNoackMap
	CmdChSwitchNotify
	CmdStartP2pDevice
	CmdStopP2pDevice
	CmdConnFailed
	CmdSetMcastRate
	CmdSetMacAcl
	CmdRadarDetect
	CmdGetProtocolFeatures
	CmdUpdateFtIes
	CmdFtEvent
	CmdCritProtocolStart
	CmdCritProtocolStop
	CmdGetCoalesce
	CmdSetCoalesce
	CmdChannelSwitch
	CmdVendor
	CmdSetQosMap
	CmdAddTxTs
	CmdDelTxTs
	CmdGetMpp
	CmdJoinOcb
	CmdLeaveOcb
	CmdChSwitchStartedNotify
	CmdTdlsChannelSwitch
	CmdTdlsCancelChannelSwitch
	CmdWiphyRegChange
	CmdAbortScan
	CmdStartNan
	CmdStopNan
	CmdAddNanFunction
	CmdDelNanFunction
	CmdChangeNanConfig
	CmdNanMatch
	__CmdAfterLast
	CmdMax = __CmdAfterLast - 1
)

// nl80211Attrs as declared in nl80211/nl80211.h:1929
type nl80211Attrs int32

// nl80211Attrs enumeration from nl80211/nl80211.h:1929
const (
	AttrUnspec = iota
	AttrWiphy
	AttrWiphyName
	AttrIfindex
	AttrIfname
	AttrIftype
	AttrMac
	AttrKeyData
	AttrKeyIdx
	AttrKeyCipher
	AttrKeySeq
	AttrKeyDefault
	AttrBeaconInterval
	AttrDtimPeriod
	AttrBeaconHead
	AttrBeaconTail
	AttrStaAid
	AttrStaFlags
	AttrStaListenInterval
	AttrStaSupportedRates
	AttrStaVlan
	AttrStaInfo
	AttrWiphyBands
	AttrMntrFlags
	AttrMeshId
	AttrStaPlinkAction
	AttrMpathNextHop
	AttrMpathInfo
	AttrBssCtsProt
	AttrBssShortPreamble
	AttrBssShortSlotTime
	AttrHtCapability
	AttrSupportedIftypes
	AttrRegAlpha2
	AttrRegRules
	AttrMeshConfig
	AttrBssBasicRates
	AttrWiphyTxqParams
	AttrWiphyFreq
	AttrWiphyChannelType
	AttrKeyDefaultMgmt
	AttrMgmtSubtype
	AttrIe
	AttrMaxNumScanSsids
	AttrScanFrequencies
	AttrScanSsids
	AttrGeneration
	AttrBss
	AttrRegInitiator
	AttrRegType
	AttrSupportedCommands
	AttrFrame
	AttrSsid
	AttrAuthType
	AttrReasonCode
	AttrKeyType
	AttrMaxScanIeLen
	AttrCipherSuites
	AttrFreqBefore
	AttrFreqAfter
	AttrFreqFixed
	AttrWiphyRetryShort
	AttrWiphyRetryLong
	AttrWiphyFragThreshold
	AttrWiphyRtsThreshold
	AttrTimedOut
	AttrUseMfp
	AttrStaFlags2
	AttrControlPort
	AttrTestdata
	AttrPrivacy
	AttrDisconnectedByAp
	AttrStatusCode
	AttrCipherSuitesPairwise
	AttrCipherSuiteGroup
	AttrWpaVersions
	AttrAkmSuites
	AttrReqIe
	AttrRespIe
	AttrPrevBssid
	AttrKey
	AttrKeys
	AttrPid
	Attr4addr
	AttrSurveyInfo
	AttrPmkid
	AttrMaxNumPmkids
	AttrDuration
	AttrCookie
	AttrWiphyCoverageClass
	AttrTxRates
	AttrFrameMatch
	AttrAck
	AttrPsState
	AttrCqm
	AttrLocalStateChange
	AttrApIsolate
	AttrWiphyTxPowerSetting
	AttrWiphyTxPowerLevel
	AttrTxFrameTypes
	AttrRxFrameTypes
	AttrFrameType
	AttrControlPortEthertype
	AttrControlPortNoEncrypt
	AttrSupportIbssRsn
	AttrWiphyAntennaTx
	AttrWiphyAntennaRx
	AttrMcastRate
	AttrOffchannelTxOk
	AttrBssHtOpmode
	AttrKeyDefaultTypes
	AttrMaxRemainOnChannelDuration
	AttrMeshSetup
	AttrWiphyAntennaAvailTx
	AttrWiphyAntennaAvailRx
	AttrSupportMeshAuth
	AttrStaPlinkState
	AttrWowlanTriggers
	AttrWowlanTriggersSupported
	AttrSchedScanInterval
	AttrInterfaceCombinations
	AttrSoftwareIftypes
	AttrRekeyData
	AttrMaxNumSchedScanSsids
	AttrMaxSchedScanIeLen
	AttrScanSuppRates
	AttrHiddenSsid
	AttrIeProbeResp
	AttrIeAssocResp
	AttrStaWme
	AttrSupportApUapsd
	AttrRoamSupport
	AttrSchedScanMatch
	AttrMaxMatchSets
	AttrPmksaCandidate
	AttrTxNoCckRate
	AttrTdlsAction
	AttrTdlsDialogToken
	AttrTdlsOperation
	AttrTdlsSupport
	AttrTdlsExternalSetup
	AttrDeviceApSme
	AttrDontWaitForAck
	AttrFeatureFlags
	AttrProbeRespOffload
	AttrProbeResp
	AttrDfsRegion
	AttrDisableHt
	AttrHtCapabilityMask
	AttrNoackMap
	AttrInactivityTimeout
	AttrRxSignalDbm
	AttrBgScanPeriod
	AttrWdev
	AttrUserRegHintType
	AttrConnFailedReason
	AttrSaeData
	AttrVhtCapability
	AttrScanFlags
	AttrChannelWidth
	AttrCenterFreq1
	AttrCenterFreq2
	AttrP2pCtwindow
	AttrP2pOppps
	AttrLocalMeshPowerMode
	AttrAclPolicy
	AttrMacAddrs
	AttrMacAclMax
	AttrRadarEvent
	AttrExtCapa
	AttrExtCapaMask
	AttrStaCapability
	AttrStaExtCapability
	AttrProtocolFeatures
	AttrSplitWiphyDump
	AttrDisableVht
	AttrVhtCapabilityMask
	AttrMdid
	AttrIeRic
	AttrCritProtId
	AttrMaxCritProtDuration
	AttrPeerAid
	AttrCoalesceRule
	AttrChSwitchCount
	AttrChSwitchBlockTx
	AttrCsaIes
	AttrCsaCOffBeacon
	AttrCsaCOffPresp
	AttrRxmgmtFlags
	AttrStaSupportedChannels
	AttrStaSupportedOperClasses
	AttrHandleDfs
	AttrSupport5Mhz
	AttrSupport10Mhz
	AttrOpmodeNotif
	AttrVendorId
	AttrVendorSubcmd
	AttrVendorData
	AttrVendorEvents
	AttrQosMap
	AttrMacHint
	AttrWiphyFreqHint
	AttrMaxApAssocSta
	AttrTdlsPeerCapability
	AttrSocketOwner
	AttrCsaCOffsetsTx
	AttrMaxCsaCounters
	AttrTdlsInitiator
	AttrUseRrm
	AttrWiphyDynAck
	AttrTsid
	AttrUserPrio
	AttrAdmittedTime
	AttrSmpsMode
	AttrOperClass
	AttrMacMask
	AttrWiphySelfManagedReg
	AttrExtFeatures
	AttrSurveyRadioStats
	AttrNetnsFd
	AttrSchedScanDelay
	AttrRegIndoor
	AttrMaxNumSchedScanPlans
	AttrMaxScanPlanInterval
	AttrMaxScanPlanIterations
	AttrSchedScanPlans
	AttrPbss
	AttrBssSelect
	AttrStaSupportP2pPs
	AttrPad
	AttrIftypeExtCapa
	AttrMuMimoGroupData
	AttrMuMimoFollowMacAddr
	AttrScanStartTimeTsf
	AttrScanStartTimeTsfBssid
	AttrMeasurementDuration
	AttrMeasurementDurationMandatory
	AttrMeshPeerAid
	AttrNanMasterPref
	AttrNanDual
	AttrNanFunc
	AttrNanMatch
	__AttrAfterLast
	Num_Attr = __AttrAfterLast
	AttrMax  = __AttrAfterLast - 1
)

// nl80211Iftype as declared in nl80211/nl80211.h:2384
type nl80211Iftype int32

// nl80211Iftype enumeration from nl80211/nl80211.h:2384
const (
	IftypeUnspecified = iota
	IftypeAdhoc
	IftypeStation
	IftypeAp
	IftypeApVlan
	IftypeWds
	IftypeMonitor
	IftypeMeshPoint
	IftypeP2pClient
	IftypeP2pGo
	IftypeP2pDevice
	IftypeOcb
	IftypeNan
	Num_Iftypes
	IftypeMax = Num_Iftypes - 1
)

// nl80211StaFlags as declared in nl80211/nl80211.h:2428
type nl80211StaFlags int32

// nl80211StaFlags enumeration from nl80211/nl80211.h:2428
const (
	__StaFlagInvalid = iota
	StaFlagAuthorized
	StaFlagShortPreamble
	StaFlagWme
	StaFlagMfp
	StaFlagAuthenticated
	StaFlagTdlsPeer
	StaFlagAssociated
	__StaFlagAfterLast
	StaFlagMax = __StaFlagAfterLast - 1
)

// nl80211StaP2pPsStatus as declared in nl80211/nl80211.h:2450
type nl80211StaP2pPsStatus int32

// nl80211StaP2pPsStatus enumeration from nl80211/nl80211.h:2450
const (
	P2pPsUnsupported = iota
	P2pPsSupported
	Num_P2pPsStatus
)

// nl80211RateInfo as declared in nl80211/nl80211.h:2505
type nl80211RateInfo int32

// nl80211RateInfo enumeration from nl80211/nl80211.h:2505
const (
	__RateInfoInvalid = iota
	RateInfoBitrate
	RateInfoMcs
	RateInfo40MhzWidth
	RateInfoShortGi
	RateInfoBitrate32
	RateInfoVhtMcs
	RateInfoVhtNss
	RateInfo80MhzWidth
	RateInfo80p80MhzWidth
	RateInfo160MhzWidth
	RateInfo10MhzWidth
	RateInfo5MhzWidth
	__RateInfoAfterLast
	RateInfoMax = __RateInfoAfterLast - 1
)

// nl80211StaBssParam as declared in nl80211/nl80211.h:2542
type nl80211StaBssParam int32

// nl80211StaBssParam enumeration from nl80211/nl80211.h:2542
const (
	__StaBssParamInvalid = iota
	StaBssParamCtsProt
	StaBssParamShortPreamble
	StaBssParamShortSlotTime
	StaBssParamDtimPeriod
	StaBssParamBeaconInterval
	__StaBssParamAfterLast
	StaBssParamMax = __StaBssParamAfterLast - 1
)

// nl80211StaInfo as declared in nl80211/nl80211.h:2620
type nl80211StaInfo int32

// nl80211StaInfo enumeration from nl80211/nl80211.h:2620
const (
	__StaInfoInvalid = iota
	StaInfoInactiveTime
	StaInfoRxBytes
	StaInfoTxBytes
	StaInfoLlid
	StaInfoPlid
	StaInfoPlinkState
	StaInfoSignal
	StaInfoTxBitrate
	StaInfoRxPackets
	StaInfoTxPackets
	StaInfoTxRetries
	StaInfoTxFailed
	StaInfoSignalAvg
	StaInfoRxBitrate
	StaInfoBssParam
	StaInfoConnectedTime
	StaInfoStaFlags
	StaInfoBeaconLoss
	StaInfoTOffset
	StaInfoLocalPm
	StaInfoPeerPm
	StaInfoNonpeerPm
	StaInfoRxBytes64
	StaInfoTxBytes64
	StaInfoChainSignal
	StaInfoChainSignalAvg
	StaInfoExpectedThroughput
	StaInfoRxDropMisc
	StaInfoBeaconRx
	StaInfoBeaconSignalAvg
	StaInfoTidStats
	StaInfoRxDuration
	StaInfoPad
	__StaInfoAfterLast
	StaInfoMax = __StaInfoAfterLast - 1
)

// nl80211TidStats as declared in nl80211/nl80211.h:2675
type nl80211TidStats int32

// nl80211TidStats enumeration from nl80211/nl80211.h:2675
const (
	__TidStatsInvalid = iota
	TidStatsRxMsdu
	TidStatsTxMsdu
	TidStatsTxMsduRetries
	TidStatsTxMsduFailed
	TidStatsPad
	Num_TidStats
	TidStatsMax = Num_TidStats - 1
)

// nl80211MpathFlags as declared in nl80211/nl80211.h:2697
type nl80211MpathFlags int32

// nl80211MpathFlags enumeration from nl80211/nl80211.h:2697
const (
	MpathFlagActive    = 1 << 0
	MpathFlagResolving = 1 << 1
	MpathFlagSnValid   = 1 << 2
	MpathFlagFixed     = 1 << 3
	MpathFlagResolved  = 1 << 4
)

// nl80211MpathInfo as declared in nl80211/nl80211.h:2724
type nl80211MpathInfo int32

// nl80211MpathInfo enumeration from nl80211/nl80211.h:2724
const (
	__MpathInfoInvalid = iota
	MpathInfoFrameQlen
	MpathInfoSn
	MpathInfoMetric
	MpathInfoExptime
	MpathInfoFlags
	MpathInfoDiscoveryTimeout
	MpathInfoDiscoveryRetries
	__MpathInfoAfterLast
	MpathInfoMax = __MpathInfoAfterLast - 1
)

// nl80211BandAttr as declared in nl80211/nl80211.h:2757
type nl80211BandAttr int32

// nl80211BandAttr enumeration from nl80211/nl80211.h:2757
const (
	__BandAttrInvalid = iota
	BandAttrFreqs
	BandAttrRates
	BandAttrHtMcsSet
	BandAttrHtCapa
	BandAttrHtAmpduFactor
	BandAttrHtAmpduDensity
	BandAttrVhtMcsSet
	BandAttrVhtCapa
	__BandAttrAfterLast
	BandAttrMax = __BandAttrAfterLast - 1
)

// nl80211FrequencyAttr as declared in nl80211/nl80211.h:2833
type nl80211FrequencyAttr int32

// nl80211FrequencyAttr enumeration from nl80211/nl80211.h:2833
const (
	__FrequencyAttrInvalid = iota
	FrequencyAttrFreq
	FrequencyAttrDisabled
	FrequencyAttrNoIr
	__FrequencyAttrNoIbss
	FrequencyAttrRadar
	FrequencyAttrMaxTxPower
	FrequencyAttrDfsState
	FrequencyAttrDfsTime
	FrequencyAttrNoHt40Minus
	FrequencyAttrNoHt40Plus
	FrequencyAttrNo80mhz
	FrequencyAttrNo160mhz
	FrequencyAttrDfsCacTime
	FrequencyAttrIndoorOnly
	FrequencyAttrIrConcurrent
	FrequencyAttrNo20mhz
	FrequencyAttrNo10mhz
	__FrequencyAttrAfterLast
	FrequencyAttrMax = __FrequencyAttrAfterLast - 1
)

// nl80211BitrateAttr as declared in nl80211/nl80211.h:2873
type nl80211BitrateAttr int32

// nl80211BitrateAttr enumeration from nl80211/nl80211.h:2873
const (
	__BitrateAttrInvalid = iota
	BitrateAttrRate
	BitrateAttr2ghzShortpreamble
	__BitrateAttrAfterLast
	BitrateAttrMax = __BitrateAttrAfterLast - 1
)

// nl80211RegInitiator as declared in nl80211/nl80211.h:2899
type nl80211RegInitiator int32

// nl80211RegInitiator enumeration from nl80211/nl80211.h:2899
const (
	RegdomSetByCore = iota
	RegdomSetByUser
	RegdomSetByDriver
	RegdomSetByCountryIe
)

// nl80211RegType as declared in nl80211/nl80211.h:2922
type nl80211RegType int32

// nl80211RegType enumeration from nl80211/nl80211.h:2922
const (
	RegdomTypeCountry = iota
	RegdomTypeWorld
	RegdomTypeCustomWorld
	RegdomTypeIntersection
)

// nl80211RegRuleAttr as declared in nl80211/nl80211.h:2954
type nl80211RegRuleAttr int32

// nl80211RegRuleAttr enumeration from nl80211/nl80211.h:2954
const (
	__RegRuleAttrInvalid = iota
	AttrRegRuleFlags
	AttrFreqRangeStart
	AttrFreqRangeEnd
	AttrFreqRangeMaxBw
	AttrPowerRuleMaxAntGain
	AttrPowerRuleMaxEirp
	AttrDfsCacTime
	__RegRuleAttrAfterLast
	RegRuleAttrMax = __RegRuleAttrAfterLast - 1
)

// nl80211SchedScanMatchAttr as declared in nl80211/nl80211.h:2989
type nl80211SchedScanMatchAttr int32

// nl80211SchedScanMatchAttr enumeration from nl80211/nl80211.h:2989
const (
	__SchedScanMatchAttrInvalid = iota
	SchedScanMatchAttrSsid
	SchedScanMatchAttrRssi
	__SchedScanMatchAttrAfterLast
	SchedScanMatchAttrMax = __SchedScanMatchAttrAfterLast - 1
)

// nl80211RegRuleFlags as declared in nl80211/nl80211.h:3026
type nl80211RegRuleFlags int32

// nl80211RegRuleFlags enumeration from nl80211/nl80211.h:3026
const (
	RrfNoOfdm       = 1 << 0
	RrfNoCck        = 1 << 1
	RrfNoIndoor     = 1 << 2
	RrfNoOutdoor    = 1 << 3
	RrfDfs          = 1 << 4
	RrfPtpOnly      = 1 << 5
	RrfPtmpOnly     = 1 << 6
	RrfNoIr         = 1 << 7
	__RrfNoIbss     = 1 << 8
	RrfAutoBw       = 1 << 11
	RrfIrConcurrent = 1 << 12
	RrfNoHt40minus  = 1 << 13
	RrfNoHt40plus   = 1 << 14
	RrfNo80mhz      = 1 << 15
	RrfNo160mhz     = 1 << 16
)

// nl80211DfsRegions as declared in nl80211/nl80211.h:3061
type nl80211DfsRegions int32

// nl80211DfsRegions enumeration from nl80211/nl80211.h:3061
const (
	DfsUnset = iota
	DfsFcc   = 1
	DfsEtsi  = 2
	DfsJp    = 3
)

// nl80211UserRegHintType as declared in nl80211/nl80211.h:3085
type nl80211UserRegHintType int32

// nl80211UserRegHintType enumeration from nl80211/nl80211.h:3085
const (
	UserRegHintUser     = iota
	UserRegHintCellBase = 1
	UserRegHintIndoor   = 2
)

// nl80211SurveyInfo as declared in nl80211/nl80211.h:3118
type nl80211SurveyInfo int32

// nl80211SurveyInfo enumeration from nl80211/nl80211.h:3118
const (
	__SurveyInfoInvalid = iota
	SurveyInfoFrequency
	SurveyInfoNoise
	SurveyInfoInUse
	SurveyInfoTime
	SurveyInfoTimeBusy
	SurveyInfoTimeExtBusy
	SurveyInfoTimeRx
	SurveyInfoTimeTx
	SurveyInfoTimeScan
	SurveyInfoPad
	__SurveyInfoAfterLast
	SurveyInfoMax = __SurveyInfoAfterLast - 1
)

// nl80211MntrFlags as declared in nl80211/nl80211.h:3162
type nl80211MntrFlags int32

// nl80211MntrFlags enumeration from nl80211/nl80211.h:3162
const (
	__MntrFlagInvalid = iota
	MntrFlagFcsfail
	MntrFlagPlcpfail
	MntrFlagControl
	MntrFlagOtherBss
	MntrFlagCookFrames
	MntrFlagActive
	__MntrFlagAfterLast
	MntrFlagMax = __MntrFlagAfterLast - 1
)

// nl80211MeshPowerMode as declared in nl80211/nl80211.h:3194
type nl80211MeshPowerMode int32

// nl80211MeshPowerMode enumeration from nl80211/nl80211.h:3194
const (
	MeshPowerUnknown = iota
	MeshPowerActive
	MeshPowerLightSleep
	MeshPowerDeepSleep
	__MeshPowerAfterLast
	MeshPowerMax = __MeshPowerAfterLast - 1
)

// nl80211MeshconfParams as declared in nl80211/nl80211.h:3312
type nl80211MeshconfParams int32

// nl80211MeshconfParams enumeration from nl80211/nl80211.h:3312
const (
	__MeshconfInvalid = iota
	MeshconfRetryTimeout
	MeshconfConfirmTimeout
	MeshconfHoldingTimeout
	MeshconfMaxPeerLinks
	MeshconfMaxRetries
	MeshconfTtl
	MeshconfAutoOpenPlinks
	MeshconfHwmpMaxPreqRetries
	MeshconfPathRefreshTime
	MeshconfMinDiscoveryTimeout
	MeshconfHwmpActivePathTimeout
	MeshconfHwmpPreqMinInterval
	MeshconfHwmpNetDiamTrvsTime
	MeshconfHwmpRootmode
	MeshconfElementTtl
	MeshconfHwmpRannInterval
	MeshconfGateAnnouncements
	MeshconfHwmpPerrMinInterval
	MeshconfForwarding
	MeshconfRssiThreshold
	MeshconfSyncOffsetMaxNeighbor
	MeshconfHtOpmode
	MeshconfHwmpPathToRootTimeout
	MeshconfHwmpRootInterval
	MeshconfHwmpConfirmationInterval
	MeshconfPowerMode
	MeshconfAwakeWindow
	MeshconfPlinkTimeout
	__MeshconfAttrAfterLast
	MeshconfAttrMax = __MeshconfAttrAfterLast - 1
)

// nl80211MeshSetupParams as declared in nl80211/nl80211.h:3397
type nl80211MeshSetupParams int32

// nl80211MeshSetupParams enumeration from nl80211/nl80211.h:3397
const (
	__MeshSetupInvalid = iota
	MeshSetupEnableVendorPathSel
	MeshSetupEnableVendorMetric
	MeshSetupIe
	MeshSetupUserspaceAuth
	MeshSetupUserspaceAmpe
	MeshSetupEnableVendorSync
	MeshSetupUserspaceMpm
	MeshSetupAuthProtocol
	__MeshSetupAttrAfterLast
	MeshSetupAttrMax = __MeshSetupAttrAfterLast - 1
)

// nl80211TxqAttr as declared in nl80211/nl80211.h:3427
type nl80211TxqAttr int32

// nl80211TxqAttr enumeration from nl80211/nl80211.h:3427
const (
	__TxqAttrInvalid = iota
	TxqAttrAc
	TxqAttrTxop
	TxqAttrCwmin
	TxqAttrCwmax
	TxqAttrAifs
	__TxqAttrAfterLast
	TxqAttrMax = __TxqAttrAfterLast - 1
)

// nl80211Ac as declared in nl80211/nl80211.h:3440
type nl80211Ac int32

// nl80211Ac enumeration from nl80211/nl80211.h:3440
const (
	AcVo = iota
	AcVi
	AcBe
	AcBk
	NumAcs
)

// nl80211ChannelType as declared in nl80211/nl80211.h:3464
type nl80211ChannelType int32

// nl80211ChannelType enumeration from nl80211/nl80211.h:3464
const (
	ChanNoHt = iota
	ChanHt20
	ChanHt40minus
	ChanHt40plus
)

// nl80211ChanWidth as declared in nl80211/nl80211.h:3490
type nl80211ChanWidth int32

// nl80211ChanWidth enumeration from nl80211/nl80211.h:3490
const (
	ChanWidth20Noht = iota
	ChanWidth20
	ChanWidth40
	ChanWidth80
	ChanWidth80p80
	ChanWidth160
	ChanWidth5
	ChanWidth10
)

// nl80211BssScanWidth as declared in nl80211/nl80211.h:3510
type nl80211BssScanWidth int32

// nl80211BssScanWidth enumeration from nl80211/nl80211.h:3510
const (
	BssChanWidth20 = iota
	BssChanWidth10
	BssChanWidth5
)

// nl80211Bss as declared in nl80211/nl80211.h:3565
type nl80211Bss int32

// nl80211Bss enumeration from nl80211/nl80211.h:3565
const (
	__BssInvalid = iota
	BssBssid
	BssFrequency
	BssTsf
	BssBeaconInterval
	BssCapability
	BssInformationElements
	BssSignalMbm
	BssSignalUnspec
	BssStatus
	BssSeenMsAgo
	BssBeaconIes
	BssChanWidth
	BssBeaconTsf
	BssPrespData
	BssLastSeenBoottime
	BssPad
	BssParentTsf
	BssParentBssid
	__BssAfterLast
	BssMax = __BssAfterLast - 1
)

// nl80211BssStatus as declared in nl80211/nl80211.h:3603
type nl80211BssStatus int32

// nl80211BssStatus enumeration from nl80211/nl80211.h:3603
const (
	BssStatusAuthenticated = iota
	BssStatusAssociated
	BssStatusIbssJoined
)

// nl80211AuthType as declared in nl80211/nl80211.h:3623
type nl80211AuthType int32

// nl80211AuthType enumeration from nl80211/nl80211.h:3623
const (
	AuthtypeOpenSystem = iota
	AuthtypeSharedKey
	AuthtypeFt
	AuthtypeNetworkEap
	AuthtypeSae
	__AuthtypeNum
	AuthtypeMax = __AuthtypeNum - 1
	AuthtypeAutomatic
)

// nl80211KeyType as declared in nl80211/nl80211.h:3643
type nl80211KeyType int32

// nl80211KeyType enumeration from nl80211/nl80211.h:3643
const (
	KeytypeGroup = iota
	KeytypePairwise
	KeytypePeerkey
	Num_Keytypes
)

// nl80211Mfp as declared in nl80211/nl80211.h:3656
type nl80211Mfp int32

// nl80211Mfp enumeration from nl80211/nl80211.h:3656
const (
	MfpNo = iota
	MfpRequired
)

// nl80211WpaVersions as declared in nl80211/nl80211.h:3661
type nl80211WpaVersions int32

// nl80211WpaVersions enumeration from nl80211/nl80211.h:3661
const (
	WpaVersion1 = 1 << 0
	WpaVersion2 = 1 << 1
)

// nl80211KeyDefaultTypes as declared in nl80211/nl80211.h:3675
type nl80211KeyDefaultTypes int32

// nl80211KeyDefaultTypes enumeration from nl80211/nl80211.h:3675
const (
	__KeyDefaultTypeInvalid = iota
	KeyDefaultTypeUnicast
	KeyDefaultTypeMulticast
	Num_KeyDefaultTypes
)

// nl80211KeyAttributes as declared in nl80211/nl80211.h:3705
type nl80211KeyAttributes int32

// nl80211KeyAttributes enumeration from nl80211/nl80211.h:3705
const (
	__KeyInvalid = iota
	KeyData
	KeyIdx
	KeyCipher
	KeySeq
	KeyDefault
	KeyDefaultMgmt
	KeyType
	KeyDefaultTypes
	__KeyAfterLast
	KeyMax = __KeyAfterLast - 1
)

// nl80211TxRateAttributes as declared in nl80211/nl80211.h:3736
type nl80211TxRateAttributes int32

// nl80211TxRateAttributes enumeration from nl80211/nl80211.h:3736
const (
	__TxrateInvalid = iota
	TxrateLegacy
	TxrateHt
	TxrateVht
	TxrateGi
	__TxrateAfterLast
	TxrateMax = __TxrateAfterLast - 1
)

// nl80211TxrateGi as declared in nl80211/nl80211.h:3759
type nl80211TxrateGi int32

// nl80211TxrateGi enumeration from nl80211/nl80211.h:3759
const (
	TxrateDefaultGi = iota
	TxrateForceSgi
	TxrateForceLgi
)

// nl80211Band as declared in nl80211/nl80211.h:3773
type nl80211Band int32

// nl80211Band enumeration from nl80211/nl80211.h:3773
const (
	Band2ghz = iota
	Band5ghz
	Band60ghz
	Num_Bands
)

// nl80211PsState as declared in nl80211/nl80211.h:3786
type nl80211PsState int32

// nl80211PsState enumeration from nl80211/nl80211.h:3786
const (
	PsDisabled = iota
	PsEnabled
)

// nl80211AttrCqm as declared in nl80211/nl80211.h:3819
type nl80211AttrCqm int32

// nl80211AttrCqm enumeration from nl80211/nl80211.h:3819
const (
	__AttrCqmInvalid = iota
	AttrCqmRssiThold
	AttrCqmRssiHyst
	AttrCqmRssiThresholdEvent
	AttrCqmPktLossEvent
	AttrCqmTxeRate
	AttrCqmTxePkts
	AttrCqmTxeIntvl
	AttrCqmBeaconLossEvent
	__AttrCqmAfterLast
	AttrCqmMax = __AttrCqmAfterLast - 1
)

// nl80211CqmRssiThresholdEvent as declared in nl80211/nl80211.h:3843
type nl80211CqmRssiThresholdEvent int32

// nl80211CqmRssiThresholdEvent enumeration from nl80211/nl80211.h:3843
const (
	CqmRssiThresholdEventLow = iota
	CqmRssiThresholdEventHigh
	CqmRssiBeaconLossEvent
)

// nl80211TxPowerSetting as declared in nl80211/nl80211.h:3856
type nl80211TxPowerSetting int32

// nl80211TxPowerSetting enumeration from nl80211/nl80211.h:3856
const (
	TxPowerAutomatic = iota
	TxPowerLimited
	TxPowerFixed
)

// nl80211PacketPatternAttr as declared in nl80211/nl80211.h:3883
type nl80211PacketPatternAttr int32

// nl80211PacketPatternAttr enumeration from nl80211/nl80211.h:3883
const (
	__PktpatInvalid = iota
	PktpatMask
	PktpatPattern
	PktpatOffset
	Num_Pktpat
	Max_Pktpat = Num_Pktpat - 1
)

// nl80211WowlanTriggers as declared in nl80211/nl80211.h:4011
type nl80211WowlanTriggers int32

// nl80211WowlanTriggers enumeration from nl80211/nl80211.h:4011
const (
	__WowlanTrigInvalid = iota
	WowlanTrigAny
	WowlanTrigDisconnect
	WowlanTrigMagicPkt
	WowlanTrigPktPattern
	WowlanTrigGtkRekeySupported
	WowlanTrigGtkRekeyFailure
	WowlanTrigEapIdentRequest
	WowlanTrig4wayHandshake
	WowlanTrigRfkillRelease
	WowlanTrigWakeupPkt80211
	WowlanTrigWakeupPkt80211Len
	WowlanTrigWakeupPkt8023
	WowlanTrigWakeupPkt8023Len
	WowlanTrigTcpConnection
	WowlanTrigWakeupTcpMatch
	WowlanTrigWakeupTcpConnlost
	WowlanTrigWakeupTcpNomoretokens
	WowlanTrigNetDetect
	WowlanTrigNetDetectResults
	Num_WowlanTrig
	Max_WowlanTrig = Num_WowlanTrig - 1
)

// nl80211WowlanTcpAttrs as declared in nl80211/nl80211.h:4129
type nl80211WowlanTcpAttrs int32

// nl80211WowlanTcpAttrs enumeration from nl80211/nl80211.h:4129
const (
	__WowlanTcpInvalid = iota
	WowlanTcpSrcIpv4
	WowlanTcpDstIpv4
	WowlanTcpDstMac
	WowlanTcpSrcPort
	WowlanTcpDstPort
	WowlanTcpDataPayload
	WowlanTcpDataPayloadSeq
	WowlanTcpDataPayloadToken
	WowlanTcpDataInterval
	WowlanTcpWakePayload
	WowlanTcpWakeMask
	Num_WowlanTcp
	Max_WowlanTcp = Num_WowlanTcp - 1
)

// nl80211AttrCoalesceRule as declared in nl80211/nl80211.h:4174
type nl80211AttrCoalesceRule int32

// nl80211AttrCoalesceRule enumeration from nl80211/nl80211.h:4174
const (
	__CoalesceRuleInvalid = iota
	AttrCoalesceRuleDelay
	AttrCoalesceRuleCondition
	AttrCoalesceRulePktPattern
	Num_AttrCoalesceRule
	AttrCoalesceRuleMax = Num_AttrCoalesceRule - 1
)

// nl80211CoalesceCondition as declared in nl80211/nl80211.h:4192
type nl80211CoalesceCondition int32

// nl80211CoalesceCondition enumeration from nl80211/nl80211.h:4192
const (
	CoalesceConditionMatch = iota
	CoalesceConditionNoMatch
)

// nl80211IfaceLimitAttrs as declared in nl80211/nl80211.h:4207
type nl80211IfaceLimitAttrs int32

// nl80211IfaceLimitAttrs enumeration from nl80211/nl80211.h:4207
const (
	IfaceLimitUnspec = iota
	IfaceLimitMax
	IfaceLimitTypes
	Num_IfaceLimit
	Max_IfaceLimit = Num_IfaceLimit - 1
)

// nl80211IfCombinationAttrs as declared in nl80211/nl80211.h:4263
type nl80211IfCombinationAttrs int32

// nl80211IfCombinationAttrs enumeration from nl80211/nl80211.h:4263
const (
	IfaceCombUnspec = iota
	IfaceCombLimits
	IfaceCombMaxnum
	IfaceCombStaApBiMatch
	IfaceCombNumChannels
	IfaceCombRadarDetectWidths
	IfaceCombRadarDetectRegions
	Num_IfaceComb
	Max_IfaceComb = Num_IfaceComb - 1
)

// nl80211PlinkState as declared in nl80211/nl80211.h:4296
type nl80211PlinkState int32

// nl80211PlinkState enumeration from nl80211/nl80211.h:4296
const (
	PlinkListen = iota
	PlinkOpnSnt
	PlinkOpnRcvd
	PlinkCnfRcvd
	PlinkEstab
	PlinkHolding
	PlinkBlocked
	Num_PlinkStates
	Max_PlinkStates = Num_PlinkStates - 1
)

// plinkActions as declared in nl80211/nl80211.h:4318
type plinkActions int32

// plinkActions enumeration from nl80211/nl80211.h:4318
const (
	PlinkActionNoAction = iota
	PlinkActionOpen
	PlinkActionBlock
	Num_PlinkActions
)

// nl80211RekeyData as declared in nl80211/nl80211.h:4340
type nl80211RekeyData int32

// nl80211RekeyData enumeration from nl80211/nl80211.h:4340
const (
	__RekeyDataInvalid = iota
	RekeyDataKek
	RekeyDataKck
	RekeyDataReplayCtr
	Num_RekeyData
	Max_RekeyData = Num_RekeyData - 1
)

// nl80211HiddenSsid as declared in nl80211/nl80211.h:4360
type nl80211HiddenSsid int32

// nl80211HiddenSsid enumeration from nl80211/nl80211.h:4360
const (
	HiddenSsidNotInUse = iota
	HiddenSsidZeroLen
	HiddenSsidZeroContents
)

// nl80211StaWmeAttr as declared in nl80211/nl80211.h:4376
type nl80211StaWmeAttr int32

// nl80211StaWmeAttr enumeration from nl80211/nl80211.h:4376
const (
	__StaWmeInvalid = iota
	StaWmeUapsdQueues
	StaWmeMaxSp
	__StaWmeAfterLast
	StaWmeMax = __StaWmeAfterLast - 1
)

// nl80211PmksaCandidateAttr as declared in nl80211/nl80211.h:4398
type nl80211PmksaCandidateAttr int32

// nl80211PmksaCandidateAttr enumeration from nl80211/nl80211.h:4398
const (
	__PmksaCandidateInvalid = iota
	PmksaCandidateIndex
	PmksaCandidateBssid
	PmksaCandidatePreauth
	Num_PmksaCandidate
	Max_PmksaCandidate = Num_PmksaCandidate - 1
)

// nl80211TdlsOperation as declared in nl80211/nl80211.h:4417
type nl80211TdlsOperation int32

// nl80211TdlsOperation enumeration from nl80211/nl80211.h:4417
const (
	TdlsDiscoveryReq = iota
	TdlsSetup
	TdlsTeardown
	TdlsEnableLink
	TdlsDisableLink
)

// nl80211FeatureFlags as declared in nl80211/nl80211.h:4526
type nl80211FeatureFlags int32

// nl80211FeatureFlags enumeration from nl80211/nl80211.h:4526
const (
	FeatureSkTxStatus             = 1 << 0
	FeatureHtIbss                 = 1 << 1
	FeatureInactivityTimer        = 1 << 2
	FeatureCellBaseRegHints       = 1 << 3
	FeatureP2pDeviceNeedsChannel  = 1 << 4
	FeatureSae                    = 1 << 5
	FeatureLowPriorityScan        = 1 << 6
	FeatureScanFlush              = 1 << 7
	FeatureApScan                 = 1 << 8
	FeatureVifTxpower             = 1 << 9
	FeatureNeedObssScan           = 1 << 10
	FeatureP2pGoCtwin             = 1 << 11
	FeatureP2pGoOppps             = 1 << 12
	FeatureAdvertiseChanLimits    = 1 << 14
	FeatureFullApClientState      = 1 << 15
	FeatureUserspaceMpm           = 1 << 16
	FeatureActiveMonitor          = 1 << 17
	FeatureApModeChanWidthChange  = 1 << 18
	FeatureDsParamSetIeInProbes   = 1 << 19
	FeatureWfaTpcIeInProbes       = 1 << 20
	FeatureQuiet                  = 1 << 21
	FeatureTxPowerInsertion       = 1 << 22
	FeatureAcktoEstimation        = 1 << 23
	FeatureStaticSmps             = 1 << 24
	FeatureDynamicSmps            = 1 << 25
	FeatureSupportsWmmAdmission   = 1 << 26
	FeatureMacOnCreate            = 1 << 27
	FeatureTdlsChannelSwitch      = 1 << 28
	FeatureScanRandomMacAddr      = 1 << 29
	FeatureSchedScanRandomMacAddr = 1 << 30
	FeatureNdRandomMacAddr        = 1 << 31
)

// nl80211ExtFeatureIndex as declared in nl80211/nl80211.h:4595
type nl80211ExtFeatureIndex int32

// nl80211ExtFeatureIndex enumeration from nl80211/nl80211.h:4595
const (
	ExtFeatureVhtIbss = iota
	ExtFeatureRrm
	ExtFeatureMuMimoAirSniffer
	ExtFeatureScanStartTime
	ExtFeatureBssParentTsf
	ExtFeatureSetScanDwell
	ExtFeatureBeaconRateLegacy
	ExtFeatureBeaconRateHt
	ExtFeatureBeaconRateVht
	Num_ExtFeatures
	Max_ExtFeatures = Num_ExtFeatures - 1
)

// nl80211ProbeRespOffloadSupportAttr as declared in nl80211/nl80211.h:4625
type nl80211ProbeRespOffloadSupportAttr int32

// nl80211ProbeRespOffloadSupportAttr enumeration from nl80211/nl80211.h:4625
const (
	ProbeRespOffloadSupportWps    = 1 << 0
	ProbeRespOffloadSupportWps2   = 1 << 1
	ProbeRespOffloadSupportP2p    = 1 << 2
	ProbeRespOffloadSupport80211u = 1 << 3
)

// nl80211ConnectFailedReason as declared in nl80211/nl80211.h:4638
type nl80211ConnectFailedReason int32

// nl80211ConnectFailedReason enumeration from nl80211/nl80211.h:4638
const (
	ConnFailMaxClients = iota
	ConnFailBlockedClient
)

// nl80211ScanFlags as declared in nl80211/nl80211.h:4667
type nl80211ScanFlags int32

// nl80211ScanFlags enumeration from nl80211/nl80211.h:4667
const (
	ScanFlagLowPriority = 1 << 0
	ScanFlagFlush       = 1 << 1
	ScanFlagAp          = 1 << 2
	ScanFlagRandomAddr  = 1 << 3
)

// nl80211AclPolicy as declared in nl80211/nl80211.h:4687
type nl80211AclPolicy int32

// nl80211AclPolicy enumeration from nl80211/nl80211.h:4687
const (
	AclPolicyAcceptUnlessListed = iota
	AclPolicyDenyUnlessListed
)

// nl80211SmpsMode as declared in nl80211/nl80211.h:4702
type nl80211SmpsMode int32

// nl80211SmpsMode enumeration from nl80211/nl80211.h:4702
const (
	SmpsOff = iota
	SmpsStatic
	SmpsDynamic
	__SmpsAfterLast
	SmpsMax = __SmpsAfterLast - 1
)

// nl80211RadarEvent as declared in nl80211/nl80211.h:4726
type nl80211RadarEvent int32

// nl80211RadarEvent enumeration from nl80211/nl80211.h:4726
const (
	RadarDetected = iota
	RadarCacFinished
	RadarCacAborted
	RadarNopFinished
)

// nl80211DfsState as declared in nl80211/nl80211.h:4744
type nl80211DfsState int32

// nl80211DfsState enumeration from nl80211/nl80211.h:4744
const (
	DfsUsable = iota
	DfsUnavailable
	DfsAvailable
)

// nl80211ProtocolFeatures as declared in nl80211/nl80211.h:4758
type nl80211ProtocolFeatures int32

// nl80211ProtocolFeatures enumeration from nl80211/nl80211.h:4758
const (
	ProtocolFeatureSplitWiphyDump = 1 << 0
)

// nl80211CritProtoId as declared in nl80211/nl80211.h:4771
type nl80211CritProtoId int32

// nl80211CritProtoId enumeration from nl80211/nl80211.h:4771
const (
	CritProtoUnspec = iota
	CritProtoDhcp
	CritProtoEapol
	CritProtoApipa
	Num_CritProto
)

// nl80211RxmgmtFlags as declared in nl80211/nl80211.h:4790
type nl80211RxmgmtFlags int32

// nl80211RxmgmtFlags enumeration from nl80211/nl80211.h:4790
const (
	RxmgmtFlagAnswered = 1 << 0
)

// nl80211TdlsPeerCapability as declared in nl80211/nl80211.h:4824
type nl80211TdlsPeerCapability int32

// nl80211TdlsPeerCapability enumeration from nl80211/nl80211.h:4824
const (
	TdlsPeerHt  = 1 << 0
	TdlsPeerVht = 1 << 1
	TdlsPeerWmm = 1 << 2
)

// nl80211SchedScanPlan as declared in nl80211/nl80211.h:4843
type nl80211SchedScanPlan int32

// nl80211SchedScanPlan enumeration from nl80211/nl80211.h:4843
const (
	__SchedScanPlanInvalid = iota
	SchedScanPlanInterval
	SchedScanPlanIterations
	__SchedScanPlanAfterLast
	SchedScanPlanMax = __SchedScanPlanAfterLast - 1
)

// nl80211BssSelectAttr as declared in nl80211/nl80211.h:4887
type nl80211BssSelectAttr int32

// nl80211BssSelectAttr enumeration from nl80211/nl80211.h:4887
const (
	__BssSelectAttrInvalid = iota
	BssSelectAttrRssi
	BssSelectAttrBandPref
	BssSelectAttrRssiAdjust
	__BssSelectAttrAfterLast
	BssSelectAttrMax = __BssSelectAttrAfterLast - 1
)

// nl80211NanDualBandConf as declared in nl80211/nl80211.h:4907
type nl80211NanDualBandConf int32

// nl80211NanDualBandConf enumeration from nl80211/nl80211.h:4907
const (
	NanBandDefault = 1 << 0
	NanBand2ghz    = 1 << 1
	NanBand5ghz    = 1 << 2
)

// nl80211NanFunctionType as declared in nl80211/nl80211.h:4922
type nl80211NanFunctionType int32

// nl80211NanFunctionType enumeration from nl80211/nl80211.h:4922
const (
	NanFuncPublish = iota
	NanFuncSubscribe
	NanFuncFollowUp
	__NanFuncTypeAfterLast
	NanFuncMaxType = __NanFuncTypeAfterLast - 1
)

// nl80211NanPublishType as declared in nl80211/nl80211.h:4940
type nl80211NanPublishType int32

// nl80211NanPublishType enumeration from nl80211/nl80211.h:4940
const (
	NanSolicitedPublish   = 1 << 0
	NanUnsolicitedPublish = 1 << 1
)

// nl80211NanFuncTermReason as declared in nl80211/nl80211.h:4954
type nl80211NanFuncTermReason int32

// nl80211NanFuncTermReason enumeration from nl80211/nl80211.h:4954
const (
	NanFuncTermReasonUserRequest = iota
	NanFuncTermReasonTtlExpired
	NanFuncTermReasonError
)

// nl80211NanFuncAttributes as declared in nl80211/nl80211.h:5006
type nl80211NanFuncAttributes int32

// nl80211NanFuncAttributes enumeration from nl80211/nl80211.h:5006
const (
	__NanFuncInvalid = iota
	NanFuncType
	NanFuncServiceId
	NanFuncPublishType
	NanFuncPublishBcast
	NanFuncSubscribeActive
	NanFuncFollowUpId
	NanFuncFollowUpReqId
	NanFuncFollowUpDest
	NanFuncCloseRange
	NanFuncTtl
	NanFuncServiceInfo
	NanFuncSrf
	NanFuncRxMatchFilter
	NanFuncTxMatchFilter
	NanFuncInstanceId
	NanFuncTermReason
	Num_NanFuncAttr
	NanFuncAttrMax = Num_NanFuncAttr - 1
)

// nl80211NanSrfAttributes as declared in nl80211/nl80211.h:5045
type nl80211NanSrfAttributes int32

// nl80211NanSrfAttributes enumeration from nl80211/nl80211.h:5045
const (
	__NanSrfInvalid = iota
	NanSrfInclude
	NanSrfBf
	NanSrfBfIdx
	NanSrfMacAddrs
	Num_NanSrfAttr
	NanSrfAttrMax = Num_NanSrfAttr - 1
)

// nl80211NanMatchAttributes as declared in nl80211/nl80211.h:5070
type nl80211NanMatchAttributes int32

// nl80211NanMatchAttributes enumeration from nl80211/nl80211.h:5070
const (
	__NanMatchInvalid = iota
	NanMatchFuncLocal
	NanMatchFuncPeer
	Num_NanMatchAttr
	NanMatchAttrMax = Num_NanMatchAttr - 1
)
