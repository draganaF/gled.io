<template>
  <div>

    <Form @submit="onSubmit">
    
      <FormGroup title="Personal Informations">

        <TextInput
          label="Name"
          v-model="user.Name"
          :isValid="validateText(user.Name)"
          :showErrorMessage="showErrorMessage"
          errorMessage="Valid name must be provided."
        />

        <TextInput
          label="LastName"
          v-model="user.LastName"
          :isValid="validateText(user.LastName)"
          :showErrorMessage="showErrorMessage"
          errorMessage="Valid last name must be provided."
        />

      </FormGroup>

      <FormGroup title="Credentials">

        <TextInput
          label="Email"
          v-model="user.Email"
          :isValid="validateEmail(user.Email)"
          :showErrorMessage="showErrorMessage"
          errorMessage="Valid email name must be provided."
        />

        <TextInput 
          type="password"
          label="Password"
          v-model="user.Password"
          :isValid="validatePassword(user.Password)"
          :showErrorMessage="showErrorMessage"
          errorMessage="Password must have at least 8 characters, special character and a number."
        />

        <TextInput 
          type="password"
          label="Confirm Password"
          v-model="user.RepeatedPassword"
          :isValid="user.Password === user.RepeatedPassword"
          :showErrorMessage="showErrorMessage"
          errorMessage="Passwords don't match."
        />


      </FormGroup>

      <Button @click="showErrorMessage = true" type="submit">Submit</Button>
    </Form>

  </div>
</template>

<script>
import TextInput from '../../../components/Form/TextInput.vue';
import Form from '../../../components/Form/Form.vue';
import FormGroup from '../../../components/Form/FormGroup.vue';
import { validateText, validateEmail, validatePassword } from '../../../utils/validation';
import Button from '../../../components/Form/Button.vue';
import { mapActions, mapGetters } from 'vuex';
import toastr from 'toastr'

const initialUser = {
  Name: '',
  LastName: '',
  Email: '',
  Password: '',
  RepeatedPassword: ''
};

export default {
  components: { TextInput, Form, FormGroup, Button },

  props: ["role"],

  data: () =>({
    user: { ...initialUser },
    showErrorMessage: false,
  }),

  computed: {
    ...mapGetters({
      result: 'users/getResult'
    })
  },

  watch: {
    result({ ok, message, label }) {
      if(label !== 'create') return;

      if(ok) {
        toastr.success(message);
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {

    ...mapActions({
      createUser: 'users/createUser'
    }),

    validateText,
    validateEmail,
    validatePassword,

    onSubmit() {
      this.createUser({ ...this.user, Role: this.role });
    }
  },
}
</script>

<style>
</style>