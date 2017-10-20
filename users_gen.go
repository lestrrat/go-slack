package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = objects.EpochTime(0)

// UsersDeletePhotoCall is created by UsersService.DeletePhoto method call
type UsersDeletePhotoCall struct {
	service *UsersService
}

// UsersGetPresenceCall is created by UsersService.GetPresence method call
type UsersGetPresenceCall struct {
	service *UsersService
	user    string
}

// UsersIdentityCall is created by UsersService.Identity method call
type UsersIdentityCall struct {
	service *UsersService
}

// UsersInfoCall is created by UsersService.Info method call
type UsersInfoCall struct {
	service       *UsersService
	includeLocale bool
	user          string
}

// UsersListCall is created by UsersService.List method call
type UsersListCall struct {
	service       *UsersService
	includeLocale bool
	limit         int
	presence      bool
}

// UsersSetActiveCall is created by UsersService.SetActive method call
type UsersSetActiveCall struct {
	service *UsersService
}

// UsersSetPresenceCall is created by UsersService.SetPresence method call
type UsersSetPresenceCall struct {
	service  *UsersService
	presence string
}

// DeletePhoto creates a UsersDeletePhotoCall object in preparation for accessing the users.deletePhoto endpoint
func (s *UsersService) DeletePhoto() *UsersDeletePhotoCall {
	var call UsersDeletePhotoCall
	call.service = s
	return &call
}

// Values returns the UsersDeletePhotoCall object as url.Values
func (c *UsersDeletePhotoCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

// Do executes the call to access users.deletePhoto endpoint
func (c *UsersDeletePhotoCall) Do(ctx context.Context) error {
	const endpoint = "users.deletePhoto"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		SlackResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to users.deletePhoto`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// GetPresence creates a UsersGetPresenceCall object in preparation for accessing the users.getPresence endpoint
func (s *UsersService) GetPresence(user string) *UsersGetPresenceCall {
	var call UsersGetPresenceCall
	call.service = s
	call.user = user
	return &call
}

// Values returns the UsersGetPresenceCall object as url.Values
func (c *UsersGetPresenceCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.user) <= 0 {
		return nil, errors.New(`missing required parameter user`)
	}
	v.Set("user", c.user)
	return v, nil
}

// Do executes the call to access users.getPresence endpoint
func (c *UsersGetPresenceCall) Do(ctx context.Context) (*objects.UserPresence, error) {
	const endpoint = "users.getPresence"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*objects.UserPresence
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.getPresence`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserPresence, nil
}

// Identity creates a UsersIdentityCall object in preparation for accessing the users.identity endpoint
func (s *UsersService) Identity() *UsersIdentityCall {
	var call UsersIdentityCall
	call.service = s
	return &call
}

// Values returns the UsersIdentityCall object as url.Values
func (c *UsersIdentityCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

// Do executes the call to access users.identity endpoint
func (c *UsersIdentityCall) Do(ctx context.Context) (*objects.UserProfile, *objects.Team, error) {
	const endpoint = "users.identity"
	v, err := c.Values()
	if err != nil {
		return nil, nil, err
	}
	var res struct {
		SlackResponse
		*objects.UserProfile `json:"user"`
		*objects.Team        `json:"team"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, nil, errors.Wrap(err, `failed to post to users.identity`)
	}
	if !res.OK {
		return nil, nil, errors.New(res.Error.String())
	}

	return res.UserProfile, res.Team, nil
}

// Info creates a UsersInfoCall object in preparation for accessing the users.info endpoint
func (s *UsersService) Info(user string) *UsersInfoCall {
	var call UsersInfoCall
	call.service = s
	call.user = user
	return &call
}

// IncludeLocale sets the value for optional includeLocale parameter
func (c *UsersInfoCall) IncludeLocale(includeLocale bool) *UsersInfoCall {
	c.includeLocale = includeLocale
	return c
}

// Values returns the UsersInfoCall object as url.Values
func (c *UsersInfoCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeLocale {
		v.Set("include_locale", "true")
	}

	if len(c.user) <= 0 {
		return nil, errors.New(`missing required parameter user`)
	}
	v.Set("user", c.user)
	return v, nil
}

// Do executes the call to access users.info endpoint
func (c *UsersInfoCall) Do(ctx context.Context) (*objects.User, error) {
	const endpoint = "users.info"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*objects.User
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.info`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.User, nil
}

// List creates a UsersListCall object in preparation for accessing the users.list endpoint
func (s *UsersService) List() *UsersListCall {
	var call UsersListCall
	call.service = s
	return &call
}

// IncludeLocale sets the value for optional includeLocale parameter
func (c *UsersListCall) IncludeLocale(includeLocale bool) *UsersListCall {
	c.includeLocale = includeLocale
	return c
}

// Limit sets the value for optional limit parameter
func (c *UsersListCall) Limit(limit int) *UsersListCall {
	c.limit = limit
	return c
}

// Presence sets the value for optional presence parameter
func (c *UsersListCall) Presence(presence bool) *UsersListCall {
	c.presence = presence
	return c
}

// Values returns the UsersListCall object as url.Values
func (c *UsersListCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeLocale {
		v.Set("include_locale", "true")
	}

	if c.limit > 0 {
		v.Set("limit", strconv.Itoa(c.limit))
	}

	if c.presence {
		v.Set("presence", "true")
	}
	return v, nil
}

// Do executes the call to access users.list endpoint
func (c *UsersListCall) Do(ctx context.Context) (objects.UserList, error) {
	const endpoint = "users.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		objects.UserList `json:"members"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.list`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserList, nil
}

// SetActive creates a UsersSetActiveCall object in preparation for accessing the users.setActive endpoint
func (s *UsersService) SetActive() *UsersSetActiveCall {
	var call UsersSetActiveCall
	call.service = s
	return &call
}

// Values returns the UsersSetActiveCall object as url.Values
func (c *UsersSetActiveCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

// Do executes the call to access users.setActive endpoint
func (c *UsersSetActiveCall) Do(ctx context.Context) error {
	const endpoint = "users.setActive"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		SlackResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to users.setActive`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// SetPresence creates a UsersSetPresenceCall object in preparation for accessing the users.setPresence endpoint
func (s *UsersService) SetPresence(presence string) *UsersSetPresenceCall {
	var call UsersSetPresenceCall
	call.service = s
	call.presence = presence
	return &call
}

// Values returns the UsersSetPresenceCall object as url.Values
func (c *UsersSetPresenceCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.presence) <= 0 {
		return nil, errors.New(`missing required parameter presence`)
	}
	v.Set("presence", c.presence)
	return v, nil
}

// Do executes the call to access users.setPresence endpoint
func (c *UsersSetPresenceCall) Do(ctx context.Context) error {
	const endpoint = "users.setPresence"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		SlackResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to users.setPresence`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}