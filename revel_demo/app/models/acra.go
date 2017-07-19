package models

import (
	"github.com/revel/revel"
    "time"
)

type Acra struct {
    Id int64 `xorm:"pk autoincr"`
    Logcat string `xorm:"Text" json:"LOGCAT"`
    DeviceFeatures string `xorm:"Text" json:"DEVICE_FEATURES"`
    AppVersionCode int64 `xorm:"Int" json:"APP_VERSION_CODE"`
    UserCrashDate time.Time `xorm:"DateTime" json:"USER_CRASH_DATE"`
    CusrtomData string `xorm:"Text" json:"CUSTOM_DATA"`
    AvailableMemSize int64 `xorm:"Int" json:"AVAILABLE_MEM_SIZE"`
    Brand string `xorm:"char(255)" json:"BRAND"`
    InstallationId string `xorm:"char(255)" json:"INSTALLATION_ID"`
    SharedPreferences string `xorm:"Text" json:"SHARED_PREFERENCES"`
    FilePath string `xorm:"char(255)" json:"FILE_PATH"`
    StackTrace string `xorm:"Text" json:"STACK_TRACE"`
    PhoneModel string `xorm:"char(255)" json:"PHONE_MODEL"`
    TotalMemSize int64 `xorm:"Int" json:"TOTAL_MEM_SIZE"`
    PackageName string `xorm:"char(255)" json:"PACKAGE_NAME"`
    IsSilent bool `xorm:"Bool" json:"IS_SILENT"`
    DumpsysMeminfo string `xorm:"Text" json:"DUMPSYS_MEMINFO"`
    InitialConfiguration string `xorm:"Text" json:"INITIAL_CONFIGURATION"`
    BuildConfig string `xorm:"Text" json:"BUILD_CONFIG"`
    Product string `xorm:"char(255)" json:"PRODUCT"`
    Display string `xorm:"Text" json:"DISPLAY"`
    ReportId string `xorm:"char(255)" json:"REPORT_ID"`
    AppVersionName string `xorm:"char(32)" json:"APP_VERSION_NAME"`
    CrashConfiguration string `xorm:"Text" json:"CRASH_CONFIGURATION"`
    Build string `xorm:"Text" json:"BUILD"`
    Environment string `xorm:"Text" json:"ENVIRONMENT"`
    UserEmail string `xorm:"char(255)" json:"USER_EMAIL"`
    AndroidVersion string `xorm:"char(32)" json:"ANDROID_VERSION"`
    UserAppStartDate time.Time `xorm:"DateTime" json:"USER_APP_START_DATE"`
}

func (a Acra) InsertOne(acra Acra) bool {
    has, err := engine.Table("ACRA").Insert(acra)
    if err != nil {
        revel.WARN.Println(has)
        revel.WARN.Printf("Error %v", err)
        return false
    }
    return true
}
