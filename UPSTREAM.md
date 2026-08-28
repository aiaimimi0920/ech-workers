# Upstream Tracking

## Repositories

- Upstream: https://github.com/hhsw2015/ech-workers
- EasyProxy fork: https://github.com/aiaimimi0920/ech-workers
- Root integration path: upstreams/ech-workers

## Baseline

- Audited upstream main: 46c6c71
- EasyProxy integration branch: main

The upstream main branch currently has a different source layout from the Go
helper integrated by EasyProxy. Do not replace the fork tree wholesale during
an upstream sync.

## EasyProxy Delta

- Go helper split into config, ECH, proxy, and WebSocket packages.
- Local HTTP and SOCKS5 listener used by the EasyProxy connector runtime.
- Container and root-monorepo build integration.
- Go toolchain declaration aligned with the EasyProxy build.

## Sync Policy

1. Fetch upstream into a dedicated sync branch.
2. Review source and protocol changes path by path.
3. Port applicable fixes into the EasyProxy layout.
4. Run go test ./... and the root container smoke tests.
5. Update the EasyProxy root submodule pointer only after integration tests pass.

Never force-push an unreviewed upstream tree onto the integration branch.
