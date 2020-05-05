// import { SingoClient } from "./client";
//
// const videos = new Map<string, HTMLVideoElement>();
// const btn = document.querySelector('#join') as HTMLButtonElement;
// const mutebtn = document.querySelector('#mute') as HTMLButtonElement;
// const videobtn = document.querySelector('#video') as HTMLButtonElement;
// const myScreen = document.querySelector('#my-screen') as HTMLVideoElement;
// let client: SingoClient;
//
// async function f() {
//   try {
//     btn.disabled = true;
//     const c = new SingoClient(myScreen);
//     client = c;
//     await c.joinRoom('hoge');
//     c.onTrack = ((clientId, stream) => {
//       console.log(clientId, stream);
//       const elId = `#partner-${clientId}`;
//       const pre = document.getElementById(elId);
//       pre?.parentNode.removeChild(pre);
//
//       const $video = document.createElement('video') as HTMLVideoElement;
//       $video.id = elId;
//       $video.setAttribute('autoplay', 'true')
//       $video.setAttribute('playsinline', 'true')
//       $video.setAttribute('muted', 'true')
//       const pa = document.querySelector('#partners');
//       pa.appendChild($video);
//       $video.srcObject = stream;
//     });
//     c.onLeave = ((clientId) => {
//       const elId = `#partner-${clientId}`;
//       const pre = document.getElementById(elId);
//       pre?.parentNode.removeChild(pre);
//     });
//   } catch (e) {
//     console.error(e)
//   }
// }
//
// btn.onclick = f;
//
// let mute = false;
// mutebtn.onclick = () => {
//   mute = !mute;
//   client.changeAudioTrackEnabled(mute)
// };
//
// let video = true;
// videobtn.onclick = () => {
//   video = !video;
//   client.changeVideoTrackEnabled(video);
// };
