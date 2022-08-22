<template>
  <div>
    <Form @submit="emitSearchEvent">
      <FormRow>
        <div class="col-4">
          <TextInput label="Movie Name" v-model="params.Name" />
        </div>

        <div class="col-4">
          <TextInput label="Movie Director" v-model="params.Director" />
        </div>

        <div class="col-4">
          <TextInput label="Actor" v-model="params.Actor" />
        </div>
      </FormRow>
      <FormRow>
        <div class="col-3">
          <NumberInput label="Duration" v-model="params.Duration" />
        </div>

        <div class="col-3">
          <NumberInput label="Year" v-model="params.Year" />
        </div>

        <div class="col-3">
          <TextInput label="Language" v-model="params.Language" />
        </div>

        <div class="col-3">
          <TextInput label="Country" v-model="params.Country" />
        </div>
      </FormRow>
      <FormRow>
        <div class="col-12">
          <SelectOptionInput
            label="Genre"
            :options="genres"
            v-model="params.Genre"
          />
        </div>
      </FormRow>
      <FormRow>
        <div class="col-6">
          <DateTimePicker
            label="Choose Start Range"
            v-model="params.DateFrom"
            id="dp1"
          />
        </div>
        <div class="col-6">
          <DateTimePicker 
            label="Choose End Range" 
            v-model="params.DateTo"
            id="dp2"
          />
        </div>
        </FormRow>
        <FormRow>
        <div class="col-6">
          <NumberInput
            label="Choose Score"
            v-model="params.Score"
          />
        </div>
        <div class="col-6">
          <TextInput
            label="Cinema hall name"
            v-model="params.CinemaHallName"
          />
          
        </div>
      </FormRow>
      <Button type="submit">Search</Button>
    </Form>
  </div>
</template>
<script>
import SelectOptionInput from "../../../components/Form/SelectOptionInput.vue";
import DateTimePicker from "../../../components/Form/DateTimePicker.vue";
import Form from "../../../components/Form/Form.vue";
import FormRow from "../../../components/Form/FormRow.vue";
import TextInput from "../../../components/Form/TextInput.vue";
import Button from "../../../components/Form/Button.vue";
import NumberInput from "../../../components/Form/NumberInput.vue";

const genres = [
  { value: 0, label: "Action" },
  { value: 1, label: "SciFi" },
  { value: 2, label: "Adventure" },
  { value: 3, label: "Comedy" },
  { value: 4, label: "Romantic" },
  { value: 5, label: "Horror" },
  { value: 6, label: "Triller" },
  { value: 7, label: "Drama" },
  { value: 8, label: "Fantasy" },
  { value: 9, label: "Animated" },
  { value: 10, label: "Crime" },
  { value: 11, label: "Western" },
  { value: 12, label: "History" },
  { value: 13, label: "Biography" },
  { value: 14, label: "Documentary" },
];
export default {
  components: { Form, FormRow, TextInput, SelectOptionInput, Button, DateTimePicker, NumberInput },

  props: ["existingParams"],

  data: () => ({
    params: {},
    genres: genres,
  }),

  methods: {
    emitSearchEvent() {
      this.$emit("search", this.params);
    },
  },

  mounted() {
    this.params = { ...this.existingParams };
  },
};
</script>
