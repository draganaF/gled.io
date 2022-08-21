<template>
  <div>

    <Form
      @submit="emitSearchEvent"
    >
      <FormRow>
        <div class="col-4">
          <TextInput
            label="First Name"
            v-model="params.Name"
          />
        </div>

        <div class="col-4">
          <TextInput
            label="Last Name"
            v-model="params.LastName"
          />
        </div>

        <div class="col-4">
          <TextInput
            label="Email"
            v-model="params.Email"
          />
        </div>

      </FormRow>

      <FormRow>
        <div class="col-6">
          <NumberInput
            label="From Bought Tickets"
            v-model="params.FromNumberOfBoughtTickets"
          />
        </div>

        <div class="col-6">
          <NumberInput
            label="To Bought Tickets"
            v-model="params.ToNumberOfBoughtTickets"
          />
        </div>
      </FormRow>

      <FormRow>
        <div class="col-6">
          <NumberInput
            label="From Reserved Tickets"
            v-model="params.FromNumberOfResevedTickets"
          />
        </div>

        <div class="col-6">
          <NumberInput
            label="To Reserved Tickets"
            v-model="params.ToNumberOfResevedTickets"
          />
        </div>
      </FormRow>

      <Button type="submit">Search</Button>

    </Form>

  </div>

</template>

<script>
import Form from '../../../components/Form/Form.vue';
import FormRow from '../../../components/Form/FormRow.vue';
import TextInput from '../../../components/Form/TextInput.vue';
import NumberInput from '../../../components/Form/NumberInput.vue';
import Button from '../../../components/Form/Button.vue';

export default {

  components: { Form, FormRow, TextInput, NumberInput, Button },

  props: ['existingParams'],

  data: () => ({
    params: {}
  }),

  methods: {

    emitSearchEvent() {
      this.$emit('search', this.params);
    }
  },

  mounted() {
    this.params = { ...this.existingParams };
  }

}

</script>

<style scoped>
</style>