<template>
  <div>

    <Form>
    
      <FormGroup title="Additional Informations">

        <NumberInput
          label="Sold Tickets"
          v-model="user.NumberOfSoldTickets"
          :disabled=true
        />
      
      </FormGroup>

    </Form>

  </div>
</template>

<script>
import Form from '../../../components/Form/Form.vue';
import FormGroup from '../../../components/Form/FormGroup.vue';
import NumberInput from '../../../components/Form/NumberInput.vue';

export default {
  components: { Form, FormGroup, NumberInput },

  props: ["user"],
}
</script>

<style>
</style>