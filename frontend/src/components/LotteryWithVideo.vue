<template>
  <div class="lottery-container">
    <video ref="backgroundVideo" autoplay loop muted playsinline preload="auto" class="background-video"
      @canplay="onBackgroundVideoCanPlay">
      <source :src="backgroundVideoSrc" type="video/mp4" />
    </video>

    <div v-if="showLotteryVideo" class="video-container">
      <video ref="lotteryVideo" @ended="onLotteryVideoEnded" preload="auto" class="lottery-video">
        <source :src="lotteryVideoSrc" type="video/mp4" />
        您的浏览器不支持视频标签。
      </video>
    </div>

    <div v-if="showResult" class="result-container">
      <video ref="resultVideo" autoplay loop muted playsinline preload="auto" class="result-video">
        <source :src="resultVideoSrc" type="video/mp4" />
      </video>
      <div class="overlay-container" :class="{ visible: overlayVisible }">
        <img :src="overlayImageSrc" alt="Overlay" class="overlay-image" />
      </div>
      <div class="winner-name" :class="{ visible: nameVisible }">{{ winnerName }}</div>
    </div>

    <div class="controls">
      <label for="fileInput" class="file-input-label">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
          stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
          <polyline points="17 8 12 3 7 8"></polyline>
          <line x1="12" y1="3" x2="12" y2="15"></line>
        </svg>
      </label>
      <input id="fileInput" type="file" accept=".txt" @change="handleFileUpload" class="file-input" />
      <button @click="startLottery" :disabled="!hasNames || isLotteryInProgress" class="lottery-btn">
        开始抽奖
      </button>
    </div>

    <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
  </div>
</template>

<script>
let GetNameList, RandomName;

try {
  const wailsApp = window.go.main.App;
  GetNameList = wailsApp.GetNameList;
  RandomName = wailsApp.RandomName;
} catch (error) {
  console.error("无法导入 Wails 函数:", error);
}

export default {
  data() {
    return {
      names: [],
      winnerName: '',
      hasNames: false,
      isLotteryInProgress: false,
      showLotteryVideo: false,
      showResult: false,
      overlayVisible: false,
      nameVisible: false,
      backgroundVideoSrc: 'assets/videos/原神祈愿.mp4',
      lotteryVideoSrc: 'assets/videos/单抽出金.mp4',
      resultVideoSrc: 'assets/videos/原神祈愿.mp4',
      overlayImageSrc: 'assets/images/bg-catalyst.webp',
      errorMessage: '',
      videoLoaded: false,
      resultVideoPreloaded: false,
    };
  },
  mounted() {
    this.preloadAssets();
    const video = this.$refs.backgroundVideo;
    if (video) {
      video.play().catch(error => {
        console.error("背景视频播放错误:", error);
      });
    }
  },
  methods: {
    preloadAssets() {
      Promise.all([
        this.preloadVideo(this.backgroundVideoSrc),
        this.preloadVideo(this.resultVideoSrc),
        this.preloadImage(this.overlayImageSrc),
      ])
        .then(() => {
          console.log("所有资产已预加载");
          this.resultVideoPreloaded = true;
        })
        .catch(error => {
          console.error("资产预加载错误:", error);
        });
    },
    preloadVideo(src) {
      return new Promise((resolve, reject) => {
        const video = document.createElement('video');
        video.src = src;
        video.preload = 'auto';
        video.oncanplaythrough = () => resolve();
        video.onerror = () => reject();
      });
    },
    preloadImage(src) {
      return new Promise((resolve, reject) => {
        const img = new Image();
        img.src = src;
        img.onload = () => resolve();
        img.onerror = () => reject();
      });
    },
    onBackgroundVideoCanPlay() {
      this.videoLoaded = true;
      const video = this.$refs.backgroundVideo;
      if (video && !video.paused) {
        video.play().catch(error => {
          console.error("背景视频播放错误:", error);
        });
      }
    },
    async handleFileUpload(event) {
      try {
        const file = event.target.files[0];
        const text = await file.text();
        if (typeof GetNameList === 'function') {
          const nameList = await GetNameList(text);
          this.names = nameList;
          this.hasNames = this.names.length > 0;
          this.errorMessage = '';
        } else {
          throw new Error("GetNameList 函数不可用");
        }
      } catch (error) {
        console.error("文件上传错误:", error);
        this.errorMessage = "无法处理名单文件。请确保文件格式正确。";
      }
    },
    async startLottery() {
      if (this.isLotteryInProgress) return;

      this.isLotteryInProgress = true;
      this.showLotteryVideo = true;
      this.showResult = false;
      this.errorMessage = '';

      this.$nextTick(() => {
        const video = this.$refs.lotteryVideo;
        if (video) {
          video.play().catch(error => {
            console.error("视频播放错误:", error);
            this.errorMessage = "无法播放视频。请检查视频文件。";
            this.onLotteryVideoEnded();
          });
        } else {
          this.errorMessage = "视频元素不可用。";
          this.onLotteryVideoEnded();
        }
      });

      try {
        if (typeof RandomName === 'function') {
          this.winnerName = await RandomName();
        } else {
          throw new Error("RandomName 函数不可用");
        }
      } catch (error) {
        console.error("抽奖错误:", error);
        this.errorMessage = "抽奖过程中出现错误。请重试。";
        this.winnerName = "抽奖失败";
      }
    },
    onLotteryVideoEnded() {
      this.showLotteryVideo = false;
      this.showResult = true;
      this.isLotteryInProgress = false;

      if (this.resultVideoPreloaded) {
        const video = this.$refs.resultVideo;
        if (video) {
          video.play().catch(error => {
            console.error("结果视频播放错误:", error);
          });
        }
      }

      requestAnimationFrame(() => {
        setTimeout(() => {
          this.overlayVisible = true;
          this.nameVisible = true;
        }, 300);
      });
    }
  }
};
</script>

<style scoped>
html,
body {
  margin: 0;
  padding: 0;
  overflow: hidden;
}

.lottery-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  background: black;
}

.background-video,
.result-video {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.overlay-container {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 5;
  opacity: 0;
  transition: opacity 0.5s ease-in-out;
}

.overlay-container.visible {
  opacity: 1;
}

.overlay-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  filter: drop-shadow(0 0 10px rgb(182, 171, 113));
}

.winner-name {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 64px;
  font-weight: bold;
  color: #000000;
  text-shadow: 2px 2px 4px rgba(255, 255, 255, 0.5);
  z-index: 10;
  opacity: 0;
  transition: opacity 0.5s ease-in-out;
}

.winner-name.visible {
  opacity: 1;
}

.video-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: black;
  z-index: 10;
}

.lottery-video {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.result-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 5;
}

.controls {
  position: fixed;
  bottom: 20px;
  right: 20px;
  display: flex;
  align-items: center;
  z-index: 20;
}

.file-input {
  display: none;
}

.file-input-label {
  cursor: pointer;
  margin-right: 10px;
  padding: 5px;
  background-color: rgba(52, 152, 219, 0.7);
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: background-color 0.3s;
}

.file-input-label:hover {
  background-color: rgba(41, 128, 185, 0.7);
}

.lottery-btn {
  padding: 10px 20px;
  font-size: 18px;
  background-color: rgba(231, 76, 60, 0.7);
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.lottery-btn:hover:not(:disabled) {
  background-color: rgba(192, 57, 43, 0.7);
}

.lottery-btn:disabled {
  background-color: rgba(149, 165, 166, 0.7);
  cursor: not-allowed;
}

.error-message {
  position: fixed;
  bottom: 20px;
  left: 20px;
  color: #e74c3c;
  background-color: rgba(255, 255, 255, 0.8);
  padding: 10px;
  border-radius: 5px;
  z-index: 30;
}
</style>