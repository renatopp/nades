<script setup>
import { defineModel, defineEmits, defineProps, ref, onMounted, computed, watch } from "vue";
import Image from "primevue/image";

const props = defineProps(["label", "url", "crosshair", "magnify"]);
const url = ref(props.url);
const label = ref(props.label);
const crosshair = ref(props.crosshair || false);
const magnify = ref(props.magnify || false);

watch(props, () => {
	url.value = props.url
	label.value = props.label
	crosshair.value = props.crosshair
	magnify.value = props.magnify
}, { immediate: true, deep: true })

</script>

<template>
	<Image :src="url" :alt="label" class="r2p-image" :class="{ crosshair, magnify }" />
</template>

<style>
.r2p-image {
	position: relative;
	overflow: hidden;
	transition: .2s ease-in-out all;
}

.r2p-image.magnify:hover img {
	transition: .1s ease-in-out all;
	transform: scale(2.5);
}

.r2p-image:hover button {
	display: hidden !important;
	opacity: 0 !important;
}

.r2p-image {
	overflow: hidden;
	transition: .2s ease-in-out all;
}

.r2p-image.crosshair::after {
	content: url('../assets/images/crosshair.png');
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%) scale(0.15);
	transition: .2s ease-in-out all;
}

.r2p-image.crosshair:hover::after {
	transition: .1s ease-in-out all;
	opacity: 0.3;
}
</style>