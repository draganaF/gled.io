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
          :disabled="!isEdit"
        />

        <TextInput
          label="LastName"
          v-model="user.LastName"
          :isValid="validateText(user.LastName)"
          :showErrorMessage="showErrorMessage"
          errorMessage="Valid last name must be provided."
          :disabled="!isEdit"
        />

      <Button 
        v-if="isEdit"
        @click="showErrorMessage = true" type="submit"
      >
        Submit
      </Button>
      
      </FormGroup>

    </Form>

  </div>
</template>

<script>
import TextInput from '../../../components/Form/TextInput.vue';
import Form from '../../../components/Form/Form.vue';
import FormGroup from '../../../components/Form/FormGroup.vue';
import { validateText } from '../../../utils/validation';
import Button from '../../../components/Form/Button.vue';
import { mapActions, mapGetters } from 'vuex';
import toastr from 'toastr'

export default {
  components: { TextInput, Form, FormGroup, Button },

  props: ["existingUser", "isEdit"],

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
      if(label !== 'update') return;

      if(ok) {
        toastr.success(message);
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {

    ...mapActions({
      updateUser: 'users/updateUser'
    }),

    validateText,

    onSubmit() {
      this.updateUser(this.user);
    }
  },
}
</script>

<style>
</style>