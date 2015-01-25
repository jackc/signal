# Signal

Signal is a Go implementation of the signal pattern. It was inspired by [js-signals](https://github.com/millermedeiros/js-signals). It uses [gen](https://github.com/clipperhouse/gen) to generate type safe signal types.

## Usage

See [demo/main.go](demo/main.go) for a simple example. See [demo/event_signal.go](demo/event_signal.go) for and example of the generated signal type.

To add to your project:

    $ gen add github.com/jackc/signal

Add the type annotation to the type you wish to signal with:

    // +gen signal
    type Event struct {
      ID   int
      Type string
    }

Generate the type:

    $ gen

You will now have the file event_signal.go containing type EventSignal has Add, Remove, and Dispatch methods.


