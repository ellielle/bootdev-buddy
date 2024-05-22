# Boot.dev Buddy

This app provides a GUI to collect and show data from [Boot.dev](https://boot.dev/). It can be used as a companion app while doing coursework to monitor on-going events, buff timers, courses or lessons added since the app was last opened, and much more.

Note: This project isn't affiliated with Boot.dev in any way. I'm just a student that wanted to make something I thought was cool.

## Logging In

Boot.dev uses a one-time password system for logging in via it's CLI. Using this same OTP, BDB can authenticate the user and retrieve their Boot.dev data.

> Currently logging into the app is session-based. You'll need another OTP each time you open it.

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
