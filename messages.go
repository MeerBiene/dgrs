package dgrs

import "github.com/bwmarrin/discordgo"

// SetMessage sets the given message object to the cache.
func (s *State) SetMessage(message *discordgo.Message) (err error) {
	err = s.set(joinKeys(KeyMessage, message.ChannelID, message.ID), message, s.getLifetime(message))
	return
}

// Message tries to retrieve a message by the given channel and message ID.
//
// If no message was found and FetchAndStore is enabled, the object
// will be tried to be retrieved from the API. When this was successful,
// it is stored in the cache and the object is returned.
//
// Otherwise, if the object was not found in the cache and FetchAndStore
// is disabled, nil is returned.
func (s *State) Message(channelID, messageID string) (v *discordgo.Message, err error) {
	v = &discordgo.Message{}
	ok, err := s.get(joinKeys(KeyMessage, channelID, messageID), v)
	if !ok {
		if s.options.FetchAndStore {
			if v, err = s.session.ChannelMessage(channelID, messageID); v != nil && err == nil {
				err = s.SetMessage(v)
			}
		} else {
			v = nil
		}
	}
	return
}

// Messages returns a list of messages of the given channel
// which are stored in the cache at the given moment.
func (s *State) Messages(channelID string, forceFetch ...bool) (v []*discordgo.Message, err error) {
	v = make([]*discordgo.Message, 0)
	if err = s.list(joinKeys(KeyMessage, channelID, "*"), &v); err != nil {
		return
	}

	if (len(v) == 0 || optBool(forceFetch)) && s.options.FetchAndStore {
		var last string
		var ms []*discordgo.Message
		for ms == nil || len(ms) > 0 {
			if ms != nil {
				last = ms[len(ms)-1].ID
			}
			if ms, err = s.session.ChannelMessages(channelID, 100, last, "", ""); err != nil {
				return
			}
			v = append(v, ms...)
			for _, m := range ms {
				if err = s.SetMessage(m); err != nil {
					return
				}
			}
		}
	}

	return
}

// RemoveMessage removes a message object from the cache by
// the given channel and message ID.
func (s *State) RemoveMessage(channelID, messageID string) (err error) {
	return s.del(joinKeys(KeyMessage, channelID, messageID))
}
