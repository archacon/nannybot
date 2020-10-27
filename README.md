<img align="right" src="http://bwmarrin.github.io/discordgo/img/discordgo.png">

## NannyBot

Nannybot uses Reddit's API to pull reddit posts in predetermined lists of

threads and refines the list to meet a "like" threshold. Nanny bot will post a

random image from the lists when "Get meme" is typed in the discord channel it

has joined

**Join [Discord Gophers](https://discord.gg/0f1SbxBZjYoCtNPP)
Discord chat channel for support.**

### Build

This assumes you already have a working Go environment setup and that
DiscordGo is correctly installed on your system.

From within the pingpong example folder, run the below command to compile the
example.

```sh
go build
```

### Usage

This example uses bot tokens for authentication only. While user/password is
supported by DiscordGo, it is not recommended for bots.

```
./nannybot --help
Usage of ./nannybot:
  -t string
        Bot Token
```

The below example shows how to start the bot

```sh
./nannybot -t YOUR_BOT_TOKEN
Bot is now running.  Press CTRL-C to exit.
```
git