<template>
  <NuxtLayout name="404">
    <div>
    <canvas ref="canvas" class="w-full h-screen" ></canvas>
  </div>
  </NuxtLayout>
</template>

<script>
import { Rive, Layout } from "@rive-app/canvas";
export default {
	name: "Rive",
	props: {
		src: String,
		fit: String,
		alignment: String,
	},
	mounted() {
		new Rive({
			canvas: this.$refs.canvas,
			src: this.$props.src,
			layout: new Layout({
				fit: this.$props.fit,
				alignment: this.$props.alignment,
			}),
			autoplay: true,
		});
	},
};
</script>
