import Client from "./client";

const videos = new Map<string, HTMLVideoElement>();
const btn = document.querySelector('button');

async function f() {
  try {
    btn.disabled = true;
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
    });
    c.onLeave = ((clientId) => {
      const elId = `#partner-${clientId}`;
      const pre = document.getElementById(elId);
      pre?.parentNode.removeChild(pre);
    });
  } catch (e) {
    document.querySelector('#error').innerHTML = e
  }
}

btn.onclick = f;
