<template>
  <div>

    <Form @submit="onSubmit">

      <FormGroup title="Credentials">

        <TextInput
          label="Email"
          v-model="user.Email"
          :disabled=true
        />

        <TextInput 
          type="password"
          label="New Password"
          v-model="user.NewPassword"
          :isValid="validatePassword(user.NewPassword)"
          :showErrorMessage="showErrorMessage"
          errorMessage="Password must have at least 8 characters, special character and a number."
        />

        <TextInput 
          type="password"
          label="Confirm Password"
          v-model="user.RepeatedPassword"
          :isValid="user.NewPassword === user.RepeatedPassword"
          :showErrorMessage="showErrorMessage"
          errorMessage="Passwords don't match."
        />

        <TextInput 
          type="password"
          label="Old Password"
          v-model="user.OldPassword"
        />

      <Button @click="showErrorMessage = true" type="submit">Submit</Button>

      </FormGroup>

    </Form>

  </div>
</template>

<script>
import TextInput from '../../../components/Form/TextInput.vue';
import Form from '../../../components/Form/Form.vue';
import FormGroup from '../../../components/Form/FormGroup.vue';
import { validatePassword } from '../../../utils/validation';
import Button from '../../../components/Form/Button.vue';
import { mapActions, mapGetters } from 'vuex';
import toastr from 'toastr'


export default {
  components: { TextInput, Form, FormGroup, Button },

  props: ["existingUser"],

  data: () =>({
    user: { },
    showErrorMessage: false,
  }),

  computed: {
    ...mapGetters({
      result: 'users/getResult'
    })
  },

  watch: {
    existingUser(user) {
      this.user = { ...user }
    },
    
    result({ ok, message, label }) {
      if(label !== 'updatePassword') return;

      if(ok) {
        toastr.success(message);
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {

    ...mapActions({
      updateUserPassword: 'users/updateUserPassword'
    }),

    validatePassword,

    onSubmit() {
      this.showErrorMessage = false;
      this.updateUserPassword(this.user);
    }
  },
}
</script>

<style>
</style>