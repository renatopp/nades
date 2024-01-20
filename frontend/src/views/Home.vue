<script setup>
import Card from 'primevue/card';
import Button from 'primevue/button'
import rawMaps from '../assets/maps.json'

function getMaps() {
  const maps = []
  for (const id in rawMaps) {
    const map = rawMaps[id]
    if (!map.enabled) continue

    maps.push({
      id,
      ...rawMaps[id]
    })
  }

  return maps
}

function getBackground(id) {
	return new URL(`../assets/images/maps/${id}/bg.webp`, import.meta.url).href
}

function getLogo(id) {
	return new URL(`../assets/images/maps/${id}/logo.webp`, import.meta.url).href
}

</script>

<template>
	<div class="home flex align-items-center justify-content-center h-screen w-screen">
		<div class="main flex align-items-center justify-content-center h-screen w-screen">
			<router-link class="banner flex-1" v-for="(map, i) in getMaps()" :key="i" :to="'/map/' + map.id">
				<div class="bg" v-bind:style="{ backgroundImage: 'url(\'' + getBackground(map.id) + ' \')' }"></div>
				<img :src="getLogo(map.id)" class="logo" />
			</router-link>
		</div>
	</div>
</template>

<style scoped>
.main {
	max-width: 1200px;
	padding: 0 20px;
}

.banner {
	position: relative;
	max-width: 300px;
	aspect-ratio: 1/2.5;
	display: flex;
	align-items: center;
	transform: skew(-10deg);
	overflow: hidden;
	/* border: 1px solid #333; */
	margin: 0 2px;
}

.bg {
	position: absolute;
	top: -50%;
	left: -50%;
	right: -50%;
	bottom: -50%;
	background-size: cover;
	background-position: center;
	background-repeat: no-repeat;
	transition: all 0.1s ease-in-out;
	z-index: 0;
	transform: skew(10deg);
}

.logo {
	position: absolute;
	width: 90%;
	max-width: 100px;
	height: auto;
	z-index: 1;
	top: 50%;
	left: 58%;
	transform: skew(10deg) translate(-50%, -50%);
}

.banner:hover {
	position: relative;
	cursor: pointer;
	transform: scale(1.05) skew(-10deg);
	transition: all 0.2s ease-in-out;
	z-index: 1;
}

/* .banner:hover .logo {
	transform: scale(1.05) skew(10deg) translate(-50%, -50%);
	transition: all 0.2s ease-in-out;
} */

.banner43 {
	width: auto;
	height: 300px;
	padding: 0 20px;
	display: flex;
	align-items: center;
	justify-content: center;
	transition: all 0.1s ease-in-out;
	transform: skew(-10deg);
}

.banner43>* {
	transform: skew(10deg);
}

.banner43:hover {
	position: relative;
	cursor: pointer;
	transform: scale(1.05) skew(-10deg);
	transition: all 0.2s ease-in-out;
	z-index: 1;
}
</style>
