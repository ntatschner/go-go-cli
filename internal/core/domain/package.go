package domain

// RuleType defines types of rules for application detection or requirements.
type RuleType string

const (
    Requirement RuleType = "requirement"
    Detection   RuleType = "detection"
)

// MobileApp represents the configuration for a mobile application.
type MobileApp struct {
    DisplayName                  string
    Description                  string
    Publisher                    string
    LargeIcon                    MimeContent
    IsFeatured                   bool
    PrivacyInformationURL        string
    InformationURL               string
    Owner                        string
    Developer                    string
    Notes                        string
    PublishingState              PublishingState
    CommittedContentVersion      string
    FileName                     string
    Size                         int64
    InstallCommandLine           string
    UninstallCommandLine         string
    ApplicableArchitectures      WindowsArchitecture
    MinimumFreeDiskSpaceInMB     int32
    MinimumMemoryInMB            int32
    MinimumNumberOfProcessors    int32
    MinimumCpuSpeedInMHz         int32
    Rules                        []Rule
    InstallExperience            InstallExperience
    ReturnCodes                  []Win32LobAppReturnCode
    MsiInformation               MsiInformation
    SetupFilePath                string
    MinimumSupportedWindowsRelease string
}

// MimeContent represents MIME content for icons or other binary data.
type MimeContent struct {
    Type  string
    Value []byte
}

// InstallExperience defines the installation experience for the app.
type InstallExperience struct {
    RunAsAccount          RunAsAccountType
    DeviceRestartBehavior RestartBehavior
}

// Win32LobAppReturnCode represents return codes for the Win32 app.
type Win32LobAppReturnCode struct {
    ReturnCode int32
    Type       ReturnCodeType
}

// MsiInformation represents MSI package information.
type MsiInformation struct {
    ProductCode    string
    ProductVersion string
    UpgradeCode    string
    RequiresReboot bool
    PackageType    MsiPackageType
    ProductName    string
    Publisher      string
}

// MobileAppService interface for managing mobile apps.
type MobileAppService interface {
    CreateMobileApp(app MobileApp) error
}

// Adapter for Microsoft Graph SDK.
type GraphSDKAdapter struct {
    // Graph client and other necessary fields
}

// CreateMobileApp implements MobileAppService, using the Microsoft Graph SDK.
func (g *GraphSDKAdapter) CreateMobileApp(app MobileApp) error {
    // Convert domain.MobileApp to graphmodels.MobileApp and use the SDK to create the app.
    return nil
}

// PublishingState represents the state of publishing for the mobile app.
type PublishingState string

const (
    Processing PublishingState = "processing"
    Published  PublishingState = "published"
    Failed     PublishingState = "failed"
)

// WindowsArchitecture defines the supported architectures for the app.
type WindowsArchitecture string

const (
    X86   WindowsArchitecture = "x86"
    X64   WindowsArchitecture = "x64"
    ARM   WindowsArchitecture = "arm"
    ARM64 WindowsArchitecture = "arm64"
)

// RunAsAccountType defines the type of account to run the app installation as.
type RunAsAccountType string

const (
    User  RunAsAccountType = "user"
    Admin RunAsAccountType = "admin"
)

// RestartBehavior defines the behavior of device restart during app installation.
type RestartBehavior string

const (
    Allow   RestartBehavior = "allow"
    Suppress RestartBehavior = "suppress"
    Force   RestartBehavior = "force"
)

// ReturnCodeType defines the type of return code.
type ReturnCodeType string

const (
    Success ReturnCodeType = "success"
    Failure ReturnCodeType = "failure"
)

// MsiPackageType defines the type of MSI package.
type MsiPackageType string

const (
    PerUser    MsiPackageType = "perUser"
    PerMachine MsiPackageType = "perMachine"
)