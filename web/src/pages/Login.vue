<template>
  <v-main>
    <page-container>
      <v-container class="d-flex fluid fill-height align-center justify-center">
        <v-col cols="12" class="v-col-md-9 v-col-lg-6 v-col-xl-5">
          <v-card class="elevation-12 pa-sm-6 pa-md-12">
            <v-card-content>
              <form ref="form" @submit.prevent="onSubmit">
                <v-row justify="center">
                  <v-col cols="12">
                    <v-row>
                      <v-col cols="12">
                        <h1 class="text-h4 font-weight-bold text-primary">
                          Špajza - Prihlásenie
                        </h1>
                      </v-col>
                      <v-col cols="12">
                        <v-text-field
                          name="username"
                          label="Meno"
                          type="text"
                          v-model="username"
                          placeholder="Meno"
                          messages="* Povinné"
                          append-inner-icon="mdi-outline-person"
                          :disabled="isLoading"
                          autocomplete="email"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                        <v-text-field
                          name="password"
                          label="Heslo"
                          v-model="password"
                          type="password"
                          placeholder="Heslo"
                          messages="* Povinné"
                          autocomplete="password"
                          :disabled="isLoading"
                          append-inner-icon="mdi-form-textbox-password"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="6">
                        <router-link to="/" class="text-primary"
                          >Registrácia</router-link
                        >
                      </v-col>
                      <v-col cols="6">
                        <v-btn
                          color="primary"
                          width="100%"
                          type="submit"
                          :disabled="isDisabled"
                        >
                          Prihlásiť sa
                        </v-btn>
                      </v-col>
                    </v-row>
                  </v-col>
                </v-row>
              </form>
            </v-card-content>
          </v-card>
        </v-col>
      </v-container>
    </page-container>
  </v-main>
</template>
<script lang="ts" setup>
import PageContainer from "@/components/common/PageContainer.vue";
import { useAuthenticationStore } from "@/auth/authentication";
import { computed, ref } from "vue";
import { useRouter } from "vue-router";

const isLoading = ref(false);
const username = ref("");
const password = ref("");
const errorMsg = ref("");

const auth = useAuthenticationStore();
const router = useRouter();
const isDisabled = computed(() => {
  return (
    isLoading.value ||
    (username?.value?.length ?? 0) < 3 ||
    (password?.value?.length ?? 0) < 3
  );
});

const onSubmit = async () => {
  if (isDisabled.value) return;
  try {
    isLoading.value = true;
    const resp = await auth.login(username.value, password.value);
    isLoading.value = false;
    if (resp.ok) {
      router.replace(auth.returnUrl);
      return;
    }
  } catch (error) {
    console.log(error);
  } finally {
    errorMsg.value = "Prihlásenie sa nepodarilo";
  }
};
</script>
