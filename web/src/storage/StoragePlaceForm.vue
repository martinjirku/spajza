<template>
  <Form
    :key="storagePlaceId"
    @submit="onSubmit"
    ref="formRef"
    :initial-values="category"
    :validation-schema="schema"
    lang="sk"
    v-slot="{ meta: { valid, dirty }, resetForm }"
  >
    <span class="text-h6">{{
      storagePlaceId !== -1 ? "Zmena miesta uloženia" : "Nové miesto uloženia"
    }}</span>
    <Field name="title" v-slot="{ errorMessage, value, field }">
      <q-input
        label="Názov"
        :model-value="(value as string)"
        v-bind="field"
        type="text"
        :error="!!errorMessage"
        :error-message="errorMessage"
      ></q-input>
    </Field>
    <Field name="code" v-slot="{ errorMessage, value, field }">
      <q-input
        label="Kód"
        :model-value="(value as string)"
        v-bind="field"
        type="text"
        :error="!!errorMessage"
        :error-message="errorMessage"
      ></q-input>
    </Field>

    <div class="row q-col-gutter-sm">
      <div class="col-6">
        <q-btn class="full-width" flat type="button" @click="resetForm()">
          Vyčistiť
        </q-btn>
      </div>
      <div class="col-6">
        <q-btn
          color="primary"
          class="full-width"
          type="submit"
          :disable="!valid || !dirty || isLoading"
          >{{ storagePlaceId !== -1 ? "Uložiť" : "Vytvoriť" }}</q-btn
        >
      </div>
    </div>
  </Form>
</template>
<script lang="ts" setup>
import { FormContext, SubmissionHandler } from "vee-validate";
import { ref, watch, computed } from "vue";
import {
  useStoragePlacesMutation,
  useStoragePlaces,
} from "./StoragePlaceQuery";
import { schema } from "./StoragaPlace";
import { Field, Form } from "vee-validate";
import { StoragePlace } from "@api/storagePlace";

const { storagePlaceId } = defineProps({
  storagePlaceId: {
    type: Number,
    required: true,
  },
});
const emit = defineEmits<{
  (e: "submitted", value: [StoragePlace, boolean]): void;
}>();
const formRef = ref<FormContext<Record<string, any>>>();

const { data: storagePlaces } = useStoragePlaces();
const category = computed(() => {
  return (
    storagePlaces.value?.find((c) => c.storagePlaceId === storagePlaceId) ?? {}
  );
});
const { mutateAsync, isLoading } = useStoragePlacesMutation();

watch([category], ([category]) => {
  setTimeout(() => {
    formRef.value?.resetForm(category);
  }, 0);
});

const onSubmit = ((values) => {
  mutateAsync(values as StoragePlace).then((data) => {
    const updated = values.storagePlaceId !== -1;
    emit("submitted", [data, updated]);
  });
}) as SubmissionHandler;
</script>
