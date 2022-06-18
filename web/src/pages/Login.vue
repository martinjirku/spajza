<template>
  <page-container>
    <div class="row full-height full-width q-pa-md justify-center">
      <div class="login-box self-center col-12 col-md-8 col-lg-6 col-xl-5">
        <form @submit.prevent="onSubmit">
          <div class="row q-px-xl q-py-lg">
            <div class="col-12">
              <h4 class="font-weight-bold text-primary">
                Špajza - Prihlásenie
              </h4>
            </div>
            <div class="col-12">
              <q-input
                for="username"
                v-model="values.username"
                type="text"
                label="Meno"
                aria-label="Meno"
                autocomplete="email"
                :disabled="isSubmitting"
                :error="username.meta.dirty && !username.meta.valid"
                :error-message="errors.username"
              ></q-input>
            </div>
            <div class="col-12">
              <q-input
                for="password"
                v-model="values.password"
                type="password"
                label="Heslo"
                aria-label="Meno"
                autocomplete="password"
                :disabled="isSubmitting"
                :error="password.meta.dirty && !password.meta.valid"
                :error-message="errors.password"
              ></q-input>
            </div>
            <div class="col-12 text-negative">
              <span>{{ errorMsg }}</span>
            </div>
            <div class="col-6 justify-center">
              <router-link to="/" class="text-primary">Registrácia</router-link>
            </div>
            <div class="col-6">
              <q-btn
                class="full-width"
                color="primary"
                width="100%"
                type="submit"
                :disabled="isSubmitting"
                :loading="isSubmitting"
              >
                Prihlásiť sa
              </q-btn>
            </div>
          </div>
        </form>
      </div>
    </div>
  </page-container>
</template>
<style lang="scss" scoped>
.login-box {
  background-color: rgb(255, 255, 255, 0.8);
}
</style>

<script lang="ts" setup>
import PageContainer from "@components/common/PageContainer.vue";
import { useAuthenticationStore } from "@auth/authentication";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useField, useForm } from "vee-validate";
import { string, object } from "yup";

type FormState = { username: string; password: string };
const usernameRules = string().required().max(255).email();
const passwordRules = string().required().min(4).max(255);

const validationSchema = object<FormState>({
  username: usernameRules,
  password: passwordRules,
});

const { handleSubmit, values, errors, validateField, isSubmitting, meta } =
  useForm<FormState>({ validationSchema });
const username = useField("username");
const password = useField("password");
const errorMsg = ref("");

const auth = useAuthenticationStore();
const router = useRouter();

const onSubmit = handleSubmit(async (values) => {
  try {
    const resp = await auth.login(values.username, values.password);
    if (resp.ok) {
      router.replace(auth.returnUrl ?? "/");
      return;
    }
  } catch (error) {
    console.log(error);
  } finally {
    errorMsg.value = "*Prihlásenie sa nepodarilo, skúste znova.";
  }
});
</script>
