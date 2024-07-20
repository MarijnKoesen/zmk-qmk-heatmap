# ZMK Heatmap Generator

A heatmap generator for ZMK keyboards.

Generating a heatmap is done in two steps:

1) **Extract your keymap**: in order to properly parse combo's as combo's instead of two keystrokes we need to load
   your keymap. 

   Generate your keymap with https://github.com/hnaderi/keymap-drawer/

   ```
   keymap parse -c 10 -z <your-zmk-config-project>/config/corne.keymap > keymap.yaml 
   ```

1) **Process your keystrokes**: this is the process of listening to all of your keystrokes and
   collection the amount each key is pressed and storing this in a file. Since we only store
   the amount each key is pressed it should be safe and no passwords or any other sensitive
   information will be stored or derivable.

   First make sure that `USB_LOGGING` is enabled on your keyboard, see: [#Enable USB logging in your keyboard](Enable USB logging in your keyboard) below.

   Next start the collection by running the command and point to your `keymap.yaml`:
   ```
   zmk-heatmap collect --keymap=keymap.yaml
   ```

1) **Generating the heatmap**: taking all the keystrokes into account the heatmap can now be
   created:
   ```
   zmk-heatmap generate
   ```

1) **Create the keymap config**: copy the output from the `zmk-heatmap generate` and add it in the `svg_extra_style`.

   For an example see: [https://github.com/MarijnKoesen/zmk-config/blob/main/.keymap.config.ferris-sweep.yaml](https://github.com/MarijnKoesen/zmk-config/blob/main/.keymap.config.ferris-sweep.yaml)

   I'm planning to integrate this nicer in the future, but for now it's a bit of a manual copy/paste.

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
