// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"
)

// DiscordSession is an autogenerated mock type for the DiscordSession type
type DiscordSession struct {
	mock.Mock
}

// AddHandler provides a mock function with given fields: _a0
func (_m *DiscordSession) AddHandler(_a0 interface{}) func() {
	ret := _m.Called(_a0)

	var r0 func()
	if rf, ok := ret.Get(0).(func(interface{}) func()); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}

// Channel provides a mock function with given fields: channelID
func (_m *DiscordSession) Channel(channelID string) (*discordgo.Channel, error) {
	ret := _m.Called(channelID)

	var r0 *discordgo.Channel
	if rf, ok := ret.Get(0).(func(string) *discordgo.Channel); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelMessage provides a mock function with given fields: channelID, messageID
func (_m *DiscordSession) ChannelMessage(channelID string, messageID string) (*discordgo.Message, error) {
	ret := _m.Called(channelID, messageID)

	var r0 *discordgo.Message
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.Message); ok {
		r0 = rf(channelID, messageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(channelID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelMessages provides a mock function with given fields: channelID, limit, beforeID, afterID, aroundID
func (_m *DiscordSession) ChannelMessages(channelID string, limit int, beforeID string, afterID string, aroundID string) ([]*discordgo.Message, error) {
	ret := _m.Called(channelID, limit, beforeID, afterID, aroundID)

	var r0 []*discordgo.Message
	if rf, ok := ret.Get(0).(func(string, int, string, string, string) []*discordgo.Message); ok {
		r0 = rf(channelID, limit, beforeID, afterID, aroundID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, string, string, string) error); ok {
		r1 = rf(channelID, limit, beforeID, afterID, aroundID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Guild provides a mock function with given fields: guildID
func (_m *DiscordSession) Guild(guildID string) (*discordgo.Guild, error) {
	ret := _m.Called(guildID)

	var r0 *discordgo.Guild
	if rf, ok := ret.Get(0).(func(string) *discordgo.Guild); ok {
		r0 = rf(guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Guild)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GuildChannels provides a mock function with given fields: guildID
func (_m *DiscordSession) GuildChannels(guildID string) ([]*discordgo.Channel, error) {
	ret := _m.Called(guildID)

	var r0 []*discordgo.Channel
	if rf, ok := ret.Get(0).(func(string) []*discordgo.Channel); ok {
		r0 = rf(guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GuildEmojis provides a mock function with given fields: guildID
func (_m *DiscordSession) GuildEmojis(guildID string) ([]*discordgo.Emoji, error) {
	ret := _m.Called(guildID)

	var r0 []*discordgo.Emoji
	if rf, ok := ret.Get(0).(func(string) []*discordgo.Emoji); ok {
		r0 = rf(guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GuildMember provides a mock function with given fields: guildID, memberID
func (_m *DiscordSession) GuildMember(guildID string, memberID string) (*discordgo.Member, error) {
	ret := _m.Called(guildID, memberID)

	var r0 *discordgo.Member
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.Member); ok {
		r0 = rf(guildID, memberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(guildID, memberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GuildMembers provides a mock function with given fields: guildID, after, limit
func (_m *DiscordSession) GuildMembers(guildID string, after string, limit int) ([]*discordgo.Member, error) {
	ret := _m.Called(guildID, after, limit)

	var r0 []*discordgo.Member
	if rf, ok := ret.Get(0).(func(string, string, int) []*discordgo.Member); ok {
		r0 = rf(guildID, after, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(guildID, after, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GuildRoles provides a mock function with given fields: guildID
func (_m *DiscordSession) GuildRoles(guildID string) ([]*discordgo.Role, error) {
	ret := _m.Called(guildID)

	var r0 []*discordgo.Role
	if rf, ok := ret.Get(0).(func(string) []*discordgo.Role); ok {
		r0 = rf(guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User provides a mock function with given fields: userID
func (_m *DiscordSession) User(userID string) (*discordgo.User, error) {
	ret := _m.Called(userID)

	var r0 *discordgo.User
	if rf, ok := ret.Get(0).(func(string) *discordgo.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
