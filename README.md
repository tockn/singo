# singo

![icon](https://user-images.githubusercontent.com/29294540/81082118-cdf66a00-8f2d-11ea-97a8-eb31e6c2fe8c.jpg)

## singoとは

singoはフルメッシュP2Pによる複数人通信が可能なWebRTCシグナリングサーバです。  
JavaScript (TypeScript)によるSDKを提供しています。また、サンプルとして簡単なビデオ通話システムも用意しています。

![image](https://user-images.githubusercontent.com/29294540/81088076-752acf80-8f35-11ea-8f82-b320352e3a21.png)

## 使い方

### シグナリングサーバのみを立ち上げる

#### Dockerを使う場合

```
$ make docker-run
```

### ローカルでビルドして使う場合

```
$ go version
> go version go1.14 darwin/amd64

$ make run
```

### シグナリングサーバとサンプルのビデオ通話システムを立ち上げる

#### Dockerを使う場合

```
$ make docker-run-example
```

### ローカルでビルドして使う場合

```
$ go version
> go version go1.14 darwin/amd64

$ make run-example
```

## 処理の流れ

### シーケンス図

![sequence](https://github.com/tockn/singo/blob/master/sequence.png?raw=true)

クライアントとsingoはWebSocketでコネクションを張り、各メッセージをやり取りします。主要なメッセージとして以下があります。

##### クライアント側から

[定義](https://github.com/tockn/singo/blob/master/handler/message.go)

- join
  - roomに入るときに送ります。roomIDをつけます
- offer
  - 特定のclientに対してSDP Offerを送信します
- answer
  - 特定のclientに対してSDP Answerを送信します

##### singo側から

[定義](https://github.com/tockn/singo/blob/master/model/message.go)

- notify-client-id
  - joinしたclientに対してclient idを通知します
- offer
  - 特定のclientに対してSDP Offerを送信します
- answer
  - 特定のclientに対してSDP Answerを送信します
- new-client
  - roomに新たにjoinしてきたclient情報を通知します
- leave-client
  - roomから退出したclient情報を通知します
