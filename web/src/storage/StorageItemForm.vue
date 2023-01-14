<template>
  <Form
    @submit="onSubmit"
    ref="formRef"
    :validation-schema="schema"
    :initial-values="{ amount: 100, unit: 'gram' }"
    lang="sk"
    v-slot="{ meta: { valid, dirty }, errors, resetForm }"
  >
    <div v-if="isLoading" class="fit row q-pt-xl">
      <q-space></q-space>
      <div class="items-center">
        <q-spinner-ball class="self-center" size="80px" color="grey-1" />
      </div>
      <q-space></q-space>
    </div>
    <div v-if="!isLoading" class="row q-col-gutter-md">
      <div class="col-6">
        <Field
          name="ean"
          v-slot="{ errorMessage, value, field: { value: _, ...field } }"
        >
          <barcode-input
            label="Ean kód"
            v-model="(value as string)"
            v-bind="field"
            :error="!!errorMessage"
            :error-message="errorMessage"
          ></barcode-input>
        </Field>
        <Field
          name="categoryId"
          v-slot="{ errorMessage, value, field: { value: _, ...field } }"
          @update:model-value="onCategoryChange"
        >
          <q-select
            label="Kategória"
            v-model="(value as number)"
            :options="categoryOptions"
            v-bind="field"
            map-options
            emit-value
            :error="!!errorMessage"
            :error-message="errorMessage"
          >
            <template v-slot:no-option>
              <q-item>
                <q-item-section class="text-grey">
                  Žiadni rodičia
                </q-item-section>
              </q-item>
            </template>
          </q-select>
        </Field>
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
        <Field
          name="expirationDate"
          v-slot="{ errorMessage, value, field: { value: _, ...field } }"
        >
          <q-input
            label="Dátum spotreby"
            v-model="(value as string)"
            v-bind="field"
            readonly
            :error="!!errorMessage"
            :error-message="errorMessage"
          >
            <template v-slot:append>
              <q-icon name="event" class="cursor-pointer"> </q-icon>
            </template>
            <q-popup-proxy transition-show="scale" transition-hide="scale">
              <q-date
                v-model="(value as string)"
                v-bind="field"
                years-in-month-view
              >
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="Close" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-input>
        </Field>
      </div>
      <div class="col-6">
        <Field
          name="storagePlaceId"
          v-slot="{ errorMessage, value, field: { value: _, ...field } }"
        >
          <q-select
            label="Miesto uloženia"
            v-model="(value as number)"
            :options="storagePlaceOptions"
            v-bind="field"
            map-options
            emit-value
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
          </q-select>
        </Field>
        <div class="row">
          <div class="col-5">
            <Field
              name="amount"
              v-slot="{ errorMessage, value, field: { value: _, ...field } }"
            >
              <q-input
                label="Množstvo"
                v-model.number="(value as number)"
                v-bind="field"
                input-class="text-right"
                type="number"
                map-options
                emit-value
                :error="!!errorMessage"
                :error-message="errorMessage"
              >
              </q-input>
            </Field>
          </div>
          <div class="col-7">
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
      </div>
    </div>

    <div class="row q-col-gutter-sm q-mt-md">
      <q-space></q-space>
      <div class="col-3">
        <q-btn class="full-width" flat type="button" @click="resetForm()">
          Vyčistiť
        </q-btn>
      </div>
      <div class="col-3">
        <q-btn
          color="primary"
          class="full-width"
          type="submit"
          :disable="isLoading || !valid || !dirty"
        >
          Vytvoriť
        </q-btn>
      </div>
    </div>
  </Form>
</template>
<script lang="ts" setup>
import { FormContext, SubmissionHandler } from "vee-validate";
import { computed, ref, watch } from "vue";
import { useNewStorageItemMutation } from "./StorageQuery";
import { schema } from "./StorageItem";
import { Field, Form } from "vee-validate";
import { NewStorageItem, StorageItem } from "@api/storage";
import { useCategories } from "@categories/CategoryQuery";
import { createTreeLikeCategoryOptions } from "@categories/Category";
import { useStoragePlaces } from "./StoragePlaceQuery";
import { createStoragePlaceOptions } from "./StoragaPlace";
import { useUnits } from "@categories/UnitQuery";
import { createUnitOptions } from "@units/units";
import BarcodeInput from "@components/common/BarcodeInput.vue";

type ScreenSection = "barcode" | "form";

const emit = defineEmits<{
  (e: "submitted", value: StorageItem): void;
}>();

type Props = {
  barcodePreload: boolean;
};
const props = withDefaults(defineProps<Props>(), {
  barcodePreload: false,
});

const formRef = ref<FormContext<NewStorageItem>>();

const { data: dataCategories, isLoading: isCategoryLoading } = useCategories();
const { data: storagePlaces, isLoading: isStoragePlacesLoading } =
  useStoragePlaces();
const { data: units, isLoading: isUnitsLoading } = useUnits();
const { mutateAsync, isLoading: isSubmittingNewStorageItem } =
  useNewStorageItemMutation();
const selectedCategoryId = ref<number>();

const onCategoryChange = (categoryId: number) => {
  const category = dataCategories.value?.find((c) => c.id === categoryId);
  selectedCategoryId.value = categoryId;
  if (!category) return;
  formRef.value?.setFieldValue("title", category.title);
  formRef.value?.setFieldValue("unit", category.defaultUnit);
};

const categoryOptions = computed(() =>
  createTreeLikeCategoryOptions(dataCategories.value)
);
const storagePlaceOptions = computed(() =>
  createStoragePlaceOptions(storagePlaces.value)
);

const unitOptions = computed(() => {
  const category = dataCategories.value?.find(
    (c) => c.id === selectedCategoryId.value
  );
  const defaultUnit = units.value?.find(
    (u) => u.name === category?.defaultUnit
  );
  return createUnitOptions(
    units.value?.filter(
      (u) =>
        u.quantity === defaultUnit?.quantity ||
        defaultUnit?.quantity === undefined
    )
  );
}, {});

const isLoading = computed(() => {
  return (
    isCategoryLoading.value ||
    isStoragePlacesLoading.value ||
    isUnitsLoading.value
  );
});

watch(isLoading, (prev, next) => {
  if (isLoading.value === false) {
    formRef.value?.resetForm();
  }
});

const onSubmit = ((value) => {
  mutateAsync(value as NewStorageItem).then(({ data }) => {
    emit("submitted", data);
  });
}) as SubmissionHandler;
</script>
