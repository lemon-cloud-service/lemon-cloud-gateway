package handler

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_core"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-gateway/define"
	"github.com/lemon-cloud-service/lemon-cloud-gateway/manager"
	"github.com/lemon-cloud-service/lemon-cloud-gateway/plugins"
	_ "github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/micro/v2/api"
	"github.com/micro/micro/v2/cmd"
	"os"
)

func SystemStart() {
	var err error
	// 打印系统信息
	define.PrintSystemInfo()

	// 从磁盘中读取配置文件
	lccu_log.Info("Start reading configuration files...")
	err = manager.ConfigManagerInstance().Init()
	if err != nil {
		lccu_log.Error("System start failed. Error reading configuration file: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Configuration file read completed")

	// 调用核心服务启动函数
	lccu_log.Info("Start configuring the registry...")
	err = lccc_core.CoreService().Start(&lccc_core.CoreStartParams{
		RunGrpcService:                 false,
		ServiceGeneralConfig:           manager.ConfigManagerInstance().GeneralConfig(),
		ServiceBaseInfo:                define.GetServiceBaseInfo(),
		ServiceApplicationInfo:         define.GetServiceApplicationInfo(),
		GrpcServiceImplRegisterHandler: nil,
		SystemSettingsDefine:           define.GetSystemSettings(),
	})
	if err != nil {
		lccu_log.Error("System start failed. Reason: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Registry config completed")

	// 初始化各种插件，然后启动micro的命令行组件
	plugins.Init()
	api.Name = fmt.Sprintf("%v.%v", manager.ConfigManagerInstance().GeneralConfig().Service.Namespace, define.GetServiceBaseInfo().ServiceKey)
	api.Address = fmt.Sprintf(":%d", manager.ConfigManagerInstance().GatewayConfig().Gateway.Server.HttpPort)
	svcAddress := fmt.Sprintf(":%d", manager.ConfigManagerInstance().GeneralConfig().Service.Port)
	cmd.Init(
		micro.Version(define.SYSTEM_INFO_VERSION),
		micro.Address(svcAddress),
		micro.Registry(lccc_core.CoreService().GenerateMicroRegistry(manager.ConfigManagerInstance().GeneralConfig())))
}
