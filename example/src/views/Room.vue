<template>
  <div class="room" ref="room">
    <video-menu :on-leave-clicked="leave" />
    <div class="screens">
      <div class="screen">
        <video-screen ref="myScreen" :stream="myStream" screen-id="myscreen" />
      </div>
      <div class="screen">
        <video-screen ref="myScreen" :stream="myStream" screen-id="myscreen" />
      </div>
      <div class="screen">
        <video-screen ref="myScreen" :stream="myStream" screen-id="myscreen" />
      </div>
      <div class="screen">
        <video-screen ref="myScreen" :stream="myStream" screen-id="myscreen" />
      </div>

      <div class="screen" v-for="(s, i) in partnerStreams" :key="i">
        <video-screen :stream="s" :screen-id="i" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { SingoClient } from "singo-sdk";
import VideoScreen from "@/components/VideoScreen.vue";
import VideoMenu from "@/components/VideoMenu.vue";
import { disableBodyScroll } from "body-scroll-lock";
@Component({
  components: { VideoMenu, VideoScreen }
})
export default class Room extends Vue {
  private client: SingoClient;
  private myStream: MediaStream = new MediaStream();
  private partnerStreamMap: Map<string, MediaStream> = new Map();
  private partnerStreams: MediaStream[] = [];
  private ref: Element;

  get roomId(): string {
    return this.$route.params.id;
  }

  async mounted() {
    this.ref = this.$refs.room as Element;
    disableBodyScroll(this.ref);
    const sc = this.$refs.myScreen as HTMLVideoElement;
    this.client = new SingoClient(sc, {
      SignalingServerEndpoint: "ws://localhost:5000"
    });
    await this.client.joinRoom(this.roomId);
    this.myStream = this.client.stream;

    this.client.onTrack = (clientId, stream) => {
      this.partnerStreamMap.set(clientId, stream);
      this.updatePartnerStream();
    };

    this.client.onLeave = clientId => {
      this.partnerStreamMap.delete(clientId);
      this.updatePartnerStream();
    };
  }

  private updatePartnerStream() {
    const ps = [] as MediaStream[];
    this.partnerStreamMap.forEach(stream => {
      ps.push(stream);
    });
    this.partnerStreams = ps;
  }

  private leave() {
    if (!confirm("本当に退出しますか？")) return;
    this.$router.push("/");
  }
}
</script>

<style scoped>
.meeting {
}
.screens {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  width: 100%;
}
.screen {
  font-size: 0;
}
</style>
