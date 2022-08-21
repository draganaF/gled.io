<template>

  <div class="col-8">
    <Card title="Registration">
      <RegistrationForm :role=0 />
    </Card>
  </div>

</template>

<script>

import RegistrationForm from '../shared/RegistrationForm.vue'
import Card from '../../../components/Card/Card.vue';

export default {
  components: {
    RegistrationForm,
    Card,
  },

  mounted() {
    console.log("Registraion page has mounted")
  }
}
</script>

<style>
</style>
