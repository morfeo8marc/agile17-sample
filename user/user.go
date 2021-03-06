package user

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/pkg/errors"
)

// User represents a user
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   uint32 `json:"age"`
	Title string `json:"title"`
}

// NewFromJSON returns a new user from the deserialized json, or an error
func NewFromJSON(s []byte) (*User, error) {
	u := &User{}
	err := json.Unmarshal(s, u)
	return u, err
}

// NewUser represents an updated User
type NewUser struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Level        uint32 `jon:"level"`
	BusinessUnit string `json:"unit"`
}

// NewUserFromUser transforms an old style user into a new style user
func NewUserFromUser(u *User) (*NewUser, error) {
	newU := &NewUser{
		ID:   u.ID,
		Name: u.Name,
	}
	departments := map[string][]string{
		"engineering": []string{"engineering", "developer", "tester"},
		"executive":   []string{"cfo", "cto", "ceo"},
		"operations":  []string{"sre", "dba", "administrator"},
	}
	levels := map[uint32][]string{
		0: []string{"junior", "entry-level"},
		1: []string{"mid-level"},
		2: []string{"senior"},
		3: []string{"principle", "lead"},
	}
	dep, err := lookupKey(departments, u.Title)
	if err != nil {
		return nil, errors.Wrap(err, "could not get department")
	}
	newU.BusinessUnit = dep

	newU.Level = lookupIntKey(levels, u.Title)
	return newU, nil
}

func lookupKey(kv map[string][]string, s string) (string, error) {
	s = strings.ToLower(s)
	words := strings.Split(s, " ")
	for _, w := range words {
		for k, val := range kv {
			for _, v := range val {
				if w == v {
					return k, nil
				}
			}
		}
	}
	return "", errors.New("key not found for " + s)
}

func lookupIntKey(kv map[uint32][]string, s string) uint32 {
	s = strings.ToLower(s)
	words := strings.Split(s, " ")
	for _, w := range words {
		for k, val := range kv {
			for _, v := range val {
				if w == v {
					return k
				}
			}
		}
	}
	return 0
}

func RandomUser() *User {
	return &User{
		ID:    randomString(5),
		Name:  randomString(10),
		Age:   uint32(rand.Intn(64) + 18),
		Title: randomTitle(),
	}
}

func RandomNewUser() *NewUser {
	units := []string{"engineering", "executive", "operations"}
	return &NewUser{
		ID:           randomString(5),
		Name:         randomString(10),
		Level:        uint32(rand.Intn(4)),
		BusinessUnit: units[rand.Intn(len(units))],
	}
}

func randomString(size int) string {
	const dict = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const dictSize = len(dict)
	out := make([]byte, size)
	for i := 0; i < size; i++ {
		out[i] = dict[rand.Intn(dictSize)]
	}
	return string(out)
}

func randomTitle() string {
	jobs := []string{
		"engineering",
		"developer",
		"tester",
		"cfo",
		"cto",
		"ceo",
		"sre",
		"dba",
		"administrator",
	}
	levels := []string{
		"",
		"junior",
		"entry-level",
		"mid-level",
		"senior",
		"principle",
		"lead",
	}
	jobsLen := len(jobs)
	levelLen := len(levels)
	return fmt.Sprintf(
		"%s %s",
		levels[rand.Intn(levelLen)],
		jobs[rand.Intn(jobsLen)],
	)
}
