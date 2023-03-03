package public

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
)

func GenOrderMap() {
	text := `{
        "class": "bigip.dns",
        "group_id": "aa13f1cf-b96d-11ed-89a1-a497b115c6b6",
        "project_id": "56",
        "VIRTUAL_SERVER": {
            "address": "192.168.113.25",
            "class": "dns.server_virtual_server",
            "condition": true,
            "name": "test_rollback_VS25",
            "partition": "Common",
            "port": 80,
            "raw_configs": {
                "monitors": [
                    "/Common/sunrun_test_f5_monitor"
                ]
            },
            "server_name": "test"
        },
        "POL1": {
            "class": "dns.pool",
            "name": "pool0_test25.ordermap.cn",
            "partition": "Common",
            "type": "a",
            "load_balancing_preferred": "round-robin",
            "load_balancing_alternate": "round-robin",
            "load_balancing_fallback": "global-availability"
        },
        "POL_MEMBERS1": {
            "class": "dns.pool_member",
            "name": "pool0_test25.ordermap.cn",
            "type": "a",
            "members": [
                {
                    "full_path": "/Common/test:test_rollback_VS25"
                }
            ]
        },
        "WIDE_IP": {
            "class": "dns.wide_ip",
            "name": "test25.orderMap.cn",
            "partition": "Common",
            "condition": true
        }
    }`

	var objs orderedmap.OrderedMap

	if err := json.Unmarshal([]byte(text), &objs); err != nil {
		fmt.Println(err)
		return
	}

	objs_len := []string{"VIRTUAL_SERVER", "POL1", "POL_MEMBERS1", "WIDE_IP"}

	reversed := orderedmap.New()
	for i := len(objs_len) - 1; i >= 0; i-- {
		val, _ := objs.Get(objs_len[i])
		reversed.Set(objs_len[i], val)
	}

	var err error
	var data []byte
	if data, err = json.Marshal(reversed); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
