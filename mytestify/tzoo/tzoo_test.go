package tzoo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockClient struct{ mock.Mock }

func newMockClient() *mockClient { return &mockClient{} }

func (c *mockClient) ListAnimalFacts(
	q AnimalFactsQuery,
) (*AnimalFactsResponse, error) {
	args := c.Called(q)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*AnimalFactsResponse), args.Error(1)
}

var firstPageAPIReq = AnimalFactsQuery{
	AnimalName: "sloth",
	PageToken:  "",
}

func reqSlothFacts(q AnimalFactsQuery) bool {
	switch strings.ToLower(q.AnimalName) {
	case "sloth",
		"lemur: caffeine-free edition",
		"descendants of the giants who made today's avocados happen":

		return true
	default:
		return false
	}
}

var page1 = AnimalFactsResponse{
	Facts: []string{
		"Sloths' slowness is actually used as a form of camouflage",
		"Baby sloths make the cutest li'l squeak 🥰",
		"Sloths need their beauty sleep, plz send noise-cancelling headphones",
	},
	AreThereMore:  true,
	NextPageToken: "some more facts on the next page",
}

var page2 = AnimalFactsResponse{
	Facts: []string{"Sloths' favorite snack is hibiscus flowers"},
}

func TestGetSlothsFavoriteSnack(t *testing.T) {
	c := newMockClient()
	c.On(
		"ListAnimalFacts", firstPageAPIReq,
	).Return(&AnimalFactsResponse{
		Facts: []string{
			"Sloths' slowness is actually used as a form of camouflage",
			"Baby sloths make the cutest li'l squeak 🥰",
			"Sloths' favorite snack is hibiscus flowers",
		},
		AreThereMore:  false,
		NextPageToken: "",
	}, nil)

	favSnack, err := getSlothsFavoriteSnack(c)
	if err != nil {
		t.Fatalf("got error getting sloths' favorite snack: %v", err)
	}

	if favSnack != "hibiscus flowers" {
		t.Errorf(
			"expected favorite snack to be hibiscus flowers, got %s",
			favSnack,
		)
	}
}

func TestGetSlothsFavoriteSnack500Error(t *testing.T) {
	c := newMockClient()
	c.On("ListAnimalFacts", firstPageAPIReq).Return(
		(*AnimalFactsResponse)(nil), &ErrorResponse{
			StatusCode: 500,
			Message:    "server error",
		},
	)

	_, err := getSlothsFavoriteSnack(c)
	if err == nil {
		t.Fatal("got nil error from getSlothsFavoriteSnack")
	}

	errRes, ok := err.(*ErrorResponse)

	if !ok {
		t.Fatalf("expected error to be ErrorResponse, got %T", err)
	}
	if status := errRes.StatusCode; status != 500 {
		t.Errorf("expected 500, got %d status code", status)
	}
}

func TestGetSlothsFavoriteSnackNotFound(t *testing.T) {
	c := newMockClient()
	c.On(
		"ListAnimalFacts", firstPageAPIReq,
	).Return(&AnimalFactsResponse{
		Facts: []string{
			"Sloths' slowness is actually used as a form of camouflage",
			"Baby sloths make the cutest li'l squeak 🥰",
			"Sloths need their beauty sleep, plz send noise-cancelling headphones",
		},
		AreThereMore:  false,
		NextPageToken: "",
	}, nil)

	if _, err := getSlothsFavoriteSnack(c); err != errFactNotFound {
		t.Fatalf("should have gotten errFactNotFound, got %v", err)
	}
}

func TestGetSlothsFavoriteSnackOnPage2(t *testing.T) {
	c := newMockClient()
	c.On("ListAnimalFacts", mock.MatchedBy(reqSlothFacts)).
		Return(&page1, nil).
		Once()
	c.On("ListAnimalFacts", mock.MatchedBy(reqSlothFacts)).
		Return(&page2, nil).
		Once()

	favSnack, err := getSlothsFavoriteSnack(c)
	if err != nil {
		t.Fatalf("got error getting sloths' favorite snack: %v", err)
	}

	if favSnack != "hibiscus flowers" {
		t.Errorf(
			"expected favorite snack to be hibiscus flowers, got %s",
			favSnack,
		)
	}

	c.AssertNumberOfCalls(t, "ListAnimalFacts", 2)
}
