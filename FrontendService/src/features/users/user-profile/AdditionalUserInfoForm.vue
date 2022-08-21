<template>
  <div>

    <Form>
    
      <FormGroup title="Additional Informations">

        <FormRow>
          <div class="col-6">
            <NumberInput
              label="Balance"
              v-model="user.Total"
              :disabled=true
            />
          </div>

          <div class="col-6">
            <NumberInput
              label="Negative Points"
              v-model="user.NegativePoints"
              :disabled=true
            />
          </div>
        </FormRow>

        <FormRow>
          <div class="col-6">
            <NumberInput
              label="Bought Tickets"
              v-model="user.NumberOfBoughtTickets"
              :disabled=true
            />
          </div>

          <div class="col-6">
            <NumberInput
              label="Reserved Tickets"
              v-model="user.NumberOfReservedTickets"
              :disabled=true
            />
          </div>
        </FormRow>
      
      </FormGroup>

    </Form>

  </div>
</template>

<script>
import Form from '../../../components/Form/Form.vue';
import FormGroup from '../../../components/Form/FormGroup.vue';
import FormRow from '../../../components/Form/FormRow.vue';
import NumberInput from '../../../components/Form/NumberInput.vue';

export default {
  components: { Form, FormGroup, FormRow, NumberInput },

  props: ["user"],
}
</script>

<style>
</style>