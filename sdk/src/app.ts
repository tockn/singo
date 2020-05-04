// import Client from "./client";
//
// const videos = new Map<string, HTMLVideoElement>();
// const btn = document.querySelector('#join') as HTMLButtonElement;
// const myScreen = document.querySelector('#my-screen') as HTMLVideoElement;
//
// async function f() {
//   try {
//     btn.disabled = true;
//     const c = new Client(myScreen);
//     await c.joinRoom('hoge');
//     c.onTrack = ((clientId, stream) => {
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
//       $video.volume = 0;
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
