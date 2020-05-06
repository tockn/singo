# singo JavaScript SDK

singoのJavaScript SDKです。

## Install

```
npm i --save singo-sdk
```
or
```
yarn add singo-sdk
```

## Usage (example)

```js
const myScreen = document.querySelector('#my-screen')

// コンストラクタで自分の画面のvideo要素を渡します。
const client = new SingoClient(myScreen, {
  signalingServerEndpoint: 'ws://localhost:5000'
})

// joinRoomメソッドで特定の名前のroomへ入ります。
// ここではvideoタグを追加して表示しています
await c.joinRoom('room name')

// onTrackは新たなclientのstreamを受け取ったときに呼ばれます
c.onTrack = ((clientId, stream) => {
  const elId = `#partner-${clientId}`;
  const pre = document.getElementById(elId);
  pre?.parentNode.removeChild(pre);

  const $video = document.createElement('video') as HTMLVideoElement;
  $video.id = elId;
  $video.setAttribute('autoplay', 'true')
  $video.setAttribute('playsinline', 'true')
  $video.setAttribute('muted', 'true')
  const pa = document.querySelector('#partners');
  pa.appendChild($video);
  $video.srcObject = stream;
});

// onLeaveはclientが退出したときに呼ばれます。
// ここでは退出したclientのvideoタグを削除しています。
c.onLeave = ((clientId) => {
  const elId = `#partner-${clientId}`;
  const pre = document.getElementById(elId);
  pre?.parentNode.removeChild(pre);
});
```
