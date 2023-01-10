package tzoo

import (
	"errors"
	"regexp"
)

var errFactNotFound = errors.New("fact not found")

var favSnackMatcher = regexp.MustCompile(
	"favorite snack is (.*)",
)

func getSlothsFavoriteSnack(c ZooClient) (string, error) {
	// Until we're at the last page of facts, call ListAnimalFacts
	// with the current page token to paginate through the list,
	// exiting when either the response's AreThereMore field is
	// false, or we find out what sloths' favorite snack is.
	var pageToken string
	for {
		res, err := c.ListAnimalFacts(AnimalFactsQuery{
			AnimalName: "sloth",
			PageToken:  pageToken,
		})
		if err != nil {
			return "", err
		}

		// check if any facts match the "favorite snack is"
		// regex and if so, return the match
		for _, f := range res.Facts {
			match := favSnackMatcher.FindStringSubmatch(f)
			if len(match) < 2 {
				continue
			}
			return match[1], nil
		}

		// check the response to see if there are any more
		// pages of facts about sloths
		pageToken = res.NextPageToken

		if !res.AreThereMore {
			break
		}
	}

	// otherwise if the fact about sloths' favorite snack
	// isn't in the zoo API, return errFactNotFound.
	return "", errFactNotFound
}
