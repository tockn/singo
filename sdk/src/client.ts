export default class Client {
  private readonly endpoint: string;
  private ws: WebSocket;
  private pcs: Map<string, RTCPeerConnection> = new Map<string, RTCPeerConnection>();
  private connected = false;
  private clientId: string;
  public stream: MediaStream;

  constructor(options?: ClientOptions) {
    this.endpoint = options?.SignalingServerEndpoint || 'ws://192.168.0.140:5000';
  }

  public async createNewPeer() {
    const pc = new RTCPeerConnection();
    this.pcs.set('hoge', pc);

    const stream = await navigator.mediaDevices.getUserMedia({
      video: true,
      audio: true,
    });
    this.stream = stream;

    for (const track of stream.getTracks()) {
      pc.addTrack(track, stream);
    }

    pc.onicecandidate = async (ev) => {
      if (!ev.candidate) {
        this.sendOffer();
      }
    };

    if ('ontrack' in pc) {
      pc.ontrack = async (ev) => {
          const $v = document.querySelector('#partner') as HTMLVideoElement;
          $v.srcObject = ev.streams[0];
          $v.volume = 0;
          $v.play()
      }
    } else {
      // @ts-ignore
      pc.onaddstream = (async ev => {
          const $v = document.querySelector('#partner') as HTMLVideoElement;
          $v.srcObject = ev.stream;
          $v.volume = 0;
          $v.play()
      });
    }
  }

  public async joinRoom(roomID: string): Promise<void> {
    return new Promise<void>(((resolve, reject) => {
      this.ws = new WebSocket(`${this.endpoint}/connect`);
      this.ws.onmessage = async (e: MessageEvent) => {
        await this.handleMessage(JSON.parse(e.data))
      };
      this.ws.onopen = (e: Event) => {
        this.connected = true;
        this.ws.send(JSON.stringify({
          type: 'join',
          payload: {
            room_id: roomID
          }
        }));
        resolve();
      };
    }));
  }

  public async createOffer() {
    const pc = this.pcs.get('hoge');
    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);
    // Trickle ICEならここで初期ICEを送る
  }

  public async sendOffer() {
    const pc = this.pcs.get('hoge');
    this.ws.send(JSON.stringify({
      type: MessageType.Offer,
      payload: {
        sdp: pc.localDescription.sdp,
      }
    }));
  }

  private async handleMessage(data: Message) {
    switch (data.type) {
      case MessageType.NotifyClientId:
        this.handleMessageNotifyClientId(data.payload);
        break;
      case MessageType.Offer:
        await this.handleMessageOffer(data.payload);
        break;
      case MessageType.Answer:
        await this.handleMessageAnswer(data.payload);
        break;
    }
  }

  private handleMessageNotifyClientId(payload: any) {
    this.clientId = payload.client_id;
  }

  private async handleMessageOffer(payload: any) {
    const pc = this.pcs.get('hoge');
    await pc.setRemoteDescription({
      type: "offer",
      sdp: payload.sdp,
    });
    const answer = await pc.createAnswer();
    this.ws.send(JSON.stringify({
      type: MessageType.Answer,
      payload: {
        sdp: answer.sdp
      }
    }));
    await pc.setLocalDescription(answer)
  }

  private async handleMessageAnswer(payload: any) {
    const pc = this.pcs.get('hoge');
    await pc.setRemoteDescription({
      type: "answer",
      sdp: payload.sdp,
    });
  }
}

interface ClientOptions {
  SignalingServerEndpoint: string
}

enum MessageType {
  NotifyClientId = 'notify-client-id',
  Error = 'error',
  Offer = 'offer',
  Answer = 'answer'
}

interface Message {
  type: MessageType
  payload: any
}

// stream取得でビデオ通話やり取り
// 複数人対応のために、pc作成タイミングとかpcsとか色々
