<template>
  <q-input v-bind="attrs" type="text" v-model="model">
    <template v-slot:append>
      <div>
        <q-icon name="sym_o_barcode">
          <q-popup-proxy
            ref="popupProxy"
            transition-show="scale"
            transition-hide="scale"
            cover
            :breakpoint="600"
            @show="onShow"
            @before-hide="onHide"
          >
            <div class="window">
              <div :id="attrs.name"></div>
            </div>
          </q-popup-proxy>
        </q-icon>
      </div>
    </template>
  </q-input>
</template>
<style scoped>
.window {
  width: 400px;
  max-width: 80vw;
  max-height: 60vw;
}
</style>
<script lang="ts">
export default {
  inheritAttrs: false,
};
</script>
<script lang="ts" setup>
import { useAttrs, defineProps, ref } from "vue";
import { QInputProps as InterfaceRenamed, QPopupProxy } from "quasar";
import { Html5QrcodeScanner } from "html5-qrcode";

interface Props extends InterfaceRenamed {}
const props = defineProps<Props>();
const attrs = useAttrs() as unknown as InterfaceRenamed;
const popupProxy = ref();
const model = ref<string>();
const scanner = ref<Html5QrcodeScanner>();

const onShow = () => {
  if (!attrs.name) {
    return;
  }
  scanner.value = new Html5QrcodeScanner(
    attrs.name ?? "random",
    {
      fps: 25,
      qrbox: (viewfinderWidth: number, viewfinderHeight: number) => ({
        width: viewfinderWidth * 0.7,
        height: viewfinderHeight * 0.7,
      }),
      aspectRatio: 2,
      showTorchButtonIfSupported: true,
      showZoomSliderIfSupported: true,
      useBarCodeDetectorIfSupported: true,
    },
    false
  );

  scanner.value.render((decodedText, decodedResult) => {
    popupProxy.value?.hide();
    model.value = decodedText;
  }, undefined);
};

const onHide = () => {
  scanner.value?.clear();
};
</script>
