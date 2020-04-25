package encmgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// StxEncMgrMetrics is used to marshal and unmarshall metic data from daemon service
type StxEncMgrMetrics struct {
	Enclosures []struct {
		Attributes struct {
			ID         string `json:"id"`
			Model      string `json:"model"`
			Serial     string `json:"serial"`
			SasAddress string `json:"sas_address"`
		} `json:"attributes"`
		Elements []struct {
			ArrayDevice struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Array Device,omitempty"`
			PowerSupply struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Power Supply,omitempty"`
			CoolingElement struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Cooling Element,omitempty"`
			TemperatureSensor struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Temperature sensor,omitempty"`
			DoorLockSensor struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Door Lock Sensor,omitempty"`
			AudibleAlarm struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Audible Alarm,omitempty"`
			EnclosureServicesControllerElectronics struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Enclosure Services Controller Electronics,omitempty"`
			Enclosure struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Enclosure,omitempty"`
			VoltageSensor struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Voltage Sensor,omitempty"`
			CurrentSensor struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Current Sensor,omitempty"`
			SASExpander struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"SAS Expander,omitempty"`
			SASConnector struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"SAS Connector,omitempty"`
			SBBMidplaneInterconnect struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"SBB Midplane Interconnect,omitempty"`
			EnclosureElectronicsPower struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Enclosure Electronics Power,omitempty"`
			EnclosureSettings struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Enclosure Settings,omitempty"`
			EnclosureElectronicsDiagnostics struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Enclosure Electronics Diagnostics,omitempty"`
			Sideplane struct {
				Device []struct {
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
				} `json:"Device"`
			} `json:"Sideplane,omitempty"`
		} `json:"elements"`
	} `json:"enclosures"`
}

// WriteJSONReportToFile - dumps as JSON EncMgr object
func WriteJSONReportToFile(rpt *StxEncMgrMetrics, filePath string) error {
	jsonData, err := json.MarshalIndent(rpt, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	return err
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
