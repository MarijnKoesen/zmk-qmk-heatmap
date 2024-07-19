# ZMK Heatmap Generator

A heatmap generator for ZMK keyboards.

Generating a heatmap is done in two steps:

1) **Process your keystrokes**: this is the process of listening to all of your keystrokes and
   collection the amount each key is pressed and storing this in a file. Since we only store
   the amount each key is pressed it should be safe and no passwords or any other sensitive
   information will be stored or derivable.

   First make sure that `USB_LOGGING` is enabled on your keyboard, see: [#Enable USB logging in your keyboard](Enable USB logging in your keyboard) below.

   Next start the collection by running the command pointing at your keyboard:
   ```
   $ zmk-heatmap collect -k /dev/tty.usbmodem142101
   ```

2) **Generating the heatmap**: taking all the keystrokes into account the heatmap can now be
   created:
   ```
   $ zmk-heatmap generate'`,
   ```


## Enable USB logging in your keyboard 

In order to generate the heatmap we need to know which keys are pressed. We cannot do this by listening to the keystrokes alone as we'd loose any information on layers etc.

So this is done by listening to the [`USB_LOGGING`](https://zmk.dev/docs/development/usb-logging) output from your ZMK keyboard.

To enable this option set the following parameters in your `config/*.conf` file:

```
CONFIG_ZMK_USB_LOGGING=y
CONFIG_LOG_BUFFER_SIZE=31768
CONFIG_LOG_PROCESS_THREAD_SLEEP_MS=10
CONFIG_LOG_SPEED=y
```

Next build the firmware and flash your keyboard.


## Example heatmap

![Example Heatmap](/doc/example-heatmap.svg)