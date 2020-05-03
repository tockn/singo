import Client from "./client";

const videos = new Map<string, HTMLVideoElement>();

async function f() {
  try {
    const c = new Client();
    await c.joinRoom('hoge');
    c.onTrack = ((clientId, stream) => {
      const elId = `#partner-${clientId}`;
      const pre = document.getElementById(elId);
      pre?.parentNode.removeChild(pre);

      const $video = document.createElement('video') as HTMLVideoElement;
      $video.id = elId;
      const pa = document.querySelector('#partners');
      pa.appendChild($video);
      $video.srcObject = stream;
      $video.volume = 0;
      $video.play();
    })
  } catch (e) {
    document.querySelector('#error').innerHTML = e
  }
}
const btn = document.querySelector('button');
btn.onclick = f;
