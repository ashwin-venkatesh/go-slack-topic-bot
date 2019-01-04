# go-slack-topic-bot

Some utilities for making a [Slack](https://slack.com/) bot that manages a channel's topic.


## Usage

Use this programmatically from your own programs. See the ([examples](examples/cfbosh.go)) for how you can declare topics you care about. The following environment variables are used in the examples:

 * `SLACK_CHANNEL` - the channel ID where the topic will be updated (not the name; e.g. `CA1B2C3D4`)
 * `SLACK_TOKEN` - an API token for accessing Slack (you may want to create one [here](https://apps.slack.com/apps/A0F7YS25R-bots))


## License

[MIT License](LICENSE)

## Windows Bots

### Adding the new team member to the bot

In order to add a new person to the team, you will need
to be signed in to the Pivotal and Cloud Foundry Slack orgs. 

Once there, navigate to the user profile for this new team member
and in the url you should see this at the end of the path: `team/USER_ID/`.

For the Pivotal `pcf-windows` slack channel, navigate to
`windows-bots/pcf-windows.go` in this repo. Add the user to the list:
```go
People: map[string]string{
  "Name": "PIVOTAL_USER_ID",
  ...
}
```

For the Cloud Foundry `garden-windows` slack channel, navigate to
`windows-bots/garden-windows.go` in this repo. Add the user to the list:
```go
People: map[string]string{
  "Name": "CLOUD_FOUNDRY_USER_ID",
  ...
}
```

### Redeploying the bot
