package api

type GetChallengesParams struct {
	Name        *string `schema:"name,omitempty"`
	MaxAttempts *int    `schema:"max_attempts,omitempty"`
	Value       *int    `schema:"value,omitempty"`
	Category    *int    `schema:"category,omitempty"`
	Type        *string `schema:"type,omitempty"`
	State       *string `schema:"state,omitempty"`
	Q           *string `schema:"q,omitempty"`
	// TODO cover "view=admin" parameter that shows hidden challenges
}

func (client *Client) GetChallenges(params *GetChallengesParams, opts ...Option) ([]*Challenge, error) {
	challs := []*Challenge{}
	if err := get(client, "/challenges", params, &challs, opts...); err != nil {
		return nil, err
	}
	return challs, nil
}

type PostChallengesParams struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Initial     int    `json:"initial"`
	Decay       *int   `json:"decay,omitempty"`
	Minimum     *int   `json:"minimum,omitempty"`
	State       string `json:"state"`
	Type        string `json:"type"`
	Anonymize	string `json:"anonymize"`
	ConnectionInfo	string `json:"connectioninfo"`
	MaxAttempts	string `json:"maxattempts"`
	Value		string `json:"value"`
}

func (client *Client) PostChallenges(params *PostChallengesParams, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := post(client, "/challenges", params, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

type PostChallengesAttemptParams struct {
	ChallengeID int    `json:"challenge_id"`
	Submission  string `json:"submission"`
}

func (client *Client) PostChallengesAttempt(params *PostChallengesAttemptParams, opts ...Option) (*Attempt, error) {
	attempt := &Attempt{}
	if err := post(client, "/challenges/attempt", params, &attempt, opts...); err != nil {
		return nil, err
	}
	return attempt, nil
}

func (client *Client) GetChallengesTypes(opts ...Option) (map[string]*Type, error) {
	types := map[string]*Type{}
	if err := get(client, "/challenges/types", nil, &types, opts...); err != nil {
		return nil, err
	}
	return types, nil
}

func (client *Client) GetChallenge(id string, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := get(client, "/challenges/"+id, nil, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

type PatchChallengeParams struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Initial     string `json:"initial"`      // XXX should be int
	Decay       string `json:"decay"`        // XXX should be int
	Minimum     string `json:"minimum"`      // XXX should be int
	MaxAttempts string `json:"max_attempts"` // XXX should be int
	State       string `json:"state"`
}

func (client *Client) DeleteChallenge(id string, opts ...Option) error {
	return delete(client, "/challenges/"+id, nil, nil, opts...)
}

func (client *Client) PatchChallenge(id string, params *PatchChallengeParams, opts ...Option) (*Challenge, error) {
	ch := &Challenge{}
	if err := patch(client, "/challenges/"+id, params, &ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *Client) GetChallengeFiles(id string, opts ...Option) ([]*File, error) {
	cf := []*File{}
	if err := get(client, "/challenges/"+id+"/files", nil, &cf, opts...); err != nil {
		return nil, err
	}
	return cf, nil
}

func (client *Client) GetChallengeFlags(id string, opts ...Option) ([]*Flag, error) {
	cf := []*Flag{}
	if err := get(client, "/challenges/"+id+"/flags", nil, &cf, opts...); err != nil {
		return nil, err
	}
	return cf, nil
}

func (client *Client) GetChallengeHints(id string, opts ...Option) ([]*Hint, error) {
	ch := []*Hint{}
	if err := get(client, "/challenges/"+id+"/hints", nil, &ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *Client) GetChallengeRequirements(id string, opts ...Option) (*Requirements, error) {
	req := &Requirements{}
	if err := get(client, "/challenges/"+id+"/requirements", nil, &req, opts...); err != nil {
		return nil, err
	}
	return req, nil
}

// TODO find content to determine model
func (client *Client) GetChallengeSolves(id string, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := get(client, "/challenges/"+id+"/solves", nil, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

func (client *Client) GetChallengeTags(id string, opts ...Option) ([]*Tag, error) {
	ct := []*Tag{}
	if err := get(client, "/challenges/"+id+"/tags", nil, &ct, opts...); err != nil {
		return nil, err
	}
	return ct, nil
}

func (client *Client) GetChallengeTopics(id string, opts ...Option) ([]*Topic, error) {
	ct := []*Topic{}
	if err := get(client, "/challenges/"+id+"/topics", nil, &ct, opts...); err != nil {
		return nil, err
	}
	return ct, nil
}
