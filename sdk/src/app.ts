import Client from "./client";

async function f() {
  const c = new Client();
  await c.connect();
  c.joinRoom('hoge');
}
f();
