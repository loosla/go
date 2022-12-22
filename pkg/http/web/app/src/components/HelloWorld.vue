<template>
  <v-container fill-height fluid>
    <v-row align="center" justify="center">
      <v-col cols="6">
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-text-field v-model="url" :rules="urlRules" label="Url" required></v-text-field>

          <v-text-field v-model="email" :rules="emailRules" label="E-mail" required></v-text-field>

          <v-text-field v-model="password" label="Password" required></v-text-field>

          <v-btn :disabled="!valid" color="success" class="mr-4" @click="save">
            Save
          </v-btn>

          <v-btn color="error" class="mr-4" @click="cancel">
            Cancel
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    valid: true,
    url: '',
    urlRules: [
      v => !!v || 'Url is required',
      v => (v && v.length >= 3) || 'Name must be more than 3 characters',
    ],
    email: '',
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    password: '',
  }),

  methods: {
    save() {
      this.$refs.form.validate()
    },
    cancel() {
      this.$refs.form.reset()
    },
  },
}
</script>