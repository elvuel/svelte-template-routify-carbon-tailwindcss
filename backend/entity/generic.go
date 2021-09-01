package entity

import (
	"time"
)

type Organization struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	UniqueID  string     `json:"unique_id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	ID             uint   `json:"id" gorm:"primarykey"`
	UniqueID       string `json:"unique_id"`
	OrganizationID uint   `json:"organization_id"`

	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Description  string `json:"description"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Role struct {
	ID             uint   `json:"id" gorm:"primarykey"`
	UniqueID       string `json:"unique_id"`
	OrganizationID uint   `json:"organization_id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserRole struct {
	ID             uint   `json:"id" gorm:"primarykey"`
	UniqueID       string `json:"unique_id"`
	OrganizationID uint   `json:"organization_id"`

	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Permission struct {
	ID             uint   `json:"id" gorm:"primarykey"`
	UniqueID       string `json:"unique_id"`
	OrganizationID uint   `json:"organization_id"`

	Resource string `json:"resource"`
	Action   string `json:"action"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type RolePermission struct {
	ID             uint   `json:"id" gorm:"primarykey"`
	UniqueID       string `json:"unique_id"`
	OrganizationID uint   `json:"organization_id"`

	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Menu struct {
	ID             uint   `json:"id" gorm:"primarykey"`
	UniqueID       string `json:"unique_id"`
	OrganizationID uint   `json:"organization_id"`

	ParentID    uint   `json:"parent_id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	LocaleKey   string `json:"locale_key"`
	Position    int    `json:"position"`
	Icon        string `json:"icon"`
	Description string `json:"description"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CarbonIcon struct {
	ID uint `json:"id" gorm:"primarykey"`

	Name     string `json:"name"`
	Catalog  string `json:"catalog"`
	Icon     string `json:"icon"`
	Size     string `json:"size"` // 16, 20, 24, 32
	Disabled bool   `json:"disabled"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	// {"Alphanumeric":["LetterAa","LetterBb","LetterCc","LetterDd","LetterEe","LetterFf","LetterGg","LetterHh","LetterIi","LetterJj","LetterKk","LetterLl","LetterMm","LetterNn","LetterOo","LetterPp","LetterQq","LetterRr","LetterSs","LetterTt","LetterUu","LetterVv","LetterWw","LetterXx","LetterYy","LetterZz","Number_0","Number_1","Number_2","Number_3","Number_4","Number_5","Number_6","Number_7","Number_8","Number_9","NumberSmall_0","NumberSmall_1","NumberSmall_2","NumberSmall_3","NumberSmall_4","NumberSmall_5","NumberSmall_6","NumberSmall_7","NumberSmall_8","NumberSmall_9","Omega","Sigma"]}
	// {"App Catalogue":["CloudAlerting","CloudAuditing","CloudDataOps","CloudLogging","CloudMonitoring","CloudServiceManagement","IbmCloudFoundry_1","CloudRegistry","CloudSatellite","CloudServices","ContentDeliveryNetwork","DirectLink","DnsServices","FlowLogsVpc","IbmCloudDedicatedHost","IbmCloudInternetServices","IbmCloudSubnets","IbmCloudVpcEndpoints","IbmSecurityServices","InfrastructureClassic","LoadBalancerVpc","VirtualPrivateCloudAlt","VlanIbm"]}
	// {"Commerce":["Account","Apple","Basketball","Bat","BeeBat","Building","Corn","Currency","CurrencyBaht","CurrencyDollar","CurrencyEuro","CurrencyLira","CurrencyPound","CurrencyRuble","CurrencyRupee","CurrencyShekel","CurrencyWon","CurrencyYen","Delivery","DeliveryAdd","DeliveryParcel","Enterprise","Finance","Fish","FishMultiple","Fragile","FruitBowl","Gift","Industry","InventoryManagement","KeepDry","Money","Monster","NoodleBowl","PiggyBank","PiggyBankSlot","Purchase","Receipt","ShoppingBag","ShoppingCart","ShoppingCartArrowDown","ShoppingCartArrowUp","ShoppingCartClear","ShoppingCartError","ShoppingCartMinus","ShoppingCartPlus","ShoppingCatalog","Soccer","Sprout","StackLimitation","Store","Strawberry","Tennis","TennisBall","ThisSideUp","Tree","TwoPersonLift","Wallet","Wheat","WirelessCheckout"]}
	// {"Controls":["CenterToFit","Dashboard","DashboardReference","FitToScreen","Forward_10","Forward_30","Forward_5","NextFilled","NextOutline","OpenPanelBottom","OpenPanelFilledBottom","OpenPanelFilledLeft","OpenPanelFilledRight","OpenPanelFilledTop","OpenPanelLeft","OpenPanelRight","OpenPanelTop","Pause","PauseFilled","PauseOutline","PauseOutlineFilled","Play","PlayFilled","PlayFilledAlt","PlayOutline","PlayOutlineFilled","Power","PreviousFilled","PreviousOutline","Recording","RecordingFilled","RecordingFilledAlt","Rewind_10","Rewind_30","Rewind_5","Settings","SettingsAdjust","SettingsCheck","ShrinkScreen","ShrinkScreenFilled","SkipBack","SkipBackFilled","SkipBackOutline","SkipBackOutlineFilled","SkipBackOutlineSolid","SkipBackSolidFilled","SkipForward","SkipForwardFilled","SkipForwardOutline","SkipForwardOutlineFilled","SkipForwardOutlineSolid","SkipForwardSolidFilled","Stop","StopFilled","StopFilledAlt","StopOutline","StopOutlineFilled","VideoAdd","VideoChat","VolumeDown","VolumeDownAlt","VolumeDownFilled","VolumeDownFilledAlt","VolumeMute","VolumeMuteFilled","VolumeUp","VolumeUpAlt","VolumeUpFilled","VolumeUpFilledAlt"]}
	// {"Data":["Activity","Analytics","AnalyticsReference","AppConnectivity","Archive","Blockchain","Boolean","BoxPlot","Calculation","CalculationAlt","CalendarHeatMap","CharacterPatterns","Chart_3D","ChartArea","ChartAreaSmooth","ChartAreaStepper","ChartAverage","ChartBar","ChartBarFloating","ChartBarOverlay","ChartBarStacked","ChartBarTarget","ChartBubble","ChartBubblePacked","ChartBullet","ChartCandlestick","ChartClusterBar","ChartColumn","ChartColumnFloating","ChartColumnTarget","ChartCombo","ChartComboStacked","ChartCustom","ChartErrorBar","ChartErrorBarAlt","ChartEvaluation","ChartHighLow","ChartHistogram","ChartLine","ChartLineData","ChartLineSmooth","ChartMarimekko","ChartMaximum","ChartMedian","ChartMinimum","ChartMultiLine","ChartMultitype","ChartNetwork","ChartParallel","ChartPie","ChartPoint","ChartPopulation","ChartRadar","ChartRadial","ChartRelationship","ChartRing","ChartRiver","ChartRose","ChartScatter","ChartSpiral","ChartStacked","ChartStepper","ChartSunburst","ChartTSne","ChartTreemap","ChartVennDiagram","ChartViolinPlot","ChartWaterfall","ChartWinLoss","ChartChoroplethMap","CircleFilled","CirclePacking","ClassifierLanguage","ColumnDependency","Concept","ContainerSoftware","Covariate","CrossTab","Data_1","Data_2","DataBase","DataBaseAlt","DataCheck","DataConnected","DataError","DataFormat","DataReference","DataSet","DataStructured","DataUnstructured","DataView","DataBin","DataClass","DataCollection","DataPlayer","DataRefinery","DataRefineryReference","DataShare","DataTable","DataTableReference","DataVisualization_1","DataVisualization_2","DataVisualization_3","DataVisualization_4","DataStore","DecisionTree","Diagram","DiagramReference","DocumentSentiment","DoubleInteger","DriverAnalysis","Factor","Flow","FlowConnection","FlowData","FlowModeler","FlowStream","FlowStreamReference","GraphicalDataFlow","Growth","Hashtag","HeatMap","HeatMap_02","HeatMap_03","HeatMapStocks","HybridNetworking","ImproveRelevance","IncreaseLevel","Interactions","JoinFull","JoinInner","JoinLeft","JoinOuter","JoinRight","Linux","LinuxAlt","ManagedSolutions","MathCurve","ModelBuilder","ModelBuilderReference","NameSpace","Network_1","Network_2","Network_3","Network_3Reference","Network_4","Nominal","ObjectStorage","Ordinal","Parameter","ParentChild","PhraseSentiment","PresentationFile","ProgressBar","ProgressBarRound","QqPlot","QuadrantPlot","QueryQueue","Report","ReportData","RowCollapse","RowExpand","SankeyDiagram","SankeyDiagramAlt","ScatterMatrix","Schematics","ShareKnowledge","SkillLevel","SkillLevelAdvanced","SkillLevelBasic","SkillLevelIntermediate","StemLeafPlot","StoragePool","StorageRequest","StringInteger","StringText","SummaryKpi","SysProvision","Table","TableBuilt","TableSplit","Term","TextLink","TextLinkAnalysis","TextMining","TextMiningApplier","TimePlot","TreeView","TreeViewAlt","TypePattern","Types","VirtualColumn","VirtualColumnKey","VirtualMachine","VirtualPrivateCloud","VisualRecognition","VmdkDisk","WordCloud"]}
	// {"File":["4K","4KFilled","Api","Api_1","App","Application","Badge","Book","Cad","Catalog","Categories","Category","CategoryAdd","CategoryAnd","CategoryNew","CategoryNewEach","Cda","Certificate","CertificateCheck","Classification","ClosedCaption","ClosedCaptionAlt","ClosedCaptionFilled","IbmCloudFoundry_2","ContentView","Copy","CopyFile","CopyLink","Course","Csv","Cube","CubeView","Doc","Document","DocumentAdd","DocumentAttachment","DocumentAudio","DocumentBlank","DocumentDownload","DocumentExport","DocumentImport","DocumentPdf","DocumentProtected","DocumentSecurity","DocumentSigned","DocumentSketch","DocumentSubtract","DocumentTasks","DocumentUnknown","DocumentUnprotected","DocumentVideo","DocumentView","DocumentWordProcessor","DocumentWordProcessorReference","Download","Dvr","Export","Folder","FolderAdd","FolderDetails","FolderDetailsReference","FolderMoveTo","FolderOff","FolderOpen","FolderParent","FolderShared","Folders","Gamification","GeneratePdf","Gif","Hd","HdFilled","Hdr","Html","HtmlReference","Http","Iso","IsoFilled","IsoOutline","Jpg","Json","JsonReference","Legend","License","LicenseDraft","LicenseGlobal","LicenseMaintenance","LicenseMaintenanceDraft","LicenseThirdParty","LicenseThirdPartyDraft","ListDropdown","MediaLibrary","Model","ModelAlt","ModelReference","Mov","Mp3","Mp4","Mpeg","Mpg2","Music","MusicAdd","MusicRemove","NoTicket","NonCertified","Notebook","NotebookReference","PanelExpansion","Paste","Pdf","PdfReference","Playlist","Png","Policy","PopUp","Ppt","Product","Raw","Result","ResultNew","ResultOld","Roadmap","Rule","RuleCancelled","RuleDraft","RuleFilled","RuleTest","Scale","Script","ScriptReference","Sdk","Security","Sql","StarReview","Svg","TableOfContents","Task","TaskView","Template","Ticket","Tif","Tsv","Txt","TxtReference","Upload","Version","Websheet","Wmv","Workspace","Xls","Xml","Zip","ZipReference"]}
	// {"Formatting":["AlignHorizontalCenter","AlignHorizontalLeft","AlignHorizontalRight","AlignVerticalBottom","AlignVerticalCenter","AlignVerticalTop","AlignBoxBottomCenter","AlignBoxBottomLeft","AlignBoxBottomRight","AlignBoxMiddleCenter","AlignBoxMiddleLeft","AlignBoxMiddleRight","AlignBoxTopCenter","AlignBoxTopLeft","AlignBoxTopRight","Attachment","BorderBottom","BorderFull","BorderLeft","BorderNone","BorderRight","BorderTop","BrightnessContrast","BringToFront","ColorPalette","ColorSwitch","ColumnDelete","ColumnInsert","Contrast","Corner","Crop","Cut","CutOut","DistributeHorizontalCenter","DistributeHorizontalLeft","DistributeHorizontalRight","DistributeVerticalBottom","DistributeVerticalCenter","DistributeVerticalTop","DocumentHorizontal","DocumentVertical","DotMark","Draw","DropPhoto","DropPhotoFilled","Edit","EditOff","Erase","Eyedropper","Gradient","Image","ImageCopy","ImageReference","Language","Lasso","LassoPolygon","Link","ListBoxes","ListBulleted","ListChecked","ListNumbered","MagicWand","MagicWandFilled","NoImage","Opacity","PageBreak","PageNumber","PaintBrush","PaintBrushAlt","Paragraph","Pen","PenFountain","Percentage","PercentageFilled","Quotes","ReflectHorizontal","ReflectVertical","RotateClockwise","RotateClockwiseAlt","RotateClockwiseAltFilled","RotateClockwiseFilled","RotateCounterclockwise","RotateCounterclockwiseAlt","RotateCounterclockwiseAltFilled","RotateCounterclockwiseFilled","RowDelete","RowInsert","Save","Select_01","Select_02","SendToBack","ShapeExcept","ShapeExclude","ShapeIntersect","ShapeJoin","ShapeUnite","SpellCheck","SprayPaint","TextAlignCenter","TextAlignJustify","TextAlignLeft","TextAlignMixed","TextAlignRight","TextAllCaps","TextBold","TextClearFormat","TextColor","TextCreation","TextFill","TextFont","TextFootnote","TextHighlight","TextIndent","TextIndentLess","TextIndentMore","TextItalic","TextKerning","TextLeading","TextLineSpacing","TextNewLine","TextScale","TextSelection","TextSmallCaps","TextStrikethrough","TextSubscript","TextSuperscript","TextTracking","TextUnderline","TextVerticalAlignment","TextWrap","TrashCan","Unlink"]}
	// {"Health":["3DCursor","3DCursorAlt","3DCurveAutoColon","3DCurveAutoVessels","3DCurveManual","3DIca","3DMprToggle","3DPrintMesh","3DSoftware","ThirdPartyConnect","AiResults","AiResultsHigh","AiResultsLow","AiResultsMedium","AiResultsUrgent","AiResultsVeryHigh","AiStatus","AiStatusComplete","AiStatusFailed","AiStatusInProgress","AiStatusQueued","AiStatusRejected","Angle","AnnotationVisibility","ArrowAnnotation","AutoScroll","BrushFreehand","BrushPolygon","CdArchive","CdCreateArchive","CdCreateExchange","Chemistry","ChemistryReference","CircleMeasurement","CobbAngle","ContourFinding","Coronavirus","CrossReference","CutInHalf","Denominate","Dicom_6000","DicomOverlay","Dna","DownloadStudy","EdgeEnhancement","EdgeEnhancement_01","EdgeEnhancement_02","EdgeEnhancement_03","Erase_3D","ExamMode","FusionBlender","HangingProtocol","HealthCross","HoleFilling","HoleFillingCursor","Ica_2D","ImageMedical","InteractiveSegmentationCursor","LaunchStudy_1","LaunchStudy_2","LaunchStudy_3","Magnify","Mammogram","Medication","MedicationAlert","MedicationReminder","Nominate","PageScroll","PetImageB","PetImageO","Pills","PillsAdd","PillsSubtract","PointerText","RegionAnalysisArea","RegionAnalysisVolume","Registration","Rotate_180","Rotate_360","SaveAnnotation","SaveImage","SaveSeries","ScalpelCursor","ScalpelLasso","ScalpelSelect","Smoothing","SmoothingCursor","SpineLabel","SplitDiscard","StackedMove","StackedScrolling_1","StackedScrolling_2","StressBreathEditor","StudyNext","StudyPrevious","StudySkip","SubVolume","TextAnnotationToggle","Threshold","ThumbnailPreview","WindowAuto","WindowBase","WindowOverlay","WindowPreset","ZoomPan"]}
	// {"IBM":["Bee","Carbon","CloudApp","EdtLoop","IbmCloud","IbmSecurity","Watson","IbmWatsonMachineLearning"]}
	// {"Instructments":["Binoculars","Box","Bullhorn","Calibrate","Cursor_1","Cursor_2","Meter","MeterAlt","Microscope","Portfolio","Radar","Ruler","RulerAlt","Scales","ScalesTipped","Scalpel","Stamp","Stethoscope","ToolBox","Tools","ToolsAlt","Umbrella"]}
	// {"Navigation":["Add","AddAlt","AddFilled","Apps","ArrowDown","ArrowDownLeft","ArrowDownRight","ArrowLeft","ArrowRight","ArrowUp","ArrowUpLeft","ArrowUpRight","CaretDown","CaretLeft","CaretRight","CaretUp","ChevronDown","ChevronLeft","ChevronRight","ChevronUp","Close","CloseFilled","CloseOutline","DownToBottom","Draggable","Home","Menu","OverflowMenuHorizontal","OverflowMenuVertical","PageFirst","PageLast","SelectWindow","Subtract","SubtractAlt","Switcher","UpToTop","ZoomArea","ZoomFit","ZoomIn","ZoomInArea","ZoomOut","ZoomOutArea","ZoomReset"]}
	// {"Operations":["AddComment","Area","AreaCustom","ArrowShiftDown","ArrowsHorizontal","ArrowsVertical","Automatic","Bookmark","BookmarkAdd","BookmarkFilled","Branch","CaretSort","CaretSortDown","CaretSortUp","CenterCircle","CenterSquare","ChangeCatalog","Chat","ChatLaunch","ChevronMini","ChevronSort","ChevronSortDown","ChevronSortUp","Choices","CircleDash","Clean","CollapseAll","CollapseCategories","Commit","Compare","Connect","ConnectRecursive","ConnectSource","ConnectTarget","ConvertToCloud","Deploy","DeployRules","DragHorizontal","DragVertical","DrillBack","DrillDown","DrillThrough","Exit","ExpandAll","ExpandCategories","Explore","Fade","Filter","FilterEdit","FilterRemove","FilterReset","Flag","FlagFilled","Fork","FunctionMath","ImageSearch","Insert","InsertPage","InsertSyntax","Intersect","JumpLink","Launch","Layers","Login","Logout","Loop","MacCommand","MacOption","MacShift","ManageProtection","Maximize","Migrate","MigrateAlt","Minimize","Move","NewTab","NotSent","NotSentFilled","OperationsField","OperationsRecord","Overlay","Package","PanHorizontal","PanVertical","PauseFuture","PausePast","Pin","PinFilled","Query","RecentlyViewed","Recommend","Redo","Renew","Repeat","RepeatOne","Replicate","Reply","ReplyAll","RequestQuote","Reset","ResetAlt","Restart","Rotate","Run","SaveModel","Scan","ScanDisabled","Search","SearchLocate","Send","SendAlt","SendAltFilled","SendFilled","Shuffle","SortAscending","SortDescending","Split","SplitScreen","Tag","TagEdit","TagGroup","Translate","Transpose","Undo","Upgrade","ViewNext","WorkspaceImport","XAxis","YAxis","ZAxis"]}
	// {"Research":["Barrier","BlochSphere","CcX","CircuitComposer","ComposerEdit","CU1","CU3","CY","CZ","H","HintonPlot","Id","Matrix","Operation","OperationGauge","OperationIf","S","SAlt","T","TAlt","U1","U2","U3","X","Y","Z"]}
	// {"Senses":["Cognitive","Hearing","Idea","Movement","Smell","Taste","Touch_1","Touch_1Filled","Touch_1Down","Touch_1DownFilled","Touch_2","Touch_2Filled","TouchInteraction"]}
	// {"Social":["LogoDigg","LogoDiscord","LogoFacebook","LogoFlickr","LogoGitHub","LogoGlassdoor","LogoInstagram","LogoLinkedIn","LogoLivestream","LogoMedium","LogoPinterest","LogoQuora","LogoSkype","LogoSlack","LogoSnapchat","LogoTumblr","LogoTwitter","LogoWechat","LogoXing","LogoYelp","LogoYouTube"]}
	// {"Status":["Checkmark","CheckmarkFilled","CheckmarkFilledError","CheckmarkFilledWarning","CheckmarkOutline","CheckmarkOutlineError","CheckmarkOutlineWarning","ConditionPoint","ConditionWaitPoint","CriticalSeverity","Error","ErrorFilled","ErrorOutline","Caution","CautionInverted","CircleFill","CircleStroke","Critical","SquareFill","Help","HelpFilled","InProgress","InProgressError","InProgressWarning","Incomplete","IncompleteCancel","IncompleteError","IncompleteWarning","Information","InformationFilled","InformationSquare","InformationSquareFilled","LowSeverity","Misuse","MisuseOutline","Pending","PendingFilled","Queued","Undefined","UndefinedFilled","Unknown","UnknownFilled","Warning","WarningAlt","WarningAltFilled","WarningAltInverted","WarningAltInvertedFilled","WarningFilled","WarningHex","WarningHexFilled","WarningOther","WarningSquare","WarningSquareFilled"]}
	// {"Systems":["ApplicationMobile","ApplicationVirtual","ApplicationWeb","Autoscaling","BareMetalServer","BastionHost","BlockStorage","BlockStorageAlt","ChatOperational","CodeSigningService","ContainerRegistry","ContainerServices","DataCenter","DataAccessor","DataBackup","DataBlob","DataDiode","EdgeNodeAlt","FileStorage","Firewall","FirewallClassic","FloatingIp","Gateway","GatewayApi","GatewayMail","GatewayPublic","GatewaySecurity","GatewayUserAccess","GatewayVpn","GroupAccess","GroupAccount","GroupResource","GroupSecurity","Gui","GuiManagement","HardwareSecurityModule","HybridNetworkingAlt","IdManagement","ImageService","InstanceBx","InstanceClassic","InstanceCx","InstanceMx","InstanceVirtual","IntrusionPrevention","Kubernetes","LoadBalancerApplication","LoadBalancerClassic","LoadBalancerGlobal","LoadBalancerListener","LoadBalancerLocal","LoadBalancerNetwork","LoadBalancerPool","MessageQueue","MobilityServices","NetworkEnterprise","NetworkOverlay","NetworkPublic","ObjectStorageAlt","PointOfPresence","Router","RouterVoice","RouterWifi","SecurityServices","ServerDns","ServerProxy","ServerTime","Slisor","SubnetAclRules","SwitchLayer_2","SwitchLayer_3","TwoFactorAuthentication","VehicleApi","VehicleConnected","VehicleInsights","VehicleServices","VirtualDesktop","Vlan","VpnConnection","VpnPolicy","WifiNotSecure","WifiSecure","WifiBridge","WifiBridgeAlt"]}
	// {"Technology":["Aperture","Asset","At","AudioConsole","AugmentedReality","Barcode","BatteryCharging","BatteryEmpty","BatteryFull","BatteryHalf","BatteryLow","BatteryQuarter","Blog","Bluetooth","BluetoothOff","Bot","BreakingChange","BuildingInsights_1","BuildingInsights_2","BuildingInsights_3","Calculator","CalculatorCheck","Camera","CameraAction","CellTower","ChatBot","Chip","CloudDownload","CloudUpload","Code","CodeHide","CodeReference","ConnectionReceive","ConnectionSend","ConnectionTwoWay","ConnectionSignal","ConnectionSignalOff","Debug","DeploymentPattern","DeploymentPolicy","DeskAdjustable","Development","Devices","Drone","DroneDelivery","DroneFront","DroneVideo","EdgeCluster","EdgeDevice","EdgeNode","EdgeService","Email","EmailNew","Equalizer","FetchUpload","FetchUploadCloud","FingerprintRecognition","Forum","Function","GameConsole","GameWireless","Headphones","Headset","Integration","IoTConnect","IoTPlatform","Keyboard","Laptop","LocationCurrent","LogoJupyter","LogoKeybase","LogoOpenShift","LogoPython","LogoRScript","LogoVMware","MachineLearning","MachineLearningModel","MailAll","MailReply","Mobile","MobileAdd","MobileAudio","MobileCheck","MobileDownload","MobileLandscape","Outage","Password","PhoneIp","Plug","PlugFilled","Printer","QrCode","Radio","RadioCombat","RadioPushToTalk","Rocket","Rss","Satellite","Share","SignalStrength","SimCard","Tablet","TabletLandscape","Terminal","TransmissionLte","Usb","Voicemail","Vpn","Watch","WiFi","WifiController","WiFiOff","Wikis"]}
	// {"Time":["Alarm","AlarmAdd","AlarmSubtract","Calendar","CalendarSettings","CalendarTools","Event","EventSchedule","Hourglass","Reminder","ReminderMedical","Snooze","Time","Timer"]}
	// {"Toggle":["Asleep","AsleepFilled","Awake","CarouselHorizontal","CarouselVertical","Checkbox","CheckboxChecked","CheckboxCheckedFilled","CheckboxIndeterminate","CheckboxIndeterminateFilled","Column","Favorite","FavoriteFilled","FavoriteHalf","Flash","FlashFilled","FlashOff","FlashOffFilled","Grid","Light","LightFilled","List","Locked","Microphone","MicrophoneFilled","MicrophoneOff","MicrophoneOffFilled","Notification","NotificationFilled","NotificationNew","NotificationOff","NotificationOffFilled","Phone","PhoneApplication","PhoneBlock","PhoneBlockFilled","PhoneFilled","PhoneIncoming","PhoneIncomingFilled","PhoneOff","PhoneOffFilled","PhoneOutgoing","PhoneOutgoingFilled","PhoneSettings","PhoneVoice","PhoneVoiceFilled","RadioButton","RadioButtonChecked","Row","Screen","ScreenOff","Star","StarFilled","StarHalf","Thumbnail_1","Thumbnail_2","ThumbsDown","ThumbsDownFilled","ThumbsUp","ThumbsUpFilled","Trophy","TrophyFilled","Unlocked","Video","VideoFilled","VideoOff","VideoOffFilled","View","ViewFilled","ViewMode_1","ViewMode_2","ViewOff","ViewOffFilled"]}
	// {"Travel":["AirlineDigitalGate","AirlineManageGates","AirlinePassengerCare","AirlineRapidBoard","Airport_01","Airport_02","AirportLocation","Arrival","BaggageClaim","Bar","Bicycle","Bus","CabinCare","CabinCareAlert","CabinCareAlt","Cafe","Campsite","Car","CarFront","ChargingStation","ChargingStationFilled","Compass","Construction","Crossroads","CrowdReport","CrowdReportFilled","Cyclist","DeliveryTruck","Departure","DirectionBearRight_01","DirectionBearRight_01Filled","DirectionBearRight_02","DirectionBearRight_02Filled","DirectionCurve","DirectionCurveFilled","DirectionFork","DirectionForkFilled","DirectionLoopLeft","DirectionLoopLeftFilled","DirectionLoopRight","DirectionLoopRightFilled","DirectionMerge","DirectionMergeFilled","DirectionRight_01","DirectionRight_01Filled","DirectionRight_02","DirectionRight_02Filled","DirectionRotaryFirstRight","DirectionRotaryFirstRightFilled","DirectionRotaryRight","DirectionRotaryRightFilled","DirectionRotaryStraight","DirectionRotaryStraightFilled","DirectionSharpTurn","DirectionSharpTurnFilled","DirectionStraight","DirectionStraightFilled","DirectionStraightRight","DirectionStraightRightFilled","DirectionUTurn","DirectionUTurnFilled","Earth","EarthAmericas","EarthAmericasFilled","EarthEuropeAfrica","EarthEuropeAfricaFilled","EarthFilled","EarthSoutheastAsia","EarthSoutheastAsiaFilled","Fire","FlaggingTaxi","FlightInternational","FlightRoster","FlightSchedule","GasStation","GasStationFilled","Globe","Harbor","Helicopter","HelpDesk","Hospital","HospitalBed","Hotel","Lifesaver","Location","LocationCompany","LocationCompanyFilled","LocationFilled","LocationHazard","LocationHazardFilled","LocationHeart","LocationHeartFilled","LocationPerson","LocationPersonFilled","LocationStar","LocationStarFilled","Map","MapBoundary","Milestone","Monument","Mountain","NavaidCivil","NavaidDme","NavaidHelipad","NavaidMilitary","NavaidMilitaryCivil","NavaidNdb","NavaidNdbDme","NavaidPrivate","NavaidSeaplane","NavaidTacan","NavaidVhfor","NavaidVor","NavaidVordme","NavaidVortac","PalmTree","PassengerDrinks","PassengerPlus","Pedestrian","PedestrianFamily","PedestrianChild","PicnicArea","Plane","PlanePrivate","PlaneSea","Police","Restaurant","RestaurantFine","Road","RoadWeather","Scooter","ScooterFront","ServiceDesk","Shuttle","StopSign","StopSignFilled","Swim","Taxi","Theater","TrafficEvent","TrafficFlow","TrafficFlowIncident","TrafficIncident","TrafficWeatherIncident","TrafficCone","Train","TrainHeart","TrainProfile","TrainSpeed","TrainTicket","TrainTime","Tram","Van","Worship","WorshipChristian","WorshipJewish","WorshipMuslim"]}
	// {"User":["Accessibility","AccessibilityAlt","AccessibilityColor","AccessibilityColorFilled","Collaborate","Cough","Credentials","DogWalker","Education","Events","EventsAlt","FaceActivated","FaceActivatedAdd","FaceActivatedFilled","FaceAdd","FaceCool","FaceDissatisfied","FaceDissatisfiedFilled","FaceDizzy","FaceDizzyFilled","FaceMask","FaceNeutral","FaceNeutralFilled","FacePending","FacePendingFilled","FaceSatisfied","FaceSatisfiedFilled","FaceWink","FaceWinkFilled","Friendship","GenderFemale","GenderMale","Group","GroupPresentation","Identification","Partnership","Person","PersonFavorite","Sight","Transgender","User","UserAccess","UserActivity","UserAdmin","UserAvatar","UserAvatarFilled","UserAvatarFilledAlt","UserCertification","UserData","UserFavorite","UserFavoriteAlt","UserFavoriteAltFilled","UserFilled","UserFollow","UserIdentification","UserMilitary","UserMultiple","UserOnline","UserProfile","UserRole","UserServiceDesk","UserSettings","UserSimulation","UserSpeaker","UserXRay","UserProfileAlt","VoiceActivate"]}
	// {"Weather":["AccumulationIce","AccumulationPrecipitation","AccumulationRain","AccumulationSnow","AgricultureAnalytics","Buoy","CarbonAccounting","Cloud","CloudCeiling","Cloudy","CropGrowth","CropHealth","DewPoint","DewPointFilled","Drought","Earthquake","EnergyRenewable","Flood","FloodWarning","Fog","ForecastHail","ForecastHail_30","ForecastLightning","ForecastLightning_30","Hail","Haze","HazeNight","Humidity","HumidityAlt","Hurricane","IceAccretion","IceVision","Lightning","MarineWarning","MixedRainHail","Moon","Moonrise","Moonset","MostlyCloudy","MostlyCloudyNight","NotAvailable","ObservedHail","ObservedLightning","OutlookSevere","PartlyCloudy","PartlyCloudyNight","Pest","Pressure","PressureFilled","RadarEnhanced","RadarWeather","Rain","RainDrizzle","RainHeavy","RainScattered","RainScatteredNight","RainDrop","RefEvapotranspiration","SailboatCoastal","SailboatOffshore","SatelliteRadar","SatelliteWeather","Sleet","Smoke","Snow","SnowBlizzard","SnowHeavy","SnowScattered","SnowScatteredNight","SnowDensity","Snowflake","SoilMoisture","SoilMoistureField","SoilMoistureGlobal","SoilTemperature","SoilTemperatureField","SoilTemperatureGlobal","SolarPanel","StayInside","StormTracker","Sun","Sunrise","Sunset","Temperature","TemperatureCelsius","TemperatureCelsiusAlt","TemperatureFahrenheit","TemperatureFahrenheitAlt","TemperatureFeelsLike","TemperatureFrigid","TemperatureHot","TemperatureInversion","TemperatureMax","TemperatureMin","TemperatureWater","Thunderstorm","ThunderstormScattered","ThunderstormScatteredNight","ThunderstormSevere","ThunderstormStrong","Tides","Tornado","TornadoWarning","TropicalStorm","TropicalStormModelTracks","TropicalStormTracks","TropicalWarning","Tsunami","UvIndex","UvIndexAlt","UvIndexFilled","WaveDirection","WaveHeight","WavePeriod","WeatherFrontCold","WeatherFrontStationary","WeatherFrontWarm","WeatherStation","WindGusts","WindPower","WindStream","Windy","WindyDust","WindySnow","WindyStrong","WinterWarning","WintryMix"]}
}

// "sql:type:string;max-len:50;unique=false;default=sss;null=true;empty=false" "validation:required;max:50"

// TypeString ignores: Validate,DefaultFunc,GoType,Descriptior
const TypeStringField = `
entfield:TypeString;
Unique=(true/false);
Sensitive=(true/false);
Match=(string_to_regexp);
MinLen=(string_to_int);
NotEmpty=(true/false);
MaxLen=(string_to_in);
Default=(string);
Nillable=(true/false);
Optional=(true/false);
Immutable=(true/false);
Comment=(string);
StructTag=(string);
StorageKey=(string);
SchemaType=(JSONString_to_map[string][string]};
Annotations=([]string_to_schema.Annotation)
`

// TypeText ignores: Validate,DefaultFunc,GoType,Descriptior
const TypeTextField = `
entfield:TypeText;
Unique=(true/false);
Sensitive=(true/false);
Match=(string_to_regexp);
MinLen=(string_to_int);
NotEmpty=(true/false);
MaxLen=(string_to_in);
Default=(string);
Nillable=(true/false);
Optional=(true/false);
Immutable=(true/false);
Comment=(string);
StructTag=(string);
StorageKey=(string);
SchemaType=(JSONString_to_map[string][string]};
Annotations=([]string_to_schema.Annotation)
`

// TypeBytesField ignores: Validate,DefaultFunc,GoType,Descriptior
const TypeBytesField = `
entfield:TypeBytes;
Default=(string_to_bytes);
Nillable=(true/false);
Optional=(true/false);
Unique=(true/false);
Immutable=(true/false);
Comment=(string);
StructTag=(string);
MaxLen=(string_to_in);
StorageKey=(string);
SchemaType=(JSONString_to_map[string][string]};
Annotations=([]string_to_schema.Annotation);
`

// TypeBool ignores: GoType,Descriptior
const TypeBoolField = `
entfield:TypeBool;
Default=(true/false);
Nillable=(true/false);
Optional=(true/false);
Immutable=(true/false);
Comment=(string);
StructTag=(string);
StorageKey=(string);
MaxLen=(string_to_in);
Annotations=([]string_to_schema.Annotation);
`

// TypeTime ignores: GoType,Descriptior
const TypeTimeField = `
entfield:TypeTime;
Default=(string_expr);
UpdateDefault=(string_expr);
Nillable=(true/false);
Optional=(true/false);
Immutable=(true/false);
Comment=(string);
StructTag=(string);
StorageKey=(string);
SchemaType=(JSONString_to_map[string][string]};
Annotations=([]string_to_schema.Annotation);
`

// TypeJSONField ignores: Validate,DefaultFunc,GoType,Descriptior
const TypeJSONField = `
entfield:TypeJSON;
Optional=(true/false);
Immutable=(true/false);
Comment=(string);
StructTag=(string);
StorageKey=(string);
SchemaType=(JSONString_to_map[string][string]};
Annotations=([]string_to_schema.Annotation);
`

const TypeEnumField = `
entfield:TypeEnum;
Values=(...[]string);
NamedValues=(...(map[string]string) to []string);
Default=(string);
StorageKey=(string);
Optional=(true/false);
Immutable=(true/false);
Comment=(string);
Nillable=(true/false);
StructTag=(string);
SchemaType=(JSONString_to_map[string][string]};
Annotations=([]string_to_schema.Annotation);
`

type EntField interface {
	Type() string
}

type EntModel struct {
	Name   string `json:"name"`
	Fields []*EntField
}
