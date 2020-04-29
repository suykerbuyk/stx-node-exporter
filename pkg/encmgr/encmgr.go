package encmgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type EncAttributes struct {
	ID         string `json:"id"`
	Model      string `json:"model"`
	Serial     string `json:"serial"`
	SasAddress string `json:"sas_address"`
}
type EncArrayDevice struct {
	Type              int    `json:"type"`
	TypeStr           string `json:"type_str"`
	Number            int    `json:"number"`
	RawBytes          string `json:"raw_bytes"`
	GlobalStatus      int    `json:"global_status,omitempty"`
	GlobalStatusStr   string `json:"global_status_str,omitempty"`
	Swap              int    `json:"swap"`
	Disabled          int    `json:"disabled"`
	Prdfail           int    `json:"prdfail"`
	RebuildRemapAbort int    `json:"rebuild_remap_abort"`
	RebuildRemap      int    `json:"rebuild_remap"`
	InFailedArray     int    `json:"in_failed_array"`
	InCriticalArray   int    `json:"in_critical_array"`
	ConsistencyCheck  int    `json:"consistency_check"`
	Report            int    `json:"report"`
	Identify          int    `json:"identify"`
	Remove            int    `json:"remove"`
	DoNotRemove       int    `json:"do_not_remove"`
	DeviceOff         int    `json:"device_off"`
	FaultReq          int    `json:"fault_req"`
	FaultSensed       int    `json:"fault_sensed"`
	Status            int    `json:"status,omitempty"`
	StatusStr         string `json:"status_str,omitempty"`
}
type ArrayDeviceArray struct {
	Device []EncArrayDevice `json:"Device,omitempty"`
}

type EncPowerSupplyDevice struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	Identify        int    `json:"identify"`
	DcFailure       int    `json:"dc_failure"`
	AcFailure       int    `json:"ac_failure"`
	TempWarning     int    `json:"temp_warning"`
	OverTempFailure int    `json:"over_temp_failure"`
	Off             int    `json:"off"`
	Failure         int    `json:"failure"`
	HotSwap         int    `json:"hot_swap"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type PowerSupplyArray struct {
	Device []EncPowerSupplyDevice `json:"Device,omitempty"`
}
type EncCoolingDevice struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	ActualSpeed     int    `json:"actual_speed"`
	ActualSpeedCode int    `json:"actual_speed_code"`
	Off             int    `json:"off"`
	Failure         int    `json:"failure"`
	HotSwap         int    `json:"hot_swap"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type CoolingDeviceArray struct {
	Device []EncCoolingDevice `json:Device,omitempty"`
}
type EncTemperatureDevice struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	TempCelcius     int    `json:"temp_celcius"`
	UnderTempWarn   int    `json:"under_temp_warn"`
	UnderTempFail   int    `json:"under_temp_fail"`
	OverTempWarn    int    `json:"over_temp_warn"`
	OverTempFail    int    `json:"over_temp_fail"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type TemperatureDeviceArray struct {
	Device []EncTemperatureDevice `json:"Device,omitempty"`
}
type EncDoorLockSensor struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	Open            int    `json:"open"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type DoorLockSensorArray struct {
	Device []EncDoorLockSensor `json:"Device,omitempty"`
}
type EncAudioAlarm struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	Unrecoverable   int    `json:"unrecoverable"`
	Critical        int    `json:"critical"`
	NonCritical     int    `json:"non_critical"`
	Informational   int    `json:"informational"`
	Remind          int    `json:"remind"`
	Muted           int    `json:"muted"`
	ReqMuted        int    `json:"req_muted"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type AudioAlarmArray struct {
	Device []EncAudioAlarm `json:"Device,omitEmpty"`
}
type EncControllerElectronics struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	Failure         int    `json:"failure"`
	Identify        int    `json:"identify"`
	Report          int    `json:"report"`
	HotSwap         int    `json:"hot_swap"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type ControllerElectronicsArray struct {
	Device []EncControllerElectronics `json:"Device,omitEmpty"`
}
type EncEnclosure struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	Identify        int    `json:"identify"`
	WarnIndication  int    `json:"warn_indication"`
	FailIndication  int    `json:"fail_indication"`
	WarnReq         int    `json:"warn_req"`
	FailReq         int    `json:"fail_req"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type EnclosureArray struct {
	Device []EncEnclosure `json:"Device,omitEmpty"`
}
type EncVoltageSensor struct {
	Type            int     `json:"type"`
	TypeStr         string  `json:"type_str"`
	Number          int     `json:"number"`
	RawBytes        string  `json:"raw_bytes"`
	GlobalStatus    int     `json:"global_status,omitempty"`
	GlobalStatusStr string  `json:"global_status_str,omitempty"`
	Swap            int     `json:"swap"`
	Disabled        int     `json:"disabled"`
	Prdfail         int     `json:"prdfail"`
	Voltage         float64 `json:"voltage"`
	Status          int     `json:"status,omitempty"`
	StatusStr       string  `json:"status_str,omitempty"`
}
type VoltageSensorArray struct {
	Device []EncVoltageSensor `json:"Device,omitempty"`
}

type EncCurrentSensor struct {
	Type            int     `json:"type"`
	TypeStr         string  `json:"type_str"`
	Number          int     `json:"number"`
	RawBytes        string  `json:"raw_bytes"`
	GlobalStatus    int     `json:"global_status,omitempty"`
	GlobalStatusStr string  `json:"global_status_str,omitempty"`
	Swap            int     `json:"swap"`
	Disabled        int     `json:"disabled"`
	Prdfail         int     `json:"prdfail"`
	Current         float64 `json:"current"`
	Status          int     `json:"status,omitempty"`
	StatusStr       string  `json:"status_str,omitempty"`
}
type CurrentSensorArray struct {
	Device []EncCurrentSensor `json:"Device,omitempty"`
}
type EncSASExpander struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type SASExpanderArray struct {
	Device []EncSASExpander `json:"Device,omitempty"`
}
type EncSASConnector struct {
	Type              int    `json:"type"`
	TypeStr           string `json:"type_str"`
	Number            int    `json:"number"`
	RawBytes          string `json:"raw_bytes"`
	GlobalStatus      int    `json:"global_status,omitempty"`
	GlobalStatusStr   string `json:"global_status_str,omitempty"`
	Swap              int    `json:"swap"`
	Disabled          int    `json:"disabled"`
	Prdfail           int    `json:"prdfail"`
	ConnectorType     int    `json:"connector_type"`
	ConnectorTypeStr  string `json:"connector_type_str"`
	Identify          int    `json:"identify"`
	ConnectorPhysLink int    `json:"connector_phys_link"`
	Failure           int    `json:"failure"`
	Status            int    `json:"status,omitempty"`
	StatusStr         string `json:"status_str,omitempty"`
}
type SASConnectorArray struct {
	Device []EncSASConnector `json:"Device,omitempty"`
}
type EncSBBMidplaneInterconnect struct {
	Type                    int    `json:"type"`
	TypeStr                 string `json:"type_str"`
	Number                  int    `json:"number"`
	RawBytes                string `json:"raw_bytes"`
	GlobalStatus            int    `json:"global_status,omitempty"`
	GlobalStatusStr         string `json:"global_status_str,omitempty"`
	Swap                    int    `json:"swap"`
	Disabled                int    `json:"disabled"`
	Prdfail                 int    `json:"prdfail"`
	I2C0Fail                int    `json:"i2c_0_fail"`
	I2C1Fail                int    `json:"i2c_1_fail"`
	I2C2Fail                int    `json:"i2c_2_fail"`
	SgpioFail               int    `json:"sgpio_fail"`
	MidVpdReadFail          int    `json:"mid_vpd_read_fail"`
	MidVpdMismatch          int    `json:"mid_vpd_mismatch"`
	MidVpdRecoveryStatus    int    `json:"mid_vpd_recovery_status"`
	MidVpdRecoveryStatusStr string `json:"mid_vpd_recovery_status_str"`
	Status                  int    `json:"status,omitempty"`
	StatusStr               string `json:"status_str,omitempty"`
}
type SBBMidplaneInterconnectArray struct {
	Device []EncSBBMidplaneInterconnect `json:"Device,omitEmpty"`
}
type EncEnclosureElectronicsPower struct {
	Type                    int    `json:"type"`
	TypeStr                 string `json:"type_str"`
	Number                  int    `json:"number"`
	RawBytes                string `json:"raw_bytes"`
	GlobalStatus            int    `json:"global_status,omitempty"`
	GlobalStatusStr         string `json:"global_status_str,omitempty"`
	Swap                    int    `json:"swap"`
	Disabled                int    `json:"disabled"`
	Prdfail                 int    `json:"prdfail"`
	GemReadyValue           int    `json:"gem_ready_value"`
	GemReadyValueStr        string `json:"gem_ready_value_str"`
	GemsatResetCount        int    `json:"gemsat_reset_count"`
	GemsatWatchdog          int    `json:"gemsat_watchdog"`
	PowerCtrlActivationCode int    `json:"power_ctrl_activation_code"`
	LastResetType           int    `json:"last_reset_type"`
	LastResetTypeStr        string `json:"last_reset_type_str"`
	ResetCount              int    `json:"reset_count"`
	Status                  int    `json:"status,omitempty"`
	StatusStr               string `json:"status_str,omitempty"`
}
type EnclosureElectronicsPowerArray struct {
	Device []EncEnclosureElectronicsPower `json:"Device,omitEmpty"`
}
type EncEnclosureSetting struct {
	Type                     int    `json:"type"`
	TypeStr                  string `json:"type_str"`
	Number                   int    `json:"number"`
	RawBytes                 string `json:"raw_bytes"`
	GlobalStatus             int    `json:"global_status,omitempty"`
	GlobalStatusStr          string `json:"global_status_str,omitempty"`
	Swap                     int    `json:"swap"`
	Disabled                 int    `json:"disabled"`
	Prdfail                  int    `json:"prdfail"`
	EnclosureSettingsChanged int    `json:"enclosure_settings_changed"`
	EnclosureIDActive        int    `json:"enclosure_id_active"`
	OpsPanelRemoved          int    `json:"ops_panel_removed"`
	EnclosureIDNotSupported  int    `json:"enclosure_id_not_supported"`
	ReqEnclosureIDFail       int    `json:"req_enclosure_id_fail"`
	EnclosureID              int    `json:"enclosure_id"`
	Status                   int    `json:"status,omitempty"`
	StatusStr                string `json:"status_str,omitempty"`
}
type EnclosureSettingArray struct {
	Device []EncEnclosureSetting `json:"Device,omitempty"`
}
type EncEnclosureElectronicsDiagnostic struct {
	Type                   int    `json:"type"`
	TypeStr                string `json:"type_str"`
	Number                 int    `json:"number"`
	RawBytes               string `json:"raw_bytes"`
	GlobalStatus           int    `json:"global_status,omitempty"`
	GlobalStatusStr        string `json:"global_status_str,omitempty"`
	Swap                   int    `json:"swap"`
	Disabled               int    `json:"disabled"`
	Prdfail                int    `json:"prdfail"`
	RAMLogFull             int    `json:"ram_log_full"`
	RAMLogFullWarn         int    `json:"ram_log_full_warn"`
	NonVolatileLogFull     int    `json:"non_volatile_log_full"`
	NonVolatileLogFullWarn int    `json:"non_volatile_log_full_warn"`
	UserFull               int    `json:"user_full"`
	SystemNew              int    `json:"system_new"`
	UserNew                int    `json:"user_new"`
	Status                 int    `json:"status,omitempty"`
	StatusStr              string `json:"status_str,omitempty"`
}
type EnclosureElectronicsDiagnosticArray struct {
	Device []EncEnclosureElectronicsDiagnostic `json:"Device,omitempty"`
}
type EncSidePlane struct {
	Type            int    `json:"type"`
	TypeStr         string `json:"type_str"`
	Number          int    `json:"number"`
	RawBytes        string `json:"raw_bytes"`
	GlobalStatus    int    `json:"global_status,omitempty"`
	GlobalStatusStr string `json:"global_status_str,omitempty"`
	Swap            int    `json:"swap"`
	Disabled        int    `json:"disabled"`
	Prdfail         int    `json:"prdfail"`
	CableFaultReq   int    `json:"cable_fault_req"`
	Fault           int    `json:"fault"`
	CableFault      int    `json:"cable_fault"`
	Powered         int    `json:"powered"`
	CoverRemoved    int    `json:"cover_removed"`
	FailReq         int    `json:"fail_req"`
	IdentifyReq     int    `json:"identify_req"`
	Status          int    `json:"status,omitempty"`
	StatusStr       string `json:"status_str,omitempty"`
}
type EncSidePlaneArray struct {
	Device []EncSidePlane `json:"Device,omitempty"`
}

// StxEncMgrMetrics is used to marshal and unmarshall metic data from daemon service
// it is very important the element structs be pointers as to avoid instantiating empty
// instances.
type StxEncMgrMetrics struct {
	Enclosures []struct {
		Attributes EncAttributes `json:"attributes"`
		Elements   struct {
			ArrayDevices                    *ArrayDeviceArray                    `json:"Array Device,omitempty"`
			PowerSupplies                   *PowerSupplyArray                    `json:"Power Supply,omitempty"`
			CoolingDevices                  *CoolingDeviceArray                  `json:"Cooling Element,omitempty"`
			TemperatureDevices              *TemperatureDeviceArray              `json:"Temperature sensor,omitempty"`
			DoorLockSensors                 *DoorLockSensorArray                 `json:"Door Lock Sensor,omitempty"`
			AudioAlarms                     *AudioAlarmArray                     `json:"Audible Alarm,omitempty"`
			ControllerElectronics           *ControllerElectronicsArray          `json:"Enclosure Services Controller Electronics,omitempty"`
			Enclosures                      *EnclosureArray                      `json:"Enclosure,omitempty"`
			VoltageSensors                  *VoltageSensorArray                  `json:"Voltage Sensor,omitempty"`
			CurrentSensors                  *CurrentSensorArray                  `json:"Current Sensor,omitempty"`
			SASExpanders                    *SASExpanderArray                    `json:"SAS Expander,omitempty"`
			SASConnectors                   *SASConnectorArray                   `json:"SAS Connector,omitempty"`
			SBBMidplaneInterconnects        *SBBMidplaneInterconnectArray        `json:"SBB Midplane Interconnect,omitempty"`
			EnclosureElectronicsPower       *EnclosureElectronicsPowerArray      `json:"Enclosure Electronics Power,omitempty"`
			EnclosureSettings               *EnclosureSettingArray               `json:"Enclosure Settings,omitempty"`
			EnclosureElectronicsDiagnostics *EnclosureElectronicsDiagnosticArray `json:"Enclosure Electronics Diagnostics,omitempty"`
			Sideplanes                      *EncSidePlaneArray                   `json:"Sideplane,omitempty"`
		} `json:"elements"`
	} `json:"enclosures"`
}

// StxEncMetricsFromFile - returns object with data from file
func StxEncMetricsFromFile(enc *StxEncMgrMetrics, filePath string) error {
	jsonSourceFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonSourceFile.Close()
	byteValues, err := ioutil.ReadAll(jsonSourceFile)
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(byteValues), &enc)
	return err
}

// PrintJSONReport - unmarshalls and dumps enclosure report to console.
func PrintJSONReport(enc *StxEncMgrMetrics) error {
	jsonData, err := json.MarshalIndent(&enc, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", jsonData)
	return nil
}

// ReadFromJSONFile - Deserialize JSON file into object
func (enc *StxEncMgrMetrics) ReadFromJSONFile(filePath string) error {
	jsonSourceFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonSourceFile.Close()
	byteValues, err := ioutil.ReadAll(jsonSourceFile)
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(byteValues), &enc)
	enc.sanitizeJSON()
	return err
}

// ReadFromNetwork - Deserialize from network uri
func (enc *StxEncMgrMetrics) ReadFromNetwork(uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byteValues, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(byteValues), &enc)
	enc.sanitizeJSON()
	return err
}

// WriteToJSONFile - Serialize object to JSON file
func (enc *StxEncMgrMetrics) WriteToJSONFile(filePath string) error {
	jsonData, err := json.MarshalIndent(enc, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	return err
}

// sanitizeJSON - Clean up the values of superfolous/illegal chars
func (enc *StxEncMgrMetrics) sanitizeJSON() {
	for idx := range enc.Enclosures {
		enc.Enclosures[idx].Attributes.Model = strings.Replace(enc.Enclosures[idx].Attributes.Model, "\"", "", -1)
	}
}
