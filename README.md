# Boot.dev Buddy

[![CI](https://github.com/ellielle/bootdev-buddy/actions/workflows/ci.yml/badge.svg)](https://github.com/ellielle/bootdev-buddy/actions/workflows/ci.yml)

This app provides a GUI to collect and show data from [Boot.dev](https://boot.dev/). It can be used as a companion app while doing coursework to monitor on-going events, their buffs, and much more.

<details>
	
<summary>	
Early UI during a boss fight
</summary
	
![2024-06-10_16-16-26](https://github.com/ellielle/bootdev-buddy/assets/40385743/42309740-4705-4183-8a3d-8d20a33d7297)
 
</details>

> Note: This project isn't affiliated with Boot.dev in any way. I'm just a student that wanted to make something I thought was cool.

## Logging In

Boot.dev uses a one-time password system for logging in via it's CLI. Using this same OTP, BDB can authenticate the user and retrieve their Boot.dev data.

Access tokens are only valid for 60 minutes. Once invalid, the app will attempt to renew the token. If it fails, you may need get another OTP and sign in again. If that fails, restart the app.

> This app creates a file in its working directory to store your current tokens called `.bootdevbuddy.json`. If you have issues with sign-in or are removing the app, delete this file.

1. In the app, click on the text that says to `click here` to get your one-time password.

2. In the new tab that opens, click the 📋(clipboard) icon to copy your OTP

3. Paste the token into the awaiting box, and click the `Sign in` button.

## Planned Features

- [x] Sign in using Boot.dev's authentication
- [x] In-memory cache of requests made
- [x] Automatic sign-in on load
- [-] View a tally of lessons completed, and see how many of each course are done
- [ ] Fancier and more formatted stats
- [ ] View other mages' profiles on leaderboard by click
- [ ] Buff timers
- [ ] Boss event monitoring / participation
- [ ] More(?)


## Live Development

If you want to use dev mode in the browser and not with a GUI app, set `StartHidden` to `true` in `main.go`.

```go
func main() {
	app := NewApp()
	err := wails.Run(&options.App{
        ...
		StartHidden: true,  // Me
        ...
    }
}
```

1. Run `wails dev` to start up the dev server. This will build your Go backend and Svelte frontend, and start a server at `http://localhost:34115`

2. `app.go` contains App methods that will be callable in the Svelte app. Outside of that, the separation is pretty apparent.

## Building

To build a redistributable, production mode package, use `wails build`.

## Why?

I'm a fan of monitoring software, making it look nice, and making it easy to get the information you want at a glance. I thought having a buddy app to the [Boot.dev](https://boot.dev) curriculum that could give a student more sense of community, while also monitoring your own progress, would be fun.
