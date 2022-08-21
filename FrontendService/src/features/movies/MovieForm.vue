<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-6 ml-auto mr-auto">
          <Card title="Create new Movie">
            <Form>
              <FormRow>
                <div class="col-6">
                  <TextInput
                  label="Movie Name"
                  v-model="Name"/>
                </div>
                <div class="col-6">
                  <TextInput
                  label="Director"
                  v-model="Director"/>
                </div>
              </FormRow>
              <FormRow>
                <div class="col-12">
                  <TextInput
                  label="Plot"
                  v-model="Plot"/>
                </div>  
              </FormRow>
              <FormRow>
                <div class="col-12">
                  <TextInput
                  label="Actors"
                  v-model="Actors"/>
                </div>
              </FormRow>
              <FormRow>
                <div class="col-12">
                  <TextInput
                  label="Picture"
                  v-model="Picture"/>
                </div>
              </FormRow>
              <FromRow>
                <div class="col-12">
                  <SelectOptionInput
                  label="Genre"
                  :options="genres"
                  v-model="Genre"/>
                </div>
              </FromRow>
              
              <FormRow>
                <div class="col-6">
                  <TextInput
                  label="Duration"
                  v-model="Duration"/>
                </div>
                <div class="col-6">
                  <TextInput
                  label="Year"
                  v-model="Year"/>
                </div>
              </FormRow>
              <FormRow>
                <div class="col-6">
                  <TextInput
                  label="Country"
                  v-model="Country"/>
                </div>
                <div class="col-6">
                  <TextInput
                  label="Language"
                  v-model="Language"/>
                </div>
              </FormRow>
              <Button @click="handleCreate">Create</Button>
            </Form>
          </Card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Form from "../../components/Form/Form.vue";
import FormRow from "../../components/Form/FormRow.vue";
import Button from "../../components/Form/Button.vue";
import toastr from "toastr";
import { mapActions, mapGetters } from "vuex";
import Card from "../../components/Card/Card.vue";
import TextInput from "../../components/Form/TextInput.vue";
import SelectOptionInput from "../../components/Form/SelectOptionInput.vue";

const genres = [
    { value: 0, label: 'Action' },
    { value: 1, label: 'SciFi' },
    { value: 2, label: 'Adventure' },
    { value: 3, label: 'Comedy' },
    { value: 4, label: 'Romantic' },
    { value: 5, label: 'Horror' },
    { value: 6, label: 'Triller' },
    { value: 7, label: 'Drama' },
    { value: 8, label: 'Fantasy' },
    { value: 9, label: 'Animated' },
    { value: 10, label: 'Crime' },
    { value: 11, label: 'Western' },
    { value: 12, label: 'History' },
    { value: 13, label: 'Biography' },
    { value: 14, label: 'Documentary'}
]
export default {
  components: {
    Form,
    FormRow,
    Button,
    Card,
    TextInput,
    SelectOptionInput
},

  data: function () {
    return {
      Name: "",
      Plot: "",
      Genre:"",
      Director: "",
      Actors: "",
      Language: "",
      Picture: "",
      Duration: 0,
      Year: 0,
      Country: "",
      genres: genres
    };
  },

  computed: {
    ...mapGetters({
      result: "movies/getResult",
    }),
  },

  watch: {
    result({ ok, message, label }) {
      if (label !== "create") return;
      if (ok) {
        this.$router.push("/movies");
      } else {
        toastr.error(message);
      }
    },
  },

  methods: {
    ...mapActions({
      create: "movies/createMovie",
    }),

    handleCreate() {
      const movie = {
        Name: this.Name,
        Plot: this.Plot,
        Genre: this.Genre,
        Director: this.Director,
        Actors: this.Actors,
        Language: this.Language,
        Picture: this.Picture,
        Duration: this.Duration,
        Year: this.Year,
        Country: this.Country
      };
      console.log(movie)
      this.create(movie);
    },
  },
  mounted() {
    this.fetchMovies();
    this.fetchHalls();
  },
};
</script>
