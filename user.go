package discordgo

import "time"

// UserFlags is the flags of "user" (see UserFlags* consts)
// https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlags int

// Valid UserFlags values
const (
	UserFlagDiscordEmployee           UserFlags = 1 << 0
	UserFlagDiscordPartner            UserFlags = 1 << 1
	UserFlagHypeSquadEvents           UserFlags = 1 << 2
	UserFlagBugHunterLevel1           UserFlags = 1 << 3
	UserFlagHouseBravery              UserFlags = 1 << 6
	UserFlagHouseBrilliance           UserFlags = 1 << 7
	UserFlagHouseBalance              UserFlags = 1 << 8
	UserFlagEarlySupporter            UserFlags = 1 << 9
	UserFlagTeamUser                  UserFlags = 1 << 10
	UserFlagSystem                    UserFlags = 1 << 12
	UserFlagBugHunterLevel2           UserFlags = 1 << 14
	UserFlagVerifiedBot               UserFlags = 1 << 16
	UserFlagVerifiedBotDeveloper      UserFlags = 1 << 17
	UserFlagDiscordCertifiedModerator UserFlags = 1 << 18
	UserFlagBotHTTPInteractions       UserFlags = 1 << 19
	UserFlagSpammer                   UserFlags = 1 << 20
	UserFlagActiveDeveloper           UserFlags = 1 << 22
)

// UserPremiumType is the premium type of a user (see UserPremiumType* consts)
// https://discord.com/developers/docs/resources/user#user-object-premium-types
type UserPremiumType int

// Valid UserPremiumType values
const (
	UserPremiumTypeNone         UserPremiumType = 0
	UserPremiumTypeNitroClassic UserPremiumType = 1
	UserPremiumTypeNitro        UserPremiumType = 2
	UserPremiumTypeNitroBasic   UserPremiumType = 3
)

// UserSettingsType is the type of user settings (see UserSettingsType* consts)
type UserSettingsType int

// Valid UserSettingsType values
const (
	UserSettingsTypePreloadedUserSettings UserSettingsType = 1
	UserSettingsTypeFrecencyUserSettings  UserSettingsType = 2
	UserSettingsTypeTestSettings          UserSettingsType = 3
)

// A User stores all data for an individual Discord user.
type User struct {
	// The ID of the user.
	ID string `json:"id"`

	// The email of the user. This is only present when
	// the application possesses the email scope for the user.
	Email string `json:"email"`

	// The user's phone number
	Phone string `json:"phone"`

	// The user's username.
	Username string `json:"username"`

	// The user's display name
	DisplayName string `json:"global_name"`

	// The hash of the user's avatar. Use Session.UserAvatar
	// to retrieve the avatar itself.
	Avatar string `json:"avatar"`

	// The user's chosen language option.
	Locale string `json:"locale"`

	// The discriminator of the user (4 numbers after name).
	Discriminator string `json:"discriminator"`

	// The token of the user. This is only present for
	// the user represented by the current session.
	Token string `json:"token"`

	// Whether the user's email is verified.
	Verified bool `json:"verified"`

	// Whether the user has multi-factor authentication enabled.
	MFAEnabled bool `json:"mfa_enabled"`

	// The hash of the user's banner image.
	Banner string `json:"banner"`

	// User's banner color, encoded as hexadecimal color code
	BannerColor string `json:"banner_color"`

	// User's banner color, encoded as an integer representation of hexadecimal color code
	AccentColor int `json:"accent_color"`

	// Whether the user is a bot.
	Bot bool `json:"bot"`

	// The public flags on a user's account.
	// This is a combination of bit masks; the presence of a certain flag can
	// be checked by performing a bitwise AND between this int and the flag.
	PublicFlags UserFlags `json:"public_flags"`

	// The type of Nitro subscription on a user's account.
	// Only available when the request is authorized via a Bearer token.
	PremiumType UserPremiumType `json:"premium_type"`

	// Whether the user is an Official Discord System user (part of the urgent message system).
	System bool `json:"system"`

	// The flags on a user's account.
	// Only available when the request is authorized via a Bearer token.
	Flags int `json:"flags"`

	// "About Me" section of the user.
	// Only available on client state user.
	Bio string `json:"bio"`

	// Pronouns of the user.
	// Only available on client state user.
	Pronouns string `json:"pronouns"`
}

// String returns a unique identifier of the form username#discriminator
// NOTE: deprecated, use user.Username or user.DisplayName instead
func (u *User) String() string {
	return u.Username + "#" + u.Discriminator
}

// Mention return a string which mentions the user
func (u *User) Mention() string {
	return "<@" + u.ID + ">"
}

// AvatarURL returns a URL to the user's avatar.
//    size:    The size of the user's avatar as a power of two
//             if size is an empty string, no size parameter will
//             be added to the URL.
func (u *User) AvatarURL(size string) string {
	return avatarURL(u.Avatar, EndpointDefaultUserAvatar(u.Discriminator),
		EndpointUserAvatar(u.ID, u.Avatar), EndpointUserAvatarAnimated(u.ID, u.Avatar), size)
}

// BannerURL returns the URL of the users's banner image.
//    size:    The size of the desired banner image as a power of two
//             Image size can be any power of two between 16 and 4096.
func (u *User) BannerURL(size string) string {
	return bannerURL(u.Banner, EndpointUserBanner(u.ID, u.Banner), EndpointUserBannerAnimated(u.ID, u.Banner), size)
}

type Profile struct {
	// The user who owns the profile.
	User *User `json:"user"`

	// The user's connected accounts.
	Connections []*UserConnection `json:"connections"`

	PremiumSince *time.Time `json:"premium_since"`

	BoostingSince *time.Time `json:"premium_guild_since"`

	MutualGuilds []struct {
		ID       string `json:"id"`
		Nickname string `json:"nick"`
	} `json:"mutual_guilds"`

	UserProfile map[string]any `json:"user_profile"`
}
