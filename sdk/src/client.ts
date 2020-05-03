import userEnv from 'userEnv'
export default class Client {
  private readonly endpoint: string;
  private ws: WebSocket;
  private pcs: Map<string, RTCPeerConnection> = new Map<string, RTCPeerConnection>();
  private connected = false;
  private clientId: string;
  public stream: MediaStream;

  public onTrack: ((clientId: string, stream: MediaStream) => any);
  public onLeave: ((clientId: string) => any);

  constructor(options?: ClientOptions) {
    this.endpoint = options?.SignalingServerEndpoint || userEnv.wsUrl;
  }

  public async createNewPeer(clientId: string): Promise<RTCPeerConnection> {
    const pc = new RTCPeerConnection(
      {'iceServers':[
        {'urls': 'stun:stun.l.google.com:19302'},
        {'urls': 'stun:stun1.l.google.com:19302'},
        {'urls': 'stun:stun2.l.google.com:19302'}
      ]}
    );
    this.pcs.set(clientId, pc);

    if (!this.stream) {
      const stream = await navigator.mediaDevices.getUserMedia(
        {
          'audio': true,
          'video': {
            'width': {
              'max': 640
            },
            'height': {
              'max': 480
            },
            'frameRate': {
              'max': 20
            }
          }
        }
      );
      this.stream = stream;
      const $video = document.querySelector('#my-screen') as HTMLVideoElement;
      $video.srcObject = stream;
      $video.volume = 0;
      await $video.play();
    }

    for (const track of this.stream.getTracks()) {
      pc.addTrack(track, this.stream);
    }

    // TODO 複数のVideo対応する
    if ('ontrack' in pc) {
      pc.ontrack = async (ev) => {
        this.onTrack(clientId, ev.streams[0]);
      }
    } else {
      // @ts-ignore
      pc.onaddstream = (async ev => {
        this.onTrack(clientId, ev.stream);
      });
    }

    return pc;
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

  private async handleMessage(data: Message) {
    switch (data.type) {
      case MessageType.NotifyClientId:
        this.handleMessageNotifyClientId(data.payload);
        break;
      case MessageType.NewClient:
        this.handleNewClient(data.payload);
        break;
      case MessageType.LeaveClient:
        this.handleLeaveClient(data.payload);
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

  private async handleNewClient(payload: any) {
    const clientId = payload.client_id;
    const pc = await this.createNewPeer(clientId);
    await this.createOffer(clientId);
    pc.onicecandidate = async (ev) => {
      if (!ev.candidate) {
        this.sendOffer(clientId);
      }
    };
  }

  private async handleLeaveClient(payload: any) {
    const clientId = payload.client_id;
    const pc = this.pcs.get(clientId);
    pc.close();
    this.pcs.delete(clientId);
    this.onLeave(clientId);
  }

  public async createOffer(clientId: string) {
    const pc = this.pcs.get(clientId);
    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);
    // Trickle ICEならここで初期ICEを送る
  }

  public async sendOffer(clientId: string) {
    const pc = this.pcs.get(clientId);
    this.ws.send(JSON.stringify({
      type: MessageType.Offer,
      payload: {
        sdp: pc.localDescription.sdp,
        client_id: clientId,
      }
    }));
  }

  private async handleMessageOffer(payload: any) {
    const clientId = payload.client_id as string;
    const pc = await this.createNewPeer(clientId);
    pc.onicecandidate = async (ev) => {
      if (!ev.candidate) {
        await this.sendAnswer(clientId);
      }
    };
    console.log(clientId, payload)
    await pc.setRemoteDescription({
      type: 'offer',
      sdp: payload.sdp,
    });
    await this.createAnswer(clientId);
  }

  public async createAnswer(clientId: string) {
    const pc = this.pcs.get(clientId);
    const answer = await pc.createAnswer();
    await pc.setLocalDescription(answer);
    // Trickle ICEならここで初期ICEを送る
  }

  private sendAnswer(clientId: string) {
    const pc = this.pcs.get(clientId);
    this.ws.send(JSON.stringify({
      type: MessageType.Answer,
      payload: {
        sdp: pc.localDescription.sdp,
        client_id: clientId,
      }
    }));
  }

  private async handleMessageAnswer(payload: any) {
    const pc = this.pcs.get(payload.client_id);
    await pc.setRemoteDescription({
      type: 'answer',
      sdp: payload.sdp,
    });
  }
}

interface ClientOptions {
  SignalingServerEndpoint: string
}

enum MessageType {
  NotifyClientId = 'notify-client-id',
    NewClient = 'new-client',
    LeaveClient = 'leave-client',
    Error = 'error',
    Offer = 'offer',
    Answer = 'answer'
}

interface Message {
  type: MessageType
  payload: any
}
