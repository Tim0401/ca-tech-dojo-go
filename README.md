# ca-tech-dojo-go

https://techbowl.co.jp/techtrain/missions/12  
https://github.com/CyberAgentHack/techtrain-mission  

## 概要ブログ記事

https://dreamer.cyou/articles/5fd78d59551c83001c0c4197  

## API仕様

https://github.com/CyberAgentHack/techtrain-mission/blob/master/api-document.yaml  
https://github.com/CyberAgentHack/techtrain-mission  の内容が基本なので、SwaggerEditorで確認しながらPostman等で叩く。  
https://editor.swagger.io/  

### 追加仕様

ガチャで取得したキャラに応じてランキング機能を実装。  
取得用データはRedisで保持し、goのコマンドを叩いてデータを更新する。(コマンドは下の項に記載)  

|url|説明|param|
|--|--|--|
|/ranking/user?$Top=5&$Skip=0|ランキングを取得する|$Top:取得数(デフォルト:100) $Skip:取得をスキップする行数(デフォルト:0 → 1位から取得)|

### 実行

```bash
docker-compose up -d   // Dockerコンテナ群起動
bash init-mysql.sh     // db初期化
```

|service|host:port|
|--|--|
|Mysql|localhost:3306|
|phpmyadmin|localhost:8081|
|Redis|localhost:6379|
|redisinsight|localhost:8001|

アプリケーション起動  
`cmd/cago`にて
```
go run main.go
```

ランキングRedis更新  
`cmd/update`にて
```
go run update.go
```


### 負荷計測

vegetaを使用して負荷とレスポンス速度の計測  

`vegeta`にて  
```
vegeta attack -rate=10 -duration=5s -targets=gacha.txt | vegeta report
```

## 以下参考サイト

## 環境構築

### go

デバッグ環境  
https://qiita.com/momotaro98/items/7fbcad57a9d8488fe999  
https://qiita.com/momotaro98/items/10ae87b21903dd54601c  

### mysql

https://qiita.com/A-Kira/items/f401aea261693c395966  


## 実装

### net/http

https://qiita.com/convto/items/2822d029349cb1b4df93  
https://qiita.com/taizo/items/bf1ec35a65ad5f608d45  

### middleware

https://qiita.com/tnakata/items/ea962f1cdad21c2f68aa  
https://journal.lampetty.net/entry/implementing-middleware-with-http-package-in-go  

### DB

https://qiita.com/tenntenn/items/dddb13c15643454a7c3b  
https://precure-3dprinter.hatenablog.jp/entry/2018/11/22/Golang%E3%81%A7%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E3%82%92%E4%BD%BF%E3%81%86%E8%A9%B1  
https://qiita.com/mayah/items/a235a52a336095545e9d  
https://blog.withnic.net/2018/08/golang%E3%81%A7transaction%E5%87%A6%E7%90%86%E3%82%92%E8%80%83%E3%81%88%E3%82%8B/  

### mysql

https://www.sambaiz.net/article/189/  
https://akrfjmt.hatenablog.com/entry/2018/05/15/014455  

### redis 

https://qiita.com/akubi0w1/items/8701c05fe7186ceee632  
https://qiita.com/gold-kou/items/966d9a0332f4e110c4f8  
https://qiita.com/riverplus/items/c1df770838d3868c3a13  

### 外部設定ファイル

https://qiita.com/taka23kz/items/2f7aca381b3b27670fd0  

### context

https://qiita.com/yoshinori_hisakawa/items/50966e9ba2627e5ac124  


### jsonの受け取り

https://qiita.com/nyamage/items/e07de57d486238567ba7  
https://qiita.com/nayuneko/items/2ec20ba69804e8bf7ca3  

### 構造体

https://qiita.com/k-penguin-sato/items/62dfe0f93f56e4bf9157  

### error

https://qiita.com/sonatard/items/9c9faf79ac03c20f4ae1  
https://qiita.com/sonatard/items/5afd13a7640e628ee4d2  

### pprof

https://qiita.com/momotaro98/items/bd24a5d4603e378cc357  

### 負荷

https://qiita.com/chidakiyo/items/f8cdfac7683216a29c56  

## 設計

### DI

https://qiita.com/yoshinori_hisakawa/items/a944115eb77ed9247794  

### クリーンアーキテクチャ Laravel

https://qiita.com/nrslib/items/aa49d10dd2bcb3110f22  

### クリーンアーキテクチャ Go

https://qiita.com/muroon/items/8add8da911341312176d   
https://github.com/muroon/memo_sample  
