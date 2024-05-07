# hatop
HAProxy Statistics

```
HAProxy Statistics (refreshing every 2 seconds):
PxName | SvName | Scur | Status | LastChk
----------------------------------------------------------------------
front | FRONTEND | 0 | OPEN | 0
servers | app1 | 0 | UP |
servers | app2 | 0 | UP |
servers | app3 | 0 | UP |
servers | app4 | 0 | UP |
servers | app5 | 0 | UP |
servers | app6 | 0 | UP |
servers | app7 | 0 | UP |
servers | app8 | 0 | UP |
servers | BACKEND | 0 | UP |
unauthorized | BACKEND | 0 | UP |
----------------------------------------------------------------------
Queue: 0
```

# Installation

```bash
go build -o hatop
```

# Usage

```bash
sudo ./hatop
```