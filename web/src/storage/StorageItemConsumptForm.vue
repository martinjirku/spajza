<template>
  <Form
    @submit="onSubmit"
    ref="formRef"
    class="q-ma-md"
    :validation-schema="schema"
    :initial-values="{ amount: props.defaultValue, unit: props.defaultUnit }"
    lang="sk"
    v-slot="{ meta: { valid, dirty } }"
  >
    <div v-if="!isLoading" class="row">
      <div class="col-12">
        <Field
          name="amount"
          v-slot="{
            errorMessage,
            value,
            field: { value: _, 'onUpdate:modelValue': modelValue, ...field },
          }"
        >
          <q-input
            label="Množstvo"
            v-model.number="(value as number)"
            @update:model-value="
              (v) =>
                modelValue?.(typeof v === 'number' ? v : parseFloat(v ?? '0'))
            "
            type="number"
            input-class="text-right"
            :error="!!errorMessage"
            :error-message="errorMessage"
          >
          </q-input>
        </Field>
      </div>
      <div class="col-12">
        <Field
          name="unit"
          v-slot="{ errorMessage, value, field: { value: _, ...field } }"
        >
          <q-select
            label="Jednotky"
            v-model="(value as number)"
            :options="unitOptions"
            stack-label
            map-options
            :use-input="false"
            emit-value
            use-chips
            v-bind="field"
            :error="!!errorMessage"
            :error-message="errorMessage"
          >
            <template v-slot:no-option>
              <q-item>
                <q-item-section class="text-grey">
                  Neexistujú žiadne miesta
                </q-item-section>
              </q-item>
            </template>
            <template v-slot:selected-item="scope">
              <q-chip
                v-if="scope"
                square
                dense
                color="white"
                text-color="primary"
              >
                <q-avatar
                  color="primary"
                  text-color="white"
                  :icon="scope.opt.icon"
                />
                <span class="q-mx-xs">
                  {{ scope.opt.label }}
                </span>
              </q-chip>
              <q-badge v-else>prázdny</q-badge>
            </template>
          </q-select>
        </Field>
      </div>
    </div>

    <div class="row q-mt-md">
      <div class="col-6">
        <q-btn class="full-width" flat type="button" @click="onClose">
          Zrušiť
        </q-btn>
      </div>
      <div class="col-6">
        <q-btn
          color="primary"
          class="full-width"
          type="submit"
          :disable="isLoading || !valid"
        >
          Upotrebiť
        </q-btn>
      </div>
    </div>
  </Form>
</template>
<script lang="ts" setup>
import { FormContext, SubmissionHandler } from "vee-validate";
import { computed, ref } from "vue";
import { useConsumptMutation } from "./StorageQuery";
import { Field, Form } from "vee-validate";
import { Consumption } from "@api/storage";
import { useUnits } from "@categories/UnitQuery";
import { createUnitOptions } from "@units/units";
import { number, object, string } from "yup";

const props = defineProps({
  defaultUnit: {
    type: String,
    required: true,
  },
  defaultValue: {
    type: Number,
    required: true,
  },
  id: {
    type: Number,
    required: true,
  },
  close: {
    type: Function,
  },
});

const schema = object({
  amount: number()
    .transform((v) => (Number.isNaN(v) ? undefined : v))
    .required(),
  unit: string().required(),
});

const emit = defineEmits<{ (e: "submitted", value: Consumption): void }>();

const formRef = ref<FormContext<Consumption>>();

const { data: units, isLoading: isUnitsLoading } = useUnits();
const { mutateAsync } = useConsumptMutation();

const unitOptions = computed(() => {
  const unit = units.value?.find((u) => u.name === props.defaultUnit);
  return createUnitOptions(
    units.value?.filter((u) => u.quantity === unit?.quantity)
  );
}, {});

const isLoading = computed(() => {
  return isUnitsLoading.value;
});

const onSubmit = ((value) => {
  mutateAsync({ storageItemId: props.id, ...(value as Consumption) }).then(
    (data) => {
      emit("submitted", data);
      props.close?.();
    }
  );
}) as SubmissionHandler;

const onClose = () => {
  props.close?.();
  formRef.value?.resetForm();
};
</script>
