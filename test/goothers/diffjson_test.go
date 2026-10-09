//go:build integration

package goothers

import (
	"godemo/internal/goothers"
	"testing"
)

func TestDiffJson(t *testing.T) {

	DiffCitrixJson()

	DiffHorizonJson()

	DiffYKJson()

	// DiffZDNSJson()
	//

	DiffH3cJson()

	DiffHillstoneJson()
}

func DiffNginxJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_nginx.json",
		"/mnt/d/Company/provider_json/remote_nginx.json",
	)
}

func DiffCitrixJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_citrix.json",
		"/mnt/d/Company/provider_json/remote_citrix.json",
	)
}

func DiffHillstoneNGFWJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_hillstone_ngfw.json",
		"/mnt/d/Company/provider_json/remote_hillstone_ngfw.json",
	)
}

func DiffHorizonJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_horizon.json",
		"/mnt/d/Company/provider_json/remote_horizon.json",
	)
}

func DiffYKJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_yk.json",
		"/mnt/d/Company/provider_json/remote_yk.json",
	)
}

func DiffZDNSJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_zdns.json",
		"/mnt/d/Company/provider_json/remote_zdns.json",
	)
}

func DiffH3cJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_h3c.json",
		"/mnt/d/Company/provider_json/remote_h3c.json",
	)
}

func DiffHillstoneJson() {
	goothers.DiffJson(
		"/mnt/d/Company/provider_json/local_hillstone.json",
		"/mnt/d/Company/provider_json/remote_hillstone.json",
	)

}
