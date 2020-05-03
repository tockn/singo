import Client from "./client";

async function f() {
  try {
    const c = new Client();
    await c.createNewPeer();
    await c.joinRoom('hoge');
    const btn = document.querySelector('button');
    btn.onclick = async () => {
      await c.createOffer();
    };
    const $video = document.querySelector('#my-screen') as HTMLVideoElement;
    $video.srcObject = c.stream;
    $video.volume = 0;
    await $video.play();
  } catch (e) {
    document.querySelector('#error').innerHTML = e
  }
}
f();
