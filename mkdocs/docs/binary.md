# Obtaining The Binary

In order to install/use the HelperNode, you'll need to obtain the binary by visiting the [releases page](https://github.com/RedHatOfficial/ocp4-helpernode/releases) on GitHub.

Current release is `v2beta1`

```shell
wget https://github.com/RedHatOfficial/ocp4-helpernode/releases/download/v2beta1/helpernodectl
```

Copy the binary into your `$PATH`, and make it executable.

```shell
sudo cp helpernodectl /usr/local/bin/
sudo chmod +x /usr/local/bin/helpernodectl
```

# Global Flags

Currently there are two global flags: `--config` and `--log-level`

```shell
Global Flags:
  -c, --config string      config file (default is $HOME/.helpernodectl.yaml)
      --log-level string   log level (e.g. "debug | info | warn | error") (default "info")
```
