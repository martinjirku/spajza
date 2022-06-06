<template>
  <v-main>
    <page-container>
      <v-container class="d-flex fluid fill-height align-center justify-center">
        <v-col cols="12" class="v-col-md-9 v-col-lg-6 v-col-xl-5">
          <v-card class="elevation-12 pa-sm-6 pa-md-12">
            <v-card-content>
              <form ref="form" @submit.prevent="">
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
                          placeholder="Meno"
                          messages="* Povinné"
                          append-inner-icon="mdi-outline-person"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                        <v-text-field
                          name="password"
                          label="Heslo"
                          type="password"
                          placeholder="Heslo"
                          messages="* Povinné"
                          append-inner-icon="mdi-form-textbox-password"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="6">
                        <router-link to="/registracia" class="text-primary"
                          >Registrácia</router-link
                        >
                      </v-col>
                      <v-col cols="6">
                        <v-btn color="primary" width="100%">
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
</script>
