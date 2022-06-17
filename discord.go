// Discordgo - Discord bindings for Go
// Available at https://github.com/bwmarrin/discordgo

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains high level helper functions and easy entry points for the
// entire discordgo package.  These functions are being developed and are very
// experimental at this point.  They will most likely change so please use the
// low level functions if that's a problem.

// Package discordgo provides Discord binding for Go
package discordgo

import (
	"net/http"
	"time"
)

// VERSION of DiscordGo, follows Semantic Versioning. (http://semver.org/)
const VERSION = "0.25.0"

// New creates a new Discord session with provided token.
// If the token is for a bot, it must be prefixed with "Bot "
// 		e.g. "Bot ..."
// Or if it is an OAuth2 token, it must be prefixed with "Bearer "
//		e.g. "Bearer ..."
func New(token string) (s *Session, err error) {

	// Create an empty Session interface.
	s = &Session{
		State:                  NewState(),
		Ratelimiter:            NewRatelimiter(),
		StateEnabled:           true,
		Compress:               true,
		ShouldReconnectOnError: true,
		ShouldRetryOnRateLimit: true,
		ShardID:                0,
		ShardCount:             1,
		MaxRestRetries:         3,
		Client:                 &http.Client{Timeout: (20 * time.Second)},
		UserAgent:              "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36",
		sequence:               new(int64),
		LastHeartbeatAck:       time.Now().UTC(),
	}

	// Initilize the Identify Package with defaults
	// These can be modified prior to calling Open()
	s.Identify.Compress = true
	s.Identify.Capabilities = 509
	s.Identify.Properties.Browser = "Chrome"
	s.Identify.Properties.BrowserUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36"
	s.Identify.Properties.BrowserVersion = "101.0.4951.67"
	s.Identify.Properties.ClientBuildNumber = 132647
	s.Identify.Properties.ClientEventSource = nil
	s.Identify.Properties.Device = ""
	s.Identify.Properties.OS = "Windows"
	s.Identify.Properties.OSVersion = "10"
	s.Identify.Properties.Referrer = ""
	s.Identify.Properties.ReferrerCurrent = ""
	s.Identify.Properties.ReferringDomain = ""
	s.Identify.Properties.ReferringDomainCurrent = ""
	s.Identify.Properties.ReleaseChannel = "stable"
	s.Identify.Properties.SystemLocale = "en-US"
	s.Identify.Token = token
	s.Token = token

	return
}
