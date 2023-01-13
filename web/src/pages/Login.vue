<template>
  <page-container>
    <div
      id="g_id_onload"
      :data-client_id="google.clientId"
      :data-login_uri="google.loginUri"
      :data-ux_mode="google.uxMode"
      data-context="signin"
      data-itp_support="true"
    ></div>
    <div class="row full-height full-width q-pa-md justify-center">
      <div class="login-box self-center col-12 col-md-8 col-lg-6 col-xl-5">
        <form @submit.prevent="onSubmit">
          <div class="row q-px-xl q-py-lg">
            <div class="col-12">
              <h4 class="font-weight-bold text-primary">
                Špajza - Prihlásenie
              </h4>
            </div>
            <!-- <div class="col-12">
              <div
                class="g_id_signin"
                data-type="standard"
                data-shape="rectangular"
                data-theme="outline"
                data-text="signin_with"
                data-size="medium"
                data-logo_alignment="center"
                data-width="100%"
              ></div>
            </div> -->
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
  background-color: var(--bg-semitransparent);
}

.google-btn iframe {
  width: 100% !important;
}
.google-btn > div {
  height: 100% !important;
}
.google-btn > div > div {
  height: 100% !important;
}
.google-btn [role="button"] {
  height: 100% !important;
}
</style>

<script lang="ts" setup>
import PageContainer from "@components/common/PageContainer.vue";
import { useAuthenticationStore } from "@auth/authentication";
import { ref, onBeforeMount } from "vue";
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
const google = ref({
  clientId:
    "1089933053808-glhibpnso4vbc38beorao10b30p64d84.apps.googleusercontent.com",
  loginUri: `${
    window.location.origin
  }/api/user/auth/google?redirect=${encodeURIComponent(
    auth?.returnUrl ?? "/"
  )}`,
  uxMode: "popup",
});
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

onBeforeMount(() => {
  (window as any).google?.accounts.id.initialize({
    client_id: google.value.clientId,
    login_uri: google.value.loginUri,
    ux_mode: google.value.uxMode,
  });
  (window as any).google?.accounts.id.prompt();
});
</script>
