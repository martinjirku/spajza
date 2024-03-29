createUnitOptions
<template>
  <Form
    :key="categoryId"
    @submit="onSubmit"
    ref="formRef"
    :initial-values="category"
    :validation-schema="schema"
    lang="sk"
    v-slot="{ meta: { valid, dirty }, resetForm }"
  >
    <span class="text-h6">{{
      categoryId !== -1 ? "Zmena kategórie" : "Nová kategória"
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
    <Field
      name="defaultUnit"
      v-slot="{ errorMessage, value, field: { value: _, ...field } }"
    >
      <q-select
        label="Jednotka"
        v-model="(value as string)"
        :options="unitOptions"
        v-bind="field"
        stack-label
        input-debounce="10"
        map-options
        emit-value
        use-chips
        use-input
        @filter="filterUnits"
        :error="!!errorMessage"
        :error-message="errorMessage"
      >
        <template v-slot:selected-item="scope">
          <q-chip
            v-if="scope"
            square
            color="white"
            text-color="primary"
            class="q-mb-none q-mt-sm q-ml-xs q-mr-none"
          >
            <q-avatar
              color="primary"
              text-color="white"
              :icon="scope.opt.icon"
            />
            <span class="q-mx-sm">
              {{ scope.opt.label }}
            </span>
          </q-chip>
          <q-badge v-else>prázdny</q-badge>
        </template>
        <template v-slot:no-option>
          <q-item>
            <q-item-section class="text-grey"> Žiadne jednotky </q-item-section>
          </q-item>
        </template>
      </q-select>
    </Field>

    <Field
      name="path"
      v-slot="{ errorMessage, value, field: { value: _, ...field } }"
    >
      <q-select
        label="Hlavná kategória"
        v-model="(value as string)"
        :options="parentsOptions"
        v-bind="field"
        map-options
        emit-value
        clearable
        :error="!!errorMessage"
        :error-message="errorMessage"
      >
        <template v-slot:no-option>
          <q-item>
            <q-item-section class="text-grey"> Žiadni rodičia </q-item-section>
          </q-item>
        </template>
      </q-select>
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
          >{{ categoryId !== -1 ? "Uložiť" : "Vytvoriť" }}</q-btn
        >
      </div>
    </div>
  </Form>
</template>
<script lang="ts" setup>
import { FormContext, SubmissionHandler, useForm } from "vee-validate";
import { ref, watch, computed } from "vue";
import { createParentOptions, ParentOption, schema } from "./Category";
import { useCategories, useCategoryMutation } from "./CategoryQuery";
import { useUnits } from "./UnitQuery";
import { Field, Form } from "vee-validate";
import { Category } from "@api/category";
import { createUnitOptions } from "@units/units";
const { categoryId } = defineProps({
  categoryId: {
    type: Number,
    required: true,
  },
});
const emit = defineEmits<{
  (e: "submitted", value: [Category, boolean]): void;
}>();
const formRef = ref<FormContext<Record<string, any>>>();

const { data: categories } = useCategories();
const category = computed(() => {
  return categories.value?.find((c) => c.id === categoryId) ?? {};
});
const { mutateAsync, isLoading } = useCategoryMutation();
const { data: units } = useUnits();
const unitOptions = ref(createUnitOptions(units.value));

watch([units, () => categoryId], () => {
  unitOptions.value = createUnitOptions(units.value);
});
watch([category], ([category]) => {
  setTimeout(() => {
    formRef.value?.resetForm(category);
  }, 0);
});

const indexedParents = computed(() =>
  categories.value?.reduce((r, c) => {
    r[c.id] = c;
    return r;
  }, {} as Record<string, Category>)
);
const parentsOptions = computed<ParentOption[]>(() =>
  createParentOptions(categories.value, indexedParents.value)
);

const filterUnits = (val: string, update: Function) => {
  if (val === "") {
    update(() => {
      unitOptions.value = createUnitOptions(units.value);
    });
    return;
  }
  update(() => {
    unitOptions.value = createUnitOptions(
      units.value?.filter((i) => {
        return (
          i.name.toLowerCase().indexOf(val.toLowerCase()) > -1 ||
          i.symbol.toLowerCase().indexOf(val.toLowerCase()) > -1 ||
          i.names.some((n) => n.toLowerCase().indexOf(val.toLowerCase()) > -1)
        );
      }) ?? []
    );
  });
};

const onSubmit = ((values) => {
  mutateAsync(values as Category).then((data) => {
    const updated = values.id !== -1;
    emit("submitted", [data, updated]);
  });
}) as SubmissionHandler;
</script>
