package dgrs

import "github.com/bwmarrin/discordgo"

// SetUser sets the given user object to the cache.
func (s *State) SetUser(user *discordgo.User) (err error) {
	err = s.set(joinKeys(KeyUser, user.ID), user, s.getLifetime(user))
	return
}

// User tries to retrieve a user by the given user ID.
//
// If no user was found and FetchAndStore is enabled, the object
// will be tried to be retrieved from the API. When this was successful,
// it is stored in the cache and the object is returned.
//
// Otherwise, if the object was not found in the cache and FetchAndStore
// is disabled, nil is returned.
func (s *State) User(id string) (v *discordgo.User, err error) {
	v = &discordgo.User{}
	ok, err := s.get(joinKeys(KeyUser, id), v)
	if !ok {
		if s.options.FetchAndStore {
			if v, err = s.session.User(id); v != nil && err == nil {
				err = s.SetUser(v)
			}
		} else {
			v = nil
		}
	}
	return
}

// Users returns a list of users which are stored
// in the cache at the given moment.
func (s *State) Users() (v []*discordgo.User, err error) {
	v = make([]*discordgo.User, 0)
	err = s.list(joinKeys(KeyUser, "*"), &v)
	return
}

// RemoveUser removes a user object from the cache by the given ID.
func (s *State) RemoveUser(id string) (err error) {
	return s.del(joinKeys(KeyUser, id))
}

// SelfUser returns the current user object of the authenticated
// account.
//
// This object is retrieved on receiving the 'Ready' event.
func (s *State) SelfUser() (v *discordgo.User, err error) {
	return s.User("@me")
}

// SetSelfUser allows to set a custom user object as self
// user to the cache.
func (s *State) SetSelfUser(user *discordgo.User) (err error) {
	err = s.set(joinKeys(KeyUser, "@me"), user, s.getLifetime(user))
	return
}
