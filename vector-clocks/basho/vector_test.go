package easy

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("not equal")
	}
}

func Test_basho_story(t *testing.T) {
	/*
		Alice, Ben, Cathy, and Dave are planning to meet next week for
		dinner.
	*/
	alice := newEvent("alice")
	ben := newEvent("ben")
	cathy := newEvent("cathy")
	dave := newEvent("dave")
	/*
		Alice says, “Let’s meet Wednesday,” and tags that value as the first version of the
		message that she has seen
	*/
	alice.setVal("wednesday")
	send(alice, ben)
	send(alice, dave)
	send(alice, cathy)
	/*
		Now Dave and Ben start talking. Ben suggests Tuesday
	*/
	ben.setVal("tuesday")
	send(ben, dave)
	/*
		Dave replies, confirming Tuesday
	*/
	dave.setVal("tuesday")
	send(dave, ben)
	/*
		Now Cathy gets into the act, suggesting Thursday
	*/
	cathy.setVal("thursday")
	send(cathy, dave)
	/*
		Luckily, Dave’s a reasonable guy, and chooses Thursday
	*/
	dave.setVal("thursday")
	send(dave, cathy)
	/*
		So now when Alice asks Ben and Cathy for the latest decision,
		the replies she receive are, from Ben Tuesday, Cathy Thursday
	*/
	assertEqual(t, ben.value, "tuesday")
	assertEqual(t, cathy.value, "thursday")

	/*
		From this, she can tell that Dave intended his correspondence
		with Cathy to override the decision he made with Ben. All Alice
		has to do is show Ben the vector clock from Cathy’s message,
		and Ben will know that he has been overruled
	*/

	send(cathy, alice)
	send(alice, ben)

	/*
		Consensus on thursday
	*/

	assertEqual(t, alice.value, "thursday")
	assertEqual(t, ben.value, "thursday")
	assertEqual(t, cathy.value, "thursday")
	assertEqual(t, dave.value, "thursday")
}
