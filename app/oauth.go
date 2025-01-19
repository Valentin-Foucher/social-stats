package app

import (
	"fmt"

	"github.com/Valentin-Foucher/social-stats/configs"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/strava"
)

func InitializeOauth2Providers(baseUrl string, conf *configs.Configuration) {
	url := fmt.Sprintf("https://%s/api/auth", baseUrl)
	goth.UseProviders(
		// // Use twitterv2 instead of twitter if you only have access to the Essential API Level
		// // the twitter provider uses a v1.1 API that is not available to the Essential Level
		// twitterv2.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s/api/auth/twitterv2/callback", url)),
		// // If you'd like to use authenticate instead of authorize in TwitterV2 provider, use this instead.
		// // twitterv2.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s/api/auth/twitterv2/callback", url)),

		// twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s/api/auth/twitter/callback", url)),
		// // If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// // twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s/api/auth/twitter/callback", url)),

		// tiktok.New(os.Getenv("TIKTOK_KEY"), os.Getenv("TIKTOK_SECRET"), fmt.Sprintf("%s/api/auth/tiktok/callback", url)),
		// facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), fmt.Sprintf("%s/api/auth/facebook/callback", url)),
		// google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), fmt.Sprintf("%s/api/auth/google/callback", url)),
		// gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), fmt.Sprintf("%s/api/auth/gplus/callback", url)),
		// github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s/api/auth/github/callback", url)),
		// spotify.New(os.Getenv("SPOTIFY_KEY"), os.Getenv("SPOTIFY_SECRET"), fmt.Sprintf("%s/api/auth/spotify/callback", url)),
		// linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), fmt.Sprintf("%s/api/auth/linkedin/callback", url)),
		// twitch.New(os.Getenv("TWITCH_KEY"), os.Getenv("TWITCH_SECRET"), fmt.Sprintf("%s/api/auth/twitch/callback", url)),
		// instagram.New(os.Getenv("INSTAGRAM_KEY"), os.Getenv("INSTAGRAM_SECRET"), fmt.Sprintf("%s/api/auth/instagram/callback", url)),
		// amazon.New(os.Getenv("AMAZON_KEY"), os.Getenv("AMAZON_SECRET"), fmt.Sprintf("%s/api/auth/amazon/callback", url)),
		// slack.New(os.Getenv("SLACK_KEY"), os.Getenv("SLACK_SECRET"), fmt.Sprintf("%s/api/auth/slack/callback", url)),
		// // By default paypal production auth urls will be used, please set PAYPAL_ENV=sandbox as environment variable for testing
		// // in sandbox environment
		// paypal.New(os.Getenv("PAYPAL_KEY"), os.Getenv("PAYPAL_SECRET"), fmt.Sprintf("%s/api/auth/paypal/callback", url)),
		// steam.New(os.Getenv("STEAM_KEY"), fmt.Sprintf("%s/api/auth/steam/callback", url)),
		// uber.New(os.Getenv("UBER_KEY"), os.Getenv("UBER_SECRET"), fmt.Sprintf("%s/api/auth/uber/callback", url)),
		// soundcloud.New(os.Getenv("SOUNDCLOUD_KEY"), os.Getenv("SOUNDCLOUD_SECRET"), fmt.Sprintf("%s/api/auth/soundcloud/callback", url)),
		strava.New(conf.Strava.PublicKey, conf.Strava.SecretKey, fmt.Sprintf("%s/strava/callback", url)),
	)
}
