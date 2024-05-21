# Boot.dev Buddy

This app provides a GUI to collect and show data from [Boot.dev](https://boot.dev/). It can be used as a companion app while doing coursework to monitor on-going events, buff timers, courses or lessons added since the app was last opened, and much more.

Note: This project isn't affiliated with Boot.dev in any way. I'm just a student that wanted to make something I thought was cool.

## Logging In

Boot.dev uses a one-time password system for logging in via it's CLI. Using this same OTP, BDB can authenticate the user and retrieve their Boot.dev data.

Note: The JWT they use expires after an hour. Currently that will require obtaining another OTP.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Why?

I'm a fan of monitoring software, making it look nice, and making it easy to get the information you want at a glance. I thought having a buddy app to the [Boot.dev](https://boot.dev) curriculum that could give a student more sense of community, while also monitoring your own progress, would be fun.
