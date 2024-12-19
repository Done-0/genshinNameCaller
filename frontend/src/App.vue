<template>  
  <div   
    class="container"   
    @dblclick.stop="toggleWindowFullscreen"  
  >  
    <LotteryWithVideo />  
    <div v-if="isTransitioning" class="overlay"></div>  
  </div>  
</template>

<script>  
import LotteryWithVideo from './components/LotteryWithVideo.vue';  
import { ref, onMounted, onUnmounted } from 'vue';
import { WindowFullscreen, WindowUnfullscreen  } from './wailsjs/runtime/runtime';

export default {  
  components: {  
    LotteryWithVideo  
  },  
  setup() {  
    const isTransitioning = ref(false);  
    const isFullscreen = ref(false);  

    const preventSelection = (e) => {  
      e.preventDefault();  
      return false;  
    };  

    const toggleWindowFullscreen = async (event) => {  
      event.stopPropagation();  
      event.preventDefault();  

      try {
        const currentFullscreen = isFullscreen.value;

        if (currentFullscreen) {
          WindowUnfullscreen();
          isFullscreen.value = false;
        } else {
          WindowFullscreen();
          isFullscreen.value = true;
        }

        startTransition();
      } catch (error) {
        console.error('Fullscreen toggle error:', error);
      }
    };

    const startTransition = () => {  
      isTransitioning.value = true;  
      setTimeout(() => {  
        isTransitioning.value = false;  
      }, 500);  
    };  

    onMounted(() => {  
      document.addEventListener('selectstart', preventSelection);  
      document.addEventListener('mousedown', preventSelection);  
      
      document.body.style.userSelect = 'none';  
      document.body.style.WebkitUserSelect = 'none';  
      document.body.style.MozUserSelect = 'none';  
      document.body.style.msUserSelect = 'none';  
      document.body.style.cursor = 'default';  
    });  

    onUnmounted(() => {  
      document.removeEventListener('selectstart', preventSelection);  
      document.removeEventListener('mousedown', preventSelection);  
    });  

    return {  
      isTransitioning,  
      isFullscreen,  
      toggleWindowFullscreen  
    };  
  }  
}  
</script>  

<style scoped>  
html, body {  
  user-select: none !important;  
  -webkit-user-select: none !important;  
  -moz-user-select: none !important;  
  -ms-user-select: none !important;  
  cursor: default;  
  overflow: hidden;  
  margin: 0;  
  padding: 0;  
}  

* {  
  -webkit-touch-callout: none;  
}  

.container {  
  position: relative;  
  width: 100%;  
  height: 100%;  
  user-select: none;  
  pointer-events: auto;  
}  

.overlay {  
  position: absolute;  
  top: 0;  
  left: 0;  
  width: 100%;  
  height: 100%;  
  background-color: rgba(0, 0, 0, 0.5);  
  transition: opacity 0.5s ease-in-out;  
  opacity: 1;  
  pointer-events: none;  
}  

.container:not(.transitioning) .overlay {  
  opacity: 0;  
}  
</style>