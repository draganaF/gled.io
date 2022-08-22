<template>
  <div>

    <Form @submit="onSubmit">

      <NumberInput
        :min="1"
        :max="5"
        label="Score"
        v-model="recension.score"
        :isValid="recension.score >= 1 && recension.score <=5"
        :showErrorMessage="showErrorMessage"
        errorMessage="Valid score must be provided."
      />

      <TextArea
        label="Message"
        v-model="recension.message"
        :isValid="validateText(recension.message)"
        :showErrorMessage="showErrorMessage"
        errorMessage="Valid message must be provided."
      />

      <Button @click="showErrorMessage = true" type="submit">Submit</Button>
      
    </Form>

  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import toastr from 'toastr'
import { validateText } from '../../../utils/validation';
import { getUserIdFromToken } from '../../../utils/token';
import Form from '../../../components/Form/Form.vue';
import TextArea from '../../../components/Form/TextArea.vue';
import NumberInput from '../../../components/Form/NumberInput.vue';
import Button from '../../../components/Form/Button.vue';

const initialRecension = {
  message: '',
  score: 1
}

export default {
  
  components: { Form, TextArea, NumberInput, Button },

  props: ['movieId'],

  data: () => ({
    recension: { ...initialRecension },
    showErrorMessage: false,
  }),

  computed: {
    ...mapGetters({
      result: 'recensions/getResult'
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
      createRecension: 'recensions/createRecension',
    }),

    validateText,

    onSubmit() {
      this.createRecension({
        ...this.recension,
        movie_id: this.movieId,
        user_id: +getUserIdFromToken()
      });
    }
  }


}
</script>

<style scoped>
</style>