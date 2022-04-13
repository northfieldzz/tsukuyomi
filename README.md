# Tsukuyomi
私にとっては神様になる  

## development

```shell script
$ go mod download
$ go run tsukuyomi.go
```

### Environment  

- VERSION  
  このAPIのバージョン指定  

- PORT   
  リスニングポート指定  

- GIN_MODE=release  
  デプロイモードの選択　指定なしでdebug  
  
- DATABASE_HOST  
  DATABASE_PORT  
  DATABASE_NAME  
  DATABASE_USER  
  DATABASE_PASSWORD  
  DATABASE_SSLMODE  
  データベース接続関連  

- DISCORD_TOKEN  
  discord api で使用するToken  
  
- LINEBOT_SECRET_KEY  
  LINEBOT_CHANNEL_ACCESS_TOKEN  
  line messaging apiで使用する認証系  
  
- GOOGLE_API_KEY   
  GCP Youtube APIの認証キー  
