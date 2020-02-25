package manager

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_define"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_model"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_config"
	"github.com/lemon-cloud-service/lemon-cloud-gateway/define"
	"sync"
)

type ConfigManager struct {
	generalConfig *lccc_model.GeneralConfig
	gatewayConfig *define.GatewayConfig
}

var configManagerInstance *ConfigManager
var configManagerOnce sync.Once

func ConfigManagerInstance() *ConfigManager {
	configManagerOnce.Do(func() {
		configManagerInstance = &ConfigManager{}
	})
	return configManagerInstance
}

func (cm *ConfigManager) Init() error {
	if !cm.InitializationStatus() {
		cm.generalConfig = &lccc_model.GeneralConfig{}
		cm.gatewayConfig = &define.GatewayConfig{}
		if err := lccu_config.LoadYamlConfigFile(lccc_define.FILE_PATH_GENERAL_CONFIG_FILE, cm.generalConfig); err != nil {
			return err
		}
		if err := lccu_config.LoadYamlConfigFile(lccc_define.FILE_PATH_GENERAL_CONFIG_FILE, cm.gatewayConfig); err != nil {
			return err
		}
	}
	return nil
}

func (cm *ConfigManager) GeneralConfig() *lccc_model.GeneralConfig {
	return cm.generalConfig
}

func (cm *ConfigManager) GatewayConfig() *define.GatewayConfig {
	return cm.gatewayConfig
}

func (cm *ConfigManager) InitializationStatus() bool {
	return cm.generalConfig != nil
}
