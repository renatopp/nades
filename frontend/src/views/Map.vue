<script setup>
import { onMounted, ref } from 'vue'
import InputText from 'primevue/inputtext';
import Textarea from 'primevue/textarea';
import Button from 'primevue/button';
import Panel from 'primevue/panel';
import Image from 'primevue/image';
import SelectButton from 'primevue/selectbutton';
import FileUpload from 'primevue/fileupload';

import Radar from '../components/Radar.vue'
import Screenshot from '../components/Screenshot.vue'
import { UpdateVariant } from '../../wailsjs/go/main/App';
import { useConfirm } from 'primevue/useconfirm';

const confirm = useConfirm();

var radar = ref()
var map = ref(null)
var selectedNade = ref(null)
var selectedVariant = ref(null)

onMounted(() => {
	map.value = router.currentRoute.value.params.mapId
})

function selectNade(nade) {
	selectedNade.value = nade
	selectedVariant.value = null
}

function selectVariant(variant, nade) {
	// selectedNade.value = nade
	selectedVariant.value = variant
}

function updateVariant() {
	UpdateVariant(selectedVariant.value)
}

function deleteVariant() {
	radar.value.deleteVariantNode()
}

function deleteNade() {
	radar.value.deleteNadeNode()
}

</script>

<template>
	<div class="operation">
		<router-link class="link" :to="'/'">
			<i class="pi pi-angle-left" style="font-size: 1.5rem"></i>
			<!-- {{ map }} -->
		</router-link>
	</div>

	<div class="flex align-items-stretch justify-content-center h-screen w-screen">
		<div class="radar flex-1 align-items-center">
			<Suspense>
				<Radar ref="radar" :map="map" @nade-selected="selectNade" @variant-selected="selectVariant" v-if="map" />
			</Suspense>
		</div>
		<div class="form flex-1 align-items-center" v-if="map">

			<!-- NO NADE SELECTED -->
			<div v-if="!selectedNade">
				<h4>Select a nade</h4>
			</div>

			<!-- NO VARIANT SELECTED -->
			<div v-if="selectedNade && !selectedVariant" class="flex align-items-center justify-content-center text-center">
				<div class="py-8">
					<p class="text-300">No Variant Yet</p>
					<p class="text-200 text-sm">Right click in the map to create a variant.</p>

					<Button type="button" label="Delete" icon="pi pi-trash" severity="danger" size="small" text
						@click="deleteNade" />
				</div>
			</div>

			<!-- VARIANT SELECTED -->
			<div v-if="selectedNade && selectedVariant">
				<Screenshot label="Throw Image" v-model="selectedVariant.throwImage" @change="updateVariant" />

				<div>
					<div class="flex">
						<InputText class="title transparent flex-1" type="text" v-model="selectedVariant.name"
							@change="updateVariant" />
						<Button type="button" label="Delete" icon="pi pi-trash" severity="danger" size="small" text
							@click="deleteVariant" />
					</div>

					<div class="flex">
						<Textarea class="transparent flex-1" v-model="selectedVariant.description" @change="updateVariant"
							autoResize />
					</div>
				</div>

				<div class="flex gap-2">
					<div class="flex-1">
						<Screenshot label="Lineup Image" v-model="selectedVariant.lineupImage" @change="updateVariant" />
					</div>
					<div class="flex-1">
						<Screenshot label="Position Image" v-model="selectedVariant.positionImage" @change="updateVariant" />
					</div>
				</div>
			</div>

			<!-- <NadeList @selected="select" v-hide="selectedNade" :map="map" /> -->
			<!-- <VariantList @deselected="select(null)" v-show="selectedNade" :nade="selectedNade" /> -->
		</div>
	</div>
</template>

<style scoped>
.operation {
	position: absolute;
	top: 10px;
	left: 10px;
	font-size: 2rem;
	font-family: 'Teko', sans-serif;
}

.operation .link {
	color: #fff;
	text-decoration: none;
	opacity: 0.3;
	transition: .1s ease-in-out all;
}

.operation .link:hover {
	opacity: 0.5;
	transition: .3s ease-in-out all;
}

.form {
	padding: 10px;
	border-left: 1px solid #444;
	max-width: 600px;
}

.screenshot {
	width: 100%;
	height: auto;
}

.title {
	/* font-size: 2rem; */
	font-weight: bold;
}

.transparent {
	background: transparent;
	border: none;
}

.ssplaceholder {
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
