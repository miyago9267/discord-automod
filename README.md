# Discord Automod

作用於方便伺服器管理員進行自動化偵測非法行為並實施處置的模版。

## 使用

### 安裝

```bash
git clone https://github.com/miyago/discord-automod.git
go mod tidy
```

### 執行

```go
go run ./cmd/main.go
```

### 設定

複製`.env.example`檔案為`.env`，並填入必要的環境變數。

## 功能

- 自動timeout那些很喜歡 `@everyone` 的人

## TODO

- 自動化偵測非法行為
