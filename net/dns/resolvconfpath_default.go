// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gokrazy

package dns

const (
	resolvConf = "/data/adb/tailscale/tmp/resolv.conf"
	backupConf = "/data/adb/tailscale/tmp/resolv.pre-tailscale-backup.conf"
)
