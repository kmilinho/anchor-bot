package twapi

type TwUser struct {
	ScreenName string
}

type Tweet struct {
	Text string
}

func (*TwUser) ShowTweets () []Tweet{
	return nil
}