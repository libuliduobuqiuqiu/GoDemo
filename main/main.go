package main

import (
	"fmt"
	"git.gzsunrun.cn/ad/gofx/storage/tag"
	"github.com/oleiade/reflections"
)

type Device struct {
	ID                   string `json:"id" db:"id"`
	AvailabilityID       string `json:"availability_zone_id" db:"availability_zone_id"`
	AvailabilityZoneName string `json:"availability_zone_name" db:"availability_zone_name,out"`
	RegionName           string `json:"region_name" db:"region_name,out"`
	Contact              string `json:"contact" db:"contact"`
	Address              string `json:"address" db:"address"`
	ApiPort              int    `json:"api_port" db:"api_port"`
	SSHPort              int    `json:"ssh_port" db:"ssh_port"`
	ProjectID            string `json:"project_id" db:"project_id"`
	MStartDate           string `json:"maintenance_start_date" db:"maintenance_start_date"`
	MEndData             string `json:"maintenance_end_date" db:"maintenance_end_date"`
	Buyer                string `json:"buyer" db:"buyer"`
	UserName             string `json:"username" db:"username"`
	DeviceGroupID        string `json:"device_group_id" db:"device_group_id"`
	DeviceGroupName      string `json:"device_group_name" db:"-"` //设备组名称
	SynGroupID           string `json:"syn_group_id" db:"syn_group_id"`
	UpdateTime           string `json:"update_time" db:"update_time"`     //UpdateTime : 配置最后更新时间
	DGSynStatus          string `json:"dg_syn_status" db:"dg_syn_status"` //设备组同步状态
	DgSynTime            string `json:"dg_syn_time" db:"dg_syn_time"`     //设备组同步时间
	SGSynStatus          string `json:"sg_syn_status" db:"sg_syn_status"` //同步组同步状态
	SGSynTime            string `json:"sg_syn_time" db:"sg_syn_time"`     //同步组同步时间
	StateSynTime         string `json:"state_syn_time" db:"state_syn_time"`
	StateSynStatus       string `json:"state_syn_status" db:"state_syn_status"`
	Sync                 bool   `json:"sync" db:"-"` //是否同步

	SynDeviceGroup    bool              `json:"-" db:"-"` //是否同步设备组信息
	SynSyncGroup      bool              `json:"-" db:"-"` //是否同步同步组信息
	SynDeviceInfo     bool              `json:"-" db:"-"` //是否同步设备信息
	SyncConfigToGroup bool              `json:"-" db:"-"` //强制同步到组
	DeviceRunningConf string            `json:"-" db:"-"` //配置文件内容
	SynZoneRunner     bool              `json:"-" db:"-"` //是否同步ZoneRunner模块
	Devices           map[string]string `json:"devices"`
}

func main() {
	//r := GinDemo.GinInit()
	// r.Run(":8088")
	//SqlxDemo.InitDB()
	//musicNames := []string{"zhangsan", "linshukai", "lsk"}
	//musicList, err := SqlxDemo.QueryByNames(musicNames)
	//if err != nil {
	//	fmt.Println("通过Name查询Music数据失败：", err)
	//}
	//for _, musicItem := range musicList {
	//	fmt.Println(musicItem)
	//}

	a := Device{RegionName: "linshukai"}
	val, err := reflections.GetField(a, "RegionName")

	if err == nil {
		fmt.Println(val)
	}

	dbs, err := tag.NewParser(a).Prase()
	fmt.Println(dbs)

	taglen := len(dbs)
	for i := 0; i < taglen; i++ {
		fmt.Println(dbs[i].FieldName, dbs[i].In)

		if !dbs[i].In {
			continue
		}

		val, _ := tag.GetVal(a, dbs[i].FieldName)
		fmt.Println("Val: ", taglen, val)
	}

}
