# gotbot

#### A generic telegram bot api client made with go.

### Introduction

This package was a result of my first telegram bot implementation. When I started
that project I planned to only be concerned with my bot's functionality and not
the details of handling requests and responses from the telegram bot api server.
But after browsing for telegram api client packages written with `go`, I wasn't
satisfied with what I found and I finally decided it was better to write one myself.

Since the telegram bot api is very vast, only some of the functionalities are currently
implemented. But a contribution from you would very much be appreciated to minimize this limitation
as much as possible.

### Getting started

in order to start using this package, use the following command to add this package
to your project.

```go
go get github.com/roskee/gotbot
```

once you have successfully installed the package, you can use the methods and models specified
in the package to set up and manage your bot.

### Basics

Assuming you have created your bot on telegram and have
your [token](https://core.telegram.org/bots/api#authorizing-your-bot).
you can use the `NewBot` function to connect your bot.  
[Learn more about creating your first bot](https://core.telegram.org/bots/features#botfather)

```go
apiToken := os.Getenv("API_KEY")
bot := gotbot.NewBot(apiToken)
```

Then you can use the `RegisterMethod` to register your first command.  
[Learn more about commands](https://core.telegram.org/bots/api#setmycommands)

```go
err := bot.RegisterMethod("start", "start command", func(update entity.Update) {
fmt.Println("start command was sent")
// do you thing... maybe reply to the user
})
```

Finally, you have to start your server to actually listen for user interactions with your bot.
There are two mutually exclusive ways of doing this.

- Polling with [getUpdates](https://core.telegram.org/bots/api#getupdates) method or,
- Registering a [webhook](https://core.telegram.org/bots/api#setwebhook)

#### polling

you can call `getUpdates` method periodically to get all user interactions since your last call to it.
your server will indefinitely (with the given interval) ask the telegram bot api for new updates.

```go
err = bot.Poll(5 * time.Second)
```

#### webhook

you can also register a webhook url for the telegram bot api server to call whenever there is a new update.
This method only works with [these security configurations](https://core.telegram.org/bots/webhooks)

```go
webhook := entity.Webhook{
    URL: "https://myurl.com/updatesCallback",
}
err = bot.Listen(webhook)
```

**Best Wishes!**  
*Kirubel Adamu*