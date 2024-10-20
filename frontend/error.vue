<!-- <script setup lang="ts">
import type { NuxtError } from "nuxt/app";

const props = defineProps({
	error: Object as () => NuxtError,
});

import { Rive, Layout } from "@rive-app/canvas";

const name = "Rive";
const src = String;
const fit = String;
const alignment = String;

mounted(() => {
	new Rive({
		canvas: this.$refs.canvas,
		src: this.$props.src,
		layout: new Layout({
      fit: this.$props.fit,
      alignment: this.$props.alignment,
    }),
    autoplay: true,
  });
});
</script>
<template>
  <div>
    <NuxtLayout>
    <div>
    <canvas ref="canvas" class="w-full h-screen" ></canvas>
  </div>
  </NuxtLayout>
    
    <!-- <h1>{{ error?.statusCode }}</h1>
    <NuxtLink to="/">Go back home</NuxtLink> -->
  </div>
</template> -->

<template>
  <div>
    <canvas ref="canvas" width="400" height="400"></canvas>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { Rive } from '@rive-app/canvas';

const props = defineProps<{
	error: Object as () => NuxtError,
  src: string;
}>();

const canvas = ref<HTMLCanvasElement | null>(null);

onMounted(() => {
  if (canvas.value) {
    new Rive({
      canvas: canvas.value,
      src: props.src,
      autoplay: true,
    });
  }
});
</script>
