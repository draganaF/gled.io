<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-6 ml-auto mr-auto">
          <Card title="Create new Movie">
            <Form @submit="onSubmit($event)">
              <FormRow>
                <div class="col-6">
                  <TextInput 
                  label="Movie Name" 
                  v-model="movie.Name" 
                  :isValid="validateAlphanumericalWord(movie.Name)"
                  :showErrorMessage="showErrorMessage"
                  errorMessage="Valid score must be provided." />
                </div>
                <div class="col-6">
                  <TextInput 
                  label="Director" 
                  v-model="movie.Director"
                  :isValid="validateText(movie.Director)" 
                  :showErrorMessage="showErrorMessage"
                  errorMessage="Valid score must be provided."/>
                </div>
              </FormRow>
              <FormRow>
                <div class="col-12">
                  <TextInput 
                  label="Plot" 
                  v-model="movie.Plot"
                  :isValid="validateAlphanumericalWord(movie.Plot)"
                  :showErrorMessage="showErrorMessage"
                  errorMessage="Valid score must be provided." />
                </div>
              </FormRow>
              <FormRow>
                <div class="col-12">
                  <TextInput 
                  label="Actors" 
                  v-model="movie.Actors"
                  :isValid="validateAlphanumericalWord(movie.Actors)" 
                  :showErrorMessage="showErrorMessage"
                  errorMessage="Valid score must be provided."/>

                </div>
              </FormRow>
              <FormRow>
                <div class="col-12">
                  <TextInput label="Picture" v-model="movie.Picture" />
                </div>
              </FormRow>
              <FormRow>
                <div class="col-12">
                  <SelectOptionInput
                    label="Genre"
                    :options="genres"
                    v-model="movie.Genre"
                  />
                </div>
              </FormRow>

              <FormRow>
                <div class="col-6">
                  <NumberInput 
                  label="Duration" 
                  v-model="movie.Duration" 
                  :isValid="movie.Duration >=0"/>
                </div>
                <div class="col-6">
                  <NumberInput 
                  label="Year" 
                  v-model="movie.Year"
                  :isValid="movie.Year >=0 && movie.Year <=2023" />
                </div>
              </FormRow>
              <FormRow>
                <div class="col-6">
                  <TextInput 
                  label="Country" 
                  v-model="movie.Country"
                  :isValid="validateText(movie.Country)" />
                </div>
                <div class="col-6">
                  <TextInput 
                  label="Language" 
                  v-model="movie.Language"
                  :isValid="validateText(movie.Language)" />
                </div>
              </FormRow>
              <Button @click="showErrorMessage = true" type="submit">{{ isEdit ? 'Update' : 'Create' }}</Button>
            </Form>
          </Card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Form from "../../../components/Form/Form.vue";
import FormRow from "../../../components/Form/FormRow.vue";
import Button from "../../../components/Form/Button.vue";
import toastr from "toastr";
import { mapActions, mapGetters } from "vuex";
import { validateText, validateAlphanumericalWord } from '../../../utils/validation';

import Card from "../../../components/Card/Card.vue";
import TextInput from "../../../components/Form/TextInput.vue";
import SelectOptionInput from "../../../components/Form/SelectOptionInput.vue";
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

const initialMovie = {
  Name: "",
  Plot: "",
  Genre: "",
  Director: "",
  Actors: "",
  Language: "",
  Picture: "",
  Duration: 0,
  Year: 0,
  Country: "",
};
export default {
  components: {
    Form,
    FormRow,
    Button,
    Card,
    TextInput,
    SelectOptionInput,
    NumberInput
  },
  
  data: function () {
    return {
      isEdit: false,
      movie: { ...initialMovie },
      genres: genres,
      showErrorMessage: false,
    };
  },

  computed: {
    ...mapGetters({
      result: "movies/getResult",
      existingMovie: "movies/getMovie"
    }),
  },

  watch: {
    result({ ok, message, label }) {
      
      if (ok) {
        toastr.success(message);
        this.$router.push("/movies");
      } else {
        toastr.error(message);
      }
    },

    isEdit() {
      this.setEdit();
    },

    existingMovie() {
      this.setEdit();
    },
  },

  methods: {
    ...mapActions({
      create: "movies/createMovie",
      update: "movies/updateMovie",
      fetchMovie: "movies/fetchMovieById"
    }),
    validateText,
    validateAlphanumericalWord,
    setEdit() {
      
      if (!this.isEdit) {
        this.movie = { ...initialMovie };
        return;
      }
      if (this.existingMovie) {
        this.movie = this.existingMovie;
      }
    },
    onSubmit(e) {
      e.preventDefault();
      const movieObject = {
        Name: this.movie.Name,
        Plot: this.movie.Plot,
        Genre: this.movie.Genre,
        Director: this.movie.Director,
        Actors: this.movie.Actors,
        Language: this.movie.Language,
        Picture: this.movie.Picture,
        Duration: +this.movie.Duration,
        Year: +this.movie.Year,
        Country: this.movie.Country,
      };

      if (!this.isEdit) {
        this.create(movieObject);
      } else {
        movieObject.Id = this.existingMovie.Id;
        this.update(movieObject);
      }
    }
  },
  mounted() {
    if (this.$route.params.id) {

      this.isEdit = true
      this.fetchMovie(this.$route.params.id)
      
    } 
  }
};
</script>
