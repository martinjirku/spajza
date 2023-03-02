<template>
  <div :id="props.id"></div>
</template>
<script lang="ts" setup>
import { Html5QrcodeScanner } from "html5-qrcode";
import { QrDimensionFunction, QrDimensions } from "html5-qrcode/esm/core";
import { onMounted } from "vue";

type Props = {
  id: string;
};
const props = defineProps<Props>();
onMounted(() => {
  const dimensions: QrDimensionFunction = (
    viewfinderWidth: number,
    viewfinderHeight: number
  ) => ({
    width: viewfinderWidth * 0.8,
    height: viewfinderHeight * 0.8,
  });
  const scanner = new Html5QrcodeScanner(
    props.id,
    {
      fps: 10,
      qrbox: dimensions,
      aspectRatio: 3,
      experimentalFeatures: { useBarCodeDetectorIfSupported: true },
      showTorchButtonIfSupported: true,
    },
    false
  );

  scanner.render((decodedText, decodedResult) => {
    console.log(decodedText, decodedResult);
    scanner.clear();
  }, undefined);
});
</script>
