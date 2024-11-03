## 0.2.0 (Unreleased)

### :dependabot: **Dependencies**

* deps: bumps actions/download-artifact from 4.1.7 to 4.1.8 (GH-47)
* deps: bumps dependabot/fetch-metadata from 2.1.0 to 2.2.0 (GH-46)
* deps: bumps github.com/creativeprojects/go-selfupdate from 1.2.0 to 1.4.0 (GH-49)
* deps: bumps github.com/go-resty/resty/v2 from 2.13.1 to 2.15.3 (GH-50)
* deps: bumps github.com/redis/go-redis/v9 from 9.5.1 to 9.7.0 (GH-51)
* deps: bumps github.com/spf13/cobra from 1.8.0 to 1.8.1 (GH-44)
* deps: bumps github.com/spf13/viper from 1.18.2 to 1.19.0 (GH-41)
* deps: bumps goreleaser/goreleaser-action from 5 to 6 (GH-43)

## 0.1.0 (May 20, 2024)

### :rocket: **New Features**

* `api` - Add new endpoint `/namespaces/links` to get links from all namespaces. (GH-32)
* `glctl` - Add new flag `-A` or `--all-namespaces` in command `get link` to get links from all namespaces. (GH-32)
* `sdk` - Add new method `GetLinksAllNamespace` to get links from all namespaces. (GH-32)
* `server/short` - Now return 404 HTML page when not found if webbrowser is used or return 404 JSON when CURL is used. (GH-19)

### :dependabot: **Dependencies**

* deps: bumps actions/cache from 3 to 4 (GH-13)
* deps: bumps actions/download-artifact from 4.1.1 to 4.1.7 (GH-35)
* deps: bumps actions/setup-go from 4 to 5 (GH-14)
* deps: bumps dependabot/fetch-metadata from 1.6.0 to 2.1.0 (GH-36)
* deps: bumps github.com/creativeprojects/go-selfupdate from 1.1.3 to 1.2.0 (GH-37)
* deps: bumps github.com/go-resty/resty/v2 from 2.10.0 to 2.13.1 (GH-39)
* deps: bumps github.com/google/uuid from 1.5.0 to 1.6.0 (GH-21)
* deps: bumps github.com/labstack/echo/v4 from 4.11.4 to 4.12.0 (GH-31)
* deps: bumps github.com/redis/go-redis/v9 from 9.3.1 to 9.5.1 (GH-40)
* deps: bumps go.etcd.io/bbolt from 1.3.8 to 1.3.10 (GH-38)
* deps: bumps golangci/golangci-lint-action from 3 to 6 (GH-34)
* deps: bumps peter-evans/repository-dispatch from 2 to 3 (GH-22)

## 0.0.19 (January 20, 2024)

### :rocket: **New Features**

* `docs/api` - Add complete API documentation and swagger generation. (GH-10)

### :tada: **Improvements**

* `glctl/update` - Now `glctl` can update itself. Run `glctl update` to update to the latest version. (GH-9)

## 0.0.18 (January 13, 2024)
