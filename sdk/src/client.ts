export default class Client {
  private readonly endpoint: string;
  private ws: WebSocket;
  private pcs: Map<string, RTCPeerConnection> = new Map<string, RTCPeerConnection>();
  private connected = false;
  private clientId: string;

  constructor(options?: ClientOptions) {
    this.endpoint = options?.SignalingServerEndpoint || 'ws://127.0.0.1:5000';
  }

  public async connect(): Promise<void> {
    return new Promise<void>(((resolve, reject) => {
      this.ws = new WebSocket(`${this.endpoint}/connect`);
      this.ws.onmessage = async (e: MessageEvent) => {
        await this.handleMessage(JSON.parse(e.data))
      };
      this.ws.onopen = (e: Event) => {
        this.connected = true;
        resolve();
      };
    }));
  }

  public async joinRoom(roomID: string) {
    this.ws.send(JSON.stringify({
      type: 'join',
      payload: {
        room_id: roomID
      }
    }));

    const pc = new RTCPeerConnection();
    pc.onicecandidate = async (ev) => {
      if (ev.candidate === null) {
        return
      }
    };
    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);
    this.ws.send(JSON.stringify({
      type: MessageType.Offer,
      payload: {
        sdp: offer.sdp,
      }
    }));
    this.pcs.set('hoge', pc)
  }

  private async handleMessage(data: Message) {
    console.log(data, data.type, data.type === MessageType.Offer);
    switch (data.type) {
      case MessageType.NotifyClientId:
        this.handleMessageNotifyClientId(data.payload);
        break;
      case MessageType.Offer:
        console.log(data.payload);
        await this.handleMessageOffer(data.payload);
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
