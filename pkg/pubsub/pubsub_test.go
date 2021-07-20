// +build ignore unit

package pubsub

import (
	"testing"

	"github.com/helmwave/helmwave/pkg/release/uniqname"
	"github.com/stretchr/testify/suite"
)

type PubSubTestSuite struct {
	suite.Suite
	ps *ReleasePubSub
}

func (s *PubSubTestSuite) SetupTest() {
	s.ps = NewReleasePubSub()
}

func (s *PubSubTestSuite) TestPublishFalied() {
	release := uniqname.UniqName("blabla1")
	ch := s.ps.Subscribe(release)

	s.ps.PublishFailed(release)

	status := <-ch
	s.Equal(ReleaseFailed, status)
}

func (s *PubSubTestSuite) TestPublishSuccess() {
	release := uniqname.UniqName("blabla2")
	ch := s.ps.Subscribe(release)

	s.ps.PublishSuccess(release)

	status := <-ch
	s.Equal(ReleaseSuccess, status)
}

func (s *PubSubTestSuite) TestSubscribe() {
	release := uniqname.UniqName("blabla3")
	ch1 := s.ps.Subscribe(release)

	s.NotNil(s.ps.subs)
	s.NotNil(s.ps.subs[release])
	s.Len(s.ps.subs[release], 1)
	s.EqualValues(ch1, (<-chan ReleaseStatus)(s.ps.subs[release][0]))

	ch2 := s.ps.Subscribe(release)
	s.Len(s.ps.subs[release], 2)
	s.EqualValues(ch2, (<-chan ReleaseStatus)(s.ps.subs[release][1]))
}

func TestPubSubTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(PubSubTestSuite))
}