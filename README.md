Pravah.io
=========

### Development

```
const ActivationThresh = 1
```

should be set in `protocol/identify/obsaddr.go`. This is a trust threshold. Observed addr is only added if `ActivationThresh` number of peers report `this` address within `ttl`.