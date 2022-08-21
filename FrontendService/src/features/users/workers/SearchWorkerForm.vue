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
            label="From Sold Tickets"
            v-model="params.FromNumberOfSoldTickets"
          />
        </div>

        <div class="col-6">
          <NumberInput
            label="To Sold Tickets"
            v-model="params.ToNumberOfSoldTickets"
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