# debug-cli

DigiSyn Link Network Debug CLI

**Download latest Binary here:** https://github.com/DigiSynlink/debug-cli/releases/tag/dev

## Quick Example

For more info, run without any args to see help info

### For Windows User Interface name

Use your devices' friendly name. 

```powershell
$ .\debug-cli.exe interface list
time="2022-08-05T00:46:36-07:00" level=info msg="Devices found:"
- 'Ethernet'
- 'Wi-Fi'
- 'Local Area Connection* 1'
- 'Local Area Connection* 2'
- 'Loopback Pseudo-Interface 1'
- 'vEthernet (WSL)'
```

```powershell
.\debug-cli.exe --lvl debug device discover --interface "vEthernet (WSL)"
```

### Find Device on Interface

```bash
./debug-cli device discover --interface en7
```

```bash
INFO[0000] Boardcast address: 255.255.255.255
INFO[0000] Auto-binding to interface: en7
INFO[0000] Binding to address: 169.254.172.166
INFO[0000] Sending boardcast...
INFO[0000] Waiting for a response... 1min Deadline
INFO[0000] Found device: dmx208 at: 169.254.242.216/16
```

### Passively Receive Announcement

```bash
./debug-cli device listen --interface en7
```

```bash
INFO[0000] Listening on interface: en7
INFO[0000] Announce address: 239.255.255.254:9900
WARN[0000] To stop Listen, press Ctrl+C
INFO[0000] Waiting for a response... 1min Deadline
INFO[0005] Device: DMX-2085 frint(169.254.242.216) is online
INFO[0005] Current Online Machines:
INFO[0005] 	DMX-2085 frint: {announce yes DLA04-31801 DMX-2085 frint DLA04 both 4 IN1,IN2,IN3,IN4 Out1,Out2,Out3,Out4 169.254.242.216  254666093}
```