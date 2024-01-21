<script setup>
import { onMounted, onUnmounted, ref } from 'vue'
import { DeleteNade, DeleteVariant, GetAllNades, GetAllVariants, InsertNade, InsertVariant, UpdateNade, UpdateVariant } from '../../wailsjs/go/main/App'
import { useConfirm } from "primevue/useconfirm";
import ContextMenu from 'primevue/contextmenu'
import ConfirmDialog from 'primevue/confirmdialog';
import Image from 'primevue/image';
import 'konva'
import Konva from 'konva';

const emit = defineEmits(['nadeSelected', 'variantSelected'])

const props = defineProps(['map'])
const map = props.map

defineExpose({
	deleteNadeNode,
	deleteVariantNode,
})

const confirm = useConfirm();

let stage;
let layerMap;
let layerIcons;
let size;
const WIDTH = 1200;
const HEIGHT = 1200;

let mode = 'nade' // nade, variant
let loading = ref(false)
let nadeNodes = []
let variantNodes = []
let selectedNadeNode = null
let selectedVariantNode = null

let contextPosition = { x: 0, y: 0 }
const nadeContext = ref()
const nadeContextActions = ref([
	{ label: 'Smoke', icon: 'pi pi-plus', command: () => createNadeOnMouse('smoke') },
	{ label: 'Molly', icon: 'pi pi-plus', command: () => createNadeOnMouse('molly') },
	{ label: 'Flash', icon: 'pi pi-plus', command: () => createNadeOnMouse('flash') },
	{ label: 'HE', icon: 'pi pi-plus', command: () => createNadeOnMouse('he') },
])
const variantContext = ref()
const variantContextActions = ref([
	{ label: 'New Variant', icon: 'pi pi-plus', command: () => createVariantToMouse() },
	// { label: 'Delete', icon: 'pi pi-trash', command: () => console.log('Delete') },
])

let contextNadeNode = null
const nadeNodeContext = ref()
const nadeNodeContextActions = ref([
	{ label: 'Delete', icon: 'pi pi-trash', command: () => deleteNadeNode() },
])

let contextVariantNode = null
const variantNodeContext = ref()
const variantNodeContextActions = ref([
	{ label: 'Delete', icon: 'pi pi-trash', command: () => deleteVariantNode() },
])

onMounted(async () => {
	startup()
	onResize()
	window.addEventListener('resize', onResize)

	await load()
});

onUnmounted(async () => {
	window.removeEventListener('resize', onResize)
})

function getRadarBgUrl(id) {
	return new URL(`../assets/images/maps/${id}/radar.png`, import.meta.url).href
}

function getIconUrl(id) {
	return new URL(`../assets/images/icons/${id}.png`, import.meta.url).href
}

async function load() {
	loading.value = true
	const res = await GetAllNades(map)
	loading.value = false

	res.forEach(nade => {
		addNade(nade)
	})
}

function startup() {
	stage = new Konva.Stage({
		container: 'radar',
		width: WIDTH,
		height: HEIGHT
	});

	layerMap = new Konva.Layer();
	layerIcons = new Konva.Layer();

	Konva.Image.fromURL(getRadarBgUrl(map), function (node) {
		node.setAttrs({
			x: 0,
			y: 0,
			width: WIDTH,
			height: HEIGHT,
			bg: true
		});
		node.cache()
		node.filters([Konva.Filters.Grayscale, Konva.Filters.Brighten]);
		node.brightness(-.4)
		layerMap.add(node);
	});

	stage.add(layerMap);
	stage.add(layerIcons);

	layerMap.draw()
	layerIcons.draw()

	stage.on('click', onClick)
	stage.on('dblclick', onDoubleClick)
	stage.on('contextmenu', onContext)
	stage.on('wheel', onZoom)
}

function onResize() {
	// const container = document.querySelector('#radar');
	const width = window.innerWidth / 2
	const height = window.innerHeight

	size = width;
	if (width > height) {
		size = height;
	}

	stage.width(size);
	stage.height(size);
	stage.scale({
		x: size / WIDTH,
		y: size / HEIGHT
	})
}

function onClick(e) {
	contextVariantNode = null
	contextNadeNode = null

	if (e.evt.which === 3) return // right click

	e.evt.preventDefault()

	if (mode === 'nade') {
		if (!e.target.getAttr('nade')) return
		selectNadeNode(e.target)

	} else if (mode === 'variant') {
		if (variantContext.value?.visible || nadeNodeContext.value?.visible || variantNodeContext.value?.visible) {
			variantContext.value?.hide()
			nadeNodeContext.value?.hide()
			variantNodeContext.value?.hide()
			return
		}

		if (!e.target.getAttr('nade')) {
			setMode('nade')
			return
		}

		if (e.target.getAttr('variant')) {
			selectVariantNode(e.target)
			return
		}
	}
}

function onDoubleClick(e) {
	// e.evt.preventDefault()
	// if (e.target !== stage && !e.target.getAttr('bg')) {
	// 	return
	// }

	// contextPosition = stage.getPointerPosition()
	// addVariantToMouse()
}

function onContext(e) {
	contextVariantNode = null
	contextNadeNode = null

	e.evt.preventDefault()
	if (e.target !== stage && !e.target.getAttr('bg')) {
		if (e.target.getAttr('variant')) {
			contextVariantNode = e.target
			variantNodeContext.value.show(e.evt)

		} else if (e.target.getAttr('nade')) {
			contextNadeNode = e.target
			nadeNodeContext.value.show(e.evt)
		}

		return
	}

	contextPosition = stage.getPointerPosition()

	if (mode === 'nade') {
		nadeContext.value.show(e.evt)
	} else if (mode === 'variant') {
		variantContext.value.show(e.evt)
	}
}

function onZoom(e) {
	e.evt.preventDefault();
}

function clear() {
	nadeNodes = []
	variantNodes = []
	layerIcons.destroyChildren()
	layerIcons.draw()
}

async function selectNadeNode(node) {
	selectedNadeNode = node
	if (node !== null) {
		setMode('variant')
	}
	emit('nadeSelected', node?.getAttr('nade'))
}

async function selectVariantNode(node) {
  if (selectedVariantNode !== null) {
    selectedVariantNode.strokeWidth(0)
  }

	selectedVariantNode = node
	selectedVariantNode?.strokeWidth(2)
	emit('variantSelected', node?.getAttr('variant'), node?.getAttr('nade'))
}

async function setMode(m) {
	clear()
	mode = m

	if (mode === 'nade') {
		selectNadeNode(null)

		loading.value = true
		const res = await GetAllNades(map)
		loading.value = false

		res.forEach(nade => {
			addNade(nade)
		})

	} else if (mode === 'variant') {
		const nade = selectedNadeNode.getAttr('nade')

		loading.value = true
		const res = await GetAllVariants(nade.id)
		loading.value = false

		addNade(nade)

		selectedVariantNode = null
		for (const variant of res) {
			const v = await addVariant(variant)
			if (selectedVariantNode === null) {
				selectVariantNode(v)
			}
		}
	}
}

async function createNadeOnMouse(type) {
	const nade = {
		id: null,
		map: map,
		type: type,
		x: pos(contextPosition.x),
		y: pos(contextPosition.y),
	}

	nade.id = await InsertNade(nade)

	return addNade(nade, true)
}

async function addNade(nade, select) {
	const x = (sop(nade.x) / stage.scaleX()) - (stage.x() / stage.scaleX())
	const y = (sop(nade.y) / stage.scaleY()) - (stage.y() / stage.scaleY())

	Konva.Image.fromURL(getIconUrl(nade.type), (node) => {
		node.setAttrs({
			x: x,
			y: y,
			width: 64,
			height: 64,
			offsetX: 32,
			offsetY: 32,
			draggable: true,
			nade: nade
		});
		node.on('dragmove', (e) => {
      selectedNadeNode = node
      updateLinks(null, node)
    })
		node.on('dragend', (e) => {
      selectedNadeNode = node
			nade.x = pos(e.target.x() * stage.scaleX() + stage.x())
			nade.y = pos(e.target.y() * stage.scaleY() + stage.y())
			UpdateNade(nade)
      updateLinks(null, node)
		})

		layerIcons.add(node);
		layerIcons.draw();

		nadeNodes.push(node)

    if (select) {
      selectNadeNode(node)
    }
	});
}

async function createVariantToMouse() {
	const variant = {
		id: null,
		nadeId: selectedNadeNode.getAttr('nade').id,
		name: 'New Variant',
		description: 'No description',
		x: pos(contextPosition.x),
		y: pos(contextPosition.y),
	}

	variant.id = await InsertVariant(variant)

	addVariant(variant)
}

async function addVariant(variant) {
	const x = (sop(variant.x) / stage.scaleX()) - (stage.x() / stage.scaleX())
	const y = (sop(variant.y) / stage.scaleY()) - (stage.y() / stage.scaleY())

  const link = new Konva.Line({
    points: [0, 0, 0, 0],
    stroke: 'white',
    strokeWidth: 1,
    lineCap: 'round',
    lineJoin: 'round',
  });

	var circle = new Konva.Circle({
		x,
		y,
		radius: 12,
		fill: 'red',
		draggable: true,
		variant: variant,
    stroke: 'white',
		nade: selectedNadeNode.getAttr('nade'),
    link, 
	});
  circle.on('dragmove', (e) => {
    updateLinks(circle)
  })
	circle.on('dragend', (e) => {
		variant.x = pos(e.target.x() * stage.scaleX() + stage.x())
		variant.y = pos(e.target.y() * stage.scaleY() + stage.y())
		UpdateVariant(variant)
    updateLinks(circle)
	})

  link.points([selectedNadeNode.x(), selectedNadeNode.y(), circle.x(), circle.y()])

	layerIcons.add(link);
	layerIcons.add(circle);
	layerIcons.draw();
	variantNodes.push(circle)

	selectVariantNode(circle)

	return circle
}

async function deleteNadeNode() {
	if (contextNadeNode === null) {
		contextNadeNode = selectedNadeNode
	}

	const nade = contextNadeNode.getAttr('nade')
	await confirm.require({
		message: `Are you sure you want to delete this nade?`,
		header: 'Delete Nade',
		icon: 'pi pi-exclamation-triangle',
		accept: async () => {
			nadeNodes = nadeNodes.filter(n => n !== contextNadeNode)
			contextNadeNode.destroy()
			layerIcons.draw()
			await DeleteNade(nade.id)
			selectNadeNode(null)
			selectVariantNode(null)
			setMode('nade')
		},
		reject: () => { }
	});
}

async function deleteVariantNode() {
	if (contextVariantNode === null) {
		contextVariantNode = selectedVariantNode
	}

	const nade = contextVariantNode.getAttr('nade')
	const variant = contextVariantNode.getAttr('variant')
	await confirm.require({
		message: `Are you sure you want to delete the variant "${variant.name}"?`,
		header: 'Delete Variant',
		icon: 'pi pi-exclamation-triangle',
		accept: async () => {
			variantNodes = variantNodes.filter(n => n !== contextVariantNode)
			contextVariantNode.destroy()
			layerIcons.draw()
			await DeleteVariant(variant.id)
			selectVariantNode(null)
			setMode('variant')
		},
		reject: () => { }
	});
}

async function updateLinks(variant, node) {
  if (!selectedNadeNode) return
  const variants = !!variant? [variant] : variantNodes
  node = !!node? node : selectedNadeNode

  for (const variant of variants) {
    const link = variant.getAttr('link')

    link.points([node.x(), node.y(), variant.x(), variant.y()])
  }
  
  layerIcons.draw()
}

function pos(x) {
	return Math.round(1000 * x / size)
}
function sop(x) {
	return Math.floor(x * size / 1000)
}
</script>

<template>
	<div id="radar" class="radar"></div>
	<ContextMenu ref="nadeContext" :model="nadeContextActions" />
	<ContextMenu ref="nadeNodeContext" :model="nadeNodeContextActions" />
	<ContextMenu ref="variantContext" :model="variantContextActions" />
	<ContextMenu ref="variantNodeContext" :model="variantNodeContextActions" />
	<ConfirmDialog></ConfirmDialog>
</template>

<style scoped>
.radar {
	display: flex;
	align-items: center;
	justify-content: center;
	overflow: hidden;

	width: 100%;
	height: 100%;
}
</style>
