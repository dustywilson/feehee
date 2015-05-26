package feehee

import "testing"

func TestEmail(t *testing.T) {
	startingGoodAddress := `dusty@scj.io`
	goodAddresses := []string{
		`dustyw@scjalliance.com`,
		`dusty@linux.com`,
		`dusty+test@linux.com`,
	}
	email, err := NewEmailAddress(startingGoodAddress)
	if err != nil {
		t.Errorf("%s should have been a good address, but wasn't", startingGoodAddress)
	}
	lastEmailAddress := email.String()
	for _, goodAddress := range goodAddresses {
		_, err := NewEmailAddress(goodAddress)
		if err != nil {
			t.Errorf("%s should have been a good address, but wasn't", goodAddress)
			continue
		}
		err = email.Update(goodAddress)
		if err != nil {
			t.Errorf("%s should have been a good address, but wasn't", goodAddress)
			continue
		}
		if email.String() == lastEmailAddress {
			t.Errorf("%s should have been an accepted update, but wasn't", goodAddress)
			continue
		}
		lastEmailAddress = email.String()
	}

	badAddresses := []string{
		`dustyw@scj@alliance.com`, // too many @ symbols
		`dusty@linu.x.com`,        // no MX
		`dusty@example.com`,       // no MX
		`dusty@examplecom`,        // no dot in domain
		`dustyexample.com`,        // no @ symbol
		`dusty.example.com`,       // no @ symbol
	}
	for _, badAddress := range badAddresses {
		_, err := NewEmailAddress(badAddress)
		if err == nil {
			t.Errorf("%s should have been a bad address, but wasn't", badAddress)
			continue
		}
		err = email.Update(badAddress)
		if err == nil {
			t.Errorf("%s should have been a bad address, but wasn't", badAddress)
			continue
		}
		if email.String() != lastEmailAddress {
			t.Errorf("%s should not have been an accepted update", badAddress)
			continue
		}
		lastEmailAddress = email.String()
	}
}
