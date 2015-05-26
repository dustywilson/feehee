package feehee

import "testing"

func TestPerson(t *testing.T) {
	startingGoodUsername := `dustyw`
	goodUsernames := []string{
		`dusty`,
		`dwilson`,
		`dustywilson`,
		`d`,
		`dustywilsondustywilsondustywilsondustywilsondustywilsondustywilson`,
	}
	username, err := NewUsername(startingGoodUsername)
	if err != nil {
		t.Errorf("%s should have been a good username, but wasn't", startingGoodUsername)
	}
	lastUsername := username.String()
	for _, goodUsername := range goodUsernames {
		_, err := NewUsername(goodUsername)
		if err != nil {
			t.Errorf("%s should have been a good username, but wasn't", goodUsername)
			continue
		}
		err = username.Update(goodUsername)
		if err != nil {
			t.Errorf("%s should have been a good username, but wasn't", goodUsername)
			continue
		}
		if username.String() == lastUsername {
			t.Errorf("%s should have been an accepted update, but wasn't", goodUsername)
			continue
		}
		lastUsername = username.String()
	}

	badAddresses := []string{
		`dustyw@scjalliance.com`, // no @ symbol permitted
		`dusty wilson`,           // no spaces permitted
		`dusty!wilson`,           // no exclamation mark permitted
	}
	for _, badAddress := range badAddresses {
		_, err := NewUsername(badAddress)
		if err == nil {
			t.Errorf("%s should have been a bad username, but wasn't", badAddress)
			continue
		}
		err = username.Update(badAddress)
		if err == nil {
			t.Errorf("%s should have been a bad username, but wasn't", badAddress)
			continue
		}
		if username.String() != lastUsername {
			t.Errorf("%s should not have been an accepted update", badAddress)
			continue
		}
		lastUsername = username.String()
	}
}
