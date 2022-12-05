package ADops

import (
	"fmt"
	"github.com/oleiade/reflections"
	"strings"
	"sunrun/public"
)

type VirtualServer struct {
	ID                       string `json:"id" db:"id"`                                   //编号
	Name                     string `json:"name" db:"name"`                               //名称
	Availability             string `json:"availability" db:"availability,option"`        //vs状态
	ProviderName             string `json:"provider_name" db:"provider_name"`             //产品名
	DeviceGroupID            string `json:"device_group_id" db:"device_group_id"`         //设备组ID
	DeviceGroupName          string `json:"device_group_name" db:"device_group_name,out"` //设备组名字
	FullPath                 string `json:"full_path" db:"full_path"`                     //全路径
	Partition                string `json:"partition" db:"partition"`                     //分区名称
	State                    string `json:"state" db:"state"`                             //开启状态
	Type                     string `json:"type" db:"type"`                               //类型
	DestAddress              string `json:"dest_address" db:"dest_address"`               //目标地址
	Protocol                 string `json:"protocol" db:"protocol"`                       //协议
	Port                     int    `json:"port" db:"port"`                               //端口号
	VirtualAddressId         string `json:"virtual_address_id" db:"virtual_address_id"`   //虚拟地址id
	public.RemarkStruct             //备注
	public.DescriptionStruct        //描述
}

func getResolveField(i interface{}, s string) (interface{}, string, error) {
	l := strings.Split(s, ".")
	fmt.Printf("l: %s \n", l)
	_, tag, err := resolveField(i, l, 0)
	if err != nil {
		return nil, "", err
	}

	return nil, tag, nil
}

func resolveField(i interface{}, l []string, s int) (interface{}, string, error) {
	ni, err := reflections.GetField(i, l[s])
	if err != nil {
		return nil, "", err
	}
	if len(l)-1 == s {
		tag, e := reflections.GetFieldTag(i, l[s], "db")
		return ni, tag, e
	}

	return resolveField(ni, l, s+1)
}

func WithResolveSelectFields(i interface{}, s ...string) []string {
	var result []string

	for _, v := range s {
		_, tag, err := getResolveField(i, v)
		fmt.Printf("Value: %s, Tag: %s \n", v, tag)
		if err != nil {
			return nil
		}

		if tag == "" || tag == "-" {
			continue
		}

		index := strings.Index(tag, ",")
		if index > 0 {
			tag = tag[:index]
		}
		result = append(result, tag)

	}
	return result
}

func CreateVirtualServer() {
	a := VirtualServer{}
	fmt.Println(a)

	t := WithResolveSelectFields(a, "ID", "Name", "ProviderName", "DeviceGroupName", "DestAddress", "Port",
		"Protocol", "Type", "ProjectID", "DeviceGroupID", "Availability", "State", "FullPath", "Remark")
	fmt.Println(t)
}
