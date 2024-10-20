
<template>
  <div>
    <canvas ref="canvas" height="1000" width="1000"></canvas>
  </div>
</template>

<script setup lang="ts">
import type { NuxtError } from '#app';
import { ref, onMounted } from 'vue';
import { Rive } from '@rive-app/react-canvas';

const props = defineProps({
	error: Object as () => NuxtError,
  src: { type: String, required: true },
});

const canvas = ref<HTMLCanvasElement | null>(null);

onMounted(() => {
  if (canvas.value) {
    new Rive({
      canvas: canvas.value,
      src: '/llg_404_test.riv',
    });
  }
});
</script>
