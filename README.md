# Hex Plugin - WinRM

Hex Plugin which executes commands on Windows servers via WinRM.

```
{
  "rule": "example winrm rule",
  "match": "C Drive",
  "actions": [
    {
      "type": "hex-winrm",
      "command": "dir C:\",
      "config": {
        "server": "127.0.0.1",
        "port": "5985",
        "login": "hexbot",
        "pass": "${HEXBOT_WINRM_PASSWORD}"
      }
    }
  ]
}
```
