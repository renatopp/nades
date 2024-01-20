<script setup>
import { defineModel, defineEmits, defineProps } from 'vue'
import Button from 'primevue/button';
import Panel from 'primevue/panel';
import Image from 'primevue/image';
import SelectButton from 'primevue/selectbutton';
import FileUpload from 'primevue/fileupload';
import imageCompression from 'browser-image-compression'

const model = defineModel()
const emit = defineEmits(['change'])
const props = defineProps(['label'])
const label = props.label

async function upload(e) {
	var file = e.files[0];
	if (!file) {
		return;
	}

	let blob = await fetch(file.objectURL).then((r) => r.blob()); //blob:url
	let reader = new FileReader();
	reader.readAsDataURL(blob);
	reader.onloadend = function () {
		const base64data = reader.result;
		setImage(base64data)
	};
}

async function paste(e) {
	try {
		const clipboard = await navigator.clipboard.read();
		if (!clipboard) {
			return;
		}

		const file = clipboard[0];
		if (!file?.types[0].includes('image')) {
			return;
		}

		console.log(typeof file)
		console.log(Object.keys(file))

		let blob = await file.getType(file.types[0])
		let b64 = await compress(blob)
		setImage(b64)
	} catch (error) {

	}
}

async function compress(blob) {
	return new Promise((resolve, reject) => {
		let reader = new FileReader()
		reader.readAsDataURL(blob)
		reader.onloadend = function () {
			const base64data = reader.result;
			resolve(base64data)
		}
	})
}

async function del() {
	setImage(null)
}

async function setImage(img) {
	model.value = img
	emit('change')
}


</script>

<template>
	<div class="ssplaceholder flex flex-column flex-nowrap align-items-center justify-content-center align-content-center"
		v-if="!model">
		<div class="text-300 pb-2">{{ label }}</div>
		<div class="flex flex-wrap">
			<FileUpload class="file-select" mode="basic" accept="image/*" :maxFileSize="10000000" auto customUpload
				@uploader="upload">
				Select
			</FileUpload>
			<span class="px-2">or</span>
			<Button type="button" label="Paste" class="file-paste" icon="pi pi-file-import" @click="paste" />
		</div>
	</div>
	<div class="sswrapper" v-else>
		<Image :src="model" width="100%" class="screenshot" preview />
		<Button class="ssaction" type="button" icon="pi pi-minus-circle" severity="danger" size="small" text @click="del" />
	</div>
</template>


<style scoped>
.ssplaceholder {
	padding: 10px;
	width: 100%;
	aspect-ratio: auto 16/9;
	background: #222;
}

.sswrapper {
	position: relative;
}

.ssaction {
	position: absolute;
	top: 5px;
	right: 5px;
}
</style>