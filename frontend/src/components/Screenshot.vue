<script setup>
import { defineModel, defineEmits, defineProps, ref, onMounted } from "vue";
import LZString from "lz-string";
import Button from "primevue/button";
import Panel from "primevue/panel";
import Image from "primevue/image";
import SelectButton from "primevue/selectbutton";
import FileUpload from "primevue/fileupload";
import ProgressBar from 'primevue/progressbar';
import imageCompression from "browser-image-compression";

const props = defineProps(["label", "url"]);
const emit = defineEmits(["change"]);
const model = defineModel();
const label = props.label;
const url = props.url;

const uploading = ref(false)
const progress = ref(0)

async function upload(e) {
  var file = e.files[0];
  if (!file) {
    return;
  }

  uploading.value = true
  let blob = await fetch(file.objectURL).then((r) => r.blob()); //blob:url
  let b64 = await compress(blob);
  await uploadImage(b64);
  uploading.value = false
}

async function paste(e) {
  try {
    const clipboard = await navigator.clipboard.read();
    if (!clipboard) {
      return;
    }
    
    const file = clipboard[0];
    const f = file.types.findIndex((t) => t.includes("image"));

    if (!file?.types[f]?.includes("image")) {
      return;
    }
    
    uploading.value = true
    let blob = await file.getType(file.types[f]);
    let b64 = await compress(blob);
    await uploadImage(b64);
  } catch (error) {
    console.log(error)

  } finally {
    uploading.value = false
    progress.value = 0
  }
}

async function compress(blob) {
  console.log("original size:", blob.size / 1024 / 1024 + " MB");
  blob = await imageCompression(blob, {
    maxSizeMB: 1,
    maxWidthOrHeight: 1280,
    useWebWorker: true,
    onProgress: (p) => progress.value = p,
  })
  console.log("final size:", blob.size / 1024 / 1024 + " MB");

  return blob
  // const bf = await blob.arrayBuffer()
  // const u = new Uint8Array(bf)
  // return `[${u.toString()}]`

  // return new Promise((resolve, reject) => {
  //   let reader = new FileReader();
  //   reader.readAsDataURL(blob);
  //   reader.onloadend = function () {
  //     let base64 = reader.result;
  //     console.log(base64.length)
      
  //     // base64 = LZString.compressToBase64(base64)
  //     console.log(base64.length)
  //     resolve(base64);
  //   };
  // });
} 

async function del() {
  // uploadImage(null);
  model.value = false;
  emit("change");
}

async function uploadImage(img) {
  const ext = getExtension(img);

  const formData = new FormData();
  formData.append('file', img);

  const res = await fetch(`http://localhost:26023/upload/${url}.${ext}`, {
    method: "POST",
    mode: 'no-cors',
    headers: {
      "Content-Type": "application/json",

    },
    body: formData,
  })

  model.value = ext;
  emit("change");
}

function getImage() {
  // return LZString.decompressFromBase64(model.value)
  // return model.value
  return `http://localhost:26023/static/`+url+`.`+model.value
}

function getExtension(blob) {
  const mime = blob.type;
  return mime.split('/')[1];
}
</script>

<template>
  <div
    class="ssplaceholder flex flex-column flex-nowrap align-items-center justify-content-center align-content-center"
    v-if="uploading"  
  >
    <ProgressBar :value="progress" :displayValue="true" style="width:70%" />
  </div>
  <div
    class="ssplaceholder flex flex-column flex-nowrap align-items-center justify-content-center align-content-center"
    v-else-if="!model"
  >
    <div class="text-300 pb-2">{{ label }}</div>
    <div class="flex flex-wrap">
      <FileUpload
        class="file-select"
        mode="basic"
        accept="image/*"
        :maxFileSize="10000000"
        auto
        customUpload
        @uploader="upload"
      >
        Select
      </FileUpload>
      <span class="px-2">or</span>
      <Button
        type="button"
        label="Paste"
        class="file-paste"
        icon="pi pi-file-import"
        @click="paste"
      />
    </div>
  </div>
  <div class="sswrapper" v-else>
    <Image :src="getImage()" width="100%" class="screenshot" preview />
    <Button
      class="ssaction"
      type="button"
      icon="pi pi-minus-circle"
      severity="danger"
      size="small"
      text
      @click="del"
    />
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
