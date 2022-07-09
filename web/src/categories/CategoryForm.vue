<template>
  <Form
    :key="categoryId"
    @submit="onSubmit"
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
          :disable="!valid || !dirty"
          >{{ categoryId !== -1 ? "Uložiť" : "Vytvoriť" }}</q-btn
        >
      </div>
    </div>
  </Form>
</template>
<script lang="ts" setup>
import { SubmissionHandler } from "vee-validate";
import { ref, watch, defineProps, computed } from "vue";
import { createUnits, schema } from "./Category";
import { useCategories } from "./CategoryQuery";
import { useUnits } from "./UnitQuery";
import { Field, Form } from "vee-validate";
import { setLocale } from "yup";
const { categoryId } = defineProps({
  categoryId: {
    type: Number,
    required: true,
  },
});

const category = computed(() => {
  return categories.value?.find((c) => c.id === categoryId) ?? {};
});
const { data: categories } = useCategories();

const { data: units } = useUnits();
const unitOptions = ref(createUnits(units.value));

watch([units, () => categoryId], () => {
  unitOptions.value = createUnits(units.value);
});

const filterUnits = (val: string, update: Function) => {
  if (val === "") {
    update(() => {
      unitOptions.value = createUnits(units.value);
    });
    return;
  }
  update(() => {
    unitOptions.value = createUnits(
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
  console.log(values);
}) as SubmissionHandler;
</script>
